# Amazon DocumentDB Elastic Clusters

API version: 2022-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/docdb-elastic/2022-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ApplyPendingMaintenanceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applyAction` | `string` | yes |
| `applyOn` | `string` | no |
| `optInType` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePendingMaintenanceAction` | `ResourcePendingMaintenanceAction` | yes |

## CopyClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `copyTags` | `boolean` | no |
| `kmsKeyId` | `string` | no |
| `snapshotArn` | `string` | yes |
| `tags` | `Map<string>` | no |
| `targetSnapshotName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `ClusterSnapshot` | yes |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminUserName` | `string` | yes |
| `adminUserPassword` | `string` | yes |
| `authType` | `string` | yes |
| `backupRetentionPeriod` | `integer` | no |
| `clientToken` | `string` | no |
| `clusterName` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `preferredBackupWindow` | `string` | no |
| `preferredMaintenanceWindow` | `string` | no |
| `shardCapacity` | `integer` | yes |
| `shardCount` | `integer` | yes |
| `shardInstanceCount` | `integer` | no |
| `subnetIds` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

## CreateClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | yes |
| `snapshotName` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `ClusterSnapshot` | yes |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

## DeleteClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `ClusterSnapshot` | yes |

## GetCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

## GetClusterSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshot` | `ClusterSnapshot` | yes |

## GetPendingMaintenanceAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePendingMaintenanceAction` | `ResourcePendingMaintenanceAction` | yes |

## ListClusterSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `snapshotType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `snapshots` | `List<ClusterSnapshotInList>` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusters` | `List<ClusterInList>` | no |
| `nextToken` | `string` | no |

## ListPendingMaintenanceActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resourcePendingMaintenanceActions` | `List<ResourcePendingMaintenanceAction>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RestoreClusterFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `shardCapacity` | `integer` | no |
| `shardInstanceCount` | `integer` | no |
| `snapshotArn` | `string` | yes |
| `subnetIds` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

## StartCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

## StopCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

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


## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adminUserPassword` | `string` | no |
| `authType` | `string` | no |
| `backupRetentionPeriod` | `integer` | no |
| `clientToken` | `string` | no |
| `clusterArn` | `string` | yes |
| `preferredBackupWindow` | `string` | no |
| `preferredMaintenanceWindow` | `string` | no |
| `shardCapacity` | `integer` | no |
| `shardCount` | `integer` | no |
| `shardInstanceCount` | `integer` | no |
| `subnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | yes |

