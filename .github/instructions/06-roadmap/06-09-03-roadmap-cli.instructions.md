---
applyTo: ".github/instructions/manual/**"
---

# 06.09.03 Roadmap: CLI

### CLI

```bash
xaligo render
xaligo validate
```

Required contract:

- `xaligo render <input.xal> --format svg|pptx`; omitted format defaults to
  `svg`.
- `xaligo render markdown <input.md>` renders fenced `xal` blocks to SVG and
  embeds image references.
- Reject retired format names (`excalidraw`, `pdf`, `excel`, `xlsx`, `xyflow`,
  and `isoflow`) as unknown formats.
- Keep format conversion under `xaligo render --format ...`; `generate` should
  remain focused on source `.xal` generation.
- `validate` must reuse parser/layout validation rather than duplicate parsing.

---
