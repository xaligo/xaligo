# Amazon Data Lifecycle Manager

API version: 2018-01-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dlm/2018-01-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionRoleArn` | `string` | yes |
| `Description` | `string` | yes |
| `State` | `string` | yes |
| `PolicyDetails` | `PolicyDetails` | no |
| `Tags` | `Map<string>` | no |
| `DefaultPolicy` | `string` | no |
| `CreateInterval` | `integer` | no |
| `RetainInterval` | `integer` | no |
| `CopyTags` | `boolean` | no |
| `ExtendDeletion` | `boolean` | no |
| `CrossRegionCopyTargets` | `List<CrossRegionCopyTarget>` | no |
| `Exclusions` | `Exclusions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | no |

## DeleteLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetLifecyclePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyIds` | `List<string>` | no |
| `State` | `string` | no |
| `ResourceTypes` | `List<string>` | no |
| `TargetTags` | `List<string>` | no |
| `TagsToAdd` | `List<string>` | no |
| `DefaultPolicyType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `List<LifecyclePolicySummary>` | no |

## GetLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `LifecyclePolicy` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyId` | `string` | yes |
| `ExecutionRoleArn` | `string` | no |
| `State` | `string` | no |
| `Description` | `string` | no |
| `PolicyDetails` | `PolicyDetails` | no |
| `CreateInterval` | `integer` | no |
| `RetainInterval` | `integer` | no |
| `CopyTags` | `boolean` | no |
| `ExtendDeletion` | `boolean` | no |
| `CrossRegionCopyTargets` | `List<CrossRegionCopyTarget>` | no |
| `Exclusions` | `Exclusions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


