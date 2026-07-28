---
applyTo: "**"
---

# Instruction index

This is the only always-loaded instruction. Read the task route, then open only
the listed numbered files. Re-evaluate when scope changes. Numbers are the
canonical reading order; within a chapter read `.00` before relevant sections.

## Task routes

| Task | Read in order |
|---|---|
| Any repository change | 01, 02, 03 |
| Scope/status/planning | 04, 05, 06 |
| `.xal` syntax or behavior | 07, then 08 for implementation |
| Go/TypeScript structure | 08, 09 |
| PPTX/connectors/routing | relevant 07 sections, then 08 and 10 |
| Diagram authoring | relevant 07 sections, then 11 |

Preserve unrelated worktree changes. Implementation requests allow local
commits; never push, publish, tag, open a PR, or rewrite history without explicit
permission. Detail files use a nonmatching `applyTo` intentionally and are
loaded through this index.

## Numbered contents

### 01 General

- 01.00 [Overview](01-00-general-overview.instructions.md)
- 01.01 [Project](01-01-general-project.instructions.md)
- 01.02 [Directory structure](01-02-general-directory-structure.instructions.md)
- 01.03 [Architecture rules](01-03-general-architecture-rules.instructions.md)
- 01.04 [Testing rules](01-04-general-testing-rules.instructions.md)
- 01.05 [Assets and configuration](01-05-general-assets-and-configuration.instructions.md)
- 01.06 [Conventions](01-06-general-conventions.instructions.md)
- 01.07 [Verification](01-07-general-verification.instructions.md)


### 02 Agent guide

- 02.00 [Overview](02-00-agent-guide-overview.instructions.md)
- 02.01 [Project summary](02-01-agent-guide-project-summary.instructions.md)
- 02.02 [Working rules](02-02-agent-guide-working-rules.instructions.md)
- 02.03 [Common commands](02-03-agent-guide-common-commands.instructions.md)
- 02.04 [Shared Use-Case APIs](02-04-agent-guide-shared-use-case-apis.instructions.md)
- 02.05 [Asset workflow](02-05-agent-guide-asset-workflow.instructions.md)
- 02.06 [Services CSV](02-06-agent-guide-services-csv.instructions.md)
- 02.07 [Completion checklist](02-07-agent-guide-completion-checklist.instructions.md)


### 03 Development flow

- 03.00 [Overview](03-00-development-flow-overview.instructions.md)
- 03.01 [Authorization and Scope](03-01-development-flow-authorization-and-scope.instructions.md)
- 03.02 [Before Changing Files](03-02-development-flow-before-changing-files.instructions.md)
- 03.03 [Change Slices](03-03-development-flow-change-slices.instructions.md)
- 03.04 [Commit Cadence](03-04-development-flow-commit-cadence.instructions.md)
- 03.05 [Working-Tree Safety](03-05-development-flow-working-tree-safety.instructions.md)
- 03.06 [Verification Cadence](03-06-development-flow-verification-cadence.instructions.md)
- 03.07 [Documentation and Architecture Diagrams](03-07-development-flow-documentation-and-architecture-diagrams.instructions.md)
- 03.08 [Completion Audit](03-08-development-flow-completion-audit.instructions.md)


### 04 Feature catalog

- 04.00 [Overview](04-00-feature-catalog-overview.instructions.md)
- 04.01 [Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)](04-01-feature-catalog-group-1-core-dsl-document-envelope-xal-1xxxxxx.instructions.md)
- 04.02 [Group 2 — Layout & Composition Primitives (`XAL-2xxxxxx`)](04-02-feature-catalog-group-2-layout-composition-primitives-xal-2xxxxx.instructions.md)
- 04.03 [Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)](04-03-feature-catalog-group-3-aws-architecture-primitives-icon-catalog.instructions.md)
- 04.04 [Group 4 — Connections & Routing (`XAL-4xxxxxx`)](04-04-feature-catalog-group-4-connections-routing-xal-4xxxxxx.instructions.md)
- 04.05 [Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)](04-05-feature-catalog-group-5-content-elements-tables-databases-uml-xa.instructions.md)
- 04.06 [Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)](04-06-feature-catalog-group-6-output-formats-rendering-xal-6xxxxxx.instructions.md)
- 04.07 [Group 7 — Presentation, Theming & Physical Page Fitting (`XAL-7xxxxxx`)](04-07-feature-catalog-group-7-presentation-theming-physical-page-fitti.instructions.md)
- 04.08 [Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)](04-08-feature-catalog-group-8-cli-commands-developer-tooling-xal-8xxxx.instructions.md)
- 04.09 [Group 9 — Validation, Diagnostics, Diff & External Integration (`XAL-9xxxxxx`)](04-09-feature-catalog-group-9-validation-diagnostics-diff-external-int.instructions.md)


