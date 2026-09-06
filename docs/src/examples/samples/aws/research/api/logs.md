# Amazon CloudWatch Logs

API version: 2014-03-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/logs/2014-03-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateKmsKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `kmsKeyId` | `string` | yes |
| `resourceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSourceToS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `dataSource` | `DataSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | no |

## CancelExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `importStatistics` | `ImportStatistics` | no |
| `importStatus` | `string` | no |
| `creationTime` | `long` | no |
| `lastUpdatedTime` | `long` | no |

## CreateDelivery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliverySourceName` | `string` | yes |
| `deliveryDestinationArn` | `string` | yes |
| `recordFields` | `List<string>` | no |
| `fieldDelimiter` | `string` | no |
| `s3DeliveryConfiguration` | `S3DeliveryConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delivery` | `Delivery` | no |

## CreateExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskName` | `string` | no |
| `logGroupName` | `string` | yes |
| `logStreamNamePrefix` | `string` | no |
| `from` | `long` | yes |
| `to` | `long` | yes |
| `destination` | `string` | yes |
| `destinationPrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## CreateImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importSourceArn` | `string` | yes |
| `importRoleArn` | `string` | yes |
| `importFilter` | `ImportFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `importDestinationArn` | `string` | no |
| `creationTime` | `long` | no |

## CreateLogAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupArnList` | `List<string>` | yes |
| `detectorName` | `string` | no |
| `evaluationFrequency` | `string` | no |
| `filterPattern` | `string` | no |
| `kmsKeyId` | `string` | no |
| `anomalyVisibilityTime` | `long` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorArn` | `string` | no |

## CreateLogGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `logGroupClass` | `string` | no |
| `deletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLogStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `logStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLookupTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableName` | `string` | yes |
| `description` | `string` | no |
| `tableBody` | `string` | no |
| `queryId` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | no |
| `createdAt` | `long` | no |

## CreateScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `queryLanguage` | `string` | yes |
| `queryString` | `string` | yes |
| `logGroupIdentifiers` | `List<string>` | no |
| `scheduleExpression` | `string` | yes |
| `timezone` | `string` | no |
| `startTimeOffset` | `long` | no |
| `endTimeOffset` | `long` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `scheduleStartTime` | `long` | no |
| `scheduleEndTime` | `long` | no |
| `executionRoleArn` | `string` | yes |
| `state` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledQueryArn` | `string` | no |
| `state` | `string` | no |

## DeleteAccountPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataProtectionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDelivery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeliveryDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeliveryDestinationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeliverySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIndexPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationName` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLogAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLogGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLogStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `logStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLookupTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMetricFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `filterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `success` | `boolean` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | no |
| `resourceArn` | `string` | no |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRetentionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriptionFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `filterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSyslogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `vpcEndpointId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyType` | `string` | yes |
| `policyName` | `string` | no |
| `accountIdentifiers` | `List<string>` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountPolicies` | `List<AccountPolicy>` | no |
| `nextToken` | `string` | no |

## DescribeConfigurationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `string` | no |
| `logTypes` | `List<string>` | no |
| `resourceTypes` | `List<string>` | no |
| `deliveryDestinationTypes` | `List<string>` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationTemplates` | `List<ConfigurationTemplate>` | no |
| `nextToken` | `string` | no |

## DescribeDeliveries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveries` | `List<Delivery>` | no |
| `nextToken` | `string` | no |

## DescribeDeliveryDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestinations` | `List<DeliveryDestination>` | no |
| `nextToken` | `string` | no |

## DescribeDeliverySources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliverySources` | `List<DeliverySource>` | no |
| `nextToken` | `string` | no |

## DescribeDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationNamePrefix` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinations` | `List<Destination>` | no |
| `nextToken` | `string` | no |

## DescribeExportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |
| `statusCode` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportTasks` | `List<ExportTask>` | no |
| `nextToken` | `string` | no |

## DescribeFieldIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifiers` | `List<string>` | yes |
| `indexCategories` | `List<string>` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fieldIndexes` | `List<FieldIndex>` | no |
| `nextToken` | `string` | no |

## DescribeImportTaskBatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | yes |
| `batchImportStatus` | `List<string>` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importSourceArn` | `string` | no |
| `importId` | `string` | no |
| `importBatches` | `List<ImportBatch>` | no |
| `nextToken` | `string` | no |

## DescribeImportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importId` | `string` | no |
| `importStatus` | `string` | no |
| `importSourceArn` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imports` | `List<Import>` | no |
| `nextToken` | `string` | no |

## DescribeIndexPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifiers` | `List<string>` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexPolicies` | `List<IndexPolicy>` | no |
| `nextToken` | `string` | no |

## DescribeLogGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIdentifiers` | `List<string>` | no |
| `logGroupNamePrefix` | `string` | no |
| `logGroupNamePattern` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `includeLinkedAccounts` | `boolean` | no |
| `logGroupClass` | `string` | no |
| `logGroupIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroups` | `List<LogGroup>` | no |
| `nextToken` | `string` | no |

## DescribeLogStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `logGroupIdentifier` | `string` | no |
| `logStreamNamePrefix` | `string` | no |
| `orderBy` | `string` | no |
| `descending` | `boolean` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logStreams` | `List<LogStream>` | no |
| `nextToken` | `string` | no |

## DescribeLookupTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableNamePrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTables` | `List<LookupTable>` | no |
| `nextToken` | `string` | no |

## DescribeMetricFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `filterNamePrefix` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `metricName` | `string` | no |
| `metricNamespace` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `metricFilters` | `List<MetricFilter>` | no |
| `nextToken` | `string` | no |

## DescribeQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `queryLanguage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queries` | `List<QueryInfo>` | no |
| `nextToken` | `string` | no |

## DescribeQueryDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryLanguage` | `string` | no |
| `queryDefinitionNamePrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryDefinitions` | `List<QueryDefinition>` | no |
| `nextToken` | `string` | no |

## DescribeResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `resourceArn` | `string` | no |
| `policyScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicies` | `List<ResourcePolicy>` | no |
| `nextToken` | `string` | no |

## DescribeSubscriptionFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `filterNamePrefix` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionFilters` | `List<SubscriptionFilter>` | no |
| `nextToken` | `string` | no |

## DisassociateKmsKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `resourceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSourceFromS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | no |

## FilterLogEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `logGroupIdentifier` | `string` | no |
| `logStreamNames` | `List<string>` | no |
| `logStreamNamePrefix` | `string` | no |
| `startTime` | `long` | no |
| `endTime` | `long` | no |
| `filterPattern` | `string` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `startFromHead` | `boolean` | no |
| `interleaved` | `boolean` | no |
| `unmask` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<FilteredLogEvent>` | no |
| `searchedLogStreams` | `List<SearchedLogStream>` | no |
| `nextToken` | `string` | no |

## GetDataProtectionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | no |
| `policyDocument` | `string` | no |
| `lastUpdatedTime` | `long` | no |

## GetDelivery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `delivery` | `Delivery` | no |

## GetDeliveryDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestination` | `DeliveryDestination` | no |

## GetDeliveryDestinationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | no |

## GetDeliverySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliverySource` | `DeliverySource` | no |

## GetIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationName` | `string` | no |
| `integrationType` | `string` | no |
| `integrationStatus` | `string` | no |
| `integrationDetails` | `IntegrationDetails` | no |

## GetLogAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `detectorName` | `string` | no |
| `logGroupArnList` | `List<string>` | no |
| `evaluationFrequency` | `string` | no |
| `filterPattern` | `string` | no |
| `anomalyDetectorStatus` | `string` | no |
| `kmsKeyId` | `string` | no |
| `creationTimeStamp` | `long` | no |
| `lastModifiedTimeStamp` | `long` | no |
| `anomalyVisibilityTime` | `long` | no |

## GetLogEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `logGroupIdentifier` | `string` | no |
| `logStreamName` | `string` | yes |
| `startTime` | `long` | no |
| `endTime` | `long` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `startFromHead` | `boolean` | no |
| `unmask` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<OutputLogEvent>` | no |
| `nextForwardToken` | `string` | no |
| `nextBackwardToken` | `string` | no |

## GetLogFields

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSourceName` | `string` | yes |
| `dataSourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logFields` | `List<LogFieldsListItem>` | no |

## GetLogGroupFields

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | no |
| `time` | `long` | no |
| `logGroupIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupFields` | `List<LogGroupField>` | no |

## GetLogObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `unmask` | `boolean` | no |
| `logObjectPointer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fieldStream` | `GetLogObjectResponseStream` | no |

## GetLogRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logRecordPointer` | `string` | yes |
| `unmask` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logRecord` | `Map<string>` | no |

## GetLookupTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | no |
| `lookupTableName` | `string` | no |
| `description` | `string` | no |
| `tableBody` | `string` | no |
| `sizeBytes` | `long` | no |
| `lastUpdatedTime` | `long` | no |
| `kmsKeyId` | `string` | no |

## GetQueryResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryLanguage` | `string` | no |
| `results` | `List<List<ResultField>>` | no |
| `statistics` | `QueryStatistics` | no |
| `status` | `string` | no |
| `encryptionKey` | `string` | no |
| `nextToken` | `string` | no |

## GetScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledQueryArn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `queryLanguage` | `string` | no |
| `queryString` | `string` | no |
| `logGroupIdentifiers` | `List<string>` | no |
| `scheduleExpression` | `string` | no |
| `timezone` | `string` | no |
| `startTimeOffset` | `long` | no |
| `endTimeOffset` | `long` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `state` | `string` | no |
| `scheduleType` | `string` | no |
| `lastTriggeredTime` | `long` | no |
| `lastExecutionStatus` | `string` | no |
| `scheduleStartTime` | `long` | no |
| `scheduleEndTime` | `long` | no |
| `executionRoleArn` | `string` | no |
| `creationTime` | `long` | no |
| `lastUpdatedTime` | `long` | no |

## GetScheduledQueryHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `startTime` | `long` | yes |
| `endTime` | `long` | yes |
| `executionStatuses` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `scheduledQueryArn` | `string` | no |
| `triggerHistory` | `List<TriggerHistoryRecord>` | no |
| `nextToken` | `string` | no |

## GetStorageTierPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageTier` | `string` | no |
| `lastUpdatedTime` | `long` | no |

## GetTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | no |
| `creationTime` | `long` | no |
| `lastModifiedTime` | `long` | no |
| `transformerConfig` | `List<Processor>` | no |

## ListAggregateLogGroupSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIdentifiers` | `List<string>` | no |
| `includeLinkedAccounts` | `boolean` | no |
| `logGroupClass` | `string` | no |
| `logGroupNamePattern` | `string` | no |
| `dataSources` | `List<DataSourceFilter>` | no |
| `groupBy` | `string` | yes |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregateLogGroupSummaries` | `List<AggregateLogGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListAnomalies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorArn` | `string` | no |
| `suppressionState` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalies` | `List<Anomaly>` | no |
| `nextToken` | `string` | no |

## ListIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationNamePrefix` | `string` | no |
| `integrationType` | `string` | no |
| `integrationStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationSummaries` | `List<IntegrationSummary>` | no |

## ListLogAnomalyDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterLogGroupArn` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectors` | `List<AnomalyDetector>` | no |
| `nextToken` | `string` | no |

## ListLogGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupNamePattern` | `string` | no |
| `logGroupClass` | `string` | no |
| `includeLinkedAccounts` | `boolean` | no |
| `accountIdentifiers` | `List<string>` | no |
| `nextToken` | `string` | no |
| `limit` | `integer` | no |
| `dataSources` | `List<DataSourceFilter>` | no |
| `fieldIndexNames` | `List<string>` | no |
| `logGroupTags` | `List<TagFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroups` | `List<LogGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListLogGroupsForQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifiers` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListScheduledQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `state` | `string` | no |
| `scheduleType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `scheduledQueries` | `List<ScheduledQuerySummary>` | no |

## ListSourcesForS3TableIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<S3TableIntegrationSource>` | no |
| `nextToken` | `string` | no |

## ListSyslogConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | no |
| `vpcEndpointId` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `syslogConfigurations` | `List<SyslogConfiguration>` | no |
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

## ListTagsLogGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutAccountPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | yes |
| `policyDocument` | `string` | yes |
| `policyType` | `string` | yes |
| `scope` | `string` | no |
| `selectionCriteria` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountPolicy` | `AccountPolicy` | no |

## PutBearerTokenAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `bearerTokenAuthenticationEnabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDataProtectionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `policyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | no |
| `policyDocument` | `string` | no |
| `lastUpdatedTime` | `long` | no |

## PutDeliveryDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `outputFormat` | `string` | no |
| `deliveryDestinationConfiguration` | `DeliveryDestinationConfiguration` | no |
| `deliveryDestinationType` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestination` | `DeliveryDestination` | no |

## PutDeliveryDestinationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryDestinationName` | `string` | yes |
| `deliveryDestinationPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `Policy` | no |

## PutDeliverySource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `resourceArn` | `string` | yes |
| `logType` | `string` | yes |
| `tags` | `Map<string>` | no |
| `deliverySourceConfiguration` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliverySource` | `DeliverySource` | no |

## PutDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationName` | `string` | yes |
| `targetArn` | `string` | yes |
| `roleArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destination` | `Destination` | no |

## PutDestinationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationName` | `string` | yes |
| `accessPolicy` | `string` | yes |
| `forceUpdate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutIndexPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `policyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexPolicy` | `IndexPolicy` | no |

## PutIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationName` | `string` | yes |
| `resourceConfig` | `ResourceConfig` | yes |
| `integrationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationName` | `string` | no |
| `integrationStatus` | `string` | no |

## PutLogEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `logStreamName` | `string` | yes |
| `logEvents` | `List<InputLogEvent>` | yes |
| `sequenceToken` | `string` | no |
| `entity` | `Entity` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextSequenceToken` | `string` | no |
| `rejectedLogEventsInfo` | `RejectedLogEventsInfo` | no |
| `rejectedEntityInfo` | `RejectedEntityInfo` | no |

## PutLogGroupDeletionProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `deletionProtectionEnabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMetricFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `filterName` | `string` | yes |
| `filterPattern` | `string` | yes |
| `metricTransformations` | `List<MetricTransformation>` | yes |
| `applyOnTransformedLogs` | `boolean` | no |
| `fieldSelectionCriteria` | `string` | no |
| `emitSystemFieldDimensions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutQueryDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryLanguage` | `string` | no |
| `name` | `string` | yes |
| `queryDefinitionId` | `string` | no |
| `logGroupNames` | `List<string>` | no |
| `queryString` | `string` | yes |
| `clientToken` | `string` | no |
| `parameters` | `List<QueryParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryDefinitionId` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyName` | `string` | no |
| `policyDocument` | `string` | no |
| `resourceArn` | `string` | no |
| `expectedRevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `ResourcePolicy` | no |
| `revisionId` | `string` | no |

## PutRetentionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `retentionInDays` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutStorageTierPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageTier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `storageTier` | `string` | no |
| `lastUpdatedTime` | `long` | no |

## PutSubscriptionFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `filterName` | `string` | yes |
| `filterPattern` | `string` | yes |
| `destinationArn` | `string` | yes |
| `roleArn` | `string` | no |
| `distribution` | `string` | no |
| `applyOnTransformedLogs` | `boolean` | no |
| `fieldSelectionCriteria` | `string` | no |
| `emitSystemFields` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutSyslogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `vpcEndpointId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifier` | `string` | yes |
| `transformerConfig` | `List<Processor>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartLiveTail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupIdentifiers` | `List<string>` | yes |
| `logStreamNames` | `List<string>` | no |
| `logStreamNamePrefixes` | `List<string>` | no |
| `logEventFilterPattern` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseStream` | `StartLiveTailResponseStream` | no |

## StartQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryLanguage` | `string` | no |
| `logGroupName` | `string` | no |
| `logGroupNames` | `List<string>` | no |
| `logGroupIdentifiers` | `List<string>` | no |
| `startTime` | `long` | yes |
| `endTime` | `long` | yes |
| `queryString` | `string` | yes |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | no |

## StopQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `success` | `boolean` | no |

## TagLogGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestMetricFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterPattern` | `string` | yes |
| `logEventMessages` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `matches` | `List<MetricFilterMatchRecord>` | no |

## TestTransformer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformerConfig` | `List<Processor>` | yes |
| `logEventMessages` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `transformedLogs` | `List<TransformedLogRecord>` | no |

## UntagLogGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logGroupName` | `string` | yes |
| `tags` | `List<string>` | yes |

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


## UpdateAnomaly

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyId` | `string` | no |
| `patternId` | `string` | no |
| `anomalyDetectorArn` | `string` | yes |
| `suppressionType` | `string` | no |
| `suppressionPeriod` | `SuppressionPeriod` | no |
| `baseline` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDeliveryConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `recordFields` | `List<string>` | no |
| `fieldDelimiter` | `string` | no |
| `s3DeliveryConfiguration` | `S3DeliveryConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLogAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `anomalyDetectorArn` | `string` | yes |
| `evaluationFrequency` | `string` | no |
| `filterPattern` | `string` | no |
| `anomalyVisibilityTime` | `long` | no |
| `enabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLookupTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | yes |
| `description` | `string` | no |
| `tableBody` | `string` | no |
| `queryId` | `string` | no |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lookupTableArn` | `string` | no |
| `lastUpdatedTime` | `long` | no |

## UpdateScheduledQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `queryLanguage` | `string` | yes |
| `queryString` | `string` | yes |
| `logGroupIdentifiers` | `List<string>` | no |
| `scheduleExpression` | `string` | yes |
| `timezone` | `string` | no |
| `startTimeOffset` | `long` | no |
| `endTimeOffset` | `long` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `scheduleStartTime` | `long` | no |
| `scheduleEndTime` | `long` | no |
| `executionRoleArn` | `string` | yes |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scheduledQueryArn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `queryLanguage` | `string` | no |
| `queryString` | `string` | no |
| `logGroupIdentifiers` | `List<string>` | no |
| `scheduleExpression` | `string` | no |
| `timezone` | `string` | no |
| `startTimeOffset` | `long` | no |
| `endTimeOffset` | `long` | no |
| `destinationConfiguration` | `DestinationConfiguration` | no |
| `state` | `string` | no |
| `scheduleType` | `string` | no |
| `lastTriggeredTime` | `long` | no |
| `lastExecutionStatus` | `string` | no |
| `scheduleStartTime` | `long` | no |
| `scheduleEndTime` | `long` | no |
| `executionRoleArn` | `string` | no |
| `creationTime` | `long` | no |
| `lastUpdatedTime` | `long` | no |

