# AWSBillingConductor

API version: 2021-07-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/billingconductor/2021-07-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## AssociatePricingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `PricingRuleArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## BatchAssociateResourcesToCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |
| `BillingPeriodRange` | `CustomLineItemBillingPeriodRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfullyAssociatedResources` | `List<AssociateResourceResponseElement>` | no |
| `FailedAssociatedResources` | `List<AssociateResourceResponseElement>` | no |

## BatchDisassociateResourcesFromCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |
| `BillingPeriodRange` | `CustomLineItemBillingPeriodRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfullyDisassociatedResources` | `List<DisassociateResourceResponseElement>` | no |
| `FailedDisassociatedResources` | `List<DisassociateResourceResponseElement>` | no |

## CreateBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `AccountGrouping` | `AccountGrouping` | yes |
| `ComputationPreference` | `ComputationPreference` | yes |
| `PrimaryAccountId` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreateCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `BillingGroupArn` | `string` | yes |
| `BillingPeriodRange` | `CustomLineItemBillingPeriodRange` | no |
| `Tags` | `Map<string>` | no |
| `ChargeDetails` | `CustomLineItemChargeDetails` | yes |
| `AccountId` | `string` | no |
| `ComputationRule` | `string` | no |
| `PresentationDetails` | `PresentationObject` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreatePricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `PricingRuleArns` | `List<string>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreatePricingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Scope` | `string` | yes |
| `Type` | `string` | yes |
| `ModifierPercentage` | `double` | no |
| `Service` | `string` | no |
| `Tags` | `Map<string>` | no |
| `BillingEntity` | `string` | no |
| `Tiering` | `CreateTieringInput` | no |
| `UsageType` | `string` | no |
| `Operation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeleteBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeleteCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `BillingPeriodRange` | `CustomLineItemBillingPeriodRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeletePricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeletePricingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DisassociateAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DisassociatePricingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `PricingRuleArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## GetBillingGroupCostReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `BillingPeriodRange` | `BillingPeriodRange` | no |
| `GroupBy` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingGroupCostReportResults` | `List<BillingGroupCostReportResultElement>` | no |
| `NextToken` | `string` | no |

## ListAccountAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `Filters` | `ListAccountAssociationsFilter` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkedAccounts` | `List<AccountAssociationsListElement>` | no |
| `NextToken` | `string` | no |

## ListBillingGroupCostReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `ListBillingGroupCostReportsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingGroupCostReports` | `List<BillingGroupCostReportElement>` | no |
| `NextToken` | `string` | no |

## ListBillingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `ListBillingGroupsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingGroups` | `List<BillingGroupListElement>` | no |
| `NextToken` | `string` | no |

## ListCustomLineItemVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `ListCustomLineItemVersionsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomLineItemVersions` | `List<CustomLineItemVersionListElement>` | no |
| `NextToken` | `string` | no |

## ListCustomLineItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `ListCustomLineItemsFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomLineItems` | `List<CustomLineItemListElement>` | no |
| `NextToken` | `string` | no |

## ListPricingPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `Filters` | `ListPricingPlansFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingPlans` | `List<PricingPlanListElement>` | no |
| `NextToken` | `string` | no |

## ListPricingPlansAssociatedWithPricingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingRuleArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingRuleArn` | `string` | no |
| `PricingPlanArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPricingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `Filters` | `ListPricingRulesFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingRules` | `List<PricingRuleListElement>` | no |
| `NextToken` | `string` | no |

## ListPricingRulesAssociatedToPricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingPlanArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `PricingPlanArn` | `string` | no |
| `PricingRuleArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListResourcesAssociatedToCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | no |
| `Arn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `ListResourcesAssociatedToCustomLineItemFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssociatedResources` | `List<ListResourcesAssociatedToCustomLineItemResponseElement>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateBillingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `ComputationPreference` | `ComputationPreference` | no |
| `Description` | `string` | no |
| `AccountGrouping` | `UpdateBillingGroupAccountGrouping` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `PrimaryAccountId` | `string` | no |
| `PricingPlanArn` | `string` | no |
| `Size` | `long` | no |
| `LastModifiedTime` | `long` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `AccountGrouping` | `UpdateBillingGroupAccountGrouping` | no |

## UpdateCustomLineItem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ChargeDetails` | `UpdateCustomLineItemChargeDetails` | no |
| `BillingPeriodRange` | `CustomLineItemBillingPeriodRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `BillingGroupArn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ChargeDetails` | `ListCustomLineItemChargeDetails` | no |
| `LastModifiedTime` | `long` | no |
| `AssociationSize` | `long` | no |

## UpdatePricingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Size` | `long` | no |
| `LastModifiedTime` | `long` | no |

## UpdatePricingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `ModifierPercentage` | `double` | no |
| `Tiering` | `UpdateTieringInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Scope` | `string` | no |
| `Type` | `string` | no |
| `ModifierPercentage` | `double` | no |
| `Service` | `string` | no |
| `AssociatedPricingPlanCount` | `long` | no |
| `LastModifiedTime` | `long` | no |
| `BillingEntity` | `string` | no |
| `Tiering` | `UpdateTieringInput` | no |
| `UsageType` | `string` | no |
| `Operation` | `string` | no |

