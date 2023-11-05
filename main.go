package main

import (
	"context"
	"os"

	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

func main() {

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx := context.Background()

	indexer := indexer.NewIndexer(logger, ctx)

}
