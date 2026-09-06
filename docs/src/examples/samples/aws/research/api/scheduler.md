# Amazon EventBridge Scheduler

API version: 2021-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/scheduler/2021-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionAfterCompletion` | `string` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `EndDate` | `timestamp` | no |
| `FlexibleTimeWindow` | `FlexibleTimeWindow` | yes |
| `GroupName` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `Name` | `string` | yes |
| `ScheduleExpression` | `string` | yes |
| `ScheduleExpressionTimezone` | `string` | no |
| `StartDate` | `timestamp` | no |
| `State` | `string` | no |
| `Target` | `Target` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleArn` | `string` | yes |

## CreateScheduleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleGroupArn` | `string` | yes |

## DeleteSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `GroupName` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionAfterCompletion` | `string` | no |
| `Arn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `Description` | `string` | no |
| `EndDate` | `timestamp` | no |
| `FlexibleTimeWindow` | `FlexibleTimeWindow` | no |
| `GroupName` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `LastModificationDate` | `timestamp` | no |
| `Name` | `string` | no |
| `ScheduleExpression` | `string` | no |
| `ScheduleExpressionTimezone` | `string` | no |
| `StartDate` | `timestamp` | no |
| `State` | `string` | no |
| `Target` | `Target` | no |

## GetScheduleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `LastModificationDate` | `timestamp` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## ListScheduleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NamePrefix` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ScheduleGroups` | `List<ScheduleGroupSummary>` | yes |

## ListSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `State` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Schedules` | `List<ScheduleSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionAfterCompletion` | `string` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `EndDate` | `timestamp` | no |
| `FlexibleTimeWindow` | `FlexibleTimeWindow` | yes |
| `GroupName` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `Name` | `string` | yes |
| `ScheduleExpression` | `string` | yes |
| `ScheduleExpressionTimezone` | `string` | no |
| `StartDate` | `timestamp` | no |
| `State` | `string` | no |
| `Target` | `Target` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduleArn` | `string` | yes |

