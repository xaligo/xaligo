---
applyTo: ".github/instructions/manual/**"
---

# 08.07 Architecture: Renderer-neutral text contract

## Renderer-neutral text contract

Every text draw operation carries its resolved box plus a text-layout policy:

```text
wrap | no-wrap
fit: none | shrink
text overflow: visible | clip
line height
content padding
semantic role
```

The semantic role distinguishes ordinary labels, item labels, group headers,
ports, connector labels, and other future text without requiring encoders to
infer behavior from element IDs. Glyph overflow must either be included in
bounds/obstacle calculations or removed by the declared fit/clip policy.

Layout and canonical-scene values are expressed in layout pixels. The current
shared presentation plan stores geometry and padding in inches and font sizes in
points. For effective PPI `p`, conversion is `inch = px / p` and
`pt = px * 72 / p`; paper fitting changes `p` once and both formulas use that
same value. Fixed physical sizes, such as an explicitly specified PPTX label
size, must be represented as an intentional semantic policy rather than an
incidental conversion constant.
