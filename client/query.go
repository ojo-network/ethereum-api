package client

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

// QuerySpotPrice queries the spot price of a pool
func (c *Client) QuerySpotPrice(ctx context.Context, pool Pool) indexer.SpotPrice {
	poolCaller, err := abi.NewPoolCaller(common.HexToAddress(pool.Address), c.ethClient)
	if err != nil {
		c.logger.Error().Err(err).Msgf("error initializing %s pool caller", pool.ExchangePair)
		return indexer.SpotPrice{}
	}
	slot0, err := poolCaller.Slot0(nil)
	if err != nil {
		c.logger.Error().Err(err).Msgf("error getting %s pool balance", pool.ExchangePair)
		return indexer.SpotPrice{}
	}
	return indexer.SpotPrice{
		BlockNum:     c.QueryBlockNumber(ctx),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: pool.ExchangePair,
		Price:        pool.SqrtPriceX96ToDec(slot0.SqrtPriceX96),
	}
}

// QueryBlockNumber queries the current ethereum block number
func (c *Client) QueryBlockNumber(ctx context.Context) indexer.BlockNum {
	blockNum, err := c.ethClient.BlockNumber(ctx)
	if err != nil {
		c.logger.Error().Err(err).Msg("error retrieving ethereum block number")
		return 0
	}
	return indexer.BlockNum(blockNum)
}
