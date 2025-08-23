.PHONY: help
help:
	@grep -h -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: format
format:
	@go fmt

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test ./...

.PHONY: build
build:
	@go build -v ./main.go

.PHONY: run
run:
	@go run -v ./main.go
