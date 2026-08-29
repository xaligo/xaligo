---
applyTo: ".github/instructions/manual/**"
---

# 11.02 Diagram creation: Step 2 — Create services.csv

## Step 2 — Create services.csv

`services.csv` lists the services to include in the diagram.

**Format:** `id,OfficialName,Abbreviation,Summary,Usage,Notes`

- Column 1 (`id`) as a number → icon is fetched from service-catalog.csv.
- Lines starting with `#` are treated as comments and ignored.
- For `xaligo render --services`, every non-comment row must have a positive
  numeric `id` and a non-empty `OfficialName`; duplicate IDs are rejected before
  rendering.
- `Abbreviation`, when set, is used as the **icon label inside the diagram** and in the standalone legend icon below the frame.
  - Takes priority over the built-in abbreviation table in
    `internal/entity/service.go`.
  - When empty, the built-in table is used as fallback, then the official name.
- `OfficialName` is displayed as the full-name text in legends.

```csv
# 3-tier Architecture service list — IDs must match <item> tags in the .xal file
# Format: id,OfficialName,Abbreviation,Summary,Usage,Notes
1179,Amazon Route 53,R53,DNS web service,Domain name resolution and health checks,
1581,Amazon VPC Internet Gateway,IGW,Internet connectivity,Inbound/outbound internet traffic,
1182,Elastic Load Balancing,ELB,Load balancing service,Distribute traffic across EC2 instances,
27,Amazon EC2,EC2,Virtual server,Application tier,
1582,Amazon VPC NAT Gateway,NATGW,NAT gateway,Outbound internet for private subnets,
110,Amazon Aurora,Aurora,Relational database,High-performance managed DB,
113,Amazon ElastiCache,EC,In-memory caching,Session and query cache,
```

> **Note:** rendering warns to stderr when an `<item id="N">` in the .xal
> is not listed in services.csv, or when a services.csv entry has no corresponding
> `<item>` in the diagram.  Keep both files in sync to suppress these warnings.

Reference: [docs/src/examples/samples/services.csv](../../docs/src/examples/samples/services.csv)

---
