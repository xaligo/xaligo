# AmazonConnectCampaignServiceV2

API version: 2024-04-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connectcampaignsv2/2024-04-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `connectInstanceId` | `string` | yes |
| `channelSubtypeConfig` | `ChannelSubtypeConfig` | no |
| `type` | `string` | no |
| `source` | `Source` | no |
| `connectCampaignFlowArn` | `string` | no |
| `schedule` | `Schedule` | no |
| `entryLimitsConfig` | `EntryLimitsConfig` | no |
| `communicationTimeConfig` | `CommunicationTimeConfig` | no |
| `communicationLimitsOverride` | `CommunicationLimitsConfig` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |

## DeleteCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCampaignChannelSubtypeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `channelSubtype` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCampaignCommunicationLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `config` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCampaignCommunicationTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `config` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCampaignEntryLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectInstanceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `campaignDeletionPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectInstanceIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `integrationIdentifier` | `IntegrationIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstanceOnboardingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaign` | `Campaign` | no |

## GetCampaignState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `state` | `string` | no |

## GetCampaignStateBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRequests` | `List<SuccessfulCampaignStateResponse>` | no |
| `failedRequests` | `List<FailedCampaignStateResponse>` | no |

## GetConnectInstanceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceConfig` | `InstanceConfig` | no |

## GetInstanceCommunicationLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `communicationLimitsConfig` | `InstanceCommunicationLimitsConfig` | no |

## GetInstanceOnboardingJobStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceOnboardingJobStatus` | `InstanceOnboardingJobStatus` | no |

## ListCampaigns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `CampaignFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `campaignSummaryList` | `List<CampaignSummary>` | no |

## ListConnectInstanceIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `integrationSummaryList` | `List<IntegrationSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PauseCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConnectInstanceIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `integrationConfig` | `IntegrationConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutInstanceCommunicationLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `communicationLimitsConfig` | `InstanceCommunicationLimitsConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutOutboundRequestBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `outboundRequests` | `List<OutboundRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRequests` | `List<SuccessfulRequest>` | no |
| `failedRequests` | `List<FailedRequest>` | no |

## PutProfileOutboundRequestBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `profileOutboundRequests` | `List<ProfileOutboundRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRequests` | `List<SuccessfulProfileOutboundRequest>` | no |
| `failedRequests` | `List<FailedProfileOutboundRequest>` | no |

## ResumeCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartInstanceOnboardingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |
| `encryptionConfig` | `EncryptionConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceOnboardingJobStatus` | `InstanceOnboardingJobStatus` | no |

## StopCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignChannelSubtypeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `channelSubtypeConfig` | `ChannelSubtypeConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignCommunicationLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `communicationLimitsOverride` | `CommunicationLimitsConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignCommunicationTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `communicationTimeConfig` | `CommunicationTimeConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignEntryLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `entryLimitsConfig` | `EntryLimitsConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignFlowAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `connectCampaignFlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `schedule` | `Schedule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaignSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `source` | `Source` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


