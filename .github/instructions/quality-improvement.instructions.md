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
9. Determinism, performance, and resource safety.
10. Accessibility and theme readability.
11. API, CLI, and preview compatibility.

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
| 1 | Canonical V1 document envelope | in-progress | `TestV1ParseValidatesCanonicalEnvelopeHierarchy`; canonical sample validates and renders two SVG artifacts | Finish migration-warning/docs-image audit, then close Q01. |
| 2 | UML diagrams | in-progress | Activity swimlane slice adds parser/layout/scene support for `partition`, `lanes="vertical"`, `theme="xaligo"`, and `route="loop"`; target SVG renders from `docs/src/examples/targets/uml-activity-atm-swimlane.xal` | Finish activity visual parity and cross-format checks, then continue per-diagram semantic/design passes. |
| 3 | Data registry and imports | not-started | table data, database schema, SQL import commits exist | Inventory import coverage and diagnostics after the UML precision/design pass starts. |
| 4 | Tables | not-started | `docs/src/examples/samples/tables.xal` | Audit table layout, pipe styling, text fitting, and docs image quality. |
| 5 | Relational databases | not-started | `docs/src/examples/samples/databases.xal` | Audit entity layout, key styling, relation routing, and SQL import behavior. |
| 6 | Frame-scoped references | not-started | frame-qualified endpoint support exists | Audit duplicate ID handling and user-correctable diagnostics. |
| 7 | Cross-frame page links | not-started | `docs/src/examples/samples/page-links.xal` | Audit terminal side selection, labels, insets, and invalid edge handling. |
| 8 | Frame metadata bands | not-started | `docs/src/examples/samples/frame-metadata.xal` | Audit metadata layout, reservation strips, wrapping, and collisions. |
| 9 | Page-oriented outputs | not-started | SVG/PPTX/PDF/Excel pagination commits exist | Audit default artifacts, `--combine-frames`, safe IDs, and crops. |
| 10 | Renderer matrix | not-started | SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, Isoflow encoders | Define representative cross-format checks for changed shared contracts. |
| 11 | CI, release, and tooling | not-started | `VERSION`, workflows, RTK and security instructions | Audit CI gates, version policy, npm lockfile policy, and vendored dependency handling. |
| 12 | Diagnostics and error UX | not-started | `internal/usecase/diagnostics.go`, diagnostics tests | Audit severity, source positions, aggregation, cancellation, and actionable messages. |
| 13 | Determinism, concurrency, and performance | not-started | render determinism and concurrency tests | Audit repeatability, ordering, cancellation, limits, and representative performance. |
| 14 | CLI, API, and preview contracts | not-started | controllers, use-case API tests, preview tests | Audit option parity, exit behavior, live preview, and backward compatibility. |
| 15 | Themes and accessibility | not-started | theme application and renderer theme tests | Audit contrast, semantic distinction, text alternatives, and light/dark parity. |
| 16 | Structural diff | not-started | diff engine and highlight tests | Audit matching, moved/modified classification, highlights, and pagination interaction. |
| 17 | Reproducibility and artifact integrity | not-started | deterministic render and generated docs assets | Audit byte stability, clean regeneration, package contents, and stale artifacts. |
| 18 | Configuration, logging, and observability | not-started | config, options, and shared logging code | Audit defaults, precedence, redaction, error context, and noisy-output behavior. |

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
| Q01.1 | done | Confirm `<xaligo version="1">`, `<data>`, `<frames>`, and `<frame>` parse as the canonical V1 document shape. | `TestV1ParseValidatesCanonicalEnvelopeHierarchy`; `go run ./cmd validate docs/src/examples/samples/canonical-v1-envelope.xal` |
| Q01.2 | not-started | Confirm legacy `<frame>` and `<frames>` inputs remain compatible and emit the intended migration warning. | diagnostics tests covering warning text and source position |
| Q01.3 | done | Validate duplicate, missing, or unsafe frame IDs and document-level IDs. | `TestV1ParseValidatesCanonicalEnvelopeHierarchy` |
| Q01.4 | in-progress | Confirm canonical samples render as separate artifacts and combined compatibility output. | canonical sample renders `canonical-v1-envelope-overview.svg` and `canonical-v1-envelope-database-detail.svg`; combined compatibility still pending |
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
| Q05.1 | not-started | Freeze the supported UML sample matrix and map each sample to its owning syntax, parser, scene, layout, routing, and renderer responsibilities. | inventory of `docs/src/examples/samples/uml-*.xal`; owner map recorded in this file or docs |
| Q05.2 | not-started | Establish a per-UML visual baseline before edits: validate and render every `uml-*.xal` sample to SVG, then identify overlap, spacing, typography, connector, and semantic-notation gaps. | `go run ./cmd validate docs/src/examples/samples/uml-all.xal`; SVG render set under `output/uml-quality/` |
| Q05.3 | in-progress | Improve activity-diagram semantic accuracy: activity partitions/swimlanes, initial/final nodes, actions, object nodes, decisions, forks, joins, merges, responsibilities, constraints, guards, `control-flow`, `object-flow`, cross-frame page links, and backward loop edges. | `TestUMLActivityPartitionsNormalizeV1EngineParseUML`; `TestUMLActivityElementsSupportCrossFramePageLinks`; `TestRenderActivityCrossFrameTargetExample`; `route="loop"` retained as UML route metadata; target samples render |
| Q05.4 | in-progress | Improve activity-diagram design quality: xaligo-logo color theme, vertical/horizontal lane headers, left-to-right/top-to-bottom reading flow, diamond/bar/final-node proportions, lane spacing, label placement, no redundant container title/border, and control-vs-object-flow distinction. | `TestUMLActivitySwimlanesReachEditableScene`; `TestUMLActivityHorizontalSwimlanesReachEditableScene`; `TestUMLActivityHidesRedundantContainerBorderAndTitle`; regenerated `output/uml-activity-atm-swimlane-rendered.svg` visual reference |
| Q05.5 | in-progress | Improve class-diagram semantic accuracy: class boxes, attributes, operations, visibility, stereotypes, abstract/static markers, inheritance, realization, association, aggregation, composition, dependency, and multiplicity. | `TestUMLClassStereotypeAndModifiersReachEditableScene`; `TestUMLAggregationAndCompositionRemainHeadlessAtDestination`; `uml-class.xal` includes stereotype/abstract/static classifier examples |
| Q05.6 | in-progress | Improve class-diagram design quality: compartment rhythm, long member wrapping, stereotype readability, relation label spacing, crow-foot/diamond marker clarity, and dense-layout crossing reduction. | Class classifiers use compact multi-row placement, cyan headers with white text, white bodies, real attribute/operation dividers, and deep-blue relation lines; UML class containers hide redundant border/title; package grids auto-balance to the frame and expand to assigned cells; `uml-class.xal` includes frame metadata and a cross-frame repository detail arrow; `TestUMLClassPackagesAutoBalanceToFrameArea`, `TestUMLClassHidesRedundantContainerBorderAndTitle`, `TestUMLClassDiagramSupportsPackageGroups`, and `TestUMLClassPackageSVGKeepsReadableCompartmentText` cover the behavior |
| Q05.7 | in-progress | Improve sequence-diagram semantic accuracy: lifelines, participants, activation bars, sync/async messages, returns, self messages, create/delete, ordering, and message labels. | `TestUMLSequenceMessagesRenderActivationBars`; `TestUMLSequenceMessagesUseResponseAndStopNotation`; `TestUMLSequenceMessagesDistinguishSyncAndAsyncNotation`; `TestUMLSequenceOrderControlsVerticalMessageAnchors`; `TestUMLSequenceOrderAnchorsRemainTopToBottomForEveryConnectionSide`; validation rejects non-destructive `destroy-message` labels and messages leaving inactive lifelines; `uml-sequence.xal` includes message/create/return examples |
| Q05.8 | in-progress | Improve sequence-diagram design quality: timeline spacing, activation contrast, message arrow style, return-line distinction, lifeline header readability, and vertical density. | Participants and lifelines render as compact cyan headers with deep-blue dashed vertical axes; sequence messages render destination activation bars with deep-blue outlines and white bodies; activation bars extend through related calls, returns, and cleanup messages, and fully contained activations merge into their covering bar; synchronous messages use filled triangle arrows; asynchronous messages use open arrows in SVG/plan output; return messages are dashed response connectors; destroy messages add destination stop markers; sequence metadata is retained in editable scene output |
| Q05.9 | in-progress | Improve state-machine semantic accuracy: initial/final states, composite states, transitions, events, guards, effects, entry/do/exit/internal actions, history states, and invalid transition diagnostics. | `TestUMLStateMachineFinalRendersFinalDot`; `TestUMLStateMachineConceptLabelsReachEditableScene`; `uml-state-machine.xal` validates and renders event/guard/action labels plus state behavior compartments |
| Q05.10 | in-progress | Improve state-machine design quality: state shape proportions, grid placement, nested-state padding, transition bend points, guard/effect label placement, class-aligned color theme, optional element-name visibility, and final-node readability. | `TestUMLStateMachinePseudostatesKeepCompactProportions`; `TestUMLStateMachineRowsSeparateBranches`; `TestUMLStateMachineGridColumnsAlignRelatedRows`; `TestUMLStateMachineContainerRowsAndColumnsPlaceStates`; `TestUMLStateMachineConnectorsAvoidIntermediateStates`; `TestUMLStateMachineBentConnectorsAvoidIntermediateStates`; `TestUMLStateMachineBentRouteAvoidanceStaysInsideFrame`; `TestUMLStateMachineDistantConnectorsUseOuterDetours`; `TestUMLStateMachineSampleSVGRoutesStayInsideFrameAndAvoidStates`; `TestUMLStateMachineCanHideElementNames`; `TestUMLRelationBendsReachEditableScene`; `TestUMLRelationLabelsAvoidEndpointItems`; state-machine final nodes render as compact UML final states with an inner dot; layout-only `container`/`row`/`col` tags provide grid placement and relation-aware column inference remains available; class/activity/state-machine relations support child `bend` tags; state-machine transition routes avoid intermediate state bodies, stay inside same-frame bounds, and can use larger outside detours for distant or tight bent routes; SVG draw-plan routing also treats state-machine nodes as obstacles and clamps routes to the page frame; UML relation labels avoid endpoint items; state names use cyan headers with white text by default and can be suppressed with `show-element-names="false"` while retaining compartments; state bodies are white with row and key/value dividers, and state boxes use uniform sizing where practical; `uml-state-machine.xal` models the EC order status flow with container-separated normal/refund/return/cancel branches and regenerated `docs/src/images/uml-state-machine.svg` |
| Q05.11 | not-started | Improve component/deployment/package/composite-structure semantic accuracy: components, ports, interfaces, nodes, artifacts, packages, dependencies, connectors, and containment rules. | `uml-component.xal`, `uml-deployment.xal`, `uml-package.xal`, and `uml-composite-structure.xal` tests |
| Q05.12 | not-started | Improve component/deployment/package/composite-structure design quality: UML adornment placement, nested group padding, port readability, dependency routing, and package tab proportions. | rendered SVG set plus shared group/port geometry assertions |
| Q05.13 | not-started | Improve interaction, communication, timing, object, use-case, and profile diagrams for supported UML notation and explicit unsupported-profile diagnostics. | `uml-interaction-overview.xal`, `uml-communication.xal`, `uml-timing.xal`, `uml-object.xal`, `uml-use-case.xal`, and `uml-profile.xal` tests |
| Q05.14 | in-progress | Normalize shared UML visual language across all diagram types: typography scale, stroke weights, marker sizes, semantic colors, label backgrounds, anchor profiles, and light/dark contrast. | All UML containers omit redundant outer border/title; visible titles/versions move to frame metadata in the individual UML samples and `uml-all.xal`; `TestUMLSequenceHidesRedundantContainerBorderAndTitle` extends the chrome rule beyond activity/class diagrams; `TestUMLCommonRelationAnchorsUseShapeProfiles` fixes rectangle-like endpoints to five anchors per side and diamond-like endpoints to vertices in editable scene output |
| Q05.15 | in-progress | Verify UML cross-format parity for every shared scene/plan change: SVG baseline, Excalidraw editability, PPTX plan/export, PDF/Excel page projection, and XYFlow/Isoflow applicability. | `TestUMLCommonRelationAnchorsReachDrawPlanRoutes` verifies shared plan route endpoints use the same rectangle and diamond UML anchor profiles as Excalidraw bindings; broader renderer matrix remains pending |
| Q05.16 | in-progress | Refresh UML documentation and examples after implementation slices, including generated images and command accuracy. | Individual `uml-*.xal` samples and `uml-all.xal` use frame metadata for titles/versions; docs/spec describe `<uml>` as a semantic container and frame metadata as the visible title surface; UML reference/spec document common rectangle 20-point and diamond 4-vertex relation endpoint profiles |

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

