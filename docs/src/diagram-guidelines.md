# Beautiful Diagram Guidelines

A good architecture diagram should be understandable before it is examined in
detail. The reader should be able to identify the system boundary, major
components, and primary flow within a few seconds. This guide describes a
repeatable way to achieve that result with Xaligo.

## Start with the message

Decide what the diagram must explain before adding elements. A diagram should
normally answer one main question, such as:

- Where does a request enter, and which components process it?
- Which resources belong to each network or security boundary?
- How is data replicated, stored, or delivered?
- Which components are global, regional, VPC-level, or AZ-specific?

If one page must answer several unrelated questions, split it into an overview
and one or more detail frames. Identified child frames naturally become
separate SVG artifacts or PowerPoint slides.

## Build a clear visual hierarchy

Arrange the diagram from large concepts to small details:

1. Put the system or cloud boundary at the outside.
2. Nest regions, networks, zones, and logical groups inside it.
3. Place services and components inside the smallest truthful boundary.
4. Add connections only after the component layout is stable.

Use AWS-specific group tags only for their actual meaning. For example,
`<public-subnet>` represents an internet-routable subnet; it is not a generic
blue box. Use `<generic-group>` for application tiers, security services,
observability, CI/CD, or other logical collections.

AWS service placement should reflect its real scope:

| Scope | Recommended placement | Examples |
|---|---|---|
| Global | Inside the cloud, outside regions | Route 53, CloudFront, IAM |
| Regional | Inside a region, outside the VPC | S3, Lambda, CloudWatch |
| VPC | Inside the VPC, outside availability zones | Internet gateway, load balancer |
| AZ | Inside an availability zone and subnet | EC2, NAT gateway, database instance |

See [AWS Groups](xal/groups.md) for the available containers and their
semantics.

## Prefer a simple reading direction

Choose one dominant flow direction and keep it consistent. Left-to-right works
well for request and data flows; top-to-bottom works well for layered systems.
Avoid making the reader alternate between both directions.

- Put entry points near the top or left edge.
- Put processing components in the middle.
- Put databases, queues, and external destinations near the bottom or right.
- Align siblings and give repeated structures the same size.
- Keep similar gaps between peers; larger gaps should indicate a change of
  boundary or responsibility.
- Prefer whitespace over decorative elements. Empty space makes grouping and
  flow easier to perceive.

Use fixed layouts for small, predictable groups and grid or flex layouts for
repeated structures. Use absolute positioning only when the intended geometry
cannot be expressed structurally; extensive manual coordinates are difficult
to maintain and tend to produce uneven diagrams. See [Layout](xal/layout.md)
for the supported layout models.

## Keep labels concise and consistent

Labels should describe the reader-facing role of an element, not repeat every
implementation detail.

- Use short nouns for components: `API`, `Worker`, `Orders DB`.
- Use verb phrases for important flows: `Publish event`, `Read replica`.
- Use one capitalization style throughout the diagram.
- Avoid unexplained abbreviations. If an abbreviation is necessary, define it
  in the legend or surrounding text.
- Keep group titles distinct from child labels.
- Put long explanations in frame metadata, a Markdown caption, or adjacent
  documentation instead of shrinking text to fit a box.

When AWS icons are used, supply `services.csv` so labels and legends use the
official service names consistently. See [Services and Legends](services.md).

## Use color with restraint

Color should encode meaning, not fill empty space. Start with the selected
light or dark theme, then add accent colors only for information the reader
must distinguish.

A practical palette uses:

- one neutral treatment for normal structure;
- one primary accent for the main path;
- one secondary accent for an alternate or asynchronous path;
- red, orange, or yellow only for risk, failure, warning, or trust boundaries.

Do not rely on color alone. Pair color with labels, line styles, arrowheads, or
group boundaries so the diagram remains understandable when printed in
grayscale or viewed by readers with color-vision differences. Reuse exactly
the same color for the same meaning on every frame.

## Make connections easy to follow

Connections should explain relationships without becoming the dominant visual
feature.

