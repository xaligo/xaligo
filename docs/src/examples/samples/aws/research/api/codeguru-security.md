# Amazon CodeGuru Security

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codeguru-security/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingIdentifiers` | `List<FindingIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<Finding>` | yes |
| `failedFindings` | `List<BatchGetFindingsError>` | yes |

## CreateScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `resourceId` | `ResourceId` | yes |
| `scanName` | `string` | yes |
| `scanType` | `string` | no |
| `analysisType` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |
| `runId` | `string` | yes |
| `resourceId` | `ResourceId` | yes |
| `scanState` | `string` | yes |
| `scanNameArn` | `string` | no |

## CreateUploadUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3Url` | `string` | yes |
| `requestHeaders` | `Map<string>` | yes |
| `codeArtifactId` | `string` | yes |

## GetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionConfig` | `EncryptionConfig` | yes |

## GetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<Finding>` | no |
| `nextToken` | `string` | no |

## GetMetricsSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `date` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricsSummary` | `MetricsSummary` | no |

## GetScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |
| `runId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanName` | `string` | yes |
| `runId` | `string` | yes |
| `scanState` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `analysisType` | `string` | yes |
| `updatedAt` | `timestamp` | no |
| `numberOfRevisions` | `long` | no |
| `scanNameArn` | `string` | no |
| `errorMessage` | `string` | no |

## ListFindingsMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `startDate` | `timestamp` | yes |
| `endDate` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsMetrics` | `List<AccountFindingsMetric>` | no |
| `nextToken` | `string` | no |

## ListScans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<ScanSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

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


## UpdateAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionConfig` | `EncryptionConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionConfig` | `EncryptionConfig` | yes |

