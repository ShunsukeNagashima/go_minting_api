.PHONY: help generate-api generate-models
.DEFAULT_GOAL := help

help: ## Show options
				@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | \
								awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

generate-api: ## Generate api codes by OpenAPI
	oapi-codegen -config docs/configs/chi-server.config.yaml docs/openapi.yaml

generate-models: ## Generate models codes by OpenAPI
	oapi-codegen -config docs/configs/models.config.yaml docs/openapi.yaml
