# AWS Cost Explorer Service

API version: 2017-10-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ce/2017-10-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAnomalyMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyMonitor` | `AnomalyMonitor` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |

## CreateAnomalySubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalySubscription` | `AnomalySubscription` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |

## CreateCostCategoryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EffectiveStart` | `string` | no |
| `RuleVersion` | `string` | yes |
| `Rules` | `List<CostCategoryRule>` | yes |
| `DefaultValue` | `string` | no |
| `SplitChargeRules` | `List<CostCategorySplitChargeRule>` | no |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | no |
| `EffectiveStart` | `string` | no |

## DeleteAnomalyMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAnomalySubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCostCategoryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | no |
| `EffectiveEnd` | `string` | no |

## DescribeCostCategoryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | yes |
| `EffectiveOn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategory` | `CostCategory` | no |

## GetAnomalies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | no |
| `DateInterval` | `AnomalyDateInterval` | yes |
| `Feedback` | `string` | no |
| `TotalImpact` | `TotalImpactFilter` | no |
| `NextPageToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Anomalies` | `List<Anomaly>` | yes |
| `NextPageToken` | `string` | no |

## GetAnomalyMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArnList` | `List<string>` | no |
| `NextPageToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyMonitors` | `List<AnomalyMonitor>` | yes |
| `NextPageToken` | `string` | no |

## GetAnomalySubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArnList` | `List<string>` | no |
| `MonitorArn` | `string` | no |
| `NextPageToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalySubscriptions` | `List<AnomalySubscription>` | yes |
| `NextPageToken` | `string` | no |

## GetApproximateUsageRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Granularity` | `string` | yes |
| `Services` | `List<string>` | no |
| `ApproximationDimension` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Services` | `Map<long>` | no |
| `TotalRecords` | `long` | no |
| `LookbackPeriod` | `DateInterval` | no |

## GetCommitmentPurchaseAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EstimatedCompletionTime` | `string` | yes |
| `AnalysisCompletionTime` | `string` | no |
| `AnalysisStartedTime` | `string` | yes |
| `AnalysisId` | `string` | yes |
| `AnalysisStatus` | `string` | yes |
| `ErrorCode` | `string` | no |
| `AnalysisDetails` | `AnalysisDetails` | no |
| `CommitmentPurchaseAnalysisConfiguration` | `CommitmentPurchaseAnalysisConfiguration` | yes |

## GetCostAndUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Granularity` | `string` | yes |
| `Filter` | `Expression` | no |
| `Metrics` | `List<string>` | yes |
| `GroupBy` | `List<GroupDefinition>` | no |
| `BillingViewArn` | `string` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageToken` | `string` | no |
| `GroupDefinitions` | `List<GroupDefinition>` | no |
| `ResultsByTime` | `List<ResultByTime>` | no |
| `DimensionValueAttributes` | `List<DimensionValuesWithAttributes>` | no |

## GetCostAndUsageComparisons

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingViewArn` | `string` | no |
| `BaselineTimePeriod` | `DateInterval` | yes |
| `ComparisonTimePeriod` | `DateInterval` | yes |
| `MetricForComparison` | `string` | yes |
| `Filter` | `Expression` | no |
| `GroupBy` | `List<GroupDefinition>` | no |
| `MaxResults` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostAndUsageComparisons` | `List<CostAndUsageComparison>` | no |
| `TotalCostAndUsage` | `Map<ComparisonMetricValue>` | no |
| `NextPageToken` | `string` | no |

## GetCostAndUsageWithResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Granularity` | `string` | yes |
| `Filter` | `Expression` | yes |
| `Metrics` | `List<string>` | no |
| `GroupBy` | `List<GroupDefinition>` | no |
| `BillingViewArn` | `string` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageToken` | `string` | no |
| `GroupDefinitions` | `List<GroupDefinition>` | no |
| `ResultsByTime` | `List<ResultByTime>` | no |
| `DimensionValueAttributes` | `List<DimensionValuesWithAttributes>` | no |

## GetCostCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchString` | `string` | no |
| `TimePeriod` | `DateInterval` | yes |
| `CostCategoryName` | `string` | no |
| `Filter` | `Expression` | no |
| `SortBy` | `List<SortDefinition>` | no |
| `BillingViewArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageToken` | `string` | no |
| `CostCategoryNames` | `List<string>` | no |
| `CostCategoryValues` | `List<string>` | no |
| `ReturnSize` | `integer` | yes |
| `TotalSize` | `integer` | yes |

