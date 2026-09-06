# Amazon Elastic Block Store

API version: 2019-11-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ebs/2019-11-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CompleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `ChangedBlocksCount` | `integer` | yes |
| `Checksum` | `string` | no |
| `ChecksumAlgorithm` | `string` | no |
| `ChecksumAggregationMethod` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## GetSnapshotBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `BlockIndex` | `integer` | yes |
| `BlockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataLength` | `integer` | no |
| `BlockData` | `blob` | no |
| `Checksum` | `string` | no |
| `ChecksumAlgorithm` | `string` | no |

## ListChangedBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirstSnapshotId` | `string` | no |
| `SecondSnapshotId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartingBlockIndex` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangedBlocks` | `List<ChangedBlock>` | no |
| `ExpiryTime` | `timestamp` | no |
| `VolumeSize` | `long` | no |
| `BlockSize` | `integer` | no |
| `NextToken` | `string` | no |

## ListSnapshotBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartingBlockIndex` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blocks` | `List<Block>` | no |
| `ExpiryTime` | `timestamp` | no |
| `VolumeSize` | `long` | no |
| `BlockSize` | `integer` | no |
| `NextToken` | `string` | no |

## PutSnapshotBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `BlockIndex` | `integer` | yes |
| `BlockData` | `blob` | yes |
| `DataLength` | `integer` | yes |
| `Progress` | `integer` | no |
| `Checksum` | `string` | yes |
| `ChecksumAlgorithm` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Checksum` | `string` | no |
| `ChecksumAlgorithm` | `string` | no |

## StartSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeSize` | `long` | yes |
| `ParentSnapshotId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyArn` | `string` | no |
| `Timeout` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `SnapshotId` | `string` | no |
| `OwnerId` | `string` | no |
| `Status` | `string` | no |
| `StartTime` | `timestamp` | no |
| `VolumeSize` | `long` | no |
| `BlockSize` | `integer` | no |
| `Tags` | `List<Tag>` | no |
| `ParentSnapshotId` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `SseType` | `string` | no |

