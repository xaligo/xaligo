# Amazon Kinesis Analytics

API version: 2015-08-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kinesisanalytics/2015-08-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddApplicationCloudWatchLoggingOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `CloudWatchLoggingOption` | `CloudWatchLoggingOption` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddApplicationInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `Input` | `Input` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddApplicationInputProcessingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `InputId` | `string` | yes |
| `InputProcessingConfiguration` | `InputProcessingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddApplicationOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `Output` | `Output` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddApplicationReferenceDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `ReferenceDataSource` | `ReferenceDataSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `ApplicationDescription` | `string` | no |
| `Inputs` | `List<Input>` | no |
| `Outputs` | `List<Output>` | no |
| `CloudWatchLoggingOptions` | `List<CloudWatchLoggingOption>` | no |
| `ApplicationCode` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSummary` | `ApplicationSummary` | yes |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CreateTimestamp` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationCloudWatchLoggingOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `CloudWatchLoggingOptionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationInputProcessingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `InputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `OutputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationReferenceDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `ReferenceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationDetail` | `ApplicationDetail` | yes |

## DiscoverInputSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | no |
| `RoleARN` | `string` | no |
| `InputStartingPositionConfiguration` | `InputStartingPositionConfiguration` | no |
| `S3Configuration` | `S3Configuration` | no |
| `InputProcessingConfiguration` | `InputProcessingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSchema` | `SourceSchema` | no |
| `ParsedInputRecords` | `List<List<string>>` | no |
| `ProcessedInputRecords` | `List<string>` | no |
| `RawInputRecords` | `List<string>` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `ExclusiveStartApplicationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSummaries` | `List<ApplicationSummary>` | yes |
| `HasMoreApplications` | `boolean` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## StartApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `InputConfigurations` | `List<InputConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `CurrentApplicationVersionId` | `long` | yes |
| `ApplicationUpdate` | `ApplicationUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


