---
applyTo: ".github/instructions/manual/**"
---

# 07.03.01 XAL specification: Frame and physical-page contract

### Frame and physical-page contract

An identified child `<frame>` is the V1 physical page unit. Frames are emitted
in source order after the complete document scene and all cross-frame links
have been resolved.

| Format | Default mapping |
|---|---|
| SVG | One `.svg` artifact per frame |
| PPTX | One slide per frame |
| PDF | One page per frame |
| Excel | One worksheet per frame, containing the frame's SVG image |
| Excalidraw, XYFlow, Isoflow | One logical document containing all frames |

SVG, PPTX, PDF, and Excel omit the page-frame outline in both default and
`--combine-frames` output. Frame geometry remains authoritative for page size,
cropping, endpoint ownership, and the outer logical page edge used to select a
cross-frame page-link side and tangent anchor. The drawable frame terminal may
sit on a parallel inward inset line. A default page-local SVG uses the exact
frame rectangle as its canvas and clip boundary; PDF pages and Excel page
images inherit that strict crop. Combined SVG compatibility output retains
marker-safe bounds expansion. Excalidraw retains editable frame structure with
transparent page-frame strokes.

For a document with one child frame, SVG writes exactly the requested output
path. For multiple child frames, an output request such as `diagram.svg`
produces `diagram-<safe-frame-id>.svg` for each frame. The safe ID retains ASCII
letters, digits, `_`, and `-`; every run of other characters becomes one `-`,
leading and trailing `-` are removed, and an empty result falls back to
`frame-<source-order>`. Two IDs that resolve to the same output filename are an
error. SVG does not create an implicit archive.

`--combine-frames` is the explicit compatibility option for page-oriented
formats. It restores the historical single canvas, single slide, single PDF
page, or single Excel worksheet. It does not change Excalidraw, XYFlow, or
Isoflow because those formats are already single logical documents.
