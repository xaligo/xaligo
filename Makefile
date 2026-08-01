.PHONY: help build build-wasm test security-setup security-check fmt tidy run init clean

BIN_DIR  := .bin
BINARY   := $(BIN_DIR)/xaligo
TOOLS_BIN_DIR := $(BIN_DIR)/tools
GOVULNCHECK   := $(TOOLS_BIN_DIR)/govulncheck
GOVULNCHECK_VERSION := v1.6.0
PPTX_EXPORTER_DIR := external/pptx-exporter
WASM_OUT      := $(PPTX_EXPORTER_DIR)/wasm
VERSION  := $(shell sed -n '1{s/^v//;p;q;}' VERSION)
LDFLAGS  := -X github.com/xaligo/xaligo/internal/controller.version=$(VERSION)

help: ## Show commands
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: build-wasm ## Build CLI binary and PPTX exporter WASM bundle
	@mkdir -p $(BIN_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd
	@echo "Built: $(BINARY)"

build-wasm: ## Build TS/WASI PPTX exporter into external/pptx-exporter/wasm/
	@mkdir -p $(WASM_OUT)
	npm --prefix $(PPTX_EXPORTER_DIR) run build:pptx-exporter-wasm
	@echo "Built: $(WASM_OUT)/xaligo.wasm"

test: ## Run tests
	go test ./...

security-setup: ## Install security scanners and prepare npm audit metadata
	@mkdir -p $(TOOLS_BIN_DIR)
	GOBIN=$(CURDIR)/$(TOOLS_BIN_DIR) go install golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)
	npm install --package-lock-only --ignore-scripts

security-check: ## Scan Go and npm dependencies for known vulnerabilities
	@test -x $(GOVULNCHECK) || { echo "Run 'make security-setup' first." >&2; exit 1; }
	$(GOVULNCHECK) ./...
	npm audit --audit-level=low

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
