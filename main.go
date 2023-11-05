package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ojo-network/indexer/indexer"
	"github.com/rs/zerolog"
)

func main() {

	// Listen for user interrupt
	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt, syscall.SIGTERM)

	// Create logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Create context
	ctx := context.Background()

	indexer := indexer.NewIndexer(logger, ctx)
}
