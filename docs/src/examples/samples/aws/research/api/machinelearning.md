# Amazon Machine Learning

API version: 2014-12-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/machinelearning/2014-12-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `ResourceType` | `string` | no |

## CreateBatchPrediction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | yes |
| `BatchPredictionName` | `string` | no |
| `MLModelId` | `string` | yes |
| `BatchPredictionDataSourceId` | `string` | yes |
| `OutputUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | no |

## CreateDataSourceFromRDS

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |
| `DataSourceName` | `string` | no |
| `RDSData` | `RDSDataSpec` | yes |
| `RoleARN` | `string` | yes |
| `ComputeStatistics` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |

## CreateDataSourceFromRedshift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |
| `DataSourceName` | `string` | no |
| `DataSpec` | `RedshiftDataSpec` | yes |
| `RoleARN` | `string` | yes |
| `ComputeStatistics` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |

## CreateDataSourceFromS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |
| `DataSourceName` | `string` | no |
| `DataSpec` | `S3DataSpec` | yes |
| `ComputeStatistics` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |

## CreateEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |
| `EvaluationName` | `string` | no |
| `MLModelId` | `string` | yes |
| `EvaluationDataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | no |

## CreateMLModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |
| `MLModelName` | `string` | no |
| `MLModelType` | `string` | yes |
| `Parameters` | `Map<string>` | no |
| `TrainingDataSourceId` | `string` | yes |
| `Recipe` | `string` | no |
| `RecipeUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |

## CreateRealtimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |
| `RealtimeEndpointInfo` | `RealtimeEndpointInfo` | no |

## DeleteBatchPrediction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | no |

## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |

## DeleteEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | no |

## DeleteMLModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |

## DeleteRealtimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |
| `RealtimeEndpointInfo` | `RealtimeEndpointInfo` | no |

## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagKeys` | `List<string>` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `ResourceType` | `string` | no |

## DescribeBatchPredictions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterVariable` | `string` | no |
| `EQ` | `string` | no |
| `GT` | `string` | no |
| `LT` | `string` | no |
| `GE` | `string` | no |
| `LE` | `string` | no |
| `NE` | `string` | no |
| `Prefix` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<BatchPrediction>` | no |
| `NextToken` | `string` | no |

## DescribeDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterVariable` | `string` | no |
| `EQ` | `string` | no |
| `GT` | `string` | no |
| `LT` | `string` | no |
| `GE` | `string` | no |
| `LE` | `string` | no |
| `NE` | `string` | no |
| `Prefix` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<DataSource>` | no |
| `NextToken` | `string` | no |

## DescribeEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterVariable` | `string` | no |
| `EQ` | `string` | no |
| `GT` | `string` | no |
| `LT` | `string` | no |
| `GE` | `string` | no |
| `LE` | `string` | no |
| `NE` | `string` | no |
| `Prefix` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<Evaluation>` | no |
| `NextToken` | `string` | no |

## DescribeMLModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterVariable` | `string` | no |
| `EQ` | `string` | no |
| `GT` | `string` | no |
| `LT` | `string` | no |
| `GE` | `string` | no |
| `LE` | `string` | no |
| `NE` | `string` | no |
| `Prefix` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<MLModel>` | no |
| `NextToken` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `ResourceType` | `string` | no |
| `Tags` | `List<Tag>` | no |

## GetBatchPrediction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | no |
| `MLModelId` | `string` | no |
| `BatchPredictionDataSourceId` | `string` | no |
| `InputDataLocationS3` | `string` | no |
| `CreatedByIamUser` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `OutputUri` | `string` | no |
| `LogUri` | `string` | no |
| `Message` | `string` | no |
| `ComputeTime` | `long` | no |
| `FinishedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |
| `TotalRecordCount` | `long` | no |
| `InvalidRecordCount` | `long` | no |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |
| `Verbose` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |
| `DataLocationS3` | `string` | no |
| `DataRearrangement` | `string` | no |
| `CreatedByIamUser` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `DataSizeInBytes` | `long` | no |
| `NumberOfFiles` | `long` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `LogUri` | `string` | no |
| `Message` | `string` | no |
| `RedshiftMetadata` | `RedshiftMetadata` | no |
| `RDSMetadata` | `RDSMetadata` | no |
| `RoleARN` | `string` | no |
| `ComputeStatistics` | `boolean` | no |
| `ComputeTime` | `long` | no |
| `FinishedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |
| `DataSourceSchema` | `string` | no |

## GetEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | no |
| `MLModelId` | `string` | no |
| `EvaluationDataSourceId` | `string` | no |
| `InputDataLocationS3` | `string` | no |
| `CreatedByIamUser` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `PerformanceMetrics` | `PerformanceMetrics` | no |
| `LogUri` | `string` | no |
| `Message` | `string` | no |
| `ComputeTime` | `long` | no |
| `FinishedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |

## GetMLModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |
| `Verbose` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |
| `TrainingDataSourceId` | `string` | no |
| `CreatedByIamUser` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `SizeInBytes` | `long` | no |
| `EndpointInfo` | `RealtimeEndpointInfo` | no |
| `TrainingParameters` | `Map<string>` | no |
| `InputDataLocationS3` | `string` | no |
| `MLModelType` | `string` | no |
| `ScoreThreshold` | `float` | no |
| `ScoreThresholdLastUpdatedAt` | `timestamp` | no |
| `LogUri` | `string` | no |
| `Message` | `string` | no |
| `ComputeTime` | `long` | no |
| `FinishedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |
| `Recipe` | `string` | no |
| `Schema` | `string` | no |

## Predict

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |
| `Record` | `Map<string>` | yes |
| `PredictEndpoint` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Prediction` | `Prediction` | no |

## UpdateBatchPrediction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | yes |
| `BatchPredictionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchPredictionId` | `string` | no |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | yes |
| `DataSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceId` | `string` | no |

## UpdateEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |
| `EvaluationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | no |

## UpdateMLModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | yes |
| `MLModelName` | `string` | no |
| `ScoreThreshold` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MLModelId` | `string` | no |

