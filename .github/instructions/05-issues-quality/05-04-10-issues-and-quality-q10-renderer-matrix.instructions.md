---
applyTo: ".github/instructions/manual/**"
---

# 05.04.10 Issues and quality: Q10 Renderer Matrix

### Q10 Renderer Matrix

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q10.1 | not-started | Define the minimal representative sample set for SVG, PPTX, and Markdown SVG embedding. | matrix documented in this file or docs |
| Q10.2 | not-started | Verify shared scene/plan changes are visible consistently across applicable encoders. | focused format matrix command set |
| Q10.3 | not-started | Confirm Markdown output references the same SVG artifacts and metadata behavior as direct SVG rendering. | Markdown and SVG structural assertions |
| Q10.4 | not-started | Confirm PPTX plan parity and external TypeScript exporter behavior. | Go/WASM plan tests and `npm --prefix external/pptx-exporter test` |
| Q10.5 | not-started | Confirm retired renderer dependencies and APIs stay absent from native and browser-WASM builds. | dependency and public-surface guard |
