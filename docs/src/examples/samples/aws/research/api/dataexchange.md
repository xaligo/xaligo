# AWS Data Exchange

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/dataexchange/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptDataGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SenderPrincipal` | `string` | no |
| `ReceiverPrincipal` | `string` | yes |
| `Description` | `string` | no |
| `AcceptanceState` | `string` | yes |
| `AcceptedAt` | `timestamp` | no |
| `EndsAt` | `timestamp` | no |
| `GrantDistributionScope` | `string` | yes |
| `DataSetId` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |

## CancelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDataGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `GrantDistributionScope` | `string` | yes |
| `ReceiverPrincipal` | `string` | yes |
| `SourceDataSetId` | `string` | yes |
| `EndsAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SenderPrincipal` | `string` | yes |
| `ReceiverPrincipal` | `string` | yes |
| `Description` | `string` | no |
| `AcceptanceState` | `string` | yes |
| `AcceptedAt` | `timestamp` | no |
| `EndsAt` | `timestamp` | no |
| `GrantDistributionScope` | `string` | yes |
| `DataSetId` | `string` | yes |
| `SourceDataSetId` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## CreateDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetType` | `string` | yes |
| `Description` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetType` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Origin` | `string` | no |
| `OriginDetails` | `OriginDetails` | no |
| `SourceId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |

## CreateEventAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `Action` | yes |
| `Event` | `Event` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `Action` | no |
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Event` | `Event` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetConfiguration` | `AssetConfiguration` | no |
| `Details` | `RequestDetails` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetConfiguration` | `AssetConfiguration` | no |
| `CreatedAt` | `timestamp` | no |
| `Details` | `ResponseDetails` | no |
| `Errors` | `List<JobError>` | no |
| `Id` | `string` | no |
| `State` | `string` | no |
| `Type` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## CreateRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Comment` | `string` | no |
| `DataSetId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Comment` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Finalized` | `boolean` | no |
| `Id` | `string` | no |
| `SourceId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |
| `RevocationComment` | `string` | no |
| `Revoked` | `boolean` | no |
| `RevokedAt` | `timestamp` | no |

## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetDetails` | `AssetDetails` | no |
| `AssetType` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `RevisionId` | `string` | no |
| `SourceId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |

## GetDataGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SenderPrincipal` | `string` | yes |
| `ReceiverPrincipal` | `string` | yes |
| `Description` | `string` | no |
| `AcceptanceState` | `string` | yes |
| `AcceptedAt` | `timestamp` | no |
| `EndsAt` | `timestamp` | no |
| `GrantDistributionScope` | `string` | yes |
| `DataSetId` | `string` | yes |
| `SourceDataSetId` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## GetDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetType` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Origin` | `string` | no |
| `OriginDetails` | `OriginDetails` | no |
| `SourceId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |

## GetEventAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `Action` | no |
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Event` | `Event` | no |
| `Id` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetConfiguration` | `AssetConfiguration` | no |
| `CreatedAt` | `timestamp` | no |
| `Details` | `ResponseDetails` | no |
| `Errors` | `List<JobError>` | no |
| `Id` | `string` | no |
| `State` | `string` | no |
| `Type` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## GetReceivedDataGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SenderPrincipal` | `string` | no |
| `ReceiverPrincipal` | `string` | yes |
| `Description` | `string` | no |
| `AcceptanceState` | `string` | yes |
| `AcceptedAt` | `timestamp` | no |
| `EndsAt` | `timestamp` | no |
| `GrantDistributionScope` | `string` | yes |
| `DataSetId` | `string` | yes |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `UpdatedAt` | `timestamp` | yes |

## GetRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Comment` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Finalized` | `boolean` | no |
| `Id` | `string` | no |
| `SourceId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UpdatedAt` | `timestamp` | no |
| `RevocationComment` | `string` | no |
| `Revoked` | `boolean` | no |
| `RevokedAt` | `timestamp` | no |

## ListDataGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantSummaries` | `List<DataGrantSummaryEntry>` | no |
| `NextToken` | `string` | no |

## ListDataSetRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Revisions` | `List<RevisionEntry>` | no |

## ListDataSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Origin` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSets` | `List<DataSetEntry>` | no |
| `NextToken` | `string` | no |

## ListEventActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSourceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventActions` | `List<EventActionEntry>` | no |
| `NextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Jobs` | `List<JobEntry>` | no |
| `NextToken` | `string` | no |

## ListReceivedDataGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AcceptanceState` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataGrantSummaries` | `List<ReceivedDataGrantSummariesEntry>` | no |
| `NextToken` | `string` | no |

## ListRevisionAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Assets` | `List<AssetEntry>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## RevokeRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `RevisionId` | `string` | yes |
| `RevocationComment` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Comment` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Finalized` | `boolean` | no |
| `Id` | `string` | no |
| `SourceId` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `RevocationComment` | `string` | no |
| `Revoked` | `boolean` | no |
| `RevokedAt` | `timestamp` | no |

## SendApiAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `string` | no |
| `QueryStringParameters` | `Map<string>` | no |
| `AssetId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `RequestHeaders` | `Map<string>` | no |
| `Method` | `string` | no |
| `Path` | `string` | no |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `string` | no |
| `ResponseHeaders` | `Map<string>` | no |

## SendDataSetNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `ScopeDetails` | no |
| `ClientToken` | `string` | no |
| `Comment` | `string` | no |
| `DataSetId` | `string` | yes |
| `Details` | `NotificationDetails` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssetId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `Name` | `string` | yes |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetDetails` | `AssetDetails` | no |
| `AssetType` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `RevisionId` | `string` | no |
| `SourceId` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## UpdateDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | yes |
| `Description` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssetType` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Origin` | `string` | no |
| `OriginDetails` | `OriginDetails` | no |
| `SourceId` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## UpdateEventAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `Action` | no |
| `EventActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `Action` | no |
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Event` | `Event` | no |
| `Id` | `string` | no |
| `UpdatedAt` | `timestamp` | no |

## UpdateRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Comment` | `string` | no |
| `DataSetId` | `string` | yes |
| `Finalized` | `boolean` | no |
| `RevisionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Comment` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DataSetId` | `string` | no |
| `Finalized` | `boolean` | no |
| `Id` | `string` | no |
| `SourceId` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `RevocationComment` | `string` | no |
| `Revoked` | `boolean` | no |
| `RevokedAt` | `timestamp` | no |

