---
applyTo: ".github/instructions/manual/**"
---

# 02.07 Agent guide: Completion checklist

## Completion checklist

1. Format changed Go files with `gofmt`.
2. Run `go test ./...` and `go build ./...`.
3. Cross-build `cmd/wasm` only when changing that adapter or a shared V1
   browser path it consumes. Native V2 changes do not require this check.
4. For PPTX exporter changes, run its Cargo tests and `make build-exporter`.
5. Run `git diff --check` and inspect `git status --short`.
6. Update the DSL spec, architecture, README, or roadmap when their contract
   changed.

Unit tests belong in `test/unit`, mirroring the source tree they cover.
Black-box API and adapter tests belong in `test/integration`. Prefer testing
observable behavior over exposing package-private helpers only for tests.
