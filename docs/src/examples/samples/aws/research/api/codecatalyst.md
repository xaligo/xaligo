# Amazon CodeCatalyst

API version: 2022-09-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codecatalyst/2022-09-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `expiresTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `secret` | `string` | yes |
| `name` | `string` | yes |
| `expiresTime` | `timestamp` | yes |
| `accessTokenId` | `string` | yes |

## CreateDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `repositories` | `List<RepositoryInput>` | no |
| `clientToken` | `string` | no |
| `alias` | `string` | no |
| `ides` | `List<IdeConfiguration>` | no |
| `instanceType` | `string` | yes |
| `inactivityTimeoutMinutes` | `integer` | no |
| `persistentStorage` | `PersistentStorageConfiguration` | yes |
| `vpcConnectionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `vpcConnectionName` | `string` | no |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `displayName` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | no |
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |

## CreateSourceRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

## CreateSourceRepositoryBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `sourceRepositoryName` | `string` | yes |
| `name` | `string` | yes |
| `headCommitId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ref` | `string` | no |
| `name` | `string` | no |
| `lastUpdatedTime` | `timestamp` | no |
| `headCommitId` | `string` | no |

## DeleteAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |

## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `name` | `string` | yes |
| `displayName` | `string` | no |

## DeleteSourceRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |

## DeleteSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `displayName` | `string` | no |

## GetDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `lastUpdatedTime` | `timestamp` | yes |
| `creatorId` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `repositories` | `List<DevEnvironmentRepositorySummary>` | yes |
| `alias` | `string` | no |
| `ides` | `List<Ide>` | no |
| `instanceType` | `string` | yes |
| `inactivityTimeoutMinutes` | `integer` | yes |
| `persistentStorage` | `PersistentStorage` | yes |
| `vpcConnectionName` | `string` | no |

## GetProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | no |
| `name` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |

## GetSourceRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `lastUpdatedTime` | `timestamp` | yes |
| `createdTime` | `timestamp` | yes |

## GetSourceRepositoryCloneUrls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `sourceRepositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `https` | `string` | yes |

## GetSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `regionName` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |

## GetSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptionType` | `string` | no |
| `awsAccountName` | `string` | no |
| `pendingSubscriptionType` | `string` | no |
| `pendingSubscriptionStartTime` | `timestamp` | no |

## GetUserDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `userName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | no |
| `userName` | `string` | no |
| `displayName` | `string` | no |
| `primaryEmail` | `EmailAddress` | no |
| `version` | `string` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `id` | `string` | yes |
| `projectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `sourceRepositoryName` | `string` | no |
| `sourceBranchName` | `string` | no |
| `definition` | `WorkflowDefinition` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |
| `runMode` | `string` | yes |
| `status` | `string` | yes |

## GetWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `id` | `string` | yes |
| `projectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `workflowId` | `string` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<WorkflowRunStatusReason>` | no |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | no |
| `lastUpdatedTime` | `timestamp` | yes |

## ListAccessTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AccessTokenSummary>` | yes |
| `nextToken` | `string` | no |

## ListDevEnvironmentSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `devEnvironmentId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DevEnvironmentSessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListDevEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | no |
| `filters` | `List<Filter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DevEnvironmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListEventLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `startTime` | `timestamp` | yes |
| `endTime` | `timestamp` | yes |
| `eventName` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<EventLogEntry>` | yes |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filters` | `List<ProjectListFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ProjectSummary>` | no |

## ListSourceRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ListSourceRepositoriesItem>` | no |
| `nextToken` | `string` | no |

## ListSourceRepositoryBranches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `sourceRepositoryName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<ListSourceRepositoryBranchesItem>` | yes |

## ListSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<SpaceSummary>` | no |

## ListWorkflowRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `workflowId` | `string` | no |
| `projectName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortBy` | `List<WorkflowRunSortCriteria>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<WorkflowRunSummary>` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortBy` | `List<WorkflowSortCriteria>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<WorkflowSummary>` | no |

## StartDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `ides` | `List<IdeConfiguration>` | no |
| `instanceType` | `string` | no |
| `inactivityTimeoutMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |

## StartDevEnvironmentSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `sessionConfiguration` | `DevEnvironmentSessionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessDetails` | `DevEnvironmentAccessDetails` | yes |
| `sessionId` | `string` | no |
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |

## StartWorkflowRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `workflowId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `workflowId` | `string` | yes |

## StopDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |

## StopDevEnvironmentSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `sessionId` | `string` | yes |

## UpdateDevEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `id` | `string` | yes |
| `alias` | `string` | no |
| `ides` | `List<IdeConfiguration>` | no |
| `instanceType` | `string` | no |
| `inactivityTimeoutMinutes` | `integer` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `spaceName` | `string` | yes |
| `projectName` | `string` | yes |
| `alias` | `string` | no |
| `ides` | `List<IdeConfiguration>` | no |
| `instanceType` | `string` | no |
| `inactivityTimeoutMinutes` | `integer` | no |
| `clientToken` | `string` | no |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `spaceName` | `string` | no |
| `name` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |

## UpdateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |

## VerifySession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identity` | `string` | no |

