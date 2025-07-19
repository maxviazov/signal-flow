package streamers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/maxviazov/signal-flow/internal/config"
	"github.com/maxviazov/signal-flow/pkg/streamer"
	"github.com/rs/zerolog"
	"time"
)

// StreamerClient handles the connection and communication with the Alpaca WebSocket API.
type StreamerClient struct {
	cfg    config.AlpacaConfig
	logger *zerolog.Logger
	conn   *websocket.Conn
}

// subscribeRequest defines the structure for the subscription message.
type subscribeRequest struct {
	Action string   `json:"action"`
	Trades []string `json:"trades,omitempty"`
	Quotes []string `json:"quotes,omitempty"`
	Bars   []string `json:"bars,omitempty"`
}

// TradeUpdate defines the structure for a trade message from Alpaca.
type TradeUpdate struct {
	Type      string  `json:"T"`
	Symbol    string  `json:"S"`
	Price     float64 `json:"p"`
	Size      int64   `json:"s"`
	Timestamp string  `json:"t"`
	Condition string  `json:"c,omitempty"`
}

// authRequest defines the structure for the authentication message.
type authRequest struct {
	Action string `json:"action"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// responseMsg is a generic struct to parse the type of response message from Alpaca.
type responseMsg struct {
	Type string `json:"T"`
	Msg  string `json:"msg"`
}

// New creates a new Alpaca client that implements the streamer.Streamer interface.
func New(cfg config.AlpacaConfig, logger *zerolog.Logger) streamer.Streamer {
	return &StreamerClient{
		cfg:    cfg,
		logger: logger,
	}
}

// Connect establishes a WebSocket connection, sets up handlers, and authenticates.
func (c *StreamerClient) Connect() error {
	c.logger.Info().Str("url", c.cfg.StreamURL).Msg("Connecting to Alpaca WebSocket...")

	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.StreamURL, nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to establish WebSocket connection")
		return err
	}
	c.conn = conn
	c.logger.Info().Msg("WebSocket connection established")

	// Set a Ping handler to respond to the server's health checks.
	c.conn.SetPingHandler(func(appData string) error {
		c.logger.Info().Str("payload", appData).Msg("Ping received")
		err := c.conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(5*time.Second))
		if err != nil {
			c.logger.Error().Err(err).Msg("Failed to send pong")
		} else {
			c.logger.Info().Msg("Pong sent")
		}
		return err
	})

	// Set a Close handler for graceful shutdown logging.
	c.conn.SetCloseHandler(func(code int, text string) error {
		c.logger.Info().Int("code", code).Str("text", text).Msg("WebSocket connection closed")
		return nil
	})

	// Authenticate
	authMsg := authRequest{
		Action: "auth",
		Key:    c.cfg.APIKey,
		Secret: c.cfg.APISecret,
	}

	c.logger.Info().Msg("Sending authentication request...")
	if err := c.conn.WriteJSON(authMsg); err != nil {
		c.logger.Error().Err(err).Msg("Failed to send authentication message")
		_ = c.conn.Close()
		return err
	}

	// Wait for the authentication confirmation before proceeding.
	if err := c.waitForResponse("authenticated"); err != nil {
		c.logger.Error().Err(err).Msg("Did not receive authentication confirmation")
		return err
	}

	c.logger.Info().Msg("Successfully authenticated with Alpaca")
	return nil
}

// Subscribe sends a subscription request and waits for confirmation.
func (c *StreamerClient) Subscribe(symbols []string) error {
	c.logger.Info().Strs("symbols", symbols).Msg("Subscribing to trades")
	msg := subscribeRequest{
		Action: "subscribe",
		Trades: symbols,
	}

	if err := c.conn.WriteJSON(msg); err != nil {
		c.logger.Error().Err(err).Msg("Failed to subscribe to trades")
		_ = c.conn.Close()
		return err
	}

	// Wait for the subscription confirmation.
	if err := c.waitForResponse("subscribed"); err != nil {
		c.logger.Error().Err(err).Msg("Did not receive subscription confirmation")
		return err
	}

	c.logger.Info().Msg("Successfully subscribed to trades")
	return nil
}

// Listen starts an infinite loop to read market data from the WebSocket.
func (c *StreamerClient) Listen() error {
	c.logger.Info().Msg("Starting to listen for market data...")
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.logger.Error().Err(err).Msg("Error during message reading")
			return err
		}
		fmt.Printf("\033[33m%s\\033[0m\\n\", string(message)")
		// For now, we just log the raw message.
		// In the future, this is where you will parse the TradeUpdate and save it.
		c.logger.Info().RawJSON("market_data", message).Msg("Received data")
	}
}

// waitForResponse reads messages from the connection until it finds one
// with a "msg" field matching the expectedMsg.
func (c *StreamerClient) waitForResponse(expectedMsg string) error {
	for {
		_, msgBytes, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}

		c.logger.Debug().RawJSON("response", msgBytes).Msg("Checking response from server")

		var responses []responseMsg
		if err := json.Unmarshal(msgBytes, &responses); err != nil {
			c.logger.Warn().Err(err).Msg("Could not unmarshal response as an array, trying as single object")
			var singleResponse responseMsg
			if err2 := json.Unmarshal(msgBytes, &singleResponse); err2 != nil {
				c.logger.Warn().Err(err2).Msg("Could not unmarshal response as single object, skipping message")
				continue
			}
			responses = append(responses, singleResponse)
		}

		for _, r := range responses {
			if r.Type == "success" && r.Msg == expectedMsg {
				c.logger.Info().Str("expected", expectedMsg).Msg("Got expected response from server")
				return nil
			}
			if r.Type == "error" {
				return fmt.Errorf("received error from streamers: %s", string(msgBytes))
			}
		}
	}
}
