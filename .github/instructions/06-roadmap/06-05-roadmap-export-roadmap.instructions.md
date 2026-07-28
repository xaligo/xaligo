---
applyTo: ".github/instructions/manual/**"
---

# 06.05 Roadmap: Export Roadmap

## Export Roadmap

Primary formats include SVG, Excalidraw, PPTX, PDF, and Excel. The page-oriented
formats use identified child frames as their default physical page boundary:
SVG files, PPTX slides, PDF pages, and Excel worksheets respectively.
`--combine-frames` retains the previous single-canvas behavior. Add or continue:

- XYFlow export for React Flow-style GUI editors. (initial implementation complete)
- Isoflow export for isometric and 2.5D integrations. (initial upstream model export complete)
- Generic tiling of one oversized frame across several pages. Frame pagination
  is implemented and is distinct from this remaining tiling work.

Both exports should consume the shared resolved model; they must not become
alternative parsers for `.xal`.
