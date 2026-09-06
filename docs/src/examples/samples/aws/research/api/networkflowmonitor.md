# Network Flow Monitor

API version: 2023-04-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/networkflowmonitor/2023-04-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `localResources` | `List<MonitorLocalResource>` | yes |
| `remoteResources` | `List<MonitorRemoteResource>` | no |
| `scopeArn` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `monitorStatus` | `string` | yes |
| `localResources` | `List<MonitorLocalResource>` | yes |
| `remoteResources` | `List<MonitorRemoteResource>` | yes |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## CreateScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targets` | `List<TargetResource>` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `status` | `string` | yes |
| `scopeArn` | `string` | yes |
| `tags` | `Map<string>` | no |

## DeleteMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `monitorStatus` | `string` | yes |
| `localResources` | `List<MonitorLocalResource>` | yes |
| `remoteResources` | `List<MonitorRemoteResource>` | yes |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetQueryResultsMonitorTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `queryId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unit` | `string` | no |
| `topContributors` | `List<MonitorTopContributorsRow>` | no |
| `nextToken` | `string` | no |

## GetQueryResultsWorkloadInsightsTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `topContributors` | `List<WorkloadInsightsTopContributorsRow>` | no |
| `nextToken` | `string` | no |

## GetQueryResultsWorkloadInsightsTopContributorsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unit` | `string` | yes |
| `datapoints` | `List<WorkloadInsightsTopContributorsDataPoint>` | yes |
| `nextToken` | `string` | no |

## GetQueryStatusMonitorTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## GetQueryStatusWorkloadInsightsTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## GetQueryStatusWorkloadInsightsTopContributorsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## GetScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `status` | `string` | yes |
| `scopeArn` | `string` | yes |
| `targets` | `List<TargetResource>` | yes |
| `tags` | `Map<string>` | no |

## ListMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `monitorStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitors` | `List<MonitorSummary>` | yes |
| `nextToken` | `string` | no |

## ListScopes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopes` | `List<ScopeSummary>` | yes |
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

## StartQueryMonitorTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `metricName` | `string` | yes |
| `destinationCategory` | `string` | yes |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

## StartQueryWorkloadInsightsTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `metricName` | `string` | yes |
| `destinationCategory` | `string` | yes |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

## StartQueryWorkloadInsightsTopContributorsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `metricName` | `string` | yes |
| `destinationCategory` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

## StopQueryMonitorTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopQueryWorkloadInsightsTopContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopQueryWorkloadInsightsTopContributorsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `queryId` | `string` | yes |

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


## UpdateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorName` | `string` | yes |
| `localResourcesToAdd` | `List<MonitorLocalResource>` | no |
| `localResourcesToRemove` | `List<MonitorLocalResource>` | no |
| `remoteResourcesToAdd` | `List<MonitorRemoteResource>` | no |
| `remoteResourcesToRemove` | `List<MonitorRemoteResource>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorArn` | `string` | yes |
| `monitorName` | `string` | yes |
| `monitorStatus` | `string` | yes |
| `localResources` | `List<MonitorLocalResource>` | yes |
| `remoteResources` | `List<MonitorRemoteResource>` | yes |
| `createdAt` | `timestamp` | yes |
| `modifiedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## UpdateScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `resourcesToAdd` | `List<TargetResource>` | no |
| `resourcesToDelete` | `List<TargetResource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scopeId` | `string` | yes |
| `status` | `string` | yes |
| `scopeArn` | `string` | yes |
| `tags` | `Map<string>` | no |

