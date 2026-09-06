# AWS Security Agent

API version: 2025-09-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/securityagent/2025-09-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `artifactContent` | `blob` | yes |
| `artifactType` | `string` | yes |
| `fileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifactId` | `string` | yes |

## BatchCreateSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `securityRequirements` | `List<CreateSecurityRequirementEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityRequirements` | `List<BatchCreateSecurityRequirementResult>` | yes |
| `errors` | `List<BatchSecurityRequirementError>` | yes |

## BatchDeleteCodeReviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleted` | `List<string>` | no |
| `failed` | `List<DeleteCodeReviewFailure>` | no |

## BatchDeletePentests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleted` | `List<Pentest>` | no |
| `failed` | `List<DeletePentestFailure>` | no |

## BatchDeleteSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `securityRequirementNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deletedSecurityRequirementNames` | `List<string>` | yes |
| `errors` | `List<BatchSecurityRequirementError>` | yes |

## BatchDeleteThreatModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleted` | `List<string>` | no |
| `failed` | `List<DeleteThreatModelFailure>` | no |

## BatchGetAgentSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaces` | `List<AgentSpace>` | no |
| `notFound` | `List<string>` | no |

## BatchGetArtifactMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `artifactIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifactMetadataList` | `List<ArtifactMetadataItem>` | yes |

## BatchGetCodeReviewJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `codeReviewJobTaskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewJobTasks` | `List<CodeReviewJobTask>` | no |
| `notFound` | `List<string>` | no |

## BatchGetCodeReviewJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewJobIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewJobs` | `List<CodeReviewJob>` | no |
| `notFound` | `List<string>` | no |

## BatchGetCodeReviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviews` | `List<CodeReview>` | no |
| `notFound` | `List<string>` | no |

## BatchGetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findings` | `List<Finding>` | no |
| `notFound` | `List<string>` | no |

## BatchGetPentestJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `taskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<Task>` | no |
| `notFound` | `List<string>` | no |

## BatchGetPentestJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestJobIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestJobs` | `List<PentestJob>` | no |
| `notFound` | `List<string>` | no |

## BatchGetPentests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentests` | `List<Pentest>` | no |
| `notFound` | `List<string>` | no |

## BatchGetSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `securityRequirementNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityRequirements` | `List<BatchGetSecurityRequirementResult>` | yes |
| `errors` | `List<BatchSecurityRequirementError>` | yes |

## BatchGetTargetDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomains` | `List<TargetDomain>` | no |
| `notFound` | `List<string>` | no |

## BatchGetThreatModelJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `threatModelJobTaskIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelJobTasks` | `List<ThreatModelJobTask>` | no |
| `notFound` | `List<string>` | no |

## BatchGetThreatModelJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelJobIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelJobs` | `List<ThreatModelJob>` | no |
| `notFound` | `List<string>` | no |

## BatchGetThreatModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModels` | `List<ThreatModel>` | no |
| `notFound` | `List<string>` | no |

## BatchGetThreats

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatIds` | `List<string>` | yes |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threats` | `List<Threat>` | no |
| `notFound` | `List<string>` | no |

## BatchUpdateSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `securityRequirements` | `List<UpdateSecurityRequirementEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updatedSecurityRequirementNames` | `List<string>` | yes |
| `errors` | `List<BatchSecurityRequirementError>` | yes |

## CreateAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `awsResources` | `AWSResources` | no |
| `targetDomainIds` | `List<string>` | no |
| `codeReviewSettings` | `CodeReviewSettings` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `awsResources` | `AWSResources` | no |
| `targetDomainIds` | `List<string>` | no |
| `codeReviewSettings` | `CodeReviewSettings` | no |
| `kmsKeyId` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `idcInstanceArn` | `string` | no |
| `roleArn` | `string` | no |
| `defaultKmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

## CreateCodeReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `assets` | `Assets` | yes |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `codeRemediationStrategy` | `string` | no |
| `validationMode` | `string` | no |
| `maxTaskHours` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewId` | `string` | yes |
| `title` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `assets` | `Assets` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `agentSpaceId` | `string` | no |
| `codeRemediationStrategy` | `string` | no |
| `validationMode` | `string` | no |
| `maxTaskHours` | `double` | no |

## CreateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provider` | `string` | yes |
| `input` | `ProviderInput` | yes |
| `integrationDisplayName` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `privateConnectionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationId` | `string` | yes |

## CreateMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `membershipId` | `string` | yes |
| `memberType` | `string` | yes |
| `config` | `MembershipConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePentest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `assets` | `Assets` | no |
| `excludeRiskTypes` | `List<string>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `vpcConfig` | `VpcConfig` | no |
| `networkTrafficConfig` | `NetworkTrafficConfig` | no |
| `codeRemediationStrategy` | `string` | no |
| `disableManagedSkills` | `List<string>` | no |
| `maxTaskHours` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestId` | `string` | no |
| `title` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `assets` | `Assets` | no |
| `excludeRiskTypes` | `List<string>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `agentSpaceId` | `string` | no |

## CreatePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnectionName` | `string` | yes |
| `mode` | `PrivateConnectionMode` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `status` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateSecurityRequirementPack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `status` | `string` | yes |
| `kmsKeyId` | `string` | no |

## CreateTargetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainName` | `string` | yes |
| `verificationMethod` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | yes |
| `domainName` | `string` | yes |
| `verificationStatus` | `string` | yes |
| `verificationStatusReason` | `string` | no |
| `verificationDetails` | `VerificationDetails` | no |
| `createdAt` | `timestamp` | no |
| `verifiedAt` | `timestamp` | no |

## CreateThreat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `threatJobId` | `string` | yes |
| `title` | `string` | no |
| `statement` | `string` | no |
| `severity` | `string` | no |
| `comments` | `string` | no |
| `stride` | `List<string>` | no |
| `threatSource` | `string` | no |
| `prerequisites` | `string` | no |
| `threatAction` | `string` | no |
| `threatImpact` | `string` | no |
| `impactedGoal` | `List<string>` | no |
| `impactedAssets` | `List<string>` | no |
| `anchor` | `ThreatAnchorShape` | no |
| `evidence` | `List<ThreatEvidenceShape>` | no |
| `recommendation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatId` | `string` | yes |
| `threatJobId` | `string` | yes |
| `title` | `string` | no |
| `statement` | `string` | no |
| `severity` | `string` | no |
| `status` | `string` | no |
| `comments` | `string` | no |
| `stride` | `List<string>` | no |
| `threatSource` | `string` | no |
| `prerequisites` | `string` | no |
| `threatAction` | `string` | no |
| `threatImpact` | `string` | no |
| `impactedGoal` | `List<string>` | no |
| `impactedAssets` | `List<string>` | no |
| `anchor` | `ThreatAnchorShape` | no |
| `evidence` | `List<ThreatEvidenceShape>` | no |
| `recommendation` | `string` | no |
| `createdBy` | `string` | no |
| `updatedBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CreateThreatModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `description` | `string` | no |
| `assets` | `Assets` | no |
| `scopeDocs` | `List<DocumentInfo>` | no |
| `serviceRole` | `string` | yes |
| `logConfig` | `CloudWatchLog` | no |
| `reportDestination` | `ReportDestination` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelId` | `string` | yes |
| `title` | `string` | no |
| `agentSpaceId` | `string` | no |
| `description` | `string` | no |
| `assets` | `Assets` | no |
| `scopeDocs` | `List<DocumentInfo>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## DeleteAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `artifactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `membershipId` | `string` | yes |
| `memberType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `status` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## DeleteSecurityRequirementPack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTargetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | no |

## DescribePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnectionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `status` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `domain` | `string` | yes |
| `applicationName` | `string` | no |
| `idcConfiguration` | `IdCConfiguration` | no |
| `roleArn` | `string` | no |
| `defaultKmsKeyId` | `string` | no |

## GetArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `artifactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `artifactId` | `string` | yes |
| `artifact` | `Artifact` | yes |
| `fileName` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## GetIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationId` | `string` | yes |
| `installationId` | `string` | yes |
| `provider` | `string` | yes |
| `providerType` | `string` | yes |
| `displayName` | `string` | no |
| `kmsKeyId` | `string` | no |
| `targetUrl` | `string` | no |
| `privateConnectionName` | `string` | no |

## GetSecurityRequirementPack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `vendorName` | `string` | no |
| `managementType` | `string` | yes |
| `status` | `string` | yes |
| `importStatus` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `kmsKeyId` | `string` | no |

## ImportSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `input` | `ImportSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `importStatus` | `string` | yes |

## InitiateProviderRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provider` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `redirectTo` | `string` | yes |
| `csrfState` | `string` | yes |

## ListAgentSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceSummaries` | `List<AgentSpaceSummary>` | no |
| `nextToken` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationSummaries` | `List<ApplicationSummary>` | yes |
| `nextToken` | `string` | no |

## ListArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifactSummaries` | `List<ArtifactSummary>` | yes |
| `nextToken` | `string` | no |

## ListCodeReviewJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `codeReviewJobId` | `string` | no |
| `stepName` | `string` | no |
| `categoryName` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewJobTaskSummaries` | `List<CodeReviewJobTaskSummary>` | no |
| `nextToken` | `string` | no |

## ListCodeReviewJobsForCodeReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `codeReviewId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewJobSummaries` | `List<CodeReviewJobSummary>` | no |
| `nextToken` | `string` | no |

## ListCodeReviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewSummaries` | `List<CodeReviewSummary>` | no |
| `nextToken` | `string` | no |

## ListDiscoveredEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `pentestJobId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `prefix` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `discoveredEndpoints` | `List<DiscoveredEndpoint>` | no |
| `nextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `pentestJobId` | `string` | no |
| `codeReviewJobId` | `string` | no |
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |
| `riskType` | `string` | no |
| `riskLevel` | `string` | no |
| `status` | `string` | no |
| `confidence` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingsSummaries` | `List<FindingSummary>` | no |
| `nextToken` | `string` | no |

## ListIntegratedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `integrationId` | `string` | no |
| `resourceType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integratedResourceSummaries` | `List<IntegratedResourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `IntegrationFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `integrationSummaries` | `List<IntegrationSummary>` | yes |
| `nextToken` | `string` | no |

## ListMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `memberType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `membershipSummaries` | `List<MembershipSummary>` | yes |
| `nextToken` | `string` | no |

## ListPentestJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `pentestJobId` | `string` | no |
| `stepName` | `string` | no |
| `categoryName` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSummaries` | `List<TaskSummary>` | no |
| `nextToken` | `string` | no |

## ListPentestJobsForPentest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `pentestId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestJobSummaries` | `List<PentestJobSummary>` | no |
| `nextToken` | `string` | no |

## ListPentests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestSummaries` | `List<PentestSummary>` | no |
| `nextToken` | `string` | no |

## ListPrivateConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnections` | `List<PrivateConnectionSummary>` | yes |
| `nextToken` | `string` | no |

## ListSecurityRequirementPacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `ListSecurityRequirementPackFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityRequirementPackSummaries` | `List<SecurityRequirementPackSummary>` | yes |
| `nextToken` | `string` | no |

## ListSecurityRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityRequirementSummaries` | `List<SecurityRequirementSummary>` | yes |
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

## ListTargetDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainSummaries` | `List<TargetDomainSummary>` | no |
| `nextToken` | `string` | no |

## ListThreatModelJobTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `threatModelJobId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelJobTaskSummaries` | `List<ThreatModelJobTaskSummary>` | no |
| `nextToken` | `string` | no |

## ListThreatModelJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `threatModelId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelJobSummaries` | `List<ThreatModelJobSummary>` | no |
| `nextToken` | `string` | no |

## ListThreatModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelSummaries` | `List<ThreatModelSummary>` | no |
| `nextToken` | `string` | no |

## ListThreats

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatJobId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threats` | `List<ThreatSummary>` | no |
| `nextToken` | `string` | no |

## StartCodeRemediation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `pentestJobId` | `string` | no |
| `codeReviewJobId` | `string` | no |
| `findingIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartCodeReviewJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `codeReviewId` | `string` | yes |
| `diffSource` | `DiffSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `codeReviewId` | `string` | yes |
| `codeReviewJobId` | `string` | yes |
| `agentSpaceId` | `string` | no |

## StartPentestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `pentestId` | `string` | yes |
| `jobType` | `string` | no |
| `selectedFindingIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `pentestId` | `string` | no |
| `pentestJobId` | `string` | no |
| `agentSpaceId` | `string` | no |

## StartThreatModelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `threatModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `title` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `threatModelId` | `string` | no |
| `threatModelJobId` | `string` | yes |
| `agentSpaceId` | `string` | no |

## StopCodeReviewJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `codeReviewJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopPentestJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `pentestJobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopThreatModelJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `threatModelJobId` | `string` | yes |

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


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `awsResources` | `AWSResources` | no |
| `targetDomainIds` | `List<string>` | no |
| `codeReviewSettings` | `CodeReviewSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `awsResources` | `AWSResources` | no |
| `targetDomainIds` | `List<string>` | no |
| `codeReviewSettings` | `CodeReviewSettings` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `roleArn` | `string` | no |
| `defaultKmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |

## UpdateCodeReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `title` | `string` | no |
| `assets` | `Assets` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `codeRemediationStrategy` | `string` | no |
| `validationMode` | `string` | no |
| `maxTaskHours` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeReviewId` | `string` | yes |
| `title` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `assets` | `Assets` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `agentSpaceId` | `string` | no |
| `codeRemediationStrategy` | `string` | no |
| `validationMode` | `string` | no |
| `maxTaskHours` | `double` | no |

## UpdateFinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `findingId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `riskType` | `string` | no |
| `riskLevel` | `string` | no |
| `riskScore` | `string` | no |
| `attackScript` | `string` | no |
| `reasoning` | `string` | no |
| `status` | `string` | no |
| `customerNote` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIntegratedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `integrationId` | `string` | yes |
| `items` | `List<IntegratedResourceInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePentest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `title` | `string` | no |
| `assets` | `Assets` | no |
| `excludeRiskTypes` | `List<string>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `vpcConfig` | `VpcConfig` | no |
| `networkTrafficConfig` | `NetworkTrafficConfig` | no |
| `codeRemediationStrategy` | `string` | no |
| `disableManagedSkills` | `List<string>` | no |
| `maxTaskHours` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pentestId` | `string` | no |
| `title` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `assets` | `Assets` | no |
| `excludeRiskTypes` | `List<string>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `agentSpaceId` | `string` | no |

## UpdatePrivateConnectionCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnectionName` | `string` | yes |
| `certificate` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `status` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## UpdateSecurityRequirementPack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |

## UpdateTargetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | yes |
| `verificationMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | yes |
| `domainName` | `string` | yes |
| `verificationStatus` | `string` | yes |
| `verificationStatusReason` | `string` | no |
| `verificationDetails` | `VerificationDetails` | no |
| `createdAt` | `timestamp` | no |
| `verifiedAt` | `timestamp` | no |

## UpdateThreat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `title` | `string` | no |
| `status` | `string` | no |
| `comments` | `string` | no |
| `statement` | `string` | no |
| `severity` | `string` | no |
| `threatSource` | `string` | no |
| `prerequisites` | `string` | no |
| `threatAction` | `string` | no |
| `threatImpact` | `string` | no |
| `impactedGoal` | `List<string>` | no |
| `impactedAssets` | `List<string>` | no |
| `anchor` | `ThreatAnchorShape` | no |
| `evidence` | `List<ThreatEvidenceShape>` | no |
| `recommendation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatId` | `string` | yes |
| `threatJobId` | `string` | yes |
| `title` | `string` | no |
| `statement` | `string` | no |
| `severity` | `string` | no |
| `status` | `string` | no |
| `comments` | `string` | no |
| `stride` | `List<string>` | no |
| `threatSource` | `string` | no |
| `prerequisites` | `string` | no |
| `threatAction` | `string` | no |
| `threatImpact` | `string` | no |
| `impactedGoal` | `List<string>` | no |
| `impactedAssets` | `List<string>` | no |
| `anchor` | `ThreatAnchorShape` | no |
| `evidence` | `List<ThreatEvidenceShape>` | no |
| `recommendation` | `string` | no |
| `createdBy` | `string` | no |
| `updatedBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## UpdateThreatModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelId` | `string` | yes |
| `agentSpaceId` | `string` | yes |
| `title` | `string` | no |
| `description` | `string` | no |
| `assets` | `Assets` | no |
| `scopeDocs` | `List<DocumentInfo>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `threatModelId` | `string` | yes |
| `title` | `string` | no |
| `agentSpaceId` | `string` | no |
| `description` | `string` | no |
| `assets` | `Assets` | no |
| `scopeDocs` | `List<DocumentInfo>` | no |
| `serviceRole` | `string` | no |
| `logConfig` | `CloudWatchLog` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## VerifyTargetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetDomainId` | `string` | no |
| `domainName` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `verifiedAt` | `timestamp` | no |
| `status` | `string` | no |
| `verificationStatusReason` | `string` | no |

