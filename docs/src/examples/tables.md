# Tables

This example combines concise pipe rows with an explicit tagged row. Both are
normalized into the same V1 table model and rendered through the shared output
pipeline.

![V1 table example](../images/tables.svg)

Source: [samples/tables.xal](samples/tables.xal).
See [.xal Tables](../xal/tables.md) for the detailed table model.

```bash
xaligo validate docs/src/examples/samples/tables.xal
xaligo render docs/src/examples/samples/tables.xal \
  --format svg \
  -o output/tables.svg
```
