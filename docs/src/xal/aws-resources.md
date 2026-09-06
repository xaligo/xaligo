# AWS resource tags

All bundled AWS service/resource/category icons have dedicated tags. The
[component catalog](../examples/samples/aws/README.md) contains 877 tags,
including 21 groups, with editable XAL, rendered SVG, and parameter/design notes
for each component. All 1,875 AWS catalog entries are covered; size variants
share one tag. Tabler and Yamaha icons remain available through ordinary items.

```xml
<xaligo version="2">
  <frames>
    <frame id="application" width="900" height="500">
      <vpc id="network" title="Application VPC" width="760" height="360" margin="48">
        <private-subnet id="private" title="Private subnet" width="560" height="220">
          <aws-ec2 id="api" label="API / アプリ" instance-type="t3.micro" />
        </private-subnet>
        <vpc-endpoint id="s3-access" side="right" anchor="0.65"
                      endpoint-type="gateway" service-name="S3" />
      </vpc>
      <connections>
        <connection src="api" dst="s3-access" kind="traffic" />
      </connections>
    </frame>
  </frames>
</xaligo>
```

V1 accepts the same icon/group/boundary tags with `version="1"`; native ALB/NLB
components below require V2. The original tag and
semantic ID remain available for connections, source navigation, and indexing.
Existing numeric `<item id="27"/>` syntax is unchanged.

## Naming and design

Service/resource tags use `aws-` plus a normalized catalog name, for example
`aws-ec2`, `aws-lambda`, `aws-s3-bucket`, and `aws-vpc-nat-gateway`. Category
markers use `aws-category-…`. A same-name resource icon receives a `-resource`
suffix when a service already owns the tag. Existing group names and
`vpc-endpoint` are retained. The checked-in catalog manifest is the exact list;
names are not guessed dynamically at render time.

Ordinary resources are empty, connectable icon-and-label nodes. The official
icon and its colors are fixed by the profile; `id` identifies the diagram
instance, not a catalog number. Labels are wrapped using half/full-width-aware
text measurement before layout. The default icon is 48 px, with a 160 px label
slot and a height that includes every annotation line. A numeric item still
uses its existing compact item-grid behavior.

| Parameter | Behavior |
|---|---|
| `id` | Required, unique, whitespace-free connection identifier |
| `label`, `title`, `name` | Display override, in that order; empty `label` hides the name |
| `size` | Positive finite square icon size, default 48 px |
| `label-width` | Default node/label width, default 160 px; at least `size + 12` |
| `width`, `height` | Explicit layout slot; must contain the icon and text |
| `detail` | Additional free-form diagram annotation |
| `show-details` | `false` hides annotation text; defaults to showing supplied values |
| `visible` | `false` hides the component |

Every tag's README lists its supported functional annotations. These values
affect the visible label (or group header); they do not create AWS resources,
validate an AWS account, or model the entire service API. Enum, boolean,
nonnegative-integer, and CIDR annotations are validated according to their
declared types. Runtime names, instance types, and other evolving AWS values
are deliberately text annotations rather than a stale allowlist.

Examples include EC2 `instance-type`/`state`, Lambda `runtime`/`memory-mb`,
RDS `engine`/`multi-az`, S3 bucket `bucket-name`/`versioning`, DynamoDB
`billing-mode`/`partition-key`, SQS `queue-type`, and VPC/subnet `cidr`.
Other icons expose category-appropriate annotations such as workload, stored
data, message, protected resource, input/output, or managed target.

## Network boundaries

