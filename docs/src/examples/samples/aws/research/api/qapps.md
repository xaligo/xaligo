# QApps

API version: 2023-11-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/qapps/2023-11-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateLibraryItemReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateQAppWithUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchCreateCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `categories` | `List<BatchCreateCategoryInputCategory>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDeleteCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `categories` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchUpdateCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `categories` | `List<CategoryInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLibraryItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |
| `appVersion` | `integer` | yes |
| `categories` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryItemId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `ratingCount` | `integer` | yes |
| `isVerified` | `boolean` | no |

## CreatePresignedUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `cardId` | `string` | yes |
| `appId` | `string` | yes |
| `fileContentsSha256` | `string` | yes |
| `fileName` | `string` | yes |
| `scope` | `string` | yes |
| `sessionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileId` | `string` | yes |
| `presignedUrl` | `string` | yes |
| `presignedUrlFields` | `Map<string>` | yes |
| `presignedUrlExpiration` | `timestamp` | yes |

## CreateQApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |
| `appDefinition` | `AppDefinitionInput` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `appArn` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |
| `initialPrompt` | `string` | no |
| `appVersion` | `integer` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `updatedBy` | `string` | yes |
| `requiredCapabilities` | `List<string>` | no |

## DeleteLibraryItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeQAppPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `appId` | `string` | no |
| `permissions` | `List<PermissionOutput>` | no |

## DisassociateLibraryItemReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateQAppFromUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportQAppSessionData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `csvFileLink` | `string` | yes |
| `expiresAt` | `timestamp` | yes |
| `sessionArn` | `string` | yes |

## GetLibraryItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |
| `appId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryItemId` | `string` | yes |
| `appId` | `string` | yes |
| `appVersion` | `integer` | yes |
| `categories` | `List<Category>` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `ratingCount` | `integer` | yes |
| `isRatedByUser` | `boolean` | no |
| `userCount` | `integer` | no |
| `isVerified` | `boolean` | no |

## GetQApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |
| `appVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `appArn` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |
| `initialPrompt` | `string` | no |
| `appVersion` | `integer` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `updatedBy` | `string` | yes |
| `requiredCapabilities` | `List<string>` | no |
| `appDefinition` | `AppDefinition` | yes |

## GetQAppSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |
| `sessionName` | `string` | no |
| `appVersion` | `integer` | no |
| `latestPublishedAppVersion` | `integer` | no |
| `status` | `string` | yes |
| `cardStatus` | `Map<CardStatus>` | yes |
| `userIsHost` | `boolean` | no |

## GetQAppSessionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |
| `sessionName` | `string` | no |
| `sharingConfiguration` | `SessionSharingConfiguration` | yes |
| `sessionOwner` | `boolean` | no |

## ImportDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `cardId` | `string` | yes |
| `appId` | `string` | yes |
| `fileContentsBase64` | `string` | yes |
| `fileName` | `string` | yes |
| `scope` | `string` | yes |
| `sessionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileId` | `string` | no |

## ListCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `categories` | `List<Category>` | no |

## ListLibraryItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |
| `categoryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryItems` | `List<LibraryItemMember>` | no |
| `nextToken` | `string` | no |

## ListQAppSessionData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |
| `sessionData` | `List<QAppSessionData>` | no |
| `nextToken` | `string` | no |

## ListQApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apps` | `List<UserAppItem>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PredictQApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `options` | `PredictQAppInputOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `PredictAppDefinition` | yes |
| `problemStatement` | `string` | yes |

## StartQAppSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |
| `appVersion` | `integer` | yes |
| `initialValues` | `List<CardValue>` | no |
| `sessionId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |

## StopQAppSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLibraryItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |
| `status` | `string` | no |
| `categories` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `libraryItemId` | `string` | yes |
| `appId` | `string` | yes |
| `appVersion` | `integer` | yes |
| `categories` | `List<Category>` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `ratingCount` | `integer` | yes |
| `isRatedByUser` | `boolean` | no |
| `userCount` | `integer` | no |
| `isVerified` | `boolean` | no |

## UpdateLibraryItemMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `libraryItemId` | `string` | yes |
| `isVerified` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |
| `title` | `string` | no |
| `description` | `string` | no |
| `appDefinition` | `AppDefinitionInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `appArn` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |
| `initialPrompt` | `string` | no |
| `appVersion` | `integer` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `updatedBy` | `string` | yes |
| `requiredCapabilities` | `List<string>` | no |

## UpdateQAppPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `appId` | `string` | yes |
| `grantPermissions` | `List<PermissionInput>` | no |
| `revokePermissions` | `List<PermissionInput>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `appId` | `string` | no |
| `permissions` | `List<PermissionOutput>` | no |

## UpdateQAppSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |
| `values` | `List<CardValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |

## UpdateQAppSessionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionName` | `string` | no |
| `sharingConfiguration` | `SessionSharingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `sessionArn` | `string` | yes |
| `sessionName` | `string` | no |
| `sharingConfiguration` | `SessionSharingConfiguration` | yes |

