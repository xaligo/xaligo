# Amazon Personalize

API version: 2018-05-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/personalize/2018-05-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBatchInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `solutionVersionArn` | `string` | yes |
| `filterArn` | `string` | no |
| `numResults` | `integer` | no |
| `jobInput` | `BatchInferenceJobInput` | yes |
| `jobOutput` | `BatchInferenceJobOutput` | yes |
| `roleArn` | `string` | yes |
| `batchInferenceJobConfig` | `BatchInferenceJobConfig` | no |
| `tags` | `List<Tag>` | no |
| `batchInferenceJobMode` | `string` | no |
| `themeGenerationConfig` | `ThemeGenerationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchInferenceJobArn` | `string` | no |

## CreateBatchSegmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `solutionVersionArn` | `string` | yes |
| `filterArn` | `string` | no |
| `numResults` | `integer` | no |
| `jobInput` | `BatchSegmentJobInput` | yes |
| `jobOutput` | `BatchSegmentJobOutput` | yes |
| `roleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchSegmentJobArn` | `string` | no |

## CreateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `solutionVersionArn` | `string` | yes |
| `minProvisionedTPS` | `integer` | no |
| `campaignConfig` | `CampaignConfig` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | no |

## CreateDataDeletionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `dataSource` | `DataSource` | yes |
| `roleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataDeletionJobArn` | `string` | no |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `schemaArn` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `datasetType` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | no |

## CreateDatasetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `datasetArn` | `string` | yes |
| `ingestionMode` | `string` | no |
| `roleArn` | `string` | yes |
| `jobOutput` | `DatasetExportJobOutput` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetExportJobArn` | `string` | no |

## CreateDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `roleArn` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `domain` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `domain` | `string` | no |

## CreateDatasetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `datasetArn` | `string` | yes |
| `dataSource` | `DataSource` | yes |
| `roleArn` | `string` | no |
| `tags` | `List<Tag>` | no |
| `importMode` | `string` | no |
| `publishAttributionMetricsToS3` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetImportJobArn` | `string` | no |

## CreateEventTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTrackerArn` | `string` | no |
| `trackingId` | `string` | no |

## CreateFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `filterExpression` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterArn` | `string` | no |

## CreateMetricAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `metrics` | `List<MetricAttribute>` | yes |
| `metricsOutputConfig` | `MetricAttributionOutput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributionArn` | `string` | no |

## CreateRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `datasetGroupArn` | `string` | yes |
| `recipeArn` | `string` | yes |
| `recommenderConfig` | `RecommenderConfig` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | no |

## CreateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `schema` | `string` | yes |
| `domain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaArn` | `string` | no |

## CreateSolution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `performHPO` | `boolean` | no |
| `performAutoML` | `boolean` | no |
| `performAutoTraining` | `boolean` | no |
| `performIncrementalUpdate` | `boolean` | no |
| `recipeArn` | `string` | no |
| `datasetGroupArn` | `string` | yes |
| `eventType` | `string` | no |
| `solutionConfig` | `SolutionConfig` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | no |

## CreateSolutionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `solutionArn` | `string` | yes |
| `trainingMode` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | no |

## DeleteCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTrackerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMetricAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSolution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `algorithmArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `algorithm` | `Algorithm` | no |

## DescribeBatchInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchInferenceJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchInferenceJob` | `BatchInferenceJob` | no |

## DescribeBatchSegmentJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchSegmentJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchSegmentJob` | `BatchSegmentJob` | no |

## DescribeCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaign` | `Campaign` | no |

## DescribeDataDeletionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataDeletionJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataDeletionJob` | `DataDeletionJob` | no |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataset` | `Dataset` | no |

## DescribeDatasetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetExportJob` | `DatasetExportJob` | no |

## DescribeDatasetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroup` | `DatasetGroup` | no |

## DescribeDatasetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetImportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetImportJob` | `DatasetImportJob` | no |

## DescribeEventTracker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTrackerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTracker` | `EventTracker` | no |

## DescribeFeatureTransformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `featureTransformationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `featureTransformation` | `FeatureTransformation` | no |

## DescribeFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `Filter` | no |

## DescribeMetricAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttribution` | `MetricAttribution` | no |

## DescribeRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipe` | `Recipe` | no |

## DescribeRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommender` | `Recommender` | no |

## DescribeSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schema` | `DatasetSchema` | no |

## DescribeSolution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solution` | `Solution` | no |

## DescribeSolutionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersion` | `SolutionVersion` | no |

## GetSolutionMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | no |
| `metrics` | `Map<double>` | no |

## ListBatchInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchInferenceJobs` | `List<BatchInferenceJobSummary>` | no |
| `nextToken` | `string` | no |

## ListBatchSegmentJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchSegmentJobs` | `List<BatchSegmentJobSummary>` | no |
| `nextToken` | `string` | no |

## ListCampaigns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaigns` | `List<CampaignSummary>` | no |
| `nextToken` | `string` | no |

## ListDataDeletionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataDeletionJobs` | `List<DataDeletionJobSummary>` | no |
| `nextToken` | `string` | no |

## ListDatasetExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetExportJobs` | `List<DatasetExportJobSummary>` | no |
| `nextToken` | `string` | no |

## ListDatasetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroups` | `List<DatasetGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListDatasetImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetImportJobs` | `List<DatasetImportJobSummary>` | no |
| `nextToken` | `string` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasets` | `List<DatasetSummary>` | no |
| `nextToken` | `string` | no |

## ListEventTrackers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTrackers` | `List<EventTrackerSummary>` | no |
| `nextToken` | `string` | no |

## ListFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<FilterSummary>` | no |
| `nextToken` | `string` | no |

## ListMetricAttributionMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributionArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metrics` | `List<MetricAttribute>` | no |
| `nextToken` | `string` | no |

## ListMetricAttributions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributions` | `List<MetricAttributionSummary>` | no |
| `nextToken` | `string` | no |

## ListRecipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipeProvider` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `domain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recipes` | `List<RecipeSummary>` | no |
| `nextToken` | `string` | no |

## ListRecommenders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenders` | `List<RecommenderSummary>` | no |
| `nextToken` | `string` | no |

## ListSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemas` | `List<DatasetSchemaSummary>` | no |
| `nextToken` | `string` | no |

## ListSolutionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersions` | `List<SolutionVersionSummary>` | no |
| `nextToken` | `string` | no |

## ListSolutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetGroupArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutions` | `List<SolutionSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## StartRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | no |

## StopRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | no |

## StopSolutionVersionCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | yes |
| `solutionVersionArn` | `string` | no |
| `minProvisionedTPS` | `integer` | no |
| `campaignConfig` | `CampaignConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | no |

## UpdateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `schemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | no |

## UpdateMetricAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addMetrics` | `List<MetricAttribute>` | no |
| `removeMetrics` | `List<string>` | no |
| `metricsOutputConfig` | `MetricAttributionOutput` | no |
| `metricAttributionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricAttributionArn` | `string` | no |

## UpdateRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | yes |
| `recommenderConfig` | `RecommenderConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommenderArn` | `string` | no |

## UpdateSolution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | yes |
| `performAutoTraining` | `boolean` | no |
| `performIncrementalUpdate` | `boolean` | no |
| `solutionUpdateConfig` | `SolutionUpdateConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `solutionArn` | `string` | no |

