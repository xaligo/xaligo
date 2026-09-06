# TrustedAdvisor Public API

API version: 2022-09-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/trustedadvisor/2022-09-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchUpdateRecommendationResourceExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationResourceExclusions` | `List<RecommendationResourceExclusion>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchUpdateRecommendationResourceExclusionErrors` | `List<UpdateRecommendationResourceExclusionError>` | yes |

## GetOrganizationRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationRecommendationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationRecommendation` | `OrganizationRecommendation` | no |

## GetRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationIdentifier` | `string` | yes |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendation` | `Recommendation` | no |

## ListChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `pillar` | `string` | no |
| `awsService` | `string` | no |
| `source` | `string` | no |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `checkSummaries` | `List<CheckSummary>` | yes |

## ListOrganizationRecommendationAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `organizationRecommendationIdentifier` | `string` | yes |
| `affectedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `accountRecommendationLifecycleSummaries` | `List<AccountRecommendationLifecycleSummary>` | yes |

## ListOrganizationRecommendationResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |
| `exclusionStatus` | `string` | no |
| `regionCode` | `string` | no |
| `organizationRecommendationIdentifier` | `string` | yes |
| `affectedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `organizationRecommendationResourceSummaries` | `List<OrganizationRecommendationResourceSummary>` | yes |

## ListOrganizationRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `pillar` | `string` | no |
| `awsService` | `string` | no |
| `source` | `string` | no |
| `checkIdentifier` | `string` | no |
| `afterLastUpdatedAt` | `timestamp` | no |
| `beforeLastUpdatedAt` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `organizationRecommendationSummaries` | `List<OrganizationRecommendationSummary>` | yes |

## ListRecommendationResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |
| `exclusionStatus` | `string` | no |
| `regionCode` | `string` | no |
| `recommendationIdentifier` | `string` | yes |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationResourceSummaries` | `List<RecommendationResourceSummary>` | yes |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `pillar` | `string` | no |
| `awsService` | `string` | no |
| `source` | `string` | no |
| `checkIdentifier` | `string` | no |
| `afterLastUpdatedAt` | `timestamp` | no |
| `beforeLastUpdatedAt` | `timestamp` | no |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationSummaries` | `List<RecommendationSummary>` | yes |

## ListRecommendationsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `awsResourceArn` | `string` | yes |
| `pillar` | `string` | no |
| `status` | `string` | no |
| `checkArn` | `string` | no |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `recommendationForResourceSummaries` | `List<RecommendationForResourceSummary>` | yes |

## UpdateOrganizationRecommendationLifecycle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleStage` | `string` | yes |
| `updateReason` | `string` | no |
| `updateReasonCode` | `string` | no |
| `organizationRecommendationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRecommendationLifecycle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleStage` | `string` | yes |
| `updateReason` | `string` | no |
| `updateReasonCode` | `string` | no |
| `recommendationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


