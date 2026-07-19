# UML

V1 supports a strict xaligo UML profile for the fourteen UML diagram families.
Use this page as the authoring overview; keep detailed tag tables, validation
rules, and complete sample links in the [UML Reference](../reference/uml.md).

## Minimal shape

`<uml>` is a semantic component inside a frame. It contains exactly one
diagram-kind child such as `<class-diagram>` or `<sequence-diagram>`.

```xml
<frame id="sequence" title="Checkout Sequence" version="2026.07" margin-top="48">
  <metadata align="right" width="156" key-width="56" font-size="12" />
  <uml id="checkout-sequence">
    <sequence-diagram>
      <participant id="customer" title="Customer" />
      <lifeline id="api" title="Ordering API" />
      <message src="customer" dst="api" order="1" title="checkout(cart)" />
      <return-message src="api" dst="customer" order="2" title="ok" />
    </sequence-diagram>
  </uml>
</frame>
```

The `<uml>` container does not render its own border or title. Put visible
titles, versions, owners, and review state in the owning frame metadata.

## Authoring rules

- Element IDs are frame-local connection references; UML relations use those
  local IDs directly.
- Exactly one diagram family is allowed inside each `<uml>` component.
- Every diagram family has a closed element/relation vocabulary.
- Relation endpoints must reference existing local UML element IDs.
- `direction="right|down"` selects layout direction; sequence and timing
  diagrams default to horizontal.
- Sequence/communication messages require unique positive dot-separated
  `order` values such as `1`, `2`, and `1.1`.
- Non-participant sequence sources must already be active before sending a
  message or return.

## Where to go next

- [UML Reference](../reference/uml.md): complete diagram-family matrix,
  relation semantics, sequence activation rules, compartments, reusable models,
  and sample source links.
- [UML Examples](../examples/uml.md): compact visual gallery and sample index.
- [Structured diagram design](../design/structured-diagrams.md): product and
  implementation direction for tables, databases, and UML.