---
applyTo: ".github/instructions/manual/**"
---

# 06.08 Roadmap: Rebaselined Implementation Order

## Rebaselined Implementation Order

Use this order when starting new roadmap work from the current repository state:

1. Complete the shared geometry/text correctness gate and its cross-renderer
   regression tests.
2. Move mixed item-grid occupancy into resolved layout and finish neutral
   scene/plan naming.
3. Complete the repository-layer WASM PPTX exporter contract by providing
   the native Rust exporter C ABI; keep Go free of PPTX/OOXML writer code.
4. Harden shared network routing with cross-renderer visual regression tests.
5. Build the VS Code preview on the reusable HTTP/SSE protocol exposed by
   `xaligo serve`.
