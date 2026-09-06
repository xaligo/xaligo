# AWS IoT Secure Tunneling

API version: 2018-10-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotsecuretunneling/2018-10-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CloseTunnel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelId` | `string` | yes |
| `delete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeTunnel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnel` | `Tunnel` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListTunnels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelSummaries` | `List<TunnelSummary>` | no |
| `nextToken` | `string` | no |

## OpenTunnel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `destinationConfig` | `DestinationConfig` | no |
| `timeoutConfig` | `TimeoutConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelId` | `string` | no |
| `tunnelArn` | `string` | no |
| `sourceAccessToken` | `string` | no |
| `destinationAccessToken` | `string` | no |

## RotateTunnelAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelId` | `string` | yes |
| `clientMode` | `string` | yes |
| `destinationConfig` | `DestinationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tunnelArn` | `string` | no |
| `sourceAccessToken` | `string` | no |
| `destinationAccessToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


