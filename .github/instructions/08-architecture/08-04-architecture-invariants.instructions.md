---
applyTo: ".github/instructions/manual/**"
---

# 08.04 Architecture: Invariants

## Invariants

1. `.xal` is the only source DSL. Do not add adapter-specific parsers.
2. Mode selects visual semantics; format selects output serialization.
3. Format-rendering production paths call parser and layout through
   `internal/usecase`. Adapters use an injected `usecase.RenderUsecase`;
   controllers use separate narrow use cases for diagnostics and other
   supporting workflows.
4. Routing and connector behavior belongs in shared scene/plan layers, not in
   individual output adapters.
5. Filesystem-less environments provide an `AssetSource`; they do not fork the
   render pipeline.
6. Native configuration remains the default when `RenderOptions.Assets` is nil.
7. The supported engine formats are a closed set: SVG and PPTX for V1/V2, plus
   terminal text for V2 only. Markdown is an SVG-embedding workflow. Any other
   format requires a new explicit product-scope decision before implementation.
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
22. Language versions are selected by the document root and version, never by
    parser fallback. `<frame>`/`<frames>` are legacy V1,
    `<xaligo version="1">` is canonical V1, and
    `<xaligo version="2">` is V2. A V1 reader rejects the V2 version before
    interpreting nested syntax.
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
    one artifact per frame and PPTX maps one frame to one slide.
    `CombineFrames` explicitly requests one combined SVG canvas or PPTX slide.
    Both encoders omit the page-frame outline in default and combined output;
    the frame remains a logical crop/page-link boundary rather than a visible
    rectangle.
    A default page-local SVG uses the exact logical frame rectangle as its
    canvas and clip boundary, without adding stroke/marker safety padding.
    Combined SVG output retains marker-safe canvas expansion.
27. Page projection happens only after the complete document scene, connector
    routing, and cross-frame link semantics are resolved. A per-frame encoder
    consumes an ordered `DocumentPlan` projection; it must not parse, lay out,
    route, or infer crop geometry independently. A one-frame SVG render returns
    the exact requested output path, while a multi-frame SVG render uses stable
    frame-derived artifact IDs and rejects filename collisions.
28. Retired Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow adapters,
    dependencies, aliases, assets, and browser globals must remain absent.
    The browser adapter exposes SVG rendering, PPTX-plan construction, and
    diagnostics only.
29. A frame metadata tag band is resolved once in the V1 shared layout and
    presentation scene as page-owned decoration. The resolved metadata
    `row-gap`, which defaults to 4 layout pixels, is both the inter-row spacing
    and the metadata page-edge inset. The selected top/bottom band edge and
    both horizontal row bounds are inset by that value, and every row wraps and
    aligns within `frame width - 2 * row-gap`; frame padding, content margins,
    and content-box offsets do not replace or add to this inset. The full-width
    reservation strip still runs from the outer logical frame edge to the final
    content-box boundary and is at least
    `row-gap + complete band height + 8` layout pixels deep.
    The inset is measured from the logical frame edge before any common PPTX
    slide centering; it is not an export `--paper-margin`. Normal items, text,
    local/UML connector paths and labels, and cross-frame page links cannot
    enter it. Legacy/automatic page-link side selection remaps a reserved edge
    to the nearest safe edge and clamps left/right terminals outside the strip.
    An explicit cross-frame `src-frame-side`, `dst-frame-side`,
    `src-frame-anchor`, or
    `dst-frame-anchor` that selects the reservation is instead a validation
    error. When metadata is enabled, its resolved `row-gap` is also the inward
    normal inset for page-link terminals on all four sides; without metadata,
    the terminal inset is 4 layout pixels. A resolved zero `row-gap` retains
    the outer logical frame edge. An explicit frame side/anchor requires the
    inset to fit that side's normal dimension and its actual terminal to avoid
    the reservation; failures are reported at the connection source position.
    Without an explicit frame terminal, validation only requires a non-empty
    set of sides satisfying those rules. Shared scene construction uses actual
    endpoint visual geometry to retain a safe preference or choose the nearest
    safe side; it does not use validator `Box` geometry to predict that side. A
    safe selected `left`/`right` terminal is not rejected for an unused
    top/bottom inset line. The inset is never implicitly clamped. Page-link
    labels stay adjacent to the final inset terminal with a 4-layout-pixel gap
    while avoiding metadata and endpoint geometry. These rules, text metrics,
    layer order, and per-page ownership are
    encoder-independent. SVG and PPTX consume that shared result; Markdown
    inherits it through embedded SVG artifacts.
30. Cross-frame connector geometry distinguishes the item endpoint from the
    logical page terminal. `src-side`/`dst-side` and
    `src-anchor`/`dst-anchor` bind the endpoint; `src-frame-side`/
    `dst-frame-side` and `src-frame-anchor`/`dst-frame-anchor` select the
    owning frame side independently. The outer logical frame edge supplies the
    side and 10/30/50/70/90-percent tangent coordinate, but the drawable frame
    terminal is on the parallel page-terminal inset line. Applying the inset
    changes only the normal coordinate; an explicit frame anchor retains its
    tangent coordinate. The endpoint- and frame-terminal-adjacent route
    segments are perpendicular to their respective sides. Frame-terminal
    attributes are invalid on same-frame connections. At zero inset, an owning
    frame endpoint that coincides with an explicit frame anchor is a
    source-positioned validation error; explicit endpoint anchors keep their
    slot, while explicit endpoint sides and automatic endpoint sides resolve to
    their center slot for this check. Automatic left/right coincidence
    avoidance uses the corner gutter and metadata clearance when possible, but
    a tiny safe range falls back to the full non-reserved interval without
    leaving the frame.
