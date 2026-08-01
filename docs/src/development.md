# Development

The canonical implementation pipeline is:

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
| `external/pptx-exporter` | TypeScript/WASM package and PPTX adapter |

Verification commands:

```bash
make test-engine
go build ./...
git diff --check
```

PPTX exporter builds:

```bash
make build-wasm
npm run build --workspace=@xaligo/xaligo-external
```

Generated binaries, `node_modules`, `output`, mdBook build output, WASM
artifacts, and TypeScript `dist` output should not be committed.

GitHub Pages is published by `.github/workflows/pages.yml` when documentation
changes are merged to `main`. The workflow builds `docs/` with mdBook and
deploys `docs/book` through GitHub Actions Pages.
