# Amazon MemoryDB

API version: 2021-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/memorydb/2021-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchUpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterNames` | `List<string>` | yes |
| `ServiceUpdate` | `ServiceUpdateRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedClusters` | `List<Cluster>` | no |
| `UnprocessedClusters` | `List<UnprocessedCluster>` | no |

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

## CreateACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACLName` | `string` | yes |
| `UserNames` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACL` | `ACL` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeType` | `string` | yes |
| `MultiRegionClusterName` | `string` | no |
| `ParameterGroupName` | `string` | no |
| `Description` | `string` | no |
| `NumShards` | `integer` | no |
| `NumReplicasPerShard` | `integer` | no |
| `SubnetGroupName` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `MaintenanceWindow` | `string` | no |
| `Port` | `integer` | no |
| `SnsTopicArn` | `string` | no |
| `TLSEnabled` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `SnapshotArns` | `List<string>` | no |
| `SnapshotName` | `string` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `Tags` | `List<Tag>` | no |
| `SnapshotWindow` | `string` | no |
| `ACLName` | `string` | yes |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `DataTiering` | `boolean` | no |
| `NetworkType` | `string` | no |
| `IpDiscovery` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## CreateMultiRegionCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionClusterNameSuffix` | `string` | yes |
| `Description` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `NodeType` | `string` | yes |
| `MultiRegionParameterGroupName` | `string` | no |
| `NumShards` | `integer` | no |
| `TLSEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionCluster` | `MultiRegionCluster` | no |

## CreateParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `Family` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroup` | `ParameterGroup` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `SnapshotName` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## CreateSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | yes |
| `Description` | `string` | no |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroup` | `SubnetGroup` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AuthenticationMode` | `AuthenticationMode` | yes |
| `AccessString` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## DeleteACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACLName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACL` | `ACL` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `MultiRegionClusterName` | `string` | no |
| `FinalSnapshotName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DeleteMultiRegionCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionCluster` | `MultiRegionCluster` | no |

## DeleteParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroup` | `ParameterGroup` | no |

## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshot` | `Snapshot` | no |

## DeleteSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroup` | `SubnetGroup` | no |

## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## DescribeACLs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACLName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACLs` | `List<ACL>` | no |
| `NextToken` | `string` | no |

## DescribeClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ShowShardDetails` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Clusters` | `List<Cluster>` | no |

## DescribeEngineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `ParameterGroupFamily` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DefaultOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `EngineVersions` | `List<EngineVersionInfo>` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceName` | `string` | no |
| `SourceType` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Duration` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Events` | `List<Event>` | no |

## DescribeMultiRegionClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionClusterName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ShowClusterDetails` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MultiRegionClusters` | `List<MultiRegionCluster>` | no |

## DescribeMultiRegionParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionParameterGroupName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MultiRegionParameterGroups` | `List<MultiRegionParameterGroup>` | no |

## DescribeMultiRegionParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionParameterGroupName` | `string` | yes |
| `Source` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MultiRegionParameters` | `List<MultiRegionParameter>` | no |

## DescribeParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ParameterGroups` | `List<ParameterGroup>` | no |

## DescribeParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Parameters` | `List<Parameter>` | no |

## DescribeReservedNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservationId` | `string` | no |
| `ReservedNodesOfferingId` | `string` | no |
| `NodeType` | `string` | no |
| `Duration` | `string` | no |
| `OfferingType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedNodes` | `List<ReservedNode>` | no |

## DescribeReservedNodesOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodesOfferingId` | `string` | no |
| `NodeType` | `string` | no |
| `Duration` | `string` | no |
| `OfferingType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedNodesOfferings` | `List<ReservedNodesOffering>` | no |

## DescribeServiceUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceUpdateName` | `string` | no |
| `ClusterNames` | `List<string>` | no |
| `Status` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServiceUpdates` | `List<ServiceUpdate>` | no |

## DescribeSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | no |
| `SnapshotName` | `string` | no |
| `Source` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ShowDetail` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Snapshots` | `List<Snapshot>` | no |

## DescribeSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SubnetGroups` | `List<SubnetGroup>` | no |

## DescribeUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `NextToken` | `string` | no |

## FailoverShard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `ShardName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ListAllowedMultiRegionClusterUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScaleUpNodeTypes` | `List<string>` | no |
| `ScaleDownNodeTypes` | `List<string>` | no |

## ListAllowedNodeTypeUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScaleUpNodeTypes` | `List<string>` | no |
| `ScaleDownNodeTypes` | `List<string>` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## PurchaseReservedNodesOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNodesOfferingId` | `string` | yes |
| `ReservationId` | `string` | no |
| `NodeCount` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedNode` | `ReservedNode` | no |

## ResetParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `AllParameters` | `boolean` | no |
| `ParameterNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroup` | `ParameterGroup` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## UpdateACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACLName` | `string` | yes |
| `UserNamesToAdd` | `List<string>` | no |
| `UserNamesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ACL` | `ACL` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `Description` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `MaintenanceWindow` | `string` | no |
| `SnsTopicArn` | `string` | no |
| `SnsTopicStatus` | `string` | no |
| `ParameterGroupName` | `string` | no |
| `SnapshotWindow` | `string` | no |
| `SnapshotRetentionLimit` | `integer` | no |
| `NodeType` | `string` | no |
| `Engine` | `string` | no |
| `EngineVersion` | `string` | no |
| `ReplicaConfiguration` | `ReplicaConfigurationRequest` | no |
| `ShardConfiguration` | `ShardConfigurationRequest` | no |
| `ACLName` | `string` | no |
| `IpDiscovery` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## UpdateMultiRegionCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionClusterName` | `string` | yes |
| `NodeType` | `string` | no |
| `Description` | `string` | no |
| `EngineVersion` | `string` | no |
| `ShardConfiguration` | `ShardConfigurationRequest` | no |
| `MultiRegionParameterGroupName` | `string` | no |
| `UpdateStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionCluster` | `MultiRegionCluster` | no |

## UpdateParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `ParameterNameValues` | `List<ParameterNameValue>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroup` | `ParameterGroup` | no |

## UpdateSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | yes |
| `Description` | `string` | no |
| `SubnetIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroup` | `SubnetGroup` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AuthenticationMode` | `AuthenticationMode` | no |
| `AccessString` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

