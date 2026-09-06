# Migration Hub Strategy Recommendations

API version: 2020-02-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/migrationhubstrategy/2020-02-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetApplicationComponentDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentDetail` | `ApplicationComponentDetail` | no |
| `associatedApplications` | `List<AssociatedApplication>` | no |
| `associatedServerIds` | `List<string>` | no |
| `moreApplicationResource` | `boolean` | no |

## GetApplicationComponentStrategies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentStrategies` | `List<ApplicationComponentStrategy>` | no |

## GetAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentTargets` | `List<AssessmentTarget>` | no |
| `dataCollectionDetails` | `DataCollectionDetails` | no |
| `id` | `string` | no |

## GetImportFileTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `completionTime` | `timestamp` | no |
| `id` | `string` | no |
| `importName` | `string` | no |
| `inputS3Bucket` | `string` | no |
| `inputS3Key` | `string` | no |
| `numberOfRecordsFailed` | `integer` | no |
| `numberOfRecordsSuccess` | `integer` | no |
| `startTime` | `timestamp` | no |
| `status` | `string` | no |
| `statusReportS3Bucket` | `string` | no |
| `statusReportS3Key` | `string` | no |

## GetLatestAssessmentId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## GetPortfolioPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationMode` | `string` | no |
| `applicationPreferences` | `ApplicationPreferences` | no |
| `databasePreferences` | `DatabasePreferences` | no |
| `prioritizeBusinessGoals` | `PrioritizeBusinessGoals` | no |

## GetPortfolioSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentSummary` | `AssessmentSummary` | no |

## GetRecommendationReportDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `recommendationReportDetails` | `RecommendationReportDetails` | no |

## GetServerDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedApplications` | `List<AssociatedApplication>` | no |
| `nextToken` | `string` | no |
| `serverDetail` | `ServerDetail` | no |

## GetServerStrategies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serverStrategies` | `List<ServerStrategy>` | no |

## ListAnalyzableServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sort` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `analyzableServers` | `List<AnalyzableServerSummary>` | no |
| `nextToken` | `string` | no |

## ListApplicationComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentCriteria` | `string` | no |
| `filterValue` | `string` | no |
| `groupIdFilter` | `List<Group>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sort` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationComponentInfos` | `List<ApplicationComponentDetail>` | no |
| `nextToken` | `string` | no |

## ListCollectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Collectors` | `List<Collector>` | no |
| `nextToken` | `string` | no |

## ListImportFileTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `taskInfos` | `List<ImportFileTaskInformation>` | no |

## ListServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filterValue` | `string` | no |
| `groupIdFilter` | `List<Group>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serverCriteria` | `string` | no |
| `sort` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serverInfos` | `List<ServerDetail>` | no |

## PutPortfolioPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationMode` | `string` | no |
| `applicationPreferences` | `ApplicationPreferences` | no |
| `databasePreferences` | `DatabasePreferences` | no |
| `prioritizeBusinessGoals` | `PrioritizeBusinessGoals` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentDataSourceType` | `string` | no |
| `assessmentTargets` | `List<AssessmentTarget>` | no |
| `s3bucketForAnalysisData` | `string` | no |
| `s3bucketForReportData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | no |

## StartImportFileTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3Bucket` | `string` | yes |
| `dataSourceType` | `string` | no |
| `groupId` | `List<Group>` | no |
| `name` | `string` | yes |
| `s3bucketForReportData` | `string` | no |
| `s3key` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## StartRecommendationReportGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `groupIdFilter` | `List<Group>` | no |
| `outputFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## StopAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `assessmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplicationComponentConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appType` | `string` | no |
| `applicationComponentId` | `string` | yes |
| `configureOnly` | `boolean` | no |
| `inclusionStatus` | `string` | no |
| `secretsManagerKey` | `string` | no |
| `sourceCodeList` | `List<SourceCode>` | no |
| `strategyOption` | `StrategyOption` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serverId` | `string` | yes |
| `strategyOption` | `StrategyOption` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


