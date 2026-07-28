---
applyTo: ".github/instructions/manual/**"
---

# 05.04.10 Issues and quality: Q10 Renderer Matrix

### Q10 Renderer Matrix

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q10.1 | not-started | Define the minimal representative sample set for SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow. | matrix documented in this file or docs |
| Q10.2 | not-started | Verify shared scene/plan changes are visible consistently across applicable encoders. | focused format matrix command set |
| Q10.3 | not-started | Confirm Excalidraw editability and metadata survive renderer changes. | JSON structural assertions |
| Q10.4 | not-started | Confirm PPTX plan parity and external TypeScript exporter behavior. | Go/WASM plan tests and `npm --prefix external test` |
| Q10.5 | not-started | Confirm native-only PDF/Excel dependencies stay out of browser-WASM builds. | `GOOS=js GOARCH=wasm go list -deps ./cmd/wasm` guard |
