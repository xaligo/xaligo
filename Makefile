.PHONY: help build build-engine build-wasm test test-engine security-setup security-check fmt tidy run init clean

BIN_DIR  := .bin
BINARY   := $(BIN_DIR)/xaligo
TOOLS_BIN_DIR := $(BIN_DIR)/tools
GOVULNCHECK   := $(TOOLS_BIN_DIR)/govulncheck
GOVULNCHECK_VERSION := v1.6.0
PPTX_EXPORTER_DIR := external/pptx-exporter
WASM_OUT      := $(PPTX_EXPORTER_DIR)/wasm
ENGINE_DIR    := external/engine
ENGINE_PACKAGE := xaligo-engine-ffi
ENGINE_STATICLIB ?= $(ENGINE_DIR)/target/release/libxaligo_engine.a
ENGINE_LINK_DIR := $(ENGINE_DIR)/lib
ENGINE_LINK_LIB := $(ENGINE_LINK_DIR)/libxaligo_engine.a
ENGINE_BUILD_TAG := xaligo_engine
VERSION  := $(shell sed -n '1{s/^v//;p;q;}' VERSION)
LDFLAGS  := -X github.com/xaligo/xaligo/internal/controller.version=$(VERSION)

help: ## Show commands
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: build-engine build-wasm ## Build the single CLI binary with Rust engine and PPTX exporter
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 go build -tags $(ENGINE_BUILD_TAG) -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd
	@echo "Built: $(BINARY)"

build-engine: ## Build and stage the Rust static library for cgo
	cargo build --manifest-path $(ENGINE_DIR)/Cargo.toml --package $(ENGINE_PACKAGE) --release
	@mkdir -p $(ENGINE_LINK_DIR)
	install -m 0644 $(ENGINE_STATICLIB) $(ENGINE_LINK_LIB)
	@echo "Built: $(ENGINE_LINK_LIB)"

build-wasm: ## Build TS/WASI PPTX exporter into external/pptx-exporter/wasm/
	@mkdir -p $(WASM_OUT)
	npm --prefix $(PPTX_EXPORTER_DIR) run build:pptx-exporter-wasm
	@echo "Built: $(WASM_OUT)/xaligo.wasm"

test: test-engine ## Run tests
	go test ./...

test-engine: build-engine ## Test Rust crates and the linked cgo engine path
	cargo test --manifest-path $(ENGINE_DIR)/Cargo.toml --workspace
	CGO_ENABLED=1 go test -tags $(ENGINE_BUILD_TAG) ./... -count=1

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
	cargo fmt --manifest-path $(ENGINE_DIR)/Cargo.toml --all

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
	rm -rf $(ENGINE_DIR)/target
	rm -rf $(ENGINE_LINK_DIR)
