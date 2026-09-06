# Amazon Textract

API version: 2018-06-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/textract/2018-06-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AnalyzeDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Document` | `Document` | yes |
| `FeatureTypes` | `List<string>` | yes |
| `HumanLoopConfig` | `HumanLoopConfig` | no |
| `QueriesConfig` | `QueriesConfig` | no |
| `AdaptersConfig` | `AdaptersConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `Blocks` | `List<Block>` | no |
| `HumanLoopActivationOutput` | `HumanLoopActivationOutput` | no |
| `AnalyzeDocumentModelVersion` | `string` | no |

## AnalyzeExpense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Document` | `Document` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `ExpenseDocuments` | `List<ExpenseDocument>` | no |

## AnalyzeID

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentPages` | `List<Document>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityDocuments` | `List<IdentityDocument>` | no |
| `DocumentMetadata` | `DocumentMetadata` | no |
| `AnalyzeIDModelVersion` | `string` | no |

## CreateAdapter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterName` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Description` | `string` | no |
| `FeatureTypes` | `List<string>` | yes |
| `AutoUpdate` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |

## CreateAdapterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `DatasetConfig` | `AdapterVersionDatasetConfig` | yes |
| `KMSKeyId` | `string` | no |
| `OutputConfig` | `OutputConfig` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |
| `AdapterVersion` | `string` | no |

## DeleteAdapter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAdapterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |
| `AdapterVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetectDocumentText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Document` | `Document` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `Blocks` | `List<Block>` | no |
| `DetectDocumentTextModelVersion` | `string` | no |

## GetAdapter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |
| `AdapterName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Description` | `string` | no |
| `FeatureTypes` | `List<string>` | no |
| `AutoUpdate` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetAdapterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |
| `AdapterVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |
| `AdapterVersion` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `FeatureTypes` | `List<string>` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `DatasetConfig` | `AdapterVersionDatasetConfig` | no |
| `KMSKeyId` | `string` | no |
| `OutputConfig` | `OutputConfig` | no |
| `EvaluationMetrics` | `List<AdapterVersionEvaluationMetric>` | no |
| `Tags` | `Map<string>` | no |

## GetDocumentAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `JobStatus` | `string` | no |
| `NextToken` | `string` | no |
| `Blocks` | `List<Block>` | no |
| `Warnings` | `List<Warning>` | no |
| `StatusMessage` | `string` | no |
| `AnalyzeDocumentModelVersion` | `string` | no |

## GetDocumentTextDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `JobStatus` | `string` | no |
| `NextToken` | `string` | no |
| `Blocks` | `List<Block>` | no |
| `Warnings` | `List<Warning>` | no |
| `StatusMessage` | `string` | no |
| `DetectDocumentTextModelVersion` | `string` | no |

## GetExpenseAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `JobStatus` | `string` | no |
| `NextToken` | `string` | no |
| `ExpenseDocuments` | `List<ExpenseDocument>` | no |
| `Warnings` | `List<Warning>` | no |
| `StatusMessage` | `string` | no |
| `AnalyzeExpenseModelVersion` | `string` | no |

## GetLendingAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `JobStatus` | `string` | no |
| `NextToken` | `string` | no |
| `Results` | `List<LendingResult>` | no |
| `Warnings` | `List<Warning>` | no |
| `StatusMessage` | `string` | no |
| `AnalyzeLendingModelVersion` | `string` | no |

## GetLendingAnalysisSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentMetadata` | `DocumentMetadata` | no |
| `JobStatus` | `string` | no |
| `Summary` | `LendingSummary` | no |
| `Warnings` | `List<Warning>` | no |
| `StatusMessage` | `string` | no |
| `AnalyzeLendingModelVersion` | `string` | no |

## ListAdapterVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |
| `AfterCreationTime` | `timestamp` | no |
| `BeforeCreationTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterVersions` | `List<AdapterVersionOverview>` | no |
| `NextToken` | `string` | no |

## ListAdapters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AfterCreationTime` | `timestamp` | no |
| `BeforeCreationTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Adapters` | `List<AdapterOverview>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartDocumentAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentLocation` | `DocumentLocation` | yes |
| `FeatureTypes` | `List<string>` | yes |
| `ClientRequestToken` | `string` | no |
| `JobTag` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `OutputConfig` | `OutputConfig` | no |
| `KMSKeyId` | `string` | no |
| `QueriesConfig` | `QueriesConfig` | no |
| `AdaptersConfig` | `AdaptersConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartDocumentTextDetection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentLocation` | `DocumentLocation` | yes |
| `ClientRequestToken` | `string` | no |
| `JobTag` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `OutputConfig` | `OutputConfig` | no |
| `KMSKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartExpenseAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentLocation` | `DocumentLocation` | yes |
| `ClientRequestToken` | `string` | no |
| `JobTag` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `OutputConfig` | `OutputConfig` | no |
| `KMSKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## StartLendingAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DocumentLocation` | `DocumentLocation` | yes |
| `ClientRequestToken` | `string` | no |
| `JobTag` | `string` | no |
| `NotificationChannel` | `NotificationChannel` | no |
| `OutputConfig` | `OutputConfig` | no |
| `KMSKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAdapter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | yes |
| `Description` | `string` | no |
| `AdapterName` | `string` | no |
| `AutoUpdate` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdapterId` | `string` | no |
| `AdapterName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `Description` | `string` | no |
| `FeatureTypes` | `List<string>` | no |
| `AutoUpdate` | `string` | no |

