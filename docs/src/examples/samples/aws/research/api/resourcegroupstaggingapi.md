# AWS Resource Groups Tagging API

API version: 2017-01-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/resourcegroupstaggingapi/2017-01-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeReportCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `S3Location` | `string` | no |
| `ErrorMessage` | `string` | no |

## GetComplianceSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetIdFilters` | `List<string>` | no |
| `RegionFilters` | `List<string>` | no |
| `ResourceTypeFilters` | `List<string>` | no |
| `TagKeyFilters` | `List<string>` | no |
| `GroupBy` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `PaginationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryList` | `List<Summary>` | no |
| `PaginationToken` | `string` | no |

## GetResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |
| `TagFilters` | `List<TagFilter>` | no |
| `ResourcesPerPage` | `integer` | no |
| `TagsPerPage` | `integer` | no |
| `ResourceTypeFilters` | `List<string>` | no |
| `IncludeComplianceDetails` | `boolean` | no |
| `ExcludeCompliantResources` | `boolean` | no |
| `ResourceARNList` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |
| `ResourceTagMappingList` | `List<ResourceTagMapping>` | no |

## GetTagKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |
| `TagKeys` | `List<string>` | no |

## GetTagValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |
| `Key` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PaginationToken` | `string` | no |
| `TagValues` | `List<string>` | no |

## ListRequiredTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequiredTags` | `List<RequiredTag>` | no |
| `NextToken` | `string` | no |

## StartReportCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3Bucket` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARNList` | `List<string>` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedResourcesMap` | `Map<FailureInfo>` | no |

## UntagResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARNList` | `List<string>` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedResourcesMap` | `Map<FailureInfo>` | no |

