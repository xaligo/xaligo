# Amazon Forecast Service

API version: 2018-06-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/forecast/2018-06-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAutoPredictor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorName` | `string` | yes |
| `ForecastHorizon` | `integer` | no |
| `ForecastTypes` | `List<string>` | no |
| `ForecastDimensions` | `List<string>` | no |
| `ForecastFrequency` | `string` | no |
| `DataConfig` | `DataConfig` | no |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `ReferencePredictorArn` | `string` | no |
| `OptimizationMetric` | `string` | no |
| `ExplainPredictor` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `MonitorConfig` | `MonitorConfig` | no |
| `TimeAlignmentBoundary` | `TimeAlignmentBoundary` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetName` | `string` | yes |
| `Domain` | `string` | yes |
| `DatasetType` | `string` | yes |
| `DataFrequency` | `string` | no |
| `Schema` | `Schema` | yes |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | no |

## CreateDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupName` | `string` | yes |
| `Domain` | `string` | yes |
| `DatasetArns` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupArn` | `string` | no |

## CreateDatasetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobName` | `string` | yes |
| `DatasetArn` | `string` | yes |
| `DataSource` | `DataSource` | yes |
| `TimestampFormat` | `string` | no |
| `TimeZone` | `string` | no |
| `UseGeolocationForTimeZone` | `boolean` | no |
| `GeolocationFormat` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Format` | `string` | no |
| `ImportMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobArn` | `string` | no |

## CreateExplainability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityName` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `ExplainabilityConfig` | `ExplainabilityConfig` | yes |
| `DataSource` | `DataSource` | no |
| `Schema` | `Schema` | no |
| `EnableVisualization` | `boolean` | no |
| `StartDateTime` | `string` | no |
| `EndDateTime` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityArn` | `string` | no |

## CreateExplainabilityExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExportName` | `string` | yes |
| `ExplainabilityArn` | `string` | yes |
| `Destination` | `DataDestination` | yes |
| `Tags` | `List<Tag>` | no |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExportArn` | `string` | no |

## CreateForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastName` | `string` | yes |
| `PredictorArn` | `string` | yes |
| `ForecastTypes` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `TimeSeriesSelector` | `TimeSeriesSelector` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastArn` | `string` | no |

## CreateForecastExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobName` | `string` | yes |
| `ForecastArn` | `string` | yes |
| `Destination` | `DataDestination` | yes |
| `Tags` | `List<Tag>` | no |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobArn` | `string` | no |

## CreateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | no |

## CreatePredictor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorName` | `string` | yes |
| `AlgorithmArn` | `string` | no |
| `ForecastHorizon` | `integer` | yes |
| `ForecastTypes` | `List<string>` | no |
| `PerformAutoML` | `boolean` | no |
| `AutoMLOverrideStrategy` | `string` | no |
| `PerformHPO` | `boolean` | no |
| `TrainingParameters` | `Map<string>` | no |
| `EvaluationParameters` | `EvaluationParameters` | no |
| `HPOConfig` | `HyperParameterTuningJobConfig` | no |
| `InputDataConfig` | `InputDataConfig` | yes |
| `FeaturizationConfig` | `FeaturizationConfig` | yes |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `Tags` | `List<Tag>` | no |
| `OptimizationMetric` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | no |

## CreatePredictorBacktestExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobName` | `string` | yes |
| `PredictorArn` | `string` | yes |
| `Destination` | `DataDestination` | yes |
| `Tags` | `List<Tag>` | no |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobArn` | `string` | no |

## CreateWhatIfAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalysisName` | `string` | yes |
| `ForecastArn` | `string` | yes |
| `TimeSeriesSelector` | `TimeSeriesSelector` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalysisArn` | `string` | no |

## CreateWhatIfForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastName` | `string` | yes |
| `WhatIfAnalysisArn` | `string` | yes |
| `TimeSeriesTransformations` | `List<TimeSeriesTransformation>` | no |
| `TimeSeriesReplacementsDataSource` | `TimeSeriesReplacementsDataSource` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastArn` | `string` | no |

## CreateWhatIfForecastExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExportName` | `string` | yes |
| `WhatIfForecastArns` | `List<string>` | yes |
| `Destination` | `DataDestination` | yes |
| `Tags` | `List<Tag>` | no |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExportArn` | `string` | no |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDatasetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExplainability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExplainabilityExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteForecastExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePredictor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePredictorBacktestExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceTree

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWhatIfAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalysisArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWhatIfForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWhatIfForecastExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAutoPredictor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | no |
| `PredictorName` | `string` | no |
| `ForecastHorizon` | `integer` | no |
| `ForecastTypes` | `List<string>` | no |
| `ForecastFrequency` | `string` | no |
| `ForecastDimensions` | `List<string>` | no |
| `DatasetImportJobArns` | `List<string>` | no |
| `DataConfig` | `DataConfig` | no |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `ReferencePredictorSummary` | `ReferencePredictorSummary` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `OptimizationMetric` | `string` | no |
| `ExplainabilityInfo` | `ExplainabilityInfo` | no |
| `MonitorInfo` | `MonitorInfo` | no |
| `TimeAlignmentBoundary` | `TimeAlignmentBoundary` | no |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetArn` | `string` | no |
| `DatasetName` | `string` | no |
| `Domain` | `string` | no |
| `DatasetType` | `string` | no |
| `DataFrequency` | `string` | no |
| `Schema` | `Schema` | no |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |

## DescribeDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupName` | `string` | no |
| `DatasetGroupArn` | `string` | no |
| `DatasetArns` | `List<string>` | no |
| `Domain` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |

## DescribeDatasetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobName` | `string` | no |
| `DatasetImportJobArn` | `string` | no |
| `DatasetArn` | `string` | no |
| `TimestampFormat` | `string` | no |
| `TimeZone` | `string` | no |
| `UseGeolocationForTimeZone` | `boolean` | no |
| `GeolocationFormat` | `string` | no |
| `DataSource` | `DataSource` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `FieldStatistics` | `Map<Statistics>` | no |
| `DataSize` | `double` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `Format` | `string` | no |
| `ImportMode` | `string` | no |

## DescribeExplainability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityArn` | `string` | no |
| `ExplainabilityName` | `string` | no |
| `ResourceArn` | `string` | no |
| `ExplainabilityConfig` | `ExplainabilityConfig` | no |
| `EnableVisualization` | `boolean` | no |
| `DataSource` | `DataSource` | no |
| `Schema` | `Schema` | no |
| `StartDateTime` | `string` | no |
| `EndDateTime` | `string` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `Message` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |

## DescribeExplainabilityExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExportArn` | `string` | no |
| `ExplainabilityExportName` | `string` | no |
| `ExplainabilityArn` | `string` | no |
| `Destination` | `DataDestination` | no |
| `Message` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `Format` | `string` | no |

## DescribeForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastArn` | `string` | no |
| `ForecastName` | `string` | no |
| `ForecastTypes` | `List<string>` | no |
| `PredictorArn` | `string` | no |
| `DatasetGroupArn` | `string` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `TimeSeriesSelector` | `TimeSeriesSelector` | no |

## DescribeForecastExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobArn` | `string` | no |
| `ForecastExportJobName` | `string` | no |
| `ForecastArn` | `string` | no |
| `Destination` | `DataDestination` | no |
| `Message` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `Format` | `string` | no |

## DescribeMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | no |
| `MonitorArn` | `string` | no |
| `ResourceArn` | `string` | no |
| `Status` | `string` | no |
| `LastEvaluationTime` | `timestamp` | no |
| `LastEvaluationState` | `string` | no |
| `Baseline` | `Baseline` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `EstimatedEvaluationTimeRemainingInMinutes` | `long` | no |

