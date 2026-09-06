# AWS Backup Search

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/backupsearch/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetSearchJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `SearchScopeSummary` | `SearchScopeSummary` | no |
| `CurrentSearchProgress` | `CurrentSearchProgress` | no |
| `StatusMessage` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `CompletionTime` | `timestamp` | no |
| `Status` | `string` | yes |
| `SearchScope` | `SearchScope` | yes |
| `ItemFilters` | `ItemFilters` | yes |
| `CreationTime` | `timestamp` | yes |
| `SearchJobIdentifier` | `string` | yes |
| `SearchJobArn` | `string` | yes |

## GetSearchResultExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobIdentifier` | `string` | yes |
| `ExportJobArn` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `CompletionTime` | `timestamp` | no |
| `StatusMessage` | `string` | no |
| `ExportSpecification` | `ExportSpecification` | no |
| `SearchJobArn` | `string` | no |

## ListSearchJobBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobIdentifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<SearchJobBackupsResult>` | yes |
| `NextToken` | `string` | no |

## ListSearchJobResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobIdentifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<ResultItem>` | yes |
| `NextToken` | `string` | no |

## ListSearchJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByStatus` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobs` | `List<SearchJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListSearchResultExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `SearchJobIdentifier` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobs` | `List<ExportJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartSearchJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |
| `Name` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `ClientToken` | `string` | no |
| `SearchScope` | `SearchScope` | yes |
| `ItemFilters` | `ItemFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobArn` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `SearchJobIdentifier` | `string` | no |

## StartSearchResultExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobIdentifier` | `string` | yes |
| `ExportSpecification` | `ExportSpecification` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobArn` | `string` | no |
| `ExportJobIdentifier` | `string` | yes |

## StopSearchJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchJobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


