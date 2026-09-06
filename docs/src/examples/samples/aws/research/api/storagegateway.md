# AWS Storage Gateway

API version: 2013-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/storagegateway/2013-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationKey` | `string` | yes |
| `GatewayName` | `string` | yes |
| `GatewayTimezone` | `string` | yes |
| `GatewayRegion` | `string` | yes |
| `GatewayType` | `string` | no |
| `TapeDriveType` | `string` | no |
| `MediumChangerType` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## AddCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `DiskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |

## AddUploadBuffer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `DiskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## AddWorkingStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `DiskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## AssignTapePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | yes |
| `PoolId` | `string` | yes |
| `BypassGovernanceRetention` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## AssociateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `Password` | `string` | yes |
| `ClientToken` | `string` | yes |
| `GatewayARN` | `string` | yes |
| `LocationARN` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `AuditDestinationARN` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |
| `EndpointNetworkConfiguration` | `EndpointNetworkConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARN` | `string` | no |

## AttachVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TargetName` | `string` | no |
| `VolumeARN` | `string` | yes |
| `NetworkInterfaceId` | `string` | yes |
| `DiskId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |
| `TargetARN` | `string` | no |

## CancelArchival

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## CancelCacheReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | no |

## CancelRetrieval

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## CreateCachediSCSIVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `VolumeSizeInBytes` | `long` | yes |
| `SnapshotId` | `string` | no |
| `TargetName` | `string` | yes |
| `SourceVolumeARN` | `string` | no |
| `NetworkInterfaceId` | `string` | yes |
| `ClientToken` | `string` | yes |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |
| `TargetARN` | `string` | no |

## CreateNFSFileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `NFSFileShareDefaults` | `NFSFileShareDefaults` | no |
| `GatewayARN` | `string` | yes |
| `EncryptionType` | `string` | no |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `Role` | `string` | yes |
| `LocationARN` | `string` | yes |
| `DefaultStorageClass` | `string` | no |
| `ObjectACL` | `string` | no |
| `ClientList` | `List<string>` | no |
| `Squash` | `string` | no |
| `ReadOnly` | `boolean` | no |
| `GuessMIMETypeEnabled` | `boolean` | no |
| `RequesterPays` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `FileShareName` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |
| `NotificationPolicy` | `string` | no |
| `VPCEndpointDNSName` | `string` | no |
| `BucketRegion` | `string` | no |
| `AuditDestinationARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |

## CreateSMBFileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | yes |
| `GatewayARN` | `string` | yes |
| `EncryptionType` | `string` | no |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `Role` | `string` | yes |
| `LocationARN` | `string` | yes |
| `DefaultStorageClass` | `string` | no |
| `ObjectACL` | `string` | no |
| `ReadOnly` | `boolean` | no |
| `GuessMIMETypeEnabled` | `boolean` | no |
| `RequesterPays` | `boolean` | no |
| `SMBACLEnabled` | `boolean` | no |
| `AccessBasedEnumeration` | `boolean` | no |
| `AdminUserList` | `List<string>` | no |
| `ValidUserList` | `List<string>` | no |
| `InvalidUserList` | `List<string>` | no |
| `AuditDestinationARN` | `string` | no |
| `Authentication` | `string` | no |
| `CaseSensitivity` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `FileShareName` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |
| `NotificationPolicy` | `string` | no |
| `VPCEndpointDNSName` | `string` | no |
| `BucketRegion` | `string` | no |
| `OplocksEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |
| `SnapshotDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |
| `SnapshotId` | `string` | no |

## CreateSnapshotFromVolumeRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |
| `SnapshotDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `VolumeARN` | `string` | no |
| `VolumeRecoveryPointTime` | `string` | no |

## CreateStorediSCSIVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `DiskId` | `string` | yes |
| `SnapshotId` | `string` | no |
| `PreserveExistingData` | `boolean` | yes |
| `TargetName` | `string` | yes |
| `NetworkInterfaceId` | `string` | yes |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |
| `VolumeSizeInBytes` | `long` | no |
| `TargetARN` | `string` | no |

## CreateTapePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `StorageClass` | `string` | yes |
| `RetentionLockType` | `string` | no |
| `RetentionLockTimeInDays` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolARN` | `string` | no |

