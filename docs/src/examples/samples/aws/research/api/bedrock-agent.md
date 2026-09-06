# Agents for Amazon Bedrock

API version: 2023-06-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-agent/2023-06-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAgentCollaborator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `agentDescriptor` | `AgentDescriptor` | yes |
| `collaboratorName` | `string` | yes |
| `collaborationInstruction` | `string` | yes |
| `relayConversationHistory` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentCollaborator` | `AgentCollaborator` | yes |

## AssociateAgentKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `description` | `string` | yes |
| `knowledgeBaseState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentKnowledgeBase` | `AgentKnowledgeBase` | yes |

## CreateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentName` | `string` | yes |
| `clientToken` | `string` | no |
| `instruction` | `string` | no |
| `foundationModel` | `string` | no |
| `description` | `string` | no |
| `orchestrationType` | `string` | no |
| `customOrchestration` | `CustomOrchestration` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `agentResourceRoleArn` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `promptOverrideConfiguration` | `PromptOverrideConfiguration` | no |
| `guardrailConfiguration` | `GuardrailConfiguration` | no |
| `memoryConfiguration` | `MemoryConfiguration` | no |
| `agentCollaboration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agent` | `Agent` | yes |

## CreateAgentActionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `actionGroupName` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `parentActionGroupSignature` | `string` | no |
| `parentActionGroupSignatureParams` | `Map<string>` | no |
| `actionGroupExecutor` | `ActionGroupExecutor` | no |
| `apiSchema` | `APISchema` | no |
| `actionGroupState` | `string` | no |
| `functionSchema` | `FunctionSchema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentActionGroup` | `AgentActionGroup` | yes |

## CreateAgentAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentAliasName` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `routingConfiguration` | `List<AgentAliasRoutingConfigurationListItem>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAlias` | `AgentAlias` | yes |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `dataSourceConfiguration` | `DataSourceConfiguration` | yes |
| `dataDeletionPolicy` | `string` | no |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `vectorIngestionConfiguration` | `VectorIngestionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | yes |

## CreateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `definition` | `FlowDefinition` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `version` | `string` | yes |
| `definition` | `FlowDefinition` | no |

## CreateFlowAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<FlowAliasRoutingConfigurationListItem>` | yes |
| `concurrencyConfiguration` | `FlowAliasConcurrencyConfiguration` | no |
| `flowIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<FlowAliasRoutingConfigurationListItem>` | yes |
| `concurrencyConfiguration` | `FlowAliasConcurrencyConfiguration` | no |
| `flowId` | `string` | yes |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateFlowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `version` | `string` | yes |
| `definition` | `FlowDefinition` | no |

## CreateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `knowledgeBaseConfiguration` | `KnowledgeBaseConfiguration` | yes |
| `storageConfiguration` | `StorageConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBase` | yes |

## CreatePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `version` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## CreatePromptVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptIdentifier` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `version` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentStatus` | `string` | yes |

## DeleteAgentActionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `actionGroupId` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAgentAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentAliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentAliasId` | `string` | yes |
| `agentAliasStatus` | `string` | yes |

## DeleteAgentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `agentStatus` | `string` | yes |

## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `status` | `string` | yes |

## DeleteFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

## DeleteFlowAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `aliasIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowId` | `string` | yes |
| `id` | `string` | yes |

## DeleteFlowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `flowVersion` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `version` | `string` | yes |

## DeleteKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `status` | `string` | yes |

## DeleteKnowledgeBaseDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `clientToken` | `string` | no |
| `documentIdentifiers` | `List<DocumentIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentDetails` | `List<KnowledgeBaseDocumentDetail>` | no |

## DeletePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptIdentifier` | `string` | yes |
| `promptVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `version` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `revisionId` | `string` | no |

## DisassociateAgentCollaborator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `collaboratorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAgentKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agent` | `Agent` | yes |

## GetAgentActionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `actionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentActionGroup` | `AgentActionGroup` | yes |

## GetAgentAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentAliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAlias` | `AgentAlias` | yes |

## GetAgentCollaborator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `collaboratorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentCollaborator` | `AgentCollaborator` | yes |

## GetAgentKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentKnowledgeBase` | `AgentKnowledgeBase` | yes |

## GetAgentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentVersion` | `AgentVersion` | yes |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | yes |

## GetFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `version` | `string` | yes |
| `definition` | `FlowDefinition` | no |
| `validations` | `List<FlowValidation>` | no |

## GetFlowAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `aliasIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<FlowAliasRoutingConfigurationListItem>` | yes |
| `concurrencyConfiguration` | `FlowAliasConcurrencyConfiguration` | no |
| `flowId` | `string` | yes |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetFlowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `flowVersion` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `version` | `string` | yes |
| `definition` | `FlowDefinition` | no |

## GetIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `ingestionJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionJob` | `IngestionJob` | yes |

## GetKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBase` | yes |

## GetKnowledgeBaseDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `documentIdentifiers` | `List<DocumentIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentDetails` | `List<KnowledgeBaseDocumentDetail>` | no |

## GetPrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptIdentifier` | `string` | yes |
| `promptVersion` | `string` | no |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `version` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |
| `revisionId` | `string` | yes |

## IngestKnowledgeBaseDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `clientToken` | `string` | no |
| `documents` | `List<KnowledgeBaseDocument>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentDetails` | `List<KnowledgeBaseDocumentDetail>` | no |

## ListAgentActionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionGroupSummaries` | `List<ActionGroupSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAliasSummaries` | `List<AgentAliasSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentCollaborators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentCollaboratorSummaries` | `List<AgentCollaboratorSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentKnowledgeBases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentKnowledgeBaseSummaries` | `List<AgentKnowledgeBaseSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentVersionSummaries` | `List<AgentVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSummaries` | `List<AgentSummary>` | yes |
| `nextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceSummaries` | `List<DataSourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListFlowAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowAliasSummaries` | `List<FlowAliasSummary>` | yes |
| `nextToken` | `string` | no |

## ListFlowVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowVersionSummaries` | `List<FlowVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowSummaries` | `List<FlowSummary>` | yes |
| `nextToken` | `string` | no |

## ListIngestionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `filters` | `List<IngestionJobFilter>` | no |
| `sortBy` | `IngestionJobSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionJobSummaries` | `List<IngestionJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListKnowledgeBaseDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentDetails` | `List<KnowledgeBaseDocumentDetail>` | yes |
| `nextToken` | `string` | no |

## ListKnowledgeBases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseSummaries` | `List<KnowledgeBaseSummary>` | yes |
| `nextToken` | `string` | no |

## ListPrompts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptSummaries` | `List<PromptSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PrepareAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentStatus` | `string` | yes |
| `agentVersion` | `string` | yes |
| `preparedAt` | `timestamp` | yes |

## PrepareFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | yes |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `revisionId` | `string` | yes |

## StartIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionJob` | `IngestionJob` | yes |

## StopIngestionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `ingestionJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ingestionJob` | `IngestionJob` | yes |

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


## UpdateAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentName` | `string` | yes |
| `instruction` | `string` | no |
| `foundationModel` | `string` | yes |
| `description` | `string` | no |
| `orchestrationType` | `string` | no |
| `customOrchestration` | `CustomOrchestration` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `agentResourceRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `promptOverrideConfiguration` | `PromptOverrideConfiguration` | no |
| `guardrailConfiguration` | `GuardrailConfiguration` | no |
| `memoryConfiguration` | `MemoryConfiguration` | no |
| `agentCollaboration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agent` | `Agent` | yes |

## UpdateAgentActionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `actionGroupId` | `string` | yes |
| `actionGroupName` | `string` | yes |
| `description` | `string` | no |
| `parentActionGroupSignature` | `string` | no |
| `parentActionGroupSignatureParams` | `Map<string>` | no |
| `actionGroupExecutor` | `ActionGroupExecutor` | no |
| `actionGroupState` | `string` | no |
| `apiSchema` | `APISchema` | no |
| `functionSchema` | `FunctionSchema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentActionGroup` | `AgentActionGroup` | yes |

## UpdateAgentAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentAliasId` | `string` | yes |
| `agentAliasName` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<AgentAliasRoutingConfigurationListItem>` | no |
| `aliasInvocationState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentAlias` | `AgentAlias` | yes |

## UpdateAgentCollaborator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `collaboratorId` | `string` | yes |
| `agentDescriptor` | `AgentDescriptor` | yes |
| `collaboratorName` | `string` | yes |
| `collaborationInstruction` | `string` | yes |
| `relayConversationHistory` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentCollaborator` | `AgentCollaborator` | yes |

## UpdateAgentKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentId` | `string` | yes |
| `agentVersion` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `description` | `string` | no |
| `knowledgeBaseState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentKnowledgeBase` | `AgentKnowledgeBase` | yes |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `dataSourceConfiguration` | `DataSourceConfiguration` | yes |
| `dataDeletionPolicy` | `string` | no |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `vectorIngestionConfiguration` | `VectorIngestionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | yes |

## UpdateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `definition` | `FlowDefinition` | no |
| `flowIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `customerEncryptionKeyArn` | `string` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `version` | `string` | yes |
| `definition` | `FlowDefinition` | no |

## UpdateFlowAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<FlowAliasRoutingConfigurationListItem>` | yes |
| `concurrencyConfiguration` | `FlowAliasConcurrencyConfiguration` | no |
| `flowIdentifier` | `string` | yes |
| `aliasIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<FlowAliasRoutingConfigurationListItem>` | yes |
| `concurrencyConfiguration` | `FlowAliasConcurrencyConfiguration` | no |
| `flowId` | `string` | yes |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `knowledgeBaseConfiguration` | `KnowledgeBaseConfiguration` | yes |
| `storageConfiguration` | `StorageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBase` | yes |

## UpdatePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `promptIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `customerEncryptionKeyArn` | `string` | no |
| `defaultVariant` | `string` | no |
| `variants` | `List<PromptVariant>` | no |
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `version` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## ValidateFlowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `definition` | `FlowDefinition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `validations` | `List<FlowValidation>` | yes |

