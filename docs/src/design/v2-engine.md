# V2 Generic Engine

The V2 calculation engine is a Rust library linked into the
`xaligo` executable. It does not run as a daemon, sidecar, dynamic plugin, or
separate command:

```text
Go use case -> cgo -> C ABI v5 -> Rust staticlib
```

The public Go use-case boundary remains:

```go
type EngineUsecase interface {
    Resolve(context.Context, entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error)
    RenderSVG(context.Context, entity.EngineDocumentSpec) ([]byte, error)
    NormalizeSVG(context.Context, []byte) (entity.EngineSVG, error)
}
```

The generic calculation core and this ABI are implemented. The
`<xaligo version="2">` frontend lowers directly to `EngineDocumentSpec`, and
the ordinary render use case projects its resolved result through the shared
SVG/PPTX document plan. Its authoring profile preserves V1-style rows, columns,
AWS groups, catalog items, ports, and connections while keeping generic V2
parameters available as explicit extensions. Profile normalization is one
linear tree walk; it does not reparse XML or run the V1 renderer. Catalog icon
data and short labels share one load per render; the bundled default catalog is
cached across renders.

## Rust source structure

The engine is one staticlib crate and follows the responsibility structure of
[`ryo-arima/vem/src`](https://github.com/ryo-arima/vem/tree/main/src):

```text
external/engine/src/
├── lib.rs                  staticlib export boundary
├── mod.rs                  module registry
├── base.rs                 decode/execute/encode composition root
├── cnf/engine.rs           ABI constants, limits, and defaults
├── cnf/engine_abi.rs       generated ABI field constants
├── ent/model/              generic document and normalized SVG models
├── ent/model/aws/          typed ALB, NLB, and listener definitions
├── ent/request/engine.rs   binary request decoding
├── ent/response/engine.rs  binary response encoding
├── usc/engine.rs           operation dispatch
├── usc/aws/                per-service validation and visual composition
├── usc/cancel.rs           call-scoped cooperative cancellation
├── usc/layout.rs           layout resolution entry point and shared state
├── usc/layout_flow.rs      stack, grid, and absolute placement
├── usc/layout_geometry.rs  bounds and geometry helpers
├── usc/layout_routing.rs   ports and line routing
├── usc/layout_validation.rs input and constraint validation
├── usc/layout_tests.rs     layout regression tests
├── usc/svg.rs              SVG normalization and projection
├── rep.rs                  reserved for future encoders such as PPTX
├── ctl/engine.rs           panic-safe C ABI and owned buffers
└── util/
    ├── serialize.rs        explicit ABI response serialization
    ├── deserialize.rs      explicit bounded ABI request deserialization
    ├── clone/debug/eq.rs   explicit model trait implementations
    ├── default.rs          explicit neutral-value defaults
    ├── mcode.rs            engine message codes and log levels
    ├── logger.rs           Go-compatible environment-configured logging
    └── error.rs            layout and SVG errors
```

`external/engine/abi/fields.csv` is the single field-index schema. Running
`make generate-engine-abi` regenerates both Go and Rust constants without
introducing runtime reflection, serde, JSON, or arbitrary maps.

The module hierarchy is declared only in
[`external/engine/src/mod.rs`](../../../external/engine/src/mod.rs), including
the inline AWS namespaces. There are no nested AWS `mod.rs` files. The component
enum lives in
[`ent/model/aws/component.rs`](../../../external/engine/src/ent/model/aws/component.rs),
and the composition entry point lives in
[`usc/aws/composition.rs`](../../../external/engine/src/usc/aws/composition.rs).
Re-exports preserve the existing `ent::model::aws::Component` and
`usc::aws::compose` paths; this organization does not change rendering or ABI.

This preserves the `cnf / ent / usc / ctl / util` dependency vocabulary while
placing calculation behavior in cohesive use-case files and without copying
VEM's CLI-specific `main.rs`. `lib.rs` is the corresponding library entry
point. The C symbols and Go `EngineUsecase` contract are unchanged by this
source-only reorganization.

The current data path is `ctl -> usc -> ctl`. The `rep` layer deliberately has
no implementation because calculation results return directly through `ctl`.
If PPTX package generation later moves into Rust, its external-representation
writer belongs in a flat `rep/pptx_*.rs` slice and may be called from `usc`;
layout, routing, and validation remain in `usc`. Generic algorithm files stay
shallow; service components deliberately use `usc/aws/<component>.rs`.

Engine-owned models intentionally have no `derive` or serde annotations.
Standard traits and the binary ABI codecs are implemented explicitly under
`util`, following VEM's implementation pattern. `serialize.rs` and
`deserialize.rs` continue to read and write the fixed-width little-endian ABI;
they do not add a JSON or generic-map boundary.

Grouped imports use a vertical, one-item-per-line form throughout the Rust
crate. This makes imports stable and reviewable as layer dependencies change.

## Rust logging

`util/mcode.rs` defines structured engine message codes and
`util/logger.rs` follows the Go shared logger contract: level filtering,
optional structured JSON, component/service metadata, caller metadata, error
field extraction, and `XALIGO_LOG_*` configuration. The static library adapts
that contract in two ways: stderr is the default so Go protocol stdout remains
machine-usable, and a Rust fatal log never terminates the embedding Go process.
Engine failures are still returned through the typed ABI.

Default-level execution is silent because the composition root emits only
debug lifecycle events. Logs do not include `.xal` contents or absolute caller
paths.

## Generic concepts

The engine accepts only the following calculation concepts:

| Concept | Calculation responsibility |
|---|---|
| `Frame` | Canvas, page, and top-level containment geometry |
| `Group` | Nested content box and child layout |
| `Capture` | Generic emphasis or annotation boundary |
| `Item` | Atomic or composed visual slot |
| `Port` | Addressable endpoint placed inside or as an icon intersecting one selected owner border |
| `Line` | Straight or orthogonal route between IDs |
| `Text` | Renderer-neutral intrinsic text measurement and label data |
| `Spacer` | Non-drawing layout participant |

These concepts remain the shared layout/routing vocabulary. Native AWS
components pass a closed, typed component model across the ABI and expand once
into these generic parts before layout. Source tags and profile IDs do not
drive generic algorithms or either encoder.

## Native AWS components

ALB/NLB service definitions live in `ent/model/aws/{alb,nlb,listener}.rs` and
their validation/design in `usc/aws/`. The latter controls domain-tag text
measurement, listener-card placement and security badges. It clones and expands
the source document without changing authored IDs. Both SVG and PPTX consume
the same result. There are no watermarks or renderer-specific drawing strings.

ALB feature models add rules, conditions/matches, actions/weighted forwards/JWT
claims, transforms/rewrites, target groups/targets, typed options and inherited
presentation controls. Their corresponding `usc/aws/` files validate and
measure the feature tree before generic layout. A closed option schema checks
each setting's owner and type; it is not an AWS deployment validator. Hidden
cards do not allocate layout space, but are still validated. Connections to
hidden IDs are projected to their visible ancestors. All modules are declared
centrally in `external/engine/src/mod.rs`, without AWS-local `mod.rs` files.

This is a statically linked built-in composition stage, not executable profile
callbacks. Other AWS catalog tags still use their icon/group/boundary profiles;
native models for all services are not yet implemented. See the [component
syntax and limits](../xal/aws-resources.md#native-albnlb-components-v2).

## Layout and routing

The implemented policies are:

- vertical and horizontal fixed-before-flex allocation;
- per-child weights, margins, dimensions, intrinsic sizes, offsets, and
  min/max constraints;
- generic grids, including a 12-column configuration and row/column spans;
- adaptive item grids that select rows and columns from the available aspect
  ratio and shrink icons when label space requires it;
- absolute placement;
- nested content boxes with padding, gap, alignment, justification, and
  `error` or `visible` overflow;
- owner-relative ports with side, anchor, tangential offset, explicit size,
  boundary-icon overlap validation, and a generic shape-less boundary-icon
  form; and
- straight and deterministic orthogonal routes with generic obstacle scoring,
  endpoint decorations, line styles, and labels.

Every element carries independent typed parameters. Pointer fields on the Go
request preserve unset values separately from explicit `0` and `false`.
Invalid enum values, unknown ABI versions, non-finite numbers, bad ranges,
duplicate IDs, malformed parent indexes, excessive nesting, invalid spans, and
missing line endpoints fail before a renderer receives geometry.

The Rust SVG projection consumes the same immutable resolved document returned
by `Resolve`. It draws generic shapes, text, ports, lines, labels, and endpoint
decorations without recomputing layout or routing. SVG registration continues
to pass through the separate safe normalization operation before SQLite stores
it.

## ABI v5

ABI v5 is a bounded binary contract. All fixed-width values are little-endian;
variable text is UTF-8 with explicit lengths. An input tree is flattened once
in pre-order and each record stores its parent index. Optional numeric and
boolean fields use presence bitsets, so an omitted value cannot collapse into
its zero value. Named string slots carry IDs, text, colors, icon references,
and endpoints; arbitrary maps and renderer JSON are not accepted. Version 3 added
closed AWS component kind/domain/protocol/mTLS/reference fields, a listener
port, and optional backend TLS/mTLS flags. Version 4 adds the optional listener
`show-title` boolean, preserving unset versus false. Version 5 adds bounded
feature kind/subtype/name/value/aux/order slots and detail-level/show/hide
controls. Rust decodes them into closed service models; no arbitrary option
maps, renderer strings or plugin callbacks cross the ABI. The XLE2/XLR2 family
magic remains; the independent version field is 5 and rejects older layouts
before decoding.

The resolved response remains in deterministic pre-order and contains:

- parent identity and generic concept;
- final finite `(x, y, width, height)` geometry;
- resolved renderer-neutral visual and text values;
- resolved text box plus selected icon reference and icon box; and
- final line points, style, decorations, label, and label position.

The engine limits one request to 10,000 elements, 128 nesting levels, 16 MiB
of ABI input, 32 MiB of ABI output, and 2 MiB for one normalized SVG. Rust owns
the response allocation; Go copies it before invoking the matching C free
function.

Go creates a C-owned atomic cancellation handle for each context-aware engine
call. Rust checks it at bounded layout and routing boundaries and returns the
ordinary typed error response; no callback, daemon, or subprocess is involved.

Frontend elements retain source-span IDs and parameter provenance in Go.
Calculation errors are exposed as structured diagnostics and mapped back to
the originating span for CLI and LSP consumers without sending source
contents through the ABI.

## `.xal` lowering and efficiency

The envelope frontend parses the original `.xal` bytes once and lowers the
typed concepts directly to `EngineDocumentSpec`. The same parsed concept tree
can supply diagnostics, LSP symbols, RAG rows, and engine input; those
consumers do not serialize and parse an intermediate representation.

All-item groups select the generic adaptive-grid policy. Its candidate-column
scan is linear in the number of slots, uses the resolved content aspect ratio
and label reservation, and preserves stable source-order tie breaking. Frame
metadata is lowered as ordinary invisible layout containers plus styled text
cells, so its reserved strip and content offset use the same core allocator
instead of renderer-specific coordinates.

V1-profile groups retain the V1 header-tag geometry: 32-pixel group icons,
left-aligned labels, border-top alignment, frame-metadata clearance, and
collision avoidance against preceding headers and group boundaries. Collision
queries use bounded passes over narrow vertical buckets rather than scanning
every prior element. Group SVG colors are normalized once and retained in a
size- and entry-bounded cache.

`ProjectConcept` is a Go type alias of the closed `EngineConcept` vocabulary,
so project analysis, LSP, RAG rows, and future engine lowering do not need
another domain-name-to-concept conversion.

The ABI representation is designed for that reuse:

```text
typed concept tree
  -> one pre-order flatten pass                 O(n)
  -> parent-indexed contiguous ABI records      O(n)
  -> Rust ID validation + child adjacency       O(n)
  -> recursive layout in stable source order
  -> line endpoint lookup by ID                 O(1) average per endpoint
```

This removes a JSON/map conversion and avoids rebuilding hierarchy in every
output adapter. Routing currently evaluates generic obstacle candidates per
line; dense diagrams therefore retain a line-by-obstacle cost and should be
profiled separately from parsing and layout.

### Reference benchmark

The repository includes opt-in benchmarks based on both complex-hybrid sample
versions:

```bash
make build
CGO_ENABLED=1 go test \
  -tags 'xaligo_engine xaligo_exporter sqlite_fts5 sqlite_omit_load_extension' \
  ./test/integration -run '^$' \
  -bench '^BenchmarkComplexHybridV2(RenderSVGEndToEnd|FrontendLower)$' \
  -benchmem -benchtime=100x
```

A 2026-08-30 warm-cache reference run on Apple M2 (`darwin/arm64`) produced:

| Stage | Time/op | Go B/op | Go allocs/op |
|---|---:|---:|---:|
| V2 compatibility frontend lower | 0.314 ms | 340,176 | 4,611 |
| V2 complete SVG render | 3.21 ms | 2,334,155 | 7,199 |

These rows include the V1-style adaptive item grid, frame metadata composition,
group-header collision handling, catalog labels, tinted group assets, generic
routing, and SVG embedding used by the V2 sample. Go's `B/op` does not account
for allocations made by Rust. Use the benchmark to detect regressions within a
row; renderer and routing differences still make a V1/V2 wall-clock comparison
unsuitable as a component benchmark.

Project intelligence remains a separate workflow. Initial RAG discovery
indexes only Markdown under `docs/`. A `.xal` concept tree is analyzed only
when an editor, agent, or future V2 frontend explicitly supplies that document;
it is not added to the initial RAG corpus.

## Compatibility

Existing callers that provide the original flat `Direction`, `Gap`, and
`Elements` fields continue to work. A blank element concept is treated as a
generic `Item`, and vertical or horizontal allocation retains the original
fixed/flexible behavior. New callers can add hierarchy and typed parameters
without changing the three use-case methods.
