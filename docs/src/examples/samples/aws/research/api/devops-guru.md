# Amazon DevOps Guru

API version: 2020-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/devops-guru/2020-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Config` | `NotificationChannelConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

## DeleteInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenReactiveInsights` | `integer` | yes |
| `OpenProactiveInsights` | `integer` | yes |
| `MetricsAnalyzed` | `integer` | yes |
| `ResourceHours` | `long` | yes |
| `AnalyzedResourceCount` | `long` | no |

## DescribeAccountOverview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromTime` | `timestamp` | yes |
| `ToTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReactiveInsights` | `integer` | yes |
| `ProactiveInsights` | `integer` | yes |
| `MeanTimeToRecoverInMilliseconds` | `long` | yes |

## DescribeAnomaly

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveAnomaly` | `ProactiveAnomaly` | no |
| `ReactiveAnomaly` | `ReactiveAnomaly` | no |

## DescribeEventSourcesConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSources` | `EventSourcesConfig` | no |

## DescribeFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightFeedback` | `InsightFeedback` | no |

## DescribeInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveInsight` | `ProactiveInsight` | no |
| `ReactiveInsight` | `ReactiveInsight` | no |

## DescribeOrganizationHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `OrganizationalUnitIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OpenReactiveInsights` | `integer` | yes |
| `OpenProactiveInsights` | `integer` | yes |
| `MetricsAnalyzed` | `integer` | yes |
| `ResourceHours` | `long` | yes |

## DescribeOrganizationOverview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromTime` | `timestamp` | yes |
| `ToTime` | `timestamp` | no |
| `AccountIds` | `List<string>` | no |
| `OrganizationalUnitIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReactiveInsights` | `integer` | yes |
| `ProactiveInsights` | `integer` | yes |

## DescribeOrganizationResourceCollectionHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationResourceCollectionType` | `string` | yes |
| `AccountIds` | `List<string>` | no |
| `OrganizationalUnitIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFormation` | `List<CloudFormationHealth>` | no |
| `Service` | `List<ServiceHealth>` | no |
| `Account` | `List<AccountHealth>` | no |
| `NextToken` | `string` | no |
| `Tags` | `List<TagHealth>` | no |

## DescribeResourceCollectionHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCollectionType` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFormation` | `List<CloudFormationHealth>` | no |
| `Service` | `List<ServiceHealth>` | no |
| `NextToken` | `string` | no |
| `Tags` | `List<TagHealth>` | no |

## DescribeServiceIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceIntegration` | `ServiceIntegrationConfig` | no |

## GetCostEstimation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCollection` | `CostEstimationResourceCollectionFilter` | no |
| `Status` | `string` | no |
| `Costs` | `List<ServiceResourceCost>` | no |
| `TimeRange` | `CostEstimationTimeRange` | no |
| `TotalCost` | `double` | no |
| `NextToken` | `string` | no |

## GetResourceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCollectionType` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCollection` | `ResourceCollectionFilter` | no |
| `NextToken` | `string` | no |

## ListAnomaliesForInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `StartTimeRange` | `StartTimeRange` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |
| `Filters` | `ListAnomaliesForInsightFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveAnomalies` | `List<ProactiveAnomalySummary>` | no |
| `ReactiveAnomalies` | `List<ReactiveAnomalySummary>` | no |
| `NextToken` | `string` | no |

## ListAnomalousLogGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `AnomalousLogGroups` | `List<AnomalousLogGroup>` | yes |
| `NextToken` | `string` | no |

## ListEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ListEventsFilters` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<Event>` | yes |
| `NextToken` | `string` | no |

## ListInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusFilter` | `ListInsightsStatusFilter` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveInsights` | `List<ProactiveInsightSummary>` | no |
| `ReactiveInsights` | `List<ReactiveInsightSummary>` | no |
| `NextToken` | `string` | no |

## ListMonitoredResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ListMonitoredResourcesFilters` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoredResourceIdentifiers` | `List<MonitoredResourceIdentifier>` | yes |
| `NextToken` | `string` | no |

## ListNotificationChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<NotificationChannel>` | no |
| `NextToken` | `string` | no |

## ListOrganizationInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusFilter` | `ListInsightsStatusFilter` | yes |
| `MaxResults` | `integer` | no |
| `AccountIds` | `List<string>` | no |
| `OrganizationalUnitIds` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveInsights` | `List<ProactiveOrganizationInsightSummary>` | no |
| `ReactiveInsights` | `List<ReactiveOrganizationInsightSummary>` | no |
| `NextToken` | `string` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `NextToken` | `string` | no |
| `Locale` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recommendations` | `List<Recommendation>` | no |
| `NextToken` | `string` | no |

## PutFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightFeedback` | `InsightFeedback` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTimeRange` | `StartTimeRange` | yes |
| `Filters` | `SearchInsightsFilters` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveInsights` | `List<ProactiveInsightSummary>` | no |
| `ReactiveInsights` | `List<ReactiveInsightSummary>` | no |
| `NextToken` | `string` | no |

## SearchOrganizationInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |
| `StartTimeRange` | `StartTimeRange` | yes |
| `Filters` | `SearchOrganizationInsightsFilters` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProactiveInsights` | `List<ProactiveInsightSummary>` | no |
| `ReactiveInsights` | `List<ReactiveInsightSummary>` | no |
| `NextToken` | `string` | no |

## StartCostEstimation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCollection` | `CostEstimationResourceCollectionFilter` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventSourcesConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSources` | `EventSourcesConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |
| `ResourceCollection` | `UpdateResourceCollectionFilter` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceIntegration` | `UpdateServiceIntegrationConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


