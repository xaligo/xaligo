# xaligo

![Gopher and Crawfish by a river](docs/images/Gemini_Generated_Image_1lec2o1lec2o1lec.png)

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![CC BY 3.0](https://img.shields.io/badge/illustration-CC%20BY%203.0-lightgrey.svg)](https://creativecommons.org/licenses/by/3.0/)

> The Go Gopher was designed by [Renée French](https://reneefrench.blogspot.com/).  
> Licensed under [CC BY 3.0](https://creativecommons.org/licenses/by/3.0/).  
> This illustration is a derivative work inspired by the original Go Gopher design.

A Diagram-as-Code engine that renders the Vue-style `.xal` DSL to Excalidraw,
SVG, and PPTX. It includes Vuetify-style layout, AWS/network icon catalogs,
orthogonal routing, route/traffic layers, line jumps, and automatic junctions.

> **For automated agents (codex-cli, GPT, etc.):** see [docs/agents/general.md](docs/agents/general.md) for detailed step-by-step procedures.

## Installation

### CLI (Go binary)

```bash
git clone https://github.com/ryo-arima/xaligo
cd xaligo
go mod tidy
make build        # produces .bin/xaligo
```

### npm / WebAssembly (TypeScript / VS Code extension)

A WebAssembly package is provided for environments where spawning a child process is not feasible
(e.g., a VS Code web extension or an extension running in the extension host).

```bash
npm install @ryo-arima/xaligo
```

After building the WASM artifact (`make build-wasm`), use it from TypeScript:

```typescript
import { loadXaligo } from "@ryo-arima/xaligo";

const xaligo = await loadXaligo();                         // loads xaligo.wasm on first call

// Convert .xal DSL string → Excalidraw JSON string
const json = await xaligo.render(xalSource);

// Convert with a services CSV for the legend
const json = await xaligo.renderWithServices(xalSource, servicesCsv);

// Render directly to PPTX bytes through PptxGenJS
const pptx = await xaligo.renderPptx(xalSource, { theme: "dark" });
```

**Build the WASM artifact:**

```bash
make build-wasm   # outputs packages/xaligo/wasm/xaligo.wasm
                  #         and packages/xaligo/wasm/wasm_exec.js
```

## npm Package Layout

```
packages/
└── xaligo/   @ryo-arima/xaligo   — WASM + TypeScript wrapper
```

## Commands

| Command | Description |
|---|---|
| `xaligo render <file.xal> --format excalidraw -o <out.excalidraw> [--mode standard\|network\|aws] [--theme light\|dark] [--services <csv>]` | Convert .xal → .excalidraw |
| `xaligo render <file.xal> --format svg -o <out.svg> [--mode standard\|network\|aws] [--theme light\|dark] [--services <csv>]` | Convert .xal → SVG |
| `xaligo render <file.xal> --format pptx -o <out.pptx> [--mode standard\|network\|aws] [--theme light\|dark] [--services <csv>] [pptx flags]` | Convert .xal → .pptx via the WASM PPTX exporter |
| `xaligo generate xal [flags] -o <out.xal>` | Auto-generate an AWS infrastructure hierarchy .xal |
| `xaligo validate <file.xal>` | Validate .xal syntax and layout |
| `xaligo add service --name <name> --file <file>` | Add a single AWS service icon to an existing file |
| `xaligo add service --list <csv> --file <file>` | Bulk-add AWS service icons to an existing file |
| `xaligo init [-o <dir>]` | Generate a sample.xal |
| `xaligo version` | Print version |

`mode` controls layout/presentation semantics while `format` controls the output
container. `standard`, `network`, and `aws` currently share the resolved 2D
pipeline; `aws-2.5d` and `topology` are reserved for later roadmap phases.

Native CLI PPTX output requires the separately configured WASI exporter
`pptx_exporter.wasm`. The npm/WASM API can generate PPTX now through
PptxGenJS; Excalidraw and SVG do not require the PPTX exporter.

### generate xal flags

```
--clouds N                         Number of AWS Cloud blocks (default 1)
--accounts N                       Number of Account blocks (default 1)
--regions N                        Number of Region blocks (default 1)
--azs N                            Number of Availability Zones (default 2)
--az-layout grid|staggered         AZ placement style (default grid)
--subnets N                        Number of subnets (default 2)
--spacing vertical|horizontal|both Spacing direction (default both)
--start top|left                   Drawing start position (default top)
--paper A4                         Paper size
--orientation portrait|landscape   Paper orientation (default landscape)
-o <file>                          Output file path
```

## Quick Start

### Option A — Start from a hand-crafted diagram

```bash
# 1. Find the catalog IDs for the services you need
grep -i "ec2\|rds\|cloudfront" etc/resources/aws/service-index.csv

# 2. Create services.csv (id,OfficialName,Abbreviation,Summary,Usage,Notes)
#    See examples/services.csv for reference

# 3. Write your .xal layout file
#    See examples/sample.xal for a 3-tier architecture example

# 4. Generate
mkdir -p output
.bin/xaligo render examples/sample.xal \
  --format excalidraw \
  -o output/sample.excalidraw \
  --services examples/services.csv

# Optional native CLI PPTX export (requires pptx_exporter.wasm)
.bin/xaligo render examples/sample.xal --format pptx \
  -o output/sample.pptx --services examples/services.csv
```

PPTX flags: `--title`, `--author`, `--company`, `--subject`, `--compression true|false`, `--px-per-inch`.

## Go API

The root package exposes the same parse/layout/render pipeline used by the CLI:

```go
svg, err := xaligo.RenderSVG(ctx, source, xaligo.RenderOptions{
    Mode:  xaligo.ModeNetwork,
    Theme: "dark",
})

err = xaligo.Validate(ctx, source)
```

`Render`, `RenderExcalidraw`, `RenderSVG`, and `RenderPPTX` are available now.
`RenderXYFlow` and `RenderIsoflow` return `ErrNotImplemented` until their
roadmap phases are implemented.

### Option B — Auto-generate an AWS hierarchy

```bash
# Generate a .xal for an AWS configuration
.bin/xaligo generate xal --clouds 1 --accounts 1 --regions 2 --azs 2 \
  --az-layout staggered --subnets 2 --spacing both --start top \
  --paper A4 --orientation landscape -o output/infra.xal

# Convert to .excalidraw
.bin/xaligo render output/infra.xal \
  --format excalidraw \
  -o output/infra.excalidraw \
  --services examples/services.csv
```

Import the generated `.excalidraw` file into [Excalidraw](https://excalidraw.com).

## .xal DSL

### Root structure

```xml
<frame width="1122" height="794" class="pa-4" item-size="48">
  <!-- place elements here -->
</frame>
```

### Layout tags

| Tag | Description |
|---|---|
| `<frame>` | Root tag. Specifies width, height, and padding |
| `<container>` | Vertical stack container (`layout="horizontal"` for horizontal layout) |
| `<row>` | 12-column grid row |
| `<col>` | Column inside `<row>` (`span` sets width) |
| `<spacer>` / `<blank>` | Empty layout cell. Use in rows, stacks, or item grids to reserve blank space |

### AWS group tags

Tags rendered with AWS architecture diagram group border styles.

| Tag | Display name | Border color |
|---|---|---|
| `<aws-cloud>` | AWS Cloud | `#000000` |
| `<aws-account>` | AWS Account | `#E7008A` |
| `<region>` | Region | `#00A1C9` |
| `<availability-zone>` | Availability Zone | `#00A1C9` |
| `<vpc>` | VPC | `#8C4FFF` |
| `<public-subnet>` | Public Subnet | `#3F8624` |
| `<private-subnet>` | Private Subnet | `#00A1C9` |
| `<security-group>` | Security Group | `#CC0000` |
| `<auto-scaling-group>` | Auto Scaling Group | `#E7601B` |
| `<server-contents>` | Server Contents | `#7A7C7F` |
| `<corporate-data-center>` | Corporate Data Center | `#7A7C7F` |
| others | See xal-spec for details | — |

### `<item>` tag

Embeds a catalog icon by specifying its ID from `service-catalog.csv`.
Omitting or leaving `id` empty makes the element a spacer (no icon rendered).

Tabler Icons are available through the same catalog mechanism as AWS icons.
Their stable catalog IDs start at `100000`; see
`etc/resources/aws/service-index.csv` for the name-to-ID mapping.

```xml
<item id="104109" /> <!-- Tabler: server -->
```

Refresh the vendored SVG files and catalog entries with
`npm run import:tabler-icons` after updating `@tabler/icons`.
Tabler Icons are distributed under the MIT license included with the assets.

Yamaha Network Diagram Icons use the same catalog mechanism with IDs starting
at `200000`. The original SVG files are redistributed unchanged under CC BY-ND
4.0; attribution is included alongside the assets.

```xml
<item id="200000" /> <!-- Yamaha network icon -->
```

Refresh them from Yamaha's official ZIP with `npm run import:yamaha-icons`.
See the bundled `ATTRIBUTION.txt` for Yamaha's CC BY-ND 4.0 terms.

```xml
<public-subnet title="Public Subnet">
  <item id="1178" />   <!-- with icon -->
  <item />             <!-- spacer (empty slot) -->
  <item id="1189" />   <!-- with icon -->
</public-subnet>
```

### `<connection>` tag

Draws an elbowed arrow between `<item>` elements. Must be a direct child of `<frame>`.

```xml
<frame width="1122" height="794" class="pa-4">
  <!-- ... layout elements ... -->

  <!-- list connections at the end of frame -->
  <connection src="1178" dst="1189" />
</frame>
```

| Attribute | Description |
|---|---|
| `src` | Catalog ID of the arrow start item |
| `dst` | Catalog ID of the arrow end item |
| `kind` | `route` (thin structural path) or `traffic` (strong directional flow) |
| `color` | Per-line CSS/hex stroke color |
| `stroke-width` | Per-line stroke width; defaults to 1 for route and 2 for traffic |
| `stroke-style` | `solid`, `dashed`, or `dotted` |
| `start-arrowhead` / `end-arrowhead` | Independently set each end to `none`, `arrow`, `triangle`, `stealth`, `diamond`, or `oval` |
| `arrowhead` | Backward-compatible alias for `end-arrowhead` |

Connections are always rendered in **elbowed (right-angle)** style. Route lines
use circular endpoints by default and are drawn below traffic lines; traffic
sharing a route is assigned a separate candidate lane where space permits.
Routes that fan out from or converge on the same side of an item automatically
share a short trunk and render a circular junction at the branch point.

At an interior crossing, SVG and PPTX place a 6px background-colored rectangle
between the lower and upper lines. The color follows the uppermost opaque
container at that point and otherwise uses the slide background. The mask stays
smaller than the standard 8px lane gap so adjacent lines are not erased.
Endpoint touches and parallel overlaps are not treated as crossings.
Start and end points connect to the **midpoint of the nearest edge** of the icon image or label text element.  
When the connection direction is downward, the label text element edge is used; otherwise the icon image edge is used.  
Edges are fixed with Excalidraw's `fixedPoint` binding, so arrows snap correctly when the file is opened.

### Key attributes

| Attribute | Target | Description |
|---|---|---|
| `title` | any | Display label |
| `layout="horizontal"` | container tags | Arrange children horizontally |
| `layout="staggered"` | AWS group tags | Stack children with depth offset |
| `row="N"` | child in vertical stack | Height ratio (flex-grow equivalent) |
| `col="N"` | child in `layout="horizontal"` | Width ratio (flex-grow equivalent) |
| `width="N"` / `height="N"` | non-root child | Fixed child size in px |
| `content-width="N"` / `content-height="N"` | containers/groups | Shrink the inner layout area, leaving blank space around it |
| `align="top-left"` etc. | containers/groups and item grids | Align inner content or item grid (`top|middle|bottom` + `left|center|right`; item grids also support `spread`) |
| `gap="N"` | container tags | Child spacing (px) |
| `border="none"` | any | Hide border |
| `visible="false"` | any | Hide only this component (children are still rendered) |
| `item-size="N"` | `<frame>` | Override icon size for all `<item>` elements in this file (px) |
| `class` | any | Vuetify-style spacing class |

### Spacing classes

Unit is `8px`. Multiple classes are space-separated: `class="pa-4 ml-2"`

| Pattern | Description |
|---|---|
| `pa-{n}` / `ma-{n}` | padding / margin all sides |
| `px-{n}` / `py-{n}` | padding left+right / top+bottom |
| `mx-{n}` / `my-{n}` | margin left+right / top+bottom |
| `pt/pr/pb/pl-{n}` | padding per side |
| `mt/mr/mb/ml-{n}` | margin per side |

On the root `<frame>`, margin is treated as outer content whitespace: the paper
frame keeps its full size, and the diagram content is inset from the paper edge.

## Sample DSL

See [examples/sample.xal](examples/sample.xal) for a full 3-tier architecture example.
The snippet below shows the essential structure:

```xml
<frame width="1200" height="820" class="pa-4">
  <aws-cloud title="AWS Cloud">
    <aws-account title="Production Account">
      <region title="ap-northeast-1">
        <vpc title="VPC (10.0.0.0/16)">
          <availability-zone title="AZ: ap-northeast-1a">

            <!-- Tier 1: Presentation (public) -->
            <public-subnet title="Tier 1 — Presentation" row="3">
              <item id="1179" />  <!-- Route 53 -->
              <item id="1581" />  <!-- Internet Gateway -->
              <item id="1182" />  <!-- ELB -->
            </public-subnet>

            <!-- Tier 2: Application (private) -->
            <private-subnet title="Tier 2 — Application" row="2">
              <item id="27" />    <!-- EC2 -->
              <item id="1582" />  <!-- NAT Gateway -->
            </private-subnet>

            <!-- Tier 3: Data (private) -->
            <private-subnet title="Tier 3 — Data" row="2">
              <item id="110" />   <!-- Aurora -->
              <item id="113" />   <!-- ElastiCache -->
            </private-subnet>

          </availability-zone>
        </vpc>
      </region>
    </aws-account>
  </aws-cloud>

  <!-- connections must be direct children of <frame>, placed last -->
  <connection src="1179" dst="1182" />
  <connection src="1182" dst="27" />
  <connection src="27" dst="110" />
  <connection src="27" dst="113" />
</frame>
```

## Configuration

You can customize paths and defaults in `etc/resources/aws/app.yaml` (all values are optional; defaults are used when the file is absent).

```yaml
paths:
  asset_package:       etc/resources/aws/svg
  service_catalog_csv: etc/resources/aws/service-catalog.csv
  output_frames:       output/aws-frames
  pptx_exporter_wasm:  packages/xaligo/wasm/pptx_exporter.wasm

legend:
  offset_x:  120
  offset_y:  0
  icon_size: 32
  font_size: 12

item:
  icon_size: 48   # default icon size for <item> elements (px)
```

## Build & Test

```bash
make build        # build .bin/xaligo (native Go binary)
make build-wasm   # build WASM artifact + copy wasm_exec.js into packages/xaligo/wasm/
make run          # examples/sample.xal → output/sample.excalidraw
make clean        # remove .bin/, output/, and WASM artifacts
go test ./...
```

## License

[MIT](LICENSE)
