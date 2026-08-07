---
applyTo: ".github/instructions/manual/**"
---

# 04.06 Feature catalog: Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

## Group 6 — Output Formats & Rendering (`XAL-6xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-6000010 | Excalidraw output | Excluded unless justified | Retired public output; the remaining Excalidraw-shaped V1 scene is an internal compatibility stage only. |
| XAL-6000020 | SVG output | Implemented | Standalone SVG, one file per frame by default or one combined canvas with `--combine-frames`. |
| XAL-6000030 | PowerPoint (PPTX) output | Implemented | PPTX presentation built from a shared Go draw plan and executed by the statically linked Rust `pptx` exporter, one slide per frame by default. |
| XAL-6000040 | PDF output | Excluded unless justified | Retired output and native PDF dependencies removed. |
| XAL-6000050 | Excel output | Excluded unless justified | Retired output, XLSX alias, and spreadsheet dependencies removed. |
| XAL-6000060 | XYFlow (React Flow) output | Excluded unless justified | Retired graph-integration output and public APIs removed. |
| XAL-6000070 | Isoflow output | Excluded unless justified | Retired isometric output, manifest, generated icons, and public APIs removed. |
| XAL-6000080 | Frame-to-physical-page mapping contract | Implemented | Shared default mapping of one identified child frame to one SVG file or PPTX slide. |
| XAL-6000090 | `--combine-frames` compatibility mode | Implemented | Combines all frames onto one SVG canvas or PPTX slide. |
| XAL-6000100 | Safe frame-ID output naming | Implemented | Multi-frame SVG output derives `<stem>-<safe-frame-id>.svg` names, with deterministic collision detection. |
| XAL-6000110 | SVG legend placement | Implemented | `--svg-legend-position top|right|bottom|left` controls where the services.csv-derived legend renders. |
| XAL-6000120 | PPTX legend slides | Implemented | Dedicated legend slide(s) appended after diagram slides, using a 4-column icon/abbreviation/official-name layout. |
| XAL-6000130 | PDF page cropping | Excluded unless justified | Retired with PDF output. |
| XAL-6000140 | Excel worksheet SVG embedding | Excluded unless justified | Retired with Excel output. |
| XAL-6000150 | Single-logical-document output invariance | Excluded unless justified | Retired with Excalidraw, XYFlow, and Isoflow outputs. |
| XAL-6000160 | Isoflow generic-endpoint projection | Excluded unless justified | Retired with Isoflow output. |
| XAL-6000170 | Oversized-frame page tiling | Planned | Generic tiling of one oversized frame across several physical pages/slides, distinct from the existing one-frame-per-page mapping. |
| XAL-6000180 | Renderer capability/projection contract | Planned | A typed SVG/PPTX capability declaration so any intentionally lossy projection is explicit and tested. |
