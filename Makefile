
.PHONY: setup-dynamo
setup-dynamo: ## local setup for dynamodb
	@docker pull amazon/dynamodb-local

.PHONY: run-dynamo
run-dynamo: ## run dynamodb container
	@docker run -d --rm --name dynamodb -p 8000:8000 amazon/dynamodb-local
	@export DYNAMODB_ENDPOINT=http://localhost:8000

.PHONY: stop-dynamo
stop-dynamo: ## stop dynamodb container
	@docker container stop dynamodb

.PHONY: dynamo-admin
dynamo-admin: ## run dynamodb-admin
	@dynamodb-admin

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
