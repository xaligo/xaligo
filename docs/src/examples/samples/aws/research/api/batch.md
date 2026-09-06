# AWS Batch

API version: 2016-08-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/batch/2016-08-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateComputeEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironmentName` | `string` | yes |
| `type` | `string` | yes |
| `state` | `string` | no |
| `unmanagedvCpus` | `integer` | no |
| `computeResources` | `ComputeResource` | no |
| `serviceRole` | `string` | no |
| `tags` | `Map<string>` | no |
| `eksConfiguration` | `EksConfiguration` | no |
| `context` | `string` | no |
| `ecsSettings` | `EcsSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironmentName` | `string` | no |
| `computeEnvironmentArn` | `string` | no |

## CreateConsumableResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResourceName` | `string` | yes |
| `totalQuantity` | `long` | no |
| `resourceType` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResourceName` | `string` | yes |
| `consumableResourceArn` | `string` | yes |

## CreateJobQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueueName` | `string` | yes |
| `state` | `string` | no |
| `schedulingPolicyArn` | `string` | no |
| `priority` | `integer` | yes |
| `computeEnvironmentOrder` | `List<ComputeEnvironmentOrder>` | no |
| `serviceEnvironmentOrder` | `List<ServiceEnvironmentOrder>` | no |
| `jobQueueType` | `string` | no |
| `tags` | `Map<string>` | no |
| `jobStateTimeLimitActions` | `List<JobStateTimeLimitAction>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueueName` | `string` | yes |
| `jobQueueArn` | `string` | yes |

## CreateQuotaShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareName` | `string` | yes |
| `jobQueue` | `string` | yes |
| `capacityLimits` | `List<QuotaShareCapacityLimit>` | yes |
| `resourceSharingConfiguration` | `QuotaShareResourceSharingConfiguration` | yes |
| `preemptionConfiguration` | `QuotaSharePreemptionConfiguration` | yes |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareName` | `string` | no |
| `quotaShareArn` | `string` | no |

## CreateSchedulingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `quotaSharePolicy` | `QuotaSharePolicy` | no |
| `fairsharePolicy` | `FairsharePolicy` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## CreateServiceEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironmentName` | `string` | yes |
| `serviceEnvironmentType` | `string` | yes |
| `state` | `string` | no |
| `capacityLimits` | `List<CapacityLimit>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironmentName` | `string` | yes |
| `serviceEnvironmentArn` | `string` | yes |

## DeleteComputeEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConsumableResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteJobQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQuotaShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchedulingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeComputeEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironments` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironments` | `List<ComputeEnvironmentDetail>` | no |
| `nextToken` | `string` | no |

## DescribeConsumableResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResourceName` | `string` | yes |
| `consumableResourceArn` | `string` | yes |
| `totalQuantity` | `long` | no |
| `inUseQuantity` | `long` | no |
| `availableQuantity` | `long` | no |
| `resourceType` | `string` | no |
| `createdAt` | `long` | no |
| `tags` | `Map<string>` | no |

## DescribeJobDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDefinitions` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `jobDefinitionName` | `string` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDefinitions` | `List<JobDefinition>` | no |
| `nextToken` | `string` | no |

## DescribeJobQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueues` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueues` | `List<JobQueueDetail>` | no |
| `nextToken` | `string` | no |

## DescribeJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobDetail>` | no |

## DescribeQuotaShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareName` | `string` | no |
| `quotaShareArn` | `string` | no |
| `jobQueueArn` | `string` | no |
| `capacityLimits` | `List<QuotaShareCapacityLimit>` | no |
| `resourceSharingConfiguration` | `QuotaShareResourceSharingConfiguration` | no |
| `preemptionConfiguration` | `QuotaSharePreemptionConfiguration` | no |
| `state` | `string` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |

## DescribeSchedulingPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schedulingPolicies` | `List<SchedulingPolicyDetail>` | no |

## DescribeServiceEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironments` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironments` | `List<ServiceEnvironmentDetail>` | no |
| `nextToken` | `string` | no |

## DescribeServiceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attempts` | `List<ServiceJobAttemptDetail>` | no |
| `capacityUsage` | `List<ServiceJobCapacityUsageDetail>` | no |
| `createdAt` | `long` | no |
| `isTerminated` | `boolean` | no |
| `jobArn` | `string` | no |
| `jobId` | `string` | yes |
| `jobName` | `string` | yes |
| `jobQueue` | `string` | yes |
| `latestAttempt` | `LatestServiceJobAttempt` | no |
| `retryStrategy` | `ServiceJobRetryStrategy` | no |
| `scheduledAt` | `long` | no |
| `schedulingPriority` | `integer` | no |
| `serviceRequestPayload` | `string` | no |
| `serviceJobType` | `string` | yes |
| `shareIdentifier` | `string` | no |
| `quotaShareName` | `string` | no |
| `preemptionConfiguration` | `ServiceJobPreemptionConfiguration` | no |
| `preemptionSummary` | `ServiceJobPreemptionSummary` | no |
| `startedAt` | `long` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `stoppedAt` | `long` | no |
| `tags` | `Map<string>` | no |
| `timeoutConfig` | `ServiceJobTimeout` | no |

## GetJobQueueSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `frontOfQueue` | `FrontOfQueueDetail` | no |
| `frontOfQuotaShares` | `FrontOfQuotaSharesDetail` | no |
| `queueUtilization` | `QueueSnapshotUtilizationDetail` | no |

## ListConsumableResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<KeyValuesPair>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResources` | `List<ConsumableResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | no |
| `arrayJobId` | `string` | no |
| `multiNodeJobId` | `string` | no |
| `jobStatus` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<KeyValuesPair>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaryList` | `List<JobSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobsByConsumableResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResource` | `string` | yes |
| `filters` | `List<KeyValuesPair>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<ListJobsByConsumableResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListQuotaShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShares` | `List<QuotaShareDetail>` | no |
| `nextToken` | `string` | no |

## ListSchedulingPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schedulingPolicies` | `List<SchedulingPolicyListingDetail>` | no |
| `nextToken` | `string` | no |

## ListServiceJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | no |
| `jobStatus` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<KeyValuesPair>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaryList` | `List<ServiceJobSummary>` | yes |
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

## RegisterJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDefinitionName` | `string` | yes |
| `type` | `string` | yes |
| `parameters` | `Map<string>` | no |
| `schedulingPriority` | `integer` | no |
| `containerProperties` | `ContainerProperties` | no |
| `nodeProperties` | `NodeProperties` | no |
| `retryStrategy` | `RetryStrategy` | no |
| `propagateTags` | `boolean` | no |
| `timeout` | `JobTimeout` | no |
| `tags` | `Map<string>` | no |
| `platformCapabilities` | `List<string>` | no |
| `eksProperties` | `EksProperties` | no |
| `ecsProperties` | `EcsProperties` | no |
| `consumableResourceProperties` | `ConsumableResourceProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobDefinitionName` | `string` | yes |
| `jobDefinitionArn` | `string` | yes |
| `revision` | `integer` | yes |

## SubmitJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `jobQueue` | `string` | yes |
| `shareIdentifier` | `string` | no |
| `schedulingPriorityOverride` | `integer` | no |
| `arrayProperties` | `ArrayProperties` | no |
| `dependsOn` | `List<JobDependency>` | no |
| `jobDefinition` | `string` | yes |
| `parameters` | `Map<string>` | no |
| `containerOverrides` | `ContainerOverrides` | no |
| `nodeOverrides` | `NodeOverrides` | no |
| `retryStrategy` | `RetryStrategy` | no |
| `propagateTags` | `boolean` | no |
| `timeout` | `JobTimeout` | no |
| `tags` | `Map<string>` | no |
| `eksPropertiesOverride` | `EksPropertiesOverride` | no |
| `ecsPropertiesOverride` | `EcsPropertiesOverride` | no |
| `consumableResourcePropertiesOverride` | `ConsumableResourceProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobName` | `string` | yes |
| `jobId` | `string` | yes |

## SubmitServiceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `jobQueue` | `string` | yes |
| `retryStrategy` | `ServiceJobRetryStrategy` | no |
| `schedulingPriority` | `integer` | no |
| `serviceRequestPayload` | `string` | yes |
| `serviceJobType` | `string` | yes |
| `shareIdentifier` | `string` | no |
| `quotaShareName` | `string` | no |
| `preemptionConfiguration` | `ServiceJobPreemptionConfiguration` | no |
| `timeoutConfig` | `ServiceJobTimeout` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobName` | `string` | yes |
| `jobId` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateServiceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `reason` | `string` | yes |

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


## UpdateComputeEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironment` | `string` | yes |
| `state` | `string` | no |
| `unmanagedvCpus` | `integer` | no |
| `computeResources` | `ComputeResourceUpdate` | no |
| `serviceRole` | `string` | no |
| `updatePolicy` | `UpdatePolicy` | no |
| `context` | `string` | no |
| `ecsSettings` | `EcsSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `computeEnvironmentName` | `string` | no |
| `computeEnvironmentArn` | `string` | no |

## UpdateConsumableResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResource` | `string` | yes |
| `operation` | `string` | no |
| `quantity` | `long` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consumableResourceName` | `string` | yes |
| `consumableResourceArn` | `string` | yes |
| `totalQuantity` | `long` | no |

## UpdateJobQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueue` | `string` | yes |
| `state` | `string` | no |
| `schedulingPolicyArn` | `string` | no |
| `priority` | `integer` | no |
| `computeEnvironmentOrder` | `List<ComputeEnvironmentOrder>` | no |
| `serviceEnvironmentOrder` | `List<ServiceEnvironmentOrder>` | no |
| `jobStateTimeLimitActions` | `List<JobStateTimeLimitAction>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobQueueName` | `string` | no |
| `jobQueueArn` | `string` | no |

## UpdateQuotaShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareArn` | `string` | yes |
| `capacityLimits` | `List<QuotaShareCapacityLimit>` | no |
| `resourceSharingConfiguration` | `QuotaShareResourceSharingConfiguration` | no |
| `preemptionConfiguration` | `QuotaSharePreemptionConfiguration` | no |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quotaShareName` | `string` | no |
| `quotaShareArn` | `string` | no |

## UpdateSchedulingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `quotaSharePolicy` | `QuotaSharePolicy` | no |
| `fairsharePolicy` | `FairsharePolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironment` | `string` | yes |
| `state` | `string` | no |
| `capacityLimits` | `List<CapacityLimit>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceEnvironmentName` | `string` | yes |
| `serviceEnvironmentArn` | `string` | yes |

## UpdateServiceJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |
| `schedulingPriority` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobName` | `string` | no |
| `jobId` | `string` | no |

