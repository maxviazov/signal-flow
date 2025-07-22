# SignalFlow

[![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen)](https://golang.org/)
[![License](https://img.shields.io/github/license/maxviazov/signal-flow)](LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/commits/main)
[![Issues](https://img.shields.io/github/issues/maxviazov/signal-flow)](https://github.com/maxviazov/signal-flow/issues)
[![Stars](https://img.shields.io/github/stars/maxviazov/signal-flow?style=social)](https://github.com/maxviazov/signal-flow/stargazers)

> **SignalFlow** — event-driven платформа для анализа рыночных данных в реальном времени, обогащённая ИИ-анализом новостей.

---

## 📚 Table of Contents

- [🔍 About the Project](#-about-the-project)
- [📐 Architecture](#-architecture)
- [🧰 Tech Stack](#-tech-stack)
- [🚀 Getting Started](#-getting-started)
- [🗺️ Roadmap](#-roadmap)
- [🤝 Contribution](#-contribution)
- [🪪 License](#-license)

---

## 🔍 About the Project

SignalFlow агрегирует рыночные данные и потоковые новости, обрабатывает их в реальном времени и предоставляет удобный интерфейс для аналитики.  
Система проектировалась с учётом модульности, расширяемости и надёжности.

---

## 📐 Architecture

```mermaid
flowchart LR
  A[Alpaca Market Data] --> B[Stream Client]
  B --> C[Kafka (в будущем)]
  C --> D[PostgreSQL Storage]
  D --> E[Analytics Engine]
  E --> F[REST/WebSocket API]
```

> В текущей версии реализован прямой поток данных → хранилище → логгер

---

## 🧰 Tech Stack

| Layer          | Technology                |
|----------------|---------------------------|
| Language       | Go                        |
| Configuration  | Viper + YAML              |
| Logging        | Zerolog                   |
| API/Streaming  | WebSockets (Alpaca)       |
| Storage        | PostgreSQL (pgxpool)      |
| Container      | Docker + Compose          |
| Linting        | golangci-lint             |

---

## 🚀 Getting Started

### 🔧 Requirements

- Go 1.21+
- Docker & Docker Compose
- Alpaca API ключи

### 📦 Setup

```bash
git clone https://github.com/maxviazov/signal-flow.git
cd signal-flow
cp config.yaml.example config.yaml
# укажи переменные в .env или config.yaml
make build
make run
```

---

## 🗺️ Roadmap

- [x] Конфигурация через Viper
- [x] Логгирование Zerolog
- [x] WebSocket-интеграция с Alpaca
- [ ] Подключение Kafka
- [ ] Интеграция Redis Pub/Sub
- [ ] Интерфейс аналитики на WebSocket API
- [ ] Веб-дашборд на React (возможно)

---

## 🤝 Contribution

Pull requests приветствуются! Для крупных изменений — сначала открой issue.  
Форматируй код с `go fmt`, соблюдай стиль.

---

## 🪪 License

Этот проект лицензирован под [MIT](LICENSE).
