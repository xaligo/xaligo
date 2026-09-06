# Amazon Lex Model Building Service

API version: 2017-04-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lex-models/2017-04-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBotVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `checksum` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `intents` | `List<Intent>` | no |
| `clarificationPrompt` | `Prompt` | no |
| `abortStatement` | `Statement` | no |
| `status` | `string` | no |
| `failureReason` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `voiceId` | `string` | no |
| `checksum` | `string` | no |
| `version` | `string` | no |
| `locale` | `string` | no |
| `childDirected` | `boolean` | no |
| `enableModelImprovements` | `boolean` | no |
| `detectSentiment` | `boolean` | no |

## CreateIntentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `checksum` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `slots` | `List<Slot>` | no |
| `sampleUtterances` | `List<string>` | no |
| `confirmationPrompt` | `Prompt` | no |
| `rejectionStatement` | `Statement` | no |
| `followUpPrompt` | `FollowUpPrompt` | no |
| `conclusionStatement` | `Statement` | no |
| `dialogCodeHook` | `CodeHook` | no |
| `fulfillmentActivity` | `FulfillmentActivity` | no |
| `parentIntentSignature` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |

## CreateSlotTypeVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `checksum` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `enumerationValues` | `List<EnumerationValue>` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `valueSelectionStrategy` | `string` | no |
| `parentSlotTypeSignature` | `string` | no |
| `slotTypeConfigurations` | `List<SlotTypeConfiguration>` | no |

## DeleteBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `botName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBotChannelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBotVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntentVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlotTypeVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUtterances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `versionOrAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `intents` | `List<Intent>` | no |
| `enableModelImprovements` | `boolean` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `clarificationPrompt` | `Prompt` | no |
| `abortStatement` | `Statement` | no |
| `status` | `string` | no |
| `failureReason` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `voiceId` | `string` | no |
| `checksum` | `string` | no |
| `version` | `string` | no |
| `locale` | `string` | no |
| `childDirected` | `boolean` | no |
| `detectSentiment` | `boolean` | no |

## GetBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `botName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botName` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `checksum` | `string` | no |
| `conversationLogs` | `ConversationLogsResponse` | no |

## GetBotAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `nameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BotAliases` | `List<BotAliasMetadata>` | no |
| `nextToken` | `string` | no |

## GetBotChannelAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `botAlias` | `string` | no |
| `botName` | `string` | no |
| `createdDate` | `timestamp` | no |
| `type` | `string` | no |
| `botConfiguration` | `Map<string>` | no |
| `status` | `string` | no |
| `failureReason` | `string` | no |

## GetBotChannelAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `nameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botChannelAssociations` | `List<BotChannelAssociation>` | no |
| `nextToken` | `string` | no |

## GetBotVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bots` | `List<BotMetadata>` | no |
| `nextToken` | `string` | no |

## GetBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `nameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bots` | `List<BotMetadata>` | no |
| `nextToken` | `string` | no |

## GetBuiltinIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signature` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signature` | `string` | no |
| `supportedLocales` | `List<string>` | no |
| `slots` | `List<BuiltinIntentSlot>` | no |

## GetBuiltinIntents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `locale` | `string` | no |
| `signatureContains` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intents` | `List<BuiltinIntentMetadata>` | no |
| `nextToken` | `string` | no |

## GetBuiltinSlotTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `locale` | `string` | no |
| `signatureContains` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypes` | `List<BuiltinSlotTypeMetadata>` | no |
| `nextToken` | `string` | no |

## GetExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |
| `resourceType` | `string` | yes |
| `exportType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `version` | `string` | no |
| `resourceType` | `string` | no |
| `exportType` | `string` | no |
| `exportStatus` | `string` | no |
| `failureReason` | `string` | no |
| `url` | `string` | no |

## GetImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `resourceType` | `string` | no |
| `mergeStrategy` | `string` | no |
| `importId` | `string` | no |
| `importStatus` | `string` | no |
| `failureReason` | `List<string>` | no |
| `createdDate` | `timestamp` | no |

## GetIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `slots` | `List<Slot>` | no |
| `sampleUtterances` | `List<string>` | no |
| `confirmationPrompt` | `Prompt` | no |
| `rejectionStatement` | `Statement` | no |
| `followUpPrompt` | `FollowUpPrompt` | no |
| `conclusionStatement` | `Statement` | no |
| `dialogCodeHook` | `CodeHook` | no |
| `fulfillmentActivity` | `FulfillmentActivity` | no |
| `parentIntentSignature` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |

## GetIntentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intents` | `List<IntentMetadata>` | no |
| `nextToken` | `string` | no |

## GetIntents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `nameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intents` | `List<IntentMetadata>` | no |
| `nextToken` | `string` | no |

## GetMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationId` | `string` | no |
| `v1BotName` | `string` | no |
| `v1BotVersion` | `string` | no |
| `v1BotLocale` | `string` | no |
| `v2BotId` | `string` | no |
| `v2BotRole` | `string` | no |
| `migrationStatus` | `string` | no |
| `migrationStrategy` | `string` | no |
| `migrationTimestamp` | `timestamp` | no |
| `alerts` | `List<MigrationAlert>` | no |

## GetMigrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortByAttribute` | `string` | no |
| `sortByOrder` | `string` | no |
| `v1BotNameContains` | `string` | no |
| `migrationStatusEquals` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationSummaries` | `List<MigrationSummary>` | no |
| `nextToken` | `string` | no |

## GetSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `enumerationValues` | `List<EnumerationValue>` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `valueSelectionStrategy` | `string` | no |
| `parentSlotTypeSignature` | `string` | no |
| `slotTypeConfigurations` | `List<SlotTypeConfiguration>` | no |

## GetSlotTypeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypes` | `List<SlotTypeMetadata>` | no |
| `nextToken` | `string` | no |

## GetSlotTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `nameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypes` | `List<SlotTypeMetadata>` | no |
| `nextToken` | `string` | no |

## GetUtterancesView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botVersions` | `List<string>` | yes |
| `statusType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | no |
| `utterances` | `List<UtteranceList>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## PutBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `intents` | `List<Intent>` | no |
| `enableModelImprovements` | `boolean` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `clarificationPrompt` | `Prompt` | no |
| `abortStatement` | `Statement` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `voiceId` | `string` | no |
| `checksum` | `string` | no |
| `processBehavior` | `string` | no |
| `locale` | `string` | yes |
| `childDirected` | `boolean` | yes |
| `detectSentiment` | `boolean` | no |
| `createVersion` | `boolean` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `intents` | `List<Intent>` | no |
| `enableModelImprovements` | `boolean` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `clarificationPrompt` | `Prompt` | no |
| `abortStatement` | `Statement` | no |
| `status` | `string` | no |
| `failureReason` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `voiceId` | `string` | no |
| `checksum` | `string` | no |
| `version` | `string` | no |
| `locale` | `string` | no |
| `childDirected` | `boolean` | no |
| `createVersion` | `boolean` | no |
| `detectSentiment` | `boolean` | no |
| `tags` | `List<Tag>` | no |

## PutBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `botVersion` | `string` | yes |
| `botName` | `string` | yes |
| `checksum` | `string` | no |
| `conversationLogs` | `ConversationLogsRequest` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botName` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `checksum` | `string` | no |
| `conversationLogs` | `ConversationLogsResponse` | no |
| `tags` | `List<Tag>` | no |

## PutIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `slots` | `List<Slot>` | no |
| `sampleUtterances` | `List<string>` | no |
| `confirmationPrompt` | `Prompt` | no |
| `rejectionStatement` | `Statement` | no |
| `followUpPrompt` | `FollowUpPrompt` | no |
| `conclusionStatement` | `Statement` | no |
| `dialogCodeHook` | `CodeHook` | no |
| `fulfillmentActivity` | `FulfillmentActivity` | no |
| `parentIntentSignature` | `string` | no |
| `checksum` | `string` | no |
| `createVersion` | `boolean` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `slots` | `List<Slot>` | no |
| `sampleUtterances` | `List<string>` | no |
| `confirmationPrompt` | `Prompt` | no |
| `rejectionStatement` | `Statement` | no |
| `followUpPrompt` | `FollowUpPrompt` | no |
| `conclusionStatement` | `Statement` | no |
| `dialogCodeHook` | `CodeHook` | no |
| `fulfillmentActivity` | `FulfillmentActivity` | no |
| `parentIntentSignature` | `string` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `createVersion` | `boolean` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |

## PutSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `enumerationValues` | `List<EnumerationValue>` | no |
| `checksum` | `string` | no |
| `valueSelectionStrategy` | `string` | no |
| `createVersion` | `boolean` | no |
| `parentSlotTypeSignature` | `string` | no |
| `slotTypeConfigurations` | `List<SlotTypeConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `enumerationValues` | `List<EnumerationValue>` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `checksum` | `string` | no |
| `valueSelectionStrategy` | `string` | no |
| `createVersion` | `boolean` | no |
| `parentSlotTypeSignature` | `string` | no |
| `slotTypeConfigurations` | `List<SlotTypeConfiguration>` | no |

## StartImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | yes |
| `resourceType` | `string` | yes |
| `mergeStrategy` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `resourceType` | `string` | no |
| `mergeStrategy` | `string` | no |
| `importId` | `string` | no |
| `importStatus` | `string` | no |
| `tags` | `List<Tag>` | no |
| `createdDate` | `timestamp` | no |

## StartMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `v1BotName` | `string` | yes |
| `v1BotVersion` | `string` | yes |
| `v2BotName` | `string` | yes |
| `v2BotRole` | `string` | yes |
| `migrationStrategy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `v1BotName` | `string` | no |
| `v1BotVersion` | `string` | no |
| `v1BotLocale` | `string` | no |
| `v2BotId` | `string` | no |
| `v2BotRole` | `string` | no |
| `migrationId` | `string` | no |
| `migrationStrategy` | `string` | no |
| `migrationTimestamp` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


