# Amazon Neptune

API version: 2014-10-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/neptune/2014-10-31/service-2.json).

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

## CreateDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `BackupRetentionPeriod` | `integer` | no |
| `CharacterSetName` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
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
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DeletionProtection` | `boolean` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `GlobalClusterIdentifier` | `string` | no |
| `StorageType` | `string` | no |
| `NetworkType` | `string` | no |

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
| `OptionGroupName` | `string` | no |
| `CharacterSetName` | `string` | no |
| `PubliclyAccessible` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `DBClusterIdentifier` | `string` | yes |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `StorageEncrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `Domain` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `MonitoringRoleArn` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `PromotionTier` | `integer` | no |
| `Timezone` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DeletionProtection` | `boolean` | no |

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
| `DeletionProtection` | `boolean` | no |
| `DatabaseName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `StorageEncrypted` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalCluster` | `GlobalCluster` | no |

## DeleteDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |
| `SkipFinalSnapshot` | `boolean` | no |
| `FinalDBSnapshotIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

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

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstance` | `DBInstance` | no |

## DeleteDBParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `DBEngineVersions` | `List<DBEngineVersion>` | no |

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

## DescribeGlobalClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalClusterIdentifier` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `GlobalClusters` | `List<GlobalCluster>` | no |

## DescribeOrderableDBInstanceOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `EngineVersion` | `string` | no |
| `DBInstanceClass` | `string` | no |
| `LicenseModel` | `string` | no |
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

## DescribeValidDBInstanceModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBInstanceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ValidDBInstanceModificationsMessage` | `ValidDBInstanceModificationsMessage` | no |

## FailoverDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | no |
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
| `CloudwatchLogsExportConfiguration` | `CloudwatchLogsExportConfiguration` | no |
| `EngineVersion` | `string` | no |
| `AllowMajorVersionUpgrade` | `boolean` | no |
| `DBInstanceParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `StorageType` | `string` | no |
| `NetworkType` | `string` | no |

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
| `OptionGroupName` | `string` | no |
| `NewDBInstanceIdentifier` | `string` | no |
| `StorageType` | `string` | no |
| `TdeCredentialArn` | `string` | no |
| `TdeCredentialPassword` | `string` | no |
| `CACertificateIdentifier` | `string` | no |
| `Domain` | `string` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `MonitoringInterval` | `integer` | no |
| `DBPortNumber` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `MonitoringRoleArn` | `string` | no |
| `DomainIAMRoleName` | `string` | no |
| `PromotionTier` | `integer` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `EnablePerformanceInsights` | `boolean` | no |
| `PerformanceInsightsKMSKeyId` | `string` | no |
| `CloudwatchLogsExportConfiguration` | `CloudwatchLogsExportConfiguration` | no |
| `DeletionProtection` | `boolean` | no |

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

## PromoteReadReplicaDBCluster

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
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DBClusterParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `CopyTagsToSnapshot` | `boolean` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `StorageType` | `string` | no |
| `NetworkType` | `string` | no |

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
| `SourceDBClusterIdentifier` | `string` | yes |
| `RestoreToTime` | `timestamp` | no |
| `UseLatestRestorableTime` | `boolean` | no |
| `Port` | `integer` | no |
| `DBSubnetGroupName` | `string` | no |
| `OptionGroupName` | `string` | no |
| `VpcSecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `EnableIAMDatabaseAuthentication` | `boolean` | no |
| `EnableCloudwatchLogsExports` | `List<string>` | no |
| `DBClusterParameterGroupName` | `string` | no |
| `DeletionProtection` | `boolean` | no |
| `ServerlessV2ScalingConfiguration` | `ServerlessV2ScalingConfiguration` | no |
| `StorageType` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## StartDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

## StopDBCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBClusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DBCluster` | `DBCluster` | no |

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

