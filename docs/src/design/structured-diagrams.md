# Structured Diagrams: Tables, Databases, and UML

> Status: V1 implementation design. Canonical V1 uses
> `<xaligo version="1">`. Historical root `<frame>` and `<frames>` documents
> remain accepted and produce a migration warning. V2 remains reserved for
> `<scene version="2">`.

This design turns the table, relational-database, and UML discussions into one
coherent target while keeping their semantic processing independent. It
defines the document envelope, shared data boundary, component syntax,
normalization rules, validation ownership, and phased delivery. Rendering and
layout details may evolve without changing these boundaries.

## Design goals

- Keep one `.xal` source document capable of containing architecture, table,
  database, and UML views.
- Reuse data definitions across frames without coupling a definition to one
  visual notation.
- Offer concise pipe syntax and lossless explicit tag syntax where both are
  appropriate.
- Import common data and schema files through deterministic, safe adapters.
- Preserve the existing concepts of frames, item/grid placement, explicit
  geometry, ports, and routed lines without making every diagram use one
  semantic processor.
- Lower every frontend into typed, renderer-neutral models before layout and
  encoding.

## Document envelope

The target hierarchy is:

```text
xaligo
├─ metadata
├─ imports
├─ data                 reusable document-wide definitions
├─ styles
└─ frames
   └─ frame             page/layout boundary
      ├─ existing layout and architecture components
      ├─ table
      ├─ database
      └─ uml
         └─ exactly one UML diagram-kind child
```

Example:

```xml
<xaligo version="1">
  <data>
    <table-data id="service-inventory" src="services.csv" />
    <database-schema id="application-schema" src="schema.sql"
                     format="sql" dialect="postgresql" />
    <uml-model id="domain-model">
      <!-- typed UML definitions -->
    </uml-model>
  </data>

  <frames>
    <frame id="inventory">
      <table data="service-inventory" />
    </frame>
    <frame id="erd">
      <database data="application-schema" notation="crow-foot" />
    </frame>
    <frame id="domain">
      <uml>
        <class-diagram data="domain-model" />
      </uml>
    </frame>
  </frames>
</xaligo>
```

`<data>` owns reusable meaning; `<frame>` owns presentation and page layout.
A data definition may be projected into several views, such as a database
schema rendered as an ER diagram, a general table, or a UML class diagram.
Such projections are explicit mappings, not implicit reinterpretation.

Every frame has a stable, document-unique `id`. A frame may contain several
components arranged with the normal row, column, container, and grid controls.
The component selects its semantic processor; `frame` never carries values
such as `type="uml-class"`.

## Processing boundaries

General tables, relational schemas, and UML are separate frontends:

```text
source document
  -> envelope parser and root/version dispatch
  -> include/import resolution
  -> data registry and reference resolution
  -> component dispatch
       table    -> table parser/validator/layout
       database -> RDB parser/validator/ER layout
       uml      -> UML-kind parser/validator/layout
  -> renderer-neutral draw operations
  -> shared format encoders
```

They may share primitives such as boxes, compartments, text policy, ports,
orthogonal edges, themes, and output encoders. They do not share a catch-all
semantic AST or force UML/RDB behavior through the existing architecture
layout engine. The common boundary is a typed document registry plus neutral
resolved drawing contracts.

The envelope parser selects exactly one versioned frontend from the root and
version pair. It must not retry parsers after an error. `<xaligo version="1">`
selects V1; `<scene version="2">` remains the reject-safe V2 boundary.

## General tables

`<table>` represents data, not a visual placement grid. Simple tables use a
GFM-like pipe form:

```xml
<table id="services" title="Services" grid="all" striped="true">
  | Service | Role     | Port |
  |:--------|:---------|-----:|
  | Nginx   | Proxy    |  443 |
  | API     | Backend  | 8080 |
</table>
```

Alignment comes from the separator row. The initial profile supports plain
text cells and escaped pipes; it does not treat arbitrary Markdown or HTML as
cell content.

