# Amazon Pinpoint

API version: 2016-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pinpoint/2016-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateApplicationRequest` | `CreateApplicationRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationResponse` | `ApplicationResponse` | yes |

## CreateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `WriteCampaignRequest` | `WriteCampaignRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignResponse` | `CampaignResponse` | yes |

## CreateEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailTemplateRequest` | `EmailTemplateRequest` | yes |
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateTemplateMessageBody` | `CreateTemplateMessageBody` | yes |

## CreateExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ExportJobRequest` | `ExportJobRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobResponse` | `ExportJobResponse` | yes |

## CreateImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ImportJobRequest` | `ImportJobRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobResponse` | `ImportJobResponse` | yes |

## CreateInAppTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InAppTemplateRequest` | `InAppTemplateRequest` | yes |
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateCreateMessageBody` | `TemplateCreateMessageBody` | yes |

## CreateJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `WriteJourneyRequest` | `WriteJourneyRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyResponse` | `JourneyResponse` | yes |

## CreatePushTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PushNotificationTemplateRequest` | `PushNotificationTemplateRequest` | yes |
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateTemplateMessageBody` | `CreateTemplateMessageBody` | yes |

## CreateRecommenderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateRecommenderConfiguration` | `CreateRecommenderConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderConfigurationResponse` | `RecommenderConfigurationResponse` | yes |

## CreateSegment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `WriteSegmentRequest` | `WriteSegmentRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentResponse` | `SegmentResponse` | yes |

## CreateSmsTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSTemplateRequest` | `SMSTemplateRequest` | yes |
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateTemplateMessageBody` | `CreateTemplateMessageBody` | yes |

## CreateVoiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `VoiceTemplateRequest` | `VoiceTemplateRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateTemplateMessageBody` | `CreateTemplateMessageBody` | yes |

## DeleteAdmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ADMChannelResponse` | `ADMChannelResponse` | yes |

## DeleteApnsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSChannelResponse` | `APNSChannelResponse` | yes |

## DeleteApnsSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSSandboxChannelResponse` | `APNSSandboxChannelResponse` | yes |

## DeleteApnsVoipChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipChannelResponse` | `APNSVoipChannelResponse` | yes |

## DeleteApnsVoipSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipSandboxChannelResponse` | `APNSVoipSandboxChannelResponse` | yes |

## DeleteApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationResponse` | `ApplicationResponse` | yes |

## DeleteBaiduChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaiduChannelResponse` | `BaiduChannelResponse` | yes |

## DeleteCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignResponse` | `CampaignResponse` | yes |

## DeleteEmailChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailChannelResponse` | `EmailChannelResponse` | yes |

## DeleteEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointResponse` | `EndpointResponse` | yes |

## DeleteEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStream` | `EventStream` | yes |

## DeleteGcmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GCMChannelResponse` | `GCMChannelResponse` | yes |

## DeleteInAppTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## DeleteJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyResponse` | `JourneyResponse` | yes |

## DeletePushTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## DeleteRecommenderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderConfigurationResponse` | `RecommenderConfigurationResponse` | yes |

## DeleteSegment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SegmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentResponse` | `SegmentResponse` | yes |

## DeleteSmsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSChannelResponse` | `SMSChannelResponse` | yes |

## DeleteSmsTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## DeleteUserEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointsResponse` | `EndpointsResponse` | yes |

## DeleteVoiceChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceChannelResponse` | `VoiceChannelResponse` | yes |

## DeleteVoiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## GetAdmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ADMChannelResponse` | `ADMChannelResponse` | yes |

## GetApnsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSChannelResponse` | `APNSChannelResponse` | yes |

## GetApnsSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSSandboxChannelResponse` | `APNSSandboxChannelResponse` | yes |

## GetApnsVoipChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipChannelResponse` | `APNSVoipChannelResponse` | yes |

## GetApnsVoipSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipSandboxChannelResponse` | `APNSVoipSandboxChannelResponse` | yes |

## GetApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationResponse` | `ApplicationResponse` | yes |

## GetApplicationDateRangeKpi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndTime` | `timestamp` | no |
| `KpiName` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `StartTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDateRangeKpiResponse` | `ApplicationDateRangeKpiResponse` | yes |

## GetApplicationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSettingsResource` | `ApplicationSettingsResource` | yes |

## GetApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationsResponse` | `ApplicationsResponse` | yes |

## GetBaiduChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaiduChannelResponse` | `BaiduChannelResponse` | yes |

## GetCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignResponse` | `CampaignResponse` | yes |

## GetCampaignActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivitiesResponse` | `ActivitiesResponse` | yes |

## GetCampaignDateRangeKpi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |
| `EndTime` | `timestamp` | no |
| `KpiName` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `StartTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignDateRangeKpiResponse` | `CampaignDateRangeKpiResponse` | yes |

## GetCampaignVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |
| `Version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignResponse` | `CampaignResponse` | yes |

## GetCampaignVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignsResponse` | `CampaignsResponse` | yes |

## GetCampaigns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignsResponse` | `CampaignsResponse` | yes |

## GetChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelsResponse` | `ChannelsResponse` | yes |

## GetEmailChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailChannelResponse` | `EmailChannelResponse` | yes |

## GetEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailTemplateResponse` | `EmailTemplateResponse` | yes |

## GetEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointResponse` | `EndpointResponse` | yes |

## GetEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStream` | `EventStream` | yes |

## GetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobResponse` | `ExportJobResponse` | yes |

## GetExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobsResponse` | `ExportJobsResponse` | yes |

## GetGcmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GCMChannelResponse` | `GCMChannelResponse` | yes |

## GetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobResponse` | `ImportJobResponse` | yes |

## GetImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobsResponse` | `ImportJobsResponse` | yes |

## GetInAppMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InAppMessagesResponse` | `InAppMessagesResponse` | yes |

## GetInAppTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InAppTemplateResponse` | `InAppTemplateResponse` | yes |

## GetJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyResponse` | `JourneyResponse` | yes |

## GetJourneyDateRangeKpi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndTime` | `timestamp` | no |
| `JourneyId` | `string` | yes |
| `KpiName` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `StartTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyDateRangeKpiResponse` | `JourneyDateRangeKpiResponse` | yes |

## GetJourneyExecutionActivityMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyActivityId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyExecutionActivityMetricsResponse` | `JourneyExecutionActivityMetricsResponse` | yes |

## GetJourneyExecutionMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyExecutionMetricsResponse` | `JourneyExecutionMetricsResponse` | yes |

## GetJourneyRunExecutionActivityMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyActivityId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyRunExecutionActivityMetricsResponse` | `JourneyRunExecutionActivityMetricsResponse` | yes |

## GetJourneyRunExecutionMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyRunExecutionMetricsResponse` | `JourneyRunExecutionMetricsResponse` | yes |

## GetJourneyRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyRunsResponse` | `JourneyRunsResponse` | yes |

## GetPushTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PushNotificationTemplateResponse` | `PushNotificationTemplateResponse` | yes |

## GetRecommenderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderConfigurationResponse` | `RecommenderConfigurationResponse` | yes |

## GetRecommenderConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListRecommenderConfigurationsResponse` | `ListRecommenderConfigurationsResponse` | yes |

## GetSegment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SegmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentResponse` | `SegmentResponse` | yes |

## GetSegmentExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `SegmentId` | `string` | yes |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobsResponse` | `ExportJobsResponse` | yes |

## GetSegmentImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `SegmentId` | `string` | yes |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobsResponse` | `ImportJobsResponse` | yes |

## GetSegmentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SegmentId` | `string` | yes |
| `Version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentResponse` | `SegmentResponse` | yes |

## GetSegmentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `SegmentId` | `string` | yes |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentsResponse` | `SegmentsResponse` | yes |

## GetSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentsResponse` | `SegmentsResponse` | yes |

## GetSmsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSChannelResponse` | `SMSChannelResponse` | yes |

## GetSmsTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSTemplateResponse` | `SMSTemplateResponse` | yes |

## GetUserEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointsResponse` | `EndpointsResponse` | yes |

## GetVoiceChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceChannelResponse` | `VoiceChannelResponse` | yes |

## GetVoiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceTemplateResponse` | `VoiceTemplateResponse` | yes |

## ListJourneys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `PageSize` | `string` | no |
| `Token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneysResponse` | `JourneysResponse` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagsModel` | `TagsModel` | yes |

## ListTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `TemplateName` | `string` | yes |
| `TemplateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateVersionsResponse` | `TemplateVersionsResponse` | yes |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `string` | no |
| `Prefix` | `string` | no |
| `TemplateType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplatesResponse` | `TemplatesResponse` | yes |

## PhoneNumberValidate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumberValidateRequest` | `NumberValidateRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NumberValidateResponse` | `NumberValidateResponse` | yes |

## PutEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `WriteEventStream` | `WriteEventStream` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStream` | `EventStream` | yes |

## PutEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EventsRequest` | `EventsRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventsResponse` | `EventsResponse` | yes |

## RemoveAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `AttributeType` | `string` | yes |
| `UpdateAttributesRequest` | `UpdateAttributesRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttributesResource` | `AttributesResource` | yes |

## SendMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MessageRequest` | `MessageRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageResponse` | `MessageResponse` | yes |

## SendOTPMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SendOTPMessageRequestParameters` | `SendOTPMessageRequestParameters` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageResponse` | `MessageResponse` | yes |

## SendUsersMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SendUsersMessageRequest` | `SendUsersMessageRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SendUsersMessageResponse` | `SendUsersMessageResponse` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagsModel` | `TagsModel` | yes |

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


## UpdateAdmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ADMChannelRequest` | `ADMChannelRequest` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ADMChannelResponse` | `ADMChannelResponse` | yes |

## UpdateApnsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSChannelRequest` | `APNSChannelRequest` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSChannelResponse` | `APNSChannelResponse` | yes |

## UpdateApnsSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSSandboxChannelRequest` | `APNSSandboxChannelRequest` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSSandboxChannelResponse` | `APNSSandboxChannelResponse` | yes |

## UpdateApnsVoipChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipChannelRequest` | `APNSVoipChannelRequest` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipChannelResponse` | `APNSVoipChannelResponse` | yes |

## UpdateApnsVoipSandboxChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipSandboxChannelRequest` | `APNSVoipSandboxChannelRequest` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APNSVoipSandboxChannelResponse` | `APNSVoipSandboxChannelResponse` | yes |

## UpdateApplicationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `WriteApplicationSettingsRequest` | `WriteApplicationSettingsRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSettingsResource` | `ApplicationSettingsResource` | yes |

## UpdateBaiduChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `BaiduChannelRequest` | `BaiduChannelRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaiduChannelResponse` | `BaiduChannelResponse` | yes |

## UpdateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CampaignId` | `string` | yes |
| `WriteCampaignRequest` | `WriteCampaignRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignResponse` | `CampaignResponse` | yes |

## UpdateEmailChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EmailChannelRequest` | `EmailChannelRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailChannelResponse` | `EmailChannelResponse` | yes |

## UpdateEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateNewVersion` | `boolean` | no |
| `EmailTemplateRequest` | `EmailTemplateRequest` | yes |
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndpointId` | `string` | yes |
| `EndpointRequest` | `EndpointRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateEndpointsBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EndpointBatchRequest` | `EndpointBatchRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateGcmChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `GCMChannelRequest` | `GCMChannelRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GCMChannelResponse` | `GCMChannelResponse` | yes |

## UpdateInAppTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateNewVersion` | `boolean` | no |
| `InAppTemplateRequest` | `InAppTemplateRequest` | yes |
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateJourney

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `WriteJourneyRequest` | `WriteJourneyRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyResponse` | `JourneyResponse` | yes |

## UpdateJourneyState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `JourneyId` | `string` | yes |
| `JourneyStateRequest` | `JourneyStateRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JourneyResponse` | `JourneyResponse` | yes |

## UpdatePushTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateNewVersion` | `boolean` | no |
| `PushNotificationTemplateRequest` | `PushNotificationTemplateRequest` | yes |
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateRecommenderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderId` | `string` | yes |
| `UpdateRecommenderConfiguration` | `UpdateRecommenderConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderConfigurationResponse` | `RecommenderConfigurationResponse` | yes |

## UpdateSegment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SegmentId` | `string` | yes |
| `WriteSegmentRequest` | `WriteSegmentRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentResponse` | `SegmentResponse` | yes |

## UpdateSmsChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `SMSChannelRequest` | `SMSChannelRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSChannelResponse` | `SMSChannelResponse` | yes |

## UpdateSmsTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateNewVersion` | `boolean` | no |
| `SMSTemplateRequest` | `SMSTemplateRequest` | yes |
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateTemplateActiveVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateActiveVersionRequest` | `TemplateActiveVersionRequest` | yes |
| `TemplateName` | `string` | yes |
| `TemplateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## UpdateVoiceChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `VoiceChannelRequest` | `VoiceChannelRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceChannelResponse` | `VoiceChannelResponse` | yes |

## UpdateVoiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateNewVersion` | `boolean` | no |
| `TemplateName` | `string` | yes |
| `Version` | `string` | no |
| `VoiceTemplateRequest` | `VoiceTemplateRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageBody` | `MessageBody` | yes |

## VerifyOTPMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `VerifyOTPMessageRequestParameters` | `VerifyOTPMessageRequestParameters` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationResponse` | `VerificationResponse` | yes |

