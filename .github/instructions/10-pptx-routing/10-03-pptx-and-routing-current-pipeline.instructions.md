---
applyTo: ".github/instructions/manual/**"
---

# 10.03 PPTX and routing: Current Pipeline

## Current Pipeline

```text
.xal DSL
  -> Go parse and numeric-domain validation (typed normalization is the target)
  -> resolved layout and canonical scene
  -> ordered page-oriented Go document plan (neutral-schema migration remains)
  -> internal repository encoder (SVG), or
  -> Go repository -> WASM command -> external controller -> use case -> repository
  -> SVG | .pptx
```

Geometry belongs on the Go side. The WASM export module should only translate
the resolved plan into PPTX bytes. The current Excalidraw-shaped V1 scene is a
temporary internal compatibility serialization, not an output, target
architecture name, or ownership boundary for the shared plan.
