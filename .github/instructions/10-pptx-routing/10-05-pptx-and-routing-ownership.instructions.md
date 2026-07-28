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
| WASM-compatible PPTX drawing/export | `external` TypeScript package and implementation |
| PPTX WASI command entry | `external/command.ts` |
| Public browser/JavaScript API bridge | `cmd/wasm/main.go` |
