---
applyTo: ".github/instructions/manual/**"
---

# 01.01 General: Project

## Project

`xaligo` is a Go CLI and WebAssembly/TypeScript package that converts the
`.xal` diagram DSL to SVG and PPTX. It also renders fenced `xal` blocks in
Markdown to SVG and embeds the resulting image references. No other output
format is supported or planned without an explicit product-scope decision.

```text
module: github.com/xaligo/xaligo
Go:     1.26
```

Read `../03-development-flow/03-00-development-flow-overview.instructions.md` for task slicing, verification, and
local commit workflow; `../06-roadmap/06-00-roadmap-overview.instructions.md` for product direction;
`../07-xal-spec/07-01-xal-specification-overview.instructions.md` for DSL behavior; and
`../08-architecture/08-00-architecture-overview.instructions.md` for implementation boundaries. Read
`../09-coding/09-00-coding-overview.instructions.md` before changing Go or TypeScript source.
