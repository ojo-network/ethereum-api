package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ojo-network/ethereum-api/abi"
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
					spotPrice := c.QuerySpotPrice(p, blockNum)
					c.logger.Info().Interface("spotPrice", spotPrice).Msg("spot price received")
					c.indexer.AddPrice(spotPrice)
				}
			}
		}
	}()
}

// QuerySpotPrice queries the spot price of a pool
func (c *Client) QuerySpotPrice(p pool.Pool, blockNum uint64) indexer.SpotPrice {
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
		BlockNum:     indexer.BlockNum(blockNum),
		Timestamp:    utils.CurrentUnixTime(),
		ExchangePair: p.ExchangePair(),
		Price:        p.SqrtPriceX96ToDec(slot0.SqrtPriceX96),
	}
}
