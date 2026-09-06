---
applyTo: ".github/instructions/manual/**"
---

# 07.15.03 XAL specification: Dedicated AWS resource tags

The checked-in AWS profile registry covers every bundled AWS service,
resource, category, and group icon. Size variants share a dedicated tag;
same-name service/resource icons remain distinct. Existing group names and
`vpc-endpoint` are stable; other tags use `aws-` and category tags use
`aws-category-`. The manifest is authoritative, not heuristic tag matching
at runtime. Numeric `<item>` and non-AWS catalogs retain their old contracts.

Keep profiles declarative in `internal/core/profiles/aws`. Generate registry
data and the documentation manifest with `npm run generate:aws-tags`; no
network or engine callbacks belong in a profile. Both Go frontends preserve
semantic IDs and tags while converting resource/group/boundary profiles into
their engine's generic concepts. V2 ALB/NLB native components are the explicit
exception defined below and in 08.10; they use typed Rust service models.

Unless the native component form below is selected, dedicated
service/resource/category tags are empty leaves with required unique
connection IDs. Their fixed official icon defaults to 48px. A 160px default
label slot, explicit half/full-width-aware wrapping, and reserved annotation
height prevent clipped labels. `label`/`title`/`name` override the display name;
`detail` adds text; `show-details="false"` hides supplied annotations. `size`
and `label-width` must be positive finite values. Explicit width/height must
contain the icon and label. Icon overrides and boundary attributes are invalid
on an ordinary resource. Invisible resources do not draw themselves.

Functional annotations are declared per tag/family with enum, boolean, text,
nonnegative-integer, and CIDR schemas. They are diagram annotations, not AWS API
configuration or deployment validation. Only explicitly supplied values are
shown. Group annotations join the single-line header. Boundary annotations
remain source/project metadata; boundary icons remain text-free.

`vpc-endpoint`, `aws-vpc-internet-gateway`, and `aws-vpc-vpn-gateway` are direct
VPC-child boundary ports using the existing side/anchor/offset/size contract.
Default and explicit instances of the same side share auto-distribution.
Interior resources (including the subnet-scoped NAT example) do not snap to
borders. Endpoint border placement is a logical authoring convention, not a
claim that interface endpoint ENIs live on a VPC's physical perimeter.
All other scope metadata is advisory because many icons represent services,
policies, features, or logical concepts rather than deployable instances.

For every tag, maintain `docs/src/examples/samples/aws/<tag>/sample.xal`, its
rendered `sample.svg`, and a README specifying implemented parameters/design
and review notes. Bootstrap only missing sources; never overwrite subsequent
per-component edits. Refresh SVGs from current XAL via
`npm run generate:aws-samples -- --render [--tag=<tag>]`. Cover all catalog IDs,
both frontends, both renderers, LSP completions, validation, and stale sample
SVGs with regression tests. These source-controlled SVGs are intentional user
deliverables, not build artifacts.

## Functional research and review sheets

The sample library includes a research classification for every tag. Preserve
the distinction between an exact resource-schema mapping, a service-context
schema, an API model, software/workflow concepts, and a category/group/symbol.
Do not claim a feature icon is a deployable resource merely because its parent
service has a CloudFormation type. Preserve lifecycle warnings for retired
services and separate current service capability from historical icon names.

Research inputs are factual, offline CloudFormation and AWS SDK model snapshots
under `docs/src/examples/samples/aws/research`; retain source/version/date/hash
provenance. They are documentation inputs, not engine runtime dependencies.
`generate:aws-designs -- --update` is an explicit bulk migration which backs up
overwritten sources. Ordinary `generate:aws-samples -- --render` never updates
the editable sources, and renders every `.xal`/`.svg` pair in each tag folder.
V2 examples have one frame per file; alternative configurations use extra pairs.

Configuration cards are editable generic rectangles with semantic IDs. Their
AWS schema field names are NOT automatically supported XAL attributes. README
parameter tables must identify implemented tag attributes separately. Size
cards for all wrapped lines, including half/full-width labels. Configuration
references use dashed lines; request traffic uses solid lines. Do not derive
ownership or network connectivity from membership in an API namespace.

ALB/NLB tags expose scheme, IP family, subnets, security groups, listener
protocol/port, server certificate/TLS policy, target group/type/protocol/port,
targets and health-check annotations. Ports are integers in 1..65535. ALB has
listener-rules, mutual-tls-mode and trust-store; NLB has ALPN and client-IP
preservation annotations. NLB rejects ALB-style trust stores and mTLS listener
termination. ALB passthrough sends certificate information in HTTP headers;
NLB TCP passthrough transports encrypted bytes for backend TLS/mTLS. Validate
incompatible supplied annotations in the shared Go profile, while allowing
partial diagrams. Do not present these checks as complete AWS validation.

## Native ALB/NLB components (V2 only)

The ALB/NLB tags select component form when they contain `aws-listener`
children, supply `domain`, or use `view="component"`. Icon-only leaf form
remains compatible with V1; compound form and `aws-listener` must explicitly
fail on standalone V1 with a version-2 requirement, never fall back silently.

