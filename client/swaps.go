package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ojo-network/ethereum-api/abi"
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
				}
			}
		}
	}()
}

// WatchUniswapSwapEvent watches for swap events on a uniswap pool
func (c *Client) WatchUniswapSwapEvent(p pool.Pool) error {
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
			c.logger.Info().Interface("uniswap swap", swap).Msg("uniswap swap event received")
			c.indexer.AddSwap(swap)
			c.indexer.AddPrice(spotPrice)
		}
	}
}

// WatchAlgebraSwapEvent watches for swap events on an alegbra pool
func (c *Client) WatchAlgebraSwapEvent(p pool.Pool) error {
	poolFilterer, err := abi.NewAlgebraPoolFilterer(common.HexToAddress(p.Address), c.ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *abi.AlgebraPoolSwap)
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
