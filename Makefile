.PHONY: help build run test clean lint fmt vet install-tools migrate-up migrate-down

# Variables
APP_NAME=my-go-driver
MAIN_PATH=./cmd/api
BUILD_DIR=./bin
GO=go
GOFLAGS=-v

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install dependencies
	$(GO) mod download
	$(GO) mod tidy

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

run: ## Run the application
	$(GO) run $(MAIN_PATH)/main.go

dev: ## Run the application in development mode with auto-reload (requires air)
	air

test: ## Run tests
	$(GO) test -v -race -coverprofile=coverage.out ./...

test-coverage: test ## Run tests with coverage report
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build files
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@$(GO) clean

fmt: ## Format code
	$(GO) fmt ./...

vet: ## Run go vet
	$(GO) vet ./...

lint: ## Run golangci-lint
	golangci-lint run

install-tools: ## Install development tools
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/air-verse/air@latest
	$(GO) install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

docker-build: ## Build Docker image
	docker build -t $(APP_NAME):latest .

docker-run: ## Run Docker container
	docker-compose up -d

docker-stop: ## Stop Docker container
	docker-compose down

# Database migrations
migrate-up: ## Run all database migrations
	@echo "Running database migrations..."
	migrate -path migrations -database "$(DB_URL)" up

migrate-down: ## Rollback last migration
	@echo "Rolling back last migration..."
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-drop: ## Drop all tables (DANGEROUS!)
	@echo "Dropping all tables..."
	migrate -path migrations -database "$(DB_URL)" drop -f

migrate-force: ## Force set migration version (usage: make migrate-force version=1)
	migrate -path migrations -database "$(DB_URL)" force $(version)

migrate-version: ## Show current migration version
	migrate -path migrations -database "$(DB_URL)" version

migrate-create: ## Create a new migration (usage: make migrate-create name=create_users_table)
	@echo "Creating migration: $(name)"
	migrate create -ext sql -dir migrations -seq $(name)

install-migrate: ## Install golang-migrate tool
	@echo "Installing golang-migrate..."
	@go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.DEFAULT_GOAL := help
