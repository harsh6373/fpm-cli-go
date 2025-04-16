# Project settings
BINARY_NAME=fpm-cli
PKG=github.com/harsh6373/fpm-cli-go

# Output directories
BUILD_DIR=bin
SRC_DIR=.
VERSION ?= $(shell git describe --tags --always)

# Platforms to build
PLATFORMS=\
  "linux amd64" \
  "darwin amd64" \
  "darwin arm64" \
  "windows amd64"

# ===== Commands =====

all: build

# Build for current platform
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run tests
test:
	go test ./... -v

# Clean build output
clean:
	rm -rf $(BUILD_DIR)

# Format code
fmt:
	go fmt ./...

# Install locally
install:
	go install $(PKG)@latest

# Generate shell completions
completion:
	go run main.go completion bash > completions/$(BINARY_NAME).bash
	go run main.go completion zsh > completions/_$(BINARY_NAME)
	go run main.go completion fish > completions/$(BINARY_NAME).fish
	go run main.go completion powershell > completions/$(BINARY_NAME).ps1

# Build for all platforms
release:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
	  os=$$(echo $$platform | cut -d' ' -f1); \
	  arch=$$(echo $$platform | cut -d' ' -f2); \
	  output_name=$(BUILD_DIR)/$(BINARY_NAME)-$$os-$$arch; \
	  if [ $$os = "windows" ]; then output_name=$$output_name.exe; fi; \
	  echo "🔨 Building $$output_name"; \
	  GOOS=$$os GOARCH=$$arch go build -ldflags="-X main.version=$(VERSION)" -o $$output_name $(SRC_DIR); \
	done

.PHONY: all build test clean fmt install completion release
