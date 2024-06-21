package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	balancerpool "github.com/ojo-network/ethereum-api/abi/balancer/pool"
	"github.com/ojo-network/ethereum-api/abi/camelot"
	"github.com/ojo-network/ethereum-api/abi/uniswap"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

// PollSpotPrices polls for the spot price of a pool
func (c *Client) PollSpotPrices(pools []pool.Pool) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			case blockNum := <-c.newBlock:
				for _, p := range pools {
					var spotPrice indexer.SpotPrice
					switch c.poolContract {
					case pool.PoolUniswap:
						spotPrice = c.QueryUniswapSpotPrice(p, blockNum)
						c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolAlgebra:
						spotPrice = c.QueryAlgebraSpotPrice(p, blockNum)
						c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolBalancer:
						spotPrice = c.QueryAlgebraSpotPrice(p, blockNum)
						c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					}
				}
			}
		}
	}()
}

// QueryUniswapSpotPrice queries the spot price of a uniswap pool
func (c *Client) QueryUniswapSpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
	poolCaller, err := uniswap.NewPoolCaller(common.HexToAddress(p.Address), c.ethClient)
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
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        p.SqrtPriceX96ToDec(slot0.SqrtPriceX96),
	}
}

// QueryAlgebraSpotPrice queries the spot price of an alegbra pool
func (c *Client) QueryAlgebraSpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
	poolCaller, err := camelot.NewAlgebraPoolCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		c.reportError(fmt.Errorf("error initializing %s pool caller: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	globalState, err := poolCaller.GlobalState(nil)
	if err != nil {
		c.reportError(fmt.Errorf("error getting %s pool balance: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        p.SqrtPriceX96ToDec(globalState.Price),
	}
}

// QueryBalancerSpotPrice queries the spot price of a balancer pool
func (c *Client) QueryBalancerSpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
	poolCaller, err := balancerpool.NewPoolCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		c.reportError(fmt.Errorf("error initializing %s pool caller: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	poolRate, err := poolCaller.GetRate(nil)
	if err != nil {
		c.reportError(fmt.Errorf("error getting %s pool balance: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        p.SqrtPriceX96ToDec(poolRate),
	}
}
