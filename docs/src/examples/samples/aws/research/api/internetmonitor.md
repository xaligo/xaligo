# Amazon CloudWatch Internet Monitor

API version: 2021-06-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/internetmonitor/2021-06-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `Resources` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `MaxCityNetworksToMonitor` | `integer` | no |
| `InternetMeasurementsLogDelivery` | `InternetMeasurementsLogDelivery` | no |
| `TrafficPercentageToMonitor` | `integer` | no |
| `HealthEventsConfig` | `HealthEventsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Status` | `string` | yes |

## DeleteMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetHealthEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `EventId` | `string` | yes |
| `LinkedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventArn` | `string` | yes |
| `EventId` | `string` | yes |
| `StartedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | yes |
| `ImpactedLocations` | `List<ImpactedLocation>` | yes |
| `Status` | `string` | yes |
| `PercentOfTotalTrafficImpacted` | `double` | no |
| `ImpactType` | `string` | yes |
| `HealthScoreThreshold` | `double` | no |

## GetInternetEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventId` | `string` | yes |
| `EventArn` | `string` | yes |
| `StartedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |
| `ClientLocation` | `ClientLocation` | yes |
| `EventType` | `string` | yes |
| `EventStatus` | `string` | yes |

## GetMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `LinkedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `MonitorArn` | `string` | yes |
| `Resources` | `List<string>` | yes |
| `Status` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `ProcessingStatus` | `string` | no |
| `ProcessingStatusInfo` | `string` | no |
| `Tags` | `Map<string>` | no |
| `MaxCityNetworksToMonitor` | `integer` | no |
| `InternetMeasurementsLogDelivery` | `InternetMeasurementsLogDelivery` | no |
| `TrafficPercentageToMonitor` | `integer` | no |
| `HealthEventsConfig` | `HealthEventsConfig` | no |

## GetQueryResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `QueryId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fields` | `List<QueryField>` | yes |
| `Data` | `List<List<string>>` | yes |
| `NextToken` | `string` | no |

## GetQueryStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `QueryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## ListHealthEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `EventStatus` | `string` | no |
| `LinkedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthEvents` | `List<HealthEvent>` | yes |
| `NextToken` | `string` | no |

## ListInternetEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `EventStatus` | `string` | no |
| `EventType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InternetEvents` | `List<InternetEventSummary>` | yes |
| `NextToken` | `string` | no |

## ListMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `MonitorStatus` | `string` | no |
| `IncludeLinkedAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Monitors` | `List<Monitor>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `QueryType` | `string` | yes |
| `FilterParameters` | `List<FilterParameter>` | no |
| `LinkedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |

## StopQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `QueryId` | `string` | yes |

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


## UpdateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorName` | `string` | yes |
| `ResourcesToAdd` | `List<string>` | no |
| `ResourcesToRemove` | `List<string>` | no |
| `Status` | `string` | no |
| `ClientToken` | `string` | no |
| `MaxCityNetworksToMonitor` | `integer` | no |
| `InternetMeasurementsLogDelivery` | `InternetMeasurementsLogDelivery` | no |
| `TrafficPercentageToMonitor` | `integer` | no |
| `HealthEventsConfig` | `HealthEventsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitorArn` | `string` | yes |
| `Status` | `string` | yes |

