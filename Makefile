PROJECT_DIR := $(shell pwd)
PROJECT_BIN := $(PROJECT_DIR)/bin

$(shell [-f bin] || mkdir -p $(PROJECT_BIN))

.PHONY:build
build:
	go build -o $(PROJECT_BIN)/newsbot ./cmd

.DEFAULT_GOAL := build