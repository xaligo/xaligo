# AWS Lambda Core

API version: 2026-04-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lambda-core/2026-04-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateNetworkConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Configuration` | `NetworkConnectorConfiguration` | yes |
| `OperatorRole` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Id` | `string` | yes |
| `Configuration` | `NetworkConnectorConfiguration` | no |
| `OperatorRole` | `string` | no |
| `State` | `string` | no |

## DeleteNetworkConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Id` | `string` | yes |
| `Configuration` | `NetworkConnectorConfiguration` | no |
| `OperatorRole` | `string` | no |
| `State` | `string` | no |

## GetNetworkConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Id` | `string` | yes |
| `Version` | `long` | no |
| `Configuration` | `NetworkConnectorConfiguration` | no |
| `OperatorRole` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `LastModified` | `timestamp` | no |

## ListNetworkConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkConnectors` | `List<NetworkConnectorSummary>` | yes |
| `NextMarker` | `string` | no |

## UpdateNetworkConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Configuration` | `NetworkConnectorConfiguration` | no |
| `OperatorRole` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Id` | `string` | yes |
| `OperatorRole` | `string` | no |
| `Configuration` | `NetworkConnectorConfiguration` | no |
| `State` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastModified` | `timestamp` | no |

