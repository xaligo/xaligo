# Amazon Lex Model Building V2

API version: 2020-08-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lexv2-models/2020-08-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateCustomVocabularyItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `customVocabularyItemList` | `List<NewCustomVocabularyItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `errors` | `List<FailedCustomVocabularyItem>` | no |
| `resources` | `List<CustomVocabularyItem>` | no |

## BatchDeleteCustomVocabularyItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `customVocabularyItemList` | `List<CustomVocabularyEntryId>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `errors` | `List<FailedCustomVocabularyItem>` | no |
| `resources` | `List<CustomVocabularyItem>` | no |

## BatchUpdateCustomVocabularyItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `customVocabularyItemList` | `List<CustomVocabularyItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `errors` | `List<FailedCustomVocabularyItem>` | no |
| `resources` | `List<CustomVocabularyItem>` | no |

## BuildBotLocale

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botLocaleStatus` | `string` | no |
| `lastBuildSubmittedDateTime` | `timestamp` | no |

## CreateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `dataPrivacy` | `DataPrivacy` | yes |
| `idleSessionTTLInSeconds` | `integer` | yes |
| `botTags` | `Map<string>` | no |
| `testBotAliasTags` | `Map<string>` | no |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |
| `errorLogSettings` | `ErrorLogSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botName` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `dataPrivacy` | `DataPrivacy` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `botStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `botTags` | `Map<string>` | no |
| `testBotAliasTags` | `Map<string>` | no |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |
| `errorLogSettings` | `ErrorLogSettings` | no |

## CreateBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasName` | `string` | yes |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botAliasLocaleSettings` | `Map<BotAliasLocaleSettings>` | no |
| `conversationLogSettings` | `ConversationLogSettings` | no |
| `sentimentAnalysisSettings` | `SentimentAnalysisSettings` | no |
| `botId` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | no |
| `botAliasName` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botAliasLocaleSettings` | `Map<BotAliasLocaleSettings>` | no |
| `conversationLogSettings` | `ConversationLogSettings` | no |
| `sentimentAnalysisSettings` | `SentimentAnalysisSettings` | no |
| `botAliasStatus` | `string` | no |
| `botId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## CreateBotLocale

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `description` | `string` | no |
| `nluIntentConfidenceThreshold` | `double` | yes |
| `voiceSettings` | `VoiceSettings` | no |
| `unifiedSpeechSettings` | `UnifiedSpeechSettings` | no |
| `audioFillerSettings` | `AudioFillerSettings` | no |
| `speechRecognitionSettings` | `SpeechRecognitionSettings` | no |
| `generativeAISettings` | `GenerativeAISettings` | no |
| `speechDetectionSensitivity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeName` | `string` | no |
| `localeId` | `string` | no |
| `description` | `string` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `voiceSettings` | `VoiceSettings` | no |
| `unifiedSpeechSettings` | `UnifiedSpeechSettings` | no |
| `audioFillerSettings` | `AudioFillerSettings` | no |
| `speechRecognitionSettings` | `SpeechRecognitionSettings` | no |
| `botLocaleStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `generativeAISettings` | `GenerativeAISettings` | no |
| `speechDetectionSensitivity` | `string` | no |

## CreateBotReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `replicaRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `replicaRegion` | `string` | no |
| `sourceRegion` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `botReplicaStatus` | `string` | no |

## CreateBotVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `description` | `string` | no |
| `botVersionLocaleSpecification` | `Map<BotVersionLocaleDetails>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botVersionLocaleSpecification` | `Map<BotVersionLocaleDetails>` | no |
| `botStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |

## CreateExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceSpecification` | `ExportResourceSpecification` | yes |
| `fileFormat` | `string` | yes |
| `filePassword` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |
| `resourceSpecification` | `ExportResourceSpecification` | no |
| `fileFormat` | `string` | no |
| `exportStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |

## CreateIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentName` | `string` | yes |
| `intentDisplayName` | `string` | no |
| `description` | `string` | no |
| `parentIntentSignature` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |
| `dialogCodeHook` | `DialogCodeHookSettings` | no |
| `fulfillmentCodeHook` | `FulfillmentCodeHookSettings` | no |
| `intentConfirmationSetting` | `IntentConfirmationSetting` | no |
| `intentClosingSetting` | `IntentClosingSetting` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `initialResponseSetting` | `InitialResponseSetting` | no |
| `qnAIntentConfiguration` | `QnAIntentConfiguration` | no |
| `qInConnectIntentConfiguration` | `QInConnectIntentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | no |
| `intentName` | `string` | no |
| `intentDisplayName` | `string` | no |
| `description` | `string` | no |
| `parentIntentSignature` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |
| `dialogCodeHook` | `DialogCodeHookSettings` | no |
| `fulfillmentCodeHook` | `FulfillmentCodeHookSettings` | no |
| `intentConfirmationSetting` | `IntentConfirmationSetting` | no |
| `intentClosingSetting` | `IntentClosingSetting` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `initialResponseSetting` | `InitialResponseSetting` | no |
| `qnAIntentConfiguration` | `QnAIntentConfiguration` | no |
| `qInConnectIntentConfiguration` | `QInConnectIntentConfiguration` | no |

## CreateResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `revisionId` | `string` | no |

## CreateResourcePolicyStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `statementId` | `string` | yes |
| `effect` | `string` | yes |
| `principal` | `List<Principal>` | yes |
| `action` | `List<string>` | yes |
| `condition` | `Map<Map<string>>` | no |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `revisionId` | `string` | no |

## CreateSlot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotName` | `string` | yes |
| `description` | `string` | no |
| `slotTypeId` | `string` | no |
| `valueElicitationSetting` | `SlotValueElicitationSetting` | yes |
| `obfuscationSetting` | `ObfuscationSetting` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `intentId` | `string` | yes |
| `multipleValuesSetting` | `MultipleValuesSetting` | no |
| `subSlotSetting` | `SubSlotSetting` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | no |
| `slotName` | `string` | no |
| `description` | `string` | no |
| `slotTypeId` | `string` | no |
| `valueElicitationSetting` | `SlotValueElicitationSetting` | no |
| `obfuscationSetting` | `ObfuscationSetting` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `multipleValuesSetting` | `MultipleValuesSetting` | no |
| `subSlotSetting` | `SubSlotSetting` | no |

## CreateSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeName` | `string` | yes |
| `description` | `string` | no |
| `slotTypeValues` | `List<SlotTypeValue>` | no |
| `valueSelectionSetting` | `SlotValueSelectionSetting` | no |
| `parentSlotTypeSignature` | `string` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `externalSourceSetting` | `ExternalSourceSetting` | no |
| `compositeSlotTypeSetting` | `CompositeSlotTypeSetting` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | no |
| `slotTypeName` | `string` | no |
| `description` | `string` | no |
| `slotTypeValues` | `List<SlotTypeValue>` | no |
| `valueSelectionSetting` | `SlotValueSelectionSetting` | no |
| `parentSlotTypeSignature` | `string` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `externalSourceSetting` | `ExternalSourceSetting` | no |
| `compositeSlotTypeSetting` | `CompositeSlotTypeSetting` | no |

## CreateTestSetDiscrepancyReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |
| `target` | `TestSetDiscrepancyReportResourceTarget` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetDiscrepancyReportId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `testSetId` | `string` | no |
| `target` | `TestSetDiscrepancyReportResourceTarget` | no |

## CreateUploadUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `uploadUrl` | `string` | no |

## DeleteBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botStatus` | `string` | no |

## DeleteBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | yes |
| `botId` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | no |
| `botId` | `string` | no |
| `botAliasStatus` | `string` | no |

## DeleteBotAnalyzerRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAnalyzerRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBotLocale

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botLocaleStatus` | `string` | no |

## DeleteBotReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `replicaRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `replicaRegion` | `string` | no |
| `botReplicaStatus` | `string` | no |

## DeleteBotVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `botStatus` | `string` | no |

## DeleteCustomVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `customVocabularyStatus` | `string` | no |

## DeleteExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |
| `exportStatus` | `string` | no |

## DeleteImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `importStatus` | `string` | no |

## DeleteIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `revisionId` | `string` | no |

## DeleteResourcePolicyStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `statementId` | `string` | yes |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `revisionId` | `string` | no |

## DeleteSlot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `intentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `skipResourceInUseCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTestSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUtterances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `localeId` | `string` | no |
| `sessionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botName` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `dataPrivacy` | `DataPrivacy` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `botStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |
| `failureReasons` | `List<string>` | no |
| `errorLogSettings` | `ErrorLogSettings` | no |

## DescribeBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | yes |
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | no |
| `botAliasName` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botAliasLocaleSettings` | `Map<BotAliasLocaleSettings>` | no |
| `conversationLogSettings` | `ConversationLogSettings` | no |
| `sentimentAnalysisSettings` | `SentimentAnalysisSettings` | no |
| `botAliasHistoryEvents` | `List<BotAliasHistoryEvent>` | no |
| `botAliasStatus` | `string` | no |
| `botId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `parentBotNetworks` | `List<ParentBotNetwork>` | no |

## DescribeBotAnalyzerRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAnalyzerRequestId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botAnalyzerStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `botAnalyzerRecommendationList` | `List<BotAnalyzerRecommendation>` | no |
| `nextToken` | `string` | no |

## DescribeBotLocale

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `localeName` | `string` | no |
| `description` | `string` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `voiceSettings` | `VoiceSettings` | no |
| `unifiedSpeechSettings` | `UnifiedSpeechSettings` | no |
| `audioFillerSettings` | `AudioFillerSettings` | no |
| `speechRecognitionSettings` | `SpeechRecognitionSettings` | no |
| `intentsCount` | `integer` | no |
| `slotTypesCount` | `integer` | no |
| `botLocaleStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `lastBuildSubmittedDateTime` | `timestamp` | no |
| `botLocaleHistoryEvents` | `List<BotLocaleHistoryEvent>` | no |
| `recommendedActions` | `List<string>` | no |
| `generativeAISettings` | `GenerativeAISettings` | no |
| `speechDetectionSensitivity` | `string` | no |

## DescribeBotRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `botRecommendationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationStatus` | `string` | no |
| `botRecommendationId` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `transcriptSourceSetting` | `TranscriptSourceSetting` | no |
| `encryptionSetting` | `EncryptionSetting` | no |
| `botRecommendationResults` | `BotRecommendationResults` | no |

## DescribeBotReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `replicaRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `replicaRegion` | `string` | no |
| `sourceRegion` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `botReplicaStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |

## DescribeBotResourceGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `generationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `generationId` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `generationStatus` | `string` | no |
| `generationInputPrompt` | `string` | no |
| `generatedBotLocaleUrl` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `modelArn` | `string` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## DescribeBotVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botName` | `string` | no |
| `botVersion` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `dataPrivacy` | `DataPrivacy` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `botStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `creationDateTime` | `timestamp` | no |
| `parentBotNetworks` | `List<ParentBotNetwork>` | no |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |

## DescribeCustomVocabularyMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `customVocabularyStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## DescribeExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |
| `resourceSpecification` | `ExportResourceSpecification` | no |
| `fileFormat` | `string` | no |
| `exportStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `downloadUrl` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## DescribeImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `resourceSpecification` | `ImportResourceSpecification` | no |
| `importedResourceId` | `string` | no |
| `importedResourceName` | `string` | no |
| `mergeStrategy` | `string` | no |
| `importStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## DescribeIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | no |
| `intentName` | `string` | no |
| `intentDisplayName` | `string` | no |
| `description` | `string` | no |
| `parentIntentSignature` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |
| `dialogCodeHook` | `DialogCodeHookSettings` | no |
| `fulfillmentCodeHook` | `FulfillmentCodeHookSettings` | no |
| `slotPriorities` | `List<SlotPriority>` | no |
| `intentConfirmationSetting` | `IntentConfirmationSetting` | no |
| `intentClosingSetting` | `IntentClosingSetting` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `initialResponseSetting` | `InitialResponseSetting` | no |
| `qnAIntentConfiguration` | `QnAIntentConfiguration` | no |
| `qInConnectIntentConfiguration` | `QInConnectIntentConfiguration` | no |

## DescribeResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `policy` | `string` | no |
| `revisionId` | `string` | no |

## DescribeSlot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `intentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | no |
| `slotName` | `string` | no |
| `description` | `string` | no |
| `slotTypeId` | `string` | no |
| `valueElicitationSetting` | `SlotValueElicitationSetting` | no |
| `obfuscationSetting` | `ObfuscationSetting` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `multipleValuesSetting` | `MultipleValuesSetting` | no |
| `subSlotSetting` | `SubSlotSetting` | no |

## DescribeSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | no |
| `slotTypeName` | `string` | no |
| `description` | `string` | no |
| `slotTypeValues` | `List<SlotTypeValue>` | no |
| `valueSelectionSetting` | `SlotValueSelectionSetting` | no |
| `parentSlotTypeSignature` | `string` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `externalSourceSetting` | `ExternalSourceSetting` | no |
| `compositeSlotTypeSetting` | `CompositeSlotTypeSetting` | no |

## DescribeTestExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `testExecutionStatus` | `string` | no |
| `testSetId` | `string` | no |
| `testSetName` | `string` | no |
| `target` | `TestExecutionTarget` | no |
| `apiMode` | `string` | no |
| `testExecutionModality` | `string` | no |
| `failureReasons` | `List<string>` | no |

## DescribeTestSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | no |
| `testSetName` | `string` | no |
| `description` | `string` | no |
| `modality` | `string` | no |
| `status` | `string` | no |
| `roleArn` | `string` | no |
| `numTurns` | `integer` | no |
| `storageLocation` | `TestSetStorageLocation` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## DescribeTestSetDiscrepancyReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetDiscrepancyReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetDiscrepancyReportId` | `string` | no |
| `testSetId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `target` | `TestSetDiscrepancyReportResourceTarget` | no |
| `testSetDiscrepancyReportStatus` | `string` | no |
| `lastUpdatedDataTime` | `timestamp` | no |
| `testSetDiscrepancyTopErrors` | `TestSetDiscrepancyErrors` | no |
| `testSetDiscrepancyRawOutputUrl` | `string` | no |
| `failureReasons` | `List<string>` | no |

## DescribeTestSetGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetGenerationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetGenerationId` | `string` | no |
| `testSetGenerationStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `testSetId` | `string` | no |
| `testSetName` | `string` | no |
| `description` | `string` | no |
| `storageLocation` | `TestSetStorageLocation` | no |
| `generationDataSource` | `TestSetGenerationDataSource` | no |
| `roleArn` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## GenerateBotElement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentId` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |

## GetTestExecutionArtifactsUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | no |
| `downloadArtifactsUrl` | `string` | no |

## ListAggregatedUtterances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | yes |
| `aggregationDuration` | `UtteranceAggregationDuration` | yes |
| `sortBy` | `AggregatedUtterancesSortBy` | no |
| `filters` | `List<AggregatedUtterancesFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botAliasId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `aggregationDuration` | `UtteranceAggregationDuration` | no |
| `aggregationWindowStartTime` | `timestamp` | no |
| `aggregationWindowEndTime` | `timestamp` | no |
| `aggregationLastRefreshedDateTime` | `timestamp` | no |
| `aggregatedUtterancesSummaries` | `List<AggregatedUtterancesSummary>` | no |
| `nextToken` | `string` | no |

## ListBotAliasReplicas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `replicaRegion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `sourceRegion` | `string` | no |
| `replicaRegion` | `string` | no |
| `botAliasReplicaSummaries` | `List<BotAliasReplicaSummary>` | no |
| `nextToken` | `string` | no |

## ListBotAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasSummaries` | `List<BotAliasSummary>` | no |
| `nextToken` | `string` | no |
| `botId` | `string` | no |

## ListBotAnalyzerHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `localeId` | `string` | no |
| `botVersion` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `localeId` | `string` | no |
| `botVersion` | `string` | no |
| `botAnalyzerHistoryList` | `List<BotAnalyzerHistorySummary>` | no |
| `nextToken` | `string` | no |

