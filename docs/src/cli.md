# Command Line

## Render

```bash
xaligo render <file.xal> --format <format> -o <output>
```

Supported formats:

| Format | Output |
|---|---|
| `excalidraw` | Editable Excalidraw JSON |
| `svg` | Standalone SVG; one file per frame by default |
| `pptx` | PowerPoint presentation; one slide per frame by default |
| `pdf` | PDF document; one page per frame by default |
| `excel` | Excel workbook; one frame SVG per worksheet by default |
| `xlsx` | Alias for `excel` |
| `xyflow` | React Flow / XYFlow JSON |
| `isoflow` | Isoflow-compatible model JSON |

UML input (`<uml>...</uml>`) currently rejects `--format excalidraw` because
the editable UML Excalidraw export is disabled. Use `svg`, `pdf`, `pptx`,
`excel`, `xyflow`, or `isoflow` for UML diagrams.

### Frames and physical pages

Identified child frames are physical pages for SVG, PPTX, PDF, and Excel. They
are emitted in source order.

| Format | Multiple-frame default | `--combine-frames` |
|---|---|---|
| SVG | Separate `<output-stem>-<safe-frame-id>.svg` files | One SVG canvas |
| PPTX | One slide per frame in one presentation | One diagram slide |
| PDF | One page per frame in one document | One PDF page |
| Excel | One worksheet per frame in one workbook | One worksheet |

Excalidraw, XYFlow, and Isoflow always remain one logical document, so
`--combine-frames` does not change them. Live SVG preview also uses the combined
canvas so every frame remains visible in one browser view.

For a one-frame SVG document, `-o` is the exact output filename. For several
frames, `-o output/diagram.svg` produces names such as
`output/diagram-overview.svg` and `output/diagram-detail.svg`. Safe frame IDs
retain ASCII letters, digits, `_`, and `-`; other character runs become `-`.
Leading and trailing `-` are removed, an empty result falls back to the frame's
source order, and a filename collision is an error. SVG output is not wrapped
in an implicit ZIP archive.

Common render flags:

| Flag | Description |
|---|---|
| `--mode standard|network|aws` | Accepted V1 rendering mode. These three values currently have identical V1 rendering semantics and share the same resolved 2D pipeline |
| `--theme light|dark` | Output theme |
| `--services <csv>` | Service metadata and label overrides |
| `--svg-legend-position top|right|bottom|left` | SVG legend placement |
| `--arrow-style thin|standard|triangle|stealth|arrow|diamond|oval|none` | SVG/PPTX/PDF/Excel Plan default used when a connection omits its arrowhead. Explicit DSL arrowheads on normal/traffic connections and explicit stroke widths take precedence; routes require effective arrowheads to be `none` |
| `--combine-frames` | Preserve the compatibility single-canvas/page form for SVG, PPTX, PDF, and Excel |

`aws-2.5d` and `topology` are recognized roadmap modes but currently return a
not-implemented error. Any other mode, format, theme, orientation, paper size,
arrow-style option, or SVG legend-position value outside its documented enum
returns an error.

`--arrow-style` belongs to the shared physical Plan used by SVG, PPTX, PDF, and
Excel. The editable Excalidraw, XYFlow, and Isoflow V1 outputs consume the
resolved DSL scene directly and therefore use the DSL connection defaults
instead.

Physical-page and PPTX flags:

| Flag | Description |
|---|---|
| `--paper A5|A4|A3|A2|A1|Letter|Legal|Tabloid` | Physical page/slide paper size |
| `--orientation portrait|landscape` | Physical page/slide orientation |
| `--paper-margin <inches>` | Margin applied before fitting |
| `--paper-margin-top/right/bottom/left <inches>` | Per-side margin override |
| `--px-per-inch <number>` | Layout scaling base |
| `--title`, `--author`, `--company`, `--subject` | PPTX presentation metadata |
| `--compression true|false` | PPTX compression |

`--title` sets package-level PPTX metadata. It is unrelated to the visible
`title` attribute on a page `<frame>` and never creates a frame tag.

## Structural Diff

```bash
xaligo diff before.xal after.xal -o output/architecture
```

The command compares parsed `.xal` data structures rather than source lines and
writes two SVG images:

- `output/architecture-removed.svg` renders the old document and highlights
  removed elements plus the old side of modified or moved elements in pale red.
- `output/architecture-added.svg` renders the new document and highlights
  added elements plus the new side of modified or moved elements in pale green.

XML formatting, comments, attribute order, parser-private attributes, and the
equivalent V1 forms with an omitted version or `version="1"` do not create a
diff when that version is on the document root. A `version` on an identified
child frame is visible page content and is compared normally, including the
literal value `1`. Matching prefers unique `name`, `ref`, and `id` values, then exact
subtrees, followed by deterministic order-aware structural matching. Give
elements an explicit `id`, `name`, or `ref` when moves must remain identifiable
across different parents.

`-o` is an output prefix, not a single output filename. A trailing `.svg` is
removed before `-removed.svg` and `-added.svg` are appended. The default prefix
is `output`. `--theme`, `--mode`, and `--px-per-inch` are applied identically to
both images. No difference is a successful result and still produces two
unhighlighted SVGs.

See the [structural diff sample](samples.md#structural-diff) for a complete
before/after pair and the generated images.

## Generate

Generate a starter `.xal` hierarchy:

```bash
xaligo generate xal -o generated.xal --paper A4 --orientation landscape
```

Useful generation flags:

| Flag | Default | Description |
|---|---:|---|
| `--clouds`, `--accounts`, `--regions`, `--azs` | `1`, `1`, `1`, `2` | AWS hierarchy counts |
| `--az-layout grid|staggered` | `grid` | Availability Zone placement |
| `--subnets` | `2` | Number of subnets |
| `--spacing vertical|horizontal|both` | `both` | Spacing direction |
| `--start top|left` | `top` | Drawing start position |
| `--paper` | `A4` | Paper size |
| `--orientation` | `landscape` | Page orientation |

Only `--output` is required. The generated file uses the canonical
`<xaligo version="1"><frames>...</frames></xaligo>` document envelope.

## Other Commands

| Command | Description |
|---|---|
| `xaligo diff <before.xal> <after.xal> -o <prefix>` | Render paired structural-diff SVGs |
| `xaligo validate <file.xal>` | Validate syntax, layout, and connection references |
| `xaligo render markdown <file.md>` | Embed rendered `xal` code blocks as SVG images into a Markdown file |
| `xaligo serve <file.xal\|file.md>` | Serve a live preview; `.xal` previews one combined SVG, `.md`/`.markdown` previews the full document with diagrams embedded inline; `--port` overrides the configured `serve.port` (default `8080`), and `--paper`/`--orientation` fix the preview to a physical page size |
| `xaligo add service --name <name> --file <file>` | Add a service icon |
| `xaligo add service --list <csv> --file <file>` | Bulk-add service icons |
| `xaligo init [-o <dir>]` | Generate a sample `.xal` file |
| `xaligo version` | Print version |
