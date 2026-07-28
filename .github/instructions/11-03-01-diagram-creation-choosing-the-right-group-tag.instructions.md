---
applyTo: ".github/instructions/manual/**"
---

# 11.03.01 Diagram creation: Choosing the right group tag

### Choosing the right group tag

Use AWS-specific group tags only when the content matches the tag's meaning.
For logical groupings that do not correspond to a specific AWS construct, use `<generic-group>`.

| Tag | When to use |
|---|---|
| `<public-subnet>` | Items that belong to a public (internet-routable) subnet |
| `<private-subnet>` | Items that belong to a private subnet |
| `<security-group>` | Resources sharing an EC2 security group |
| `<auto-scaling-group>` | An EC2 Auto Scaling group |
| `<generic-group>` | Any logical grouping that does not fit the above (security services, storage tiers, CI/CD, etc.) |
| `<capture>` | A border-only structural annotation group (e.g. highlighting a "hot path") that participates in normal nested layout without conveying AWS/architectural semantics |

> **Incorrect:** using `<public-subnet title="Security &amp; Identity">` for IAM / WAF — these are not subnet resources.
> **Correct:** use `<generic-group title="Security &amp; Identity">` instead.
