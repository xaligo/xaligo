# AWS Billing

API version: 2023-09-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/billing/2023-09-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateSourceViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `sourceViews` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## CreateBillingView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `sourceViews` | `List<string>` | yes |
| `dataFilterExpression` | `Expression` | no |
| `clientToken` | `string` | no |
| `resourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | no |

## DeleteBillingView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DisassociateSourceViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `sourceViews` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## GetBillingPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `features` | `List<string>` | yes |
| `filters` | `List<BillingFeatureFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingPreferences` | `List<BillingPreferenceSummary>` | yes |
| `nextToken` | `string` | no |

## GetBillingView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingView` | `BillingViewElement` | yes |

## GetCreditAllocationHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `creditId` | `long` | no |
| `startDate` | `timestamp` | yes |
| `endDate` | `timestamp` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creditAllocationHistoryList` | `List<CreditAllocationHistoryEntry>` | no |
| `partialResults` | `boolean` | yes |
| `failedMonths` | `List<string>` | no |
| `nextToken` | `string` | no |

## GetCredits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `startDate` | `timestamp` | yes |
| `endDate` | `timestamp` | no |
| `payerAccountFlag` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credits` | `List<CreditData>` | no |

## GetEnterpriseSupportChargeSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingMonth` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payerAccountId` | `string` | yes |
| `billingMonth` | `string` | yes |
| `billingPeriodStartDate` | `timestamp` | yes |
| `billingPeriodEndDate` | `timestamp` | yes |
| `isEstimated` | `boolean` | yes |
| `billDate` | `timestamp` | yes |
| `supportCharge` | `string` | yes |
| `totalSupportCharge` | `string` | yes |
| `supportDiscount` | `string` | yes |
| `totalSupportEligibleSpend` | `string` | yes |
| `totalSupportEligibleUsageSpend` | `string` | yes |
| `totalSupportEligibleReservedInstanceSpend` | `string` | yes |
| `totalSupportEligibleSavingsPlanSpend` | `string` | yes |
| `supportChargePercentage` | `string` | yes |
| `supportEffectivePricingPlan` | `PricingPlan` | yes |

## GetEnterpriseSupportContractDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingMonth` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isContractActive` | `boolean` | no |
| `supportAllocationMethod` | `string` | yes |
| `supportReservedInstanceAmortizationStartDate` | `timestamp` | no |
| `supportReservedInstanceTreatmentMethod` | `string` | no |
| `supportSavingsPlansAmortizationStartDate` | `timestamp` | no |
| `supportSavingsPlansTreatmentMethod` | `string` | no |
| `supportProrateStartDate` | `timestamp` | no |
| `contractPayerAccountIds` | `List<ContractAccount>` | yes |
| `chargedPayerAccountIds` | `List<ChargeAccount>` | yes |
| `additionalSupportCharge` | `List<AdditionalCharge>` | no |
| `additionalSupportEligibleUsageSpend` | `List<AdditionalCharge>` | no |
| `pricingPlans` | `List<PricingPlan>` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | no |

## ListBillingViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activeTimeRange` | `ActiveTimeRange` | no |
| `arns` | `List<string>` | no |
| `billingViewTypes` | `List<string>` | no |
| `names` | `List<StringSearch>` | no |
| `ownerAccountId` | `string` | no |
| `sourceAccountId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingViews` | `List<BillingViewListElement>` | yes |
| `nextToken` | `string` | no |

## ListEnterpriseSupportLinkedAccountCharges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billingMonth` | `string` | yes |
| `accountId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkedAccount` | `List<LinkedAccountCharge>` | yes |
| `nextToken` | `string` | no |

## ListSourceViewsForBillingView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceViews` | `List<string>` | yes |
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

## RedeemCredits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promoCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateBillingPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `feature` | `string` | yes |
| `billingPreferencesPerKey` | `List<BillingPreferenceForKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBillingView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `dataFilterExpression` | `Expression` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `updatedAt` | `timestamp` | no |

