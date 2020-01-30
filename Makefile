.PHONY: default
default: build test run

.PHONY: build
build:
	@go build -o bin/git2go-test

.PHONY: test
test:
	@go test -v -cover

.PHONY: _run
_run:
	@go run main.go

.PHONY: run
run: build _run