### Q12 Diagnostics and Error UX

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q12.1 | not-started | Audit diagnostic severity, line/column positions, element context, and stable wording. | diagnostics unit tests |
| Q12.2 | not-started | Confirm multiple independent input errors are reported without hiding the first actionable cause. | aggregate-diagnostics tests |
| Q12.3 | not-started | Confirm warnings are non-blocking while errors consistently fail validate and render. | validate/diagnose agreement tests |
| Q12.4 | not-started | Verify canceled contexts stop diagnostics and rendering with wrapped context errors. | cancellation tests |
| Q12.5 | not-started | Review CLI and preview presentation of diagnostics for concise, user-correctable output. | controller/preview tests and manual command check |

### Q13 Determinism, Concurrency, and Performance

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q13.1 | not-started | Confirm identical source, options, assets, and environment produce byte-stable output where the format allows it. | render determinism tests |
| Q13.2 | not-started | Confirm parallel jobs preserve document/page/artifact order and do not share mutable render state. | concurrency and race-enabled tests |
| Q13.3 | not-started | Verify cancellation propagates through I/O and orchestration without goroutine or temporary-file leaks. | cancellation/leak tests |
| Q13.4 | not-started | Establish representative render benchmarks for complex architecture, tables, database, UML, and multi-frame documents. | Go benchmarks with recorded baseline |
| Q13.5 | not-started | Define and test safe behavior for extreme node counts, text lengths, dimensions, ratios, and import sizes. | bounded stress tests and finite-geometry assertions |

