.DEFAULT_GOAL := help
ARG =

.PHONY: help
help: ## help for skeleton
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# cli
.PHONY: cli
cli: ## go run cmd/cli/main.go
	go run cmd/skeleton/main.go $(ARG)

# go test
.PHONY: go_test
go_test: ## go test ./...
	GO_ENV=test go test -tags=unit ./...

# go bench
.PHONY: go_bench
go_bench: ## go test -tags=unit -bench=. -benchmem ./...
	GO_ENV=test go test -tags=unit -bench=. -benchmem ./...
	@echo '実行回数 実行時間(ns/op) メモリ割当容量(B/op) メモリ割当回数(allocs/op)'

# go mod tidy
.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy

# lint
.PHONY: lint
lint: ## golangci-lint run
	golangci-lint run

# wire
.PHONY: wire
wire: ## go generate ./cmd/...
	go generate ./cmd/...

# gen
.PHONY: gen
gen: ## go generate ./...
	go generate ./...
