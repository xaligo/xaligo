# Amazon Cognito Sync

API version: 2014-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cognito-sync/2014-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BulkPublish

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |

## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dataset` | `Dataset` | no |

## DescribeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dataset` | `Dataset` | no |

## DescribeIdentityPoolUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolUsage` | `IdentityPoolUsage` | no |

## DescribeIdentityUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityUsage` | `IdentityUsage` | no |

## GetBulkPublishDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `BulkPublishStartTime` | `timestamp` | no |
| `BulkPublishCompleteTime` | `timestamp` | no |
| `BulkPublishStatus` | `string` | no |
| `FailureMessage` | `string` | no |

## GetCognitoEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `Map<string>` | no |

## GetIdentityPoolConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `PushSync` | `PushSync` | no |
| `CognitoStreams` | `CognitoStreams` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Datasets` | `List<Dataset>` | no |
| `Count` | `integer` | no |
| `NextToken` | `string` | no |

## ListIdentityPoolUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolUsages` | `List<IdentityPoolUsage>` | no |
| `MaxResults` | `integer` | no |
| `Count` | `integer` | no |
| `NextToken` | `string` | no |

## ListRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |
| `LastSyncCount` | `long` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SyncSessionToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | no |
| `NextToken` | `string` | no |
| `Count` | `integer` | no |
| `DatasetSyncCount` | `long` | no |
| `LastModifiedBy` | `string` | no |
| `MergedDatasetNames` | `List<string>` | no |
| `DatasetExists` | `boolean` | no |
| `DatasetDeletedAfterRequestedSyncCount` | `boolean` | no |
| `SyncSessionToken` | `string` | no |

## RegisterDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `Platform` | `string` | yes |
| `Token` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceId` | `string` | no |

## SetCognitoEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `Events` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityPoolConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `PushSync` | `PushSync` | no |
| `CognitoStreams` | `CognitoStreams` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `PushSync` | `PushSync` | no |
| `CognitoStreams` | `CognitoStreams` | no |

## SubscribeToDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnsubscribeFromDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | yes |
| `DatasetName` | `string` | yes |
| `DeviceId` | `string` | no |
| `RecordPatches` | `List<RecordPatch>` | no |
| `SyncSessionToken` | `string` | yes |
| `ClientContext` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<Record>` | no |

