APP_DIR := ./src/cmd
IMAGE_NAME := beershop
IMAGE_TAG := latest

.PHONY: run-dev build run-prod stop clear

run-dev:
	air

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