# Agent Registry Control

API version: 2025-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/agent-registry-control/2025-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `discoveryConfiguration` | `DiscoveryConfiguration` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |
| `autoDetectionConfiguration` | `AutoDetectionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |

## CreateRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `recordType` | `string` | yes |
| `descriptors` | `Descriptors` | yes |
| `recordVersion` | `string` | no |
| `clientToken` | `string` | no |
| `provenance` | `List<Provenance>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recordArn` | `string` | yes |
| `status` | `string` | yes |

## DeleteRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `registryId` | `string` | yes |
| `registryArn` | `string` | yes |
| `discoveryConfiguration` | `DiscoveryConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `autoDetection` | `AutoDetection` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `recordType` | `string` | yes |
| `descriptors` | `Descriptors` | no |
| `recordVersion` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `statusReason` | `string` | no |
| `provenance` | `List<Provenance>` | no |
| `createdByAutoDetection` | `boolean` | no |
| `createdBy` | `string` | no |

## ListRegistries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<RegistryFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registries` | `List<RegistrySummary>` | yes |
| `nextToken` | `string` | no |

## ListRegistryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `List<RegistryRecordFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<RegistryRecordSummary>` | yes |
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

## SubmitRegistryRecordForApproval

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

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


## UpdateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `name` | `string` | no |
| `description` | `UpdatedDescription` | no |
| `discoveryConfiguration` | `UpdatedDiscoveryConfiguration` | no |
| `approvalConfiguration` | `UpdatedApprovalConfiguration` | no |
| `autoDetectionConfiguration` | `UpdatedAutoDetectionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `registryId` | `string` | yes |
| `registryArn` | `string` | yes |
| `discoveryConfiguration` | `DiscoveryConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `autoDetection` | `AutoDetection` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | no |
| `displayName` | `UpdatedDisplayName` | no |
| `description` | `UpdatedDescription` | no |
| `recordType` | `string` | no |
| `descriptors` | `UpdatedDescriptors` | no |
| `recordVersion` | `string` | no |
| `triggerSynchronization` | `boolean` | no |
| `provenance` | `List<Provenance>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `recordType` | `string` | yes |
| `descriptors` | `Descriptors` | no |
| `recordVersion` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `statusReason` | `string` | no |
| `provenance` | `List<Provenance>` | no |
| `createdByAutoDetection` | `boolean` | no |
| `createdBy` | `string` | no |

## UpdateRegistryRecordStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

