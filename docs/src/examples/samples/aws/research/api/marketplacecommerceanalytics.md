# AWS Marketplace Commerce Analytics

API version: 2015-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplacecommerceanalytics/2015-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GenerateDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSetType` | `string` | yes |
| `dataSetPublicationDate` | `timestamp` | yes |
| `roleNameArn` | `string` | yes |
| `destinationS3BucketName` | `string` | yes |
| `destinationS3Prefix` | `string` | no |
| `snsTopicArn` | `string` | yes |
| `customerDefinedValues` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSetRequestId` | `string` | no |

## StartSupportDataExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSetType` | `string` | yes |
| `fromDate` | `timestamp` | yes |
| `roleNameArn` | `string` | yes |
| `destinationS3BucketName` | `string` | yes |
| `destinationS3Prefix` | `string` | no |
| `snsTopicArn` | `string` | yes |
| `customerDefinedValues` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSetRequestId` | `string` | no |

