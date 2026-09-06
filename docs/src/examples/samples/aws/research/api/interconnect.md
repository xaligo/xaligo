# Interconnect

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/interconnect/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptConnectionProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachPoint` | `AttachPoint` | yes |
| `activationKey` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connection` | `Connection` | no |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `bandwidth` | `string` | yes |
| `attachPoint` | `AttachPoint` | yes |
| `environmentId` | `string` | yes |
| `remoteAccount` | `RemoteAccountIdentifier` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connection` | `Connection` | no |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connection` | `Connection` | yes |

## DescribeConnectionProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activationKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bandwidth` | `string` | yes |
| `environmentId` | `string` | yes |
| `provider` | `Provider` | yes |
| `location` | `string` | yes |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connection` | `Connection` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | yes |

## ListAttachPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachPoints` | `List<AttachPointDescriptor>` | yes |
| `nextToken` | `string` | no |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |
| `environmentId` | `string` | no |
| `provider` | `Provider` | no |
| `attachPoint` | `AttachPoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connections` | `List<ConnectionSummary>` | no |
| `nextToken` | `string` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `provider` | `Provider` | no |
| `location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<Environment>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `bandwidth` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connection` | `Connection` | no |

