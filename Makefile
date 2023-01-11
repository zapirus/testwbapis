.PHONY: build
build:
	go build -v ./cmd/goapitest

.DEFAULT_GOAL := build