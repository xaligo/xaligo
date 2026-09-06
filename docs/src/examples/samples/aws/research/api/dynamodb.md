# Amazon DynamoDB

API version: 2012-08-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dynamodb/2012-08-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statements` | `List<BatchStatementRequest>` | yes |
| `ReturnConsumedCapacity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Responses` | `List<BatchStatementResponse>` | no |
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |

## BatchGetItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestItems` | `Map<KeysAndAttributes>` | yes |
| `ReturnConsumedCapacity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Responses` | `Map<List<Map<AttributeValue>>>` | no |
| `UnprocessedKeys` | `Map<KeysAndAttributes>` | no |
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |

## BatchWriteItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestItems` | `Map<List<WriteRequest>>` | yes |
| `ReturnConsumedCapacity` | `string` | no |
| `ReturnItemCollectionMetrics` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedItems` | `Map<List<WriteRequest>>` | no |
| `ItemCollectionMetrics` | `Map<List<ItemCollectionMetrics>>` | no |
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |

## CreateBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `BackupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupDetails` | `BackupDetails` | no |

## CreateGlobalTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | yes |
| `ReplicationGroup` | `List<Replica>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableDescription` | `GlobalTableDescription` | no |

## CreateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttributeDefinitions` | `List<AttributeDefinition>` | no |
| `TableName` | `string` | yes |
| `KeySchema` | `List<KeySchemaElement>` | no |
| `LocalSecondaryIndexes` | `List<LocalSecondaryIndex>` | no |
| `GlobalSecondaryIndexes` | `List<GlobalSecondaryIndex>` | no |
| `BillingMode` | `string` | no |
| `ProvisionedThroughput` | `ProvisionedThroughput` | no |
| `StreamSpecification` | `StreamSpecification` | no |
| `SSESpecification` | `SSESpecification` | no |
| `Tags` | `List<Tag>` | no |
| `TableClass` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `WarmThroughput` | `WarmThroughput` | no |
| `ResourcePolicy` | `string` | no |
| `OnDemandThroughput` | `OnDemandThroughput` | no |
| `GlobalTableSourceArn` | `string` | no |
| `GlobalTableSettingsReplicationMode` | `string` | no |
| `VectorIndexes` | `List<VectorIndex>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableDescription` | `TableDescription` | no |

## DeleteBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupDescription` | `BackupDescription` | no |

## DeleteItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `Key` | `Map<AttributeValue>` | yes |
| `Expected` | `Map<ExpectedAttributeValue>` | no |
| `ConditionalOperator` | `string` | no |
| `ReturnValues` | `string` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `ReturnItemCollectionMetrics` | `string` | no |
| `ConditionExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |
| `ReturnValuesOnConditionCheckFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |
| `ItemCollectionMetrics` | `ItemCollectionMetrics` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ExpectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevisionId` | `string` | no |

## DeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableDescription` | `TableDescription` | no |

## DescribeBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupDescription` | `BackupDescription` | no |

## DescribeContinuousBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousBackupsDescription` | `ContinuousBackupsDescription` | no |

## DescribeContributorInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `IndexName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `IndexName` | `string` | no |
| `ContributorInsightsRuleList` | `List<string>` | no |
| `ContributorInsightsStatus` | `string` | no |
| `LastUpdateDateTime` | `timestamp` | no |
| `FailureException` | `FailureException` | no |
| `ContributorInsightsMode` | `string` | no |

## DescribeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | yes |

## DescribeExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportDescription` | `ExportDescription` | no |

## DescribeGlobalTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableDescription` | `GlobalTableDescription` | no |

## DescribeGlobalTableSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | no |
| `ReplicaSettings` | `List<ReplicaSettingsDescription>` | no |

## DescribeImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportTableDescription` | `ImportTableDescription` | yes |

## DescribeKinesisStreamingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `KinesisDataStreamDestinations` | `List<KinesisDataStreamDestination>` | no |

## DescribeLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountMaxReadCapacityUnits` | `long` | no |
| `AccountMaxWriteCapacityUnits` | `long` | no |
| `TableMaxReadCapacityUnits` | `long` | no |
| `TableMaxWriteCapacityUnits` | `long` | no |

## DescribeTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `TableDescription` | no |

## DescribeTableReplicaAutoScaling

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableAutoScalingDescription` | `TableAutoScalingDescription` | no |

## DescribeTimeToLive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeToLiveDescription` | `TimeToLiveDescription` | no |

## DisableKinesisStreamingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `StreamArn` | `string` | yes |
| `EnableKinesisStreamingConfiguration` | `EnableKinesisStreamingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `StreamArn` | `string` | no |
| `DestinationStatus` | `string` | no |
| `EnableKinesisStreamingConfiguration` | `EnableKinesisStreamingConfiguration` | no |

## EnableKinesisStreamingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `StreamArn` | `string` | yes |
| `EnableKinesisStreamingConfiguration` | `EnableKinesisStreamingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `StreamArn` | `string` | no |
| `DestinationStatus` | `string` | no |
| `EnableKinesisStreamingConfiguration` | `EnableKinesisStreamingConfiguration` | no |

## ExecuteStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statement` | `string` | yes |
| `Parameters` | `List<AttributeValue>` | no |
| `ConsistentRead` | `boolean` | no |
| `NextToken` | `string` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `Limit` | `integer` | no |
| `ReturnValuesOnConditionCheckFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Map<AttributeValue>>` | no |
| `NextToken` | `string` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |
| `LastEvaluatedKey` | `Map<AttributeValue>` | no |

## ExecuteTransaction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactStatements` | `List<ParameterizedStatement>` | yes |
| `ClientRequestToken` | `string` | no |
| `ReturnConsumedCapacity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Responses` | `List<ItemResponse>` | no |
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |

## ExportTableToPointInTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableArn` | `string` | yes |
| `ExportTime` | `timestamp` | no |
| `ClientToken` | `string` | no |
| `S3Bucket` | `string` | yes |
| `S3BucketOwner` | `string` | no |
| `S3Prefix` | `string` | no |
| `S3SseAlgorithm` | `string` | no |
| `S3SseKmsKeyId` | `string` | no |
| `ExportFormat` | `string` | no |
| `ExportType` | `string` | no |
| `IncrementalExportSpecification` | `IncrementalExportSpecification` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportDescription` | `ExportDescription` | no |

## GetItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `Key` | `Map<AttributeValue>` | yes |
| `AttributesToGet` | `List<string>` | no |
| `ConsistentRead` | `boolean` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `ProjectionExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Item` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## ImportTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `S3BucketSource` | `S3BucketSource` | yes |
| `InputFormat` | `string` | yes |
| `InputFormatOptions` | `InputFormatOptions` | no |
| `InputCompressionType` | `string` | no |
| `TableCreationParameters` | `TableCreationParameters` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportTableDescription` | `ImportTableDescription` | yes |

## ListBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `Limit` | `integer` | no |
| `TimeRangeLowerBound` | `timestamp` | no |
| `TimeRangeUpperBound` | `timestamp` | no |
| `ExclusiveStartBackupArn` | `string` | no |
| `BackupType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupSummaries` | `List<BackupSummary>` | no |
| `LastEvaluatedBackupArn` | `string` | no |

## ListContributorInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContributorInsightsSummaries` | `List<ContributorInsightsSummary>` | no |
| `NextToken` | `string` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportSummaries` | `List<ExportSummary>` | no |
| `NextToken` | `string` | no |

## ListGlobalTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExclusiveStartGlobalTableName` | `string` | no |
| `Limit` | `integer` | no |
| `RegionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTables` | `List<GlobalTable>` | no |
| `LastEvaluatedGlobalTableName` | `string` | no |

## ListImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableArn` | `string` | no |
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportSummaryList` | `List<ImportSummary>` | no |
| `NextToken` | `string` | no |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExclusiveStartTableName` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableNames` | `List<string>` | no |
| `LastEvaluatedTableName` | `string` | no |

## ListTagsOfResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## PutItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `Item` | `Map<AttributeValue>` | yes |
| `Expected` | `Map<ExpectedAttributeValue>` | no |
| `ReturnValues` | `string` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `ReturnItemCollectionMetrics` | `string` | no |
| `ConditionalOperator` | `string` | no |
| `ConditionExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |
| `ReturnValuesOnConditionCheckFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |
| `ItemCollectionMetrics` | `ItemCollectionMetrics` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |
| `ExpectedRevisionId` | `string` | no |
| `ConfirmRemoveSelfResourceAccess` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevisionId` | `string` | no |