Explicit tags provide stable IDs, rich content, spans, and cell styling:

```xml
<table id="services" title="Services" grid="all">
  <header>
    <cell id="service-heading">Service</cell>
    <cell>Role</cell>
    <cell align="right">Port</cell>
  </header>
  <row id="api-row">
    <cell id="api"><item id="27" />API</cell>
    <cell>Backend</cell>
    <cell align="right">8080</cell>
  </row>
</table>
```

Pipe rows and explicit rows may coexist. Source order is preserved. Both forms
must declare the same effective column count; conflicting header definitions,
duplicate row/cell IDs, and invalid spans are errors. Both lower into the same
typed table model before layout.

Presentation properties include `title`, `grid`, `striped`, `compact`, column
widths, horizontal/vertical alignment, overflow policy, six-digit text/fill/
border colors, font family, and font size. Table styles inherit through rows to
cells; explicit row/cell values override them, while `header-*` table
attributes style pipe headers. `rowspan`, `colspan`, rich cell children, and
cell endpoints belong to the explicit form.

`<grid>` remains a layout primitive for arranging diagram components. A grid
must not be inferred to be a data table merely because it looks tabular.

## Relational database design

`<database>` is the displayed schema component; `<database-schema>` is the
reusable data definition. `<entity>` represents an RDB relation and is not an
alias for a general table.

```xml
<data>
  <database-schema id="application-schema">
    <entity id="roles" name="roles">
      <column name="id" type="bigint" primary-key="true"
              nullable="false" />
      <column name="name" type="varchar(100)" unique="true" />
    </entity>
    <entity id="users" name="users">
      <column name="id" type="bigint" primary-key="true"
              nullable="false" />
      <column name="role_id" type="bigint" nullable="false" />
      <foreign-key id="fk-users-role" columns="role_id"
                   references="roles.id" on-delete="restrict" />
    </entity>
  </database-schema>
</data>

<frames>
  <frame id="database">
    <database data="application-schema" notation="crow-foot" />
  </frame>
</frames>
```

An entity may also use a schema-specific pipe form:

```xml
<entity id="users" name="users">
  | Name    | Type         | Key | Nullable | References |
  |:--------|:-------------|:----|:---------|:-----------|
  | id      | bigint       | PK  | false    |            |
  | email   | varchar(255) | UQ  | false    |            |
  | role_id | bigint       | FK  | false    | roles.id   |
</entity>
```

The header names are part of the profile and map to typed column fields. Tags
remain authoritative for composite primary/foreign keys, indexes, checks,
referential actions, comments, and dialect-specific details. Mixed syntax
lowers once into the same RDB model and rejects contradictory declarations.

Foreign keys can generate relations. An explicit relation may add visual
meaning but cannot contradict schema constraints without an explicit
projection mode. Endpoint references use typed paths such as
`users.role_id`; table cell endpoints use a separate syntax so they cannot be
mistaken for RDB columns.

RDB validation includes unique entity/column IDs, referenced entity and column
existence, compatible key arity and types, one effective primary key, valid
index/check definitions, and consistent null/default declarations. Diagnostics
retain the originating file and source range.

## UML

`<uml>` is the common UML component. It contains exactly one diagram-kind
child, which selects the semantic processor:

```text
uml
├─ class-diagram
├─ object-diagram
├─ component-diagram
├─ deployment-diagram
├─ package-diagram
├─ composite-structure-diagram
├─ profile-diagram
├─ use-case-diagram
├─ activity-diagram
├─ state-machine-diagram
├─ sequence-diagram
├─ communication-diagram
├─ interaction-overview-diagram
└─ timing-diagram
```

Example:

```xml
<uml id="domain-view" title="Domain Model" theme="default">
  <class-diagram data="domain-model" layout="elk" direction="right" />
</uml>
```

No child and multiple diagram-kind children are validation errors. Unknown
kinds are rejected instead of being treated as generic groups. Common `uml`
properties include `id`, `title`, `theme`, `style`, `data`, and visibility.
The child may override inherited data or style values.

