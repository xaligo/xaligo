# Amazon DynamoDB Accelerator (DAX)

API version: 2017-04-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dax/2017-04-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeType` | `string` | yes |
| `Description` | `string` | no |
| `ReplicationFactor` | `integer` | yes |
| `AvailabilityZones` | `List<string>` | no |
| `SubnetGroupName` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `NotificationTopicArn` | `string` | no |
| `IamRoleArn` | `string` | yes |
| `ParameterGroupName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SSESpecification` | `SSESpecification` | no |
| `ClusterEndpointEncryptionType` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## CreateParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroup` | `ParameterGroup` | no |

## CreateSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | yes |
| `Description` | `string` | no |
| `SubnetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroup` | `SubnetGroup` | no |

## DecreaseReplicationFactor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NewReplicationFactor` | `integer` | yes |
| `AvailabilityZones` | `List<string>` | no |
| `NodeIdsToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DeleteParameterGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionMessage` | `string` | no |

## DeleteSubnetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionMessage` | `string` | no |

## DescribeClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterNames` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Clusters` | `List<Cluster>` | no |

## DescribeDefaultParameters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Parameters` | `List<Parameter>` | no |

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

## DescribeParameterGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParameterGroupNames` | `List<string>` | no |
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
| `Source` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Parameters` | `List<Parameter>` | no |

## DescribeSubnetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetGroupNames` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SubnetGroups` | `List<SubnetGroup>` | no |

## IncreaseReplicationFactor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NewReplicationFactor` | `integer` | yes |
| `AvailabilityZones` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## RebootNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `Description` | `string` | no |
| `PreferredMaintenanceWindow` | `string` | no |
| `NotificationTopicArn` | `string` | no |
| `NotificationTopicStatus` | `string` | no |
| `ParameterGroupName` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

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