### Q14 CLI, API, and Preview Contracts

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q14.1 | not-started | Audit CLI flags, defaults, aliases, validation, exit codes, stdout/stderr, and output path behavior. | controller/command integration tests |
| Q14.2 | not-started | Confirm constructor-injected use-case APIs and convenience methods return equivalent contracts. | use-case API tests |
| Q14.3 | not-started | Verify native and embedded asset sources produce equivalent output at matching settings. | native/embedded parity tests |
| Q14.4 | not-started | Audit live preview initial render, reload events, diagnostics, browser refresh, and file-change recovery. | preview tests and local serve check |
| Q14.5 | not-started | Confirm legacy mode/format aliases and public option defaults remain backward compatible. | compatibility tests and docs review |

### Q15 Themes and Accessibility

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q15.1 | not-started | Audit light/dark theme contrast for text, borders, fills, connectors, metadata, and diff highlights. | contrast calculation plus theme renders |
| Q15.2 | not-started | Confirm semantic distinctions do not depend on color alone when line style, marker, label, or shape can carry meaning. | renderer structural assertions and design review |
| Q15.3 | not-started | Verify text remains readable at documented output sizes and under zoom/scaling across formats. | SVG/PPTX/PDF visual checks |
| Q15.4 | not-started | Audit SVG titles/descriptions, meaningful identifiers, and document structure available to assistive tooling where supported. | SVG structure tests |
| Q15.5 | not-started | Verify full-width, combining, RTL, and uncommon glyph fallback behavior is documented or safely handled. | typography tests with representative strings |

