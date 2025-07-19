APP_NAME := cchat
VERSION := v0.1.0
DIST_DIR := dist
README := README.md
LICENSE := LICENSE

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64

.PHONY: all build clean

all: clean build

build:
	@mkdir -p $(DIST_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$${platform%/*}; \
		ARCH=$${platform#*/}; \
		BINARY_NAME="$(APP_NAME)-$${OS}-$${ARCH}"; \
		echo "Building $$BINARY_NAME..."; \
		GOOS=$${OS} GOARCH=$${ARCH} go build -o $(DIST_DIR)/$$BINARY_NAME main.go; \
		echo "Packaging $$BINARY_NAME.tar.gz..."; \
		tar -czf $(DIST_DIR)/$$BINARY_NAME.tar.gz -C $(DIST_DIR) $$BINARY_NAME $(README) $(LICENSE); \
		rm $(DIST_DIR)/$$BINARY_NAME; \
	done
	@echo "✅ Build and packaging done!"

clean:
	@rm -rf $(DIST_DIR)
