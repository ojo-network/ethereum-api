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

	// Parse the config
	cfg, err := config.ParseConfig("config.yaml")
	if err != nil {
		logger.Fatal().Err(err).Msg("error parsing config")
	}

	// Create new websocket and REST server for each exchange
	for j, exchange := range cfg.Exchanges {
		// Only register base health check on one server to avoid multiple registrations panic
		var registerBaseHealthCheck bool
		if j == 0 {
			registerBaseHealthCheck = true
		} else {
			registerBaseHealthCheck = false
		}

		// Initialize indexer
		i := indexer.NewIndexer(logger, ctx)

		// Start and maintain connection to blockchain nodes
		client.MaintainConnection(exchange, i, ctx, logger)

		s, err := server.NewServer(logger, cfg.Server, cfg.AssetPairs(exchange), registerBaseHealthCheck)
		if err != nil {
			logger.Error().Err(err).Msg("error creating server")
			cancel()
		}

		go func(exchangeName string) {
			err = s.StartWebsocketAPI(ctx, logger, i, exchangeName)
			if err != nil {
				logger.Error().Err(err).Msg("websocket api error")
			}
		}(string(exchange.Name))
	}

	// Cancel the context on user interrupt
	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt, syscall.SIGTERM)
	<-userInterrupt
	logger.Info().Msg("user interrupt")
	cancel()
}
