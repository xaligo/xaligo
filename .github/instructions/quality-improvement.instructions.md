---
applyTo: "**/*.{go,ts,xal,md,yml,yaml,json}"
---

# Quality Improvement Preconditions

Use this file when planning, implementing, reviewing, or verifying feature
quality work. The goal is to raise each Phase 2 feature from "renders" to a
documented, tested, readable, and cross-format-stable user experience.

## Quality Rubric

Evaluate every feature slice against these gates before considering it done:

1. Specification correctness.
2. Validate/render agreement.
3. Layout correctness.
4. Design quality.
5. Cross-format consistency.
6. Regression tests.
7. Samples and documentation.
8. CI, security, and release readiness.

## Feature Quality Workflow

For each feature, proceed in this order:

1. Identify the authoritative contract in the DSL specification,
   architecture instructions, design documentation, examples, and tests.
2. Establish one observable invariant or failing behavior before editing.
3. Prefer a focused regression test that can fail for the diagnosed cause.
4. Fix the earliest shared layer that owns the behavior. Do not hide shared
   parser, layout, routing, scene, or plan defects inside one renderer.
5. Validate the touched slice with the narrowest useful command, then broaden
   only when the change affects shared contracts or multiple formats.
6. Update examples, documentation, and generated source-controlled assets when
   their visible behavior or contract changes.
7. Commit one coherent feature-quality slice after the required checks pass.

## Design Quality Gate

Treat design quality as a first-class feature requirement, not a cosmetic
follow-up. A feature is not complete only because it validates and renders.

Check these visual qualities for every rendered feature:

- Visual hierarchy: titles, group labels, node labels, metadata tags, and
  connector labels should communicate reading order and importance clearly.
- Spacing and rhythm: node gaps, group padding, table cells, metadata rows,
  connector-label offsets, and page-link terminals should feel intentional and
  consistent with the diagram type.
- Typography: text size, wrapping, fitting, clipping, and full-width character
  handling should remain readable across English, Japanese, symbols, and long
  identifiers.
- Shape language: UML, table, database, AWS, metadata, and page-link elements
  should look meaningfully distinct while still belonging to one visual system.
- Connector readability: route, traffic, UML control flow, object flow,
  relation, and page-link lines should remain distinguishable and avoid
  labels, ports, metadata bands, and important shapes.
- Color and contrast: themes, diff highlights, semantic markers, and metadata
  tags should preserve readable contrast in SVG, Excalidraw, PPTX, PDF, Excel,
  XYFlow, and Isoflow projections where applicable.
- Documentation polish: examples should be good enough for users to copy and
  for documentation to present without extra manual cleanup.

## Phase 2 Quality Order

Harden Phase 2 features in this order unless a blocking defect requires a
smaller local detour:

1. Canonical V1 document envelope: `<xaligo version="1">`, `<data>`,
   `<frames>`, identified `<frame>`, legacy compatibility, and migration
   diagnostics.
2. Data registry and imports: reusable table data, database schemas, SQL
   imports, relative paths, provenance, duplicate IDs, and missing resources.
3. Tables: standard tables, pipe tables, styling, cell text, headers, column
   sizing, empty cells, and imported data.
4. Relational databases: entities, columns, primary keys, foreign keys,
   composite keys, SQL imports, relation routing, and notation-specific labels.
5. UML diagrams: class, sequence, activity, state, component, strict profiles,
   node shapes, flow types, guards, lifelines, activations, and nested metadata.
6. Frame-scoped references: local endpoints, `frame-id.endpoint-id`, duplicate
   IDs across frames, and user-correctable diagnostics.
7. Cross-frame page links: source/destination page terminals, automatic and
   explicit side selection, anchors, labels, inset behavior, and invalid
   metadata-reserved edges.
8. Frame metadata bands: visible `id`, `title`, `version`, custom metadata,
   row gaps, widths, wrapping, top/bottom placement, and collision avoidance.
9. Page-oriented outputs: one frame per SVG artifact, PPTX slide, PDF page, and
   Excel worksheet; `--combine-frames`; safe artifact IDs; crop boundaries; and
   filename collision handling.
10. Renderer matrix: SVG baseline behavior, Excalidraw editability, PPTX plan
    parity, PDF/Excel native constraints, XYFlow/Isoflow projections, and
    browser-WASM dependency boundaries.
11. CI, release, and tooling: Go, TypeScript, docs, security, version gates,
    RTK operation, npm lockfile policy, vendored dependencies, and generated
    artifact policy.

## First Quality Slice

Start with `docs/src/examples/samples/uml-activity.xal` when beginning the
feature-by-feature quality pass. It is a canonical V1 UML activity sample and
should verify:

- `initial`, `action`, `object-node`, `decision`, `fork`, `join`, `merge`, and
  `final` nodes.
- `control-flow` and `object-flow` rendering and visual distinction.
- `guard`, `title`, nested `responsibility`, and nested `constraint` behavior.
- Left-to-right activity readability.
- Decision diamonds, fork/join bars, merge nodes, final nodes, and connector
  labels avoiding collisions.
- SVG output quality sufficient for documentation, with additional format
  checks added when shared contracts or encoder behavior changes.

## Definition of Done

A feature-quality slice is done only when:

- The specification and documentation match the implementation.
- Representative valid samples pass validation and render successfully.
- Representative invalid inputs fail with source-positioned, user-correctable
  diagnostics.
- Regression tests cover the observable behavior being protected.
- Design quality has been reviewed for hierarchy, spacing, typography, shape
  language, connector readability, and contrast.
- Cross-format checks have been run when the feature affects shared scene,
  routing, plan, pagination, or renderer contracts.
- Required security and repository checks have passed.
- The slice is committed without unrelated working-tree changes.