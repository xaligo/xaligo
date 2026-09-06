# AWS Elemental MediaPackage v2

API version: 2022-12-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediapackagev2/2022-12-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelHarvestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `HarvestJobName` | `string` | yes |
| `ETag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `ClientToken` | `string` | no |
| `InputType` | `string` | no |
| `Description` | `string` | no |
| `InputSwitchConfiguration` | `InputSwitchConfiguration` | no |
| `OutputHeaderConfiguration` | `OutputHeaderConfiguration` | no |
| `OutputLockingMode` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelName` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `IngestEndpoints` | `List<IngestEndpoint>` | no |
| `InputType` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |
| `InputSwitchConfiguration` | `InputSwitchConfiguration` | no |
| `OutputHeaderConfiguration` | `OutputHeaderConfiguration` | no |
| `OutputLockingMode` | `string` | no |

## CreateChannelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `Arn` | `string` | yes |
| `EgressDomain` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `ETag` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateHarvestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Description` | `string` | no |
| `HarvestedManifests` | `HarvestedManifests` | yes |
| `ScheduleConfiguration` | `HarvesterScheduleConfiguration` | yes |
| `Destination` | `Destination` | yes |
| `ClientToken` | `string` | no |
| `HarvestJobName` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Destination` | `Destination` | yes |
| `HarvestJobName` | `string` | yes |
| `HarvestedManifests` | `HarvestedManifests` | yes |
| `Description` | `string` | no |
| `ScheduleConfiguration` | `HarvesterScheduleConfiguration` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `ErrorMessage` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `ContainerType` | `string` | yes |
| `Segment` | `Segment` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `HlsManifests` | `List<CreateHlsManifestConfiguration>` | no |
| `LowLatencyHlsManifests` | `List<CreateLowLatencyHlsManifestConfiguration>` | no |
| `DashManifests` | `List<CreateDashManifestConfiguration>` | no |
| `MssManifests` | `List<CreateMssManifestConfiguration>` | no |
| `ForceEndpointErrorConfiguration` | `ForceEndpointErrorConfiguration` | no |
| `UriSeparator` | `string` | no |
| `StreamNameOutputMode` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `ContainerType` | `string` | yes |
| `Segment` | `Segment` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `HlsManifests` | `List<GetHlsManifestConfiguration>` | no |
| `LowLatencyHlsManifests` | `List<GetLowLatencyHlsManifestConfiguration>` | no |
| `DashManifests` | `List<GetDashManifestConfiguration>` | no |
| `MssManifests` | `List<GetMssManifestConfiguration>` | no |
| `ForceEndpointErrorConfiguration` | `ForceEndpointErrorConfiguration` | no |
| `UriSeparator` | `string` | no |
| `StreamNameOutputMode` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOriginEndpointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelName` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `ResetAt` | `timestamp` | no |
| `Description` | `string` | no |
| `IngestEndpoints` | `List<IngestEndpoint>` | no |
| `InputType` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |
| `InputSwitchConfiguration` | `InputSwitchConfiguration` | no |
| `OutputHeaderConfiguration` | `OutputHeaderConfiguration` | no |
| `OutputLockingMode` | `string` | no |

## GetChannelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `Arn` | `string` | yes |
| `EgressDomain` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `Policy` | `string` | yes |

## GetHarvestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `HarvestJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Destination` | `Destination` | yes |
| `HarvestJobName` | `string` | yes |
| `HarvestedManifests` | `HarvestedManifests` | yes |
| `Description` | `string` | no |
| `ScheduleConfiguration` | `HarvesterScheduleConfiguration` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `ErrorMessage` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `ContainerType` | `string` | yes |
| `Segment` | `Segment` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `ResetAt` | `timestamp` | no |
| `Description` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `HlsManifests` | `List<GetHlsManifestConfiguration>` | no |
| `LowLatencyHlsManifests` | `List<GetLowLatencyHlsManifestConfiguration>` | no |
| `DashManifests` | `List<GetDashManifestConfiguration>` | no |
| `MssManifests` | `List<GetMssManifestConfiguration>` | no |
| `ForceEndpointErrorConfiguration` | `ForceEndpointErrorConfiguration` | no |
| `UriSeparator` | `string` | no |
| `StreamNameOutputMode` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetOriginEndpointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Policy` | `string` | yes |
| `CdnAuthConfiguration` | `CdnAuthConfiguration` | no |

## ListChannelGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ChannelGroupListConfiguration>` | no |
| `NextToken` | `string` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ChannelListConfiguration>` | no |
| `NextToken` | `string` | no |

## ListHarvestJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | no |
| `OriginEndpointName` | `string` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<HarvestJob>` | no |
| `NextToken` | `string` | no |

## ListOriginEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<OriginEndpointListConfiguration>` | no |
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

## PutChannelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutOriginEndpointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Policy` | `string` | yes |
| `CdnAuthConfiguration` | `CdnAuthConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetChannelState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `Arn` | `string` | yes |
| `ResetAt` | `timestamp` | yes |

## ResetOriginEndpointState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `Arn` | `string` | yes |
| `ResetAt` | `timestamp` | yes |

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
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `ETag` | `string` | no |
| `Description` | `string` | no |
| `InputSwitchConfiguration` | `InputSwitchConfiguration` | no |
| `OutputHeaderConfiguration` | `OutputHeaderConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelName` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `IngestEndpoints` | `List<IngestEndpoint>` | no |
| `InputType` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |
| `InputSwitchConfiguration` | `InputSwitchConfiguration` | no |
| `OutputHeaderConfiguration` | `OutputHeaderConfiguration` | no |
| `OutputLockingMode` | `string` | no |

## UpdateChannelGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ETag` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `Arn` | `string` | yes |
| `EgressDomain` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `ContainerType` | `string` | yes |
| `Segment` | `Segment` | no |
| `Description` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `HlsManifests` | `List<CreateHlsManifestConfiguration>` | no |
| `LowLatencyHlsManifests` | `List<CreateLowLatencyHlsManifestConfiguration>` | no |
| `DashManifests` | `List<CreateDashManifestConfiguration>` | no |
| `MssManifests` | `List<CreateMssManifestConfiguration>` | no |
| `ForceEndpointErrorConfiguration` | `ForceEndpointErrorConfiguration` | no |
| `UriSeparator` | `string` | no |
| `StreamNameOutputMode` | `string` | no |
| `ETag` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ChannelGroupName` | `string` | yes |
| `ChannelName` | `string` | yes |
| `OriginEndpointName` | `string` | yes |
| `ContainerType` | `string` | yes |
| `Segment` | `Segment` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ModifiedAt` | `timestamp` | yes |
| `Description` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `HlsManifests` | `List<GetHlsManifestConfiguration>` | no |
| `LowLatencyHlsManifests` | `List<GetLowLatencyHlsManifestConfiguration>` | no |
| `MssManifests` | `List<GetMssManifestConfiguration>` | no |
| `ForceEndpointErrorConfiguration` | `ForceEndpointErrorConfiguration` | no |
| `UriSeparator` | `string` | no |
| `StreamNameOutputMode` | `string` | no |
| `ETag` | `string` | no |
| `Tags` | `Map<string>` | no |
| `DashManifests` | `List<GetDashManifestConfiguration>` | no |

