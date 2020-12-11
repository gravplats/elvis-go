.PHONY: help
help:
	@echo "Basic commands"
	@echo "  make build"
	@echo "  make test"

.PHONY: build
build:
	go build -o dist/elvis

.PHONY: test
test:
	go test -cover ./...