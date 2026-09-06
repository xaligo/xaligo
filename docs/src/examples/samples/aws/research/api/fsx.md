# Amazon FSx

API version: 2018-03-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/fsx/2018-03-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateFileSystemAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileSystemId` | `string` | yes |
| `Aliases` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<Alias>` | no |

## CancelDataRepositoryTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Lifecycle` | `string` | no |
| `TaskId` | `string` | no |

## CopyBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `SourceBackupId` | `string` | yes |
| `SourceRegion` | `string` | no |
| `KmsKeyId` | `string` | no |
| `CopyTags` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backup` | `Backup` | no |

## CopySnapshotAndUpdateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `VolumeId` | `string` | yes |
| `SourceSnapshotARN` | `string` | yes |
| `CopyStrategy` | `string` | no |
| `Options` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | no |
| `Lifecycle` | `string` | no |
| `AdministrativeActions` | `List<AdministrativeAction>` | no |

## CreateAndAttachS3AccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `OpenZFSConfiguration` | `CreateAndAttachS3AccessPointOpenZFSConfiguration` | no |
| `OntapConfiguration` | `CreateAndAttachS3AccessPointOntapConfiguration` | no |
| `S3AccessPoint` | `CreateAndAttachS3AccessPointS3Configuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3AccessPointAttachment` | `S3AccessPointAttachment` | no |

## CreateBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `VolumeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backup` | `Backup` | no |

## CreateDataRepositoryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `FileSystemPath` | `string` | no |
| `DataRepositoryPath` | `string` | yes |
| `BatchImportMetaDataOnCreate` | `boolean` | no |
| `ImportedFileChunkSize` | `integer` | no |
| `S3` | `S3DataRepositoryConfiguration` | no |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `DataRepositoryAssociation` | no |

## CreateDataRepositoryTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `Paths` | `List<string>` | no |
| `FileSystemId` | `string` | yes |
| `Report` | `CompletionReport` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CapacityToRelease` | `long` | no |
| `ReleaseConfiguration` | `ReleaseConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataRepositoryTask` | `DataRepositoryTask` | no |

## CreateFileCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileCacheType` | `string` | yes |
| `FileCacheTypeVersion` | `string` | yes |
| `StorageCapacity` | `integer` | yes |
| `SubnetIds` | `List<string>` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `CopyTagsToDataRepositoryAssociations` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `LustreConfiguration` | `CreateFileCacheLustreConfiguration` | no |
| `DataRepositoryAssociations` | `List<FileCacheDataRepositoryAssociation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCache` | `FileCacheCreating` | no |

## CreateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileSystemType` | `string` | yes |
| `StorageCapacity` | `integer` | no |
| `StorageType` | `string` | no |
| `SubnetIds` | `List<string>` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `WindowsConfiguration` | `CreateFileSystemWindowsConfiguration` | no |
| `LustreConfiguration` | `CreateFileSystemLustreConfiguration` | no |
| `OntapConfiguration` | `CreateFileSystemOntapConfiguration` | no |
| `FileSystemTypeVersion` | `string` | no |
| `OpenZFSConfiguration` | `CreateFileSystemOpenZFSConfiguration` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystem` | `FileSystem` | no |

## CreateFileSystemFromBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `SubnetIds` | `List<string>` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `WindowsConfiguration` | `CreateFileSystemWindowsConfiguration` | no |
| `LustreConfiguration` | `CreateFileSystemLustreConfiguration` | no |
| `StorageType` | `string` | no |
| `KmsKeyId` | `string` | no |
| `FileSystemTypeVersion` | `string` | no |
| `OpenZFSConfiguration` | `CreateFileSystemOpenZFSConfiguration` | no |
| `StorageCapacity` | `integer` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystem` | `FileSystem` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `Name` | `string` | yes |
| `VolumeId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateStorageVirtualMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveDirectoryConfiguration` | `CreateSvmActiveDirectoryConfiguration` | no |
| `ClientRequestToken` | `string` | no |
| `FileSystemId` | `string` | yes |
| `Name` | `string` | yes |
| `SvmAdminPassword` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `RootVolumeSecurityStyle` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageVirtualMachine` | `StorageVirtualMachine` | no |

## CreateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `VolumeType` | `string` | yes |
| `Name` | `string` | yes |
| `OntapConfiguration` | `CreateOntapVolumeConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `OpenZFSConfiguration` | `CreateOpenZFSVolumeConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volume` | `Volume` | no |

## CreateVolumeFromBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Name` | `string` | yes |
| `OntapConfiguration` | `CreateOntapVolumeConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volume` | `Volume` | no |

## DeleteBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | no |
| `Lifecycle` | `string` | no |

## DeleteDataRepositoryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `DeleteDataInFileSystem` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |
| `Lifecycle` | `string` | no |
| `DeleteDataInFileSystem` | `boolean` | no |

## DeleteFileCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCacheId` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCacheId` | `string` | no |
| `Lifecycle` | `string` | no |

## DeleteFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `WindowsConfiguration` | `DeleteFileSystemWindowsConfiguration` | no |
| `LustreConfiguration` | `DeleteFileSystemLustreConfiguration` | no |
| `OpenZFSConfiguration` | `DeleteFileSystemOpenZFSConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | no |
| `Lifecycle` | `string` | no |
| `WindowsResponse` | `DeleteFileSystemWindowsResponse` | no |
| `LustreResponse` | `DeleteFileSystemLustreResponse` | no |
| `OpenZFSResponse` | `DeleteFileSystemOpenZFSResponse` | no |

## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `SnapshotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `Lifecycle` | `string` | no |

## DeleteStorageVirtualMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `StorageVirtualMachineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageVirtualMachineId` | `string` | no |
| `Lifecycle` | `string` | no |

## DeleteVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `VolumeId` | `string` | yes |
| `OntapConfiguration` | `DeleteVolumeOntapConfiguration` | no |
| `OpenZFSConfiguration` | `DeleteVolumeOpenZFSConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | no |
| `Lifecycle` | `string` | no |
| `OntapResponse` | `DeleteVolumeOntapResponse` | no |

## DescribeBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backups` | `List<Backup>` | no |
| `NextToken` | `string` | no |

## DescribeDataRepositoryAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<DataRepositoryAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeDataRepositoryTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskIds` | `List<string>` | no |
| `Filters` | `List<DataRepositoryTaskFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataRepositoryTasks` | `List<DataRepositoryTask>` | no |
| `NextToken` | `string` | no |

## DescribeFileCaches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCacheIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCaches` | `List<FileCache>` | no |
| `NextToken` | `string` | no |

## DescribeFileSystemAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileSystemId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<Alias>` | no |
| `NextToken` | `string` | no |

## DescribeFileSystems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystems` | `List<FileSystem>` | no |
| `NextToken` | `string` | no |

## DescribeS3AccessPointAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `Filters` | `List<S3AccessPointAttachmentsFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3AccessPointAttachments` | `List<S3AccessPointAttachment>` | no |
| `NextToken` | `string` | no |

## DescribeSharedVpcConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableFsxRouteTableUpdatesFromParticipantAccounts` | `string` | no |

## DescribeSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotIds` | `List<string>` | no |
| `Filters` | `List<SnapshotFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeShared` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshots` | `List<Snapshot>` | no |
| `NextToken` | `string` | no |

## DescribeStorageVirtualMachines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageVirtualMachineIds` | `List<string>` | no |
| `Filters` | `List<StorageVirtualMachineFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageVirtualMachines` | `List<StorageVirtualMachine>` | no |
| `NextToken` | `string` | no |

## DescribeVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeIds` | `List<string>` | no |
| `Filters` | `List<VolumeFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volumes` | `List<Volume>` | no |
| `NextToken` | `string` | no |

## DetachAndDeleteS3AccessPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Lifecycle` | `string` | no |
| `Name` | `string` | no |

## DisassociateFileSystemAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileSystemId` | `string` | yes |
| `Aliases` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<Alias>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ReleaseFileSystemNfsV3Locks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystem` | `FileSystem` | no |

## RestoreVolumeFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `VolumeId` | `string` | yes |
| `SnapshotId` | `string` | yes |
| `Options` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | no |
| `Lifecycle` | `string` | no |
| `AdministrativeActions` | `List<AdministrativeAction>` | no |

## StartMisconfiguredStateRecovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `FileSystemId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystem` | `FileSystem` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataRepositoryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `ImportedFileChunkSize` | `integer` | no |
| `S3` | `S3DataRepositoryConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `DataRepositoryAssociation` | no |

## UpdateFileCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCacheId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `LustreConfiguration` | `UpdateFileCacheLustreConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileCache` | `FileCache` | no |

## UpdateFileSystem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystemId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `StorageCapacity` | `integer` | no |
| `WindowsConfiguration` | `UpdateFileSystemWindowsConfiguration` | no |
| `LustreConfiguration` | `UpdateFileSystemLustreConfiguration` | no |
| `OntapConfiguration` | `UpdateFileSystemOntapConfiguration` | no |
| `OpenZFSConfiguration` | `UpdateFileSystemOpenZFSConfiguration` | no |
| `StorageType` | `string` | no |
| `FileSystemTypeVersion` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileSystem` | `FileSystem` | no |

## UpdateSharedVpcConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableFsxRouteTableUpdatesFromParticipantAccounts` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableFsxRouteTableUpdatesFromParticipantAccounts` | `string` | no |

## UpdateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `Name` | `string` | yes |
| `SnapshotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## UpdateStorageVirtualMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveDirectoryConfiguration` | `UpdateSvmActiveDirectoryConfiguration` | no |
| `ClientRequestToken` | `string` | no |
| `StorageVirtualMachineId` | `string` | yes |
| `SvmAdminPassword` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageVirtualMachine` | `StorageVirtualMachine` | no |

## UpdateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `VolumeId` | `string` | yes |
| `OntapConfiguration` | `UpdateOntapVolumeConfiguration` | no |
| `Name` | `string` | no |
| `OpenZFSConfiguration` | `UpdateOpenZFSVolumeConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volume` | `Volume` | no |

