# UML Reference

This is the detailed reference for the V1 UML profile. Keep complete syntax,
validation rules, and fine-grained sample links here. Overview pages and example
pages should link here instead of repeating these tables.

## Component and IDs

`<uml>` must be inside a frame, requires a frame-local unique `id`, and contains
exactly one diagram-kind child.

```xml
<frame id="domain-page" title="Domain model" version="2026.07" margin-top="48">
  <metadata align="right" width="156" key-width="56" font-size="12" />
  <uml id="domain">
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
</frame>
```

Element IDs are public frame-level connection references. UML-native relations
use local IDs directly. A frame-level `<connection>` uses the same ID in the
owning frame, or `frame-id.local-id` across frames.

| Connection location | Endpoint syntax |
|---|---|
| Same frame | `local-id` |
| Another frame | `frame-id.local-id` |

UML element IDs must be unique among frame connection endpoint IDs and public
refs. UML component IDs and local IDs cannot contain whitespace, `.` or `/`.
The opaque scoped internal ID is generated and must not be written in source.

`direction` accepts only `right` or `down`. Sequence and timing diagrams default
to horizontal; the remaining families default to vertical. The `<uml>` container
is semantic only and does not render an outer border or title. Use frame
metadata for visible titles, versions, owners, and review state.

The generic `<element>` and `<relation>` spellings are not valid in the strict
profile. Use the typed element and relation tags listed below.

UML diagrams currently do not support user-facing Excalidraw export. Rendering
UML with `--format excalidraw` returns an error; use SVG, PDF, PPTX, Excel,
XYFlow, or Isoflow output instead.

## Diagram Families

| Diagram kind | Elements | Relations | Required and constrained semantics | Sample |
|---|---|---|---|---|
| `class-diagram` | `package`, `class`, `interface`, `enumeration` | `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` | At least one classifier. Packages group classifiers and package-local relations. Aggregation/composition are class to class; generalization joins equal classifier kinds; realization is class to interface. | [uml-class.xal](../examples/samples/uml-class.xal) |
| `object-diagram` | `object` | `link`, `dependency` | At least one object; every relation joins objects. | [uml-object.xal](../examples/samples/uml-object.xal) |
| `component-diagram` | `component`, `interface`, `port`, `artifact` | `dependency`, `realization`, `association`, `assembly`, `delegation` | At least one component. Component boxes render with a cyan header, white body, left-aligned white component name, and no two-rectangle adornment. A component may contain only child `interface` entries; each child interface renders as a small white boundary port box with the interface text inside and outside circle symbols, clears the component header, and sits mostly inside the component with only a small outside protrusion. Matching component associations keep the interface user endpoint on the source component's nearest non-interface-side anchor, selected from 15 top/right/bottom component anchors, add a left-bulging caller-side semicircle as a forked connection endpoint, bind the destination endpoint to a same-named interface circle, and approach the circle horizontally from the outside-left side of left-side interface circles at the circle center height; the line ends at the semicircle's left bend, and the semicircle radius is 2px larger than and center-aligned with the interface circle. When multiple associations target the same component interface, the destination renders one circle per incoming association with enough vertical spacing to prevent caller-side semicircle overlap and groups them back to the interface box with a bracket-style stem. Realization is component to interface. Assembly uses port/interface endpoints and includes a port. Delegation starts at a port. | [uml-component.xal](../examples/samples/uml-component.xal), [uml-component-connected.xal](../examples/samples/uml-component-connected.xal), [uml-component-connected-complex.xal](../examples/samples/uml-component-connected-complex.xal) |
| `deployment-diagram` | `node`, `artifact`, `component` | `deployment`, `communication-path`, `dependency` | At least one node. Deployment is artifact/component to node; communication-path is node to node. | [uml-deployment.xal](../examples/samples/uml-deployment.xal) |
| `package-diagram` | `package`, `class`, `interface`, `component` | `dependency`, `package-import`, `package-merge` | At least one package. Import and merge join packages. | [uml-package.xal](../examples/samples/uml-package.xal) |
| `composite-structure-diagram` | `structure`, `collaboration`, `part`, `port`, `component` | `connector`, `assembly`, `delegation`, `dependency` | At least one part/port. Connector joins parts/ports; assembly joins ports; delegation starts at a port. | [uml-composite-structure.xal](../examples/samples/uml-composite-structure.xal) |
| `profile-diagram` | `profile`, `stereotype`, `metaclass` | `extension`, `reference`, `generalization` | At least one profile and stereotype. Extension is stereotype to metaclass; generalization joins stereotypes. | [uml-profile.xal](../examples/samples/uml-profile.xal) |
| `use-case-diagram` | `actor`, `use-case`, `system-boundary` | `association`, `include`, `extend`, `generalization` | At least one use case. Association joins actor and use-case; include/extend join use-cases; generalization joins equal actor/use-case kinds. | [uml-use-case.xal](../examples/samples/uml-use-case.xal) |
| `activity-diagram` | `initial`, `final`, `activity`, `action`, `decision`, `merge`, `fork`, `join`, `object-node`; optional `partition` containers | `control-flow`, `object-flow` | At least one activity/action. Control-flow excludes object-node; object-flow includes one. Control-node degrees are validated. | [uml-activity.xal](../examples/samples/uml-activity.xal) |
| `state-machine-diagram` | `initial`, `final`, `state`, `history`, `choice`, `fork`, `join`; optional layout-only `container`/`row`/`col` | `transition` | At least one state. Initial/final direction and pseudostate degrees are validated. State entry/do/exit/internal/region compartments and transition event/guard/action labels render visibly. Initial/final/choice/history pseudostates keep compact proportions, and final states render the standard inner dot. `show-element-names="false"` hides compartment element names such as `entry` and `do` while retaining state titles, compartment values, and transition labels; an element can override it with `show-element-names="true"`. A layout-only `container` can group elements into `row` and `col` child tags. | [uml-state-machine.xal](../examples/samples/uml-state-machine.xal) |
| `sequence-diagram` | `participant`, `lifeline` | `message`, `return-message`, `create-message`, `destroy-message` | At least one participant/lifeline. Every message has a unique `order`; non-participant sources must already be active; create/destroy cannot target themselves. | [uml-sequence.xal](../examples/samples/uml-sequence.xal) |
| `communication-diagram` | `object`, `participant` | `link`, `message` | At least two participants, one link, and one message. Each ordered message needs a link between the same unordered pair. | [uml-communication.xal](../examples/samples/uml-communication.xal) |
| `interaction-overview-diagram` | `initial`, `final`, `interaction`, `decision`, `fork`, `join` | `control-flow` | At least one interaction. Initial/final direction and control-node degrees are validated. | [uml-interaction-overview.xal](../examples/samples/uml-interaction-overview.xal) |
| `timing-diagram` | `lifeline`, `time-state` | `transition`, `occurrence`, `duration` | At least one lifeline/time-state. Owned intervals cannot overlap; transitions are chronological within one lifeline. | [uml-timing.xal](../examples/samples/uml-timing.xal) |

