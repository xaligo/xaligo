# AWS Application Discovery Service

API version: 2015-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/discovery/2015-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateConfigurationItemsToApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationConfigurationId` | `string` | yes |
| `configurationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDeleteAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleteAgents` | `List<DeleteAgent>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteAgentError>` | no |

## BatchDeleteImportData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importTaskIds` | `List<string>` | yes |
| `deleteHistory` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteImportDataError>` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `wave` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationId` | `string` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationIds` | `List<string>` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationIds` | `List<string>` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentIds` | `List<string>` | no |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentsInfo` | `List<AgentInfo>` | no |
| `nextToken` | `string` | no |

## DescribeBatchDeleteConfigurationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `BatchDeleteConfigurationTask` | no |

## DescribeConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<Map<string>>` | no |

## DescribeContinuousExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportIds` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `descriptions` | `List<ContinuousExportDescription>` | no |
| `nextToken` | `string` | no |

## DescribeExportConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportIds` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportsInfo` | `List<ExportInfo>` | no |
| `nextToken` | `string` | no |

## DescribeExportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportIds` | `List<string>` | no |
| `filters` | `List<ExportFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportsInfo` | `List<ExportInfo>` | no |
| `nextToken` | `string` | no |

## DescribeImportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<ImportTaskFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tasks` | `List<ImportTask>` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<TagFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<ConfigurationTag>` | no |
| `nextToken` | `string` | no |

## DisassociateConfigurationItemsFromApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationConfigurationId` | `string` | yes |
| `configurationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |

## GetDiscoverySummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `servers` | `long` | no |
| `applications` | `long` | no |
| `serversMappedToApplications` | `long` | no |
| `serversMappedtoTags` | `long` | no |
| `agentSummary` | `CustomerAgentInfo` | no |
| `connectorSummary` | `CustomerConnectorInfo` | no |
| `meCollectorSummary` | `CustomerMeCollectorInfo` | no |
| `agentlessCollectorSummary` | `CustomerAgentlessCollectorInfo` | no |

## ListConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationType` | `string` | yes |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `orderBy` | `List<OrderByElement>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<Map<string>>` | no |
| `nextToken` | `string` | no |

## ListServerNeighbors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationId` | `string` | yes |
| `portInformationNeeded` | `boolean` | no |
| `neighborConfigurationIds` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `neighbors` | `List<NeighborConnectionDetail>` | yes |
| `nextToken` | `string` | no |
| `knownDependencyCount` | `long` | no |

## StartBatchDeleteConfigurationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationType` | `string` | yes |
| `configurationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskId` | `string` | no |

## StartContinuousExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |
| `s3Bucket` | `string` | no |
| `startTime` | `timestamp` | no |
| `dataSource` | `string` | no |
| `schemaStorageConfig` | `Map<string>` | no |

## StartDataCollectionByAgentIds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentsConfigurationStatus` | `List<AgentConfigurationStatus>` | no |

## StartExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportDataFormat` | `List<string>` | no |
| `filters` | `List<ExportFilter>` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `preferences` | `ExportPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | no |

## StartImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientRequestToken` | `string` | no |
| `name` | `string` | yes |
| `importUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `ImportTask` | no |

## StopContinuousExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | no |
| `stopTime` | `timestamp` | no |

## StopDataCollectionByAgentIds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentsConfigurationStatus` | `List<AgentConfigurationStatus>` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `wave` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


