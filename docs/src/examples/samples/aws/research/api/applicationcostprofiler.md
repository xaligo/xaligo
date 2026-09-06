# AWS Application Cost Profiler

API version: 2020-09-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/applicationcostprofiler/2020-09-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

## GetReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportDescription` | `string` | yes |
| `reportFrequency` | `string` | yes |
| `format` | `string` | yes |
| `destinationS3Location` | `S3Location` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdated` | `timestamp` | yes |

## ImportApplicationUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceS3Location` | `SourceS3Location` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |

## ListReportDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportDefinitions` | `List<ReportDefinition>` | no |
| `nextToken` | `string` | no |

## PutReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportDescription` | `string` | yes |
| `reportFrequency` | `string` | yes |
| `format` | `string` | yes |
| `destinationS3Location` | `S3Location` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

## UpdateReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | yes |
| `reportDescription` | `string` | yes |
| `reportFrequency` | `string` | yes |
| `format` | `string` | yes |
| `destinationS3Location` | `S3Location` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportId` | `string` | no |

