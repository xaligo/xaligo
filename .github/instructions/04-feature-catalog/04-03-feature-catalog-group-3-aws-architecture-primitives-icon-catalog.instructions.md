---
applyTo: ".github/instructions/manual/**"
---

# 04.03 Feature catalog: Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)

## Group 3 — AWS Architecture Primitives & Icon Catalog (`XAL-3xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-3000010 | `aws-cloud` group tag | Implemented | Top-level AWS Cloud boundary grouping primitive with automatic containment layout. |
| XAL-3000020 | `region` group tag | Implemented | AWS Region boundary grouping primitive. |
| XAL-3000030 | `vpc` group tag | Implemented | VPC boundary grouping primitive. |
| XAL-3000040 | `availability-zone` group tag | Implemented | Availability Zone boundary grouping primitive. |
| XAL-3000050 | `public-subnet` / `private-subnet` group tags | Implemented | Subnet-scoped grouping primitives for AZ-local resources. |
| XAL-3000060 | `security-group` group tag | Implemented | Grouping primitive for resources sharing an EC2 security group. |
| XAL-3000070 | `auto-scaling-group` group tag | Implemented | Grouping primitive for an EC2 Auto Scaling group. |
| XAL-3000080 | `generic-group` tag | Implemented | Non-AWS-specific logical grouping primitive for content that does not match a dedicated AWS group tag. |
| XAL-3000090 | AWS Architecture-Service-Icons catalog | Implemented | Bundled AWS service icon set looked up by numeric catalog ID from `service-catalog.csv`. |
| XAL-3000100 | Tabler icon catalog | Implemented | Imported and cataloged Tabler icon set, refreshed via `npm run import:tabler-icons`. |
| XAL-3000110 | Yamaha icon catalog | Implemented | Imported and cataloged Yamaha icon set, refreshed via `npm run import:yamaha-icons`. |
| XAL-3000120 | Isoflow icon manifest | Excluded unless justified | Retired with Isoflow output; manifest, generator, and generated icon set removed. |
| XAL-3000130 | Service ID lookup catalogs | Implemented | `service-index.csv` for quick ID lookup and `service-catalog.csv` for the full catalog with embedded icon data. |
| XAL-3000140 | `services.csv` label overrides | Implemented | `id,OfficialName,Abbreviation,Summary,Usage,Notes` per-diagram service list driving icon abbreviation overrides. |
| XAL-3000150 | Service legend generation | Implemented | SVG and PPTX output render a service legend derived from `services.csv`, with configurable SVG legend position. |
| XAL-3000160 | Icon license and attribution bundling | Implemented | Bundled license and attribution files for AWS, Tabler, and Yamaha icon sets are preserved alongside the generated catalogs. |