Component attributes: required unique `id`, optional single-line `domain`,
`view="component"`, `width`, `height`, `x`, `y`, `dx`, `dy`, margin/class,
fill/stroke/stroke-width/corner-radius/opacity/layer/visible. At least one and
at most 32 direct listener children; no arbitrary text or generic children.
ALB additionally accepts `aws-target-group` and `aws-option` children.
Leaf annotation parameters are not accepted in component form. The native
stage measures header width with half/full-width-aware text metrics, places
the fixed official icon and domain tag at the upper left, and wraps listener
cards into rows. Card width/height fit visible content rather than fixed slots;
absent references and hidden headings reserve no height. No watermark is rendered. Explicit dimensions must contain
the header and cards; weight and listener geometry overrides are unsupported.

`aws-listener` is a connectable child with required `id`, `protocol`,
and integer `port` in 1..65535. Optional parameters: `mtls` = off/verify/passthrough,
`certificate`, `trust-store`, `target-group`, `backend-tls`, `backend-mtls`,
and `visible`, `show-title`. The default title is “Listener”; `show-title="false"`
hides only this title and removes its height. IDs and ports must be unique within their respective scopes.
ALB supports HTTP/HTTPS; mTLS requires HTTPS, verify requires a trust-store,
and a trust-store is invalid in other modes. NLB supports TCP/TLS/UDP/TCP_UDP/
QUIC/TCP_QUIC and rejects ALB-style mTLS/trust stores. NLB backend mTLS requires
TCP passthrough and backend TLS. ALB cannot present a backend client certificate.
Certificates require HTTPS/TLS. TLS ON/OFF indicates listener TLS termination;
Exactly one TLS badge and one mTLS badge are stacked vertically, green for ON
and gray for OFF. Backend TLS/mTLS flags remain validated target metadata,
not additional listener badges; show target security on the connected target.
ALB mTLS PASS means certificate forwarding in HTTP headers, not raw TLS relay.
Certificate/CA presence is compactly indicated, not proof of validity. Reference
values are diagram metadata, not AWS lookups; full values stay in source, long
target-group display names may be abbreviated. Target links are authored
explicitly with connections to semantic IDs. ALB listeners accept dedicated
rules and option cards. NLB does not accept ALB feature children.

## Native ALB rules and abstraction (V2)

Maintain the implemented grammar, examples and complete closed option list in
`docs/src/xal/alb.md` and `alb-options.md`. Dedicated child tags cover rules,
conditions/matches, actions/weighted targets/JWT claims, transforms/rewrites,
target groups/registered targets, and options. All IDs remain stable and
connectable. Six condition types, six action types and two transform types are
supported. Validate ownership, priorities/action order, comparison limits,
TLS/authentication compatibility and supplied typed settings before rendering,
including hidden nodes. Allow partial diagrams and external target references;
do not claim complete AWS provisioning or regex/traffic evaluation.

Render fine-grained `aws-option` children as one two-column Setting/Value table
per owner, not as independent cards. Keep every source option ID as the semantic
row owner, wrap each cell independently, derive row height from the taller cell,
and omit hidden rows without retaining space. SVG and PPTX consume the same
generic cell shapes and text.

Target groups use AWS registration types `ip`, `instance`, or `lambda`.
`aws-target-service` is a separate logical workload child for ECS, EKS, EC2,
Lambda, or generic IP services; it may contain actual `aws-registered-target`
children and service-owned options. Preserve this distinction. Validate ECS
Fargate/awsvpc as IP targets, ECS bridge/host as instance targets, and EKS pod
IP versus worker-node/NodePort modes. Display known services with their fixed
official icon. Service membership is a diagram relationship, not an assertion
that ECS/EKS service objects themselves are registered with ELB.

`detail-level=summary|standard|detailed`, comma-separated `show`/`hide`, and
per-child `visible=false` control abstraction. Unset levels inherit; parent
category suppression takes precedence over a child's visibility. Standard
hides option cards; summary also hides rules/references/targets; detailed
shows supplied fields. Native composition measures only visible cards, removes
hidden subtree allocation, and remaps hidden connection endpoints to the
nearest visible ancestor. Same-anchor collapsed lines are hidden. Keep source
models unchanged. Old compact ALB/NLB rendering remains when no new features
or presentation controls are supplied. Reference values remain source metadata,
not secret storage or AWS API lookups. TLS policies, certificates, CA/CRL and
IP/AZ placements are reference/option cards, not separately deployed resources.

Service definitions and composition must remain in per-component files under
`ent/model/aws` and `usc/aws`, declared only in `external/engine/src/mod.rs`.
Use the versioned typed ABI (currently v5); no arbitrary maps or service logic
in generic layout, SVG/PPTX encoders, or Go profile callbacks. Update ABI tests,
SVG/PPTX shared-plan tests and every affected editable sample/SVG pair.
