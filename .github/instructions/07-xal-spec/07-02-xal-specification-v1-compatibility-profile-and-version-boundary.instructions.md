---
applyTo: ".github/instructions/manual/**"
---

# 07.02 XAL specification: V1 Compatibility Profile and Version Boundary

## V1 Compatibility Profile and Version Boundary

Canonical V1 source explicitly sets `version="1"` on `<xaligo>`. An
unversioned `<xaligo>` defaults to V1 with a warning. A `version` value other
than `1` is invalid. Legacy `<frame>` and `<frames>` roots accept the historical
V1 version rules but always emit a warning recommending the canonical envelope.
This document-root `version` selects the DSL and is not visible page metadata.
By contrast, a non-empty `version` on an identified `<frame>` that is a direct
child of the document-root `<frames>` is that page's visible content revision;
it does not select a language version. Structural diff ignores only the
document-root DSL version and compares child-frame content revisions normally.

V2 uses a distinct, reject-safe root:

```xml
<scene version="2">
  ...
</scene>
```

`<scene>` requires `version="2"`; an unversioned `<scene>` is invalid. A V1
reader recognizes `<xaligo>`, `<frame>`, and `<frames>`, but rejects a
V2 document at the root instead of partially rendering V2 syntax as V1. Do not
use `<frame version="2">` or `<frames version="2">`.

A V2 implementation must accept this V1 profile as input, preserve its
defaults and compatibility behavior, and lower it directly to the shared typed
model. It must not rewrite V1 XML into V2 XML, parse the document twice, or
invoke V1 through a serialized intermediate representation. V1 has no
dependency on, and no obligation to understand, V2.

Canonical V1 source uses lowercase XML tag names, attribute names, and enum
tokens exactly as documented here. Historical case-insensitive or directional
aliases that are not listed in this specification are accepted implementation
details, not part of the frozen compatibility profile. A V2 compatibility
frontend canonicalizes the documented V1 values once at its input boundary.
