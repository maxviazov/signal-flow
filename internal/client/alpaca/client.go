package alpaca

import (
	"github.com/gorilla/websocket"
	"github.com/maxviazov/signal-flow/internal/config"
	"github.com/rs/zerolog"
)

type Client struct {
	cfg    config.AlpacaConfig
	logger *zerolog.Logger
	conn   *websocket.Conn // Assuming websocket is used for Alpaca API
}

func New(cfg config.AlpacaConfig, logger *zerolog.Logger) (*Client, error) {
	return &Client{
		cfg:    cfg,
		logger: logger,
	}, nil
}

type authRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`    // API Key
	Secret string `json:"secret"` // API Secret
}

func (c *Client) Connect() error {
	c.logger.Info().Str("url", c.cfg.StreamURL).Msg("Connecting to Alpaca WebSocket")
	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.StreamURL, nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to connect to Alpaca WebSocket")
		return err
	}
	c.conn = conn

	authMsg := authRequest{
		Action: c.cfg.Action,
		Key:    c.cfg.APIKey,
		Secret: c.cfg.APISecret,
	}
	// Send authentication message
	c.logger.Info().Msg("Connected to Alpaca WebSocket")
	err = c.conn.WriteJSON(authMsg)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to connect to Alpaca WebSocket")
		c.conn.Close()
		c.logger.Info().Msg("Disconnected from Alpaca WebSocket")
		return err
	}
	c.logger.Info().Msg("Authentication message sent to Alpaca WebSocket")
	return nil
}
