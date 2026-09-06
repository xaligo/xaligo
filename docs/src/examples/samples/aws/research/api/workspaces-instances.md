# Amazon Workspaces Instances

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workspaces-instances/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |
| `VolumeId` | `string` | yes |
| `Device` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | yes |
| `ClientToken` | `string` | no |
| `Encrypted` | `boolean` | no |
| `Iops` | `integer` | no |
| `KmsKeyId` | `string` | no |
| `SizeInGB` | `integer` | no |
| `SnapshotId` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `Throughput` | `integer` | no |
| `VolumeType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | no |

## CreateWorkspaceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ManagedInstance` | `ManagedInstanceRequest` | yes |
| `BillingConfiguration` | `BillingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | no |

## DeleteVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspaceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |
| `VolumeId` | `string` | yes |
| `Device` | `string` | no |
| `DisassociateMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetWorkspaceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceErrors` | `List<WorkspaceInstanceError>` | no |
| `EC2InstanceErrors` | `List<EC2InstanceError>` | no |
| `ProvisionState` | `string` | no |
| `WorkspaceInstanceId` | `string` | no |
| `EC2ManagedInstance` | `EC2ManagedInstance` | no |
| `BillingConfiguration` | `BillingConfiguration` | no |

## ListInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `InstanceConfigurationFilter` | `InstanceConfigurationFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<InstanceTypeInfo>` | yes |
| `NextToken` | `string` | no |

## ListRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<Region>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListWorkspaceInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionStates` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstances` | `List<WorkspaceInstance>` | yes |
| `NextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceInstanceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


