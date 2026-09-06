# Application Migration Service

API version: 2020-02-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mgn/2020-02-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ArchiveApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `applicationAggregatedStatus` | `ApplicationAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |
| `waveID` | `string` | no |

## ArchiveWave

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `waveAggregatedStatus` | `WaveAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |

## AssociateApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `applicationIDs` | `List<string>` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `sourceServerIDs` | `List<string>` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ChangeServerLifeCycleState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `lifeCycle` | `ChangeServerLifeCycleStateSourceServerLifecycle` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `applicationAggregatedStatus` | `ApplicationAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |
| `waveID` | `string` | no |

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `ssmInstanceID` | `string` | yes |
| `tags` | `Map<string>` | no |
| `ssmCommandConfig` | `ConnectorSsmCommandConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorID` | `string` | no |
| `name` | `string` | no |
| `ssmInstanceID` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |
| `ssmCommandConfig` | `ConnectorSsmCommandConfig` | no |

## CreateLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |
| `tags` | `Map<string>` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `associatePublicIpAddress` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `smallVolumeMaxSize` | `long` | no |
| `smallVolumeConf` | `LaunchTemplateDiskConf` | no |
| `largeVolumeConf` | `LaunchTemplateDiskConf` | no |
| `enableParametersEncryption` | `boolean` | no |
| `parametersEncryptionKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |
| `tags` | `Map<string>` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `associatePublicIpAddress` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `smallVolumeMaxSize` | `long` | no |
| `smallVolumeConf` | `LaunchTemplateDiskConf` | no |
| `largeVolumeConf` | `LaunchTemplateDiskConf` | no |
| `enableParametersEncryption` | `boolean` | no |
| `parametersEncryptionKey` | `string` | no |

## CreateNetworkMigrationDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `sourceConfigurations` | `List<SourceConfiguration>` | no |
| `targetS3Configuration` | `TargetS3Configuration` | yes |
| `targetNetwork` | `TargetNetwork` | yes |
| `targetDeployment` | `string` | no |
| `vpcProvisioningStrategy` | `string` | no |
| `cidrMappings` | `List<CidrMapping>` | no |
| `tags` | `Map<string>` | no |
| `scopeTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `networkMigrationDefinitionID` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `sourceConfigurations` | `List<SourceConfiguration>` | no |
| `targetS3Configuration` | `TargetS3Configuration` | no |
| `targetNetwork` | `TargetNetwork` | no |
| `targetDeployment` | `string` | no |
| `vpcProvisioningStrategy` | `string` | no |
| `cidrMappings` | `List<CidrMapping>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `scopeTags` | `Map<string>` | no |

## CreateReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stagingAreaSubnetId` | `string` | yes |
| `associateDefaultSecurityGroup` | `boolean` | yes |
| `replicationServersSecurityGroupsIDs` | `List<string>` | yes |
| `replicationServerInstanceType` | `string` | yes |
| `useDedicatedReplicationServer` | `boolean` | yes |
| `defaultLargeStagingDiskType` | `string` | yes |
| `ebsEncryption` | `string` | yes |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | yes |
| `dataPlaneRouting` | `string` | yes |
| `createPublicIP` | `boolean` | yes |
| `stagingAreaTags` | `Map<string>` | yes |
| `useFipsEndpoint` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

## CreateWave

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `waveAggregatedStatus` | `WaveAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkMigrationDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVcenterClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vcenterClientID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWave

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeJobLogItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<JobLog>` | no |
| `nextToken` | `string` | no |

## DescribeJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeJobsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Job>` | no |
| `nextToken` | `string` | no |

## DescribeLaunchConfigurationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateIDs` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<LaunchConfigurationTemplate>` | no |
| `nextToken` | `string` | no |

## DescribeReplicationConfigurationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateIDs` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ReplicationConfigurationTemplate>` | no |
| `nextToken` | `string` | no |

## DescribeSourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeSourceServersRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SourceServer>` | no |
| `nextToken` | `string` | no |

## DescribeVcenterClients

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<VcenterClient>` | no |
| `nextToken` | `string` | no |

## DisassociateApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `applicationIDs` | `List<string>` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `sourceServerIDs` | `List<string>` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisconnectFromService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## FinalizeCutover

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## GetLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |

## GetNetworkMigrationDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `networkMigrationDefinitionID` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `sourceConfigurations` | `List<SourceConfiguration>` | no |
| `targetS3Configuration` | `TargetS3Configuration` | no |
| `targetNetwork` | `TargetNetwork` | no |
| `targetDeployment` | `string` | no |
| `vpcProvisioningStrategy` | `string` | no |
| `cidrMappings` | `List<CidrMapping>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `scopeTags` | `Map<string>` | no |

## GetNetworkMigrationMapperSegmentConstruct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |
| `networkMigrationExecutionID` | `string` | yes |
| `segmentID` | `string` | yes |
| `constructID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `construct` | `NetworkMigrationMapperSegmentConstruct` | no |

## GetReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

## InitializeService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListApplicationsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Application>` | no |
| `nextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListConnectorsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Connector>` | no |
| `nextToken` | `string` | no |

