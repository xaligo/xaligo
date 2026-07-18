# UML

V1 supports a strict xaligo profile for all fourteen UML 2.x diagram families.
Each family has a closed element and relation vocabulary, and the parser checks
owners, endpoints, control-node degrees, message order, communication links,
and timing intervals before the shared layout runs.

The result is a portable UML projection: SVG, Excalidraw, PPTX, XYFlow, and
Isoflow consume the same shapes and connectors, while editable output retains
the UML diagram, element, relation, owner, order, multiplicity, guard, and time
metadata that its target schema can represent.

## Component and IDs

`<uml>` must be inside a frame, requires a frame-local unique `id`, and contains
exactly one diagram-kind child.

```xml
<uml id="domain" title="Domain model">
  <class-diagram direction="right">
    <interface id="repository" title="OrderRepository">
      <operation>save(order: Order)</operation>
    </interface>
    <class id="sql-repository" title="SqlOrderRepository">
      <operation>save(order: Order)</operation>
    </class>
    <realization src="sql-repository" dst="repository" />
  </class-diagram>
</uml>
```

Element IDs are local to one UML component. UML-native relations use those
local IDs directly. A normal frame-level `<connection>` uses the generated
public reference instead:

| Connection location | Endpoint syntax |
|---|---|
| Same frame | `uml-id/local-id` |
| Another frame | `frame-id.uml-id/local-id` |

For example, use `domain/sql-repository` in the same frame and
`detail.domain/sql-repository` from another frame. UML IDs and local IDs cannot
contain whitespace, `.` or `/`; the separators are reserved. The opaque
hex-scoped internal ID must not be written in source.

`direction` accepts only `right` or `down`. `right` selects horizontal xaligo
layout. Sequence and timing diagrams also default to horizontal; the remaining
families default to vertical. This selects the shared layout direction, not UML
event ordering.

Elements accept normal xaligo presentation attributes such as `width`,
`height`, `color`, `background-color`, `border-color`, `font-family`, and
`font-size`. Their UML defaults are Helvetica at 14 px. Relation color, width,
endpoint side/anchor, bends, coordinate scale, and grid use the normal
connection rules. Element `name` is display text only; it does not become a
frame-level connection alias. Use `uml-id/local-id` instead.

The generic `<element>` and `<relation>` spellings are not valid in the strict
profile. Use the typed tags in the following table.

## Diagram families

| Diagram kind | Elements | Relations | Required and constrained semantics |
|---|---|---|---|
| `class-diagram` | `class`, `interface`, `enumeration` | `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` | At least one classifier. Aggregation/composition are class to class; generalization joins equal classifier kinds; realization is class to interface. |
| `object-diagram` | `object` | `link`, `dependency` | At least one object; every relation joins objects. |
| `component-diagram` | `component`, `interface`, `port`, `artifact` | `dependency`, `realization`, `association`, `assembly`, `delegation` | At least one component. Realization is component to interface. Assembly uses port/interface endpoints and includes a port. Delegation starts at a port. |
| `deployment-diagram` | `node`, `artifact`, `component` | `deployment`, `communication-path`, `dependency` | At least one node. Deployment is artifact/component to node; communication-path is node to node. |
| `package-diagram` | `package`, `class`, `interface`, `component` | `dependency`, `package-import`, `package-merge` | At least one package. Import and merge join packages. |
| `composite-structure-diagram` | `structure`, `collaboration`, `part`, `port`, `component` | `connector`, `assembly`, `delegation`, `dependency` | At least one part/port. Connector joins parts/ports; assembly joins ports; delegation starts at a port. |
| `profile-diagram` | `profile`, `stereotype`, `metaclass` | `extension`, `reference`, `generalization` | At least one profile and stereotype. Extension is stereotype to metaclass; generalization joins stereotypes. |
| `use-case-diagram` | `actor`, `use-case`, `system-boundary` | `association`, `include`, `extend`, `generalization` | At least one use-case. Association joins actor and use-case; include/extend join use-cases; generalization joins equal actor/use-case kinds. |
| `activity-diagram` | `initial`, `final`, `activity`, `action`, `decision`, `merge`, `fork`, `join`, `object-node` | `control-flow`, `object-flow` | At least one activity/action. Control-flow excludes object-node; object-flow includes one. Control-node degrees are validated. |
| `state-machine-diagram` | `initial`, `final`, `state`, `history`, `choice`, `fork`, `join` | `transition` | At least one state. Initial/final direction and pseudostate degrees are validated. |
| `sequence-diagram` | `participant`, `lifeline` | `message`, `return-message`, `create-message`, `destroy-message` | At least one participant/lifeline. Every message has a unique `order`; create/destroy cannot target itself. |
| `communication-diagram` | `object`, `participant` | `link`, `message` | At least two participants, one link, and one message. Each ordered message needs a link between the same unordered pair. |
| `interaction-overview-diagram` | `initial`, `final`, `interaction`, `decision`, `fork`, `join` | `control-flow` | At least one interaction. Initial/final direction and control-node degrees are validated. |
| `timing-diagram` | `lifeline`, `time-state` | `transition`, `occurrence`, `duration` | At least one lifeline/time-state. Owned intervals cannot overlap; transitions are chronological within one lifeline. |

## Ownership

`owner` is a forward-reference-capable local element ID. Its accepted uses are
closed:

| Element | `owner` | Allowed kinds |
|---|---|---|
| Component-diagram `port` | required | `component` |
| Composite `part` | required | `structure`, `component`, `collaboration` |
| Composite `port` | required | `structure`, `part`, `component`, `collaboration` |
| Use-case `use-case` | optional | `system-boundary` |
| Timing `time-state` | required | `lifeline` |

