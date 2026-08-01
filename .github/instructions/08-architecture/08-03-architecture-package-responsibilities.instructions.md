---
applyTo: ".github/instructions/manual/**"
---

# 08.03 Architecture: Package responsibilities

## Package responsibilities

| Path | Responsibility |
|---|---|
| `internal/entity` | Independent entity layer containing cross-layer structures |
| `internal/usecase` | Render orchestration, context checks, repository port adaptation, and future parallel scheduling |
| `internal/usecase/v1/engine` | Synchronous V1 parser, validation, layout, scene, routing, and plan calculations; no repository or scheduling ownership |
| `internal/repository` | Filesystem, catalog, HTTP preview, and output-format encoding/export adapters |
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | JavaScript-global adapter over shared use cases and embedded assets |
| `external/engine` | Rust generic layout/SVG workspace and versioned in-process ABI |
| `external/pptx-exporter` | TypeScript PPTX adapter: `command.ts`, `controller`, `entity`, `repository`, and `usecase` |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |
