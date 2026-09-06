# FinSpace Public API

API version: 2020-07-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/finspace-data/2020-07-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateUserToPermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |
| `userId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |

## CreateChangeset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetId` | `string` | yes |
| `changeType` | `string` | yes |
| `sourceParams` | `Map<string>` | yes |
| `formatParams` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |
| `changesetId` | `string` | no |

## CreateDataView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetId` | `string` | yes |
| `autoUpdate` | `boolean` | no |
| `sortColumns` | `List<string>` | no |
| `partitionColumns` | `List<string>` | no |
| `asOfTimestamp` | `long` | no |
| `destinationTypeParams` | `DataViewDestinationTypeParams` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |
| `dataViewId` | `string` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetTitle` | `string` | yes |
| `kind` | `string` | yes |
| `datasetDescription` | `string` | no |
| `ownerInfo` | `DatasetOwnerInfo` | no |
| `permissionGroupParams` | `PermissionGroupParams` | yes |
| `alias` | `string` | no |
| `schemaDefinition` | `SchemaUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |

## CreatePermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `applicationPermissions` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `emailAddress` | `string` | yes |
| `type` | `string` | yes |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `apiAccess` | `string` | no |
| `apiAccessPrincipalArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |

## DeletePermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | no |

## DisableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |

## DisassociateUserFromPermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |
| `userId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |

## EnableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |

## GetChangeset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `changesetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changesetId` | `string` | no |
| `changesetArn` | `string` | no |
| `datasetId` | `string` | no |
| `changeType` | `string` | no |
| `sourceParams` | `Map<string>` | no |
| `formatParams` | `Map<string>` | no |
| `createTime` | `long` | no |
| `status` | `string` | no |
| `errorInfo` | `ChangesetErrorInfo` | no |
| `activeUntilTimestamp` | `long` | no |
| `activeFromTimestamp` | `long` | no |
| `updatesChangesetId` | `string` | no |
| `updatedByChangesetId` | `string` | no |

## GetDataView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataViewId` | `string` | yes |
| `datasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoUpdate` | `boolean` | no |
| `partitionColumns` | `List<string>` | no |
| `datasetId` | `string` | no |
| `asOfTimestamp` | `long` | no |
| `errorInfo` | `DataViewErrorInfo` | no |
| `lastModifiedTime` | `long` | no |
| `createTime` | `long` | no |
| `sortColumns` | `List<string>` | no |
| `dataViewId` | `string` | no |
| `dataViewArn` | `string` | no |
| `destinationTypeParams` | `DataViewDestinationTypeParams` | no |
| `status` | `string` | no |

## GetDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |
| `datasetArn` | `string` | no |
| `datasetTitle` | `string` | no |
| `kind` | `string` | no |
| `datasetDescription` | `string` | no |
| `createTime` | `long` | no |
| `lastModifiedTime` | `long` | no |
| `schemaDefinition` | `SchemaUnion` | no |
| `alias` | `string` | no |
| `status` | `string` | no |

## GetExternalDataViewAccessDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataViewId` | `string` | yes |
| `datasetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | no |
| `s3Location` | `S3Location` | no |

## GetPermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroup` | `PermissionGroup` | no |

## GetProgrammaticAccessCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `durationInMinutes` | `long` | no |
| `environmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `Credentials` | no |
| `durationInMinutes` | `long` | no |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `status` | `string` | no |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `emailAddress` | `string` | no |
| `type` | `string` | no |
| `apiAccess` | `string` | no |
| `apiAccessPrincipalArn` | `string` | no |
| `createTime` | `long` | no |
| `lastEnabledTime` | `long` | no |
| `lastDisabledTime` | `long` | no |
| `lastModifiedTime` | `long` | no |
| `lastLoginTime` | `long` | no |

## GetWorkingLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `locationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3Uri` | `string` | no |
| `s3Path` | `string` | no |
| `s3Bucket` | `string` | no |

## ListChangesets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changesets` | `List<ChangesetSummary>` | no |
| `nextToken` | `string` | no |

## ListDataViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dataViews` | `List<DataViewSummary>` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasets` | `List<Dataset>` | no |
| `nextToken` | `string` | no |

## ListPermissionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroups` | `List<PermissionGroup>` | no |
| `nextToken` | `string` | no |

## ListPermissionGroupsByUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroups` | `List<PermissionGroupByUser>` | no |
| `nextToken` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `users` | `List<User>` | no |
| `nextToken` | `string` | no |

## ListUsersByPermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `users` | `List<UserByPermissionGroup>` | no |
| `nextToken` | `string` | no |

## ResetUserPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `temporaryPassword` | `string` | no |

## UpdateChangeset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetId` | `string` | yes |
| `changesetId` | `string` | yes |
| `sourceParams` | `Map<string>` | yes |
| `formatParams` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changesetId` | `string` | no |
| `datasetId` | `string` | no |

## UpdateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetId` | `string` | yes |
| `datasetTitle` | `string` | yes |
| `kind` | `string` | yes |
| `datasetDescription` | `string` | no |
| `alias` | `string` | no |
| `schemaDefinition` | `SchemaUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | no |

## UpdatePermissionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `applicationPermissions` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionGroupId` | `string` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `type` | `string` | no |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `apiAccess` | `string` | no |
| `apiAccessPrincipalArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |

