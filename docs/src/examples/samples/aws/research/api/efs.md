# Amazon Elastic File System

API version: 2015-02-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/efs/2015-02-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `FileSystemId` | `string` | yes |
| `PosixUser` | `PosixUser` | no |
| `RootDirectory` | `RootDirectory` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `AccessPointId` | `string` | no |
| `AccessPointArn` | `string` | no |
| `FileSystemId` | `string` | no |
| `PosixUser` | `PosixUser` | no |
| `RootDirectory` | `RootDirectory` | no |
| `OwnerId` | `string` | no |
| `LifeCycleState` | `string` | no |

## CreateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationToken` | `string` | yes |
| `PerformanceMode` | `string` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `ThroughputMode` | `string` | no |
| `ProvisionedThroughputInMibps` | `double` | no |
| `AvailabilityZoneName` | `string` | no |
| `Backup` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerId` | `string` | yes |
| `CreationToken` | `string` | yes |
| `FileSystemId` | `string` | yes |
| `FileSystemArn` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LifeCycleState` | `string` | yes |
| `Name` | `string` | no |
| `NumberOfMountTargets` | `integer` | yes |
| `SizeInBytes` | `FileSystemSize` | yes |
| `PerformanceMode` | `string` | yes |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `ThroughputMode` | `string` | no |
| `ProvisionedThroughputInMibps` | `double` | no |
| `AvailabilityZoneName` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `Tags` | `List<Tag>` | yes |
| `FileSystemProtection` | `FileSystemProtectionDescription` | no |

## CreateMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `SubnetId` | `string` | yes |
| `IpAddress` | `string` | no |
| `Ipv6Address` | `string` | no |
| `IpAddressType` | `string` | no |
| `SecurityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerId` | `string` | no |
| `MountTargetId` | `string` | yes |
| `FileSystemId` | `string` | yes |
| `SubnetId` | `string` | yes |
| `LifeCycleState` | `string` | yes |
| `IpAddress` | `string` | no |
| `Ipv6Address` | `string` | no |
| `NetworkInterfaceId` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `AvailabilityZoneName` | `string` | no |
| `VpcId` | `string` | no |

## CreateReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceFileSystemId` | `string` | yes |
| `Destinations` | `List<DestinationToCreate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceFileSystemId` | `string` | yes |
| `SourceFileSystemRegion` | `string` | yes |
| `SourceFileSystemArn` | `string` | yes |
| `OriginalSourceFileSystemArn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `Destinations` | `List<Destination>` | yes |
| `SourceFileSystemOwnerId` | `string` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMountTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MountTargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceFileSystemId` | `string` | yes |
| `DeletionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccessPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccessPointId` | `string` | no |
| `FileSystemId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPoints` | `List<AccessPointDescription>` | no |
| `NextToken` | `string` | no |

## DescribeAccountPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdPreference` | `ResourceIdPreference` | no |
| `NextToken` | `string` | no |

## DescribeBackupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPolicy` | `BackupPolicy` | no |

## DescribeFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | no |
| `Policy` | `string` | no |

## DescribeFileSystems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |
| `CreationToken` | `string` | no |
| `FileSystemId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `FileSystems` | `List<FileSystemDescription>` | no |
| `NextMarker` | `string` | no |

## DescribeLifecycleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecyclePolicies` | `List<LifecyclePolicy>` | no |

## DescribeMountTargetSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MountTargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroups` | `List<string>` | yes |

## DescribeMountTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |
| `FileSystemId` | `string` | no |
| `MountTargetId` | `string` | no |
| `AccessPointId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MountTargets` | `List<MountTargetDescription>` | no |
| `NextMarker` | `string` | no |

## DescribeReplicationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Replications` | `List<ReplicationConfigurationDescription>` | no |
| `NextToken` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Tags` | `List<Tag>` | yes |
| `NextMarker` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ModifyMountTargetSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MountTargetId` | `string` | yes |
| `SecurityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccountPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdPreference` | `ResourceIdPreference` | no |

## PutBackupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `BackupPolicy` | `BackupPolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupPolicy` | `BackupPolicy` | no |

## PutFileSystemPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `Policy` | `string` | yes |
| `BypassPolicyLockoutSafetyCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | no |
| `Policy` | `string` | no |

## PutLifecycleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `LifecyclePolicies` | `List<LifecyclePolicy>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecyclePolicies` | `List<LifecyclePolicy>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `ThroughputMode` | `string` | no |
| `ProvisionedThroughputInMibps` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerId` | `string` | yes |
| `CreationToken` | `string` | yes |
| `FileSystemId` | `string` | yes |
| `FileSystemArn` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LifeCycleState` | `string` | yes |
| `Name` | `string` | no |
| `NumberOfMountTargets` | `integer` | yes |
| `SizeInBytes` | `FileSystemSize` | yes |
| `PerformanceMode` | `string` | yes |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `ThroughputMode` | `string` | no |
| `ProvisionedThroughputInMibps` | `double` | no |
| `AvailabilityZoneName` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `Tags` | `List<Tag>` | yes |
| `FileSystemProtection` | `FileSystemProtectionDescription` | no |

## UpdateFileSystemProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `ReplicationOverwriteProtection` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationOverwriteProtection` | `string` | no |

