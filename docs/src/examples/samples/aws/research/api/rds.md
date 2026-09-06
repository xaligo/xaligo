# Amazon Relational Database Service

API version: 2014-10-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rds/2014-10-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddRoleToDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `FeatureName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddRoleToDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `FeatureName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddSourceIdentifierToSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApplyPendingMaintenanceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ApplyAction` | `string` | yes |
| `OptInType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePendingMaintenanceActions` | `ResourcePendingMaintenanceActions` | no |

## AuthorizeDBSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroupName` | `string` | yes |
| `CIDRIP` | `string` | no |
| `EC2SecurityGroupName` | `string` | no |
| `EC2SecurityGroupId` | `string` | no |
| `EC2SecurityGroupOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroup` | `DBSecurityGroup` | no |

## BacktrackDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `BacktrackTo` | `timestamp` | yes |
| `Force` | `boolean` | no |
| `UseEarliestTimeOnPointInTimeUnavailable` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
| `BacktrackIdentifier` | `string` | no |
| `BacktrackTo` | `timestamp` | no |
| `BacktrackedFrom` | `timestamp` | no |
| `BacktrackRequestCreationTime` | `timestamp` | no |
| `Status` | `string` | no |

## CancelExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskIdentifier` | `string` | no |
| `SourceArn` | `string` | no |
| `ExportOnly` | `List<string>` | no |
| `SnapshotTime` | `timestamp` | no |
| `TaskStartTime` | `timestamp` | no |
| `TaskEndTime` | `timestamp` | no |
| `S3Bucket` | `string` | no |
| `S3Prefix` | `string` | no |
| `IamRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Status` | `string` | no |
| `PercentProgress` | `integer` | no |
| `TotalExtractedDataInGB` | `integer` | no |
| `FailureCause` | `string` | no |
| `WarningMessage` | `string` | no |
| `SourceType` | `string` | no |

## CopyDBClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBClusterParameterGroupIdentifier` | `string` | yes |
| `TargetDBClusterParameterGroupIdentifier` | `string` | yes |
| `TargetDBClusterParameterGroupDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroup` | `DBClusterParameterGroup` | no |

## CopyDBClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBClusterSnapshotIdentifier` | `string` | yes |
| `TargetDBClusterSnapshotIdentifier` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `PreSignedUrl` | `string` | no |
| `CopyTags` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshot` | `DBClusterSnapshot` | no |

## CopyDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBParameterGroupIdentifier` | `string` | yes |
| `TargetDBParameterGroupIdentifier` | `string` | yes |
| `TargetDBParameterGroupDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroup` | `DBParameterGroup` | no |

## CopyDBSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBSnapshotIdentifier` | `string` | yes |
| `TargetDBSnapshotIdentifier` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CopyTags` | `boolean` | no |
| `PreSignedUrl` | `string` | no |
| `OptionGroupName` | `string` | no |
| `TargetCustomAvailabilityZone` | `string` | no |
| `SnapshotTarget` | `string` | no |
| `CopyOptionGroup` | `boolean` | no |
| `SnapshotAvailabilityZone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshot` | `DBSnapshot` | no |

## CopyOptionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceOptionGroupIdentifier` | `string` | yes |
| `TargetOptionGroupIdentifier` | `string` | yes |
| `TargetOptionGroupDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroup` | `OptionGroup` | no |

## CreateBlueGreenDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeploymentName` | `string` | yes |
| `Source` | `string` | yes |
| `TargetEngineVersion` | `string` | no |
| `TargetDBParameterGroupName` | `string` | no |
| `TargetDBClusterParameterGroupName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `TargetDBInstanceClass` | `string` | no |
| `UpgradeTargetStorageConfig` | `boolean` | no |
| `TargetIops` | `integer` | no |
| `TargetStorageType` | `string` | no |
| `TargetAllocatedStorage` | `integer` | no |
| `TargetStorageThroughput` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeployment` | `BlueGreenDeployment` | no |

## CreateCustomDBEngineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | yes |
| `DatabaseInstallationFilesS3BucketName` | `string` | no |
| `DatabaseInstallationFilesS3Prefix` | `string` | no |
| `DatabaseInstallationFiles` | `List<string>` | no |
| `ImageId` | `string` | no |
| `KMSKeyId` | `string` | no |
| `SourceCustomDbEngineVersionIdentifier` | `string` | no |
| `UseAwsProvidedLatestImage` | `boolean` | no |
| `Description` | `string` | no |
| `Manifest` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `MajorEngineVersion` | `string` | no |
| `EngineVersion` | `string` | no |
| `DatabaseInstallationFilesS3BucketName` | `string` | no |
| `DatabaseInstallationFilesS3Prefix` | `string` | no |
| `DatabaseInstallationFiles` | `List<string>` | no |
| `CustomDBEngineVersionManifest` | `string` | no |
| `DBParameterGroupFamily` | `string` | no |
| `DBEngineDescription` | `string` | no |
| `DBEngineVersionArn` | `string` | no |
| `DBEngineVersionDescription` | `string` | no |
| `DefaultCharacterSet` | `CharacterSet` | no |
| `FailureReason` | `string` | no |
| `Image` | `CustomDBEngineVersionAMI` | no |
| `DBEngineMediaType` | `string` | no |
| `KMSKeyId` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `SupportedCharacterSets` | `List<CharacterSet>` | no |
| `SupportedNcharCharacterSets` | `List<CharacterSet>` | no |
| `ValidUpgradeTarget` | `List<UpgradeTarget>` | no |
| `SupportedTimezones` | `List<Timezone>` | no |
| `ExportableLogTypes` | `List<string>` | no |
| `SupportsLogExportsToCloudwatchLogs` | `boolean` | no |
| `SupportsReadReplica` | `boolean` | no |
| `SupportedEngineModes` | `List<string>` | no |
| `SupportedFeatureNames` | `List<string>` | no |
| `Status` | `string` | no |
| `SupportsParallelQuery` | `boolean` | no |
| `SupportsGlobalDatabases` | `boolean` | no |
| `TagList` | `List<Tag>` | no |
| `SupportsBabelfish` | `boolean` | no |
| `SupportsLimitlessDatabase` | `boolean` | no |
| `SupportsCertificateRotationWithoutRestart` | `boolean` | no |
| `SupportedCACertificateIdentifiers` | `List<string>` | no |
| `SupportsLocalWriteForwarding` | `boolean` | no |
| `SupportsIntegrations` | `boolean` | no |
| `ServerlessV2FeaturesSupport` | `ServerlessV2FeaturesSupport` | no |

## CreateDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `CharacterSetName` | `string` | no |
| `DatabaseName` | `string` | no |
| `DBClusterIdentifier` | `string` | yes |
| `DBClusterParameterGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `DBSubnetGroupName` | `string` | no |
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | no |
| `Port` | `integer` | no |
| `MasterUsername` | `string` | no |
| `MasterUserPassword` | `string` | no |
| `OptionGroupName` | `string` | no |
| `PreferredBackupWindow` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `ReplicationSourceIdentifier` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `StorageEncrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `PreSignedUrl` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `BacktrackWindow` | `long` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `EngineMode` | `string` | no |
| `ScalingConfiguration` | `ScalingConfiguration` | no |
| `RdsCustomClusterConfiguration` | `RdsCustomClusterConfiguration` | no |
| `DBClusterInstanceClass` | `string` | no |
| `AllocatedStorage` | `integer` | no |
| `StorageType` | `string` | no |
| `Iops` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `DeletionProtection` | `boolean` | no |
| `GlobalClusterIdentifier` | `string` | no |
| `EnableHttpEndpoint` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `EnableGlobalWriteForwarding` | `boolean` | no |
| `NetworkType` | `string` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `EnableLimitlessDatabase` | `boolean` | no |
| `ClusterScalabilityType` | `string` | no |
| `DBSystemId` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `EnableLocalWriteForwarding` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `CACertificateIdentifier` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MasterUserAuthenticationType` | `string` | no |
| `WithExpressConfiguration` | `boolean` | no |
| `AssociatedRoles` | `List<DBClusterAssociatedRole>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## CreateDBClusterEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `DBClusterEndpointIdentifier` | `string` | yes |
| `EndpointType` | `string` | yes |
| `StaticMembers` | `List<string>` | no |
| `ExcludedMembers` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterEndpointIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `DBClusterEndpointResourceIdentifier` | `string` | no |
| `Endpoint` | `string` | no |
| `Status` | `string` | no |
| `EndpointType` | `string` | no |
| `CustomEndpointType` | `string` | no |
| `StaticMembers` | `List<string>` | no |
| `ExcludedMembers` | `List<string>` | no |
| `DBClusterEndpointArn` | `string` | no |

## CreateDBClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | yes |
| `DBParameterGroupFamily` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroup` | `DBClusterParameterGroup` | no |

## CreateDBClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotIdentifier` | `string` | yes |
| `DBClusterIdentifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshot` | `DBClusterSnapshot` | no |

## CreateDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBName` | `string` | no |
| `DBInstanceIdentifier` | `string` | yes |
| `AllocatedStorage` | `integer` | no |
| `DBInstanceClass` | `string` | yes |
| `Engine` | `string` | yes |
| `MasterUsername` | `string` | no |
| `MasterUserPassword` | `string` | no |
| `DBSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `AvailabilityZone` | `string` | no |
| `DBSubnetGroupName` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `DBParameterGroupName` | `string` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `Port` | `integer` | no |
| `MultiAZ` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `LicenseModel` | `string` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `CharacterSetName` | `string` | no |
| `NcharCharacterSetName` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `DBClusterIdentifier` | `string` | no |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `StorageEncrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `Domain` | `string` | no |
| `DomainFqdn` | `string` | no |
| `DomainOu` | `string` | no |
| `DomainAuthSecretArn` | `string` | no |
| `DomainDnsIps` | `List<string>` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `PromotionTier` | `integer` | no |
| `Timezone` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `DeletionProtection` | `boolean` | no |
| `MaxAllocatedStorage` | `integer` | no |
| `EnableCustomerOwnedIp` | `boolean` | no |
| `NetworkType` | `string` | no |
| `BackupTarget` | `string` | no |
| `CustomIamInstanceProfile` | `string` | no |
| `DBSystemId` | `string` | no |
| `CACertificateIdentifier` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `MultiTenant` | `boolean` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `EngineLifecycleSupport` | `string` | no |
| `AdditionalStorageVolumes` | `List<AdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MasterUserAuthenticationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## CreateDBInstanceReadReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `SourceDBInstanceIdentifier` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `Port` | `integer` | no |
| `MultiAZ` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `DBParameterGroupName` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `DBSubnetGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `StorageType` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `PreSignedUrl` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `UseDefaultProcessorFeatures` | `boolean` | no |
| `DeletionProtection` | `boolean` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `DomainFqdn` | `string` | no |
| `DomainOu` | `string` | no |
| `DomainAuthSecretArn` | `string` | no |
| `DomainDnsIps` | `List<string>` | no |
| `ReplicaMode` | `string` | no |
| `EnableCustomerOwnedIp` | `boolean` | no |
| `NetworkType` | `string` | no |
| `MaxAllocatedStorage` | `integer` | no |
| `BackupTarget` | `string` | no |
| `CustomIamInstanceProfile` | `string` | no |
| `AllocatedStorage` | `integer` | no |
| `SourceDBClusterIdentifier` | `string` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `UpgradeStorageConfig` | `boolean` | no |
| `CACertificateIdentifier` | `string` | no |
| `AdditionalStorageVolumes` | `List<AdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## CreateDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |
| `DBParameterGroupFamily` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroup` | `DBParameterGroup` | no |

## CreateDBProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `EngineFamily` | `string` | yes |
| `DefaultAuthScheme` | `string` | no |
| `Auth` | `List<UserAuthConfig>` | no |
| `RoleArn` | `string` | yes |
| `VpcSubnetIds` | `List<string>` | yes |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `RequireTLS` | `boolean` | no |
| `IdleClientTimeout` | `integer` | no |
| `DebugLogging` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `EndpointNetworkType` | `string` | no |
| `TargetConnectionNetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxy` | `DBProxy` | no |

## CreateDBProxyEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `DBProxyEndpointName` | `string` | yes |
| `VpcSubnetIds` | `List<string>` | yes |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `TargetRole` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `EndpointNetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpoint` | `DBProxyEndpoint` | no |

## CreateDBSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroupName` | `string` | yes |
| `DBSecurityGroupDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroup` | `DBSecurityGroup` | no |

## CreateDBShardGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupIdentifier` | `string` | yes |
| `DBClusterIdentifier` | `string` | yes |
| `ComputeRedundancy` | `integer` | no |
| `MaxACU` | `double` | yes |
| `MinACU` | `double` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupResourceId` | `string` | no |
| `DBShardGroupIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `MaxACU` | `double` | no |
| `MinACU` | `double` | no |
| `ComputeRedundancy` | `integer` | no |
| `Status` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Endpoint` | `string` | no |
| `DBShardGroupArn` | `string` | no |
| `TagList` | `List<Tag>` | no |

## CreateDBSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotIdentifier` | `string` | yes |
| `DBInstanceIdentifier` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshot` | `DBSnapshot` | no |

## CreateDBSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroupName` | `string` | yes |
| `DBSubnetGroupDescription` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroup` | `DBSubnetGroup` | no |

## CreateEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | yes |
| `SourceType` | `string` | no |
| `EventCategories` | `List<string>` | no |
| `SourceIds` | `List<string>` | no |
| `Enabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## CreateGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |
| `SourceDBClusterIdentifier` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `DatabaseName` | `string` | no |
| `StorageEncrypted` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## CreateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `TargetArn` | `string` | yes |
| `IntegrationName` | `string` | yes |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `DataFilter` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `IntegrationArn` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataFilter` | `string` | no |
| `Description` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `Errors` | `List<IntegrationError>` | no |

## CreateOptionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupName` | `string` | yes |
| `EngineName` | `string` | yes |
| `MajorEngineVersion` | `string` | yes |
| `OptionGroupDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroup` | `OptionGroup` | no |

## CreateTenantDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `TenantDBName` | `string` | yes |
| `MasterUsername` | `string` | yes |
| `MasterUserPassword` | `string` | no |
| `CharacterSetName` | `string` | no |
| `NcharCharacterSetName` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantDatabase` | `TenantDatabase` | no |

## DeleteBlueGreenDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeploymentIdentifier` | `string` | yes |
| `DeleteTarget` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeployment` | `BlueGreenDeployment` | no |

## DeleteCustomDBEngineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `MajorEngineVersion` | `string` | no |
| `EngineVersion` | `string` | no |
| `DatabaseInstallationFilesS3BucketName` | `string` | no |
| `DatabaseInstallationFilesS3Prefix` | `string` | no |
| `DatabaseInstallationFiles` | `List<string>` | no |
| `CustomDBEngineVersionManifest` | `string` | no |
| `DBParameterGroupFamily` | `string` | no |
| `DBEngineDescription` | `string` | no |
| `DBEngineVersionArn` | `string` | no |
| `DBEngineVersionDescription` | `string` | no |
| `DefaultCharacterSet` | `CharacterSet` | no |
| `FailureReason` | `string` | no |
| `Image` | `CustomDBEngineVersionAMI` | no |
| `DBEngineMediaType` | `string` | no |
| `KMSKeyId` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `SupportedCharacterSets` | `List<CharacterSet>` | no |
| `SupportedNcharCharacterSets` | `List<CharacterSet>` | no |
| `ValidUpgradeTarget` | `List<UpgradeTarget>` | no |
| `SupportedTimezones` | `List<Timezone>` | no |
| `ExportableLogTypes` | `List<string>` | no |
| `SupportsLogExportsToCloudwatchLogs` | `boolean` | no |
| `SupportsReadReplica` | `boolean` | no |
| `SupportedEngineModes` | `List<string>` | no |
| `SupportedFeatureNames` | `List<string>` | no |
| `Status` | `string` | no |
| `SupportsParallelQuery` | `boolean` | no |
| `SupportsGlobalDatabases` | `boolean` | no |
| `TagList` | `List<Tag>` | no |
| `SupportsBabelfish` | `boolean` | no |
| `SupportsLimitlessDatabase` | `boolean` | no |
| `SupportsCertificateRotationWithoutRestart` | `boolean` | no |
| `SupportedCACertificateIdentifiers` | `List<string>` | no |
| `SupportsLocalWriteForwarding` | `boolean` | no |
| `SupportsIntegrations` | `boolean` | no |
| `ServerlessV2FeaturesSupport` | `ServerlessV2FeaturesSupport` | no |

## DeleteDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `SkipFinalSnapshot` | `boolean` | no |
| `FinalDBSnapshotIdentifier` | `string` | no |
| `DeleteAutomatedBackups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## DeleteDBClusterAutomatedBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbClusterResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterAutomatedBackup` | `DBClusterAutomatedBackup` | no |

## DeleteDBClusterEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterEndpointIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterEndpointIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `DBClusterEndpointResourceIdentifier` | `string` | no |
| `Endpoint` | `string` | no |
| `Status` | `string` | no |
| `EndpointType` | `string` | no |
| `CustomEndpointType` | `string` | no |
| `StaticMembers` | `List<string>` | no |
| `ExcludedMembers` | `List<string>` | no |
| `DBClusterEndpointArn` | `string` | no |

## DeleteDBClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDBClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshot` | `DBClusterSnapshot` | no |

## DeleteDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `SkipFinalSnapshot` | `boolean` | no |
| `FinalDBSnapshotIdentifier` | `string` | no |
| `DeleteAutomatedBackups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## DeleteDBInstanceAutomatedBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbiResourceId` | `string` | no |
| `DBInstanceAutomatedBackupsArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceAutomatedBackup` | `DBInstanceAutomatedBackup` | no |

## DeleteDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDBProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxy` | `DBProxy` | no |

## DeleteDBProxyEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpoint` | `DBProxyEndpoint` | no |

## DeleteDBSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDBShardGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupResourceId` | `string` | no |
| `DBShardGroupIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `MaxACU` | `double` | no |
| `MinACU` | `double` | no |
| `ComputeRedundancy` | `integer` | no |
| `Status` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Endpoint` | `string` | no |
| `DBShardGroupArn` | `string` | no |
| `TagList` | `List<Tag>` | no |

## DeleteDBSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshot` | `DBSnapshot` | no |

## DeleteDBSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## DeleteGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `IntegrationArn` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataFilter` | `string` | no |
| `Description` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `Errors` | `List<IntegrationError>` | no |

## DeleteOptionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTenantDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `TenantDBName` | `string` | yes |
| `SkipFinalSnapshot` | `boolean` | no |
| `FinalDBSnapshotIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantDatabase` | `TenantDatabase` | no |

## DeregisterDBProxyTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `TargetGroupName` | `string` | no |
| `DBInstanceIdentifiers` | `List<string>` | no |
| `DBClusterIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountQuotas` | `List<AccountQuota>` | no |

## DescribeBlueGreenDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeploymentIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeployments` | `List<BlueGreenDeployment>` | no |
| `Marker` | `string` | no |

## DescribeCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultCertificateForNewLaunches` | `string` | no |
| `Certificates` | `List<Certificate>` | no |
| `Marker` | `string` | no |

## DescribeDBClusterAutomatedBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbClusterResourceId` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusterAutomatedBackups` | `List<DBClusterAutomatedBackup>` | no |

## DescribeDBClusterBacktracks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `BacktrackIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusterBacktracks` | `List<DBClusterBacktrack>` | no |

## DescribeDBClusterEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
| `DBClusterEndpointIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusterEndpoints` | `List<DBClusterEndpoint>` | no |

## DescribeDBClusterParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusterParameterGroups` | `List<DBClusterParameterGroup>` | no |

## DescribeDBClusterParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | yes |
| `Source` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<Parameter>` | no |
| `Marker` | `string` | no |

## DescribeDBClusterSnapshotAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotAttributesResult` | `DBClusterSnapshotAttributesResult` | no |

## DescribeDBClusterSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
| `DBClusterSnapshotIdentifier` | `string` | no |
| `SnapshotType` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `IncludeShared` | `boolean` | no |
| `IncludePublic` | `boolean` | no |
| `DbClusterResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusterSnapshots` | `List<DBClusterSnapshot>` | no |

## DescribeDBClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `IncludeShared` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBClusters` | `List<DBCluster>` | no |

## DescribeDBEngineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `DBParameterGroupFamily` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `DefaultOnly` | `boolean` | no |
| `ListSupportedCharacterSets` | `boolean` | no |
| `ListSupportedTimezones` | `boolean` | no |
| `IncludeAll` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBEngineVersions` | `List<DBEngineVersion>` | no |

## DescribeDBInstanceAutomatedBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DbiResourceId` | `string` | no |
| `DBInstanceIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `DBInstanceAutomatedBackupsArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBInstanceAutomatedBackups` | `List<DBInstanceAutomatedBackup>` | no |

## DescribeDBInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBInstances` | `List<DBInstance>` | no |

## DescribeDBLogFiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `FilenameContains` | `string` | no |
| `FileLastWritten` | `long` | no |
| `FileSize` | `long` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DescribeDBLogFiles` | `List<DescribeDBLogFilesDetails>` | no |
| `Marker` | `string` | no |

## DescribeDBMajorEngineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `MajorEngineVersion` | `string` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBMajorEngineVersions` | `List<DBMajorEngineVersion>` | no |
| `Marker` | `string` | no |

## DescribeDBParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBParameterGroups` | `List<DBParameterGroup>` | no |

## DescribeDBParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |
| `Source` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<Parameter>` | no |
| `Marker` | `string` | no |

## DescribeDBProxies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxies` | `List<DBProxy>` | no |
| `Marker` | `string` | no |

## DescribeDBProxyEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | no |
| `DBProxyEndpointName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpoints` | `List<DBProxyEndpoint>` | no |
| `Marker` | `string` | no |

## DescribeDBProxyTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `TargetGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroups` | `List<DBProxyTargetGroup>` | no |
| `Marker` | `string` | no |

## DescribeDBProxyTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `TargetGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<DBProxyTarget>` | no |
| `Marker` | `string` | no |

## DescribeDBRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastUpdatedAfter` | `timestamp` | no |
| `LastUpdatedBefore` | `timestamp` | no |
| `Locale` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBRecommendations` | `List<DBRecommendation>` | no |
| `Marker` | `string` | no |

## DescribeDBSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBSecurityGroups` | `List<DBSecurityGroup>` | no |

## DescribeDBShardGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroups` | `List<DBShardGroup>` | no |
| `Marker` | `string` | no |

## DescribeDBSnapshotAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotAttributesResult` | `DBSnapshotAttributesResult` | no |

## DescribeDBSnapshotTenantDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | no |
| `DBSnapshotIdentifier` | `string` | no |
| `SnapshotType` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `DbiResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBSnapshotTenantDatabases` | `List<DBSnapshotTenantDatabase>` | no |

## DescribeDBSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | no |
| `DBSnapshotIdentifier` | `string` | no |
| `SnapshotType` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `IncludeShared` | `boolean` | no |
| `IncludePublic` | `boolean` | no |
| `DbiResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBSnapshots` | `List<DBSnapshot>` | no |

## DescribeDBSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBSubnetGroups` | `List<DBSubnetGroup>` | no |

## DescribeEngineDefaultClusterParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupFamily` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineDefaults` | `EngineDefaults` | no |

## DescribeEngineDefaultParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupFamily` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineDefaults` | `EngineDefaults` | no |

## DescribeEventCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceType` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventCategoriesMapList` | `List<EventCategoriesMap>` | no |

## DescribeEventSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `EventSubscriptionsList` | `List<EventSubscription>` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceIdentifier` | `string` | no |
| `SourceType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Duration` | `integer` | no |
| `EventCategories` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Events` | `List<Event>` | no |

## DescribeExportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskIdentifier` | `string` | no |
| `SourceArn` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `SourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ExportTasks` | `List<ExportTask>` | no |

## DescribeGlobalClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `GlobalClusters` | `List<GlobalCluster>` | no |

## DescribeIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Integrations` | `List<Integration>` | no |

## DescribeOptionGroupOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineName` | `string` | yes |
| `MajorEngineVersion` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupOptions` | `List<OptionGroupOption>` | no |
| `Marker` | `string` | no |

## DescribeOptionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `EngineName` | `string` | no |
| `MajorEngineVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupsList` | `List<OptionGroup>` | no |
| `Marker` | `string` | no |

## DescribeOrderableDBInstanceOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `LicenseModel` | `string` | no |
| `AvailabilityZoneGroup` | `string` | no |
| `Vpc` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrderableDBInstanceOptions` | `List<OrderableDBInstanceOption>` | no |
| `Marker` | `string` | no |

## DescribePendingMaintenanceActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PendingMaintenanceActions` | `List<ResourcePendingMaintenanceActions>` | no |
| `Marker` | `string` | no |

## DescribeReservedDBInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedDBInstanceId` | `string` | no |
| `ReservedDBInstancesOfferingId` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `Duration` | `string` | no |
| `ProductDescription` | `string` | no |
| `OfferingType` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `LeaseId` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedDBInstances` | `List<ReservedDBInstance>` | no |

## DescribeReservedDBInstancesOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedDBInstancesOfferingId` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `Duration` | `string` | no |
| `ProductDescription` | `string` | no |
| `OfferingType` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedDBInstancesOfferings` | `List<ReservedDBInstancesOffering>` | no |

## DescribeServerlessV2PlatformVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessV2PlatformVersion` | `string` | no |
| `Engine` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DefaultOnly` | `boolean` | no |
| `IncludeAll` | `boolean` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ServerlessV2PlatformVersions` | `List<ServerlessV2PlatformVersionInfo>` | no |

## DescribeSourceRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `SourceRegions` | `List<SourceRegion>` | no |

## DescribeTenantDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | no |
| `TenantDBName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `TenantDatabases` | `List<TenantDatabase>` | no |

## DescribeValidDBInstanceModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ValidDBInstanceModificationsMessage` | `ValidDBInstanceModificationsMessage` | no |

## DisableHttpEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `HttpEndpointEnabled` | `boolean` | no |

## DownloadDBLogFilePortion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `LogFileName` | `string` | yes |
| `Marker` | `string` | no |
| `NumberOfLines` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogFileData` | `string` | no |
| `Marker` | `string` | no |
| `AdditionalDataPending` | `boolean` | no |

## EnableHttpEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `HttpEndpointEnabled` | `boolean` | no |

## FailoverDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `TargetDBInstanceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## FailoverGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |
| `TargetDbClusterIdentifier` | `string` | yes |
| `AllowDataLoss` | `boolean` | no |
| `Switchover` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ModifyActivityStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `AuditPolicyState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |
| `KinesisStreamName` | `string` | no |
| `Status` | `string` | no |
| `Mode` | `string` | no |
| `EngineNativeAuditFieldsIncluded` | `boolean` | no |
| `PolicyStatus` | `string` | no |

## ModifyCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateIdentifier` | `string` | no |
| `RemoveCustomerOverride` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `Certificate` | no |

## ModifyCurrentDBClusterCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `Capacity` | `integer` | no |
| `SecondsBeforeTimeout` | `integer` | no |
| `TimeoutAction` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
| `PendingCapacity` | `integer` | no |
| `CurrentCapacity` | `integer` | no |
| `SecondsBeforeTimeout` | `integer` | no |
| `TimeoutAction` | `string` | no |

## ModifyCustomDBEngineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | yes |
| `Description` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `MajorEngineVersion` | `string` | no |
| `EngineVersion` | `string` | no |
| `DatabaseInstallationFilesS3BucketName` | `string` | no |
| `DatabaseInstallationFilesS3Prefix` | `string` | no |
| `DatabaseInstallationFiles` | `List<string>` | no |
| `CustomDBEngineVersionManifest` | `string` | no |
| `DBParameterGroupFamily` | `string` | no |
| `DBEngineDescription` | `string` | no |
| `DBEngineVersionArn` | `string` | no |
| `DBEngineVersionDescription` | `string` | no |
| `DefaultCharacterSet` | `CharacterSet` | no |
| `FailureReason` | `string` | no |
| `Image` | `CustomDBEngineVersionAMI` | no |
| `DBEngineMediaType` | `string` | no |
| `KMSKeyId` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `SupportedCharacterSets` | `List<CharacterSet>` | no |
| `SupportedNcharCharacterSets` | `List<CharacterSet>` | no |
| `ValidUpgradeTarget` | `List<UpgradeTarget>` | no |
| `SupportedTimezones` | `List<Timezone>` | no |
| `ExportableLogTypes` | `List<string>` | no |
| `SupportsLogExportsToCloudwatchLogs` | `boolean` | no |
| `SupportsReadReplica` | `boolean` | no |
| `SupportedEngineModes` | `List<string>` | no |
| `SupportedFeatureNames` | `List<string>` | no |
| `Status` | `string` | no |
| `SupportsParallelQuery` | `boolean` | no |
| `SupportsGlobalDatabases` | `boolean` | no |
| `TagList` | `List<Tag>` | no |
| `SupportsBabelfish` | `boolean` | no |
| `SupportsLimitlessDatabase` | `boolean` | no |
| `SupportsCertificateRotationWithoutRestart` | `boolean` | no |
| `SupportedCACertificateIdentifiers` | `List<string>` | no |
| `SupportsLocalWriteForwarding` | `boolean` | no |
| `SupportsIntegrations` | `boolean` | no |
| `ServerlessV2FeaturesSupport` | `ServerlessV2FeaturesSupport` | no |

## ModifyDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `NewDBClusterIdentifier` | `string` | no |
| `ApplyImmediately` | `boolean` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `DBClusterParameterGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Port` | `integer` | no |
| `MasterUserPassword` | `string` | no |
| `OptionGroupName` | `string` | no |
| `PreferredBackupWindow` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `BacktrackWindow` | `long` | no |
| `CloudwatchLogsExportConfiguration` | `CloudwatchLogsExportConfiguration` | no |
| `EngineVersion` | `string` | no |
| `AllowMajorVersionUpgrade` | `boolean` | no |
| `DBInstanceParameterGroupName` | `string` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `ScalingConfiguration` | `ScalingConfiguration` | no |
| `DeletionProtection` | `boolean` | no |
| `EnableHttpEndpoint` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `EnableGlobalWriteForwarding` | `boolean` | no |
| `DBClusterInstanceClass` | `string` | no |
| `AllocatedStorage` | `integer` | no |
| `StorageType` | `string` | no |
| `Iops` | `integer` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `NetworkType` | `string` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `RotateMasterUserPassword` | `boolean` | no |
| `EnableLocalWriteForwarding` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `EngineMode` | `string` | no |
| `AllowEngineModeChange` | `boolean` | no |
| `AwsBackupRecoveryPointArn` | `string` | no |
| `EnableLimitlessDatabase` | `boolean` | no |
| `CACertificateIdentifier` | `string` | no |
| `MasterUserAuthenticationType` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## ModifyDBClusterEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterEndpointIdentifier` | `string` | yes |
| `EndpointType` | `string` | no |
| `StaticMembers` | `List<string>` | no |
| `ExcludedMembers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterEndpointIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `DBClusterEndpointResourceIdentifier` | `string` | no |
| `Endpoint` | `string` | no |
| `Status` | `string` | no |
| `EndpointType` | `string` | no |
| `CustomEndpointType` | `string` | no |
| `StaticMembers` | `List<string>` | no |
| `ExcludedMembers` | `List<string>` | no |
| `DBClusterEndpointArn` | `string` | no |

## ModifyDBClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | yes |
| `Parameters` | `List<Parameter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | no |

## ModifyDBClusterSnapshotAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotIdentifier` | `string` | yes |
| `AttributeName` | `string` | yes |
| `ValuesToAdd` | `List<string>` | no |
| `ValuesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterSnapshotAttributesResult` | `DBClusterSnapshotAttributesResult` | no |

## ModifyDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `AllocatedStorage` | `integer` | no |
| `DBInstanceClass` | `string` | no |
| `DBSubnetGroupName` | `string` | no |
| `DBSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `ApplyImmediately` | `boolean` | no |
| `MasterUserPassword` | `string` | no |
| `DBParameterGroupName` | `string` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AllowMajorVersionUpgrade` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `LicenseModel` | `string` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `NewDBInstanceIdentifier` | `string` | no |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `CACertificateIdentifier` | `string` | no |
| `Domain` | `string` | no |
| `DomainFqdn` | `string` | no |
| `DomainOu` | `string` | no |
| `DomainAuthSecretArn` | `string` | no |
| `DomainDnsIps` | `List<string>` | no |
| `DisableDomain` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `DBPortNumber` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `MonitoringRoleArn` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `PromotionTier` | `integer` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `CloudwatchLogsExportConfiguration` | `CloudwatchLogsExportConfiguration` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `UseDefaultProcessorFeatures` | `boolean` | no |
| `DeletionProtection` | `boolean` | no |
| `MaxAllocatedStorage` | `integer` | no |
| `CertificateRotationRestart` | `boolean` | no |
| `ReplicaMode` | `string` | no |
| `AutomationMode` | `string` | no |
| `ResumeFullAutomationModeMinutes` | `integer` | no |
| `EnableCustomerOwnedIp` | `boolean` | no |
| `NetworkType` | `string` | no |
| `AwsBackupRecoveryPointArn` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `RotateMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `MultiTenant` | `boolean` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `Engine` | `string` | no |
| `AdditionalStorageVolumes` | `List<ModifyAdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MasterUserAuthenticationType` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## ModifyDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |
| `Parameters` | `List<Parameter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | no |

## ModifyDBProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `NewDBProxyName` | `string` | no |
| `DefaultAuthScheme` | `string` | no |
| `Auth` | `List<UserAuthConfig>` | no |
| `RequireTLS` | `boolean` | no |
| `IdleClientTimeout` | `integer` | no |
| `DebugLogging` | `boolean` | no |
| `RoleArn` | `string` | no |
| `SecurityGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxy` | `DBProxy` | no |

## ModifyDBProxyEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpointName` | `string` | yes |
| `NewDBProxyEndpointName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyEndpoint` | `DBProxyEndpoint` | no |

## ModifyDBProxyTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupName` | `string` | yes |
| `DBProxyName` | `string` | yes |
| `ConnectionPoolConfig` | `ConnectionPoolConfiguration` | no |
| `NewName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyTargetGroup` | `DBProxyTargetGroup` | no |

## ModifyDBRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationId` | `string` | yes |
| `Locale` | `string` | no |
| `Status` | `string` | no |
| `RecommendedActionUpdates` | `List<RecommendedActionUpdate>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBRecommendation` | `DBRecommendation` | no |

## ModifyDBShardGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupIdentifier` | `string` | yes |
| `MaxACU` | `double` | no |
| `MinACU` | `double` | no |
| `ComputeRedundancy` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupResourceId` | `string` | no |
| `DBShardGroupIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `MaxACU` | `double` | no |
| `MinACU` | `double` | no |
| `ComputeRedundancy` | `integer` | no |
| `Status` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Endpoint` | `string` | no |
| `DBShardGroupArn` | `string` | no |
| `TagList` | `List<Tag>` | no |

## ModifyDBSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotIdentifier` | `string` | yes |
| `EngineVersion` | `string` | no |
| `OptionGroupName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshot` | `DBSnapshot` | no |

## ModifyDBSnapshotAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotIdentifier` | `string` | yes |
| `AttributeName` | `string` | yes |
| `ValuesToAdd` | `List<string>` | no |
| `ValuesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSnapshotAttributesResult` | `DBSnapshotAttributesResult` | no |

## ModifyDBSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroupName` | `string` | yes |
| `DBSubnetGroupDescription` | `string` | no |
| `SubnetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSubnetGroup` | `DBSubnetGroup` | no |

## ModifyEventSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SnsTopicArn` | `string` | no |
| `SourceType` | `string` | no |
| `EventCategories` | `List<string>` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## ModifyGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |
| `NewGlobalClusterIdentifier` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AllowMajorVersionUpgrade` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## ModifyIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationIdentifier` | `string` | yes |
| `IntegrationName` | `string` | no |
| `DataFilter` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `TargetArn` | `string` | no |
| `IntegrationName` | `string` | no |
| `IntegrationArn` | `string` | no |
| `KMSKeyId` | `string` | no |
| `AdditionalEncryptionContext` | `Map<string>` | no |
| `Status` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataFilter` | `string` | no |
| `Description` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `Errors` | `List<IntegrationError>` | no |

## ModifyOptionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroupName` | `string` | yes |
| `OptionsToInclude` | `List<OptionConfiguration>` | no |
| `OptionsToRemove` | `List<string>` | no |
| `ApplyImmediately` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptionGroup` | `OptionGroup` | no |

## ModifyTenantDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `TenantDBName` | `string` | yes |
| `MasterUserPassword` | `string` | no |
| `NewTenantDBName` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `RotateMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantDatabase` | `TenantDatabase` | no |

## PromoteReadReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## PromoteReadReplicaDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## PurchaseReservedDBInstancesOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedDBInstancesOfferingId` | `string` | yes |
| `ReservedDBInstanceId` | `string` | no |
| `DBInstanceCount` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedDBInstance` | `ReservedDBInstance` | no |

## RebootDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## RebootDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `ForceFailover` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## RebootDBShardGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBShardGroupResourceId` | `string` | no |
| `DBShardGroupIdentifier` | `string` | no |
| `DBClusterIdentifier` | `string` | no |
| `MaxACU` | `double` | no |
| `MinACU` | `double` | no |
| `ComputeRedundancy` | `integer` | no |
| `Status` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Endpoint` | `string` | no |
| `DBShardGroupArn` | `string` | no |
| `TagList` | `List<Tag>` | no |

## RegisterDBProxyTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyName` | `string` | yes |
| `TargetGroupName` | `string` | no |
| `DBInstanceIdentifiers` | `List<string>` | no |
| `DBClusterIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBProxyTargets` | `List<DBProxyTarget>` | no |

## RemoveFromGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |
| `DbClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## RemoveRoleFromDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `FeatureName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveRoleFromDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `FeatureName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveSourceIdentifierFromSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionName` | `string` | yes |
| `SourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSubscription` | `EventSubscription` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetDBClusterParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | yes |
| `ResetAllParameters` | `boolean` | no |
| `Parameters` | `List<Parameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterParameterGroupName` | `string` | no |

## ResetDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |
| `ResetAllParameters` | `boolean` | no |
| `Parameters` | `List<Parameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | no |

## RestoreDBClusterFromS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `CharacterSetName` | `string` | no |
| `DatabaseName` | `string` | no |
| `DBClusterIdentifier` | `string` | yes |
| `DBClusterParameterGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `DBSubnetGroupName` | `string` | no |
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | no |
| `Port` | `integer` | no |
| `MasterUsername` | `string` | yes |
| `MasterUserPassword` | `string` | no |
| `OptionGroupName` | `string` | no |
| `PreferredBackupWindow` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `StorageEncrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `SourceEngine` | `string` | yes |
| `SourceEngineVersion` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `S3Prefix` | `string` | no |
| `S3IngestionRoleArn` | `string` | yes |
| `BacktrackWindow` | `long` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DeletionProtection` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `StorageType` | `string` | no |
| `NetworkType` | `string` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `AssociatedRoles` | `List<DBClusterAssociatedRole>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## RestoreDBClusterFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `DBClusterIdentifier` | `string` | yes |
| `SnapshotIdentifier` | `string` | yes |
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | no |
| `Port` | `integer` | no |
| `DBSubnetGroupName` | `string` | no |
| `DatabaseName` | `string` | no |
| `OptionGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `BacktrackWindow` | `long` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `EngineMode` | `string` | no |
| `ScalingConfiguration` | `ScalingConfiguration` | no |
| `DBClusterParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `DBClusterInstanceClass` | `string` | no |
| `StorageType` | `string` | no |
| `Iops` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `NetworkType` | `string` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `RdsCustomClusterConfiguration` | `RdsCustomClusterConfiguration` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `EnableVPCNetworking` | `boolean` | no |
| `EnableInternetAccessGateway` | `boolean` | no |
| `AssociatedRoles` | `List<DBClusterAssociatedRole>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## RestoreDBClusterToPointInTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `RestoreType` | `string` | no |
| `SourceDBClusterIdentifier` | `string` | no |
| `RestoreToTime` | `timestamp` | no |
| `UseLatestRestorableTime` | `boolean` | no |
| `Port` | `integer` | no |
| `DBSubnetGroupName` | `string` | no |
| `OptionGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `BacktrackWindow` | `long` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DBClusterParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `DBClusterInstanceClass` | `string` | no |
| `StorageType` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Iops` | `integer` | no |
| `NetworkType` | `string` | no |
| `SourceDbClusterResourceId` | `string` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `ScalingConfiguration` | `ScalingConfiguration` | no |
| `EngineMode` | `string` | no |
| `RdsCustomClusterConfiguration` | `RdsCustomClusterConfiguration` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `EnableVPCNetworking` | `boolean` | no |
| `EnableInternetAccessGateway` | `boolean` | no |
| `AssociatedRoles` | `List<DBClusterAssociatedRole>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## RestoreDBInstanceFromDBSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `DBSnapshotIdentifier` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `Port` | `integer` | no |
| `AvailabilityZone` | `string` | no |
| `DBSubnetGroupName` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `PubliclyAccessible` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `LicenseModel` | `string` | no |
| `DBName` | `string` | no |
| `Engine` | `string` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Domain` | `string` | no |
| `DomainFqdn` | `string` | no |
| `DomainOu` | `string` | no |
| `DomainAuthSecretArn` | `string` | no |
| `DomainDnsIps` | `List<string>` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `DomainIAMRoleName` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `UseDefaultProcessorFeatures` | `boolean` | no |
| `DBParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `EnableCustomerOwnedIp` | `boolean` | no |
| `NetworkType` | `string` | no |
| `BackupTarget` | `string` | no |
| `CustomIamInstanceProfile` | `string` | no |
| `AllocatedStorage` | `integer` | no |
| `DBClusterSnapshotIdentifier` | `string` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `CACertificateIdentifier` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `AdditionalStorageVolumes` | `List<AdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## RestoreDBInstanceFromS3

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBName` | `string` | no |
| `DBInstanceIdentifier` | `string` | yes |
| `AllocatedStorage` | `integer` | no |
| `DBInstanceClass` | `string` | yes |
| `Engine` | `string` | yes |
| `MasterUsername` | `string` | no |
| `MasterUserPassword` | `string` | no |
| `DBSecurityGroups` | `List<string>` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `AvailabilityZone` | `string` | no |
| `DBSubnetGroupName` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `DBParameterGroupName` | `string` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `Port` | `integer` | no |
| `MultiAZ` | `boolean` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `LicenseModel` | `string` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `StorageType` | `string` | no |
| `StorageEncrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `SourceEngine` | `string` | yes |
| `SourceEngineVersion` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `S3Prefix` | `string` | no |
| `S3IngestionRoleArn` | `string` | yes |
| `DatabaseInsightsMode` | `string` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `PerformanceInsightsRetentionPeriod` | `integer` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `UseDefaultProcessorFeatures` | `boolean` | no |
| `DeletionProtection` | `boolean` | no |
| `MaxAllocatedStorage` | `integer` | no |
| `NetworkType` | `string` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `CACertificateIdentifier` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `AdditionalStorageVolumes` | `List<AdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## RestoreDBInstanceToPointInTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBInstanceIdentifier` | `string` | no |
| `TargetDBInstanceIdentifier` | `string` | yes |
| `RestoreTime` | `timestamp` | no |
| `UseLatestRestorableTime` | `boolean` | no |
| `DBInstanceClass` | `string` | no |
| `Port` | `integer` | no |
| `AvailabilityZone` | `string` | no |
| `DBSubnetGroupName` | `string` | no |
| `MultiAZ` | `boolean` | no |
| `PubliclyAccessible` | `boolean` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `LicenseModel` | `string` | no |
| `DBName` | `string` | no |
| `Engine` | `string` | no |
| `Iops` | `integer` | no |
| `StorageThroughput` | `integer` | no |
| `OptionGroupName` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Domain` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `DomainFqdn` | `string` | no |
| `DomainOu` | `string` | no |
| `DomainAuthSecretArn` | `string` | no |
| `DomainDnsIps` | `List<string>` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `ProcessorFeatures` | `List<ProcessorFeature>` | no |
| `UseDefaultProcessorFeatures` | `boolean` | no |
| `DBParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `SourceDbiResourceId` | `string` | no |
| `MaxAllocatedStorage` | `integer` | no |
| `EnableCustomerOwnedIp` | `boolean` | no |
| `NetworkType` | `string` | no |
| `SourceDBInstanceAutomatedBackupsArn` | `string` | no |
| `BackupTarget` | `string` | no |
| `CustomIamInstanceProfile` | `string` | no |
| `AllocatedStorage` | `integer` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `PreferredBackupWindow` | `string` | no |
| `DedicatedLogVolume` | `boolean` | no |
| `CACertificateIdentifier` | `string` | no |
| `EngineLifecycleSupport` | `string` | no |
| `AdditionalStorageVolumes` | `List<AdditionalStorageVolume>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ManageMasterUserPassword` | `boolean` | no |
| `MasterUserSecretKmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## RevokeDBSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroupName` | `string` | yes |
| `CIDRIP` | `string` | no |
| `EC2SecurityGroupName` | `string` | no |
| `EC2SecurityGroupId` | `string` | no |
| `EC2SecurityGroupOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBSecurityGroup` | `DBSecurityGroup` | no |

## StartActivityStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Mode` | `string` | yes |
| `KmsKeyId` | `string` | yes |
| `ApplyImmediately` | `boolean` | no |
| `EngineNativeAuditFieldsIncluded` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |
| `KinesisStreamName` | `string` | no |
| `Status` | `string` | no |
| `Mode` | `string` | no |
| `EngineNativeAuditFieldsIncluded` | `boolean` | no |
| `ApplyImmediately` | `boolean` | no |

## StartDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## StartDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## StartDBInstanceAutomatedBackupsReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBInstanceArn` | `string` | yes |
| `BackupRetentionPeriod` | `integer` | no |
| `KmsKeyId` | `string` | no |
| `PreSignedUrl` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceAutomatedBackup` | `DBInstanceAutomatedBackup` | no |

## StartExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskIdentifier` | `string` | yes |
| `SourceArn` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `IamRoleArn` | `string` | yes |
| `KmsKeyId` | `string` | yes |
| `S3Prefix` | `string` | no |
| `ExportOnly` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskIdentifier` | `string` | no |
| `SourceArn` | `string` | no |
| `ExportOnly` | `List<string>` | no |
| `SnapshotTime` | `timestamp` | no |
| `TaskStartTime` | `timestamp` | no |
| `TaskEndTime` | `timestamp` | no |
| `S3Bucket` | `string` | no |
| `S3Prefix` | `string` | no |
| `IamRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Status` | `string` | no |
| `PercentProgress` | `integer` | no |
| `TotalExtractedDataInGB` | `integer` | no |
| `FailureCause` | `string` | no |
| `WarningMessage` | `string` | no |
| `SourceType` | `string` | no |

## StopActivityStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ApplyImmediately` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |
| `KinesisStreamName` | `string` | no |
| `Status` | `string` | no |

## StopDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## StopDBInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |
| `DBSnapshotIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## StopDBInstanceAutomatedBackupsReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDBInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceAutomatedBackup` | `DBInstanceAutomatedBackup` | no |

## SwitchoverBlueGreenDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeploymentIdentifier` | `string` | yes |
| `SwitchoverTimeout` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueGreenDeployment` | `BlueGreenDeployment` | no |

## SwitchoverGlobalCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | yes |
| `TargetDbClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## SwitchoverReadReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

