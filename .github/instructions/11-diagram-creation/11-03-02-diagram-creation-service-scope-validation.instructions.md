---
applyTo: ".github/instructions/manual/**"
---

# 11.03.02 Diagram creation: Service Scope Validation

### Service Scope Validation

Before finalizing the `.xal`, verify that each service is placed at the correct scope level.
Placing a global or regional service inside an `<availability-zone>` is misleading.

| Scope | Placement in .xal | Typical services |
|---|---|---|
| **Global** | Direct child of `<aws-cloud>`, inside `<generic-group>` | Route 53, CloudFront, IAM, WAF |
| **Regional** | Inside `<region>`, outside `<vpc>`, inside `<generic-group>` | Lambda, S3, CloudWatch, SQS, SNS, EventBridge, Step Functions, CodePipeline, Macie |
| **VPC-level** | Inside `<vpc>`, outside `<availability-zone>`, inside `<generic-group>` | Internet Gateway, ELB/ALB, Secrets Manager |
| **AZ-specific** | Inside `<availability-zone>`, in `<public-subnet>` / `<private-subnet>` | EC2, NAT Gateway, RDS instance, Aurora replica, ElastiCache node, ECS task, EKS node |

> **Incorrect:** placing Route 53 or IAM inside `<availability-zone>` — these services are not AZ-bound.
> **Correct:** group them under `<generic-group title="Global Services">` as a direct child of `<aws-cloud>`.

Quick checklist:
- [ ] Global services (Route 53, CloudFront, IAM, WAF) → outside `<region>`
- [ ] Regional managed services (Lambda, S3, SQS, etc.) → inside `<region>`, outside `<vpc>`
- [ ] Network edge (IGW, ELB) → inside `<vpc>`, outside `<availability-zone>`
- [ ] Compute/DB instances → inside `<availability-zone>`
- [ ] Services not tied to a VPC → never inside `<vpc>` or `<availability-zone>`

```xml
<frame version="1" width="1440" height="900" class="pa-4">
  <aws-cloud id="aws-cloud" title="AWS Cloud">

    <!-- ✅ Global: outside <region> — not bound to any specific region -->
    <generic-group id="global-services" title="Global Services">
      <item id="1179" />  <!-- Route 53 -->
      <item id="216"  />  <!-- IAM -->
    </generic-group>

    <region id="region-apne1" title="ap-northeast-1" row="8">

      <!-- ✅ Regional: inside <region>, outside <vpc> — no VPC required -->
      <generic-group id="managed-serverless" title="Managed &amp; Serverless">
        <item id="13"   />  <!-- Lambda -->
        <item id="1020" />  <!-- S3 -->
      </generic-group>

      <vpc id="vpc-main" title="VPC (10.0.0.0/16)" row="6">

        <!-- ✅ VPC-edge: inside <vpc>, outside <availability-zone> -->
        <generic-group id="vpc-edge" title="VPC Edge">
          <item id="1581" />  <!-- Internet Gateway -->
          <item id="1182" />  <!-- ELB -->
        </generic-group>

        <row gap="8" row="5">
          <col span="6">
            <availability-zone id="az-apne1a" title="AZ: ap-northeast-1a">
              <!-- ✅ AZ-specific: public-subnet for NAT Gateway -->
              <public-subnet id="public-subnet-a" title="Public Subnet">
                <item id="1582" />  <!-- NAT Gateway -->
              </public-subnet>
              <!-- ✅ AZ-specific: compute instances in private subnet -->
              <private-subnet id="app-tier-a" title="Application Tier" row="3">
                <item id="27"  />   <!-- EC2 -->
                <item id="547" />   <!-- ECS -->
              </private-subnet>
            </availability-zone>
          </col>
          <col span="6">
            <availability-zone id="az-apne1b" title="AZ: ap-northeast-1b">
              <!-- ✅ AZ-specific: DB instances in private subnet -->
              <private-subnet id="data-tier-b" title="Data Tier">
                <item id="117" />   <!-- RDS -->
                <item id="110" />   <!-- Aurora -->
              </private-subnet>
            </availability-zone>
          </col>
        </row>

      </vpc>
    </region>
  </aws-cloud>

  <connection src="1182" dst="27" />
  <connection src="27"   dst="117" />
</frame>
```

Every `<connection>` must be a direct child of `<frame>` or a direct child of a
frame-level `<connections>` group. Each `src` / `dst` value must match exactly
one item, AWS group, rectangle, port, or identified child frame by catalog ID,
`id`, `name`, or `ref`. If the same service icon appears multiple times, give
the connected item a unique `name` or `ref` and use that value as the endpoint.

For network diagrams, define structural paths and communication flows
separately:

```xml
<connection src="client" dst="router" kind="route" />
<connection src="client" dst="router" kind="traffic" color="#2563eb" />
```

Routes have no arrowheads. V1 rejects a route whose effective start or end
arrowhead is non-`none`, including values inherited from `<connections>`.
Traffic lines are directional and, when they share the same endpoints as a
route, render beside the route lane. See
[docs/src/examples/samples/route-traffic.xal](../../docs/src/examples/samples/route-traffic.xal) for a compact
route/traffic example.

Excalidraw output uses the same orthogonal routing metadata and adds small
editable anchor grids behind item icons. These anchors keep lines from covering
icons/labels while preserving visible endpoints. When several lines would share
the same X or Y lane, the renderer offsets later lines where possible. Group
header tags are treated as route obstacles so tag labels stay readable.

Reference: [docs/src/examples/samples/sample.xal](../../docs/src/examples/samples/sample.xal)
DSL specification: [07.01 XAL overview](../07-xal-spec/07-01-xal-specification-overview.instructions.md)

---
