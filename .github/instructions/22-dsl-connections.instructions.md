---
applyTo: "**/*.xal"
---

# XAL connections

Resolve orthogonal routes, endpoint references, sides/anchors, obstacles,
lanes, junctions, jumps, labels, and cross-frame stubs in shared layers.
`kind="route"` is headless; traffic is directional.

Cross-frame links become page-local `to`/`from` stubs. Endpoint side/anchor and
frame-terminal side/anchor are independent and must honor metadata reservation
and inset rules. Exact syntax and precedence: `reference.md` section `07`
“`<connection>` Tag” and section `10`.
