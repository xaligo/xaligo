---
applyTo: ".github/instructions/manual/**"
---

# 08.10 Architecture: V2 generic engine and plugins

## Status and purpose

This file defines the V2 architecture. The generic Rust concepts, typed Go
request/response, C ABI v5, nested layout, generic port/routing calculation,
SVG projection, builtin icon registry, and LSP/RAG service boundaries are
implemented. The version-selected `<xaligo>` frontend, initial concise V1-style
authoring normalization, adaptive item-grid selection, frame-metadata
composition, and shared SVG/PPTX projection are implemented; complete profile
normalization remains target work. This is not a description of the current V1
implementation. V2 adopts the single-binary,
shared-service, Rust-engine, and external-runtime boundary in
`08-11-architecture-single-binary-service-and-external-runtimes.instructions.md`
and the SQLite SVG registry in
`08-12-architecture-embedded-svg-asset-registry.instructions.md`.

V2 preserves the existing layout concepts, including frames, nested groups,
captures, items, ports, lines, fixed-versus-flexible allocation, the 12-column
grid, item grids, spacing, overflow validation, and shared connector routing.
Generic layout, routing, and encoders remain domain-neutral. Built-in AWS
service components now have a typed native composition stage before that
shared pipeline, as explicitly requested for per-service design control.

In this design, a plugin is a separable package of declarative concept data,
parameter defaults and constraints, aliases, styles, icons, and attribution.
A plugin is not an executable layout or routing extension.

## Supported output boundary

V2 has exactly two output encoders: SVG and PPTX. Markdown support is a host
workflow that renders fenced `xal` blocks through the SVG artifact API and
embeds image references. Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow are
outside the V2 product and plugin contracts. Plugins may tune generic concepts
and parameters, but may not add encoders, format-specific fields, or output
callbacks.

## Pipeline

```text
.xal source
   -> one root/version dispatch
   -> one envelope frontend with V1 | V2 normalization mode
   -> registered vocabulary/profile normalization
   -> typed version-neutral DocumentSpec
   -> versioned Go/Rust engine request
   -> optional native AWS component validation/composition
   -> Rust parameter validation and intrinsic measurement
   -> Rust generic layout and constraint resolution
   -> Rust generic port and line routing
   -> typed immutable ResolvedDocument response
   -> renderer-neutral draw/document plan
   -> SVG encoder | PPTX exporter boundary
```

The envelope frontend parses the original bytes once and lowers directly to
`DocumentSpec` using the selected normalization mode. It does not call the V1
engine, rewrite XML, retry another parser, or serialize through an intermediate
V1 representation.

## Generic calculation concepts

The V2 core calculates only a closed set of renderer- and domain-neutral
concepts:

- `Frame`: physical-page, canvas, crop, clip, and page-link boundary.
- `Group`: nested boundary with a child content box and layout policy.
- `Capture`: annotation or emphasis boundary that participates in ordinary
  containment and connection semantics.
- `Item`: atomic or composed visual content with intrinsic or allocated size.
- `Port`: independently addressable connection terminal owned by another
  element.
- `Line`: connection between endpoints, including route, decorations, and
  labels.
- `Text`: measurable text content with renderer-neutral wrapping, fitting,
  clipping, and semantic-role policies.
- `Spacer`: non-drawing layout participant.

Additional reusable primitives may be added only when their calculation
semantics are domain-neutral. A domain feature should first be expressed by
composition of existing concepts. For example, an AWS VPC is a configured
`Group`; a UML class is a composition of groups, items, and text compartments;
and a UML inheritance relation is a configured `Line`.

Generic layout/routing calculations must never branch on source tags, plugin IDs, icon
namespaces, `aws`, `uml`, or another domain name.

### Native AWS composition boundary

Per-service data definitions belong in `external/engine/src/ent/model/aws/`;
validation and visual composition belong in `external/engine/src/usc/aws/`.
This built-in, statically linked stage is an explicit exception to the former
all-domains-declarative-only rule, not a dynamic plugin callback mechanism.
ALB, NLB, listener and ALB rule/condition/action/transform/target/service/option models
are implemented. Other catalog tags retain
their generic profile behavior; do not claim all services have native models.
Go parses source once and passes typed fields through ABI v5. Rust expands
components into ordinary generic elements, then uses the same layout/router
and resolved SVG/PPTX plan. No XML, arbitrary maps, network access, renderer
branches, watermark, or source rewriting belongs in this stage.
Composition owns listener-card dimensions, header/domain-tag measurement,
per-service TLS/mTLS rules, and stable generated IDs. It must preserve authored
component/listener IDs and not mutate the source document. Generated decoration
parts are internal; native presentation controls select their visibility,
not renderer-specific styling. Component
geometry/attribute constraints are documented in 07.15.03.

