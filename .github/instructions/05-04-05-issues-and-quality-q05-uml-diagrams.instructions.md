---
applyTo: ".github/instructions/manual/**"
---

# 05.04.05 Issues and quality: Q05 UML Diagrams

### Q05 UML Diagrams

| Task | Status | Scope | Verification target |
|---|---|---|---|
| Q05.2 | in-progress | Establish a per-UML visual baseline before edits: validate and render each supported individual `uml-*.xal` sample to SVG, then identify overlap, spacing, typography, connector, and semantic-notation gaps. | all retained UML samples validate; tracked SVG baselines exist, with component and activity visual audits completed. |
| Q05.3 | in-progress | Improve activity-diagram semantic accuracy: initial/final nodes, actions, object nodes, decisions, forks, joins, merges, responsibilities, constraints, guards, `control-flow`, and `object-flow`. | vertical and horizontal partition samples validate; focused parser/scene tests cover partitions, loop routes, and activity metadata. |
| Q05.4 | in-progress | Improve activity-diagram design quality: left-to-right reading flow, diamond/bar/final-node proportions, lane spacing, label placement, and control-vs-object-flow distinction. | vertical and horizontal activity SVG baselines are regenerated; focused geometry tests cover lane placement and loop routing. |
| Q05.6 | in-progress | Improve class-diagram design quality: compartment rhythm, long member wrapping, stereotype readability, relation label spacing, crow-foot/diamond marker clarity, and dense-layout crossing reduction. | class SVG/PPTX review and text-fit tests; relation-label spacing fixed so labels avoid unrelated classifier boxes in the same frame, not only their own connection endpoints (`TestUMLRelationLabelsAvoidUnrelatedClassifierBoxes`, regenerated `docs/src/images/uml-class.svg`); compartment rhythm, long member wrapping, crow-foot/diamond marker clarity, and dense-layout crossing reduction remain to be reviewed |
| Q05.7 | not-started | Improve sequence-diagram semantic accuracy: lifelines, participants, activation bars, sync/async messages, returns, self messages, create/delete, ordering, and message labels. | `uml-sequence.xal` layout tests and route ordering assertions |
| Q05.8 | not-started | Improve sequence-diagram design quality: timeline spacing, activation contrast, message arrow style, return-line distinction, lifeline header readability, and vertical density. | sequence SVG/PPTX review and geometry assertions |
| Q05.9 | not-started | Improve state-machine semantic accuracy: initial/final states, composite states, transitions, events, guards, effects, entry/do/exit actions, and invalid transition diagnostics. | `uml-state-machine.xal` valid/invalid tests |
| Q05.10 | not-started | Improve state-machine design quality: state shape proportions, nested-state padding, transition bend points, guard/effect label placement, and final-node readability. | state-machine SVG review and collision assertions |
| Q05.11 | in-progress | Improve component-diagram semantic and design quality: boundary interfaces, shared interface-name widths, compact automatic height, explicit sizing, fan-out, and connector routing. | component parser/layout/scene/routing tests plus `uml-component.xal` and its regenerated SVG baseline |
| Q05.12 | not-started | Plan future communication-diagram support only when ordered topology semantics are explicitly required beyond sequence/state-machine diagrams. | plan entry only; no implementation, samples, or generated assets until restarted |
| Q05.13 | not-started | Keep object, use-case, profile, and interaction-overview out of the supported UML set unless a non-substitutable use case is identified. | unsupported parser diagnostics and documentation review |
| Q05.17 | not-started | Accepted addition: design and implement timing-diagram support (`timing-diagram` selector) — lifeline/state-timeline elements, state/value change events, duration constraints, and time-axis layout — as a new UML family alongside class/component/activity/state-machine/sequence. | new `uml-timing.xal` sample; parser/scene/layout/routing tests; diagram-kind vocabulary and relation-projection tables in `07-00-xal-specification-overview.instructions.md` updated once implemented |
| Q05.14 | not-started | Normalize shared UML visual language across all diagram types: typography scale, stroke weights, marker sizes, semantic colors, label backgrounds, and light/dark contrast. | design review checklist plus theme render comparisons |
| Q05.15 | not-started | Verify UML cross-format parity for every shared scene/plan change: SVG baseline, Excalidraw editability, PPTX plan/export, PDF/Excel page projection, and XYFlow/Isoflow applicability. | focused renderer matrix for changed UML contracts |
