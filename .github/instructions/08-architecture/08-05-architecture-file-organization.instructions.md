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
`rag`, `render`, `serve`, `validate`, and `version` in
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
C header, single Rust staticlib crate, and generated ignored link directory
stay together in `external/engine`; do not recreate an `internal/engineffi`
package. Rust source follows the same responsibility layering as
`ryo-arima/vem/src`: `cnf` owns constants and limits, `ent` owns model,
request, and response data, `usc` owns operation orchestration and its
cohesive flow, geometry, routing, validation, and SVG calculation slices,
`ctl` owns ingress/egress through the C ABI, `rep` is intentionally empty until
Rust owns an external representation such as PPTX package generation, and `util` owns shared
technical helpers, and `base.rs` is the composition root. `lib.rs` exports only
the static-library boundary. Do not restore separate layout, SVG, or FFI crates
or place all responsibilities back into a monolithic `lib.rs` or `usc/engine.rs`.

Keep `ctl`, `usc`, and `rep` shallow. Express cohesive slices with filenames
such as `layout_flow.rs`, `layout_routing.rs`, and a future `pptx_package.rs`;
do not create responsibility subdirectories below those layer directories.

The Rust engine follows VEM's explicit implementation convention. Keep entity
declarations free of `derive`-generated implementations. Implement `Clone`,
`Debug`, `Default`, and equality in the matching `util` responsibility files,
and implement ABI serialization/deserialization explicitly in
`util/serialize.rs` and `util/deserialize.rs`. These codec files operate on the
bounded little-endian ABI; they must not introduce serde, JSON, or arbitrary
maps at the Go/Rust boundary. Required ABI/compiler/tooling attributes,
including `repr(C)`, exported-symbol, lint, and targeted rustfmt attributes,
are not model implementations and remain permitted. Test attributes and test
module wiring remain exclusively under `test/`, never in product source.

Write every grouped Rust import vertically, with one imported item per line:

```rust
use module::{
    First,
    Second,
};
```

Single-item imports may use the ungrouped form. A targeted `rustfmt::skip` may
be placed on a short grouped import only when stable rustfmt would otherwise
collapse this required layout.

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
  `internal/entity` or `external/exporter/src/ent`; this rule does not move
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
