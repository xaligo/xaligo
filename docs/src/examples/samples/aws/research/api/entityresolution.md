# AWS EntityResolution

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/entityresolution/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddPolicyStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `statementId` | `string` | yes |
| `effect` | `string` | yes |
| `action` | `List<string>` | yes |
| `principal` | `List<string>` | yes |
| `condition` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `token` | `string` | yes |
| `policy` | `string` | no |

## BatchDeleteUniqueId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `inputSource` | `string` | no |
| `uniqueIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `errors` | `List<DeleteUniqueIdError>` | yes |
| `deleted` | `List<DeletedUniqueId>` | yes |
| `disconnectedUniqueIds` | `List<string>` | yes |

## CreateIdMappingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdMappingWorkflowInputSource>` | yes |
| `outputSourceConfig` | `List<IdMappingWorkflowOutputSource>` | no |
| `idMappingTechniques` | `IdMappingTechniques` | yes |
| `incrementalRunConfig` | `IdMappingIncrementalRunConfig` | no |
| `roleArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `workflowArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdMappingWorkflowInputSource>` | yes |
| `outputSourceConfig` | `List<IdMappingWorkflowOutputSource>` | no |
| `idMappingTechniques` | `IdMappingTechniques` | yes |
| `incrementalRunConfig` | `IdMappingIncrementalRunConfig` | no |
| `roleArn` | `string` | no |

## CreateIdNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdNamespaceInputSource>` | no |
| `idMappingWorkflowProperties` | `List<IdNamespaceIdMappingWorkflowProperties>` | no |
| `type` | `string` | yes |
| `roleArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |
| `idNamespaceArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdNamespaceInputSource>` | no |
| `idMappingWorkflowProperties` | `List<IdNamespaceIdMappingWorkflowProperties>` | no |
| `type` | `string` | yes |
| `roleArn` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## CreateMatchingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<InputSource>` | yes |
| `outputSourceConfig` | `List<OutputSource>` | yes |
| `resolutionTechniques` | `ResolutionTechniques` | yes |
| `incrementalRunConfig` | `IncrementalRunConfig` | no |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `workflowArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<InputSource>` | yes |
| `outputSourceConfig` | `List<OutputSource>` | yes |
| `resolutionTechniques` | `ResolutionTechniques` | yes |
| `incrementalRunConfig` | `IncrementalRunConfig` | no |
| `roleArn` | `string` | yes |

## CreateSchemaMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |
| `description` | `string` | no |
| `mappedInputFields` | `List<SchemaInputAttribute>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |
| `schemaArn` | `string` | yes |
| `description` | `string` | yes |
| `mappedInputFields` | `List<SchemaInputAttribute>` | yes |

## DeleteIdMappingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | yes |

## DeleteIdNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | yes |

## DeleteMatchingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | yes |

## DeletePolicyStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `statementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `token` | `string` | yes |
| `policy` | `string` | no |

## DeleteSchemaMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | yes |

## GenerateMatchId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `records` | `List<Record>` | yes |
| `processingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `matchGroups` | `List<MatchGroup>` | yes |
| `failedRecords` | `List<FailedRecord>` | yes |

## GetIdMappingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | no |
| `metrics` | `IdMappingJobMetrics` | no |
| `errorDetails` | `ErrorDetails` | no |
| `outputSourceConfig` | `List<IdMappingJobOutputSource>` | no |
| `jobType` | `string` | no |

## GetIdMappingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `workflowArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdMappingWorkflowInputSource>` | yes |
| `outputSourceConfig` | `List<IdMappingWorkflowOutputSource>` | no |
| `idMappingTechniques` | `IdMappingTechniques` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `incrementalRunConfig` | `IdMappingIncrementalRunConfig` | no |
| `roleArn` | `string` | no |
| `tags` | `Map<string>` | no |

## GetIdNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |
| `idNamespaceArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdNamespaceInputSource>` | no |
| `idMappingWorkflowProperties` | `List<IdNamespaceIdMappingWorkflowProperties>` | no |
| `type` | `string` | yes |
| `roleArn` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetMatchId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `record` | `Map<string>` | yes |
| `applyNormalization` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `matchId` | `string` | no |
| `matchRule` | `string` | no |

## GetMatchingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `status` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | no |
| `metrics` | `JobMetrics` | no |
| `errorDetails` | `ErrorDetails` | no |
| `outputSourceConfig` | `List<JobOutputSource>` | no |

## GetMatchingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `workflowArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<InputSource>` | yes |
| `outputSourceConfig` | `List<OutputSource>` | yes |
| `resolutionTechniques` | `ResolutionTechniques` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `incrementalRunConfig` | `IncrementalRunConfig` | no |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `token` | `string` | yes |
| `policy` | `string` | no |

## GetProviderService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `providerName` | `string` | yes |
| `providerServiceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `providerName` | `string` | yes |
| `providerServiceName` | `string` | yes |
| `providerServiceDisplayName` | `string` | yes |
| `providerServiceType` | `string` | yes |
| `providerServiceArn` | `string` | yes |
| `providerConfigurationDefinition` | `Document` | no |
| `providerIdNameSpaceConfiguration` | `ProviderIdNameSpaceConfiguration` | no |
| `providerJobConfiguration` | `Document` | no |
| `providerEndpointConfiguration` | `ProviderEndpointConfiguration` | yes |
| `anonymizedOutput` | `boolean` | yes |
| `providerEntityOutputDefinition` | `Document` | yes |
| `providerIntermediateDataAccessConfiguration` | `ProviderIntermediateDataAccessConfiguration` | no |
| `providerComponentSchema` | `ProviderComponentSchema` | no |

## GetSchemaMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |
| `schemaArn` | `string` | yes |
| `description` | `string` | no |
| `mappedInputFields` | `List<SchemaInputAttribute>` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |
| `hasWorkflows` | `boolean` | yes |

## ListIdMappingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSummary>` | no |
| `nextToken` | `string` | no |

## ListIdMappingWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowSummaries` | `List<IdMappingWorkflowSummary>` | no |
| `nextToken` | `string` | no |

## ListIdNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceSummaries` | `List<IdNamespaceSummary>` | no |
| `nextToken` | `string` | no |

## ListMatchingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSummary>` | no |
| `nextToken` | `string` | no |

## ListMatchingWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowSummaries` | `List<MatchingWorkflowSummary>` | no |
| `nextToken` | `string` | no |

## ListProviderServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `providerName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `providerServiceSummaries` | `List<ProviderServiceSummary>` | no |
| `nextToken` | `string` | no |

## ListSchemaMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaList` | `List<SchemaMappingSummary>` | no |
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

## PutPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `token` | `string` | no |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `token` | `string` | yes |
| `policy` | `string` | no |

## StartIdMappingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `outputSourceConfig` | `List<IdMappingJobOutputSource>` | no |
| `jobType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `outputSourceConfig` | `List<IdMappingJobOutputSource>` | no |
| `jobType` | `string` | no |

## StartMatchingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |

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


## UpdateIdMappingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdMappingWorkflowInputSource>` | yes |
| `outputSourceConfig` | `List<IdMappingWorkflowOutputSource>` | no |
| `idMappingTechniques` | `IdMappingTechniques` | yes |
| `incrementalRunConfig` | `IdMappingIncrementalRunConfig` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `workflowArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdMappingWorkflowInputSource>` | yes |
| `outputSourceConfig` | `List<IdMappingWorkflowOutputSource>` | no |
| `idMappingTechniques` | `IdMappingTechniques` | yes |
| `incrementalRunConfig` | `IdMappingIncrementalRunConfig` | no |
| `roleArn` | `string` | no |

## UpdateIdNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdNamespaceInputSource>` | no |
| `idMappingWorkflowProperties` | `List<IdNamespaceIdMappingWorkflowProperties>` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idNamespaceName` | `string` | yes |
| `idNamespaceArn` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<IdNamespaceInputSource>` | no |
| `idMappingWorkflowProperties` | `List<IdNamespaceIdMappingWorkflowProperties>` | no |
| `type` | `string` | yes |
| `roleArn` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateMatchingWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<InputSource>` | yes |
| `outputSourceConfig` | `List<OutputSource>` | yes |
| `resolutionTechniques` | `ResolutionTechniques` | yes |
| `incrementalRunConfig` | `IncrementalRunConfig` | no |
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowName` | `string` | yes |
| `description` | `string` | no |
| `inputSourceConfig` | `List<InputSource>` | yes |
| `outputSourceConfig` | `List<OutputSource>` | yes |
| `resolutionTechniques` | `ResolutionTechniques` | yes |
| `incrementalRunConfig` | `IncrementalRunConfig` | no |
| `roleArn` | `string` | yes |

## UpdateSchemaMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |
| `description` | `string` | no |
| `mappedInputFields` | `List<SchemaInputAttribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schemaName` | `string` | yes |
| `schemaArn` | `string` | yes |
| `description` | `string` | no |
| `mappedInputFields` | `List<SchemaInputAttribute>` | yes |