## CreateTapeWithBarcode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeSizeInBytes` | `long` | yes |
| `TapeBarcode` | `string` | yes |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `PoolId` | `string` | no |
| `Worm` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## CreateTapes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeSizeInBytes` | `long` | yes |
| `ClientToken` | `string` | yes |
| `NumTapesToCreate` | `integer` | yes |
| `TapeBarcodePrefix` | `string` | yes |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `PoolId` | `string` | no |
| `Worm` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARNs` | `List<string>` | no |

## DeleteAutomaticTapeCreationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## DeleteBandwidthRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `BandwidthType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## DeleteCacheReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | no |

## DeleteChapCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetARN` | `string` | yes |
| `InitiatorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetARN` | `string` | no |
| `InitiatorName` | `string` | no |

## DeleteFileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `ForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |

## DeleteGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## DeleteSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |

## DeleteTape

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeARN` | `string` | yes |
| `BypassGovernanceRetention` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## DeleteTapeArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | yes |
| `BypassGovernanceRetention` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## DeleteTapePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolARN` | `string` | no |

## DeleteVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |

## DescribeAvailabilityMonitorTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Status` | `string` | no |
| `StartTime` | `timestamp` | no |

## DescribeBandwidthRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `AverageUploadRateLimitInBitsPerSec` | `long` | no |
| `AverageDownloadRateLimitInBitsPerSec` | `long` | no |

## DescribeBandwidthRateLimitSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `BandwidthRateLimitIntervals` | `List<BandwidthRateLimitInterval>` | no |

## DescribeCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `DiskIds` | `List<string>` | no |
| `CacheAllocatedInBytes` | `long` | no |
| `CacheUsedPercentage` | `double` | no |
| `CacheDirtyPercentage` | `double` | no |
| `CacheHitPercentage` | `double` | no |
| `CacheMissPercentage` | `double` | no |

## DescribeCacheReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportInfo` | `CacheReportInfo` | no |

## DescribeCachediSCSIVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARNs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachediSCSIVolumes` | `List<CachediSCSIVolume>` | no |

## DescribeChapCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChapCredentials` | `List<ChapInfo>` | no |

## DescribeFileSystemAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARNList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationInfoList` | `List<FileSystemAssociationInfo>` | no |

## DescribeGatewayInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `GatewayId` | `string` | no |
| `GatewayName` | `string` | no |
| `GatewayTimezone` | `string` | no |
| `GatewayState` | `string` | no |
| `GatewayNetworkInterfaces` | `List<NetworkInterface>` | no |
| `GatewayType` | `string` | no |
| `NextUpdateAvailabilityDate` | `string` | no |
| `LastSoftwareUpdate` | `string` | no |
| `Ec2InstanceId` | `string` | no |
| `Ec2InstanceRegion` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `VPCEndpoint` | `string` | no |
| `CloudWatchLogGroupARN` | `string` | no |
| `HostEnvironment` | `string` | no |
| `EndpointType` | `string` | no |
| `SoftwareUpdatesEndDate` | `string` | no |
| `DeprecationDate` | `string` | no |
| `GatewayCapacity` | `string` | no |
| `SupportedGatewayCapacities` | `List<string>` | no |
| `HostEnvironmentId` | `string` | no |
| `SoftwareVersion` | `string` | no |

## DescribeMaintenanceStartTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `HourOfDay` | `integer` | no |
| `MinuteOfHour` | `integer` | no |
| `DayOfWeek` | `integer` | no |
| `DayOfMonth` | `integer` | no |
| `Timezone` | `string` | no |
| `SoftwareUpdatePreferences` | `SoftwareUpdatePreferences` | no |

## DescribeNFSFileShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARNList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NFSFileShareInfoList` | `List<NFSFileShareInfo>` | no |

## DescribeSMBFileShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARNList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMBFileShareInfoList` | `List<SMBFileShareInfo>` | no |

## DescribeSMBSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `DomainName` | `string` | no |
| `ActiveDirectoryStatus` | `string` | no |
| `SMBGuestPasswordSet` | `boolean` | no |
| `SMBSecurityStrategy` | `string` | no |
| `FileSharesVisible` | `boolean` | no |
| `SMBLocalGroups` | `SMBLocalGroups` | no |

## DescribeSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |
| `StartAt` | `integer` | no |
| `RecurrenceInHours` | `integer` | no |
| `Description` | `string` | no |
| `Timezone` | `string` | no |
| `Tags` | `List<Tag>` | no |

## DescribeStorediSCSIVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARNs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorediSCSIVolumes` | `List<StorediSCSIVolume>` | no |

## DescribeTapeArchives

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARNs` | `List<string>` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeArchives` | `List<TapeArchive>` | no |
| `Marker` | `string` | no |

## DescribeTapeRecoveryPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `TapeRecoveryPointInfos` | `List<TapeRecoveryPointInfo>` | no |
| `Marker` | `string` | no |

## DescribeTapes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `TapeARNs` | `List<string>` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tapes` | `List<Tape>` | no |
| `Marker` | `string` | no |

## DescribeUploadBuffer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `DiskIds` | `List<string>` | no |
| `UploadBufferUsedInBytes` | `long` | no |
| `UploadBufferAllocatedInBytes` | `long` | no |

## DescribeVTLDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `VTLDeviceARNs` | `List<string>` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `VTLDevices` | `List<VTLDevice>` | no |
| `Marker` | `string` | no |

## DescribeWorkingStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `DiskIds` | `List<string>` | no |
| `WorkingStorageUsedInBytes` | `long` | no |
| `WorkingStorageAllocatedInBytes` | `long` | no |

## DetachVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |
| `ForceDetach` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |

## DisableGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## DisassociateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARN` | `string` | yes |
| `ForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARN` | `string` | no |

## EvictFilesFailingUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `ForceRemove` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationId` | `string` | no |

## JoinDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `DomainName` | `string` | yes |
| `OrganizationalUnit` | `string` | no |
| `DomainControllers` | `List<string>` | no |
| `TimeoutInSeconds` | `integer` | no |
| `UserName` | `string` | yes |
| `Password` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `ActiveDirectoryStatus` | `string` | no |

## ListAutomaticTapeCreationPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomaticTapeCreationPolicyInfos` | `List<AutomaticTapeCreationPolicyInfo>` | no |

## ListCacheReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportList` | `List<CacheReportInfo>` | no |
| `Marker` | `string` | no |

## ListFileShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `NextMarker` | `string` | no |
| `FileShareInfoList` | `List<FileShareInfo>` | no |

## ListFileSystemAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Limit` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `NextMarker` | `string` | no |
| `FileSystemAssociationSummaryList` | `List<FileSystemAssociationSummary>` | no |

## ListGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateways` | `List<GatewayInfo>` | no |
| `Marker` | `string` | no |

## ListLocalDisks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Disks` | `List<Disk>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |
| `Marker` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListTapePools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolARNs` | `List<string>` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolInfos` | `List<PoolInfo>` | no |
| `Marker` | `string` | no |

## ListTapes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARNs` | `List<string>` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeInfos` | `List<TapeInfo>` | no |
| `Marker` | `string` | no |

## ListVolumeInitiators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Initiators` | `List<string>` | no |

## ListVolumeRecoveryPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `VolumeRecoveryPointInfos` | `List<VolumeRecoveryPointInfo>` | no |

## ListVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Marker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `Marker` | `string` | no |
| `VolumeInfos` | `List<VolumeInfo>` | no |

## NotifyWhenUploaded

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |
| `NotificationId` | `string` | no |

## RefreshCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `FolderList` | `List<string>` | no |
| `Recursive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |
| `NotificationId` | `string` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |

## ResetCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## RetrieveTapeArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | yes |
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## RetrieveTapeRecoveryPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | yes |
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TapeARN` | `string` | no |

## SetLocalConsolePassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `LocalConsolePassword` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## SetSMBGuestPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `Password` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## ShutdownGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## StartAvailabilityMonitorTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## StartCacheReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `Role` | `string` | yes |
| `LocationARN` | `string` | yes |
| `BucketRegion` | `string` | yes |
| `VPCEndpointDNSName` | `string` | no |
| `InclusionFilters` | `List<CacheReportFilter>` | no |
| `ExclusionFilters` | `List<CacheReportFilter>` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheReportARN` | `string` | no |

## StartGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateAutomaticTapeCreationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomaticTapeCreationRules` | `List<AutomaticTapeCreationRule>` | yes |
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateBandwidthRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `AverageUploadRateLimitInBitsPerSec` | `long` | no |
| `AverageDownloadRateLimitInBitsPerSec` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateBandwidthRateLimitSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `BandwidthRateLimitIntervals` | `List<BandwidthRateLimitInterval>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateChapCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetARN` | `string` | yes |
| `SecretToAuthenticateInitiator` | `string` | yes |
| `InitiatorName` | `string` | yes |
| `SecretToAuthenticateTarget` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetARN` | `string` | no |
| `InitiatorName` | `string` | no |

## UpdateFileSystemAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARN` | `string` | yes |
| `UserName` | `string` | no |
| `Password` | `string` | no |
| `AuditDestinationARN` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemAssociationARN` | `string` | no |

## UpdateGatewayInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `GatewayName` | `string` | no |
| `GatewayTimezone` | `string` | no |
| `CloudWatchLogGroupARN` | `string` | no |
| `GatewayCapacity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |
| `GatewayName` | `string` | no |

## UpdateGatewaySoftwareNow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateMaintenanceStartTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `HourOfDay` | `integer` | no |
| `MinuteOfHour` | `integer` | no |
| `DayOfWeek` | `integer` | no |
| `DayOfMonth` | `integer` | no |
| `SoftwareUpdatePreferences` | `SoftwareUpdatePreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateNFSFileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `EncryptionType` | `string` | no |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `NFSFileShareDefaults` | `NFSFileShareDefaults` | no |
| `DefaultStorageClass` | `string` | no |
| `ObjectACL` | `string` | no |
| `ClientList` | `List<string>` | no |
| `Squash` | `string` | no |
| `ReadOnly` | `boolean` | no |
| `GuessMIMETypeEnabled` | `boolean` | no |
| `RequesterPays` | `boolean` | no |
| `FileShareName` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |
| `NotificationPolicy` | `string` | no |
| `AuditDestinationARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |

## UpdateSMBFileShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | yes |
| `EncryptionType` | `string` | no |
| `KMSEncrypted` | `boolean` | no |
| `KMSKey` | `string` | no |
| `DefaultStorageClass` | `string` | no |
| `ObjectACL` | `string` | no |
| `ReadOnly` | `boolean` | no |
| `GuessMIMETypeEnabled` | `boolean` | no |
| `RequesterPays` | `boolean` | no |
| `SMBACLEnabled` | `boolean` | no |
| `AccessBasedEnumeration` | `boolean` | no |
| `AdminUserList` | `List<string>` | no |
| `ValidUserList` | `List<string>` | no |
| `InvalidUserList` | `List<string>` | no |
| `AuditDestinationARN` | `string` | no |
| `CaseSensitivity` | `string` | no |
| `FileShareName` | `string` | no |
| `CacheAttributes` | `CacheAttributes` | no |
| `NotificationPolicy` | `string` | no |
| `OplocksEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileShareARN` | `string` | no |

## UpdateSMBFileShareVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `FileSharesVisible` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateSMBLocalGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `SMBLocalGroups` | `SMBLocalGroups` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateSMBSecurityStrategy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | yes |
| `SMBSecurityStrategy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayARN` | `string` | no |

## UpdateSnapshotSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | yes |
| `StartAt` | `integer` | yes |
| `RecurrenceInHours` | `integer` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeARN` | `string` | no |

## UpdateVTLDeviceType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VTLDeviceARN` | `string` | yes |
| `DeviceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VTLDeviceARN` | `string` | no |

