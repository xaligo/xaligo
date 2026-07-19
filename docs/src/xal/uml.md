# UML

V1 supports a strict xaligo UML profile for class, component, activity,
state-machine, and sequence diagrams. Use this page as the authoring overview;
see the
[UML examples](../examples/uml.md) for rendered sample sources.

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
- `direction="right|down"` selects layout direction; sequence diagrams default
  to horizontal.
- Component diagrams may use `component`, `interface`, `port`, and `artifact`
  elements with `dependency`, `realization`, `association`, `assembly`, and
  `delegation` relations.
- Component boxes automatically grow with their interface rows and incoming
  association fan-out. Set `component-width="280"` or
  `component-height="180"` on `<component-diagram>` for diagram defaults, and
  override one box with `<component width="340" height="220">`. Omitting both
  height attributes keeps compact automatic height enabled.
- Set `interface-width="88"` on a `<component>` to use that width for every
  interface-name box inside the component. Interface descriptions continue to
  use the remaining horizontal space.
- Activity diagrams may group elements in `<partition id="..." title="...">`
  swimlanes and set `lanes="vertical|horizontal"` with `theme="xaligo"`.
- Sequence messages require unique positive dot-separated `order` values such
  as `1`, `2`, and `1.1`.
- Non-participant sequence sources must already be active before sending a
  message or return.

## Where to go next

- [UML Examples](../examples/uml.md): compact visual gallery and sample index.
- [Structured diagram design](../design/structured-diagrams.md): product and
  implementation direction for tables, databases, and UML.