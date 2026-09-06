# AWS Billing and Cost Management Recommended Actions

API version: 2024-11-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bcm-recommended-actions/2024-11-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ListRecommendedActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `RequestFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedActions` | `List<RecommendedAction>` | yes |
| `nextToken` | `string` | no |

