---
applyTo: ".github/instructions/manual/**"
---

# 05.06 Issues and quality: Phase 2 Quality Order

## Phase 2 Quality Order

Harden Phase 2 features in this order unless a blocking defect requires a
smaller local detour:

1. Canonical V1 document envelope: `<xaligo version="1">`, `<data>`,
   `<frames>`, identified `<frame>`, legacy compatibility, and migration
   diagnostics.
2. UML diagrams: first semantic accuracy for each UML kind, then design
   polish for that kind, then shared UML visual-language and cross-format
   parity. Prioritize activity, class, sequence, state machine, timing,
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
9. Page-oriented outputs: one frame per SVG artifact or PPTX slide;
   `--combine-frames`; safe artifact IDs; crop boundaries; and
   filename collision handling.
10. Renderer matrix: SVG baseline behavior, PPTX plan parity, Markdown SVG
    embedding, retired-format absence, and browser-WASM dependency boundaries.
11. CI, release, and tooling: Go, Rust, npm, docs, security, version gates,
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
