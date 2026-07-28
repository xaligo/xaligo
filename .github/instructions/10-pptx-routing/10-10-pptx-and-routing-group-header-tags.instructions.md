---
applyTo: ".github/instructions/manual/**"
---

# 10.10 PPTX and routing: Group Header Tags

## Group Header Tags

- Group header tag labels are single-line in every output whose text engine can
  represent that policy. The shared draw plan marks their semantic role,
  wrapping, fitting, clipping, line height, and padding; the TS drawing layer
  consumes those values rather than inferring behavior from an element ID.
- Excalidraw scene generation must reserve conservative tag label width before
  PPTX export. `groupLabelCharW` is intentionally larger than the average
  Excalidraw text metric so PowerPoint no-wrap text stays inside the tag
  background.
- When changing group tag font size, font family, padding, or tag geometry,
  update both the scene width estimate and the group-header regression tests.
