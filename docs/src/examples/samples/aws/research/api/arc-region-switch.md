# ARC - Region switch

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/arc-region-switch/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ApprovePlanExecutionStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `stepName` | `string` | yes |
| `approval` | `string` | yes |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `workflows` | `List<Workflow>` | yes |
| `executionRole` | `string` | yes |
| `recoveryTimeObjectiveMinutes` | `integer` | no |
| `associatedAlarms` | `Map<AssociatedAlarm>` | no |
| `triggers` | `List<Trigger>` | no |
| `reportConfiguration` | `ReportConfiguration` | no |
| `name` | `string` | yes |
| `regions` | `List<string>` | yes |
| `recoveryApproach` | `string` | yes |
| `primaryRegion` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plan` | `Plan` | no |

## DeletePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plan` | `Plan` | no |

## GetPlanEvaluationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `lastEvaluationTime` | `timestamp` | no |
| `lastEvaluatedVersion` | `string` | no |
| `region` | `string` | no |
| `evaluationState` | `string` | no |
| `warnings` | `List<ResourceWarning>` | no |
| `nextToken` | `string` | no |

## GetPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `version` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `comment` | `string` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | no |
| `mode` | `string` | yes |
| `executionState` | `string` | yes |
| `executionAction` | `string` | yes |
| `executionRegion` | `string` | yes |
| `recoveryExecutionId` | `string` | no |
| `stepStates` | `List<StepState>` | no |
| `plan` | `Plan` | no |
| `actualRecoveryTime` | `string` | no |
| `generatedReportDetails` | `List<GeneratedReport>` | no |
| `nextToken` | `string` | no |

## GetPlanInRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plan` | `Plan` | no |

## ListPlanExecutionEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ExecutionEvent>` | no |
| `nextToken` | `string` | no |

## ListPlanExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AbbreviatedExecution>` | no |
| `nextToken` | `string` | no |

## ListPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plans` | `List<AbbreviatedPlan>` | no |
| `nextToken` | `string` | no |

## ListPlansInRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plans` | `List<AbbreviatedPlan>` | no |
| `nextToken` | `string` | no |

## ListRoute53HealthChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `hostedZoneId` | `string` | no |
| `recordName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `healthChecks` | `List<Route53HealthCheck>` | no |
| `nextToken` | `string` | no |

## ListRoute53HealthChecksInRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `hostedZoneId` | `string` | no |
| `recordName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `healthChecks` | `List<Route53HealthCheck>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceTags` | `Map<string>` | no |

## StartPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `targetRegion` | `string` | yes |
| `action` | `string` | yes |
| `mode` | `string` | no |
| `comment` | `string` | no |
| `latestVersion` | `string` | no |
| `recoveryExecutionId` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |
| `plan` | `string` | no |
| `planVersion` | `string` | no |
| `activateRegion` | `string` | no |
| `deactivateRegion` | `string` | no |

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
| `resourceTagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `description` | `string` | no |
| `workflows` | `List<Workflow>` | yes |
| `executionRole` | `string` | yes |
| `recoveryTimeObjectiveMinutes` | `integer` | no |
| `associatedAlarms` | `Map<AssociatedAlarm>` | no |
| `triggers` | `List<Trigger>` | no |
| `reportConfiguration` | `ReportConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `plan` | `Plan` | no |

## UpdatePlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `action` | `string` | yes |
| `comment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePlanExecutionStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `planArn` | `string` | yes |
| `executionId` | `string` | yes |
| `comment` | `string` | yes |
| `stepName` | `string` | yes |
| `actionToTake` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


