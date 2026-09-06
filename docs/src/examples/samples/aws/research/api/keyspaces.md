# Amazon Keyspaces

API version: 2022-02-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/keyspaces/2022-02-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateKeyspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `replicationSpecification` | `ReplicationSpecification` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

## CreateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |
| `schemaDefinition` | `SchemaDefinition` | yes |
| `comment` | `Comment` | no |
| `capacitySpecification` | `CapacitySpecification` | no |
| `encryptionSpecification` | `EncryptionSpecification` | no |
| `pointInTimeRecovery` | `PointInTimeRecovery` | no |
| `ttl` | `TimeToLive` | no |
| `defaultTimeToLive` | `integer` | no |
| `tags` | `List<Tag>` | no |
| `clientSideTimestamps` | `ClientSideTimestamps` | no |
| `autoScalingSpecification` | `AutoScalingSpecification` | no |
| `replicaSpecifications` | `List<ReplicaSpecification>` | no |
| `cdcSpecification` | `CdcSpecification` | no |
| `warmThroughputSpecification` | `WarmThroughputSpecification` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

## CreateType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldDefinitions` | `List<FieldDefinition>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceArn` | `string` | yes |
| `typeName` | `string` | yes |

## DeleteKeyspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `typeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceArn` | `string` | yes |
| `typeName` | `string` | yes |

## GetKeyspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `resourceArn` | `string` | yes |
| `replicationStrategy` | `string` | yes |
| `replicationRegions` | `List<string>` | no |
| `replicationGroupStatuses` | `List<ReplicationGroupStatus>` | no |

## GetTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |
| `resourceArn` | `string` | yes |
| `creationTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `schemaDefinition` | `SchemaDefinition` | no |
| `capacitySpecification` | `CapacitySpecificationSummary` | no |
| `encryptionSpecification` | `EncryptionSpecification` | no |
| `pointInTimeRecovery` | `PointInTimeRecoverySummary` | no |
| `ttl` | `TimeToLive` | no |
| `defaultTimeToLive` | `integer` | no |
| `comment` | `Comment` | no |
| `clientSideTimestamps` | `ClientSideTimestamps` | no |
| `replicaSpecifications` | `List<ReplicaSpecificationSummary>` | no |
| `latestStreamArn` | `string` | no |
| `cdcSpecification` | `CdcSpecificationSummary` | no |
| `warmThroughputSpecification` | `WarmThroughputSpecificationSummary` | no |

## GetTableAutoScalingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |
| `resourceArn` | `string` | yes |
| `autoScalingSpecification` | `AutoScalingSpecification` | no |
| `replicaSpecifications` | `List<ReplicaAutoScalingSpecification>` | no |

## GetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `typeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldDefinitions` | `List<FieldDefinition>` | no |
| `lastModifiedTimestamp` | `timestamp` | no |
| `status` | `string` | no |
| `directReferringTables` | `List<string>` | no |
| `directParentTypes` | `List<string>` | no |
| `maxNestingDepth` | `integer` | no |
| `keyspaceArn` | `string` | yes |

## ListKeyspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `keyspaces` | `List<KeyspaceSummary>` | yes |

## ListTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `keyspaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tables` | `List<TableSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tags` | `List<Tag>` | no |

## ListTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `keyspaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `types` | `List<string>` | yes |

## RestoreTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceKeyspaceName` | `string` | yes |
| `sourceTableName` | `string` | yes |
| `targetKeyspaceName` | `string` | yes |
| `targetTableName` | `string` | yes |
| `restoreTimestamp` | `timestamp` | no |
| `capacitySpecificationOverride` | `CapacitySpecification` | no |
| `encryptionSpecificationOverride` | `EncryptionSpecification` | no |
| `pointInTimeRecoveryOverride` | `PointInTimeRecovery` | no |
| `tagsOverride` | `List<Tag>` | no |
| `autoScalingSpecification` | `AutoScalingSpecification` | no |
| `replicaSpecifications` | `List<ReplicaSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restoredTableARN` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateKeyspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `replicationSpecification` | `ReplicationSpecification` | yes |
| `clientSideTimestamps` | `ClientSideTimestamps` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

## UpdateTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |
| `addColumns` | `List<ColumnDefinition>` | no |
| `capacitySpecification` | `CapacitySpecification` | no |
| `encryptionSpecification` | `EncryptionSpecification` | no |
| `pointInTimeRecovery` | `PointInTimeRecovery` | no |
| `ttl` | `TimeToLive` | no |
| `defaultTimeToLive` | `integer` | no |
| `clientSideTimestamps` | `ClientSideTimestamps` | no |
| `autoScalingSpecification` | `AutoScalingSpecification` | no |
| `replicaSpecifications` | `List<ReplicaSpecification>` | no |
| `cdcSpecification` | `CdcSpecification` | no |
| `warmThroughputSpecification` | `WarmThroughputSpecification` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

