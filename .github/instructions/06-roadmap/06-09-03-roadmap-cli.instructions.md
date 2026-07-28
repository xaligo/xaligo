---
applyTo: ".github/instructions/manual/**"
---

# 06.09.03 Roadmap: CLI

### CLI

```bash
xaligo render
xaligo validate
```

Required compatibility:

- Keep existing `xaligo render <input.xal> -o <out.excalidraw>` working.
- Add `xaligo render <input.xal> --format excalidraw|svg|pptx|pdf|excel`;
  accept `xlsx` as the Excel alias.
- Keep format conversion under `xaligo render --format ...`; `generate` should
  remain focused on source `.xal` generation.
- `validate` must reuse parser/layout validation rather than duplicate parsing.

---
