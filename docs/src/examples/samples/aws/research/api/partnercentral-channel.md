# Partner Central Channel API

API version: 2024-03-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/partnercentral-channel/2024-03-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptChannelHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelHandshakeDetail` | `AcceptChannelHandshakeDetail` | no |

## CancelChannelHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelHandshakeDetail` | `CancelChannelHandshakeDetail` | no |

## CreateChannelHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `handshakeType` | `string` | yes |
| `catalog` | `string` | yes |
| `associatedResourceIdentifier` | `string` | yes |
| `payload` | `ChannelHandshakePayload` | no |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelHandshakeDetail` | `CreateChannelHandshakeDetail` | no |

## CreateProgramManagementAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `program` | `string` | yes |
| `displayName` | `string` | yes |
| `accountId` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `programManagementAccountDetail` | `CreateProgramManagementAccountDetail` | no |

## CreateRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `associationType` | `string` | yes |
| `programManagementAccountIdentifier` | `string` | yes |
| `associatedAccountId` | `string` | yes |
| `displayName` | `string` | yes |
| `resaleAccountModel` | `string` | no |
| `sector` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |
| `requestedSupportPlan` | `SupportPlan` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationshipDetail` | `CreateRelationshipDetail` | no |

## DeleteProgramManagementAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |
| `programManagementAccountIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `programManagementAccountIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationshipDetail` | `RelationshipDetail` | no |

## ListChannelHandshakes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `handshakeType` | `string` | yes |
| `catalog` | `string` | yes |
| `participantType` | `string` | yes |
| `maxResults` | `integer` | no |
| `statuses` | `List<string>` | no |
| `associatedResourceIdentifiers` | `List<string>` | no |
| `handshakeTypeFilters` | `ListChannelHandshakesTypeFilters` | no |
| `handshakeTypeSort` | `ListChannelHandshakesTypeSort` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ChannelHandshakeSummary>` | no |
| `nextToken` | `string` | no |

## ListProgramManagementAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `maxResults` | `integer` | no |
| `displayNames` | `List<string>` | no |
| `programs` | `List<string>` | no |
| `accountIds` | `List<string>` | no |
| `statuses` | `List<string>` | no |
| `sort` | `ListProgramManagementAccountsSortBase` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ProgramManagementAccountSummary>` | no |
| `nextToken` | `string` | no |

## ListRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `maxResults` | `integer` | no |
| `associatedAccountIds` | `List<string>` | no |
| `associationTypes` | `List<string>` | no |
| `displayNames` | `List<string>` | no |
| `programManagementAccountIdentifiers` | `List<string>` | no |
| `sort` | `ListRelationshipsSortBase` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RelationshipSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## RejectChannelHandshake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelHandshakeDetail` | `RejectChannelHandshakeDetail` | no |

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


## UpdateProgramManagementAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |
| `displayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `programManagementAccountDetail` | `UpdateProgramManagementAccountDetail` | no |

## UpdateRelationship

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalog` | `string` | yes |
| `identifier` | `string` | yes |
| `programManagementAccountIdentifier` | `string` | yes |
| `revision` | `string` | no |
| `displayName` | `string` | no |
| `requestedSupportPlan` | `SupportPlan` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relationshipDetail` | `UpdateRelationshipDetail` | no |

