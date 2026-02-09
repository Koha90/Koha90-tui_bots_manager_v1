build:
	@go build -o bin/tui_bot_manager cmd/main.go

run: build
	@./bin/tui_bot_manager

test:
	@go test ./...
