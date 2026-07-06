# Live Preview

Use `serve` for a browser-based SVG preview with automatic reload:

```bash
xaligo serve docs/src/examples/samples/junctions.xal --mode network
```

Open:

```text
http://127.0.0.1:8080
```

Endpoints:

| Endpoint | Purpose |
|---|---|
| `/` | Browser preview page |
| `/diagram.svg` | Current SVG or HTTP 422 with diagnostics |
| `/api/status` | JSON version, render error, and diagnostics |
| `/events` | Server-Sent Events for editor integrations |
| `/healthz` | Health check |

Parse and layout errors are shown in the preview without stopping the watcher.
