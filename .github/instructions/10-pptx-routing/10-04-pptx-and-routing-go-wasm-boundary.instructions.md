---
applyTo: ".github/instructions/manual/**"
---

# 10.04 PPTX and routing: Go / Native Rust Boundary

## Go / Native Rust Boundary

The adopted integration style is Go invoking the statically linked native Rust
PPTX exporter from the repository layer through cgo and a narrow C ABI.

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

`cmd/wasm` is outside this boundary. It is a retained V1 browser compatibility
adapter, cannot execute the native V2 engine, and must not become a second PPTX
runtime or a required V2/release gate.

Other integration styles are not the current implementation target:

| Style | Status |
|---|---|
| stdin/stdout JSON-RPC | Candidate for long-running/high-volume workflows |
| HTTP API | Candidate for service/BFF separation |
| gRPC | Candidate for high-performance typed service boundaries |
| Node.js subprocess | Not a target for PPTX export |
| Embedded JS engine (`goja`, V8) | Not a target for PPTX export |

Do not spend implementation time replacing the repository-layer exporter with
`goja` or V8 unless that architecture is explicitly re-approved.
