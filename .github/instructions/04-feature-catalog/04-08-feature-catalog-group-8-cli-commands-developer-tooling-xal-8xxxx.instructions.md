---
applyTo: ".github/instructions/manual/**"
---

# 04.08 Feature catalog: Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)

## Group 8 — CLI Commands & Developer Tooling (`XAL-8xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-8000010 | `xaligo render` command | Implemented | Renders a `.xal` source file into any supported output format via `--format`. |
| XAL-8000020 | `xaligo validate` command | Implemented | Validates `.xal` syntax, layout, and connection references without producing output. |
| XAL-8000030 | `xaligo serve` command | Implemented | Serves a live-reloading preview over HTTP, polling the source file and re-rendering on change; a `.xal` source previews one combined SVG canvas, a `.md`/`.markdown` source previews the full Markdown document with rendered `xal` code blocks embedded inline; `--paper`/`--orientation` fix the preview to a specific physical page size and orientation at server startup. |
| XAL-8000040 | `xaligo diff` command | Implemented | Renders paired removed/added structural-diff SVGs comparing two `.xal` documents. |
| XAL-8000050 | `xaligo generate xal` command | Implemented | Scaffolds a `.xal` file with a configurable AWS Cloud/Account/Region/VPC/AZ/Subnet hierarchy. |
| XAL-8000060 | `xaligo add service` (single mode) | Implemented | Appends one named AWS service icon and a legend entry into an existing `.excalidraw` file. |
| XAL-8000070 | `xaligo add service --list` (batch mode) | Implemented | Bulk-adds every service in a `services.csv` list into an existing `.excalidraw` file. |
| XAL-8000080 | `xaligo init` command | Implemented | Writes a minimal starter `sample.xal` file to help new users get started. |
| XAL-8000090 | `xaligo version` command | Implemented | Reports the resolved build/release version, with an embedded-value, env-var, and `VERSION`-file fallback chain. |
| XAL-8000100 | Man-page-style detailed CLI help | Implemented | Every command and subcommand exposes a detailed `Long` description and runnable `Examples` in its `--help` output, mirrored in the CLI reference documentation. |
| XAL-8000110 | Structured CLI logging | Implemented | Stable `share.NewMCode`-tagged debug/info/error log lines throughout controllers and use cases for traceable CLI diagnostics. |
| XAL-8000120 | VS Code extension integration | Planned | A separate-repository VS Code extension providing `.xal` syntax highlighting, a live Preview Panel, and source-positioned diagnostics, built on this repository's reusable render/validate APIs and HTTP/SSE preview protocol. |
| XAL-8000130 | Multi-view live preview | Planned | Extending `xaligo serve`/preview beyond SVG-first to Excalidraw, XYFlow, Isoflow, and 2.5D preview views. |
| XAL-8000140 | `xaligo render markdown` command | Implemented | Reads a Markdown file, renders every fenced ` ```xal ` code block to SVG through the shared render pipeline, and writes a new Markdown file with a `![](path.svg)` image reference per rendered frame in place of each code block; generated SVGs default beside the source Markdown file and `--svg-dir`/`--output` override the locations; `--paper`/`--orientation` fit each rendered diagram to a physical page size, matching `render --format svg`. |
