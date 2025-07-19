package main

import (
	"fmt"
	"github.com/maxviazov/signal-flow/internal/client/streamers"
	"github.com/maxviazov/signal-flow/internal/config"
	"github.com/maxviazov/signal-flow/pkg/logger"
	"log"
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

	appLogger.Info().Msg("Logger initialized successfully")
	appLogger.Info().Msg("Starting sf-ingestor service...")
	appLogger.Info().Msgf("Alpaca Base URL: %s", cfg.Alpaca.BaseURL)
	appLogger.Info().Msg("Service started successfully")

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
