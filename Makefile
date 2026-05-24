.PHONY: build run test

build:
	go build -o bin/cube ./cmd

run:
	go run ./cmd

test:
	go test -v ./...
