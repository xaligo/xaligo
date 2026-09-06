# Agents for Amazon Bedrock Runtime

API version: 2023-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-agent-runtime/2023-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AgenticRetrieveStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agenticRetrieveConfiguration` | `AgenticRetrieveConfiguration` | yes |
| `generateResponse` | `boolean` | no |
| `memoryConfiguration` | `AgenticRetrieveMemoryConfiguration` | no |
| `messages` | `List<AgenticRetrieveMessage>` | yes |
| `nextToken` | `string` | no |
| `policyConfiguration` | `AgenticRetrievePolicyConfiguration` | no |
| `retrievers` | `List<AgenticRetriever>` | yes |
| `userContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stream` | `AgenticRetrieveStreamResponseOutput` | yes |

## CheckIngestedDocumentAcl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `documentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `userContext` | `UserContext` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hasAccess` | `boolean` | yes |

## CreateInvocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `invocationId` | `string` | no |
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAt` | `timestamp` | yes |
| `invocationId` | `string` | yes |
| `sessionId` | `string` | yes |

## CreateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionKeyArn` | `string` | no |
| `sessionMetadata` | `Map<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAt` | `timestamp` | yes |
| `sessionArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionStatus` | `string` | yes |

## DeleteAgentMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAliasId` | `string` | yes |
| `agentId` | `string` | yes |
| `memoryId` | `string` | no |
| `sessionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EndSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionStatus` | `string` | yes |

## GenerateQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryGenerationInput` | `QueryGenerationInput` | yes |
| `transformationConfiguration` | `TransformationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queries` | `List<GeneratedQuery>` | no |

## GetAgentMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAliasId` | `string` | yes |
| `agentId` | `string` | yes |
| `maxItems` | `integer` | no |
| `memoryId` | `string` | yes |
| `memoryType` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryContents` | `List<Memory>` | no |
| `nextToken` | `string` | no |

## GetDocumentContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `documentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `outputFormat` | `string` | no |
| `userContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentContentLength` | `long` | no |
| `mimeType` | `string` | yes |
| `presignedUrl` | `string` | yes |

## GetExecutionFlowSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionIdentifier` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customerEncryptionKeyArn` | `string` | no |
| `definition` | `string` | yes |
| `executionRoleArn` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |
| `flowVersion` | `string` | yes |

## GetFlowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionIdentifier` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endedAt` | `timestamp` | no |
| `errors` | `List<FlowExecutionError>` | no |
| `executionArn` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |
| `flowVersion` | `string` | yes |
| `startedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## GetIngestedDocumentAcl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `documentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentAcl` | `DocumentAcl` | yes |

## GetInvocationStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationIdentifier` | `string` | yes |
| `invocationStepId` | `string` | yes |
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationStep` | `InvocationStep` | yes |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAt` | `timestamp` | yes |
| `encryptionKeyArn` | `string` | no |
| `lastUpdatedAt` | `timestamp` | yes |
| `sessionArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionMetadata` | `Map<string>` | no |
| `sessionStatus` | `string` | yes |

## InvokeAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAliasId` | `string` | yes |
| `agentId` | `string` | yes |
| `bedrockModelConfigurations` | `BedrockModelConfigurations` | no |
| `enableTrace` | `boolean` | no |
| `endSession` | `boolean` | no |
| `inputText` | `string` | no |
| `memoryId` | `string` | no |
| `promptCreationConfigurations` | `PromptCreationConfigurations` | no |
| `sessionId` | `string` | yes |
| `sessionState` | `SessionState` | no |
| `sourceArn` | `string` | no |
| `streamingConfigurations` | `StreamingConfigurations` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `completion` | `ResponseStream` | yes |
| `contentType` | `string` | yes |
| `memoryId` | `string` | no |
| `sessionId` | `string` | yes |

## InvokeFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `enableTrace` | `boolean` | no |
| `executionId` | `string` | no |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |
| `inputs` | `List<FlowInput>` | yes |
| `modelPerformanceConfiguration` | `ModelPerformanceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |
| `responseStream` | `FlowResponseStream` | yes |

## InvokeInlineAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionGroups` | `List<AgentActionGroup>` | no |
| `agentCollaboration` | `string` | no |
| `agentName` | `string` | no |
| `bedrockModelConfigurations` | `InlineBedrockModelConfigurations` | no |
| `collaboratorConfigurations` | `List<CollaboratorConfiguration>` | no |
| `collaborators` | `List<Collaborator>` | no |
| `customOrchestration` | `CustomOrchestration` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `enableTrace` | `boolean` | no |
| `endSession` | `boolean` | no |
| `foundationModel` | `string` | yes |
| `guardrailConfiguration` | `GuardrailConfigurationWithArn` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `inlineSessionState` | `InlineSessionState` | no |
| `inputText` | `string` | no |
| `instruction` | `string` | yes |
| `knowledgeBases` | `List<KnowledgeBase>` | no |
| `orchestrationType` | `string` | no |
| `promptCreationConfigurations` | `PromptCreationConfigurations` | no |
| `promptOverrideConfiguration` | `PromptOverrideConfiguration` | no |
| `sessionId` | `string` | yes |
| `streamingConfigurations` | `StreamingConfigurations` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `completion` | `InlineAgentResponseStream` | yes |
| `contentType` | `string` | yes |
| `sessionId` | `string` | yes |

## ListFlowExecutionEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventType` | `string` | yes |
| `executionIdentifier` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecutionEvents` | `List<FlowExecutionEvent>` | yes |
| `nextToken` | `string` | no |

## ListFlowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowAliasIdentifier` | `string` | no |
| `flowIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecutionSummaries` | `List<FlowExecutionSummary>` | yes |
| `nextToken` | `string` | no |

## ListInvocationSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationStepSummaries` | `List<InvocationStepSummary>` | yes |
| `nextToken` | `string` | no |

## ListInvocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationSummaries` | `List<InvocationSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sessionSummaries` | `List<SessionSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## OptimizePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `input` | `InputPrompt` | yes |
| `targetModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `optimizedPrompt` | `OptimizedPromptStream` | yes |

## PutInvocationStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationIdentifier` | `string` | yes |
| `invocationStepId` | `string` | no |
| `invocationStepTime` | `timestamp` | yes |
| `payload` | `InvocationStepPayload` | yes |
| `sessionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invocationStepId` | `string` | yes |

## Rerank

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `queries` | `List<RerankQuery>` | yes |
| `rerankingConfiguration` | `RerankingConfiguration` | yes |
| `sources` | `List<RerankSource>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `results` | `List<RerankResult>` | yes |

## Retrieve

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailConfiguration` | `GuardrailConfiguration` | no |
| `knowledgeBaseId` | `string` | yes |
| `nextToken` | `string` | no |
| `retrievalConfiguration` | `KnowledgeBaseRetrievalConfiguration` | no |
| `retrievalQuery` | `KnowledgeBaseQuery` | yes |
| `userContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailAction` | `string` | no |
| `nextToken` | `string` | no |
| `retrievalResults` | `List<KnowledgeBaseRetrievalResult>` | yes |

## RetrieveAndGenerate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `input` | `RetrieveAndGenerateInput` | yes |
| `retrieveAndGenerateConfiguration` | `RetrieveAndGenerateConfiguration` | no |
| `sessionConfiguration` | `RetrieveAndGenerateSessionConfiguration` | no |
| `sessionId` | `string` | no |
| `userContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `citations` | `List<Citation>` | no |
| `guardrailAction` | `string` | no |
| `output` | `RetrieveAndGenerateOutput` | yes |
| `sessionId` | `string` | yes |

## RetrieveAndGenerateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `input` | `RetrieveAndGenerateInput` | yes |
| `retrieveAndGenerateConfiguration` | `RetrieveAndGenerateConfiguration` | no |
| `sessionConfiguration` | `RetrieveAndGenerateSessionConfiguration` | no |
| `sessionId` | `string` | no |
| `userContext` | `UserContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `stream` | `RetrieveAndGenerateStreamResponseOutput` | yes |

## StartFlowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowAliasIdentifier` | `string` | yes |
| `flowExecutionName` | `string` | no |
| `flowIdentifier` | `string` | yes |
| `inputs` | `List<FlowInput>` | yes |
| `modelPerformanceConfiguration` | `ModelPerformanceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | no |

## StopFlowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionIdentifier` | `string` | yes |
| `flowAliasIdentifier` | `string` | yes |
| `flowIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | no |
| `status` | `string` | yes |

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


## UpdateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionIdentifier` | `string` | yes |
| `sessionMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `sessionArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionStatus` | `string` | yes |

