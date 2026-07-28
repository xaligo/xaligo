---
applyTo: "**"
---

# Instruction index

This index and `operations.instructions.md` are the only always-loaded files.
Search by task terms, then open only matching entries. Re-evaluate when scope
changes. Numbers are the canonical reading order.

## Task routes

| Task | Read in order |
|---|---|
| Any repository change | keyword-matched sections in 01–03 |
| Scope/status/planning | keyword-matched sections in 04–06 |
| `.xal` syntax or behavior | matching 07 section; add matching 08 section for implementation |
| Go/TypeScript structure | matching 08–09 sections |
| PPTX/connectors/routing | matching 07, 08, and 10 sections |
| Diagram authoring | matching 07 and 11 sections |

Preserve unrelated worktree changes. Implementation requests allow local
commits; never push, publish, tag, open a PR, or rewrite history without explicit
permission. Detail files use a nonmatching `applyTo` intentionally and are
loaded through this index.

## Directories

| Range | Directory | Contents |
|---|---|---|
| 01 | `01-general/` | Project-wide rules |
| 02 | `02-agent-guide/` | Repository orientation and commands |
| 03 | `03-development-flow/` | Change, verification, and commit workflow |
| 04 | `04-feature-catalog/` | Stable feature IDs and status |
| 05 | `05-issues-quality/` | Quality backlog and completion gates |
| 06 | `06-roadmap/` | Product direction and delivery phases |
| 07 | `07-xal-spec/` | Authoritative DSL specification |
| 08 | `08-architecture/` | Pipeline and package boundaries |
| 09 | `09-coding/` | Source organization and identifiers |
| 10 | `10-pptx-routing/` | PPTX, routing, and page fitting |
| 11 | `11-diagram-creation/` | Diagram authoring workflow |

## Numbered contents

Search this file directly, for example: `rg -i "metadata|page-link|WASM" .github/instructions/index.instructions.md`.

### AI operations

- [Development command workflow](operations.instructions.md) — keywords: AI, operation, development, commands, minimal-context, rg, sed, git

### 01 General

- 01.01 [Project](01-general/01-01-general-project.instructions.md) — keywords: repository, project, rules, module
- 01.02 [Directory structure](01-general/01-02-general-directory-structure.instructions.md) — keywords: repository, project, rules, Directory, structure, paths, packages
- 01.03 [Architecture rules](01-general/01-03-general-architecture-rules.instructions.md) — keywords: repository, project, rules, Architecture, design, dependencies
- 01.04 [Testing rules](01-general/01-04-general-testing-rules.instructions.md) — keywords: repository, project, rules, Testing, tests, validation
- 01.05 [Assets and configuration](01-general/01-05-general-assets-and-configuration.instructions.md) — keywords: repository, project, rules, Assets, configuration, catalog, YAML
- 01.06 [Conventions](01-general/01-06-general-conventions.instructions.md) — keywords: repository, project, rules, Conventions, naming, style
- 01.07 [Verification](01-general/01-07-general-verification.instructions.md) — keywords: repository, project, rules, Verification, tests, validation


### 02 Agent guide

- 02.00 [Overview](02-agent-guide/02-00-agent-guide-overview.instructions.md) — keywords: agent, commands, API
- 02.01 [Project summary](02-agent-guide/02-01-agent-guide-project-summary.instructions.md) — keywords: agent, commands, API, Project, summary, module, repository
- 02.02 [Working rules](02-agent-guide/02-02-agent-guide-working-rules.instructions.md) — keywords: agent, commands, API, Working
- 02.03 [Common commands](02-agent-guide/02-03-agent-guide-common-commands.instructions.md) — keywords: agent, commands, API, Common, CLI, command
- 02.04 [Shared Use-Case APIs](02-agent-guide/02-04-agent-guide-shared-use-case-apis.instructions.md) — keywords: agent, commands, API, Use-Case, APIs
- 02.05 [Asset workflow](02-agent-guide/02-05-agent-guide-asset-workflow.instructions.md) — keywords: agent, commands, API, Asset, workflow, assets, catalog
- 02.06 [Services CSV](02-agent-guide/02-06-agent-guide-services-csv.instructions.md) — keywords: agent, commands, API, Services, CSV, assets, catalog, service
- 02.07 [Completion checklist](02-agent-guide/02-07-agent-guide-completion-checklist.instructions.md) — keywords: agent, commands, API, Completion, checklist


### 03 Development flow

- 03.00 [Overview](03-development-flow/03-00-development-flow-overview.instructions.md) — keywords: workflow, git, verification
- 03.01 [Authorization and Scope](03-development-flow/03-01-development-flow-authorization-and-scope.instructions.md) — keywords: workflow, git, verification, Authorization, Scope, permissions, boundaries
- 03.02 [Before Changing Files](03-development-flow/03-02-development-flow-before-changing-files.instructions.md) — keywords: workflow, git, verification, Before, Changing, Files
- 03.03 [Change Slices](03-development-flow/03-03-development-flow-change-slices.instructions.md) — keywords: workflow, git, verification, Change, Slices
- 03.04 [Commit Cadence](03-development-flow/03-04-development-flow-commit-cadence.instructions.md) — keywords: workflow, git, verification, Commit, Cadence, staging, history
- 03.05 [Working-Tree Safety](03-development-flow/03-05-development-flow-working-tree-safety.instructions.md) — keywords: workflow, git, verification, Working-Tree, Safety, staging, history
- 03.06 [Verification Cadence](03-development-flow/03-06-development-flow-verification-cadence.instructions.md) — keywords: workflow, git, verification, Cadence, tests, validation
- 03.07 [Documentation and Architecture Diagrams](03-development-flow/03-07-development-flow-documentation-and-architecture-diagrams.instructions.md) — keywords: workflow, git, verification, Documentation, Architecture, Diagrams, design, dependencies
- 03.08 [Completion Audit](03-development-flow/03-08-development-flow-completion-audit.instructions.md) — keywords: workflow, git, verification, Completion, Audit


