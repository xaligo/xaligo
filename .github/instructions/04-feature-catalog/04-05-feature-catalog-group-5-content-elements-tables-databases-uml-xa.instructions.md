---
applyTo: ".github/instructions/manual/**"
---

# 04.05 Feature catalog: Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)

## Group 5 — Content Elements: Tables, Databases, UML (`XAL-5xxxxxx`)

| ID | Feature | Status | Summary |
|---|---|---|---|
| XAL-5000010 | `<table>` GFM-like pipe syntax | Implemented | Markdown-style pipe table with header/alignment row, normalized to the same typed rows as tag syntax. |
| XAL-5000020 | `<table>` explicit tag syntax | Implemented | `<header>`/`<row>`/`<cell>` children as an alternative to pipe syntax, normalized identically before layout. |
| XAL-5000030 | Table style inheritance | Implemented | Table-level `color`/`background-color`/`border-color`/`font-*` and `header-*` overrides, with `cell > header/row > table > built-in` precedence. |
| XAL-5000040 | Imported table schemas | Implemented | Table structure imported from an external schema source instead of authored inline in `.xal`. |
| XAL-5000050 | `<database-schema>` reusable definitions | Implemented | Document-level `<data>` schema definitions referenced by one or more frames via `<database data="schema-id">`. |
| XAL-5000060 | `<entity>`/`<column>` relational definitions | Implemented | Typed entity/column definitions supporting `name`, `type`, `primary-key`, `nullable`, `unique`, and `default`. |
| XAL-5000070 | `<foreign-key>` relation generation | Implemented | Single-column foreign key with a `references="entity.column"` target that generates an inter-entity relation. |
| XAL-5000080 | Relational key/nullability styling | Implemented | Primary/foreign key and nullability/data-type visual styling on entity column rows. |
| XAL-5000090 | Crow-foot relation notation | Implemented | Relation endpoints render with crow-foot cardinality notation on database diagrams. |
| XAL-5000100 | `<uml>` shared diagram component | Implemented | Common V1 component adapting UML families to xaligo's shared layout, shape, connector, and output pipeline. |
| XAL-5000110 | UML class diagrams | Implemented | `class`, `interface`, `enumeration` elements with `association`, `aggregation`, `composition`, `generalization`, `realization`, `dependency` relations. |
| XAL-5000120 | UML component diagrams | Implemented | `component`, `interface`, `port`, `artifact` elements with `dependency`, `realization`, `association`, `assembly`, `delegation` relations. |
| XAL-5000130 | UML activity diagrams | Implemented | `initial`/`final`/`activity`/`action`/`decision`/`merge`/`fork`/`join`/`object-node` elements with `control-flow`/`object-flow` relations. |
| XAL-5000140 | UML activity partitions/swimlanes | Implemented | `<partition>` swimlanes with `lanes="vertical|horizontal"` and the `theme="xaligo"` swimlane visual theme. |
| XAL-5000150 | UML state-machine diagrams | Implemented | `state`/`history`/`choice`/`fork`/`join`/`initial`/`final` elements with `transition` relations, events, guards, and effects. |
| XAL-5000160 | UML sequence diagrams | Implemented | `participant`/`lifeline` elements with `message`/`return-message`/`create-message`/`destroy-message` relations and explicit diagram-unique `order`. |
| XAL-5000170 | UML element compartments | Implemented | Typed ordered text compartments per element kind (`attribute`, `operation`, `constraint`, `note`, `entry`/`do`/`exit`, etc.). |
| XAL-5000180 | UML relation attributes | Implemented | Shared relation `label`/`guard`/`route` hint and `src-multiplicity`/`dst-multiplicity` on association-family relations. |
| XAL-5000190 | UML relation projection | Implemented | Fixed lowering of UML relation kinds to the shared orthogonal connector model (dashed/solid line, source diamond, no-arrowhead). |
| XAL-5000200 | Reusable `<uml-model>` definitions | Implemented | Document-level `<data>` UML models selected by one or more diagram-kind children via `data="model-id"`. |
| XAL-5000210 | Public UML connection endpoint references | Implemented | `uml-id/local-id` (same frame) and `frame-id.uml-id/local-id` (cross-frame) public endpoint forms for normal `<connection>` tags. |
| XAL-5000220 | Deliberately lossy V1 UML projection | Implemented | Documented, intentional projection limits (no dashed lifelines/activations/combined fragments, flattened compartments) applied consistently across every output format. |
| XAL-5000230 | UML timing diagrams | Planned | Accepted new UML family (Q05.17): `timing-diagram` selector with lifeline/state-timeline elements, state/value-change events, duration constraints, and time-axis layout. |
