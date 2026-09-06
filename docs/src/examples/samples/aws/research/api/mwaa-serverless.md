# AmazonMWAAServerless

API version: 2024-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mwaa-serverless/2024-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ClientToken` | `string` | no |
| `DefinitionS3Location` | `DefinitionS3Location` | yes |
| `Code` | `Code` | no |
| `RoleArn` | `string` | yes |
| `Description` | `string` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EngineVersion` | `integer` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `TriggerMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `CreatedAt` | `timestamp` | no |
| `RevisionId` | `string` | no |
| `WorkflowStatus` | `string` | no |
| `WorkflowVersion` | `string` | no |
| `IsLatestVersion` | `boolean` | no |
| `Warnings` | `List<string>` | no |

## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `WorkflowVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `WorkflowVersion` | `string` | no |

## GetTaskInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `TaskInstanceId` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `RunId` | `string` | yes |
| `TaskInstanceId` | `string` | yes |
| `WorkflowVersion` | `string` | no |
| `Status` | `string` | no |
| `DurationInSeconds` | `integer` | no |
| `OperatorName` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `StartedAt` | `timestamp` | no |
| `AttemptNumber` | `integer` | no |
| `ErrorMessage` | `string` | no |
| `TaskId` | `string` | no |
| `LogStream` | `string` | no |
| `Xcom` | `Map<string>` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `WorkflowVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `WorkflowVersion` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `ModifiedAt` | `timestamp` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EngineVersion` | `integer` | no |
| `WorkflowStatus` | `string` | no |
| `DefinitionS3Location` | `DefinitionS3Location` | no |
| `Code` | `Code` | no |
| `CodeSnapshottedAt` | `timestamp` | no |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |
| `RoleArn` | `string` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `TriggerMode` | `string` | no |
| `WorkflowDefinition` | `string` | no |

## GetWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | no |
| `WorkflowVersion` | `string` | no |
| `RunId` | `string` | no |
| `RunType` | `string` | no |
| `OverrideParameters` | `Map<Document>` | no |
| `RunDetail` | `WorkflowRunDetail` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTaskInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `RunId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskInstances` | `List<TaskInstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkflowRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `WorkflowArn` | `string` | yes |
| `WorkflowVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowRuns` | `List<WorkflowRunSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkflowVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `WorkflowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowVersions` | `List<WorkflowVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workflows` | `List<WorkflowSummary>` | yes |
| `NextToken` | `string` | no |

## StartWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `OverrideParameters` | `Map<Document>` | no |
| `WorkflowVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RunId` | `string` | no |
| `Status` | `string` | no |
| `StartedAt` | `timestamp` | no |

## StopWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `RunId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | no |
| `WorkflowVersion` | `string` | no |
| `RunId` | `string` | no |
| `Status` | `string` | no |

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


## UpdateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `DefinitionS3Location` | `DefinitionS3Location` | yes |
| `Code` | `Code` | no |
| `RoleArn` | `string` | yes |
| `Description` | `string` | no |
| `LoggingConfiguration` | `LoggingConfiguration` | no |
| `EngineVersion` | `integer` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `TriggerMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowArn` | `string` | yes |
| `ModifiedAt` | `timestamp` | no |
| `WorkflowVersion` | `string` | no |
| `Warnings` | `List<string>` | no |

