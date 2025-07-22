# SignalFlow

[![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen)](https://golang.org/)
[![License](https://img.shields.io/github/license/maxviazov/signal-flow)](LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/commits/main)
[![Issues](https://img.shields.io/github/issues/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/issues)
[![Stars](https://img.shields.io/github/stars/maxviazov/signal-flow?style=social)](https://github.com/maxviazov/signal-flow/stargazers)

---

**SignalFlow** is an event-driven platform for real-time market data ingestion and analysis, enriched with AI-based news sentiment features.

---

## 📚 Table of Contents

- [About the Project](#about-the-project)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Roadmap](#roadmap)
- [Contribution](#contribution)
- [License](#license)

---

## 🧠 About the Project

SignalFlow provides a scalable infrastructure to consume, process, and store real-time financial data from sources like Alpaca, with optional extensions for news analytics using AI.

---

## 🏗 Architecture

- **sf-ingestor**: main entrypoint service for consuming and logging streamed data
- **internal/**: modular implementation for configuration, DB, and client interaction
- **pkg/**: utilities such as logger and streamer abstraction
- **docker-compose**: PostgreSQL, services orchestration for local development

---

## 🧰 Tech Stack

| Layer           | Technology                     |
|----------------|---------------------------------|
| Language        | Go                              |
| Configuration   | Viper + YAML                    |
| Logging         | Zerolog                         |
| API/Streaming   | WebSockets (Alpaca)             |
| Storage         | PostgreSQL (pgxpool)            |
| Containerization| Docker + Compose                |
| Linting         | golangci-lint                   |

---

## 🚀 Getting Started

### Requirements

- Go 1.21+
- Docker & Docker Compose
- Alpaca API keys

### Setup

```bash
git clone https://github.com/maxviazov/signal-flow.git
cd signal-flow

cp config.yaml.example config.yaml
# or use environment variables via .env

make build
make run
```

---

## 🛣 Roadmap

- [x] Viper-based configuration
- [x] Zerolog logging
- [x] Alpaca WebSocket integration
- [ ] Kafka support
- [ ] Redis Pub/Sub middleware
- [ ] Public WebSocket API
- [ ] React-based dashboard (maybe 😉)

---

## 🤝 Contribution

Pull requests are welcome! For major changes, please open an issue first and follow the project code style via `go fmt`.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---