ROOT := $(realpath .)

.PHONY: up
## up: build and run cmd service in local
up: clean format docs
	go run ./cmd/main.go

.PHONY: docs
## docs: generate swagger docs
docs:
	swag fmt
	swag init --generalInfo ./main.go --dir ./cmd,./service/api

.PHONY: clean
## clean: remove old binaries and unused file
clean:
	go clean
	go clean -testcache
	go mod tidy

.PHONY: format
## format: go format
format:
	go fmt ./...