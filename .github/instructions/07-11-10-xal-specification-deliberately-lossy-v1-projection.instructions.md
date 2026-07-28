---
applyTo: ".github/instructions/manual/**"
---

# 07.11.10 XAL specification: Deliberately lossy V1 projection

### Deliberately lossy V1 projection

V1 preserves the selected UML family, element kind, relation kind, owner, and
relation label in the shared semantic scene, then projects them into the
capabilities common to xaligo outputs:

- `initial` and `final` become ellipses;
- `decision`, `merge`, `choice`, and `history` become diamonds;
- every other element becomes an editable rectangle whose ordered
  compartments are flattened into its visible text;
- every relation becomes a shared orthogonal connector with a separate label;
  aggregation and composition currently share the same diamond projection;
- sequence order is retained in labels and metadata and controls top-to-bottom
  message anchors, but V1 does not draw dashed lifelines, activations, combined
  fragments, or a separate vertical event axis;
- current V1 UML elements do not support semantic ownership.

SVG, Excalidraw, PPTX, PDF, Excel, XYFlow, and Isoflow all consume this same resolved
geometry. Excalidraw-compatible output carries xaligo UML custom data for
editing. XYFlow retains UML node and relation fields in node/edge `data` and
records the projected node shape. Isoflow projects every connected UML shape
to a labeled generic endpoint icon because its upstream connector schema has no
arbitrary UML data field. Another target schema may omit marker details it
cannot represent. An encoder must use native target constructs where available
and must not add private schema-breaking fields. The output is not XMI and is not a
lossless UML interchange representation.
