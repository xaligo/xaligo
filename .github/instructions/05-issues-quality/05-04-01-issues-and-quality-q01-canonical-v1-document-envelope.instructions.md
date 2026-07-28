---
applyTo: ".github/instructions/manual/**"
---

# 05.04.01 Issues and quality: Q01 Canonical V1 Document Envelope

### Q01 Canonical V1 Document Envelope

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q01.2 | not-started | Confirm legacy `<frame>` and `<frames>` inputs remain compatible and emit the intended migration warning. | diagnostics tests covering warning text and source position |
| Q01.4 | in-progress | Confirm canonical samples render as separate artifacts and combined compatibility output. | canonical sample renders `canonical-v1-envelope-overview.svg` and `canonical-v1-envelope-database-detail.svg`; combined compatibility still pending |
| Q01.5 | not-started | Review canonical-envelope docs for current behavior, command accuracy, and image freshness. | `mdbook build docs` and regenerated SVG comparison |
