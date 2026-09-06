# AWS Marketplace Reporting Service

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-reporting/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetBuyerDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboardIdentifier` | `string` | yes |
| `embeddingDomains` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `embedUrl` | `string` | yes |
| `dashboardIdentifier` | `string` | yes |
| `embeddingDomains` | `List<string>` | yes |

