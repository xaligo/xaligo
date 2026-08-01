---
applyTo: ".github/instructions/manual/**"
---

# 08.10 Architecture: V2 generic engine and plugins

## Status and purpose

This file defines the planned V2 architecture. It is a target contract, not a
description of the current V1 implementation. V2 adopts the single-binary,
shared-service, Rust-engine, and external-runtime boundary in
`08-11-architecture-single-binary-service-and-external-runtimes.instructions.md`
and the SQLite SVG registry in
`08-12-architecture-embedded-svg-asset-registry.instructions.md`.

V2 preserves the existing layout concepts, including frames, nested groups,
captures, items, ports, lines, fixed-versus-flexible allocation, the 12-column
grid, item grids, spacing, overflow validation, and shared connector routing.
It removes domain names and domain-specific calculations from the engine so
AWS, UML, and future vocabularies use the same calculation pipeline.

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
   -> native V2 frontend | frozen V1 compatibility frontend
   -> registered vocabulary/profile normalization
   -> typed version-neutral DocumentSpec
   -> versioned Go/Rust engine request
   -> Rust parameter validation and intrinsic measurement
   -> Rust generic layout and constraint resolution
   -> Rust generic port and line routing
   -> typed immutable ResolvedDocument response
   -> renderer-neutral draw/document plan
   -> SVG encoder | PPTX exporter boundary
```

The native V2 and V1 compatibility frontends each parse the original bytes
once and lower directly to `DocumentSpec`. They do not call the V1 engine,
rewrite XML, retry another parser, or serialize through an intermediate V1
scene.

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

Core calculations must never branch on source tags, plugin IDs, icon
namespaces, `aws`, `uml`, or another domain name.

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
│   └── crates/
│       ├── layout-engine/ generic measurement, layout, and routing
│       ├── svg-engine/    SVG validation, normalization, and projection
│       └── ffi/           versioned C ABI static-library boundary
└── pptx-exporter/         TypeScript/PptxGenJS PPTX adapter
```

The intended dependency and data direction is:

```text
native frontend ─────────┐
V1 compatibility ────────┼─> normalized typed IR
builtin/AWS/UML profiles ┘             |
                                       v
                             internal/usecase/v2/engine
                                       |
                              direct in-process ABI
                                       |
                                       v
                         external/engine Rust workspace
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
details. The Go host serializes normalized typed data; the Rust engine never
branches on profile IDs or source tag names.

The engine ABI version is independent from the XAL document version. Reject an
unknown ABI before calculation. Requests and responses use fixed-width,
little-endian fields with finite numeric validation and explicit optional-value
flags. Arbitrary maps and renderer-specific JSON must not cross the ABI.

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
plugin then selects that behavior through data. Do not solve the gap with an
AWS- or UML-specific conditional, callback, or hook.

Profiles are registered explicitly at a composition boundary. Do not use Go
`init` functions for registration. Reject duplicate profile IDs, concept
names, aliases, parameter definitions, and icon references. Freeze the
registry before normalization begins so calculation is deterministic.

## Neutral input and resolved output

The input to calculation is a typed tree of generic elements. Domain source
names are retained only as optional provenance and must not drive calculation.

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
Generated parts remain individually addressable for parameter overrides.

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
- ports: owning element, side, anchor, offset, size, visibility, and label; and
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
- generic item-grid selection and intrinsic-size fitting;
- stack, horizontal/vertical flow, grid, absolute, lane, layered, and timeline
  policies when implemented;
- margin, padding, gap, content-box, alignment, and overflow resolution;
- text and icon intrinsic measurement;
- port and anchor resolution;
- orthogonal and other shared routing policies;
- obstacle collection, collision avoidance, junctions, and line labels;
- frame-local and cross-frame connection geometry; and
- deterministic source ordering and stable tie-breaking.

AWS and UML profiles may select and tune these policies but cannot calculate
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
modules. The Rust engine remains generic and accepts resolved profile data;
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

- Rust layout and SVG crates compile and test without importing AWS or UML;
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
- generated semantic parts have stable IDs and accept independent overrides;
- invalid parameters fail before layout or rendering;
- equivalent normalized documents from builtin, AWS, and UML profiles resolve
  to equal geometry;
- registry and layout results are deterministic;
- missing-icon policies behave identically in SVG and PPTX; and
- the V1 compatibility frontend matches frozen V1 neutral-model and
  resolved-geometry goldens.
