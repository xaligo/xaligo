# AWS Data Pipeline

API version: 2012-10-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/datapipeline/2012-10-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `parameterValues` | `List<ParameterValue>` | no |
| `startTimestamp` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `uniqueId` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |

## DeactivatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `cancelActive` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeObjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `objectIds` | `List<string>` | yes |
| `evaluateExpressions` | `boolean` | no |
| `marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineObjects` | `List<PipelineObject>` | yes |
| `marker` | `string` | no |
| `hasMoreResults` | `boolean` | no |

## DescribePipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineDescriptionList` | `List<PipelineDescription>` | yes |

## EvaluateExpression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `objectId` | `string` | yes |
| `expression` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatedExpression` | `string` | yes |

## GetPipelineDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineObjects` | `List<PipelineObject>` | no |
| `parameterObjects` | `List<ParameterObject>` | no |
| `parameterValues` | `List<ParameterValue>` | no |

## ListPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineIdList` | `List<PipelineIdName>` | yes |
| `marker` | `string` | no |
| `hasMoreResults` | `boolean` | no |

## PollForTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workerGroup` | `string` | yes |
| `hostname` | `string` | no |
| `instanceIdentity` | `InstanceIdentity` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskObject` | `TaskObject` | no |

## PutPipelineDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `pipelineObjects` | `List<PipelineObject>` | yes |
| `parameterObjects` | `List<ParameterObject>` | no |
| `parameterValues` | `List<ParameterValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `validationErrors` | `List<ValidationError>` | no |
| `validationWarnings` | `List<ValidationWarning>` | no |
| `errored` | `boolean` | yes |

## QueryObjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `query` | `Query` | no |
| `sphere` | `string` | yes |
| `marker` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `marker` | `string` | no |
| `hasMoreResults` | `boolean` | no |

## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReportTaskProgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `fields` | `List<Field>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `canceled` | `boolean` | yes |

## ReportTaskRunnerHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskrunnerId` | `string` | yes |
| `workerGroup` | `string` | no |
| `hostname` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `terminate` | `boolean` | yes |

## SetStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `objectIds` | `List<string>` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetTaskStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `taskStatus` | `string` | yes |
| `errorId` | `string` | no |
| `errorMessage` | `string` | no |
| `errorStackTrace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ValidatePipelineDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipelineId` | `string` | yes |
| `pipelineObjects` | `List<PipelineObject>` | yes |
| `parameterObjects` | `List<ParameterObject>` | no |
| `parameterValues` | `List<ParameterValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `validationErrors` | `List<ValidationError>` | no |
| `validationWarnings` | `List<ValidationWarning>` | no |
| `errored` | `boolean` | yes |

