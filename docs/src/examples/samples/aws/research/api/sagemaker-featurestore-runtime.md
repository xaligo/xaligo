# Amazon SageMaker Feature Store Runtime

API version: 2020-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker-featurestore-runtime/2020-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifiers` | `List<BatchGetRecordIdentifier>` | yes |
| `ExpirationTimeResponse` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<BatchGetRecordResultDetail>` | yes |
| `Errors` | `List<BatchGetRecordError>` | yes |
| `UnprocessedIdentifiers` | `List<BatchGetRecordIdentifier>` | yes |

## BatchWriteRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<BatchWriteRecordEntry>` | yes |
| `TtlDuration` | `TtlDuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchWriteRecordError>` | yes |
| `UnprocessedEntries` | `List<BatchWriteRecordEntry>` | yes |

## DeleteRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `RecordIdentifierValueAsString` | `string` | yes |
| `EventTime` | `string` | yes |
| `TargetStores` | `List<string>` | no |
| `DeletionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `RecordIdentifierValueAsString` | `string` | yes |
| `FeatureNames` | `List<string>` | no |
| `ExpirationTimeResponse` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Record` | `List<FeatureValue>` | no |
| `ExpiresAt` | `string` | no |

## ListRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeSoftDeletedRecords` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecordIdentifiers` | `List<string>` | yes |
| `NextToken` | `string` | no |

## PutRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `Record` | `List<FeatureValue>` | yes |
| `TargetStores` | `List<string>` | no |
| `TtlDuration` | `TtlDuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `RecordIdentifierValueAsString` | `string` | yes |
| `Features` | `List<FeatureValue>` | yes |
| `TargetStores` | `List<string>` | no |
| `TtlDuration` | `TtlDuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


