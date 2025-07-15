# SignalFlow

Event-driven platform for real-time market data analysis, enriched with AI-driven news sentiment.

## Table of Contents

- [About The Project](#about-the-project)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Roadmap](#roadmap)
- [License](#license)

## About The Project

SignalFlow is a backend system designed to ingest real-time market data, calculate technical indicators, and enrich them
with AI-powered news analysis. The primary goal is to provide a robust foundation for building automated trading
strategies or advanced market monitoring tools.

This project serves as a portfolio piece demonstrating a modern, event-driven microservices architecture using Go,
Python, and Google Cloud Platform.

## Architecture

The system is composed of several independent microservices communicating via a message broker:

- **`sf-ingestor`**: A Go service responsible for connecting to market data WebSocket APIs and publishing raw ticks to a
  message queue.
- **`sf-calculator`**: A Go service that consumes raw data, calculates technical indicators, stores data in a
  time-series database, and publishes trading signals.
- **`sf-news-analyzer`**: A Python service that periodically fetches financial news, analyzes sentiment using an LLM,
  and publishes the analysis.

*A diagram will be added here later.*

## Tech Stack

- **Languages**: Go, Python
- **Cloud Platform**: Google Cloud Platform (GCP)
- **Messaging**: Google Cloud Pub/Sub
- **Database**: TimescaleDB (PostgreSQL extension for time-series data)
- **AI**: Google Gemini API
- **Deployment**: Cloud Run, Docker
- **CI/CD**: GitHub Actions

## Getting Started

To get a local copy up and running, follow these steps.

### Prerequisites

- Go (latest version)
- Docker & Docker Compose
- `make` utility

### Installation

1. Clone the repo
   ```sh
   git clone [https://github.com/your_username/signal-flow.git](https://github.com/maxviazov/signal-flow.git)