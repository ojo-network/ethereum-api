package client

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
	"github.com/rs/zerolog"
)

const maxErrorsPerMinute = 10

type Client struct {
	nodeUrl   string
	ethClient *ethclient.Client
	indexer   *indexer.Indexer

	ctx        context.Context
	cancelFunc context.CancelFunc

	resetErrorsAt int64
	numErrors     int64

	newBlock chan uint64

	logger zerolog.Logger
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
		newBlock:      make(chan uint64),
		logger:        logger,
	}, nil
}

// WatchSwapsAndPollPrices starts the necessary go routines to watches for swap events,
// polls for spot prices, and submit those prices to the indexer. It blocks until the
// client context is cancelled.
func (c *Client) WatchSwapsAndPollPrices(pools []pool.Pool) {
	c.WatchForBlocksAndRestart()
	c.PollSpotPrices(pools)

	for _, pool := range pools {
		c.WatchSwapsAndRestart(pool)
	}
	<-c.ctx.Done()
}

// reportError is called anytime there is an error with the ethereum client. If there are
// more than maxErrorsPerMinute errors in a one minute period, the client context is cancelled.
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
