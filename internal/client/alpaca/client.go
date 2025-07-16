package alpaca

import "github.com/maxviazov/signal-flow/internal/config"

type Client struct {
	apiKey    string
	apiSecret string
	baseURL   string
}

func New(cfg *config.AlpacaConfig) *Client {
	return &Client{
		apiKey:    cfg.APIKey,
		apiSecret: cfg.APISecret,
		baseURL:   cfg.BaseURL,
	}
}
