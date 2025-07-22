[![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen)](https://golang.org/)
[![License](https://img.shields.io/github/license/maxviazov/signal-flow)](LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/commits/main)
[![Issues](https://img.shields.io/github/issues/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/issues)
[![Stars](https://img.shields.io/github/stars/maxviazov/signal-flow?style=social)](https://github.com/maxviazov/signal-flow/stargazers)

# 📡 signal-flow

**signal-flow** is a personal project I'm building to experiment with streaming market data in real time using Go, Alpaca WebSocket API, and PostgreSQL.

The goal is to ingest live trade data, log it, and (soon) store and process it for analytics, AI-driven sentiment analysis, and potential signal generation.

> ⚠️ This project is a work-in-progress — but it already has a solid structure for scaling into a real-time data platform.

---

## 💡 What it does (right now)

- Connects to Alpaca’s WebSocket (v2/iex)
- Authenticates with API key + secret (via `.env` or config)
- Subscribes to trades for selected tickers (like AAPL, GOOGL, TSLA)
- Receives and logs raw JSON market data using `zerolog`
- Handles shutdown gracefully (`SIGINT`, `SIGTERM`)
- Uses `pgxpool` for PostgreSQL connection management (writing coming next!)
- All configuration is powered by `Viper`

---

## 🧱 Tech stack

| What | Why |
|------|-----|
| Go | Fast, typed, great for stream processing |
| WebSocket | Real-time data from Alpaca |
| PostgreSQL | For storing structured trade data |
| Zerolog | Minimalistic, structured logging |
| Viper | Flexible config loading (yaml + env) |
| Docker Compose | Local dev with Postgres (coming soon) |

---

## 🧰 Project structure

```bash
.
├── cmd/sf-ingestor      # App entrypoint
├── internal/
│   ├── config           # Viper config
│   ├── client/streamers # Alpaca WebSocket client
│   └── repository/      # Postgres connection
├── pkg/logger           # Custom zerolog wrapper
├── config.yaml.example  # Config template
├── logs/                # App logs
└── test/                # Unit tests (loggers, utils)
```

---

## 🚀 Run it locally

1. Clone the repo
2. Create `.env` or copy and fill `config.yaml.example`
3. Run the app:

```bash
go run ./cmd/sf-ingestor
```

Coming soon: `docker-compose` for local PostgreSQL + API

---

## 🧪 Sample config

```yaml
streamers:
  base_url: "https://paper-api.alpaca.markets"
  stream_url: "wss://stream.data.alpaca.markets/v2/iex"
  api_key: "${ALPACA_API_KEY}"
  api_secret: "${ALPACA_API_SECRET}"

log:
  level_console: "info"
  level_file: "debug"

postgres:
  postgres_host: "localhost"
  postgres_port: 5432
  postgres_user: "${POSTGRES_USER}"
  postgres_password: "${POSTGRES_PASSWORD}"
  postgres_db: "signal_flow"
```

---

## 📌 Roadmap

- [x] Connect to Alpaca WS and authenticate
- [x] Subscribe to trades and log incoming data
- [ ] Parse & store trade messages in Postgres
- [ ] Stream processed data into downstream services (Kafka?)
- [ ] Add sentiment analysis from news
- [ ] Expose REST or gRPC API to consume signals
- [ ] Dockerize + CI

---

## 👨‍💻 About me

I'm [Max](https://github.com/maxviazov) — a backend developer from Israel passionate about real-time systems, clean Go code, and data processing.  
Feel free to reach out, open issues, or suggest features 🙌

---

> Built with ☕ and way too many logging statements.  
> Stay tuned!