## Target package and language boundaries

```text
internal/
├── entity/engine.go       typed Go-side engine request and response
├── usecase/v2/engine.go   cancellation, lifecycle, and Rust invocation
├── parser/                V1 compatibility and native V2 frontends (target)
├── ir/                    version-neutral normalized document (target)
└── core/profiles/         builtin/AWS/UML declarative profiles (target)

external/
├── engine/
│   ├── api.go             Go-side availability contract
│   ├── bridge_cgo.go      cgo C-ABI invocation adapter
│   ├── bridge_stub.go     non-native/browser build fallback
│   ├── Cargo.toml
│   ├── include/xaligo_engine.h
│   ├── lib/               generated ignored static-library link location
│   └── src/
│       ├── lib.rs         staticlib export boundary
│       ├── mod.rs         module registry
│       ├── base.rs        composition root
│       ├── cnf/           ABI constants, limits, and defaults
│       ├── ent/           model plus ABI request/response entities
│       ├── usc/           operation dispatch plus flat layout_*, routing, validation, and SVG files
│       ├── rep.rs         reserved external-representation layer; currently no implementation
│       ├── ctl/           versioned C ABI controller
│       └── util/          explicit codecs/traits, message codes, logger, and errors
└── exporter/              Rust `pptx` adapter with a C ABI
```

The intended dependency and data direction is:

```text
V2 normalization ────────┐
V1 normalization ────────┼─> normalized typed IR
builtin/AWS/UML profiles ┘             |
                                       v
                             internal/usecase/v2/engine
                                       |
                              direct in-process ABI
                                       |
                                       v
                       external/engine Rust staticlib crate
```

The parent `internal/usecase` composition owns root dispatch and profile
registration. `internal/usecase/v2` owns V2 I/O adaptation, cancellation,
stage ordering, Rust library invocation, and concurrency policy. Rust engine
calculations are synchronous and contain no
repository access, command dispatch, daemon discovery, sockets, HTTP calls, or
process spawning. A native build compiles the Rust engine as a platform-native
`staticlib` and links it into the Go executable through cgo. Go invokes the
versioned `extern "C"` functions directly; it does not launch Cargo, load a
sidecar library, or start an engine subprocess at runtime.

Builtin, AWS, and UML profiles are separate declarative data boundaries from
their first implementation. They depend only on the stable V2 profile/model
contract and their own data/assets. They must not depend on repositories, the
V1 implementation, one another, private builtin assets, or Rust implementation
details. The Go host serializes normalized typed data; Rust composition selects
closed AWS component kinds, never arbitrary profile IDs or source tag names.

The current request path is `ctl -> usc -> ctl`: `ctl` receives and returns
versioned ABI data, while `usc` validates and calculates the result. `rep` does
not sit on that path and must not receive calculation logic merely to populate
the layer. It is reserved for a future Rust-owned external representation such
as PPTX package generation; such an encoder may be called by `usc` through
`rep`, while layout and routing remain owned by `usc`. Keep all three layers
shallow and express responsibility through filenames such as `layout_flow.rs`.

The engine ABI version is independent from the XAL document version. Reject an
unknown ABI before calculation. Requests and responses use fixed-width,
little-endian fields with finite numeric validation and explicit optional-value
flags. Arbitrary maps and renderer-specific JSON must not cross the ABI.

ABI field indexes have one authoritative source at
`external/engine/abi/fields.csv`. `make generate-engine-abi` produces the Go
and Rust constants consumed by the explicit codecs. Do not hand-maintain a
second field-order list.

Rust entities do not use derive or serde annotations to generate behavior.
Shared standard-trait implementations live explicitly in `util/clone.rs`,
`util/debug.rs`, `util/default.rs`, and `util/eq.rs`; the request decoder and
response encoder live in `util/deserialize.rs` and `util/serialize.rs`. This is
an organization rule, not permission to serialize the neutral model as JSON.

