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

	// Listen for user interrupt
	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt, syscall.SIGTERM)

	// Create logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Create context
	ctx, cancel := context.WithCancel(context.Background())

	// Cancel the context on user interrupt
	go func() {
		<-userInterrupt
		logger.Info().Msg("user interrupt")
		cancel()
	}()

	cfg, err := config.ParseConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("error parsing config")
	}

	// Initialize indexer
	i := indexer.NewIndexer(logger, ctx)

	// Start websocket server and broadcast prices
	serverCfg := server.ServerConfig{
		ListenAddr:        "0.0.0.0:5005",
		WriteTimeout:      "20s",
		ReadTimeout:       "20s",
		BroadcastInterval: "5s",
	}
	s, err := server.NewServer(logger, serverCfg)
	if err != nil {
		logger.Error().Err(err).Msg("error creating server")
		cancel()
	}

	// g, _ := errgroup.WithContext(ctx)

	// websocket server process
	// g.Go(func() error {
	// 	defer cancel()
	// 	err := s.StartWebsocketAPI(ctx, logger, i)
	// 	logger.Info().Msg("websocket server exited")
	// 	return err
	// })

	// client process watching for swap events
	//g.Go(func() error {
	//defer cancel()
	c, err := client.NewClient(cfg.NodeUrl, i, logger)
	if err != nil {
		logger.Error().Err(err).Msg("error creating ethereum client")
		cancel()
	}
	c.WatchAndRestartPools(ctx, cfg.Pools)
	//logger.Info().Msg("ethereum client exited")
	//return err
	//})

	// err = g.Wait()
	// if err != nil {
	// 	logger.Error().Err(err).Send()
	// }

	err = s.StartWebsocketAPI(ctx, logger, i)
	if err != nil {
		logger.Error().Err(err).Msg("websocket api error")
	}
}
