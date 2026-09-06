# Amazon Keyspaces Streams

API version: 2024-09-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/keyspacesstreams/2024-09-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shardIterator` | `string` | yes |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `changeRecords` | `List<Record>` | no |
| `nextShardIterator` | `string` | no |
| `iteratorDescription` | `IteratorDescription` | no |

## GetShardIterator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamArn` | `string` | yes |
| `shardId` | `string` | yes |
| `shardIteratorType` | `string` | yes |
| `sequenceNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `shardIterator` | `string` | no |

## GetStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `shardFilter` | `ShardFilter` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streamArn` | `string` | yes |
| `streamLabel` | `string` | yes |
| `streamStatus` | `string` | yes |
| `streamViewType` | `string` | yes |
| `creationRequestDateTime` | `timestamp` | yes |
| `keyspaceName` | `string` | yes |
| `tableName` | `string` | yes |
| `shards` | `List<Shard>` | no |
| `nextToken` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyspaceName` | `string` | no |
| `tableName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `streams` | `List<Stream>` | no |
| `nextToken` | `string` | no |