### 05 Issues and quality

- 05.00 [Overview](05-00-issues-and-quality-overview.instructions.md)
- 05.01 [Quality Rubric](05-01-issues-and-quality-quality-rubric.instructions.md)
- 05.02 [Feature Quality Workflow](05-02-issues-and-quality-feature-quality-workflow.instructions.md)
- 05.03 [Status Tracking](05-03-issues-and-quality-status-tracking.instructions.md)
- 05.04 [Detailed Task Backlog](05-04-issues-and-quality-detailed-task-backlog.instructions.md)
  - 05.04.01 [Q01 Canonical V1 Document Envelope](05-04-01-issues-and-quality-q01-canonical-v1-document-envelope.instructions.md)
  - 05.04.02 [Q02 Data Registry and Imports](05-04-02-issues-and-quality-q02-data-registry-and-imports.instructions.md)
  - 05.04.03 [Q03 Tables](05-04-03-issues-and-quality-q03-tables.instructions.md)
  - 05.04.04 [Q04 Relational Databases](05-04-04-issues-and-quality-q04-relational-databases.instructions.md)
  - 05.04.05 [Q05 UML Diagrams](05-04-05-issues-and-quality-q05-uml-diagrams.instructions.md)
  - 05.04.06 [Q06 Frame-Scoped References](05-04-06-issues-and-quality-q06-frame-scoped-references.instructions.md)
  - 05.04.07 [Q07 Cross-Frame Page Links](05-04-07-issues-and-quality-q07-cross-frame-page-links.instructions.md)
  - 05.04.08 [Q08 Frame Metadata Bands](05-04-08-issues-and-quality-q08-frame-metadata-bands.instructions.md)
  - 05.04.09 [Q09 Page-Oriented Outputs](05-04-09-issues-and-quality-q09-page-oriented-outputs.instructions.md)
  - 05.04.10 [Q10 Renderer Matrix](05-04-10-issues-and-quality-q10-renderer-matrix.instructions.md)
  - 05.04.11 [Q11 CI, Release, and Tooling](05-04-11-issues-and-quality-q11-ci-release-and-tooling.instructions.md)
  - 05.04.12 [Q12 Diagnostics and Error UX](05-04-12-issues-and-quality-q12-diagnostics-and-error-ux.instructions.md)
  - 05.04.13 [Q13 Determinism, Concurrency, and Performance](05-04-13-issues-and-quality-q13-determinism-concurrency-and-performance.instructions.md)
  - 05.04.14 [Q14 CLI, API, and Preview Contracts](05-04-14-issues-and-quality-q14-cli-api-and-preview-contracts.instructions.md)
  - 05.04.15 [Q15 Themes and Accessibility](05-04-15-issues-and-quality-q15-themes-and-accessibility.instructions.md)
  - 05.04.16 [Q16 Structural Diff](05-04-16-issues-and-quality-q16-structural-diff.instructions.md)
  - 05.04.17 [Q17 Reproducibility and Artifact Integrity](05-04-17-issues-and-quality-q17-reproducibility-and-artifact-integrity.instructions.md)
  - 05.04.18 [Q18 Configuration, Logging, and Observability](05-04-18-issues-and-quality-q18-configuration-logging-and-observability.instructions.md)
- 05.05 [Design Quality Gate](05-05-issues-and-quality-design-quality-gate.instructions.md)
- 05.06 [Phase 2 Quality Order](05-06-issues-and-quality-phase-2-quality-order.instructions.md)
- 05.07 [First UML Quality Slice](05-07-issues-and-quality-first-uml-quality-slice.instructions.md)
- 05.08 [Definition of Done](05-08-issues-and-quality-definition-of-done.instructions.md)


### 06 Roadmap

