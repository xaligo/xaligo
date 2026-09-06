# AWS re:Post Private

API version: 2022-05-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/repostspace/2022-05-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchAddChannelRoleToAccessors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelId` | `string` | yes |
| `accessorIds` | `List<string>` | yes |
| `channelRole` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addedAccessorIds` | `List<string>` | yes |
| `errors` | `List<BatchError>` | yes |

## BatchAddRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `accessorIds` | `List<string>` | yes |
| `role` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addedAccessorIds` | `List<string>` | yes |
| `errors` | `List<BatchError>` | yes |

## BatchRemoveChannelRoleFromAccessors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelId` | `string` | yes |
| `accessorIds` | `List<string>` | yes |
| `channelRole` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `removedAccessorIds` | `List<string>` | yes |
| `errors` | `List<BatchError>` | yes |

## BatchRemoveRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `accessorIds` | `List<string>` | yes |
| `role` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `removedAccessorIds` | `List<string>` | yes |
| `errors` | `List<BatchError>` | yes |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelName` | `string` | yes |
| `channelDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | yes |

## CreateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `subdomain` | `string` | yes |
| `tier` | `string` | yes |
| `description` | `string` | no |
| `userKMSKey` | `string` | no |
| `tags` | `Map<string>` | no |
| `roleArn` | `string` | no |
| `supportedEmailDomains` | `SupportedEmailDomainsParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |

## DeleteSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `adminId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelId` | `string` | yes |
| `channelName` | `string` | yes |
| `channelDescription` | `string` | no |
| `createDateTime` | `timestamp` | yes |
| `deleteDateTime` | `timestamp` | no |
| `channelRoles` | `Map<List<string>>` | no |
| `channelStatus` | `string` | yes |

## GetSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `configurationStatus` | `string` | yes |
| `clientId` | `string` | yes |
| `identityStoreId` | `string` | no |
| `applicationArn` | `string` | no |
| `description` | `string` | no |
| `vanityDomainStatus` | `string` | yes |
| `vanityDomain` | `string` | yes |
| `randomDomain` | `string` | yes |
| `customerRoleArn` | `string` | no |
| `createDateTime` | `timestamp` | yes |
| `deleteDateTime` | `timestamp` | no |
| `tier` | `string` | yes |
| `storageLimit` | `long` | yes |
| `userAdmins` | `List<string>` | no |
| `groupAdmins` | `List<string>` | no |
| `roles` | `Map<List<string>>` | no |
| `userKMSKey` | `string` | no |
| `userCount` | `integer` | no |
| `contentSize` | `long` | no |
| `supportedEmailDomains` | `SupportedEmailDomainsStatus` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channels` | `List<ChannelData>` | yes |
| `nextToken` | `string` | no |

## ListSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaces` | `List<SpaceData>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RegisterAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `adminId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendInvites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `accessorIds` | `List<string>` | yes |
| `title` | `string` | yes |
| `body` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `channelId` | `string` | yes |
| `channelName` | `string` | yes |
| `channelDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceId` | `string` | yes |
| `description` | `string` | no |
| `tier` | `string` | no |
| `roleArn` | `string` | no |
| `supportedEmailDomains` | `SupportedEmailDomainsParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


