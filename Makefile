.PHONY: build dev test

build:
	go run ./cmd/site

dev:
	go tool modd

test:
	go test ./cmd/site
