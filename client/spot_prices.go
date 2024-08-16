package client

import (
	"fmt"
	"math/big"

	sdkmath "cosmossdk.io/math"

	"github.com/ethereum/go-ethereum/common"
	balancerpool "github.com/ojo-network/ethereum-api/abi/balancer/pool"
	"github.com/ojo-network/ethereum-api/abi/camelot"
	"github.com/ojo-network/ethereum-api/abi/curve"
	"github.com/ojo-network/ethereum-api/abi/pancake"
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
						c.logger.Info().Interface("Uniswap spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolAlgebra:
						spotPrice = c.QueryAlgebraSpotPrice(p, blockNum)
						c.logger.Info().Interface("Alegbra spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolBalancer:
						spotPrice = c.QueryBalancerSpotPrice(p, blockNum)
						c.logger.Info().Interface("Balancer spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolPancake:
						spotPrice = c.QueryPancakeSpotPrice(p, blockNum)
						c.logger.Info().Interface("Pancake spotPrice", spotPrice).Msg("spot price received")
						c.indexer.AddPrice(spotPrice)
					case pool.PoolCurve:
						spotPrice = c.QueryCurveSpotPrice(p, blockNum)
						c.logger.Info().Interface("Curve spotPrice", spotPrice).Msg("spot price received")
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

	poolRate, err := poolCaller.GetTokenRate(nil, common.HexToAddress(p.BaseAddress))
	if err != nil {
		c.reportError(fmt.Errorf("error getting %s token rate from pool: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        sdkmath.LegacyNewDecFromBigIntWithPrec(poolRate, 18),
	}
}

// QueryPancakeSpotPrice queries the spot price of a pancake pool
func (c *Client) QueryPancakeSpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
	pancakeCaller, err := pancake.NewPancakeCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		c.reportError(fmt.Errorf("error initializing %s pool caller: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	slot0, err := pancakeCaller.Slot0(nil)
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

// QueryCurveSpotPrice queries the spot price of a curve pool
func (c *Client) QueryCurveSpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
	curveCaller, err := curve.NewCurveCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		c.reportError(fmt.Errorf("error initializing %s pool caller: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}

	// price comes inverted
	poolPriceInverted, err := curveCaller.LastPrice(nil, big.NewInt(0))
	if err != nil {
		c.reportError(fmt.Errorf("error getting %s token last price from pool: %w", p.ExchangePair(), err))
		return indexer.SpotPrice{}
	}
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(36), nil)
	poolPrice := new(big.Int).Quo(scale, poolPriceInverted)

	return indexer.SpotPrice{
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        sdkmath.LegacyNewDecFromBigIntWithPrec(poolPrice, 18),
	}
}
