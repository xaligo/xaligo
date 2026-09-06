# Amazon Connect Contact Lens

API version: 2020-08-21. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connect-contact-lens/2020-08-21/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ListRealtimeContactAnalysisSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Segments` | `List<RealtimeContactAnalysisSegment>` | yes |
| `NextToken` | `string` | no |