## DescribePredictor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | no |
| `PredictorName` | `string` | no |
| `AlgorithmArn` | `string` | no |
| `AutoMLAlgorithmArns` | `List<string>` | no |
| `ForecastHorizon` | `integer` | no |
| `ForecastTypes` | `List<string>` | no |
| `PerformAutoML` | `boolean` | no |
| `AutoMLOverrideStrategy` | `string` | no |
| `PerformHPO` | `boolean` | no |
| `TrainingParameters` | `Map<string>` | no |
| `EvaluationParameters` | `EvaluationParameters` | no |
| `HPOConfig` | `HyperParameterTuningJobConfig` | no |
| `InputDataConfig` | `InputDataConfig` | no |
| `FeaturizationConfig` | `FeaturizationConfig` | no |
| `EncryptionConfig` | `EncryptionConfig` | no |
| `PredictorExecutionDetails` | `PredictorExecutionDetails` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `IsAutoPredictor` | `boolean` | no |
| `DatasetImportJobArns` | `List<string>` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `OptimizationMetric` | `string` | no |

## DescribePredictorBacktestExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobArn` | `string` | no |
| `PredictorBacktestExportJobName` | `string` | no |
| `PredictorArn` | `string` | no |
| `Destination` | `DataDestination` | no |
| `Message` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `Format` | `string` | no |

## DescribeWhatIfAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalysisArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalysisName` | `string` | no |
| `WhatIfAnalysisArn` | `string` | no |
| `ForecastArn` | `string` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `TimeSeriesSelector` | `TimeSeriesSelector` | no |

## DescribeWhatIfForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastName` | `string` | no |
| `WhatIfForecastArn` | `string` | no |
| `WhatIfAnalysisArn` | `string` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModificationTime` | `timestamp` | no |
| `TimeSeriesTransformations` | `List<TimeSeriesTransformation>` | no |
| `TimeSeriesReplacementsDataSource` | `TimeSeriesReplacementsDataSource` | no |
| `ForecastTypes` | `List<string>` | no |

## DescribeWhatIfForecastExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExportArn` | `string` | no |
| `WhatIfForecastExportName` | `string` | no |
| `WhatIfForecastArns` | `List<string>` | no |
| `Destination` | `DataDestination` | no |
| `Message` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `EstimatedTimeRemainingInMinutes` | `long` | no |
| `LastModificationTime` | `timestamp` | no |
| `Format` | `string` | no |

## GetAccuracyMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorEvaluationResults` | `List<EvaluationResult>` | no |
| `IsAutoPredictor` | `boolean` | no |
| `AutoMLOverrideStrategy` | `string` | no |
| `OptimizationMetric` | `string` | no |

## ListDatasetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroups` | `List<DatasetGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListDatasetImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetImportJobs` | `List<DatasetImportJobSummary>` | no |
| `NextToken` | `string` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Datasets` | `List<DatasetSummary>` | no |
| `NextToken` | `string` | no |

## ListExplainabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Explainabilities` | `List<ExplainabilitySummary>` | no |
| `NextToken` | `string` | no |

## ListExplainabilityExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplainabilityExports` | `List<ExplainabilityExportSummary>` | no |
| `NextToken` | `string` | no |

## ListForecastExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastExportJobs` | `List<ForecastExportJobSummary>` | no |
| `NextToken` | `string` | no |

## ListForecasts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Forecasts` | `List<ForecastSummary>` | no |
| `NextToken` | `string` | no |

## ListMonitorEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `MonitorArn` | `string` | yes |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PredictorMonitorEvaluations` | `List<PredictorMonitorEvaluation>` | no |

## ListMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Monitors` | `List<MonitorSummary>` | no |
| `NextToken` | `string` | no |

## ListPredictorBacktestExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredictorBacktestExportJobs` | `List<PredictorBacktestExportJobSummary>` | no |
| `NextToken` | `string` | no |

## ListPredictors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Predictors` | `List<PredictorSummary>` | no |
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

## ListWhatIfAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfAnalyses` | `List<WhatIfAnalysisSummary>` | no |
| `NextToken` | `string` | no |

## ListWhatIfForecastExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastExports` | `List<WhatIfForecastExportSummary>` | no |
| `NextToken` | `string` | no |

## ListWhatIfForecasts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecasts` | `List<WhatIfForecastSummary>` | no |
| `NextToken` | `string` | no |

## ResumeResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetGroupArn` | `string` | yes |
| `DatasetArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


