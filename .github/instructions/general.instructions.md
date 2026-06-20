---
applyTo: "**"
---

# xaligo — General Coding Guidelines

## Project Overview

`xaligo` is a Go CLI/Node package that converts a Vue-style custom DSL (`.xal` files) into
Excalidraw JSON and PPTX files, using Vuetify-style grid / padding / margin / container layout.
It also provides an `add service` command for appending AWS service icons to existing `.excalidraw` files.

## Module Information

```
module: github.com/ryo-arima/xaligo
Go:     1.22
Key dependencies: github.com/spf13/cobra v1.8.1
                  gopkg.in/yaml.v3 v3.0.1
```

## Directory Structure

```
xaligo/
├── cmd/
│   ├── main.go                   # Native CLI entry point
│   └── wasm/                     # Go/WASM bridge used by the Node package/PPTX exporter
├── pkg/
│   ├── command.go                # Root cobra command (wires subcommands)
│   └── controller/
│       ├── render.go             # xaligo render <input.xal> -o output.excalidraw
│       ├── init.go               # xaligo init -o <dir>  (generates sample.xal)
│       ├── version.go            # xaligo version
│       └── add.go                # xaligo add service [flags]
├── internal/
│   ├── model/
│   │   └── ast.go                # DSL AST: Document, Node
│   ├── parser/
│   │   └── parser.go             # XML-based DSL parser
│   ├── layout/
│   │   └── layout.go             # Vuetify-style layout engine
│   ├── excalidraw/
│   │   └── scene.go              # Excalidraw JSON builder (for render command)
│   ├── entity/
│   │   ├── scene.go              # Scene struct (for add command)
│   │   └── service.go            # ServiceEntry struct
│   ├── pptxplan/                 # PPTX paper scaling, routing, draw plan, legend data
│   ├── repository/
│   │   ├── builder.go            # MakeText / MakeImage element builders
│   │   ├── scene.go              # ReadScene / WriteScene
│   │   ├── icon.go               # SvgToDataURL / FileID / LoadFromCSV / SVGBGColor
│   │   ├── service_list.go       # ReadServiceList (CSV/TXT parser)
│   │   └── pptx.go               # Repository-layer WASM PPTX exporter invocation
│   └── config/
│       └── config.go             # Config struct + findProjectRoot + etc/resources/aws/app.yaml loading
├── examples/
│   └── sample.xal               # Sample DSL file
├── etc/
│   └── resources/
│       └── aws/
│           ├── app.yaml         # Path / legend size settings (defaults apply when absent)
│           ├── service-catalog.csv  # Full SVG icon catalog
│           ├── svg/             # AWS icon SVGs (Architecture-Service/Group/Resource/Category-Icons)
│           └── templates/
│               ├── excalidraw/  # Per-AWS-group-tag templates (.excalidraw)
│               └── xal/         # Per-AWS-group-tag templates (.xal)
├── scripts/
│   ├── gen_service_catalog.py   # Regenerate service-catalog.csv
│   └── gen_group_templates.py   # Regenerate etc/resources/aws/templates/{excalidraw,xal}/
├── Makefile
├── packages/
│   └── xaligo/                   # npm package, WASM assets, PPTX export support
├── go.mod / go.sum
└── README.md
```

## Architecture Guidelines

- Use `roadmap.instructions.md` as the long-term implementation direction when
  prioritizing renderer, routing, preview, extension, and export work.
