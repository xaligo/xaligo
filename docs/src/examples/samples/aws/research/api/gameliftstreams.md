# Amazon GameLift Streams

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/gameliftstreams/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddStreamGroupLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `LocationConfigurations` | `List<LocationConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Locations` | `List<LocationState>` | yes |

## AssociateApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `ApplicationIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ApplicationArns` | `List<string>` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | yes |
| `RuntimeEnvironment` | `RuntimeEnvironment` | yes |
| `ExecutablePath` | `string` | yes |
| `ApplicationSourceUri` | `string` | yes |
| `ApplicationLogPaths` | `List<string>` | no |
| `ApplicationLogOutputUri` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `RuntimeEnvironment` | `RuntimeEnvironment` | no |
| `ExecutablePath` | `string` | no |
| `ApplicationLogPaths` | `List<string>` | no |
| `ApplicationLogOutputUri` | `string` | no |
| `ApplicationSourceUri` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `ReplicationStatuses` | `List<ReplicationStatus>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `AssociatedStreamGroups` | `List<string>` | no |

## CreateStreamGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | yes |
| `StreamClass` | `string` | yes |
| `DefaultApplicationIdentifier` | `string` | no |
| `LocationConfigurations` | `List<LocationConfiguration>` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `DefaultApplication` | `DefaultApplication` | no |
| `LocationStates` | `List<LocationState>` | no |
| `StreamClass` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ExpiresAt` | `timestamp` | no |
| `AssociatedApplications` | `List<string>` | no |

## CreateStreamSessionAdminShell

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamSessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |
| `StreamUrl` | `string` | no |
| `TokenValue` | `string` | no |

## CreateStreamSessionConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Identifier` | `string` | yes |
| `StreamSessionIdentifier` | `string` | yes |
| `SignalRequest` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SignalResponse` | `string` | no |

## CreateStreamUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `ApplicationIdentifier` | `string` | yes |
| `Protocol` | `string` | yes |
| `UrlExpiresAfterMinutes` | `integer` | yes |
| `UsageLimit` | `integer` | no |
| `Description` | `string` | no |
| `Locations` | `List<string>` | yes |
| `SessionLengthSeconds` | `integer` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `StreamUrlId` | `string` | no |
| `StreamUrl` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `ExpiresAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `UsageLimit` | `integer` | no |
| `RemainingUses` | `integer` | no |
| `StreamGroupArn` | `string` | no |
| `ApplicationArn` | `string` | no |
| `Protocol` | `string` | no |
| `Locations` | `List<string>` | no |
| `SessionLengthSeconds` | `integer` | no |
| `Description` | `string` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStreamGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `ApplicationIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ApplicationArns` | `List<string>` | no |

## ExportStreamSessionFiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamSessionIdentifier` | `string` | yes |
| `OutputUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `RuntimeEnvironment` | `RuntimeEnvironment` | no |
| `ExecutablePath` | `string` | no |
| `ApplicationLogPaths` | `List<string>` | no |
| `ApplicationLogOutputUri` | `string` | no |
| `ApplicationSourceUri` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `ReplicationStatuses` | `List<ReplicationStatus>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `AssociatedStreamGroups` | `List<string>` | no |

## GetStreamGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `DefaultApplication` | `DefaultApplication` | no |
| `LocationStates` | `List<LocationState>` | no |
| `StreamClass` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ExpiresAt` | `timestamp` | no |
| `AssociatedApplications` | `List<string>` | no |

## GetStreamSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamSessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `StreamGroupId` | `string` | no |
| `UserId` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `Protocol` | `string` | no |
| `Location` | `string` | no |
| `SignalRequest` | `string` | no |
| `SignalResponse` | `string` | no |
| `ConnectionTimeoutSeconds` | `integer` | no |
| `SessionLengthSeconds` | `integer` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `PerformanceStatsConfiguration` | `PerformanceStatsConfiguration` | no |
| `LogFileLocationUri` | `string` | no |
| `WebSdkProtocolUrl` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ApplicationArn` | `string` | no |
| `ExportFilesMetadata` | `ExportFilesMetadata` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |

## GetStreamUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamUrlIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `StreamUrlId` | `string` | no |
| `StreamUrl` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `ExpiresAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `UsageLimit` | `integer` | no |
| `RemainingUses` | `integer` | no |
| `StreamGroupArn` | `string` | no |
| `ApplicationArn` | `string` | no |
| `Protocol` | `string` | no |
| `Locations` | `List<string>` | no |
| `SessionLengthSeconds` | `integer` | no |
| `Description` | `string` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |
| `StreamSessions` | `List<StreamSessionSummary>` | no |

## ListApplicationShaderCaches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ShaderCacheSummary>` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ApplicationSummary>` | no |
| `NextToken` | `string` | no |

## ListStreamGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<StreamGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListStreamSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ExportFilesStatus` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<StreamSessionSummary>` | no |
| `NextToken` | `string` | no |

## ListStreamSessionsByAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ExportFilesStatus` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<StreamSessionSummary>` | no |
| `NextToken` | `string` | no |

## ListStreamUrls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `StreamGroupIdentifier` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<StreamUrlSummary>` | no |
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

## RemoveStreamGroupLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Locations` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeStreamUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamUrlIdentifier` | `string` | yes |
| `RevocationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartStreamSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Identifier` | `string` | yes |
| `Protocol` | `string` | yes |
| `SignalRequest` | `string` | yes |
| `ApplicationIdentifier` | `string` | yes |
| `UserId` | `string` | no |
| `Locations` | `List<string>` | no |
| `ConnectionTimeoutSeconds` | `integer` | no |
| `SessionLengthSeconds` | `integer` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `PerformanceStatsConfiguration` | `PerformanceStatsConfiguration` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `StreamGroupId` | `string` | no |
| `UserId` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `Protocol` | `string` | no |
| `Location` | `string` | no |
| `SignalRequest` | `string` | no |
| `SignalResponse` | `string` | no |
| `ConnectionTimeoutSeconds` | `integer` | no |
| `SessionLengthSeconds` | `integer` | no |
| `AdditionalLaunchArgs` | `List<string>` | no |
| `AdditionalEnvironmentVariables` | `Map<string>` | no |
| `PerformanceStatsConfiguration` | `PerformanceStatsConfiguration` | no |
| `LogFileLocationUri` | `string` | no |
| `WebSdkProtocolUrl` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ApplicationArn` | `string` | no |
| `ExportFilesMetadata` | `ExportFilesMetadata` | no |
| `RoleArn` | `string` | no |
| `DisplayConfiguration` | `DisplayConfiguration` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateStreamSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamSessionIdentifier` | `string` | yes |

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


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Description` | `string` | no |
| `ApplicationLogPaths` | `List<string>` | no |
| `ApplicationLogOutputUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `RuntimeEnvironment` | `RuntimeEnvironment` | no |
| `ExecutablePath` | `string` | no |
| `ApplicationLogPaths` | `List<string>` | no |
| `ApplicationLogOutputUri` | `string` | no |
| `ApplicationSourceUri` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `ReplicationStatuses` | `List<ReplicationStatus>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `AssociatedStreamGroups` | `List<string>` | no |

## UpdateStreamGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `LocationConfigurations` | `List<LocationConfiguration>` | no |
| `Description` | `string` | no |
| `DefaultApplicationIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `DefaultApplication` | `DefaultApplication` | no |
| `LocationStates` | `List<LocationState>` | no |
| `StreamClass` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ExpiresAt` | `timestamp` | no |
| `AssociatedApplications` | `List<string>` | no |