## ListExportErrors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ExportTaskError>` | no |
| `nextToken` | `string` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListExportsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ExportTask>` | no |
| `nextToken` | `string` | no |

## ListImportErrors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ImportTaskError>` | no |
| `nextToken` | `string` | no |

## ListImportFileEnrichments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListImportFileEnrichmentsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ImportFileEnrichment>` | no |
| `nextToken` | `string` | no |

## ListImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListImportsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ImportTask>` | no |
| `nextToken` | `string` | no |

## ListManagedAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ManagedAccount>` | yes |
| `nextToken` | `string` | no |

## ListNetworkMigrationAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationAnalysesFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationAnalysisJobDetails>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationAnalysisResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationAnalysisResultsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationAnalysisResult>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationCodeGenerationSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationCodeGenerationSegmentsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationCodeGenerationSegment>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationCodeGenerations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationCodeGenerationsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationCodeGenerationJobDetails>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListNetworkMigrationDefinitionsRequestFilters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationDefinitionSummary>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationDeployedStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationDeployedStackDetails>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationDeployerJobFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationDeployerJobDetails>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationExecutionRequestFilters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationExecution>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationMapperSegmentConstructs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `segmentID` | `string` | yes |
| `filters` | `ListNetworkMigrationMapperSegmentConstructsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationMapperSegmentConstruct>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationMapperSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationMapperSegmentsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationMapperSegment>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationMappingUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationMappingUpdatesFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationMappingUpdateJobDetails>` | no |
| `nextToken` | `string` | no |

## ListNetworkMigrationMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `filters` | `ListNetworkMigrationMappingsFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NetworkMigrationMappingJobDetails>` | no |
| `nextToken` | `string` | no |

## ListSourceServerActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `filters` | `SourceServerActionsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SourceServerActionDocument>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTemplateActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `filters` | `TemplateActionsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TemplateActionDocument>` | no |
| `nextToken` | `string` | no |

## ListWaves

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `ListWavesRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Wave>` | no |
| `nextToken` | `string` | no |

## MarkAsArchived

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## PauseReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## PutSourceServerAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `actionName` | `string` | yes |
| `documentIdentifier` | `string` | yes |
| `order` | `integer` | yes |
| `actionID` | `string` | yes |
| `documentVersion` | `string` | no |
| `active` | `boolean` | no |
| `timeoutSeconds` | `integer` | no |
| `mustSucceedForCutover` | `boolean` | no |
| `parameters` | `Map<List<SsmParameterStoreParameter>>` | no |
| `externalParameters` | `Map<SsmExternalParameter>` | no |
| `description` | `string` | no |
| `category` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionID` | `string` | no |
| `actionName` | `string` | no |
| `documentIdentifier` | `string` | no |
| `order` | `integer` | no |
| `documentVersion` | `string` | no |
| `active` | `boolean` | no |
| `timeoutSeconds` | `integer` | no |
| `mustSucceedForCutover` | `boolean` | no |
| `parameters` | `Map<List<SsmParameterStoreParameter>>` | no |
| `externalParameters` | `Map<SsmExternalParameter>` | no |
| `description` | `string` | no |
| `category` | `string` | no |

## PutTemplateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `actionName` | `string` | yes |
| `documentIdentifier` | `string` | yes |
| `order` | `integer` | yes |
| `actionID` | `string` | yes |
| `documentVersion` | `string` | no |
| `active` | `boolean` | no |
| `timeoutSeconds` | `integer` | no |
| `mustSucceedForCutover` | `boolean` | no |
| `parameters` | `Map<List<SsmParameterStoreParameter>>` | no |
| `operatingSystem` | `string` | no |
| `externalParameters` | `Map<SsmExternalParameter>` | no |
| `description` | `string` | no |
| `category` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionID` | `string` | no |
| `actionName` | `string` | no |
| `documentIdentifier` | `string` | no |
| `order` | `integer` | no |
| `documentVersion` | `string` | no |
| `active` | `boolean` | no |
| `timeoutSeconds` | `integer` | no |
| `mustSucceedForCutover` | `boolean` | no |
| `parameters` | `Map<List<SsmParameterStoreParameter>>` | no |
| `operatingSystem` | `string` | no |
| `externalParameters` | `Map<SsmExternalParameter>` | no |
| `description` | `string` | no |
| `category` | `string` | no |

## RemoveSourceServerAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `actionID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTemplateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `actionID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResumeReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## RetryDataReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## StartCutover

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerIDs` | `List<string>` | yes |
| `tags` | `Map<string>` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StartExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3Bucket` | `string` | yes |
| `s3Key` | `string` | yes |
| `s3BucketOwner` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportTask` | `ExportTask` | no |

## StartImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `s3BucketSource` | `S3BucketSource` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importTask` | `ImportTask` | no |

## StartImportFileEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `s3BucketSource` | `EnrichmentSourceS3Configuration` | yes |
| `s3BucketTarget` | `EnrichmentTargetS3Configuration` | yes |
| `ipAssignmentStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartNetworkMigrationAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartNetworkMigrationCodeGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `codeGenerationOutputFormatTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartNetworkMigrationDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartNetworkMigrationMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `securityGroupMappingStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartNetworkMigrationMappingUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationExecutionID` | `string` | yes |
| `networkMigrationDefinitionID` | `string` | yes |
| `constructs` | `List<StartNetworkMigrationMappingUpdateConstruct>` | no |
| `segments` | `List<StartNetworkMigrationMappingUpdateSegment>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |

## StartReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## StartTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerIDs` | `List<string>` | yes |
| `tags` | `Map<string>` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StopReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateTargetInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerIDs` | `List<string>` | yes |
| `tags` | `Map<string>` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## UnarchiveApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `applicationAggregatedStatus` | `ApplicationAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |
| `waveID` | `string` | no |

## UnarchiveWave

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `waveAggregatedStatus` | `WaveAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `applicationAggregatedStatus` | `ApplicationAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |
| `waveID` | `string` | no |

## UpdateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorID` | `string` | yes |
| `name` | `string` | no |
| `ssmCommandConfig` | `ConnectorSsmCommandConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorID` | `string` | no |
| `name` | `string` | no |
| `ssmInstanceID` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |
| `ssmCommandConfig` | `ConnectorSsmCommandConfig` | no |

## UpdateLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `name` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |

## UpdateLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `associatePublicIpAddress` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `smallVolumeMaxSize` | `long` | no |
| `smallVolumeConf` | `LaunchTemplateDiskConf` | no |
| `largeVolumeConf` | `LaunchTemplateDiskConf` | no |
| `enableParametersEncryption` | `boolean` | no |
| `parametersEncryptionKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `postLaunchActions` | `PostLaunchActions` | no |
| `enableMapAutoTagging` | `boolean` | no |
| `mapAutoTaggingMpeID` | `string` | no |
| `tags` | `Map<string>` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `associatePublicIpAddress` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `bootMode` | `string` | no |
| `smallVolumeMaxSize` | `long` | no |
| `smallVolumeConf` | `LaunchTemplateDiskConf` | no |
| `largeVolumeConf` | `LaunchTemplateDiskConf` | no |
| `enableParametersEncryption` | `boolean` | no |
| `parametersEncryptionKey` | `string` | no |

## UpdateNetworkMigrationDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sourceConfigurations` | `List<SourceConfiguration>` | no |
| `targetS3Configuration` | `TargetS3ConfigurationUpdate` | no |
| `targetNetwork` | `TargetNetworkUpdate` | no |
| `targetDeployment` | `string` | no |
| `vpcProvisioningStrategy` | `string` | no |
| `cidrMappings` | `List<CidrMapping>` | no |
| `scopeTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `networkMigrationDefinitionID` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `sourceConfigurations` | `List<SourceConfiguration>` | no |
| `targetS3Configuration` | `TargetS3Configuration` | no |
| `targetNetwork` | `TargetNetwork` | no |
| `targetDeployment` | `string` | no |
| `vpcProvisioningStrategy` | `string` | no |
| `cidrMappings` | `List<CidrMapping>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `scopeTags` | `Map<string>` | no |

## UpdateNetworkMigrationMapperSegment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkMigrationDefinitionID` | `string` | yes |
| `networkMigrationExecutionID` | `string` | yes |
| `segmentID` | `string` | yes |
| `scopeTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | no |
| `networkMigrationExecutionID` | `string` | no |
| `networkMigrationDefinitionID` | `string` | no |
| `segmentID` | `string` | no |
| `segmentType` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `logicalID` | `string` | no |
| `checksum` | `Checksum` | no |
| `outputS3Configuration` | `S3Configuration` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `scopeTags` | `Map<string>` | no |
| `targetAccount` | `string` | no |
| `referencedSegments` | `List<string>` | no |

## UpdateReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `accountID` | `string` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

## UpdateReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `useFipsEndpoint` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `internetProtocol` | `string` | no |
| `storeSnapshotOnLocalZone` | `boolean` | no |
| `storageConfiguration` | `StorageConfiguration` | no |

## UpdateSourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountID` | `string` | no |
| `sourceServerID` | `string` | yes |
| `connectorAction` | `SourceServerConnectorAction` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `platform` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## UpdateSourceServerReplicationType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `replicationType` | `string` | yes |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `isArchived` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `launchedInstance` | `LaunchedInstance` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `replicationType` | `string` | no |
| `vcenterClientID` | `string` | no |
| `applicationID` | `string` | no |
| `userProvidedID` | `string` | no |
| `fqdnForActionFramework` | `string` | no |
| `connectorAction` | `SourceServerConnectorAction` | no |

## UpdateWave

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `accountID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `waveID` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `isArchived` | `boolean` | no |
| `waveAggregatedStatus` | `WaveAggregatedStatus` | no |
| `creationDateTime` | `string` | no |
| `lastModifiedDateTime` | `string` | no |
| `tags` | `Map<string>` | no |

