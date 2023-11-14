package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ojo-network/ethereum-api/client"
	"github.com/ojo-network/ethereum-api/config"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/server"
	"github.com/rs/zerolog"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel the context on user interrupt
	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-userInterrupt
		logger.Info().Msg("user interrupt")
		cancel()
	}()

	// Parse the config
	cfg, err := config.ParseConfig("config.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("error parsing config")
	}

	// Initialize indexer
	i := indexer.NewIndexer(logger, ctx)

	// Start and maintain connection to ethereum node
	client.MaintainConnection(cfg, i, ctx, logger)

	// Create new websocket and REST server
	s, err := server.NewServer(logger, cfg.Server, cfg.AssetPairs())
	if err != nil {
		logger.Error().Err(err).Msg("error creating server")
		cancel()
	}

	// Start websocket server and block until error or context is cancelled
	err = s.StartWebsocketAPI(ctx, logger, i)
	if err != nil {
		logger.Error().Err(err).Msg("websocket api error")
	}
}
