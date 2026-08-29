---
applyTo: ".github/instructions/manual/**"
---

# 10.11 PPTX and routing: Item Labels

## Item Labels

- Item icon size defaults to 32px in native CLI config.
- Item label font is 8pt at the default 96 PPI. At another effective PPI it is
  `10.666...px * 72 / effectivePPI` points so its ratio to the icon is stable.
- The layout-pixel font size for item labels is
  `8pt * 96 / 72 = 10.666...px`.
- Item label boxes are 14px high.
- Do not shrink label boxes to text metrics if it breaks PowerPoint placement.
