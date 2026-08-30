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
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | Source-only legacy V1 JavaScript-global compatibility adapter over shared use cases and embedded assets; no native V2 execution |
| `external/engine` | One layered Rust staticlib crate (`cnf`, `ent`, `usc`, `ctl`, reserved `rep`, and `util`), C header, and Go/cgo adapter for the versioned in-process ABI |
| `external/exporter` | Rust PPTX adapter with the same `cnf`, `ent`, `usc`, `rep`, `ctl`, and `util` layering as `external/engine`; linked into its static library |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |
