# Services and Legends

`--services <csv>` provides label overrides and legend metadata.

The CSV format is:

```text
id,OfficialName,Abbreviation,Summary,Usage,Notes
```

Rules:

- Blank lines and lines beginning with `#` are ignored.
- `id` must be a positive numeric service catalog ID.
- `OfficialName` is required.
- `Abbreviation` is optional. When present, it is used for item labels and
  legend abbreviations.
- Duplicate IDs are rejected.

Example:

```csv
id,OfficialName,Abbreviation,Summary,Usage,Notes
1178,Amazon CloudFront,CF,CDN,Edge delivery,
27,Amazon EC2,EC2,Compute,Application servers,
117,Amazon RDS,RDS,Database,Relational database,
```

SVG output draws a service legend when `--services` is supplied. PPTX output
adds separate legend slide(s) after all frame/diagram slides.
