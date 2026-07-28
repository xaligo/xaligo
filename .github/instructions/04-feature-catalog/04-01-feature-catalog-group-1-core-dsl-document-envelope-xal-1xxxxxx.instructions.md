---
applyTo: ".github/instructions/manual/**"
---

# 04.01 Feature catalog: Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)

## Group 1 — Core DSL & Document Envelope (`XAL-1xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-1000010 | Vue-style `.xal` XML DSL | Implemented | Vue-style layout DSL with XML syntax, parsed with `encoding/xml`, handling attributes, nested tags, and text content. |
| XAL-1000020 | Canonical V1 envelope | Implemented | `<xaligo version="1">` containing document-wide `<data>` and exactly one `<frames>` page collection. |
| XAL-1000030 | Legacy root compatibility | Implemented | Historical `<frame>`/`<frames>` document roots remain readable and renderable but emit a migration warning recommending the canonical envelope. |
| XAL-1000040 | V2 reject-safe root boundary | Implemented | `<scene version="2">` is a distinct root; a V1 reader rejects it outright instead of partially rendering V2 syntax as V1. |
| XAL-1000050 | Shared V1 pipeline | Implemented | One `.xal -> parser -> layout -> shared scene/plan -> encoder` pipeline shared by every output format, with no per-format parallel parser or layout path. |
| XAL-1000060 | `<frames>` page collection | Implemented | Groups page `<frame>` children with `gap` and optional `layout="vertical"`; every child `<frame>` requires a non-empty `id`. |
| XAL-1000070 | Frame sizing and spacing attributes | Implemented | `width`, `height`, `class`, `margin`/`margin-*`, `content-width`/`content-height`, `align`, and `item-size` on a page or nested frame. |
| XAL-1000080 | Numeric and geometry validation contract | Implemented | Finite-number and per-attribute domain checks (positive sizes, span ≤ 12, non-negative gaps, etc.) applied before layout, shared by `validate` and every render format. |
| XAL-1000090 | Child containment overflow policy | Implemented | `overflow="error"` (default, rejects out-of-bounds children) vs `overflow="visible"` (permits them while keeping sibling cursors deterministic). |
| XAL-1000100 | Frame metadata tag band | Implemented | `<metadata>`/`<entry>` key-value tag band exposing `id`, `title`, page-content `version`, and custom entries as a reserved strip other content must avoid. |
| XAL-1000110 | Page-content `version` attribute | Implemented | A non-empty `version` on an identified child `<frame>` is the page's visible content revision, distinct from the document-root DSL version selector. |
| XAL-1000120 | Document-level `<data>` container | Implemented | Reusable definitions (UML models, database schemas) declared once under `<data>` and referenced by `data="id"` from one or more frames. |
| XAL-1000130 | Resolved shared text-layout policy | Implemented | Role-based wrap/fit/overflow/line-height rules (group header, ordinary label, item label, port label, connector label) applied consistently by every encoder. |
| XAL-1000140 | V2 `<scene version="2">` native engine | Planned | A distinct, reject-safe V2 root and native frontend that lowers directly to the same typed version-neutral model as V1, sharing renderers/encoders downstream; V1 remains a frozen compatibility profile inside V2. |
| XAL-1000150 | Typed normalized layout structure | Planned | Replace repeated reads from `Node.Attrs` with first-class resolved-layout data so validation and rendering share one geometry source of truth. |
| XAL-1000160 | Format-neutral shared scene/plan schema | Planned | Migrate the canonical scene/plan schema away from Excalidraw/PPTX-shaped fields entirely, keeping only compatibility aliases for existing public callers. |