## ListBotLocales

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `sortBy` | `BotLocaleSortBy` | no |
| `filters` | `List<BotLocaleFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `nextToken` | `string` | no |
| `botLocaleSummaries` | `List<BotLocaleSummary>` | no |

## ListBotRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationSummaries` | `List<BotRecommendationSummary>` | no |
| `nextToken` | `string` | no |

## ListBotReplicas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `sourceRegion` | `string` | no |
| `botReplicaSummaries` | `List<BotReplicaSummary>` | no |

## ListBotResourceGenerations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `sortBy` | `GenerationSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `generationSummaries` | `List<GenerationSummary>` | no |
| `nextToken` | `string` | no |

## ListBotVersionReplicas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `replicaRegion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `BotVersionReplicaSortBy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `sourceRegion` | `string` | no |
| `replicaRegion` | `string` | no |
| `botVersionReplicaSummaries` | `List<BotVersionReplicaSummary>` | no |
| `nextToken` | `string` | no |

## ListBotVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `sortBy` | `BotVersionSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersionSummaries` | `List<BotVersionSummary>` | no |
| `nextToken` | `string` | no |

## ListBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortBy` | `BotSortBy` | no |
| `filters` | `List<BotFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botSummaries` | `List<BotSummary>` | no |
| `nextToken` | `string` | no |

## ListBuiltInIntents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `localeId` | `string` | yes |
| `sortBy` | `BuiltInIntentSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `builtInIntentSummaries` | `List<BuiltInIntentSummary>` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

## ListBuiltInSlotTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `localeId` | `string` | yes |
| `sortBy` | `BuiltInSlotTypeSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `builtInSlotTypeSummaries` | `List<BuiltInSlotTypeSummary>` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

## ListCustomVocabularyItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `customVocabularyItems` | `List<CustomVocabularyItem>` | no |
| `nextToken` | `string` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `sortBy` | `ExportSortBy` | no |
| `filters` | `List<ExportFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `exportSummaries` | `List<ExportSummary>` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

## ListImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `sortBy` | `ImportSortBy` | no |
| `filters` | `List<ImportFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `importSummaries` | `List<ImportSummary>` | no |
| `nextToken` | `string` | no |
| `localeId` | `string` | no |

## ListIntentMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `metrics` | `List<AnalyticsIntentMetric>` | yes |
| `binBy` | `List<AnalyticsBinBySpecification>` | no |
| `groupBy` | `List<AnalyticsIntentGroupBySpecification>` | no |
| `filters` | `List<AnalyticsIntentFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `results` | `List<AnalyticsIntentResult>` | no |
| `nextToken` | `string` | no |

## ListIntentPaths

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `intentPath` | `string` | yes |
| `filters` | `List<AnalyticsPathFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodeSummaries` | `List<AnalyticsIntentNodeSummary>` | no |

## ListIntentStageMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `metrics` | `List<AnalyticsIntentStageMetric>` | yes |
| `binBy` | `List<AnalyticsBinBySpecification>` | no |
| `groupBy` | `List<AnalyticsIntentStageGroupBySpecification>` | no |
| `filters` | `List<AnalyticsIntentStageFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `results` | `List<AnalyticsIntentStageResult>` | no |
| `nextToken` | `string` | no |

## ListIntents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `sortBy` | `IntentSortBy` | no |
| `filters` | `List<IntentFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentSummaries` | `List<IntentSummary>` | no |
| `nextToken` | `string` | no |

## ListRecommendedIntents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `botRecommendationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationId` | `string` | no |
| `summaryList` | `List<RecommendedIntentSummary>` | no |
| `nextToken` | `string` | no |

## ListSessionAnalyticsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `sortBy` | `SessionDataSortBy` | no |
| `filters` | `List<AnalyticsSessionFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `nextToken` | `string` | no |
| `sessions` | `List<SessionSpecification>` | no |

## ListSessionMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `metrics` | `List<AnalyticsSessionMetric>` | yes |
| `binBy` | `List<AnalyticsBinBySpecification>` | no |
| `groupBy` | `List<AnalyticsSessionGroupBySpecification>` | no |
| `filters` | `List<AnalyticsSessionFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `results` | `List<AnalyticsSessionResult>` | no |
| `nextToken` | `string` | no |

## ListSlotTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `sortBy` | `SlotTypeSortBy` | no |
| `filters` | `List<SlotTypeFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `slotTypeSummaries` | `List<SlotTypeSummary>` | no |
| `nextToken` | `string` | no |

## ListSlots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `intentId` | `string` | yes |
| `sortBy` | `SlotSortBy` | no |
| `filters` | `List<SlotFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentId` | `string` | no |
| `slotSummaries` | `List<SlotSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTestExecutionResultItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | yes |
| `resultFilterBy` | `TestExecutionResultFilterBy` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionResults` | `TestExecutionResultItems` | no |
| `nextToken` | `string` | no |

## ListTestExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortBy` | `TestExecutionSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutions` | `List<TestExecutionSummary>` | no |
| `nextToken` | `string` | no |

## ListTestSetRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetRecords` | `List<TestSetTurnRecord>` | no |
| `nextToken` | `string` | no |

## ListTestSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortBy` | `TestSetSortBy` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSets` | `List<TestSetSummary>` | no |
| `nextToken` | `string` | no |

## ListUtteranceAnalyticsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `sortBy` | `UtteranceDataSortBy` | no |
| `filters` | `List<AnalyticsUtteranceFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `nextToken` | `string` | no |
| `utterances` | `List<UtteranceSpecification>` | no |

## ListUtteranceMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `startDateTime` | `timestamp` | yes |
| `endDateTime` | `timestamp` | yes |
| `metrics` | `List<AnalyticsUtteranceMetric>` | yes |
| `binBy` | `List<AnalyticsBinBySpecification>` | no |
| `groupBy` | `List<AnalyticsUtteranceGroupBySpecification>` | no |
| `attributes` | `List<AnalyticsUtteranceAttribute>` | no |
| `filters` | `List<AnalyticsUtteranceFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `results` | `List<AnalyticsUtteranceResult>` | no |
| `nextToken` | `string` | no |

## SearchAssociatedTranscripts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `botRecommendationId` | `string` | yes |
| `searchOrder` | `string` | no |
| `filters` | `List<AssociatedTranscriptFilter>` | yes |
| `maxResults` | `integer` | no |
| `nextIndex` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationId` | `string` | no |
| `nextIndex` | `integer` | no |
| `associatedTranscripts` | `List<AssociatedTranscript>` | no |
| `totalResults` | `integer` | no |

## StartBotAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `analysisScope` | `string` | yes |
| `localeId` | `string` | no |
| `botVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botAnalyzerStatus` | `string` | no |
| `botAnalyzerRequestId` | `string` | no |
| `creationDateTime` | `timestamp` | no |

## StartBotRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `transcriptSourceSetting` | `TranscriptSourceSetting` | yes |
| `encryptionSetting` | `EncryptionSetting` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationStatus` | `string` | no |
| `botRecommendationId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `transcriptSourceSetting` | `TranscriptSourceSetting` | no |
| `encryptionSetting` | `EncryptionSetting` | no |

## StartBotResourceGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `generationInputPrompt` | `string` | yes |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `generationInputPrompt` | `string` | no |
| `generationId` | `string` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `generationStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |

## StartImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |
| `resourceSpecification` | `ImportResourceSpecification` | yes |
| `mergeStrategy` | `string` | yes |
| `filePassword` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `resourceSpecification` | `ImportResourceSpecification` | no |
| `mergeStrategy` | `string` | no |
| `importStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |

## StartTestExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |
| `target` | `TestExecutionTarget` | yes |
| `apiMode` | `string` | yes |
| `testExecutionModality` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testExecutionId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `testSetId` | `string` | no |
| `target` | `TestExecutionTarget` | no |
| `apiMode` | `string` | no |
| `testExecutionModality` | `string` | no |

## StartTestSetGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetName` | `string` | yes |
| `description` | `string` | no |
| `storageLocation` | `TestSetStorageLocation` | yes |
| `generationDataSource` | `TestSetGenerationDataSource` | yes |
| `roleArn` | `string` | yes |
| `testSetTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetGenerationId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `testSetGenerationStatus` | `string` | no |
| `testSetName` | `string` | no |
| `description` | `string` | no |
| `storageLocation` | `TestSetStorageLocation` | no |
| `generationDataSource` | `TestSetGenerationDataSource` | no |
| `roleArn` | `string` | no |
| `testSetTags` | `Map<string>` | no |

## StopBotAnalyzer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAnalyzerRequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botAnalyzerStatus` | `string` | no |
| `botAnalyzerRequestId` | `string` | no |

## StopBotRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `botRecommendationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationStatus` | `string` | no |
| `botRecommendationId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `Map<string>` | yes |

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


## UpdateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botName` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `dataPrivacy` | `DataPrivacy` | yes |
| `idleSessionTTLInSeconds` | `integer` | yes |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |
| `errorLogSettings` | `ErrorLogSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botName` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `dataPrivacy` | `DataPrivacy` | no |
| `idleSessionTTLInSeconds` | `integer` | no |
| `botStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `botType` | `string` | no |
| `botMembers` | `List<BotMember>` | no |
| `errorLogSettings` | `ErrorLogSettings` | no |

## UpdateBotAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | yes |
| `botAliasName` | `string` | yes |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botAliasLocaleSettings` | `Map<BotAliasLocaleSettings>` | no |
| `conversationLogSettings` | `ConversationLogSettings` | no |
| `sentimentAnalysisSettings` | `SentimentAnalysisSettings` | no |
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botAliasId` | `string` | no |
| `botAliasName` | `string` | no |
| `description` | `string` | no |
| `botVersion` | `string` | no |
| `botAliasLocaleSettings` | `Map<BotAliasLocaleSettings>` | no |
| `conversationLogSettings` | `ConversationLogSettings` | no |
| `sentimentAnalysisSettings` | `SentimentAnalysisSettings` | no |
| `botAliasStatus` | `string` | no |
| `botId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## UpdateBotLocale

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `description` | `string` | no |
| `nluIntentConfidenceThreshold` | `double` | yes |
| `voiceSettings` | `VoiceSettings` | no |
| `unifiedSpeechSettings` | `UnifiedSpeechSettings` | no |
| `audioFillerSettings` | `AudioFillerSettings` | no |
| `speechRecognitionSettings` | `SpeechRecognitionSettings` | no |
| `generativeAISettings` | `GenerativeAISettings` | no |
| `speechDetectionSensitivity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `localeName` | `string` | no |
| `description` | `string` | no |
| `nluIntentConfidenceThreshold` | `double` | no |
| `voiceSettings` | `VoiceSettings` | no |
| `unifiedSpeechSettings` | `UnifiedSpeechSettings` | no |
| `audioFillerSettings` | `AudioFillerSettings` | no |
| `speechRecognitionSettings` | `SpeechRecognitionSettings` | no |
| `botLocaleStatus` | `string` | no |
| `failureReasons` | `List<string>` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `recommendedActions` | `List<string>` | no |
| `generativeAISettings` | `GenerativeAISettings` | no |
| `speechDetectionSensitivity` | `string` | no |

## UpdateBotRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `botRecommendationId` | `string` | yes |
| `encryptionSetting` | `EncryptionSetting` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `botRecommendationStatus` | `string` | no |
| `botRecommendationId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `transcriptSourceSetting` | `TranscriptSourceSetting` | no |
| `encryptionSetting` | `EncryptionSetting` | no |

## UpdateExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | yes |
| `filePassword` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |
| `resourceSpecification` | `ExportResourceSpecification` | no |
| `fileFormat` | `string` | no |
| `exportStatus` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

