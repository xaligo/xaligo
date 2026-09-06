# AWS Performance Insights

API version: 2018-02-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pi/2018-02-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreatePerformanceAnalysisReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisReportId` | `string` | no |

## DeletePerformanceAnalysisReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `AnalysisReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDimensionKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Metric` | `string` | yes |
| `PeriodInSeconds` | `integer` | no |
| `GroupBy` | `DimensionGroup` | yes |
| `AdditionalMetrics` | `List<string>` | no |
| `PartitionBy` | `DimensionGroup` | no |
| `Filter` | `Map<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlignedStartTime` | `timestamp` | no |
| `AlignedEndTime` | `timestamp` | no |
| `PartitionKeys` | `List<ResponsePartitionKey>` | no |
| `Keys` | `List<DimensionKeyDescription>` | no |
| `NextToken` | `string` | no |

## GetDimensionKeyDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `Group` | `string` | yes |
| `GroupIdentifier` | `string` | yes |
| `RequestedDimensions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dimensions` | `List<DimensionKeyDetail>` | no |

## GetPerformanceAnalysisReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `AnalysisReportId` | `string` | yes |
| `TextFormat` | `string` | no |
| `AcceptLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisReport` | `AnalysisReport` | no |

## GetResourceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `Features` | `Map<FeatureMetadata>` | no |

## GetResourceMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `MetricQueries` | `List<MetricQuery>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `PeriodInSeconds` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PeriodAlignment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlignedStartTime` | `timestamp` | no |
| `AlignedEndTime` | `timestamp` | no |
| `Identifier` | `string` | no |
| `MetricList` | `List<MetricKeyDataPoints>` | no |
| `NextToken` | `string` | no |

## ListAvailableResourceDimensions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `Metrics` | `List<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AuthorizedActions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDimensions` | `List<MetricDimensionGroups>` | no |
| `NextToken` | `string` | no |

## ListAvailableResourceMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `MetricTypes` | `List<string>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metrics` | `List<ResponseResourceMetric>` | no |
| `NextToken` | `string` | no |

## ListPerformanceAnalysisReportRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `AnalysisReportId` | `string` | yes |
| `RecommendationIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recommendations` | `List<Recommendation>` | no |
| `NextToken` | `string` | no |

## ListPerformanceAnalysisReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `Identifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ListTags` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnalysisReports` | `List<AnalysisReportSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | yes |
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


