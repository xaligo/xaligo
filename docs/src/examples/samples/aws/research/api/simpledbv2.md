# Amazon SimpleDB v2

API version: 2025-09-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/simpledbv2/2025-09-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportArn` | `string` | yes |
| `clientToken` | `string` | yes |
| `exportStatus` | `string` | yes |
| `domainName` | `string` | yes |
| `requestedAt` | `timestamp` | yes |
| `s3Bucket` | `string` | yes |
| `s3KeyPrefix` | `string` | no |
| `s3SseAlgorithm` | `string` | no |
| `s3SseKmsKeyId` | `string` | no |
| `s3BucketOwner` | `string` | no |
| `failureCode` | `string` | no |
| `failureMessage` | `string` | no |
| `exportManifest` | `string` | no |
| `itemsCount` | `long` | no |
| `exportDataCutoffTime` | `timestamp` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportSummaries` | `List<ExportSummary>` | yes |
| `nextToken` | `string` | no |

## StartDomainExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `domainName` | `string` | yes |
| `s3Bucket` | `string` | yes |
| `s3KeyPrefix` | `string` | no |
| `s3SseAlgorithm` | `string` | no |
| `s3SseKmsKeyId` | `string` | no |
| `s3BucketOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `exportArn` | `string` | yes |
| `requestedAt` | `timestamp` | yes |

