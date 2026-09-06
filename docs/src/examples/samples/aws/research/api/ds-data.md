# AWS Directory Service Data

API version: 2023-05-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ds-data/2023-05-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddGroupMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `GroupName` | `string` | yes |
| `MemberName` | `string` | yes |
| `MemberRealm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `GroupScope` | `string` | no |
| `GroupType` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `SAMAccountName` | `string` | no |
| `SID` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `EmailAddress` | `string` | no |
| `GivenName` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `SAMAccountName` | `string` | yes |
| `Surname` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `SAMAccountName` | `string` | no |
| `SID` | `string` | no |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `OtherAttributes` | `List<string>` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `DistinguishedName` | `string` | no |
| `GroupScope` | `string` | no |
| `GroupType` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | no |
| `SID` | `string` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `OtherAttributes` | `List<string>` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `DistinguishedName` | `string` | no |
| `EmailAddress` | `string` | no |
| `Enabled` | `boolean` | no |
| `GivenName` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | no |
| `SID` | `string` | no |
| `Surname` | `string` | no |
| `UserPrincipalName` | `string` | no |

## DisableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListGroupMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `MemberRealm` | `string` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `MemberRealm` | `string` | no |
| `Members` | `List<Member>` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `Groups` | `List<GroupSummary>` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

## ListGroupsForMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `MemberRealm` | `string` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `SAMAccountName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `Groups` | `List<GroupSummary>` | no |
| `MemberRealm` | `string` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `Users` | `List<UserSummary>` | no |

## RemoveGroupMember

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `GroupName` | `string` | yes |
| `MemberName` | `string` | yes |
| `MemberRealm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `SearchAttributes` | `List<string>` | yes |
| `SearchString` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `Groups` | `List<Group>` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |

## SearchUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `SearchAttributes` | `List<string>` | yes |
| `SearchString` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `NextToken` | `string` | no |
| `Realm` | `string` | no |
| `Users` | `List<User>` | no |

## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `GroupScope` | `string` | no |
| `GroupType` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `SAMAccountName` | `string` | yes |
| `UpdateType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `EmailAddress` | `string` | no |
| `GivenName` | `string` | no |
| `OtherAttributes` | `Map<AttributeValue>` | no |
| `SAMAccountName` | `string` | yes |
| `Surname` | `string` | no |
| `UpdateType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


