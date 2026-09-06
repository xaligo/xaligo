# AWS CodeBuild

API version: 2016-10-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codebuild/2016-10-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteBuilds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buildsDeleted` | `List<string>` | no |
| `buildsNotDeleted` | `List<BuildNotDeleted>` | no |

## BatchGetBuildBatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buildBatches` | `List<BuildBatch>` | no |
| `buildBatchesNotFound` | `List<string>` | no |

## BatchGetBuilds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `builds` | `List<Build>` | no |
| `buildsNotFound` | `List<string>` | no |

## BatchGetCommandExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandboxId` | `string` | yes |
| `commandExecutionIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandExecutions` | `List<CommandExecution>` | no |
| `commandExecutionsNotFound` | `List<string>` | no |

## BatchGetFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleets` | `List<Fleet>` | no |
| `fleetsNotFound` | `List<string>` | no |

## BatchGetProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projects` | `List<Project>` | no |
| `projectsNotFound` | `List<string>` | no |

## BatchGetReportGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroupArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroups` | `List<ReportGroup>` | no |
| `reportGroupsNotFound` | `List<string>` | no |

## BatchGetReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reports` | `List<Report>` | no |
| `reportsNotFound` | `List<string>` | no |

## BatchGetSandboxes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandboxes` | `List<Sandbox>` | no |
| `sandboxesNotFound` | `List<string>` | no |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `baseCapacity` | `integer` | yes |
| `environmentType` | `string` | yes |
| `computeType` | `string` | yes |
| `computeConfiguration` | `ComputeConfiguration` | no |
| `scalingConfiguration` | `ScalingConfigurationInput` | no |
| `overflowBehavior` | `string` | no |
| `vpcConfig` | `VpcConfig` | no |
| `proxyConfiguration` | `ProxyConfiguration` | no |
| `imageId` | `string` | no |
| `fleetServiceRole` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleet` | `Fleet` | no |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `source` | `ProjectSource` | yes |
| `secondarySources` | `List<ProjectSource>` | no |
| `sourceVersion` | `string` | no |
| `secondarySourceVersions` | `List<ProjectSourceVersion>` | no |
| `artifacts` | `ProjectArtifacts` | yes |
| `secondaryArtifacts` | `List<ProjectArtifacts>` | no |
| `cache` | `ProjectCache` | no |
| `environment` | `ProjectEnvironment` | yes |
| `serviceRole` | `string` | yes |
| `timeoutInMinutes` | `integer` | no |
| `queuedTimeoutInMinutes` | `integer` | no |
| `encryptionKey` | `string` | no |
| `tags` | `List<Tag>` | no |
| `vpcConfig` | `VpcConfig` | no |
| `badgeEnabled` | `boolean` | no |
| `logsConfig` | `LogsConfig` | no |
| `fileSystemLocations` | `List<ProjectFileSystemLocation>` | no |
| `buildBatchConfig` | `ProjectBuildBatchConfig` | no |
| `concurrentBuildLimit` | `integer` | no |
| `autoRetryLimit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `Project` | no |

## CreateReportGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `exportConfig` | `ReportExportConfig` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroup` | `ReportGroup` | no |

## CreateWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `branchFilter` | `string` | no |
| `filterGroups` | `List<List<WebhookFilter>>` | no |
| `buildType` | `string` | no |
| `manualCreation` | `boolean` | no |
| `scopeConfiguration` | `ScopeConfiguration` | no |
| `pullRequestBuildPolicy` | `PullRequestBuildPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | no |

## DeleteBuildBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `buildsDeleted` | `List<string>` | no |
| `buildsNotDeleted` | `List<BuildNotDeleted>` | no |

## DeleteFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReportGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `deleteReports` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSourceCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## DeleteWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCodeCoverages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `sortBy` | `string` | no |
| `minLineCoveragePercentage` | `double` | no |
| `maxLineCoveragePercentage` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `codeCoverages` | `List<CodeCoverage>` | no |

## DescribeTestCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `TestCaseFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `testCases` | `List<TestCase>` | no |

## GetReportGroupTrend

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroupArn` | `string` | yes |
| `numOfReports` | `integer` | no |
| `trendField` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stats` | `ReportGroupTrendStats` | no |
| `rawData` | `List<ReportWithRawData>` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |

## ImportSourceCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `username` | `string` | no |
| `token` | `string` | yes |
| `serverType` | `string` | yes |
| `authType` | `string` | yes |
| `shouldOverwrite` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## InvalidateProjectCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListBuildBatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `BuildBatchFilter` | no |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListBuildBatchesForProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | no |
| `filter` | `BuildBatchFilter` | no |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListBuilds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListBuildsForProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListCommandExecutionsForSandbox

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandboxId` | `string` | yes |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandExecutions` | `List<CommandExecution>` | no |
| `nextToken` | `string` | no |

## ListCuratedEnvironmentImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `platforms` | `List<EnvironmentPlatform>` | no |

## ListFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `sortBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `fleets` | `List<string>` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `projects` | `List<string>` | no |

## ListReportGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortOrder` | `string` | no |
| `sortBy` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reportGroups` | `List<string>` | no |

## ListReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `ReportFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reports` | `List<string>` | no |

## ListReportsForReportGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroupArn` | `string` | yes |
| `nextToken` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `ReportFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reports` | `List<string>` | no |

## ListSandboxes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListSandboxesForProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListSharedProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `projects` | `List<string>` | no |

## ListSharedReportGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sortOrder` | `string` | no |
| `sortBy` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `reportGroups` | `List<string>` | no |

## ListSourceCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceCredentialsInfos` | `List<SourceCredentialsInfo>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |

## RetryBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `idempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `build` | `Build` | no |

## RetryBuildBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `idempotencyToken` | `string` | no |
| `retryType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buildBatch` | `BuildBatch` | no |

## StartBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `secondarySourcesOverride` | `List<ProjectSource>` | no |
| `secondarySourcesVersionOverride` | `List<ProjectSourceVersion>` | no |
| `sourceVersion` | `string` | no |
| `artifactsOverride` | `ProjectArtifacts` | no |
| `secondaryArtifactsOverride` | `List<ProjectArtifacts>` | no |
| `environmentVariablesOverride` | `List<EnvironmentVariable>` | no |
| `sourceTypeOverride` | `string` | no |
| `sourceLocationOverride` | `string` | no |
| `sourceAuthOverride` | `SourceAuth` | no |
| `gitCloneDepthOverride` | `integer` | no |
| `gitSubmodulesConfigOverride` | `GitSubmodulesConfig` | no |
| `buildspecOverride` | `string` | no |
| `insecureSslOverride` | `boolean` | no |
| `reportBuildStatusOverride` | `boolean` | no |
| `buildStatusConfigOverride` | `BuildStatusConfig` | no |
| `environmentTypeOverride` | `string` | no |
| `imageOverride` | `string` | no |
| `computeTypeOverride` | `string` | no |
| `certificateOverride` | `string` | no |
| `cacheOverride` | `ProjectCache` | no |
| `serviceRoleOverride` | `string` | no |
| `privilegedModeOverride` | `boolean` | no |
| `timeoutInMinutesOverride` | `integer` | no |
| `queuedTimeoutInMinutesOverride` | `integer` | no |
| `encryptionKeyOverride` | `string` | no |
| `idempotencyToken` | `string` | no |
| `logsConfigOverride` | `LogsConfig` | no |
| `registryCredentialOverride` | `RegistryCredential` | no |
| `imagePullCredentialsTypeOverride` | `string` | no |
| `debugSessionEnabled` | `boolean` | no |
| `fleetOverride` | `ProjectFleet` | no |
| `autoRetryLimitOverride` | `integer` | no |
| `hostKernelOverride` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `build` | `Build` | no |

## StartBuildBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `secondarySourcesOverride` | `List<ProjectSource>` | no |
| `secondarySourcesVersionOverride` | `List<ProjectSourceVersion>` | no |
| `sourceVersion` | `string` | no |
| `artifactsOverride` | `ProjectArtifacts` | no |
| `secondaryArtifactsOverride` | `List<ProjectArtifacts>` | no |
| `environmentVariablesOverride` | `List<EnvironmentVariable>` | no |
| `sourceTypeOverride` | `string` | no |
| `sourceLocationOverride` | `string` | no |
| `sourceAuthOverride` | `SourceAuth` | no |
| `gitCloneDepthOverride` | `integer` | no |
| `gitSubmodulesConfigOverride` | `GitSubmodulesConfig` | no |
| `buildspecOverride` | `string` | no |
| `insecureSslOverride` | `boolean` | no |
| `reportBuildBatchStatusOverride` | `boolean` | no |
| `environmentTypeOverride` | `string` | no |
| `imageOverride` | `string` | no |
| `computeTypeOverride` | `string` | no |
| `certificateOverride` | `string` | no |
| `cacheOverride` | `ProjectCache` | no |
| `serviceRoleOverride` | `string` | no |
| `privilegedModeOverride` | `boolean` | no |
| `buildTimeoutInMinutesOverride` | `integer` | no |
| `queuedTimeoutInMinutesOverride` | `integer` | no |
| `encryptionKeyOverride` | `string` | no |
| `idempotencyToken` | `string` | no |
| `logsConfigOverride` | `LogsConfig` | no |
| `registryCredentialOverride` | `RegistryCredential` | no |
| `imagePullCredentialsTypeOverride` | `string` | no |
| `buildBatchConfigOverride` | `ProjectBuildBatchConfig` | no |
| `debugSessionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buildBatch` | `BuildBatch` | no |

## StartCommandExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandboxId` | `string` | yes |
| `command` | `string` | yes |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `commandExecution` | `CommandExecution` | no |

## StartSandbox

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | no |
| `idempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandbox` | `Sandbox` | no |

## StartSandboxConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandboxId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ssmSession` | `SSMSession` | no |

## StopBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `build` | `Build` | no |

## StopBuildBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `buildBatch` | `BuildBatch` | no |

## StopSandbox

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sandbox` | `Sandbox` | no |

## UpdateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `baseCapacity` | `integer` | no |
| `environmentType` | `string` | no |
| `computeType` | `string` | no |
| `computeConfiguration` | `ComputeConfiguration` | no |
| `scalingConfiguration` | `ScalingConfigurationInput` | no |
| `overflowBehavior` | `string` | no |
| `vpcConfig` | `VpcConfig` | no |
| `proxyConfiguration` | `ProxyConfiguration` | no |
| `imageId` | `string` | no |
| `fleetServiceRole` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleet` | `Fleet` | no |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `source` | `ProjectSource` | no |
| `secondarySources` | `List<ProjectSource>` | no |
| `sourceVersion` | `string` | no |
| `secondarySourceVersions` | `List<ProjectSourceVersion>` | no |
| `artifacts` | `ProjectArtifacts` | no |
| `secondaryArtifacts` | `List<ProjectArtifacts>` | no |
| `cache` | `ProjectCache` | no |
| `environment` | `ProjectEnvironment` | no |
| `serviceRole` | `string` | no |
| `timeoutInMinutes` | `integer` | no |
| `queuedTimeoutInMinutes` | `integer` | no |
| `encryptionKey` | `string` | no |
| `tags` | `List<Tag>` | no |
| `vpcConfig` | `VpcConfig` | no |
| `badgeEnabled` | `boolean` | no |
| `logsConfig` | `LogsConfig` | no |
| `fileSystemLocations` | `List<ProjectFileSystemLocation>` | no |
| `buildBatchConfig` | `ProjectBuildBatchConfig` | no |
| `concurrentBuildLimit` | `integer` | no |
| `autoRetryLimit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `Project` | no |

## UpdateProjectVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `projectVisibility` | `string` | yes |
| `resourceAccessRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | no |
| `publicProjectAlias` | `string` | no |
| `projectVisibility` | `string` | no |

## UpdateReportGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `exportConfig` | `ReportExportConfig` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reportGroup` | `ReportGroup` | no |

## UpdateWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectName` | `string` | yes |
| `branchFilter` | `string` | no |
| `rotateSecret` | `boolean` | no |
| `filterGroups` | `List<List<WebhookFilter>>` | no |
| `buildType` | `string` | no |
| `pullRequestBuildPolicy` | `PullRequestBuildPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | no |

