package client

import (
	"context"
	"time"

	"github.com/ojo-network/ethereum-api/config"
	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

const sleepDurationAfterAllNodesFail = 2 * time.Minute

// MaintainConnection maintains a connection to an ethereum node, using the nodeUrls
// in the order provided, until a connection is established. Continues this process
// after receiving an error from the client until the context is cancelled.
func MaintainConnection(
	cfg *config.Config,
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
				if nodeIndex > len(cfg.NodeUrls)-1 {
					time.Sleep(sleepDurationAfterAllNodesFail)
					nodeIndex = 0
				}
				ethClient, err := NewClient(cfg.NodeUrls[nodeIndex], indexer, ctx, logger)
				if err != nil {
					logger.Error().Err(err).Msgf("error connecting to ethereum node %s", cfg.NodeUrls[nodeIndex])
				} else {
					ethClient.WatchSwapsAndPollPrices(cfg.Pools)
				}
				nodeIndex++
			}
		}
	}()
}
