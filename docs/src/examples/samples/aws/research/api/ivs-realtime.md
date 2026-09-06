# Amazon Interactive Video Service RealTime

API version: 2020-07-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ivs-realtime/2020-07-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateEncoderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `video` | `Video` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encoderConfiguration` | `EncoderConfiguration` | no |

## CreateIngestConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `stageArn` | `string` | no |
| `userId` | `string` | no |
| `attributes` | `Map<string>` | no |
| `ingestProtocol` | `string` | yes |
| `insecureIngest` | `boolean` | no |
| `redundantIngest` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestConfiguration` | `IngestConfiguration` | no |

## CreateParticipantToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `duration` | `integer` | no |
| `userId` | `string` | no |
| `attributes` | `Map<string>` | no |
| `capabilities` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `participantToken` | `ParticipantToken` | no |

## CreateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `participantTokenConfigurations` | `List<ParticipantTokenConfiguration>` | no |
| `tags` | `Map<string>` | no |
| `autoParticipantRecordingConfiguration` | `AutoParticipantRecordingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stage` | `Stage` | no |
| `participantTokens` | `List<ParticipantToken>` | no |

## CreateStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `s3` | `S3StorageConfiguration` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageConfiguration` | `StorageConfiguration` | no |

## DeleteEncoderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIngestConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisconnectParticipant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `participantId` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetComposition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `composition` | `Composition` | no |

## GetEncoderConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encoderConfiguration` | `EncoderConfiguration` | no |

## GetIngestConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestConfiguration` | `IngestConfiguration` | no |

## GetParticipant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `participantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `participant` | `Participant` | no |

## GetPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKey` | `PublicKey` | no |

## GetStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stage` | `Stage` | no |

## GetStageSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageSession` | `StageSession` | no |

## GetStorageConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageConfiguration` | `StorageConfiguration` | no |

## ImportPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKeyMaterial` | `string` | yes |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKey` | `PublicKey` | no |

## ListCompositions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterByStageArn` | `string` | no |
| `filterByEncoderConfigurationArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `compositions` | `List<CompositionSummary>` | yes |
| `nextToken` | `string` | no |

## ListEncoderConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encoderConfigurations` | `List<EncoderConfigurationSummary>` | yes |
| `nextToken` | `string` | no |

## ListIngestConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterByStageArn` | `string` | no |
| `filterByState` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestConfigurations` | `List<IngestConfigurationSummary>` | yes |
| `nextToken` | `string` | no |

## ListParticipantEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `participantId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<Event>` | yes |
| `nextToken` | `string` | no |

## ListParticipantReplicas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceStageArn` | `string` | yes |
| `participantId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicas` | `List<ParticipantReplica>` | yes |
| `nextToken` | `string` | no |

## ListParticipants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `filterByUserId` | `string` | no |
| `filterByPublished` | `boolean` | no |
| `filterByState` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filterByRecordingState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `participants` | `List<ParticipantSummary>` | yes |
| `nextToken` | `string` | no |

## ListPublicKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `publicKeys` | `List<PublicKeySummary>` | yes |
| `nextToken` | `string` | no |

## ListStageSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageSessions` | `List<StageSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListStages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stages` | `List<StageSummary>` | yes |
| `nextToken` | `string` | no |

## ListStorageConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageConfigurations` | `List<StorageConfigurationSummary>` | yes |
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

## StartComposition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stageArn` | `string` | yes |
| `idempotencyToken` | `string` | no |
| `layout` | `LayoutConfiguration` | no |
| `destinations` | `List<DestinationConfiguration>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `composition` | `Composition` | no |

## StartParticipantReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceStageArn` | `string` | yes |
| `destinationStageArn` | `string` | yes |
| `participantId` | `string` | yes |
| `reconnectWindowSeconds` | `integer` | no |
| `attributes` | `Map<string>` | no |

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

## StopComposition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopParticipantReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceStageArn` | `string` | yes |
| `destinationStageArn` | `string` | yes |
| `participantId` | `string` | yes |

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


## UpdateIngestConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `stageArn` | `string` | no |
| `redundantIngest` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestConfiguration` | `IngestConfiguration` | no |

## UpdateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `autoParticipantRecordingConfiguration` | `AutoParticipantRecordingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stage` | `Stage` | no |

