# Managed Streaming for Kafka

API version: 2018-11-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kafka/2018-11-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchAssociateScramSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `SecretArnList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `UnprocessedScramSecrets` | `List<UnprocessedScramSecret>` | no |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `ClusterArn` | `string` | yes |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `IcebergDestinationConfiguration` | `IcebergDestinationConfiguration` | no |
| `S3DestinationConfiguration` | `S3DestinationConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `TopicConfigurationList` | `List<TopicConfiguration>` | yes |
| `LoggingInfo` | `ChannelLoggingInfo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterOperationArn` | `string` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerNodeGroupInfo` | `BrokerNodeGroupInfo` | yes |
| `Rebalancing` | `Rebalancing` | no |
| `ClientAuthentication` | `ClientAuthentication` | no |
| `ClusterName` | `string` | yes |
| `ConfigurationInfo` | `ConfigurationInfo` | no |
| `EncryptionInfo` | `EncryptionInfo` | no |
| `EnhancedMonitoring` | `string` | no |
| `OpenMonitoring` | `OpenMonitoringInfo` | no |
| `KafkaVersion` | `string` | yes |
| `LoggingInfo` | `LoggingInfo` | no |
| `NumberOfBrokerNodes` | `integer` | yes |
| `Tags` | `Map<string>` | no |
| `StorageMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterName` | `string` | no |
| `State` | `string` | no |

## CreateClusterV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Provisioned` | `ProvisionedRequest` | no |
| `Serverless` | `ServerlessRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterName` | `string` | no |
| `State` | `string` | no |
| `ClusterType` | `string` | no |

## CreateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `KafkaVersions` | `List<string>` | no |
| `Name` | `string` | yes |
| `ServerProperties` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LatestRevision` | `ConfigurationRevision` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## CreateReplicator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `KafkaClusters` | `List<KafkaCluster>` | yes |
| `ReplicationInfoList` | `List<ReplicationInfo>` | yes |
| `ReplicatorName` | `string` | yes |
| `ServiceExecutionRoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `LogDelivery` | `LogDelivery` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicatorArn` | `string` | no |
| `ReplicatorName` | `string` | no |
| `ReplicatorState` | `string` | no |

## CreateTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `TopicName` | `string` | yes |
| `PartitionCount` | `integer` | yes |
| `ReplicationFactor` | `integer` | yes |
| `Configs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |
| `TopicName` | `string` | no |
| `Status` | `string` | no |

## CreateVpcConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetClusterArn` | `string` | yes |
| `Authentication` | `string` | yes |
| `VpcId` | `string` | yes |
| `ClientSubnets` | `List<string>` | yes |
| `SecurityGroups` | `List<string>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectionArn` | `string` | no |
| `State` | `string` | no |
| `Authentication` | `string` | no |
| `VpcId` | `string` | no |
| `ClientSubnets` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `State` | `string` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterOperationArn` | `string` | no |

## DeleteClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `State` | `string` | no |

## DeleteReplicator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CurrentVersion` | `string` | no |
| `ReplicatorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicatorArn` | `string` | no |
| `ReplicatorState` | `string` | no |

## DeleteTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `TopicName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |
| `TopicName` | `string` | no |
| `Status` | `string` | no |

## DeleteVpcConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectionArn` | `string` | no |
| `State` | `string` | no |

## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterInfo` | `ClusterInfo` | no |

## DescribeClusterV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterInfo` | `Cluster` | no |

## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelName` | `string` | yes |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `IcebergDestinationConfiguration` | `IcebergDestinationConfiguration` | no |
| `S3DestinationConfiguration` | `S3DestinationConfiguration` | no |
| `Status` | `string` | yes |
| `DestinationType` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `TopicConfigurationList` | `List<TopicConfiguration>` | yes |
| `LoggingInfo` | `ChannelLoggingInfo` | no |
| `StateInfo` | `ChannelStateInfo` | no |
| `ClusterOperationArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeClusterOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationInfo` | `ClusterOperationInfo` | no |

## DescribeClusterOperationV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationInfo` | `ClusterOperationV2` | no |

## DescribeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Description` | `string` | no |
| `KafkaVersions` | `List<string>` | no |
| `LatestRevision` | `ConfigurationRevision` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## DescribeConfigurationRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Revision` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Description` | `string` | no |
| `Revision` | `long` | no |
| `ServerProperties` | `blob` | no |

## DescribeReplicator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicatorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | no |
| `CurrentVersion` | `string` | no |
| `IsReplicatorReference` | `boolean` | no |
| `KafkaClusters` | `List<KafkaClusterDescription>` | no |
| `ReplicationInfoList` | `List<ReplicationInfoDescription>` | no |
| `ReplicatorArn` | `string` | no |
| `ReplicatorDescription` | `string` | no |
| `ReplicatorName` | `string` | no |
| `ReplicatorResourceArn` | `string` | no |
| `ReplicatorState` | `string` | no |
| `ServiceExecutionRoleArn` | `string` | no |
| `StateInfo` | `ReplicationStateInfo` | no |
| `Tags` | `Map<string>` | no |
| `LogDelivery` | `LogDelivery` | no |

## DescribeTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `TopicName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |
| `TopicName` | `string` | no |
| `ReplicationFactor` | `integer` | no |
| `PartitionCount` | `integer` | no |
| `Configs` | `string` | no |
| `Status` | `string` | no |

## DescribeTopicPartitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `TopicName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Partitions` | `List<TopicPartitionInfo>` | no |
| `NextToken` | `string` | no |

## DescribeVpcConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectionArn` | `string` | no |
| `TargetClusterArn` | `string` | no |
| `State` | `string` | no |
| `Authentication` | `string` | no |
| `VpcId` | `string` | no |
| `Subnets` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## BatchDisassociateScramSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `SecretArnList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `UnprocessedScramSecrets` | `List<UnprocessedScramSecret>` | no |

## GetBootstrapBrokers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BootstrapBrokerString` | `string` | no |
| `BootstrapBrokerStringTls` | `string` | no |
| `BootstrapBrokerStringSaslScram` | `string` | no |
| `BootstrapBrokerStringSaslIam` | `string` | no |
| `BootstrapBrokerStringPublicTls` | `string` | no |
| `BootstrapBrokerStringPublicSaslScram` | `string` | no |
| `BootstrapBrokerStringPublicSaslIam` | `string` | no |
| `BootstrapBrokerStringVpcConnectivityTls` | `string` | no |
| `BootstrapBrokerStringVpcConnectivitySaslScram` | `string` | no |
| `BootstrapBrokerStringVpcConnectivitySaslIam` | `string` | no |
| `BootstrapBrokerStringIpv6` | `string` | no |
| `BootstrapBrokerStringTlsIpv6` | `string` | no |
| `BootstrapBrokerStringSaslScramIpv6` | `string` | no |
| `BootstrapBrokerStringSaslIamIpv6` | `string` | no |

## GetCompatibleKafkaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleKafkaVersions` | `List<CompatibleKafkaVersion>` | no |

## GetClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CurrentVersion` | `string` | no |
| `Policy` | `string` | no |

## ListClusterOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationInfoList` | `List<ClusterOperationInfo>` | no |
| `NextToken` | `string` | no |

## ListClusterOperationsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterOperationInfoList` | `List<ClusterOperationV2Summary>` | no |
| `NextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterNameFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterInfoList` | `List<ClusterInfo>` | no |
| `NextToken` | `string` | no |

## ListClustersV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterNameFilter` | `string` | no |
| `ClusterTypeFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterInfoList` | `List<Cluster>` | no |
| `NextToken` | `string` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TopicNameFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelInfo>` | no |
| `NextToken` | `string` | no |

## ListConfigurationRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Revisions` | `List<ConfigurationRevision>` | no |

## ListConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configurations` | `List<Configuration>` | no |
| `NextToken` | `string` | no |

## ListKafkaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KafkaVersions` | `List<KafkaVersion>` | no |
| `NextToken` | `string` | no |

## ListNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NodeInfoList` | `List<NodeInfo>` | no |

## ListReplicators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ReplicatorNameFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Replicators` | `List<ReplicatorSummary>` | no |

## ListScramSecrets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SecretArnList` | `List<string>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListClientVpcConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpcConnections` | `List<ClientVpcConnection>` | no |
| `NextToken` | `string` | no |

## ListTopics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TopicNameFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Topics` | `List<TopicInfo>` | no |
| `NextToken` | `string` | no |

## ListVpcConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnections` | `List<VpcConnection>` | no |
| `NextToken` | `string` | no |

## RejectClientVpcConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `VpcConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutClusterPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | no |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CurrentVersion` | `string` | no |

## RebootBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerIds` | `List<string>` | yes |
| `ClusterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBrokerCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `TargetNumberOfBrokerNodes` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateBrokerType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `TargetInstanceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateBrokerStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `TargetBrokerEBSVolumeInfo` | `List<BrokerEBSVolumeInfo>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `ServerProperties` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `LatestRevision` | `ConfigurationRevision` | no |

## UpdateConnectivity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `ConnectivityInfo` | `ConnectivityInfo` | no |
| `CurrentVersion` | `string` | yes |
| `ZookeeperAccess` | `ZookeeperAccess` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterArn` | `string` | yes |
| `IcebergDestinationUpdate` | `IcebergDestinationUpdate` | no |
| `S3DestinationUpdate` | `S3DestinationUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ClusterOperationArn` | `string` | no |

## UpdateClusterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `ConfigurationInfo` | `ConfigurationInfo` | yes |
| `CurrentVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateClusterKafkaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `ConfigurationInfo` | `ConfigurationInfo` | no |
| `CurrentVersion` | `string` | yes |
| `TargetKafkaVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateMonitoring

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `EnhancedMonitoring` | `string` | no |
| `OpenMonitoring` | `OpenMonitoringInfo` | no |
| `LoggingInfo` | `LoggingInfo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateRebalancing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `Rebalancing` | `Rebalancing` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateReplicationInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumerGroupReplication` | `ConsumerGroupReplicationUpdate` | no |
| `CurrentVersion` | `string` | yes |
| `ReplicatorArn` | `string` | yes |
| `SourceKafkaClusterArn` | `string` | no |
| `SourceKafkaClusterId` | `string` | no |
| `TargetKafkaClusterArn` | `string` | no |
| `TargetKafkaClusterId` | `string` | no |
| `TopicReplication` | `TopicReplicationUpdate` | no |
| `LogDelivery` | `LogDelivery` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplicatorArn` | `string` | no |
| `ReplicatorState` | `string` | no |

## UpdateSecurity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientAuthentication` | `ClientAuthentication` | no |
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `EncryptionInfo` | `EncryptionInfo` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateStorage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `ProvisionedThroughput` | `ProvisionedThroughput` | no |
| `StorageMode` | `string` | no |
| `VolumeSizeGB` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | no |
| `ClusterOperationArn` | `string` | no |

## UpdateTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `TopicName` | `string` | yes |
| `Configs` | `string` | no |
| `PartitionCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |
| `TopicName` | `string` | no |
| `Status` | `string` | no |

