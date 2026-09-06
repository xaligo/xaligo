# AWS CloudTrail

API version: 2013-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudtrail/2013-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagsList` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | no |
| `QueryId` | `string` | yes |
| `EventDataStoreOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | yes |
| `QueryStatus` | `string` | yes |
| `EventDataStoreOwnerAccountId` | `string` | no |

## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Source` | `string` | yes |
| `Destinations` | `List<Destination>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Name` | `string` | no |
| `Source` | `string` | no |
| `Destinations` | `List<Destination>` | no |
| `Tags` | `List<Tag>` | no |

## CreateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RefreshSchedule` | `RefreshSchedule` | no |
| `TagsList` | `List<Tag>` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `Widgets` | `List<RequestWidget>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardArn` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Widgets` | `List<Widget>` | no |
| `TagsList` | `List<Tag>` | no |
| `RefreshSchedule` | `RefreshSchedule` | no |
| `TerminationProtectionEnabled` | `boolean` | no |

## CreateEventDataStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `TagsList` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `StartIngestion` | `boolean` | no |
| `BillingMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `TagsList` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `KmsKeyId` | `string` | no |
| `BillingMode` | `string` | no |

## CreateTrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `S3KeyPrefix` | `string` | no |
| `SnsTopicName` | `string` | no |
| `IncludeGlobalServiceEvents` | `boolean` | no |
| `IsMultiRegionTrail` | `boolean` | no |
| `EnableLogFileValidation` | `boolean` | no |
| `CloudWatchLogsLogGroupArn` | `string` | no |
| `CloudWatchLogsRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `IsOrganizationTrail` | `boolean` | no |
| `TagsList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `S3BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `SnsTopicName` | `string` | no |
| `SnsTopicARN` | `string` | no |
| `IncludeGlobalServiceEvents` | `boolean` | no |
| `IsMultiRegionTrail` | `boolean` | no |
| `TrailARN` | `string` | no |
| `LogFileValidationEnabled` | `boolean` | no |
| `CloudWatchLogsLogGroupArn` | `string` | no |
| `CloudWatchLogsRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `IsOrganizationTrail` | `boolean` | no |

## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventDataStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterOrganizationDelegatedAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DelegatedAdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | no |
| `QueryId` | `string` | no |
| `QueryAlias` | `string` | no |
| `RefreshId` | `string` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | no |
| `QueryString` | `string` | no |
| `QueryStatus` | `string` | no |
| `QueryStatistics` | `QueryStatisticsForDescribeQuery` | no |
| `ErrorMessage` | `string` | no |
| `DeliveryS3Uri` | `string` | no |
| `DeliveryStatus` | `string` | no |
| `Prompt` | `string` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

## DescribeTrails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trailNameList` | `List<string>` | no |
| `includeShadowTrails` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trailList` | `List<Trail>` | no |

## DisableFederation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `FederationStatus` | `string` | no |

## EnableFederation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |
| `FederationRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `FederationStatus` | `string` | no |
| `FederationRoleArn` | `string` | no |

## GenerateQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStores` | `List<string>` | yes |
| `Prompt` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryStatement` | `string` | no |
| `QueryAlias` | `string` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

## GetChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Name` | `string` | no |
| `Source` | `string` | no |
| `SourceConfig` | `SourceConfig` | no |
| `Destinations` | `List<Destination>` | no |
| `IngestionStatus` | `IngestionStatus` | no |

## GetDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardArn` | `string` | no |
| `Type` | `string` | no |
| `Status` | `string` | no |
| `Widgets` | `List<Widget>` | no |
| `RefreshSchedule` | `RefreshSchedule` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `LastRefreshId` | `string` | no |
| `LastRefreshFailureReason` | `string` | no |
| `TerminationProtectionEnabled` | `boolean` | no |

## GetEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | no |
| `EventDataStore` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `EventDataStoreArn` | `string` | no |
| `MaxEventSize` | `string` | no |
| `ContextKeySelectors` | `List<ContextKeySelector>` | no |
| `AggregationConfigurations` | `List<AggregationConfiguration>` | no |

## GetEventDataStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `KmsKeyId` | `string` | no |
| `BillingMode` | `string` | no |
| `FederationStatus` | `string` | no |
| `FederationRoleArn` | `string` | no |
| `PartitionKeys` | `List<PartitionKey>` | no |

## GetEventSelectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `EventSelectors` | `List<EventSelector>` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |

## GetImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | no |
| `Destinations` | `List<string>` | no |
| `ImportSource` | `ImportSource` | no |
| `StartEventTime` | `timestamp` | no |
| `EndEventTime` | `timestamp` | no |
| `ImportStatus` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `ImportStatistics` | `ImportStatistics` | no |

## GetInsightSelectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | no |
| `EventDataStore` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `InsightSelectors` | `List<InsightSelector>` | no |
| `EventDataStoreArn` | `string` | no |
| `InsightsDestination` | `string` | no |

## GetQueryResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | no |
| `QueryId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxQueryResults` | `integer` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryStatus` | `string` | no |
| `QueryStatistics` | `QueryStatistics` | no |
| `QueryResultRows` | `List<List<Map<string>>>` | no |
| `NextToken` | `string` | no |
| `ErrorMessage` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourcePolicy` | `string` | no |
| `DelegatedAdminResourcePolicy` | `string` | no |

## GetTrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Trail` | `Trail` | no |

## GetTrailStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IsLogging` | `boolean` | no |
| `LatestDeliveryError` | `string` | no |
| `LatestNotificationError` | `string` | no |
| `LatestDeliveryTime` | `timestamp` | no |
| `LatestNotificationTime` | `timestamp` | no |
| `StartLoggingTime` | `timestamp` | no |
| `StopLoggingTime` | `timestamp` | no |
| `LatestCloudWatchLogsDeliveryError` | `string` | no |
| `LatestCloudWatchLogsDeliveryTime` | `timestamp` | no |
| `LatestDigestDeliveryTime` | `timestamp` | no |
| `LatestDigestDeliveryError` | `string` | no |
| `LatestDeliveryAttemptTime` | `string` | no |
| `LatestNotificationAttemptTime` | `string` | no |
| `LatestNotificationAttemptSucceeded` | `string` | no |
| `LatestDeliveryAttemptSucceeded` | `string` | no |
| `TimeLoggingStarted` | `string` | no |
| `TimeLoggingStopped` | `string` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<Channel>` | no |
| `NextToken` | `string` | no |

## ListDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Dashboards` | `List<DashboardDetail>` | no |
| `NextToken` | `string` | no |

## ListEventDataStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStores` | `List<EventDataStore>` | no |
| `NextToken` | `string` | no |

## ListImportFailures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<ImportFailureListItem>` | no |
| `NextToken` | `string` | no |

## ListImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `Destination` | `string` | no |
| `ImportStatus` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Imports` | `List<ImportsListItem>` | no |
| `NextToken` | `string` | no |

## ListInsightsData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightSource` | `string` | yes |
| `DataType` | `string` | yes |
| `Dimensions` | `Map<string>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<Event>` | no |
| `NextToken` | `string` | no |

## ListInsightsMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | no |
| `EventSource` | `string` | yes |
| `EventName` | `string` | yes |
| `InsightType` | `string` | yes |
| `ErrorCode` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Period` | `integer` | no |
| `DataType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `EventSource` | `string` | no |
| `EventName` | `string` | no |
| `InsightType` | `string` | no |
| `ErrorCode` | `string` | no |
| `Timestamps` | `List<timestamp>` | no |
| `Values` | `List<double>` | no |
| `NextToken` | `string` | no |

## ListPublicKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyList` | `List<PublicKey>` | no |
| `NextToken` | `string` | no |

## ListQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `QueryStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queries` | `List<Query>` | no |
| `NextToken` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdList` | `List<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTagList` | `List<ResourceTag>` | no |
| `NextToken` | `string` | no |

## ListTrails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Trails` | `List<TrailInfo>` | no |
| `NextToken` | `string` | no |

## LookupEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LookupAttributes` | `List<LookupAttribute>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `EventCategory` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<Event>` | no |
| `NextToken` | `string` | no |

## PutEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | no |
| `EventDataStore` | `string` | no |
| `MaxEventSize` | `string` | no |
| `ContextKeySelectors` | `List<ContextKeySelector>` | no |
| `AggregationConfigurations` | `List<AggregationConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `EventDataStoreArn` | `string` | no |
| `MaxEventSize` | `string` | no |
| `ContextKeySelectors` | `List<ContextKeySelector>` | no |
| `AggregationConfigurations` | `List<AggregationConfiguration>` | no |

## PutEventSelectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | yes |
| `EventSelectors` | `List<EventSelector>` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `EventSelectors` | `List<EventSelector>` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |

## PutInsightSelectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailName` | `string` | no |
| `InsightSelectors` | `List<InsightSelector>` | yes |
| `EventDataStore` | `string` | no |
| `InsightsDestination` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrailARN` | `string` | no |
| `InsightSelectors` | `List<InsightSelector>` | no |
| `EventDataStoreArn` | `string` | no |
| `InsightsDestination` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ResourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourcePolicy` | `string` | no |
| `DelegatedAdminResourcePolicy` | `string` | no |

## RegisterOrganizationDelegatedAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagsList` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestoreEventDataStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `KmsKeyId` | `string` | no |
| `BillingMode` | `string` | no |

## SearchSampleQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchPhrase` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchResults` | `List<SearchSampleQueriesSearchResult>` | no |
| `NextToken` | `string` | no |

## StartDashboardRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | yes |
| `QueryParameterValues` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshId` | `string` | no |

## StartEventDataStoreIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destinations` | `List<string>` | no |
| `ImportSource` | `ImportSource` | no |
| `StartEventTime` | `timestamp` | no |
| `EndEventTime` | `timestamp` | no |
| `ImportId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | no |
| `Destinations` | `List<string>` | no |
| `ImportSource` | `ImportSource` | no |
| `StartEventTime` | `timestamp` | no |
| `EndEventTime` | `timestamp` | no |
| `ImportStatus` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |

## StartLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryStatement` | `string` | no |
| `DeliveryS3Uri` | `string` | no |
| `QueryAlias` | `string` | no |
| `QueryParameters` | `List<string>` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryId` | `string` | no |
| `EventDataStoreOwnerAccountId` | `string` | no |

## StopEventDataStoreIngestion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportId` | `string` | no |
| `ImportSource` | `ImportSource` | no |
| `Destinations` | `List<string>` | no |
| `ImportStatus` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `StartEventTime` | `timestamp` | no |
| `EndEventTime` | `timestamp` | no |
| `ImportStatistics` | `ImportStatistics` | no |

## StopLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `string` | yes |
| `Destinations` | `List<Destination>` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelArn` | `string` | no |
| `Name` | `string` | no |
| `Source` | `string` | no |
| `Destinations` | `List<Destination>` | no |

## UpdateDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardId` | `string` | yes |
| `Widgets` | `List<RequestWidget>` | no |
| `RefreshSchedule` | `RefreshSchedule` | no |
| `TerminationProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardArn` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Widgets` | `List<Widget>` | no |
| `RefreshSchedule` | `RefreshSchedule` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |

## UpdateEventDataStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStore` | `string` | yes |
| `Name` | `string` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `BillingMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDataStoreArn` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `AdvancedEventSelectors` | `List<AdvancedEventSelector>` | no |
| `MultiRegionEnabled` | `boolean` | no |
| `OrganizationEnabled` | `boolean` | no |
| `RetentionPeriod` | `integer` | no |
| `TerminationProtectionEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `UpdatedTimestamp` | `timestamp` | no |
| `KmsKeyId` | `string` | no |
| `BillingMode` | `string` | no |
| `FederationStatus` | `string` | no |
| `FederationRoleArn` | `string` | no |

## UpdateTrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `S3BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `SnsTopicName` | `string` | no |
| `IncludeGlobalServiceEvents` | `boolean` | no |
| `IsMultiRegionTrail` | `boolean` | no |
| `EnableLogFileValidation` | `boolean` | no |
| `CloudWatchLogsLogGroupArn` | `string` | no |
| `CloudWatchLogsRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `IsOrganizationTrail` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `S3BucketName` | `string` | no |
| `S3KeyPrefix` | `string` | no |
| `SnsTopicName` | `string` | no |
| `SnsTopicARN` | `string` | no |
| `IncludeGlobalServiceEvents` | `boolean` | no |
| `IsMultiRegionTrail` | `boolean` | no |
| `TrailARN` | `string` | no |
| `LogFileValidationEnabled` | `boolean` | no |
| `CloudWatchLogsLogGroupArn` | `string` | no |
| `CloudWatchLogsRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `IsOrganizationTrail` | `boolean` | no |

