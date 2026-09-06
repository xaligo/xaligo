# Access Analyzer

API version: 2019-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/accessanalyzer/2019-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ApplyArchiveRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `ruleName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelPolicyGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckAccessNotGranted

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyDocument` | `string` | yes |
| `access` | `List<Access>` | yes |
| `policyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `string` | no |
| `message` | `string` | no |
| `reasons` | `List<ReasonSummary>` | no |

## CheckNoNewAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `newPolicyDocument` | `string` | yes |
| `existingPolicyDocument` | `string` | yes |
| `policyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `string` | no |
| `message` | `string` | no |
| `reasons` | `List<ReasonSummary>` | no |

## CheckNoPublicAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyDocument` | `string` | yes |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `string` | no |
| `message` | `string` | no |
| `reasons` | `List<ReasonSummary>` | no |

## CreateAccessPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `configurations` | `Map<Configuration>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

## CreateAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `type` | `string` | yes |
| `archiveRules` | `List<InlineArchiveRule>` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |
| `configuration` | `AnalyzerConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## CreateArchiveRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `ruleName` | `string` | yes |
| `filter` | `Map<Criterion>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateServiceLinkedAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `archiveRules` | `List<InlineArchiveRule>` | no |
| `clientToken` | `string` | no |
| `configuration` | `AnalyzerConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## DeleteAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteArchiveRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `ruleName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceLinkedAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GenerateFindingRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPreviewId` | `string` | yes |
| `analyzerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPreview` | `AccessPreview` | yes |

## GetAnalyzedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resource` | `AnalyzedResource` | no |

## GetAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzer` | `AnalyzerSummary` | yes |

## GetArchiveRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `ruleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `archiveRule` | `ArchiveRuleSummary` | yes |

## GetFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `finding` | `Finding` | no |

## GetFindingRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `id` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startedAt` | `timestamp` | yes |
| `completedAt` | `timestamp` | no |
| `nextToken` | `string` | no |
| `error` | `RecommendationError` | no |
| `resourceArn` | `string` | yes |
| `recommendedSteps` | `List<RecommendedStep>` | no |
| `recommendationType` | `string` | yes |
| `status` | `string` | yes |

## GetFindingV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `id` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzedAt` | `timestamp` | yes |
| `createdAt` | `timestamp` | yes |
| `error` | `string` | no |
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `resource` | `string` | no |
| `resourceType` | `string` | yes |
| `resourceOwnerAccount` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `findingDetails` | `List<FindingDetails>` | yes |
| `findingType` | `string` | no |

## GetFindingsStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsStatistics` | `List<FindingsStatistics>` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetGeneratedPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `includeResourcePlaceholders` | `boolean` | no |
| `includeServiceLevelTemplate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDetails` | `JobDetails` | yes |
| `generatedPolicyResult` | `GeneratedPolicyResult` | yes |

## ListAccessPreviewFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPreviewId` | `string` | yes |
| `analyzerArn` | `string` | yes |
| `filter` | `Map<Criterion>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<AccessPreviewFinding>` | yes |
| `nextToken` | `string` | no |

## ListAccessPreviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPreviews` | `List<AccessPreviewSummary>` | yes |
| `nextToken` | `string` | no |

## ListAnalyzedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `resourceType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzedResources` | `List<AnalyzedResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListAnalyzers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzers` | `List<AnalyzerSummary>` | yes |
| `nextToken` | `string` | no |

## ListArchiveRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `archiveRules` | `List<ArchiveRuleSummary>` | yes |
| `nextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `filter` | `Map<Criterion>` | no |
| `sort` | `SortCriteria` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<FindingSummary>` | yes |
| `nextToken` | `string` | no |

## ListFindingsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `filter` | `Map<Criterion>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sort` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<FindingSummaryV2>` | yes |
| `nextToken` | `string` | no |

## ListPolicyGenerations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerations` | `List<PolicyGeneration>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## StartPolicyGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerationDetails` | `PolicyGenerationDetails` | yes |
| `cloudTrailDetails` | `CloudTrailDetails` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

## StartResourceScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `resourceArn` | `string` | yes |
| `resourceOwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

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


## UpdateAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `configuration` | `AnalyzerConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `AnalyzerConfiguration` | no |

## UpdateArchiveRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerName` | `string` | yes |
| `ruleName` | `string` | yes |
| `filter` | `Map<Criterion>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzerArn` | `string` | yes |
| `status` | `string` | yes |
| `ids` | `List<string>` | no |
| `resourceArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ValidatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `locale` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `policyDocument` | `string` | yes |
| `policyType` | `string` | yes |
| `validatePolicyResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<ValidatePolicyFinding>` | yes |
| `nextToken` | `string` | no |