## Presentation Defaults

Elements accept normal xaligo presentation attributes such as `width`, `height`,
`color`, `background-color`, `border-color`, `font-family`, and `font-size`.
Their UML defaults are Helvetica at 14 px. Relation color, width, endpoint
side/anchor, bends, coordinate scale, and grid use the normal connection rules.
Element `name` is display text only; use the element `id` for references.

State-machine diagrams additionally accept `show-element-names="false"` on
`<state-machine-diagram>` to suppress state compartment element names such as
`entry`, `do`, `exit`, `internal`, `note`, and `region`. The setting does not
hide state titles, compartment values, transition labels, or shape geometry.
Set `show-element-names="true"` on a specific state-machine element to restore
its compartment element names when the diagram default is hidden.

UML relation endpoints use a UML-specific default anchor profile before they
enter the shared routing pipeline. Rectangle-like UML elements snap to five
inset anchors on each side (`top-1` through `top-5`, and the matching right,
bottom, and left names), giving 20 perimeter anchors. Diamond-like elements
such as `choice`, `decision`, `merge`, and `history` use the four vertices;
explicit diamond anchors choose the vertex for that side. Sequence messages are
the exception: their endpoints remain ordered by message time along the
lifeline.

Class diagrams use compact classifier placement, cyan headers, white bodies,
and deep-blue relation lines. Packages render with the general-group visual
language. Use `grid="N"` on `<class-diagram>` or class-diagram `<package>` to
choose the classifier/package column count.

Sequence diagrams render participants and lifelines as compact top headers with
dashed vertical lifeline axes. Messages are ordered from top to bottom by
numeric `order`. `message`, `create-message`, and `destroy-message` draw
destination activation bars; `return-message` is a dashed response connector and
does not start a new activation. Activation bars extend through related calls,
returns, and cleanup messages until the lifeline returns. Fully contained
activations merge into their covering bar so nested self-calls do not add inner
strokes. `message mode="async"` uses an open arrowhead. `destroy-message` draws
a stop marker and its label must clearly describe destruction, deletion,
disposal, removal, or termination.

State-machine diagrams can use a layout-only `<container>` whose child `<row>`
and `<col>` tags place states on a shared grid without adding UML elements.
When an element is not placed in a container column, the row layout reuses
nearby connected columns where possible before assigning the next free column,
so branch states stay close to the state that leads to them. State compartments
render UML state behavior rows for `entry`, `do`, `exit`, `internal`, `region`, and `note`;
`note` is useful for simple descriptive states that do not need actions. Each
state has a cyan name header with white text, a white body, horizontal row
dividers, and a vertical key/value divider. Transition labels use
`event [guard] / action-or-effect`. By default, state machines use the same
xaligo palette as class and sequence diagrams: deep-blue borders/text/relations,
white bodies, cyan state-name headers, and a deep-blue initial dot.
Presentation attributes such as `background-color`, `border-color`, and `color`
can still override individual shapes.

