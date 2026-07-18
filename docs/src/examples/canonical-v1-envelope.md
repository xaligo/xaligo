# Canonical V1 Envelope

This sample uses the canonical V1 document hierarchy:

```text
xaligo
├─ data
└─ frames
   ├─ overview
   └─ database-detail
```

The two identified frames share one document and use a cross-frame connection.
`<data>` is present as the document-wide definition registry; table, database,
and UML definitions will populate it as those V1 processors are delivered.

![Canonical V1 envelope](../images/canonical-v1-envelope.svg)

Source:

```xml
{{#include samples/canonical-v1-envelope.xal}}
```

Validate and render it with:

```bash
xaligo validate docs/src/examples/samples/canonical-v1-envelope.xal
xaligo render docs/src/examples/samples/canonical-v1-envelope.xal \
  --format svg \
  -o output/canonical-v1-envelope.svg
```
