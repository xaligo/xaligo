# Amazon Lex Runtime Service

API version: 2016-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lex-runtime/2016-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `userId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | no |
| `botAlias` | `string` | no |
| `userId` | `string` | no |
| `sessionId` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `userId` | `string` | yes |
| `checkpointLabelFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recentIntentSummaryView` | `List<IntentSummary>` | no |
| `sessionAttributes` | `Map<string>` | no |
| `sessionId` | `string` | no |
| `dialogAction` | `DialogAction` | no |
| `activeContexts` | `List<ActiveContext>` | no |

## PostContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `userId` | `string` | yes |
| `sessionAttributes` | `string` | no |
| `requestAttributes` | `string` | no |
| `contentType` | `string` | yes |
| `accept` | `string` | no |
| `inputStream` | `blob` | yes |
| `activeContexts` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `intentName` | `string` | no |
| `nluIntentConfidence` | `string` | no |
| `alternativeIntents` | `string` | no |
| `slots` | `string` | no |
| `sessionAttributes` | `string` | no |
| `sentimentResponse` | `string` | no |
| `message` | `string` | no |
| `encodedMessage` | `string` | no |
| `messageFormat` | `string` | no |
| `dialogState` | `string` | no |
| `slotToElicit` | `string` | no |
| `inputTranscript` | `string` | no |
| `encodedInputTranscript` | `string` | no |
| `audioStream` | `blob` | no |
| `botVersion` | `string` | no |
| `sessionId` | `string` | no |
| `activeContexts` | `string` | no |

## PostText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `userId` | `string` | yes |
| `sessionAttributes` | `Map<string>` | no |
| `requestAttributes` | `Map<string>` | no |
| `inputText` | `string` | yes |
| `activeContexts` | `List<ActiveContext>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `intentName` | `string` | no |
| `nluIntentConfidence` | `IntentConfidence` | no |
| `alternativeIntents` | `List<PredictedIntent>` | no |
| `slots` | `Map<string>` | no |
| `sessionAttributes` | `Map<string>` | no |
| `message` | `string` | no |
| `sentimentResponse` | `SentimentResponse` | no |
| `messageFormat` | `string` | no |
| `dialogState` | `string` | no |
| `slotToElicit` | `string` | no |
| `responseCard` | `ResponseCard` | no |
| `sessionId` | `string` | no |
| `botVersion` | `string` | no |
| `activeContexts` | `List<ActiveContext>` | no |

## PutSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | yes |
| `botAlias` | `string` | yes |
| `userId` | `string` | yes |
| `sessionAttributes` | `Map<string>` | no |
| `dialogAction` | `DialogAction` | no |
| `recentIntentSummaryView` | `List<IntentSummary>` | no |
| `accept` | `string` | no |
| `activeContexts` | `List<ActiveContext>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `intentName` | `string` | no |
| `slots` | `string` | no |
| `sessionAttributes` | `string` | no |
| `message` | `string` | no |
| `encodedMessage` | `string` | no |
| `messageFormat` | `string` | no |
| `dialogState` | `string` | no |
| `slotToElicit` | `string` | no |
| `audioStream` | `blob` | no |
| `sessionId` | `string` | no |
| `activeContexts` | `string` | no |

