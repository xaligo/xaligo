---
applyTo: ".github/instructions/manual/**"
---

# 08.02 Architecture: Language-version boundary

## Language-version boundary

`<xaligo version="1">` selects canonical V1. Root `<frame>` and `<frames>` are
legacy V1 compatibility inputs and emit a migration warning. Native V2
uses the reject-safe `<scene version="2">` root. The V1 parser is not extended
to recognize `<scene>` and does not import or call V2 code.

The parent use-case boundary owns one lightweight root/version dispatch before
engine selection. It must inspect the first XML start element once, reject
contradictory root/version pairs, and pass the original bytes to exactly one
frontend. It must not select a version by retrying another parser after an
error.

V2 provides two frontends: its native `<scene version="2">` frontend and a V1
compatibility frontend that implements the frozen V1 behavior. Both lower
directly to one typed, version-neutral model consumed by V2 layout, routing,
and format encoders. The V1 compatibility path must not rewrite XML, parse a
document twice, serialize through an intermediate V1 scene, or invoke a full
V1 renderer and then reverse-engineer its output. This one-way relationship
allows V2 to render V1 while V1 remains unaware of V2.
