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
	Alpaca AlpacaConfig `mapstructure:"alpaca"`
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
	BaseURL string `mapstructure:"base_url" validate:"required,url"` // Ensure BaseURL is provided
}

func NewConfig() (*Config, error) {
	v := viper.New()
	validate := validator.New()

	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("yaml")   // type of the config file (yaml, json, etc.)
	v.AddConfigPath(".")      // path to look for the config file in

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv() // read in environment variables that match

	cfg.Alpaca.APIKey = v.GetString("alpaca.api_key")
	cfg.Alpaca.APISecret = v.GetString("alpaca.api_secret")

	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}
