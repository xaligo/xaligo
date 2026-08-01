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
<xaligo version="1">
  <frames>
    <frame id="getting-started" title="Getting Started" version="2026.07"
           width="1122" height="794" class="pa-4" margin-top="48">
      <metadata align="right" width="156" key-width="56" font-size="12" />

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
  </frames>
</xaligo>
```

Render it:

```bash
xaligo render diagram.xal --format svg -o diagram.svg
```

SVG is the default, so this is equivalent:

```bash
xaligo render diagram.xal -o diagram.svg
```

PPTX export uses the bundled or configured WASM exporter:

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx \
  --paper A3 --orientation landscape \
  --paper-margin-top 0.75 --paper-margin-bottom 0.75
```

## Render Multiple Frames

An identified child frame is one physical output page. A multi-frame document
therefore creates one SVG file per frame or one PPTX slide per frame. With
`overview` and `detail`
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

## Render Markdown

Fenced `xal` blocks can be replaced with generated SVG image references:

```bash
xaligo render markdown guide.md
```

See [Markdown rendering](rendering/markdown.md) for output-location options.

## Validate

Run validation before rendering large diagrams:

```bash
xaligo validate diagram.xal
```

Validation catches malformed XML, invalid layout references, duplicate
connection references, nested connection tags, and unresolved endpoints.
