# AWS Amplify

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/amplify/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `repository` | `string` | no |
| `platform` | `string` | no |
| `computeRoleArn` | `string` | no |
| `iamServiceRoleArn` | `string` | no |
| `oauthToken` | `string` | no |
| `accessToken` | `string` | no |
| `environmentVariables` | `Map<string>` | no |
| `enableBranchAutoBuild` | `boolean` | no |
| `enableBranchAutoDeletion` | `boolean` | no |
| `enableBasicAuth` | `boolean` | no |
| `basicAuthCredentials` | `string` | no |
| `customRules` | `List<CustomRule>` | no |
| `tags` | `Map<string>` | no |
| `buildSpec` | `string` | no |
| `customHeaders` | `string` | no |
| `enableAutoBranchCreation` | `boolean` | no |
| `autoBranchCreationPatterns` | `List<string>` | no |
| `autoBranchCreationConfig` | `AutoBranchCreationConfig` | no |
| `jobConfig` | `JobConfig` | no |
| `cacheConfig` | `CacheConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## CreateBackendEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |
| `stackName` | `string` | no |
| `deploymentArtifacts` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `backendEnvironment` | `BackendEnvironment` | yes |

## CreateBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `description` | `string` | no |
| `stage` | `string` | no |
| `framework` | `string` | no |
| `enableNotification` | `boolean` | no |
| `enableAutoBuild` | `boolean` | no |
| `enableSkewProtection` | `boolean` | no |
| `environmentVariables` | `Map<string>` | no |
| `basicAuthCredentials` | `string` | no |
| `enableBasicAuth` | `boolean` | no |
| `enablePerformanceMode` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `buildSpec` | `string` | no |
| `ttl` | `string` | no |
| `displayName` | `string` | no |
| `enablePullRequestPreview` | `boolean` | no |
| `pullRequestEnvironmentName` | `string` | no |
| `backendEnvironmentArn` | `string` | no |
| `backend` | `Backend` | no |
| `computeRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `Branch` | yes |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `fileMap` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | no |
| `fileUploadUrls` | `Map<string>` | yes |
| `zipUploadUrl` | `string` | yes |

## CreateDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `domainName` | `string` | yes |
| `enableAutoSubDomain` | `boolean` | no |
| `subDomainSettings` | `List<SubDomainSetting>` | yes |
| `autoSubDomainCreationPatterns` | `List<string>` | no |
| `autoSubDomainIAMRole` | `string` | no |
| `certificateSettings` | `CertificateSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainAssociation` | `DomainAssociation` | yes |

## CreateWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | yes |

## DeleteApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## DeleteBackendEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `backendEnvironment` | `BackendEnvironment` | yes |

## DeleteBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `Branch` | yes |

## DeleteDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainAssociation` | `DomainAssociation` | yes |

## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummary` | `JobSummary` | yes |

## DeleteWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhookId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | yes |

## GenerateAccessLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `domainName` | `string` | yes |
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `logUrl` | `string` | no |

## GetApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## GetArtifactUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifactId` | `string` | yes |
| `artifactUrl` | `string` | yes |

## GetBackendEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `backendEnvironment` | `BackendEnvironment` | yes |

## GetBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `Branch` | yes |

## GetDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainAssociation` | `DomainAssociation` | yes |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | yes |

## GetWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhookId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | yes |

## ListApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apps` | `List<App>` | yes |
| `nextToken` | `string` | no |

## ListArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifacts` | `List<Artifact>` | yes |
| `nextToken` | `string` | no |

## ListBackendEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `environmentName` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `backendEnvironments` | `List<BackendEnvironment>` | yes |
| `nextToken` | `string` | no |

## ListBranches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branches` | `List<Branch>` | yes |
| `nextToken` | `string` | no |

## ListDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainAssociations` | `List<DomainAssociation>` | yes |
| `nextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaries` | `List<JobSummary>` | yes |
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

## ListWebhooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhooks` | `List<Webhook>` | yes |
| `nextToken` | `string` | no |

## StartDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | no |
| `sourceUrl` | `string` | no |
| `sourceUrlType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummary` | `JobSummary` | yes |

## StartJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | no |
| `jobType` | `string` | yes |
| `jobReason` | `string` | no |
| `commitId` | `string` | no |
| `commitMessage` | `string` | no |
| `commitTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummary` | `JobSummary` | yes |

## StopJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummary` | `JobSummary` | yes |

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


## UpdateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `platform` | `string` | no |
| `computeRoleArn` | `string` | no |
| `iamServiceRoleArn` | `string` | no |
| `environmentVariables` | `Map<string>` | no |
| `enableBranchAutoBuild` | `boolean` | no |
| `enableBranchAutoDeletion` | `boolean` | no |
| `enableBasicAuth` | `boolean` | no |
| `basicAuthCredentials` | `string` | no |
| `customRules` | `List<CustomRule>` | no |
| `buildSpec` | `string` | no |
| `customHeaders` | `string` | no |
| `enableAutoBranchCreation` | `boolean` | no |
| `autoBranchCreationPatterns` | `List<string>` | no |
| `autoBranchCreationConfig` | `AutoBranchCreationConfig` | no |
| `repository` | `string` | no |
| `oauthToken` | `string` | no |
| `accessToken` | `string` | no |
| `jobConfig` | `JobConfig` | no |
| `cacheConfig` | `CacheConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `app` | `App` | yes |

## UpdateBranch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `branchName` | `string` | yes |
| `description` | `string` | no |
| `framework` | `string` | no |
| `stage` | `string` | no |
| `enableNotification` | `boolean` | no |
| `enableAutoBuild` | `boolean` | no |
| `enableSkewProtection` | `boolean` | no |
| `environmentVariables` | `Map<string>` | no |
| `basicAuthCredentials` | `string` | no |
| `enableBasicAuth` | `boolean` | no |
| `enablePerformanceMode` | `boolean` | no |
| `buildSpec` | `string` | no |
| `ttl` | `string` | no |
| `displayName` | `string` | no |
| `enablePullRequestPreview` | `boolean` | no |
| `pullRequestEnvironmentName` | `string` | no |
| `backendEnvironmentArn` | `string` | no |
| `backend` | `Backend` | no |
| `computeRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `Branch` | yes |

## UpdateDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appId` | `string` | yes |
| `domainName` | `string` | yes |
| `enableAutoSubDomain` | `boolean` | no |
| `subDomainSettings` | `List<SubDomainSetting>` | no |
| `autoSubDomainCreationPatterns` | `List<string>` | no |
| `autoSubDomainIAMRole` | `string` | no |
| `certificateSettings` | `CertificateSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainAssociation` | `DomainAssociation` | yes |

## UpdateWebhook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhookId` | `string` | yes |
| `branchName` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhook` | `Webhook` | yes |