### 04 Feature catalog

- 04.00 [Overview](04-feature-catalog/04-00-feature-catalog-overview.instructions.md) — keywords: feature, status, XAL-ID
- 04.01 [Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)](04-feature-catalog/04-01-feature-catalog-group-1-core-dsl-document-envelope-xal-1xxxxxx.instructions.md) — keywords: feature, status, XAL-ID, Core, DSL, Document, Envelope, XAL-1xxxxxx
- 04.02 [Group 2 — Layout & Composition Primitives (`XAL-2xxxxxx`)](04-feature-catalog/04-02-feature-catalog-group-2-layout-composition-primitives-xal-2xxxxx.instructions.md) — keywords: feature, status, XAL-ID, Layout, Composition, Primitives, XAL-2xxxxxx, capability
- 04.03 [Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)](04-feature-catalog/04-03-feature-catalog-group-3-aws-architecture-primitives-icon-catalog.instructions.md) — keywords: feature, status, XAL-ID, AWS, Architecture, Primitives, Icon, Catalog
- 04.04 [Group 4 — Connections & Routing (`XAL-4xxxxxx`)](04-feature-catalog/04-04-feature-catalog-group-4-connections-routing-xal-4xxxxxx.instructions.md) — keywords: feature, status, XAL-ID, Connections, Routing, XAL-4xxxxxx, capability, implemented
- 04.05 [Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)](04-feature-catalog/04-05-feature-catalog-group-5-content-elements-tables-databases-uml-xa.instructions.md) — keywords: feature, status, XAL-ID, Content, Elements, Tables, Databases, UML
- 04.06 [Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)](04-feature-catalog/04-06-feature-catalog-group-6-output-formats-rendering-xal-6xxxxxx.instructions.md) — keywords: feature, status, XAL-ID, Output, Formats, Rendering, XAL-6xxxxxx, capability
- 04.07 [Group 7 — Presentation, Theming & Physical Page Fitting (`XAL-7xxxxxx`)](04-feature-catalog/04-07-feature-catalog-group-7-presentation-theming-physical-page-fitti.instructions.md) — keywords: feature, status, XAL-ID, Presentation, Theming, Physical, Page, Fitting
- 04.08 [Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)](04-feature-catalog/04-08-feature-catalog-group-8-cli-commands-developer-tooling-xal-8xxxx.instructions.md) — keywords: feature, status, XAL-ID, CLI, Commands, Developer, Tooling, XAL-8xxxxxx
- 04.09 [Group 9 — Validation, Diagnostics, Diff & External Integration (`XAL-9xxxxxx`)](04-feature-catalog/04-09-feature-catalog-group-9-validation-diagnostics-diff-external-int.instructions.md) — keywords: feature, status, XAL-ID, Validation, Diagnostics, Diff, External, Integration


### 05 Issues and quality

