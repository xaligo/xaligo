# API and Configuration

## Public Boundaries

External consumers should use the CLI, the preview protocol, or the
TypeScript/WASM package. In-repository adapters use `internal/usecase` for the
same parse, layout, render, and validation pipeline.

Available render paths include:

- `Render`
- `RenderExcalidraw`
- `RenderSVG`
- `RenderArtifacts`
- `RenderPPTX`
- `RenderPDF`
- `RenderExcel`
- `RenderXYFlow`
- `RenderIsoflow`
- `Validate`
- `Diagnose`

`RenderArtifacts` is the SVG multi-artifact boundary: it returns one ordered
artifact per identified frame by default. `RenderSVG` remains convenient for a
single frame or for `RenderOptions.CombineFrames`; it reports an error instead
of silently discarding extra frame artifacts. PPTX, PDF, and Excel are container
formats, so their render methods return one byte sequence containing one slide,
page, or worksheet per frame. Excalidraw, XYFlow, and Isoflow always return one
logical document.

The CLI accepts `xlsx` as an alias, but API callers set the canonical
`FormatExcel` (`excel`) value.

Editor integrations should prefer diagnostics from the validation API because
they include source-positioned line and column information.

## Configuration

Native runs can customize resource paths and defaults with
`etc/resources/aws/app.yaml`. All values are optional.

```yaml
paths:
  asset_package:       etc/resources/aws/svg
  service_catalog_csv: etc/resources/aws/service-catalog.csv
  output_frames:       output/aws-frames
  pptx_exporter_wasm:  external/pptx-exporter/wasm/xaligo.wasm

legend:
  offset_x:  120
  offset_y:  0
  icon_size: 32
  font_size: 12

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
