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
<frame width="1122" height="794" class="pa-4">
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

PPTX export uses the bundled or configured WASM exporter:

```bash
xaligo render diagram.xal --format pptx -o diagram.pptx \
  --paper A3 --orientation landscape \
  --paper-margin-top 0.75 --paper-margin-bottom 0.75
```

## Validate

Run validation before rendering large diagrams:

```bash
xaligo validate diagram.xal
```

Validation catches malformed XML, invalid layout references, duplicate
connection references, nested connection tags, and unresolved endpoints.
