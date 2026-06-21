---
applyTo: "**/*.{go,ts,md,yml,yaml,json}"
---

# Architecture

This document defines the implementation boundaries of xaligo. Product
direction lives in `roadmap.instructions.md`; DSL behavior lives in
`xal-spec.instructions.md`.

## Core pipeline

```text
.xal source
   -> internal/usecase parser functions
   -> internal/usecase layout functions
    -> resolved Excalidraw scene
    -> shared draw plan / integration encoder
    -> SVG | Excalidraw | PPTX | XYFlow | Isoflow
```

The flat `internal/usecase` package is the shared rendering boundary.
Format-rendering adapters (CLI, preview server, and WASM) call this use case
instead of assembling a parallel parser/layout/render pipeline. Utility
commands such as `generate xal` and `add service` may use their focused internal
builders and repositories directly.

## Package responsibilities

| Path | Responsibility |
|---|---|
| `internal/entity` | Independent entity layer containing cross-layer structures |
| `internal/usecase` | Parser, layout, rendering, validation, preview, and shared plan calculations; organized by filenames and constructor-injected API |
| `internal/repository` | Native filesystem/catalog/PPTX adapter operations |
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | JavaScript-global adapter over shared use cases and embedded assets |
| `external` | TypeScript external adapter layer mirroring `internal`: `command.ts`, `controller`, `entity`, `repository`, `usecase` |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |

## Invariants

1. `.xal` is the only source DSL. Do not add adapter-specific parsers.
2. Mode selects visual semantics; format selects output serialization.
3. Format-rendering production paths call parser and layout through
   `internal/usecase`. Adapters use injected `usecase.API`, `Render`, `BuildPPTXPlan`,
   `Validate`, or `Diagnose`.
4. Routing and connector behavior belongs in shared scene/plan layers, not in
   individual output adapters.
5. Filesystem-less environments provide an `AssetSource`; they do not fork the
   render pipeline.
6. Native configuration remains the default when `RenderOptions.Assets` is nil.
7. New formats require a `Format` value, shared render function, CLI wiring,
   tests, and adapter documentation.
8. Errors are returned and wrapped with context. Core packages do not panic.

## Dependency direction

```text
cmd / internal/controller / cmd/wasm / TypeScript
                         |
                         v
                     internal/usecase
                         |
                         v
                    internal/*
```

Entity and use-case packages must not depend on CLI, preview, WASM, or
TypeScript adapters. Encoders consume entity structures and must not depend on
use-case implementations merely to access types.

## Verification

Run after structural changes:

```bash
go test ./...
go build ./...
GOOS=js GOARCH=wasm go build -o /tmp/xaligo.wasm ./cmd/wasm
npm install
npm run build --workspace=@ryo-arima/xaligo
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.
