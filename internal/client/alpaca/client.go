package alpaca

import (
	"github.com/gorilla/websocket"
	"github.com/maxviazov/signal-flow/internal/config"
	"github.com/rs/zerolog"
)

type Client struct {
	cfg    config.AlpacaConfig
	logger *zerolog.Logger
	conn   *websocket.Conn
}

func New(cfg config.AlpacaConfig, logger *zerolog.Logger) *Client {
	return &Client{
		cfg:    cfg,
		logger: logger,
	}
}

// authRequest defines the structure for the authentication message.
type authRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Connect establishes a WebSocket connection and authenticates with Alpaca.
func (c *Client) Connect() error {
	c.logger.Info().Str("url", c.cfg.StreamURL).Msg("Connecting to Alpaca WebSocket...")

	// Establish the connection
	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.StreamURL, nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to establish WebSocket connection")
		return err
	}
	c.conn = conn
	c.logger.Info().Msg("WebSocket connection established")

	// Prepare and send the authentication message
	authMsg := authRequest{
		Action: "auth", // Use the constant "auth" for the action
		Key:    c.cfg.APIKey,
		Secret: c.cfg.APISecret,
	}

	c.logger.Info().Msg("Sending authentication request...")
	err = c.conn.WriteJSON(authMsg)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to send authentication message")
		_ = c.conn.Close() // Attempt to close the connection
		return err
	}

	c.logger.Info().Msg("Successfully authenticated with Alpaca")
	return nil
}
