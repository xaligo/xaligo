# AWS Billing and Cost Management Pricing Calculator

API version: 2024-06-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bcm-pricing-calculator/2024-06-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateBillScenarioCommitmentModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `commitmentModifications` | `List<BatchCreateBillScenarioCommitmentModificationEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BatchCreateBillScenarioCommitmentModificationItem>` | no |
| `errors` | `List<BatchCreateBillScenarioCommitmentModificationError>` | no |

## BatchCreateBillScenarioUsageModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `usageModifications` | `List<BatchCreateBillScenarioUsageModificationEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BatchCreateBillScenarioUsageModificationItem>` | no |
| `errors` | `List<BatchCreateBillScenarioUsageModificationError>` | no |

## BatchCreateWorkloadEstimateUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadEstimateId` | `string` | yes |
| `usage` | `List<BatchCreateWorkloadEstimateUsageEntry>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BatchCreateWorkloadEstimateUsageItem>` | no |
| `errors` | `List<BatchCreateWorkloadEstimateUsageError>` | no |

## BatchDeleteBillScenarioCommitmentModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteBillScenarioCommitmentModificationError>` | no |

## BatchDeleteBillScenarioUsageModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteBillScenarioUsageModificationError>` | no |

## BatchDeleteWorkloadEstimateUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadEstimateId` | `string` | yes |
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteWorkloadEstimateUsageError>` | no |

## BatchUpdateBillScenarioCommitmentModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `commitmentModifications` | `List<BatchUpdateBillScenarioCommitmentModificationEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillScenarioCommitmentModificationItem>` | no |
| `errors` | `List<BatchUpdateBillScenarioCommitmentModificationError>` | no |

## BatchUpdateBillScenarioUsageModification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `usageModifications` | `List<BatchUpdateBillScenarioUsageModificationEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillScenarioUsageModificationItem>` | no |
| `errors` | `List<BatchUpdateBillScenarioUsageModificationError>` | no |

## BatchUpdateWorkloadEstimateUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadEstimateId` | `string` | yes |
| `usage` | `List<BatchUpdateWorkloadEstimateUsageEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<WorkloadEstimateUsageItem>` | no |
| `errors` | `List<BatchUpdateWorkloadEstimateUsageError>` | no |

## CreateBillEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | no |
| `failureMessage` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `costSummary` | `BillEstimateCostSummary` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |
| `costCategoryGroupSharingPreferenceEffectiveDate` | `timestamp` | no |

## CreateBillScenario

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `failureMessage` | `string` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |

## CreateWorkloadEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `clientToken` | `string` | no |
| `rateType` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `rateType` | `string` | no |
| `rateTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `totalCost` | `double` | no |
| `costCurrency` | `string` | no |
| `failureMessage` | `string` | no |

## DeleteBillEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBillScenario

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkloadEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetBillEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | no |
| `failureMessage` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `costSummary` | `BillEstimateCostSummary` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |
| `costCategoryGroupSharingPreferenceEffectiveDate` | `timestamp` | no |

## GetBillScenario

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `failureMessage` | `string` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |

## GetPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managementAccountRateTypeSelections` | `List<string>` | no |
| `memberAccountRateTypeSelections` | `List<string>` | no |
| `standaloneAccountRateTypeSelections` | `List<string>` | no |

## GetWorkloadEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `rateType` | `string` | no |
| `rateTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `totalCost` | `double` | no |
| `costCurrency` | `string` | no |
| `failureMessage` | `string` | no |

## ListBillEstimateCommitments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billEstimateId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillEstimateCommitmentSummary>` | no |
| `nextToken` | `string` | no |

## ListBillEstimateInputCommitmentModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billEstimateId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillEstimateInputCommitmentModificationSummary>` | no |
| `nextToken` | `string` | no |

## ListBillEstimateInputUsageModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billEstimateId` | `string` | yes |
| `filters` | `List<ListUsageFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillEstimateInputUsageModificationSummary>` | no |
| `nextToken` | `string` | no |

## ListBillEstimateLineItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billEstimateId` | `string` | yes |
| `filters` | `List<ListBillEstimateLineItemsFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillEstimateLineItemSummary>` | no |
| `nextToken` | `string` | no |

## ListBillEstimates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<ListBillEstimatesFilter>` | no |
| `createdAtFilter` | `FilterTimestamp` | no |
| `expiresAtFilter` | `FilterTimestamp` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillEstimateSummary>` | no |
| `nextToken` | `string` | no |

## ListBillScenarioCommitmentModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillScenarioCommitmentModificationItem>` | no |
| `nextToken` | `string` | no |

## ListBillScenarioUsageModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `billScenarioId` | `string` | yes |
| `filters` | `List<ListUsageFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillScenarioUsageModificationItem>` | no |
| `nextToken` | `string` | no |

## ListBillScenarios

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<ListBillScenariosFilter>` | no |
| `createdAtFilter` | `FilterTimestamp` | no |
| `expiresAtFilter` | `FilterTimestamp` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<BillScenarioSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListWorkloadEstimateUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadEstimateId` | `string` | yes |
| `filters` | `List<ListUsageFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<WorkloadEstimateUsageItem>` | no |
| `nextToken` | `string` | no |

## ListWorkloadEstimates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAtFilter` | `FilterTimestamp` | no |
| `expiresAtFilter` | `FilterTimestamp` | no |
| `filters` | `List<ListWorkloadEstimatesFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<WorkloadEstimateSummary>` | no |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBillEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `expiresAt` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | no |
| `failureMessage` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `costSummary` | `BillEstimateCostSummary` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |
| `costCategoryGroupSharingPreferenceEffectiveDate` | `timestamp` | no |

## UpdateBillScenario

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `expiresAt` | `timestamp` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `billInterval` | `BillInterval` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `failureMessage` | `string` | no |
| `groupSharingPreference` | `string` | no |
| `costCategoryGroupSharingPreferenceArn` | `string` | no |

## UpdatePreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managementAccountRateTypeSelections` | `List<string>` | no |
| `memberAccountRateTypeSelections` | `List<string>` | no |
| `standaloneAccountRateTypeSelections` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managementAccountRateTypeSelections` | `List<string>` | no |
| `memberAccountRateTypeSelections` | `List<string>` | no |
| `standaloneAccountRateTypeSelections` | `List<string>` | no |

## UpdateWorkloadEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `expiresAt` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `createdAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |
| `rateType` | `string` | no |
| `rateTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `totalCost` | `double` | no |
| `costCurrency` | `string` | no |
| `failureMessage` | `string` | no |

