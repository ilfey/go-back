.PHONY: vet fmt test build

vet:
	go vet ./...

fmt:
	go fmt ./...

test:
	go test -v -race -timeout 30s ./...

build:
	go build -v ./cmd/server

.DEFAULT_GOAL := build