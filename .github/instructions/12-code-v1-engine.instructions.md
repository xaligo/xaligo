---
applyTo: "internal/usecase/v1/engine/**/*.go"
---

# V1 engine

The engine is synchronous: no repositories, `context.Context`, goroutines,
channels, worker pools, or concurrency policy. Parent use cases own I/O,
cancellation, ordering, partitioning, and parallelism.

Every package-scope identifier is
`<base>V1Engine<FileBaseCamelCase>`; locals, fields, imports, package names,
and `init` are exempt. Moving a declaration requires updating its suffix and
references. Exact naming examples: `reference.md` section `09`.
