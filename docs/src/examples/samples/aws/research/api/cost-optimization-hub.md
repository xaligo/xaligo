# Cost Optimization Hub

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cost-optimization-hub/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsEstimationMode` | `string` | no |
| `memberAccountDiscountVisibility` | `string` | no |
| `preferredCommitment` | `PreferredCommitment` | no |

## GetRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | no |
| `resourceId` | `string` | no |
| `resourceArn` | `string` | no |
| `accountId` | `string` | no |
| `currencyCode` | `string` | no |
| `recommendationLookbackPeriodInDays` | `integer` | no |
| `costCalculationLookbackPeriodInDays` | `integer` | no |
| `estimatedSavingsPercentage` | `double` | no |
| `estimatedSavingsOverCostCalculationLookbackPeriod` | `double` | no |
| `currentResourceType` | `string` | no |
| `recommendedResourceType` | `string` | no |
| `region` | `string` | no |
| `source` | `string` | no |
| `lastRefreshTimestamp` | `timestamp` | no |
| `estimatedMonthlySavings` | `double` | no |
| `estimatedMonthlyCost` | `double` | no |
| `implementationEffort` | `string` | no |
| `restartNeeded` | `boolean` | no |
| `actionType` | `string` | no |
| `rollbackPossible` | `boolean` | no |
| `currentResourceDetails` | `ResourceDetails` | no |
| `recommendedResourceDetails` | `ResourceDetails` | no |
| `tags` | `List<Tag>` | no |

## ListEfficiencyMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `groupBy` | `string` | no |
| `granularity` | `string` | yes |
| `timePeriod` | `TimePeriod` | yes |
| `maxResults` | `integer` | no |
| `orderBy` | `OrderBy` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `efficiencyMetricsByGroup` | `List<EfficiencyMetricsByGroup>` | no |
| `nextToken` | `string` | no |

## ListEnrollmentStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeOrganizationInfo` | `boolean` | no |
| `accountId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AccountEnrollmentStatus>` | no |
| `includeMemberAccounts` | `boolean` | no |
| `nextToken` | `string` | no |

## ListRecommendationSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `Filter` | no |
| `groupBy` | `string` | yes |
| `maxResults` | `integer` | no |
| `metrics` | `List<string>` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `estimatedTotalDedupedSavings` | `double` | no |
| `items` | `List<RecommendationSummary>` | no |
| `groupBy` | `string` | no |
| `currencyCode` | `string` | no |
| `metrics` | `SummaryMetricsResult` | no |
| `nextToken` | `string` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `Filter` | no |
| `orderBy` | `OrderBy` | no |
| `includeAllRecommendations` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Recommendation>` | no |
| `nextToken` | `string` | no |

## UpdateEnrollmentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `includeMemberAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## UpdatePreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsEstimationMode` | `string` | no |
| `memberAccountDiscountVisibility` | `string` | no |
| `preferredCommitment` | `PreferredCommitment` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsEstimationMode` | `string` | no |
| `memberAccountDiscountVisibility` | `string` | no |
| `preferredCommitment` | `PreferredCommitment` | no |

