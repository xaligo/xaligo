---
applyTo: ".github/instructions/manual/**"
---

# 06.17 Roadmap: v0.9 PowerPoint Export

## v0.9 PowerPoint Export

Status: partially implemented ahead of schedule. Go-side geometry/routing plan
generation exists, and the Rust/WASI `pptx` adapter generates PPTX through the
development path. The required long-term gap is `xaligo.wasm`, invoked
from the Go repository layer with resolved plan JSON.
