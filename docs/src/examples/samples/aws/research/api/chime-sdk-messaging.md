# Amazon Chime SDK Messaging

API version: 2021-05-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime-sdk-messaging/2021-05-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelFlowArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchCreateChannelMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `Type` | `string` | no |
| `MemberArns` | `List<string>` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BatchChannelMemberships` | `BatchChannelMemberships` | no |
| `Errors` | `List<BatchCreateChannelMembershipError>` | no |

## ChannelFlowCallback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackId` | `string` | yes |
| `ChannelArn` | `string` | yes |
| `DeleteResource` | `boolean` | no |
| `ChannelMessage` | `ChannelMessageCallback` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `CallbackId` | `string` | no |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `Name` | `string` | yes |
| `Mode` | `string` | no |
| `Privacy` | `string` | no |
| `Metadata` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ChimeBearer` | `string` | yes |
| `ChannelId` | `string` | no |
| `MemberArns` | `List<string>` | no |
| `ModeratorArns` | `List<string>` | no |
| `ElasticChannelConfiguration` | `ElasticChannelConfiguration` | no |
| `ExpirationSettings` | `ExpirationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |

## CreateChannelBan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Member` | `Identity` | no |

## CreateChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `Processors` | `List<Processor>` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | no |

## CreateChannelMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `Type` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Member` | `Identity` | no |
| `SubChannelId` | `string` | no |

## CreateChannelModerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelModeratorArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `ChannelModerator` | `Identity` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelBan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MessageId` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelModerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelModeratorArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMessagingStreamingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `Channel` | no |

## DescribeChannelBan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelBan` | `ChannelBan` | no |

## DescribeChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlow` | `ChannelFlow` | no |

## DescribeChannelMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelMembership` | `ChannelMembership` | no |

## DescribeChannelMembershipForAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `AppInstanceUserArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelMembership` | `ChannelMembershipForAppInstanceUserSummary` | no |

## DescribeChannelModeratedByAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `AppInstanceUserArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `ChannelModeratedByAppInstanceUserSummary` | no |

## DescribeChannelModerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelModeratorArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelModerator` | `ChannelModerator` | no |

## DisassociateChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChannelFlowArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetChannelMembershipPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Member` | `Identity` | no |
| `Preferences` | `ChannelMembershipPreferences` | no |

## GetChannelMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MessageId` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelMessage` | `ChannelMessage` | no |

## GetChannelMessageStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MessageId` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ChannelMessageStatusStructure` | no |

## GetMessagingSessionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoint` | `MessagingSessionEndpoint` | no |

## GetMessagingStreamingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingConfigurations` | `List<StreamingConfiguration>` | no |

## ListChannelBans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `NextToken` | `string` | no |
| `ChannelBans` | `List<ChannelBanSummary>` | no |

## ListChannelFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlows` | `List<ChannelFlowSummary>` | no |
| `NextToken` | `string` | no |

## ListChannelMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `Type` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `ChannelMemberships` | `List<ChannelMembershipSummary>` | no |
| `NextToken` | `string` | no |

## ListChannelMembershipsForAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelMemberships` | `List<ChannelMembershipForAppInstanceUserSummary>` | no |
| `NextToken` | `string` | no |

## ListChannelMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `SortOrder` | `string` | no |
| `NotBefore` | `timestamp` | no |
| `NotAfter` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `NextToken` | `string` | no |
| `ChannelMessages` | `List<ChannelMessageSummary>` | no |
| `SubChannelId` | `string` | no |

## ListChannelModerators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `NextToken` | `string` | no |
| `ChannelModerators` | `List<ChannelModeratorSummary>` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `Privacy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelSummary>` | no |
| `NextToken` | `string` | no |

## ListChannelsAssociatedWithChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelAssociatedWithFlowSummary>` | no |
| `NextToken` | `string` | no |

## ListChannelsModeratedByAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelModeratedByAppInstanceUserSummary>` | no |
| `NextToken` | `string` | no |

## ListSubChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `SubChannels` | `List<SubChannelSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutChannelExpirationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChimeBearer` | `string` | no |
| `ExpirationSettings` | `ExpirationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `ExpirationSettings` | `ExpirationSettings` | no |

## PutChannelMembershipPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MemberArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `Preferences` | `ChannelMembershipPreferences` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Member` | `Identity` | no |
| `Preferences` | `ChannelMembershipPreferences` | no |

## PutMessagingStreamingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `StreamingConfigurations` | `List<StreamingConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingConfigurations` | `List<StreamingConfiguration>` | no |

## RedactChannelMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MessageId` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `MessageId` | `string` | no |
| `SubChannelId` | `string` | no |

## SearchChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChimeBearer` | `string` | no |
| `Fields` | `List<SearchField>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelSummary>` | no |
| `NextToken` | `string` | no |

## SendChannelMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `Content` | `string` | yes |
| `Type` | `string` | yes |
| `Persistence` | `string` | yes |
| `Metadata` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `ChimeBearer` | `string` | yes |
| `PushNotification` | `PushNotificationConfiguration` | no |
| `MessageAttributes` | `Map<MessageAttributeValue>` | no |
| `SubChannelId` | `string` | no |
| `ContentType` | `string` | no |
| `Target` | `List<Target>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `MessageId` | `string` | no |
| `Status` | `ChannelMessageStatusStructure` | no |
| `SubChannelId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `Name` | `string` | no |
| `Mode` | `string` | no |
| `Metadata` | `string` | no |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |

## UpdateChannelFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | yes |
| `Processors` | `List<Processor>` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelFlowArn` | `string` | no |

## UpdateChannelMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `MessageId` | `string` | yes |
| `Content` | `string` | yes |
| `Metadata` | `string` | no |
| `ChimeBearer` | `string` | yes |
| `SubChannelId` | `string` | no |
| `ContentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `MessageId` | `string` | no |
| `Status` | `ChannelMessageStatusStructure` | no |
| `SubChannelId` | `string` | no |

## UpdateChannelReadMarker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | yes |
| `ChimeBearer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |

