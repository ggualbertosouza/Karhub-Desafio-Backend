APP_DIR := ./src/cmd
IMAGE_NAME := beershop
IMAGE_TAG := latest
DEV_COMPOSE := compose.dev.yml

.PHONY: run-dev build run-prod stop clear db-dev-up db-dev-down db-dev-logs

# Development
run-dev:
	air

run-tests:
	go test -v ./...

db-dev-up:
	@echo "Starting development database..."
	docker compose -f $(DEV_COMPOSE) up -d

db-dev-down:
	@echo "Stopping development database..."
	docker compose -f $(DEV_COMPOSE) down

db-dev-logs:
	docker compose -f $(DEV_COMPOSE) logs -f postgres

# Production
build:
	@echo "Building image"
	docker build \
		-f ./docker/app/Dockerfile \
		-t $(IMAGE_NAME):$(IMAGE_TAG) \
		.

run-prod: stop build
	@echo "Running container"
	docker run -d \
		--env-file .env \
		-p 3001:3001 \
		--name $(IMAGE_NAME) \
		$(IMAGE_NAME):$(IMAGE_TAG)

stop:
	@echo "Stopping container"
	- docker stop $(IMAGE_NAME)
	- docker rm $(IMAGE_NAME)

clear: stop
	@echo "Cleaning image"
	- docker rmi $(IMAGE_NAME):$(IMAGE_TAG)

logs:
	@echo "Opening logs"
	docker logs -f $(IMAGE_NAME)