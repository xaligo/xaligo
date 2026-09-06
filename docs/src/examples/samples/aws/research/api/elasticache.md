# Amazon ElastiCache

API version: 2015-02-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/elasticache/2015-02-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## AuthorizeCacheSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroupName` | `string` | yes |
| `EC2SecurityGroupName` | `string` | yes |
| `EC2SecurityGroupOwnerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroup` | `CacheSecurityGroup` | no |

## BatchApplyUpdateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupIds` | `List<string>` | no |
| `CacheClusterIds` | `List<string>` | no |
| `ServiceUpdateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedUpdateActions` | `List<ProcessedUpdateAction>` | no |
| `UnprocessedUpdateActions` | `List<UnprocessedUpdateAction>` | no |

## BatchStopUpdateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupIds` | `List<string>` | no |
| `CacheClusterIds` | `List<string>` | no |
| `ServiceUpdateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedUpdateActions` | `List<ProcessedUpdateAction>` | no |
| `UnprocessedUpdateActions` | `List<UnprocessedUpdateAction>` | no |

## CompleteMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## CopyServerlessCacheSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceServerlessCacheSnapshotName` | `string` | yes |
| `TargetServerlessCacheSnapshotName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshot` | `ServerlessCacheSnapshot` | no |

## CopySnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceSnapshotName` | `string` | yes |
| `TargetSnapshotName` | `string` | yes |
| `TargetBucket` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateCacheCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | yes |
| `ReplicationGroupId` | `string` | no |
| `AZMode` | `string` | no |
| `PreferredAvailabilityZone` | `string` | no |
| `PreferredAvailabilityZones` | `List<string>` | no |
| `NumCacheNodes` | `integer` | no |
| `CacheNodeType` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `CacheParameterGroupName` | `string` | no |
| `CacheSubnetGroupName` | `string` | no |
| `CacheSecurityGroupNames` | `List<string>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `SnapshotArns` | `List<string>` | no |
| `SnapshotName` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `Port` | `integer` | no |
| `NotificationTopicArn` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `SnapshotWindow` | `string` | no |
| `AuthToken` | `string` | no |
| `OutpostMode` | `string` | no |
| `PreferredOutpostArn` | `string` | no |
| `PreferredOutpostArns` | `List<string>` | no |
| `LogDeliveryConfigurations` | `List<LogDeliveryConfigurationRequest>` | no |
| `TransitEncryptionEnabled` | `boolean` | no |
| `NetworkType` | `string` | no |
| `IpDiscovery` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheCluster` | `CacheCluster` | no |

## CreateCacheParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | yes |
| `CacheParameterGroupFamily` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroup` | `CacheParameterGroup` | no |

## CreateCacheSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroupName` | `string` | yes |
| `Description` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroup` | `CacheSecurityGroup` | no |

## CreateCacheSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroupName` | `string` | yes |
| `CacheSubnetGroupDescription` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroup` | `CacheSubnetGroup` | no |

## CreateGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupIdSuffix` | `string` | yes |
| `GlobalReplicationGroupDescription` | `string` | no |
| `PrimaryReplicationGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## CreateReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `ReplicationGroupDescription` | `string` | yes |
| `GlobalReplicationGroupId` | `string` | no |
| `PrimaryClusterId` | `string` | no |
| `AutomaticFailoverEnabled` | `boolean` | no |
| `MultiAZEnabled` | `boolean` | no |
| `NumCacheClusters` | `integer` | no |
| `PreferredCacheClusterAZs` | `List<string>` | no |
| `NumNodeGroups` | `integer` | no |
| `ReplicasPerNodeGroup` | `integer` | no |
| `NodeGroupConfiguration` | `List<NodeGroupConfiguration>` | no |
| `CacheNodeType` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `CacheParameterGroupName` | `string` | no |
| `CacheSubnetGroupName` | `string` | no |
| `CacheSecurityGroupNames` | `List<string>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `SnapshotArns` | `List<string>` | no |
| `SnapshotName` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `Port` | `integer` | no |
| `NotificationTopicArn` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `SnapshotWindow` | `string` | no |
| `AuthToken` | `string` | no |
| `TransitEncryptionEnabled` | `boolean` | no |
| `AtRestEncryptionEnabled` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `UserGroupIds` | `List<string>` | no |
| `LogDeliveryConfigurations` | `List<LogDeliveryConfigurationRequest>` | no |
| `DataTieringEnabled` | `boolean` | no |
| `NetworkType` | `string` | no |
| `IpDiscovery` | `string` | no |
| `TransitEncryptionMode` | `string` | no |
| `ClusterMode` | `string` | no |
| `ServerlessCacheSnapshotName` | `string` | no |
| `Durability` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## CreateServerlessCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheName` | `string` | yes |
| `Description` | `string` | no |
| `Engine` | `string` | yes |
| `MajorEngineVersion` | `string` | no |
| `CacheUsageLimits` | `CacheUsageLimits` | no |
| `KmsKeyId` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SnapshotArnsToRestore` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `UserGroupId` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `DailySnapshotTime` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCache` | `ServerlessCache` | no |

## CreateServerlessCacheSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshotName` | `string` | yes |
| `ServerlessCacheName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshot` | `ServerlessCacheSnapshot` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | no |
| `CacheClusterId` | `string` | no |
| `SnapshotName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `UserName` | `string` | yes |
| `Engine` | `string` | yes |
| `Passwords` | `List<string>` | no |
| `AccessString` | `string` | yes |
| `NoPasswordRequired` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `AuthenticationMode` | `AuthenticationMode` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `UserName` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `MinimumEngineVersion` | `string` | no |
| `AccessString` | `string` | no |
| `UserGroupIds` | `List<string>` | no |
| `Authentication` | `Authentication` | no |
| `ARN` | `string` | no |

## CreateUserGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | yes |
| `Engine` | `string` | yes |
| `UserIds` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `UserIds` | `List<string>` | no |
| `MinimumEngineVersion` | `string` | no |
| `PendingChanges` | `UserGroupPendingChanges` | no |
| `ReplicationGroups` | `List<string>` | no |
| `ServerlessCaches` | `List<string>` | no |
| `ARN` | `string` | no |

## DecreaseNodeGroupsInGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `NodeGroupCount` | `integer` | yes |
| `GlobalNodeGroupsToRemove` | `List<string>` | no |
| `GlobalNodeGroupsToRetain` | `List<string>` | no |
| `ApplyImmediately` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## DecreaseReplicaCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `NewReplicaCount` | `integer` | no |
| `ReplicaConfiguration` | `List<ConfigureShard>` | no |
| `ReplicasToRemove` | `List<string>` | no |
| `ApplyImmediately` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## DeleteCacheCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | yes |
| `FinalSnapshotIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheCluster` | `CacheCluster` | no |

## DeleteCacheParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCacheSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCacheSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `RetainPrimaryReplicationGroup` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## DeleteReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `RetainPrimaryCluster` | `boolean` | no |
| `FinalSnapshotIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## DeleteServerlessCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheName` | `string` | yes |
| `FinalSnapshotName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCache` | `ServerlessCache` | no |

## DeleteServerlessCacheSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshot` | `ServerlessCacheSnapshot` | no |

## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `UserName` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `MinimumEngineVersion` | `string` | no |
| `AccessString` | `string` | no |
| `UserGroupIds` | `List<string>` | no |
| `Authentication` | `Authentication` | no |
| `ARN` | `string` | no |

## DeleteUserGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `UserIds` | `List<string>` | no |
| `MinimumEngineVersion` | `string` | no |
| `PendingChanges` | `UserGroupPendingChanges` | no |
| `ReplicationGroups` | `List<string>` | no |
| `ServerlessCaches` | `List<string>` | no |
| `ARN` | `string` | no |

## DescribeCacheClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `ShowCacheNodeInfo` | `boolean` | no |
| `ShowCacheClustersNotInReplicationGroups` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `CacheClusters` | `List<CacheCluster>` | no |

## DescribeCacheEngineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `CacheParameterGroupFamily` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `DefaultOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `CacheEngineVersions` | `List<CacheEngineVersion>` | no |

## DescribeCacheParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `CacheParameterGroups` | `List<CacheParameterGroup>` | no |

## DescribeCacheParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | yes |
| `Source` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Parameters` | `List<Parameter>` | no |
| `CacheNodeTypeSpecificParameters` | `List<CacheNodeTypeSpecificParameter>` | no |

## DescribeCacheSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `CacheSecurityGroups` | `List<CacheSecurityGroup>` | no |

## DescribeCacheSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroupName` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `CacheSubnetGroups` | `List<CacheSubnetGroup>` | no |

## DescribeEngineDefaultParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupFamily` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineDefaults` | `EngineDefaults` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceIdentifier` | `string` | no |
| `SourceType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Duration` | `integer` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Events` | `List<Event>` | no |

## DescribeGlobalReplicationGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |
| `ShowMemberInfo` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `GlobalReplicationGroups` | `List<GlobalReplicationGroup>` | no |

## DescribeReplicationGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReplicationGroups` | `List<ReplicationGroup>` | no |

## DescribeReservedCacheNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCacheNodeId` | `string` | no |
| `ReservedCacheNodesOfferingId` | `string` | no |
| `CacheNodeType` | `string` | no |
| `Duration` | `string` | no |
| `ProductDescription` | `string` | no |
| `OfferingType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedCacheNodes` | `List<ReservedCacheNode>` | no |

## DescribeReservedCacheNodesOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCacheNodesOfferingId` | `string` | no |
| `CacheNodeType` | `string` | no |
| `Duration` | `string` | no |
| `ProductDescription` | `string` | no |
| `OfferingType` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ReservedCacheNodesOfferings` | `List<ReservedCacheNodesOffering>` | no |

## DescribeServerlessCacheSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheName` | `string` | no |
| `ServerlessCacheSnapshotName` | `string` | no |
| `SnapshotType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServerlessCacheSnapshots` | `List<ServerlessCacheSnapshot>` | no |

## DescribeServerlessCaches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServerlessCaches` | `List<ServerlessCache>` | no |

## DescribeServiceUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceUpdateName` | `string` | no |
| `ServiceUpdateStatus` | `List<string>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `ServiceUpdates` | `List<ServiceUpdate>` | no |

## DescribeSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | no |
| `CacheClusterId` | `string` | no |
| `SnapshotName` | `string` | no |
| `SnapshotSource` | `string` | no |
| `Marker` | `string` | no |
| `MaxRecords` | `integer` | no |
| `ShowNodeGroupConfig` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `Snapshots` | `List<Snapshot>` | no |

## DescribeUpdateActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceUpdateName` | `string` | no |
| `ReplicationGroupIds` | `List<string>` | no |
| `CacheClusterIds` | `List<string>` | no |
| `Engine` | `string` | no |
| `ServiceUpdateStatus` | `List<string>` | no |
| `ServiceUpdateTimeRange` | `TimeRangeFilter` | no |
| `UpdateActionStatus` | `List<string>` | no |
| `ShowNodeLevelUpdateStatus` | `boolean` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `UpdateActions` | `List<UpdateAction>` | no |

## DescribeUserGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroups` | `List<UserGroup>` | no |
| `Marker` | `string` | no |

## DescribeUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `UserId` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxRecords` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `Marker` | `string` | no |

## DisassociateGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `ReplicationGroupId` | `string` | yes |
| `ReplicationGroupRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## ExportServerlessCacheSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshotName` | `string` | yes |
| `S3BucketName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheSnapshot` | `ServerlessCacheSnapshot` | no |

## FailoverGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `PrimaryRegion` | `string` | yes |
| `PrimaryReplicationGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## IncreaseNodeGroupsInGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `NodeGroupCount` | `integer` | yes |
| `RegionalConfigurations` | `List<RegionalConfiguration>` | no |
| `ApplyImmediately` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## IncreaseReplicaCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `NewReplicaCount` | `integer` | no |
| `ReplicaConfiguration` | `List<ConfigureShard>` | no |
| `ApplyImmediately` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## ListAllowedNodeTypeModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | no |
| `ReplicationGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScaleUpModifications` | `List<string>` | no |
| `ScaleDownModifications` | `List<string>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ModifyCacheCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | yes |
| `NumCacheNodes` | `integer` | no |
| `CacheNodeIdsToRemove` | `List<string>` | no |
| `AZMode` | `string` | no |
| `NewAvailabilityZones` | `List<string>` | no |
| `CacheSecurityGroupNames` | `List<string>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `NotificationTopicArn` | `string` | no |
| `CacheParameterGroupName` | `string` | no |
| `NotificationTopicStatus` | `string` | no |
| `ApplyImmediately` | `boolean` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `SnapshotWindow` | `string` | no |
| `CacheNodeType` | `string` | no |
| `AuthToken` | `string` | no |
| `AuthTokenUpdateStrategy` | `string` | no |
| `LogDeliveryConfigurations` | `List<LogDeliveryConfigurationRequest>` | no |
| `IpDiscovery` | `string` | no |
| `ScaleConfig` | `ScaleConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheCluster` | `CacheCluster` | no |

## ModifyCacheParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | yes |
| `ParameterNameValues` | `List<ParameterNameValue>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | no |

## ModifyCacheSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroupName` | `string` | yes |
| `CacheSubnetGroupDescription` | `string` | no |
| `SubnetIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSubnetGroup` | `CacheSubnetGroup` | no |

## ModifyGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `ApplyImmediately` | `boolean` | yes |
| `CacheNodeType` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `CacheParameterGroupName` | `string` | no |
| `GlobalReplicationGroupDescription` | `string` | no |
| `AutomaticFailoverEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## ModifyReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `ReplicationGroupDescription` | `string` | no |
| `PrimaryClusterId` | `string` | no |
| `SnapshottingClusterId` | `string` | no |
| `AutomaticFailoverEnabled` | `boolean` | no |
| `MultiAZEnabled` | `boolean` | no |
| `NodeGroupId` | `string` | no |
| `CacheSecurityGroupNames` | `List<string>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `NotificationTopicArn` | `string` | no |
| `CacheParameterGroupName` | `string` | no |
| `NotificationTopicStatus` | `string` | no |
| `ApplyImmediately` | `boolean` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `SnapshotWindow` | `string` | no |
| `CacheNodeType` | `string` | no |
| `AuthToken` | `string` | no |
| `AuthTokenUpdateStrategy` | `string` | no |
| `UserGroupIdsToAdd` | `List<string>` | no |
| `UserGroupIdsToRemove` | `List<string>` | no |
| `RemoveUserGroups` | `boolean` | no |
| `LogDeliveryConfigurations` | `List<LogDeliveryConfigurationRequest>` | no |
| `IpDiscovery` | `string` | no |
| `TransitEncryptionEnabled` | `boolean` | no |
| `TransitEncryptionMode` | `string` | no |
| `ClusterMode` | `string` | no |
| `Durability` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## ModifyReplicationGroupShardConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `NodeGroupCount` | `integer` | yes |
| `ApplyImmediately` | `boolean` | yes |
| `ReshardingConfiguration` | `List<ReshardingConfiguration>` | no |
| `NodeGroupsToRemove` | `List<string>` | no |
| `NodeGroupsToRetain` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## ModifyServerlessCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCacheName` | `string` | yes |
| `Description` | `string` | no |
| `CacheUsageLimits` | `CacheUsageLimits` | no |
| `RemoveUserGroup` | `boolean` | no |
| `UserGroupId` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `DailySnapshotTime` | `string` | no |
| `Engine` | `string` | no |
| `MajorEngineVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServerlessCache` | `ServerlessCache` | no |

## ModifyUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `AccessString` | `string` | no |
| `AppendAccessString` | `string` | no |
| `Passwords` | `List<string>` | no |
| `NoPasswordRequired` | `boolean` | no |
| `AuthenticationMode` | `AuthenticationMode` | no |
| `Engine` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `UserName` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `MinimumEngineVersion` | `string` | no |
| `AccessString` | `string` | no |
| `UserGroupIds` | `List<string>` | no |
| `Authentication` | `Authentication` | no |
| `ARN` | `string` | no |

## ModifyUserGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | yes |
| `UserIdsToAdd` | `List<string>` | no |
| `UserIdsToRemove` | `List<string>` | no |
| `Engine` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserGroupId` | `string` | no |
| `Status` | `string` | no |
| `Engine` | `string` | no |
| `UserIds` | `List<string>` | no |
| `MinimumEngineVersion` | `string` | no |
| `PendingChanges` | `UserGroupPendingChanges` | no |
| `ReplicationGroups` | `List<string>` | no |
| `ServerlessCaches` | `List<string>` | no |
| `ARN` | `string` | no |

## PurchaseReservedCacheNodesOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCacheNodesOfferingId` | `string` | yes |
| `ReservedCacheNodeId` | `string` | no |
| `CacheNodeCount` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCacheNode` | `ReservedCacheNode` | no |

## RebalanceSlotsInGlobalReplicationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroupId` | `string` | yes |
| `ApplyImmediately` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalReplicationGroup` | `GlobalReplicationGroup` | no |

## RebootCacheCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheClusterId` | `string` | yes |
| `CacheNodeIdsToReboot` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheCluster` | `CacheCluster` | no |

## RemoveTagsFromResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ResetCacheParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | yes |
| `ResetAllParameters` | `boolean` | no |
| `ParameterNameValues` | `List<ParameterNameValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheParameterGroupName` | `string` | no |

## RevokeCacheSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroupName` | `string` | yes |
| `EC2SecurityGroupName` | `string` | yes |
| `EC2SecurityGroupOwnerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CacheSecurityGroup` | `CacheSecurityGroup` | no |

## StartMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `CustomerNodeEndpointList` | `List<CustomerNodeEndpoint>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## TestFailover

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `NodeGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

## TestMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroupId` | `string` | yes |
| `CustomerNodeEndpointList` | `List<CustomerNodeEndpoint>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicationGroup` | `ReplicationGroup` | no |