Rust logging uses message codes and the same `XALIGO_LOG_LEVEL`,
`XALIGO_LOG_STRUCTURED`, `XALIGO_LOG_CALLER`, and `XALIGO_LOG_OUTPUT`
configuration vocabulary as the Go logger. The in-process static library uses
stderr as its default diagnostic destination, emits no source document
contents or sensitive absolute paths, and never exits the Go host. Failures
continue to cross the typed ABI response; logging is supplementary and debug
lifecycle logging remains filtered at the default info level.

## Declarative plugin contract

A profile may declare:

- namespaced concept names and aliases;
- mapping from each concept to one generic concept or a declarative composition
  of generic concepts;
- default parameter values;
- parameter types, finite ranges, allowed enum values, and inheritance rules;
- stable named parts for elements created by a declarative composition;
- icon aliases, legacy catalog IDs, intrinsic icon metadata, and fallback
  icons;
- renderer-neutral colors, line decorations, shapes, and text roles; and
- license and attribution metadata for bundled assets.

A profile must not declare or register:

- layout, sizing, measurement, placement, routing, or obstacle callbacks;
- renderer- or encoder-specific draw operations;
- repository or filesystem access;
- domain-specific post-processing of resolved geometry; or
- format-specific fallback behavior.

If a domain needs behavior that existing core concepts and parameters cannot
express, add or generalize a reusable core concept or algorithm first. The
plugin then selects that behavior through data. The built-in native AWS stage
above is the sole service-specific exception; do not add domain branches to
generic layout, routing, or encoders.

Profiles are registered explicitly at a composition boundary. Do not use Go
`init` functions for registration. Reject duplicate profile IDs, concept
names, aliases, parameter definitions, and icon references. Freeze the
registry before normalization begins so calculation is deterministic.

## Neutral input and resolved output

The input is a typed tree of generic elements with optional closed AWS component
models. Domain source names are optional provenance, not calculation selectors.

```go
type ElementSpec struct {
	ID       ElementID
	Concept  ConceptRef
	Box      BoxOverrides
	Layout   LayoutOverrides
	Visual   VisualOverrides
	Text     *TextOverrides
	Icon     *IconOverrides
	Routing  *RoutingOverrides
	Children []ElementSpec
}
```

The exact Go declarations may evolve during implementation, but the following
separation is mandatory:

- `ElementSpec` contains unresolved defaults, inheritance, and element-level
  overrides.
- parameter resolution produces typed concrete values with source/provenance
  information;
- layout and routing produce an immutable `ResolvedDocument`; and
- renderers and encoders consume the resolved result without recalculating
  dimensions, grids, anchors, paths, text policy, or icon placement.

Every source-authored element has a stable semantic ID. Every part generated by
a declarative composition has a deterministic semantic part ID derived from
its owning element and profile part name, not from renderer-generated IDs.
Declarative parts remain individually addressable for parameter overrides;
native AWS decoration overrides are not exposed yet.

## Fine-grained per-element parameters

Every concrete frame, group, capture, item, port, line, text element, spacer,
and generated semantic part may override its tunable parameters independently.
Plugins supply defaults; they do not make ordinary layout, visual, text, icon,
or routing parameters plugin-global.

At minimum, the typed parameter system must cover:

- geometry: position, width, height, intrinsic size, minimum/maximum size,
  margin, padding, offset, and overflow;
- allocation: fixed/flexible sizing, weight, row/column span, grid rows and
  columns, gap, direction, alignment, and justification;
- visuals: shape, fill, stroke, stroke width/style, corner radius, opacity,
  visibility, and layer;
- text: font, size, color, alignment, wrapping, fitting, clipping, line height,
  padding, and semantic role;
- icons: namespaced reference, size, scale, color, placement, offset, and
  missing-icon policy;
- ports: owning element, side, anchor, tangential offset, size, visibility,
  label, and an icon-only boundary presentation composed from generic visual
  and icon parameters; and
- lines: source/target endpoint, source/target side and anchor, routing policy,
  obstacle margin, line style, endpoint decorations, label, and label
  placement.

Use typed optional values, or an equivalent representation, so an unset value
is distinct from explicit `0`, `false`, an empty string, or another valid zero
value. Preserve source positions and parameter provenance for diagnostics.
Arbitrary string maps must not cross into layout, routing, scene, or encoder
calculation.

The canonical parameter precedence, from lowest to highest, is:

