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
