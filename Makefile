# Makefile for Mars Rover Go project
# Version: 1.0.0
# Author: Alejandro Sosa <alesjohnson@hotmail.com>

# Project information
VERSION := 1.0.0

# Default target: display help
.PHONY: all
all: help

# Target: Install all Go dependencies
.PHONY: setup
setup:
	@echo "🚀 MARS ROVERS - Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed."

# Target: Build the HTTP server
.PHONY: build
build:
	@echo "🚀 MARS ROVERS - Building the HTTP server..."
	go build -o bin/mars-rover-http ./cmd/mars-rover-http
	@echo "Build complete."

# Target: Start the HTTP server
.PHONY: start
start: build
	@echo "🚀 MARS ROVERS - Starting the HTTP server on port 8080..."
	./bin/mars-rover-http

# Target: Display help message
.PHONY: help
help:
	@echo "╔═══════════════════════════════════════════╗"
	@echo "║                                           ║"
	@echo "║            M A R S  R O V E R S           ║"
	@echo "║                                           ║"
	@echo "║            Version: $(VERSION)		    ║"
	@echo "║                                           ║"
	@echo "╚═══════════════════════════════════════════╝"
	@echo
	@echo "Available commands:"
	@echo "  setup   - Install all Go dependencies"
	@echo "  build   - Build the HTTP server"
	@echo "  start   - Start the HTTP server on port 8080"
	@echo "  help    - Display this help message"
	@echo
	@echo "For more information, please refer to the README.md file."
