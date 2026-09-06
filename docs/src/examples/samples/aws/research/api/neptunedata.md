# Amazon NeptuneData

API version: 2023-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/neptunedata/2023-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelGremlinQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CancelLoaderJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CancelMLDataProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |
| `clean` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CancelMLModelTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |
| `clean` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CancelMLModelTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |
| `clean` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## CancelOpenCypherQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |
| `silent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `payload` | `boolean` | no |

## CreateMLEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `mlModelTrainingJobId` | `string` | no |
| `mlModelTransformJobId` | `string` | no |
| `update` | `boolean` | no |
| `neptuneIamRoleArn` | `string` | no |
| `modelName` | `string` | no |
| `instanceType` | `string` | no |
| `instanceCount` | `integer` | no |
| `volumeEncryptionKMSKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `creationTimeInMillis` | `long` | no |

## DeleteMLEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |
| `clean` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## DeletePropertygraphStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |
| `status` | `string` | no |
| `payload` | `DeleteStatisticsValueMap` | no |

## DeleteSparqlStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |
| `status` | `string` | no |
| `payload` | `DeleteStatisticsValueMap` | no |

## ExecuteFastReset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `token` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `FastResetToken` | no |

## ExecuteGremlinExplainQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gremlinQuery` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `output` | `blob` | no |

## ExecuteGremlinProfileQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gremlinQuery` | `string` | yes |
| `results` | `boolean` | no |
| `chop` | `integer` | no |
| `serializer` | `string` | no |
| `indexOps` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `output` | `blob` | no |

## ExecuteGremlinQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gremlinQuery` | `string` | yes |
| `serializer` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `status` | `GremlinQueryStatusAttributes` | no |
| `result` | `Document` | no |
| `meta` | `Document` | no |

## ExecuteOpenCypherExplainQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `openCypherQuery` | `string` | yes |
| `parameters` | `string` | no |
| `explainMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `blob` | yes |

## ExecuteOpenCypherQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `openCypherQuery` | `string` | yes |
| `parameters` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `results` | `Document` | yes |

## GetEngineStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `startTime` | `string` | no |
| `dbEngineVersion` | `string` | no |
| `role` | `string` | no |
| `dfeQueryEngine` | `string` | no |
| `gremlin` | `QueryLanguageVersion` | no |
| `sparql` | `QueryLanguageVersion` | no |
| `opencypher` | `QueryLanguageVersion` | no |
| `labMode` | `Map<string>` | no |
| `rollingBackTrxCount` | `integer` | no |
| `rollingBackTrxEarliestStartTime` | `string` | no |
| `features` | `Map<Document>` | no |
| `settings` | `Map<string>` | no |

## GetGremlinQueryStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | no |
| `queryString` | `string` | no |
| `queryEvalStats` | `QueryEvalStats` | no |

## GetLoaderJobStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loadId` | `string` | yes |
| `details` | `boolean` | no |
| `errors` | `boolean` | no |
| `page` | `integer` | no |
| `errorsPerPage` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `Document` | yes |

## GetMLDataProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `id` | `string` | no |
| `processingJob` | `MlResourceDefinition` | no |

## GetMLEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `id` | `string` | no |
| `endpoint` | `MlResourceDefinition` | no |
| `endpointConfig` | `MlConfigDefinition` | no |

## GetMLModelTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `id` | `string` | no |
| `processingJob` | `MlResourceDefinition` | no |
| `hpoJob` | `MlResourceDefinition` | no |
| `modelTransformJob` | `MlResourceDefinition` | no |
| `mlModels` | `List<MlConfigDefinition>` | no |

## GetMLModelTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `id` | `string` | no |
| `baseProcessingJob` | `MlResourceDefinition` | no |
| `remoteModelTransformJob` | `MlResourceDefinition` | no |
| `models` | `List<MlConfigDefinition>` | no |

## GetOpenCypherQueryStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `queryId` | `string` | no |
| `queryString` | `string` | no |
| `queryEvalStats` | `QueryEvalStats` | no |

## GetPropertygraphStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `Statistics` | yes |

## GetPropertygraphStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `long` | no |
| `iteratorType` | `string` | no |
| `commitNum` | `long` | no |
| `opNum` | `long` | no |
| `encoding` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lastEventId` | `Map<string>` | yes |
| `lastTrxTimestampInMillis` | `long` | yes |
| `format` | `string` | yes |
| `records` | `List<PropertygraphRecord>` | yes |
| `totalRecords` | `integer` | yes |

## GetPropertygraphSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |
| `payload` | `PropertygraphSummaryValueMap` | no |

## GetRDFGraphSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `integer` | no |
| `payload` | `RDFGraphSummaryValueMap` | no |

## GetSparqlStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `Statistics` | yes |

## GetSparqlStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `long` | no |
| `iteratorType` | `string` | no |
| `commitNum` | `long` | no |
| `opNum` | `long` | no |
| `encoding` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lastEventId` | `Map<string>` | yes |
| `lastTrxTimestampInMillis` | `long` | yes |
| `format` | `string` | yes |
| `records` | `List<SparqlRecord>` | yes |
| `totalRecords` | `integer` | yes |

## ListGremlinQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeWaiting` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acceptedQueryCount` | `integer` | no |
| `runningQueryCount` | `integer` | no |
| `queries` | `List<GremlinQueryStatus>` | no |

## ListLoaderJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `limit` | `integer` | no |
| `includeQueuedLoads` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `LoaderIdResult` | yes |

## ListMLDataProcessingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxItems` | `integer` | no |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |

## ListMLEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxItems` | `integer` | no |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |

## ListMLModelTrainingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxItems` | `integer` | no |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |

## ListMLModelTransformJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxItems` | `integer` | no |
| `neptuneIamRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |

## ListOpenCypherQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `includeWaiting` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acceptedQueryCount` | `integer` | no |
| `runningQueryCount` | `integer` | no |
| `queries` | `List<GremlinQueryStatus>` | no |

## ManagePropertygraphStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `RefreshStatisticsIdMap` | no |

## ManageSparqlStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `RefreshStatisticsIdMap` | no |

## StartLoaderJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `source` | `string` | yes |
| `format` | `string` | yes |
| `s3BucketRegion` | `string` | yes |
| `iamRoleArn` | `string` | yes |
| `mode` | `string` | no |
| `failOnError` | `boolean` | no |
| `parallelism` | `string` | no |
| `parserConfiguration` | `Map<string>` | no |
| `updateSingleCardinalityProperties` | `boolean` | no |
| `queueRequest` | `boolean` | no |
| `dependencies` | `List<string>` | no |
| `userProvidedEdgeIds` | `boolean` | no |
| `edgeOnlyLoad` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `payload` | `Map<string>` | yes |

## StartMLDataProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `previousDataProcessingJobId` | `string` | no |
| `inputDataS3Location` | `string` | yes |
| `processedDataS3Location` | `string` | yes |
| `sagemakerIamRoleArn` | `string` | no |
| `neptuneIamRoleArn` | `string` | no |
| `processingInstanceType` | `string` | no |
| `processingInstanceVolumeSizeInGB` | `integer` | no |
| `processingTimeOutInSeconds` | `integer` | no |
| `modelType` | `string` | no |
| `configFileName` | `string` | no |
| `subnets` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `volumeEncryptionKMSKey` | `string` | no |
| `s3OutputEncryptionKMSKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `creationTimeInMillis` | `long` | no |

## StartMLModelTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `previousModelTrainingJobId` | `string` | no |
| `dataProcessingJobId` | `string` | yes |
| `trainModelS3Location` | `string` | yes |
| `sagemakerIamRoleArn` | `string` | no |
| `neptuneIamRoleArn` | `string` | no |
| `baseProcessingInstanceType` | `string` | no |
| `trainingInstanceType` | `string` | no |
| `trainingInstanceVolumeSizeInGB` | `integer` | no |
| `trainingTimeOutInSeconds` | `integer` | no |
| `maxHPONumberOfTrainingJobs` | `integer` | no |
| `maxHPOParallelTrainingJobs` | `integer` | no |
| `subnets` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `volumeEncryptionKMSKey` | `string` | no |
| `s3OutputEncryptionKMSKey` | `string` | no |
| `enableManagedSpotTraining` | `boolean` | no |
| `customModelTrainingParameters` | `CustomModelTrainingParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `creationTimeInMillis` | `long` | no |

## StartMLModelTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `dataProcessingJobId` | `string` | no |
| `mlModelTrainingJobId` | `string` | no |
| `trainingJobName` | `string` | no |
| `modelTransformOutputS3Location` | `string` | yes |
| `sagemakerIamRoleArn` | `string` | no |
| `neptuneIamRoleArn` | `string` | no |
| `customModelTransformParameters` | `CustomModelTransformParameters` | no |
| `baseProcessingInstanceType` | `string` | no |
| `baseProcessingInstanceVolumeSizeInGB` | `integer` | no |
| `subnets` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `volumeEncryptionKMSKey` | `string` | no |
| `s3OutputEncryptionKMSKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `creationTimeInMillis` | `long` | no |

