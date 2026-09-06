# Nova Act Service

API version: 2025-08-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/nova-act/2025-08-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `sessionId` | `string` | yes |
| `task` | `string` | yes |
| `toolSpecs` | `List<ToolSpec>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actId` | `string` | yes |
| `status` | `string` | yes |

## CreateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |

## CreateWorkflowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `exportConfig` | `WorkflowExportConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## CreateWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `modelId` | `string` | yes |
| `clientToken` | `string` | no |
| `logGroupName` | `string` | no |
| `clientInfo` | `ClientInfo` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowRunId` | `string` | yes |
| `status` | `string` | yes |

## DeleteWorkflowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## GetWorkflowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `exportConfig` | `WorkflowExportConfig` | no |
| `status` | `string` | yes |

## GetWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowRunArn` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `status` | `string` | yes |
| `startedAt` | `timestamp` | yes |
| `endedAt` | `timestamp` | no |
| `modelId` | `string` | yes |
| `logGroupName` | `string` | no |

## InvokeActStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `sessionId` | `string` | yes |
| `actId` | `string` | yes |
| `callResults` | `List<CallResult>` | yes |
| `previousStepId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `calls` | `List<Call>` | yes |
| `stepId` | `string` | yes |

## ListActs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | no |
| `sessionId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actSummaries` | `List<ActSummary>` | yes |
| `nextToken` | `string` | no |

## ListModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCompatibilityVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelSummaries` | `List<ModelSummary>` | yes |
| `modelAliases` | `List<ModelAlias>` | yes |
| `compatibilityInformation` | `CompatibilityInformation` | yes |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionSummaries` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListWorkflowDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionSummaries` | `List<WorkflowDefinitionSummary>` | yes |
| `nextToken` | `string` | no |

## ListWorkflowRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowRunSummaries` | `List<WorkflowRunSummary>` | yes |
| `nextToken` | `string` | no |

## UpdateAct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `sessionId` | `string` | yes |
| `actId` | `string` | yes |
| `status` | `string` | yes |
| `error` | `ActError` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowDefinitionName` | `string` | yes |
| `workflowRunId` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


