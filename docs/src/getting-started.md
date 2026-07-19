# Getting Started

## Install

Install the CLI from npm:

```bash
npm install -g @xaligo/xaligo
xaligo version
```

Or build from source:

```bash
git clone https://github.com/xaligo/xaligo
cd xaligo
go mod tidy
make build
.bin/xaligo version
```

## Render A Diagram

Create a `.xal` file:

```xml
<frame version="1" width="1122" height="794" class="pa-4">
  <aws-cloud id="cloud" title="AWS Cloud">
    <region id="region-apne1" title="ap-northeast-1">
      <vpc id="prod-vpc" title="Production VPC">
        <public-subnet id="public" title="Public Subnet">
          <item id="1178" name="web" />
        </public-subnet>
        <private-subnet id="app" title="Application Subnet">
          <item id="27" name="app-server" />
        </private-subnet>
      </vpc>
    </region>
  </aws-cloud>

  <connections kind="traffic" color="#2563eb">
    <connection src="web" dst="app-server" />
  </connections>
</frame>
```

Render it:

```bash
xaligo render diagram.xal --format excalidraw -o diagram.excalidraw
xaligo render diagram.xal --format svg -o diagram.svg
```

PDF and Excel use the same source:

```bash
xaligo render diagram.xal --format pdf -o diagram.pdf
xaligo render diagram.xal --format excel -o diagram.xlsx
```

PPTX export uses the bundled or configured WASM exporter:

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx \
  --paper A3 --orientation landscape \
  --paper-margin-top 0.75 --paper-margin-bottom 0.75
```

## Render Multiple Frames

An identified child frame is one physical output page. A multi-frame document
therefore creates one SVG file per frame, one PPTX slide per frame, one PDF page
per frame, or one Excel worksheet per frame. With `overview` and `detail`
frames, this SVG command:

```bash
xaligo render diagram.xal --format svg -o diagram.svg
```

writes `diagram-overview.svg` and `diagram-detail.svg`. A one-frame document
still writes the exact `-o` path. Use the compatibility form when a single
canvas is required:

```bash
xaligo render diagram.xal --format svg -o diagram.svg --combine-frames
```

See [Cross-Frame Page Links](examples/page-links.md) for a complete multi-frame
source and the matching `to ...` / `from ...` page-link labels.

## Validate

Run validation before rendering large diagrams:

```bash
xaligo validate diagram.xal
```

Validation catches malformed XML, invalid layout references, duplicate
connection references, nested connection tags, and unresolved endpoints.
