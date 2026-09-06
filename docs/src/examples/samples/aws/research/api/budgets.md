# AWS Budgets

API version: 2016-10-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/budgets/2016-10-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `Budget` | `Budget` | yes |
| `NotificationsWithSubscribers` | `List<NotificationWithSubscribers>` | no |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBudgetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `NotificationType` | `string` | yes |
| `ActionType` | `string` | yes |
| `ActionThreshold` | `ActionThreshold` | yes |
| `Definition` | `Definition` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `ApprovalModel` | `string` | yes |
| `Subscribers` | `List<Subscriber>` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |

## CreateNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |
| `Subscribers` | `List<Subscriber>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |
| `Subscriber` | `Subscriber` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBudgetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Action` | `Action` | yes |

## DeleteNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |
| `Subscriber` | `Subscriber` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ShowFilterExpression` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Budget` | `Budget` | no |

## DescribeBudgetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Action` | `Action` | yes |

## DescribeBudgetActionHistories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |
| `TimePeriod` | `TimePeriod` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionHistories` | `List<ActionHistory>` | yes |
| `NextToken` | `string` | no |

## DescribeBudgetActionsForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<Action>` | yes |
| `NextToken` | `string` | no |

## DescribeBudgetActionsForBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<Action>` | yes |
| `NextToken` | `string` | no |

## DescribeBudgetNotificationsForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BudgetNotificationsForAccount` | `List<BudgetNotificationsForAccount>` | no |
| `NextToken` | `string` | no |

## DescribeBudgetPerformanceHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `TimePeriod` | `TimePeriod` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BudgetPerformanceHistory` | `BudgetPerformanceHistory` | no |
| `NextToken` | `string` | no |

## DescribeBudgets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ShowFilterExpression` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Budgets` | `List<Budget>` | no |
| `NextToken` | `string` | no |

## DescribeNotificationsForBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Notifications` | `List<Notification>` | no |
| `NextToken` | `string` | no |

## DescribeSubscribersForNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscribers` | `List<Subscriber>` | no |
| `NextToken` | `string` | no |

## ExecuteBudgetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |
| `ExecutionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |
| `ExecutionType` | `string` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTags` | `List<ResourceTag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `ResourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `NewBudget` | `Budget` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBudgetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `ActionId` | `string` | yes |
| `NotificationType` | `string` | no |
| `ActionThreshold` | `ActionThreshold` | no |
| `Definition` | `Definition` | no |
| `ExecutionRoleArn` | `string` | no |
| `ApprovalModel` | `string` | no |
| `Subscribers` | `List<Subscriber>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `OldAction` | `Action` | yes |
| `NewAction` | `Action` | yes |

## UpdateNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `OldNotification` | `Notification` | yes |
| `NewNotification` | `Notification` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `BudgetName` | `string` | yes |
| `Notification` | `Notification` | yes |
| `OldSubscriber` | `Subscriber` | yes |
| `NewSubscriber` | `Subscriber` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