- 06.00 [Overview](06-00-roadmap-overview.instructions.md)
- 06.01 [Product Architecture Preconditions](06-01-roadmap-product-architecture-preconditions.instructions.md)
  - 06.01.01 [V1 Structured-Diagram Profile](06-01-01-roadmap-v1-structured-diagram-profile.instructions.md)
  - 06.01.02 [Common DSL and Go Core](06-01-02-roadmap-common-dsl-and-go-core.instructions.md)
  - 06.01.03 [V1 Compatibility and V2 Input](06-01-03-roadmap-v1-compatibility-and-v2-input.instructions.md)
  - 06.01.04 [Mode and Format Are Independent](06-01-04-roadmap-mode-and-format-are-independent.instructions.md)
  - 06.01.05 [Shared Rendering APIs](06-01-05-roadmap-shared-rendering-apis.instructions.md)
  - 06.01.06 [Rendering Correctness Gate](06-01-06-roadmap-rendering-correctness-gate.instructions.md)
- 06.02 [Delivery Phases](06-02-roadmap-delivery-phases.instructions.md)
  - 06.02.01 [Phase 1: Basic Output](06-02-01-roadmap-phase-1-basic-output.instructions.md)
  - 06.02.02 [Phase 2: Network Diagram Features](06-02-02-roadmap-phase-2-network-diagram-features.instructions.md)
  - 06.02.03 [Phase 3: Live Preview](06-02-03-roadmap-phase-3-live-preview.instructions.md)
- 06.03 [VS Code Extension Preconditions](06-03-roadmap-vs-code-extension-preconditions.instructions.md)
- 06.04 [AWS 2.5D Mode](06-04-roadmap-aws-2-5d-mode.instructions.md)
- 06.05 [Export Roadmap](06-05-roadmap-export-roadmap.instructions.md)
- 06.06 [Long-Term Product Position](06-06-roadmap-long-term-product-position.instructions.md)
- 06.07 [Current State](06-07-roadmap-current-state.instructions.md)
- 06.08 [Rebaselined Implementation Order](06-08-roadmap-rebaselined-implementation-order.instructions.md)
- 06.09 [v0.1 Foundation](06-09-roadmap-v0-1-foundation.instructions.md)
  - 06.09.01 [Rendering Engine Refactoring](06-09-01-roadmap-rendering-engine-refactoring.instructions.md)
  - 06.09.02 [Public API](06-09-02-roadmap-public-api.instructions.md)
  - 06.09.03 [CLI](06-09-03-roadmap-cli.instructions.md)
- 06.10 [v0.2 SVG Renderer](06-10-roadmap-v0-2-svg-renderer.instructions.md)
  - 06.10.01 [SVG Export](06-10-01-roadmap-svg-export.instructions.md)
  - 06.10.02 [Supported Elements](06-10-02-roadmap-supported-elements.instructions.md)
  - 06.10.03 [Themes](06-10-03-roadmap-themes.instructions.md)
- 06.11 [v0.3 Network Diagram Features](06-11-roadmap-v0-3-network-diagram-features.instructions.md)
  - 06.11.01 [Route Connector](06-11-01-roadmap-route-connector.instructions.md)
  - 06.11.02 [Connector Model](06-11-02-roadmap-connector-model.instructions.md)
  - 06.11.03 [Orthogonal Routing](06-11-03-roadmap-orthogonal-routing.instructions.md)
  - 06.11.04 [Route / Traffic Separation](06-11-04-roadmap-route-traffic-separation.instructions.md)
  - 06.11.05 [DSL](06-11-05-roadmap-dsl.instructions.md)
- 06.12 [v0.4 Advanced Routing](06-12-roadmap-v0-4-advanced-routing.instructions.md)
  - 06.12.01 [Edge Offset](06-12-01-roadmap-edge-offset.instructions.md)
  - 06.12.02 [Layer Routing](06-12-02-roadmap-layer-routing.instructions.md)
  - 06.12.03 [Junction Generation](06-12-03-roadmap-junction-generation.instructions.md)
- 06.13 [v0.5 Line Jumps](06-13-roadmap-v0-5-line-jumps.instructions.md)
  - 06.13.01 [Bridge / Jump Lines](06-13-01-roadmap-bridge-jump-lines.instructions.md)
  - 06.13.02 [Features](06-13-02-roadmap-features.instructions.md)
- 06.14 [v0.6 Live Preview](06-14-roadmap-v0-6-live-preview.instructions.md)
  - 06.14.01 [xaligo serve](06-14-01-roadmap-xaligo-serve.instructions.md)
  - 06.14.02 [Features](06-14-02-roadmap-features.instructions.md)
  - 06.14.03 [Backend Stack](06-14-03-roadmap-backend-stack.instructions.md)
  - 06.14.04 [Frontend Stack](06-14-04-roadmap-frontend-stack.instructions.md)
  - 06.14.05 [Preview Flow](06-14-05-roadmap-preview-flow.instructions.md)
- 06.15 [v0.7 VS Code Extension](06-15-roadmap-v0-7-vs-code-extension.instructions.md)
  - 06.15.01 [Language Support](06-15-01-roadmap-language-support.instructions.md)
  - 06.15.02 [Features](06-15-02-roadmap-features.instructions.md)
  - 06.15.03 [Preview Panel](06-15-03-roadmap-preview-panel.instructions.md)
  - 06.15.04 [Live Preview](06-15-04-roadmap-live-preview.instructions.md)
- 06.16 [v0.8 Excalidraw Integration](06-16-roadmap-v0-8-excalidraw-integration.instructions.md)
  - 06.16.01 [Excalidraw Preview](06-16-01-roadmap-excalidraw-preview.instructions.md)
  - 06.16.02 [Features](06-16-02-roadmap-features.instructions.md)
  - 06.16.03 [Excalidraw Export](06-16-03-roadmap-excalidraw-export.instructions.md)
- 06.17 [v0.9 PowerPoint Export](06-17-roadmap-v0-9-powerpoint-export.instructions.md)
  - 06.17.01 [PPTX Export](06-17-01-roadmap-pptx-export.instructions.md)
  - 06.17.02 [Supported Features](06-17-02-roadmap-supported-features.instructions.md)
- 06.18 [v1.0](06-18-roadmap-v1-0.instructions.md)
  - 06.18.01 [VS Code Marketplace Release](06-18-01-roadmap-vs-code-marketplace-release.instructions.md)
- 06.19 [Future Vision](06-19-roadmap-future-vision.instructions.md)
  - 06.19.01 [AWS Architecture Mode](06-19-01-roadmap-aws-architecture-mode.instructions.md)
  - 06.19.02 [Network Diagram Mode](06-19-02-roadmap-network-diagram-mode.instructions.md)
  - 06.19.03 [Infrastructure as Diagram](06-19-03-roadmap-infrastructure-as-diagram.instructions.md)
- 06.20 [Project Goal](06-20-roadmap-project-goal.instructions.md)


### 07 XAL specification

- 07.00 [Overview](07-00-xal-specification-overview.instructions.md)
- 07.01 [Overview](07-01-xal-specification-overview.instructions.md)
- 07.02 [V1 Compatibility Profile and Version Boundary](07-02-xal-specification-v1-compatibility-profile-and-version-boundary.instructions.md)
- 07.03 [Root Tag](07-03-xal-specification-root-tag.instructions.md)
  - 07.03.01 [Frame and physical-page contract](07-03-01-xal-specification-frame-and-physical-page-contract.instructions.md)
  - 07.03.02 [Frame metadata tag band](07-03-02-xal-specification-frame-metadata-tag-band.instructions.md)
- 07.04 [Numeric and Geometry Contract](07-04-xal-specification-numeric-and-geometry-contract.instructions.md)
  - 07.04.01 [Fixed and flexible child allocation](07-04-01-xal-specification-fixed-and-flexible-child-allocation.instructions.md)
- 07.05 [Layout Tags](07-05-xal-specification-layout-tags.instructions.md)
  - 07.05.01 [`<container>`](07-05-01-xal-specification-container.instructions.md)
  - 07.05.02 [`<row>`](07-05-02-xal-specification-row.instructions.md)
  - 07.05.03 [`<col>`](07-05-03-xal-specification-col.instructions.md)
