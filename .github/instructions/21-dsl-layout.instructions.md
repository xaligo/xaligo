---
applyTo: "**/*.xal"
---

# XAL layout

Validation and rendering use the same layout and geometry checks. Fixed
children reserve space before flexible weights. Default overflow is `error`;
`visible` preserves deterministic cursor advancement. Drawable geometry must
be finite and positive.

Parent layout owns allocation, content boxes, containment, item-grid occupancy,
and text policy. Encoders consume resolved results. Exact tags, spacing,
domains, and algorithms: `reference.md` section `07`, “Numeric and Geometry
Contract” through “Layout Calculation Rules”.
