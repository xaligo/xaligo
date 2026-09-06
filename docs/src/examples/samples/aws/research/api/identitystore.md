# AWS SSO Identity Store

API version: 2020-06-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/identitystore/2020-06-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `IdentityStoreId` | `string` | yes |

## CreateGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MemberId` | `MemberId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MembershipId` | `string` | yes |
| `IdentityStoreId` | `string` | yes |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserName` | `string` | no |
| `Name` | `Name` | no |
| `DisplayName` | `string` | no |
| `NickName` | `string` | no |
| `ProfileUrl` | `string` | no |
| `Emails` | `List<Email>` | no |
| `Addresses` | `List<Address>` | no |
| `PhoneNumbers` | `List<PhoneNumber>` | no |
| `UserType` | `string` | no |
| `Title` | `string` | no |
| `PreferredLanguage` | `string` | no |
| `Locale` | `string` | no |
| `Timezone` | `string` | no |
| `Photos` | `List<Photo>` | no |
| `Website` | `string` | no |
| `Birthdate` | `string` | no |
| `Roles` | `List<Role>` | no |
| `Extensions` | `Map<AttributeValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MembershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `DisplayName` | `string` | no |
| `ExternalIds` | `List<ExternalId>` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `UpdatedBy` | `string` | no |
| `IdentityStoreId` | `string` | yes |

## DescribeGroupMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MembershipId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MembershipId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MemberId` | `MemberId` | yes |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `UpdatedBy` | `string` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |
| `Extensions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |
| `UserName` | `string` | no |
| `ExternalIds` | `List<ExternalId>` | no |
| `Name` | `Name` | no |
| `DisplayName` | `string` | no |
| `NickName` | `string` | no |
| `ProfileUrl` | `string` | no |
| `Emails` | `List<Email>` | no |
| `Addresses` | `List<Address>` | no |
| `PhoneNumbers` | `List<PhoneNumber>` | no |
| `UserType` | `string` | no |
| `Title` | `string` | no |
| `PreferredLanguage` | `string` | no |
| `Locale` | `string` | no |
| `Timezone` | `string` | no |
| `UserStatus` | `string` | no |
| `Photos` | `List<Photo>` | no |
| `Website` | `string` | no |
| `Birthdate` | `string` | no |
| `Roles` | `List<Role>` | no |
| `CreatedAt` | `timestamp` | no |
| `CreatedBy` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `UpdatedBy` | `string` | no |
| `Extensions` | `Map<AttributeValue>` | no |

## GetGroupId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `AlternateIdentifier` | `AlternateIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `IdentityStoreId` | `string` | yes |

## GetGroupMembershipId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MemberId` | `MemberId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MembershipId` | `string` | yes |
| `IdentityStoreId` | `string` | yes |

## GetUserId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `AlternateIdentifier` | `AlternateIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |

## IsMemberInGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MemberId` | `MemberId` | yes |
| `GroupIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<GroupMembershipExistenceResult>` | yes |

## ListGroupMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupMemberships` | `List<GroupMembership>` | yes |
| `NextToken` | `string` | no |

## ListGroupMembershipsForMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MemberId` | `MemberId` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupMemberships` | `List<GroupMembership>` | yes |
| `NextToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<Group>` | yes |
| `NextToken` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `Extensions` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | yes |
| `NextToken` | `string` | no |

## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `GroupId` | `string` | yes |
| `Operations` | `List<AttributeOperation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityStoreId` | `string` | yes |
| `UserId` | `string` | yes |
| `Operations` | `List<AttributeOperation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