- Connect the nearest sensible sides of components.
- Prefer short orthogonal routes with few bends.
- Avoid diagonal lines crossing labels or icons.
- Keep parallel flows aligned and separated.
- Place heavily connected components near the center of their neighbors.
- Use explicit ports when several relationships meet one component.
- Label only relationships whose purpose is not obvious.

Use `kind="route"` for structural network paths and `kind="traffic"` for
directional communication. Do not use an arrowhead on a route. When several
lines cross, first improve component placement; use line jumps only after the
layout itself is clear. Detailed controls are documented under
[Connections](xal/connections.md) and the [Arrow Reference](reference/arrows/index.md).

## Keep detail proportional to the output

Design for the final viewing size. A diagram that looks comfortable when
zoomed in may be unreadable on an A4 page or presentation slide.

- Use one frame per page or slide for normal documents.
- Prefer landscape orientation for wide request flows and multi-column network
  diagrams.
- Prefer portrait orientation for layered or sequential flows.
- Use paper margins to keep content away from slide and print edges.
- Avoid using `--combine-frames` for unrelated pages.
- Check that text remains readable at 100% zoom and in a slide-show view.

SVG and PPTX share resolved geometry, but both outputs must still be reviewed.
PowerPoint fonts and editable shapes can expose spacing issues that are less
obvious in SVG.

## A practical authoring workflow

Work from structure toward presentation:

1. Write the diagram's question and intended audience.
2. List the essential components; remove anything that does not support the
   question.
3. Create truthful boundaries and place components at the correct scope.
4. Establish a single reading direction and align sibling groups.
5. Add the main flow, then secondary connections.
6. Apply concise labels and a restrained semantic palette.
7. Validate and render both required output formats.
8. Review at the final viewing size and simplify again.

Useful commands:

```bash
xaligo validate architecture.xal
xaligo render architecture.xal --format svg -o architecture.svg
xaligo render architecture.xal --format pptx -o architecture.pptx \
  --paper A3 --orientation landscape
xaligo serve architecture.xal
```

For AWS diagrams, include the service catalog:

```bash
xaligo render architecture.xal --format svg -o architecture.svg \
  --services services.csv --svg-legend-position right
```

Use live preview while adjusting hierarchy and spacing, but make the final
decision from the exported SVG and PPTX rather than the editor source alone.

## Common problems

| Problem | Better approach |
|---|---|
| Every available service is shown | Keep only components needed to explain the selected scenario |
| Different concepts use identical boxes | Use truthful boundaries, icons, labels, and a small number of semantic styles |
| Global services appear inside an AZ | Move each service to its real infrastructure scope |
| Lines cross through labels and icons | Reorder components, shorten routes, and introduce explicit ports |
| Many colors have no defined meaning | Reduce the palette and document each accent's purpose |
| Text is made tiny to fit | Shorten labels, enlarge the frame, or split the diagram |
| Manual coordinates dominate the source | Replace repeated coordinates with grid, flex, rows, or columns |
| One frame mixes overview and implementation detail | Create separate overview and detail frames |

## Final review checklist

Before publishing, verify:

- [ ] The title or surrounding text states the diagram's purpose.
- [ ] The main boundary and primary flow are recognizable at a glance.
- [ ] Every component is inside the correct semantic and infrastructure scope.
- [ ] Siblings are aligned, evenly spaced, and free from overlap.
- [ ] Labels are concise, consistent, and readable at the final size.
- [ ] Icons, labels, and `services.csv` entries agree.
- [ ] Connections avoid components, labels, group headers, and unnecessary
      crossings.
- [ ] Color meanings are consistent and are not the only distinguishing cue.
- [ ] `xaligo validate` succeeds without unexplained warnings.
- [ ] SVG and PPTX outputs have both been inspected.
- [ ] The `.xal` source is committed with any generated documentation asset.

For complete examples, see [Samples](samples.md), [Examples](examples/index.md),
and the [XAL DSL overview](xal/overview.md).
