---
applyTo: ".github/instructions/manual/**"
---

# 08.08 Architecture: Dependency direction

## Dependency direction

```text
internal command / controller / cmd/wasm
                  |
                  v
   internal/usecase orchestration
          /                 \
         v                   v
internal/usecase/       internal/repository
  v1/engine               interfaces and
     |                   implementations
     v                         |
 internal/entity <-------------+

external/pptx-exporter/command.ts
        |
        v
external/pptx-exporter/controller
        |
        v
external/pptx-exporter/usecase
        |
        v
external/pptx-exporter/repository
```

Entity and use-case packages must not depend on CLI, preview, WASM, or
TypeScript adapters. Encoders consume entity structures and must not depend on
use-case implementations merely to access types.