## Query

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `IndexName` | `string` | no |
| `Select` | `string` | no |
| `AttributesToGet` | `List<string>` | no |
| `Limit` | `integer` | no |
| `ConsistentRead` | `boolean` | no |
| `KeyConditions` | `Map<Condition>` | no |
| `QueryFilter` | `Map<Condition>` | no |
| `ConditionalOperator` | `string` | no |
| `ScanIndexForward` | `boolean` | no |
| `ExclusiveStartKey` | `Map<AttributeValue>` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `ProjectionExpression` | `string` | no |
| `FilterExpression` | `string` | no |
| `KeyConditionExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Map<AttributeValue>>` | no |
| `Count` | `integer` | no |
| `ScannedCount` | `integer` | no |
| `LastEvaluatedKey` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |

## RestoreTableFromBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetTableName` | `string` | yes |
| `BackupArn` | `string` | yes |
| `BillingModeOverride` | `string` | no |
| `GlobalSecondaryIndexOverride` | `List<GlobalSecondaryIndex>` | no |
| `LocalSecondaryIndexOverride` | `List<LocalSecondaryIndex>` | no |
| `ProvisionedThroughputOverride` | `ProvisionedThroughput` | no |
| `OnDemandThroughputOverride` | `OnDemandThroughput` | no |
| `SSESpecificationOverride` | `SSESpecification` | no |
| `VectorIndexOverride` | `List<VectorIndex>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableDescription` | `TableDescription` | no |

## RestoreTableToPointInTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceTableArn` | `string` | no |
| `SourceTableName` | `string` | no |
| `TargetTableName` | `string` | yes |
| `UseLatestRestorableTime` | `boolean` | no |
| `RestoreDateTime` | `timestamp` | no |
| `BillingModeOverride` | `string` | no |
| `GlobalSecondaryIndexOverride` | `List<GlobalSecondaryIndex>` | no |
| `LocalSecondaryIndexOverride` | `List<LocalSecondaryIndex>` | no |
| `ProvisionedThroughputOverride` | `ProvisionedThroughput` | no |
| `OnDemandThroughputOverride` | `OnDemandThroughput` | no |
| `SSESpecificationOverride` | `SSESpecification` | no |
| `VectorIndexOverride` | `List<VectorIndex>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableDescription` | `TableDescription` | no |

## Scan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `IndexName` | `string` | no |
| `AttributesToGet` | `List<string>` | no |
| `Limit` | `integer` | no |
| `Select` | `string` | no |
| `ScanFilter` | `Map<Condition>` | no |
| `ConditionalOperator` | `string` | no |
| `ExclusiveStartKey` | `Map<AttributeValue>` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `TotalSegments` | `integer` | no |
| `Segment` | `integer` | no |
| `ProjectionExpression` | `string` | no |
| `FilterExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |
| `ConsistentRead` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Map<AttributeValue>>` | no |
| `Count` | `integer` | no |
| `ScannedCount` | `integer` | no |
| `LastEvaluatedKey` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |

## SearchVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `IndexName` | `string` | yes |
| `ReturnConsumedCapacity` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |
| `ProjectionExpression` | `string` | no |
| `SearchVector` | `List<AttributeValue>` | yes |
| `SearchConditionExpression` | `string` | no |
| `TopK` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumedCapacity` | `VectorCapacity` | no |
| `SearchResults` | `List<SearchResultItem>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TransactGetItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactItems` | `List<TransactGetItem>` | yes |
| `ReturnConsumedCapacity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |
| `Responses` | `List<ItemResponse>` | no |

## TransactWriteItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransactItems` | `List<TransactWriteItem>` | yes |
| `ReturnConsumedCapacity` | `string` | no |
| `ReturnItemCollectionMetrics` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumedCapacity` | `List<ConsumedCapacity>` | no |
| `ItemCollectionMetrics` | `Map<List<ItemCollectionMetrics>>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContinuousBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `PointInTimeRecoverySpecification` | `PointInTimeRecoverySpecification` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousBackupsDescription` | `ContinuousBackupsDescription` | no |

## UpdateContributorInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `IndexName` | `string` | no |
| `ContributorInsightsAction` | `string` | yes |
| `ContributorInsightsMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `IndexName` | `string` | no |
| `ContributorInsightsStatus` | `string` | no |
| `ContributorInsightsMode` | `string` | no |

## UpdateGlobalTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | yes |
| `ReplicaUpdates` | `List<ReplicaUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableDescription` | `GlobalTableDescription` | no |

## UpdateGlobalTableSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | yes |
| `GlobalTableBillingMode` | `string` | no |
| `GlobalTableProvisionedWriteCapacityUnits` | `long` | no |
| `GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate` | `AutoScalingSettingsUpdate` | no |
| `GlobalTableGlobalSecondaryIndexSettingsUpdate` | `List<GlobalTableGlobalSecondaryIndexSettingsUpdate>` | no |
| `ReplicaSettingsUpdate` | `List<ReplicaSettingsUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalTableName` | `string` | no |
| `ReplicaSettings` | `List<ReplicaSettingsDescription>` | no |

## UpdateItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `Key` | `Map<AttributeValue>` | yes |
| `AttributeUpdates` | `Map<AttributeValueUpdate>` | no |
| `Expected` | `Map<ExpectedAttributeValue>` | no |
| `ConditionalOperator` | `string` | no |
| `ReturnValues` | `string` | no |
| `ReturnConsumedCapacity` | `string` | no |
| `ReturnItemCollectionMetrics` | `string` | no |
| `UpdateExpression` | `string` | no |
| `ConditionExpression` | `string` | no |
| `ExpressionAttributeNames` | `Map<string>` | no |
| `ExpressionAttributeValues` | `Map<AttributeValue>` | no |
| `ReturnValuesOnConditionCheckFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<AttributeValue>` | no |
| `ConsumedCapacity` | `ConsumedCapacity` | no |
| `ItemCollectionMetrics` | `ItemCollectionMetrics` | no |

## UpdateKinesisStreamingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `StreamArn` | `string` | yes |
| `UpdateKinesisStreamingConfiguration` | `UpdateKinesisStreamingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `StreamArn` | `string` | no |
| `DestinationStatus` | `string` | no |
| `UpdateKinesisStreamingConfiguration` | `UpdateKinesisStreamingConfiguration` | no |

## UpdateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttributeDefinitions` | `List<AttributeDefinition>` | no |
| `TableName` | `string` | yes |
| `BillingMode` | `string` | no |
| `ProvisionedThroughput` | `ProvisionedThroughput` | no |
| `GlobalSecondaryIndexUpdates` | `List<GlobalSecondaryIndexUpdate>` | no |
| `StreamSpecification` | `StreamSpecification` | no |
| `SSESpecification` | `SSESpecification` | no |
| `ReplicaUpdates` | `List<ReplicationGroupUpdate>` | no |
| `TableClass` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `MultiRegionConsistency` | `string` | no |
| `GlobalTableWitnessUpdates` | `List<GlobalTableWitnessGroupUpdate>` | no |
| `OnDemandThroughput` | `OnDemandThroughput` | no |
| `WarmThroughput` | `WarmThroughput` | no |
| `GlobalTableSettingsReplicationMode` | `string` | no |
| `VectorIndexUpdates` | `List<VectorIndexUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableDescription` | `TableDescription` | no |

## UpdateTableReplicaAutoScaling

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalSecondaryIndexUpdates` | `List<GlobalSecondaryIndexAutoScalingUpdate>` | no |
| `TableName` | `string` | yes |
| `ProvisionedWriteCapacityAutoScalingUpdate` | `AutoScalingSettingsUpdate` | no |
| `ReplicaUpdates` | `List<ReplicaAutoScalingUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableAutoScalingDescription` | `TableAutoScalingDescription` | no |

## UpdateTimeToLive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `TimeToLiveSpecification` | `TimeToLiveSpecification` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeToLiveSpecification` | `TimeToLiveSpecification` | no |

