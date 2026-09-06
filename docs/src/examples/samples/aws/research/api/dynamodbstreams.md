# Amazon DynamoDB Streams

API version: 2012-08-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dynamodbstreams/2012-08-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamArn` | `string` | yes |
| `Limit` | `integer` | no |
| `ExclusiveStartShardId` | `string` | no |
| `ShardFilter` | `ShardFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamDescription` | `StreamDescription` | no |

## GetRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardIterator` | `string` | yes |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | no |
| `NextShardIterator` | `string` | no |

## GetShardIterator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamArn` | `string` | yes |
| `ShardId` | `string` | yes |
| `ShardIteratorType` | `string` | yes |
| `SequenceNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ShardIterator` | `string` | no |

## ListStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TableName` | `string` | no |
| `Limit` | `integer` | no |
| `ExclusiveStartStreamArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Streams` | `List<Stream>` | no |
| `LastEvaluatedStreamArn` | `string` | no |