V1 implements all fourteen diagram-kind selectors through shared typed element,
compartment, relation, interaction, and timeline primitives. The selected kind
is retained as semantic metadata while each element and relation is validated
before lowering to renderer-neutral shapes and connectors.

Existing layout ideas are reused at the component boundary: `<uml>` can sit in
rows, columns, grids, or containers and returns resolved width, height, draw
operations, anchors, and links. Inside `<uml>`, its selected processor owns the
layout. Existing item icons may be embedded where a UML element explicitly
permits them. Existing line concepts inform edge style, anchors, bends,
routing, and line jumps, but UML relationships keep their own semantic kinds.

## Imports and overrides

Imports are declared in `<data>` or the document-level `<imports>` registry:

```xml
<data>
  <table-data id="services" src="services.csv" format="csv" />
  <table-data id="endpoints" src="endpoints.yaml" format="yaml"
              path="services" />
  <database-schema id="schema" src="schema.sql" format="sql"
                   dialect="postgresql" />
</data>
```

V1 implements CSV/TSV, JSON, and YAML table adapters with an injected
filesystem and paths relative to the input document. SQL DDL and DBML remain
planned. OpenAPI may later project schemas into tables or UML components. Spreadsheet import is deferred
until its type, formula, merged-cell, and formatting loss policy is specified.

Resolution order is:

1. parse the envelope and declarations;
2. expand includes while detecting cycles;
3. load imports under the caller's asset/file policy;
4. parse each import into its typed source model;
5. apply explicit inline overrides;
6. resolve references and projections;
7. perform component-specific semantic validation;
8. lay out and render frames.

Overrides address stable semantic IDs. They never depend on imported row
position. The initial profile uses the deterministic precedence `explicit tag
override > pipe declaration > imported value > default`. A conflicting value
at the same precedence is an error.

Import resolution records source URI/path, format, dialect, and source ranges
so diagnostics and round trips can preserve provenance. A bundled/normalized
form may embed resolved data, but normal editing retains external references.

Arbitrary command execution is outside the import contract. Network and
filesystem access are capability-injected, disabled unless supplied by the
caller, and subject to path/scheme policies. Credentials are references to
environment or secret providers and are never embedded in rendered output.

## Typed models and renderer contract

The document model contains metadata, imports, data definitions, styles, and
frames. It does not flatten all meanings into generic `Node` and `Edge` maps.
Table, RDB, and each UML family retain typed semantic models until their own
layout stage.

After layout, processors emit shared renderer-neutral operations for shapes,
text, compartments, icons, ports, and semantic edges. Operations retain stable
source IDs, semantic kind, parentage, text policy, geometry, style, and
optional endpoint metadata. Output encoders project these operations into SVG,
PPTX, Excalidraw, XYFlow, or another supported capability set without becoming
alternate semantic models.

## Delivery sequence

1. Add the `<xaligo version="1">` envelope, legacy-root warning, data registry,
   frames, imports, and typed
   component dispatch without changing current V1/V2 behavior.
2. Deliver general tables with pipe/tag normalization and SVG output.
3. Add CSV, JSON, and YAML table imports. (Implemented in V1.)
4. Deliver RDB entities, keys, relations, Crow's Foot rendering, and SQL DDL
   import for PostgreSQL, MySQL, and SQLite. (Entities, keys, relations, and
   common SQL DDL import are implemented in V1; Crow's Foot remains.)
5. Deliver all fourteen UML diagram families through common typed primitives
   and diagram-kind validation. (Implemented in V1.)
6. Add DBML/OpenAPI projections, richer table cells, and additional encoders.
7. Add explicit normalized/bundled output and GUI round-trip contracts.

Every phase requires parser/validation tests, model golden tests, resolved
geometry tests, cross-format capability tests, source-positioned import
diagnostics, and V1/V2 root rejection regression tests.
