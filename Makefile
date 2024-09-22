APP_NAME := entra
SRC_DIR := ./src
BUILD_DIR := ./build

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Build and cp binary to ~/.local/bin
release: build
	@echo "Rease $(APP_NAME)..."
	cp ./build/$(APP_NAME) ~/.local/bin/$(APP_NAME)

.PHONY: all build release clean