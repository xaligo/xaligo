# IAM Toolbox (Preview)

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iam-toolbox/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRequestAuthorizationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizationId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestContext` | `Map<Document>` | yes |
| `evaluations` | `List<Evaluation>` | yes |
| `policies` | `List<PolicyInfo>` | yes |
| `nextToken` | `string` | no |

