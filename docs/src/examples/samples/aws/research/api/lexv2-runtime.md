# Amazon Lex Runtime V2

API version: 2020-08-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lexv2-runtime/2020-08-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `botAliasId` | `string` | no |
| `localeId` | `string` | no |
| `sessionId` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | no |
| `messages` | `List<Message>` | no |
| `interpretations` | `List<Interpretation>` | no |
| `sessionState` | `SessionState` | no |

## PutSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |
| `messages` | `List<Message>` | no |
| `sessionState` | `SessionState` | yes |
| `requestAttributes` | `Map<string>` | no |
| `responseContentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `messages` | `string` | no |
| `sessionState` | `string` | no |
| `requestAttributes` | `string` | no |
| `sessionId` | `string` | no |
| `audioStream` | `blob` | no |

## RecognizeText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |
| `text` | `string` | yes |
| `sessionState` | `SessionState` | no |
| `requestAttributes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `messages` | `List<Message>` | no |
| `sessionState` | `SessionState` | no |
| `interpretations` | `List<Interpretation>` | no |
| `requestAttributes` | `Map<string>` | no |
| `sessionId` | `string` | no |
| `recognizedBotMember` | `RecognizedBotMember` | no |

## RecognizeUtterance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |
| `sessionState` | `string` | no |
| `requestAttributes` | `string` | no |
| `requestContentType` | `string` | yes |
| `responseContentType` | `string` | no |
| `inputStream` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inputMode` | `string` | no |
| `contentType` | `string` | no |
| `messages` | `string` | no |
| `interpretations` | `string` | no |
| `sessionState` | `string` | no |
| `requestAttributes` | `string` | no |
| `sessionId` | `string` | no |
| `inputTranscript` | `string` | no |
| `audioStream` | `blob` | no |
| `recognizedBotMember` | `string` | no |

## StartConversation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | yes |
| `botAliasId` | `string` | yes |
| `localeId` | `string` | yes |
| `sessionId` | `string` | yes |
| `conversationMode` | `string` | no |
| `requestEventStream` | `StartConversationRequestEventStream` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseEventStream` | `StartConversationResponseEventStream` | no |

