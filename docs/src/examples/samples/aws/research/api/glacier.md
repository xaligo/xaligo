# Amazon Glacier

API version: 2012-06-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/glacier/2012-06-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AbortMultipartUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `uploadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AbortVaultLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddTagsToVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CompleteMultipartUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `uploadId` | `string` | yes |
| `archiveSize` | `string` | no |
| `checksum` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | no |
| `checksum` | `string` | no |
| `archiveId` | `string` | no |

## CompleteVaultLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `lockId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | no |

## DeleteArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `archiveId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `JobDescription` | `string` | no |
| `Action` | `string` | no |
| `ArchiveId` | `string` | no |
| `VaultARN` | `string` | no |
| `CreationDate` | `string` | no |
| `Completed` | `boolean` | no |
| `StatusCode` | `string` | no |
| `StatusMessage` | `string` | no |
| `ArchiveSizeInBytes` | `long` | no |
| `InventorySizeInBytes` | `long` | no |
| `SNSTopic` | `string` | no |
| `CompletionDate` | `string` | no |
| `SHA256TreeHash` | `string` | no |
| `ArchiveSHA256TreeHash` | `string` | no |
| `RetrievalByteRange` | `string` | no |
| `Tier` | `string` | no |
| `InventoryRetrievalParameters` | `InventoryRetrievalJobDescription` | no |
| `JobOutputPath` | `string` | no |
| `SelectParameters` | `SelectParameters` | no |
| `OutputLocation` | `OutputLocation` | no |

## DescribeVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VaultARN` | `string` | no |
| `VaultName` | `string` | no |
| `CreationDate` | `string` | no |
| `LastInventoryDate` | `string` | no |
| `NumberOfArchives` | `long` | no |
| `SizeInBytes` | `long` | no |

## GetDataRetrievalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `DataRetrievalPolicy` | no |

## GetJobOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `jobId` | `string` | yes |
| `range` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | no |
| `checksum` | `string` | no |
| `status` | `integer` | no |
| `contentRange` | `string` | no |
| `acceptRanges` | `string` | no |
| `contentType` | `string` | no |
| `archiveDescription` | `string` | no |

## GetVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `VaultAccessPolicy` | no |

## GetVaultLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `State` | `string` | no |
| `ExpirationDate` | `string` | no |
| `CreationDate` | `string` | no |

## GetVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vaultNotificationConfig` | `VaultNotificationConfig` | no |

## InitiateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `jobParameters` | `JobParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | no |
| `jobId` | `string` | no |
| `jobOutputPath` | `string` | no |

## InitiateMultipartUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `archiveDescription` | `string` | no |
| `partSize` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | no |
| `uploadId` | `string` | no |

## InitiateVaultLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `policy` | `VaultLockPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lockId` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `limit` | `string` | no |
| `marker` | `string` | no |
| `statuscode` | `string` | no |
| `completed` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobList` | `List<GlacierJobDescription>` | no |
| `Marker` | `string` | no |

## ListMultipartUploads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `marker` | `string` | no |
| `limit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UploadsList` | `List<UploadListElement>` | no |
| `Marker` | `string` | no |

## ListParts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `uploadId` | `string` | yes |
| `marker` | `string` | no |
| `limit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultipartUploadId` | `string` | no |
| `VaultARN` | `string` | no |
| `ArchiveDescription` | `string` | no |
| `PartSizeInBytes` | `long` | no |
| `CreationDate` | `string` | no |
| `Parts` | `List<PartListElement>` | no |
| `Marker` | `string` | no |

## ListProvisionedCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedCapacityList` | `List<ProvisionedCapacityDescription>` | no |

## ListTagsForVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListVaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `marker` | `string` | no |
| `limit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VaultList` | `List<DescribeVaultOutput>` | no |
| `Marker` | `string` | no |

## PurchaseProvisionedCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityId` | `string` | no |

## RemoveTagsFromVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `TagKeys` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetDataRetrievalPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `Policy` | `DataRetrievalPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetVaultAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `policy` | `VaultAccessPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetVaultNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `vaultNotificationConfig` | `VaultNotificationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UploadArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vaultName` | `string` | yes |
| `accountId` | `string` | yes |
| `archiveDescription` | `string` | no |
| `checksum` | `string` | no |
| `body` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | no |
| `checksum` | `string` | no |
| `archiveId` | `string` | no |

## UploadMultipartPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |
| `vaultName` | `string` | yes |
| `uploadId` | `string` | yes |
| `checksum` | `string` | no |
| `range` | `string` | no |
| `body` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `checksum` | `string` | no |

