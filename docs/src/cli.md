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
| `--mode standard|network|aws` | Rendering mode. Current modes share the same resolved 2D pipeline |
| `--theme light|dark` | Output theme |
| `--services <csv>` | Service metadata and label overrides |
| `--svg-legend-position top|right|bottom|left` | SVG legend placement |

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
| `xaligo validate <file.xal>` | Validate syntax, layout, and connection references |
| `xaligo serve <file.xal>` | Serve an SVG live preview |
| `xaligo add service --name <name> --file <file>` | Add a service icon |
| `xaligo add service --list <csv> --file <file>` | Bulk-add service icons |
| `xaligo init [-o <dir>]` | Generate a sample `.xal` file |
| `xaligo version` | Print version |
