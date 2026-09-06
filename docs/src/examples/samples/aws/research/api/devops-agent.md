# AWS DevOps Agent Service

API version: 2026-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/devops-agent/2026-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `serviceId` | `string` | yes |
| `configuration` | `ServiceConfiguration` | yes |
| `capabilities` | `Map<CapabilityConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `Association` | yes |
| `webhook` | `GenericWebhook` | no |

## CreateAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `locale` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `preferences` | `Map<boolean>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpace` | `AgentSpace` | yes |
| `tags` | `Map<string>` | no |

## CreateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetType` | `string` | yes |
| `metadata` | `Document` | no |
| `content` | `AssetContent` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `asset` | `Asset` | yes |

## CreateAssetFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `path` | `string` | yes |
| `content` | `AssetFileBody` | yes |
| `metadata` | `Document` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `file` | `AssetFile` | yes |

## CreateBacklogTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `reference` | `ReferenceInput` | no |
| `taskType` | `string` | yes |
| `title` | `string` | yes |
| `description` | `string` | no |
| `priority` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `Task` | yes |

## CreateChat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `userId` | `string` | no |
| `userType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionId` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreatePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `mode` | `PrivateConnectionMode` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `status` | `string` | yes |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `type` | `string` | yes |
| `condition` | `TriggerCondition` | yes |
| `action` | `TriggerAction` | yes |
| `status` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trigger` | `Trigger` | yes |

## DeleteAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssetFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `path` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `status` | `string` | yes |

## DeleteTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `triggerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribePrivateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `status` | `string` | yes |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## DisableOperatorApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `authFlow` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableOperatorApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `authFlow` | `string` | yes |
| `operatorAppRoleArn` | `string` | yes |
| `idcInstanceArn` | `string` | no |
| `issuerUrl` | `string` | no |
| `idpClientId` | `string` | no |
| `idpClientSecret` | `string` | no |
| `provider` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `operatorAppUrl` | `string` | no |
| `iam` | `IamAuthConfiguration` | no |
| `idc` | `IdcAuthConfiguration` | no |
| `idp` | `IdpAuthConfiguration` | no |

## GetAccountUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `monthlyAccountInvestigationHours` | `UsageMetric` | no |
| `monthlyAccountEvaluationHours` | `UsageMetric` | no |
| `monthlyAccountSystemLearningHours` | `UsageMetric` | no |
| `monthlyAccountOnDemandHours` | `UsageMetric` | no |
| `usagePeriodStartTime` | `timestamp` | yes |
| `usagePeriodEndTime` | `timestamp` | yes |

## GetAgentSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpace` | `AgentSpace` | yes |
| `tags` | `Map<string>` | no |

## GetAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `assetVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `asset` | `Asset` | yes |

## GetAssetContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `assetVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `content` | `AssetZipContent` | yes |
| `version` | `integer` | yes |

## GetAssetFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `path` | `string` | yes |
| `assetVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `file` | `AssetFile` | yes |

## GetAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `Association` | yes |

## GetBacklogTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `taskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `Task` | yes |

## GetOperatorApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `operatorAppUrl` | `string` | no |
| `iam` | `IamAuthConfiguration` | no |
| `idc` | `IdcAuthConfiguration` | no |
| `idp` | `IdpAuthConfiguration` | no |

## GetRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `recommendationId` | `string` | yes |
| `recommendationVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendation` | `Recommendation` | yes |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `RegisteredService` | yes |
| `tags` | `Map<string>` | no |

## GetTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `triggerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trigger` | `Trigger` | yes |

## ListAgentSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `agentSpaces` | `List<AgentSpace>` | yes |

## ListAssetFiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `assetVersion` | `integer` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AssetFileSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AssetTypeSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AssetVersionMetadata>` | yes |
| `nextToken` | `string` | no |

## ListAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetType` | `string` | no |
| `updatedAfter` | `timestamp` | no |
| `updatedBefore` | `timestamp` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Asset>` | yes |
| `nextToken` | `string` | no |

## ListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filterServiceTypes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `associations` | `List<Association>` | yes |

## ListBacklogTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `filter` | `TaskFilter` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |
| `sortField` | `string` | no |
| `order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<Task>` | yes |
| `nextToken` | `string` | no |

## ListChats

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `userId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executions` | `List<ChatExecution>` | yes |
| `nextToken` | `string` | no |

## ListExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `taskId` | `string` | yes |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executions` | `List<Execution>` | yes |
| `nextToken` | `string` | no |

## ListGoals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `status` | `string` | no |
| `goalType` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `goals` | `List<Goal>` | yes |
| `nextToken` | `string` | no |

## ListJournalRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `executionId` | `string` | yes |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |
| `recordType` | `string` | no |
| `order` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `records` | `List<JournalRecord>` | yes |
| `nextToken` | `string` | no |

## ListPendingMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `executionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `executionId` | `string` | yes |
| `messages` | `List<PendingMessage>` | yes |
| `createdAt` | `timestamp` | yes |

## ListPrivateConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `privateConnections` | `List<PrivateConnectionSummary>` | yes |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `taskId` | `string` | no |
| `goalId` | `string` | no |
| `status` | `string` | no |
| `priority` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendations` | `List<Recommendation>` | yes |
| `nextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filterServiceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `services` | `List<RegisteredService>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## ListTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Trigger>` | yes |
| `nextToken` | `string` | no |

## ListWebhooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `webhooks` | `List<Webhook>` | yes |

## RegisterService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `string` | yes |
| `serviceDetails` | `ServiceDetails` | yes |
| `kmsKeyArn` | `string` | no |
| `privateConnectionName` | `string` | no |
| `targetUrlPrivateConnectionName` | `string` | no |
| `exchangeUrlPrivateConnectionName` | `string` | no |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceId` | `string` | no |
| `additionalStep` | `AdditionalServiceRegistrationStep` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

## SendMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `executionId` | `string` | yes |
| `content` | `string` | yes |
| `context` | `SendMessageContext` | no |
| `userId` | `string` | no |
| `assetIds` | `List<string>` | no |
| `modelTier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `SendMessageEvents` | yes |

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
| `locale` | `string` | no |
| `preferences` | `Map<boolean>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpace` | `AgentSpace` | yes |

## UpdateApprovalAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `approvalId` | `string` | yes |
| `action` | `string` | yes |
| `finalPattern` | `ApprovalPattern` | no |
| `reason` | `string` | no |
| `ttlSeconds` | `integer` | no |
| `singleUse` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `approvalId` | `string` | yes |
| `status` | `string` | yes |
| `expiresAt` | `timestamp` | no |

## UpdateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `metadata` | `Document` | no |
| `content` | `AssetContent` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `asset` | `Asset` | yes |

## UpdateAssetFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `assetId` | `string` | yes |
| `path` | `string` | yes |
| `content` | `AssetFileBody` | no |
| `metadata` | `Document` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `file` | `AssetFile` | yes |

## UpdateAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `associationId` | `string` | yes |
| `configuration` | `ServiceConfiguration` | yes |
| `capabilities` | `Map<CapabilityConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `Association` | yes |
| `webhook` | `GenericWebhook` | no |

## UpdateBacklogTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `taskId` | `string` | yes |
| `taskStatus` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `Task` | yes |

## UpdateGoal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `goalId` | `string` | yes |
| `evaluationSchedule` | `GoalScheduleInput` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `goal` | `Goal` | yes |

## UpdateOperatorAppIdpConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `idpClientSecret` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `idp` | `IdpAuthConfiguration` | yes |

## UpdatePrivateConnectionCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `certificate` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `resourceGatewayId` | `string` | no |
| `hostAddress` | `string` | no |
| `vpcId` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `status` | `string` | yes |
| `certificateExpiryTime` | `timestamp` | no |
| `dnsResolution` | `string` | no |
| `failureMessage` | `string` | no |

## UpdateRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `recommendationId` | `string` | yes |
| `status` | `string` | no |
| `additionalContext` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recommendation` | `Recommendation` | yes |

## UpdateTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |
| `triggerId` | `string` | yes |
| `status` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trigger` | `Trigger` | yes |

## ValidateAwsAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentSpaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


