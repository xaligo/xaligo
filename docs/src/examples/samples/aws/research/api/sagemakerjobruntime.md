# Sagemaker Job Runtime Service

API version: 2026-02-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemakerjobruntime/2026-02-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CompleteRollout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |
| `TrajectoryId` | `string` | yes |
| `Status` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Sample

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |
| `TrajectoryId` | `string` | yes |
| `Body` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Body` | `blob` | yes |

## SampleWithResponseStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |
| `TrajectoryId` | `string` | yes |
| `Body` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Body` | `blob` | yes |

## UpdateReward

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |
| `TrajectoryId` | `string` | yes |
| `Rewards` | `List<double>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


