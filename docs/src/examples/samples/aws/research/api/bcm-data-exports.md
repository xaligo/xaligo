# AWS Billing and Cost Management Data Exports

API version: 2023-11-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bcm-data-exports/2023-11-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Export` | `Export` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | no |

## DeleteExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | no |

## GetExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |
| `ExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionId` | `string` | no |
| `Export` | `Export` | no |
| `ExecutionStatus` | `ExecutionStatus` | no |

## GetExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Export` | `Export` | no |
| `ExportStatus` | `ExportStatus` | no |

## GetTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | yes |
| `TableProperties` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `Description` | `string` | no |
| `TableProperties` | `Map<string>` | no |
| `Schema` | `List<Column>` | no |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Executions` | `List<ExecutionReference>` | no |
| `NextToken` | `string` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Exports` | `List<ExportReference>` | no |
| `NextToken` | `string` | no |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tables` | `List<Table>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTags` | `List<ResourceTag>` | no |
| `NextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | yes |
| `Export` | `Export` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportArn` | `string` | no |

