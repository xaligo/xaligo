# Braket

API version: 2019-09-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/braket/2019-09-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `cancellationStatus` | `string` | yes |

## CancelQuantumTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTaskArn` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTaskArn` | `string` | yes |
| `cancellationStatus` | `string` | yes |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `algorithmSpecification` | `AlgorithmSpecification` | yes |
| `inputDataConfig` | `List<InputFileConfig>` | no |
| `outputDataConfig` | `JobOutputDataConfig` | yes |
| `checkpointConfig` | `JobCheckpointConfig` | no |
| `jobName` | `string` | yes |
| `roleArn` | `string` | yes |
| `stoppingCondition` | `JobStoppingCondition` | no |
| `instanceConfig` | `InstanceConfig` | yes |
| `hyperParameters` | `Map<string>` | no |
| `deviceConfig` | `DeviceConfig` | yes |
| `tags` | `Map<string>` | no |
| `associations` | `List<Association>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateQuantumTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `deviceArn` | `string` | yes |
| `deviceParameters` | `string` | no |
| `shots` | `long` | yes |
| `outputS3Bucket` | `string` | yes |
| `outputS3KeyPrefix` | `string` | yes |
| `action` | `string` | yes |
| `tags` | `Map<string>` | no |
| `jobToken` | `string` | no |
| `associations` | `List<Association>` | no |
| `experimentalCapabilities` | `ExperimentalCapabilities` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTaskArn` | `string` | yes |

## CreateSpendingLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `deviceArn` | `string` | yes |
| `spendingLimit` | `string` | yes |
| `timePeriod` | `TimePeriod` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spendingLimitArn` | `string` | yes |

## DeleteSpendingLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spendingLimitArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceArn` | `string` | yes |
| `deviceName` | `string` | yes |
| `providerName` | `string` | yes |
| `deviceType` | `string` | yes |
| `deviceStatus` | `string` | yes |
| `deviceCapabilities` | `string` | yes |
| `deviceQueueInfo` | `List<DeviceQueueInfo>` | no |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `additionalAttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `jobArn` | `string` | yes |
| `roleArn` | `string` | yes |
| `failureReason` | `string` | no |
| `jobName` | `string` | yes |
| `hyperParameters` | `Map<string>` | no |
| `inputDataConfig` | `List<InputFileConfig>` | no |
| `outputDataConfig` | `JobOutputDataConfig` | yes |
| `stoppingCondition` | `JobStoppingCondition` | no |
| `checkpointConfig` | `JobCheckpointConfig` | no |
| `algorithmSpecification` | `AlgorithmSpecification` | yes |
| `instanceConfig` | `InstanceConfig` | yes |
| `createdAt` | `timestamp` | yes |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `billableDuration` | `integer` | no |
| `deviceConfig` | `DeviceConfig` | no |
| `events` | `List<JobEventDetails>` | no |
| `tags` | `Map<string>` | no |
| `queueInfo` | `HybridJobQueueInfo` | no |
| `associations` | `List<Association>` | no |

## GetQuantumTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTaskArn` | `string` | yes |
| `additionalAttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTaskArn` | `string` | yes |
| `status` | `string` | yes |
| `failureReason` | `string` | no |
| `deviceArn` | `string` | yes |
| `deviceParameters` | `string` | yes |
| `shots` | `long` | yes |
| `outputS3Bucket` | `string` | yes |
| `outputS3Directory` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `endedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `jobArn` | `string` | no |
| `queueInfo` | `QuantumTaskQueueInfo` | no |
| `associations` | `List<Association>` | no |
| `numSuccessfulShots` | `long` | no |
| `actionMetadata` | `ActionMetadata` | no |
| `experimentalCapabilities` | `ExperimentalCapabilities` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## SearchDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<SearchDevicesFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devices` | `List<DeviceSummary>` | yes |
| `nextToken` | `string` | no |

## SearchJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<SearchJobsFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<JobSummary>` | yes |
| `nextToken` | `string` | no |

## SearchQuantumTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<SearchQuantumTasksFilter>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `quantumTasks` | `List<QuantumTaskSummary>` | yes |
| `nextToken` | `string` | no |

## SearchSpendingLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<SearchSpendingLimitsFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spendingLimits` | `List<SpendingLimitSummary>` | yes |
| `nextToken` | `string` | no |

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


## UpdateSpendingLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spendingLimitArn` | `string` | yes |
| `clientToken` | `string` | yes |
| `spendingLimit` | `string` | no |
| `timePeriod` | `TimePeriod` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


