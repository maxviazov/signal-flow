SHELL := /bin/sh

# We install the linter locally to the project to avoid permission issues and for version isolation.
GOLINT := $(CURDIR)/bin/golangci-lint
GOLINT_VERSION := v2.2.2

.PHONY: run clean lint

clean:
	@echo "==> Cleaning up..."
	@go clean -cache
	@rm -rf $(CURDIR)/bin

run:
	@echo "==> Running the application..."
	@set -a;\
	if [ -f .env ]; then . .env; fi; \
	set +a; \
	go run ./cmd/sf-ingestor/main.go

lint:
	@echo "==> Running linter..."
	@if [ ! -f $(GOLINT) ]; then \
		echo "Installing golangci-lint to $(GOLINT)..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLINT)) $(GOLINT_VERSION); \
	fi
	@$(GOLINT) run ./... --timeout 5m