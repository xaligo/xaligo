---
applyTo: ".github/instructions/manual/**"
---

# 01.03 General: Architecture rules

## Architecture rules

- Preserve `.xal -> parser -> layout -> shared scene/plan -> encoder`.
- Format-rendering adapters call `internal/usecase`; they do not create
  parallel parser or layout pipelines.
- Input/output-format-specific encoding and persistence belong to
  `internal/repository`; use-case filenames describe processing, not formats.
- `internal/entity` owns structures exchanged between layers and contains no
  application orchestration. Shared value helpers such as theme names and
  service labels may live here when they are renderer-independent.
- Calculation and orchestration belong under `internal/usecase`. Synchronous
  calculations live in `internal/usecase/v1/engine`; repository
  I/O, cancellation, stage ordering, and concurrency control remain in the
  parent `internal/usecase` package.
- `v1/engine` must not import repositories, start goroutines, own worker
  pools, or interpret contexts. Parallel execution is a caller-owned policy
  over independent engine jobs.
- Keep mode (visual semantics) independent from format (serialization).
- Keep cross-format routing and geometry in shared layers.
- `cmd` imports command/adapters only; business logic stays outside entry points.
- Native and embedded environments differ through `usecase.AssetSource`, not
  through duplicated render implementations.
- Go constructs PPTX draw plans; the statically linked Rust `pptx` adapter writes
  PPTX bytes. Do not add a second OOXML writer.
- Return context-wrapped errors. Do not panic in core code.
