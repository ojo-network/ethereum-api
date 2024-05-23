package client

import (
	"context"
	"time"

	"github.com/ojo-network/ethereum-api/config"
	"github.com/ojo-network/ethereum-api/pool"
	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

const sleepDurationAfterAllNodesFail = 2 * time.Minute

// MaintainConnection maintains a connection to an ethereum node using the nodeUrls
// in the order provided until a connection is established. It continues this process
// after receiving an error from the client until the context is cancelled.
func MaintainConnection(
	exchange config.Exchange,
	indexer *indexer.Indexer,
	ctx context.Context,
	logger zerolog.Logger,
) {
	go func() {
		var nodeIndex int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if nodeIndex > len(exchange.NodeUrls)-1 {
					time.Sleep(sleepDurationAfterAllNodesFail)
					nodeIndex = 0
				}
				ethClient, err := NewClient(
					exchange.NodeUrls[nodeIndex],
					pool.SupportedExchanges[exchange.Name],
					indexer,
					ctx,
					logger,
				)
				if err != nil {
					logger.Error().Err(err).Msgf("error connecting to ethereum node %s", exchange.NodeUrls[nodeIndex])
				} else {
					ethClient.WatchSwapsAndPollPrices(exchange.Pools)
				}
				nodeIndex++
			}
		}
	}()
}
