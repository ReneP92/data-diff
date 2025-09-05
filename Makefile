# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=data-diff
BINARY_UNIX=$(BINARY_NAME)_unix

# Version information
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_TIME ?= $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)

# Build flags
LDFLAGS=-ldflags "-X github.com/renepersau/data-diff/internal/commands.Version=$(VERSION) -X github.com/renepersau/data-diff/internal/commands.BuildTime=$(BUILD_TIME) -X github.com/renepersau/data-diff/internal/commands.GitCommit=$(GIT_COMMIT)"

.PHONY: all build clean test coverage deps fmt vet lint run install docker-build docker-run help

all: test build

build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v ./cmd/data-diff

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) -v ./cmd/data-diff

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

coverage: test-coverage

deps:
	$(GOMOD) download
	$(GOMOD) tidy

fmt:
	$(GOCMD) fmt ./...

vet:
	$(GOCMD) vet ./...

lint:
	golangci-lint run

run: build
	./$(BINARY_NAME)

install: build
	cp $(BINARY_NAME) /usr/local/bin/

docker-build:
	docker build -t $(BINARY_NAME):$(VERSION) .
	docker tag $(BINARY_NAME):$(VERSION) $(BINARY_NAME):latest

docker-run:
	docker run --rm -it $(BINARY_NAME):latest

# Development helpers
dev-setup:
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "Development environment setup complete"

# Release helpers
release-build:
	@echo "Building release binaries..."
	@mkdir -p dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 ./cmd/data-diff
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd/data-diff
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 ./cmd/data-diff
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd/data-diff

help:
	@echo "Available targets:"
	@echo "  all          - Run tests and build"
	@echo "  build        - Build the binary"
	@echo "  build-linux  - Build Linux binary"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  coverage     - Run tests with coverage"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run go vet"
	@echo "  lint         - Run linter"
	@echo "  run          - Build and run the application"
	@echo "  install      - Install binary to /usr/local/bin"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  dev-setup    - Setup development environment"
	@echo "  release-build- Build release binaries for multiple platforms"
	@echo "  help         - Show this help message"

