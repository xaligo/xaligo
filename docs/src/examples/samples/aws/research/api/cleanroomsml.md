# AWS Clean Rooms ML

API version: 2023-09-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cleanroomsml/2023-09-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelTrainedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `versionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelTrainedModelInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `trainedModelInferenceJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainingDataStartTime` | `timestamp` | no |
| `trainingDataEndTime` | `timestamp` | no |
| `name` | `string` | yes |
| `trainingDatasetArn` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceModelArn` | `string` | yes |

## CreateConfiguredAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `audienceModelArn` | `string` | yes |
| `outputConfig` | `ConfiguredAudienceModelOutputConfig` | yes |
| `description` | `string` | no |
| `sharedAudienceMetrics` | `List<string>` | yes |
| `minMatchingSeedSize` | `integer` | no |
| `audienceSizeConfig` | `AudienceSizeConfig` | no |
| `tags` | `Map<string>` | no |
| `childResourceTagOnCreatePolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

## CreateConfiguredModelAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `trainingContainerConfig` | `ContainerConfig` | no |
| `inferenceContainerConfig` | `InferenceContainerConfig` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmArn` | `string` | yes |

## CreateConfiguredModelAlgorithmAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredModelAlgorithmArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `privacyConfiguration` | `PrivacyConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmAssociationArn` | `string` | yes |

## CreateMLInputChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `configuredModelAlgorithmAssociations` | `List<string>` | yes |
| `inputChannel` | `InputChannel` | yes |
| `name` | `string` | yes |
| `retentionInDays` | `integer` | yes |
| `description` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `payerConfiguration` | `PayerConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mlInputChannelArn` | `string` | yes |

## CreateTrainedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `hyperparameters` | `Map<string>` | no |
| `environment` | `Map<string>` | no |
| `resourceConfig` | `ResourceConfig` | yes |
| `stoppingCondition` | `StoppingCondition` | no |
| `incrementalTrainingDataChannels` | `List<IncrementalTrainingDataChannel>` | no |
| `dataChannels` | `List<ModelTrainingDataChannel>` | yes |
| `trainingInputMode` | `string` | no |
| `description` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `mlModelTrainingPayerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainedModelArn` | `string` | yes |
| `versionIdentifier` | `string` | no |

## CreateTrainingDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `roleArn` | `string` | yes |
| `trainingData` | `List<Dataset>` | yes |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainingDatasetArn` | `string` | yes |

## DeleteAudienceGenerationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceGenerationJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredAudienceModelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredModelAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguredModelAlgorithmAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMLConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMLInputChannelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mlInputChannelArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrainedModelOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainedModelArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `versionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrainingDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainingDatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAudienceGenerationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceGenerationJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `audienceGenerationJobArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `configuredAudienceModelArn` | `string` | yes |
| `seedAudience` | `AudienceGenerationJobDataSource` | no |
| `includeSeedInOutput` | `boolean` | no |
| `collaborationId` | `string` | no |
| `metrics` | `AudienceQualityMetrics` | no |
| `startedBy` | `string` | no |
| `tags` | `Map<string>` | no |
| `protectedQueryIdentifier` | `string` | no |

## GetAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `trainingDataStartTime` | `timestamp` | no |
| `trainingDataEndTime` | `timestamp` | no |
| `audienceModelArn` | `string` | yes |
| `name` | `string` | yes |
| `trainingDatasetArn` | `string` | yes |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

## GetCollaborationConfiguredModelAlgorithmAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `configuredModelAlgorithmArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `creatorAccountId` | `string` | yes |
| `privacyConfiguration` | `PrivacyConfiguration` | no |

## GetCollaborationMLInputChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mlInputChannelArn` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `mlInputChannelArn` | `string` | yes |
| `name` | `string` | yes |
| `configuredModelAlgorithmAssociations` | `List<string>` | yes |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `retentionInDays` | `integer` | yes |
| `numberOfRecords` | `long` | no |
| `privacyBudgets` | `PrivacyBudgets` | no |
| `description` | `string` | no |
| `syntheticDataConfiguration` | `SyntheticDataConfiguration` | no |
| `payerConfiguration` | `PayerConfiguration` | no |
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `creatorAccountId` | `string` | yes |

## GetCollaborationTrainedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainedModelArn` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `versionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `versionIdentifier` | `string` | no |
| `incrementalTrainingDataChannels` | `List<IncrementalTrainingDataChannelOutput>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `resourceConfig` | `ResourceConfig` | no |
| `trainingInputMode` | `string` | no |
| `stoppingCondition` | `StoppingCondition` | no |
| `metricsStatus` | `string` | no |
| `metricsStatusDetails` | `string` | no |
| `logsStatus` | `string` | no |
| `logsStatusDetails` | `string` | no |
| `trainingContainerImageDigest` | `string` | no |
| `mlModelTrainingPayerAccountId` | `string` | no |
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `creatorAccountId` | `string` | yes |

## GetConfiguredAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `configuredAudienceModelArn` | `string` | yes |
| `name` | `string` | yes |
| `audienceModelArn` | `string` | yes |
| `outputConfig` | `ConfiguredAudienceModelOutputConfig` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `sharedAudienceMetrics` | `List<string>` | yes |
| `minMatchingSeedSize` | `integer` | no |
| `audienceSizeConfig` | `AudienceSizeConfig` | no |
| `tags` | `Map<string>` | no |
| `childResourceTagOnCreatePolicy` | `string` | no |

## GetConfiguredAudienceModelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |
| `configuredAudienceModelPolicy` | `string` | yes |
| `policyHash` | `string` | yes |

## GetConfiguredModelAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `configuredModelAlgorithmArn` | `string` | yes |
| `name` | `string` | yes |
| `trainingContainerConfig` | `ContainerConfig` | no |
| `inferenceContainerConfig` | `InferenceContainerConfig` | no |
| `roleArn` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

## GetConfiguredModelAlgorithmAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `configuredModelAlgorithmArn` | `string` | yes |
| `name` | `string` | yes |
| `privacyConfiguration` | `PrivacyConfiguration` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

## GetMLConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `defaultOutputLocation` | `MLOutputConfiguration` | yes |
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |

## GetMLInputChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mlInputChannelArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `mlInputChannelArn` | `string` | yes |
| `name` | `string` | yes |
| `configuredModelAlgorithmAssociations` | `List<string>` | yes |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `retentionInDays` | `integer` | yes |
| `numberOfRecords` | `long` | no |
| `privacyBudgets` | `PrivacyBudgets` | no |
| `description` | `string` | no |
| `syntheticDataConfiguration` | `SyntheticDataConfiguration` | no |
| `payerConfiguration` | `PayerConfiguration` | no |
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `inputChannel` | `InputChannel` | yes |
| `protectedQueryIdentifier` | `string` | no |
| `numberOfFiles` | `double` | no |
| `sizeInGb` | `double` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

## GetTrainedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainedModelArn` | `string` | yes |
| `membershipIdentifier` | `string` | yes |
| `versionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `collaborationIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `versionIdentifier` | `string` | no |
| `incrementalTrainingDataChannels` | `List<IncrementalTrainingDataChannelOutput>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `statusDetails` | `StatusDetails` | no |
| `configuredModelAlgorithmAssociationArn` | `string` | yes |
| `resourceConfig` | `ResourceConfig` | no |
| `trainingInputMode` | `string` | no |
| `stoppingCondition` | `StoppingCondition` | no |
| `metricsStatus` | `string` | no |
| `metricsStatusDetails` | `string` | no |
| `logsStatus` | `string` | no |
| `logsStatusDetails` | `string` | no |
| `trainingContainerImageDigest` | `string` | no |
| `mlModelTrainingPayerAccountId` | `string` | no |
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `hyperparameters` | `Map<string>` | no |
| `environment` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `dataChannels` | `List<ModelTrainingDataChannel>` | yes |

## GetTrainedModelInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `trainedModelInferenceJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `trainedModelInferenceJobArn` | `string` | yes |
| `configuredModelAlgorithmAssociationArn` | `string` | no |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `trainedModelVersionIdentifier` | `string` | no |
| `resourceConfig` | `InferenceResourceConfig` | yes |
| `outputConfiguration` | `InferenceOutputConfiguration` | yes |
| `membershipIdentifier` | `string` | yes |
| `dataSource` | `ModelInferenceDataSource` | yes |
| `containerExecutionParameters` | `InferenceContainerExecutionParameters` | no |
| `statusDetails` | `StatusDetails` | no |
| `description` | `string` | no |
| `inferenceContainerImageDigest` | `string` | no |
| `environment` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `metricsStatus` | `string` | no |
| `metricsStatusDetails` | `string` | no |
| `logsStatus` | `string` | no |
| `logsStatusDetails` | `string` | no |
| `tags` | `Map<string>` | no |
| `mlModelInferencePayerAccountId` | `string` | no |

## GetTrainingDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainingDatasetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `trainingDatasetArn` | `string` | yes |
| `name` | `string` | yes |
| `trainingData` | `List<Dataset>` | yes |
| `status` | `string` | yes |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

## ListAudienceExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `audienceGenerationJobArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `audienceExportJobs` | `List<AudienceExportJobSummary>` | yes |

## ListAudienceGenerationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `configuredAudienceModelArn` | `string` | no |
| `collaborationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `audienceGenerationJobs` | `List<AudienceGenerationJobSummary>` | yes |

## ListAudienceModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `audienceModels` | `List<AudienceModelSummary>` | yes |

## ListCollaborationConfiguredModelAlgorithmAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationConfiguredModelAlgorithmAssociations` | `List<CollaborationConfiguredModelAlgorithmAssociationSummary>` | yes |

## ListCollaborationMLInputChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationMLInputChannelsList` | `List<CollaborationMLInputChannelSummary>` | yes |

## ListCollaborationTrainedModelExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `collaborationIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `trainedModelVersionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationTrainedModelExportJobs` | `List<CollaborationTrainedModelExportJobSummary>` | yes |

## ListCollaborationTrainedModelInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `collaborationIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | no |
| `trainedModelVersionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationTrainedModelInferenceJobs` | `List<CollaborationTrainedModelInferenceJobSummary>` | yes |

## ListCollaborationTrainedModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `collaborationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `collaborationTrainedModels` | `List<CollaborationTrainedModelSummary>` | yes |

## ListConfiguredAudienceModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `configuredAudienceModels` | `List<ConfiguredAudienceModelSummary>` | yes |

## ListConfiguredModelAlgorithmAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `configuredModelAlgorithmAssociations` | `List<ConfiguredModelAlgorithmAssociationSummary>` | yes |

## ListConfiguredModelAlgorithms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `configuredModelAlgorithms` | `List<ConfiguredModelAlgorithmSummary>` | yes |

## ListMLInputChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `mlInputChannelsList` | `List<MLInputChannelSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## ListTrainedModelInferenceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `membershipIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | no |
| `trainedModelVersionIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `trainedModelInferenceJobs` | `List<TrainedModelInferenceJobSummary>` | yes |

## ListTrainedModelVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `membershipIdentifier` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `trainedModels` | `List<TrainedModelSummary>` | yes |

## ListTrainedModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `membershipIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `trainedModels` | `List<TrainedModelSummary>` | yes |

## ListTrainingDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `trainingDatasets` | `List<TrainingDatasetSummary>` | yes |

## PutConfiguredAudienceModelPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |
| `configuredAudienceModelPolicy` | `string` | yes |
| `previousPolicyHash` | `string` | no |
| `policyExistenceCondition` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelPolicy` | `string` | yes |
| `policyHash` | `string` | yes |

## PutMLConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `defaultOutputLocation` | `MLOutputConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAudienceExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `audienceGenerationJobArn` | `string` | yes |
| `audienceSize` | `AudienceSize` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAudienceGenerationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `configuredAudienceModelArn` | `string` | yes |
| `seedAudience` | `AudienceGenerationJobDataSource` | yes |
| `includeSeedInOutput` | `boolean` | no |
| `collaborationId` | `string` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `audienceGenerationJobArn` | `string` | yes |

## StartTrainedModelExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `trainedModelVersionIdentifier` | `string` | no |
| `membershipIdentifier` | `string` | yes |
| `outputConfiguration` | `TrainedModelExportOutputConfiguration` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartTrainedModelInferenceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `trainedModelArn` | `string` | yes |
| `trainedModelVersionIdentifier` | `string` | no |
| `configuredModelAlgorithmAssociationArn` | `string` | no |
| `resourceConfig` | `InferenceResourceConfig` | yes |
| `outputConfiguration` | `InferenceOutputConfiguration` | yes |
| `dataSource` | `ModelInferenceDataSource` | yes |
| `description` | `string` | no |
| `containerExecutionParameters` | `InferenceContainerExecutionParameters` | no |
| `environment` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `mlModelInferencePayerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trainedModelInferenceJobArn` | `string` | yes |

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


## UpdateConfiguredAudienceModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |
| `outputConfig` | `ConfiguredAudienceModelOutputConfig` | no |
| `audienceModelArn` | `string` | no |
| `sharedAudienceMetrics` | `List<string>` | no |
| `minMatchingSeedSize` | `integer` | no |
| `audienceSizeConfig` | `AudienceSizeConfig` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuredAudienceModelArn` | `string` | yes |

