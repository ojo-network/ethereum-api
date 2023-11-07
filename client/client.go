package client

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/ethereum-api/config"
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

func (c *Client) WatchSwapEvent(
	pool config.Pool,
	ctx context.Context,
) error {
	poolFilterer, err := abi.NewPoolFilterer(common.HexToAddress(pool.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *abi.PoolSwap)
	opts := &bind.WatchOpts{Start: nil, Context: ctx}
	subscription, err := poolFilterer.WatchSwap(opts, eventSink, nil, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			c.logger.Info().Msg("unsubscribing from swap events")
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
