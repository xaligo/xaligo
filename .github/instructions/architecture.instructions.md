---
applyTo: "**/*.{go,ts,md,yml,yaml,json}"
---

# Architecture

This document defines the implementation boundaries of xaligo. Product
direction lives in `roadmap.instructions.md`; DSL behavior lives in
`xal-spec.instructions.md`.

## Core pipeline

```text
.xal source
   -> internal/usecase orchestration
   -> internal/usecase/v1/engine parser functions
   -> validated numeric and enum attributes
   -> internal/usecase/v1/engine layout functions
   -> resolved canonical scene
   -> internal/usecase/v1/engine plan calculations
   -> internal/repository output encoder
   -> SVG | Excalidraw | PPTX | PDF | Excel | XYFlow | Isoflow
```

The parent `internal/usecase` package is the shared rendering and orchestration
boundary. Its `v1/engine` subpackage contains synchronous V1 calculation stages.
Format-rendering adapters (CLI, preview server, and WASM) call a
constructor-injected `RenderUsecase` instead of assembling a parallel
parser/layout/render pipeline. Utility
commands such as `generate xal` and `add service` may use their focused internal
builders and repositories directly.

## Language-version boundary

`<xaligo version="1">` selects canonical V1. Root `<frame>` and `<frames>` are
legacy V1 compatibility inputs and emit a migration warning. Native V2
uses the reject-safe `<scene version="2">` root. The V1 parser is not extended
to recognize `<scene>` and does not import or call V2 code.

The parent use-case boundary owns one lightweight root/version dispatch before
engine selection. It must inspect the first XML start element once, reject
contradictory root/version pairs, and pass the original bytes to exactly one
frontend. It must not select a version by retrying another parser after an
error.

V2 provides two frontends: its native `<scene version="2">` frontend and a V1
compatibility frontend that implements the frozen V1 behavior. Both lower
directly to one typed, version-neutral model consumed by V2 layout, routing,
and format encoders. The V1 compatibility path must not rewrite XML, parse a
document twice, serialize through an intermediate V1 scene, or invoke a full
V1 renderer and then reverse-engineer its output. This one-way relationship
allows V2 to render V1 while V1 remains unaware of V2.

## Package responsibilities

| Path | Responsibility |
|---|---|
| `internal/entity` | Independent entity layer containing cross-layer structures |
| `internal/usecase` | Render orchestration, context checks, repository port adaptation, and future parallel scheduling |
| `internal/usecase/v1/engine` | Synchronous V1 parser, validation, layout, scene, routing, and plan calculations; no repository or scheduling ownership |
| `internal/repository` | Filesystem, catalog, HTTP preview, and output-format encoding/export adapters |
| `internal/command.go` | Root Cobra command assembly |
| `internal/controller` | Cobra CLI argument and file-I/O adapters |
| `cmd/wasm` | JavaScript-global adapter over shared use cases and embedded assets |
| `external` | TypeScript external adapter layer mirroring `internal`: `command.ts`, `controller`, `entity`, `repository`, `usecase` |
| `test/unit` | Unit tests mirroring the source tree they cover |
| `test/integration` | Black-box tests of exported APIs and adapters |
| `etc/resources/aws` | Catalogs, templates, embedded assets, and attribution |

## Invariants

1. `.xal` is the only source DSL. Do not add adapter-specific parsers.
2. Mode selects visual semantics; format selects output serialization.
3. Format-rendering production paths call parser and layout through
   `internal/usecase`. Adapters use an injected `usecase.RenderUsecase`;
   controllers use separate narrow use cases for diagnostics, scene I/O,
   catalog access, and persisted export.
4. Routing and connector behavior belongs in shared scene/plan layers, not in
   individual output adapters.
5. Filesystem-less environments provide an `AssetSource`; they do not fork the
   render pipeline.
6. Native configuration remains the default when `RenderOptions.Assets` is nil.
7. New formats require a `Format` value, shared render function, CLI wiring,
   tests, and adapter documentation.
8. Errors are returned and wrapped with context. Core packages do not panic.
9. Native CLI dependency construction belongs in `NewRootCmd`; the WASM entry
   point is its own composition root. Controllers depend on use cases, never on
   other controllers.
10. Input/output destination dependencies belong in `internal/repository` and
    must not appear as use-case filenames.
11. Validation and rendering use the same parse, normalization, and geometry
    checks. An input accepted by `validate` must not fail during rendering
    because of malformed or non-finite geometry.
12. Resolved geometry contains only finite coordinates and strictly positive
    drawable sizes. A child stays inside its parent's content box unless the
    source explicitly selects a non-containing overflow policy.
