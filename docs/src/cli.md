# Command Line

## Render

```bash
xaligo render <file.xal> --format <format> -o <output>
```

Supported formats:

| Format | Output |
|---|---|
| `excalidraw` | Editable Excalidraw JSON |
| `svg` | Standalone SVG |
| `pptx` | PowerPoint presentation |
| `xyflow` | React Flow / XYFlow JSON |
| `isoflow` | Isoflow-compatible model JSON |

Common render flags:

| Flag | Description |
|---|---|
| `--mode standard|network|aws` | Accepted V1 rendering mode. These three values currently have identical V1 rendering semantics and share the same resolved 2D pipeline |
| `--theme light|dark` | Output theme |
| `--services <csv>` | Service metadata and label overrides |
| `--svg-legend-position top|right|bottom|left` | SVG legend placement |
| `--arrow-style thin|standard|triangle|stealth|arrow|diamond|oval|none` | SVG/PPTX Plan default used when a connection omits its arrowhead. Explicit DSL arrowheads on normal/traffic connections and explicit stroke widths take precedence; routes require effective arrowheads to be `none` |

`aws-2.5d` and `topology` are recognized roadmap modes but currently return a
not-implemented error. Any other mode, format, theme, orientation, paper size,
arrow-style option, or SVG legend-position value outside its documented enum
returns an error.

`--arrow-style` belongs to the shared physical Plan used by SVG and PPTX. The
editable Excalidraw, XYFlow, and Isoflow V1 outputs consume the resolved DSL
scene directly and therefore use the DSL connection defaults instead.

PPTX-specific flags:

| Flag | Description |
|---|---|
| `--paper A5|A4|A3|A2|A1|Letter|Legal|Tabloid` | Slide paper size |
| `--orientation portrait|landscape` | Slide orientation |
| `--paper-margin <inches>` | Margin applied before fitting |
| `--paper-margin-top/right/bottom/left <inches>` | Per-side margin override |
| `--px-per-inch <number>` | Layout scaling base |
| `--title`, `--author`, `--company`, `--subject` | Presentation metadata |
| `--compression true|false` | PPTX compression |

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
diff. Matching prefers unique `name`, `ref`, and `id` values, then exact
subtrees, followed by deterministic order-aware structural matching. Give
elements an explicit `id`, `name`, or `ref` when moves must remain identifiable
across different parents.

`-o` is an output prefix, not a single output filename. A trailing `.svg` is
removed before `-removed.svg` and `-added.svg` are appended. The default prefix
is `output`. `--theme`, `--mode`, and `--px-per-inch` are applied identically to
both images. No difference is a successful result and still produces two
unhighlighted SVGs.

## Generate

Generate a starter `.xal` hierarchy:

```bash
xaligo generate xal -o generated.xal --paper A4 --orientation landscape
```

Useful generation flags:

| Flag | Description |
|---|---|
| `--clouds`, `--accounts`, `--regions`, `--azs` | AWS hierarchy counts |
| `--az-layout grid|staggered` | Availability Zone placement |
| `--subnets` | Number of subnets |
| `--spacing vertical|horizontal|both` | Spacing direction |
| `--start top|left` | Drawing start position |

## Other Commands

| Command | Description |
|---|---|
| `xaligo diff <before.xal> <after.xal> -o <prefix>` | Render paired structural-diff SVGs |
| `xaligo validate <file.xal>` | Validate syntax, layout, and connection references |
| `xaligo serve <file.xal>` | Serve an SVG live preview |
| `xaligo add service --name <name> --file <file>` | Add a service icon |
| `xaligo add service --list <csv> --file <file>` | Bulk-add service icons |
| `xaligo init [-o <dir>]` | Generate a sample `.xal` file |
| `xaligo version` | Print version |
