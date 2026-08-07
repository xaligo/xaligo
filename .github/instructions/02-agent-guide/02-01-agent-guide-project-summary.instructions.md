---
applyTo: ".github/instructions/manual/**"
---

# 02.01 Agent guide: Project summary

## Project summary

- Go 1.26 module: `github.com/xaligo/xaligo`
- CLI entry point: `cmd/main.go`
- PPTX exporter WASI command entry point: `external/exporter/src/ctl/command.rs`
- Rust package and implementation: `external/exporter`
- Shared application boundary: `internal/usecase`
- Generated CLI: `.bin/xaligo`
- Generated PPTX exporter WASM: `external/exporter/wasm/xaligo.wasm`
