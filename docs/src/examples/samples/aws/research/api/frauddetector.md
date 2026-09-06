# Amazon Fraud Detector

API version: 2019-11-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/frauddetector/2019-11-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateVariable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `variableEntries` | `List<VariableEntry>` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchCreateVariableError>` | no |

## BatchGetVariable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `variables` | `List<Variable>` | no |
| `errors` | `List<BatchGetVariableError>` | no |

## CancelBatchImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelBatchPredictionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBatchImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `inputPath` | `string` | yes |
| `outputPath` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `iamRoleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBatchPredictionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `inputPath` | `string` | yes |
| `outputPath` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `detectorName` | `string` | yes |
| `detectorVersion` | `string` | no |
| `iamRoleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDetectorVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `description` | `string` | no |
| `externalModelEndpoints` | `List<string>` | no |
| `rules` | `List<Rule>` | yes |
| `modelVersions` | `List<ModelVersion>` | no |
| `ruleExecutionMode` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | no |
| `detectorVersionId` | `string` | no |
| `status` | `string` | no |

## CreateList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `elements` | `List<string>` | no |
| `variableType` | `string` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `description` | `string` | no |
| `eventTypeName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `trainingDataSource` | `string` | yes |
| `trainingDataSchema` | `TrainingDataSchema` | yes |
| `externalEventsDetail` | `ExternalEventsDetail` | no |
| `ingestedEventsDetail` | `IngestedEventsDetail` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | no |
| `modelType` | `string` | no |
| `modelVersionNumber` | `string` | no |
| `status` | `string` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `detectorId` | `string` | yes |
| `description` | `string` | no |
| `expression` | `string` | yes |
| `language` | `string` | yes |
| `outcomes` | `List<string>` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rule` | `Rule` | no |

## CreateVariable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `dataType` | `string` | yes |
| `dataSource` | `string` | yes |
| `defaultValue` | `string` | yes |
| `description` | `string` | no |
| `variableType` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBatchImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBatchPredictionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDetectorVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEntityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `deleteAuditHistory` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventsByEventType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypeName` | `string` | no |
| `eventsDeletionStatus` | `string` | no |

## DeleteExternalModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelEndpoint` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `modelVersionNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOutcome

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rule` | `Rule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVariable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | no |
| `detectorVersionSummaries` | `List<DetectorVersionSummary>` | no |
| `nextToken` | `string` | no |
| `arn` | `string` | no |

## DescribeModelVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | no |
| `modelVersionNumber` | `string` | no |
| `modelType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelVersionDetails` | `List<ModelVersionDetail>` | no |
| `nextToken` | `string` | no |

## GetBatchImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchImports` | `List<BatchImport>` | no |
| `nextToken` | `string` | no |

## GetBatchPredictionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchPredictions` | `List<BatchPrediction>` | no |
| `nextToken` | `string` | no |

## GetDeleteEventsByEventTypeStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypeName` | `string` | no |
| `eventsDeletionStatus` | `string` | no |

## GetDetectorVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | no |
| `detectorVersionId` | `string` | no |
| `description` | `string` | no |
| `externalModelEndpoints` | `List<string>` | no |
| `modelVersions` | `List<ModelVersion>` | no |
| `rules` | `List<Rule>` | no |
| `status` | `string` | no |
| `lastUpdatedTime` | `string` | no |
| `createdTime` | `string` | no |
| `ruleExecutionMode` | `string` | no |
| `arn` | `string` | no |

## GetDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectors` | `List<Detector>` | no |
| `nextToken` | `string` | no |

## GetEntityTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entityTypes` | `List<EntityType>` | no |
| `nextToken` | `string` | no |

## GetEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `event` | `Event` | no |

## GetEventPrediction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | no |
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `entities` | `List<Entity>` | yes |
| `eventTimestamp` | `string` | yes |
| `eventVariables` | `Map<string>` | yes |
| `externalModelEndpointDataBlobs` | `Map<ModelEndpointDataBlob>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelScores` | `List<ModelScores>` | no |
| `ruleResults` | `List<RuleResult>` | no |
| `externalModelOutputs` | `List<ExternalModelOutputs>` | no |

## GetEventPredictionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |
| `predictionTimestamp` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | no |
| `eventTypeName` | `string` | no |
| `entityId` | `string` | no |
| `entityType` | `string` | no |
| `eventTimestamp` | `string` | no |
| `detectorId` | `string` | no |
| `detectorVersionId` | `string` | no |
| `detectorVersionStatus` | `string` | no |
| `eventVariables` | `List<EventVariableSummary>` | no |
| `rules` | `List<EvaluatedRule>` | no |
| `ruleExecutionMode` | `string` | no |
| `outcomes` | `List<string>` | no |
| `evaluatedModelVersions` | `List<EvaluatedModelVersion>` | no |
| `evaluatedExternalModels` | `List<EvaluatedExternalModel>` | no |
| `predictionTimestamp` | `string` | no |

## GetEventTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventTypes` | `List<EventType>` | no |
| `nextToken` | `string` | no |

## GetExternalModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelEndpoint` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `externalModels` | `List<ExternalModel>` | no |
| `nextToken` | `string` | no |

## GetKMSEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKey` | `KMSKey` | no |

## GetLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `labels` | `List<Label>` | no |
| `nextToken` | `string` | no |

## GetListElements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `elements` | `List<string>` | no |
| `nextToken` | `string` | no |

## GetListsMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lists` | `List<AllowDenyList>` | no |
| `nextToken` | `string` | no |

## GetModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `modelVersionNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | no |
| `modelType` | `string` | no |
| `modelVersionNumber` | `string` | no |
| `trainingDataSource` | `string` | no |
| `trainingDataSchema` | `TrainingDataSchema` | no |
| `externalEventsDetail` | `ExternalEventsDetail` | no |
| `ingestedEventsDetail` | `IngestedEventsDetail` | no |
| `status` | `string` | no |
| `arn` | `string` | no |

## GetModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | no |
| `modelType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `models` | `List<Model>` | no |

## GetOutcomes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `outcomes` | `List<Outcome>` | no |
| `nextToken` | `string` | no |

## GetRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | no |
| `detectorId` | `string` | yes |
| `ruleVersion` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleDetails` | `List<RuleDetail>` | no |
| `nextToken` | `string` | no |

## GetVariables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `variables` | `List<Variable>` | no |
| `nextToken` | `string` | no |

## ListEventPredictions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `FilterCondition` | no |
| `eventType` | `FilterCondition` | no |
| `detectorId` | `FilterCondition` | no |
| `detectorVersionId` | `FilterCondition` | no |
| `predictionTimeRange` | `PredictionTimeRange` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventPredictionSummaries` | `List<EventPredictionSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `nextToken` | `string` | no |

## PutDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `description` | `string` | no |
| `eventTypeName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEntityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEventType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `eventVariables` | `List<string>` | yes |
| `labels` | `List<string>` | no |
| `entityTypes` | `List<string>` | yes |
| `eventIngestion` | `string` | no |
| `tags` | `List<Tag>` | no |
| `eventOrchestration` | `EventOrchestration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutExternalModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelEndpoint` | `string` | yes |
| `modelSource` | `string` | yes |
| `invokeModelEndpointRoleArn` | `string` | yes |
| `inputConfiguration` | `ModelInputConfiguration` | yes |
| `outputConfiguration` | `ModelOutputConfiguration` | yes |
| `modelEndpointStatus` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutKMSEncryptionKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsEncryptionKeyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutOutcome

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `eventTimestamp` | `string` | yes |
| `eventVariables` | `Map<string>` | yes |
| `assignedLabel` | `string` | no |
| `labelTimestamp` | `string` | no |
| `entities` | `List<Entity>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDetectorVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |
| `externalModelEndpoints` | `List<string>` | yes |
| `rules` | `List<Rule>` | yes |
| `description` | `string` | no |
| `modelVersions` | `List<ModelVersion>` | no |
| `ruleExecutionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDetectorVersionMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |
| `description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDetectorVersionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorId` | `string` | yes |
| `detectorVersionId` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `eventTypeName` | `string` | yes |
| `assignedLabel` | `string` | yes |
| `labelTimestamp` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `elements` | `List<string>` | no |
| `description` | `string` | no |
| `updateMode` | `string` | no |
| `variableType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateModelVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `majorVersionNumber` | `string` | yes |
| `externalEventsDetail` | `ExternalEventsDetail` | no |
| `ingestedEventsDetail` | `IngestedEventsDetail` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | no |
| `modelType` | `string` | no |
| `modelVersionNumber` | `string` | no |
| `status` | `string` | no |

## UpdateModelVersionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `modelType` | `string` | yes |
| `modelVersionNumber` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRuleMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rule` | `Rule` | yes |
| `description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRuleVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rule` | `Rule` | yes |
| `description` | `string` | no |
| `expression` | `string` | yes |
| `language` | `string` | yes |
| `outcomes` | `List<string>` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rule` | `Rule` | no |

## UpdateVariable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `defaultValue` | `string` | no |
| `description` | `string` | no |
| `variableType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


