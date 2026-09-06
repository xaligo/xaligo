# AWS Savings Plans

API version: 2019-06-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/savingsplans/2019-06-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSavingsPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanOfferingId` | `string` | yes |
| `commitment` | `string` | yes |
| `upfrontPaymentAmount` | `string` | no |
| `purchaseTime` | `timestamp` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | no |

## DeleteQueuedSavingsPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeSavingsPlanRates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | yes |
| `filters` | `List<SavingsPlanRateFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | no |
| `searchResults` | `List<SavingsPlanRate>` | no |
| `nextToken` | `string` | no |

## DescribeSavingsPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanArns` | `List<string>` | no |
| `savingsPlanIds` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `states` | `List<string>` | no |
| `filters` | `List<SavingsPlanFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlans` | `List<SavingsPlan>` | no |
| `nextToken` | `string` | no |

## DescribeSavingsPlansOfferingRates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanOfferingIds` | `List<string>` | no |
| `savingsPlanPaymentOptions` | `List<string>` | no |
| `savingsPlanTypes` | `List<string>` | no |
| `products` | `List<string>` | no |
| `serviceCodes` | `List<string>` | no |
| `usageTypes` | `List<string>` | no |
| `operations` | `List<string>` | no |
| `filters` | `List<SavingsPlanOfferingRateFilterElement>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchResults` | `List<SavingsPlanOfferingRate>` | no |
| `nextToken` | `string` | no |

## DescribeSavingsPlansOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringIds` | `List<string>` | no |
| `paymentOptions` | `List<string>` | no |
| `productType` | `string` | no |
| `planTypes` | `List<string>` | no |
| `durations` | `List<long>` | no |
| `currencies` | `List<string>` | no |
| `descriptions` | `List<string>` | no |
| `serviceCodes` | `List<string>` | no |
| `usageTypes` | `List<string>` | no |
| `operations` | `List<string>` | no |
| `filters` | `List<SavingsPlanOfferingFilterElement>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchResults` | `List<SavingsPlanOffering>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ReturnSavingsPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `savingsPlanId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


