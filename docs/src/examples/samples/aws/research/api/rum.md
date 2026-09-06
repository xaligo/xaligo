# CloudWatch RUM

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rum/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateRumMetricDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |
| `MetricDefinitions` | `List<MetricDefinitionRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchCreateRumMetricDefinitionsError>` | yes |
| `MetricDefinitions` | `List<MetricDefinition>` | no |

## BatchDeleteRumMetricDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |
| `MetricDefinitionIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchDeleteRumMetricDefinitionsError>` | yes |
| `MetricDefinitionIds` | `List<string>` | no |

## BatchGetRumMetricDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDefinitions` | `List<MetricDefinition>` | no |
| `NextToken` | `string` | no |

## CreateAppMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Domain` | `string` | no |
| `DomainList` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `AppMonitorConfiguration` | `AppMonitorConfiguration` | no |
| `CwLogEnabled` | `boolean` | no |
| `CustomEvents` | `CustomEvents` | no |
| `DeobfuscationConfiguration` | `DeobfuscationConfiguration` | no |
| `Platform` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## DeleteAppMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyRevisionId` | `string` | no |

## DeleteRumMetricsDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAppMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitor` | `AppMonitor` | no |

## GetAppMonitorData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TimeRange` | `TimeRange` | yes |
| `Filters` | `List<QueryFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<string>` | no |
| `NextToken` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | no |
| `PolicyRevisionId` | `string` | no |

## ListAppMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AppMonitorSummaries` | `List<AppMonitorSummary>` | no |

## ListRumMetricsDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destinations` | `List<MetricDestinationSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PolicyDocument` | `string` | yes |
| `PolicyRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | no |
| `PolicyRevisionId` | `string` | no |

## PutRumEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `BatchId` | `string` | yes |
| `AppMonitorDetails` | `AppMonitorDetails` | yes |
| `UserDetails` | `UserDetails` | yes |
| `RumEvents` | `List<RumEvent>` | yes |
| `Alias` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRumMetricsDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |
| `IamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAppMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Domain` | `string` | no |
| `DomainList` | `List<string>` | no |
| `AppMonitorConfiguration` | `AppMonitorConfiguration` | no |
| `CwLogEnabled` | `boolean` | no |
| `CustomEvents` | `CustomEvents` | no |
| `DeobfuscationConfiguration` | `DeobfuscationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRumMetricDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppMonitorName` | `string` | yes |
| `Destination` | `string` | yes |
| `DestinationArn` | `string` | no |
| `MetricDefinition` | `MetricDefinitionRequest` | yes |
| `MetricDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


