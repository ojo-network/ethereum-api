package client

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
	"github.com/rs/zerolog"
)

const maxErrorsPerMinute = 5

type Client struct {
	nodeUrl   string
	ethClient *ethclient.Client
	indexer   *indexer.Indexer

	ctx           context.Context
	cancelFunc    context.CancelFunc
	resetErrorsAt int64
	numErrors     int64
	logger        zerolog.Logger
}

func NewClient(
	nodeUrl string,
	indexer *indexer.Indexer,
	mainCtx context.Context,
	logger zerolog.Logger,
) (*Client, error) {
	logger.Info().Msgf("connecting to ethereum node %s", nodeUrl)

	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}

	ctx, cancelFunc := context.WithCancel(mainCtx)

	return &Client{
		nodeUrl:       nodeUrl,
		ethClient:     ethClient,
		indexer:       indexer,
		ctx:           ctx,
		cancelFunc:    cancelFunc,
		resetErrorsAt: utils.CurrentUnixTime(),
		logger:        logger,
	}, nil
}

// WatchSwapsAndPollPrices starts the necessary go routines to watches for swap events,
// polls for spot prices, and submit those prices to the indexer. It blocks until the
// client context is cancelled.
func (c *Client) WatchSwapsAndPollPrices(pools []pool.Pool) {
	for _, pool := range pools {
		c.WatchAndRestart(pool)
		c.PollSpotPrice(pool)
	}
	<-c.ctx.Done()
}

// PollSpotPrice polls for the spot price of a pool
func (c *Client) PollSpotPrice(p pool.Pool) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				time.Sleep(5 * time.Second)
				spotPrice := c.QuerySpotPrice(p)
				c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
				c.indexer.AddPrice(spotPrice)
			}
		}
	}()
}

// WatchAndRestart watches for swap events and restarts the watcher if it errors
func (c *Client) WatchAndRestart(p pool.Pool) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				err := c.WatchSwapEvent(p)
				if err != nil {
					c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
				}
			}
		}
	}()
}

// WatchSwapEvent watches for swap events on a pool
func (c *Client) WatchSwapEvent(p pool.Pool) error {
	poolFilterer, err := abi.NewPoolFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *abi.PoolSwap)
	opts := &bind.WatchOpts{Start: nil, Context: c.ctx}
	c.logger.Info().Msgf("subscribing to %s swap events", p.ExchangePair())
	subscription, err := poolFilterer.WatchSwap(opts, eventSink, nil, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info().Msgf("unsubscribing from %s swap events", p.ExchangePair())
			subscription.Unsubscribe()
			return nil
		case err := <-subscription.Err():
			return err
		case event := <-eventSink:
			swap := p.ConvertEventToSwap(event)
			spotPrice := p.ConvertEventToSpotPrice(event)
			c.logger.Info().Interface("swap", swap).Msg("swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

func (c *Client) reportError(err error) {
	c.logger.Error().Err(err).Msg("ethereum client error")
	if time.Now().Unix() > c.resetErrorsAt {
		c.resetErrorsAt = time.Now().Add(time.Minute * 1).Unix()
		c.numErrors = 1
	}
	c.numErrors++
	if c.numErrors > maxErrorsPerMinute {
		if err := c.ctx.Err(); err == nil {
			c.logger.Error().Msgf("more than %d errors for a one minute period, cancelling ethereum client", maxErrorsPerMinute)
		}
		c.cancelFunc()
	}
}
