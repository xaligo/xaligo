# Live Preview

Use `serve` for a browser-based preview with automatic reload:

```bash
xaligo serve docs/src/examples/samples/junctions.xal --mode network
```

Open:

```text
http://127.0.0.1:8080
```

A `.md`/`.markdown` source previews the full Markdown document with rendered
`xal` code blocks embedded inline, the same as `render markdown`:

```bash
xaligo serve docs/guide.md
```

Use `--paper`/`--orientation` to preview how a diagram fits a specific
physical page size and orientation; changing the paper size or orientation
requires restarting the server:

```bash
xaligo serve diagram.xal --paper A4 --orientation landscape
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
