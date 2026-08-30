# Installation

## npm CLI

Install the native CLI package:

```bash
npm install -g @xaligo/xaligo
xaligo version
```

The package installs the `xaligo` launcher and runtime icon catalogs. Its
postinstall step downloads the native CLI binary and verifies its SHA-256
checksum. The Rust V2 engine and PPTX exporter are already linked into that
binary; no runtime WASM or dynamic-library artifact is installed.

## Go Build

Building the native CLI requires Go 1.26, Rust 1.85 or newer, and a C compiler
for cgo:

```bash
git clone https://github.com/xaligo/xaligo
cd xaligo
make build
.bin/xaligo version
```

`make build` compiles the Rust engine and PPTX exporter into one static library
and links it into `.bin/xaligo` through cgo. The Rust code is part of the
executable; no engine process, dynamic library, or WASM artifact is installed
alongside it.

## Native Exporter Development

Build and test the standalone Rust PPTX exporter workspace when changing that
adapter:

```bash
make build-exporter
cargo test --manifest-path test/unit/external/exporter/Cargo.toml --locked
```

The npm package exposes the native `xaligo` command, not a TypeScript/WASM
rendering API. The repository still contains `cmd/wasm` as a source-only legacy
V1 compatibility adapter, but it cannot execute the native V2 engine and is not
part of the npm or release artifact set.

Package layout:

| Path | Purpose |
|---|---|
| `@xaligo/xaligo` | Native CLI launcher, installer, and runtime assets |
| `external/exporter/` | Internal Rust PPTX exporter linked into the engine static library |
