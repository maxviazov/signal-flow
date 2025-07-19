// Package config provides configuration structures for the signal-flow application.
package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"strings"
)

type logConfig struct {
	LevelConsole string `mapstructure:"level_console" validate:"required,oneof=debug info warn error fatal panic"`
	LevelFile    string `mapstructure:"level_file" validate:"required,oneof=debug info warn error fatal panic"`
}

// Config represents the main configuration structure for the application.
// It contains all the necessary configuration sections.
type Config struct {
	// Alpaca contains the configuration for Alpaca API integration
	Alpaca AlpacaConfig `mapstructure:"streamers"`
	// Log contains the logging configuration for the application
	Log logConfig `mapstructure:"log" validate:"required"` // Ensure logConfig is provided
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

func overrideAlpacaSecretsFromEnv(v *viper.Viper, alpaca *AlpacaConfig) {
	if apiKey := v.GetString("streamers.api_key"); apiKey != "" {
		alpaca.APIKey = apiKey // Explicit override of APIKey from environment variables for security and best practice
	}
	if apiSecret := v.GetString("streamers.api_secret"); apiSecret != "" {
		alpaca.APISecret = apiSecret
	}
}

func NewConfig() (*Config, error) {
	// Create a new Viper instance for configuration management
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	// Set default values for configuration
	v.SetDefault("log.level_console", "info")
	v.SetDefault("log.level_file", "info")
	v.SetDefault("streamers.base_url", "https://paper-api.alpaca.markets/v2")
	v.SetDefault("streamers.stream_url", "wss://paper-api.streamers.markets/stream")

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	// Unmarshal the config file into the Config struct
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	// Override Alpaca secrets from environment variables if set (explicit override for security and best practice)
	overrideAlpacaSecretsFromEnv(v, &cfg.Alpaca)

	// Validate the LogConfig section
	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return &cfg, nil
}