- 05.00 [Overview](05-issues-quality/05-00-issues-and-quality-overview.instructions.md) — keywords: quality, backlog, acceptance
- 05.01 [Quality Rubric](05-issues-quality/05-01-issues-and-quality-quality-rubric.instructions.md) — keywords: quality, backlog, acceptance, Rubric, tests, validation
- 05.02 [Feature Quality Workflow](05-issues-quality/05-02-issues-and-quality-feature-quality-workflow.instructions.md) — keywords: quality, backlog, acceptance, Feature, Workflow, tests, validation
- 05.03 [Status Tracking](05-issues-quality/05-03-issues-and-quality-status-tracking.instructions.md) — keywords: quality, backlog, acceptance, Status, Tracking
- 05.04 [Detailed Task Backlog](05-issues-quality/05-04-issues-and-quality-detailed-task-backlog.instructions.md) — keywords: quality, backlog, acceptance, Task, issue, tracking
  - 05.04.01 [Q01 Canonical V1 Document Envelope](05-issues-quality/05-04-01-issues-and-quality-q01-canonical-v1-document-envelope.instructions.md) — keywords: quality, backlog, acceptance, Q01, Canonical, Document, Envelope, issue
  - 05.04.02 [Q02 Data Registry and Imports](05-issues-quality/05-04-02-issues-and-quality-q02-data-registry-and-imports.instructions.md) — keywords: quality, backlog, acceptance, Q02, Data, Registry, Imports, issue
  - 05.04.03 [Q03 Tables](05-issues-quality/05-04-03-issues-and-quality-q03-tables.instructions.md) — keywords: quality, backlog, acceptance, Q03, Tables, issue, tracking, table
  - 05.04.04 [Q04 Relational Databases](05-issues-quality/05-04-04-issues-and-quality-q04-relational-databases.instructions.md) — keywords: quality, backlog, acceptance, Q04, Relational, Databases, issue, tracking
  - 05.04.05 [Q05 UML Diagrams](05-issues-quality/05-04-05-issues-and-quality-q05-uml-diagrams.instructions.md) — keywords: quality, backlog, acceptance, Q05, UML, Diagrams, issue, tracking
  - 05.04.06 [Q06 Frame-Scoped References](05-issues-quality/05-04-06-issues-and-quality-q06-frame-scoped-references.instructions.md) — keywords: quality, backlog, acceptance, Q06, Frame-Scoped, References, permissions, boundaries
  - 05.04.07 [Q07 Cross-Frame Page Links](05-issues-quality/05-04-07-issues-and-quality-q07-cross-frame-page-links.instructions.md) — keywords: quality, backlog, acceptance, Q07, Cross-Frame, Page, Links, issue
  - 05.04.08 [Q08 Frame Metadata Bands](05-issues-quality/05-04-08-issues-and-quality-q08-frame-metadata-bands.instructions.md) — keywords: quality, backlog, acceptance, Q08, Frame, Metadata, Bands, issue
  - 05.04.09 [Q09 Page-Oriented Outputs](05-issues-quality/05-04-09-issues-and-quality-q09-page-oriented-outputs.instructions.md) — keywords: quality, backlog, acceptance, Q09, Page-Oriented, Outputs, issue, tracking
  - 05.04.10 [Q10 Renderer Matrix](05-issues-quality/05-04-10-issues-and-quality-q10-renderer-matrix.instructions.md) — keywords: quality, backlog, acceptance, Q10, Renderer, Matrix, issue, tracking
  - 05.04.11 [Q11 CI, Release, and Tooling](05-issues-quality/05-04-11-issues-and-quality-q11-ci-release-and-tooling.instructions.md) — keywords: quality, backlog, acceptance, Q11, Release, Tooling, issue, tracking
  - 05.04.12 [Q12 Diagnostics and Error UX](05-issues-quality/05-04-12-issues-and-quality-q12-diagnostics-and-error-ux.instructions.md) — keywords: quality, backlog, acceptance, Q12, Diagnostics, Error, issue, tracking
  - 05.04.13 [Q13 Determinism, Concurrency, and Performance](05-issues-quality/05-04-13-issues-and-quality-q13-determinism-concurrency-and-performance.instructions.md) — keywords: quality, backlog, acceptance, Q13, Determinism, Concurrency, Performance, issue
  - 05.04.14 [Q14 CLI, API, and Preview Contracts](05-issues-quality/05-04-14-issues-and-quality-q14-cli-api-and-preview-contracts.instructions.md) — keywords: quality, backlog, acceptance, Q14, CLI, API, Preview, Contracts
  - 05.04.15 [Q15 Themes and Accessibility](05-issues-quality/05-04-15-issues-and-quality-q15-themes-and-accessibility.instructions.md) — keywords: quality, backlog, acceptance, Q15, Themes, Accessibility, issue, tracking
  - 05.04.16 [Q16 Structural Diff](05-issues-quality/05-04-16-issues-and-quality-q16-structural-diff.instructions.md) — keywords: quality, backlog, acceptance, Q16, Structural, Diff, issue, tracking
  - 05.04.17 [Q17 Reproducibility and Artifact Integrity](05-issues-quality/05-04-17-issues-and-quality-q17-reproducibility-and-artifact-integrity.instructions.md) — keywords: quality, backlog, acceptance, Q17, Reproducibility, Artifact, Integrity, issue
  - 05.04.18 [Q18 Configuration, Logging, and Observability](05-issues-quality/05-04-18-issues-and-quality-q18-configuration-logging-and-observability.instructions.md) — keywords: quality, backlog, acceptance, Q18, Configuration, Logging, Observability, YAML
- 05.05 [Design Quality Gate](05-issues-quality/05-05-issues-and-quality-design-quality-gate.instructions.md) — keywords: quality, backlog, acceptance, Design, Gate, tests, validation
- 05.06 [Phase 2 Quality Order](05-issues-quality/05-06-issues-and-quality-phase-2-quality-order.instructions.md) — keywords: quality, backlog, acceptance, Phase, Order, tests, validation, milestone
- 05.07 [First UML Quality Slice](05-issues-quality/05-07-issues-and-quality-first-uml-quality-slice.instructions.md) — keywords: quality, backlog, acceptance, First, UML, Slice, tests, validation
- 05.08 [Definition of Done](05-issues-quality/05-08-issues-and-quality-definition-of-done.instructions.md) — keywords: quality, backlog, acceptance, Definition, Done


### 06 Roadmap

- 06.00 [Overview](06-roadmap/06-00-roadmap-overview.instructions.md) — keywords: roadmap, planning, release
- 06.01 [Product Architecture Preconditions](06-roadmap/06-01-roadmap-product-architecture-preconditions.instructions.md) — keywords: roadmap, planning, release, Product, Architecture, design, dependencies
  - 06.01.01 [V1 Structured-Diagram Profile](06-roadmap/06-01-01-roadmap-v1-structured-diagram-profile.instructions.md) — keywords: roadmap, planning, release, Structured-Diagram, Profile, milestone, version
  - 06.01.02 [Common DSL and Go Core](06-roadmap/06-01-02-roadmap-common-dsl-and-go-core.instructions.md) — keywords: roadmap, planning, release, Common, DSL, Core
  - 06.01.03 [V1 Compatibility and V2 Input](06-roadmap/06-01-03-roadmap-v1-compatibility-and-v2-input.instructions.md) — keywords: roadmap, planning, release, Compatibility, Input, milestone, version
  - 06.01.04 [Mode and Format Are Independent](06-roadmap/06-01-04-roadmap-mode-and-format-are-independent.instructions.md) — keywords: roadmap, planning, release, Mode, Format, Independent
  - 06.01.05 [Shared Rendering APIs](06-roadmap/06-01-05-roadmap-shared-rendering-apis.instructions.md) — keywords: roadmap, planning, release, Rendering, APIs
  - 06.01.06 [Rendering Correctness Gate](06-roadmap/06-01-06-roadmap-rendering-correctness-gate.instructions.md) — keywords: roadmap, planning, release, Rendering, Correctness, Gate
