# Amazon Simple Workflow Service

API version: 2012-01-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/swf/2012-01-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CountClosedWorkflowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `startTimeFilter` | `ExecutionTimeFilter` | no |
| `closeTimeFilter` | `ExecutionTimeFilter` | no |
| `executionFilter` | `WorkflowExecutionFilter` | no |
| `typeFilter` | `WorkflowTypeFilter` | no |
| `tagFilter` | `TagFilter` | no |
| `closeStatusFilter` | `CloseStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `count` | `integer` | yes |
| `truncated` | `boolean` | no |

## CountOpenWorkflowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `startTimeFilter` | `ExecutionTimeFilter` | yes |
| `typeFilter` | `WorkflowTypeFilter` | no |
| `tagFilter` | `TagFilter` | no |
| `executionFilter` | `WorkflowExecutionFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `count` | `integer` | yes |
| `truncated` | `boolean` | no |

## CountPendingActivityTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `taskList` | `TaskList` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `count` | `integer` | yes |
| `truncated` | `boolean` | no |

## CountPendingDecisionTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `taskList` | `TaskList` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `count` | `integer` | yes |
| `truncated` | `boolean` | no |

## DeleteActivityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `activityType` | `ActivityType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkflowType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowType` | `WorkflowType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateActivityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `activityType` | `ActivityType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprecateWorkflowType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowType` | `WorkflowType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeActivityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `activityType` | `ActivityType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `typeInfo` | `ActivityTypeInfo` | yes |
| `configuration` | `ActivityTypeConfiguration` | yes |

## DescribeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainInfo` | `DomainInfo` | yes |
| `configuration` | `DomainConfiguration` | yes |

## DescribeWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `execution` | `WorkflowExecution` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionInfo` | `WorkflowExecutionInfo` | yes |
| `executionConfiguration` | `WorkflowExecutionConfiguration` | yes |
| `openCounts` | `WorkflowExecutionOpenCounts` | yes |
| `latestActivityTaskTimestamp` | `timestamp` | no |
| `latestExecutionContext` | `string` | no |

## DescribeWorkflowType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowType` | `WorkflowType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `typeInfo` | `WorkflowTypeInfo` | yes |
| `configuration` | `WorkflowTypeConfiguration` | yes |

## GetWorkflowExecutionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `execution` | `WorkflowExecution` | yes |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<HistoryEvent>` | yes |
| `nextPageToken` | `string` | no |

## ListActivityTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `name` | `string` | no |
| `registrationStatus` | `string` | yes |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `typeInfos` | `List<ActivityTypeInfo>` | yes |
| `nextPageToken` | `string` | no |

## ListClosedWorkflowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `startTimeFilter` | `ExecutionTimeFilter` | no |
| `closeTimeFilter` | `ExecutionTimeFilter` | no |
| `executionFilter` | `WorkflowExecutionFilter` | no |
| `closeStatusFilter` | `CloseStatusFilter` | no |
| `typeFilter` | `WorkflowTypeFilter` | no |
| `tagFilter` | `TagFilter` | no |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionInfos` | `List<WorkflowExecutionInfo>` | yes |
| `nextPageToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextPageToken` | `string` | no |
| `registrationStatus` | `string` | yes |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainInfos` | `List<DomainInfo>` | yes |
| `nextPageToken` | `string` | no |

## ListOpenWorkflowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `startTimeFilter` | `ExecutionTimeFilter` | yes |
| `typeFilter` | `WorkflowTypeFilter` | no |
| `tagFilter` | `TagFilter` | no |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |
| `executionFilter` | `WorkflowExecutionFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionInfos` | `List<WorkflowExecutionInfo>` | yes |
| `nextPageToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<ResourceTag>` | no |

## ListWorkflowTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `name` | `string` | no |
| `registrationStatus` | `string` | yes |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `typeInfos` | `List<WorkflowTypeInfo>` | yes |
| `nextPageToken` | `string` | no |

## PollForActivityTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `taskList` | `TaskList` | yes |
| `identity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `activityId` | `string` | yes |
| `startedEventId` | `long` | yes |
| `workflowExecution` | `WorkflowExecution` | yes |
| `activityType` | `ActivityType` | yes |
| `input` | `string` | no |

## PollForDecisionTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `taskList` | `TaskList` | yes |
| `identity` | `string` | no |
| `nextPageToken` | `string` | no |
| `maximumPageSize` | `integer` | no |
| `reverseOrder` | `boolean` | no |
| `startAtPreviousStartedEvent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `startedEventId` | `long` | yes |
| `workflowExecution` | `WorkflowExecution` | yes |
| `workflowType` | `WorkflowType` | yes |
| `events` | `List<HistoryEvent>` | yes |
| `nextPageToken` | `string` | no |
| `previousStartedEventId` | `long` | no |

## RecordActivityTaskHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `details` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cancelRequested` | `boolean` | yes |

## RegisterActivityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `name` | `string` | yes |
| `version` | `string` | yes |
| `description` | `string` | no |
| `defaultTaskStartToCloseTimeout` | `string` | no |
| `defaultTaskHeartbeatTimeout` | `string` | no |
| `defaultTaskList` | `TaskList` | no |
| `defaultTaskPriority` | `string` | no |
| `defaultTaskScheduleToStartTimeout` | `string` | no |
| `defaultTaskScheduleToCloseTimeout` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `workflowExecutionRetentionPeriodInDays` | `string` | yes |
| `tags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterWorkflowType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `name` | `string` | yes |
| `version` | `string` | yes |
| `description` | `string` | no |
| `defaultTaskStartToCloseTimeout` | `string` | no |
| `defaultExecutionStartToCloseTimeout` | `string` | no |
| `defaultTaskList` | `TaskList` | no |
| `defaultTaskPriority` | `string` | no |
| `defaultChildPolicy` | `string` | no |
| `defaultLambdaRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RequestCancelWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowId` | `string` | yes |
| `runId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RespondActivityTaskCanceled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `details` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RespondActivityTaskCompleted

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `result` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RespondActivityTaskFailed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `reason` | `string` | no |
| `details` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RespondDecisionTaskCompleted

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskToken` | `string` | yes |
| `decisions` | `List<Decision>` | no |
| `executionContext` | `string` | no |
| `taskList` | `TaskList` | no |
| `taskListScheduleToStartTimeout` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SignalWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowId` | `string` | yes |
| `runId` | `string` | no |
| `signalName` | `string` | yes |
| `input` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowId` | `string` | yes |
| `workflowType` | `WorkflowType` | yes |
| `taskList` | `TaskList` | no |
| `taskPriority` | `string` | no |
| `input` | `string` | no |
| `executionStartToCloseTimeout` | `string` | no |
| `tagList` | `List<string>` | no |
| `taskStartToCloseTimeout` | `string` | no |
| `childPolicy` | `string` | no |
| `lambdaRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<ResourceTag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowId` | `string` | yes |
| `runId` | `string` | no |
| `reason` | `string` | no |
| `details` | `string` | no |
| `childPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UndeprecateActivityType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `activityType` | `ActivityType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UndeprecateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UndeprecateWorkflowType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `workflowType` | `WorkflowType` | yes |

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


