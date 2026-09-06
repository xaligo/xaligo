# Amazon Kendra Intelligent Ranking

API version: 2022-10-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kendra-ranking/2022-10-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateRescoreExecutionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `CapacityUnits` | `CapacityUnitsConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |

## DeleteRescoreExecutionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeRescoreExecutionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CapacityUnits` | `CapacityUnitsConfiguration` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `ErrorMessage` | `string` | no |

## ListRescoreExecutionPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryItems` | `List<RescoreExecutionPlanSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## Rescore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RescoreExecutionPlanId` | `string` | yes |
| `SearchQuery` | `string` | yes |
| `Documents` | `List<Document>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RescoreId` | `string` | no |
| `ResultItems` | `List<RescoreResultItem>` | no |

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


## UpdateRescoreExecutionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CapacityUnits` | `CapacityUnitsConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


