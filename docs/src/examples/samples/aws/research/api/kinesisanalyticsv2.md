# Amazon Kinesis Analytics

API version: 2018-05-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesisanalyticsv2/2018-05-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddApplicationCloudWatchLoggingOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | no |
| `CloudWatchLoggingOption` | `CloudWatchLoggingOption` | yes |
| `ConditionalToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `CloudWatchLoggingOptionDescriptions` | `List<CloudWatchLoggingOptionDescription>` | no |
| `OperationId` | `string` | no |

## AddApplicationInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `Input` | `Input` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `InputDescriptions` | `List<InputDescription>` | no |

## AddApplicationInputProcessingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `InputId` | `string` | yes |
| `InputProcessingConfiguration` | `InputProcessingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `InputId` | `string` | no |
| `InputProcessingConfigurationDescription` | `InputProcessingConfigurationDescription` | no |

## AddApplicationOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `Output` | `Output` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `OutputDescriptions` | `List<OutputDescription>` | no |

## AddApplicationReferenceDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `ReferenceDataSource` | `ReferenceDataSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `ReferenceDataSourceDescriptions` | `List<ReferenceDataSourceDescription>` | no |

## AddApplicationVpcConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | no |
| `VpcConfiguration` | `VpcConfiguration` | yes |
| `ConditionalToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `VpcConfigurationDescription` | `VpcConfigurationDescription` | no |
| `OperationId` | `string` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `ApplicationDescription` | `string` | no |
| `RuntimeEnvironment` | `string` | yes |
| `ServiceExecutionRole` | `string` | yes |
| `ApplicationConfiguration` | `ApplicationConfiguration` | no |
| `CloudWatchLoggingOptions` | `List<CloudWatchLoggingOption>` | no |
| `Tags` | `List<Tag>` | no |
| `ApplicationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDetail` | `ApplicationDetail` | yes |

## CreateApplicationPresignedUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `UrlType` | `string` | yes |
| `SessionExpirationDurationInSeconds` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrl` | `string` | no |

## CreateApplicationSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `SnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CreateTimestamp` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationCloudWatchLoggingOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | no |
| `CloudWatchLoggingOptionId` | `string` | yes |
| `ConditionalToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `CloudWatchLoggingOptionDescriptions` | `List<CloudWatchLoggingOptionDescription>` | no |
| `OperationId` | `string` | no |

## DeleteApplicationInputProcessingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `InputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |

## DeleteApplicationOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `OutputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |

## DeleteApplicationReferenceDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `ReferenceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |

## DeleteApplicationSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `SnapshotName` | `string` | yes |
| `SnapshotCreationTimestamp` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationVpcConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | no |
| `VpcConfigurationId` | `string` | yes |
| `ConditionalToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationVersionId` | `long` | no |
| `OperationId` | `string` | no |

## DescribeApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `IncludeAdditionalDetails` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDetail` | `ApplicationDetail` | yes |

## DescribeApplicationOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `OperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationOperationInfoDetails` | `ApplicationOperationInfoDetails` | no |

## DescribeApplicationSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `SnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotDetails` | `SnapshotDetails` | yes |

## DescribeApplicationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `ApplicationVersionId` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationVersionDetail` | `ApplicationDetail` | no |

## DiscoverInputSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |
| `ServiceExecutionRole` | `string` | yes |
| `InputStartingPositionConfiguration` | `InputStartingPositionConfiguration` | no |
| `S3Configuration` | `S3Configuration` | no |
| `InputProcessingConfiguration` | `InputProcessingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSchema` | `SourceSchema` | no |
| `ParsedInputRecords` | `List<List<string>>` | no |
| `ProcessedInputRecords` | `List<string>` | no |
| `RawInputRecords` | `List<string>` | no |

## ListApplicationOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `Operation` | `string` | no |
| `OperationStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationOperationInfoList` | `List<ApplicationOperationInfo>` | no |
| `NextToken` | `string` | no |

## ListApplicationSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotSummaries` | `List<SnapshotDetails>` | no |
| `NextToken` | `string` | no |

## ListApplicationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationVersionSummaries` | `List<ApplicationVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSummaries` | `List<ApplicationSummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## RollbackApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDetail` | `ApplicationDetail` | yes |
| `OperationId` | `string` | no |

## StartApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `RunConfiguration` | `RunConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## StopApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | no |
| `ApplicationConfigurationUpdate` | `ApplicationConfigurationUpdate` | no |
| `ServiceExecutionRoleUpdate` | `string` | no |
| `RunConfigurationUpdate` | `RunConfigurationUpdate` | no |
| `CloudWatchLoggingOptionUpdates` | `List<CloudWatchLoggingOptionUpdate>` | no |
| `ConditionalToken` | `string` | no |
| `RuntimeEnvironmentUpdate` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDetail` | `ApplicationDetail` | yes |
| `OperationId` | `string` | no |

## UpdateApplicationMaintenanceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `ApplicationMaintenanceConfigurationUpdate` | `ApplicationMaintenanceConfigurationUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationARN` | `string` | no |
| `ApplicationMaintenanceConfigurationDescription` | `ApplicationMaintenanceConfigurationDescription` | no |

