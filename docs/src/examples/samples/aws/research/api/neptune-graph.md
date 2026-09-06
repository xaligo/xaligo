# Amazon Neptune Graph

API version: 2023-11-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/neptune-graph/2023-11-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | yes |
| `roleArn` | `string` | yes |
| `taskId` | `string` | yes |
| `status` | `string` | yes |
| `format` | `string` | yes |
| `destination` | `string` | yes |
| `kmsKeyIdentifier` | `string` | yes |
| `parquetType` | `string` | no |
| `statusReason` | `string` | no |

## CancelImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | no |
| `taskId` | `string` | yes |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |

## CancelQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphName` | `string` | yes |
| `tags` | `Map<string>` | no |
| `publicConnectivity` | `boolean` | no |
| `kmsKeyIdentifier` | `string` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `deletionProtection` | `boolean` | no |
| `provisionedMemory` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## CreateGraphSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `snapshotName` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `sourceGraphId` | `string` | no |
| `snapshotCreateTime` | `timestamp` | no |
| `status` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |

## CreateGraphUsingImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphName` | `string` | yes |
| `tags` | `Map<string>` | no |
| `publicConnectivity` | `boolean` | no |
| `kmsKeyIdentifier` | `string` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `deletionProtection` | `boolean` | no |
| `importOptions` | `ImportOptions` | no |
| `maxProvisionedMemory` | `integer` | no |
| `minProvisionedMemory` | `integer` | no |
| `failOnError` | `boolean` | no |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `blankNodeHandling` | `string` | no |
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | no |
| `taskId` | `string` | yes |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `importOptions` | `ImportOptions` | no |

## CreatePrivateGraphEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `vpcId` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `vpcSecurityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `status` | `string` | yes |
| `vpcEndpointId` | `string` | no |

## DeleteGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `skipSnapshot` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## DeleteGraphSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `sourceGraphId` | `string` | no |
| `snapshotCreateTime` | `timestamp` | no |
| `status` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |

## DeletePrivateGraphEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `vpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `status` | `string` | yes |
| `vpcEndpointId` | `string` | no |

## ExecuteQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `queryString` | `string` | yes |
| `language` | `string` | yes |
| `parameters` | `Map<Document>` | no |
| `planCache` | `string` | no |
| `explainMode` | `string` | no |
| `queryTimeoutMilliseconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `payload` | `blob` | yes |

## GetExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | yes |
| `roleArn` | `string` | yes |
| `taskId` | `string` | yes |
| `status` | `string` | yes |
| `format` | `string` | yes |
| `destination` | `string` | yes |
| `kmsKeyIdentifier` | `string` | yes |
| `parquetType` | `string` | no |
| `statusReason` | `string` | no |
| `exportTaskDetails` | `ExportTaskDetails` | no |
| `exportFilter` | `ExportFilter` | no |

## GetGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## GetGraphSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `sourceGraphId` | `string` | no |
| `snapshotCreateTime` | `timestamp` | no |
| `status` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |

## GetGraphSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | no |
| `lastStatisticsComputationTime` | `timestamp` | no |
| `graphSummary` | `GraphDataSummary` | no |

## GetImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | no |
| `taskId` | `string` | yes |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `importOptions` | `ImportOptions` | no |
| `importTaskDetails` | `ImportTaskDetails` | no |
| `attemptNumber` | `integer` | no |
| `statusReason` | `string` | no |

## GetPrivateGraphEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `vpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `status` | `string` | yes |
| `vpcEndpointId` | `string` | no |

## GetQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `queryString` | `string` | no |
| `waited` | `integer` | no |
| `elapsed` | `integer` | no |
| `state` | `string` | no |

## ListExportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<ExportTaskSummary>` | yes |
| `nextToken` | `string` | no |

## ListGraphSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphSnapshots` | `List<GraphSnapshotSummary>` | yes |
| `nextToken` | `string` | no |

## ListGraphs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphs` | `List<GraphSummary>` | yes |
| `nextToken` | `string` | no |

## ListImportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<ImportTaskSummary>` | yes |
| `nextToken` | `string` | no |

## ListPrivateGraphEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateGraphEndpoints` | `List<PrivateGraphEndpointSummary>` | yes |
| `nextToken` | `string` | no |

## ListQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `maxResults` | `integer` | yes |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queries` | `List<QuerySummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ResetGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `skipSnapshot` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## RestoreGraphFromSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `snapshotIdentifier` | `string` | yes |
| `graphName` | `string` | yes |
| `provisionedMemory` | `integer` | no |
| `deletionProtection` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `replicaCount` | `integer` | no |
| `publicConnectivity` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## StartExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `roleArn` | `string` | yes |
| `format` | `string` | yes |
| `destination` | `string` | yes |
| `kmsKeyIdentifier` | `string` | yes |
| `parquetType` | `string` | no |
| `exportFilter` | `ExportFilter` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | yes |
| `roleArn` | `string` | yes |
| `taskId` | `string` | yes |
| `status` | `string` | yes |
| `format` | `string` | yes |
| `destination` | `string` | yes |
| `kmsKeyIdentifier` | `string` | yes |
| `parquetType` | `string` | no |
| `statusReason` | `string` | no |
| `exportFilter` | `ExportFilter` | no |

## StartGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

## StartImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `importOptions` | `ImportOptions` | no |
| `failOnError` | `boolean` | no |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `blankNodeHandling` | `string` | no |
| `graphIdentifier` | `string` | yes |
| `roleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphId` | `string` | no |
| `taskId` | `string` | yes |
| `source` | `string` | yes |
| `format` | `string` | no |
| `parquetType` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | yes |
| `importOptions` | `ImportOptions` | no |

## StopGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

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


## UpdateGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphIdentifier` | `string` | yes |
| `publicConnectivity` | `boolean` | no |
| `provisionedMemory` | `integer` | no |
| `deletionProtection` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `status` | `string` | no |
| `statusReason` | `string` | no |
| `createTime` | `timestamp` | no |
| `provisionedMemory` | `integer` | no |
| `endpoint` | `string` | no |
| `publicConnectivity` | `boolean` | no |
| `vectorSearchConfiguration` | `VectorSearchConfiguration` | no |
| `replicaCount` | `integer` | no |
| `kmsKeyIdentifier` | `string` | no |
| `sourceSnapshotId` | `string` | no |
| `deletionProtection` | `boolean` | no |
| `buildNumber` | `string` | no |

