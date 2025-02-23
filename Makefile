GO := go
SORT := goimports-reviser
BIN_DIR := bin
BINARY := $(BIN_DIR)/collections

.PHONY: all init tidy build sort-imports format lint test clean docker-up docker-down

# Default target
all: clean format lint test build

help: # Show this help message
	@printf "Usage: make \033[36m[target]\033[0m\n\n"
	@printf "\033[1mTargets:\033[0m\n"
	@awk 'BEGIN {FS = ":.*?# "} /^[a-zA-Z_-]+:.*?# / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

tidy: # Run go mod tidy
	@echo "Tidying..."
	@$(GO) mod tidy

init: tidy # Initialize the project
	@echo "Initializing..."
	@$(GO) install -v github.com/incu6us/goimports-reviser/v3@latest

build: # Build the project
	@echo "Building..."
	@$(GO) build -o $(BINARY) ./...

sort-imports: # Sort imports
	@echo "Sorting imports..."
	@$(SORT) -rm-unused -set-alias -format ./... > /dev/null 2>&1

format: sort-imports # Format the code
	@echo "Formatting..."
	@$(GO) fmt ./...

lint: # Lint the project
	@echo "Linting..."
	@$(GO) vet ./...

test: # Run tests
	@echo "Testing..."
	 @$(GO) test ./... > /dev/null 2>&1 || (echo "Tests failed" && $(GO) test ./...)

clean: # Clean the build artifacts
	@echo "Cleaning..."
	@$(RM) -rf $(BIN_DIR)

docker-up: # Start the docker-compose services
	@echo "Starting docker-compose services..."
	@docker compose up --build

docker-down: # Stop the docker-compose services
	@echo "Stopping docker-compose services..."
	@docker compose down