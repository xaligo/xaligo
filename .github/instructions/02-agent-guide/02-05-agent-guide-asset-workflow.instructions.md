---
applyTo: ".github/instructions/manual/**"
---

# 02.05 Agent guide: Asset workflow

## Asset workflow

- Quick ID lookup: `etc/resources/aws/service-index.csv`
- Full catalog: `etc/resources/aws/service-catalog.csv`
- Embedded asset declaration: `etc/resources/aws/assets.go`
- AWS/Tabler/Yamaha SVGs: `etc/resources/aws/svg`
- Isoflow icon manifest: `etc/resources/aws/isoflow-icons.json`

Use `npm run import:tabler-icons`, `npm run import:yamaha-icons`, or
`npm run generate:isoflow-icons` to refresh generated catalogs. Preserve the
bundled license and attribution files.
