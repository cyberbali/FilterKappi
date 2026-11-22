.PHONY: all build test lint vuln run clean

BIN_DIR := bin
BIN_NAME := filterkappi
BINARY := $(BIN_DIR)/$(BIN_NAME)

all: build

build:
	@echo "==> Building $(BIN_NAME)..."
	go build -trimpath -ldflags="-w -s" -o $(BINARY) ./cmd/filterkappi

test:
	@echo "==> Running tests..."
	go test -v ./...

lint:
	@echo "==> Running linters..."
	go vet ./...
	@if command -v staticcheck >/dev/null 2>&1; then \
		echo "==> Running staticcheck..."; \
		staticcheck ./...; \
	fi

vuln:
	@echo "==> Running vulnerability checks..."
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "govulncheck not installed locally; run 'go install golang.org/x/vuln/cmd/govulncheck@latest' or run in CI"; \
	fi

run: build
	@echo "==> Running $(BIN_NAME)..."
	./$(BINARY)

clean:
	@echo "==> Cleaning build artifacts..."
	rm -rf $(BIN_DIR)