### Q16 Structural Diff

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q16.1 | not-started | Audit added, removed, modified, moved, and unchanged classification independent of source ordering. | diff classification tests |
| Q16.2 | not-started | Confirm matching uses semantic identity rather than generated scene IDs or source lines. | fingerprint/matching tests |
| Q16.3 | not-started | Verify highlights remain behind labels, outside routing obstacles, and readable in light/dark themes. | diff scene/SVG tests |
| Q16.4 | not-started | Confirm multi-frame diff pagination and output pairing remain stable. | document-plan integration tests |
| Q16.5 | not-started | Review diff samples and docs for realistic change cases and visual clarity. | sample render and docs build |

### Q17 Reproducibility and Artifact Integrity

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q17.1 | not-started | Confirm checked-in documentation SVGs reproduce byte-for-byte or through a documented normalization step. | regenerate-and-compare command |
| Q17.2 | not-started | Detect stale generated images, missing source `.xal` files, and orphaned documentation assets. | source/asset inventory script or test |
| Q17.3 | not-started | Audit npm, release, and documentation package contents for required and forbidden files. | package dry run and archive listing |
| Q17.4 | not-started | Confirm generators are idempotent and do not reorder or rewrite unrelated files. | double-run clean-worktree check |
| Q17.5 | not-started | Verify timestamps, random IDs, map iteration, and platform paths do not make supported outputs unstable. | cross-run and cross-platform determinism checks |

