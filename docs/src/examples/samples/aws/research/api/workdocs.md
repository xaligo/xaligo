# Amazon WorkDocs

API version: 2016-05-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workdocs/2016-05-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AbortDocumentVersionUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ActivateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `AuthenticationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## AddResourcePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |
| `Principals` | `List<SharePrincipal>` | yes |
| `NotificationOptions` | `NotificationOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShareResults` | `List<ShareResult>` | no |

## CreateComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `ParentId` | `string` | no |
| `ThreadId` | `string` | no |
| `Text` | `string` | yes |
| `Visibility` | `string` | no |
| `NotifyCollaborators` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Comment` | `Comment` | no |

## CreateCustomMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |
| `VersionId` | `string` | no |
| `CustomMetadata` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `Name` | `string` | no |
| `ParentFolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `FolderMetadata` | no |

## CreateLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Labels` | `List<string>` | yes |
| `AuthenticationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateNotificationSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Endpoint` | `string` | yes |
| `Protocol` | `string` | yes |
| `SubscriptionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscription` | `Subscription` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | no |
| `Username` | `string` | yes |
| `EmailAddress` | `string` | no |
| `GivenName` | `string` | yes |
| `Surname` | `string` | yes |
| `Password` | `string` | yes |
| `TimeZoneId` | `string` | no |
| `StorageRule` | `StorageRuleType` | no |
| `AuthenticationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## DeactivateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `AuthenticationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteComment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `CommentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |
| `VersionId` | `string` | no |
| `Keys` | `List<string>` | no |
| `DeleteAll` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDocumentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `DeletePriorVersions` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFolderContents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `AuthenticationToken` | `string` | no |
| `Labels` | `List<string>` | no |
| `DeleteAll` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotificationSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionId` | `string` | yes |
| `OrganizationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `OrganizationId` | `string` | no |
| `ActivityTypes` | `string` | no |
| `ResourceId` | `string` | no |
| `UserId` | `string` | no |
| `IncludeIndirectActivities` | `boolean` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserActivities` | `List<Activity>` | no |
| `Marker` | `string` | no |

## DescribeComments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Comments` | `List<Comment>` | no |
| `Marker` | `string` | no |

## DescribeDocumentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |
| `Include` | `string` | no |
| `Fields` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentVersions` | `List<DocumentVersionMetadata>` | no |
| `Marker` | `string` | no |

## DescribeFolderContents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |
| `Sort` | `string` | no |
| `Order` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |
| `Type` | `string` | no |
| `Include` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Folders` | `List<FolderMetadata>` | no |
| `Documents` | `List<DocumentMetadata>` | no |
| `Marker` | `string` | no |

## DescribeGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `SearchQuery` | `string` | yes |
| `OrganizationId` | `string` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupMetadata>` | no |
| `Marker` | `string` | no |

## DescribeNotificationSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationId` | `string` | yes |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscriptions` | `List<Subscription>` | no |
| `Marker` | `string` | no |

## DescribeResourcePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |
| `PrincipalId` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principals` | `List<Principal>` | no |
| `Marker` | `string` | no |

## DescribeRootFolders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | yes |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Folders` | `List<FolderMetadata>` | no |
| `Marker` | `string` | no |

## DescribeUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `OrganizationId` | `string` | no |
| `UserIds` | `string` | no |
| `Query` | `string` | no |
| `Include` | `string` | no |
| `Order` | `string` | no |
| `Sort` | `string` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |
| `Fields` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `TotalNumberOfUsers` | `long` | no |
| `Marker` | `string` | no |

## GetCurrentUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## GetDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `IncludeCustomMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `DocumentMetadata` | no |
| `CustomMetadata` | `Map<string>` | no |

## GetDocumentPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `Limit` | `integer` | no |
| `Fields` | `string` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `ResourcePath` | no |

## GetDocumentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `Fields` | `string` | no |
| `IncludeCustomMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `DocumentVersionMetadata` | no |
| `CustomMetadata` | `Map<string>` | no |

## GetFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |
| `IncludeCustomMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `FolderMetadata` | no |
| `CustomMetadata` | `Map<string>` | no |

## GetFolderPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |
| `Limit` | `integer` | no |
| `Fields` | `string` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `ResourcePath` | no |

## GetResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `UserId` | `string` | no |
| `CollectionType` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Folders` | `List<FolderMetadata>` | no |
| `Documents` | `List<DocumentMetadata>` | no |
| `Marker` | `string` | no |

## InitiateDocumentVersionUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `ContentCreatedTimestamp` | `timestamp` | no |
| `ContentModifiedTimestamp` | `timestamp` | no |
| `ContentType` | `string` | no |
| `DocumentSizeInBytes` | `long` | no |
| `ParentFolderId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `DocumentMetadata` | no |
| `UploadMetadata` | `UploadMetadata` | no |

## RemoveAllResourcePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveResourcePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `ResourceId` | `string` | yes |
| `PrincipalId` | `string` | yes |
| `PrincipalType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestoreDocumentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `QueryText` | `string` | no |
| `QueryScopes` | `List<string>` | no |
| `OrganizationId` | `string` | no |
| `AdditionalResponseFields` | `List<string>` | no |
| `Filters` | `Filters` | no |
| `OrderBy` | `List<SearchSortResult>` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ResponseItem>` | no |
| `Marker` | `string` | no |

## UpdateDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `Name` | `string` | no |
| `ParentFolderId` | `string` | no |
| `ResourceState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDocumentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `DocumentId` | `string` | yes |
| `VersionId` | `string` | yes |
| `VersionStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFolder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `FolderId` | `string` | yes |
| `Name` | `string` | no |
| `ParentFolderId` | `string` | no |
| `ResourceState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationToken` | `string` | no |
| `UserId` | `string` | yes |
| `GivenName` | `string` | no |
| `Surname` | `string` | no |
| `Type` | `string` | no |
| `StorageRule` | `StorageRuleType` | no |
| `TimeZoneId` | `string` | no |
| `Locale` | `string` | no |
| `GrantPoweruserPrivileges` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

