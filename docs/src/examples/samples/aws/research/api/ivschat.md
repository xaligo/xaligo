# Amazon Interactive Video Service Chat

API version: 2020-07-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ivschat/2020-07-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateChatToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roomIdentifier` | `string` | yes |
| `userId` | `string` | yes |
| `capabilities` | `List<string>` | no |
| `sessionDurationInMinutes` | `integer` | no |
| `attributes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `token` | `string` | no |
| `tokenExpirationTime` | `timestamp` | no |
| `sessionExpirationTime` | `timestamp` | no |

## CreateLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `maximumMessageRatePerSecond` | `integer` | no |
| `maximumMessageLength` | `integer` | no |
| `messageReviewHandler` | `MessageReviewHandler` | no |
| `tags` | `Map<string>` | no |
| `loggingConfigurationIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `maximumMessageRatePerSecond` | `integer` | no |
| `maximumMessageLength` | `integer` | no |
| `messageReviewHandler` | `MessageReviewHandler` | no |
| `tags` | `Map<string>` | no |
| `loggingConfigurationIdentifiers` | `List<string>` | no |

## DeleteLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roomIdentifier` | `string` | yes |
| `id` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## DeleteRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisconnectUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roomIdentifier` | `string` | yes |
| `userId` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |

## GetRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `maximumMessageRatePerSecond` | `integer` | no |
| `maximumMessageLength` | `integer` | no |
| `messageReviewHandler` | `MessageReviewHandler` | no |
| `tags` | `Map<string>` | no |
| `loggingConfigurationIdentifiers` | `List<string>` | no |

## ListLoggingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingConfigurations` | `List<LoggingConfigurationSummary>` | yes |
| `nextToken` | `string` | no |

## ListRooms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `messageReviewHandlerUri` | `string` | no |
| `loggingConfigurationIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rooms` | `List<RoomSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## SendEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roomIdentifier` | `string` | yes |
| `eventName` | `string` | yes |
| `attributes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |

## UpdateRoom

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `maximumMessageRatePerSecond` | `integer` | no |
| `maximumMessageLength` | `integer` | no |
| `messageReviewHandler` | `MessageReviewHandler` | no |
| `loggingConfigurationIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `createTime` | `timestamp` | no |
| `updateTime` | `timestamp` | no |
| `maximumMessageRatePerSecond` | `integer` | no |
| `maximumMessageLength` | `integer` | no |
| `messageReviewHandler` | `MessageReviewHandler` | no |
| `tags` | `Map<string>` | no |
| `loggingConfigurationIdentifiers` | `List<string>` | no |

