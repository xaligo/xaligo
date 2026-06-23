.PHONY: help build build-wasm test fmt tidy run init clean

BIN_DIR  := .bin
BINARY   := $(BIN_DIR)/xaligo
WASM_OUT      := external/wasm

help: ## Show commands
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: build-wasm ## Build CLI binary and PPTX exporter WASM bundle
	@mkdir -p $(BIN_DIR)
	go build -o $(BINARY) ./cmd
	@echo "Built: $(BINARY)"

build-wasm: ## Build TS/WASI PPTX exporter into external/wasm/
	@mkdir -p $(WASM_OUT)
	npm --prefix external run build:pptx-exporter-wasm
	@echo "Built: $(WASM_OUT)/xaligo.wasm"

test: ## Run tests
	go test ./...

fmt: ## Format Go files
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

tidy: ## Tidy go.mod
	go mod tidy

run: build ## Render sample DSL
	$(BINARY) render examples/sample.xal -o examples/sample.excalidraw
	@echo "Generated: examples/sample.excalidraw"

init: build ## Create starter template under examples/
	$(BINARY) init -o examples

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR)
	rm -f $(WASM_OUT)/xaligo.wasm
