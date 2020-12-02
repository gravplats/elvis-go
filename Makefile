build:
	go build -o dist/elvis

test:
	go test -cover ./...