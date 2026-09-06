# AWS CodePipeline

API version: 2015-07-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codepipeline/2015-07-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcknowledgeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `nonce` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## AcknowledgeThirdPartyJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `nonce` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CreateCustomActionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `category` | `string` | yes |
| `provider` | `string` | yes |
| `version` | `string` | yes |
| `settings` | `ActionTypeSettings` | no |
| `configurationProperties` | `List<ActionConfigurationProperty>` | no |
| `inputArtifactDetails` | `ArtifactDetails` | yes |
| `outputArtifactDetails` | `ArtifactDetails` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionType` | `ActionType` | yes |
| `tags` | `List<Tag>` | no |

## CreatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `PipelineDeclaration` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `PipelineDeclaration` | no |
| `tags` | `List<Tag>` | no |

## DeleteCustomActionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `category` | `string` | yes |
| `provider` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterWebhookWithThirdParty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhookName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableStageTransition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `transitionType` | `string` | yes |
| `reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableStageTransition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `transitionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetActionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `category` | `string` | yes |
| `owner` | `string` | yes |
| `provider` | `string` | yes |
| `version` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionType` | `ActionTypeDeclaration` | no |

## GetJobDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDetails` | `JobDetails` | no |

## GetPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `PipelineDeclaration` | no |
| `metadata` | `PipelineMetadata` | no |

## GetPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecution` | `PipelineExecution` | no |

## GetPipelineState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | no |
| `pipelineVersion` | `integer` | no |
| `stageStates` | `List<StageState>` | no |
| `created` | `timestamp` | no |
| `updated` | `timestamp` | no |

## GetThirdPartyJobDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDetails` | `ThirdPartyJobDetails` | no |

## ListActionExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `filter` | `ActionExecutionFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionExecutionDetails` | `List<ActionExecutionDetail>` | no |
| `nextToken` | `string` | no |

## ListActionTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionOwnerFilter` | `string` | no |
| `nextToken` | `string` | no |
| `regionFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionTypes` | `List<ActionType>` | yes |
| `nextToken` | `string` | no |

## ListDeployActionExecutionTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | no |
| `actionExecutionId` | `string` | yes |
| `filters` | `List<TargetFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targets` | `List<DeployActionExecutionTarget>` | no |
| `nextToken` | `string` | no |

## ListPipelineExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `maxResults` | `integer` | no |
| `filter` | `PipelineExecutionFilter` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionSummaries` | `List<PipelineExecutionSummary>` | no |
| `nextToken` | `string` | no |

## ListPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelines` | `List<PipelineSummary>` | no |
| `nextToken` | `string` | no |

## ListRuleExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `filter` | `RuleExecutionFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleExecutionDetails` | `List<RuleExecutionDetail>` | no |
| `nextToken` | `string` | no |

## ListRuleTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleOwnerFilter` | `string` | no |
| `regionFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleTypes` | `List<RuleType>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `nextToken` | `string` | no |

## ListWebhooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhooks` | `List<ListWebhookItem>` | no |
| `NextToken` | `string` | no |

## OverrideStageCondition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |
| `conditionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PollForJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionTypeId` | `ActionTypeId` | yes |
| `maxBatchSize` | `integer` | no |
| `queryParam` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<Job>` | no |

## PollForThirdPartyJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionTypeId` | `ActionTypeId` | yes |
| `maxBatchSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<ThirdPartyJob>` | no |

## PutActionRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `actionName` | `string` | yes |
| `actionRevision` | `ActionRevision` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `newRevision` | `boolean` | no |
| `pipelineExecutionId` | `string` | no |

## PutApprovalResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `actionName` | `string` | yes |
| `result` | `ApprovalResult` | yes |
| `token` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvedAt` | `timestamp` | no |

## PutJobFailureResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `failureDetails` | `FailureDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutJobSuccessResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `currentRevision` | `CurrentRevision` | no |
| `continuationToken` | `string` | no |
| `executionDetails` | `ExecutionDetails` | no |
| `outputVariables` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutThirdPartyJobFailureResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `clientToken` | `string` | yes |
| `failureDetails` | `FailureDetails` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutThirdPartyJobSuccessResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `clientToken` | `string` | yes |
| `currentRevision` | `CurrentRevision` | no |
| `continuationToken` | `string` | no |
| `executionDetails` | `ExecutionDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `WebhookDefinition` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `ListWebhookItem` | no |

## RegisterWebhookWithThirdParty

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhookName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RetryStageExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |
| `retryMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | no |

## RollbackStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `stageName` | `string` | yes |
| `targetPipelineExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | yes |

## StartPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `variables` | `List<PipelineVariable>` | no |
| `clientRequestToken` | `string` | no |
| `sourceRevisions` | `List<SourceRevisionOverride>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | no |

## StopPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineName` | `string` | yes |
| `pipelineExecutionId` | `string` | yes |
| `abandon` | `boolean` | no |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineExecutionId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateActionType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionType` | `ActionTypeDeclaration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `PipelineDeclaration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `PipelineDeclaration` | no |

