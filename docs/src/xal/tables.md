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

See the [table example](../examples/tables.md) for a complete document.