- 06.02 [Delivery Phases](06-roadmap/06-02-roadmap-delivery-phases.instructions.md) — keywords: roadmap, planning, release, Delivery, Phases, milestone, version
  - 06.02.01 [Phase 1: Basic Output](06-roadmap/06-02-01-roadmap-phase-1-basic-output.instructions.md) — keywords: roadmap, planning, release, Phase, Basic, Output, milestone, version
  - 06.02.02 [Phase 2: Network Diagram Features](06-roadmap/06-02-02-roadmap-phase-2-network-diagram-features.instructions.md) — keywords: roadmap, planning, release, Phase, Network, Diagram, milestone, version
  - 06.02.03 [Phase 3: Live Preview](06-roadmap/06-02-03-roadmap-phase-3-live-preview.instructions.md) — keywords: roadmap, planning, release, Phase, Live, Preview, milestone, version
- 06.03 [VS Code Extension Preconditions](06-roadmap/06-03-roadmap-vs-code-extension-preconditions.instructions.md) — keywords: roadmap, planning, release, Code, Extension
- 06.04 [AWS 2.5D Mode](06-roadmap/06-04-roadmap-aws-2-5d-mode.instructions.md) — keywords: roadmap, planning, release, AWS, 2.5D, Mode, cloud
- 06.05 [Export Roadmap](06-roadmap/06-05-roadmap-export-roadmap.instructions.md) — keywords: roadmap, planning, release, Export, milestone, version
- 06.06 [Long-Term Product Position](06-roadmap/06-06-roadmap-long-term-product-position.instructions.md) — keywords: roadmap, planning, release, Long-Term, Product, Position
- 06.07 [Current State](06-roadmap/06-07-roadmap-current-state.instructions.md) — keywords: roadmap, planning, release, State
- 06.08 [Rebaselined Implementation Order](06-roadmap/06-08-roadmap-rebaselined-implementation-order.instructions.md) — keywords: roadmap, planning, release, Rebaselined, Implementation, Order
- 06.09 [v0.1 Foundation](06-roadmap/06-09-roadmap-v0-1-foundation.instructions.md) — keywords: roadmap, planning, release, v0.1, Foundation, module, repository, milestone
  - 06.09.01 [Rendering Engine Refactoring](06-roadmap/06-09-01-roadmap-rendering-engine-refactoring.instructions.md) — keywords: roadmap, planning, release, Rendering, Engine, Refactoring
  - 06.09.02 [Public API](06-roadmap/06-09-02-roadmap-public-api.instructions.md) — keywords: roadmap, planning, release, Public, API
  - 06.09.03 [CLI](06-roadmap/06-09-03-roadmap-cli.instructions.md) — keywords: roadmap, planning, release, CLI, command
- 06.10 [v0.2 SVG Renderer](06-roadmap/06-10-roadmap-v0-2-svg-renderer.instructions.md) — keywords: roadmap, planning, release, v0.2, SVG, Renderer, milestone, version
  - 06.10.01 [SVG Export](06-roadmap/06-10-01-roadmap-svg-export.instructions.md) — keywords: roadmap, planning, release, SVG, Export, renderer
  - 06.10.02 [Supported Elements](06-roadmap/06-10-02-roadmap-supported-elements.instructions.md) — keywords: roadmap, planning, release, Supported, Elements
  - 06.10.03 [Themes](06-roadmap/06-10-03-roadmap-themes.instructions.md) — keywords: roadmap, planning, release, Themes, theme, colors
- 06.11 [v0.3 Network Diagram Features](06-roadmap/06-11-roadmap-v0-3-network-diagram-features.instructions.md) — keywords: roadmap, planning, release, v0.3, Network, Diagram, milestone, version
  - 06.11.01 [Route Connector](06-roadmap/06-11-01-roadmap-route-connector.instructions.md) — keywords: roadmap, planning, release, Route, Connector, connection, routing
  - 06.11.02 [Connector Model](06-roadmap/06-11-02-roadmap-connector-model.instructions.md) — keywords: roadmap, planning, release, Connector, Model, connection, routing
  - 06.11.03 [Orthogonal Routing](06-roadmap/06-11-03-roadmap-orthogonal-routing.instructions.md) — keywords: roadmap, planning, release, Orthogonal, Routing, connection
  - 06.11.04 [Route / Traffic Separation](06-roadmap/06-11-04-roadmap-route-traffic-separation.instructions.md) — keywords: roadmap, planning, release, Route, Traffic, Separation, connection, routing
  - 06.11.05 [DSL](06-roadmap/06-11-05-roadmap-dsl.instructions.md) — keywords: roadmap, planning, release, DSL
- 06.12 [v0.4 Advanced Routing](06-roadmap/06-12-roadmap-v0-4-advanced-routing.instructions.md) — keywords: roadmap, planning, release, v0.4, Advanced, Routing, milestone, version
  - 06.12.01 [Edge Offset](06-roadmap/06-12-01-roadmap-edge-offset.instructions.md) — keywords: roadmap, planning, release, Edge, Offset
  - 06.12.02 [Layer Routing](06-roadmap/06-12-02-roadmap-layer-routing.instructions.md) — keywords: roadmap, planning, release, Layer, Routing, connection
  - 06.12.03 [Junction Generation](06-roadmap/06-12-03-roadmap-junction-generation.instructions.md) — keywords: roadmap, planning, release, Junction, Generation
- 06.13 [v0.5 Line Jumps](06-roadmap/06-13-roadmap-v0-5-line-jumps.instructions.md) — keywords: roadmap, planning, release, v0.5, Line, Jumps, milestone, version
  - 06.13.01 [Bridge / Jump Lines](06-roadmap/06-13-01-roadmap-bridge-jump-lines.instructions.md) — keywords: roadmap, planning, release, Bridge, Jump, Lines
  - 06.13.02 [Features](06-roadmap/06-13-02-roadmap-features.instructions.md) — keywords: roadmap, planning, release
- 06.14 [v0.6 Live Preview](06-roadmap/06-14-roadmap-v0-6-live-preview.instructions.md) — keywords: roadmap, planning, release, v0.6, Live, Preview, milestone, version
  - 06.14.01 [xaligo serve](06-roadmap/06-14-01-roadmap-xaligo-serve.instructions.md) — keywords: roadmap, planning, release, xaligo, serve, preview, SSE
  - 06.14.02 [Features](06-roadmap/06-14-02-roadmap-features.instructions.md) — keywords: roadmap, planning, release
  - 06.14.03 [Backend Stack](06-roadmap/06-14-03-roadmap-backend-stack.instructions.md) — keywords: roadmap, planning, release, Backend, Stack
  - 06.14.04 [Frontend Stack](06-roadmap/06-14-04-roadmap-frontend-stack.instructions.md) — keywords: roadmap, planning, release, Frontend, Stack
  - 06.14.05 [Preview Flow](06-roadmap/06-14-05-roadmap-preview-flow.instructions.md) — keywords: roadmap, planning, release, Preview, Flow, SSE
- 06.15 [v0.7 VS Code Extension](06-roadmap/06-15-roadmap-v0-7-vs-code-extension.instructions.md) — keywords: roadmap, planning, release, v0.7, Code, Extension, milestone, version
  - 06.15.01 [Language Support](06-roadmap/06-15-01-roadmap-language-support.instructions.md) — keywords: roadmap, planning, release, Language, Support
  - 06.15.02 [Features](06-roadmap/06-15-02-roadmap-features.instructions.md) — keywords: roadmap, planning, release
  - 06.15.03 [Preview Panel](06-roadmap/06-15-03-roadmap-preview-panel.instructions.md) — keywords: roadmap, planning, release, Preview, Panel, SSE
  - 06.15.04 [Live Preview](06-roadmap/06-15-04-roadmap-live-preview.instructions.md) — keywords: roadmap, planning, release, Live, Preview, SSE
- 06.16 [v0.8 Excalidraw Integration](06-roadmap/06-16-roadmap-v0-8-excalidraw-integration.instructions.md) — keywords: roadmap, planning, release, v0.8, Excalidraw, Integration, milestone, version
  - 06.16.01 [Excalidraw Preview](06-roadmap/06-16-01-roadmap-excalidraw-preview.instructions.md) — keywords: roadmap, planning, release, Excalidraw, Preview, SSE, scene
  - 06.16.02 [Features](06-roadmap/06-16-02-roadmap-features.instructions.md) — keywords: roadmap, planning, release
  - 06.16.03 [Excalidraw Export](06-roadmap/06-16-03-roadmap-excalidraw-export.instructions.md) — keywords: roadmap, planning, release, Excalidraw, Export, scene
- 06.17 [v0.9 PowerPoint Export](06-roadmap/06-17-roadmap-v0-9-powerpoint-export.instructions.md) — keywords: roadmap, planning, release, v0.9, PowerPoint, Export, milestone, version
  - 06.17.01 [PPTX Export](06-roadmap/06-17-01-roadmap-pptx-export.instructions.md) — keywords: roadmap, planning, release, PPTX, Export, PowerPoint
  - 06.17.02 [Supported Features](06-roadmap/06-17-02-roadmap-supported-features.instructions.md) — keywords: roadmap, planning, release, Supported
- 06.18 [v1.0](06-roadmap/06-18-roadmap-v1-0.instructions.md) — keywords: roadmap, planning, release, v1.0, milestone, version
  - 06.18.01 [VS Code Marketplace Release](06-roadmap/06-18-01-roadmap-vs-code-marketplace-release.instructions.md) — keywords: roadmap, planning, release, Code, Marketplace, CI
- 06.19 [Future Vision](06-roadmap/06-19-roadmap-future-vision.instructions.md) — keywords: roadmap, planning, release, Future, Vision
  - 06.19.01 [AWS Architecture Mode](06-roadmap/06-19-01-roadmap-aws-architecture-mode.instructions.md) — keywords: roadmap, planning, release, AWS, Architecture, Mode, design, dependencies
  - 06.19.02 [Network Diagram Mode](06-roadmap/06-19-02-roadmap-network-diagram-mode.instructions.md) — keywords: roadmap, planning, release, Network, Diagram, Mode
  - 06.19.03 [Infrastructure as Diagram](06-roadmap/06-19-03-roadmap-infrastructure-as-diagram.instructions.md) — keywords: roadmap, planning, release, Infrastructure, Diagram
- 06.20 [Project Goal](06-roadmap/06-20-roadmap-project-goal.instructions.md) — keywords: roadmap, planning, release, Project, Goal, module, repository


### 07 XAL specification

- 07.01 [Overview](07-xal-spec/07-01-xal-specification-overview.instructions.md) — keywords: XAL, DSL, XML
- 07.02 [V1 Compatibility Profile and Version Boundary](07-xal-spec/07-02-xal-specification-v1-compatibility-profile-and-version-boundary.instructions.md) — keywords: XAL, DSL, XML, Compatibility, Profile, Version, Boundary, milestone
- 07.03 [Root Tag](07-xal-spec/07-03-xal-specification-root-tag.instructions.md) — keywords: XAL, DSL, XML, Root, Tag, version
  - 07.03.01 [Frame and physical-page contract](07-xal-spec/07-03-01-xal-specification-frame-and-physical-page-contract.instructions.md) — keywords: XAL, DSL, XML, Frame, physical-page, contract, page
  - 07.03.02 [Frame metadata tag band](07-xal-spec/07-03-02-xal-specification-frame-metadata-tag-band.instructions.md) — keywords: XAL, DSL, XML, Frame, metadata, tag, band, page
- 07.04 [Numeric and Geometry Contract](07-xal-spec/07-04-xal-specification-numeric-and-geometry-contract.instructions.md) — keywords: XAL, DSL, XML, Numeric, Geometry, Contract, layout
  - 07.04.01 [Fixed and flexible child allocation](07-xal-spec/07-04-01-xal-specification-fixed-and-flexible-child-allocation.instructions.md) — keywords: XAL, DSL, XML, Fixed, flexible, child, allocation
- 07.05 [Layout Tags](07-xal-spec/07-05-xal-specification-layout-tags.instructions.md) — keywords: XAL, DSL, XML, Layout, Tags, geometry
  - 07.05.01 [`<container>`](07-xal-spec/07-05-01-xal-specification-container.instructions.md) — keywords: XAL, DSL, XML, container
  - 07.05.02 [`<row>`](07-xal-spec/07-05-02-xal-specification-row.instructions.md) — keywords: XAL, DSL, XML, row
  - 07.05.03 [`<col>`](07-xal-spec/07-05-03-xal-specification-col.instructions.md) — keywords: XAL, DSL, XML, col
- 07.06 [Custom Leaf and Container Tags](07-xal-spec/07-06-xal-specification-custom-leaf-and-container-tags.instructions.md) — keywords: XAL, DSL, XML, Custom, Leaf, Container, Tags
- 07.07 [`<rectangle>` and `<port>` Tags](07-xal-spec/07-07-xal-specification-rectangle-and-port-tags.instructions.md) — keywords: XAL, DSL, XML, rectangle, port, Tags
- 07.08 [Resolved Text Layout](07-xal-spec/07-08-xal-specification-resolved-text-layout.instructions.md) — keywords: XAL, DSL, XML, Resolved, Text, Layout, geometry
- 07.09 [`<table>` Tag](07-xal-spec/07-09-xal-specification-table-tag.instructions.md) — keywords: XAL, DSL, XML, table, Tag, cells
- 07.10 [Relational Database Tags](07-xal-spec/07-10-xal-specification-relational-database-tags.instructions.md) — keywords: XAL, DSL, XML, Relational, Database, Tags, schema, UML
- 07.11 [UML Tags](07-xal-spec/07-11-xal-specification-uml-tags.instructions.md) — keywords: XAL, DSL, XML, UML, Tags, relation
  - 07.11.01 [Component, identity, and layout contract](07-xal-spec/07-11-01-xal-specification-component-identity-and-layout-contract.instructions.md) — keywords: XAL, DSL, XML, Component, identity, layout, contract, geometry
  - 07.11.02 [Diagram-kind vocabulary](07-xal-spec/07-11-02-xal-specification-diagram-kind-vocabulary.instructions.md) — keywords: XAL, DSL, XML, Diagram-kind, vocabulary
  - 07.11.03 [Component diagram sizing](07-xal-spec/07-11-03-xal-specification-component-diagram-sizing.instructions.md) — keywords: XAL, DSL, XML, Component, diagram, sizing, layout, geometry
  - 07.11.04 [Activity partitions](07-xal-spec/07-11-04-xal-specification-activity-partitions.instructions.md) — keywords: XAL, DSL, XML, Activity, partitions, UML, relation
  - 07.11.05 [Ownership](07-xal-spec/07-11-05-xal-specification-ownership.instructions.md) — keywords: XAL, DSL, XML, Ownership
  - 07.11.06 [Element compartments](07-xal-spec/07-11-06-xal-specification-element-compartments.instructions.md) — keywords: XAL, DSL, XML, Element, compartments
  - 07.11.07 [Relation attributes, order, and time](07-xal-spec/07-11-07-xal-specification-relation-attributes-order-and-time.instructions.md) — keywords: XAL, DSL, XML, Relation, attributes, order, time, UML
  - 07.11.08 [Relation projection](07-xal-spec/07-11-08-xal-specification-relation-projection.instructions.md) — keywords: XAL, DSL, XML, Relation, projection, module, repository, UML
  - 07.11.09 [Reusable UML models](07-xal-spec/07-11-09-xal-specification-reusable-uml-models.instructions.md) — keywords: XAL, DSL, XML, Reusable, UML, models, relation
  - 07.11.10 [Deliberately lossy V1 projection](07-xal-spec/07-11-10-xal-specification-deliberately-lossy-v1-projection.instructions.md) — keywords: XAL, DSL, XML, Deliberately, lossy, projection, module, repository
- 07.12 [`<item>` Tag](07-xal-spec/07-12-xal-specification-item-tag.instructions.md) — keywords: XAL, DSL, XML, item, Tag
- 07.13 [`<spacer>` / `<blank>` Tags](07-xal-spec/07-13-xal-specification-spacer-blank-tags.instructions.md) — keywords: XAL, DSL, XML, spacer, blank, Tags
- 07.14 [`<connection>` Tag](07-xal-spec/07-14-xal-specification-connection-tag.instructions.md) — keywords: XAL, DSL, XML, connection, Tag, routing
  - 07.14.01a [`<connections>` Tag, part 1/4](07-xal-spec/07-14-01-xal-specification-connections-tag-part-01.instructions.md) — keywords: XAL, DSL, XML, connections, Tag, part, connection, routing
  - 07.14.01b [`<connections>` Tag, part 2/4](07-xal-spec/07-14-01-xal-specification-connections-tag-part-02.instructions.md) — keywords: XAL, DSL, XML, connections, Tag, part, connection, routing
  - 07.14.01c [`<connections>` Tag, part 3/4](07-xal-spec/07-14-01-xal-specification-connections-tag-part-03.instructions.md) — keywords: XAL, DSL, XML, connections, Tag, part, connection, routing
  - 07.14.01d [`<connections>` Tag, part 4/4](07-xal-spec/07-14-01-xal-specification-connections-tag-part-04.instructions.md) — keywords: XAL, DSL, XML, connections, Tag, part, connection, routing
- 07.15 [AWS Group Tags](07-xal-spec/07-15-xal-specification-aws-group-tags.instructions.md) — keywords: XAL, DSL, XML, AWS, Tags, capability, implemented, cloud
  - 07.15.01 [Layout Control Attributes (shared by all containers)](07-xal-spec/07-15-01-xal-specification-layout-control-attributes-shared-by-all-containe.instructions.md) — keywords: XAL, DSL, XML, Layout, Control, Attributes, all, containers
  - 07.15.02 [Child Size Ratio Attributes](07-xal-spec/07-15-02-xal-specification-child-size-ratio-attributes.instructions.md) — keywords: XAL, DSL, XML, Child, Size, Ratio, Attributes
- 07.16 [Spacing Classes (`class` attribute)](07-xal-spec/07-16-xal-specification-spacing-classes-class-attribute.instructions.md) — keywords: XAL, DSL, XML, Spacing, Classes, class, attribute, CI
  - 07.16.01 [All-sides shorthand](07-xal-spec/07-16-01-xal-specification-all-sides-shorthand.instructions.md) — keywords: XAL, DSL, XML, All-sides, shorthand
  - 07.16.02 [Axis shorthand](07-xal-spec/07-16-02-xal-specification-axis-shorthand.instructions.md) — keywords: XAL, DSL, XML, Axis, shorthand
  - 07.16.03 [Per-side](07-xal-spec/07-16-03-xal-specification-per-side.instructions.md) — keywords: XAL, DSL, XML, Per-side
  - 07.16.04 [Semantics](07-xal-spec/07-16-04-xal-specification-semantics.instructions.md) — keywords: XAL, DSL, XML, Semantics
- 07.17 [Layout Calculation Rules](07-xal-spec/07-17-xal-specification-layout-calculation-rules.instructions.md) — keywords: XAL, DSL, XML, Layout, Calculation, geometry
- 07.18 [Example](07-xal-spec/07-18-xal-specification-example.instructions.md) — keywords: XAL, DSL, XML, Example
- 07.19 [Constraints and Notes](07-xal-spec/07-19-xal-specification-constraints-and-notes.instructions.md) — keywords: XAL, DSL, XML, Constraints, Notes


### 08 Architecture

- 08.00 [Overview](08-architecture/08-00-architecture-overview.instructions.md) — keywords: architecture, pipeline, layers
- 08.01 [Core pipeline](08-architecture/08-01-architecture-core-pipeline.instructions.md) — keywords: architecture, pipeline, layers, Core, design, dependencies
- 08.02 [Language-version boundary](08-architecture/08-02-architecture-language-version-boundary.instructions.md) — keywords: architecture, pipeline, layers, Language-version, boundary
- 08.03 [Package responsibilities](08-architecture/08-03-architecture-package-responsibilities.instructions.md) — keywords: architecture, pipeline, layers, Package, responsibilities
- 08.04 [Invariants](08-architecture/08-04-architecture-invariants.instructions.md) — keywords: architecture, pipeline, layers, Invariants
- 08.05 [File organization](08-architecture/08-05-architecture-file-organization.instructions.md) — keywords: architecture, pipeline, layers, File, organization, paths, packages
- 08.06 [Geometry contract](08-architecture/08-06-architecture-geometry-contract.instructions.md) — keywords: architecture, pipeline, layers, Geometry, contract, layout
- 08.07 [Renderer-neutral text contract](08-architecture/08-07-architecture-renderer-neutral-text-contract.instructions.md) — keywords: architecture, pipeline, layers, Renderer-neutral, text, contract
- 08.08 [Dependency direction](08-architecture/08-08-architecture-dependency-direction.instructions.md) — keywords: architecture, pipeline, layers, Dependency, direction
- 08.09 [Verification](08-architecture/08-09-architecture-verification.instructions.md) — keywords: architecture, pipeline, layers, Verification, tests, validation


### 09 Coding

- 09.00 [Overview](09-coding/09-00-coding-overview.instructions.md) — keywords: coding, Go, TypeScript
- 09.01 [Responsibility-based files](09-coding/09-01-coding-responsibility-based-files.instructions.md) — keywords: coding, Go, TypeScript, Responsibility-based, files
  - 09.01.01 [`internal/usecase` root contract](09-coding/09-01-01-coding-internal-usecase-root-contract.instructions.md) — keywords: coding, Go, TypeScript, internal, usecase, root, contract, version
- 09.02 [V1 engine identifiers](09-coding/09-02-coding-v1-engine-identifiers.instructions.md) — keywords: coding, Go, TypeScript, engine, identifiers, naming, style, milestone
- 09.03 [Engine execution boundary](09-coding/09-03-coding-engine-execution-boundary.instructions.md) — keywords: coding, Go, TypeScript, Engine, execution, boundary
- 09.04 [Dependency direction](09-coding/09-04-coding-dependency-direction.instructions.md) — keywords: coding, Go, TypeScript, Dependency, direction
- 09.05 [Verification](09-coding/09-05-coding-verification.instructions.md) — keywords: coding, Go, TypeScript, Verification, tests, validation


