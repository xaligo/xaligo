# AWS IoT Data Plane

API version: 2015-05-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iot-data/2015-05-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `cleanSession` | `boolean` | no |
| `preventWillMessage` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThingShadow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `shadowName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | yes |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `includeSocketInformation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connected` | `boolean` | no |
| `thingName` | `string` | no |
| `cleanSession` | `boolean` | no |
| `sourceIp` | `string` | no |
| `sourcePort` | `integer` | no |
| `targetIp` | `string` | no |
| `targetPort` | `integer` | no |
| `keepAliveDuration` | `integer` | no |
| `connectedSince` | `long` | no |
| `disconnectedSince` | `long` | no |
| `disconnectReason` | `string` | no |
| `sessionExpiry` | `long` | no |
| `clientId` | `string` | no |
| `vpcEndpointId` | `string` | no |

## GetRetainedMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topic` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topic` | `string` | no |
| `payload` | `blob` | no |
| `qos` | `integer` | no |
| `lastModifiedTime` | `long` | no |
| `userProperties` | `blob` | no |

## GetThingShadow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `shadowName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | no |

## ListNamedShadowsForThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `nextToken` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<string>` | no |
| `nextToken` | `string` | no |
| `timestamp` | `long` | no |

## ListRetainedMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `retainedTopics` | `List<RetainedMessageSummary>` | no |
| `nextToken` | `string` | no |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptions` | `List<SubscriptionSummary>` | no |
| `nextToken` | `string` | no |

## Publish

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topic` | `string` | yes |
| `qos` | `integer` | no |
| `retain` | `boolean` | no |
| `payload` | `blob` | no |
| `userProperties` | `string` | no |
| `payloadFormatIndicator` | `string` | no |
| `contentType` | `string` | no |
| `responseTopic` | `string` | no |
| `correlationData` | `string` | no |
| `messageExpiry` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDirectMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `topic` | `string` | yes |
| `contentType` | `string` | no |
| `responseTopic` | `string` | no |
| `confirmation` | `boolean` | no |
| `timeout` | `integer` | no |
| `payload` | `blob` | no |
| `userProperties` | `string` | no |
| `payloadFormatIndicator` | `string` | no |
| `correlationData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `traceId` | `string` | no |

## UpdateThingShadow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `shadowName` | `string` | no |
| `payload` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | no |

