# AWS Free Tier

API version: 2023-09-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/freetier/2023-09-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetAccountActivity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityId` | `string` | yes |
| `languageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityId` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | yes |
| `status` | `string` | yes |
| `instructionsUrl` | `string` | yes |
| `reward` | `ActivityReward` | yes |
| `estimatedTimeToCompleteInMinutes` | `integer` | no |
| `expiresAt` | `timestamp` | no |
| `startedAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |

## GetAccountPlanState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `accountPlanType` | `string` | yes |
| `accountPlanStatus` | `string` | yes |
| `accountPlanRemainingCredits` | `MonetaryAmount` | no |
| `accountPlanExpirationDate` | `timestamp` | no |

## GetFreeTierUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `Expression` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `freeTierUsages` | `List<FreeTierUsage>` | yes |
| `nextToken` | `string` | no |

## ListAccountActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterActivityStatuses` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `languageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activities` | `List<ActivitySummary>` | yes |
| `nextToken` | `string` | no |

## UpgradeAccountPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountPlanType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `accountPlanType` | `string` | yes |
| `accountPlanStatus` | `string` | yes |

