# Amazon Bedrock AgentCore

API version: 2024-02-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-agentcore/2024-02-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateMemoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `records` | `List<MemoryRecordCreateInput>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRecords` | `List<MemoryRecordOutput>` | yes |
| `failedRecords` | `List<MemoryRecordOutput>` | yes |

## BatchDeleteMemoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `records` | `List<MemoryRecordDeleteInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRecords` | `List<MemoryRecordOutput>` | yes |
| `failedRecords` | `List<MemoryRecordOutput>` | yes |

## BatchUpdateMemoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `records` | `List<MemoryRecordUpdateInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulRecords` | `List<MemoryRecordOutput>` | yes |
| `failedRecords` | `List<MemoryRecordOutput>` | yes |

## CompleteResourceTokenAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userIdentifier` | `UserIdentifier` | yes |
| `sessionUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateABTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `gatewayArn` | `string` | yes |
| `variants` | `List<Variant>` | yes |
| `gatewayFilter` | `GatewayFilter` | no |
| `evaluationConfig` | `ABTestEvaluationConfig` | yes |
| `roleArn` | `string` | yes |
| `enableOnCreate` | `boolean` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |
| `abTestArn` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `actorId` | `string` | yes |
| `sessionId` | `string` | no |
| `eventTimestamp` | `timestamp` | yes |
| `payload` | `List<PayloadType>` | yes |
| `branch` | `Branch` | no |
| `clientToken` | `string` | no |
| `metadata` | `Map<MetadataValue>` | no |
| `extractionMode` | `string` | no |
| `extractionConfig` | `ExtractionConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `event` | `Event` | yes |

## CreatePaymentInstrument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentConnectorId` | `string` | yes |
| `paymentInstrumentType` | `string` | yes |
| `paymentInstrumentDetails` | `PaymentInstrumentDetails` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentInstrument` | `PaymentInstrument` | yes |

## CreatePaymentSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `limits` | `SessionLimits` | no |
| `expiryTimeInMinutes` | `integer` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentSession` | `PaymentSession` | yes |

## DeleteABTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |
| `abTestArn` | `string` | yes |
| `status` | `string` | yes |

## DeleteBatchEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |
| `batchEvaluationArn` | `string` | yes |
| `status` | `string` | yes |

## DeleteCapacityProviderSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `status` | `string` | yes |

## DeleteEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `sessionId` | `string` | yes |
| `eventId` | `string` | yes |
| `actorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |

## DeleteMemoryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `memoryRecordId` | `string` | yes |
| `namespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryRecordId` | `string` | yes |

## DeletePaymentInstrument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentConnectorId` | `string` | yes |
| `paymentInstrumentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeletePaymentSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentSessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |
| `status` | `string` | yes |

## Evaluate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorId` | `string` | yes |
| `evaluationInput` | `EvaluationInput` | yes |
| `evaluationTarget` | `EvaluationTarget` | no |
| `evaluationReferenceInputs` | `List<EvaluationReferenceInput>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluationResults` | `List<EvaluationResultContent>` | yes |

## GetABTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |
| `abTestArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `gatewayArn` | `string` | yes |
| `variants` | `List<Variant>` | yes |
| `gatewayFilter` | `GatewayFilter` | no |
| `evaluationConfig` | `ABTestEvaluationConfig` | yes |
| `roleArn` | `string` | no |
| `currentRunId` | `string` | no |
| `errorDetails` | `List<string>` | no |
| `startedAt` | `timestamp` | no |
| `stoppedAt` | `timestamp` | no |
| `maxDurationExpiresAt` | `timestamp` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `results` | `ABTestResults` | no |

## GetAgentCard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | no |
| `agentRuntimeArn` | `string` | yes |
| `qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | no |
| `agentCard` | `AgentCard` | yes |
| `statusCode` | `integer` | no |

## GetBatchEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |
| `batchEvaluationArn` | `string` | yes |
| `batchEvaluationName` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `evaluators` | `List<Evaluator>` | no |
| `insights` | `List<Insight>` | no |
| `dataSourceConfig` | `DataSourceConfig` | no |
| `outputConfig` | `OutputConfig` | no |
| `evaluationResults` | `EvaluationJobResults` | no |
| `failureAnalysisResult` | `FailureAnalysisResultContent` | no |
| `userIntentResult` | `UserIntentClusteringResultContent` | no |
| `executionSummaryResult` | `ExecutionSummaryClusteringResultContent` | no |
| `errorDetails` | `List<string>` | no |
| `description` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `kmsKeyArn` | `string` | no |

## GetBrowserSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `name` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `viewPort` | `ViewPort` | no |
| `extensions` | `List<BrowserExtension>` | no |
| `enterprisePolicies` | `List<BrowserEnterprisePolicy>` | no |
| `profileConfiguration` | `BrowserProfileConfiguration` | no |
| `sessionTimeoutSeconds` | `integer` | no |
| `status` | `string` | no |
| `streams` | `BrowserSessionStream` | no |
| `proxyConfiguration` | `ProxyConfiguration` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `sessionReplayArtifact` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetCodeInterpreterSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `name` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `sessionTimeoutSeconds` | `integer` | no |
| `status` | `string` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |

## GetEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `sessionId` | `string` | yes |
| `actorId` | `string` | yes |
| `eventId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `event` | `Event` | yes |

## GetMemoryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `memoryRecordId` | `string` | yes |
| `namespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryRecord` | `MemoryRecord` | yes |

## GetPaymentInstrument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentConnectorId` | `string` | no |
| `paymentInstrumentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentInstrument` | `PaymentInstrument` | yes |

## GetPaymentInstrumentBalance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentConnectorId` | `string` | yes |
| `paymentInstrumentId` | `string` | yes |
| `chain` | `string` | yes |
| `token` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentInstrumentId` | `string` | yes |
| `tokenBalance` | `TokenBalance` | yes |

## GetPaymentSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentSessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentSession` | `PaymentSession` | yes |

## GetRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |
| `recommendationArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `recommendationConfig` | `RecommendationConfig` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `recommendationResult` | `RecommendationResult` | no |
| `kmsKeyArn` | `string` | no |

## GetResourceApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadIdentityToken` | `string` | yes |
| `resourceCredentialProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `string` | yes |

## GetResourceOauth2Token

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadIdentityToken` | `string` | yes |
| `resourceCredentialProviderName` | `string` | yes |
| `scopes` | `List<string>` | yes |
| `oauth2Flow` | `string` | yes |
| `sessionUri` | `string` | no |
| `resourceOauth2ReturnUrl` | `string` | no |
| `forceAuthentication` | `boolean` | no |
| `customParameters` | `Map<string>` | no |
| `customState` | `string` | no |
| `resources` | `List<string>` | no |
| `audiences` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizationUrl` | `string` | no |
| `accessToken` | `string` | no |
| `sessionUri` | `string` | no |
| `sessionStatus` | `string` | no |

## GetResourcePaymentToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadIdentityToken` | `string` | yes |
| `resourceCredentialProviderName` | `string` | yes |
| `paymentTokenRequest` | `PaymentTokenRequestInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentTokenResponse` | `PaymentTokenResponseOutput` | yes |

## GetWorkloadAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadAccessToken` | `string` | yes |

## GetWorkloadAccessTokenForJWT

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `userToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadAccessToken` | `string` | yes |

## GetWorkloadAccessTokenForUserId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadName` | `string` | yes |
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadAccessToken` | `string` | yes |

## IngestData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `source` | `ContentSource` | yes |
| `contentTimestamp` | `timestamp` | yes |
| `actorId` | `string` | yes |
| `sessionId` | `string` | no |
| `extractionConfig` | `ExtractionConfig` | no |
| `metadata` | `Map<MetadataValue>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |

## InvokeAgentRuntime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `accept` | `string` | no |
| `mcpSessionId` | `string` | no |
| `runtimeSessionId` | `string` | no |
| `mcpProtocolVersion` | `string` | no |
| `mcpMethod` | `string` | no |
| `mcpName` | `string` | no |
| `runtimeUserId` | `string` | no |
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `traceState` | `string` | no |
| `baggage` | `string` | no |
| `agentRuntimeArn` | `string` | yes |
| `qualifier` | `string` | no |
| `accountId` | `string` | no |
| `payload` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | no |
| `mcpSessionId` | `string` | no |
| `mcpProtocolVersion` | `string` | no |
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `traceState` | `string` | no |
| `baggage` | `string` | no |
| `contentType` | `string` | yes |
| `response` | `blob` | no |
| `statusCode` | `integer` | no |

## InvokeAgentRuntimeCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `accept` | `string` | no |
| `runtimeSessionId` | `string` | no |
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `traceState` | `string` | no |
| `baggage` | `string` | no |
| `agentRuntimeArn` | `string` | yes |
| `qualifier` | `string` | no |
| `accountId` | `string` | no |
| `body` | `InvokeAgentRuntimeCommandRequestBody` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | no |
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `traceState` | `string` | no |
| `baggage` | `string` | no |
| `contentType` | `string` | yes |
| `statusCode` | `integer` | no |
| `stream` | `InvokeAgentRuntimeCommandStreamOutput` | yes |

## InvokeBrowser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `action` | `BrowserAction` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `BrowserActionResult` | yes |
| `sessionId` | `string` | yes |

## InvokeCodeInterpreter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | no |
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `name` | `string` | yes |
| `arguments` | `ToolArguments` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | no |
| `stream` | `CodeInterpreterStreamOutput` | yes |

## InvokeHarness

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessArn` | `string` | yes |
| `qualifier` | `string` | no |
| `runtimeSessionId` | `string` | yes |
| `runtimeUserId` | `string` | no |
| `traceParent` | `string` | no |
| `traceState` | `string` | no |
| `traceId` | `string` | no |
| `baggage` | `string` | no |
| `messages` | `List<HarnessMessage>` | yes |
| `model` | `HarnessModelConfiguration` | no |
| `systemPrompt` | `List<HarnessSystemContentBlock>` | no |
| `tools` | `List<HarnessTool>` | no |
| `skills` | `List<HarnessSkill>` | no |
| `allowedTools` | `List<string>` | no |
| `maxIterations` | `integer` | no |
| `maxTokens` | `integer` | no |
| `timeoutSeconds` | `integer` | no |
| `actorId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stream` | `InvokeHarnessStreamOutput` | yes |

## ListABTests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTests` | `List<ABTestSummary>` | yes |
| `nextToken` | `string` | no |

## ListActors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actorSummaries` | `List<ActorSummary>` | yes |
| `nextToken` | `string` | no |

## ListBatchEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluations` | `List<BatchEvaluationSummary>` | yes |
| `nextToken` | `string` | no |

## ListBrowserSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BrowserSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListCodeInterpreterSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<CodeInterpreterSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `sessionId` | `string` | yes |
| `actorId` | `string` | yes |
| `includePayloads` | `boolean` | no |
| `filter` | `FilterInput` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<Event>` | yes |
| `nextToken` | `string` | no |

## ListMemoryExtractionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `filter` | `ExtractionJobFilterInput` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<ExtractionJobMetadata>` | yes |
| `nextToken` | `string` | no |

## ListMemoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `namespace` | `string` | no |
| `namespacePath` | `string` | no |
| `memoryStrategyId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `metadataFilters` | `List<MemoryMetadataFilterExpression>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryRecordSummaries` | `List<MemoryRecordSummary>` | yes |
| `nextToken` | `string` | no |

## ListPaymentInstruments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentConnectorId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentInstruments` | `List<PaymentInstrumentSummary>` | yes |
| `nextToken` | `string` | no |

## ListPaymentSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentSessions` | `List<PaymentSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `statusFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationSummaries` | `List<RecommendationSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `actorId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filter` | `SessionFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionSummaries` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## ProcessPayment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `agentName` | `string` | no |
| `paymentManagerArn` | `string` | yes |
| `paymentSessionId` | `string` | yes |
| `paymentInstrumentId` | `string` | yes |
| `paymentType` | `string` | yes |
| `paymentInput` | `PaymentInput` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `processPaymentId` | `string` | yes |
| `paymentManagerArn` | `string` | yes |
| `paymentSessionId` | `string` | yes |
| `paymentInstrumentId` | `string` | yes |
| `paymentType` | `string` | yes |
| `status` | `string` | yes |
| `paymentOutput` | `PaymentOutput` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## RetrieveMemoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `namespace` | `string` | no |
| `namespacePath` | `string` | no |
| `searchCriteria` | `SearchCriteria` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryRecordSummaries` | `List<MemoryRecordSummary>` | yes |
| `nextToken` | `string` | no |

## SaveBrowserSessionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `profileIdentifier` | `string` | yes |
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileIdentifier` | `string` | yes |
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## SearchRegistryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchQuery` | `string` | yes |
| `registryIds` | `List<string>` | yes |
| `maxResults` | `integer` | no |
| `filters` | `MetadataFilterExpression` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<RegistryRecordSummary>` | yes |

## StartBatchEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationName` | `string` | yes |
| `evaluators` | `List<Evaluator>` | no |
| `insights` | `List<Insight>` | no |
| `dataSourceConfig` | `DataSourceConfig` | yes |
| `clientToken` | `string` | no |
| `evaluationMetadata` | `EvaluationMetadata` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `description` | `string` | no |
| `outputConfig` | `OutputConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |
| `batchEvaluationArn` | `string` | yes |
| `batchEvaluationName` | `string` | yes |
| `evaluators` | `List<Evaluator>` | no |
| `insights` | `List<Insight>` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `outputConfig` | `OutputConfig` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |
| `description` | `string` | no |

## StartBrowserSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `browserIdentifier` | `string` | yes |
| `name` | `string` | no |
| `sessionTimeoutSeconds` | `integer` | no |
| `viewPort` | `ViewPort` | no |
| `extensions` | `List<BrowserExtension>` | no |
| `profileConfiguration` | `BrowserProfileConfiguration` | no |
| `proxyConfiguration` | `ProxyConfiguration` | no |
| `enterprisePolicies` | `List<BrowserEnterprisePolicy>` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `streams` | `BrowserSessionStream` | no |

## StartCodeInterpreterSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `codeInterpreterIdentifier` | `string` | yes |
| `name` | `string` | no |
| `sessionTimeoutSeconds` | `integer` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## StartMemoryExtractionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `extractionJob` | `ExtractionJob` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

## StartRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `recommendationConfig` | `RecommendationConfig` | yes |
| `kmsKeyArn` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationId` | `string` | yes |
| `recommendationArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `recommendationConfig` | `RecommendationConfig` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## StopBatchEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `batchEvaluationId` | `string` | yes |
| `batchEvaluationArn` | `string` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |

## StopBrowserSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## StopCodeInterpreterSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `traceId` | `string` | no |
| `traceParent` | `string` | no |
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## StopRuntimeSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | yes |
| `agentRuntimeArn` | `string` | yes |
| `qualifier` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeSessionId` | `string` | no |
| `statusCode` | `integer` | no |

## UpdateABTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |
| `clientToken` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `variants` | `List<Variant>` | no |
| `gatewayFilter` | `GatewayFilter` | no |
| `evaluationConfig` | `ABTestEvaluationConfig` | no |
| `roleArn` | `string` | no |
| `executionStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `abTestId` | `string` | yes |
| `abTestArn` | `string` | yes |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateBrowserStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `streamUpdate` | `StreamUpdate` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserIdentifier` | `string` | yes |
| `sessionId` | `string` | yes |
| `streams` | `BrowserSessionStream` | yes |
| `updatedAt` | `timestamp` | yes |

