# Relational Database

This example defines one reusable schema in `<data>`, projects it through a
`<database>` component, and generates a relation from the foreign key.

![V1 relational database example](../images/databases.svg)

```xml
{{#include samples/databases.xal}}
```

```bash
xaligo validate docs/src/examples/samples/databases.xal
xaligo render docs/src/examples/samples/databases.xal \
  --format svg \
  -o output/databases.svg
```
