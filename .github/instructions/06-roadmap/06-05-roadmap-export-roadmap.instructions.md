---
applyTo: ".github/instructions/manual/**"
---

# 06.05 Roadmap: Export Roadmap

## Export Roadmap

The cross-version formats are SVG and PPTX; V2 additionally supports terminal
text. Identified child frames are the default
physical page boundary: one SVG artifact or one PPTX slide per frame.
`--combine-frames` retains the combined-canvas/slide behavior. Markdown embeds
the SVG artifacts produced by this same pipeline.

Excalidraw, PDF, Excel/XLSX, XYFlow, and Isoflow export are retired. Do not
retain their encoders, dependencies, assets, aliases, CLI flags, or public API
surface. Generic tiling of one oversized frame across several SVG artifacts or
PPTX slides remains a possible future feature.