1. safe V2 core default;
2. generic base-concept default;
3. plugin concept default;
4. document or theme default;
5. inherited parent value, only for a parameter declared inheritable and only
   when the child has no direct value;
6. matching style/class overrides in deterministic declaration order; and
7. the concrete element's explicit value.

An element-level explicit value is therefore always the final authority within
the parameter's valid type and range. Structural identity and core geometry
safety invariants are contracts rather than tunable parameters and cannot be
overridden.

Resolve defaults and overrides once before measurement. Reject unknown
parameters, invalid enum values, non-finite values, invalid ranges, and
conflicting constraints with source-positioned diagnostics. Do not silently
clamp a value unless the parameter contract explicitly defines clamping.

## Layout and routing ownership

The core owns all calculations, including:

- fixed-before-flexible child allocation;
- the existing 12-column row/column grid behavior;
- generic adaptive item-grid selection and intrinsic-size fitting through a
  linear candidate-column scan;
- profile-selected group-header geometry and overlap avoidance through bounded
  passes over vertically bucketed prior headers and boundaries;
- stack, horizontal/vertical flow, grid, absolute, lane, layered, and timeline
  policies when implemented;
- margin, padding, gap, content-box, alignment, and overflow resolution;
- text and icon intrinsic measurement;
- port and anchor resolution, including strict owner-contained ports and
  shape-less icon ports that intersect one selected owner border while staying
  within its tangential extent;
- orthogonal and other shared routing policies;
- obstacle collection, collision avoidance, junctions, and line labels;
- frame-local and cross-frame connection geometry; and
- deterministic source ordering and stable tie-breaking.

Declarative AWS and UML profiles may select and tune these policies but cannot calculate
columns, cells, bounds, coordinates, anchors, routes, or labels themselves.
Equivalent normalized generic input must produce equivalent resolved geometry
regardless of which profile vocabulary created it.

## Builtin generic icons

V2 includes a default, domain-neutral icon profile so the engine remains useful
without AWS or UML. Its stable namespace is `builtin`. The initial catalog
should cover generic concepts such as `generic`, `user`, `service`,
`application`, `server`, `database`, `storage`, `network`, `cloud`, `queue`,
`gateway`, `document`, and `terminal`.

An icon definition provides renderer-neutral data and metadata, including its
namespaced reference, media type or generic vector description, view box,
intrinsic dimensions, tint policy, and attribution. The core measures and
places icons but does not interpret their domain meaning.

AWS and UML assets live with their respective profiles. Removing either
profile must remove its catalog and attribution without removing or changing
the builtin catalog. Domain profiles may select a builtin icon as a declared
fallback. Missing-icon behavior is a generic per-element parameter such as
`error`, `fallback`, or `hide`; `hide` preserves the resolved layout slot.

## Extraction boundary

The in-repository builtin, AWS, and UML profile packages are the initial plugin
boundaries. A later physical extraction may produce independent profile
modules. Generic Rust algorithms accept resolved profile data;
profiles do not link Rust callbacks or executable extensions. Before a split,
stabilize the public profile/model and engine ABI contracts. Do not expose
repository, V1, command, SQLite, or encoder types merely to make an extracted
profile compile.

The core must run and render generic diagrams with only `builtin` registered.
Registering or removing AWS or UML changes only the accepted vocabulary,
defaults, styles, and available assets. It must not change the core algorithm
set or the geometry of an already normalized generic document.

## Verification gates

V2 implementation slices must include tests that establish:

- generic documents remain unchanged with the built-in AWS composition stage present;
- Go invokes the Rust engine in-process without a subprocess or daemon;
- the generated Rust static library is linked only by matching native target
  builds and is not committed as a repository artifact;
- ABI version, malformed payload, non-finite number, and range failures are
  rejected deterministically;
- builtin-only rendering supports generic frames, groups, captures, items,
  ports, icons, and lines;
- profile schemas contain only known generic concepts and parameters;
- explicit element values override every lower-precedence default;
- unset values remain distinct from explicit zero/false values;
- generated semantic parts have stable IDs; declarative parts accept independent overrides;
- invalid parameters fail before layout or rendering;
- equivalent normalized documents from builtin, AWS, and UML profiles resolve
  to equal geometry;
- registry and layout results are deterministic;
- missing-icon policies behave identically in SVG and PPTX; and
- the V1 compatibility frontend matches frozen V1 neutral-model and
  resolved-geometry goldens.
