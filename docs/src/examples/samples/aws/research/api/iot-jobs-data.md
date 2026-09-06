# AWS IoT Jobs Data Plane

API version: 2017-09-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iot-jobs-data/2017-09-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `thingName` | `string` | yes |
| `includeJobDocument` | `boolean` | no |
| `executionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `execution` | `JobExecution` | no |

## GetPendingJobExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inProgressJobs` | `List<JobExecutionSummary>` | no |
| `queuedJobs` | `List<JobExecutionSummary>` | no |

## StartCommandExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetArn` | `string` | yes |
| `commandArn` | `string` | yes |
| `parameters` | `Map<CommandParameterValue>` | no |
| `executionTimeoutSeconds` | `long` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |

## StartNextPendingJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `thingName` | `string` | yes |
| `statusDetails` | `Map<string>` | no |
| `stepTimeoutInMinutes` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `execution` | `JobExecution` | no |

## UpdateJobExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `thingName` | `string` | yes |
| `status` | `string` | yes |
| `statusDetails` | `Map<string>` | no |
| `stepTimeoutInMinutes` | `long` | no |
| `expectedVersion` | `long` | no |
| `includeJobExecutionState` | `boolean` | no |
| `includeJobDocument` | `boolean` | no |
| `executionNumber` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionState` | `JobExecutionState` | no |
| `jobDocument` | `string` | no |

