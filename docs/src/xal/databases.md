# Relational Databases

V1 stores reusable relational schemas under the document-wide `<data>` node
and renders them with a `<database>` component.

```xml
<data>
  <database-schema id="application-schema">
    <entity id="roles" name="roles">
      <column name="id" type="bigint" primary-key="true" nullable="false" />
    </entity>
    <entity id="users" name="users">
      <column name="id" type="bigint" primary-key="true" nullable="false" />
      <column name="role_id" type="bigint" nullable="false" />
      <foreign-key columns="role_id" references="roles.id" />
    </entity>
  </database-schema>
</data>
```

Reference the schema from a frame:

```xml
<database data="application-schema" title="Application Schema" />
```

Each entity renders its column name, type, and constraints. V1 constraint
attributes are `primary-key`, `nullable`, `unique`, and `default`. Composite
primary keys use an explicit declaration:

```xml
<primary-key columns="tenant_id,id" />
```

Foreign keys use equally sized comma-separated local and qualified referenced
columns. They generate one entity-to-entity relation regardless of arity:

```xml
<foreign-key columns="tenant_id,role_id"
             references="roles.tenant_id,roles.id"
             on-delete="cascade" />
```

`on-delete` and `on-update` accept `cascade`, `restrict`, `no-action`,
`set-null`, or `set-default`. The normalized relation retains its source and
referenced column lists and referential actions for downstream renderers.

## SQL DDL imports

Import a schema relative to the `.xal` input file:

```xml
<data>
  <database-schema id="application-schema" src="schema.sql"
                   format="sql" dialect="postgresql" />
</data>
```

The V1 common profile accepts `postgresql` (and the `postgres` alias), `mysql`,
and `sqlite`. It reads semicolon-terminated `CREATE TABLE` statements, quoted
or schema-qualified identifiers, column types, `NOT NULL`, `UNIQUE`, inline or
table primary keys, inline or table foreign keys, composite keys,
`REFERENCES`, `ON DELETE`, and `ON UPDATE`. Dialect-specific indexes, generated
expressions, checks, and table options are currently ignored when they do not
affect the normalized keys.

Validation rejects duplicate schema/entity/column IDs, missing references,
unknown children, inline entities combined with `data`, malformed foreign
keys, mismatched composite-key arity, missing columns, and invalid referential
actions. Rich indexes, checks, and Crow's Foot endpoint markers remain
subsequent V1 work.

See the [RDB example](../examples/databases.md).
