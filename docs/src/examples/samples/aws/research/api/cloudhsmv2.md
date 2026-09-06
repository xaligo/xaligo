# AWS CloudHSM V2

API version: 2017-04-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudhsmv2/2017-04-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CopyBackupToRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationRegion` | `string` | yes |
| `BackupId` | `string` | yes |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationBackup` | `DestinationBackup` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupRetentionPolicy` | `BackupRetentionPolicy` | no |
| `HsmType` | `string` | yes |
| `SourceBackupId` | `string` | no |
| `SubnetIds` | `List<string>` | yes |
| `NetworkType` | `string` | no |
| `TagList` | `List<Tag>` | no |
| `Mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## CreateHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `AvailabilityZone` | `string` | yes |
| `IpAddress` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Hsm` | `Hsm` | no |

## DeleteBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backup` | `Backup` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DeleteHsm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `HsmId` | `string` | no |
| `EniId` | `string` | no |
| `EniIp` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmId` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |

## DescribeBackups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `Map<List<string>>` | no |
| `Shared` | `boolean` | no |
| `SortAscending` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backups` | `List<Backup>` | no |
| `NextToken` | `string` | no |

## DescribeClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `Map<List<string>>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Clusters` | `List<Cluster>` | no |
| `NextToken` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## InitializeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SignedCert` | `string` | yes |
| `TrustAnchor` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `StateMessage` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | yes |
| `NextToken` | `string` | no |

## ModifyBackupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |
| `NeverExpires` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backup` | `Backup` | no |

## ModifyCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HsmType` | `string` | no |
| `BackupRetentionPolicy` | `BackupRetentionPolicy` | no |
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |

## RestoreBackup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BackupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Backup` | `Backup` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagList` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeyList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


