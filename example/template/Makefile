.PHONY: help
help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# docker-compose
.PHONY: up
up: ## docker-compose up
	docker-compose up

.PHONY: reup
reup: ## docker-compose up -d --build
	docker-compose up --build

.PHONY: down
down: ## docker-compose down
	docker-compose down

.PHONY: rm
rm: ## docker-compose down and remove
	docker-compose down -v && docker-compose rm -v

.PHONY: ps
ps: ## docker-compose ps
	docker-compose ps
