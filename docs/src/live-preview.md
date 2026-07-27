# Live Preview

Use `serve` for a browser-based preview with automatic reload:

```bash
xaligo serve docs/src/examples/samples/junctions.xal --mode network
```

Open:

```text
http://127.0.0.1:8080
```

The default port is `8080`. Set `serve.port` in
`etc/resources/aws/app.yaml`, or override it for one process with `--port`:

```bash
xaligo serve docs/src/examples/samples/junctions.xal --port 9090
```

The existing `--address` flag remains available for choosing both the listen
host and port. When both flags are present, `--port` replaces only the port
part of `--address`.

A `.md`/`.markdown` source previews the full Markdown document with rendered
`xal` code blocks embedded inline, the same as `render markdown`:

```bash
xaligo serve docs/src/examples/samples/markdown-preview.md
```

Markdown raw HTML is not executed. Rendered diagrams are isolated as SVG image
documents, and relative Markdown images are served only when they resolve to an
image file inside the Markdown file's directory. Parent-directory traversal,
escaping symlinks, and remote image loading are blocked by the preview sandbox
and Content Security Policy.

Use `--paper`/`--orientation` to preview how a diagram fits a specific
physical page size and orientation; changing the paper size or orientation
requires restarting the server:

```bash
xaligo serve docs/src/examples/samples/markdown-preview.md \
  --paper A4 \
  --orientation landscape
```

Endpoints:

| Endpoint | Purpose |
|---|---|
| `/` | Browser preview page |
| `/diagram.svg` | Current SVG, or HTTP 422 with diagnostics (`.xal` source) |
| `/content.html` | Current rendered HTML document, or HTTP 422 with diagnostics (`.md`/`.markdown` source) |
| `/api/status` | JSON version, render error, and diagnostics |
| `/events` | Server-Sent Events for editor integrations |
| `/healthz` | Health check |

Parse and layout errors are shown in the preview without stopping the watcher.
