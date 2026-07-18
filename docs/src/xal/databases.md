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

Each entity renders its column name, type, and constraints. Initial V1
constraint attributes are `primary-key`, `nullable`, `unique`, and `default`.
A single-column `<foreign-key>` uses `columns="local_column"` and
`references="entity.column"`; it generates an entity-to-entity relation.

Validation rejects duplicate schema/entity/column IDs, missing references,
unknown children, inline entities combined with `data`, and malformed foreign
keys. Composite keys, indexes, checks, referential actions, SQL imports, and
Crow's Foot endpoint markers remain subsequent V1 work.

See the [RDB example](../examples/databases.md).
