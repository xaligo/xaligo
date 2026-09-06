# Amazon Translate

API version: 2017-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/translate/2017-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateParallelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ParallelDataConfig` | `ParallelDataConfig` | yes |
| `EncryptionKey` | `EncryptionKey` | no |
| `ClientToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Status` | `string` | no |

## DeleteParallelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Status` | `string` | no |

## DeleteTerminology

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeTextTranslationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextTranslationJobProperties` | `TextTranslationJobProperties` | no |

## GetParallelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParallelDataProperties` | `ParallelDataProperties` | no |
| `DataLocation` | `ParallelDataDataLocation` | no |
| `AuxiliaryDataLocation` | `ParallelDataDataLocation` | no |
| `LatestUpdateAttemptAuxiliaryDataLocation` | `ParallelDataDataLocation` | no |

## GetTerminology

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TerminologyDataFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminologyProperties` | `TerminologyProperties` | no |
| `TerminologyDataLocation` | `TerminologyDataLocation` | no |
| `AuxiliaryDataLocation` | `TerminologyDataLocation` | no |

## ImportTerminology

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MergeStrategy` | `string` | yes |
| `Description` | `string` | no |
| `TerminologyData` | `TerminologyData` | yes |
| `EncryptionKey` | `EncryptionKey` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminologyProperties` | `TerminologyProperties` | no |
| `AuxiliaryDataLocation` | `TerminologyDataLocation` | no |

## ListLanguages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayLanguageCode` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Languages` | `List<Language>` | no |
| `DisplayLanguageCode` | `string` | no |
| `NextToken` | `string` | no |

## ListParallelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParallelDataPropertiesList` | `List<ParallelDataProperties>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTerminologies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminologyPropertiesList` | `List<TerminologyProperties>` | no |
| `NextToken` | `string` | no |

## ListTextTranslationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `TextTranslationJobFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TextTranslationJobPropertiesList` | `List<TextTranslationJobProperties>` | no |
| `NextToken` | `string` | no |

## StartTextTranslationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | no |
| `InputDataConfig` | `InputDataConfig` | yes |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `DataAccessRoleArn` | `string` | yes |
| `SourceLanguageCode` | `string` | yes |
| `TargetLanguageCodes` | `List<string>` | yes |
| `TerminologyNames` | `List<string>` | no |
| `ParallelDataNames` | `List<string>` | no |
| `ClientToken` | `string` | yes |
| `Settings` | `TranslationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## StopTextTranslationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobStatus` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TranslateDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Document` | `Document` | yes |
| `TerminologyNames` | `List<string>` | no |
| `SourceLanguageCode` | `string` | yes |
| `TargetLanguageCode` | `string` | yes |
| `Settings` | `TranslationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranslatedDocument` | `TranslatedDocument` | yes |
| `SourceLanguageCode` | `string` | yes |
| `TargetLanguageCode` | `string` | yes |
| `AppliedTerminologies` | `List<AppliedTerminology>` | no |
| `AppliedSettings` | `TranslationSettings` | no |

## TranslateText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Text` | `string` | yes |
| `TerminologyNames` | `List<string>` | no |
| `SourceLanguageCode` | `string` | yes |
| `TargetLanguageCode` | `string` | yes |
| `Settings` | `TranslationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TranslatedText` | `string` | yes |
| `SourceLanguageCode` | `string` | yes |
| `TargetLanguageCode` | `string` | yes |
| `AppliedTerminologies` | `List<AppliedTerminology>` | no |
| `AppliedSettings` | `TranslationSettings` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateParallelData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ParallelDataConfig` | `ParallelDataConfig` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Status` | `string` | no |
| `LatestUpdateAttemptStatus` | `string` | no |
| `LatestUpdateAttemptAt` | `timestamp` | no |