- 07.06 [Custom Leaf and Container Tags](07-06-xal-specification-custom-leaf-and-container-tags.instructions.md)
- 07.07 [`<rectangle>` and `<port>` Tags](07-07-xal-specification-rectangle-and-port-tags.instructions.md)
- 07.08 [Resolved Text Layout](07-08-xal-specification-resolved-text-layout.instructions.md)
- 07.09 [`<table>` Tag](07-09-xal-specification-table-tag.instructions.md)
- 07.10 [Relational Database Tags](07-10-xal-specification-relational-database-tags.instructions.md)
- 07.11 [UML Tags](07-11-xal-specification-uml-tags.instructions.md)
  - 07.11.01 [Component, identity, and layout contract](07-11-01-xal-specification-component-identity-and-layout-contract.instructions.md)
  - 07.11.02 [Diagram-kind vocabulary](07-11-02-xal-specification-diagram-kind-vocabulary.instructions.md)
  - 07.11.03 [Component diagram sizing](07-11-03-xal-specification-component-diagram-sizing.instructions.md)
  - 07.11.04 [Activity partitions](07-11-04-xal-specification-activity-partitions.instructions.md)
  - 07.11.05 [Ownership](07-11-05-xal-specification-ownership.instructions.md)
  - 07.11.06 [Element compartments](07-11-06-xal-specification-element-compartments.instructions.md)
  - 07.11.07 [Relation attributes, order, and time](07-11-07-xal-specification-relation-attributes-order-and-time.instructions.md)
  - 07.11.08 [Relation projection](07-11-08-xal-specification-relation-projection.instructions.md)
  - 07.11.09 [Reusable UML models](07-11-09-xal-specification-reusable-uml-models.instructions.md)
  - 07.11.10 [Deliberately lossy V1 projection](07-11-10-xal-specification-deliberately-lossy-v1-projection.instructions.md)
- 07.12 [`<item>` Tag](07-12-xal-specification-item-tag.instructions.md)
- 07.13 [`<spacer>` / `<blank>` Tags](07-13-xal-specification-spacer-blank-tags.instructions.md)
- 07.14 [`<connection>` Tag](07-14-xal-specification-connection-tag.instructions.md)
  - 07.14.01a [`<connections>` Tag, part 1/4](07-14-01-xal-specification-connections-tag-part-01.instructions.md)
  - 07.14.01b [`<connections>` Tag, part 2/4](07-14-01-xal-specification-connections-tag-part-02.instructions.md)
  - 07.14.01c [`<connections>` Tag, part 3/4](07-14-01-xal-specification-connections-tag-part-03.instructions.md)
  - 07.14.01d [`<connections>` Tag, part 4/4](07-14-01-xal-specification-connections-tag-part-04.instructions.md)
- 07.15 [AWS Group Tags](07-15-xal-specification-aws-group-tags.instructions.md)
  - 07.15.01 [Layout Control Attributes (shared by all containers)](07-15-01-xal-specification-layout-control-attributes-shared-by-all-containe.instructions.md)
  - 07.15.02 [Child Size Ratio Attributes](07-15-02-xal-specification-child-size-ratio-attributes.instructions.md)
- 07.16 [Spacing Classes (`class` attribute)](07-16-xal-specification-spacing-classes-class-attribute.instructions.md)
  - 07.16.01 [All-sides shorthand](07-16-01-xal-specification-all-sides-shorthand.instructions.md)
  - 07.16.02 [Axis shorthand](07-16-02-xal-specification-axis-shorthand.instructions.md)
  - 07.16.03 [Per-side](07-16-03-xal-specification-per-side.instructions.md)
  - 07.16.04 [Semantics](07-16-04-xal-specification-semantics.instructions.md)
- 07.17 [Layout Calculation Rules](07-17-xal-specification-layout-calculation-rules.instructions.md)
- 07.18 [Example](07-18-xal-specification-example.instructions.md)
- 07.19 [Constraints and Notes](07-19-xal-specification-constraints-and-notes.instructions.md)


### 08 Architecture

- 08.00 [Overview](08-00-architecture-overview.instructions.md)
- 08.01 [Core pipeline](08-01-architecture-core-pipeline.instructions.md)
- 08.02 [Language-version boundary](08-02-architecture-language-version-boundary.instructions.md)
- 08.03 [Package responsibilities](08-03-architecture-package-responsibilities.instructions.md)
- 08.04 [Invariants](08-04-architecture-invariants.instructions.md)
- 08.05 [File organization](08-05-architecture-file-organization.instructions.md)
- 08.06 [Geometry contract](08-06-architecture-geometry-contract.instructions.md)
- 08.07 [Renderer-neutral text contract](08-07-architecture-renderer-neutral-text-contract.instructions.md)
- 08.08 [Dependency direction](08-08-architecture-dependency-direction.instructions.md)
- 08.09 [Verification](08-09-architecture-verification.instructions.md)


