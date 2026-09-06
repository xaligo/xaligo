# Amazon Kinesis Video Streams

API version: 2017-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesisvideo/2017-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSignalingChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | yes |
| `ChannelType` | `string` | no |
| `SingleMasterConfiguration` | `SingleMasterConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | no |

## CreateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceName` | `string` | no |
| `StreamName` | `string` | yes |
| `MediaType` | `string` | no |
| `KmsKeyId` | `string` | no |
| `DataRetentionInHours` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `StreamStorageConfiguration` | `StreamStorageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |

## DeleteEdgeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSignalingChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `CurrentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | yes |
| `CurrentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeEdgeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `SyncStatus` | `string` | no |
| `FailedStatusDetails` | `string` | no |
| `EdgeConfig` | `EdgeConfig` | no |
| `EdgeAgentStatus` | `EdgeAgentStatus` | no |

## DescribeImageGenerationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageGenerationConfiguration` | `ImageGenerationConfiguration` | no |

## DescribeMappedResourceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MappedResourceConfigurationList` | `List<MappedResourceConfigurationListItem>` | no |
| `NextToken` | `string` | no |

## DescribeMediaStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | no |
| `ChannelARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaStorageConfiguration` | `MediaStorageConfiguration` | no |

## DescribeNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationConfiguration` | `NotificationConfiguration` | no |

## DescribeSignalingChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelName` | `string` | no |
| `ChannelARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelInfo` | `ChannelInfo` | no |

## DescribeStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamInfo` | `StreamInfo` | no |

## DescribeStreamStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamStorageConfiguration` | `StreamStorageConfiguration` | no |

## GetDataEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `APIName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataEndpoint` | `string` | no |

## GetSignalingChannelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `SingleMasterChannelEndpointConfiguration` | `SingleMasterChannelEndpointConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceEndpointList` | `List<ResourceEndpointListItem>` | no |

## ListEdgeAgentConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubDeviceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeConfigs` | `List<ListEdgeAgentConfigurationsEdgeConfig>` | no |
| `NextToken` | `string` | no |

## ListSignalingChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChannelNameCondition` | `ChannelNameCondition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelInfoList` | `List<ChannelInfo>` | no |
| `NextToken` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `StreamNameCondition` | `StreamNameCondition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamInfoList` | `List<StreamInfo>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Tags` | `Map<string>` | no |

## ListTagsForStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StreamARN` | `string` | no |
| `StreamName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Tags` | `Map<string>` | no |

## StartEdgeConfigurationUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `EdgeConfig` | `EdgeConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `SyncStatus` | `string` | no |
| `FailedStatusDetails` | `string` | no |
| `EdgeConfig` | `EdgeConfig` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `StreamName` | `string` | no |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeyList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamARN` | `string` | no |
| `StreamName` | `string` | no |
| `TagKeyList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataRetention

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `CurrentVersion` | `string` | yes |
| `Operation` | `string` | yes |
| `DataRetentionChangeInHours` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateImageGenerationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `ImageGenerationConfiguration` | `ImageGenerationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMediaStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `MediaStorageConfiguration` | `MediaStorageConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `NotificationConfiguration` | `NotificationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSignalingChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelARN` | `string` | yes |
| `CurrentVersion` | `string` | yes |
| `SingleMasterConfiguration` | `SingleMasterConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `CurrentVersion` | `string` | yes |
| `DeviceName` | `string` | no |
| `MediaType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStreamStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamName` | `string` | no |
| `StreamARN` | `string` | no |
| `CurrentVersion` | `string` | yes |
| `StreamStorageConfiguration` | `StreamStorageConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