13. Parent layout owns child allocation. Fixed main-axis sizes are reserved
    before flexible weights divide the remaining space; a child cannot silently
    replace its allocation after sibling positions have been calculated.
14. Text wrapping, fitting, clipping, line height, padding, and semantic role
    are part of the renderer-neutral draw contract. Encoders translate that
    contract and do not invent format-specific text behavior.
15. Geometry and typography use one effective layout-pixel transform. Changing
    PPI or paper fitting must not scale text independently from its shapes.
16. Item-grid placement is resolved before scene encoding and participates in
    the same occupancy and overflow checks as rectangles and other children.
17. Format dispatch has one use-case owner. Commands and controllers collect
    inputs and persist outputs; they do not maintain a second format switch or
    call an external output repository directly.
18. Shared types and functions use format-neutral names. Format-specific names
    may remain only as compatibility aliases at public boundaries. Neutral
    schemas must not require a renderer's JSON fields or generated-ID patterns;
    renaming an adapter-shaped type alone is not completion. Semantic element
    kind and parentage are explicit data; adapters must not reconstruct current
    scene hierarchy from rectangle containment.
19. Interfaces and constructors/factories live in the responsibility file that
    contains their concrete implementation. Do not create declaration-only
    `interface.go`, `interfaces.go`, `constructor.go`, `constructors.go`, or
    equivalent TypeScript files. Place an interface beside the concrete type
    and methods that implement it, and place `NewX`, `createX`, or another
    factory beside the type and behavior it constructs.
20. Layer components do not depend on peer components in the same layer. A
    repository must not construct, retain, or call another repository; a use
    case must not call another independently constructed use case; and a
    controller must not call another controller. Coordination between multiple
    repositories belongs to a use case, and coordination between multiple use
    cases belongs to a controller or composition/public-API boundary.
21. `internal/usecase/v1/engine` exposes synchronous calculation stages. It
    does not import `internal/repository`, interpret `context.Context`, create
    goroutines/channels/worker pools, or choose concurrency limits. The parent
    use case owns I/O, cancellation checks, job partitioning, result ordering,
    and any future parallel execution. Order-dependent routing within one plan
    remains sequential.
22. Language versions are selected by a root/version pair, never by parser
    fallback. `<frame>`/`<frames>` are V1 and `<scene version="2">` is V2, so a
    V1 reader rejects V2 before interpreting nested syntax.
23. V2 renders V1 through a frozen V1 compatibility frontend that lowers once
    into the typed neutral model. V1 has no V2 dependency, and neither XML
    rewriting, double parsing, nor renderer-output round-tripping is allowed.
24. Output schemas are capability projections of the shared semantic model.
    An encoder may omit a value its target schema cannot represent, but it must
    not invent private schema extensions or become an intermediate semantic
    model. Lossy capabilities are documented and tested explicitly.
25. Structural diff compares parsed `.xal` trees, never source lines or
    positional scene/plan IDs. The old side highlights removed and previous
    modified/moved nodes in pale red; the new side highlights added and current
    modified/moved nodes in pale green. Highlight overlays are added after
    layout and route resolution and must not become routing obstacles.
26. An identified child `<frame>` is one physical page by default. SVG emits
    one artifact per frame, PPTX maps one frame to one slide, PDF maps one
    frame to one page, and Excel maps one frame to one worksheet containing
    that frame's SVG image. `CombineFrames` is the explicit compatibility
    policy that preserves the former single-canvas, single-slide, single-page,
    or single-sheet result. Excalidraw, XYFlow, and Isoflow remain single
    logical documents and do not split by frame. Page-oriented encoders omit
    the page-frame outline in default and combined output; the frame remains a
    logical crop/page-link boundary rather than a visible rectangle.
    Excalidraw retains page-frame objects with transparent strokes.
27. Page projection happens only after the complete document scene, connector
    routing, and cross-frame link semantics are resolved. A per-frame encoder
    consumes an ordered `DocumentPlan` projection; it must not parse, lay out,
    route, or infer crop geometry independently. A one-frame SVG render returns
    the exact requested output path, while a multi-frame SVG render uses stable
    frame-derived artifact IDs and rejects filename collisions.
28. Native PDF and Excel encoders remain behind `!js` build constraints. The
    browser adapter uses lightweight `js` repository stubs because those
    formats are not exposed there; native canvas, font, and spreadsheet
    dependencies must not enter the browser-WASM dependency graph.
