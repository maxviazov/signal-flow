package streamer

type Streamer interface {
	// Connect establishes a WebSocket connection to the streaming service.
	Connect() error
	Subscribe(symbols []string) error
	Listen() error
}
