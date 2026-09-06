# AWS Cloud Control API

API version: 2021-09-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudcontrol/2021-09-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelResourceRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressEvent` | `ProgressEvent` | no |

## CreateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `TypeVersionId` | `string` | no |
| `RoleArn` | `string` | no |
| `ClientToken` | `string` | no |
| `DesiredState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressEvent` | `ProgressEvent` | no |

## DeleteResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `TypeVersionId` | `string` | no |
| `RoleArn` | `string` | no |
| `ClientToken` | `string` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressEvent` | `ProgressEvent` | no |

## GetResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `TypeVersionId` | `string` | no |
| `RoleArn` | `string` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | no |
| `ResourceDescription` | `ResourceDescription` | no |

## GetResourceRequestStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressEvent` | `ProgressEvent` | no |
| `HooksProgressEvent` | `List<HookProgressEvent>` | no |

## ListResourceRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceRequestStatusFilter` | `ResourceRequestStatusFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceRequestStatusSummaries` | `List<ProgressEvent>` | no |
| `NextToken` | `string` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `TypeVersionId` | `string` | no |
| `RoleArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceModel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | no |
| `ResourceDescriptions` | `List<ResourceDescription>` | no |
| `NextToken` | `string` | no |

## UpdateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | yes |
| `TypeVersionId` | `string` | no |
| `RoleArn` | `string` | no |
| `ClientToken` | `string` | no |
| `Identifier` | `string` | yes |
| `PatchDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressEvent` | `ProgressEvent` | no |

