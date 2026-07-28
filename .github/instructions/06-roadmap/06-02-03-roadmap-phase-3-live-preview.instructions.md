---
applyTo: ".github/instructions/manual/**"
---

# 06.02.03 Roadmap: Phase 3: Live Preview

### Phase 3: Live Preview

Status: initial implementation complete. `xaligo serve` polls `.xal` sources,
renders through the public SVG API, reports source-positioned diagnostics, and
publishes SSE reload events. Browser polish remains.

- Add `xaligo serve` on top of public render/validate APIs. (implemented)
- Watch `.xal` files and automatically re-render. (implemented)
- Serve an SVG-first browser preview with incremental refresh. (implemented)
- Keep the protocol reusable by the VS Code extension.
