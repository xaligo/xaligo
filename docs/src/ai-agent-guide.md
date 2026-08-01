# AI Agent Guide

This page is the compact entry point for AI assistants, coding agents, and
other generated-code tools working with xaligo.

## Read First

Use this order when collecting context:

1. [Introduction](introduction.md)
2. [Getting Started](getting-started.md)
3. [.xal DSL Overview](xal/overview.md)
4. [Layout](xal/layout.md)
5. [Connections](xal/connections.md)
6. [Examples](examples/index.md)
7. [Development](development.md)

For repository work, also read the applicable files under
`.github/instructions/` before planning or editing. Those files are the
source of truth for implementation constraints.

## Core Model

xaligo has one source DSL: `.xal`. Keep every renderer on the same pipeline:

```text
.xal -> parser -> layout -> shared scene/plan -> encoder
```

Do not add format-specific parsers or hidden alternate layout paths. SVG and
PPTX consume the shared parser, layout, and document plan; Markdown embeds SVG
artifacts. Other outputs are outside the current product contract.

## Important Paths

| Path | Purpose |
|---|---|
| `internal/usecase/` | Parser, layout, validation, shared scene and plan logic |
| `internal/repository/` | Filesystem, catalog, preview, and output-format adapters |
| `cmd/` | Native CLI entry point |
| `cmd/wasm/` | JavaScript/WASM adapter |
| `external/engine/` | Rust layout/SVG engine workspace |
| `external/pptx-exporter/` | TypeScript PPTX exporter implementation |
| `docs/src/examples/samples/` | Source `.xal` and CSV examples |
| `docs/src/examples/previews/` | Rendered SVG previews used by the book |
| `etc/resources/aws/` | Catalogs, icons, templates, and attribution |

Generated outputs belong in `output/` or another temporary directory, not in
the documentation source tree unless they are intentional documentation assets
such as SVG previews.

## Useful Examples

- [Line Variants](examples/line-variants.md) shows route, traffic, and style
  overrides.
- [Route and Traffic Separation](examples/route-traffic.md) shows structural
  routes plus directional traffic lanes.
- [Automatic Route Junctions](examples/junctions.md) shows shared route trunks
  and fan-out junctions.

The corresponding source files live in `docs/src/examples/samples/`.

## Common Commands

```bash
go test ./...
go build ./...
mdbook build docs
git diff --check
```

Render a documentation example:

```bash
.bin/xaligo render docs/src/examples/samples/route-traffic.xal \
  --format svg \
  --mode network \
  -o output/route-traffic.svg
```

## Editing Rules

- Preserve the shared `.xal` pipeline.
- Keep mode and output format independent.
- Keep render geometry and routing in shared use-case layers.
- Keep generated binaries, package artifacts, and ad hoc renders out of the
  repository.
- Update documentation links when moving examples or changing CLI paths.
- Prefer focused tests or render checks for behavior changes.
