# Compute Optimizer Automation

API version: 2025-09-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/compute-optimizer-automation/2025-09-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `errors` | `List<string>` | no |

## CreateAutomationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `ruleType` | `string` | yes |
| `organizationConfiguration` | `OrganizationConfiguration` | no |
| `priority` | `string` | no |
| `recommendedActionTypes` | `List<string>` | yes |
| `criteria` | `Criteria` | no |
| `schedule` | `Schedule` | yes |
| `status` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | no |
| `ruleId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `ruleType` | `string` | no |
| `ruleRevision` | `long` | no |
| `organizationConfiguration` | `OrganizationConfiguration` | no |
| `priority` | `string` | no |
| `recommendedActionTypes` | `List<string>` | no |
| `criteria` | `Criteria` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `tags` | `List<Tag>` | no |
| `createdTimestamp` | `timestamp` | no |

## DeleteAutomationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | yes |
| `ruleRevision` | `long` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | no |
| `errors` | `List<string>` | no |

## GetAutomationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | no |
| `eventDescription` | `string` | no |
| `eventType` | `string` | no |
| `eventStatus` | `string` | no |
| `eventStatusReason` | `string` | no |
| `resourceArn` | `string` | no |
| `resourceId` | `string` | no |
| `recommendedActionId` | `string` | no |
| `accountId` | `string` | no |
| `region` | `string` | no |
| `ruleId` | `string` | no |
| `resourceType` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `completedTimestamp` | `timestamp` | no |
| `estimatedMonthlySavings` | `EstimatedMonthlySavings` | no |

## GetAutomationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | no |
| `ruleId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `ruleType` | `string` | no |
| `ruleRevision` | `long` | no |
| `accountId` | `string` | no |
| `organizationConfiguration` | `OrganizationConfiguration` | no |
| `priority` | `string` | no |
| `recommendedActionTypes` | `List<string>` | no |
| `criteria` | `Criteria` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `tags` | `List<Tag>` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastUpdatedTimestamp` | `timestamp` | no |

## GetEnrollmentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `organizationRuleMode` | `string` | no |
| `lastUpdatedTimestamp` | `timestamp` | no |

## ListAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<AccountInfo>` | yes |
| `nextToken` | `string` | no |

## ListAutomationEventSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automationEventSteps` | `List<AutomationEventStep>` | no |
| `nextToken` | `string` | no |

## ListAutomationEventSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<AutomationEventFilter>` | no |
| `startDateInclusive` | `string` | no |
| `endDateExclusive` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automationEventSummaries` | `List<AutomationEventSummary>` | no |
| `nextToken` | `string` | no |

## ListAutomationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<AutomationEventFilter>` | no |
| `startTimeInclusive` | `timestamp` | no |
| `endTimeExclusive` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automationEvents` | `List<AutomationEvent>` | no |
| `nextToken` | `string` | no |

## ListAutomationRulePreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleType` | `string` | yes |
| `organizationScope` | `OrganizationScope` | no |
| `recommendedActionTypes` | `List<string>` | yes |
| `criteria` | `Criteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `previewResults` | `List<PreviewResult>` | no |
| `nextToken` | `string` | no |

## ListAutomationRulePreviewSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleType` | `string` | yes |
| `organizationScope` | `OrganizationScope` | no |
| `recommendedActionTypes` | `List<string>` | yes |
| `criteria` | `Criteria` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `previewResultSummaries` | `List<PreviewResultSummary>` | no |
| `nextToken` | `string` | no |

## ListAutomationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automationRules` | `List<AutomationRule>` | no |
| `nextToken` | `string` | no |

## ListRecommendedActionSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<RecommendedActionFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedActionSummaries` | `List<RecommendedActionSummary>` | no |
| `nextToken` | `string` | no |

## ListRecommendedActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<RecommendedActionFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedActions` | `List<RecommendedAction>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## RollbackAutomationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | no |
| `eventStatus` | `string` | no |

## StartAutomationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedActionId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendedActionId` | `string` | no |
| `eventId` | `string` | no |
| `eventStatus` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `ruleRevision` | `long` | yes |
| `tags` | `List<Tag>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `ruleRevision` | `long` | yes |
| `tagKeys` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAutomationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | yes |
| `ruleRevision` | `long` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `ruleType` | `string` | no |
| `organizationConfiguration` | `OrganizationConfiguration` | no |
| `priority` | `string` | no |
| `recommendedActionTypes` | `List<string>` | no |
| `criteria` | `Criteria` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleArn` | `string` | no |
| `ruleRevision` | `long` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `ruleType` | `string` | no |
| `organizationConfiguration` | `OrganizationConfiguration` | no |
| `priority` | `string` | no |
| `recommendedActionTypes` | `List<string>` | no |
| `criteria` | `Criteria` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `createdTimestamp` | `timestamp` | no |
| `lastUpdatedTimestamp` | `timestamp` | no |

## UpdateEnrollmentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `lastUpdatedTimestamp` | `timestamp` | yes |