`vpc-endpoint`, `aws-vpc-internet-gateway`, and `aws-vpc-vpn-gateway` are
icon-only VPC boundary ports. Their `side`, `anchor`, `offset`, and `size`
parameters use the [boundary placement contract](groups.md#aws-boundary-resources).
They do not consume normal row/column/grid space. Functional annotations on
these ports are retained as source/project metadata, not painted over the
boundary icon. NAT gateways and network interfaces remain interior icons.

The border position is a logical diagram convention. Interface endpoints
create network interfaces in selected subnets; gateway endpoints instead use
route-table associations. Do not infer physical interface placement from a
boundary icon. See [AWS PrivateLink concepts](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
and [gateway endpoints](https://docs.aws.amazon.com/vpc/latest/privatelink/gateway-endpoints.html).

Scope entries in the catalog are authoring recommendations, not forced
containment rules. Service, feature, policy, category, and physical-resource
icons do not all have a single deployment scope. The samples are an initial
component-by-component development baseline; advanced service behaviors are
not implied by the existence of a tag.

## Updating a component

Edit `docs/src/examples/samples/aws/<tag>/sample.xal` and its README, then run:

```sh
npm run generate:aws-samples -- --render --tag=vpc-endpoint
```

The generator preserves existing XAL and README files. SVGs are regenerated
from the current XAL. `npm run generate:aws-tags` refreshes the declarative
registry and manifest from bundled catalog data without network access;
`--check` checks that generated registry files are current. All sample pairs
are tested against both engines (including the explicit V1 rejection of native
components), and V2 SVGs are checked for staleness.

## Functional designs and research

Every component now has a functional review sheet with source links, typed
configuration fields, nested settings and related resource/API references.
The [research manifest](../examples/samples/aws/research/README.md) records
coverage, source versions and hashes. A service-context schema is explicitly
different from a schema owned by the pictured feature. Retired products and
historical icon names are marked; availability is not inferred from a catalog.

The detail cards use ordinary, editable XAL rectangles. Their field names do
not become new XAL attributes automatically. All 877 tag directories have
updated samples; ALB and NLB also contain separate connected examples:

- [ALB: HTTPS listener, rules, targets, client trust store](../examples/samples/aws/aws-elastic-load-balancing-application-load-balancer/README.md).
- [NLB: TLS termination and TCP passthrough](../examples/samples/aws/aws-elastic-load-balancing-network-load-balancer/README.md).

In icon-only form these tags accept `listener-protocol`, `listener-port`, `certificate`,
`tls-policy`, `target-group`, `target-type`, `target-protocol`, `target-port`,
`targets`, `health-check`, `scheme`, `ip-address-type`, `subnets`, and
`security-groups`. ALB additionally supports `listener-rules`,
`mutual-tls-mode="off|passthrough|verify"`, and `trust-store`. NLB supports
`alpn-policy` and `client-ip-preservation`, not ALB-style trust stores.
The component README lists protocol enums and examples. Ports must be
integers from 1 through 65535; incompatible supplied annotations are rejected.
These remain diagram annotations, not complete deployment validation.

## Native ALB/NLB components (V2)

ALB and NLB now have compact, service-specific frames: official icon and domain
tag at upper left, port cards inside, and TLS/mTLS ON/OFF badges. No watermark.
Fine design control belongs to `external/engine/src/usc/aws/`; service models
belong to `external/engine/src/ent/model/aws/`. Other services retain their
research review sheets pending component-by-component development.

```xml
<aws-elastic-load-balancing-application-load-balancer id="alb" domain="api.example.test">
  <aws-listener id="http" protocol="HTTP" port="80" target-group="web" />
  <aws-listener id="https" protocol="HTTPS" port="443" certificate="server-cert" target-group="api" />
  <aws-listener id="admin" protocol="HTTPS" port="8443" mtls="verify"
                certificate="server-cert" trust-store="client-ca" target-group="admin" />
</aws-elastic-load-balancing-application-load-balancer>
```

Component form is selected by children, `domain`, or `view="component"` and
requires 1–32 direct listeners with unique ports. `width` controls wrapping;
omitted dimensions fit the cards and measured domain. Too-small explicit
dimensions are errors. Geometry is controlled on the parent, not its listeners.
Cards automatically fit their protocol, references, target name, and visible
title. Unused reference rows are omitted; hiding the title also reduces height.
TLS and mTLS each have one badge, stacked vertically (green when ON, gray when
OFF). The default child-card heading is `Listener`; hide only that heading with
`<aws-listener id="tcp" protocol="TCP" port="443" show-title="false" />`.
Parent attributes also include `id`, `x/y`, `dx/dy`, margin/class,
fill/stroke/stroke-width/corner-radius/opacity/layer/visible. Icon-only
annotation attributes cannot be mixed into component form.

| Listener parameter | Meaning |
|---|---|
| `id`, `protocol`, `port` | Required connection ID, protocol, integer 1–65535 |
| `mtls` | `off` (default), ALB `verify`, or ALB `passthrough` |
| `certificate` | Server-certificate reference; HTTPS/TLS only |
| `trust-store` | Client CA reference; required for ALB verify, invalid otherwise |
| `target-group` | Display reference; short name shown, full value retained in source |
| `backend-tls`, `backend-mtls` | Optional validated target security metadata; no extra listener badges |
| `show-title` | Defaults to `true`; `false` hides “Listener” and removes its layout space |
| `visible` | Hide the listener and its decorations with `false` |

ALB accepts HTTP/HTTPS. NLB accepts TCP/TLS/UDP/TCP_UDP/QUIC/TCP_QUIC.
TLS badges indicate **listener TLS termination**, not whether TCP carries
encrypted bytes. NLB TCP with `backend-tls="true" backend-mtls="true"` displays
only listener TLS/mTLS OFF; indicate backend TLS/mTLS on the connected target.
NLB cannot terminate mTLS;
ALB passthrough forwards client certificates in HTTP headers. Certificate and
CA badges indicate supplied references, not verified validity.
[ALB mTLS](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/mutual-authentication.html),
[NLB listeners](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html).

Connect traffic explicitly to listener or rule IDs. ALB also supports native
rules, all six condition/action types, both rewrite types, target groups,
registered targets, and typed configuration cards. See [ALB rules and detail
levels](alb.md) and the [option reference](alb-options.md). Use
`detail-level="summary"` for an automatically sized overview, or `show`, `hide`,
and per-card `visible="false"` for selective detail. Certificate/CA/CRL,
TLS-policy and subnet/IP placement remain references or option cards, not
independent provisioned resources. This is diagram validation, not AWS
provisioning validation. Empty icon-only ALB/NLB forms remain V1-compatible;
the compound forms and their dedicated child tags fail explicitly on V1.

Use the normal rendering command to preserve edits and refresh every SVG in
one tag directory. To intentionally replace the entire generated design
library, run `npm run generate:aws-designs -- --update`, followed by
`npm run generate:aws-samples -- --render`. The explicit migration saves
overwritten sources in a temporary backup directory. `--check` verifies
reproducibility against the checked-in research inputs without network access.
