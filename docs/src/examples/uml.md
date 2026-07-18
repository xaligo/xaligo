# UML Diagrams

V1 supports all fourteen UML 2.x diagram families through the common `<uml>`
component. Exactly one diagram-kind child selects the diagram processor.

![All UML diagram kinds](../images/uml-all.svg)

```xml
<uml id="domain" title="Domain Model">
  <class-diagram direction="right">
    <class id="user" title="User">
      <attribute>id: bigint</attribute>
      <operation>login()</operation>
    </class>
    <class id="role" title="Role" />
    <association src="user" dst="role" title="has" />
  </class-diagram>
</uml>
```

Reusable definitions can be placed in `<data><uml-model id="...">` and
selected with `data="model-id"` on a diagram-kind child. See the complete
[fourteen-diagram sample](samples/uml-all.xal).
