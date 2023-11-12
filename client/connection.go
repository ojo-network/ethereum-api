package client

import (
	"fmt"

	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

func Connect(
	nodeUrls []string,
	indexer *indexer.Indexer,
	logger zerolog.Logger,
) (*Client, error) {
	for _, nodeUrl := range nodeUrls {
		ethClient, err := NewClient(nodeUrl, indexer, logger)
		if err != nil {
			logger.Error().Err(err).Msgf("error connecting to ethereum node %s", nodeUrl)
			continue
		}
		return ethClient, nil
	}
	return nil, fmt.Errorf("error connecting to all ethereum nodes")
}
