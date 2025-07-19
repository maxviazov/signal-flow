// Package config provides configuration structures for the signal-flow application.
package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

type LogConfig struct {
	LevelConsole string `mapstructure:"level_console" validate:"required,oneof=debug info warn error fatal panic"` // Log level for console output
	LevelFile    string `mapstructure:"level_file" validate:"required,oneof=debug info warn error fatal panic"`    // Log level for file output
}

// Config represents the main configuration structure for the application.
// It contains all the necessary configuration sections.
type Config struct {
	// Alpaca contains the configuration for Alpaca API integration
	Alpaca AlpacaConfig `mapstructure:"streamers"`
	// Log contains the logging configuration for the application
	Log LogConfig `mapstructure:"log" validate:"required"` // Ensure LogConfig is provided
}

// AlpacaConfig holds the configuration settings for Alpaca API connection.
// It includes authentication credentials and endpoint information.
type AlpacaConfig struct {
	// APIKey is the API key for authenticating with Alpaca API
	APIKey string `mapstructure:"api_key" validate:"required"` // Ensure APIKey is provided
	// APISecret is the API secret for authenticating with Alpaca API
	APISecret string `mapstructure:"api_secret" validate:"required"` // Ensure APISecret is provided
	// BaseURL is the base URL for Alpaca API endpoints
	BaseURL   string `mapstructure:"base_url" validate:"required,url"`   // Ensure BaseURL is provided
	StreamURL string `mapstructure:"stream_url" validate:"required,url"` // Ensure StreamURL is provided
}

func NewConfig() (*Config, error) {
	v := viper.New()
	validate := validator.New()

	// Set up Viper to read configuration from a file
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // type of the config file (yaml, json, etc.)
	v.AddConfigPath(".")
	v.SetDefault("log.level_console", "info")                                        // Default log level for console output
	v.SetDefault("log.level_file", "info")                                           // Default log level for file output
	v.SetDefault("streamers.base_url", "https://paper-api.alpaca.markets/v2")        // Default base URL for Alpaca API
	v.SetDefault("streamers.stream_url", "wss://paper-api.streamers.markets/stream") // Default stream URL for Alpaca API

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	// Unmarshal the config file into the Config struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv() // read in environment variables that match

	// Override config values with environment variables
	cfg.Alpaca.BaseURL = v.GetString("streamers.base_url")
	if cfg.Alpaca.BaseURL == "" {
		cfg.Alpaca.BaseURL = "https://paper-api.alpaca.markets/v2" // Default to paper trading URL
	}
	// Read APIKey and APISecret from environment variables if set
	cfg.Alpaca.APIKey = v.GetString("streamers.api_key")
	cfg.Alpaca.APISecret = v.GetString("streamers.api_secret")

	// Validate the LogConfig section
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &cfg, nil
}
