# AWS Billing and Cost Management Dashboards

API version: 2025-08-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bcm-dashboards/2025-08-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `widgets` | `List<Widget>` | yes |
| `resourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## CreateScheduledReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledReport` | `ScheduledReportInput` | yes |
| `resourceTags` | `List<ResourceTag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DeleteDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DeleteScheduledReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## ExecuteScheduledReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `healthStatus` | `HealthStatus` | no |
| `executionTriggered` | `boolean` | no |

## GetDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `widgets` | `List<Widget>` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policyDocument` | `string` | yes |

## GetScheduledReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledReport` | `ScheduledReport` | yes |

## ListDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dashboards` | `List<DashboardReference>` | yes |
| `nextToken` | `string` | no |

## ListScheduledReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledReports` | `List<ScheduledReportSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceTags` | `List<ResourceTag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `resourceTags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `resourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `widgets` | `List<Widget>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## UpdateScheduledReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `dashboardArn` | `string` | no |
| `scheduledReportExecutionRoleArn` | `string` | no |
| `scheduleConfig` | `ScheduleConfig` | no |
| `widgetIds` | `List<string>` | no |
| `widgetDateRangeOverride` | `DateTimeRange` | no |
| `clearWidgetIds` | `boolean` | no |
| `clearWidgetDateRangeOverride` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

