# QBusiness

API version: 2023-11-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/qbusiness/2023-11-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociatePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `statementId` | `string` | yes |
| `actions` | `List<string>` | yes |
| `conditions` | `List<PermissionCondition>` | no |
| `principal` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statement` | `string` | no |

## BatchDeleteDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `documents` | `List<DeleteDocument>` | yes |
| `dataSourceSyncId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedDocuments` | `List<FailedDocument>` | no |

## BatchPutDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `documents` | `List<Document>` | yes |
| `roleArn` | `string` | no |
| `dataSourceSyncId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failedDocuments` | `List<FailedDocument>` | no |

## CancelSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `subscriptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionArn` | `string` | no |
| `currentSubscription` | `SubscriptionDetails` | no |
| `nextSubscription` | `SubscriptionDetails` | no |

## Chat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | no |
| `userGroups` | `List<string>` | no |
| `conversationId` | `string` | no |
| `parentMessageId` | `string` | no |
| `clientToken` | `string` | no |
| `inputStream` | `ChatInputStream` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `outputStream` | `ChatOutputStream` | no |

## ChatSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | no |
| `userGroups` | `List<string>` | no |
| `userMessage` | `string` | no |
| `attachments` | `List<AttachmentInput>` | no |
| `actionExecution` | `ActionExecution` | no |
| `authChallengeResponse` | `AuthChallengeResponse` | no |
| `conversationId` | `string` | no |
| `parentMessageId` | `string` | no |
| `attributeFilter` | `AttributeFilter` | no |
| `chatMode` | `string` | no |
| `chatModeConfiguration` | `ChatModeConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conversationId` | `string` | no |
| `systemMessage` | `string` | no |
| `systemMessageId` | `string` | no |
| `userMessageId` | `string` | no |
| `actionReview` | `ActionReview` | no |
| `authChallengeRequest` | `AuthChallengeRequest` | no |
| `sourceAttributions` | `List<SourceAttribution>` | no |
| `failedAttachments` | `List<AttachmentOutput>` | no |

## CheckDocumentAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `userId` | `string` | yes |
| `documentId` | `string` | yes |
| `dataSourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userGroups` | `List<AssociatedGroup>` | no |
| `userAliases` | `List<AssociatedUser>` | no |
| `hasAccess` | `boolean` | no |
| `documentAcl` | `DocumentAcl` | no |

## CreateAnonymousWebExperienceUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `webExperienceId` | `string` | yes |
| `sessionDurationInMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anonymousUrl` | `string` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | yes |
| `roleArn` | `string` | no |
| `identityType` | `string` | no |
| `iamIdentityProviderArn` | `string` | no |
| `identityCenterInstanceArn` | `string` | no |
| `clientIdsForOIDC` | `List<string>` | no |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |
| `attachmentsConfiguration` | `AttachmentsConfiguration` | no |
| `qAppsConfiguration` | `QAppsConfiguration` | no |
| `personalizationConfiguration` | `PersonalizationConfiguration` | no |
| `quickSightConfiguration` | `QuickSightConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `applicationArn` | `string` | no |

## CreateChatResponseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `displayName` | `string` | yes |
| `clientToken` | `string` | no |
| `responseConfigurations` | `Map<ResponseConfiguration>` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `chatResponseConfigurationId` | `string` | yes |
| `chatResponseConfigurationArn` | `string` | yes |

## CreateDataAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `principal` | `string` | yes |
| `actionConfigurations` | `List<ActionConfiguration>` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `authenticationDetail` | `DataAccessorAuthenticationDetail` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataAccessorId` | `string` | yes |
| `idcApplicationArn` | `string` | yes |
| `dataAccessorArn` | `string` | yes |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `displayName` | `string` | yes |
| `configuration` | `DataSourceConfiguration` | yes |
| `vpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `syncSchedule` | `string` | no |
| `roleArn` | `string` | no |
| `clientToken` | `string` | no |
| `documentEnrichmentConfiguration` | `DocumentEnrichmentConfiguration` | no |
| `mediaExtractionConfiguration` | `MediaExtractionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | no |
| `dataSourceArn` | `string` | no |

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `displayName` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | no |
| `tags` | `List<Tag>` | no |
| `capacityConfiguration` | `IndexCapacityConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexId` | `string` | no |
| `indexArn` | `string` | no |

## CreatePlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `displayName` | `string` | yes |
| `type` | `string` | yes |
| `authConfiguration` | `PluginAuthConfiguration` | yes |
| `serverUrl` | `string` | no |
| `customPluginConfiguration` | `CustomPluginConfiguration` | no |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pluginId` | `string` | no |
| `pluginArn` | `string` | no |
| `buildStatus` | `string` | no |

