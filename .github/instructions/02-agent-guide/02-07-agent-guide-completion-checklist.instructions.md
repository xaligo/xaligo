---
applyTo: ".github/instructions/manual/**"
---

# 02.07 Agent guide: Completion checklist

## Completion checklist

1. Format changed Go files with `gofmt`.
2. Run `go test ./...` and `go build ./...`.
3. For shared render use-case or asset changes, cross-build `cmd/wasm`.
4. For PPTX exporter changes, run its Cargo tests and `make build-wasm`.
5. Run `git diff --check` and inspect `git status --short`.
6. Update the DSL spec, architecture, README, or roadmap when their contract
   changed.

Unit tests belong in `test/unit`, mirroring the source tree they cover.
Black-box API and adapter tests belong in `test/integration`. Prefer testing
observable behavior over exposing package-private helpers only for tests.