### 09 Coding

- 09.00 [Overview](09-00-coding-overview.instructions.md)
- 09.01 [Responsibility-based files](09-01-coding-responsibility-based-files.instructions.md)
  - 09.01.01 [`internal/usecase` root contract](09-01-01-coding-internal-usecase-root-contract.instructions.md)
- 09.02 [V1 engine identifiers](09-02-coding-v1-engine-identifiers.instructions.md)
- 09.03 [Engine execution boundary](09-03-coding-engine-execution-boundary.instructions.md)
- 09.04 [Dependency direction](09-04-coding-dependency-direction.instructions.md)
- 09.05 [Verification](09-05-coding-verification.instructions.md)


### 10 PPTX and routing

- 10.00 [Overview](10-00-pptx-and-routing-overview.instructions.md)
- 10.01 [Brainstorm Reference](10-01-pptx-and-routing-brainstorm-reference.instructions.md)
- 10.02 [Confirmed Decisions](10-02-pptx-and-routing-confirmed-decisions.instructions.md)
- 10.03 [Current Pipeline](10-03-pptx-and-routing-current-pipeline.instructions.md)
- 10.04 [Go / WASM Boundary](10-04-pptx-and-routing-go-wasm-boundary.instructions.md)
- 10.05 [Ownership](10-05-pptx-and-routing-ownership.instructions.md)
- 10.06 [Paper / Scaling](10-06-pptx-and-routing-paper-scaling.instructions.md)
- 10.07 [Routing Rules](10-07-pptx-and-routing-routing-rules.instructions.md)
  - 10.07.01 [Cross-frame page links](10-07-01-pptx-and-routing-cross-frame-page-links.instructions.md)
- 10.08 [Advanced Routing Features](10-08-pptx-and-routing-advanced-routing-features.instructions.md)
  - 10.08.01 [Line Jumps](10-08-01-pptx-and-routing-line-jumps.instructions.md)
  - 10.08.02 [Route / Traffic Separation](10-08-02-pptx-and-routing-route-traffic-separation.instructions.md)
  - 10.08.03 [Route Connectors](10-08-03-pptx-and-routing-route-connectors.instructions.md)
- 10.09 [Connector Style Options](10-09-pptx-and-routing-connector-style-options.instructions.md)
- 10.10 [Group Header Tags](10-10-pptx-and-routing-group-header-tags.instructions.md)
- 10.11 [Item Labels](10-11-pptx-and-routing-item-labels.instructions.md)
- 10.12 [Layout / Whitespace](10-12-pptx-and-routing-layout-whitespace.instructions.md)
- 10.13 [Legend Pages](10-13-pptx-and-routing-legend-pages.instructions.md)
- 10.14 [Verification Checklist](10-14-pptx-and-routing-verification-checklist.instructions.md)


### 11 Diagram creation

- 11.00 [Overview](11-00-diagram-creation-overview.instructions.md)
- 11.01 [Step 1 — Find Service IDs](11-01-diagram-creation-step-1-find-service-ids.instructions.md)
- 11.02 [Step 2 — Create services.csv](11-02-diagram-creation-step-2-create-services-csv.instructions.md)
- 11.03 [Step 3 — Create a .xal file](11-03-diagram-creation-step-3-create-a-xal-file.instructions.md)
  - 11.03.01 [Choosing the right group tag](11-03-01-diagram-creation-choosing-the-right-group-tag.instructions.md)
  - 11.03.02 [Service Scope Validation](11-03-02-diagram-creation-service-scope-validation.instructions.md)
- 11.04 [Step 4 — Render the Excalidraw file](11-04-diagram-creation-step-4-render-the-excalidraw-file.instructions.md)
- 11.05 [Command Reference](11-05-diagram-creation-command-reference.instructions.md)
- 11.06 [PPTX Notes](11-06-diagram-creation-pptx-notes.instructions.md)
