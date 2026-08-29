.PHONY: help build build-engine build-exporter generate-engine-abi test test-engine test-native security-setup security-check fmt tidy run init clean

BIN_DIR  := .bin
BINARY   := $(BIN_DIR)/xaligo
TOOLS_BIN_DIR := $(BIN_DIR)/tools
GOVULNCHECK   := $(TOOLS_BIN_DIR)/govulncheck
GOVULNCHECK_VERSION := v1.6.0
EXPORTER_DIR  := external/exporter
EXPORTER_PACKAGE := xaligo-pptx-exporter
EXPORTER_TEST_DIR := test/unit/external/exporter
ENGINE_DIR    := external/engine
ENGINE_PACKAGE := xaligo-engine-ffi
ENGINE_TEST_DIR := test/unit/external/engine
ENGINE_STATICLIB ?= $(ENGINE_DIR)/target/release/libxaligo_engine.a
ENGINE_LINK_DIR := $(ENGINE_DIR)/lib
ENGINE_LINK_LIB := $(ENGINE_LINK_DIR)/libxaligo_engine.a
NATIVE_BUILD_TAGS := xaligo_engine xaligo_exporter sqlite_fts5 sqlite_omit_load_extension
VERSION  := $(shell sed -n '1{s/^v//;p;q;}' VERSION)
LDFLAGS  := -X github.com/xaligo/xaligo/internal/controller.version=$(VERSION)

help: ## Show commands
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: build-engine ## Build the single CLI binary with Rust engine and PPTX exporter
	@mkdir -p $(BIN_DIR)
	CGO_ENABLED=1 go build -tags "$(NATIVE_BUILD_TAGS)" -ldflags "$(LDFLAGS)" -o $(BINARY) ./cmd
	@echo "Built: $(BINARY)"

build-engine: generate-engine-abi ## Build and stage the Rust static library for cgo
	cargo build --manifest-path $(ENGINE_DIR)/Cargo.toml --package $(ENGINE_PACKAGE) --release --locked
	@mkdir -p $(ENGINE_LINK_DIR)
	install -m 0644 $(ENGINE_STATICLIB) $(ENGINE_LINK_LIB)
	@echo "Built: $(ENGINE_LINK_LIB)"

generate-engine-abi: ## Generate Go and Rust ABI field constants from one schema
	go run scripts/tool/gen_engine_abi.go .

build-exporter: ## Build the Rust PPTX exporter crate
	cargo build --manifest-path $(EXPORTER_DIR)/Cargo.toml --package $(EXPORTER_PACKAGE) --release --locked
	@echo "Built: $(EXPORTER_DIR)/target/release/libxaligo_pptx_exporter.a"

test: test-native ## Run tests
	go test ./...

test-engine: test-native ## Test the linked Rust engine and native Go integration

test-native: build-engine build-exporter ## Test the Rust crates and linked cgo paths
	cargo test --manifest-path $(ENGINE_TEST_DIR)/Cargo.toml --locked
	cargo test --manifest-path $(EXPORTER_TEST_DIR)/Cargo.toml --locked
	CGO_ENABLED=1 go test -tags "$(NATIVE_BUILD_TAGS)" ./... -count=1

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
	cargo fmt --manifest-path $(EXPORTER_DIR)/Cargo.toml --all

tidy: ## Tidy go.mod
	go mod tidy

run: build ## Render sample DSL
	@mkdir -p output
	$(BINARY) render docs/src/examples/samples/sample.xal --format svg -o output/sample.svg
	@echo "Generated: output/sample.svg"

init: build ## Create starter template under output/example/
	$(BINARY) init -o output/example

clean: ## Remove build artifacts
	rm -rf $(BIN_DIR)
	rm -rf $(ENGINE_DIR)/target
	rm -rf $(EXPORTER_DIR)/target
	rm -rf $(ENGINE_LINK_DIR)