### 10 PPTX and routing

- 10.00 [Overview](10-pptx-routing/10-00-pptx-and-routing-overview.instructions.md) — keywords: PPTX, routing, connector
- 10.01 [Brainstorm Reference](10-pptx-routing/10-01-pptx-and-routing-brainstorm-reference.instructions.md) — keywords: PPTX, routing, connector, Brainstorm, Reference
- 10.02 [Confirmed Decisions](10-pptx-routing/10-02-pptx-and-routing-confirmed-decisions.instructions.md) — keywords: PPTX, routing, connector, Confirmed, Decisions, CI, release
- 10.03 [Current Pipeline](10-pptx-routing/10-03-pptx-and-routing-current-pipeline.instructions.md) — keywords: PPTX, routing, connector, Pipeline, design, dependencies
- 10.04 [Go / WASM Boundary](10-pptx-routing/10-04-pptx-and-routing-go-wasm-boundary.instructions.md) — keywords: PPTX, routing, connector, WASM, Boundary, PptxGenJS
- 10.05 [Ownership](10-pptx-routing/10-05-pptx-and-routing-ownership.instructions.md) — keywords: PPTX, routing, connector, Ownership
- 10.06 [Paper / Scaling](10-pptx-routing/10-06-pptx-and-routing-paper-scaling.instructions.md) — keywords: PPTX, routing, connector, Paper, Scaling, PPI
- 10.07 [Routing Rules](10-pptx-routing/10-07-pptx-and-routing-routing-rules.instructions.md) — keywords: PPTX, routing, connector, connection
  - 10.07.01 [Cross-frame page links](10-pptx-routing/10-07-01-pptx-and-routing-cross-frame-page-links.instructions.md) — keywords: PPTX, routing, connector, Cross-frame, page, links, frame
- 10.08 [Advanced Routing Features](10-pptx-routing/10-08-pptx-and-routing-advanced-routing-features.instructions.md) — keywords: PPTX, routing, connector, Advanced, connection
  - 10.08.01 [Line Jumps](10-pptx-routing/10-08-01-pptx-and-routing-line-jumps.instructions.md) — keywords: PPTX, routing, connector, Line, Jumps
  - 10.08.02 [Route / Traffic Separation](10-pptx-routing/10-08-02-pptx-and-routing-route-traffic-separation.instructions.md) — keywords: PPTX, routing, connector, Route, Traffic, Separation, connection
  - 10.08.03 [Route Connectors](10-pptx-routing/10-08-03-pptx-and-routing-route-connectors.instructions.md) — keywords: PPTX, routing, connector, Route, Connectors, connection
- 10.09 [Connector Style Options](10-pptx-routing/10-09-pptx-and-routing-connector-style-options.instructions.md) — keywords: PPTX, routing, connector, Style, Options, connection
- 10.10 [Group Header Tags](10-pptx-routing/10-10-pptx-and-routing-group-header-tags.instructions.md) — keywords: PPTX, routing, connector, Header, Tags, capability, implemented
- 10.11 [Item Labels](10-pptx-routing/10-11-pptx-and-routing-item-labels.instructions.md) — keywords: PPTX, routing, connector, Item, Labels
- 10.12 [Layout / Whitespace](10-pptx-routing/10-12-pptx-and-routing-layout-whitespace.instructions.md) — keywords: PPTX, routing, connector, Layout, Whitespace, geometry
- 10.13 [Legend Pages](10-pptx-routing/10-13-pptx-and-routing-legend-pages.instructions.md) — keywords: PPTX, routing, connector, Legend, Pages, frame, page
- 10.14 [Verification Checklist](10-pptx-routing/10-14-pptx-and-routing-verification-checklist.instructions.md) — keywords: PPTX, routing, connector, Verification, Checklist, tests, validation


### 11 Diagram creation

- 11.00 [Overview](11-diagram-creation/11-00-diagram-creation-overview.instructions.md) — keywords: diagram, authoring, services
- 11.01 [Step 1 — Find Service IDs](11-diagram-creation/11-01-diagram-creation-step-1-find-service-ids.instructions.md) — keywords: diagram, authoring, services, Find, Service, IDs, assets, catalog
- 11.02 [Step 2 — Create services.csv](11-diagram-creation/11-02-diagram-creation-step-2-create-services-csv.instructions.md) — keywords: diagram, authoring, services, Create, services.csv, assets, catalog, service
- 11.03 [Step 3 — Create a .xal file](11-diagram-creation/11-03-diagram-creation-step-3-create-a-xal-file.instructions.md) — keywords: diagram, authoring, services, Create, .xal, file
  - 11.03.01 [Choosing the right group tag](11-diagram-creation/11-03-01-diagram-creation-choosing-the-right-group-tag.instructions.md) — keywords: diagram, authoring, services, Choosing, right, tag, capability, implemented
  - 11.03.02 [Service Scope Validation](11-diagram-creation/11-03-02-diagram-creation-service-scope-validation.instructions.md) — keywords: diagram, authoring, services, Service, Scope, Validation, assets, catalog
- 11.04 [Step 4 — Render the Excalidraw file](11-diagram-creation/11-04-diagram-creation-step-4-render-the-excalidraw-file.instructions.md) — keywords: diagram, authoring, services, Render, Excalidraw, file, scene
- 11.05 [Command Reference](11-diagram-creation/11-05-diagram-creation-command-reference.instructions.md) — keywords: diagram, authoring, services, Command, Reference, CLI
- 11.06 [PPTX Notes](11-diagram-creation/11-06-diagram-creation-pptx-notes.instructions.md) — keywords: diagram, authoring, services, PPTX, Notes, PowerPoint
