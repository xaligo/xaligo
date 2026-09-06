# Amazon S3 Files

API version: 2025-05-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/s3files/2025-05-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |
| `fileSystemId` | `string` | yes |
| `posixUser` | `PosixUser` | no |
| `rootDirectory` | `RootDirectory` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPointArn` | `string` | yes |
| `accessPointId` | `string` | yes |
| `clientToken` | `string` | yes |
| `fileSystemId` | `string` | yes |
| `status` | `string` | yes |
| `ownerId` | `string` | yes |
| `posixUser` | `PosixUser` | no |
| `rootDirectory` | `RootDirectory` | no |
| `tags` | `List<Tag>` | no |
| `name` | `string` | no |

## CreateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucket` | `string` | yes |
| `prefix` | `string` | no |
| `clientToken` | `string` | no |
| `kmsKeyId` | `string` | no |
| `roleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `acceptBucketWarning` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTime` | `timestamp` | no |
| `fileSystemArn` | `string` | no |
| `fileSystemId` | `string` | no |
| `bucket` | `string` | no |
| `prefix` | `string` | no |
| `clientToken` | `string` | no |
| `kmsKeyId` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `roleArn` | `string` | no |
| `ownerId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `name` | `string` | no |

## CreateMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `subnetId` | `string` | yes |
| `ipv4Address` | `string` | no |
| `ipv6Address` | `string` | no |
| `ipAddressType` | `string` | no |
| `securityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `availabilityZoneId` | `string` | no |
| `ownerId` | `string` | yes |
| `mountTargetId` | `string` | yes |
| `fileSystemId` | `string` | no |
| `subnetId` | `string` | yes |
| `ipv4Address` | `string` | no |
| `ipv6Address` | `string` | no |
| `networkInterfaceId` | `string` | no |
| `vpcId` | `string` | no |
| `securityGroups` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |

## DeleteAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mountTargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPointArn` | `string` | yes |
| `accessPointId` | `string` | yes |
| `clientToken` | `string` | yes |
| `fileSystemId` | `string` | yes |
| `status` | `string` | yes |
| `ownerId` | `string` | yes |
| `posixUser` | `PosixUser` | no |
| `rootDirectory` | `RootDirectory` | no |
| `tags` | `List<Tag>` | no |
| `name` | `string` | no |

## GetFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTime` | `timestamp` | no |
| `fileSystemArn` | `string` | no |
| `fileSystemId` | `string` | no |
| `bucket` | `string` | no |
| `prefix` | `string` | no |
| `clientToken` | `string` | no |
| `kmsKeyId` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `roleArn` | `string` | no |
| `ownerId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `name` | `string` | no |

## GetFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `policy` | `string` | yes |

## GetMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mountTargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `availabilityZoneId` | `string` | no |
| `ownerId` | `string` | yes |
| `mountTargetId` | `string` | yes |
| `fileSystemId` | `string` | no |
| `subnetId` | `string` | yes |
| `ipv4Address` | `string` | no |
| `ipv6Address` | `string` | no |
| `networkInterfaceId` | `string` | no |
| `vpcId` | `string` | no |
| `securityGroups` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |

## GetSynchronizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `latestVersionNumber` | `integer` | no |
| `importDataRules` | `List<ImportDataRule>` | yes |
| `expirationDataRules` | `List<ExpirationDataRule>` | yes |

## ListAccessPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `accessPoints` | `List<ListAccessPointsDescription>` | yes |

## ListFileSystems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bucket` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `fileSystems` | `List<ListFileSystemsDescription>` | yes |

## ListMountTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | no |
| `accessPointId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `mountTargets` | `List<ListMountTargetsDescription>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `nextToken` | `string` | no |

## PutFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutSynchronizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fileSystemId` | `string` | yes |
| `latestVersionNumber` | `integer` | no |
| `importDataRules` | `List<ImportDataRule>` | yes |
| `expirationDataRules` | `List<ExpirationDataRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mountTargetId` | `string` | yes |
| `securityGroups` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `availabilityZoneId` | `string` | no |
| `ownerId` | `string` | yes |
| `mountTargetId` | `string` | yes |
| `fileSystemId` | `string` | no |
| `subnetId` | `string` | yes |
| `ipv4Address` | `string` | no |
| `ipv6Address` | `string` | no |
| `networkInterfaceId` | `string` | no |
| `vpcId` | `string` | no |
| `securityGroups` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |

