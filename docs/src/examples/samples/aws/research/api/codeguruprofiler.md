# Amazon CodeGuru Profiler

API version: 2019-07-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codeguruprofiler/2019-07-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddNotificationChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channels` | `List<Channel>` | yes |
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfiguration` | `NotificationConfiguration` | no |

## BatchGetFrameMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | no |
| `frameMetrics` | `List<FrameMetric>` | no |
| `period` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `startTime` | `timestamp` | no |
| `targetResolution` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | yes |
| `endTimes` | `List<TimestampStructure>` | yes |
| `frameMetricData` | `List<FrameMetricDatum>` | yes |
| `resolution` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `unprocessedEndTimes` | `Map<List<TimestampStructure>>` | yes |

## ConfigureAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetInstanceId` | `string` | no |
| `metadata` | `Map<string>` | no |
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `AgentConfiguration` | yes |

## CreateProfilingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentOrchestrationConfig` | `AgentOrchestrationConfig` | no |
| `clientToken` | `string` | yes |
| `computePlatform` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroup` | `ProfilingGroupDescription` | yes |

## DeleteProfilingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeProfilingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroup` | `ProfilingGroupDescription` | yes |

## GetFindingsReportAccountSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dailyReportsOnly` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reportSummaries` | `List<FindingsReportSummary>` | yes |

## GetNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfiguration` | `NotificationConfiguration` | yes |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `revisionId` | `string` | yes |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accept` | `string` | no |
| `endTime` | `timestamp` | no |
| `maxDepth` | `integer` | no |
| `period` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `startTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentEncoding` | `string` | no |
| `contentType` | `string` | yes |
| `profile` | `blob` | yes |

## GetRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | yes |
| `locale` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `startTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalies` | `List<Anomaly>` | yes |
| `profileEndTime` | `timestamp` | yes |
| `profileStartTime` | `timestamp` | yes |
| `profilingGroupName` | `string` | yes |
| `recommendations` | `List<Recommendation>` | yes |

## ListFindingsReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dailyReportsOnly` | `boolean` | no |
| `endTime` | `timestamp` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `startTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsReportSummaries` | `List<FindingsReportSummary>` | yes |
| `nextToken` | `string` | no |

## ListProfileTimes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endTime` | `timestamp` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `orderBy` | `string` | no |
| `period` | `string` | yes |
| `profilingGroupName` | `string` | yes |
| `startTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `profileTimes` | `List<ProfileTime>` | yes |

## ListProfilingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeDescription` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `profilingGroupNames` | `List<string>` | yes |
| `profilingGroups` | `List<ProfilingGroupDescription>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PostAgentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentProfile` | `blob` | yes |
| `contentType` | `string` | yes |
| `profileToken` | `string` | no |
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionGroup` | `string` | yes |
| `principals` | `List<string>` | yes |
| `profilingGroupName` | `string` | yes |
| `revisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `revisionId` | `string` | yes |

## RemoveNotificationChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | yes |
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfiguration` | `NotificationConfiguration` | no |

## RemovePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionGroup` | `string` | yes |
| `profilingGroupName` | `string` | yes |
| `revisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `revisionId` | `string` | yes |

## SubmitFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyInstanceId` | `string` | yes |
| `comment` | `string` | no |
| `profilingGroupName` | `string` | yes |
| `type` | `string` | yes |

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


## UpdateProfilingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentOrchestrationConfig` | `AgentOrchestrationConfig` | yes |
| `profilingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profilingGroup` | `ProfilingGroupDescription` | yes |

