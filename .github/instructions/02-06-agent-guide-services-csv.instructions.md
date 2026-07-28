---
applyTo: ".github/instructions/manual/**"
---

# 02.06 Agent guide: Services CSV

## Services CSV

The accepted columns are:

```text
id,OfficialName,Abbreviation,Summary,Usage,Notes
```

Pass its in-memory bytes through `RenderOptions.ServicesCSV`, or use
`--services` in the CLI. Catalog IDs and abbreviations are shared by all
renderers.
