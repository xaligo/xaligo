# AWS Groups

For dedicated service/resource icons and per-component XAL/SVG examples, see
[AWS resource tags](aws-resources.md) and the
[AWS component catalog](../examples/samples/aws/README.md).

AWS group tags render architecture-style containers with headers, icons, and
border styles. Examples include:

- `<aws-cloud>`
- `<aws-account>`
- `<region>`
- `<vpc>`
- `<availability-zone>`
- `<public-subnet>`
- `<private-subnet>`
- `<security-group>`
- `<generic-group>`
- `<capture>`

Every group tag requires a non-empty `id`. IDs must be unique among frame-like
components so groups can be used as connection endpoints.

```xml
<aws-cloud id="cloud" title="AWS Cloud">
  <region id="region-apne1" title="ap-northeast-1">
    <vpc id="prod-vpc" title="Production VPC">
      ...
    </vpc>
  </region>
</aws-cloud>
```

## Generic Group Icons

`generic-group` accepts `icon-id`, using the same catalog as `<item>`. It must
be a positive signed 32-bit decimal ID (`1..2147483647`); zero, signs,
non-decimal syntax, and out-of-range values are invalid.

```xml
<generic-group id="network-topology" title="Network Topology" icon-id="104635">
  <item id="200036" />
</generic-group>
```

## VPC endpoints on the VPC border

Use `<vpc-endpoint>` for an Amazon VPC Endpoint that should sit on the VPC
boundary. It is excluded from the VPC's normal child layout, so moving it does
not resize or reorder subnets and other resources.

```xml
<vpc id="application-vpc" title="Application VPC">
  <vpc-endpoint id="private-api" side="right" anchor="0.35" />
  <private-subnet id="application" title="Application Subnet">
    <item id="27" name="app" />
  </private-subnet>
</vpc>
```

`side` accepts `top`, `right`, `bottom`, or `left` and defaults to `right`.
Set `anchor` from `0` to `1` to slide the icon along that edge; `offset` adds a
pixel adjustment in the same direction. `size` controls the square icon size
and defaults to `48`. If several endpoints on one side omit `anchor`, they are
distributed evenly. Overlapping endpoints and positions beyond the usable
edge are rejected.

The element must be empty, must be a direct child of `<vpc>`, and requires a
unique, whitespace-free `id` used by connections. Its AWS catalog icon is
fixed to Amazon VPC Endpoints, and its center is drawn directly on the VPC
line. It is intentionally icon-only; use the service legend or nearby text
when an explanatory label is needed. A numeric `<item id="1579">` remains a
regular item inside normal layout for backward compatibility.

## Capture (Structural Annotation)

`<capture>` is a lightweight structural annotation group. It participates in
normal nested layout and places its children inside a border (and optional title
band) without implying any AWS/architectural semantics, unlike
`<generic-group>` and the AWS boundary tags above. Use it to highlight a
diagram region such as a "hot path".

`<capture>` is connectable exactly like any other group tag, including the
`frameId.id` qualified cross-frame form, so a connection to/from a `<capture>`
in another frame renders the same "to `<frame>`" / "from `<frame>`"
cross-frame page-link stubs used for any other connectable endpoint.

```xml
<container gap="24">
  <capture id="hot-path" title="Hot Path">
    <rectangle id="public-api" title="Public API" />
  </capture>
  <rectangle id="batch-export" title="Batch Export" />
</container>
<connection src="hot-path" dst="batch-export" />
```

## Text Width

Group headers estimate proportional half-width glyphs separately from East
Asian wide/full-width glyphs, which use one font-size unit. This keeps Japanese,
half-width Katakana, and mixed-width titles compact while preserving a small
cross-renderer safety allowance and 4px before the tag tip. Item labels still
count full-width characters as two display columns when calculating wrapping.
Markdown inherits SVG text layout.
