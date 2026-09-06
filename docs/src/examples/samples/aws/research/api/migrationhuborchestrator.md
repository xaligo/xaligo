# AWS Migration Hub Orchestrator

API version: 2021-08-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/migrationhuborchestrator/2021-08-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateDescription` | `string` | no |
| `templateSource` | `TemplateSource` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | no |
| `templateArn` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `templateId` | `string` | yes |
| `applicationConfigurationId` | `string` | no |
| `inputParameters` | `Map<StepInput>` | yes |
| `stepTargets` | `List<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `templateId` | `string` | no |
| `adsApplicationConfigurationId` | `string` | no |
| `workflowInputs` | `Map<StepInput>` | no |
| `stepTargets` | `List<string>` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## CreateWorkflowStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `stepGroupId` | `string` | yes |
| `workflowId` | `string` | yes |
| `stepActionType` | `string` | yes |
| `description` | `string` | no |
| `workflowStepAutomationConfiguration` | `WorkflowStepAutomationConfiguration` | no |
| `stepTarget` | `List<string>` | no |
| `outputs` | `List<WorkflowStepOutput>` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `stepGroupId` | `string` | no |
| `workflowId` | `string` | no |
| `name` | `string` | no |

## CreateWorkflowStepGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `next` | `List<string>` | no |
| `previous` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | no |
| `name` | `string` | no |
| `id` | `string` | no |
| `description` | `string` | no |
| `tools` | `List<Tool>` | no |
| `next` | `List<string>` | no |
| `previous` | `List<string>` | no |
| `creationTime` | `timestamp` | no |

## DeleteTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |

## DeleteWorkflowStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `stepGroupId` | `string` | yes |
| `workflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflowStepGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `templateArn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `inputs` | `List<TemplateInput>` | no |
| `tools` | `List<Tool>` | no |
| `creationTime` | `timestamp` | no |
| `owner` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `templateClass` | `string` | no |
| `tags` | `Map<string>` | no |

## GetTemplateStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `templateId` | `string` | yes |
| `stepGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `stepGroupId` | `string` | no |
| `templateId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `stepActionType` | `string` | no |
| `creationTime` | `string` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |
| `outputs` | `List<StepOutput>` | no |
| `stepAutomationConfiguration` | `StepAutomationConfiguration` | no |

## GetTemplateStepGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `tools` | `List<Tool>` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `templateId` | `string` | no |
| `adsApplicationConfigurationId` | `string` | no |
| `adsApplicationName` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastStartTime` | `timestamp` | no |
| `lastStopTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `tools` | `List<Tool>` | no |
| `totalSteps` | `integer` | no |
| `completedSteps` | `integer` | no |
| `workflowInputs` | `Map<StepInput>` | no |
| `tags` | `Map<string>` | no |
| `workflowBucket` | `string` | no |

## GetWorkflowStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `stepGroupId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `stepGroupId` | `string` | no |
| `workflowId` | `string` | no |
| `stepId` | `string` | no |
| `description` | `string` | no |
| `stepActionType` | `string` | no |
| `owner` | `string` | no |
| `workflowStepAutomationConfiguration` | `WorkflowStepAutomationConfiguration` | no |
| `stepTarget` | `List<string>` | no |
| `outputs` | `List<WorkflowStepOutput>` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `scriptOutputLocation` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastStartTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `noOfSrvCompleted` | `integer` | no |
| `noOfSrvFailed` | `integer` | no |
| `totalNoOfSrv` | `integer` | no |

## GetWorkflowStepGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `workflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `workflowId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `owner` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `tools` | `List<Tool>` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |

## ListPlugins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `plugins` | `List<PluginSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTemplateStepGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `templateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templateStepGroupSummary` | `List<TemplateStepGroupSummary>` | yes |

## ListTemplateSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `templateId` | `string` | yes |
| `stepGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templateStepSummaryList` | `List<TemplateStepSummary>` | no |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templateSummary` | `List<TemplateSummary>` | yes |

## ListWorkflowStepGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `workflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `workflowStepGroupsSummary` | `List<WorkflowStepGroupSummary>` | yes |

## ListWorkflowSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `workflowId` | `string` | yes |
| `stepGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `workflowStepsSummary` | `List<WorkflowStepSummary>` | yes |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `templateId` | `string` | no |
| `adsApplicationConfigurationName` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `migrationWorkflowSummary` | `List<MigrationWorkflowSummary>` | yes |

## RetryWorkflowStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `stepGroupId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stepGroupId` | `string` | no |
| `workflowId` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |

## StartWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `lastStartTime` | `timestamp` | no |

## StopWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `lastStopTime` | `timestamp` | no |

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


## UpdateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `templateName` | `string` | no |
| `templateDescription` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateId` | `string` | no |
| `templateArn` | `string` | no |
| `tags` | `Map<string>` | no |

## UpdateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `inputParameters` | `Map<StepInput>` | no |
| `stepTargets` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `templateId` | `string` | no |
| `adsApplicationConfigurationId` | `string` | no |
| `workflowInputs` | `Map<StepInput>` | no |
| `stepTargets` | `List<string>` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## UpdateWorkflowStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `stepGroupId` | `string` | yes |
| `workflowId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `stepActionType` | `string` | no |
| `workflowStepAutomationConfiguration` | `WorkflowStepAutomationConfiguration` | no |
| `stepTarget` | `List<string>` | no |
| `outputs` | `List<WorkflowStepOutput>` | no |
| `previous` | `List<string>` | no |
| `next` | `List<string>` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `stepGroupId` | `string` | no |
| `workflowId` | `string` | no |
| `name` | `string` | no |

## UpdateWorkflowStepGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `next` | `List<string>` | no |
| `previous` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowId` | `string` | no |
| `name` | `string` | no |
| `id` | `string` | no |
| `description` | `string` | no |
| `tools` | `List<Tool>` | no |
| `next` | `List<string>` | no |
| `previous` | `List<string>` | no |
| `lastModifiedTime` | `timestamp` | no |

