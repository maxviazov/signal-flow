SHELL := /bin/sh

GOLINT := /opt/homebrew/bin/golangci-lint
GOLINT_VERSION := v2.2.2

.PHONY: run clean lint

clean:
	@echo "==> Cleaning up..."
	@go clean -cache

run: clean
	@echo "==> Running the application..."
	@set -a;\
	. .env; \
	set +a; \
	go run ./cmd/sf-ingestor/main.go

lint:
	@echo "==> Running linter..."
	@if [ ! -f $(GOLINT) ]; then \
		echo "Installing golangci-lint..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLINT)) $(GOLINT_VERSION); \
	fi
	@$(GOLINT) run --timeout 5m