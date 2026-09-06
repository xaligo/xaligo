# AWS Cost and Usage Report Service

API version: 2017-01-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cur/2017-01-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseMessage` | `string` | no |

## DescribeReportDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportDefinitions` | `List<ReportDefinition>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ModifyReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | yes |
| `ReportDefinition` | `ReportDefinition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutReportDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportDefinition` | `ReportDefinition` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


