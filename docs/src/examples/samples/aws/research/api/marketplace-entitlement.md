# AWS Marketplace Entitlement Service

API version: 2017-01-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-entitlement/2017-01-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductCode` | `string` | yes |
| `Filter` | `Map<List<string>>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlements` | `List<Entitlement>` | no |
| `NextToken` | `string` | no |

