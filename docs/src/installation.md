# Installation

## npm CLI

Install the native CLI package:

```bash
npm install -g @xaligo/xaligo
xaligo version
```

The package installs the `xaligo` command, runtime icon catalogs, and the PPTX
exporter WebAssembly artifact used by the CLI.

## Go Build

Build the CLI from source:

```bash
git clone https://github.com/xaligo/xaligo
cd xaligo
go mod tidy
make build
.bin/xaligo version
```

`make build` produces `.bin/xaligo` and builds the PPTX exporter WASM artifact.

## TypeScript and WebAssembly

Install the package in environments that need the TypeScript/WASM API:

```bash
npm install @xaligo/xaligo
```

The package exposes API entry points for extension hosts and browser-like
environments where spawning the native CLI is not available.

```typescript
import { renderPptxPlan } from "@xaligo/xaligo";

const pptx = await renderPptxPlan(planJson, { title: "Architecture" });
```

Build only the PPTX exporter WASM artifact:

```bash
make build-wasm
```

Package layout:

| Path | Purpose |
|---|---|
| `@xaligo/xaligo` | CLI binary plus TypeScript/WASM API exports |
| `external/pptx-exporter/` | Internal TypeScript/PptxGenJS build workspace |
