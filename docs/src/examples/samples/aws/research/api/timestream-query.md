# Amazon Timestream Query

API version: 2018-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/timestream-query/2018-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CancellationMessage` | `string` | no |

## CreateScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `QueryString` | `string` | yes |
| `ScheduleConfiguration` | `ScheduleConfiguration` | yes |
| `NotificationConfiguration` | `NotificationConfiguration` | yes |
| `TargetConfiguration` | `TargetConfiguration` | no |
| `ClientToken` | `string` | no |
| `ScheduledQueryExecutionRoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `ErrorReportConfiguration` | `ErrorReportConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

## DeleteScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQueryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxQueryTCU` | `integer` | no |
| `QueryPricingModel` | `string` | no |
| `QueryCompute` | `QueryComputeResponse` | no |

## DescribeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | yes |

## DescribeScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQueryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQuery` | `ScheduledQueryDescription` | yes |

## ExecuteScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQueryArn` | `string` | yes |
| `InvocationTime` | `timestamp` | yes |
| `ClientToken` | `string` | no |
| `QueryInsights` | `ScheduledQueryInsights` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListScheduledQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQueries` | `List<ScheduledQuery>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `NextToken` | `string` | no |

## PrepareQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryString` | `string` | yes |
| `ValidateOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryString` | `string` | yes |
| `Columns` | `List<SelectColumn>` | yes |
| `Parameters` | `List<ParameterMapping>` | yes |

## Query

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryString` | `string` | yes |
| `ClientToken` | `string` | no |
| `NextToken` | `string` | no |
| `MaxRows` | `integer` | no |
| `QueryInsights` | `QueryInsights` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |
| `NextToken` | `string` | no |
| `Rows` | `List<Row>` | yes |
| `ColumnInfo` | `List<ColumnInfo>` | yes |
| `QueryStatus` | `QueryStatus` | no |
| `QueryInsightsResponse` | `QueryInsightsResponse` | no |

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


## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxQueryTCU` | `integer` | no |
| `QueryPricingModel` | `string` | no |
| `QueryCompute` | `QueryComputeRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxQueryTCU` | `integer` | no |
| `QueryPricingModel` | `string` | no |
| `QueryCompute` | `QueryComputeResponse` | no |

## UpdateScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledQueryArn` | `string` | yes |
| `State` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