29. A frame metadata tag band is resolved once in the V1 shared layout and
    presentation scene as page-owned decoration. It is anchored inside the
    frame padding and reuses the selected top/bottom content margin before
    shrinking the content box. Its per-row alignment, greedy wrapping and
    explicit row breaks, text metrics, layer order, per-page ownership, and
    connector avoidance are encoder-independent. SVG, PPTX, PDF, Excel, and
    Excalidraw consume that shared result; XYFlow and Isoflow may omit the
    decoration but must not reinterpret it as graph nodes or endpoints.

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

Current component prefixes are `add`, `diff`, `generate`, `init`, `render`,
`serve`, `validate`, and `version` in `internal/controller`; `render`, `diff`, `diagnostics`,
`scene_io`, `catalog`, `export`, `parser`, `layout`, `element`, `pagination`,
`plan`, `scene`, and `theme` in `internal/usecase`; and `powerpoint`, `preview`,
`isoflow`, `svg`, `pdf`, `spreadsheet`, `xyflow`, `excalidraw`, and `xaligo` in
`internal/repository`. Repository supporting files retain the same prefix, such
as `powerpoint_export.go` and `isoflow_assets.go`. Every direct
`internal/usecase/*.go` file is a complete component as specified in
`coding.instructions.md`.

Calculation files in `internal/usecase/v1/engine` use functional prefixes
such as `parse_*`, `layout_*`, `scene_*`, `route_*`, and `plan_*`. They contain
cohesive algorithm slices and do not repeat the package or architectural layer
name in filenames.

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
  `internal/entity` or `external/entity`; this rule does not move shared data
  models into implementation packages.
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

## Geometry contract

Layout is a constraint-resolution stage, not a best-effort drawing stage. It
must establish these postconditions before a scene or plan is constructed:

- every coordinate, length, weight, gap, margin, padding, and scale is finite;
- drawable width and height are greater than zero;
- row and column weights are greater than zero, and grid spans are in range;
- each content box is derived once from its allocated border box;
- fixed-size children consume space before flexible children are distributed;
- gaps are subtracted exactly once and cursors advance by the resolved size;
- containment or the selected overflow policy is recorded explicitly; and
- invalid geometry is returned as a source-positioned diagnostic, not dropped
  later by scene construction or exposed to an output encoder.

With `overflow="visible"`, fixed children still consume their resolved sizes
and advance the cursor. If they leave no positive remainder while flexible
children exist, the parent's original usable main-axis extent becomes the flex
pool. Children remain in source order, and every sibling cursor advances by the
resolved size plus its declared gap and margins, so the resulting overflow is
explicit. The default `overflow="error"` rejects the same input.

`Validate` and `Render` must both call this same stage. Encoders may reject an
I/O or serialization failure, but they must never be the first component to
discover `NaN`, `Inf`, a negative drawable size, or an impossible grid ratio.

## Renderer-neutral text contract

Every text draw operation carries its resolved box plus a text-layout policy:

```text
wrap | no-wrap
fit: none | shrink
text overflow: visible | clip
line height
content padding
semantic role
```

The semantic role distinguishes ordinary labels, item labels, group headers,
ports, connector labels, and other future text without requiring encoders to
infer behavior from element IDs. Glyph overflow must either be included in
bounds/obstacle calculations or removed by the declared fit/clip policy.

Layout and canonical-scene values are expressed in layout pixels. The current
shared presentation plan stores geometry and padding in inches and font sizes in
points. For effective PPI `p`, conversion is `inch = px / p` and
`pt = px * 72 / p`; paper fitting changes `p` once and both formulas use that
same value. Fixed physical sizes, such as an explicitly specified PPTX label
size, must be represented as an intentional semantic policy rather than an
incidental conversion constant.

## Dependency direction

```text
internal command / controller / cmd/wasm
                  |
                  v
   internal/usecase orchestration
          /                 \
         v                   v
internal/usecase/       internal/repository
  v1/engine               interfaces and
     |                   implementations
     v                         |
 internal/entity <-------------+

external/command.ts
        |
        v
external/controller
        |
        v
external/usecase
        |
        v
external/repository
```

Entity and use-case packages must not depend on CLI, preview, WASM, or
TypeScript adapters. Encoders consume entity structures and must not depend on
use-case implementations merely to access types.

## Verification

Run after structural changes:

```bash
go test ./...
go build ./...
npm install
npm run build --workspace=@ryo/xaligo-external
npm --prefix external run build:pptx-exporter-wasm
```

Generated binaries, `node_modules`, `output`, and package `dist` directories are
ignored and must not be committed.
