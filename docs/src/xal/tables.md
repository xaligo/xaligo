# Tables

V1 supports general data tables inside a frame. A table is different from the
`<grid>` layout primitive and from relational database entities.

## Pipe syntax

Use a GFM-like header, separator, and data rows:

```xml
<table id="services" title="Services">
  | Service | Role     | Port |
  |:--------|:---------|-----:|
  | API     | Backend  | 8080 |
  | DB      | Database | 5432 |
</table>
```

Separator colons select left, center, or right alignment. Every row must have
the same number of cells. Cell text is plain text; escape a literal pipe as
`\|`.

## Tagged syntax

Use tags when rows or cells need stable IDs or individual attributes:

```xml
<table id="services" title="Services">
  <header>
    <cell>Service</cell>
    <cell>Role</cell>
    <cell align="right">Port</cell>
  </header>
  <row id="worker">
    <cell>Worker</cell>
    <cell>Background jobs</cell>
    <cell align="right">9000</cell>
  </row>
</table>
```

`align` accepts `left`, `center`, `right`, or a normal vertical-horizontal
alignment such as `middle-right`. A table may combine one pipe header and its
rows with additional tagged rows. It may not define a second tagged header.

Both forms are normalized to the same V1 table rows and cells before layout.
Malformed separators, inconsistent column counts, duplicate headers, unknown
children, and empty rows are positioned validation errors.

## File imports

CSV, TSV, JSON, and YAML files can be loaded relative to the `.xal` input
file. Declare reusable data under the canonical `<data>` tag and reference it
from one or more frames:

```xml
<data>
  <table-data id="services" src="services.csv" />
</data>
<frames>
  <frame id="inventory">
    <table data="services" title="Service Inventory" />
  </frame>
</frames>
```

The format is inferred from the extension or selected with `format="csv"`,
`tsv`, `json`, `yaml`, or `yml`. CSV and TSV require a header row. JSON and
YAML require a top-level array of objects; their union of keys is sorted to
produce deterministic columns. A one-off import may use `<table
src="services.csv" />` directly.

Import paths must be relative and remain inside the caller-supplied import
filesystem. Inline rows cannot be combined with `data` or `src`. Both
`xaligo validate` and `xaligo render` resolve CLI imports relative to the
input `.xal` file.

See the [table example](../examples/tables.md) and [import example](../examples/table-imports.md)
for complete documents.
