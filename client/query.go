package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

// QuerySpotPrice queries the spot price of a pool
func (c *Client) QuerySpotPrice(p pool.Pool) indexer.SpotPrice {
	poolCaller, err := abi.NewPoolCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		c.reportError(fmt.Errorf("error initializing %s pool caller: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	slot0, err := poolCaller.Slot0(nil)
	if err != nil {
		c.reportError(fmt.Errorf("error getting %s pool balance: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	return indexer.SpotPrice{
		BlockNum:     c.QueryBlockNumber(c.ctx),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        p.SqrtPriceX96ToDec(slot0.SqrtPriceX96),
	}
}

// QueryBlockNumber queries the current ethereum block number
func (c *Client) QueryBlockNumber(ctx context.Context) indexer.BlockNum {
	blockNum, err := c.ethClient.BlockNumber(ctx)
	if err != nil {
		c.reportError(fmt.Errorf("error retrieving ethereum block number: %w", err))
		return 0
	}
	return indexer.BlockNum(blockNum)
}