## UpdateIntent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | yes |
| `intentName` | `string` | yes |
| `intentDisplayName` | `string` | no |
| `description` | `string` | no |
| `parentIntentSignature` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |
| `dialogCodeHook` | `DialogCodeHookSettings` | no |
| `fulfillmentCodeHook` | `FulfillmentCodeHookSettings` | no |
| `slotPriorities` | `List<SlotPriority>` | no |
| `intentConfirmationSetting` | `IntentConfirmationSetting` | no |
| `intentClosingSetting` | `IntentClosingSetting` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `initialResponseSetting` | `InitialResponseSetting` | no |
| `qnAIntentConfiguration` | `QnAIntentConfiguration` | no |
| `qInConnectIntentConfiguration` | `QInConnectIntentConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentId` | `string` | no |
| `intentName` | `string` | no |
| `intentDisplayName` | `string` | no |
| `description` | `string` | no |
| `parentIntentSignature` | `string` | no |
| `sampleUtterances` | `List<SampleUtterance>` | no |
| `dialogCodeHook` | `DialogCodeHookSettings` | no |
| `fulfillmentCodeHook` | `FulfillmentCodeHookSettings` | no |
| `slotPriorities` | `List<SlotPriority>` | no |
| `intentConfirmationSetting` | `IntentConfirmationSetting` | no |
| `intentClosingSetting` | `IntentClosingSetting` | no |
| `inputContexts` | `List<InputContext>` | no |
| `outputContexts` | `List<OutputContext>` | no |
| `kendraConfiguration` | `KendraConfiguration` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `initialResponseSetting` | `InitialResponseSetting` | no |
| `qnAIntentConfiguration` | `QnAIntentConfiguration` | no |
| `qInConnectIntentConfiguration` | `QInConnectIntentConfiguration` | no |

## UpdateResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `revisionId` | `string` | no |

## UpdateSlot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | yes |
| `slotName` | `string` | yes |
| `description` | `string` | no |
| `slotTypeId` | `string` | no |
| `valueElicitationSetting` | `SlotValueElicitationSetting` | yes |
| `obfuscationSetting` | `ObfuscationSetting` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `intentId` | `string` | yes |
| `multipleValuesSetting` | `MultipleValuesSetting` | no |
| `subSlotSetting` | `SubSlotSetting` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotId` | `string` | no |
| `slotName` | `string` | no |
| `description` | `string` | no |
| `slotTypeId` | `string` | no |
| `valueElicitationSetting` | `SlotValueElicitationSetting` | no |
| `obfuscationSetting` | `ObfuscationSetting` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `intentId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `multipleValuesSetting` | `MultipleValuesSetting` | no |
| `subSlotSetting` | `SubSlotSetting` | no |

## UpdateSlotType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | yes |
| `slotTypeName` | `string` | yes |
| `description` | `string` | no |
| `slotTypeValues` | `List<SlotTypeValue>` | no |
| `valueSelectionSetting` | `SlotValueSelectionSetting` | no |
| `parentSlotTypeSignature` | `string` | no |
| `botId` | `string` | yes |
| `botVersion` | `string` | yes |
| `localeId` | `string` | yes |
| `externalSourceSetting` | `ExternalSourceSetting` | no |
| `compositeSlotTypeSetting` | `CompositeSlotTypeSetting` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `slotTypeId` | `string` | no |
| `slotTypeName` | `string` | no |
| `description` | `string` | no |
| `slotTypeValues` | `List<SlotTypeValue>` | no |
| `valueSelectionSetting` | `SlotValueSelectionSetting` | no |
| `parentSlotTypeSignature` | `string` | no |
| `botId` | `string` | no |
| `botVersion` | `string` | no |
| `localeId` | `string` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |
| `externalSourceSetting` | `ExternalSourceSetting` | no |
| `compositeSlotTypeSetting` | `CompositeSlotTypeSetting` | no |

## UpdateTestSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | yes |
| `testSetName` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testSetId` | `string` | no |
| `testSetName` | `string` | no |
| `description` | `string` | no |
| `modality` | `string` | no |
| `status` | `string` | no |
| `roleArn` | `string` | no |
| `numTurns` | `integer` | no |
| `storageLocation` | `TestSetStorageLocation` | no |
| `creationDateTime` | `timestamp` | no |
| `lastUpdatedDateTime` | `timestamp` | no |

