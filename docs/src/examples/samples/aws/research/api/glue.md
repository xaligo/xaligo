# AWS Glue

API version: 2017-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/glue/2017-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateGlossaryTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `GlossaryTermIdentifiers` | `List<string>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | no |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `GlossaryTerms` | `List<string>` | no |

## BatchCreatePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionInputList` | `List<PartitionInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<PartitionError>` | no |

## BatchDeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ConnectionNameList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Succeeded` | `List<string>` | no |
| `Errors` | `Map<ErrorDetail>` | no |

## BatchDeletePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionsToDelete` | `List<PartitionValueList>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<PartitionError>` | no |

## BatchDeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TablesToDelete` | `List<string>` | yes |
| `TransactionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<TableError>` | no |

## BatchDeleteTableVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `VersionIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<TableVersionError>` | no |

## BatchGetBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |
| `IncludeBlueprint` | `boolean` | no |
| `IncludeParameterSpec` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blueprints` | `List<Blueprint>` | no |
| `MissingBlueprints` | `List<string>` | no |

## BatchGetCrawlers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Crawlers` | `List<Crawler>` | no |
| `CrawlersNotFound` | `List<string>` | no |

## BatchGetCustomEntityTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomEntityTypes` | `List<CustomEntityType>` | no |
| `CustomEntityTypesNotFound` | `List<string>` | no |

## BatchGetDataQualityResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<DataQualityResult>` | yes |
| `ResultsNotFound` | `List<string>` | no |

## BatchGetDataQualityRulesetEvaluationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Runs` | `List<DataQualityRulesetEvaluationRun>` | no |
| `RunsNotFound` | `List<string>` | no |

## BatchGetDevEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevEndpointNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevEndpoints` | `List<DevEndpoint>` | no |
| `DevEndpointsNotFound` | `List<string>` | no |

## BatchGetIterableForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | yes |
| `ItemIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<IterableFormItem>` | no |
| `Errors` | `List<ItemError>` | no |

## BatchGetJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `JobsNotFound` | `List<string>` | no |

## BatchGetPartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionsToGet` | `List<PartitionValueList>` | yes |
| `AuditContext` | `AuditContext` | no |
| `QuerySessionContext` | `QuerySessionContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Partitions` | `List<Partition>` | no |
| `UnprocessedKeys` | `List<PartitionValueList>` | no |

## BatchGetTableOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<BatchGetTableOptimizerEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableOptimizers` | `List<BatchTableOptimizer>` | no |
| `Failures` | `List<BatchGetTableOptimizerError>` | no |

## BatchGetTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TriggerNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Triggers` | `List<Trigger>` | no |
| `TriggersNotFound` | `List<string>` | no |

## BatchGetWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |
| `IncludeGraph` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workflows` | `List<Workflow>` | no |
| `MissingWorkflows` | `List<string>` | no |

## BatchPutDataQualityStatisticAnnotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InclusionAnnotations` | `List<DatapointInclusionAnnotation>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedInclusionAnnotations` | `List<AnnotationError>` | no |

## BatchStopJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobRunIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulSubmissions` | `List<BatchStopJobRunSuccessfulSubmission>` | no |
| `Errors` | `List<BatchStopJobRunError>` | no |

## BatchUpdatePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Entries` | `List<BatchUpdatePartitionRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchUpdatePartitionFailureEntry>` | no |

## CancelDataQualityRuleRecommendationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelDataQualityRulesetEvaluationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelMLTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `TaskRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |
| `TaskRunId` | `string` | no |
| `Status` | `string` | no |

## CancelStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `Id` | `integer` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckSchemaVersionValidity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataFormat` | `string` | yes |
| `SchemaDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Valid` | `boolean` | no |
| `Error` | `string` | no |

## CreateBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `BlueprintLocation` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CatalogInput` | `CatalogInput` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrokClassifier` | `CreateGrokClassifierRequest` | no |
| `XMLClassifier` | `CreateXMLClassifierRequest` | no |
| `JsonClassifier` | `CreateJsonClassifierRequest` | no |
| `CsvClassifier` | `CreateCsvClassifierRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateColumnStatisticsTaskSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Role` | `string` | yes |
| `Schedule` | `string` | no |
| `ColumnNameList` | `List<string>` | no |
| `SampleSize` | `double` | no |
| `CatalogID` | `string` | no |
| `SecurityConfiguration` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ConnectionInput` | `ConnectionInput` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateConnectionStatus` | `string` | no |

## CreateCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Role` | `string` | yes |
| `DatabaseName` | `string` | no |
| `Description` | `string` | no |
| `Targets` | `CrawlerTargets` | yes |
| `Schedule` | `string` | no |
| `Classifiers` | `List<string>` | no |
| `TablePrefix` | `string` | no |
| `SchemaChangePolicy` | `SchemaChangePolicy` | no |
| `RecrawlPolicy` | `RecrawlPolicy` | no |
| `LineageConfiguration` | `LineageConfiguration` | no |
| `LakeFormationConfiguration` | `LakeFormationConfiguration` | no |
| `Configuration` | `string` | no |
| `CrawlerSecurityConfiguration` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCustomEntityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RegexString` | `string` | yes |
| `ContextWords` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateDataQualityRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Ruleset` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `TargetTable` | `DataQualityTargetTable` | no |
| `DataQualitySecurityConfiguration` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseInput` | `DatabaseInput` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDevEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetId` | `string` | no |
| `PublicKey` | `string` | no |
| `PublicKeys` | `List<string>` | no |
| `NumberOfNodes` | `integer` | no |
| `WorkerType` | `string` | no |
| `GlueVersion` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `ExtraPythonLibsS3Path` | `string` | no |
| `ExtraJarsS3Path` | `string` | no |
| `SecurityConfiguration` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Arguments` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `Status` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetId` | `string` | no |
| `RoleArn` | `string` | no |
| `YarnEndpointAddress` | `string` | no |
| `ZeppelinRemoteSparkInterpreterPort` | `integer` | no |
| `NumberOfNodes` | `integer` | no |
| `WorkerType` | `string` | no |
| `GlueVersion` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `AvailabilityZone` | `string` | no |
| `VpcId` | `string` | no |
| `ExtraPythonLibsS3Path` | `string` | no |
| `ExtraJarsS3Path` | `string` | no |
| `FailureReason` | `string` | no |
| `SecurityConfiguration` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `Arguments` | `Map<string>` | no |

## CreateGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## CreateGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlossaryIdentifier` | `string` | yes |
| `Name` | `string` | yes |
| `ShortDescription` | `string` | no |
| `LongDescription` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `GlossaryId` | `string` | no |
| `Name` | `string` | no |
| `ShortDescription` | `string` | no |
| `LongDescription` | `string` | no |

## CreateGlueIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceArn` | `string` | yes |
| `Scopes` | `List<string>` | no |
| `UserBackgroundSessionsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |

## CreateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationName` | `string` | yes |
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `Description` | `string` | no |
| `DataFilter` | `string` | no |
| `KmsKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `IntegrationConfig` | `IntegrationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `IntegrationName` | `string` | yes |
| `Description` | `string` | no |
| `IntegrationArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `Status` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `Errors` | `List<IntegrationError>` | no |
| `DataFilter` | `string` | no |
| `IntegrationConfig` | `IntegrationConfig` | no |

## CreateIntegrationResourceProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `SourceProcessingProperties` | `SourceProcessingProperties` | no |
| `TargetProcessingProperties` | `TargetProcessingProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourcePropertyArn` | `string` | no |
| `SourceProcessingProperties` | `SourceProcessingProperties` | no |
| `TargetProcessingProperties` | `TargetProcessingProperties` | no |

## CreateIntegrationTableProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TableName` | `string` | yes |
| `SourceTableConfig` | `SourceTableConfig` | no |
| `TargetTableConfig` | `TargetTableConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `JobMode` | `string` | no |
| `JobRunQueuingEnabled` | `boolean` | no |
| `Description` | `string` | no |
| `LogUri` | `string` | no |
| `Role` | `string` | yes |
| `ExecutionProperty` | `ExecutionProperty` | no |
| `Command` | `JobCommand` | yes |
| `DefaultArguments` | `Map<string>` | no |
| `NonOverridableArguments` | `Map<string>` | no |
| `Connections` | `ConnectionsList` | no |
| `MaxRetries` | `integer` | no |
| `AllocatedCapacity` | `integer` | no |
| `Timeout` | `integer` | no |
| `MaxCapacity` | `double` | no |
| `SecurityConfiguration` | `string` | no |
| `Tags` | `Map<string>` | no |
| `NotificationProperty` | `NotificationProperty` | no |
| `GlueVersion` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `WorkerType` | `string` | no |
| `CodeGenConfigurationNodes` | `Map<CodeGenConfigurationNode>` | no |
| `ExecutionClass` | `string` | no |
| `SourceControlDetails` | `SourceControlDetails` | no |
| `MaintenanceWindow` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateMLTransform

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `InputRecordTables` | `List<GlueTable>` | yes |
| `Parameters` | `TransformParameters` | yes |
| `Role` | `string` | yes |
| `GlueVersion` | `string` | no |
| `MaxCapacity` | `double` | no |
| `WorkerType` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TransformEncryption` | `TransformEncryption` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |

## CreatePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionInput` | `PartitionInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePartitionIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionIndex` | `PartitionIndex` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryArn` | `string` | no |
| `RegistryName` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryId` | `RegistryId` | no |
| `SchemaName` | `string` | yes |
| `DataFormat` | `string` | yes |
| `Compatibility` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |
| `SchemaDefinition` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |
| `RegistryArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaArn` | `string` | no |
| `Description` | `string` | no |
| `DataFormat` | `string` | no |
| `Compatibility` | `string` | no |
| `SchemaCheckpoint` | `long` | no |
| `LatestSchemaVersion` | `long` | no |
| `NextSchemaVersion` | `long` | no |
| `SchemaStatus` | `string` | no |
| `Tags` | `Map<string>` | no |
| `SchemaVersionId` | `string` | no |
| `SchemaVersionStatus` | `string` | no |

## CreateScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DagNodes` | `List<CodeGenNode>` | no |
| `DagEdges` | `List<CodeGenEdge>` | no |
| `Language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PythonScript` | `string` | no |
| `ScalaCode` | `string` | no |

## CreateSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EncryptionConfiguration` | `EncryptionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## CreateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `Role` | `string` | yes |
| `Command` | `SessionCommand` | yes |
| `Timeout` | `integer` | no |
| `IdleTimeout` | `integer` | no |
| `DefaultArguments` | `Map<string>` | no |
| `Connections` | `ConnectionsList` | no |
| `MaxCapacity` | `double` | no |
| `NumberOfWorkers` | `integer` | no |
| `WorkerType` | `string` | no |
| `SecurityConfiguration` | `string` | no |
| `GlueVersion` | `string` | no |
| `Tags` | `Map<string>` | no |
| `RequestOrigin` | `string` | no |
| `SessionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Session` | `Session` | no |

## CreateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `Name` | `string` | no |
| `TableInput` | `TableInput` | no |
| `PartitionIndexes` | `List<PartitionIndex>` | no |
| `TransactionId` | `string` | no |
| `OpenTableFormatInput` | `OpenTableFormatInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTableOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Type` | `string` | yes |
| `TableOptimizerConfiguration` | `TableOptimizerConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `WorkflowName` | `string` | no |
| `Type` | `string` | yes |
| `Schedule` | `string` | no |
| `Predicate` | `Predicate` | no |
| `Actions` | `List<Action>` | yes |
| `Description` | `string` | no |
| `StartOnCreation` | `boolean` | no |
| `Tags` | `Map<string>` | no |
| `EventBatchingCondition` | `EventBatchingCondition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateUsageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Configuration` | `ProfileConfiguration` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateUserDefinedFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `FunctionInput` | `UserDefinedFunctionInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DefaultRunProperties` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |
| `MaxConcurrentRuns` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `AttachmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | no |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |

## DeleteBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## DeleteCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteColumnStatisticsForPartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |
| `ColumnName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteColumnStatisticsForTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `ColumnName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteColumnStatisticsTaskSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `ConnectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomEntityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## DeleteDataQualityRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDevEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlueIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `IntegrationName` | `string` | yes |
| `Description` | `string` | no |
| `IntegrationArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `Status` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `Errors` | `List<IntegrationError>` | no |
| `DataFilter` | `string` | no |

## DeleteIntegrationResourceProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegrationTableProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |

## DeleteMLTransform

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |

## DeletePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePartitionIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `IndexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryId` | `RegistryId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |
| `RegistryArn` | `string` | no |
| `Status` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyHashCondition` | `string` | no |
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `Status` | `string` | no |

## DeleteSchemaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `Versions` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaVersionErrors` | `List<SchemaVersionErrorItem>` | no |

## DeleteSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## DeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `Name` | `string` | yes |
| `TransactionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `VersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## DeleteUsageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserDefinedFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## DescribeConnectionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionType` | `string` | no |
| `Description` | `string` | no |
| `Capabilities` | `Capabilities` | no |
| `ConnectionProperties` | `Map<Property>` | no |
| `ConnectionOptions` | `Map<Property>` | no |
| `AuthenticationConfiguration` | `AuthConfiguration` | no |
| `ComputeEnvironmentConfigurations` | `Map<ComputeEnvironmentConfiguration>` | no |
| `PhysicalConnectionRequirements` | `Map<Property>` | no |
| `AthenaConnectionProperties` | `Map<Property>` | no |
| `PythonConnectionProperties` | `Map<Property>` | no |
| `SparkConnectionProperties` | `Map<Property>` | no |
| `RestConfiguration` | `RestConfiguration` | no |

## DescribeEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | yes |
| `CatalogId` | `string` | no |
| `EntityName` | `string` | yes |
| `NextToken` | `string` | no |
| `DataStoreApiVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fields` | `List<Field>` | no |
| `NextToken` | `string` | no |

## DescribeInboundIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationArn` | `string` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `TargetArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InboundIntegrations` | `List<InboundIntegration>` | no |
| `Marker` | `string` | no |

## DescribeIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Filters` | `List<IntegrationFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Integrations` | `List<Integration>` | no |
| `Marker` | `string` | no |

## DisassociateGlossaryTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `GlossaryTermIdentifiers` | `List<string>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | no |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `GlossaryTerms` | `List<string>` | no |

## GetAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `AssetTypeId` | `string` | yes |
| `GlossaryTerms` | `List<string>` | no |
| `Forms` | `Map<AssetFormEntry>` | no |
| `Attachments` | `Map<AssetFormEntry>` | no |
| `IterableForms` | `Map<IterableFormEntry>` | no |

## GetAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Forms` | `Map<AssetTypeFormReference>` | no |

## GetBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IncludeBlueprint` | `boolean` | no |
| `IncludeParameterSpec` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blueprint` | `Blueprint` | no |

## GetBlueprintRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintName` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintRun` | `BlueprintRun` | no |

## GetBlueprintRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintRuns` | `List<BlueprintRun>` | no |
| `NextToken` | `string` | no |

## GetCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `Catalog` | no |

## GetCatalogImportStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportStatus` | `CatalogImportStatus` | no |

## GetCatalogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParentCatalogId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Recursive` | `boolean` | no |
| `IncludeRoot` | `boolean` | no |
| `HasDatabases` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogList` | `List<Catalog>` | yes |
| `NextToken` | `string` | no |

## GetClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Classifier` | `Classifier` | no |

## GetClassifiers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Classifiers` | `List<Classifier>` | no |
| `NextToken` | `string` | no |

## GetColumnStatisticsForPartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |
| `ColumnNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsList` | `List<ColumnStatistics>` | no |
| `Errors` | `List<ColumnError>` | no |

## GetColumnStatisticsForTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `ColumnNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsList` | `List<ColumnStatistics>` | no |
| `Errors` | `List<ColumnError>` | no |

## GetColumnStatisticsTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskRun` | `ColumnStatisticsTaskRun` | no |

## GetColumnStatisticsTaskRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskRuns` | `List<ColumnStatisticsTaskRun>` | no |
| `NextToken` | `string` | no |

## GetColumnStatisticsTaskSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskSettings` | `ColumnStatisticsTaskSettings` | no |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Name` | `string` | yes |
| `HidePassword` | `boolean` | no |
| `ApplyOverrideForComputeEnvironment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## GetConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Filter` | `GetConnectionsFilter` | no |
| `HidePassword` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionList` | `List<Connection>` | no |
| `NextToken` | `string` | no |

## GetCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Crawler` | `Crawler` | no |

## GetCrawlerMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerNameList` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerMetricsList` | `List<CrawlerMetrics>` | no |
| `NextToken` | `string` | no |

## GetCrawlers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Crawlers` | `List<Crawler>` | no |
| `NextToken` | `string` | no |

## GetCustomEntityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `RegexString` | `string` | no |
| `ContextWords` | `List<string>` | no |

## GetDashboardUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | yes |

## GetDataCatalogEncryptionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataCatalogEncryptionSettings` | `DataCatalogEncryptionSettings` | no |

## GetDataCatalogExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportSetting` | `string` | no |
| `Status` | `string` | no |
| `EncryptionConfiguration` | `ExportEncryptionConfiguration` | no |
| `S3TableBucketArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## GetDataQualityModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticId` | `string` | no |
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `CompletedOn` | `timestamp` | no |
| `FailureReason` | `string` | no |

## GetDataQualityModelResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticId` | `string` | yes |
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompletedOn` | `timestamp` | no |
| `Model` | `List<StatisticModelResult>` | no |

## GetDataQualityResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResultId` | `string` | no |
| `ProfileId` | `string` | no |
| `Score` | `double` | no |
| `DataSource` | `DataSource` | no |
| `RulesetName` | `string` | no |
| `EvaluationContext` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `CompletedOn` | `timestamp` | no |
| `JobName` | `string` | no |
| `JobRunId` | `string` | no |
| `RulesetEvaluationRunId` | `string` | no |
| `RuleResults` | `List<DataQualityRuleResult>` | no |
| `AnalyzerResults` | `List<DataQualityAnalyzerResult>` | no |
| `Observations` | `List<DataQualityObservation>` | no |
| `AggregatedMetrics` | `DataQualityAggregatedMetrics` | no |

## GetDataQualityRuleRecommendationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |
| `DataSource` | `DataSource` | no |
| `Role` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `Status` | `string` | no |
| `ErrorString` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |
| `CompletedOn` | `timestamp` | no |
| `ExecutionTime` | `integer` | no |
| `RecommendedRuleset` | `string` | no |
| `CreatedRulesetName` | `string` | no |
| `DataQualitySecurityConfiguration` | `string` | no |
| `AdditionalRunOptions` | `DataQualityRuleRecommendationRunAdditionalRunOptions` | no |

## GetDataQualityRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Ruleset` | `string` | no |
| `TargetTable` | `DataQualityTargetTable` | no |
| `CreatedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |
| `RecommendationRunId` | `string` | no |
| `DataQualitySecurityConfiguration` | `string` | no |

## GetDataQualityRulesetEvaluationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |
| `DataSource` | `DataSource` | no |
| `Role` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `AdditionalRunOptions` | `DataQualityEvaluationRunAdditionalRunOptions` | no |
| `Status` | `string` | no |
| `ErrorString` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |
| `CompletedOn` | `timestamp` | no |
| `ExecutionTime` | `integer` | no |
| `RulesetNames` | `List<string>` | no |
| `ResultIds` | `List<string>` | no |
| `AdditionalDataSources` | `Map<DataSource>` | no |

## GetDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Database` | `Database` | no |

## GetDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceShareType` | `string` | no |
| `AttributesToGet` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseList` | `List<Database>` | yes |
| `NextToken` | `string` | no |

## GetDataflowGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PythonScript` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DagNodes` | `List<CodeGenNode>` | no |
| `DagEdges` | `List<CodeGenEdge>` | no |

## GetDevEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevEndpoint` | `DevEndpoint` | no |

## GetDevEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevEndpoints` | `List<DevEndpoint>` | no |
| `NextToken` | `string` | no |

## GetEntityRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | no |
| `CatalogId` | `string` | no |
| `EntityName` | `string` | yes |
| `NextToken` | `string` | no |
| `DataStoreApiVersion` | `string` | no |
| `ConnectionOptions` | `Map<string>` | no |
| `FilterPredicate` | `string` | no |
| `Limit` | `long` | yes |
| `OrderBy` | `string` | no |
| `SelectedFields` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | no |
| `NextToken` | `string` | no |

## GetFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Schema` | `string` | no |

## GetGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## GetGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `GlossaryId` | `string` | no |
| `Name` | `string` | no |
| `ShortDescription` | `string` | no |
| `LongDescription` | `string` | no |

## GetGlueIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationArn` | `string` | no |
| `InstanceArn` | `string` | no |
| `Scopes` | `List<string>` | no |
| `UserBackgroundSessionsEnabled` | `boolean` | no |

## GetIntegrationResourceProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourcePropertyArn` | `string` | no |
| `SourceProcessingProperties` | `SourceProcessingProperties` | no |
| `TargetProcessingProperties` | `TargetProcessingProperties` | no |

## GetIntegrationTableProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `TableName` | `string` | no |
| `SourceTableConfig` | `SourceTableConfig` | no |
| `TargetTableConfig` | `TargetTableConfig` | no |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Job` | `Job` | no |

## GetJobBookmark

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `RunId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobBookmarkEntry` | `JobBookmarkEntry` | no |

## GetJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `RunId` | `string` | yes |
| `PredecessorsIncluded` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobRun` | `JobRun` | no |

## GetJobRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobRuns` | `List<JobRun>` | no |
| `NextToken` | `string` | no |

## GetJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<Job>` | no |
| `NextToken` | `string` | no |

## GetMLTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `TaskRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |
| `TaskRunId` | `string` | no |
| `Status` | `string` | no |
| `LogGroupName` | `string` | no |
| `Properties` | `TaskRunProperties` | no |
| `ErrorString` | `string` | no |
| `StartedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |
| `CompletedOn` | `timestamp` | no |
| `ExecutionTime` | `integer` | no |

## GetMLTaskRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `TaskRunFilterCriteria` | no |
| `Sort` | `TaskRunSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskRuns` | `List<TaskRun>` | no |
| `NextToken` | `string` | no |

## GetMLTransform

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `CreatedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |
| `InputRecordTables` | `List<GlueTable>` | no |
| `Parameters` | `TransformParameters` | no |
| `EvaluationMetrics` | `EvaluationMetrics` | no |
| `LabelCount` | `integer` | no |
| `Schema` | `List<SchemaColumn>` | no |
| `Role` | `string` | no |
| `GlueVersion` | `string` | no |
| `MaxCapacity` | `double` | no |
| `WorkerType` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `MaxRetries` | `integer` | no |
| `TransformEncryption` | `TransformEncryption` | no |

## GetMLTransforms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `TransformFilterCriteria` | no |
| `Sort` | `TransformSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Transforms` | `List<MLTransform>` | yes |
| `NextToken` | `string` | no |

## GetMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `CatalogEntry` | yes |
| `Sinks` | `List<CatalogEntry>` | no |
| `Location` | `Location` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mapping` | `List<MappingEntry>` | yes |

## GetMaterializedViewRefreshTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `MaterializedViewRefreshTaskRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaterializedViewRefreshTaskRun` | `MaterializedViewRefreshTaskRun` | no |

## GetPartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |
| `AuditContext` | `AuditContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Partition` | `Partition` | no |

## GetPartitionIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartitionIndexDescriptorList` | `List<PartitionIndexDescriptor>` | no |
| `NextToken` | `string` | no |

## GetPartitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Expression` | `string` | no |
| `NextToken` | `string` | no |
| `Segment` | `Segment` | no |
| `MaxResults` | `integer` | no |
| `ExcludeColumnSchema` | `boolean` | no |
| `TransactionId` | `string` | no |
| `QueryAsOfTime` | `timestamp` | no |
| `AuditContext` | `AuditContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Partitions` | `List<Partition>` | no |
| `NextToken` | `string` | no |

## GetPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mapping` | `List<MappingEntry>` | yes |
| `Source` | `CatalogEntry` | yes |
| `Sinks` | `List<CatalogEntry>` | no |
| `Location` | `Location` | no |
| `Language` | `string` | no |
| `AdditionalPlanOptionsMap` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PythonScript` | `string` | no |
| `ScalaCode` | `string` | no |

## GetRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryId` | `RegistryId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |
| `RegistryArn` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `CreatedTime` | `string` | no |
| `UpdatedTime` | `string` | no |

## GetResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GetResourcePoliciesResponseList` | `List<GluePolicy>` | no |
| `NextToken` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyInJson` | `string` | no |
| `PolicyHash` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `UpdateTime` | `timestamp` | no |

## GetSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |
| `RegistryArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaArn` | `string` | no |
| `Description` | `string` | no |
| `DataFormat` | `string` | no |
| `Compatibility` | `string` | no |
| `SchemaCheckpoint` | `long` | no |
| `LatestSchemaVersion` | `long` | no |
| `NextSchemaVersion` | `long` | no |
| `SchemaStatus` | `string` | no |
| `CreatedTime` | `string` | no |
| `UpdatedTime` | `string` | no |

## GetSchemaByDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `SchemaDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaVersionId` | `string` | no |
| `SchemaArn` | `string` | no |
| `DataFormat` | `string` | no |
| `Status` | `string` | no |
| `CreatedTime` | `string` | no |

## GetSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | no |
| `SchemaVersionId` | `string` | no |
| `SchemaVersionNumber` | `SchemaVersionNumber` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaVersionId` | `string` | no |
| `SchemaDefinition` | `string` | no |
| `DataFormat` | `string` | no |
| `SchemaArn` | `string` | no |
| `VersionNumber` | `long` | no |
| `Status` | `string` | no |
| `CreatedTime` | `string` | no |

## GetSchemaVersionsDiff

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `FirstSchemaVersionNumber` | `SchemaVersionNumber` | yes |
| `SecondSchemaVersionNumber` | `SchemaVersionNumber` | yes |
| `SchemaDiffType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Diff` | `string` | no |

## GetSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityConfiguration` | `SecurityConfiguration` | no |

## GetSecurityConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityConfigurations` | `List<SecurityConfiguration>` | no |
| `NextToken` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Session` | `Session` | no |

## GetSessionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SparkConnect` | `SessionEndpoint` | yes |

## GetStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `Id` | `integer` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statement` | `Statement` | no |

## GetTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `Name` | `string` | yes |
| `TransactionId` | `string` | no |
| `QueryAsOfTime` | `timestamp` | no |
| `AuditContext` | `AuditContext` | no |
| `IncludeStatusDetails` | `boolean` | no |
| `AttributesToGet` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `Table` | no |

## GetTableOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | no |
| `TableName` | `string` | no |
| `TableOptimizer` | `TableOptimizer` | no |

## GetTableVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `VersionId` | `string` | no |
| `AuditContext` | `AuditContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableVersion` | `TableVersion` | no |

## GetTableVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AuditContext` | `AuditContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableVersions` | `List<TableVersion>` | no |
| `NextToken` | `string` | no |

## GetTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `Expression` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `TransactionId` | `string` | no |
| `QueryAsOfTime` | `timestamp` | no |
| `AuditContext` | `AuditContext` | no |
| `IncludeStatusDetails` | `boolean` | no |
| `AttributesToGet` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableList` | `List<Table>` | no |
| `NextToken` | `string` | no |

## GetTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## GetTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Trigger` | `Trigger` | no |

## GetTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DependentJobName` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Triggers` | `List<Trigger>` | no |
| `NextToken` | `string` | no |

## GetUnfilteredPartitionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Region` | `string` | no |
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |
| `AuditContext` | `AuditContext` | no |
| `SupportedPermissionTypes` | `List<string>` | yes |
| `QuerySessionContext` | `QuerySessionContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Partition` | `Partition` | no |
| `AuthorizedColumns` | `List<string>` | no |
| `IsRegisteredWithLakeFormation` | `boolean` | no |

## GetUnfilteredPartitionsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Region` | `string` | no |
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Expression` | `string` | no |
| `AuditContext` | `AuditContext` | no |
| `SupportedPermissionTypes` | `List<string>` | yes |
| `NextToken` | `string` | no |
| `Segment` | `Segment` | no |
| `MaxResults` | `integer` | no |
| `QuerySessionContext` | `QuerySessionContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnfilteredPartitions` | `List<UnfilteredPartition>` | no |
| `NextToken` | `string` | no |

## GetUnfilteredTableMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Region` | `string` | no |
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `Name` | `string` | yes |
| `AuditContext` | `AuditContext` | no |
| `SupportedPermissionTypes` | `List<string>` | yes |
| `ParentResourceArn` | `string` | no |
| `RootResourceArn` | `string` | no |
| `SupportedDialect` | `SupportedDialect` | no |
| `Permissions` | `List<string>` | no |
| `QuerySessionContext` | `QuerySessionContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Table` | `Table` | no |
| `AuthorizedColumns` | `List<string>` | no |
| `IsRegisteredWithLakeFormation` | `boolean` | no |
| `CellFilters` | `List<ColumnRowFilter>` | no |
| `QueryAuthorizationId` | `string` | no |
| `IsMultiDialectView` | `boolean` | no |
| `IsMaterializedView` | `boolean` | no |
| `ResourceArn` | `string` | no |
| `IsProtected` | `boolean` | no |
| `Permissions` | `List<string>` | no |
| `RowFilter` | `string` | no |

## GetUsageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Configuration` | `ProfileConfiguration` | no |
| `CreatedOn` | `timestamp` | no |
| `LastModifiedOn` | `timestamp` | no |

## GetUserDefinedFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserDefinedFunction` | `UserDefinedFunction` | no |

## GetUserDefinedFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | no |
| `Pattern` | `string` | yes |
| `FunctionType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserDefinedFunctions` | `List<UserDefinedFunction>` | no |
| `NextToken` | `string` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IncludeGraph` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workflow` | `Workflow` | no |

## GetWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |
| `IncludeGraph` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Run` | `WorkflowRun` | no |

## GetWorkflowRunProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunProperties` | `Map<string>` | no |

## GetWorkflowRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IncludeGraph` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Runs` | `List<WorkflowRun>` | no |
| `NextToken` | `string` | no |

## ImportCatalogToGlue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListAssetTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<AssetTypeItem>` | no |
| `NextToken` | `string` | no |

## ListBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blueprints` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListColumnStatisticsTaskRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskRunIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListConnectionTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionTypes` | `List<ConnectionTypeBrief>` | no |
| `NextToken` | `string` | no |

## ListCrawlers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListCrawls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `Filters` | `List<CrawlsFilter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Crawls` | `List<CrawlerHistory>` | no |
| `NextToken` | `string` | no |

## ListCustomEntityTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomEntityTypes` | `List<CustomEntityType>` | no |
| `NextToken` | `string` | no |

## ListDataQualityResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DataQualityResultFilterCriteria` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<DataQualityResultDescription>` | yes |
| `NextToken` | `string` | no |

## ListDataQualityRuleRecommendationRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DataQualityRuleRecommendationRunFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Runs` | `List<DataQualityRuleRecommendationRunDescription>` | no |
| `NextToken` | `string` | no |

## ListDataQualityRulesetEvaluationRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DataQualityRulesetEvaluationRunFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Runs` | `List<DataQualityRulesetEvaluationRunDescription>` | no |
| `NextToken` | `string` | no |

## ListDataQualityRulesets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `DataQualityRulesetFilterCriteria` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rulesets` | `List<DataQualityRulesetListDetails>` | no |
| `NextToken` | `string` | no |

## ListDataQualityStatisticAnnotations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticId` | `string` | no |
| `ProfileId` | `string` | no |
| `TimestampFilter` | `TimestampFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Annotations` | `List<StatisticAnnotation>` | no |
| `NextToken` | `string` | no |

## ListDataQualityStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticId` | `string` | no |
| `ProfileId` | `string` | no |
| `TimestampFilter` | `TimestampFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statistics` | `List<StatisticSummary>` | no |
| `NextToken` | `string` | no |

## ListDevEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevEndpointNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | no |
| `CatalogId` | `string` | no |
| `ParentEntityName` | `string` | no |
| `NextToken` | `string` | no |
| `DataStoreApiVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entities` | `List<Entity>` | no |
| `NextToken` | `string` | no |

## ListFormTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<FormTypeItem>` | yes |
| `NextToken` | `string` | no |

## ListGlossaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<GlossaryItem>` | no |
| `NextToken` | `string` | no |

## ListGlossaryTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlossaryIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<GlossaryTermItem>` | no |
| `NextToken` | `string` | no |

## ListIntegrationResourceProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Filters` | `List<IntegrationResourcePropertyFilter>` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationResourcePropertyList` | `List<IntegrationResourceProperty>` | no |
| `Marker` | `string` | no |

## ListIterableForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<IterableFormListItem>` | no |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListMLTransforms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `TransformFilterCriteria` | no |
| `Sort` | `TransformSortCriteria` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListMaterializedViewRefreshTaskRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | no |
| `TableName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaterializedViewRefreshTaskRuns` | `List<MaterializedViewRefreshTaskRun>` | no |
| `NextToken` | `string` | no |

## ListRegistries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Registries` | `List<RegistryListItem>` | no |
| `NextToken` | `string` | no |

## ListSchemaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Schemas` | `List<SchemaVersionListItem>` | no |
| `NextToken` | `string` | no |

## ListSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryId` | `RegistryId` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Schemas` | `List<SchemaListItem>` | no |
| `NextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ids` | `List<string>` | no |
| `Sessions` | `List<Session>` | no |
| `NextToken` | `string` | no |

## ListStatements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `RequestOrigin` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statements` | `List<Statement>` | no |
| `NextToken` | `string` | no |

## ListTableOptimizerRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Type` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | no |
| `TableName` | `string` | no |
| `NextToken` | `string` | no |
| `TableOptimizerRuns` | `List<TableOptimizerRun>` | no |

## ListTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DependentJobName` | `string` | no |
| `MaxResults` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TriggerNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListUsageProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Profiles` | `List<UsageProfileDefinition>` | no |
| `NextToken` | `string` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workflows` | `List<string>` | no |
| `NextToken` | `string` | no |

## ModifyIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | yes |
| `Description` | `string` | no |
| `DataFilter` | `string` | no |
| `IntegrationConfig` | `IntegrationConfig` | no |
| `IntegrationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `IntegrationName` | `string` | yes |
| `Description` | `string` | no |
| `IntegrationArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `Status` | `string` | yes |
| `CreateTime` | `timestamp` | yes |
| `Errors` | `List<IntegrationError>` | no |
| `DataFilter` | `string` | no |
| `IntegrationConfig` | `IntegrationConfig` | no |

## PutAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetTypeId` | `string` | yes |
| `Identifier` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Forms` | `Map<AssetFormEntry>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Forms` | `Map<AssetFormEntry>` | no |

## PutAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Forms` | `Map<AssetTypeFormReference>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Forms` | `Map<AssetTypeFormReference>` | no |

## PutAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | yes |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `AttachmentName` | `string` | yes |
| `Content` | `string` | yes |
| `FormTypeId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetIdentifier` | `string` | no |
| `IterableFormName` | `string` | no |
| `ItemIdentifier` | `string` | no |
| `AttachmentName` | `string` | no |
| `FormTypeId` | `string` | no |

## PutDataCatalogEncryptionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DataCatalogEncryptionSettings` | `DataCatalogEncryptionSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDataCatalogExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportSetting` | `string` | yes |
| `EncryptionConfiguration` | `ExportEncryptionConfiguration` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportSetting` | `string` | no |
| `EncryptionConfiguration` | `ExportEncryptionConfiguration` | no |

## PutDataQualityProfileAnnotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `InclusionAnnotation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Schema` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Schema` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyInJson` | `string` | yes |
| `ResourceArn` | `string` | no |
| `PolicyHashCondition` | `string` | no |
| `PolicyExistsCondition` | `string` | no |
| `EnableHybrid` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyHash` | `string` | no |

## PutSchemaVersionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | no |
| `SchemaVersionNumber` | `SchemaVersionNumber` | no |
| `SchemaVersionId` | `string` | no |
| `MetadataKeyValue` | `MetadataKeyValuePair` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `RegistryName` | `string` | no |
| `LatestVersion` | `boolean` | no |
| `VersionNumber` | `long` | no |
| `SchemaVersionId` | `string` | no |
| `MetadataKey` | `string` | no |
| `MetadataValue` | `string` | no |

## PutWorkflowRunProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |
| `RunProperties` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## QuerySchemaVersionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | no |
| `SchemaVersionNumber` | `SchemaVersionNumber` | no |
| `SchemaVersionId` | `string` | no |
| `MetadataList` | `List<MetadataKeyValuePair>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetadataInfoMap` | `Map<MetadataInfo>` | no |
| `SchemaVersionId` | `string` | no |
| `NextToken` | `string` | no |

## RegisterConnectionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionType` | `string` | yes |
| `IntegrationType` | `string` | yes |
| `Description` | `string` | no |
| `ConnectionProperties` | `ConnectionPropertiesConfiguration` | yes |
| `ConnectorAuthenticationConfiguration` | `ConnectorAuthenticationConfiguration` | yes |
| `RestConfiguration` | `RestConfiguration` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionTypeArn` | `string` | no |

## RegisterSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `SchemaDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaVersionId` | `string` | no |
| `VersionNumber` | `long` | no |
| `Status` | `string` | no |

## RemoveSchemaVersionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | no |
| `SchemaVersionNumber` | `SchemaVersionNumber` | no |
| `SchemaVersionId` | `string` | no |
| `MetadataKeyValue` | `MetadataKeyValuePair` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `RegistryName` | `string` | no |
| `LatestVersion` | `boolean` | no |
| `VersionNumber` | `long` | no |
| `SchemaVersionId` | `string` | no |
| `MetadataKey` | `string` | no |
| `MetadataValue` | `string` | no |

## ResetJobBookmark

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `RunId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobBookmarkEntry` | `JobBookmarkEntry` | no |

## ResumeWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |
| `NodeIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |
| `NodeIds` | `List<string>` | no |

## RunStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |
| `Code` | `string` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `integer` | no |

## SearchAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchText` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `SearchSort` | no |
| `FilterClause` | `SearchFilterClause` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<SearchResultItem>` | no |
| `NextToken` | `string` | no |

## SearchTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<PropertyPredicate>` | no |
| `SearchText` | `string` | no |
| `SortCriteria` | `List<SortCriterion>` | no |
| `MaxResults` | `integer` | no |
| `ResourceShareType` | `string` | no |
| `IncludeStatusDetails` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TableList` | `List<Table>` | no |

## StartBlueprintRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintName` | `string` | yes |
| `Parameters` | `string` | no |
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |

## StartColumnStatisticsTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `ColumnNameList` | `List<string>` | no |
| `Role` | `string` | yes |
| `SampleSize` | `double` | no |
| `CatalogID` | `string` | no |
| `SecurityConfiguration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ColumnStatisticsTaskRunId` | `string` | no |

## StartColumnStatisticsTaskRunSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCrawlerSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartDataQualityRuleRecommendationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSource` | `DataSource` | yes |
| `Role` | `string` | yes |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `CreatedRulesetName` | `string` | no |
| `DataQualitySecurityConfiguration` | `string` | no |
| `ClientToken` | `string` | no |
| `AdditionalRunOptions` | `DataQualityRuleRecommendationRunAdditionalRunOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |

## StartDataQualityRulesetEvaluationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSource` | `DataSource` | yes |
| `Role` | `string` | yes |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `ClientToken` | `string` | no |
| `AdditionalRunOptions` | `DataQualityEvaluationRunAdditionalRunOptions` | no |
| `RulesetNames` | `List<string>` | yes |
| `AdditionalDataSources` | `Map<DataSource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |

## StartExportLabelsTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `OutputS3Path` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskRunId` | `string` | no |

## StartImportLabelsTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `InputS3Path` | `string` | yes |
| `ReplaceAllLabels` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskRunId` | `string` | no |

## StartJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobRunQueuingEnabled` | `boolean` | no |
| `JobRunId` | `string` | no |
| `Arguments` | `Map<string>` | no |
| `AllocatedCapacity` | `integer` | no |
| `Timeout` | `integer` | no |
| `MaxCapacity` | `double` | no |
| `SecurityConfiguration` | `string` | no |
| `NotificationProperty` | `NotificationProperty` | no |
| `WorkerType` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `ExecutionClass` | `string` | no |
| `ExecutionRoleSessionPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobRunId` | `string` | no |

## StartMLEvaluationTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskRunId` | `string` | no |

## StartMLLabelingSetGenerationTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `OutputS3Path` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskRunId` | `string` | no |

## StartMaterializedViewRefreshTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `FullRefresh` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaterializedViewRefreshTaskRunId` | `string` | no |

## StartTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## StartWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunProperties` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |

## StopColumnStatisticsTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopColumnStatisticsTaskRunSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopCrawlerSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopMaterializedViewRefreshTaskRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `RequestOrigin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## StopTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## StopWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagsToAdd` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | no |
| `CatalogId` | `string` | no |
| `TestConnectionInput` | `TestConnectionInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagsToRemove` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## UpdateBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `BlueprintLocation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## UpdateCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `CatalogInput` | `CatalogInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateClassifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GrokClassifier` | `UpdateGrokClassifierRequest` | no |
| `XMLClassifier` | `UpdateXMLClassifierRequest` | no |
| `JsonClassifier` | `UpdateJsonClassifierRequest` | no |
| `CsvClassifier` | `UpdateCsvClassifierRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateColumnStatisticsForPartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValues` | `List<string>` | yes |
| `ColumnStatisticsList` | `List<ColumnStatistics>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<ColumnStatisticsError>` | no |

## UpdateColumnStatisticsForTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `ColumnStatisticsList` | `List<ColumnStatistics>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<ColumnStatisticsError>` | no |

## UpdateColumnStatisticsTaskSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Role` | `string` | no |
| `Schedule` | `string` | no |
| `ColumnNameList` | `List<string>` | no |
| `SampleSize` | `double` | no |
| `CatalogID` | `string` | no |
| `SecurityConfiguration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Name` | `string` | yes |
| `ConnectionInput` | `ConnectionInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCrawler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Role` | `string` | no |
| `DatabaseName` | `string` | no |
| `Description` | `string` | no |
| `Targets` | `CrawlerTargets` | no |
| `Schedule` | `string` | no |
| `Classifiers` | `List<string>` | no |
| `TablePrefix` | `string` | no |
| `SchemaChangePolicy` | `SchemaChangePolicy` | no |
| `RecrawlPolicy` | `RecrawlPolicy` | no |
| `LineageConfiguration` | `LineageConfiguration` | no |
| `LakeFormationConfiguration` | `LakeFormationConfiguration` | no |
| `Configuration` | `string` | no |
| `CrawlerSecurityConfiguration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCrawlerSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrawlerName` | `string` | yes |
| `Schedule` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataQualityRuleset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Ruleset` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Ruleset` | `string` | no |

## UpdateDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `Name` | `string` | yes |
| `DatabaseInput` | `DatabaseInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDevEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `PublicKey` | `string` | no |
| `AddPublicKeys` | `List<string>` | no |
| `DeletePublicKeys` | `List<string>` | no |
| `CustomLibraries` | `DevEndpointCustomLibraries` | no |
| `UpdateEtlLibraries` | `boolean` | no |
| `DeleteArguments` | `List<string>` | no |
| `AddArguments` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## UpdateGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `ShortDescription` | `string` | no |
| `LongDescription` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `GlossaryId` | `string` | no |
| `Name` | `string` | no |
| `ShortDescription` | `string` | no |
| `LongDescription` | `string` | no |

## UpdateGlueIdentityCenterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scopes` | `List<string>` | no |
| `UserBackgroundSessionsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIntegrationResourceProperty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `SourceProcessingProperties` | `SourceProcessingProperties` | no |
| `TargetProcessingProperties` | `TargetProcessingProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourcePropertyArn` | `string` | no |
| `SourceProcessingProperties` | `SourceProcessingProperties` | no |
| `TargetProcessingProperties` | `TargetProcessingProperties` | no |

## UpdateIntegrationTableProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TableName` | `string` | yes |
| `SourceTableConfig` | `SourceTableConfig` | no |
| `TargetTableConfig` | `TargetTableConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobUpdate` | `JobUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |

## UpdateJobFromSourceControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `Provider` | `string` | no |
| `RepositoryName` | `string` | no |
| `RepositoryOwner` | `string` | no |
| `BranchName` | `string` | no |
| `Folder` | `string` | no |
| `CommitId` | `string` | no |
| `AuthStrategy` | `string` | no |
| `AuthToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |

## UpdateMLTransform

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Parameters` | `TransformParameters` | no |
| `Role` | `string` | no |
| `GlueVersion` | `string` | no |
| `MaxCapacity` | `double` | no |
| `WorkerType` | `string` | no |
| `NumberOfWorkers` | `integer` | no |
| `Timeout` | `integer` | no |
| `MaxRetries` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformId` | `string` | no |

## UpdatePartition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `PartitionValueList` | `List<string>` | yes |
| `PartitionInput` | `PartitionInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryId` | `RegistryId` | yes |
| `Description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |
| `RegistryArn` | `string` | no |

## UpdateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `SchemaId` | yes |
| `SchemaVersionNumber` | `SchemaVersionNumber` | no |
| `Compatibility` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `RegistryName` | `string` | no |

## UpdateSourceControlFromJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `Provider` | `string` | no |
| `RepositoryName` | `string` | no |
| `RepositoryOwner` | `string` | no |
| `BranchName` | `string` | no |
| `Folder` | `string` | no |
| `CommitId` | `string` | no |
| `AuthStrategy` | `string` | no |
| `AuthToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |

## UpdateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `Name` | `string` | no |
| `TableInput` | `TableInput` | no |
| `SkipArchive` | `boolean` | no |
| `TransactionId` | `string` | no |
| `VersionId` | `string` | no |
| `ViewUpdateAction` | `string` | no |
| `Force` | `boolean` | no |
| `UpdateOpenTableFormatInput` | `UpdateOpenTableFormatInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTableOptimizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | yes |
| `DatabaseName` | `string` | yes |
| `TableName` | `string` | yes |
| `Type` | `string` | yes |
| `TableOptimizerConfiguration` | `TableOptimizerConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TriggerUpdate` | `TriggerUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Trigger` | `Trigger` | no |

## UpdateUsageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Configuration` | `ProfileConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## UpdateUserDefinedFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CatalogId` | `string` | no |
| `DatabaseName` | `string` | yes |
| `FunctionName` | `string` | yes |
| `FunctionInput` | `UserDefinedFunctionInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DefaultRunProperties` | `Map<string>` | no |
| `MaxConcurrentRuns` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

