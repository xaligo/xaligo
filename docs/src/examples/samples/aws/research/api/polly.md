# Amazon Polly

API version: 2016-06-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/polly/2016-06-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteLexicon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeVoices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `LanguageCode` | `string` | no |
| `IncludeAdditionalLanguageCodes` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Voices` | `List<Voice>` | no |
| `NextToken` | `string` | no |

## GetLexicon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Lexicon` | `Lexicon` | no |
| `LexiconAttributes` | `LexiconAttributes` | no |

## GetSpeechSynthesisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SynthesisTask` | `SynthesisTask` | no |

## ListLexicons

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Lexicons` | `List<LexiconDescription>` | no |
| `NextToken` | `string` | no |

## ListSpeechSynthesisTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SynthesisTasks` | `List<SynthesisTask>` | no |

## PutLexicon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Content` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartSpeechSynthesisStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | yes |
| `LanguageCode` | `string` | no |
| `LexiconNames` | `List<string>` | no |
| `OutputFormat` | `string` | yes |
| `SampleRate` | `string` | no |
| `VoiceId` | `string` | yes |
| `ActionStream` | `StartSpeechSynthesisStreamActionStream` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStream` | `StartSpeechSynthesisStreamEventStream` | no |

## StartSpeechSynthesisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `LanguageCode` | `string` | no |
| `LexiconNames` | `List<string>` | no |
| `OutputFormat` | `string` | yes |
| `OutputS3BucketName` | `string` | yes |
| `OutputS3KeyPrefix` | `string` | no |
| `SampleRate` | `string` | no |
| `SnsTopicArn` | `string` | no |
| `SpeechMarkTypes` | `List<string>` | no |
| `Text` | `string` | yes |
| `TextType` | `string` | no |
| `VoiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SynthesisTask` | `SynthesisTask` | no |

## SynthesizeSpeech

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Engine` | `string` | no |
| `LanguageCode` | `string` | no |
| `LexiconNames` | `List<string>` | no |
| `OutputFormat` | `string` | yes |
| `SampleRate` | `string` | no |
| `SpeechMarkTypes` | `List<string>` | no |
| `Text` | `string` | yes |
| `TextType` | `string` | no |
| `VoiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AudioStream` | `blob` | no |
| `ContentType` | `string` | no |
| `RequestCharacters` | `integer` | no |

