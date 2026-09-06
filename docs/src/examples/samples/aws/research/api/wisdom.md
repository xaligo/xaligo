# Amazon Connect Wisdom Service

API version: 2020-10-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/wisdom/2020-10-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `name` | `string` | yes |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `tags` | `Map<string>` | no |
| `type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistant` | `AssistantData` | no |

## CreateAssistantAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `association` | `AssistantAssociationInputData` | yes |
| `associationType` | `string` | yes |
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
| `clientToken` | `string` | no |
| `knowledgeBaseId` | `string` | yes |
| `metadata` | `Map<string>` | no |
| `name` | `string` | yes |
| `overrideLinkOutUri` | `string` | no |
| `tags` | `Map<string>` | no |
| `title` | `string` | no |
| `uploadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `ContentData` | no |

## CreateKnowledgeBase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `knowledgeBaseType` | `string` | yes |
| `name` | `string` | yes |
| `renderingConfiguration` | `RenderingConfiguration` | no |
| `serverSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | no |
| `sourceConfiguration` | `SourceConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBase` | `KnowledgeBaseData` | no |

## CreateQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channels` | `List<string>` | no |
| `clientToken` | `string` | no |
| `content` | `QuickResponseDataProvider` | yes |
| `contentType` | `string` | no |
| `description` | `string` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |
| `isActive` | `boolean` | no |
| `knowledgeBaseId` | `string` | yes |
| `language` | `string` | no |
| `name` | `string` | yes |
| `shortcutKey` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

## CreateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `SessionData` | no |

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
| `contentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJobId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |

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


## DeleteQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `quickResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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

## GetQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `quickResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

## GetRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `maxResults` | `integer` | no |
| `sessionId` | `string` | yes |
| `waitTimeSeconds` | `integer` | no |

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

## ListAssistantAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantAssociationSummaries` | `List<AssistantAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssistants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantSummaries` | `List<AssistantSummary>` | yes |
| `nextToken` | `string` | no |

## ListContents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentSummaries` | `List<ContentSummary>` | yes |
| `nextToken` | `string` | no |

## ListImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importJobSummaries` | `List<ImportJobSummary>` | yes |
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

## ListQuickResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `quickResponseSummaries` | `List<QuickResponseSummary>` | yes |

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
| `recommendationIds` | `List<string>` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<NotifyRecommendationsReceivedError>` | no |
| `recommendationIds` | `List<string>` | no |

## QueryAssistant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `queryText` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `results` | `List<ResultData>` | yes |

## RemoveKnowledgeBaseTemplateUri

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchExpression` | `SearchExpression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentSummaries` | `List<ContentSummary>` | yes |
| `nextToken` | `string` | no |

## SearchQuickResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `Map<string>` | no |
| `knowledgeBaseId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchExpression` | `QuickResponseSearchExpression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `results` | `List<QuickResponseSearchResultData>` | yes |

## SearchSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assistantId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchExpression` | `SearchExpression` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sessionSummaries` | `List<SessionSummary>` | yes |

## StartContentUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `presignedUrlTimeToLive` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `headersToInclude` | `Map<string>` | yes |
| `uploadId` | `string` | yes |
| `url` | `string` | yes |
| `urlExpiry` | `timestamp` | yes |

## StartImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `externalSourceConfiguration` | `ExternalSourceConfiguration` | no |
| `importJobType` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `metadata` | `Map<string>` | no |
| `uploadId` | `string` | yes |

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


## UpdateContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentId` | `string` | yes |
| `knowledgeBaseId` | `string` | yes |
| `metadata` | `Map<string>` | no |
| `overrideLinkOutUri` | `string` | no |
| `removeOverrideLinkOutUri` | `boolean` | no |
| `revisionId` | `string` | no |
| `title` | `string` | no |
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

## UpdateQuickResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channels` | `List<string>` | no |
| `content` | `QuickResponseDataProvider` | no |
| `contentType` | `string` | no |
| `description` | `string` | no |
| `groupingConfiguration` | `GroupingConfiguration` | no |
| `isActive` | `boolean` | no |
| `knowledgeBaseId` | `string` | yes |
| `language` | `string` | no |
| `name` | `string` | no |
| `quickResponseId` | `string` | yes |
| `removeDescription` | `boolean` | no |
| `removeGroupingConfiguration` | `boolean` | no |
| `removeShortcutKey` | `boolean` | no |
| `shortcutKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quickResponse` | `QuickResponseData` | no |

