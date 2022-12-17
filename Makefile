.PHONY: help generate-api generate-models build build-local up down logs ps
.DEFAULT_GOAL := help

DOCKER_TAG := latest
help: ## Show options
				@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | \
								awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

generate-api: ## Generate api codes by OpenAPI
	oapi-codegen -config docs/configs/chi-server.config.yaml docs/openapi.yaml

generate-models: ## Generate models codes by OpenAPI
	oapi-codegen -config docs/configs/models.config.yaml docs/openapi.yaml

build: ## Build docker image to deploy
				docker build -t shunsukenagashima/gotodo:${DOCKER_TAG} \
								--target deploy ./

build-local: ## Build docker image to local development
				docker compose build --no-cache

up: ## Do docker compose up with hot reload
				docker compose up -d

down: ## Do docker compose down
				docker compose down

logs: ## Tail docker compose logs
				docker compose logs -f

ps: ## Check container status
				docker compose ps