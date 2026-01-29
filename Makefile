APP_DIR := ./src/cmd
IMAGE_NAME := beershop
IMAGE_TAG := latest

DEV_COMPOSE := compose.dev.yml
PROD_COMPOSE := compose.prod.yml

.PHONY: run-dev run-tests db-dev-up db-dev-down db-dev-logs

run-dev:
	air

run-tests:
	go test -v ./...

db-dev-up:
	docker compose -f $(DEV_COMPOSE) up -d

db-dev-down:
	docker compose -f $(DEV_COMPOSE) down

db-dev-logs:
	docker compose -f $(DEV_COMPOSE) logs -f postgres

# Production
.PHONY: prod-up prod-down prod-logs prod-rebuild

prod-up:
	@echo "Starting production environment"
	docker compose -f $(PROD_COMPOSE) up -d

prod-down:
	@echo "Stopping production environment"
	docker compose -f $(PROD_COMPOSE) down

prod-logs:
	docker compose -f $(PROD_COMPOSE) logs -f

prod-rebuild:
	@echo "Rebuilding and starting production environment"
	docker compose -f $(PROD_COMPOSE) up -d --build