### Q18 Configuration, Logging, and Observability

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q18.1 | not-started | Audit configuration defaults, file/environment/flag precedence, and invalid-value diagnostics. | config/controller tests |
| Q18.2 | not-started | Confirm logs and wrapped errors identify the failed stage without duplicating or obscuring the root cause. | error-chain assertions and command checks |
| Q18.3 | not-started | Verify normal commands keep stdout machine-usable and send diagnostics/progress to the intended stream. | CLI output tests |
| Q18.4 | not-started | Confirm logs and diagnostics do not expose imported source contents, credentials, tokens, or sensitive paths unnecessarily. | redaction/security tests |
| Q18.5 | not-started | Define optional timing/count observability for parser, layout, routing, scene, and encoder stages without changing default output. | design decision and focused tests if implemented |

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
2. UML diagrams: first semantic accuracy for each UML kind, then design
   polish for that kind, then shared UML visual-language and cross-format
   parity. Prioritize activity, class, sequence, state machine,
   component/deployment/package/composite-structure, then the remaining UML
   samples and strict profile diagnostics.
3. Data registry and imports: reusable table data, database schemas, SQL
   imports, relative paths, provenance, duplicate IDs, and missing resources.
4. Tables: standard tables, pipe tables, styling, cell text, headers, column
   sizing, empty cells, and imported data.
5. Relational databases: entities, columns, primary keys, foreign keys,
   composite keys, SQL imports, relation routing, and notation-specific labels.
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
12. Diagnostics and error UX: severity, aggregation, source positions,
    cancellation, and presentation in CLI/preview adapters.
13. Determinism, concurrency, and performance: stable ordering and bytes,
    race safety, cancellation, benchmarks, and bounded extreme inputs.
14. CLI, API, and preview contracts: options, defaults, output behavior,
    native/embedded parity, reload behavior, and backward compatibility.
15. Themes and accessibility: contrast, non-color semantics, typography,
    assistive SVG structure, and international text behavior.
16. Structural diff: semantic matching, classification, highlights,
    multi-frame pairing, and sample quality.
17. Reproducibility and artifact integrity: generated asset freshness,
    package contents, generator idempotence, and platform stability.
18. Configuration, logging, and observability: precedence, wrapped errors,
    output streams, redaction, and optional stage telemetry.

## First UML Quality Slice

