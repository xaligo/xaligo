---
applyTo: ".github/instructions/manual/**"
---

# 08.01 Architecture: Core pipeline

## Core pipeline

```text
.xal source
   -> internal/usecase orchestration
   -> internal/usecase/v1/engine parser functions
   -> validated numeric and enum attributes
   -> internal/usecase/v1/engine layout functions
   -> resolved internal compatibility scene (V1 only)
	-> internal/usecase/v1/engine plan calculations
	-> internal/repository output encoder
	-> SVG | PPTX
```

V2 terminal output follows the native V2 branch:

```text
.xal V2 source -> V2 frontend -> Rust resolution -> immutable resolved document
   -> terminal repository -> Unicode | ASCII | ANSI text
```

The parent `internal/usecase` package is the shared rendering and orchestration
boundary. Its `v1/engine` subpackage contains synchronous V1 calculation stages.
Format-rendering adapters (CLI and preview server) call a constructor-injected
`RenderUsecase` instead of assembling a parallel parser/layout/render pipeline.
The retained legacy V1 WASM adapter follows the same rule but is not a V2
runtime. Markdown rendering calls the SVG artifact path and embeds references;
it is not a separate encoder. Utility commands such as `generate xal` may use
focused internal builders directly.
Terminal rendering is a V2-only capability projection and never enters the V1
compatibility scene or plan builder.

The V1 compatibility scene is an implementation detail while the shared plan
is migrated to the V2 neutral model. It must not be selectable as an output,
persisted by a public command, or exposed by CLI or legacy WASM format dispatch.
