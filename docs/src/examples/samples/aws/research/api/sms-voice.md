# Amazon Pinpoint SMS and Voice Service

API version: 2018-09-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sms-voice/2018-09-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestination` | `EventDestinationDefinition` | no |
| `EventDestinationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConfigurationSetEventDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDestinations` | `List<EventDestination>` | no |

## ListConfigurationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSets` | `List<string>` | no |
| `NextToken` | `string` | no |

## SendVoiceMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallerId` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `Content` | `VoiceMessageContent` | no |
| `DestinationPhoneNumber` | `string` | no |
| `OriginationPhoneNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## UpdateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestination` | `EventDestinationDefinition` | no |
| `EventDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


