# Amazon Kinesis Firehose

API version: 2015-08-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/firehose/2015-08-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `DeliveryStreamType` | `string` | no |
| `DirectPutSourceConfiguration` | `DirectPutSourceConfiguration` | no |
| `KinesisStreamSourceConfiguration` | `KinesisStreamSourceConfiguration` | no |
| `DeliveryStreamEncryptionConfigurationInput` | `DeliveryStreamEncryptionConfigurationInput` | no |
| `S3DestinationConfiguration` | `S3DestinationConfiguration` | no |
| `ExtendedS3DestinationConfiguration` | `ExtendedS3DestinationConfiguration` | no |
| `RedshiftDestinationConfiguration` | `RedshiftDestinationConfiguration` | no |
| `ElasticsearchDestinationConfiguration` | `ElasticsearchDestinationConfiguration` | no |
| `AmazonopensearchserviceDestinationConfiguration` | `AmazonopensearchserviceDestinationConfiguration` | no |
| `SplunkDestinationConfiguration` | `SplunkDestinationConfiguration` | no |
| `HttpEndpointDestinationConfiguration` | `HttpEndpointDestinationConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `AmazonOpenSearchServerlessDestinationConfiguration` | `AmazonOpenSearchServerlessDestinationConfiguration` | no |
| `MSKSourceConfiguration` | `MSKSourceConfiguration` | no |
| `SnowflakeDestinationConfiguration` | `SnowflakeDestinationConfiguration` | no |
| `IcebergDestinationConfiguration` | `IcebergDestinationConfiguration` | no |
| `DatabaseSourceConfiguration` | `DatabaseSourceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamARN` | `string` | no |

## DeleteDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `AllowForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `Limit` | `integer` | no |
| `ExclusiveStartDestinationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamDescription` | `DeliveryStreamDescription` | yes |

## ListDeliveryStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `DeliveryStreamType` | `string` | no |
| `ExclusiveStartDeliveryStreamName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamNames` | `List<string>` | yes |
| `HasMoreDeliveryStreams` | `boolean` | yes |

## ListTagsForDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `ExclusiveStartTagKey` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |
| `HasMoreTags` | `boolean` | yes |

## PutRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `Record` | `Record` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordId` | `string` | yes |
| `Encrypted` | `boolean` | no |

## PutRecordBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `Records` | `List<Record>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedPutCount` | `integer` | yes |
| `Encrypted` | `boolean` | no |
| `RequestResponses` | `List<PutRecordBatchResponseEntry>` | yes |

## StartDeliveryStreamEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `DeliveryStreamEncryptionConfigurationInput` | `DeliveryStreamEncryptionConfigurationInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopDeliveryStreamEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagDeliveryStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryStreamName` | `string` | yes |
| `CurrentDeliveryStreamVersionId` | `string` | yes |
| `DestinationId` | `string` | yes |
| `S3DestinationUpdate` | `S3DestinationUpdate` | no |
| `ExtendedS3DestinationUpdate` | `ExtendedS3DestinationUpdate` | no |
| `RedshiftDestinationUpdate` | `RedshiftDestinationUpdate` | no |
| `ElasticsearchDestinationUpdate` | `ElasticsearchDestinationUpdate` | no |
| `AmazonopensearchserviceDestinationUpdate` | `AmazonopensearchserviceDestinationUpdate` | no |
| `SplunkDestinationUpdate` | `SplunkDestinationUpdate` | no |
| `HttpEndpointDestinationUpdate` | `HttpEndpointDestinationUpdate` | no |
| `AmazonOpenSearchServerlessDestinationUpdate` | `AmazonOpenSearchServerlessDestinationUpdate` | no |
| `SnowflakeDestinationUpdate` | `SnowflakeDestinationUpdate` | no |
| `IcebergDestinationUpdate` | `IcebergDestinationUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


