# AWS Parallel Computing Service

API version: 2023-02-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pcs/2023-02-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `scheduler` | `SchedulerRequest` | yes |
| `size` | `string` | yes |
| `networking` | `NetworkingRequest` | yes |
| `slurmConfiguration` | `ClusterSlurmConfigurationRequest` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## CreateComputeNodeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `computeNodeGroupName` | `string` | yes |
| `amiId` | `string` | no |
| `subnetIds` | `List<string>` | yes |
| `purchaseOption` | `string` | no |
| `customLaunchTemplate` | `CustomLaunchTemplate` | yes |
| `iamInstanceProfileArn` | `string` | yes |
| `scalingConfiguration` | `ScalingConfigurationRequest` | yes |
| `instanceConfigs` | `List<InstanceConfig>` | yes |
| `spotOptions` | `SpotOptions` | no |
| `slurmConfiguration` | `ComputeNodeGroupSlurmConfigurationRequest` | no |
| `nodeLifecycleActions` | `NodeLifecycleActionsRequest` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeNodeGroup` | `ComputeNodeGroup` | no |

## CreateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `queueName` | `string` | yes |
| `computeNodeGroupConfigurations` | `List<ComputeNodeGroupConfiguration>` | no |
| `slurmConfiguration` | `QueueSlurmConfigurationRequest` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queue` | `Queue` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteComputeNodeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `computeNodeGroupIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `queueIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## GetComputeNodeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `computeNodeGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeNodeGroup` | `ComputeNodeGroup` | no |

## GetQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `queueIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queue` | `Queue` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusters` | `List<ClusterSummary>` | yes |
| `nextToken` | `string` | no |

## ListComputeNodeGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeNodeGroups` | `List<ComputeNodeGroupSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queues` | `List<QueueSummary>` | yes |
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

## RegisterComputeNodeGroupInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `bootstrapId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodeID` | `string` | yes |
| `sharedSecret` | `string` | yes |
| `endpoints` | `List<Endpoint>` | yes |
| `clusterName` | `string` | no |
| `computeNodeGroupId` | `string` | no |
| `computeNodeGroupName` | `string` | no |
| `nodeLifecycleActions` | `NodeLifecycleActions` | no |

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
| `clusterIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `slurmConfiguration` | `UpdateClusterSlurmConfigurationRequest` | no |
| `scheduler` | `UpdateSchedulerRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## UpdateComputeNodeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `computeNodeGroupIdentifier` | `string` | yes |
| `amiId` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `customLaunchTemplate` | `CustomLaunchTemplate` | no |
| `purchaseOption` | `string` | no |
| `spotOptions` | `SpotOptions` | no |
| `scalingConfiguration` | `ScalingConfigurationRequest` | no |
| `iamInstanceProfileArn` | `string` | no |
| `slurmConfiguration` | `UpdateComputeNodeGroupSlurmConfigurationRequest` | no |
| `nodeLifecycleActions` | `UpdateNodeLifecycleActionsRequest` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeNodeGroup` | `ComputeNodeGroup` | no |

## UpdateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterIdentifier` | `string` | yes |
| `queueIdentifier` | `string` | yes |
| `computeNodeGroupConfigurations` | `List<ComputeNodeGroupConfiguration>` | no |
| `slurmConfiguration` | `UpdateQueueSlurmConfigurationRequest` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queue` | `Queue` | no |

