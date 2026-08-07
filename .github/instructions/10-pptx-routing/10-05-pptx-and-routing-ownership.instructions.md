---
applyTo: ".github/instructions/manual/**"
---

# 10.05 PPTX and routing: Ownership

## Ownership

| Area | Owner |
|---|---|
| DSL parse/layout | `internal/usecase/v1/engine/parse_*`, `internal/usecase/v1/engine/layout_*` |
| Canonical scene and item metadata | `internal/usecase/v1/engine/scene_*` |
| Plan geometry, text layout, paper scaling, routing, legend data | `internal/usecase/v1/engine/plan_*`, `internal/usecase/v1/engine/route_*` |
| WASM exporter invocation from Go | `internal/repository/powerpoint.go` |
| WASM-compatible PPTX drawing/export | `external/exporter` Rust package and `pptx` adapter |
| PPTX WASI command entry | `external/exporter/src/ctl/command.rs` |
| Public browser/JavaScript API bridge | `cmd/wasm/main.go` |
