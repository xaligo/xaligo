# AWS Elemental MediaPackage

API version: 2017-10-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediapackage/2017-10-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ConfigureLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | yes |
| `IngressAccessLogs` | `IngressAccessLogs` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Id` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

## CreateHarvestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndTime` | `string` | yes |
| `Id` | `string` | yes |
| `OriginEndpointId` | `string` | yes |
| `S3Destination` | `S3Destination` | yes |
| `StartTime` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelId` | `string` | no |
| `CreatedAt` | `string` | no |
| `EndTime` | `string` | no |
| `Id` | `string` | no |
| `OriginEndpointId` | `string` | no |
| `S3Destination` | `S3Destination` | no |
| `StartTime` | `string` | no |
| `Status` | `string` | no |

## CreateOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `ChannelId` | `string` | yes |
| `CmafPackage` | `CmafPackageCreateOrUpdateParameters` | no |
| `DashPackage` | `DashPackage` | no |
| `Description` | `string` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | yes |
| `ManifestName` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `Origination` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TimeDelaySeconds` | `integer` | no |
| `Whitelist` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `ChannelId` | `string` | no |
| `CmafPackage` | `CmafPackage` | no |
| `CreatedAt` | `string` | no |
| `DashPackage` | `DashPackage` | no |
| `Description` | `string` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | no |
| `ManifestName` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `Origination` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TimeDelaySeconds` | `integer` | no |
| `Url` | `string` | no |
| `Whitelist` | `List<string>` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

## DescribeHarvestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelId` | `string` | no |
| `CreatedAt` | `string` | no |
| `EndTime` | `string` | no |
| `Id` | `string` | no |
| `OriginEndpointId` | `string` | no |
| `S3Destination` | `S3Destination` | no |
| `StartTime` | `string` | no |
| `Status` | `string` | no |

## DescribeOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `ChannelId` | `string` | no |
| `CmafPackage` | `CmafPackage` | no |
| `CreatedAt` | `string` | no |
| `DashPackage` | `DashPackage` | no |
| `Description` | `string` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | no |
| `ManifestName` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `Origination` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TimeDelaySeconds` | `integer` | no |
| `Url` | `string` | no |
| `Whitelist` | `List<string>` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<Channel>` | no |
| `NextToken` | `string` | no |

## ListHarvestJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncludeChannelId` | `string` | no |
| `IncludeStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HarvestJobs` | `List<HarvestJob>` | no |
| `NextToken` | `string` | no |

## ListOriginEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `OriginEndpoints` | `List<OriginEndpoint>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## RotateChannelCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

## RotateIngestEndpointCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IngestEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

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
| `Description` | `string` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `Description` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `HlsIngest` | `HlsIngest` | no |
| `Id` | `string` | no |
| `IngressAccessLogs` | `IngressAccessLogs` | no |
| `Tags` | `Map<string>` | no |

## UpdateOriginEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `CmafPackage` | `CmafPackageCreateOrUpdateParameters` | no |
| `DashPackage` | `DashPackage` | no |
| `Description` | `string` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | yes |
| `ManifestName` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `Origination` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `TimeDelaySeconds` | `integer` | no |
| `Whitelist` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `ChannelId` | `string` | no |
| `CmafPackage` | `CmafPackage` | no |
| `CreatedAt` | `string` | no |
| `DashPackage` | `DashPackage` | no |
| `Description` | `string` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | no |
| `ManifestName` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `Origination` | `string` | no |
| `StartoverWindowSeconds` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TimeDelaySeconds` | `integer` | no |
| `Url` | `string` | no |
| `Whitelist` | `List<string>` | no |