- **cmd → pkg/command.go → pkg/controller/ → internal/**: Keep dependencies unidirectional.
- `internal/` packages are only referenced from `pkg/`; never directly from `cmd/`.
- Each `controller` file exports an `Init<Cmd>Cmd() *cobra.Command` factory function, registered in `pkg/command.go`.
- Business logic stays in `internal/`; cobra flag handling is the responsibility of the `controller` layer.
- `xaligo render --format pptx` must invoke a WASM-compiled PPTX exporter from the
  repository layer (`internal/repository/pptx.go`).
- PPTX geometry is resolved by Go `internal/pptxplan`; the WASM exporter only
  turns the resolved plan into PPTX bytes.
- Repository/controller Go code must not contain a PPTX/OOXML writer. Keep Go
  limited to plan construction, WASM invocation, and writing returned bytes.
- Avoid `goja` and V8 for PPTX export execution.
- Avoid a long-term Node.js subprocess dependency in repository-layer PPTX
  export. Node may remain a development/build tool while the WASM exporter is
  being prepared.
- Keep HTTP, gRPC, and stdin/stdout RPC as future alternatives only unless the
  architecture is intentionally changed.

## Coding Conventions

- Follow standard Go `gofmt` / `golint` style.
- Package names are lowercase single words (e.g., `controller`, `repository`, `entity`).
- Wrap errors with `fmt.Errorf("<context>: %w", err)` and return them to the caller.
- Do not use `panic`. Always return errors as `error`.
- Represent Excalidraw elements as `map[string]interface{}` (for compatibility with the existing format).

## Configuration File (etc/resources/aws/app.yaml)

Loaded from `etc/resources/aws/app.yaml` at the project root (directory containing `go.mod`).
When absent, all defaults are used — the file is not required.

```yaml
paths:
  asset_package:        etc/resources/aws/svg
  service_catalog_csv:  etc/resources/aws/service-catalog.csv
  output_frames:        output/aws-frames
legend:
  offset_x:  120
  offset_y:  0
  icon_size: 32
  font_size: 12
item:
  icon_size: 32   # default max icon size for <item> elements (px). Overridable with <frame item-size="N">
```

## Icon Label Resolution

When rendering `<item>` icons, the short label below each icon is determined in the following priority order:

1. **`Abbreviation` column in services.csv** — used when `render --format excalidraw --services <csv>` is invoked and the entry for that catalog ID has a non-empty `Abbreviation`.
2. **Built-in `itemAbbreviations` table** (`internal/entity/service.go`) — fallback for any ID not covered by services.csv, and the only source when using `render` directly.

This means `services.csv` is the single source of truth for icon labels in `render --format excalidraw` workflows.
The `OfficialName` column is used for full-name legend text — never as an icon label.
In PPTX export, legend entries are rendered on separate 4-column legend slides.

## CLI Command Reference

| Command | Description |
|---|---|
| `xaligo render <file.xal> -o <out.excalidraw>` | Convert DSL to Excalidraw JSON |
| `xaligo init [-o <dir>]` | Generate `sample.xal` |
| `xaligo version` | Print version |
| `xaligo add service --name <name> --file <file>` | Add a single AWS service icon |
| `xaligo add service --list <csv> --file <file>` | Bulk-add AWS service icons |
| `xaligo generate xal --clouds N --accounts N --regions N --azs N --az-layout grid\|staggered --subnets N --spacing vertical\|horizontal\|both --start top\|left --paper A4 --orientation portrait\|landscape -o out.xal` | Generate a .xal for an AWS infrastructure hierarchy |
| `xaligo render <file.xal> --format excalidraw -o <out.excalidraw> --services <csv>` | Convert .xal to .excalidraw |
| `xaligo render <file.xal> --format pptx -o <out.pptx> --services <csv> --paper A3 --orientation landscape` | Convert .xal to PPTX when `pptx_exporter.wasm` is configured |

The npm/WASM API can generate PPTX through PptxGenJS. Native CLI PPTX output
requires the separate WASI `pptx_exporter.wasm`; Excalidraw and SVG do not.

## Build & Test

```bash
make build   # build .bin/xaligo
make build-wasm # build packages/xaligo/wasm/xaligo.wasm
npm run build --workspace packages/xaligo
make run     # examples/sample.xal → output/sample.excalidraw
make clean   # remove .bin/ and output/
go test ./...            # run all tests
go build ./...           # check for build errors
```
