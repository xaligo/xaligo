---
applyTo: "**/*.xal"
---

# XAL document

Canonical V1 is `<xaligo version="1">` with exactly one `<frames>`; identified
child frames are physical pages. Legacy roots warn. V2 is
`<scene version="2">`; never dispatch by parser fallback.

Numeric values are finite and follow their documented domains. References,
metadata, page mapping, tables, databases, UML, defaults, and style precedence
must follow `reference.md` section `07`.
