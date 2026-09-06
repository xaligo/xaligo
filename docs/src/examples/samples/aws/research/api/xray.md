# AWS X-Ray

API version: 2016-04-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/xray/2016-04-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetTraces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Traces` | `List<Trace>` | no |
| `UnprocessedTraceIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## CancelTraceRetrieval

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `FilterExpression` | `string` | no |
| `InsightsConfiguration` | `InsightsConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## CreateSamplingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRule` | `SamplingRule` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRuleRecord` | `SamplingRuleRecord` | no |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `GroupARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyName` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSamplingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | no |
| `RuleARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRuleRecord` | `SamplingRuleRecord` | no |

## GetEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncryptionConfig` | `EncryptionConfig` | no |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `GroupARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## GetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupSummary>` | no |
| `NextToken` | `string` | no |

## GetIndexingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexingRules` | `List<IndexingRule>` | no |
| `NextToken` | `string` | no |

## GetInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Insight` | `Insight` | no |

## GetInsightEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightEvents` | `List<InsightEvent>` | no |
| `NextToken` | `string` | no |

## GetInsightImpactGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `ServiceGraphStartTime` | `timestamp` | no |
| `ServiceGraphEndTime` | `timestamp` | no |
| `Services` | `List<InsightImpactGraphService>` | no |
| `NextToken` | `string` | no |

## GetInsightSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `States` | `List<string>` | no |
| `GroupARN` | `string` | no |
| `GroupName` | `string` | no |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightSummaries` | `List<InsightSummary>` | no |
| `NextToken` | `string` | no |

## GetRetrievedTracesGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalToken` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalStatus` | `string` | no |
| `Services` | `List<RetrievedService>` | no |
| `NextToken` | `string` | no |

## GetSamplingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRuleRecords` | `List<SamplingRuleRecord>` | no |
| `NextToken` | `string` | no |

## GetSamplingStatisticSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingStatisticSummaries` | `List<SamplingStatisticSummary>` | no |
| `NextToken` | `string` | no |

## GetSamplingTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingStatisticsDocuments` | `List<SamplingStatisticsDocument>` | yes |
| `SamplingBoostStatisticsDocuments` | `List<SamplingBoostStatisticsDocument>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingTargetDocuments` | `List<SamplingTargetDocument>` | no |
| `LastRuleModification` | `timestamp` | no |
| `UnprocessedStatistics` | `List<UnprocessedStatistics>` | no |
| `UnprocessedBoostStatistics` | `List<UnprocessedStatistics>` | no |

## GetServiceGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `GroupName` | `string` | no |
| `GroupARN` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Services` | `List<Service>` | no |
| `ContainsOldGroupVersions` | `boolean` | no |
| `NextToken` | `string` | no |

## GetTimeSeriesServiceStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `GroupName` | `string` | no |
| `GroupARN` | `string` | no |
| `EntitySelectorExpression` | `string` | no |
| `Period` | `integer` | no |
| `ForecastStatistics` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeSeriesServiceStatistics` | `List<TimeSeriesServiceStatistics>` | no |
| `ContainsOldGroupVersions` | `boolean` | no |
| `NextToken` | `string` | no |

## GetTraceGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Services` | `List<Service>` | no |
| `NextToken` | `string` | no |

## GetTraceSegmentDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destination` | `string` | no |
| `Status` | `string` | no |

## GetTraceSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `TimeRangeType` | `string` | no |
| `Sampling` | `boolean` | no |
| `SamplingStrategy` | `SamplingStrategy` | no |
| `FilterExpression` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceSummaries` | `List<TraceSummary>` | no |
| `ApproximateTime` | `timestamp` | no |
| `TracesProcessedCount` | `long` | no |
| `NextToken` | `string` | no |

## ListResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicies` | `List<ResourcePolicy>` | no |
| `NextToken` | `string` | no |

## ListRetrievedTraces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalToken` | `string` | yes |
| `TraceFormat` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalStatus` | `string` | no |
| `TraceFormat` | `string` | no |
| `Traces` | `List<RetrievedTrace>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## PutEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncryptionConfig` | `EncryptionConfig` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyName` | `string` | yes |
| `PolicyDocument` | `string` | yes |
| `PolicyRevisionId` | `string` | no |
| `BypassPolicyLockoutCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `ResourcePolicy` | no |

## PutTelemetryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryRecords` | `List<TelemetryRecord>` | yes |
| `EC2InstanceId` | `string` | no |
| `Hostname` | `string` | no |
| `ResourceARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTraceSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceSegmentDocuments` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedTraceSegments` | `List<UnprocessedTraceSegment>` | no |

## StartTraceRetrieval

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceIds` | `List<string>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetrievalToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `GroupARN` | `string` | no |
| `FilterExpression` | `string` | no |
| `InsightsConfiguration` | `InsightsConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## UpdateIndexingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Rule` | `IndexingRuleValueUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexingRule` | `IndexingRule` | no |

## UpdateSamplingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRuleUpdate` | `SamplingRuleUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SamplingRuleRecord` | `SamplingRuleRecord` | no |

## UpdateTraceSegmentDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destination` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destination` | `string` | no |
| `Status` | `string` | no |

