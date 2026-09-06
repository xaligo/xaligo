# Amazon Kinesis Video Signaling Channels

API version: 2019-12-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesis-video-signaling/2019-12-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetIceServerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `ClientId` | `string` | no |
| `Service` | `string` | no |
| `Username` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IceServerList` | `List<IceServer>` | no |

## SendAlexaOfferToMaster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `SenderClientId` | `string` | yes |
| `MessagePayload` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Answer` | `string` | no |

