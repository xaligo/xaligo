# AWS Signer Data Plane

API version: 2017-08-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/signer-data/2017-08-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRevocationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signatureTimestamp` | `timestamp` | yes |
| `platformId` | `string` | yes |
| `profileVersionArn` | `string` | yes |
| `jobArn` | `string` | yes |
| `certificateHashes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revokedEntities` | `List<string>` | no |

