# AWS Snow Device Management

API version: 2021-08-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/snow-device-management/2021-08-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## CreateTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `command` | `Command` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `targets` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskArn` | `string` | no |
| `taskId` | `string` | no |

## DescribeDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedWithJob` | `string` | no |
| `deviceCapacities` | `List<Capacity>` | no |
| `deviceState` | `string` | no |
| `deviceType` | `string` | no |
| `lastReachedOutAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `managedDeviceArn` | `string` | no |
| `managedDeviceId` | `string` | no |
| `physicalNetworkInterfaces` | `List<PhysicalNetworkInterface>` | no |
| `software` | `SoftwareInformation` | no |
| `tags` | `Map<string>` | no |

## DescribeDeviceEc2Instances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceIds` | `List<string>` | yes |
| `managedDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instances` | `List<InstanceSummary>` | no |

## DescribeExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedDeviceId` | `string` | yes |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `managedDeviceId` | `string` | no |
| `startedAt` | `timestamp` | no |
| `state` | `string` | no |
| `taskId` | `string` | no |

## DescribeTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `completedAt` | `timestamp` | no |
| `createdAt` | `timestamp` | no |
| `description` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |
| `targets` | `List<string>` | no |
| `taskArn` | `string` | no |
| `taskId` | `string` | no |

## ListDeviceResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedDeviceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `resources` | `List<ResourceSummary>` | no |

## ListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devices` | `List<DeviceSummary>` | no |
| `nextToken` | `string` | no |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executions` | `List<ExecutionSummary>` | no |
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
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tasks` | `List<TaskSummary>` | no |

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


