# Amazon Transcribe Service

API version: 2017-10-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/transcribe/2017-10-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCallAnalyticsCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryName` | `string` | yes |
| `Rules` | `List<Rule>` | yes |
| `Tags` | `List<Tag>` | no |
| `InputType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryProperties` | `CategoryProperties` | no |

## CreateLanguageModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LanguageCode` | `string` | yes |
| `BaseModelName` | `string` | yes |
| `ModelName` | `string` | yes |
| `InputDataConfig` | `InputDataConfig` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LanguageCode` | `string` | no |
| `BaseModelName` | `string` | no |
| `ModelName` | `string` | no |
| `InputDataConfig` | `InputDataConfig` | no |
| `ModelStatus` | `string` | no |

## CreateMedicalVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `VocabularyFileUri` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `VocabularyState` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `FailureReason` | `string` | no |

## CreateVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `Phrases` | `List<string>` | no |
| `VocabularyFileUri` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataAccessRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `VocabularyState` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `FailureReason` | `string` | no |

## CreateVocabularyFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `Words` | `List<string>` | no |
| `VocabularyFilterFileUri` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DataAccessRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | no |
| `LanguageCode` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |

## DeleteCallAnalyticsCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCallAnalyticsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAnalyticsJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLanguageModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMedicalScribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalScribeJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMedicalTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalTranscriptionJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMedicalVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranscriptionJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVocabularyFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeLanguageModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LanguageModel` | `LanguageModel` | no |

## GetCallAnalyticsCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryProperties` | `CategoryProperties` | no |

## GetCallAnalyticsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAnalyticsJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAnalyticsJob` | `CallAnalyticsJob` | no |

## GetMedicalScribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalScribeJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalScribeJob` | `MedicalScribeJob` | no |

## GetMedicalTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalTranscriptionJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalTranscriptionJob` | `MedicalTranscriptionJob` | no |

## GetMedicalVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `VocabularyState` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `DownloadUri` | `string` | no |

## GetTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranscriptionJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranscriptionJob` | `TranscriptionJob` | no |

## GetVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `VocabularyState` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `DownloadUri` | `string` | no |

## GetVocabularyFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | no |
| `LanguageCode` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `DownloadUri` | `string` | no |

## ListCallAnalyticsCategories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Categories` | `List<CategoryProperties>` | no |

## ListCallAnalyticsJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `JobNameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `CallAnalyticsJobSummaries` | `List<CallAnalyticsJobSummary>` | no |

## ListLanguageModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusEquals` | `string` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Models` | `List<LanguageModel>` | no |

## ListMedicalScribeJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `JobNameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MedicalScribeJobSummaries` | `List<MedicalScribeJobSummary>` | no |

## ListMedicalTranscriptionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `JobNameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MedicalTranscriptionJobSummaries` | `List<MedicalTranscriptionJobSummary>` | no |

## ListMedicalVocabularies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StateEquals` | `string` | no |
| `NameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `Vocabularies` | `List<VocabularyInfo>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ListTranscriptionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `JobNameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `TranscriptionJobSummaries` | `List<TranscriptionJobSummary>` | no |

## ListVocabularies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StateEquals` | `string` | no |
| `NameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `Vocabularies` | `List<VocabularyInfo>` | no |

## ListVocabularyFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VocabularyFilters` | `List<VocabularyFilterInfo>` | no |

## StartCallAnalyticsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAnalyticsJobName` | `string` | yes |
| `Media` | `Media` | yes |
| `OutputLocation` | `string` | no |
| `OutputEncryptionKMSKeyId` | `string` | no |
| `DataAccessRoleArn` | `string` | no |
| `Settings` | `CallAnalyticsJobSettings` | no |
| `Tags` | `List<Tag>` | no |
| `ChannelDefinitions` | `List<ChannelDefinition>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAnalyticsJob` | `CallAnalyticsJob` | no |

## StartMedicalScribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalScribeJobName` | `string` | yes |
| `Media` | `Media` | yes |
| `OutputBucketName` | `string` | yes |
| `OutputEncryptionKMSKeyId` | `string` | no |
| `KMSEncryptionContext` | `Map<string>` | no |
| `DataAccessRoleArn` | `string` | yes |
| `Settings` | `MedicalScribeSettings` | yes |
| `ChannelDefinitions` | `List<MedicalScribeChannelDefinition>` | no |
| `Tags` | `List<Tag>` | no |
| `MedicalScribeContext` | `MedicalScribeContext` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalScribeJob` | `MedicalScribeJob` | no |

## StartMedicalTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalTranscriptionJobName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `MediaSampleRateHertz` | `integer` | no |
| `MediaFormat` | `string` | no |
| `Media` | `Media` | yes |
| `OutputBucketName` | `string` | yes |
| `OutputKey` | `string` | no |
| `OutputEncryptionKMSKeyId` | `string` | no |
| `KMSEncryptionContext` | `Map<string>` | no |
| `Settings` | `MedicalTranscriptionSetting` | no |
| `ContentIdentificationType` | `string` | no |
| `Specialty` | `string` | yes |
| `Type` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MedicalTranscriptionJob` | `MedicalTranscriptionJob` | no |

## StartTranscriptionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranscriptionJobName` | `string` | yes |
| `LanguageCode` | `string` | no |
| `MediaSampleRateHertz` | `integer` | no |
| `MediaFormat` | `string` | no |
| `Media` | `Media` | yes |
| `OutputBucketName` | `string` | no |
| `OutputKey` | `string` | no |
| `OutputEncryptionKMSKeyId` | `string` | no |
| `KMSEncryptionContext` | `Map<string>` | no |
| `Settings` | `Settings` | no |
| `ModelSettings` | `ModelSettings` | no |
| `JobExecutionSettings` | `JobExecutionSettings` | no |
| `ContentRedaction` | `ContentRedaction` | no |
| `IdentifyLanguage` | `boolean` | no |
| `IdentifyMultipleLanguages` | `boolean` | no |
| `LanguageOptions` | `List<string>` | no |
| `Subtitles` | `Subtitles` | no |
| `Tags` | `List<Tag>` | no |
| `LanguageIdSettings` | `Map<LanguageIdSettings>` | no |
| `ToxicityDetection` | `List<ToxicityDetectionSettings>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranscriptionJob` | `TranscriptionJob` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCallAnalyticsCategory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryName` | `string` | yes |
| `Rules` | `List<Rule>` | yes |
| `InputType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CategoryProperties` | `CategoryProperties` | no |

## UpdateMedicalVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `VocabularyFileUri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `VocabularyState` | `string` | no |

## UpdateVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `Phrases` | `List<string>` | no |
| `VocabularyFileUri` | `string` | no |
| `DataAccessRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyName` | `string` | no |
| `LanguageCode` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `VocabularyState` | `string` | no |

## UpdateVocabularyFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | yes |
| `Words` | `List<string>` | no |
| `VocabularyFilterFileUri` | `string` | no |
| `DataAccessRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyFilterName` | `string` | no |
| `LanguageCode` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |

