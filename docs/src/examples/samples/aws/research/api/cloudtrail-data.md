# AWS CloudTrail Data Service

API version: 2021-08-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudtrail-data/2021-08-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## PutAuditEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `auditEvents` | `List<AuditEvent>` | yes |
| `channelArn` | `string` | yes |
| `externalId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failed` | `List<ResultErrorEntry>` | yes |
| `successful` | `List<AuditEventResultEntry>` | yes |