## GetCostComparisonDrivers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingViewArn` | `string` | no |
| `BaselineTimePeriod` | `DateInterval` | yes |
| `ComparisonTimePeriod` | `DateInterval` | yes |
| `MetricForComparison` | `string` | yes |
| `Filter` | `Expression` | no |
| `GroupBy` | `List<GroupDefinition>` | no |
| `MaxResults` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostComparisonDrivers` | `List<CostComparisonDriver>` | no |
| `NextPageToken` | `string` | no |

## GetCostForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Metric` | `string` | yes |
| `Granularity` | `string` | yes |
| `Filter` | `Expression` | no |
| `BillingViewArn` | `string` | no |
| `PredictionIntervalLevel` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Total` | `MetricValue` | no |
| `ForecastResultsByTime` | `List<ForecastResult>` | no |

## GetDimensionValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchString` | `string` | no |
| `TimePeriod` | `DateInterval` | yes |
| `Dimension` | `string` | yes |
| `Context` | `string` | no |
| `Filter` | `Expression` | no |
| `SortBy` | `List<SortDefinition>` | no |
| `BillingViewArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DimensionValues` | `List<DimensionValuesWithAttributes>` | yes |
| `ReturnSize` | `integer` | yes |
| `TotalSize` | `integer` | yes |
| `NextPageToken` | `string` | no |

## GetReservationCoverage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `GroupBy` | `List<GroupDefinition>` | no |
| `Granularity` | `string` | no |
| `Filter` | `Expression` | no |
| `Metrics` | `List<string>` | no |
| `NextPageToken` | `string` | no |
| `SortBy` | `SortDefinition` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoveragesByTime` | `List<CoverageByTime>` | yes |
| `Total` | `Coverage` | no |
| `NextPageToken` | `string` | no |

## GetReservationPurchaseRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `Service` | `string` | yes |
| `Filter` | `Expression` | no |
| `AccountScope` | `string` | no |
| `LookbackPeriodInDays` | `string` | no |
| `TermInYears` | `string` | no |
| `PaymentOption` | `string` | no |
| `ServiceSpecification` | `ServiceSpecification` | no |
| `PageSize` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `ReservationPurchaseRecommendationMetadata` | no |
| `Recommendations` | `List<ReservationPurchaseRecommendation>` | no |
| `NextPageToken` | `string` | no |

## GetReservationUtilization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `GroupBy` | `List<GroupDefinition>` | no |
| `Granularity` | `string` | no |
| `Filter` | `Expression` | no |
| `SortBy` | `SortDefinition` | no |
| `NextPageToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UtilizationsByTime` | `List<UtilizationByTime>` | yes |
| `Total` | `ReservationAggregates` | no |
| `NextPageToken` | `string` | no |

## GetRightsizingRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `Expression` | no |
| `Configuration` | `RightsizingRecommendationConfiguration` | no |
| `Service` | `string` | yes |
| `PageSize` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `RightsizingRecommendationMetadata` | no |
| `Summary` | `RightsizingRecommendationSummary` | no |
| `RightsizingRecommendations` | `List<RightsizingRecommendation>` | no |
| `NextPageToken` | `string` | no |
| `Configuration` | `RightsizingRecommendationConfiguration` | no |

## GetSavingsPlanPurchaseRecommendationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationDetailId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationDetailId` | `string` | no |
| `RecommendationDetailData` | `RecommendationDetailData` | no |

## GetSavingsPlansCoverage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `GroupBy` | `List<GroupDefinition>` | no |
| `Granularity` | `string` | no |
| `Filter` | `Expression` | no |
| `Metrics` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortBy` | `SortDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SavingsPlansCoverages` | `List<SavingsPlansCoverage>` | yes |
| `NextToken` | `string` | no |

## GetSavingsPlansPurchaseRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SavingsPlansType` | `string` | yes |
| `TermInYears` | `string` | yes |
| `PaymentOption` | `string` | yes |
| `AccountScope` | `string` | no |
| `NextPageToken` | `string` | no |
| `PageSize` | `integer` | no |
| `LookbackPeriodInDays` | `string` | yes |
| `Filter` | `Expression` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `SavingsPlansPurchaseRecommendationMetadata` | no |
| `SavingsPlansPurchaseRecommendation` | `SavingsPlansPurchaseRecommendation` | no |
| `NextPageToken` | `string` | no |

## GetSavingsPlansUtilization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Granularity` | `string` | no |
| `Filter` | `Expression` | no |
| `SortBy` | `SortDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SavingsPlansUtilizationsByTime` | `List<SavingsPlansUtilizationByTime>` | no |
| `Total` | `SavingsPlansUtilizationAggregates` | yes |

## GetSavingsPlansUtilizationDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Filter` | `Expression` | no |
| `DataType` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortBy` | `SortDefinition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SavingsPlansUtilizationDetails` | `List<SavingsPlansUtilizationDetail>` | yes |
| `Total` | `SavingsPlansUtilizationAggregates` | no |
| `TimePeriod` | `DateInterval` | yes |
| `NextToken` | `string` | no |

## GetTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchString` | `string` | no |
| `TimePeriod` | `DateInterval` | yes |
| `TagKey` | `string` | no |
| `Filter` | `Expression` | no |
| `SortBy` | `List<SortDefinition>` | no |
| `BillingViewArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageToken` | `string` | no |
| `Tags` | `List<string>` | yes |
| `ReturnSize` | `integer` | yes |
| `TotalSize` | `integer` | yes |

## GetUsageForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `DateInterval` | yes |
| `Metric` | `string` | yes |
| `Granularity` | `string` | yes |
| `Filter` | `Expression` | no |
| `BillingViewArn` | `string` | no |
| `PredictionIntervalLevel` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Total` | `MetricValue` | no |
| `ForecastResultsByTime` | `List<ForecastResult>` | no |

## ListCommitmentPurchaseAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisStatus` | `string` | no |
| `NextPageToken` | `string` | no |
| `PageSize` | `integer` | no |
| `AnalysisIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisSummaryList` | `List<AnalysisSummary>` | no |
| `NextPageToken` | `string` | no |

## ListCostAllocationTagBackfillHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackfillRequests` | `List<CostAllocationTagBackfillRequest>` | no |
| `NextToken` | `string` | no |

## ListCostAllocationTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `TagKeys` | `List<string>` | no |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostAllocationTags` | `List<CostAllocationTag>` | no |
| `NextToken` | `string` | no |

## ListCostCategoryDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EffectiveOn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SupportedResourceTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryReferences` | `List<CostCategoryReference>` | no |
| `NextToken` | `string` | no |

## ListCostCategoryResourceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryResourceAssociations` | `List<CostCategoryResourceAssociation>` | no |
| `NextToken` | `string` | no |

## ListSavingsPlansPurchaseRecommendationGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GenerationStatus` | `string` | no |
| `RecommendationIds` | `List<string>` | no |
| `PageSize` | `integer` | no |
| `NextPageToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GenerationSummaryList` | `List<GenerationSummary>` | no |
| `NextPageToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTags` | `List<ResourceTag>` | no |

## ProvideAnomalyFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyId` | `string` | yes |
| `Feedback` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyId` | `string` | yes |

## StartCommitmentPurchaseAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CommitmentPurchaseAnalysisConfiguration` | `CommitmentPurchaseAnalysisConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisId` | `string` | yes |
| `AnalysisStartedTime` | `string` | yes |
| `EstimatedCompletionTime` | `string` | yes |

## StartCostAllocationTagBackfill

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackfillFrom` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackfillRequest` | `CostAllocationTagBackfillRequest` | no |

## StartSavingsPlansPurchaseRecommendationGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationId` | `string` | no |
| `GenerationStartedTime` | `string` | no |
| `EstimatedCompletionTime` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAnomalyMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |
| `MonitorName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |

## UpdateAnomalySubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |
| `Threshold` | `double` | no |
| `Frequency` | `string` | no |
| `MonitorArnList` | `List<string>` | no |
| `Subscribers` | `List<Subscriber>` | no |
| `SubscriptionName` | `string` | no |
| `ThresholdExpression` | `Expression` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |

## UpdateCostAllocationTagsStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostAllocationTagsStatus` | `List<CostAllocationTagStatusEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<UpdateCostAllocationTagsStatusError>` | no |

## UpdateCostCategoryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | yes |
| `EffectiveStart` | `string` | no |
| `RuleVersion` | `string` | yes |
| `Rules` | `List<CostCategoryRule>` | yes |
| `DefaultValue` | `string` | no |
| `SplitChargeRules` | `List<CostCategorySplitChargeRule>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CostCategoryArn` | `string` | no |
| `EffectiveStart` | `string` | no |

