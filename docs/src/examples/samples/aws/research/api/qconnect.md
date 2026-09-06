# Amazon Q Connect

API version: 2020-10-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/qconnect/2020-10-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `versionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplateArn` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `versionNumber` | `long` | yes |

## CreateAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `configuration` | `AIAgentConfiguration` | yes |
| `visibilityStatus` | `string` | yes |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgent` | `AIAgentData` | no |

## CreateAIAgentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |
| `modifiedTime` | `timestamp` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgent` | `AIAgentData` | no |
| `versionNumber` | `long` | no |

## CreateAIGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `name` | `string` | yes |
| `blockedInputMessaging` | `string` | yes |
| `blockedOutputsMessaging` | `string` | yes |
| `visibilityStatus` | `string` | yes |
| `description` | `string` | no |
| `topicPolicyConfig` | `AIGuardrailTopicPolicyConfig` | no |
| `contentPolicyConfig` | `AIGuardrailContentPolicyConfig` | no |
| `wordPolicyConfig` | `AIGuardrailWordPolicyConfig` | no |
| `sensitiveInformationPolicyConfig` | `AIGuardrailSensitiveInformationPolicyConfig` | no |
| `contextualGroundingPolicyConfig` | `AIGuardrailContextualGroundingPolicyConfig` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrail` | `AIGuardrailData` | no |

## CreateAIGuardrailVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |
| `modifiedTime` | `timestamp` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrail` | `AIGuardrailData` | no |
| `versionNumber` | `long` | no |

## CreateAIPrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `templateConfiguration` | `AIPromptTemplateConfiguration` | yes |
| `visibilityStatus` | `string` | yes |
| `templateType` | `string` | yes |
| `modelId` | `string` | yes |
| `apiFormat` | `string` | yes |
| `tags` | `Map<string>` | no |
| `description` | `string` | no |
| `inferenceConfiguration` | `AIPromptInferenceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPrompt` | `AIPromptData` | no |

## CreateAIPromptVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |
| `modifiedTime` | `timestamp` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPrompt` | `AIPromptData` | no |
| `versionNumber` | `long` | no |

## CreateAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistant` | `AssistantData` | no |

## CreateAssistantAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `associationType` | `string` | yes |
| `association` | `AssistantAssociationInputData` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociation` | `AssistantAssociationData` | no |

## CreateContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `name` | `string` | yes |
| `title` | `string` | no |
| `overrideLinkOutUri` | `string` | no |
| `metadata` | `Map<string>` | no |
| `uploadId` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `ContentData` | no |

## CreateContentAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |
| `associationType` | `string` | yes |
| `association` | `ContentAssociationContents` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentAssociation` | `ContentAssociationData` | no |

## CreateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `knowledgeBaseType` | `string` | yes |
| `sourceConfiguration` | `SourceConfiguration` | no |
| `renderingConfiguration` | `RenderingConfiguration` | no |
| `vectorIngestionConfiguration` | `VectorIngestionConfiguration` | no |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBaseData` | no |

## CreateMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `name` | `string` | no |
| `content` | `MessageTemplateContentProvider` | no |
| `description` | `string` | no |
| `channelSubtype` | `string` | yes |
| `language` | `string` | no |
| `sourceConfiguration` | `MessageTemplateSourceConfiguration` | no |
| `defaultAttributes` | `MessageTemplateAttributes` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplate` | `MessageTemplateData` | no |

## CreateMessageTemplateAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `contentDisposition` | `string` | yes |
| `name` | `string` | yes |
| `body` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachment` | `MessageTemplateAttachment` | no |

## CreateMessageTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `messageTemplateContentSha256` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplate` | `ExtendedMessageTemplateData` | no |

## CreateQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `name` | `string` | yes |
| `content` | `QuickResponseDataProvider` | yes |
| `contentType` | `string` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |
| `description` | `string` | no |
| `shortcutKey` | `string` | no |
| `isActive` | `boolean` | no |
| `channels` | `List<string>` | no |
| `language` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

## CreateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `tagFilter` | `TagFilter` | no |
| `aiAgentConfiguration` | `Map<AIAgentConfigurationData>` | no |
| `contactArn` | `string` | no |
| `orchestratorConfigurationList` | `List<OrchestratorConfigurationEntry>` | no |
| `removeOrchestratorConfigurationList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `SessionData` | no |

## DeactivateMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `versionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplateArn` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `versionNumber` | `long` | yes |

## DeleteAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAIAgentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |
| `versionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAIGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAIGuardrailVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |
| `versionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAIPrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAIPromptVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |
| `versionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssistantAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociationId` | `string` | yes |
| `assistantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContentAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |
| `contentAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `importJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMessageTemplateAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `attachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `quickResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgent` | `AIAgentData` | no |
| `versionNumber` | `long` | no |

## GetAIGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrail` | `AIGuardrailData` | no |
| `versionNumber` | `long` | no |

## GetAIPrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPrompt` | `AIPromptData` | no |
| `versionNumber` | `long` | no |

## GetAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistant` | `AssistantData` | no |

## GetAssistantAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociationId` | `string` | yes |
| `assistantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociation` | `AssistantAssociationData` | no |

## GetContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `ContentData` | no |

## GetContentAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |
| `contentAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentAssociation` | `ContentAssociationData` | no |

## GetContentSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentSummary` | `ContentSummary` | no |

## GetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJobId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJob` | `ImportJobData` | no |

## GetKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBaseData` | no |

## GetMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplateId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplate` | `ExtendedMessageTemplateData` | no |

## GetNextMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `nextMessageToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `response` | `MessageOutput` | yes |
| `requestMessageId` | `string` | yes |
| `conversationState` | `ConversationState` | yes |
| `nextMessageToken` | `string` | no |
| `conversationSessionData` | `List<RuntimeSessionData>` | no |
| `chunkedResponseTerminated` | `boolean` | no |

## GetQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponseId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

## GetRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `maxResults` | `integer` | no |
| `waitTimeSeconds` | `integer` | no |
| `nextChunkToken` | `string` | no |
| `recommendationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendations` | `List<RecommendationData>` | yes |
| `triggers` | `List<RecommendationTrigger>` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `SessionData` | no |

## ListAIAgentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `origin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgentVersionSummaries` | `List<AIAgentVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListAIAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `origin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgentSummaries` | `List<AIAgentSummary>` | yes |
| `nextToken` | `string` | no |

## ListAIGuardrailVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrailVersionSummaries` | `List<AIGuardrailVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListAIGuardrails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrailSummaries` | `List<AIGuardrailSummary>` | yes |
| `nextToken` | `string` | no |

## ListAIPromptVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `origin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPromptVersionSummaries` | `List<AIPromptVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListAIPrompts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `origin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPromptSummaries` | `List<AIPromptSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssistantAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assistantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociationSummaries` | `List<AssistantAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssistants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantSummaries` | `List<AssistantSummary>` | yes |
| `nextToken` | `string` | no |

## ListContentAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentAssociationSummaries` | `List<ContentAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListContents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentSummaries` | `List<ContentSummary>` | yes |
| `nextToken` | `string` | no |

## ListImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJobSummaries` | `List<ImportJobSummary>` | yes |
| `nextToken` | `string` | no |

## ListKnowledgeBases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseSummaries` | `List<KnowledgeBaseSummary>` | yes |
| `nextToken` | `string` | no |

## ListMessageTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplateVersionSummaries` | `List<MessageTemplateVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListMessageTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplateSummaries` | `List<MessageTemplateSummary>` | yes |
| `nextToken` | `string` | no |

## ListMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messages` | `List<MessageOutput>` | yes |
| `nextToken` | `string` | no |

## ListModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiPromptType` | `string` | no |
| `modelLifecycle` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelSummaries` | `List<ModelSummary>` | yes |
| `nextToken` | `string` | no |

## ListQuickResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponseSummaries` | `List<QuickResponseSummary>` | yes |
| `nextToken` | `string` | no |

## ListSpans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spans` | `List<Span>` | yes |
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

## NotifyRecommendationsReceived

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `recommendationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendationIds` | `List<string>` | no |
| `errors` | `List<NotifyRecommendationsReceivedError>` | no |

## PutFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `targetId` | `string` | yes |
| `targetType` | `string` | yes |
| `contentFeedback` | `ContentFeedbackData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `assistantArn` | `string` | yes |
| `targetId` | `string` | yes |
| `targetType` | `string` | yes |
| `contentFeedback` | `ContentFeedbackData` | yes |

## QueryAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `queryText` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sessionId` | `string` | no |
| `queryCondition` | `List<QueryCondition>` | no |
| `queryInputData` | `QueryInputData` | no |
| `overrideKnowledgeBaseSearchType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<ResultData>` | yes |
| `nextToken` | `string` | no |

## RemoveAssistantAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentType` | `string` | yes |
| `orchestratorUseCase` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveKnowledgeBaseTemplateUri

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RenderMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `attributes` | `MessageTemplateAttributes` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `MessageTemplateContentProvider` | no |
| `sourceConfigurationSummary` | `MessageTemplateSourceConfigurationSummary` | no |
| `attributesNotInterpolated` | `List<string>` | no |
| `attachments` | `List<MessageTemplateAttachment>` | no |

## Retrieve

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `retrievalConfiguration` | `RetrievalConfiguration` | yes |
| `retrievalQuery` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<RetrieveResult>` | yes |

## SearchContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `knowledgeBaseId` | `string` | yes |
| `searchExpression` | `SearchExpression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentSummaries` | `List<ContentSummary>` | yes |
| `nextToken` | `string` | no |

## SearchMessageTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `searchExpression` | `MessageTemplateSearchExpression` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<MessageTemplateSearchResultData>` | yes |
| `nextToken` | `string` | no |

## SearchQuickResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `searchExpression` | `QuickResponseSearchExpression` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `attributes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `List<QuickResponseSearchResultData>` | yes |
| `nextToken` | `string` | no |

## SearchSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `assistantId` | `string` | yes |
| `searchExpression` | `SearchExpression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionSummaries` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## SendMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `type` | `string` | yes |
| `message` | `MessageInput` | yes |
| `aiAgentId` | `string` | no |
| `conversationContext` | `ConversationContext` | no |
| `configuration` | `MessageConfiguration` | no |
| `clientToken` | `string` | no |
| `orchestratorUseCase` | `string` | no |
| `metadata` | `Map<string>` | no |
| `originRequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestMessageId` | `string` | yes |
| `configuration` | `MessageConfiguration` | no |
| `nextMessageToken` | `string` | yes |

## StartContentUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `contentType` | `string` | yes |
| `presignedUrlTimeToLive` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | yes |
| `url` | `string` | yes |
| `urlExpiry` | `timestamp` | yes |
| `headersToInclude` | `Map<string>` | yes |

## StartImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `importJobType` | `string` | yes |
| `uploadId` | `string` | yes |
| `clientToken` | `string` | no |
| `metadata` | `Map<string>` | no |
| `externalSourceConfiguration` | `ExternalSourceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJob` | `ImportJobData` | no |

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


## UpdateAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `aiAgentId` | `string` | yes |
| `visibilityStatus` | `string` | yes |
| `configuration` | `AIAgentConfiguration` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiAgent` | `AIAgentData` | no |

## UpdateAIGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `aiGuardrailId` | `string` | yes |
| `visibilityStatus` | `string` | yes |
| `blockedInputMessaging` | `string` | yes |
| `blockedOutputsMessaging` | `string` | yes |
| `description` | `string` | no |
| `topicPolicyConfig` | `AIGuardrailTopicPolicyConfig` | no |
| `contentPolicyConfig` | `AIGuardrailContentPolicyConfig` | no |
| `wordPolicyConfig` | `AIGuardrailWordPolicyConfig` | no |
| `sensitiveInformationPolicyConfig` | `AIGuardrailSensitiveInformationPolicyConfig` | no |
| `contextualGroundingPolicyConfig` | `AIGuardrailContextualGroundingPolicyConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiGuardrail` | `AIGuardrailData` | no |

## UpdateAIPrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `assistantId` | `string` | yes |
| `aiPromptId` | `string` | yes |
| `visibilityStatus` | `string` | yes |
| `templateConfiguration` | `AIPromptTemplateConfiguration` | no |
| `description` | `string` | no |
| `modelId` | `string` | no |
| `inferenceConfiguration` | `AIPromptInferenceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aiPrompt` | `AIPromptData` | no |

## UpdateAssistantAIAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `aiAgentType` | `string` | yes |
| `configuration` | `AIAgentConfigurationData` | yes |
| `orchestratorUseCase` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistant` | `AssistantData` | no |

## UpdateContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `contentId` | `string` | yes |
| `revisionId` | `string` | no |
| `title` | `string` | no |
| `overrideLinkOutUri` | `string` | no |
| `removeOverrideLinkOutUri` | `boolean` | no |
| `metadata` | `Map<string>` | no |
| `uploadId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `ContentData` | no |

## UpdateKnowledgeBaseTemplateUri

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `templateUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBaseData` | no |

## UpdateMessageTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `content` | `MessageTemplateContentProvider` | no |
| `language` | `string` | no |
| `sourceConfiguration` | `MessageTemplateSourceConfiguration` | no |
| `defaultAttributes` | `MessageTemplateAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplate` | `MessageTemplateData` | no |

## UpdateMessageTemplateMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `messageTemplateId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messageTemplate` | `MessageTemplateData` | no |

## UpdateQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `quickResponseId` | `string` | yes |
| `name` | `string` | no |
| `content` | `QuickResponseDataProvider` | no |
| `contentType` | `string` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |
| `removeGroupingConfiguration` | `boolean` | no |
| `description` | `string` | no |
| `removeDescription` | `boolean` | no |
| `shortcutKey` | `string` | no |
| `removeShortcutKey` | `boolean` | no |
| `isActive` | `boolean` | no |
| `channels` | `List<string>` | no |
| `language` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

## UpdateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `description` | `string` | no |
| `tagFilter` | `TagFilter` | no |
| `aiAgentConfiguration` | `Map<AIAgentConfigurationData>` | no |
| `orchestratorConfigurationList` | `List<OrchestratorConfigurationEntry>` | no |
| `removeOrchestratorConfigurationList` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `SessionData` | no |

## UpdateSessionData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `sessionId` | `string` | yes |
| `namespace` | `string` | no |
| `data` | `List<RuntimeSessionData>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionArn` | `string` | yes |
| `sessionId` | `string` | yes |
| `namespace` | `string` | yes |
| `data` | `List<RuntimeSessionData>` | yes |

