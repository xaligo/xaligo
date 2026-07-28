---
applyTo: ".github/instructions/manual/**"
---

# 09.03 Coding: Engine execution boundary

## Engine execution boundary

- `internal/usecase/v1/engine` contains synchronous calculation stages and
  explicitly supplied synchronous dependency ports.
- It must not import concrete repositories, interpret `context.Context`, start
  goroutines, own channels or worker pools, or select concurrency limits.
- The parent `internal/usecase` package owns repository adaptation, I/O,
  cancellation checks, stage ordering, job partitioning, result ordering, and
  future parallel-process control.
- Order-dependent routing within one document or plan remains sequential.
