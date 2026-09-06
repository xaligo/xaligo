# AWS Elemental MediaPackage VOD

API version: 2018-11-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediapackage-vod/2018-11-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ConfigureLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `CreatedAt` | `string` | no |
| `DomainName` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `PackagingGroupId` | `string` | yes |
| `ResourceId` | `string` | no |
| `SourceArn` | `string` | yes |
| `SourceRoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `EgressEndpoints` | `List<EgressEndpoint>` | no |
| `Id` | `string` | no |
| `PackagingGroupId` | `string` | no |
| `ResourceId` | `string` | no |
| `SourceArn` | `string` | no |
| `SourceRoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreatePackagingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CmafPackage` | `CmafPackage` | no |
| `DashPackage` | `DashPackage` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | yes |
| `MssPackage` | `MssPackage` | no |
| `PackagingGroupId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CmafPackage` | `CmafPackage` | no |
| `CreatedAt` | `string` | no |
| `DashPackage` | `DashPackage` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `PackagingGroupId` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreatePackagingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `CreatedAt` | `string` | no |
| `DomainName` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |

## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePackagingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePackagingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `string` | no |
| `EgressEndpoints` | `List<EgressEndpoint>` | no |
| `Id` | `string` | no |
| `PackagingGroupId` | `string` | no |
| `ResourceId` | `string` | no |
| `SourceArn` | `string` | no |
| `SourceRoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribePackagingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CmafPackage` | `CmafPackage` | no |
| `CreatedAt` | `string` | no |
| `DashPackage` | `DashPackage` | no |
| `HlsPackage` | `HlsPackage` | no |
| `Id` | `string` | no |
| `MssPackage` | `MssPackage` | no |
| `PackagingGroupId` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribePackagingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApproximateAssetCount` | `integer` | no |
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `CreatedAt` | `string` | no |
| `DomainName` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |

## ListAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PackagingGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assets` | `List<AssetShallow>` | no |
| `NextToken` | `string` | no |

## ListPackagingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PackagingGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PackagingConfigurations` | `List<PackagingConfiguration>` | no |

## ListPackagingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PackagingGroups` | `List<PackagingGroup>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
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


## UpdatePackagingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApproximateAssetCount` | `integer` | no |
| `Arn` | `string` | no |
| `Authorization` | `Authorization` | no |
| `CreatedAt` | `string` | no |
| `DomainName` | `string` | no |
| `EgressAccessLogs` | `EgressAccessLogs` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |

