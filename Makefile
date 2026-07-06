.PHONY: help build build-wasm test fmt tidy run init clean

BIN_DIR  := .bin
BINARY   := $(BIN_DIR)/xaligo
WASM_OUT      := external/wasm
VERSION  := $(shell sed -n '1{s/^v//;p;q;}' VERSION)
LDFLAGS  := -X github.com/xaligo/xaligo/internal/controller.version=$(VERSION)

help: ## Show commands
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: build-wasm ## Build CLI binary and PPTX exporter WASM bundle
	@mkdir -p $(BIN_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd
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
	@mkdir -p output
	$(BINARY) render docs/src/examples/samples/sample.xal -o output/sample.excalidraw
	@echo "Generated: output/sample.excalidraw"

init: build ## Create starter template under output/example/
	$(BINARY) init -o output/example

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR)
	rm -f $(WASM_OUT)/xaligo.wasm
