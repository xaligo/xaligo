---
applyTo: ".github/instructions/manual/**"
---

# 04.06 Feature catalog: Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

## Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-6000010 | Excalidraw output | Implemented | Editable Excalidraw JSON scene with groups, connectors, and metadata tags. |
| XAL-6000020 | SVG output | Implemented | Standalone SVG, one file per frame by default or one combined canvas with `--combine-frames`. |
| XAL-6000030 | PowerPoint (PPTX) output | Implemented | PPTX presentation built from a shared Go draw plan and executed by a configured WASM/PptxGenJS exporter, one slide per frame by default. |
| XAL-6000040 | PDF output | Implemented | PDF document, one page per frame by default, inheriting SVG's strict crop boundary. |
| XAL-6000050 | Excel output | Implemented | Excel workbook, one worksheet per frame by default, each embedding that frame's rendered SVG image. |
| XAL-6000060 | XYFlow (React Flow) output | Implemented | XYFlow-compatible JSON nodes and edges for React Flow-based viewers. |
| XAL-6000070 | Isoflow output | Implemented | Isoflow-compatible model JSON for isometric network diagrams. |
| XAL-6000080 | Frame-to-physical-page mapping contract | Implemented | Shared default mapping of one identified child frame to one SVG file/PPTX slide/PDF page/Excel worksheet. |
| XAL-6000090 | `--combine-frames` compatibility mode | Implemented | Restores the historical single-canvas/page form for SVG, PPTX, PDF, and Excel; Excalidraw/XYFlow/Isoflow are unaffected since they are already single documents. |
| XAL-6000100 | Safe frame-ID output naming | Implemented | Multi-frame SVG output derives `<stem>-<safe-frame-id>.svg` names, with deterministic collision detection. |
| XAL-6000110 | SVG legend placement | Implemented | `--svg-legend-position top|right|bottom|left` controls where the services.csv-derived legend renders. |
| XAL-6000120 | PPTX legend slides | Implemented | Dedicated legend slide(s) appended after diagram slides, using a 4-column icon/abbreviation/official-name layout. |
| XAL-6000130 | PDF page cropping | Implemented | PDF pages inherit the exact per-frame SVG canvas and clip boundary as their strict crop. |
| XAL-6000140 | Excel worksheet SVG embedding | Implemented | Each Excel worksheet embeds its frame's rendered SVG image rather than reconstructing native shapes. |
| XAL-6000150 | Single-logical-document output invariance | Implemented | Excalidraw, XYFlow, and Isoflow always remain one logical document regardless of `--combine-frames`. |
| XAL-6000160 | Isoflow generic-endpoint projection | Implemented | UML and other shapes without an Isoflow-native equivalent project to a labeled generic endpoint icon. |
| XAL-6000170 | Oversized-frame page tiling | Planned | Generic tiling of one oversized frame across several physical pages/slides, distinct from the existing one-frame-per-page mapping. |
| XAL-6000180 | Renderer capability/projection contract | Planned | A typed per-format capability declaration (e.g., which formats can carry arbitrary UML/connector metadata) so unsupported projections become explicit instead of implicit. |
