---
applyTo: ".github/instructions/manual/**"
---

# 10.10 PPTX and routing: Group Header Tags

## Group Header Tags

- Group header tag labels are single-line in every output whose text engine can
  represent that policy. The shared draw plan marks their semantic role,
  wrapping, fitting, clipping, line height, and padding; the TS drawing layer
  consumes those values rather than inferring behavior from an element ID.
- Shared scene generation estimates tag labels by glyph class before SVG/PPTX
  projection: East Asian wide/full-width characters use one font-size unit,
  while proportional half-width characters use narrower factors. Keep the 8px
  text-box safety allowance and only 4px between that box and the tag tip so
  PowerPoint no-wrap text stays inside a compact tag background.
- When changing group tag font size, font family, padding, or tag geometry,
  update both the scene width estimate and the group-header regression tests.
