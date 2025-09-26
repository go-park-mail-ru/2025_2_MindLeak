.PHONY: build test run
CMD_DIR=./cmd/server

build:
		go build -v $(CMD_DIR)

.PHONY: test
test:
		go test -v ./...

run:
		go run $(CMD_DIR)

.DEFAULT_GOAL := run