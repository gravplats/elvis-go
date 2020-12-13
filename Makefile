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

.PHONY: vet
vet:
	go vet ./...

.PHONY: build-docker
docker-build:
	docker container run --rm -v "${PWD}":/usr/src/elvis -w /usr/src/elvis -e GOOS=darwin -e GOARCH=amd64 golang:1.15 make build
