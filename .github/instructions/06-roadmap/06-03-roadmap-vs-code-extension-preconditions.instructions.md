---
applyTo: ".github/instructions/manual/**"
---

# 06.03 Roadmap: VS Code Extension Preconditions

## VS Code Extension Preconditions

The VS Code extension is developed in a separate repository. This repository
owns the reusable native CLI/service boundaries and HTTP/SSE preview protocol;
do not add extension packaging or VS Code-specific parser/rendering forks here.
The retained `cmd/wasm` V1 adapter is compatibility-only and is not a V2
integration boundary.

The extension target includes:

- `.xal` syntax highlighting.
- Validation and source-positioned diagnostics through the shared native API.
- Live Preview and a Preview Panel.
- SVG preview for `.xal` and SVG-embedded Markdown preview.

The extension must call the same validation/render pipeline as the CLI. Do not
create an extension-only parser, layout engine, or hidden preview format.
