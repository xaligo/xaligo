# Amazon Interactive Video Service

API version: 2020-07-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ivs/2020-07-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessControlAllowOrigin` | `string` | no |
| `accessControlExposeHeaders` | `string` | no |
| `cacheControl` | `string` | no |
| `contentSecurityPolicy` | `string` | no |
| `strictTransportSecurity` | `string` | no |
| `xContentTypeOptions` | `string` | no |
| `xFrameOptions` | `string` | no |
| `channels` | `List<Channel>` | no |
| `errors` | `List<BatchError>` | no |

## BatchGetStreamKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessControlAllowOrigin` | `string` | no |
| `accessControlExposeHeaders` | `string` | no |
| `cacheControl` | `string` | no |
| `contentSecurityPolicy` | `string` | no |
| `strictTransportSecurity` | `string` | no |
| `xContentTypeOptions` | `string` | no |
| `xFrameOptions` | `string` | no |
| `streamKeys` | `List<StreamKey>` | no |
| `errors` | `List<BatchError>` | no |

## BatchStartViewerSessionRevocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `viewerSessions` | `List<BatchStartViewerSessionRevocationViewerSession>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessControlAllowOrigin` | `string` | no |
| `accessControlExposeHeaders` | `string` | no |
| `cacheControl` | `string` | no |
| `contentSecurityPolicy` | `string` | no |
| `strictTransportSecurity` | `string` | no |
| `xContentTypeOptions` | `string` | no |
| `xFrameOptions` | `string` | no |
| `errors` | `List<BatchStartViewerSessionRevocationError>` | no |

## CreateAdConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `mediaTailorPlaybackConfigurations` | `List<MediaTailorPlaybackConfiguration>` | yes |
| `postRollConfiguration` | `PostRollConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adConfiguration` | `AdConfiguration` | yes |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `latencyMode` | `string` | no |
| `type` | `string` | no |
| `authorized` | `boolean` | no |
| `recordingConfigurationArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `insecureIngest` | `boolean` | no |
| `preset` | `string` | no |
| `playbackRestrictionPolicyArn` | `string` | no |
| `multitrackInputConfiguration` | `MultitrackInputConfiguration` | no |
| `containerFormat` | `string` | no |
| `adConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channel` | `Channel` | no |
| `streamKey` | `StreamKey` | no |

## CreatePlaybackRestrictionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `allowedCountries` | `List<string>` | no |
| `allowedOrigins` | `List<string>` | no |
| `enableStrictOriginEnforcement` | `boolean` | no |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `playbackRestrictionPolicy` | `PlaybackRestrictionPolicy` | no |

## CreateRecordingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `destinationConfiguration` | `DestinationConfiguration` | yes |
| `tags` | `Map<string>` | no |
| `thumbnailConfiguration` | `ThumbnailConfiguration` | no |
| `recordingReconnectWindowSeconds` | `integer` | no |
| `renditionConfiguration` | `RenditionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recordingConfiguration` | `RecordingConfiguration` | no |

## CreateStreamKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamKey` | `StreamKey` | no |

## DeleteAdConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlaybackKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlaybackRestrictionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecordingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStreamKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAdConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adConfiguration` | `AdConfiguration` | no |

## GetChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channel` | `Channel` | no |

## GetPlaybackKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPair` | `PlaybackKeyPair` | no |

## GetPlaybackRestrictionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `playbackRestrictionPolicy` | `PlaybackRestrictionPolicy` | no |

## GetRecordingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recordingConfiguration` | `RecordingConfiguration` | no |

## GetStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stream` | `Stream` | no |

## GetStreamKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamKey` | `StreamKey` | no |

## GetStreamSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `streamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamSession` | `StreamSession` | no |

## ImportPlaybackKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKeyMaterial` | `string` | yes |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPair` | `PlaybackKeyPair` | no |

## InsertAdBreak

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `durationSeconds` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adBreakId` | `string` | no |

## ListAdConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adConfigurations` | `List<AdConfigurationSummary>` | yes |
| `nextToken` | `string` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterByName` | `string` | no |
| `filterByRecordingConfigurationArn` | `string` | no |
| `filterByPlaybackRestrictionPolicyArn` | `string` | no |
| `filterByAdConfigurationArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channels` | `List<ChannelSummary>` | yes |
| `nextToken` | `string` | no |

## ListPlaybackKeyPairs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyPairs` | `List<PlaybackKeyPairSummary>` | yes |
| `nextToken` | `string` | no |

## ListPlaybackRestrictionPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `playbackRestrictionPolicies` | `List<PlaybackRestrictionPolicySummary>` | yes |
| `nextToken` | `string` | no |

## ListRecordingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recordingConfigurations` | `List<RecordingConfigurationSummary>` | yes |
| `nextToken` | `string` | no |

## ListStreamKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamKeys` | `List<StreamKeySummary>` | yes |
| `nextToken` | `string` | no |

## ListStreamSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamSessions` | `List<StreamSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterBy` | `StreamFilters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streams` | `List<StreamSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## PutMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `metadata` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartViewerSessionRevocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `viewerId` | `string` | yes |
| `viewerSessionVersionsLessThanOrEqualTo` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateAdConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `mediaTailorPlaybackConfigurations` | `List<MediaTailorPlaybackConfiguration>` | no |
| `postRollConfiguration` | `PostRollConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `adConfiguration` | `AdConfiguration` | yes |

## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `latencyMode` | `string` | no |
| `type` | `string` | no |
| `authorized` | `boolean` | no |
| `recordingConfigurationArn` | `string` | no |
| `insecureIngest` | `boolean` | no |
| `preset` | `string` | no |
| `playbackRestrictionPolicyArn` | `string` | no |
| `multitrackInputConfiguration` | `MultitrackInputConfiguration` | no |
| `containerFormat` | `string` | no |
| `adConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channel` | `Channel` | no |

## UpdatePlaybackRestrictionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `allowedCountries` | `List<string>` | no |
| `allowedOrigins` | `List<string>` | no |
| `enableStrictOriginEnforcement` | `boolean` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `playbackRestrictionPolicy` | `PlaybackRestrictionPolicy` | no |

