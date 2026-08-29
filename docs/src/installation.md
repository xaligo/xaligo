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

Building the native CLI requires Go 1.26, Rust 1.85 or newer, a C compiler for
cgo, Node.js 24, npm, and Javy 9. Install the repository dependencies, then
build the CLI:

```bash
git clone https://github.com/xaligo/xaligo
cd xaligo
npm ci --ignore-scripts
make build
.bin/xaligo version
```

`make build` compiles the Rust engine as a static library, links it into
`.bin/xaligo` through cgo, and builds the PPTX exporter WASM artifact. The Rust
library is part of the executable; no engine process or dynamic library is
installed alongside it.

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
make build-exporter
```

Package layout:

| Path | Purpose |
|---|---|
| `@xaligo/xaligo` | CLI binary plus TypeScript/WASM API exports |
| `external/exporter/` | Internal Rust PPTX exporter linked into the engine static library |
