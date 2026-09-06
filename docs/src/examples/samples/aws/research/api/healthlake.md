# Amazon HealthLake

API version: 2017-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/healthlake/2017-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDataTransformationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceFormat` | `string` | yes |
| `Source` | `CreateDataTransformationProfileSource` | yes |
| `KmsKeyId` | `string` | no |
| `ProfileDescription` | `string` | no |
| `ProfileName` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `Version` | `integer` | yes |
| `SourceFormat` | `string` | yes |
| `TargetFormat` | `string` | yes |
| `ProfileName` | `string` | yes |
| `LastUpdatedAt` | `timestamp` | yes |

## CreateFHIRDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreName` | `string` | no |
| `DatastoreTypeVersion` | `string` | yes |
| `SseConfiguration` | `SseConfiguration` | no |
| `PreloadDataConfig` | `PreloadDataConfig` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `IdentityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `AnalyticsConfiguration` | `AnalyticsConfiguration` | no |
| `NlpConfiguration` | `NlpConfiguration` | no |
| `ProfileConfiguration` | `ProfileConfiguration` | no |
| `BackupConfiguration` | `BackupConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `DatastoreArn` | `string` | yes |
| `DatastoreStatus` | `string` | yes |
| `DatastoreEndpoint` | `string` | yes |

## DeleteDataTransformationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ProfileName` | `string` | no |
| `DeletionTime` | `timestamp` | yes |

## DeleteFHIRDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `DatastoreArn` | `string` | yes |
| `DatastoreStatus` | `string` | yes |
| `DatastoreEndpoint` | `string` | yes |

## DescribeDataTransformationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformationJobProperties` | `TransformationJobProperties` | yes |

## DescribeFHIRDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreProperties` | `DatastoreProperties` | yes |

## DescribeFHIRExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobProperties` | `ExportJobProperties` | yes |

## DescribeFHIRImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobProperties` | `ImportJobProperties` | yes |

## GetDataTransformationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ProfileVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `Version` | `integer` | yes |
| `SourceFormat` | `string` | yes |
| `TargetFormat` | `string` | yes |
| `ProfileMapping` | `Map<string>` | yes |
| `ProfileName` | `string` | no |
| `ProfileDescription` | `string` | no |
| `ChangeDescription` | `string` | no |
| `LastUpdatedAt` | `timestamp` | yes |

## ListDataTransformationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `JobStatus` | `string` | no |
| `JobName` | `string` | no |
| `SubmittedAfter` | `timestamp` | no |
| `SubmittedBefore` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<TransformationJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListDataTransformationProfileVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DataTransformationProfileVersionSummary>` | yes |
| `NextToken` | `string` | no |

## ListDataTransformationProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceFormat` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DataTransformationProfileSummary>` | yes |
| `NextToken` | `string` | no |

## ListFHIRDatastores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `DatastoreFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastorePropertiesList` | `List<DatastoreProperties>` | yes |
| `NextToken` | `string` | no |

## ListFHIRExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `JobName` | `string` | no |
| `JobStatus` | `string` | no |
| `SubmittedBefore` | `timestamp` | no |
| `SubmittedAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobPropertiesList` | `List<ExportJobProperties>` | yes |
| `NextToken` | `string` | no |

## ListFHIRImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `JobName` | `string` | no |
| `JobStatus` | `string` | no |
| `SubmittedBefore` | `timestamp` | no |
| `SubmittedAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobPropertiesList` | `List<ImportJobProperties>` | yes |
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

## PublishDataTransformationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `SourceFormat` | `string` | yes |
| `FromExistingVersion` | `integer` | no |
| `ChangeDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `Version` | `integer` | yes |
| `SourceFormat` | `string` | yes |
| `TargetFormat` | `string` | yes |
| `ProfileName` | `string` | no |
| `LastUpdatedAt` | `timestamp` | yes |

## RestoreFHIRDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDatastoreId` | `string` | yes |
| `RestoreConfiguration` | `RestoreConfiguration` | yes |
| `DatastoreName` | `string` | no |
| `SseConfiguration` | `SseConfiguration` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `IdentityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `AnalyticsConfiguration` | `AnalyticsConfiguration` | no |
| `NlpConfiguration` | `NlpConfiguration` | no |
| `ProfileConfiguration` | `ProfileConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `DatastoreArn` | `string` | yes |
| `DatastoreStatus` | `string` | yes |
| `DatastoreEndpoint` | `string` | yes |

## StartDataTransformationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDataConfig` | `TransformationInputDataConfig` | yes |
| `OutputDataConfig` | `TransformationOutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `ClientToken` | `string` | yes |
| `JobName` | `string` | no |
| `ProfileId` | `string` | yes |
| `DriftDetectionEnabled` | `boolean` | no |
| `ProvenanceEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `JobStatus` | `string` | yes |

## StartFHIRExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DatastoreId` | `string` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `JobStatus` | `string` | yes |
| `DatastoreId` | `string` | no |

## StartFHIRImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `InputDataConfig` | `InputDataConfig` | yes |
| `JobOutputDataConfig` | `OutputDataConfig` | yes |
| `DatastoreId` | `string` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `ValidationLevel` | `string` | no |
| `ProfileId` | `string` | no |
| `InputFormat` | `string` | no |
| `DriftDetectionEnabled` | `boolean` | no |
| `ProvenanceEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `JobStatus` | `string` | yes |
| `DatastoreId` | `string` | no |

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


## UpdateDataTransformationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ProfileMapping` | `Map<string>` | yes |
| `ChangeDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `SourceFormat` | `string` | yes |
| `TargetFormat` | `string` | yes |
| `ProfileName` | `string` | no |
| `LastUpdatedAt` | `timestamp` | yes |

## UpdateFHIRDatastore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreId` | `string` | yes |
| `DatastoreName` | `string` | no |
| `AnalyticsConfiguration` | `AnalyticsConfiguration` | no |
| `NlpConfiguration` | `NlpConfiguration` | no |
| `ProfileConfiguration` | `ProfileConfiguration` | no |
| `IdentityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `BackupConfiguration` | `BackupConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatastoreProperties` | `DatastoreProperties` | yes |

## UpdateProfileWithAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `SourceFormat` | `string` | yes |
| `InputMessage` | `AgentInputMessage` | yes |
| `ConversationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentResponse` | `AgentOutputMessage` | yes |
| `ConversationId` | `string` | yes |

