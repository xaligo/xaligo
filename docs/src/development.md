# Development

The current V1 compatibility pipeline is:

```text
.xal source
  -> internal/usecase parser
  -> internal/usecase layout
  -> shared scene and plan calculations
  -> repository encoder
```

Important package boundaries:

| Path | Responsibility |
|---|---|
| `internal/usecase` | Parser, layout, validation, scene, routing, and plans |
| `internal/usecase/v2` | V2 orchestration, cancellation, and typed Rust-engine invocation |
| `internal/entity` | Shared structures exchanged between layers |
| `internal/repository` | Filesystem and output-format adapters |
| `cmd` | CLI entry points |
| `external/engine` | Go/cgo adapter, C ABI, and Rust layout/SVG engine workspace |
| `external/exporter` | Rust PPTX adapter exposed through a C ABI |

The implemented V2 calculation boundary is:

```text
typed EngineDocumentSpec
  -> internal/usecase/v2 (cancellation and ABI adaptation)
  -> external/engine (Rust layout, routing, and SVG)
  -> immutable EngineResolvedDocument
```

ABI v2 flattens the typed tree in pre-order and retains hierarchy with parent
indexes. The future native V2 and V1 compatibility frontends should lower one
parsed `.xal` concept tree directly into this request without JSON, source
rewrites, or a renderer-shaped intermediate scene. See [V2 Generic
Engine](design/v2-engine.md).

Verification commands:

```bash
make test-engine
cargo clippy --manifest-path external/engine/Cargo.toml --workspace --all-targets --locked -- -D warnings
go build ./...
git diff --check
```

PPTX exporter builds:

```bash
make build-exporter
make build
```

Generated binaries, `node_modules`, `output`, mdBook build output, WASM
artifacts, and TypeScript `dist` output should not be committed.

GitHub Pages is published by `.github/workflows/pages.yml` when documentation
changes are merged to `main`. The workflow builds `docs/` with mdBook and
deploys `docs/book` through GitHub Actions Pages.
