# Markdown Preview Sample

This self-contained sample demonstrates a Markdown document with an embedded
`xal` diagram. It does not require a service catalog or external assets.

Preview it on A4 landscape paper:

```bash
xaligo serve docs/src/examples/samples/markdown-preview.md \
  --paper A4 \
  --orientation landscape
```

Generate a Markdown file whose `xal` block is replaced by an SVG image:

```bash
xaligo render markdown docs/src/examples/samples/markdown-preview.md \
  --output output/markdown-preview.embedded.md \
  --svg-dir output/markdown-preview
```

## Application Architecture

The browser calls the application API, which stores data in the database.

```xal
<xaligo version="1">
  <data></data>

  <frames>
    <frame id="application" width="960" height="540" class="pa-4">
      <container layout="horizontal" align="middle-center" gap="32">
        <rectangle id="browser" title="Browser" height="200" />
        <rectangle id="api" title="Application API" height="200" />
        <rectangle id="database" title="Database" height="200" />
      </container>

      <connections>
        <connection src="browser" dst="api"
                    src-side="right" dst-side="left" />
        <connection src="api" dst="database" kind="traffic"
                    src-side="right" dst-side="left" />
      </connections>
    </frame>
  </frames>
</xaligo>
```

Edit the Markdown or the `xal` block while `xaligo serve` is running. The
browser preview reloads automatically.
