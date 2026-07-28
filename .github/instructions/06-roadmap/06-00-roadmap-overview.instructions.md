---
applyTo: ".github/instructions/manual/**"
---

# 06.00 Roadmap


This roadmap is a planning precondition for future implementation work. Prefer
changes that move xaligo toward a diagram-as-code platform with a clean render
pipeline, SVG-first preview capability, network-diagram primitives, and
eventual VS Code / PPTX integration.

Implementation guidance:

- Keep the core pipeline separable as `.xal -> parser -> layout -> renderer`.
- Treat Excalidraw, SVG, PPTX, PDF, Excel, XYFlow, and Isoflow as output
  renderers/projections over a shared model where possible.
- Prioritize SVG renderer and network diagram primitives before advanced PPTX
  feature polish when the choice is otherwise ambiguous.
- Route/traffic separation, route connectors, orthogonal routing, edge offsets,
  layer routing, junctions, and line jumps are roadmap features, not one-off
  export hacks.
- Live preview and VS Code integration should build on `xaligo render` /
  `xaligo validate`, not separate hidden pipelines.
