package alpaca

import (
	"github.com/gorilla/websocket"
	"github.com/maxviazov/signal-flow/internal/config"
)

type Client struct {
	apiKey    string
	apiSecret string
	baseURL   string

	conn *websocket.Conn // Assuming websocket is used for Alpaca API
}

func New(cfg *config.AlpacaConfig) *Client {
	if cfg == nil {
		panic("config cannot be nil")
	}
	return &Client{
		apiKey:    cfg.APIKey,
		apiSecret: cfg.APISecret,
		baseURL:   cfg.BaseURL,
	}
}
