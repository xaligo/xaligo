# V2 Generic Engine

The V2 calculation core is a domain-neutral Rust library linked into the
`xaligo` executable. It does not run as a daemon, sidecar, dynamic plugin, or
separate command:

```text
Go use case -> cgo -> C ABI v2 -> Rust staticlib
```

The public Go use-case boundary remains:

```go
type EngineUsecase interface {
    Resolve(context.Context, entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error)
    RenderSVG(context.Context, entity.EngineDocumentSpec) ([]byte, error)
    NormalizeSVG(context.Context, []byte) (entity.EngineSVG, error)
}
```

The generic calculation core and this ABI are implemented. The native
`<scene version="2">` frontend and the frozen V1 compatibility frontend still
need to lower `.xal` directly to `EngineDocumentSpec`; the existing public V1
render path is not silently redirected through an incomplete adapter.

## Rust source structure

The engine is one staticlib crate and follows the responsibility structure of
[`ryo-arima/vem/src`](https://github.com/ryo-arima/vem/tree/main/src):

```text
external/engine/src/
├── lib.rs                  staticlib export boundary
├── mod.rs                  module registry
├── base.rs                 decode/execute/encode composition root
├── cnf/engine.rs           ABI constants, limits, and defaults
├── ent/model/              generic document and normalized SVG models
├── ent/request/engine.rs   binary request decoding
├── ent/response/engine.rs  binary response encoding
├── rep/layout.rs           layout, validation, ports, and routing
├── rep/svg.rs              SVG normalization and projection
├── usc/engine.rs           operation orchestration
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

This preserves the `cnf / ent / rep / usc / ctl / util` dependency vocabulary
without copying VEM's CLI-specific `main.rs`. `lib.rs` is the corresponding
library entry point. The C symbols and Go `EngineUsecase` contract are
unchanged by this source-only reorganization.

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
| `Port` | Addressable endpoint placed inside its owner |
| `Line` | Straight or orthogonal route between IDs |
| `Text` | Renderer-neutral intrinsic text measurement and label data |
| `Spacer` | Non-drawing layout participant |

AWS, UML, and future vocabularies must lower to these concepts. No profile ID,
source tag, icon namespace, `aws`, or `uml` discriminator crosses into a Rust
calculation branch.

## Layout and routing

The implemented policies are:

- vertical and horizontal fixed-before-flex allocation;
- per-child weights, margins, dimensions, intrinsic sizes, offsets, and
  min/max constraints;
- generic grids, including a 12-column configuration and row/column spans;
- absolute placement;
- nested content boxes with padding, gap, alignment, justification, and
  `error` or `visible` overflow;
- owner-relative ports with side, anchor, offset, and explicit size; and
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

## ABI v2

ABI v2 is a bounded binary contract. All fixed-width values are little-endian;
variable text is UTF-8 with explicit lengths. An input tree is flattened once
in pre-order and each record stores its parent index. Optional numeric and
boolean fields use presence bitsets, so an omitted value cannot collapse into
its zero value. Named string slots carry IDs, text, colors, icon references,
and endpoints; arbitrary maps and renderer JSON are not accepted.

The resolved response remains in deterministic pre-order and contains:

- parent identity and generic concept;
- final finite `(x, y, width, height)` geometry;
- resolved renderer-neutral visual and text values;
- selected icon reference; and
- final line points, style, decorations, label, and label position.

The engine limits one request to 10,000 elements, 128 nesting levels, 16 MiB
of ABI input, 32 MiB of ABI output, and 2 MiB for one normalized SVG. Rust owns
the response allocation; Go copies it before invoking the matching C free
function.

## `.xal` lowering and efficiency

The native frontend should parse the original `.xal` bytes once and lower the
typed concepts directly to `EngineDocumentSpec`. The same parsed concept tree
can supply diagnostics, LSP symbols, MCP inspection, and engine input, but
those consumers must not serialize and parse an intermediate scene.

`ProjectConcept` is a Go type alias of the closed `EngineConcept` vocabulary,
so project analysis, LSP, MCP, RAG rows, and future engine lowering do not need
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

The repository includes opt-in benchmarks based on
`docs/src/examples/samples/complex-hybrid-architecture.xal`:

```bash
make build-engine
GOCACHE=/tmp/xaligo-go-cache CGO_ENABLED=1 \
  go test -tags 'xaligo_engine sqlite_fts5 sqlite_omit_load_extension' \
  ./test/integration -run '^$' -bench 'BenchmarkComplexHybrid' \
  -benchmem
```

A 2026-08-02 reference run on Apple M2 (`darwin/arm64`) produced:

| Stage | Time/op | Go B/op | Go allocs/op |
|---|---:|---:|---:|
| Explicit `.xal` concept analysis | 0.997 ms | 671,388 | 7,305 |
| Existing V1 complete SVG render | 5.871 s | 922,836,714 | 2,573,735 |
| V2 sample-scale ABI + Rust resolve | 4.570 ms | 265,538 | 3,830 |
| V2 sample-scale ABI + Rust SVG | 4.672 ms | 228,940 | 16 |

This is deliberately a stage-level measurement, not a claimed V1/V2 render
speedup. The V1 row includes its complete parser, catalog, compatibility scene,
routing, and SVG work. Until the parity frontend exists, the V2 rows preserve
the sample's generic concept count, hierarchy, and line count but use neutral
benchmark parameters rather than equivalent visual assets. Go's `B/op` also
does not account for allocations made by Rust. Use the benchmark to detect
regressions within each row; compare end-to-end V1 and V2 only after both
frontends produce equivalent resolved documents.

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
