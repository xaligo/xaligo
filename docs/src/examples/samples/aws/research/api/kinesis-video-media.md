# Amazon Kinesis Video Streams Media

API version: 2017-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesis-video-media/2017-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `StartSelector` | `StartSelector` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Payload` | `blob` | no |