Every other `owner` is an error. V1 retains ownership as metadata, but the
common layout does not spatially nest an owned shape inside its owner.

## Compartments

Direct element children become ordered visible compartments. Every compartment
needs text, `title`, or `name` and cannot contain child elements.

| Element | Typed compartments |
|---|---|
| `class` | `attribute`, `operation`, `constraint`, `note` |
| `interface` | `operation`, `constraint`, `note` |
| `enumeration` | `literal`, `operation`, `note` |
| `object` | `slot`, `note` |
| `component` | `provided-interface`, `required-interface`, `responsibility`, `note` |
| `node`, `artifact` | `property`, `responsibility`, `note` |
| `package` | `responsibility`, `note` |
| `structure` | `property`, `provided-interface`, `required-interface`, `note` |
| `part` | `property`, `responsibility`, `note` |
| `profile` | `constraint`, `note` |
| `stereotype` | `property`, `constraint`, `note` |
| `metaclass` | `property`, `note` |
| `actor` | `responsibility`, `note` |
| `use-case`, `activity`, `action` | `responsibility`, `constraint`, `note` |
| `state` | `entry`, `do`, `exit`, `region`, `note` |
| `interaction` | `note` |
| `time-state` | `region`, `constraint`, `note` |

Elements absent from this table reject compartments. Generic `<compartment>`
is accepted only where a typed compartment is allowed, for compatibility; new
source should use a typed tag.

## Relations and control flow

All relations require existing local `src` and `dst`. `title` or `label`
supplies visible text. `guard` is valid only on control-flow, object-flow, and
transition. Multiplicity attributes are valid only on association,
aggregation, composition, and link:

```xml
<composition src="order" dst="line"
             src-multiplicity="1" dst-multiplicity="1..*" />
<transition src="draft" dst="confirmed" title="confirm" guard="valid" />
```

The UML relation kind fixes its stroke and arrowhead. `kind`, `stroke-style`,
`arrowhead`, `start-arrowhead`, and `end-arrowhead` cannot override it.
Include/extend/import/merge/dependency-like relations are dashed; directed
flows use destination triangles; aggregation/composition use source diamonds;
structural links have no destination marker. When their label is omitted,
include, extend, import, merge, deployment, extension, and occurrence receive a
semantic default label.

Activity, state-machine, and interaction-overview graphs enforce these flat V1
rules when a control node is present:

- `initial` has no incoming edge and at least one outgoing edge;
- `final` has no outgoing edge and at least one incoming edge;
- `decision`, `choice`, and `fork` have at least one incoming and two outgoing;
- `merge` and `join` have at least two incoming and one outgoing; and
- `history` has at least one outgoing transition.

## Messages and timing

Sequence and communication messages require a positive dot-separated `order`
without leading zeroes, such as `1`, `2`, or `1.1`. The full string is unique
within the diagram. Numeric order is prepended to the visible label and assigns
top-to-bottom connector anchors on participant/lifeline shapes; it does not
reorder declared elements or add activation boxes. Sequence messages always
anchor on a vertical edge: explicit `top` is normalized to `left`, `bottom` to
`right`, and `order` supersedes an explicit anchor slot.

A communication message must have a `link` for the same endpoint pair. Link
direction is ignored for this check.

Timing numbers are finite, unitless base-10 values in one caller-chosen unit:

| Construct | Attributes | Validation |
|---|---|---|
| `time-state` | `owner`, `from`, `to` | Lifeline owner; `0 <= from < to`; intervals for one owner cannot overlap |
| `transition` | `src`, `dst` | Time-states of one owner; source interval finishes before destination starts |
| `occurrence` | `src`, `dst`, `at` | Lifeline/time-state endpoints; `at >= 0` and inside any referenced state interval |
| `duration` | `src`, `dst`; optional `from`, `to` pair | Time-state endpoints; optional range satisfies `0 <= from < to` |

Use labels for units, for example `title="80 ms"`; numeric attributes such as
`at="20ms"` are invalid.

## Reusable models

Reusable definitions live under the document's `<data>` start/end tags:

```xml
<data>
  <uml-model id="order-objects">
    <object id="customer" title="customer: Customer"><slot>name = Alice</slot></object>
    <object id="order" title="order42: Order"><slot>status = Confirmed</slot></object>
    <link src="customer" dst="order" title="placed" />
  </uml-model>
</data>
<frames>
  <frame id="snapshot">
    <uml id="runtime"><object-diagram data="order-objects" direction="right" /></uml>
  </frame>
</frames>
```

Inline children and `data` cannot be combined. Expansion happens before the
selected family validates the model, so the model must satisfy that family's
closed vocabulary and all owner/endpoint/order/time rules.

## Shared-output projection

V1 intentionally uses the capability shared by every xaligo output:

- use-cases and initial/final nodes are ellipses;
- decision, merge, choice, and history nodes are diamonds;
- other elements are editable rectangles with flattened visible compartments;
- relations are routed orthogonal connectors with separate labels;
- aggregation and composition currently share a diamond projection;
- sequence order is metadata, text, and top-to-bottom message anchoring, not a
  separate lifeline/activation event axis;
- timing values are validated metadata, not proportional waveforms; and
- semantic ownership does not imply drawn containment.

XYFlow retains UML node/relation metadata in node and edge `data`, including
the projected node shape. Isoflow represents connected UML shapes with labeled
generic endpoint icons and connectors; its upstream schema cannot carry the
arbitrary UML metadata fields.

This is not XMI or a lossless UML interchange format. See the
[fourteen detailed examples](../examples/uml.md) for complete authoring
templates.