After the current Q01 canonical-envelope follow-ups are closed or explicitly
paused, start the UML quality pass with
`docs/src/examples/samples/uml-activity.xal`. It is the smallest UML slice that
exercises both semantic precision and visible design quality, and should
verify:

- `initial`, `action`, `object-node`, `decision`, `fork`, `join`, `merge`, and
  `final` nodes.
- `control-flow` and `object-flow` rendering and visual distinction.
- `guard`, `title`, nested `responsibility`, and nested `constraint` behavior.
- Left-to-right activity readability.
- Decision diamonds, fork/join bars, merge nodes, final nodes, and connector
  labels avoiding collisions.
- SVG output quality sufficient for documentation, with additional format
  checks added when shared contracts or encoder behavior changes.

## Target Activity Diagram Capability

This section tracks the activity-diagram target and the remaining follow-up
work beyond the initial swimlane implementation slice.

The reference capability is an ATM withdrawal activity diagram with three
vertical swimlanes for actor/system responsibility, a visible lane-header row,
start and final nodes, rounded action nodes, object-processing nodes, decision
diamonds, guard labels, repeated PIN and amount loops, and orthogonal arrows
that cross lane boundaries without obscuring labels. The resulting xaligo
sample should be able to express the same diagram structure while using xaligo
brand colors rather than the reference image's yellow/orange palette.

Implemented syntax additions preserve existing `<activity-diagram>` inputs:

```xml
<activity-diagram direction="right" theme="xaligo" lanes="vertical">
  <partition id="customer" title="Customer">
    <initial id="start" />
    <action id="start-session" title="Start session" tone="primary" />
  </partition>
  <partition id="atm" title="ATM">
    <action id="request-pin" title="Request PIN" tone="primary" />
    <decision id="pin-valid" />
    <action id="process-pin" title="Process PIN" />
  </partition>
  <partition id="bank" title="Bank">
    <action id="approve-card" title="Approve card" />
  </partition>

  <control-flow src="start" dst="start-session" />
  <control-flow src="request-pin" dst="pin-valid" guard="PIN entered" />
  <control-flow src="pin-valid" dst="request-pin" guard="invalid PIN" />
</activity-diagram>
```

The parser accepts either nested nodes inside `<partition>` or a node-level
`lane="partition-id"` reference, but not both for the same node.
Every partition requires a stable `id` and non-empty `title`. A node may belong
to at most one partition. Flow endpoints remain normal activity node IDs, and
cross-partition flow is allowed.

The activity layout reserves the top header band for lane titles, divides the
content area into equal vertical lanes, and keeps each node inside its owning
lane. Remaining follow-up work includes explicit lane weights, specialized
backward-loop routing outside the forward path, object-flow accent styling,
and fork/join bar polish.

Use the xaligo readme logo as the activity theme source:

| Token | Color | Usage |
|---|---|---|
| `xaligo-cyan` | `#08b8ea` | primary lane headers, primary action fills, selected emphasis |
| `xaligo-navy` | `#052d6e` | text, borders, arrows, decision outlines, fork/join bars |
| `xaligo-teal` | `#04b79f` | success/object-flow accent and secondary emphasis |

Recommended derived fills are `#e8f7fd` for normal activity nodes,
`#e6fbf7` for object-flow or data-processing nodes, and `#f8fbff` for lane
backgrounds. Do not use the reference image's yellow/orange palette as the
default xaligo activity theme.

The first implementation slice should make
`docs/src/examples/targets/uml-activity-atm-swimlane.xal` validate and render,
then promote or mirror it into the validated sample corpus. Add focused
parser/layout tests for partitions and loop routing, SVG geometry assertions
for lane headers and label clearance, and a visual render check for the
generated sample. Cross-format checks become mandatory when the shared scene or
draw plan changes.

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
- Determinism, cancellation, resource limits, compatibility, accessibility,
  and artifact reproducibility have been considered and tested where relevant.
- Required security and repository checks have passed.
- The slice is committed without unrelated working-tree changes.