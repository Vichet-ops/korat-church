.PHONY: help setup up stop logs db-create-migration db-migrate db-rollback rebuild-frontend

# Default target
help: ## Show available commands
	@echo "🚀 Manage App - Development Commands"
	@echo ""
	@echo "📋 Main Commands:"
	@echo "   make setup     - First time setup (migrations, seed data, etc.)"
	@echo "   make up        - Start development session (packages, migrations, containers)"
	@echo "   make stop      - Stop all services (non-destructive)"
	@echo ""
	@echo "🗄️  Database Commands:"
	@echo "   make db-create-migration name=migration_name  - Create new migration"
	@echo "   make db-migrate                               - Run pending migrations"
	@echo "   make db-rollback                              - Rollback last migration"
	@echo ""
	@echo "🌐 Your app will be available at:"
	@echo "   Frontend: http://localhost:3000"
	@echo "   Backend:  http://localhost:8080"

# First time setup - run once when cloning repo
setup: ## First time setup (migrations, seed data, environment)
	@echo "🔧 First time setup..."
	@if [ ! -f .env ]; then \
		echo "📄 Copying .env.example to .env..."; \
		cp .env.example .env; \
		echo "✅ Environment file created"; \
	else \
		echo "✅ Environment file already exists"; \
	fi
	@echo "🐳 Starting database..."
	@docker-compose up -d postgres
	@echo "⏳ Waiting for database to be ready..."
	@sleep 5
	@echo "🗄️  Running migrations..."
	@$(MAKE) db-migrate
	@echo "🌱 Seeding data..."
	@# TODO: Add seed data command when we have it
	@echo "✅ Setup complete! Run 'make up' to start development."

# Daily development startup
up: ## Start development session (packages, migrations, containers)
	@echo "🚀 Starting development session..."
	@echo "📦 Checking Go dependencies..."
	@if command -v go >/dev/null 2>&1; then \
		go mod download; \
		echo "✅ Go dependencies ready"; \
	else \
		echo "⚠️  Go not installed locally (will use Docker)"; \
	fi
	@echo "🐳 Starting database..."
	@docker-compose up -d postgres
	@echo "⏳ Waiting for database to be ready..."
	@sleep 5
	@echo "🗄️  Running any pending migrations..."
	@$(MAKE) db-migrate
	@echo "🐳 Starting all services..."
	@docker-compose up --build -d
	@echo ""
	@echo "🎉 All services started successfully!"
	@echo ""
	@echo "🌐 Your application is now running:"
	@echo "   Frontend:  http://localhost:3000"
	@echo "   Backend:   http://localhost:8080"
	@echo "   API Health: http://localhost:8080/api/health"
	@echo ""
	@echo "📋 Useful commands:"
	@echo "   make logs      - View all logs"
	@echo "   make stop      - Stop all services"
	@echo "   make rebuild-frontend - Rebuild and restart frontend container"
	@echo "   make help      - See all commands"

# Stop services (non-destructive)
stop: ## Stop all services (non-destructive)
	@echo "🛑 Stopping all services..."
	@docker-compose stop
	@echo "✅ All services stopped. Data preserved."

# View logs
logs: ## View logs from all services
	@docker-compose logs -f

# Database migration commands
db-create-migration: ## Create new migration (usage: make db-create-migration name=create_users)
	@if [ -z "$(name)" ]; then \
		echo "❌ Please provide migration name: make db-create-migration name=migration_name"; \
		exit 1; \
	fi
	@echo "📝 Creating migration: $(name)"
	@if command -v migrate >/dev/null 2>&1; then \
		migrate create -ext sql -dir migrations $(name); \
	else \
		docker run --rm -v "$(PWD)/migrations:/migrations" migrate/migrate:latest create -ext sql -dir /migrations $(name); \
	fi
	@echo "✅ Migration created in migrations/ directory"

db-migrate: ## Run pending migrations
	@echo "🗄️  Running database migrations..."
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make setup' first."; \
		exit 1; \
	fi
	@export $$(grep -v '^#' .env | grep -v '^$$' | xargs) && \
	if command -v migrate >/dev/null 2>&1; then \
		migrate -path migrations -database "$$DATABASE_URL" up; \
	else \
		docker run --rm -v "$(PWD)/migrations:/migrations" --network host migrate/migrate:latest \
		-path /migrations -database "$$DATABASE_URL" up; \
	fi
	@echo "✅ Migrations complete"

db-rollback: ## Rollback last migration
	@echo "⏪ Rolling back last migration..."
	@if [ ! -f .env ]; then \
		echo "❌ .env file not found. Run 'make setup' first."; \
		exit 1; \
	fi
	@export $$(grep -v '^#' .env | grep -v '^$$' | xargs) && \
	if command -v migrate >/dev/null 2>&1; then \
		migrate -path migrations -database "$$DATABASE_URL" down 1; \
	else \
		docker run --rm -v "$(PWD)/migrations:/migrations" --network host migrate/migrate:latest \
		-path /migrations -database "$$DATABASE_URL" down 1; \
	fi
	@echo "✅ Rollback complete"

# Rebuild frontend image and restart container
rebuild-frontend: ## Rebuild and restart frontend container
	@echo "🔁 Rebuilding frontend image..."
	@docker-compose build frontend
	@echo "🚀 Restarting frontend container..."
	@docker-compose up -d frontend
	@echo "✅ Frontend rebuilt and restarted. Visit http://localhost:3000"