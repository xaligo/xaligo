# AmazonConnectCampaignService

API version: 2021-01-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connectcampaigns/2021-01-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `connectInstanceId` | `string` | yes |
| `dialerConfig` | `DialerConfig` | yes |
| `outboundCallConfig` | `OutboundCallConfig` | yes |
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


## DeleteConnectInstanceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectInstanceId` | `string` | yes |

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


## PutDialRequestBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dialRequests` | `List<DialRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRequests` | `List<SuccessfulRequest>` | no |
| `failedRequests` | `List<FailedRequest>` | no |

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


## UpdateCampaignDialerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dialerConfig` | `DialerConfig` | yes |

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


## UpdateCampaignOutboundCallConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `connectContactFlowId` | `string` | no |
| `connectSourcePhoneNumber` | `string` | no |
| `answerMachineDetectionConfig` | `AnswerMachineDetectionConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