## CreateRetriever

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `type` | `string` | yes |
| `displayName` | `string` | yes |
| `configuration` | `RetrieverConfiguration` | yes |
| `roleArn` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `retrieverId` | `string` | no |
| `retrieverArn` | `string` | no |

## CreateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `principal` | `SubscriptionPrincipal` | yes |
| `type` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionId` | `string` | no |
| `subscriptionArn` | `string` | no |
| `currentSubscription` | `SubscriptionDetails` | no |
| `nextSubscription` | `SubscriptionDetails` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | yes |
| `userAliases` | `List<UserAlias>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateWebExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `title` | `string` | no |
| `subtitle` | `string` | no |
| `welcomeMessage` | `string` | no |
| `samplePromptsControlMode` | `string` | no |
| `origins` | `List<string>` | no |
| `roleArn` | `string` | no |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |
| `identityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `browserExtensionConfiguration` | `BrowserExtensionConfiguration` | no |
| `customizationConfiguration` | `CustomizationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webExperienceId` | `string` | no |
| `webExperienceArn` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `conversationId` | `string` | yes |
| `attachmentId` | `string` | yes |
| `userId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChatControlsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChatResponseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `chatResponseConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConversation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conversationId` | `string` | yes |
| `applicationId` | `string` | yes |
| `userId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `dataAccessorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `dataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `groupName` | `string` | yes |
| `dataSourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `pluginId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRetriever

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `retrieverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWebExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `webExperienceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociatePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `statementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `applicationId` | `string` | no |
| `applicationArn` | `string` | no |
| `identityType` | `string` | no |
| `iamIdentityProviderArn` | `string` | no |
| `identityCenterApplicationArn` | `string` | no |
| `roleArn` | `string` | no |
| `status` | `string` | no |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `error` | `ErrorDetail` | no |
| `attachmentsConfiguration` | `AppliedAttachmentsConfiguration` | no |
| `qAppsConfiguration` | `QAppsConfiguration` | no |
| `personalizationConfiguration` | `PersonalizationConfiguration` | no |
| `autoSubscriptionConfiguration` | `AutoSubscriptionConfiguration` | no |
| `clientIdsForOIDC` | `List<string>` | no |
| `quickSightConfiguration` | `QuickSightConfiguration` | no |

## GetChatControlsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseScope` | `string` | no |
| `orchestrationConfiguration` | `AppliedOrchestrationConfiguration` | no |
| `blockedPhrases` | `BlockedPhrasesConfiguration` | no |
| `topicConfigurations` | `List<TopicConfiguration>` | no |
| `creatorModeConfiguration` | `AppliedCreatorModeConfiguration` | no |
| `nextToken` | `string` | no |
| `hallucinationReductionConfiguration` | `HallucinationReductionConfiguration` | no |

## GetChatResponseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `chatResponseConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `chatResponseConfigurationId` | `string` | no |
| `chatResponseConfigurationArn` | `string` | no |
| `displayName` | `string` | no |
| `createdAt` | `timestamp` | no |
| `inUseConfiguration` | `ChatResponseConfigurationDetail` | no |
| `lastUpdateConfiguration` | `ChatResponseConfigurationDetail` | no |

## GetDataAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `dataAccessorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `dataAccessorId` | `string` | no |
| `dataAccessorArn` | `string` | no |
| `applicationId` | `string` | no |
| `idcApplicationArn` | `string` | no |
| `principal` | `string` | no |
| `actionConfigurations` | `List<ActionConfiguration>` | no |
| `authenticationDetail` | `DataAccessorAuthenticationDetail` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `dataSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `indexId` | `string` | no |
| `dataSourceId` | `string` | no |
| `dataSourceArn` | `string` | no |
| `displayName` | `string` | no |
| `type` | `string` | no |
| `configuration` | `DataSourceConfiguration` | no |
| `vpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `syncSchedule` | `string` | no |
| `roleArn` | `string` | no |
| `error` | `ErrorDetail` | no |
| `documentEnrichmentConfiguration` | `DocumentEnrichmentConfiguration` | no |
| `mediaExtractionConfiguration` | `MediaExtractionConfiguration` | no |

## GetDocumentContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `dataSourceId` | `string` | no |
| `documentId` | `string` | yes |
| `outputFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `presignedUrl` | `string` | yes |
| `mimeType` | `string` | yes |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `groupName` | `string` | yes |
| `dataSourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `GroupStatusDetail` | no |
| `statusHistory` | `List<GroupStatusDetail>` | no |

## GetIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `indexId` | `string` | no |
| `displayName` | `string` | no |
| `indexArn` | `string` | no |
| `status` | `string` | no |
| `type` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `capacityConfiguration` | `IndexCapacityConfiguration` | no |
| `documentAttributeConfigurations` | `List<DocumentAttributeConfiguration>` | no |
| `error` | `ErrorDetail` | no |
| `indexStatistics` | `IndexStatistics` | no |

## GetMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `conversationId` | `string` | yes |
| `messageId` | `string` | yes |
| `mediaId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mediaBytes` | `blob` | no |
| `mediaMimeType` | `string` | no |

## GetPlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `pluginId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `pluginId` | `string` | no |
| `displayName` | `string` | no |
| `type` | `string` | no |
| `serverUrl` | `string` | no |
| `authConfiguration` | `PluginAuthConfiguration` | no |
| `customPluginConfiguration` | `CustomPluginConfiguration` | no |
| `buildStatus` | `string` | no |
| `pluginArn` | `string` | no |
| `state` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |

## GetRetriever

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `retrieverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `retrieverId` | `string` | no |
| `retrieverArn` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `displayName` | `string` | no |
| `configuration` | `RetrieverConfiguration` | no |
| `roleArn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAliases` | `List<UserAlias>` | no |

## GetWebExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `webExperienceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |
| `webExperienceId` | `string` | no |
| `webExperienceArn` | `string` | no |
| `defaultEndpoint` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `title` | `string` | no |
| `subtitle` | `string` | no |
| `welcomeMessage` | `string` | no |
| `samplePromptsControlMode` | `string` | no |
| `origins` | `List<string>` | no |
| `roleArn` | `string` | no |
| `identityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `authenticationConfiguration` | `WebExperienceAuthConfiguration` | no |
| `error` | `ErrorDetail` | no |
| `browserExtensionConfiguration` | `BrowserExtensionConfiguration` | no |
| `customizationConfiguration` | `CustomizationConfiguration` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `applications` | `List<Application>` | no |

## ListAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `conversationId` | `string` | no |
| `userId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachments` | `List<Attachment>` | no |
| `nextToken` | `string` | no |

## ListChatResponseConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `chatResponseConfigurations` | `List<ChatResponseConfiguration>` | no |
| `nextToken` | `string` | no |

## ListConversations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `conversations` | `List<Conversation>` | no |

## ListDataAccessors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataAccessors` | `List<DataAccessor>` | no |
| `nextToken` | `string` | no |

## ListDataSourceSyncJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `statusFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `history` | `List<DataSourceSyncJob>` | no |
| `nextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSources` | `List<DataSource>` | no |
| `nextToken` | `string` | no |

## ListDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `dataSourceIds` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documentDetailList` | `List<DocumentDetails>` | no |
| `nextToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `updatedEarlierThan` | `timestamp` | yes |
| `dataSourceId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<GroupSummary>` | no |

## ListIndices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `indices` | `List<Index>` | no |

## ListMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `conversationId` | `string` | yes |
| `applicationId` | `string` | yes |
| `userId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messages` | `List<Message>` | no |
| `nextToken` | `string` | no |

## ListPluginActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `pluginId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ActionSummary>` | no |

## ListPluginTypeActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pluginType` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ActionSummary>` | no |

## ListPluginTypeMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<PluginTypeMetadataSummary>` | no |

## ListPlugins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `plugins` | `List<Plugin>` | no |

## ListRetrievers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `retrievers` | `List<Retriever>` | no |
| `nextToken` | `string` | no |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `subscriptions` | `List<Subscription>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListWebExperiences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webExperiences` | `List<WebExperience>` | no |
| `nextToken` | `string` | no |

## PutFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | no |
| `conversationId` | `string` | yes |
| `messageId` | `string` | yes |
| `messageCopiedAt` | `timestamp` | no |
| `messageUsefulness` | `MessageUsefulnessFeedback` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `groupName` | `string` | yes |
| `dataSourceId` | `string` | no |
| `type` | `string` | yes |
| `groupMembers` | `GroupMembers` | yes |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchRelevantContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `queryText` | `string` | yes |
| `contentSource` | `ContentSource` | yes |
| `attributeFilter` | `AttributeFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `relevantContent` | `List<RelevantContent>` | no |
| `nextToken` | `string` | no |

## StartDataSourceSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |

## StopDataSourceSyncJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceId` | `string` | yes |
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |

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


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `identityCenterInstanceArn` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `attachmentsConfiguration` | `AttachmentsConfiguration` | no |
| `qAppsConfiguration` | `QAppsConfiguration` | no |
| `personalizationConfiguration` | `PersonalizationConfiguration` | no |
| `autoSubscriptionConfiguration` | `AutoSubscriptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateChatControlsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `clientToken` | `string` | no |
| `responseScope` | `string` | no |
| `orchestrationConfiguration` | `OrchestrationConfiguration` | no |
| `blockedPhrasesConfigurationUpdate` | `BlockedPhrasesConfigurationUpdate` | no |
| `topicConfigurationsToCreateOrUpdate` | `List<TopicConfiguration>` | no |
| `topicConfigurationsToDelete` | `List<TopicConfiguration>` | no |
| `creatorModeConfiguration` | `CreatorModeConfiguration` | no |
| `hallucinationReductionConfiguration` | `HallucinationReductionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateChatResponseConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `chatResponseConfigurationId` | `string` | yes |
| `displayName` | `string` | no |
| `responseConfigurations` | `Map<ResponseConfiguration>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataAccessor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `dataAccessorId` | `string` | yes |
| `actionConfigurations` | `List<ActionConfiguration>` | yes |
| `authenticationDetail` | `DataAccessorAuthenticationDetail` | no |
| `displayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `displayName` | `string` | no |
| `configuration` | `DataSourceConfiguration` | no |
| `vpcConfiguration` | `DataSourceVpcConfiguration` | no |
| `description` | `string` | no |
| `syncSchedule` | `string` | no |
| `roleArn` | `string` | no |
| `documentEnrichmentConfiguration` | `DocumentEnrichmentConfiguration` | no |
| `mediaExtractionConfiguration` | `MediaExtractionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `indexId` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `capacityConfiguration` | `IndexCapacityConfiguration` | no |
| `documentAttributeConfigurations` | `List<DocumentAttributeConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `pluginId` | `string` | yes |
| `displayName` | `string` | no |
| `state` | `string` | no |
| `serverUrl` | `string` | no |
| `customPluginConfiguration` | `CustomPluginConfiguration` | no |
| `authConfiguration` | `PluginAuthConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRetriever

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `retrieverId` | `string` | yes |
| `configuration` | `RetrieverConfiguration` | no |
| `displayName` | `string` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `subscriptionId` | `string` | yes |
| `type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionArn` | `string` | no |
| `currentSubscription` | `SubscriptionDetails` | no |
| `nextSubscription` | `SubscriptionDetails` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `userId` | `string` | yes |
| `userAliasesToUpdate` | `List<UserAlias>` | no |
| `userAliasesToDelete` | `List<UserAlias>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAliasesAdded` | `List<UserAlias>` | no |
| `userAliasesUpdated` | `List<UserAlias>` | no |
| `userAliasesDeleted` | `List<UserAlias>` | no |

## UpdateWebExperience

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `webExperienceId` | `string` | yes |
| `roleArn` | `string` | no |
| `authenticationConfiguration` | `WebExperienceAuthConfiguration` | no |
| `title` | `string` | no |
| `subtitle` | `string` | no |
| `welcomeMessage` | `string` | no |
| `samplePromptsControlMode` | `string` | no |
| `identityProviderConfiguration` | `IdentityProviderConfiguration` | no |
| `origins` | `List<string>` | no |
| `browserExtensionConfiguration` | `BrowserExtensionConfiguration` | no |
| `customizationConfiguration` | `CustomizationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


