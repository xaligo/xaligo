# AWS Groups

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

Group header and item label width estimates count East Asian full-width
characters as double-width so Japanese and other full-width labels stay aligned
across Excalidraw, SVG, and PPTX.
