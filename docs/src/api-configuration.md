# API and Configuration

## Public Boundaries

External consumers should use the CLI, the preview protocol, or the
TypeScript/WASM package. In-repository adapters use `internal/usecase` for the
same parse, layout, render, and validation pipeline.

Available render paths include:

- `Render`
- `RenderSVG`
- `RenderArtifacts`
- `RenderPPTX`
- `BuildPPTXPlan`
- `Validate`
- `Diagnose`

`RenderArtifacts` is the SVG multi-artifact boundary: it returns one ordered
artifact per identified frame by default. `RenderSVG` remains convenient for a
single frame or for `RenderOptions.CombineFrames`; it reports an error instead
of silently discarding extra frame artifacts. PPTX is a container format, so
`RenderPPTX` returns one byte sequence containing one slide per frame. `Render`
accepts only `FormatSVG` and `FormatPPTX`; SVG is the default. Markdown
rendering composes `RenderArtifacts` and is exposed through the CLI.

Editor integrations should prefer diagnostics from the validation API because
they include source-positioned line and column information.

## Configuration

Native runs can customize resource paths and defaults with
`etc/resources/aws/app.yaml`. All values are optional.

```yaml
paths:
  asset_package:       etc/resources/aws/svg
  service_catalog_csv: etc/resources/aws/service-catalog.csv
  pptx_exporter_wasm:  external/exporter/wasm/xaligo.wasm
  assets_db:           xaligo-assets.db
  project_db:          .xaligo/project.db

item:
  icon_size: 32

serve:
  port: 8080
```

Native configuration remains the default when no explicit asset source is
provided. Embedded and WASM environments provide assets through their adapter
instead of forking the render pipeline.

`serve.port` selects the default HTTP port for `xaligo serve`. Valid ports are
`1` through `65535`, and the default is `8080`; an explicit `--address` keeps
its own port, while an explicit `--port` overrides the port portion of
`--address`.

The V1 compatibility default is `32` layout pixels. `item.icon_size` or an
embedded `AssetSource` supplies that render-context default, and a root
`item-size` attribute overrides it. Callers that need stable cross-environment
geometry should set the root attribute explicitly.

`paths.assets_db` selects the writable SQLite and FTS5 SVG registry. Relative
paths resolve from the project or runtime root. The first icon operation adds
the domain-neutral `builtin` catalog through the same Rust SVG validation and
normalization boundary used for user registrations. AWS and UML catalogs are
separate profile data and are not prerequisites for the builtin catalog.

`paths.project_db` selects the durable project knowledge index shared by RAG,
LSP workspace search, and MCP project tools. The database uses WAL mode and
FTS5. The initial RAG indexing pass reads only Markdown below `docs/`; `.xal`
concept rows are added only through an explicit document-analysis request.
