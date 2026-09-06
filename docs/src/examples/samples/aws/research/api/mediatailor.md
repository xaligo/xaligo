# AWS MediaTailor

API version: 2018-04-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediatailor/2018-04-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ConfigureLogsForChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `LogTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | no |
| `LogTypes` | `List<string>` | no |

## ConfigureLogsForPlaybackConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PercentEnabled` | `integer` | yes |
| `PlaybackConfigurationName` | `string` | yes |
| `EnabledLoggingStrategies` | `List<string>` | no |
| `AdsInteractionLog` | `AdsInteractionLog` | no |
| `ManifestServiceInteractionLog` | `ManifestServiceInteractionLog` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PercentEnabled` | `integer` | yes |
| `PlaybackConfigurationName` | `string` | no |
| `EnabledLoggingStrategies` | `List<string>` | no |
| `AdsInteractionLog` | `AdsInteractionLog` | no |
| `ManifestServiceInteractionLog` | `ManifestServiceInteractionLog` | no |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `FillerSlate` | `SlateSource` | no |
| `Outputs` | `List<RequestOutputItem>` | yes |
| `PlaybackMode` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Tier` | `string` | no |
| `TimeShiftConfiguration` | `TimeShiftConfiguration` | no |
| `Audiences` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `ChannelState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `FillerSlate` | `SlateSource` | no |
| `LastModifiedTime` | `timestamp` | no |
| `Outputs` | `List<ResponseOutputItem>` | no |
| `PlaybackMode` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Tier` | `string` | no |
| `TimeShiftConfiguration` | `TimeShiftConfiguration` | no |
| `Audiences` | `List<string>` | no |

## CreateLiveSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | yes |
| `LiveSourceName` | `string` | yes |
| `SourceLocationName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LiveSourceName` | `string` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreatePrefetchSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Consumption` | `PrefetchConsumption` | no |
| `Name` | `string` | yes |
| `PlaybackConfigurationName` | `string` | yes |
| `Retrieval` | `PrefetchRetrieval` | no |
| `RecurringPrefetchConfiguration` | `RecurringPrefetchConfiguration` | no |
| `ScheduleType` | `string` | no |
| `StreamId` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Consumption` | `PrefetchConsumption` | no |
| `Name` | `string` | no |
| `PlaybackConfigurationName` | `string` | no |
| `Retrieval` | `PrefetchRetrieval` | no |
| `RecurringPrefetchConfiguration` | `RecurringPrefetchConfiguration` | no |
| `ScheduleType` | `string` | no |
| `StreamId` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreaks` | `List<AdBreak>` | no |
| `ChannelName` | `string` | yes |
| `LiveSourceName` | `string` | no |
| `ProgramName` | `string` | yes |
| `ScheduleConfiguration` | `ScheduleConfiguration` | yes |
| `SourceLocationName` | `string` | yes |
| `VodSourceName` | `string` | no |
| `AudienceMedia` | `List<AudienceMedia>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreaks` | `List<AdBreak>` | no |
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LiveSourceName` | `string` | no |
| `ProgramName` | `string` | no |
| `ScheduledStartTime` | `timestamp` | no |
| `SourceLocationName` | `string` | no |
| `VodSourceName` | `string` | no |
| `ClipRange` | `ClipRange` | no |
| `DurationMillis` | `long` | no |
| `AudienceMedia` | `List<AudienceMedia>` | no |
| `Tags` | `Map<string>` | no |

## CreateSourceLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessConfiguration` | `AccessConfiguration` | no |
| `DefaultSegmentDeliveryConfiguration` | `DefaultSegmentDeliveryConfiguration` | no |
| `HttpConfiguration` | `HttpConfiguration` | yes |
| `SegmentDeliveryConfigurations` | `List<SegmentDeliveryConfiguration>` | no |
| `SourceLocationName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessConfiguration` | `AccessConfiguration` | no |
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `DefaultSegmentDeliveryConfiguration` | `DefaultSegmentDeliveryConfiguration` | no |
| `HttpConfiguration` | `HttpConfiguration` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SegmentDeliveryConfigurations` | `List<SegmentDeliveryConfiguration>` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateVodSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | yes |
| `SourceLocationName` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `VodSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `VodSourceName` | `string` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLiveSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LiveSourceName` | `string` | yes |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlaybackConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrefetchSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PlaybackConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `ProgramName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSourceLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVodSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceLocationName` | `string` | yes |
| `VodSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `ChannelState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `FillerSlate` | `SlateSource` | no |
| `LastModifiedTime` | `timestamp` | no |
| `Outputs` | `List<ResponseOutputItem>` | no |
| `PlaybackMode` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Tier` | `string` | no |
| `LogConfiguration` | `LogConfigurationForChannel` | yes |
| `TimeShiftConfiguration` | `TimeShiftConfiguration` | no |
| `Audiences` | `List<string>` | no |

## DescribeLiveSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LiveSourceName` | `string` | yes |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LiveSourceName` | `string` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `ProgramName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreaks` | `List<AdBreak>` | no |
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LiveSourceName` | `string` | no |
| `ProgramName` | `string` | no |
| `ScheduledStartTime` | `timestamp` | no |
| `SourceLocationName` | `string` | no |
| `VodSourceName` | `string` | no |
| `ClipRange` | `ClipRange` | no |
| `DurationMillis` | `long` | no |
| `AudienceMedia` | `List<AudienceMedia>` | no |
| `Tags` | `Map<string>` | no |

## DescribeSourceLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessConfiguration` | `AccessConfiguration` | no |
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `DefaultSegmentDeliveryConfiguration` | `DefaultSegmentDeliveryConfiguration` | no |
| `HttpConfiguration` | `HttpConfiguration` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SegmentDeliveryConfigurations` | `List<SegmentDeliveryConfiguration>` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeVodSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceLocationName` | `string` | yes |
| `VodSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreakOpportunities` | `List<AdBreakOpportunity>` | no |
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `VodSourceName` | `string` | no |

## GetChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetChannelSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `DurationMinutes` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Audience` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ScheduleEntry>` | no |
| `NextToken` | `string` | no |

## GetFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionId` | `string` | yes |
| `FunctionType` | `string` | yes |
| `Description` | `string` | no |
| `HttpRequestConfiguration` | `HttpRequestConfiguration` | no |
| `CustomOutputConfiguration` | `CustomOutputConfiguration` | no |
| `ConcurrentExecutorConfiguration` | `ConcurrentExecutorConfiguration` | no |
| `SequentialExecutorConfiguration` | `SequentialExecutorConfiguration` | no |
| `VastRequestConfiguration` | `VastRequestConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `Arn` | `string` | no |

## GetPlaybackConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdDecisionServerUrl` | `string` | no |
| `AvailSuppression` | `AvailSuppression` | no |
| `Bumper` | `Bumper` | no |
| `CdnConfiguration` | `CdnConfiguration` | no |
| `ConfigurationAliases` | `Map<Map<string>>` | no |
| `DashConfiguration` | `DashConfiguration` | no |
| `HlsConfiguration` | `HlsConfiguration` | no |
| `InsertionMode` | `string` | no |
| `LivePreRollConfiguration` | `LivePreRollConfiguration` | no |
| `LogConfiguration` | `LogConfiguration` | no |
| `ManifestProcessingRules` | `ManifestProcessingRules` | no |
| `Name` | `string` | no |
| `PersonalizationThresholdSeconds` | `integer` | no |
| `PlaybackConfigurationArn` | `string` | no |
| `PlaybackEndpointPrefix` | `string` | no |
| `DualStackPlaybackEndpointPrefix` | `string` | no |
| `SessionInitializationEndpointPrefix` | `string` | no |
| `DualStackSessionInitializationEndpointPrefix` | `string` | no |
| `SlateAdUrl` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TranscodeProfileName` | `string` | no |
| `VideoContentSourceUrl` | `string` | no |
| `AdConditioningConfiguration` | `AdConditioningConfiguration` | no |
| `AdDecisionServerConfiguration` | `AdDecisionServerConfiguration` | no |
| `YieldOptimizationConfiguration` | `YieldOptimizationConfiguration` | no |
| `FunctionMapping` | `Map<string>` | no |
| `AdsPersonalizationTimeouts` | `AdsPersonalizationTimeouts` | no |
| `AdsPersonalizationConcurrency` | `AdsPersonalizationConcurrency` | no |

## GetPrefetchSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `PlaybackConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Consumption` | `PrefetchConsumption` | no |
| `Name` | `string` | no |
| `PlaybackConfigurationName` | `string` | no |
| `Retrieval` | `PrefetchRetrieval` | no |
| `ScheduleType` | `string` | no |
| `RecurringPrefetchConfiguration` | `RecurringPrefetchConfiguration` | no |
| `StreamId` | `string` | no |
| `Tags` | `Map<string>` | no |

## ListAlerts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Alert>` | no |
| `NextToken` | `string` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Channel>` | no |
| `NextToken` | `string` | no |

## ListFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Function>` | no |
| `NextToken` | `string` | no |

## ListLiveSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<LiveSource>` | no |
| `NextToken` | `string` | no |

## ListPlaybackConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<PlaybackConfiguration>` | no |
| `NextToken` | `string` | no |

## ListPrefetchSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PlaybackConfigurationName` | `string` | yes |
| `ScheduleType` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<PrefetchSchedule>` | no |
| `NextToken` | `string` | no |

## ListSourceLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<SourceLocation>` | no |
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

## ListVodSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<VodSource>` | no |
| `NextToken` | `string` | no |

## PutChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionId` | `string` | yes |
| `FunctionType` | `string` | yes |
| `Description` | `string` | no |
| `HttpRequestConfiguration` | `HttpRequestConfiguration` | no |
| `CustomOutputConfiguration` | `CustomOutputConfiguration` | no |
| `ConcurrentExecutorConfiguration` | `ConcurrentExecutorConfiguration` | no |
| `SequentialExecutorConfiguration` | `SequentialExecutorConfiguration` | no |
| `VastRequestConfiguration` | `VastRequestConfiguration` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionId` | `string` | yes |
| `FunctionType` | `string` | yes |
| `Description` | `string` | no |
| `HttpRequestConfiguration` | `HttpRequestConfiguration` | no |
| `CustomOutputConfiguration` | `CustomOutputConfiguration` | no |
| `ConcurrentExecutorConfiguration` | `ConcurrentExecutorConfiguration` | no |
| `SequentialExecutorConfiguration` | `SequentialExecutorConfiguration` | no |
| `VastRequestConfiguration` | `VastRequestConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `Arn` | `string` | no |

## PutPlaybackConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdDecisionServerUrl` | `string` | no |
| `AvailSuppression` | `AvailSuppression` | no |
| `Bumper` | `Bumper` | no |
| `CdnConfiguration` | `CdnConfiguration` | no |
| `ConfigurationAliases` | `Map<Map<string>>` | no |
| `DashConfiguration` | `DashConfigurationForPut` | no |
| `InsertionMode` | `string` | no |
| `LivePreRollConfiguration` | `LivePreRollConfiguration` | no |
| `ManifestProcessingRules` | `ManifestProcessingRules` | no |
| `Name` | `string` | yes |
| `PersonalizationThresholdSeconds` | `integer` | no |
| `SlateAdUrl` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TranscodeProfileName` | `string` | no |
| `VideoContentSourceUrl` | `string` | no |
| `AdConditioningConfiguration` | `AdConditioningConfiguration` | no |
| `AdDecisionServerConfiguration` | `AdDecisionServerConfiguration` | no |
| `YieldOptimizationConfiguration` | `YieldOptimizationConfiguration` | no |
| `FunctionMapping` | `Map<string>` | no |
| `AdsPersonalizationTimeouts` | `AdsPersonalizationTimeouts` | no |
| `AdsPersonalizationConcurrency` | `AdsPersonalizationConcurrency` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdDecisionServerUrl` | `string` | no |
| `AvailSuppression` | `AvailSuppression` | no |
| `Bumper` | `Bumper` | no |
| `CdnConfiguration` | `CdnConfiguration` | no |
| `ConfigurationAliases` | `Map<Map<string>>` | no |
| `DashConfiguration` | `DashConfiguration` | no |
| `HlsConfiguration` | `HlsConfiguration` | no |
| `InsertionMode` | `string` | no |
| `LivePreRollConfiguration` | `LivePreRollConfiguration` | no |
| `LogConfiguration` | `LogConfiguration` | no |
| `ManifestProcessingRules` | `ManifestProcessingRules` | no |
| `Name` | `string` | no |
| `PersonalizationThresholdSeconds` | `integer` | no |
| `PlaybackConfigurationArn` | `string` | no |
| `PlaybackEndpointPrefix` | `string` | no |
| `DualStackPlaybackEndpointPrefix` | `string` | no |
| `SessionInitializationEndpointPrefix` | `string` | no |
| `DualStackSessionInitializationEndpointPrefix` | `string` | no |
| `SlateAdUrl` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TranscodeProfileName` | `string` | no |
| `VideoContentSourceUrl` | `string` | no |
| `AdConditioningConfiguration` | `AdConditioningConfiguration` | no |
| `AdDecisionServerConfiguration` | `AdDecisionServerConfiguration` | no |
| `YieldOptimizationConfiguration` | `YieldOptimizationConfiguration` | no |
| `FunctionMapping` | `Map<string>` | no |
| `AdsPersonalizationTimeouts` | `AdsPersonalizationTimeouts` | no |
| `AdsPersonalizationConcurrency` | `AdsPersonalizationConcurrency` | no |

## StartChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |

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


## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `FillerSlate` | `SlateSource` | no |
| `Outputs` | `List<RequestOutputItem>` | yes |
| `TimeShiftConfiguration` | `TimeShiftConfiguration` | no |
| `Audiences` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `ChannelState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `FillerSlate` | `SlateSource` | no |
| `LastModifiedTime` | `timestamp` | no |
| `Outputs` | `List<ResponseOutputItem>` | no |
| `PlaybackMode` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Tier` | `string` | no |
| `TimeShiftConfiguration` | `TimeShiftConfiguration` | no |
| `Audiences` | `List<string>` | no |

## UpdateLiveSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | yes |
| `LiveSourceName` | `string` | yes |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LiveSourceName` | `string` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreaks` | `List<AdBreak>` | no |
| `ChannelName` | `string` | yes |
| `ProgramName` | `string` | yes |
| `ScheduleConfiguration` | `UpdateProgramScheduleConfiguration` | yes |
| `AudienceMedia` | `List<AudienceMedia>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdBreaks` | `List<AdBreak>` | no |
| `Arn` | `string` | no |
| `ChannelName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `ProgramName` | `string` | no |
| `SourceLocationName` | `string` | no |
| `VodSourceName` | `string` | no |
| `LiveSourceName` | `string` | no |
| `ClipRange` | `ClipRange` | no |
| `DurationMillis` | `long` | no |
| `ScheduledStartTime` | `timestamp` | no |
| `AudienceMedia` | `List<AudienceMedia>` | no |
| `Tags` | `Map<string>` | no |

## UpdateSourceLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessConfiguration` | `AccessConfiguration` | no |
| `DefaultSegmentDeliveryConfiguration` | `DefaultSegmentDeliveryConfiguration` | no |
| `HttpConfiguration` | `HttpConfiguration` | yes |
| `SegmentDeliveryConfigurations` | `List<SegmentDeliveryConfiguration>` | no |
| `SourceLocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessConfiguration` | `AccessConfiguration` | no |
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `DefaultSegmentDeliveryConfiguration` | `DefaultSegmentDeliveryConfiguration` | no |
| `HttpConfiguration` | `HttpConfiguration` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SegmentDeliveryConfigurations` | `List<SegmentDeliveryConfiguration>` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateVodSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | yes |
| `SourceLocationName` | `string` | yes |
| `VodSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `HttpPackageConfigurations` | `List<HttpPackageConfiguration>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SourceLocationName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `VodSourceName` | `string` | no |

