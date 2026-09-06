# odb

API version: 2024-08-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/odb/2024-08-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptMarketplaceRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceRegistrationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateIamRoleToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `iamRoleArn` | `string` | yes |
| `awsIntegration` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateVirtualMachinesToExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmClusterId` | `string` | yes |
| `desiredNodeCount` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exadbVmClusterId` | `string` | yes |

## CreateAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetworkId` | `string` | no |
| `displayName` | `string` | no |
| `dbName` | `string` | no |
| `adminPassword` | `string` | no |
| `computeCount` | `double` | no |
| `dataStorageSizeInTBs` | `integer` | no |
| `dataStorageSizeInGBs` | `integer` | no |
| `dbWorkload` | `string` | no |
| `isAutoScalingEnabled` | `boolean` | no |
| `isAutoScalingForStorageEnabled` | `boolean` | no |
| `licenseModel` | `string` | no |
| `characterSet` | `string` | no |
| `ncharacterSet` | `string` | no |
| `dbVersion` | `string` | no |
| `databaseEdition` | `string` | no |
| `standbyAllowlistedIpsSource` | `string` | no |
| `autonomousMaintenanceScheduleType` | `string` | no |
| `backupRetentionPeriodInDays` | `integer` | no |
| `byolComputeCountLimit` | `double` | no |
| `cpuCoreCount` | `integer` | no |
| `customerContactsToSendToOCI` | `List<CustomerContact>` | no |
| `privateEndpointIp` | `string` | no |
| `privateEndpointLabel` | `string` | no |
| `resourcePoolLeaderId` | `string` | no |
| `resourcePoolSummary` | `ResourcePoolSummary` | no |
| `scheduledOperations` | `List<ScheduledOperationDetails>` | no |
| `standbyAllowlistedIps` | `List<string>` | no |
| `allowlistedIps` | `List<string>` | no |
| `transportableTablespace` | `TransportableTablespace` | no |
| `isBackupRetentionLocked` | `boolean` | no |
| `isLocalDataGuardEnabled` | `boolean` | no |
| `isMtlsConnectionRequired` | `boolean` | no |
| `dbToolsDetails` | `List<DatabaseTool>` | no |
| `source` | `string` | no |
| `sourceConfiguration` | `SourceConfiguration` | no |
| `encryptionKeyProvider` | `string` | no |
| `encryptionKeyConfiguration` | `EncryptionKeyConfigurationInput` | no |
| `adminPasswordSource` | `string` | no |
| `adminPasswordSourceConfiguration` | `AdminPasswordSourceConfigurationInput` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## CreateAutonomousDatabaseBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `retentionPeriodInDays` | `integer` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `autonomousDatabaseBackupId` | `string` | yes |

## CreateAutonomousDatabaseWallet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `walletType` | `string` | no |
| `password` | `string` | no |
| `passwordSource` | `string` | no |
| `passwordSourceConfiguration` | `WalletPasswordSourceConfigurationInput` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseWalletFile` | `blob` | yes |

## CreateCloudAutonomousVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `odbNetworkId` | `string` | yes |
| `displayName` | `string` | yes |
| `clientToken` | `string` | no |
| `autonomousDataStorageSizeInTBs` | `double` | yes |
| `cpuCoreCountPerNode` | `integer` | yes |
| `dbServers` | `List<string>` | no |
| `description` | `string` | no |
| `isMtlsEnabledVmCluster` | `boolean` | no |
| `licenseModel` | `string` | no |
| `maintenanceWindow` | `MaintenanceWindow` | no |
| `memoryPerOracleComputeUnitInGBs` | `integer` | yes |
| `scanListenerPortNonTls` | `integer` | no |
| `scanListenerPortTls` | `integer` | no |
| `tags` | `Map<string>` | no |
| `timeZone` | `string` | no |
| `totalContainerDatabases` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `cloudAutonomousVmClusterId` | `string` | yes |

## CreateCloudExadataInfrastructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | yes |
| `shape` | `string` | yes |
| `availabilityZone` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `tags` | `Map<string>` | no |
| `computeCount` | `integer` | yes |
| `customerContactsToSendToOCI` | `List<CustomerContact>` | no |
| `maintenanceWindow` | `MaintenanceWindow` | no |
| `storageCount` | `integer` | yes |
| `clientToken` | `string` | no |
| `databaseServerType` | `string` | no |
| `storageServerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `cloudExadataInfrastructureId` | `string` | yes |

## CreateCloudVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `cpuCoreCount` | `integer` | yes |
| `displayName` | `string` | yes |
| `giVersion` | `string` | yes |
| `hostname` | `string` | yes |
| `sshPublicKeys` | `List<string>` | yes |
| `odbNetworkId` | `string` | yes |
| `clusterName` | `string` | no |
| `dataCollectionOptions` | `DataCollectionOptions` | no |
| `dataStorageSizeInTBs` | `double` | no |
| `dbNodeStorageSizeInGBs` | `integer` | no |
| `dbServers` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `isLocalBackupEnabled` | `boolean` | no |
| `isSparseDiskgroupEnabled` | `boolean` | no |
| `licenseModel` | `string` | no |
| `memorySizeInGBs` | `integer` | no |
| `systemVersion` | `string` | no |
| `timeZone` | `string` | no |
| `clientToken` | `string` | no |
| `scanListenerPortTcp` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `cloudVmClusterId` | `string` | yes |

## CreateExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | yes |
| `enabledEcpuCount` | `integer` | yes |
| `exascaleDbStorageVaultId` | `string` | yes |
| `gridImageId` | `string` | yes |
| `hostname` | `string` | yes |
| `nodeCount` | `integer` | yes |
| `odbNetworkId` | `string` | yes |
| `shape` | `string` | yes |
| `sshPublicKeys` | `List<string>` | yes |
| `totalEcpuCount` | `integer` | yes |
| `vmFileSystemStorageTotalSizeInGBs` | `integer` | yes |
| `clusterName` | `string` | no |
| `dataCollectionOptions` | `DataCollectionOptions` | no |
| `licenseModel` | `string` | no |
| `scanListenerPortTcp` | `integer` | no |
| `scanListenerPortTcpSsl` | `integer` | no |
| `shapeAttribute` | `string` | no |
| `systemVersion` | `string` | no |
| `tags` | `Map<string>` | no |
| `timeZone` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exadbVmClusterId` | `string` | yes |

## CreateExascaleDbStorageVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | yes |
| `highCapacityDatabaseStorageTotalSizeInGBs` | `integer` | yes |
| `additionalFlashCacheInPercent` | `integer` | no |
| `autoscaleLimitInGBs` | `integer` | no |
| `availabilityZoneId` | `string` | no |
| `availabilityZone` | `string` | no |
| `description` | `string` | no |
| `isAutoscaleEnabled` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `timeZone` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exascaleDbStorageVaultId` | `string` | yes |

## CreateOdbNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | yes |
| `availabilityZone` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `clientSubnetCidr` | `string` | yes |
| `backupSubnetCidr` | `string` | no |
| `customDomainName` | `string` | no |
| `defaultDnsPrefix` | `string` | no |
| `clientToken` | `string` | no |
| `s3Access` | `string` | no |
| `zeroEtlAccess` | `string` | no |
| `stsAccess` | `string` | no |
| `kmsAccess` | `string` | no |
| `s3PolicyDocument` | `string` | no |
| `stsPolicyDocument` | `string` | no |
| `kmsPolicyDocument` | `string` | no |
| `crossRegionS3RestoreSourcesToEnable` | `List<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `odbNetworkId` | `string` | yes |

## CreateOdbPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetworkId` | `string` | yes |
| `peerNetworkId` | `string` | yes |
| `displayName` | `string` | no |
| `peerNetworkCidrsToBeAdded` | `List<string>` | no |
| `peerNetworkRouteTableIds` | `List<string>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `odbPeeringConnectionId` | `string` | yes |

## DeleteAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAutonomousDatabaseBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseBackupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudAutonomousVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudAutonomousVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudExadataInfrastructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExascaleDbStorageVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exascaleDbStorageVaultId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOdbNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetworkId` | `string` | yes |
| `deleteAssociatedResources` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOdbPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateIamRoleFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `iamRoleArn` | `string` | yes |
| `awsIntegration` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateVirtualMachinesFromExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmClusterId` | `string` | yes |
| `dbNodeIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exadbVmClusterId` | `string` | yes |

## FailoverAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `peerDbArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## GetAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabase` | `AutonomousDatabase` | yes |

## GetAutonomousDatabaseBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseBackupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseBackup` | `AutonomousDatabaseBackup` | no |

## GetAutonomousDatabaseWalletDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseWalletDetails` | `AutonomousDatabaseWalletDetails` | yes |

## GetCloudAutonomousVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudAutonomousVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudAutonomousVmCluster` | `CloudAutonomousVmCluster` | no |

## GetCloudExadataInfrastructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructure` | `CloudExadataInfrastructure` | no |

## GetCloudExadataInfrastructureUnallocatedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `dbServers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureUnallocatedResources` | `CloudExadataInfrastructureUnallocatedResources` | no |

## GetCloudVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmCluster` | `CloudVmCluster` | no |

## GetDbNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | no |
| `exadbVmClusterId` | `string` | no |
| `dbNodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbNode` | `DbNode` | no |

## GetDbServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `dbServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbServer` | `DbServer` | no |

## GetExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmCluster` | `ExadbVmCluster` | yes |

## GetExascaleDbStorageVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exascaleDbStorageVaultId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exascaleDbStorageVault` | `ExascaleDbStorageVault` | yes |

## GetOciOnboardingStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `existingTenancyActivationLink` | `string` | no |
| `newTenancyActivationLink` | `string` | no |
| `ociIdentityDomain` | `OciIdentityDomain` | no |
| `autonomousDatabaseOciIntegrationIamRoles` | `List<OciIamRole>` | no |
| `linkedOciTenancyId` | `string` | no |
| `linkedOciCompartmentId` | `string` | no |
| `subscriptionErrors` | `List<SubscriptionError>` | no |

## GetOdbNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetwork` | `OdbNetwork` | no |

## GetOdbPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbPeeringConnection` | `OdbPeeringConnection` | no |

## InitializeService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ociIdentityDomain` | `boolean` | no |
| `autonomousDatabaseOciAwsSecretsManagerIntegration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListAutonomousDatabaseBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `autonomousDatabaseId` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabaseBackups` | `List<AutonomousDatabaseBackupSummary>` | yes |

## ListAutonomousDatabaseCharacterSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `characterSetType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabaseCharacterSets` | `List<AutonomousDatabaseCharacterSetSummary>` | yes |

## ListAutonomousDatabaseClones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabaseClones` | `List<AutonomousDatabaseSummary>` | yes |

## ListAutonomousDatabasePeers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabasePeers` | `List<AutonomousDatabasePeerSummary>` | yes |

## ListAutonomousDatabaseVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `dbWorkload` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabaseVersions` | `List<AutonomousDatabaseVersionSummary>` | yes |

## ListAutonomousDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousDatabases` | `List<AutonomousDatabaseSummary>` | yes |

## ListAutonomousVirtualMachines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `cloudAutonomousVmClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `autonomousVirtualMachines` | `List<AutonomousVirtualMachineSummary>` | yes |

## ListCloudAutonomousVmClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `cloudExadataInfrastructureId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `cloudAutonomousVmClusters` | `List<CloudAutonomousVmClusterSummary>` | yes |

## ListCloudExadataInfrastructures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `cloudExadataInfrastructures` | `List<CloudExadataInfrastructureSummary>` | yes |

## ListCloudVmClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `cloudExadataInfrastructureId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `cloudVmClusters` | `List<CloudVmClusterSummary>` | yes |

## ListDbNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `cloudVmClusterId` | `string` | no |
| `exadbVmClusterId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dbNodes` | `List<DbNodeSummary>` | yes |

## ListDbServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dbServers` | `List<DbServerSummary>` | yes |

## ListDbSystemShapes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `availabilityZone` | `string` | no |
| `availabilityZoneId` | `string` | no |
| `shapeFamily` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dbSystemShapes` | `List<DbSystemShapeSummary>` | yes |

## ListExadbVmClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exascaleDbStorageVaultId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `exadbVmClusters` | `List<ExadbVmClusterSummary>` | yes |

## ListExascaleDbStorageVaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `exascaleDbStorageVaults` | `List<ExascaleDbStorageVaultSummary>` | yes |

## ListFlexComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `shape` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `flexComponents` | `List<FlexComponentSummary>` | yes |

## ListGiMinorVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `giVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `shapeFamily` | `string` | no |
| `availabilityZone` | `string` | no |
| `availabilityZoneId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `giMinorVersions` | `List<GiMinorVersionSummary>` | yes |

## ListGiVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `shape` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `giVersions` | `List<GiVersionSummary>` | yes |

## ListOdbNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `odbNetworks` | `List<OdbNetworkSummary>` | yes |

## ListOdbPeeringConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `odbNetworkId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `odbPeeringConnections` | `List<OdbPeeringConnectionSummary>` | yes |

## ListSystemVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `giVersion` | `string` | yes |
| `shape` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `systemVersions` | `List<SystemVersionSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RebootAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `isOnlineReboot` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## RebootDbNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | no |
| `exadbVmClusterId` | `string` | no |
| `dbNodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbNodeId` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## RestoreAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `timestamp` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## ShrinkAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## StartAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## StartDbNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | no |
| `exadbVmClusterId` | `string` | no |
| `dbNodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbNodeId` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## StopAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## StopDbNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudVmClusterId` | `string` | no |
| `exadbVmClusterId` | `string` | no |
| `dbNodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dbNodeId` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## SwitchoverAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `peerDbArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAutonomousDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `adminPassword` | `string` | no |
| `computeCount` | `double` | no |
| `cpuCoreCount` | `integer` | no |
| `dataStorageSizeInTBs` | `integer` | no |
| `dataStorageSizeInGBs` | `integer` | no |
| `displayName` | `string` | no |
| `dbName` | `string` | no |
| `dbVersion` | `string` | no |
| `dbWorkload` | `string` | no |
| `dbToolsDetails` | `List<DatabaseTool>` | no |
| `databaseEdition` | `string` | no |
| `licenseModel` | `string` | no |
| `isAutoScalingEnabled` | `boolean` | no |
| `isAutoScalingForStorageEnabled` | `boolean` | no |
| `isBackupRetentionLocked` | `boolean` | no |
| `isLocalDataGuardEnabled` | `boolean` | no |
| `isMtlsConnectionRequired` | `boolean` | no |
| `isRefreshableClone` | `boolean` | no |
| `isDisconnectPeer` | `boolean` | no |
| `backupRetentionPeriodInDays` | `integer` | no |
| `byolComputeCountLimit` | `double` | no |
| `localAdgAutoFailoverMaxDataLossLimit` | `integer` | no |
| `autonomousMaintenanceScheduleType` | `string` | no |
| `customerContactsToSendToOCI` | `List<CustomerContact>` | no |
| `scheduledOperations` | `List<ScheduledOperationDetails>` | no |
| `longTermBackupSchedule` | `LongTermBackupSchedule` | no |
| `openMode` | `string` | no |
| `permissionLevel` | `string` | no |
| `refreshableMode` | `string` | no |
| `privateEndpointIp` | `string` | no |
| `privateEndpointLabel` | `string` | no |
| `peerDbId` | `string` | no |
| `resourcePoolLeaderId` | `string` | no |
| `resourcePoolSummary` | `ResourcePoolSummary` | no |
| `standbyAllowlistedIpsSource` | `string` | no |
| `standbyAllowlistedIps` | `List<string>` | no |
| `allowlistedIps` | `List<string>` | no |
| `autoRefreshFrequencyInSeconds` | `integer` | no |
| `autoRefreshPointLagInSeconds` | `integer` | no |
| `timeOfAutoRefreshStart` | `timestamp` | no |
| `encryptionKeyProvider` | `string` | no |
| `encryptionKeyConfiguration` | `EncryptionKeyConfigurationInput` | no |
| `adminPasswordSource` | `string` | no |
| `adminPasswordSourceConfiguration` | `AdminPasswordSourceConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseId` | `string` | yes |
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |

## UpdateAutonomousDatabaseBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autonomousDatabaseBackupId` | `string` | yes |
| `retentionPeriodInDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `autonomousDatabaseBackupId` | `string` | yes |

## UpdateCloudExadataInfrastructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudExadataInfrastructureId` | `string` | yes |
| `maintenanceWindow` | `MaintenanceWindow` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `cloudExadataInfrastructureId` | `string` | yes |

## UpdateExadbVmCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exadbVmClusterId` | `string` | yes |
| `dataCollectionOptions` | `DataCollectionOptions` | no |
| `displayName` | `string` | no |
| `enabledEcpuCount` | `integer` | no |
| `gridImageId` | `string` | no |
| `licenseModel` | `string` | no |
| `sshPublicKeys` | `List<string>` | no |
| `systemVersion` | `string` | no |
| `totalEcpuCount` | `integer` | no |
| `updateAction` | `string` | no |
| `vmFileSystemStorageTotalSizeInGBs` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exadbVmClusterId` | `string` | yes |

## UpdateExascaleDbStorageVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exascaleDbStorageVaultId` | `string` | yes |
| `additionalFlashCacheInPercent` | `integer` | no |
| `autoscaleLimitInGBs` | `integer` | no |
| `description` | `string` | no |
| `displayName` | `string` | no |
| `highCapacityDatabaseStorageTotalSizeInGBs` | `integer` | no |
| `isAutoscaleEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `exascaleDbStorageVaultId` | `string` | yes |

## UpdateOdbNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbNetworkId` | `string` | yes |
| `displayName` | `string` | no |
| `peeredCidrsToBeAdded` | `List<string>` | no |
| `peeredCidrsToBeRemoved` | `List<string>` | no |
| `s3Access` | `string` | no |
| `zeroEtlAccess` | `string` | no |
| `stsAccess` | `string` | no |
| `kmsAccess` | `string` | no |
| `s3PolicyDocument` | `string` | no |
| `stsPolicyDocument` | `string` | no |
| `kmsPolicyDocument` | `string` | no |
| `crossRegionS3RestoreSourcesToEnable` | `List<string>` | no |
| `crossRegionS3RestoreSourcesToDisable` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `odbNetworkId` | `string` | yes |

## UpdateOdbPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `odbPeeringConnectionId` | `string` | yes |
| `displayName` | `string` | no |
| `peerNetworkCidrsToBeAdded` | `List<string>` | no |
| `peerNetworkCidrsToBeRemoved` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `odbPeeringConnectionId` | `string` | yes |

