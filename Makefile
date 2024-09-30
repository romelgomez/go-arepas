# Define variables
GO := go
PRISMA_CLIENT := github.com/steebchen/prisma-client-go
PROJECT_NAME := go-arepas
BINARY_NAME := $(PROJECT_NAME)
DATABASE_URL ?= "postgres://developer:developer@localhost:5432/go_arepas?schema=public"

# Define Prisma commands
PRISMA_GENERATE := $(GO) run $(PRISMA_CLIENT) generate
PRISMA_MIGRATE_DEV := $(GO) run $(PRISMA_CLIENT) migrate dev --schema=./prisma/schema.prisma
PRISMA_MIGRATE_DEPLOY := $(GO) run $(PRISMA_CLIENT) migrate deploy --schema=./prisma/schema.prisma

# Define targets
.PHONY: help build run test clean generate migrate-dev migrate-deploy setup start env lint fmt check

help:  ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
		helpMessage = match(lastLine, /^# (.*)/) ? substr(lastLine, RSTART + 2, RLENGTH) : ""; \
		printf "  %-20s %s\n", $$1, helpMessage; \
	} { lastLine = $$0 }' $(MAKEFILE_LIST)

# Build and run targets
build:  ## Build the Go application
	@echo "Building the application..."
	$(GO) build -o $(BINARY_NAME) .

run: build  ## Build and run the Go application
	@echo "Running the application..."
	./$(BINARY_NAME)

test:  ## Run tests
	@echo "Running tests..."
	$(GO) test -v ./...

clean:  ## Clean up generated files
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Prisma-related targets
generate:  ## Generate the Prisma client
	@echo "Generating Prisma client..."
	$(PRISMA_GENERATE)

migrate-dev:  ## Apply Prisma migrations for development
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "Error: DATABASE_URL is not set. Please set it before running this target."; \
		exit 1; \
	fi
	@echo "Applying Prisma migrations for development..."
	$(PRISMA_MIGRATE_DEV)

migrate-deploy:  ## Apply Prisma migrations for production
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "Error: DATABASE_URL is not set. Please set it before running this target."; \
		exit 1; \
	fi
	@echo "Applying Prisma migrations for production..."
	$(PRISMA_MIGRATE_DEPLOY)

# Setup and maintenance targets
setup:  ## Set up the project (download dependencies and generate Prisma client)
	@echo "Setting up the project..."
	$(GO) mod download
	$(MAKE) generate

start: generate run  ## Generate Prisma client and run the application

env:  ## Display current environment variables
	@echo "Current environment settings:"
	@echo "DATABASE_URL=$(DATABASE_URL)"

lint:  ## Run linting on the Go code
	@echo "Running golint..."
	@command -v golint >/dev/null 2>&1 || { echo >&2 "golint is not installed. Install it by running: go install golang.org/x/lint/golint@latest"; exit 1; }
	golint ./...

fmt:  ## Format the Go code
	@echo "Formatting Go code..."
	$(GO) fmt ./...

check: fmt lint test  ## Run formatting, linting, and tests
	@echo "Running code quality checks..."
