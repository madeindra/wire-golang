@PHONY: setup wire build

setup:
		@echo "Setup the project"
		@go install github.com/google/wire/cmd/wire@latest
		@go mod download
		@echo "Setup complete"

wire:
		@echo "Wire up the project"
		@wire ./...
		@echo "Wiring complete"

build:
		@echo "Build the project"
		@go build -o ./bin/user ./cmd/user
		@echo "Build complete"