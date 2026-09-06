# Amazon S3 Tables

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/s3tables/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `List<string>` | yes |

## CreateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `format` | `string` | yes |
| `metadata` | `TableMetadata` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `storageClassConfiguration` | `StorageClassConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableARN` | `string` | yes |
| `versionToken` | `string` | yes |

## CreateTableBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `storageClassConfiguration` | `StorageClassConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

## DeleteNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `versionToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableBucketEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableBucketMetricsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `versionToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTablePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTableReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |
| `versionToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `List<string>` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `ownerAccountId` | `string` | yes |
| `namespaceId` | `string` | no |
| `tableBucketId` | `string` | no |

## GetTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | no |
| `namespace` | `string` | no |
| `name` | `string` | no |
| `tableArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `tableARN` | `string` | yes |
| `namespace` | `List<string>` | yes |
| `namespaceId` | `string` | no |
| `versionToken` | `string` | yes |
| `metadataLocation` | `string` | no |
| `warehouseLocation` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `managedByService` | `string` | no |
| `modifiedAt` | `timestamp` | yes |
| `modifiedBy` | `string` | yes |
| `ownerAccountId` | `string` | yes |
| `format` | `string` | yes |
| `tableBucketId` | `string` | no |
| `managedTableInformation` | `ManagedTableInformation` | no |

## GetTableBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `ownerAccountId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `tableBucketId` | `string` | no |
| `type` | `string` | no |

## GetTableBucketEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionConfiguration` | `EncryptionConfiguration` | yes |

## GetTableBucketMaintenanceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `configuration` | `Map<TableBucketMaintenanceConfigurationValue>` | yes |

## GetTableBucketMetricsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `id` | `string` | no |

## GetTableBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `string` | yes |

## GetTableBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionToken` | `string` | yes |
| `configuration` | `TableBucketReplicationConfiguration` | yes |

## GetTableBucketStorageClass

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageClassConfiguration` | `StorageClassConfiguration` | yes |

## GetTableEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionConfiguration` | `EncryptionConfiguration` | yes |

## GetTableMaintenanceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableARN` | `string` | yes |
| `configuration` | `Map<TableMaintenanceConfigurationValue>` | yes |

## GetTableMaintenanceJobStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableARN` | `string` | yes |
| `status` | `Map<TableMaintenanceJobStatusValue>` | yes |

## GetTableMetadataLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionToken` | `string` | yes |
| `metadataLocation` | `string` | no |
| `warehouseLocation` | `string` | yes |

## GetTablePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `string` | yes |

## GetTableRecordExpirationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `TableRecordExpirationConfigurationValue` | yes |

## GetTableRecordExpirationJobStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `lastRunTimestamp` | `timestamp` | no |
| `failureMessage` | `string` | no |
| `metrics` | `TableRecordExpirationJobMetrics` | no |

## GetTableReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionToken` | `string` | yes |
| `configuration` | `TableReplicationConfiguration` | yes |

## GetTableReplicationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceTableArn` | `string` | yes |
| `destinations` | `List<ReplicationDestinationStatusModel>` | yes |

## GetTableStorageClass

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageClassConfiguration` | `StorageClassConfiguration` | yes |

## ListNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `prefix` | `string` | no |
| `continuationToken` | `string` | no |
| `maxNamespaces` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaces` | `List<NamespaceSummary>` | yes |
| `continuationToken` | `string` | no |

## ListTableBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `prefix` | `string` | no |
| `continuationToken` | `string` | no |
| `maxBuckets` | `integer` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBuckets` | `List<TableBucketSummary>` | yes |
| `continuationToken` | `string` | no |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | no |
| `prefix` | `string` | no |
| `continuationToken` | `string` | no |
| `maxTables` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tables` | `List<TableSummary>` | yes |
| `continuationToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutTableBucketEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableBucketMaintenanceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `type` | `string` | yes |
| `value` | `TableBucketMaintenanceConfigurationValue` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableBucketMetricsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `resourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableBucketReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `versionToken` | `string` | no |
| `configuration` | `TableBucketReplicationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionToken` | `string` | yes |
| `status` | `string` | yes |

## PutTableBucketStorageClass

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `storageClassConfiguration` | `StorageClassConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableMaintenanceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `value` | `TableMaintenanceConfigurationValue` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTablePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `resourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableRecordExpirationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |
| `value` | `TableRecordExpirationConfigurationValue` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTableReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableArn` | `string` | yes |
| `versionToken` | `string` | no |
| `configuration` | `TableReplicationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versionToken` | `string` | yes |
| `status` | `string` | yes |

## RenameTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `newNamespaceName` | `string` | no |
| `newName` | `string` | no |
| `versionToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateTableMetadataLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tableBucketARN` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `versionToken` | `string` | yes |
| `metadataLocation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `tableARN` | `string` | yes |
| `namespace` | `List<string>` | yes |
| `versionToken` | `string` | yes |
| `metadataLocation` | `string` | yes |

