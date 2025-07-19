# SignalFlow

Real-time market data ingestion service for financial data analysis and trading applications.

## Table of Contents

- [About The Project](#about-the-project)
- [Current Implementation](#current-implementation)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage](#usage)
- [Development](#development)
- [Roadmap](#roadmap)
- [License](#license)

## About The Project

SignalFlow is a backend system designed to ingest real-time market data from financial APIs. Currently, the project
implements a robust market data ingestion service that connects to Alpaca's WebSocket API to stream live trade data.

This project serves as a foundation for building automated trading strategies, market monitoring tools, and financial
data analysis applications. It demonstrates modern Go development practices with clean architecture, structured logging,
and containerized deployment.

## Current Implementation

The project currently consists of:

**`sf-ingestor`** - A Go service that:

- Connects to Alpaca's WebSocket API for real-time market data
- Authenticates using API key/secret credentials
- Subscribes to trade updates for specified symbols (AAPL, GOOGL, TSLA)
- Provides structured logging with both console and file output
- Handles WebSocket connection management with ping/pong and reconnection logic
- Supports flexible configuration via YAML files and environment variables

## Tech Stack

- **Language**: Go 1.24.5
- **WebSocket Client**: Gorilla WebSocket
- **Logging**: Zerolog (structured logging)
- **Configuration**: Viper (YAML + environment variables)
- **Validation**: Go Playground Validator
- **Database**: TimescaleDB (PostgreSQL extension for time-series data)
- **Containerization**: Docker & Docker Compose
- **Build Tool**: Make
- **Code Quality**: golangci-lint

## Getting Started

### Prerequisites

- Go 1.24.5 or later
- Docker & Docker Compose
- `make` utility
- Alpaca API credentials (free paper trading account)

### Installation

1. Clone the repository
   ```sh
   git clone https://github.com/maxviazov/signal-flow.git
   cd signal-flow
   ```

2. Set up environment variables
   ```sh
   cp .env.example .env
   # Edit .env with your Alpaca API credentials
   ```

3. Start the database
   ```sh
   docker-compose up -d postgres
   ```

4. Run the application
   ```sh
   make run
   ```

## Configuration

The application supports configuration through both YAML files and environment variables.

### Configuration File (config.yaml)

```yaml
alpaca:
  base_url: "https://paper-api.alpaca.markets"
  stream_url: "wss://stream.data.alpaca.markets/v2/iex"

log:
  level_console: "info"
  level_file: "debug"
```

### Environment Variables

Create a `.env` file with your Alpaca credentials:

```env
ALPACA_API_KEY=your_api_key_here
ALPACA_API_SECRET=your_api_secret_here
POSTGRES_USER=signalflow
POSTGRES_PASSWORD=your_password_here
POSTGRES_DB=signalflow
```

### Configuration Options

- **alpaca.base_url**: Alpaca API base URL (defaults to paper trading)
- **alpaca.stream_url**: WebSocket stream URL for market data
- **alpaca.api_key**: Your Alpaca API key (set via environment variable)
- **alpaca.api_secret**: Your Alpaca API secret (set via environment variable)
- **log.level_console**: Console log level (debug, info, warn, error, fatal, panic)
- **log.level_file**: File log level (debug, info, warn, error, fatal, panic)

## Usage

### Running the Service

```sh
# Run with make (loads .env automatically)
make run

# Or run directly with Go
go run ./cmd/sf-ingestor/main.go
```

### Logs

The application creates logs in two places:

- **Console**: Formatted output with timestamps
- **File**: `logs/app.log` with detailed JSON-structured logs

### Market Data

The service currently subscribes to trade updates for:

- AAPL (Apple Inc.)
- GOOGL (Alphabet Inc.)
- TSLA (Tesla Inc.)

Trade data is logged in real-time and includes:

- Symbol
- Price
- Size (volume)
- Timestamp
- Trade conditions

## Development

### Available Make Commands

```sh
make run     # Run the application with environment variables
make clean   # Clean Go cache and remove binaries
make lint    # Run golangci-lint for code quality checks
```

### Project Structure

```
.
├── cmd/sf-ingestor/          # Application entry point
├── internal/
│   ├── client/alpaca/        # Alpaca WebSocket client
│   └── config/               # Configuration management
├── pkg/logger/               # Logging utilities
├── config.yaml               # Configuration file
├── docker-compose.yml        # Database setup
└── Makefile                  # Build automation
```

### Code Quality

The project uses golangci-lint for code quality checks. Run `make lint` to check your code before committing.

## Roadmap

Future planned features:

- [ ] **sf-calculator**: Technical indicators calculation service
- [ ] **sf-news-analyzer**: News sentiment analysis with AI
- [ ] **Database Integration**: Store market data in TimescaleDB
- [ ] **Message Queue**: Implement Pub/Sub for service communication
- [ ] **REST API**: HTTP endpoints for data access
- [ ] **Cloud Deployment**: Deploy to Google Cloud Run
- [ ] **CI/CD Pipeline**: GitHub Actions for automated testing and deployment
- [ ] **Monitoring**: Metrics and health checks
- [ ] **More Data Sources**: Support for additional market data providers

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Note**: This project is currently in active development. The sf-ingestor service is functional and ready for use,
while additional services are planned for future releases.