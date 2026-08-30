---
applyTo: ".github/instructions/manual/**"
---

# 08.02 Architecture: Language-version boundary

## Language-version boundary

`<xaligo version="1">` selects canonical V1 and
`<xaligo version="2">` selects V2. Root `<frame>` and `<frames>` are legacy V1
compatibility inputs and emit a migration warning. The V1 parser rejects the
V2 version and does not import or call V2 code.

The parent use-case boundary owns one lightweight root/version dispatch before
engine selection. It must inspect the first XML start element and its version
once, reject unsupported combinations, and pass the original bytes to exactly
one frontend. It must not select a version by retrying another parser after an
error.

The V2-owned envelope frontend has version-selected normalization modes. V2
retains the concise V1 authoring profile and adds native generic parameters;
both lower directly to one typed, version-neutral model consumed by V2 layout,
routing, and format encoders. The compatibility path must not rewrite XML,
parse a document twice, serialize through an intermediate V1 representation,
or invoke a full V1 renderer and then reverse-engineer its output. This one-way
relationship allows V2 to render V1 while V1 remains unaware of V2.

The planned generic calculation core, declarative AWS/UML profile boundaries,
builtin icons, and per-element parameter contract are defined in
`08-10-architecture-v2-generic-engine-and-plugins.instructions.md`.
