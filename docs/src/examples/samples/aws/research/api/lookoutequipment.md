# Amazon Lookout for Equipment

API version: 2020-12-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lookoutequipment/2020-12-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `DatasetSchema` | `DatasetSchema` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `DatasetArn` | `string` | no |
| `Status` | `string` | no |

## CreateInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `InferenceSchedulerName` | `string` | yes |
| `DataDelayOffsetInMinutes` | `long` | no |
| `DataUploadFrequency` | `string` | yes |
| `DataInputConfiguration` | `InferenceInputConfiguration` | yes |
| `DataOutputConfiguration` | `InferenceOutputConfiguration` | yes |
| `RoleArn` | `string` | yes |
| `ServerSideKmsKeyId` | `string` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerArn` | `string` | no |
| `InferenceSchedulerName` | `string` | no |
| `Status` | `string` | no |
| `ModelQuality` | `string` | no |

## CreateLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Rating` | `string` | yes |
| `FaultCode` | `string` | no |
| `Notes` | `string` | no |
| `Equipment` | `string` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelId` | `string` | no |

## CreateLabelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `FaultCodes` | `List<string>` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | no |
| `LabelGroupArn` | `string` | no |

## CreateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `DatasetName` | `string` | yes |
| `DatasetSchema` | `DatasetSchema` | no |
| `LabelsInputConfiguration` | `LabelsInputConfiguration` | no |
| `ClientToken` | `string` | yes |
| `TrainingDataStartTime` | `timestamp` | no |
| `TrainingDataEndTime` | `timestamp` | no |
| `EvaluationDataStartTime` | `timestamp` | no |
| `EvaluationDataEndTime` | `timestamp` | no |
| `RoleArn` | `string` | no |
| `DataPreProcessingConfiguration` | `DataPreProcessingConfiguration` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `OffCondition` | `string` | no |
| `ModelDiagnosticsOutputConfiguration` | `ModelDiagnosticsOutputConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | no |
| `Status` | `string` | no |

## CreateRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `RetrainingStartDate` | `timestamp` | no |
| `RetrainingFrequency` | `string` | yes |
| `LookbackWindow` | `string` | yes |
| `PromoteMode` | `string` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `Status` | `string` | no |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `LabelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLabelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDataIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `DatasetArn` | `string` | no |
| `IngestionInputConfiguration` | `IngestionInputConfiguration` | no |
| `RoleArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `FailedReason` | `string` | no |
| `DataQualitySummary` | `DataQualitySummary` | no |
| `IngestedFilesSummary` | `IngestedFilesSummary` | no |
| `StatusDetail` | `string` | no |
| `IngestedDataSize` | `long` | no |
| `DataStartTime` | `timestamp` | no |
| `DataEndTime` | `timestamp` | no |
| `SourceDatasetArn` | `string` | no |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `DatasetArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `Schema` | `string` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `IngestionInputConfiguration` | `IngestionInputConfiguration` | no |
| `DataQualitySummary` | `DataQualitySummary` | no |
| `IngestedFilesSummary` | `IngestedFilesSummary` | no |
| `RoleArn` | `string` | no |
| `DataStartTime` | `timestamp` | no |
| `DataEndTime` | `timestamp` | no |
| `SourceDatasetArn` | `string` | no |

## DescribeInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | no |
| `ModelName` | `string` | no |
| `InferenceSchedulerName` | `string` | no |
| `InferenceSchedulerArn` | `string` | no |
| `Status` | `string` | no |
| `DataDelayOffsetInMinutes` | `long` | no |
| `DataUploadFrequency` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `DataInputConfiguration` | `InferenceInputConfiguration` | no |
| `DataOutputConfiguration` | `InferenceOutputConfiguration` | no |
| `RoleArn` | `string` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `LatestInferenceResult` | `string` | no |

## DescribeLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `LabelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | no |
| `LabelGroupArn` | `string` | no |
| `LabelId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Rating` | `string` | no |
| `FaultCode` | `string` | no |
| `Notes` | `string` | no |
| `Equipment` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## DescribeLabelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | no |
| `LabelGroupArn` | `string` | no |
| `FaultCodes` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## DescribeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `DatasetName` | `string` | no |
| `DatasetArn` | `string` | no |
| `Schema` | `string` | no |
| `LabelsInputConfiguration` | `LabelsInputConfiguration` | no |
| `TrainingDataStartTime` | `timestamp` | no |
| `TrainingDataEndTime` | `timestamp` | no |
| `EvaluationDataStartTime` | `timestamp` | no |
| `EvaluationDataEndTime` | `timestamp` | no |
| `RoleArn` | `string` | no |
| `DataPreProcessingConfiguration` | `DataPreProcessingConfiguration` | no |
| `Status` | `string` | no |
| `TrainingExecutionStartTime` | `timestamp` | no |
| `TrainingExecutionEndTime` | `timestamp` | no |
| `FailedReason` | `string` | no |
| `ModelMetrics` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `OffCondition` | `string` | no |
| `SourceModelVersionArn` | `string` | no |
| `ImportJobStartTime` | `timestamp` | no |
| `ImportJobEndTime` | `timestamp` | no |
| `ActiveModelVersion` | `long` | no |
| `ActiveModelVersionArn` | `string` | no |
| `ModelVersionActivatedAt` | `timestamp` | no |
| `PreviousActiveModelVersion` | `long` | no |
| `PreviousActiveModelVersionArn` | `string` | no |
| `PreviousModelVersionActivatedAt` | `timestamp` | no |
| `PriorModelMetrics` | `string` | no |
| `LatestScheduledRetrainingFailedReason` | `string` | no |
| `LatestScheduledRetrainingStatus` | `string` | no |
| `LatestScheduledRetrainingModelVersion` | `long` | no |
| `LatestScheduledRetrainingStartTime` | `timestamp` | no |
| `LatestScheduledRetrainingAvailableDataInDays` | `integer` | no |
| `NextScheduledRetrainingStartDate` | `timestamp` | no |
| `AccumulatedInferenceDataStartTime` | `timestamp` | no |
| `AccumulatedInferenceDataEndTime` | `timestamp` | no |
| `RetrainingSchedulerStatus` | `string` | no |
| `ModelDiagnosticsOutputConfiguration` | `ModelDiagnosticsOutputConfiguration` | no |
| `ModelQuality` | `string` | no |

## DescribeModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `ModelVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `ModelVersion` | `long` | no |
| `ModelVersionArn` | `string` | no |
| `Status` | `string` | no |
| `SourceType` | `string` | no |
| `DatasetName` | `string` | no |
| `DatasetArn` | `string` | no |
| `Schema` | `string` | no |
| `LabelsInputConfiguration` | `LabelsInputConfiguration` | no |
| `TrainingDataStartTime` | `timestamp` | no |
| `TrainingDataEndTime` | `timestamp` | no |
| `EvaluationDataStartTime` | `timestamp` | no |
| `EvaluationDataEndTime` | `timestamp` | no |
| `RoleArn` | `string` | no |
| `DataPreProcessingConfiguration` | `DataPreProcessingConfiguration` | no |
| `TrainingExecutionStartTime` | `timestamp` | no |
| `TrainingExecutionEndTime` | `timestamp` | no |
| `FailedReason` | `string` | no |
| `ModelMetrics` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `OffCondition` | `string` | no |
| `SourceModelVersionArn` | `string` | no |
| `ImportJobStartTime` | `timestamp` | no |
| `ImportJobEndTime` | `timestamp` | no |
| `ImportedDataSizeInBytes` | `long` | no |
| `PriorModelMetrics` | `string` | no |
| `RetrainingAvailableDataInDays` | `integer` | no |
| `AutoPromotionResult` | `string` | no |
| `AutoPromotionResultReason` | `string` | no |
| `ModelDiagnosticsOutputConfiguration` | `ModelDiagnosticsOutputConfiguration` | no |
| `ModelDiagnosticsResultsObject` | `S3Object` | no |
| `ModelQuality` | `string` | no |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyRevisionId` | `string` | no |
| `ResourcePolicy` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribeRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `RetrainingStartDate` | `timestamp` | no |
| `RetrainingFrequency` | `string` | no |
| `LookbackWindow` | `string` | no |
| `Status` | `string` | no |
| `PromoteMode` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## ImportDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDatasetArn` | `string` | yes |
| `DatasetName` | `string` | no |
| `ClientToken` | `string` | yes |
| `ServerSideKmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `DatasetArn` | `string` | no |
| `Status` | `string` | no |
| `JobId` | `string` | no |

## ImportModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceModelVersionArn` | `string` | yes |
| `ModelName` | `string` | no |
| `DatasetName` | `string` | yes |
| `LabelsInputConfiguration` | `LabelsInputConfiguration` | no |
| `ClientToken` | `string` | yes |
| `RoleArn` | `string` | no |
| `ServerSideKmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `InferenceDataImportStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `ModelVersionArn` | `string` | no |
| `ModelVersion` | `long` | no |
| `Status` | `string` | no |

## ListDataIngestionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DataIngestionJobSummaries` | `List<DataIngestionJobSummary>` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DatasetNameBeginsWith` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DatasetSummaries` | `List<DatasetSummary>` | no |

## ListInferenceEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InferenceSchedulerName` | `string` | yes |
| `IntervalStartTime` | `timestamp` | yes |
| `IntervalEndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `InferenceEventSummaries` | `List<InferenceEventSummary>` | no |

## ListInferenceExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InferenceSchedulerName` | `string` | yes |
| `DataStartTimeAfter` | `timestamp` | no |
| `DataEndTimeBefore` | `timestamp` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `InferenceExecutionSummaries` | `List<InferenceExecutionSummary>` | no |

## ListInferenceSchedulers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InferenceSchedulerNameBeginsWith` | `string` | no |
| `ModelName` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `InferenceSchedulerSummaries` | `List<InferenceSchedulerSummary>` | no |

## ListLabelGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupNameBeginsWith` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `LabelGroupSummaries` | `List<LabelGroupSummary>` | no |

## ListLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `IntervalStartTime` | `timestamp` | no |
| `IntervalEndTime` | `timestamp` | no |
| `FaultCode` | `string` | no |
| `Equipment` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `LabelSummaries` | `List<LabelSummary>` | no |

## ListModelVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |
| `SourceType` | `string` | no |
| `CreatedAtEndTime` | `timestamp` | no |
| `CreatedAtStartTime` | `timestamp` | no |
| `MaxModelVersion` | `long` | no |
| `MinModelVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ModelVersionSummaries` | `List<ModelVersionSummary>` | no |

## ListModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |
| `ModelNameBeginsWith` | `string` | no |
| `DatasetNameBeginsWith` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ModelSummaries` | `List<ModelSummary>` | no |

## ListRetrainingSchedulers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelNameBeginsWith` | `string` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrainingSchedulerSummaries` | `List<RetrainingSchedulerSummary>` | no |
| `NextToken` | `string` | no |

## ListSensorStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `IngestionJobId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SensorStatisticsSummaries` | `List<SensorStatisticsSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourcePolicy` | `string` | yes |
| `PolicyRevisionId` | `string` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `PolicyRevisionId` | `string` | no |

## StartDataIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `IngestionInputConfiguration` | `IngestionInputConfiguration` | yes |
| `RoleArn` | `string` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `Status` | `string` | no |

## StartInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | no |
| `ModelName` | `string` | no |
| `InferenceSchedulerName` | `string` | no |
| `InferenceSchedulerArn` | `string` | no |
| `Status` | `string` | no |

## StartRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `Status` | `string` | no |

## StopInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | no |
| `ModelName` | `string` | no |
| `InferenceSchedulerName` | `string` | no |
| `InferenceSchedulerArn` | `string` | no |
| `Status` | `string` | no |

## StopRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `Status` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateActiveModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `ModelVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | no |
| `ModelArn` | `string` | no |
| `CurrentActiveVersion` | `long` | no |
| `PreviousActiveVersion` | `long` | no |
| `CurrentActiveVersionArn` | `string` | no |
| `PreviousActiveVersionArn` | `string` | no |

## UpdateInferenceScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceSchedulerName` | `string` | yes |
| `DataDelayOffsetInMinutes` | `long` | no |
| `DataUploadFrequency` | `string` | no |
| `DataInputConfiguration` | `InferenceInputConfiguration` | no |
| `DataOutputConfiguration` | `InferenceOutputConfiguration` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLabelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelGroupName` | `string` | yes |
| `FaultCodes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `LabelsInputConfiguration` | `LabelsInputConfiguration` | no |
| `RoleArn` | `string` | no |
| `ModelDiagnosticsOutputConfiguration` | `ModelDiagnosticsOutputConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRetrainingScheduler

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `RetrainingStartDate` | `timestamp` | no |
| `RetrainingFrequency` | `string` | no |
| `LookbackWindow` | `string` | no |
| `PromoteMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


