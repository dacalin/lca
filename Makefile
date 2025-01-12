# Makefile for lca

.PHONY: build install clean

## Build the lca binary
build:
	go build -o lca ./src

## Install the lca binary to /usr/local/bin
install: build
	sudo cp lca /usr/local/bin/

## Clean build artifacts
clean:
	rm -f lca
