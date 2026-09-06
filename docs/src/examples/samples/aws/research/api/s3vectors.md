# Amazon S3 Vectors

API version: 2025-07-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/s3vectors/2025-07-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |
| `indexName` | `string` | yes |
| `dataType` | `string` | yes |
| `dimension` | `integer` | yes |
| `distanceMetric` | `string` | yes |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexArn` | `string` | yes |

## CreateVectorBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketArn` | `string` | yes |

## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVectorBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVectorBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |
| `keys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `index` | `Index` | yes |

## GetVectorBucket

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucket` | `VectorBucket` | yes |

## GetVectorBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |

## GetVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |
| `keys` | `List<string>` | yes |
| `returnData` | `boolean` | no |
| `returnMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectors` | `List<GetOutputVector>` | yes |

## ListIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `indexes` | `List<IndexSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## ListVectorBuckets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `vectorBuckets` | `List<VectorBucketSummary>` | yes |

## ListVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `segmentCount` | `integer` | no |
| `segmentIndex` | `integer` | no |
| `returnData` | `boolean` | no |
| `returnMetadata` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `vectors` | `List<ListOutputVector>` | yes |

## PutVectorBucketPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `vectorBucketArn` | `string` | no |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |
| `vectors` | `List<PutInputVector>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## QueryVectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectorBucketName` | `string` | no |
| `indexName` | `string` | no |
| `indexArn` | `string` | no |
| `topK` | `integer` | yes |
| `queryVector` | `VectorData` | yes |
| `filter` | `Document` | no |
| `returnMetadata` | `boolean` | no |
| `returnDistance` | `boolean` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vectors` | `List<QueryOutputVector>` | yes |
| `distanceMetric` | `string` | yes |
| `nextToken` | `string` | no |

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


