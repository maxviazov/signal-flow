package main

import (
	"context"
	"fmt"
	"github.com/maxviazov/signal-flow/internal/client/streamers"
	"github.com/maxviazov/signal-flow/internal/config"
	"github.com/maxviazov/signal-flow/internal/repository/postgres"
	"github.com/maxviazov/signal-flow/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting sf-ingestor service...")

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	appLogger, err := logger.NewLogger(cfg.Log.LevelConsole, cfg.Log.LevelFile)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer func() {
		if err := appLogger.Close(); err != nil {
			log.Printf("Error closing logger: %v", err)
		}
	}()

	ctx := context.Background()
	initDB, err := postgres.New(ctx, cfg)
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to initialize Postgres repository")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		appLogger.Info().Msg("Received shutdown signal, shutting down gracefully...")
		cancel()
	}()
	appLogger.Info().Msg("Logger initialized successfully")
	appLogger.Info().Msg("Starting sf-ingestor service...")
	appLogger.Info().Msgf("Alpaca Base URL: %s", cfg.Alpaca.BaseURL)
	appLogger.Info().Msg("Service started successfully")
	appLogger.Info().Msgf("Postgres: %+v", initDB)

	marketStream := streamers.New(cfg.Alpaca, &appLogger.Logger)

	if err := marketStream.Connect(); err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to connect to Alpaca WebSocket")
	}

	symbolsToTrade := []string{"AAPL", "GOOGL", "TSLA"}
	if err := marketStream.Subscribe(symbolsToTrade); err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to subscribe to symbols")
	}

	go func() {
		if err := marketStream.Listen(); err != nil {
			appLogger.Error().Err(err).Msg("Error in Listen")
		}
	}()

	select {} // Keep the main goroutine running
}
