---
applyTo: ".github/instructions/manual/**"
---

# 06.03 Roadmap: VS Code Extension Preconditions

## VS Code Extension Preconditions

The VS Code extension is developed in a separate repository. This repository
owns the reusable Go/WASM APIs and HTTP/SSE preview protocol only; do not add
extension packaging or VS Code-specific parser/rendering forks here.

The extension target includes:

- `.xal` syntax highlighting.
- Validation and source-positioned diagnostics. (Go and TypeScript/WASM APIs implemented)
- Live Preview and a Preview Panel.
- SVG preview for `.xal` and SVG-embedded Markdown preview.

The extension must call the same validation/render pipeline as the CLI. Do not
create an extension-only parser, layout engine, or hidden preview format.
