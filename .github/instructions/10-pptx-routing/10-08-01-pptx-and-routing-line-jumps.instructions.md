---
applyTo: ".github/instructions/manual/**"
---

# 10.08.01 PPTX and routing: Line Jumps

### Line Jumps

The shared draw plan implements line jumps for SVG/PPTX.

Current approach:

- Detect line segment intersections after routing.
- Determine which line is visually above the other by layer/kind/order.
- Render jumps as a 6px background-colored mask below the upper line in
  SVG/PPTX output. The mask uses the uppermost opaque container background at
  the crossing. A curved arc may replace the rectangular mask later.
Markdown inherits the SVG line-jump projection.
