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
   -> resolved canonical scene
   -> internal/usecase/v1/engine plan calculations
   -> internal/repository output encoder
   -> SVG | Excalidraw | PPTX | PDF | Excel | XYFlow | Isoflow
```

The parent `internal/usecase` package is the shared rendering and orchestration
boundary. Its `v1/engine` subpackage contains synchronous V1 calculation stages.
Format-rendering adapters (CLI, preview server, and WASM) call a
constructor-injected `RenderUsecase` instead of assembling a parallel
parser/layout/render pipeline. Utility
commands such as `generate xal` and `add service` may use their focused internal
builders and repositories directly.
