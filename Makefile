# Makefile for lca

.PHONY: build install clean

BINARIES_DIR=usr/local/bin
INSTALL_DIR=/usr/local/bin

## Build the lca binary
build:
	cd src && GOARCH=amd64 go build -o ../linux/amd64/$(BINARIES_DIR)/lca .
	cd src && GOOS=linux GOARCH=arm64 go build -o ../linux/arm64/$(BINARIES_DIR)/lca .
	cd src && GOOS=darwin GOARCH=amd64 go build -o ../darwin/amd64/$(BINARIES_DIR)/lca .
	cd src && GOOS=darwin GOARCH=arm64 go build -o ../darwin/arm64/$(BINARIES_DIR)/lca .

## Install the lca binary to /usr/local/bin
install: build
	@CURRENT_OS=$$(go env GOOS); \
	CURRENT_ARCH=$$(go env GOARCH); \
	BINARY_PATH="$$CURRENT_OS/$$CURRENT_ARCH/$(BINARIES_DIR)/lca"; \
	if [ -f "$$BINARY_PATH" ]; then \
		echo "Installing binary: $$BINARY_PATH"; \
		sudo cp "$$BINARY_PATH" $(INSTALL_DIR)/lca; \
		sudo chmod +x $(INSTALL_DIR)/lca; \
		echo "Installation complete. lca is now available."; \
	else \
		echo "Error: Binary for $$CURRENT_OS/$$CURRENT_ARCH not found at $$BINARY_PATH."; \
		exit 1; \
	fi

## Clean build artifacts
clean:
	rm -rf linux/amd64/$(BINARIES_DIR) linux/arm64/$(BINARIES_DIR) darwin/arm64/$(BINARIES_DIR) darwin/arm64/$(BINARIES_DIR)