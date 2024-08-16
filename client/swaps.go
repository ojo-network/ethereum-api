package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	balancerpool "github.com/ojo-network/ethereum-api/abi/balancer/pool"
	"github.com/ojo-network/ethereum-api/abi/balancer/vault"
	"github.com/ojo-network/ethereum-api/abi/camelot"
	"github.com/ojo-network/ethereum-api/abi/curve"
	"github.com/ojo-network/ethereum-api/abi/pancake"
	"github.com/ojo-network/ethereum-api/abi/uniswap"
	"github.com/ojo-network/ethereum-api/pool"
)

// WatchAndRestart watches for swap events and restarts the watcher if it errors
func (c *Client) WatchSwapsAndRestart(p pool.Pool) {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				switch c.poolContract {
				case pool.PoolUniswap:
					err := c.WatchUniswapSwapEvent(p)
					if err != nil {
						c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
					}
				case pool.PoolAlgebra:
					err := c.WatchAlgebraSwapEvent(p)
					if err != nil {
						c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
					}
				case pool.PoolBalancer:
					err := c.WatchBalancerSwapEvent(p)
					if err != nil {
						c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
					}
				case pool.PoolPancake:
					err := c.WatchPancakeSwapEvent(p)
					if err != nil {
						c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
					}
				case pool.PoolCurve:
					err := c.WatchCurveSwapEvent(p)
					if err != nil {
						c.reportError(fmt.Errorf("error watching %s swap events", p.ExchangePair()))
					}
				}
			}
		}
	}()
}

// WatchUniswapSwapEvent watches for swap events on a uniswap pool
func (c *Client) WatchUniswapSwapEvent(p pool.Pool) error {
	poolFilterer, err := uniswap.NewPoolFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *uniswap.PoolSwap)
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
			swap := p.ConvertUniswapEventToSwap(event)
			spotPrice := p.ConvertUniswapEventToSpotPrice(event)
			c.logger.Info().Interface("uniswap swap", swap).Msg("uniswap swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

// WatchAlgebraSwapEvent watches for swap events on an alegbra pool
func (c *Client) WatchAlgebraSwapEvent(p pool.Pool) error {
	poolFilterer, err := camelot.NewAlgebraPoolFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *camelot.AlgebraPoolSwap)
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
			swap := p.ConvertAlgebraEventToSwap(event)
			spotPrice := p.ConvertAlgebraEventToSpotPrice(event)
			c.logger.Info().Interface("alegbra swap", swap).Msg("algebra swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

// WatchBalancerSwapEvent watches for swap events on an balancer vault
func (c *Client) WatchBalancerSwapEvent(p pool.Pool) error {
	// retreive vault address that supports this pool and pool id
	poolCaller, err := balancerpool.NewPoolCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	vaultAddress, err := poolCaller.GetVault(nil)
	if err != nil {
		return err
	}
	poolId, err := poolCaller.GetPoolId(nil)
	if err != nil {
		return err
	}

	// build parameters for swap event subscription
	poolIdParam := make([][32]byte, 1)
	poolIdParam[0] = poolId
	tokenInParam := make([]common.Address, 1)
	tokenInParam[0] = common.HexToAddress(p.BaseAddress)
	tokenOutParam := []common.Address{}
	for _, tokenAddress := range p.QuoteAddresses {
		tokenOutParam = append(tokenOutParam, common.HexToAddress(tokenAddress))
	}

	vaultFilterer, err := vault.NewPoolFilterer(vaultAddress, c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *vault.PoolSwap)
	opts := &bind.WatchOpts{Start: nil, Context: c.ctx}
	c.logger.Info().Msgf("subscribing to %s swap events", p.ExchangePair())
	subscription, err := vaultFilterer.WatchSwap(opts, eventSink, poolIdParam, tokenInParam, tokenOutParam)
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
			// query rate from pool contract
			poolRate, err := poolCaller.GetTokenRate(nil, common.HexToAddress(p.BaseAddress))
			if err != nil {
				return err
			}

			swap := p.ConvertBalancerEventToSwap(event, poolRate)
			spotPrice := p.ConvertBalancerEventToSpotPrice(event, poolRate)
			c.logger.Info().Interface("balancer swap", swap).Msg("balancer swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

// WatchPancakeSwapEvent watches for swap events on a pancake pool
func (c *Client) WatchPancakeSwapEvent(p pool.Pool) error {
	pancakeFilterer, err := pancake.NewPancakeFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *pancake.PancakeSwap)
	opts := &bind.WatchOpts{Start: nil, Context: c.ctx}
	c.logger.Info().Msgf("subscribing to %s swap events", p.ExchangePair())
	subscription, err := pancakeFilterer.WatchSwap(opts, eventSink, nil, nil)
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
			swap := p.ConvertPancakeEventToSwap(event)
			spotPrice := p.ConvertPancakeEventToSpotPrice(event)
			c.logger.Info().Interface("pancake swap", swap).Msg("pancake swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

// WatchCurveSwapEvent watches for swap events on a curve pool
func (c *Client) WatchCurveSwapEvent(p pool.Pool) error {
	curveCaller, err := curve.NewCurveCaller(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	curveFilterer, err := curve.NewCurveFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *curve.CurveTokenExchange)
	opts := &bind.WatchOpts{Start: nil, Context: c.ctx}
	c.logger.Info().Msgf("subscribing to %s swap events", p.ExchangePair())
	subscription, err := curveFilterer.WatchTokenExchange(opts, eventSink, nil)
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
			// price comes inverted
			poolPriceInverted, err := curveCaller.LastPrice(nil, big.NewInt(0))
			if err != nil {
				return err
			}
			scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(36), nil)
			poolPrice := new(big.Int).Quo(scale, poolPriceInverted)

			swap := p.ConvertCurveEventToSwap(event, poolPrice)
			spotPrice := p.ConvertCurveEventToSpotPrice(event, poolPrice)
			c.logger.Info().Interface("curve swap", swap).Msg("curve swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}
