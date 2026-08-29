---
applyTo: ".github/instructions/manual/**"
---

# 10.04 PPTX and routing: Go / WASM Boundary

## Go / WASM Boundary

The adopted integration style is Go invoking a WASM-compiled PPTX exporter from
the repository layer.

Implementation preconditions:

- Go owns CLI/controller/repository orchestration.
- The native exporter must be called from `internal/repository/powerpoint.go`, not directly from
  controller or command packages.
- The exporter must be compiled into the Rust engine static library.
- Go forwards user-facing PPTX options to the native exporter through a typed
  options structure or JSON bridge.
- The native exporter consumes the resolved shared Go plan and returns PPTX
  bytes or writes them through a repository-controlled output path.
- The native exporter must not perform independent geometry, layout, or routing.
- The exporter follows `ctl -> usc -> rep`, matching the Rust
  engine's layer structure. Command/controller code must not bypass this path.
- Go repository/controller code must not implement PPTX/OOXML drawing or zip
  writing directly. Keep Go as the adapter that builds the plan, invokes the
  native exporter, and persists the returned bytes.
- Use the MIT-licensed pure-Rust `pptx` crate as the package
  writer; do not introduce `goja`, V8, or a JavaScript drawing layer.

Other integration styles are not the current implementation target:

| Style | Status |
|---|---|
| stdin/stdout JSON-RPC | Candidate for long-running/high-volume workflows |
| HTTP API | Candidate for service/BFF separation |
| gRPC | Candidate for high-performance typed service boundaries |
| Node.js subprocess | Temporary fallback only; not the target architecture |
| Embedded JS engine (`goja`, V8) | Not a target for PPTX export |

Do not spend implementation time replacing the repository-layer exporter with
`goja` or V8 unless that architecture is explicitly re-approved.
