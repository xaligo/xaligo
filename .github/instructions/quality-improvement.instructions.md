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

## Status Tracking

Maintain the feature hardening status in this file so future agents can resume
from the current quality pass without relying on conversation history.

Use these status values:

- `not-started`: the feature is in the Phase 2 scope but has not been audited.
- `auditing`: contracts, samples, tests, and visual output are being inspected.
- `needs-work`: a quality gap is known and has not been fixed.
- `in-progress`: a fix or quality improvement is actively being implemented.
- `blocked`: progress requires a decision, missing tool, or external input.
- `done`: the feature meets the Definition of Done in this file.

Update the status board when a feature-quality slice starts, when a gap is
found, when validation changes the assessment, and when a slice is committed.
Keep each row concise but evidence-based. Prefer commit hashes, test names,
sample paths, and rendered artifact paths over prose-only status notes.

| Order | Feature | Status | Evidence | Next action |
|---:|---|---|---|---|
| 1 | Canonical V1 document envelope | not-started | `docs/src/examples/samples/canonical-v1-envelope.xal` | Audit contract, validation, rendering, and migration diagnostics. |
| 2 | Data registry and imports | not-started | table data, database schema, SQL import commits exist | Inventory import coverage and diagnostics. |
| 3 | Tables | not-started | `docs/src/examples/samples/tables.xal` | Audit table layout, pipe styling, text fitting, and docs image quality. |
| 4 | Relational databases | not-started | `docs/src/examples/samples/databases.xal` | Audit entity layout, key styling, relation routing, and SQL import behavior. |
| 5 | UML diagrams | not-started | `docs/src/examples/samples/uml-activity.xal` | Start with activity diagram visual and behavioral quality. |
| 6 | Frame-scoped references | not-started | frame-qualified endpoint support exists | Audit duplicate ID handling and user-correctable diagnostics. |
| 7 | Cross-frame page links | not-started | `docs/src/examples/samples/page-links.xal` | Audit terminal side selection, labels, insets, and invalid edge handling. |
| 8 | Frame metadata bands | not-started | `docs/src/examples/samples/frame-metadata.xal` | Audit metadata layout, reservation strips, wrapping, and collisions. |
| 9 | Page-oriented outputs | not-started | SVG/PPTX/PDF/Excel pagination commits exist | Audit default artifacts, `--combine-frames`, safe IDs, and crops. |
| 10 | Renderer matrix | not-started | SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, Isoflow encoders | Define representative cross-format checks for changed shared contracts. |
| 11 | CI, release, and tooling | not-started | `VERSION`, workflows, RTK and security instructions | Audit CI gates, version policy, npm lockfile policy, and vendored dependency handling. |

When a row reaches `done`, its evidence should include the commit that closed
the slice and the highest-signal verification commands that passed. If a row
is `blocked`, the next action must state the specific decision or tool needed.

## Detailed Task Backlog

Track detailed work with stable task IDs. Update task status in the nearest
feature row above when the detailed evidence changes; split or add IDs only
when the new task can be verified and committed independently.

### Q01 Canonical V1 Document Envelope

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q01.1 | not-started | Confirm `<xaligo version="1">`, `<data>`, `<frames>`, and `<frame>` parse as the canonical V1 document shape. | parser/unit tests and `xaligo validate docs/src/examples/samples/canonical-v1-envelope.xal` |
| Q01.2 | not-started | Confirm legacy `<frame>` and `<frames>` inputs remain compatible and emit the intended migration warning. | diagnostics tests covering warning text and source position |
| Q01.3 | not-started | Validate duplicate, missing, or unsafe frame IDs and document-level IDs. | invalid-input unit tests with user-correctable diagnostics |
| Q01.4 | not-started | Confirm canonical samples render as separate artifacts and combined compatibility output. | SVG artifact integration check and docs sample render |
| Q01.5 | not-started | Review canonical-envelope docs for current behavior, command accuracy, and image freshness. | `mdbook build docs` and regenerated SVG comparison |

### Q02 Data Registry and Imports

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q02.1 | not-started | Inventory supported `<data>` children and their reference syntax. | spec/docs consistency check |
| Q02.2 | not-started | Verify relative import resolution, missing files, unsupported extensions, and provenance. | diagnostics tests with temp files |
| Q02.3 | not-started | Verify duplicate data IDs and unresolved data references across frames. | parser/validation unit tests |
| Q02.4 | not-started | Confirm imports are deterministic and do not execute commands or read outside the allowed context. | security-oriented unit tests |
| Q02.5 | not-started | Confirm imported data renders identically to equivalent inline data where applicable. | golden or structural scene comparison |

### Q03 Tables

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q03.1 | not-started | Audit standard table layout: headers, rows, empty cells, column widths, and borders. | scene/layout unit tests and SVG sample render |
| Q03.2 | not-started | Audit pipe-table parsing and styling parity with tag-based tables. | parser and render comparison tests |
| Q03.3 | not-started | Verify long English, Japanese, symbols, and narrow columns wrap or fit without overlap. | SVG text geometry assertions |
| Q03.4 | not-started | Review table visual hierarchy, spacing, typography, and docs-image polish. | design review notes plus regenerated sample SVG |
| Q03.5 | not-started | Confirm table behavior in Excalidraw and PPTX when shared table contracts change. | focused format checks |

