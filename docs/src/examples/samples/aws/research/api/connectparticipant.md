# Amazon Connect Participant Service

API version: 2018-09-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connectparticipant/2018-09-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelParticipantAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CompleteAttachmentUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentIds` | `List<string>` | yes |
| `ClientToken` | `string` | yes |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateParticipantConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `List<string>` | no |
| `ParticipantToken` | `string` | yes |
| `ConnectParticipant` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Websocket` | `Websocket` | no |
| `ConnectionCredentials` | `ConnectionCredentials` | no |
| `WebRTCConnection` | `WebRTCConnection` | no |

## DescribeView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewToken` | `string` | yes |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## DisconnectParticipant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |
| `ConnectionToken` | `string` | yes |
| `UrlExpiryInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |
| `UrlExpiry` | `string` | no |
| `AttachmentSizeInBytes` | `long` | yes |

## GetAuthenticationUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `RedirectUri` | `string` | yes |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationUrl` | `string` | no |

## GetTranscript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ScanDirection` | `string` | no |
| `SortOrder` | `string` | no |
| `StartPosition` | `StartPosition` | no |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InitialContactId` | `string` | no |
| `Transcript` | `List<Item>` | no |
| `NextToken` | `string` | no |

## SendEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | yes |
| `Content` | `string` | no |
| `ClientToken` | `string` | no |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `AbsoluteTime` | `string` | no |

## SendMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | yes |
| `Content` | `string` | yes |
| `ClientToken` | `string` | no |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `AbsoluteTime` | `string` | no |
| `MessageMetadata` | `MessageProcessingMetadata` | no |

## StartAttachmentUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | yes |
| `AttachmentSizeInBytes` | `long` | yes |
| `AttachmentName` | `string` | yes |
| `ClientToken` | `string` | yes |
| `ConnectionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | no |
| `UploadMetadata` | `UploadMetadata` | no |

