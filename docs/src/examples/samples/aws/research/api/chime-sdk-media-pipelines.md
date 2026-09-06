# Amazon Chime SDK Media Pipelines

API version: 2021-07-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime-sdk-media-pipelines/2021-07-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMediaCapturePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceType` | `string` | yes |
| `SourceArn` | `string` | yes |
| `SinkType` | `string` | yes |
| `SinkArn` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `ChimeSdkMeetingConfiguration` | `ChimeSdkMeetingConfiguration` | no |
| `SseAwsKeyManagementParams` | `SseAwsKeyManagementParams` | no |
| `SinkIamRoleArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaCapturePipeline` | `MediaCapturePipeline` | no |

## CreateMediaConcatenationPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sources` | `List<ConcatenationSource>` | yes |
| `Sinks` | `List<ConcatenationSink>` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaConcatenationPipeline` | `MediaConcatenationPipeline` | no |

## CreateMediaInsightsPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfigurationArn` | `string` | yes |
| `KinesisVideoStreamSourceRuntimeConfiguration` | `KinesisVideoStreamSourceRuntimeConfiguration` | no |
| `MediaInsightsRuntimeMetadata` | `Map<string>` | no |
| `KinesisVideoStreamRecordingSourceRuntimeConfiguration` | `KinesisVideoStreamRecordingSourceRuntimeConfiguration` | no |
| `S3RecordingSinkRuntimeConfiguration` | `S3RecordingSinkRuntimeConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipeline` | `MediaInsightsPipeline` | yes |

## CreateMediaInsightsPipelineConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfigurationName` | `string` | yes |
| `ResourceAccessRoleArn` | `string` | yes |
| `RealTimeAlertConfiguration` | `RealTimeAlertConfiguration` | no |
| `Elements` | `List<MediaInsightsPipelineConfigurationElement>` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfiguration` | `MediaInsightsPipelineConfiguration` | no |

## CreateMediaLiveConnectorPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sources` | `List<LiveConnectorSourceConfiguration>` | yes |
| `Sinks` | `List<LiveConnectorSinkConfiguration>` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaLiveConnectorPipeline` | `MediaLiveConnectorPipeline` | no |

## CreateMediaPipelineKinesisVideoStreamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamConfiguration` | `KinesisVideoStreamConfiguration` | yes |
| `PoolName` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KinesisVideoStreamPoolConfiguration` | `KinesisVideoStreamPoolConfiguration` | no |

## CreateMediaStreamPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sources` | `List<MediaStreamSource>` | yes |
| `Sinks` | `List<MediaStreamSink>` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaStreamPipeline` | `MediaStreamPipeline` | no |

## DeleteMediaCapturePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipelineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMediaInsightsPipelineConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMediaPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipelineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMediaPipelineKinesisVideoStreamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetMediaCapturePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipelineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaCapturePipeline` | `MediaCapturePipeline` | no |

## GetMediaInsightsPipelineConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfiguration` | `MediaInsightsPipelineConfiguration` | no |

## GetMediaPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipelineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipeline` | `MediaPipeline` | no |

## GetMediaPipelineKinesisVideoStreamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KinesisVideoStreamPoolConfiguration` | `KinesisVideoStreamPoolConfiguration` | no |

## GetSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpeakerSearchTask` | `SpeakerSearchTask` | no |

## GetVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `VoiceToneAnalysisTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceToneAnalysisTask` | `VoiceToneAnalysisTask` | no |

## ListMediaCapturePipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaCapturePipelines` | `List<MediaCapturePipelineSummary>` | no |
| `NextToken` | `string` | no |

## ListMediaInsightsPipelineConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfigurations` | `List<MediaInsightsPipelineConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListMediaPipelineKinesisVideoStreamPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KinesisVideoStreamPools` | `List<KinesisVideoStreamPoolSummary>` | no |
| `NextToken` | `string` | no |

## ListMediaPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaPipelines` | `List<MediaPipelineSummary>` | no |
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

## StartSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `VoiceProfileDomainArn` | `string` | yes |
| `KinesisVideoStreamSourceTaskConfiguration` | `KinesisVideoStreamSourceTaskConfiguration` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpeakerSearchTask` | `SpeakerSearchTask` | no |

## StartVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `KinesisVideoStreamSourceTaskConfiguration` | `KinesisVideoStreamSourceTaskConfiguration` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceToneAnalysisTask` | `VoiceToneAnalysisTask` | no |

## StopSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `VoiceToneAnalysisTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateMediaInsightsPipelineConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `ResourceAccessRoleArn` | `string` | yes |
| `RealTimeAlertConfiguration` | `RealTimeAlertConfiguration` | no |
| `Elements` | `List<MediaInsightsPipelineConfigurationElement>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MediaInsightsPipelineConfiguration` | `MediaInsightsPipelineConfiguration` | no |

## UpdateMediaInsightsPipelineStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `UpdateStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMediaPipelineKinesisVideoStreamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `StreamConfiguration` | `KinesisVideoStreamConfigurationUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KinesisVideoStreamPoolConfiguration` | `KinesisVideoStreamPoolConfiguration` | no |

