---
applyTo: ".github/instructions/manual/**"
---

# 06.01.03 Roadmap: V1 Compatibility and V2 Input

### V1 Compatibility and V2 Input

Keep root `<frame>` and `<frames>` as legacy V1 compatibility inputs. Canonical
V1 uses `<xaligo version="1">` with a document-wide `<data>` registry and
identified frames. Legacy roots emit a migration warning. V2 uses the same
envelope as `<xaligo version="2">`; do not place `version="2"` on legacy
`<frame>` or `<frames>` roots. Existing V1 readers reject the unsupported
document version before interpreting nested V2 syntax.

V2 must render both native V2 documents and the frozen V1 profile. Implement
that compatibility in the V2 side only: version-selected normalization lowers
the shared concise authoring profile and native V2 extensions directly to the
same typed, version-neutral model. Keep the existing V1 engine independent of
V2.

The compatibility path is complete only when it preserves V1 defaults,
fallback/error behavior, unknown nested-tag handling, connection-group
inheritance, anchor aliases, numeric catalog-ID range, and render-context item
size. Golden tests must compare V1 and V2-engine output at the neutral-model and
resolved-geometry boundaries across native and embedded targets.

Do not implement compatibility by changing root tags as strings, reparsing,
retrying parsers after syntax errors, serializing through the V1 scene, or
calling the full V1 renderer before V2. Root dispatch reads the first start
element once and selects exactly one frontend; renderers and encoders remain
shared downstream.
