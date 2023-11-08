package client

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

type Client struct {
	nodeUrl   string
	ethClient *ethclient.Client
	indexer   *indexer.Indexer
	logger    zerolog.Logger
}

func NewClient(
	nodeUrl string,
	indexer *indexer.Indexer,
	logger zerolog.Logger,
) (*Client, error) {
	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, err
	}
	return &Client{
		nodeUrl:   nodeUrl,
		ethClient: ethClient,
		indexer:   indexer,
		logger:    logger,
	}, nil
}

// WatchSwapsAndPollPrices starts the necessary go routines to watches for swap events,
// polls for spot prices, and submit those prices to the indexer
func (c *Client) WatchSwapsAndPollPrices(ctx context.Context, pools []Pool) {
	for _, pool := range pools {
		c.WatchAndRestart(ctx, pool)
		c.PollSpotPrice(ctx, pool)
	}
}

// PollSpotPrice polls for the spot price of a pool
func (c *Client) PollSpotPrice(ctx context.Context, pool Pool) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(5 * time.Second)
				spotPrice := c.QuerySpotPrice(ctx, pool)
				c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
				c.indexer.AddPrice(spotPrice)
			}
		}
	}()
}

// WatchAndRestart watches for swap events and restarts the watcher if it errors
func (c *Client) WatchAndRestart(ctx context.Context, pool Pool) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := c.WatchSwapEvent(ctx, pool)
				if err != nil {
					c.logger.Error().Err(err).Msgf("error watching %s swap events", pool.ExchangePair())
				}
			}
		}
	}()
}

// WatchSwapEvent watches for swap events on a pool
func (c *Client) WatchSwapEvent(ctx context.Context, pool Pool) error {
	poolFilterer, err := abi.NewPoolFilterer(common.HexToAddress(pool.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *abi.PoolSwap)
	opts := &bind.WatchOpts{Start: nil, Context: ctx}
	c.logger.Info().Msgf("subscribing to %s swap events", pool.ExchangePair())
	subscription, err := poolFilterer.WatchSwap(opts, eventSink, nil, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			c.logger.Info().Msgf("unsubscribing from %s swap events", pool.ExchangePair())
			subscription.Unsubscribe()
			return nil
		case err := <-subscription.Err():
			return err
		case event := <-eventSink:
			swap := pool.ConvertEventToSwap(event)
			spotPrice := pool.ConvertEventToSpotPrice(event)
			c.logger.Info().Interface("swap", swap).Msg("swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}
