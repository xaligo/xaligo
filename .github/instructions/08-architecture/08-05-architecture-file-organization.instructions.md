---
applyTo: ".github/instructions/manual/**"
---

# 08.05 Architecture: File organization

## File organization

Files are divided by cohesive implementation responsibility, not by declaration
kind. A responsibility file may contain its private types, interface, concrete
implementation, constructor/factory, constants, and methods when those
declarations exist to support that implementation.

The package directory already identifies the architectural layer, so Go file
names never repeat `controller`, `usecase`, or `repository`. Use the component
responsibility as the filename prefix:

- `<component>.go` contains the component's public interface, constructor, and
  principal concrete implementation.
- `<component>_<detail>.go` contains a cohesive private implementation slice of
  the same component.
- The public interface is `<Component>Controller`, `<Component>Usecase`, or
  `<Component>Repository`, according to its package.
- The constructor is `New<Component>Controller`, `New<Component>Usecase`, or
  `New<Component>Repository` and returns that interface.
- The concrete implementation type is unexported.

Current component prefixes are `diff`, `generate`, `icon`, `init`, `lsp`,
`mcp`, `rag`, `render`, `serve`, `validate`, and `version` in
`internal/controller`; `render`, `diff`, `diagnostics`, `project`, `parser`,
`layout`, `pagination`, `plan`, and `scene` in `internal/usecase`; and
`powerpoint`, `preview`, `svg`, `scene`, and `xaligo` in
`internal/repository`. Repository supporting files retain the same prefix,
such as `powerpoint_export.go`. Every direct
`internal/usecase/*.go` file is a complete component as specified in
`../09-coding/09-00-coding-overview.instructions.md`.

Calculation files in `internal/usecase/v1/engine` use functional prefixes
such as `parse_*`, `layout_*`, `scene_*`, `route_*`, and `plan_*`. They contain
cohesive algorithm slices and do not repeat the package or architectural layer
name in filenames.

V2 orchestration components live in `internal/usecase/v2`. The Go/cgo adapter,
C header, Rust crates, and generated ignored link directory stay together in
`external/engine`; do not recreate an `internal/engineffi` package.

- Keep a Go interface in the file containing the corresponding concrete
  implementation and its principal methods.
- Keep a Go constructor in the file containing the concrete type it returns or
  initializes.
- Keep a TypeScript interface and factory with the implementation that consumes
  or realizes that contract when the interface is implementation-specific.
- When several implementations satisfy one interface, keep the interface with
  the package's primary responsibility/implementation and keep each additional
  implementation with its own methods; do not introduce a declaration-only
  file merely to appear neutral.
- Cross-layer entity DTOs and renderer-neutral value contracts remain in
  `internal/entity` or `external/pptx-exporter/entity`; this rule does not move
  shared data models into implementation packages.
- File splitting must move complete responsibility slices. Do not split an
  interface, its constructor, and its concrete behavior into separate files.
- Place the interface, unexported concrete type, and constructor at the start of
  their responsibility file, after imports and any shared constants/log codes,
  and before implementation methods. Do not recreate a package-wide facade that
  lists methods implemented by unrelated responsibility components.
- Private files and functions may divide one concrete component's implementation
  (for example PPTX image, legend, and package helpers). Such helpers must not
  expose a peer-layer interface or constructor and do not constitute another
  repository/use-case/controller component.
