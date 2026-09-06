# AWS CodeStar Notifications

API version: 2019-10-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codestar-notifications/2019-10-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateNotificationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventTypeIds` | `List<string>` | yes |
| `Resource` | `string` | yes |
| `Targets` | `List<Target>` | yes |
| `DetailType` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeleteNotificationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeleteTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetAddress` | `string` | yes |
| `ForceUnsubscribeAll` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeNotificationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `EventTypes` | `List<EventTypeSummary>` | no |
| `Resource` | `string` | no |
| `Targets` | `List<TargetSummary>` | no |
| `DetailType` | `string` | no |
| `CreatedBy` | `string` | no |
| `Status` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastModifiedTimestamp` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## ListEventTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ListEventTypesFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventTypes` | `List<EventTypeSummary>` | no |
| `NextToken` | `string` | no |

## ListNotificationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ListNotificationRulesFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NotificationRules` | `List<NotificationRuleSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ListTargetsFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<TargetSummary>` | no |
| `NextToken` | `string` | no |

## Subscribe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Target` | `Target` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## Unsubscribe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `TargetAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotificationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `EventTypeIds` | `List<string>` | no |
| `Targets` | `List<Target>` | no |
| `DetailType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


