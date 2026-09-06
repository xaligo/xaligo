# AWS DataSync

API version: 2018-11-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/datasync/2018-11-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelTaskExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationKey` | `string` | yes |
| `AgentName` | `string` | no |
| `Tags` | `List<TagListEntry>` | no |
| `VpcEndpointId` | `string` | no |
| `SubnetArns` | `List<string>` | no |
| `SecurityGroupArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentArn` | `string` | no |

## CreateLocationAzureBlob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerUrl` | `string` | yes |
| `AuthenticationType` | `string` | yes |
| `SasConfiguration` | `AzureBlobSasConfiguration` | no |
| `BlobType` | `string` | no |
| `AccessTier` | `string` | no |
| `Subdirectory` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `Tags` | `List<TagListEntry>` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationEfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | no |
| `EfsFilesystemArn` | `string` | yes |
| `Ec2Config` | `Ec2Config` | yes |
| `Tags` | `List<TagListEntry>` | no |
| `AccessPointArn` | `string` | no |
| `FileSystemAccessRoleArn` | `string` | no |
| `InTransitEncryption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationFsxLustre

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FsxFilesystemArn` | `string` | yes |
| `SecurityGroupArns` | `List<string>` | yes |
| `Subdirectory` | `string` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationFsxOntap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Protocol` | `FsxProtocol` | yes |
| `SecurityGroupArns` | `List<string>` | yes |
| `StorageVirtualMachineArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationFsxOpenZfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FsxFilesystemArn` | `string` | yes |
| `Protocol` | `FsxProtocol` | yes |
| `SecurityGroupArns` | `List<string>` | yes |
| `Subdirectory` | `string` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationFsxWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | no |
| `FsxFilesystemArn` | `string` | yes |
| `SecurityGroupArns` | `List<string>` | yes |
| `Tags` | `List<TagListEntry>` | no |
| `User` | `string` | yes |
| `Domain` | `string` | no |
| `Password` | `string` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationHdfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | no |
| `NameNodes` | `List<HdfsNameNode>` | yes |
| `BlockSize` | `integer` | no |
| `ReplicationFactor` | `integer` | no |
| `KmsKeyProviderUri` | `string` | no |
| `QopConfiguration` | `QopConfiguration` | no |
| `AuthenticationType` | `string` | yes |
| `SimpleUser` | `string` | no |
| `KerberosPrincipal` | `string` | no |
| `KerberosKeytab` | `blob` | no |
| `KerberosKrb5Conf` | `blob` | no |
| `AgentArns` | `List<string>` | yes |
| `Tags` | `List<TagListEntry>` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationNfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | yes |
| `ServerHostname` | `string` | yes |
| `OnPremConfig` | `OnPremConfig` | yes |
| `MountOptions` | `NfsMountOptions` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationObjectStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerHostname` | `string` | yes |
| `ServerPort` | `integer` | no |
| `ServerProtocol` | `string` | no |
| `Subdirectory` | `string` | no |
| `BucketName` | `string` | yes |
| `AccessKey` | `string` | no |
| `SecretKey` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `Tags` | `List<TagListEntry>` | no |
| `ServerCertificate` | `blob` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | no |
| `S3BucketArn` | `string` | yes |
| `S3StorageClass` | `string` | no |
| `S3Config` | `S3Config` | yes |
| `AgentArns` | `List<string>` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateLocationSmb

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subdirectory` | `string` | yes |
| `ServerHostname` | `string` | yes |
| `User` | `string` | no |
| `Domain` | `string` | no |
| `Password` | `string` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |
| `AgentArns` | `List<string>` | yes |
| `MountOptions` | `SmbMountOptions` | no |
| `Tags` | `List<TagListEntry>` | no |
| `AuthenticationType` | `string` | no |
| `DnsIpAddresses` | `List<string>` | no |
| `KerberosPrincipal` | `string` | no |
| `KerberosKeytab` | `blob` | no |
| `KerberosKrb5Conf` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |

## CreateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceLocationArn` | `string` | yes |
| `DestinationLocationArn` | `string` | yes |
| `CloudWatchLogGroupArn` | `string` | no |
| `Name` | `string` | no |
| `Options` | `Options` | no |
| `Excludes` | `List<FilterRule>` | no |
| `Schedule` | `TaskSchedule` | no |
| `Tags` | `List<TagListEntry>` | no |
| `Includes` | `List<FilterRule>` | no |
| `ManifestConfig` | `ManifestConfig` | no |
| `TaskReportConfig` | `TaskReportConfig` | no |
| `TaskMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | no |

## DeleteAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentArn` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `LastConnectionTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `EndpointType` | `string` | no |
| `PrivateLinkConfig` | `PrivateLinkConfig` | no |
| `Platform` | `Platform` | no |

## DescribeLocationAzureBlob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `AuthenticationType` | `string` | no |
| `BlobType` | `string` | no |
| `AccessTier` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `ManagedSecretConfig` | `ManagedSecretConfig` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

## DescribeLocationEfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `Ec2Config` | `Ec2Config` | no |
| `CreationTime` | `timestamp` | no |
| `AccessPointArn` | `string` | no |
| `FileSystemAccessRoleArn` | `string` | no |
| `InTransitEncryption` | `string` | no |

## DescribeLocationFsxLustre

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `SecurityGroupArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |

## DescribeLocationFsxOntap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | no |
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `Protocol` | `FsxProtocol` | no |
| `SecurityGroupArns` | `List<string>` | no |
| `StorageVirtualMachineArn` | `string` | no |
| `FsxFilesystemArn` | `string` | no |

## DescribeLocationFsxOpenZfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `SecurityGroupArns` | `List<string>` | no |
| `Protocol` | `FsxProtocol` | no |
| `CreationTime` | `timestamp` | no |

## DescribeLocationFsxWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `SecurityGroupArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `User` | `string` | no |
| `Domain` | `string` | no |
| `ManagedSecretConfig` | `ManagedSecretConfig` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

## DescribeLocationHdfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `NameNodes` | `List<HdfsNameNode>` | no |
| `BlockSize` | `integer` | no |
| `ReplicationFactor` | `integer` | no |
| `KmsKeyProviderUri` | `string` | no |
| `QopConfiguration` | `QopConfiguration` | no |
| `AuthenticationType` | `string` | no |
| `SimpleUser` | `string` | no |
| `KerberosPrincipal` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `ManagedSecretConfig` | `ManagedSecretConfig` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

## DescribeLocationNfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `OnPremConfig` | `OnPremConfig` | no |
| `MountOptions` | `NfsMountOptions` | no |
| `CreationTime` | `timestamp` | no |

## DescribeLocationObjectStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `AccessKey` | `string` | no |
| `ServerPort` | `integer` | no |
| `ServerProtocol` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `ServerCertificate` | `blob` | no |
| `ManagedSecretConfig` | `ManagedSecretConfig` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

## DescribeLocationS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `S3StorageClass` | `string` | no |
| `S3Config` | `S3Config` | no |
| `AgentArns` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |

## DescribeLocationSmb

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | no |
| `LocationUri` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `User` | `string` | no |
| `Domain` | `string` | no |
| `MountOptions` | `SmbMountOptions` | no |
| `CreationTime` | `timestamp` | no |
| `DnsIpAddresses` | `List<string>` | no |
| `KerberosPrincipal` | `string` | no |
| `AuthenticationType` | `string` | no |
| `ManagedSecretConfig` | `ManagedSecretConfig` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

## DescribeTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | no |
| `Status` | `string` | no |
| `Name` | `string` | no |
| `CurrentTaskExecutionArn` | `string` | no |
| `SourceLocationArn` | `string` | no |
| `DestinationLocationArn` | `string` | no |
| `CloudWatchLogGroupArn` | `string` | no |
| `SourceNetworkInterfaceArns` | `List<string>` | no |
| `DestinationNetworkInterfaceArns` | `List<string>` | no |
| `Options` | `Options` | no |
| `Excludes` | `List<FilterRule>` | no |
| `Schedule` | `TaskSchedule` | no |
| `ErrorCode` | `string` | no |
| `ErrorDetail` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Includes` | `List<FilterRule>` | no |
| `ManifestConfig` | `ManifestConfig` | no |
| `TaskReportConfig` | `TaskReportConfig` | no |
| `ScheduleDetails` | `TaskScheduleDetails` | no |
| `TaskMode` | `string` | no |

## DescribeTaskExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutionArn` | `string` | no |
| `Status` | `string` | no |
| `Options` | `Options` | no |
| `Excludes` | `List<FilterRule>` | no |
| `Includes` | `List<FilterRule>` | no |
| `ManifestConfig` | `ManifestConfig` | no |
| `StartTime` | `timestamp` | no |
| `EstimatedFilesToTransfer` | `long` | no |
| `EstimatedBytesToTransfer` | `long` | no |
| `FilesTransferred` | `long` | no |
| `BytesWritten` | `long` | no |
| `BytesTransferred` | `long` | no |
| `BytesCompressed` | `long` | no |
| `Result` | `TaskExecutionResultDetail` | no |
| `TaskReportConfig` | `TaskReportConfig` | no |
| `FilesDeleted` | `long` | no |
| `FilesSkipped` | `long` | no |
| `FilesVerified` | `long` | no |
| `ReportResult` | `ReportResult` | no |
| `EstimatedFilesToDelete` | `long` | no |
| `TaskMode` | `string` | no |
| `FilesPrepared` | `long` | no |
| `FilesListed` | `TaskExecutionFilesListedDetail` | no |
| `FilesFailed` | `TaskExecutionFilesFailedDetail` | no |
| `EstimatedFoldersToDelete` | `long` | no |
| `EstimatedFoldersToTransfer` | `long` | no |
| `FoldersSkipped` | `long` | no |
| `FoldersPrepared` | `long` | no |
| `FoldersTransferred` | `long` | no |
| `FoldersVerified` | `long` | no |
| `FoldersDeleted` | `long` | no |
| `FoldersListed` | `TaskExecutionFoldersListedDetail` | no |
| `FoldersFailed` | `TaskExecutionFoldersFailedDetail` | no |
| `LaunchTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |

## ListAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Agents` | `List<AgentListEntry>` | no |
| `NextToken` | `string` | no |

## ListLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<LocationFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Locations` | `List<LocationListEntry>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<TagListEntry>` | no |
| `NextToken` | `string` | no |

## ListTaskExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutions` | `List<TaskExecutionListEntry>` | no |
| `NextToken` | `string` | no |

## ListTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<TaskFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tasks` | `List<TaskListEntry>` | no |
| `NextToken` | `string` | no |

## StartTaskExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |
| `OverrideOptions` | `Options` | no |
| `Includes` | `List<FilterRule>` | no |
| `Excludes` | `List<FilterRule>` | no |
| `ManifestConfig` | `ManifestConfig` | no |
| `TaskReportConfig` | `TaskReportConfig` | no |
| `Tags` | `List<TagListEntry>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutionArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<TagListEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Keys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentArn` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationAzureBlob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `AuthenticationType` | `string` | no |
| `SasConfiguration` | `AzureBlobSasConfiguration` | no |
| `BlobType` | `string` | no |
| `AccessTier` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationEfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `AccessPointArn` | `string` | no |
| `FileSystemAccessRoleArn` | `string` | no |
| `InTransitEncryption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationFsxLustre

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationFsxOntap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Protocol` | `FsxUpdateProtocol` | no |
| `Subdirectory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationFsxOpenZfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Protocol` | `FsxProtocol` | no |
| `Subdirectory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationFsxWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `Domain` | `string` | no |
| `User` | `string` | no |
| `Password` | `string` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationHdfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `NameNodes` | `List<HdfsNameNode>` | no |
| `BlockSize` | `integer` | no |
| `ReplicationFactor` | `integer` | no |
| `KmsKeyProviderUri` | `string` | no |
| `QopConfiguration` | `QopConfiguration` | no |
| `AuthenticationType` | `string` | no |
| `SimpleUser` | `string` | no |
| `KerberosPrincipal` | `string` | no |
| `KerberosKeytab` | `blob` | no |
| `KerberosKrb5Conf` | `blob` | no |
| `AgentArns` | `List<string>` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationNfs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `ServerHostname` | `string` | no |
| `OnPremConfig` | `OnPremConfig` | no |
| `MountOptions` | `NfsMountOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationObjectStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `ServerPort` | `integer` | no |
| `ServerProtocol` | `string` | no |
| `Subdirectory` | `string` | no |
| `ServerHostname` | `string` | no |
| `AccessKey` | `string` | no |
| `SecretKey` | `string` | no |
| `AgentArns` | `List<string>` | no |
| `ServerCertificate` | `blob` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `S3StorageClass` | `string` | no |
| `S3Config` | `S3Config` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLocationSmb

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationArn` | `string` | yes |
| `Subdirectory` | `string` | no |
| `ServerHostname` | `string` | no |
| `User` | `string` | no |
| `Domain` | `string` | no |
| `Password` | `string` | no |
| `CmkSecretConfig` | `CmkSecretConfig` | no |
| `CustomSecretConfig` | `CustomSecretConfig` | no |
| `AgentArns` | `List<string>` | no |
| `MountOptions` | `SmbMountOptions` | no |
| `AuthenticationType` | `string` | no |
| `DnsIpAddresses` | `List<string>` | no |
| `KerberosPrincipal` | `string` | no |
| `KerberosKeytab` | `blob` | no |
| `KerberosKrb5Conf` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |
| `Options` | `Options` | no |
| `Excludes` | `List<FilterRule>` | no |
| `Schedule` | `TaskSchedule` | no |
| `Name` | `string` | no |
| `CloudWatchLogGroupArn` | `string` | no |
| `Includes` | `List<FilterRule>` | no |
| `ManifestConfig` | `ManifestConfig` | no |
| `TaskReportConfig` | `TaskReportConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTaskExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskExecutionArn` | `string` | yes |
| `Options` | `Options` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


