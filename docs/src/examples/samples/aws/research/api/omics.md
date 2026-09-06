# Amazon Omics

API version: 2022-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/omics/2022-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AbortMultipartReadSetUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `uploadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AcceptShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shareId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## BatchDeleteReadSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |
| `sequenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<ReadSetBatchError>` | no |

## CancelAnnotationImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelRunBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelVariantImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CompleteMultipartReadSetUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `uploadId` | `string` | yes |
| `parts` | `List<CompleteReadSetUploadPartListItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `readSetId` | `string` | yes |

## CreateAnnotationStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reference` | `ReferenceItem` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `versionName` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `storeFormat` | `string` | yes |
| `storeOptions` | `StoreOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | no |
| `storeFormat` | `string` | no |
| `storeOptions` | `StoreOptions` | no |
| `status` | `string` | yes |
| `name` | `string` | yes |
| `versionName` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## CreateAnnotationStoreVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | no |
| `versionOptions` | `VersionOptions` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `versionName` | `string` | yes |
| `storeId` | `string` | yes |
| `versionOptions` | `VersionOptions` | no |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## CreateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `runConfigurations` | `RunConfigurations` | yes |
| `tags` | `Map<string>` | no |
| `requestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `uuid` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `runConfigurations` | `RunConfigurationsResponse` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## CreateMultipartReadSetUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `clientToken` | `string` | no |
| `sourceFileType` | `string` | yes |
| `subjectId` | `string` | yes |
| `sampleId` | `string` | yes |
| `generatedFrom` | `string` | no |
| `referenceArn` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `uploadId` | `string` | yes |
| `sourceFileType` | `string` | yes |
| `subjectId` | `string` | yes |
| `sampleId` | `string` | yes |
| `generatedFrom` | `string` | no |
| `referenceArn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `creationTime` | `timestamp` | yes |

## CreateReferenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `creationTime` | `timestamp` | yes |

## CreateRunCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cacheBehavior` | `string` | no |
| `cacheS3Location` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | no |
| `requestId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `cacheBucketOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateRunGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `maxCpus` | `integer` | no |
| `maxRuns` | `integer` | no |
| `maxDuration` | `integer` | no |
| `tags` | `Map<string>` | no |
| `requestId` | `string` | yes |
| `maxGpus` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateSequenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |
| `fallbackLocation` | `string` | no |
| `eTagAlgorithmFamily` | `string` | no |
| `propagatedSetLevelTags` | `List<string>` | no |
| `s3AccessConfig` | `S3AccessConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `creationTime` | `timestamp` | yes |
| `fallbackLocation` | `string` | no |
| `eTagAlgorithmFamily` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `propagatedSetLevelTags` | `List<string>` | no |
| `s3Access` | `SequenceStoreS3Access` | no |

## CreateShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `principalSubscriber` | `string` | yes |
| `shareName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shareId` | `string` | no |
| `status` | `string` | no |
| `shareName` | `string` | no |

## CreateVariantStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reference` | `ReferenceItem` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `sseConfig` | `SseConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | no |
| `status` | `string` | yes |
| `name` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `engine` | `string` | no |
| `definitionZip` | `blob` | no |
| `definitionUri` | `string` | no |
| `main` | `string` | no |
| `parameterTemplate` | `Map<WorkflowParameter>` | no |
| `storageCapacity` | `integer` | no |
| `tags` | `Map<string>` | no |
| `requestId` | `string` | yes |
| `accelerators` | `string` | no |
| `storageType` | `string` | no |
| `containerRegistryMap` | `ContainerRegistryMap` | no |
| `containerRegistryMapUri` | `string` | no |
| `readmeMarkdown` | `string` | no |
| `parameterTemplatePath` | `string` | no |
| `readmePath` | `string` | no |
| `definitionRepository` | `DefinitionRepository` | no |
| `workflowBucketOwnerId` | `string` | no |
| `readmeUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |
| `uuid` | `string` | no |

## CreateWorkflowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `versionName` | `string` | yes |
| `definitionZip` | `blob` | no |
| `definitionUri` | `string` | no |
| `accelerators` | `string` | no |
| `description` | `string` | no |
| `engine` | `string` | no |
| `main` | `string` | no |
| `parameterTemplate` | `Map<WorkflowParameter>` | no |
| `requestId` | `string` | yes |
| `storageType` | `string` | no |
| `storageCapacity` | `integer` | no |
| `tags` | `Map<string>` | no |
| `workflowBucketOwnerId` | `string` | no |
| `containerRegistryMap` | `ContainerRegistryMap` | no |
| `containerRegistryMapUri` | `string` | no |
| `readmeMarkdown` | `string` | no |
| `parameterTemplatePath` | `string` | no |
| `readmePath` | `string` | no |
| `definitionRepository` | `DefinitionRepository` | no |
| `readmeUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `workflowId` | `string` | no |
| `versionName` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |
| `uuid` | `string` | no |

## DeleteAnnotationStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteAnnotationStoreVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `versions` | `List<string>` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<VersionDeleteError>` | no |

## DeleteBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReferenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRunBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRunCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRunGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteS3AccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3AccessPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSequenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shareId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## DeleteVariantStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `versionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAnnotationImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `destinationName` | `string` | yes |
| `versionName` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | yes |
| `items` | `List<AnnotationImportItemDetail>` | yes |
| `runLeftNormalization` | `boolean` | yes |
| `formatOptions` | `FormatOptions` | yes |
| `annotationFields` | `Map<string>` | no |

## GetAnnotationStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | yes |
| `status` | `string` | yes |
| `storeArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `sseConfig` | `SseConfig` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `tags` | `Map<string>` | yes |
| `storeOptions` | `StoreOptions` | no |
| `storeFormat` | `string` | no |
| `statusMessage` | `string` | yes |
| `storeSizeBytes` | `long` | yes |
| `numVersions` | `integer` | yes |

## GetAnnotationStoreVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `versionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storeId` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |
| `versionArn` | `string` | yes |
| `name` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `tags` | `Map<string>` | yes |
| `versionOptions` | `VersionOptions` | no |
| `statusMessage` | `string` | yes |
| `versionSizeBytes` | `long` | yes |

## GetBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `uuid` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |
| `totalRuns` | `integer` | no |
| `defaultRunSetting` | `DefaultRunSetting` | no |
| `submissionSummary` | `SubmissionSummary` | no |
| `runSummary` | `RunSummary` | no |
| `creationTime` | `timestamp` | no |
| `submittedTime` | `timestamp` | no |
| `processedTime` | `timestamp` | no |
| `failedTime` | `timestamp` | no |
| `failureReason` | `string` | no |

## GetConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `uuid` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `runConfigurations` | `RunConfigurationsResponse` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetReadSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `file` | `string` | no |
| `partNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | no |

## GetReadSetActivationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | no |
| `sources` | `List<ActivateReadSetSourceItem>` | no |

## GetReadSetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `destination` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | no |
| `readSets` | `List<ExportReadSetDetail>` | no |

## GetReadSetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | no |
| `sources` | `List<ImportReadSetSourceItem>` | yes |

## GetReadSetMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `subjectId` | `string` | no |
| `sampleId` | `string` | no |
| `status` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `fileType` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `sequenceInformation` | `SequenceInformation` | no |
| `referenceArn` | `string` | no |
| `files` | `ReadSetFiles` | no |
| `statusMessage` | `string` | no |
| `creationType` | `string` | no |
| `etag` | `ETag` | no |
| `creationJobId` | `string` | no |

## GetReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |
| `range` | `string` | no |
| `partNumber` | `integer` | yes |
| `file` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | no |

## GetReferenceImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | no |
| `sources` | `List<ImportReferenceSourceItem>` | yes |

## GetReferenceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `referenceStoreId` | `string` | yes |
| `md5` | `string` | yes |
| `status` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `files` | `ReferenceFiles` | no |
| `creationType` | `string` | no |
| `creationJobId` | `string` | no |

## GetReferenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `creationTime` | `timestamp` | yes |

## GetRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `export` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `cacheId` | `string` | no |
| `cacheBehavior` | `string` | no |
| `engineVersion` | `string` | no |
| `status` | `string` | no |
| `workflowId` | `string` | no |
| `workflowType` | `string` | no |
| `runId` | `string` | no |
| `roleArn` | `string` | no |
| `name` | `string` | no |
| `runGroupId` | `string` | no |
| `batchId` | `string` | no |
| `priority` | `integer` | no |
| `definition` | `string` | no |
| `digest` | `string` | no |
| `parameters` | `RunParameters` | no |
| `storageCapacity` | `integer` | no |
| `outputUri` | `string` | no |
| `logLevel` | `string` | no |
| `resourceDigests` | `Map<string>` | no |
| `startedBy` | `string` | no |
| `creationTime` | `timestamp` | no |
| `startTime` | `timestamp` | no |
| `stopTime` | `timestamp` | no |
| `statusMessage` | `string` | no |
| `tags` | `Map<string>` | no |
| `accelerators` | `string` | no |
| `retentionMode` | `string` | no |
| `failureReason` | `string` | no |
| `logLocation` | `RunLogLocation` | no |
| `uuid` | `string` | no |
| `runOutputUri` | `string` | no |
| `storageType` | `string` | no |
| `workflowOwnerId` | `string` | no |
| `workflowVersionName` | `string` | no |
| `workflowUuid` | `string` | no |
| `networkingMode` | `string` | no |
| `scratchStorageMode` | `string` | no |
| `configuration` | `ConfigurationDetails` | no |
| `vpcConfig` | `VpcConfigResponse` | no |
| `engineSettings` | `EngineSettings` | no |

## GetRunCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `cacheBehavior` | `string` | no |
| `cacheBucketOwnerId` | `string` | no |
| `cacheS3Uri` | `string` | no |
| `creationTime` | `timestamp` | no |
| `description` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |

## GetRunGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `maxCpus` | `integer` | no |
| `maxRuns` | `integer` | no |
| `maxDuration` | `integer` | no |
| `creationTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `maxGpus` | `integer` | no |

## GetRunTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `cpus` | `integer` | no |
| `cacheHit` | `boolean` | no |
| `cacheS3Uri` | `string` | no |
| `memory` | `integer` | no |
| `creationTime` | `timestamp` | no |
| `startTime` | `timestamp` | no |
| `stopTime` | `timestamp` | no |
| `statusMessage` | `string` | no |
| `logStream` | `string` | no |
| `gpus` | `integer` | no |
| `instanceType` | `string` | no |
| `failureReason` | `string` | no |
| `imageDetails` | `ImageDetails` | no |
| `uuid` | `string` | no |

## GetS3AccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3AccessPointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3AccessPointArn` | `string` | no |
| `storeId` | `string` | no |
| `storeType` | `string` | no |
| `updateTime` | `timestamp` | no |
| `s3AccessPolicy` | `string` | yes |

## GetSequenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `creationTime` | `timestamp` | yes |
| `fallbackLocation` | `string` | no |
| `s3Access` | `SequenceStoreS3Access` | no |
| `eTagAlgorithmFamily` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `propagatedSetLevelTags` | `List<string>` | no |
| `updateTime` | `timestamp` | no |

## GetShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shareId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `share` | `ShareDetails` | no |

## GetVariantImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `destinationName` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `completionTime` | `timestamp` | no |
| `items` | `List<VariantImportItemDetail>` | yes |
| `runLeftNormalization` | `boolean` | yes |
| `annotationFields` | `Map<string>` | no |

## GetVariantStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | yes |
| `status` | `string` | yes |
| `storeArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `sseConfig` | `SseConfig` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `tags` | `Map<string>` | yes |
| `statusMessage` | `string` | yes |
| `storeSizeBytes` | `long` | yes |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `type` | `string` | no |
| `export` | `List<string>` | no |
| `workflowOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `engine` | `string` | no |
| `definition` | `string` | no |
| `main` | `string` | no |
| `digest` | `string` | no |
| `parameterTemplate` | `Map<WorkflowParameter>` | no |
| `storageCapacity` | `integer` | no |
| `creationTime` | `timestamp` | no |
| `statusMessage` | `string` | no |
| `tags` | `Map<string>` | no |
| `metadata` | `Map<string>` | no |
| `accelerators` | `string` | no |
| `storageType` | `string` | no |
| `uuid` | `string` | no |
| `containerRegistryMap` | `ContainerRegistryMap` | no |
| `readme` | `string` | no |
| `definitionRepositoryDetails` | `DefinitionRepositoryDetails` | no |
| `readmePath` | `string` | no |
| `profiles` | `List<string>` | no |
| `profileParameterTemplates` | `Map<Map<WorkflowParameter>>` | no |

## GetWorkflowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `versionName` | `string` | yes |
| `type` | `string` | no |
| `export` | `List<string>` | no |
| `workflowOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `workflowId` | `string` | no |
| `versionName` | `string` | no |
| `accelerators` | `string` | no |
| `creationTime` | `timestamp` | no |
| `description` | `string` | no |
| `definition` | `string` | no |
| `digest` | `string` | no |
| `engine` | `string` | no |
| `main` | `string` | no |
| `metadata` | `Map<string>` | no |
| `parameterTemplate` | `Map<WorkflowParameter>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `storageType` | `string` | no |
| `storageCapacity` | `integer` | no |
| `type` | `string` | no |
| `tags` | `Map<string>` | no |
| `uuid` | `string` | no |
| `workflowBucketOwnerId` | `string` | no |
| `containerRegistryMap` | `ContainerRegistryMap` | no |
| `readme` | `string` | no |
| `definitionRepositoryDetails` | `DefinitionRepositoryDetails` | no |
| `readmePath` | `string` | no |
| `profiles` | `List<string>` | no |
| `profileParameterTemplates` | `Map<Map<WorkflowParameter>>` | no |

## ListAnnotationImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |
| `filter` | `ListAnnotationImportJobsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `annotationImportJobs` | `List<AnnotationImportJobItem>` | no |
| `nextToken` | `string` | no |

## ListAnnotationStoreVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ListAnnotationStoreVersionsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `annotationStoreVersions` | `List<AnnotationStoreVersionItem>` | no |
| `nextToken` | `string` | no |

## ListAnnotationStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ListAnnotationStoresFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `annotationStores` | `List<AnnotationStoreItem>` | no |
| `nextToken` | `string` | no |

## ListBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxItems` | `integer` | no |
| `startingToken` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `runGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BatchListItem>` | no |
| `nextToken` | `string` | no |

## ListConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `startingToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ConfigurationListItem>` | no |
| `nextToken` | `string` | no |

## ListMultipartReadSetUploads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `uploads` | `List<MultipartReadSetUploadListItem>` | no |

## ListReadSetActivationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ActivateReadSetFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `activationJobs` | `List<ActivateReadSetJobItem>` | no |

## ListReadSetExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ExportReadSetFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `exportJobs` | `List<ExportReadSetJobDetail>` | no |

## ListReadSetImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sequenceStoreId` | `string` | yes |
| `filter` | `ImportReadSetFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `importJobs` | `List<ImportReadSetJobItem>` | no |

## ListReadSetUploadParts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `uploadId` | `string` | yes |
| `partSource` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ReadSetUploadPartListFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `parts` | `List<ReadSetUploadPartListItem>` | no |

## ListReadSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ReadSetFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `readSets` | `List<ReadSetListItem>` | yes |

## ListReferenceImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `referenceStoreId` | `string` | yes |
| `filter` | `ImportReferenceFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `importJobs` | `List<ImportReferenceJobItem>` | no |

## ListReferenceStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ReferenceStoreFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `referenceStores` | `List<ReferenceStoreDetail>` | yes |

## ListReferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `referenceStoreId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `ReferenceFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `references` | `List<ReferenceListItem>` | yes |

## ListRunCaches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `startingToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RunCacheListItem>` | no |
| `nextToken` | `string` | no |

## ListRunGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `startingToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RunGroupListItem>` | no |
| `nextToken` | `string` | no |

## ListRunTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | no |
| `startingToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TaskListItem>` | no |
| `nextToken` | `string` | no |

## ListRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `runGroupId` | `string` | no |
| `batchId` | `string` | no |
| `startingToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RunListItem>` | no |
| `nextToken` | `string` | no |

## ListRunsInBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchId` | `string` | yes |
| `maxItems` | `integer` | no |
| `startingToken` | `string` | no |
| `submissionStatus` | `string` | no |
| `runSettingId` | `string` | no |
| `runId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runs` | `List<RunBatchListItem>` | no |
| `nextToken` | `string` | no |

## ListSequenceStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `SequenceStoreFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sequenceStores` | `List<SequenceStoreDetail>` | yes |

## ListShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceOwner` | `string` | yes |
| `filter` | `Filter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shares` | `List<ShareDetails>` | yes |
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

## ListVariantImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |
| `filter` | `ListVariantImportJobsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `variantImportJobs` | `List<VariantImportJobItem>` | no |
| `nextToken` | `string` | no |

## ListVariantStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |
| `filter` | `ListVariantStoresFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `variantStores` | `List<VariantStoreItem>` | no |
| `nextToken` | `string` | no |

## ListWorkflowVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `type` | `string` | no |
| `workflowOwnerId` | `string` | no |
| `startingToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<WorkflowVersionListItem>` | no |
| `nextToken` | `string` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | no |
| `name` | `string` | no |
| `startingToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<WorkflowListItem>` | no |
| `nextToken` | `string` | no |

## PutS3AccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3AccessPointArn` | `string` | yes |
| `s3AccessPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3AccessPointArn` | `string` | no |
| `storeId` | `string` | no |
| `storeType` | `string` | no |

## StartAnnotationImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationName` | `string` | yes |
| `roleArn` | `string` | yes |
| `items` | `List<AnnotationImportItemSource>` | yes |
| `versionName` | `string` | no |
| `formatOptions` | `FormatOptions` | no |
| `runLeftNormalization` | `boolean` | no |
| `annotationFields` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

## StartReadSetActivationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `clientToken` | `string` | no |
| `sources` | `List<StartReadSetActivationJobSourceItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## StartReadSetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `destination` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientToken` | `string` | no |
| `sources` | `List<ExportReadSet>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `destination` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## StartReadSetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientToken` | `string` | no |
| `sources` | `List<StartReadSetImportJobSourceItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `sequenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## StartReferenceImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `referenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientToken` | `string` | no |
| `sources` | `List<StartReferenceImportJobSourceItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `referenceStoreId` | `string` | yes |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |

## StartRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | no |
| `workflowType` | `string` | no |
| `runId` | `string` | no |
| `roleArn` | `string` | yes |
| `name` | `string` | no |
| `cacheId` | `string` | no |
| `cacheBehavior` | `string` | no |
| `runGroupId` | `string` | no |
| `priority` | `integer` | no |
| `parameters` | `RunParameters` | no |
| `storageCapacity` | `integer` | no |
| `outputUri` | `string` | yes |
| `logLevel` | `string` | no |
| `tags` | `Map<string>` | no |
| `requestId` | `string` | yes |
| `retentionMode` | `string` | no |
| `storageType` | `string` | no |
| `workflowOwnerId` | `string` | no |
| `workflowVersionName` | `string` | no |
| `networkingMode` | `string` | no |
| `scratchStorageMode` | `string` | no |
| `configurationName` | `string` | no |
| `engineSettings` | `EngineSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |
| `uuid` | `string` | no |
| `runOutputUri` | `string` | no |
| `configuration` | `ConfigurationDetails` | no |
| `networkingMode` | `string` | no |

## StartRunBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchName` | `string` | no |
| `requestId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `defaultRunSetting` | `DefaultRunSetting` | yes |
| `batchRunSettings` | `BatchRunSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `uuid` | `string` | no |
| `tags` | `Map<string>` | no |

## StartVariantImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationName` | `string` | yes |
| `roleArn` | `string` | yes |
| `items` | `List<VariantImportItemSource>` | yes |
| `runLeftNormalization` | `boolean` | no |
| `annotationFields` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

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


## UpdateAnnotationStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | yes |
| `status` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |
| `storeOptions` | `StoreOptions` | no |
| `storeFormat` | `string` | no |

## UpdateAnnotationStoreVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storeId` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |
| `name` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |

## UpdateRunCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cacheBehavior` | `string` | no |
| `description` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRunGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `maxCpus` | `integer` | no |
| `maxRuns` | `integer` | no |
| `maxDuration` | `integer` | no |
| `maxGpus` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSequenceStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `fallbackLocation` | `string` | no |
| `propagatedSetLevelTags` | `List<string>` | no |
| `s3AccessConfig` | `S3AccessConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `sseConfig` | `SseConfig` | no |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | no |
| `propagatedSetLevelTags` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `fallbackLocation` | `string` | no |
| `s3Access` | `SequenceStoreS3Access` | no |
| `eTagAlgorithmFamily` | `string` | no |

## UpdateVariantStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `reference` | `ReferenceItem` | yes |
| `status` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `updateTime` | `timestamp` | yes |

## UpdateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `storageType` | `string` | no |
| `storageCapacity` | `integer` | no |
| `readmeMarkdown` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkflowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `versionName` | `string` | yes |
| `description` | `string` | no |
| `storageType` | `string` | no |
| `storageCapacity` | `integer` | no |
| `readmeMarkdown` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UploadReadSetPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sequenceStoreId` | `string` | yes |
| `uploadId` | `string` | yes |
| `partSource` | `string` | yes |
| `partNumber` | `integer` | yes |
| `payload` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checksum` | `string` | yes |