Class, activity, and state-machine relation tags can contain `<bend x="..."
y="..." />` children. Bends steer the generated orthogonal connector route and
are preserved in editable scene metadata. State-machine transition routing also
treats intermediate state and pseudostate bodies as obstacles where possible, so
default connector paths avoid cutting through nearby states. Same-frame route
points stay inside the frame bounds, and distant state pairs or tight bent routes
may use larger outside detours when an interior path would cross a state body.
UML relation labels are automatically nudged away from endpoint items when a
default label position would overlap a shape.

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

Every other `owner` is an error. Component and composite owned ports are placed
on their owner's boundary using the normal `side`, `x`, and `y` port placement
attributes. Other ownership relationships are retained as metadata and are not
spatially nested in V1.

## Compartments

Direct element children become ordered visible compartments. Every compartment
needs text, `title`, or `name` and cannot contain child elements.

| Element | Typed compartments |
|---|---|
| `class` | `attribute`, `operation`, `constraint`, `note` |
| `interface` | `operation`, `constraint`, `note` |
| `enumeration` | `literal`, `operation`, `note` |
| `object` | `slot`, `note` |
| `component` | `interface` |
| `node`, `artifact` | `property`, `responsibility`, `note` |
| `package` | `responsibility`, `note` |
| `structure` | `property`, `provided-interface`, `required-interface`, `note` |
| `part` | `property`, `responsibility`, `note` |
| `profile` | `constraint`, `note` |
| `stereotype` | `property`, `constraint`, `note` |
| `metaclass` | `property`, `note` |
| `actor` | `responsibility`, `note` |
| `use-case`, `activity`, `action` | `responsibility`, `constraint`, `note` |
| `state` | `entry`, `do`, `exit`, `internal`, `region`, `note` |
| `interaction` | `note` |
| `time-state` | `region`, `constraint`, `note` |

Elements absent from this table reject compartments. Generic `<compartment>` is
accepted only where a typed compartment is allowed, for compatibility.

## Relations and Control Flow

All relations require existing local `src` and `dst`. `title` or `label`
supplies visible text; `event` supplies the visible text when both are omitted.
`guard` is valid only on control-flow, object-flow, and transition. `effect` or
`action` appends `/ ...` to the label. Multiplicity attributes are valid only on
association, aggregation, composition, and link.

The UML relation kind fixes its stroke and arrowhead. `kind`, `stroke-style`,
`arrowhead`, `start-arrowhead`, and `end-arrowhead` cannot override it.
Include/extend/import/merge/dependency-like relations are dashed; directed flows
use destination triangles; aggregation/composition use source diamonds;
structural links have no destination marker. When their label is omitted,
include, extend, import, merge, deployment, extension, and occurrence receive a
semantic default label.

Activity, state-machine, and interaction-overview graphs enforce flat V1 rules:

- `initial` has no incoming edge and at least one outgoing edge;
- `final` has no outgoing edge and at least one incoming edge;
- `decision`, `choice`, and `fork` have at least one incoming and two outgoing;
- `merge` and `join` have at least two incoming and one outgoing; and
- `history` has at least one outgoing transition.

## Messages and Timing

Sequence and communication messages require a positive dot-separated `order`
without leading zeroes, such as `1`, `2`, or `1.1`. The full string is unique
within the diagram. Numeric order is prepended to the visible label and assigns
top-to-bottom connector anchors on participants/lifelines.

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

## Reusable Models

Reusable definitions live under document-level `<data>`:

```xml
<data>
  <uml-model id="order-objects">
    <object id="customer" title="customer: Customer" />
    <object id="order" title="order42: Order" />
    <link src="customer" dst="order" title="placed" />
  </uml-model>
</data>
```

A diagram can reference the model with `data="order-objects"`. Inline children
and `data` cannot be combined. Expansion happens before the selected family
validates the model.

## Shared-Output Projection

V1 intentionally uses the capability shared by every xaligo output:

- use cases and initial/final nodes are ellipses, with activity and
  state-machine final nodes rendered as UML final states with an inner dot;
- decision, merge, choice, and history nodes are diamonds;
- other elements are editable rectangles with flattened visible compartments;
- relations are routed orthogonal connectors with separate labels;
- aggregation and composition currently share a diamond projection;
- sequence order is metadata, text, and top-to-bottom message anchoring;
- timing values are validated metadata, not proportional waveforms; and
- semantic ownership does not imply drawn containment.

XYFlow retains UML node/relation metadata in node and edge `data`. Isoflow
represents connected UML shapes with labeled generic endpoint icons and
connectors; its upstream schema cannot carry arbitrary UML metadata fields.

This is not XMI or a lossless UML interchange format.