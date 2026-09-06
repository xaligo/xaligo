# AmazonApiGatewayManagementApi

API version: 2018-11-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/apigatewaymanagementapi/2018-11-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectedAt` | `timestamp` | no |
| `Identity` | `Identity` | no |
| `LastActiveAt` | `timestamp` | no |

## PostToConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Data` | `blob` | yes |
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


