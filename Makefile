NAME := venstar-cli
.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

check: vet lint
.PHONY:check

build:
	go build -o build/$(NAME)
.PHONY: build
