# EMR Serverless

API version: 2021-07-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/emr-serverless/2021-07-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |
| `shutdownGracePeriodInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `releaseLabel` | `string` | yes |
| `type` | `string` | yes |
| `clientToken` | `string` | yes |
| `initialCapacity` | `Map<InitialCapacityConfig>` | no |
| `maximumCapacity` | `MaximumAllowedResources` | no |
| `tags` | `Map<string>` | no |
| `autoStartConfiguration` | `AutoStartConfig` | no |
| `autoStopConfiguration` | `AutoStopConfig` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `architecture` | `string` | no |
| `imageConfiguration` | `ImageConfigurationInput` | no |
| `workerTypeSpecifications` | `Map<WorkerTypeSpecificationInput>` | no |
| `runtimeConfiguration` | `List<Configuration>` | no |
| `monitoringConfiguration` | `MonitoringConfiguration` | no |
| `diskEncryptionConfiguration` | `DiskEncryptionConfiguration` | no |
| `interactiveConfiguration` | `InteractiveConfiguration` | no |
| `schedulerConfiguration` | `SchedulerConfiguration` | no |
| `identityCenterConfiguration` | `IdentityCenterConfigurationInput` | no |
| `jobLevelCostAllocationConfiguration` | `JobLevelCostAllocationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `name` | `string` | no |
| `arn` | `string` | yes |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `Application` | yes |

## GetDashboardForJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |
| `attempt` | `integer` | no |
| `accessSystemProfileLogs` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `url` | `string` | no |

## GetJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |
| `attempt` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobRun` | `JobRun` | yes |

## GetResourceDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `resourceId` | `string` | yes |
| `resourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `url` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `Session` | yes |

## GetSessionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |
| `endpoint` | `string` | yes |
| `authToken` | `string` | yes |
| `authTokenExpiresAt` | `timestamp` | yes |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `states` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applications` | `List<ApplicationSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobRunAttempts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobRunAttempts` | `List<JobRunAttemptSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `createdAtAfter` | `timestamp` | no |
| `createdAtBefore` | `timestamp` | no |
| `states` | `List<string>` | no |
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobRuns` | `List<JobRunSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `states` | `List<string>` | no |
| `createdAtAfter` | `timestamp` | no |
| `createdAtBefore` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessions` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## StartApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `clientToken` | `string` | yes |
| `executionRoleArn` | `string` | yes |
| `executionIamPolicy` | `JobRunExecutionIamPolicy` | no |
| `jobDriver` | `JobDriver` | no |
| `configurationOverrides` | `ConfigurationOverrides` | no |
| `tags` | `Map<string>` | no |
| `executionTimeoutMinutes` | `long` | no |
| `name` | `string` | no |
| `mode` | `string` | no |
| `retryPolicy` | `RetryPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `jobRunId` | `string` | yes |
| `arn` | `string` | yes |

## StartSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `clientToken` | `string` | yes |
| `executionRoleArn` | `string` | yes |
| `configurationOverrides` | `SessionConfigurationOverrides` | no |
| `tags` | `Map<string>` | no |
| `idleTimeoutMinutes` | `long` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |
| `arn` | `string` | yes |

## StopApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `sessionId` | `string` | yes |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `clientToken` | `string` | yes |
| `initialCapacity` | `Map<InitialCapacityConfig>` | no |
| `maximumCapacity` | `MaximumAllowedResources` | no |
| `autoStartConfiguration` | `AutoStartConfig` | no |
| `autoStopConfiguration` | `AutoStopConfig` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `architecture` | `string` | no |
| `imageConfiguration` | `ImageConfigurationInput` | no |
| `workerTypeSpecifications` | `Map<WorkerTypeSpecificationInput>` | no |
| `interactiveConfiguration` | `InteractiveConfiguration` | no |
| `releaseLabel` | `string` | no |
| `runtimeConfiguration` | `List<Configuration>` | no |
| `monitoringConfiguration` | `MonitoringConfiguration` | no |
| `diskEncryptionConfiguration` | `DiskEncryptionConfiguration` | no |
| `schedulerConfiguration` | `SchedulerConfiguration` | no |
| `identityCenterConfiguration` | `IdentityCenterConfigurationInput` | no |
| `jobLevelCostAllocationConfiguration` | `JobLevelCostAllocationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `Application` | yes |

