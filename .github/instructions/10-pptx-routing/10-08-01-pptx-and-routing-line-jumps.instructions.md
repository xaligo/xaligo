---
applyTo: ".github/instructions/manual/**"
---

# 10.08.01 PPTX and routing: Line Jumps

### Line Jumps

Excalidraw does not provide reliable built-in line jumps/bridges for this
workflow. The shared draw plan therefore implements them for SVG/PPTX.

Current approach:

- Detect line segment intersections after routing.
- Determine which line is visually above the other by layer/kind/order.
- Render jumps as a 6px background-colored mask below the upper line in
  SVG/PPTX output. The mask uses the uppermost opaque container background at
  the crossing. A curved arc may replace the rectangular mask later.
- For Excalidraw output, approximate with normal lines or supported shape
  primitives when necessary.

SVG preview and PPTX can support line jumps more accurately than Excalidraw JSON.
