# PricingPlanManager

API version: 2025-08-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pricing-plan-manager/2025-08-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ApprovePaidSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## AssociateResourcesToSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `resourceArns` | `List<string>` | yes |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## CancelSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## CancelSubscriptionChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## CreateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planFamily` | `string` | yes |
| `planTier` | `string` | yes |
| `usageLevel` | `string` | no |
| `resourceArns` | `List<string>` | yes |
| `approvalMode` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## DisassociateResourcesFromSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `resourceArns` | `List<string>` | yes |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## GetSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionSummaries` | `List<SubscriptionSummary>` | yes |
| `nextToken` | `string` | no |

## UpdateSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `planTier` | `string` | yes |
| `usageLevel` | `string` | no |
| `ifMatch` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `Subscription` | yes |
| `eTag` | `string` | yes |

