// Package config provides configuration structures for the signal-flow application.
package config

// Config represents the main configuration structure for the application.
// It contains all the necessary configuration sections.
type Config struct {
	// Alpaca contains the configuration for Alpaca API integration
	Alpaca AlpacaConfig `mapstructure:"alpaca"`
}

// AlpacaConfig holds the configuration settings for Alpaca API connection.
// It includes authentication credentials and endpoint information.
type AlpacaConfig struct {
	// APIKey is the API key for authenticating with Alpaca API
	APIKey string `mapstructure:"api_key"`
	// APISecret is the API secret for authenticating with Alpaca API
	APISecret string `mapstructure:"api_secret"`
	// BaseURL is the base URL for Alpaca API endpoints
	BaseURL string `mapstructure:"base_url"`
}
