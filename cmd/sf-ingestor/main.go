package main

import (
	"fmt"
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
	//fmt.Println("Configuration loaded successfully:", cfg)
	//fmt.Printf("Alpaca Base URL: %s\n", cfg.Alpaca.BaseURL)
	//fmt.Printf("Alpaca API Key (first 5 characters): %s\n", cfg.Alpaca.APIKey[:5])

	appLogger, err := logger.NewLogger(cfg.Log.LevelConsole, cfg.Log.LevelFile)
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer appLogger.Close()

	appLogger.Info().Msg("Logger initialized successfully")
	appLogger.Info().Msg("Starting sf-ingestor service...")
	appLogger.Info().Msgf("Alpaca Base URL: %s", cfg.Alpaca.BaseURL)
	appLogger.Info().Msgf("Alpaca API Key (first 5 characters): %s", cfg.Alpaca.APIKey[:5])
	appLogger.Info().Msg("Service started successfully")

}
