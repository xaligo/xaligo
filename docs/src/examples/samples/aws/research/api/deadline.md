# AWSDeadlineCloud

API version: 2023-10-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/deadline/2023-10-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateMemberToFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `principalType` | `string` | yes |
| `identityStoreId` | `string` | yes |
| `membershipLevel` | `string` | yes |
| `principalId` | `string` | yes |
| `identityCenterRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateMemberToFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `principalType` | `string` | yes |
| `identityStoreId` | `string` | yes |
| `membershipLevel` | `string` | yes |
| `principalId` | `string` | yes |
| `identityCenterRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateMemberToJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `principalType` | `string` | yes |
| `identityStoreId` | `string` | yes |
| `membershipLevel` | `string` | yes |
| `principalId` | `string` | yes |
| `identityCenterRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateMemberToQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `principalType` | `string` | yes |
| `identityStoreId` | `string` | yes |
| `membershipLevel` | `string` | yes |
| `principalId` | `string` | yes |
| `identityCenterRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssumeFleetRoleForRead

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | yes |

## AssumeFleetRoleForWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | yes |

## AssumeQueueRoleForRead

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | yes |

## AssumeQueueRoleForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | yes |

## AssumeQueueRoleForWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `queueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentials` | `AwsCredentials` | no |

## BatchGetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetJobIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<BatchGetJobItem>` | yes |
| `errors` | `List<BatchGetJobError>` | yes |

## BatchGetJobEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `identifiers` | `List<JobEntityIdentifiersUnion>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entities` | `List<JobEntity>` | yes |
| `errors` | `List<GetJobEntityError>` | yes |

## BatchGetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetSessionIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessions` | `List<BatchGetSessionItem>` | yes |
| `errors` | `List<BatchGetSessionError>` | yes |

## BatchGetSessionAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetSessionActionIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionActions` | `List<BatchGetSessionActionItem>` | yes |
| `errors` | `List<BatchGetSessionActionError>` | yes |

## BatchGetStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetStepIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `steps` | `List<BatchGetStepItem>` | yes |
| `errors` | `List<BatchGetStepError>` | yes |

## BatchGetTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetTaskIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<BatchGetTaskItem>` | yes |
| `errors` | `List<BatchGetTaskError>` | yes |

## BatchGetWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<BatchGetWorkerIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workers` | `List<BatchGetWorkerItem>` | yes |
| `errors` | `List<BatchGetWorkerError>` | yes |

## BatchUpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `jobs` | `List<BatchUpdateJobItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchUpdateJobError>` | yes |

## BatchUpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `tasks` | `List<BatchUpdateTaskItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchUpdateTaskError>` | yes |

## CopyJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `targetS3Location` | `S3Location` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateType` | `string` | yes |

## CreateBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `displayName` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `usageTrackingResource` | `UsageTrackingResource` | yes |
| `approximateDollarLimit` | `float` | yes |
| `actions` | `List<BudgetActionToAdd>` | yes |
| `schedule` | `BudgetSchedule` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `budgetId` | `string` | yes |

## CreateFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `description` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `costScaleFactor` | `float` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `minWorkerCount` | `integer` | no |
| `maxWorkerCount` | `integer` | yes |
| `configuration` | `FleetConfiguration` | yes |
| `tags` | `Map<string>` | no |
| `hostConfiguration` | `HostConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `clientToken` | `string` | no |
| `template` | `string` | no |
| `templateType` | `string` | no |
| `priority` | `integer` | yes |
| `parameters` | `Map<JobParameter>` | no |
| `attachments` | `Attachments` | no |
| `storageProfileId` | `string` | no |
| `targetTaskRunStatus` | `string` | no |
| `maxFailedTasksCount` | `integer` | no |
| `maxRetriesPerTask` | `integer` | no |
| `maxWorkerCount` | `integer` | no |
| `sourceJobId` | `string` | no |
| `nameOverride` | `string` | no |
| `descriptionOverride` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

## CreateLicenseEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |

## CreateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `amountRequirementName` | `string` | yes |
| `maxCount` | `integer` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limitId` | `string` | yes |

## CreateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `identityCenterInstanceArn` | `string` | yes |
| `identityCenterRegion` | `string` | no |
| `subdomain` | `string` | yes |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |
| `identityCenterApplicationArn` | `string` | yes |

## CreateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `description` | `string` | no |
| `defaultBudgetAction` | `string` | no |
| `jobAttachmentSettings` | `JobAttachmentSettings` | no |
| `roleArn` | `string` | no |
| `jobRunAsUser` | `JobRunAsUser` | no |
| `requiredFileSystemLocationNames` | `List<string>` | no |
| `allowedStorageProfileIds` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `schedulingConfiguration` | `SchedulingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueId` | `string` | yes |

## CreateQueueEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `clientToken` | `string` | no |
| `priority` | `integer` | yes |
| `templateType` | `string` | yes |
| `template` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueEnvironmentId` | `string` | yes |

## CreateQueueFleetAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateQueueLimitAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `limitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateStorageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | yes |
| `osFamily` | `string` | yes |
| `fileSystemLocations` | `List<FileSystemLocation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageProfileId` | `string` | yes |

## CreateWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `hostProperties` | `HostPropertiesRequest` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workerId` | `string` | yes |

## DeleteBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `budgetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLicenseEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `limitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMeteredProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |
| `productId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueueEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `queueEnvironmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueueFleetAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueueLimitAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `limitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStorageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `storageProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `volumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMemberFromFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `principalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMemberFromFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `principalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMemberFromJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `principalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMemberFromQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `principalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `budgetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `budgetId` | `string` | yes |
| `usageTrackingResource` | `UsageTrackingResource` | yes |
| `status` | `string` | yes |
| `displayName` | `string` | yes |
| `approximateDollarLimit` | `float` | yes |
| `usages` | `ConsumedUsages` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `description` | `string` | no |
| `actions` | `List<ResponseBudgetAction>` | yes |
| `schedule` | `BudgetSchedule` | yes |
| `queueStoppedAt` | `timestamp` | no |

## GetFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `displayName` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `description` | `string` | no |
| `costScaleFactor` | `float` | yes |

## GetFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |
| `farmId` | `string` | yes |
| `displayName` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `autoScalingStatus` | `string` | no |
| `targetWorkerCount` | `integer` | no |
| `workerCount` | `integer` | yes |
| `minWorkerCount` | `integer` | yes |
| `maxWorkerCount` | `integer` | yes |
| `configuration` | `FleetConfiguration` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `description` | `string` | no |
| `hostConfiguration` | `HostConfiguration` | no |
| `capabilities` | `FleetCapabilities` | no |
| `roleArn` | `string` | yes |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `name` | `string` | yes |
| `lifecycleStatus` | `string` | yes |
| `lifecycleStatusMessage` | `string` | yes |
| `priority` | `integer` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `taskRunStatus` | `string` | no |
| `targetTaskRunStatus` | `string` | no |
| `taskRunStatusCounts` | `Map<integer>` | no |
| `taskFailureRetryCount` | `integer` | no |
| `storageProfileId` | `string` | no |
| `maxFailedTasksCount` | `integer` | no |
| `maxRetriesPerTask` | `integer` | no |
| `parameters` | `Map<JobParameter>` | no |
| `attachments` | `Attachments` | no |
| `description` | `string` | no |
| `maxWorkerCount` | `integer` | no |
| `sourceJobId` | `string` | no |

## GetLicenseEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | yes |
| `vpcId` | `string` | no |
| `dnsName` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |

## GetLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `limitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `limitId` | `string` | yes |
| `currentCount` | `integer` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `displayName` | `string` | yes |
| `amountRequirementName` | `string` | yes |
| `maxCount` | `integer` | yes |
| `description` | `string` | no |

## GetMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |
| `displayName` | `string` | yes |
| `subdomain` | `string` | yes |
| `url` | `string` | yes |
| `roleArn` | `string` | yes |
| `identityCenterInstanceArn` | `string` | yes |
| `identityCenterRegion` | `string` | no |
| `identityCenterApplicationArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## GetMonitorSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `Map<string>` | yes |

## GetQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `displayName` | `string` | yes |
| `status` | `string` | yes |
| `defaultBudgetAction` | `string` | yes |
| `blockedReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `description` | `string` | no |
| `jobAttachmentSettings` | `JobAttachmentSettings` | no |
| `roleArn` | `string` | no |
| `requiredFileSystemLocationNames` | `List<string>` | no |
| `allowedStorageProfileIds` | `List<string>` | no |
| `jobRunAsUser` | `JobRunAsUser` | no |
| `schedulingConfiguration` | `SchedulingConfiguration` | no |

## GetQueueEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `queueEnvironmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueEnvironmentId` | `string` | yes |
| `name` | `string` | yes |
| `priority` | `integer` | yes |
| `templateType` | `string` | yes |
| `template` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## GetQueueFleetAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueId` | `string` | yes |
| `fleetId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## GetQueueLimitAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `limitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueId` | `string` | yes |
| `limitId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `startedAt` | `timestamp` | yes |
| `lifecycleStatus` | `string` | yes |
| `endedAt` | `timestamp` | no |
| `targetLifecycleStatus` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `log` | `LogConfiguration` | yes |
| `hostProperties` | `HostPropertiesResponse` | no |
| `workerLog` | `LogConfiguration` | no |

## GetSessionAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `sessionActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionActionId` | `string` | yes |
| `status` | `string` | yes |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `workerUpdatedAt` | `timestamp` | no |
| `progressPercent` | `float` | no |
| `manifests` | `List<TaskRunManifestPropertiesResponse>` | no |
| `sessionId` | `string` | yes |
| `processExitCode` | `integer` | no |
| `progressMessage` | `string` | no |
| `acquiredLimits` | `List<AcquiredLimit>` | no |
| `definition` | `SessionActionDefinition` | yes |

## GetSessionsStatisticsAggregation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `aggregationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statistics` | `List<Statistics>` | no |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `nextToken` | `string` | no |

## GetStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stepId` | `string` | yes |
| `name` | `string` | yes |
| `lifecycleStatus` | `string` | yes |
| `lifecycleStatusMessage` | `string` | no |
| `taskRunStatus` | `string` | yes |
| `taskRunStatusCounts` | `Map<integer>` | yes |
| `taskFailureRetryCount` | `integer` | no |
| `targetTaskRunStatus` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `dependencyCounts` | `DependencyCounts` | no |
| `requiredCapabilities` | `StepRequiredCapabilities` | no |
| `parameterSpace` | `ParameterSpace` | no |
| `description` | `string` | no |

## GetStorageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `storageProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageProfileId` | `string` | yes |
| `displayName` | `string` | yes |
| `osFamily` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `fileSystemLocations` | `List<FileSystemLocation>` | no |

## GetStorageProfileForQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `storageProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageProfileId` | `string` | yes |
| `displayName` | `string` | yes |
| `osFamily` | `string` | yes |
| `fileSystemLocations` | `List<FileSystemLocation>` | no |

## GetTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `runStatus` | `string` | yes |
| `targetRunStatus` | `string` | no |
| `failureRetryCount` | `integer` | no |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `latestSessionActionId` | `string` | no |
| `parameters` | `Map<TaskParameterValue>` | no |

## GetVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `volumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `volumeId` | `string` | yes |
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `state` | `string` | yes |
| `sizeGiB` | `integer` | yes |
| `availabilityZoneId` | `string` | yes |
| `attachedWorkerId` | `string` | no |
| `volumeType` | `string` | yes |
| `iops` | `integer` | no |
| `throughputMiB` | `integer` | no |
| `createdAt` | `timestamp` | yes |
| `lastAssignedAt` | `timestamp` | no |
| `lastReleasedAt` | `timestamp` | no |
| `expiresAt` | `timestamp` | no |

## GetWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `hostProperties` | `HostPropertiesResponse` | no |
| `status` | `string` | yes |
| `log` | `LogConfiguration` | no |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## ListAvailableMeteredProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meteredProducts` | `List<MeteredProductSummary>` | yes |
| `nextToken` | `string` | no |

## ListBudgets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `budgets` | `List<BudgetSummary>` | yes |
| `nextToken` | `string` | no |

## ListFarmMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<FarmMember>` | yes |
| `nextToken` | `string` | no |

## ListFarms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principalId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farms` | `List<FarmSummary>` | yes |
| `nextToken` | `string` | no |

## ListFleetMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<FleetMember>` | yes |
| `nextToken` | `string` | no |

## ListFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principalId` | `string` | no |
| `displayName` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleets` | `List<FleetSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<JobMember>` | yes |
| `nextToken` | `string` | no |

## ListJobParameterDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobParameterDefinitions` | `List<JobParameterDefinition>` | yes |
| `nextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principalId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSummary>` | yes |
| `nextToken` | `string` | no |

## ListLicenseEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpoints` | `List<LicenseEndpointSummary>` | yes |
| `nextToken` | `string` | no |

## ListLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limits` | `List<LimitSummary>` | yes |
| `nextToken` | `string` | no |

## ListMeteredProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `meteredProducts` | `List<MeteredProductSummary>` | yes |
| `nextToken` | `string` | no |

## ListMonitors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitors` | `List<MonitorSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueueEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<QueueEnvironmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueueFleetAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `queueId` | `string` | no |
| `fleetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueFleetAssociations` | `List<QueueFleetAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueueLimitAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `queueId` | `string` | no |
| `limitId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queueLimitAssociations` | `List<QueueLimitAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueueMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<QueueMember>` | yes |
| `nextToken` | `string` | no |

## ListQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `principalId` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queues` | `List<QueueSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessionActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sessionId` | `string` | no |
| `taskId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionActions` | `List<SessionActionSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessions` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListSessionsForWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessions` | `List<WorkerSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListStepConsumers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumers` | `List<StepConsumer>` | yes |
| `nextToken` | `string` | no |

## ListStepDependencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dependencies` | `List<StepDependency>` | yes |
| `nextToken` | `string` | no |

## ListSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `steps` | `List<StepSummary>` | yes |
| `nextToken` | `string` | no |

## ListStorageProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageProfiles` | `List<StorageProfileSummary>` | yes |
| `nextToken` | `string` | no |

## ListStorageProfilesForQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageProfiles` | `List<StorageProfileSummary>` | yes |
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

## ListTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<TaskSummary>` | yes |
| `nextToken` | `string` | no |

## ListVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `volumes` | `List<VolumeSummary>` | yes |
| `nextToken` | `string` | no |

## ListWorkers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workers` | `List<WorkerSummary>` | yes |
| `nextToken` | `string` | no |

## PutMeteredProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `licenseEndpointId` | `string` | yes |
| `productId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `filterExpressions` | `SearchGroupedFilterExpressions` | no |
| `sortExpressions` | `List<SearchSortExpression>` | no |
| `itemOffset` | `integer` | yes |
| `pageSize` | `integer` | no |
| `queueIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSearchSummary>` | yes |
| `nextItemOffset` | `integer` | no |
| `totalResults` | `integer` | yes |

## SearchSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `filterExpressions` | `SearchGroupedFilterExpressions` | no |
| `sortExpressions` | `List<SearchSortExpression>` | no |
| `itemOffset` | `integer` | yes |
| `pageSize` | `integer` | no |
| `queueIds` | `List<string>` | yes |
| `jobId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `steps` | `List<StepSearchSummary>` | yes |
| `nextItemOffset` | `integer` | no |
| `totalResults` | `integer` | yes |

## SearchTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `filterExpressions` | `SearchGroupedFilterExpressions` | no |
| `sortExpressions` | `List<SearchSortExpression>` | no |
| `itemOffset` | `integer` | yes |
| `pageSize` | `integer` | no |
| `queueIds` | `List<string>` | yes |
| `jobId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<TaskSearchSummary>` | yes |
| `nextItemOffset` | `integer` | no |
| `totalResults` | `integer` | yes |

## SearchWorkers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `filterExpressions` | `SearchGroupedFilterExpressions` | no |
| `sortExpressions` | `List<SearchSortExpression>` | no |
| `itemOffset` | `integer` | yes |
| `pageSize` | `integer` | no |
| `fleetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workers` | `List<WorkerSearchSummary>` | yes |
| `nextItemOffset` | `integer` | no |
| `totalResults` | `integer` | yes |

## StartSessionsStatisticsAggregation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `resourceIds` | `SessionsStatisticsResources` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `timezone` | `string` | no |
| `period` | `string` | no |
| `groupBy` | `List<string>` | yes |
| `statistics` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregationId` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | no |

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


## UpdateBudget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `budgetId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `approximateDollarLimit` | `float` | no |
| `actionsToAdd` | `List<BudgetActionToAdd>` | no |
| `actionsToRemove` | `List<BudgetActionToRemove>` | no |
| `schedule` | `BudgetSchedule` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `costScaleFactor` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `minWorkerCount` | `integer` | no |
| `maxWorkerCount` | `integer` | no |
| `configuration` | `FleetConfiguration` | no |
| `hostConfiguration` | `HostConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `clientToken` | `string` | no |
| `targetTaskRunStatus` | `string` | no |
| `priority` | `integer` | no |
| `maxFailedTasksCount` | `integer` | no |
| `maxRetriesPerTask` | `integer` | no |
| `lifecycleStatus` | `string` | no |
| `maxWorkerCount` | `integer` | no |
| `name` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `limitId` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `maxCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMonitor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |
| `subdomain` | `string` | no |
| `displayName` | `string` | no |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMonitorSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monitorId` | `string` | yes |
| `settings` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `defaultBudgetAction` | `string` | no |
| `jobAttachmentSettings` | `JobAttachmentSettings` | no |
| `roleArn` | `string` | no |
| `jobRunAsUser` | `JobRunAsUser` | no |
| `requiredFileSystemLocationNamesToAdd` | `List<string>` | no |
| `requiredFileSystemLocationNamesToRemove` | `List<string>` | no |
| `allowedStorageProfileIdsToAdd` | `List<string>` | no |
| `allowedStorageProfileIdsToRemove` | `List<string>` | no |
| `schedulingConfiguration` | `SchedulingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `queueEnvironmentId` | `string` | yes |
| `clientToken` | `string` | no |
| `priority` | `integer` | no |
| `templateType` | `string` | no |
| `template` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueFleetAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `fleetId` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueLimitAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `limitId` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `sessionId` | `string` | yes |
| `clientToken` | `string` | no |
| `targetLifecycleStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `clientToken` | `string` | no |
| `targetTaskRunStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStorageProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `storageProfileId` | `string` | yes |
| `clientToken` | `string` | no |
| `displayName` | `string` | no |
| `osFamily` | `string` | no |
| `fileSystemLocationsToAdd` | `List<FileSystemLocation>` | no |
| `fileSystemLocationsToRemove` | `List<FileSystemLocation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `queueId` | `string` | yes |
| `jobId` | `string` | yes |
| `stepId` | `string` | yes |
| `taskId` | `string` | yes |
| `clientToken` | `string` | no |
| `targetRunStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `status` | `string` | no |
| `capabilities` | `WorkerCapabilities` | no |
| `hostProperties` | `HostPropertiesRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `log` | `LogConfiguration` | no |
| `hostConfiguration` | `HostConfiguration` | no |

## UpdateWorkerSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `farmId` | `string` | yes |
| `fleetId` | `string` | yes |
| `workerId` | `string` | yes |
| `updatedSessionActions` | `Map<UpdatedSessionActionInfo>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assignedSessions` | `Map<AssignedSession>` | yes |
| `cancelSessionActions` | `Map<List<string>>` | yes |
| `desiredWorkerStatus` | `string` | no |
| `updateIntervalSeconds` | `integer` | yes |

