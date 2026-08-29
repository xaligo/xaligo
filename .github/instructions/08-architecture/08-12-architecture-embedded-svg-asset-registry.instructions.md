---
applyTo: ".github/instructions/manual/**"
---

# 08.12 Architecture: Embedded SVG asset registry

## Decision

Xaligo uses an embedded SQLite registry with FTS5 for SVG assets. It does not
run a DAM product, PostgreSQL service, icon HTTP server, or another asset
process. Go owns database access and transaction policy. Rust owns SVG parsing,
safety validation, normalization, and rendering through the shared engine ABI.

The ordinary distribution consists of:

```text
xaligo
xaligo-assets.db
```

A future packaging mode may embed a seed database with `go:embed` and copy it
to a user data directory on first use. The writable database remains a normal
SQLite file; do not mutate bytes in the executable.

## Storage ownership

The initial schema is versioned by migrations and contains no renderer-specific
cache as authoritative data:

```sql
CREATE TABLE icons (
    id          INTEGER PRIMARY KEY,
    namespace   TEXT NOT NULL,
    name        TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    svg         BLOB NOT NULL,
    view_box    TEXT NOT NULL,
    width       REAL,
    height      REAL,
    checksum    BLOB NOT NULL,
    compression INTEGER NOT NULL DEFAULT 0,
    license     TEXT NOT NULL DEFAULT '',
    source      TEXT NOT NULL DEFAULT '',
    created_at  INTEGER NOT NULL,
    updated_at  INTEGER NOT NULL,
    UNIQUE(namespace, name)
);

CREATE TABLE icon_tags (
    icon_id INTEGER NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
    tag     TEXT NOT NULL,
    PRIMARY KEY (icon_id, tag)
);

CREATE TABLE icon_aliases (
    namespace TEXT NOT NULL,
    alias     TEXT NOT NULL,
    icon_id   INTEGER NOT NULL REFERENCES icons(id) ON DELETE CASCADE,
    PRIMARY KEY (namespace, alias)
);

CREATE VIRTUAL TABLE icon_search USING fts5(
    namespace,
    name,
    description,
    tags,
    aliases,
    tokenize = 'unicode61'
);
```

Migrations maintain the FTS row for every put, delete, tag, alias, and import
transaction. Search the name, namespace, description, tags, and aliases; never
index raw SVG markup. Store UTF-8 SVG bytes uncompressed initially. Compression
is a versioned storage option, not an implicit behavior selected by readers.

## API boundary

The repository component owns SQL and returns independent entities. The use
case owns validation sequencing, SVG-engine calls, checksums, and operation
coordination. Controllers own files and CLI presentation.

```go
type IconRegistryRepository interface {
	Get(context.Context, string) (Icon, error)
	Search(context.Context, string, int) ([]IconSummary, error)
	Put(context.Context, Icon) error
	Delete(context.Context, string) error
	ListNamespaces(context.Context) ([]string, error)
}
```

Stable icon identity is `<namespace>:<name>`. Namespace and name are normalized
and validated before SQL. Duplicate puts are explicit updates preserving the
stable identity. Checksums are over normalized uncompressed SVG bytes.

The target CLI is:

```text
xaligo icon add <svg> --name <namespace:name>
xaligo icon remove <namespace:name>
xaligo icon get <namespace:name>
xaligo icon search <query>
xaligo icon import <manifest-or-collection>
xaligo icon list [--namespace <name>]
xaligo icon optimize
```


## SQLite operation rules

- enable foreign keys and WAL mode when the database is writable;
- set a bounded `busy_timeout` on every connection;
- keep write transactions short and serialize bulk writes per process;
- use parameterized SQL exclusively;
- cap result limits and SVG byte sizes before allocation;
- validate schema version before ordinary operations;
- perform imports transactionally with deterministic source ordering; and
- preserve license and source attribution through every export or migration.

Read-only bundled databases may use immutable/read-only connection options and
must not attempt WAL writes. A user registry overlays bundled data by stable
identity; it does not modify the bundled seed.

## External catalogs

theSVG, selected Iconify collections, cloud-provider archives, and other icon
sets are optional build-time or explicit import sources. They are not runtime
services or library dependencies of ordinary search/render operations.

```text
external catalog or manifest
            |
       deterministic importer
            |
   SVG validation and normalization
            |
      xaligo-assets.db
```

Importers verify checksums, reject unsafe or malformed SVG, preserve licensing,
and produce the canonical registry schema. Do not download catalog content
during render, validate, LSP, or preview requests.

## Verification gates

- migrations are idempotent and reject unknown future schema versions;
- put/get/delete, aliases, tags, namespaces, and FTS search are covered with a
  temporary database;
- SQL metacharacters remain data and cannot alter queries;
- duplicate identity updates do not leave stale FTS rows;
- concurrent readers and serialized writers respect the busy timeout;
- malformed, oversized, script-bearing, or externally referencing SVG fails
  before persistence;
- checksum and search ordering are deterministic; and
- deleting an icon cascades tags, aliases, and search records.