### Q04 Relational Databases

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q04.1 | not-started | Audit entity and column layout, key styling, nullability, and data-type display. | scene/layout unit tests and database sample render |
| Q04.2 | not-started | Verify primary, foreign, and composite keys from inline and imported SQL schemas. | SQL import and schema unit tests |
| Q04.3 | not-started | Validate relation endpoints, labels, crow-foot notation, and self/multi relations. | routing/scene tests |
| Q04.4 | not-started | Check database diagrams for readable density, relation crossing, and text contrast. | SVG visual inspection and coordinate assertions where practical |
| Q04.5 | not-started | Confirm invalid schemas produce source-positioned, user-correctable diagnostics. | invalid SQL/schema tests |

### Q05 UML Diagrams

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q05.1 | not-started | Audit activity nodes: initial, action, object-node, decision, fork, join, merge, and final. | `uml-activity.xal` validate/render and scene assertions |
| Q05.2 | not-started | Verify activity `control-flow`, `object-flow`, guards, labels, responsibilities, and constraints. | parser/scene tests plus SVG label geometry |
| Q05.3 | not-started | Audit class diagrams: class boxes, attributes, operations, stereotypes, and relations. | class sample render and structural tests |
| Q05.4 | not-started | Audit sequence diagrams: lifelines, messages, activation bars, ordering, and returns. | sequence sample render and layout tests |
| Q05.5 | not-started | Audit state and component diagrams, including strict profile rejection rules. | valid/invalid UML profile tests |
| Q05.6 | not-started | Review UML visual language for consistent shapes, spacing, typography, and connector readability. | design review plus SVG/PPTX checks when affected |

### Q06 Frame-Scoped References

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q06.1 | not-started | Confirm local endpoint references resolve inside the owning frame first. | reference-resolution unit tests |
| Q06.2 | not-started | Confirm `frame-id.endpoint-id` resolves only to the intended frame endpoint. | cross-frame tests |
| Q06.3 | not-started | Validate duplicate IDs across frames, ambiguous endpoints, and missing frames. | diagnostics tests |
| Q06.4 | not-started | Confirm reference diagnostics include actionable source positions. | error-message assertions |
| Q06.5 | not-started | Check docs and examples explain local vs qualified reference behavior. | docs review and `mdbook build docs` |

### Q07 Cross-Frame Page Links

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q07.1 | not-started | Audit automatic page terminal side selection and safe side fallback. | routing tests and page-link sample render |
| Q07.2 | not-started | Verify explicit `src-frame-side`, `dst-frame-side`, anchors, and invalid same-frame usage. | validation tests |
| Q07.3 | not-started | Confirm terminal insets, labels, stubs, and arrowheads avoid endpoints and metadata strips. | SVG coordinate assertions |
| Q07.4 | not-started | Check page-link labels remain readable in multi-page SVG, PPTX, PDF, and Excel projections. | representative format checks |
| Q07.5 | not-started | Review page-link docs and samples for command accuracy and image freshness. | docs build and regenerated images |

### Q08 Frame Metadata Bands

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q08.1 | not-started | Audit visible `id`, `title`, `version`, and custom metadata rows. | metadata unit tests and sample render |
| Q08.2 | not-started | Verify `row-gap`, `width`, `key-width`, wrapping, row breaks, and top/bottom placement. | layout and SVG text tests |
| Q08.3 | not-started | Confirm reservation strips exclude normal content, labels, connectors, and page terminals. | validation/render agreement tests |
| Q08.4 | not-started | Audit collisions with AWS/general group headers and nested tags. | regression tests for nested headers |
| Q08.5 | not-started | Review metadata visual hierarchy, contrast, and docs polish. | design review and regenerated SVGs |

### Q09 Page-Oriented Outputs

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q09.1 | not-started | Verify default SVG emits one artifact per identified frame with stable safe IDs. | integration tests and output filename assertions |
| Q09.2 | not-started | Verify `--combine-frames` compatibility behavior for SVG, PPTX, PDF, and Excel. | CLI integration tests |
| Q09.3 | not-started | Confirm frame crop boundaries, hidden page-frame outlines, and marker safety behavior. | SVG/PDF geometry checks |
| Q09.4 | not-started | Verify filename collision handling and one-frame exact output path behavior. | artifact tests |
| Q09.5 | not-started | Confirm API documentation matches CLI and use-case behavior. | docs/API review |

### Q10 Renderer Matrix

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q10.1 | not-started | Define the minimal representative sample set for SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow. | matrix documented in this file or docs |
| Q10.2 | not-started | Verify shared scene/plan changes are visible consistently across applicable encoders. | focused format matrix command set |
| Q10.3 | not-started | Confirm Excalidraw editability and metadata survive renderer changes. | JSON structural assertions |
| Q10.4 | not-started | Confirm PPTX plan parity and external TypeScript exporter behavior. | Go/WASM plan tests and `npm --prefix external test` |
| Q10.5 | not-started | Confirm native-only PDF/Excel dependencies stay out of browser-WASM builds. | `GOOS=js GOARCH=wasm go list -deps ./cmd/wasm` guard |

### Q11 CI, Release, and Tooling

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q11.1 | not-started | Audit GitHub Actions for Go, TypeScript, docs, version, release, and pages jobs. | workflow review and CI run results |
| Q11.2 | not-started | Confirm `VERSION`, npm package metadata, and release tags stay coherent. | version-gate reproduction |
| Q11.3 | not-started | Confirm npm lockfile policy is reflected in workflows, Makefile, and docs. | TypeScript install/test command |
| Q11.4 | not-started | Verify RTK and security-check preconditions are current and actionable. | `make security-check` and instruction review |
| Q11.5 | not-started | Confirm generated artifacts, vendored dependencies, binaries, caches, and docs images follow repository policy. | `git status`, ignore rules, and release/package review |

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