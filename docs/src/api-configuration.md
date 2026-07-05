# API and Configuration

## Public Boundaries

External consumers should use the CLI, the preview protocol, or the
TypeScript/WASM package. In-repository adapters use `internal/usecase` for the
same parse, layout, render, and validation pipeline.

Available render paths include:

- `Render`
- `RenderExcalidraw`
- `RenderSVG`
- `RenderPPTX`
- `RenderXYFlow`
- `RenderIsoflow`
- `Validate`
- `Diagnose`

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
  pptx_exporter_wasm:  external/wasm/xaligo.wasm

legend:
  offset_x:  120
  offset_y:  0
  icon_size: 32
  font_size: 12

item:
  icon_size: 48
```

Native configuration remains the default when no explicit asset source is
provided. Embedded and WASM environments provide assets through their adapter
instead of forking the render pipeline.
