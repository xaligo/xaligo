---
applyTo: ".github/instructions/manual/**"
---

# 06.17 Roadmap: v0.9 PowerPoint Export

## v0.9 PowerPoint Export

Status: implemented. Go owns geometry/routing plan generation, and the native
Rust `pptx` adapter generates PPTX through an in-process C ABI from resolved
plan JSON.
