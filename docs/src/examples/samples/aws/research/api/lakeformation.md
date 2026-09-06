# AWS Lake Formation

API version: 2017-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lakeformation/2017-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddLFTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Resource` | `Resource` | yes |
| `LFTags` | `List<LFTagPair>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<LFTagError>` | no |

## AssumeDecoratedRoleWithSAML

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SAMLAssertion` | `string` | yes |
| `RoleArn` | `string` | yes |
| `PrincipalArn` | `string` | yes |
| `DurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyId` | `string` | no |
| `SecretAccessKey` | `string` | no |
| `SessionToken` | `string` | no |
| `Expiration` | `timestamp` | no |

## BatchGrantPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Entries` | `List<BatchPermissionsRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<BatchPermissionsFailureEntry>` | no |

## BatchRevokePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Entries` | `List<BatchPermissionsRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<BatchPermissionsFailureEntry>` | no |

## CancelTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CommitTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionStatus` | `string` | no |

## CreateDataCellsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableData` | `DataCellsFilter` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLFTag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `TagKey` | `string` | yes |
| `TagValues` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLFTagExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `CatalogId` | `string` | no |
| `Expression` | `List<LFTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLakeFormationIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `InstanceArn` | `string` | no |
| `ExternalFiltering` | `ExternalFilteringConfiguration` | no |
| `ShareRecipients` | `List<DataLakePrincipal>` | no |
| `ServiceIntegrations` | `List<ServiceIntegrationUnion>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |

## CreateLakeFormationOptIn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principal` | `DataLakePrincipal` | yes |
| `Resource` | `Resource` | yes |
| `Condition` | `Condition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataCellsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableCatalogId` | `string` | no |
| `DatabaseName` | `string` | no |
| `TableName` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLFTag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `TagKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLFTagExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLakeFormationIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLakeFormationOptIn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principal` | `DataLakePrincipal` | yes |
| `Resource` | `Resource` | yes |
| `Condition` | `Condition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteObjectsOnCancel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `TransactionId` | `string` | yes |
| `Objects` | `List<VirtualObject>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeLakeFormationIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `InstanceArn` | `string` | no |
| `ApplicationArn` | `string` | no |
| `ExternalFiltering` | `ExternalFilteringConfiguration` | no |
| `ShareRecipients` | `List<DataLakePrincipal>` | no |
| `ServiceIntegrations` | `List<ServiceIntegrationUnion>` | no |
| `ResourceShare` | `string` | no |

## DescribeResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceInfo` | `ResourceInfo` | no |

## DescribeTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionDescription` | `TransactionDescription` | no |

## ExtendTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDataCellsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableCatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataCellsFilter` | `DataCellsFilter` | no |

## GetDataLakePrincipal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | no |

## GetDataLakeSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataLakeSettings` | `DataLakeSettings` | no |

## GetEffectivePermissionsForPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Permissions` | `List<PrincipalResourcePermissions>` | no |
| `NextToken` | `string` | no |

## GetLFTag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `TagKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `TagKey` | `string` | no |
| `TagValues` | `List<string>` | no |

## GetLFTagExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CatalogId` | `string` | no |
| `Expression` | `List<LFTag>` | no |

## GetQueryState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Error` | `string` | no |
| `State` | `string` | yes |

## GetQueryStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionStatistics` | `ExecutionStatistics` | no |
| `PlanningStatistics` | `PlanningStatistics` | no |
| `QuerySubmissionTime` | `timestamp` | no |

## GetResourceLFTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Resource` | `Resource` | yes |
| `ShowAssignedLFTags` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LFTagOnDatabase` | `List<LFTagPair>` | no |
| `LFTagsOnTable` | `List<LFTagPair>` | no |
| `LFTagsOnColumns` | `List<ColumnLFTag>` | no |

## GetTableObjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `TransactionId` | `string` | no |
| `QueryAsOfTime` | `timestamp` | no |
| `PartitionPredicate` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Objects` | `List<PartitionObjects>` | no |
| `NextToken` | `string` | no |

## GetTemporaryDataLocationCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurationSeconds` | `integer` | no |
| `AuditContext` | `AuditContext` | no |
| `DataLocations` | `List<string>` | no |
| `CredentialsScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `TemporaryCredentials` | no |
| `AccessibleDataLocations` | `List<string>` | no |
| `CredentialsScope` | `string` | no |

## GetTemporaryGluePartitionCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableArn` | `string` | yes |
| `Partition` | `PartitionValueList` | yes |
| `Permissions` | `List<string>` | no |
| `DurationSeconds` | `integer` | no |
| `AuditContext` | `AuditContext` | no |
| `SupportedPermissionTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyId` | `string` | no |
| `SecretAccessKey` | `string` | no |
| `SessionToken` | `string` | no |
| `Expiration` | `timestamp` | no |

## GetTemporaryGlueTableCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableArn` | `string` | yes |
| `Permissions` | `List<string>` | no |
| `DurationSeconds` | `integer` | no |
| `AuditContext` | `AuditContext` | no |
| `SupportedPermissionTypes` | `List<string>` | no |
| `S3Path` | `string` | no |
| `QuerySessionContext` | `QuerySessionContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyId` | `string` | no |
| `SecretAccessKey` | `string` | no |
| `SessionToken` | `string` | no |
| `Expiration` | `timestamp` | no |
| `VendedS3Path` | `List<string>` | no |

## GetWorkUnitResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |
| `WorkUnitId` | `long` | yes |
| `WorkUnitToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultStream` | `blob` | no |

## GetWorkUnits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |
| `QueryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `QueryId` | `string` | yes |
| `WorkUnitRanges` | `List<WorkUnitRange>` | yes |

## GrantPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Principal` | `DataLakePrincipal` | yes |
| `Resource` | `Resource` | yes |
| `Permissions` | `List<string>` | yes |
| `Condition` | `Condition` | no |
| `PermissionsWithGrantOption` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListDataCellsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `TableResource` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataCellsFilters` | `List<DataCellsFilter>` | no |
| `NextToken` | `string` | no |

## ListLFTagExpressions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LFTagExpressions` | `List<LFTagExpression>` | no |
| `NextToken` | `string` | no |

## ListLFTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ResourceShareType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LFTags` | `List<LFTagPair>` | no |
| `NextToken` | `string` | no |

## ListLakeFormationOptIns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principal` | `DataLakePrincipal` | no |
| `Resource` | `Resource` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LakeFormationOptInsInfoList` | `List<LakeFormationOptInsInfo>` | no |
| `NextToken` | `string` | no |

## ListPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Principal` | `DataLakePrincipal` | no |
| `ResourceType` | `string` | no |
| `Resource` | `Resource` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeRelated` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrincipalResourcePermissions` | `List<PrincipalResourcePermissions>` | no |
| `NextToken` | `string` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterConditionList` | `List<FilterCondition>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceInfoList` | `List<ResourceInfo>` | no |
| `NextToken` | `string` | no |

## ListTableStorageOptimizers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `StorageOptimizerType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageOptimizerList` | `List<StorageOptimizer>` | no |
| `NextToken` | `string` | no |

## ListTransactions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `StatusFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Transactions` | `List<TransactionDescription>` | no |
| `NextToken` | `string` | no |

## PutDataLakeSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DataLakeSettings` | `DataLakeSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `UseServiceLinkedRole` | `boolean` | no |
| `RoleArn` | `string` | no |
| `WithFederation` | `boolean` | no |
| `HybridAccessEnabled` | `boolean` | no |
| `WithPrivilegedAccess` | `boolean` | no |
| `ExpectedResourceOwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveLFTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Resource` | `Resource` | yes |
| `LFTags` | `List<LFTagPair>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<LFTagError>` | no |

## RevokePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Principal` | `DataLakePrincipal` | yes |
| `Resource` | `Resource` | yes |
| `Permissions` | `List<string>` | yes |
| `Condition` | `Condition` | no |
| `PermissionsWithGrantOption` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchDatabasesByLFTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CatalogId` | `string` | no |
| `Expression` | `List<LFTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DatabaseList` | `List<TaggedDatabase>` | no |

## SearchTablesByLFTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CatalogId` | `string` | no |
| `Expression` | `List<LFTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TableList` | `List<TaggedTable>` | no |

## StartQueryPlanning

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryPlanningContext` | `QueryPlanningContext` | yes |
| `QueryString` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |

## StartTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactionId` | `string` | no |

## UpdateDataCellsFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableData` | `DataCellsFilter` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLFTag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `TagKey` | `string` | yes |
| `TagValuesToDelete` | `List<string>` | no |
| `TagValuesToAdd` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLFTagExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `CatalogId` | `string` | no |
| `Expression` | `List<LFTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLakeFormationIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ShareRecipients` | `List<DataLakePrincipal>` | no |
| `ServiceIntegrations` | `List<ServiceIntegrationUnion>` | no |
| `ApplicationStatus` | `string` | no |
| `ExternalFiltering` | `ExternalFilteringConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `WithFederation` | `boolean` | no |
| `HybridAccessEnabled` | `boolean` | no |
| `ExpectedResourceOwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTableObjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `TransactionId` | `string` | no |
| `WriteOperations` | `List<WriteOperation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTableStorageOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `StorageOptimizerConfig` | `Map<Map<string>>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `string` | no |

