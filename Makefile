GO ?= go

MAIN = "cmd/main.go"
CLI_DIR = "./bin/poke-store"

dependencies:
	$(GO) mod download
	$(GO) mod tidy

cli:
	CGO_ENABLED=0 GOOS=linux $(GO) build -o $(CLI_DIR) $(MAIN)

run:
	$(GO) run $(MAIN)