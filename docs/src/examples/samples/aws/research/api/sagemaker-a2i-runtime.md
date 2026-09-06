# Amazon Augmented AI Runtime

API version: 2019-11-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-a2i-runtime/2019-11-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteHumanLoop

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeHumanLoop

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `FailureCode` | `string` | no |
| `HumanLoopStatus` | `string` | yes |
| `HumanLoopName` | `string` | yes |
| `HumanLoopArn` | `string` | yes |
| `FlowDefinitionArn` | `string` | yes |
| `HumanLoopOutput` | `HumanLoopOutput` | no |

## ListHumanLoops

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `FlowDefinitionArn` | `string` | yes |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopSummaries` | `List<HumanLoopSummary>` | yes |
| `NextToken` | `string` | no |

## StartHumanLoop

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopName` | `string` | yes |
| `FlowDefinitionArn` | `string` | yes |
| `HumanLoopInput` | `HumanLoopInput` | yes |
| `DataAttributes` | `HumanLoopDataAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopArn` | `string` | no |

## StopHumanLoop

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanLoopName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


