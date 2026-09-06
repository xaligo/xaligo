# AWS Step Functions

API version: 2016-11-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/stepfunctions/2016-11-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateActivity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityArn` | `string` | yes |
| `creationDate` | `timestamp` | yes |

## CreateStateMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `definition` | `string` | yes |
| `roleArn` | `string` | yes |
| `type` | `string` | no |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `tags` | `List<Tag>` | no |
| `tracingConfiguration` | `TracingConfiguration` | no |
| `publish` | `boolean` | no |
| `versionDescription` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `creationDate` | `timestamp` | yes |
| `stateMachineVersionArn` | `string` | no |

## CreateStateMachineAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `name` | `string` | yes |
| `routingConfiguration` | `List<RoutingConfigurationListItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliasArn` | `string` | yes |
| `creationDate` | `timestamp` | yes |

## DeleteActivity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStateMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStateMachineAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliasArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStateMachineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeActivity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityArn` | `string` | yes |
| `name` | `string` | yes |
| `creationDate` | `timestamp` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

## DescribeExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `stateMachineArn` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `startDate` | `timestamp` | yes |
| `stopDate` | `timestamp` | no |
| `input` | `string` | no |
| `inputDetails` | `CloudWatchEventsExecutionDataDetails` | no |
| `output` | `string` | no |
| `outputDetails` | `CloudWatchEventsExecutionDataDetails` | no |
| `traceHeader` | `string` | no |
| `mapRunArn` | `string` | no |
| `error` | `string` | no |
| `cause` | `string` | no |
| `stateMachineVersionArn` | `string` | no |
| `stateMachineAliasArn` | `string` | no |
| `redriveCount` | `integer` | no |
| `redriveDate` | `timestamp` | no |
| `redriveStatus` | `string` | no |
| `redriveStatusReason` | `string` | no |

## DescribeMapRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mapRunArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mapRunArn` | `string` | yes |
| `executionArn` | `string` | yes |
| `status` | `string` | yes |
| `startDate` | `timestamp` | yes |
| `stopDate` | `timestamp` | no |
| `maxConcurrency` | `integer` | yes |
| `toleratedFailurePercentage` | `float` | yes |
| `toleratedFailureCount` | `long` | yes |
| `itemCounts` | `MapRunItemCounts` | yes |
| `executionCounts` | `MapRunExecutionCounts` | yes |
| `redriveCount` | `integer` | no |
| `redriveDate` | `timestamp` | no |

## DescribeStateMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | no |
| `definition` | `string` | yes |
| `roleArn` | `string` | yes |
| `type` | `string` | yes |
| `creationDate` | `timestamp` | yes |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `tracingConfiguration` | `TracingConfiguration` | no |
| `label` | `string` | no |
| `revisionId` | `string` | no |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `variableReferences` | `Map<List<string>>` | no |

## DescribeStateMachineAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliasArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliasArn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `routingConfiguration` | `List<RoutingConfigurationListItem>` | no |
| `creationDate` | `timestamp` | no |
| `updateDate` | `timestamp` | no |

## DescribeStateMachineForExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `name` | `string` | yes |
| `definition` | `string` | yes |
| `roleArn` | `string` | yes |
| `updateDate` | `timestamp` | yes |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `tracingConfiguration` | `TracingConfiguration` | no |
| `mapRunArn` | `string` | no |
| `label` | `string` | no |
| `revisionId` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `variableReferences` | `Map<List<string>>` | no |

## GetActivityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activityArn` | `string` | yes |
| `workerName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | no |
| `input` | `string` | no |

## GetExecutionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `reverseOrder` | `boolean` | no |
| `nextToken` | `string` | no |
| `includeExecutionData` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<HistoryEvent>` | yes |
| `nextToken` | `string` | no |

## ListActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `activities` | `List<ActivityListItem>` | yes |
| `nextToken` | `string` | no |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | no |
| `statusFilter` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `mapRunArn` | `string` | no |
| `redriveFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executions` | `List<ExecutionListItem>` | yes |
| `nextToken` | `string` | no |

## ListMapRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mapRuns` | `List<MapRunListItem>` | yes |
| `nextToken` | `string` | no |

## ListStateMachineAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliases` | `List<StateMachineAliasListItem>` | yes |
| `nextToken` | `string` | no |

## ListStateMachineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineVersions` | `List<StateMachineVersionListItem>` | yes |
| `nextToken` | `string` | no |

## ListStateMachines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachines` | `List<StateMachineListItem>` | yes |
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

## PublishStateMachineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `revisionId` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationDate` | `timestamp` | yes |
| `stateMachineVersionArn` | `string` | yes |

## RedriveExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `redriveDate` | `timestamp` | yes |

## SendTaskFailure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `error` | `string` | no |
| `cause` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendTaskHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendTaskSuccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `output` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `name` | `string` | no |
| `input` | `string` | no |
| `traceHeader` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `startDate` | `timestamp` | yes |

## StartSyncExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `name` | `string` | no |
| `input` | `string` | no |
| `traceHeader` | `string` | no |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `stateMachineArn` | `string` | no |
| `name` | `string` | no |
| `startDate` | `timestamp` | yes |
| `stopDate` | `timestamp` | yes |
| `status` | `string` | yes |
| `error` | `string` | no |
| `cause` | `string` | no |
| `input` | `string` | no |
| `inputDetails` | `CloudWatchEventsExecutionDataDetails` | no |
| `output` | `string` | no |
| `outputDetails` | `CloudWatchEventsExecutionDataDetails` | no |
| `traceHeader` | `string` | no |
| `billingDetails` | `BillingDetails` | no |

## StopExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionArn` | `string` | yes |
| `error` | `string` | no |
| `cause` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stopDate` | `timestamp` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `definition` | `string` | yes |
| `roleArn` | `string` | no |
| `input` | `string` | no |
| `inspectionLevel` | `string` | no |
| `revealSecrets` | `boolean` | no |
| `variables` | `string` | no |
| `stateName` | `string` | no |
| `mock` | `MockInput` | no |
| `context` | `string` | no |
| `stateConfiguration` | `TestStateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `output` | `string` | no |
| `error` | `string` | no |
| `cause` | `string` | no |
| `inspectionData` | `InspectionData` | no |
| `nextState` | `string` | no |
| `status` | `string` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMapRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mapRunArn` | `string` | yes |
| `maxConcurrency` | `integer` | no |
| `toleratedFailurePercentage` | `float` | no |
| `toleratedFailureCount` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStateMachine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineArn` | `string` | yes |
| `definition` | `string` | no |
| `roleArn` | `string` | no |
| `loggingConfiguration` | `LoggingConfiguration` | no |
| `tracingConfiguration` | `TracingConfiguration` | no |
| `publish` | `boolean` | no |
| `versionDescription` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateDate` | `timestamp` | yes |
| `revisionId` | `string` | no |
| `stateMachineVersionArn` | `string` | no |

## UpdateStateMachineAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stateMachineAliasArn` | `string` | yes |
| `description` | `string` | no |
| `routingConfiguration` | `List<RoutingConfigurationListItem>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateDate` | `timestamp` | yes |

## ValidateStateMachineDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `definition` | `string` | yes |
| `type` | `string` | no |
| `severity` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `result` | `string` | yes |
| `diagnostics` | `List<ValidateStateMachineDefinitionDiagnostic>` | yes |
| `truncated` | `boolean` | no |

