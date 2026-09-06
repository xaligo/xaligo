# Amazon Kinesis

API version: 2013-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesis/2013-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `Tags` | `Map<string>` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `ServiceExecutionRoleARN` | `string` | yes |
| `StreamConfigurationList` | `List<ChannelStreamConfiguration>` | yes |
| `S3DestinationConfiguration` | `S3DestinationConfiguration` | no |
| `S3TablesDestinationConfiguration` | `S3TablesDestinationConfiguration` | no |
| `EncryptionConfiguration` | `ChannelEncryptionConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `LoggingConfiguration` | `ChannelLoggingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelDescription` | `ChannelDescription` | yes |

## CreateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | yes |
| `ShardCount` | `integer` | no |
| `StreamModeDetails` | `StreamModeDetails` | no |
| `Tags` | `Map<string>` | no |
| `WarmThroughputMiBps` | `integer` | no |
| `MaxRecordSizeInKiB` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DecreaseStreamRetentionPeriod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `RetentionPeriodHours` | `integer` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `EnforceConsumerDeletion` | `boolean` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterStreamConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `ConsumerName` | `string` | no |
| `ConsumerARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MinimumThroughputBillingCommitment` | `MinimumThroughputBillingCommitmentOutput` | no |

## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelDescription` | `ChannelDescription` | yes |

## DescribeLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardLimit` | `integer` | yes |
| `OpenShardCount` | `integer` | yes |
| `OnDemandStreamCount` | `integer` | yes |
| `OnDemandStreamCountLimit` | `integer` | yes |
| `ChannelCount` | `integer` | no |
| `ChannelCountLimit` | `integer` | no |

## DescribeStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `Limit` | `integer` | no |
| `ExclusiveStartShardId` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamDescription` | `StreamDescription` | yes |

## DescribeStreamConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `ConsumerName` | `string` | no |
| `ConsumerARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumerDescription` | `ConsumerDescription` | yes |

## DescribeStreamSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamDescriptionSummary` | `StreamDescriptionSummary` | yes |

## DisableEnhancedMonitoring

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ShardLevelMetrics` | `List<string>` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `CurrentShardLevelMetrics` | `List<string>` | no |
| `DesiredShardLevelMetrics` | `List<string>` | no |
| `StreamARN` | `string` | no |

## EnableEnhancedMonitoring

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ShardLevelMetrics` | `List<string>` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `CurrentShardLevelMetrics` | `List<string>` | no |
| `DesiredShardLevelMetrics` | `List<string>` | no |
| `StreamARN` | `string` | no |

## GetRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardIterator` | `string` | yes |
| `Limit` | `integer` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | yes |
| `NextShardIterator` | `string` | no |
| `MillisBehindLatest` | `long` | no |
| `ChildShards` | `List<ChildShard>` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | yes |

## GetShardIterator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ShardId` | `string` | yes |
| `ShardIteratorType` | `string` | yes |
| `StartingSequenceNumber` | `string` | no |
| `Timestamp` | `timestamp` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardIterator` | `string` | no |

## IncreaseStreamRetentionPeriod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `RetentionPeriodHours` | `integer` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamFilter` | `List<StreamFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelSummaries` | `List<ChannelSummary>` | yes |
| `NextToken` | `string` | no |

## ListShards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `NextToken` | `string` | no |
| `ExclusiveStartShardId` | `string` | no |
| `MaxResults` | `integer` | no |
| `StreamCreationTimestamp` | `timestamp` | no |
| `ShardFilter` | `ShardFilter` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Shards` | `List<Shard>` | no |
| `NextToken` | `string` | no |

## ListStreamConsumers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StreamCreationTimestamp` | `timestamp` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Consumers` | `List<Consumer>` | no |
| `NextToken` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `ExclusiveStartStreamName` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamNames` | `List<string>` | yes |
| `HasMoreStreams` | `boolean` | yes |
| `NextToken` | `string` | no |
| `StreamSummaries` | `List<StreamSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTagsForStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ExclusiveStartTagKey` | `string` | no |
| `Limit` | `integer` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `HasMoreTags` | `boolean` | yes |

## MergeShards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ShardToMerge` | `string` | yes |
| `AdjacentShardToMerge` | `string` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `Data` | `blob` | yes |
| `PartitionKey` | `string` | yes |
| `ExplicitHashKey` | `string` | no |
| `SequenceNumberForOrdering` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardId` | `string` | yes |
| `SequenceNumber` | `string` | yes |
| `EncryptionType` | `string` | no |

## PutRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<PutRecordsRequestEntry>` | yes |
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRecordCount` | `integer` | no |
| `Records` | `List<PutRecordsResultEntry>` | yes |
| `EncryptionType` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterStreamConsumer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | yes |
| `ConsumerName` | `string` | yes |
| `StreamId` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Consumer` | `Consumer` | yes |

## RemoveTagsFromStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `TagKeys` | `List<string>` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SplitShard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `ShardToSplit` | `string` | yes |
| `NewStartingHashKey` | `string` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartStreamEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `EncryptionType` | `string` | yes |
| `KeyId` | `string` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopStreamEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `EncryptionType` | `string` | yes |
| `KeyId` | `string` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubscribeToShard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConsumerARN` | `string` | yes |
| `StreamId` | `string` | no |
| `ShardId` | `string` | yes |
| `StartingPosition` | `StartingPosition` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStream` | `SubscribeToShardEventStream` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | yes |
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagKeys` | `List<string>` | yes |
| `ResourceARN` | `string` | yes |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MinimumThroughputBillingCommitment` | `MinimumThroughputBillingCommitmentInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MinimumThroughputBillingCommitment` | `MinimumThroughputBillingCommitmentOutput` | no |

## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `S3DestinationConfiguration` | `S3DestinationUpdateInput` | no |
| `S3TablesDestinationConfiguration` | `S3TablesDestinationUpdateInput` | no |
| `LoggingConfiguration` | `ChannelLoggingUpdateInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelDescription` | `ChannelDescription` | yes |

## UpdateMaxRecordSize

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |
| `MaxRecordSizeInKiB` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateShardCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `TargetShardCount` | `integer` | yes |
| `ScalingType` | `string` | yes |
| `StreamARN` | `string` | no |
| `StreamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `CurrentShardCount` | `integer` | no |
| `TargetShardCount` | `integer` | no |
| `StreamARN` | `string` | no |

## UpdateStreamMode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | yes |
| `StreamId` | `string` | no |
| `StreamModeDetails` | `StreamModeDetails` | yes |
| `WarmThroughputMiBps` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStreamWarmThroughput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `StreamName` | `string` | no |
| `StreamId` | `string` | no |
| `WarmThroughputMiBps` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `StreamName` | `string` | no |
| `WarmThroughput` | `WarmThroughputObject` | no |

