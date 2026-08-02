---
applyTo: ".github/instructions/manual/**"
---

# 08.03 Architecture: Package responsibilities

## Package responsibilities

| Path | Responsibility |
|---|---|
| `internal/entity` | Independent entity layer containing cross-layer structures |
| `internal/usecase` | Complete render, diagnostics, diff, project-intelligence, parser, layout, pagination, plan, and scene components |
| `internal/usecase/v1/engine` | Synchronous V1 parser, validation, layout, scene, routing, and plan calculations; no repository or scheduling ownership |
| `internal/usecase/v2` | V2 Go orchestration, cancellation, typed ABI encoding, and Rust engine invocation |
| `internal/repository` | Filesystem, catalog, HTTP preview, output-format adapters, and SQLite icon/project stores |
| `internal/lsp` | LSP 3.18 stdio session state and protocol adaptation over shared use cases |
| `internal/mcp` | Stateless MCP protocol handling and tool adaptation over shared use cases |
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | JavaScript-global adapter over shared use cases and embedded assets |
| `external/engine` | Rust generic layout/SVG workspace, C header, and Go/cgo adapter for the versioned in-process ABI |
| `external/pptx-exporter` | TypeScript PPTX adapter: `command.ts`, `controller`, `entity`, `repository`, and `usecase` |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |
