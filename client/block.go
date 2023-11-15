package client

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) WatchForBlocksAndRestart() {
	go func() {
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				err := c.WatchForNewBlocks()
				if err != nil {
					c.reportError(fmt.Errorf("error watching for new blocks: %w", err))
				}
			}
		}
	}()
}

func (c *Client) WatchForNewBlocks() error {
	headers := make(chan *types.Header)
	subscription, err := c.ethClient.SubscribeNewHead(c.ctx, headers)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-subscription.Err():
			return err
		case header := <-headers:
			fmt.Println(header.Number.Uint64())
			c.newBlock <- header.Number.Uint64()
		}
	}
}
