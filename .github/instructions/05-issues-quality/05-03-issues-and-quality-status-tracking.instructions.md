---
applyTo: ".github/instructions/manual/**"
---

# 05.03 Issues and quality: Status Tracking

## Status Tracking

Maintain the feature hardening status in this file so future agents can resume
from the current quality pass without relying on conversation history.

Use these status values:

- `not-started`: the feature is in the Phase 2 scope but has not been audited.
- `auditing`: contracts, samples, tests, and visual output are being inspected.
- `needs-work`: a quality gap is known and has not been fixed.
- `in-progress`: a fix or quality improvement is actively being implemented.
- `blocked`: progress requires a decision, missing tool, or external input.

There is no `done` status value. Once a feature or task meets the Definition
of Done in this file, remove its row from the status board and backlog
instead of marking it done in place, so this file always reflects only open
issues.

Update the status board when a feature-quality slice starts, when a gap is
found, when validation changes the assessment, and when a slice is committed.
Keep each row concise but evidence-based. Prefer commit hashes, test names,
sample paths, and rendered artifact paths over prose-only status notes.

| Order | Feature | Status | Evidence | Next action |
|---:|---|---|---|---|
| 1 | Canonical V1 document envelope | in-progress | `TestV1ParseValidatesCanonicalEnvelopeHierarchy`; canonical sample validates and renders two SVG artifacts | Finish migration-warning/docs-image audit, then close Q01. |
| 2 | UML diagrams | in-progress | The supported sample matrix is frozen to class, component, activity, state-machine, and sequence; all retained samples validate under the strict profile. Timing-diagram support is accepted as an additional in-scope kind, tracked as Q05.17. | Finish per-diagram visual baselines and focused semantics, then verify cross-format parity. |
| 3 | Data registry and imports | not-started | table data, database schema, SQL import commits exist | Inventory import coverage and diagnostics after the UML precision/design pass starts. |
| 4 | Tables | not-started | `docs/src/examples/samples/tables.xal` | Audit table layout, pipe styling, text fitting, and docs image quality. |
| 5 | Relational databases | not-started | `docs/src/examples/samples/databases.xal` | Audit entity layout, key styling, relation routing, and SQL import behavior. |
| 6 | Frame-scoped references | not-started | frame-qualified endpoint support exists | Audit duplicate ID handling and user-correctable diagnostics. |
| 7 | Cross-frame page links | not-started | `docs/src/examples/samples/page-links.xal` | Audit terminal side selection, labels, insets, and invalid edge handling. |
| 8 | Frame metadata bands | not-started | `docs/src/examples/samples/frame-metadata.xal` | Audit metadata layout, reservation strips, wrapping, and collisions. |
| 9 | Page-oriented outputs | not-started | SVG/PPTX pagination commits exist | Audit default artifacts, `--combine-frames`, safe IDs, and crops. |
| 10 | Renderer matrix | not-started | SVG/PPTX encoders and Markdown embedding | Define representative cross-format checks for changed shared contracts. |
| 11 | CI, release, and tooling | not-started | `VERSION`, workflows, RTK and security instructions | Audit CI gates, version policy, npm lockfile policy, and vendored dependency handling. |
| 12 | Diagnostics and error UX | not-started | `internal/usecase/diagnostics.go`, diagnostics tests | Audit severity, source positions, aggregation, cancellation, and actionable messages. |
| 13 | Determinism, concurrency, and performance | not-started | render determinism and concurrency tests | Audit repeatability, ordering, cancellation, limits, and representative performance. |
| 14 | CLI, API, and preview contracts | not-started | controllers, use-case API tests, preview tests | Audit option parity, exit behavior, live preview, and backward compatibility. |
| 15 | Themes and accessibility | not-started | theme application and renderer theme tests | Audit contrast, semantic distinction, text alternatives, and light/dark parity. |
| 16 | Structural diff | not-started | diff engine and highlight tests | Audit matching, moved/modified classification, highlights, and pagination interaction. |
| 17 | Reproducibility and artifact integrity | not-started | deterministic render and generated docs assets | Audit byte stability, clean regeneration, package contents, and stale artifacts. |
| 18 | Configuration, logging, and observability | not-started | config, options, and shared logging code | Audit defaults, precedence, redaction, error context, and noisy-output behavior. |

When a row's feature or task meets the Definition of Done in this file,
remove the row instead of marking it done, and record the closing commit in
the relevant commit history rather than in this file. If a row is `blocked`,
the next action must state the specific decision or tool needed.
