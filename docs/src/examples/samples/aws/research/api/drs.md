# Elastic Disaster Recovery Service

API version: 2020-02-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/drs/2020-02-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateSourceNetworkStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | yes |
| `cfnStackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## CancelRecoveryPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecution` | `RecoveryPlanExecution` | yes |

## CreateExtendedSourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerArn` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServer` | `SourceServer` | no |

## CreateLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `exportBucketArn` | `string` | no |
| `postLaunchEnabled` | `boolean` | no |
| `launchIntoSourceInstance` | `boolean` | no |
| `recoveryMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplate` | `LaunchConfigurationTemplate` | no |

## CreateRecoveryPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlan` | `RecoveryPlan` | yes |

## CreateRecoveryPlanStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |
| `stepName` | `string` | yes |
| `stepOrder` | `integer` | no |
| `configuration` | `RecoveryPlanStepConfiguration` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStep` | `RecoveryPlanStep` | yes |

## CreateReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stagingAreaSubnetId` | `string` | yes |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | yes |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | yes |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | yes |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | yes |
| `pitPolicy` | `List<PITPolicyRule>` | yes |
| `tags` | `Map<string>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `tags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

## CreateSourceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcID` | `string` | yes |
| `originAccountID` | `string` | yes |
| `originRegion` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | no |

## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLaunchAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `actionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecoveryInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecoveryPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |

## DeleteRecoveryPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionArn` | `string` | yes |

## DeleteRecoveryPlanStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStepArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStepArn` | `string` | yes |

## DeleteReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSourceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeJobLogItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<JobLog>` | no |
| `nextToken` | `string` | no |

## DescribeJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeJobsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<Job>` | no |
| `nextToken` | `string` | no |

## DescribeLaunchConfigurationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateIDs` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<LaunchConfigurationTemplate>` | no |
| `nextToken` | `string` | no |

## DescribeRecoveryInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeRecoveryInstancesRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `items` | `List<RecoveryInstance>` | no |

## DescribeRecoverySnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `filters` | `DescribeRecoverySnapshotsRequestFilters` | no |
| `order` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RecoverySnapshot>` | no |
| `nextToken` | `string` | no |

## DescribeReplicationConfigurationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateIDs` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ReplicationConfigurationTemplate>` | no |
| `nextToken` | `string` | no |

## DescribeSourceNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeSourceNetworksRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SourceNetwork>` | no |
| `nextToken` | `string` | no |

## DescribeSourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `DescribeSourceServersRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SourceServer>` | no |
| `nextToken` | `string` | no |

## DisconnectRecoveryInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisconnectSourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |
| `recoveryInstanceId` | `string` | no |
| `lastLaunchResult` | `string` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `stagingArea` | `StagingArea` | no |
| `sourceCloudProperties` | `SourceCloudProperties` | no |
| `replicationDirection` | `string` | no |
| `reversedDirectionSourceServerArn` | `string` | no |
| `sourceNetworkID` | `string` | no |
| `agentVersion` | `string` | no |

## ExportSourceNetworkCfnTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `s3DestinationUrl` | `string` | no |

## GetFailbackReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |
| `name` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `usePrivateIP` | `boolean` | no |
| `internetProtocol` | `string` | no |

## GetLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `postLaunchEnabled` | `boolean` | no |
| `launchIntoInstanceProperties` | `LaunchIntoInstanceProperties` | no |
| `recoveryMode` | `string` | no |

## GetRecoveryPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlan` | `RecoveryPlan` | yes |

## GetRecoveryPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecution` | `RecoveryPlanExecution` | yes |

## GetRecoveryPlanExecutionStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStepArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStep` | `RecoveryPlanExecutionStep` | yes |

## GetRecoveryPlanStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStepArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStep` | `RecoveryPlanStep` | yes |

## GetReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

## InitializeService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListExtensibleSourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stagingAccountID` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<StagingSourceServer>` | no |
| `nextToken` | `string` | no |

## ListLaunchActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `filters` | `LaunchActionsRequestFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<LaunchAction>` | no |
| `nextToken` | `string` | no |

## ListRecoveryPlanExecutionSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionArn` | `string` | yes |
| `filter` | `ListRecoveryPlanExecutionStepsFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionSteps` | `List<RecoveryPlanExecutionStepSummary>` | yes |
| `nextToken` | `string` | no |

## ListRecoveryPlanExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutions` | `List<RecoveryPlanExecutionSummary>` | yes |
| `nextToken` | `string` | no |

## ListRecoveryPlanSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanSteps` | `List<RecoveryPlanStep>` | yes |
| `nextToken` | `string` | no |

## ListRecoveryPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlans` | `List<RecoveryPlanSummary>` | yes |
| `nextToken` | `string` | no |

## ListStagingAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<Account>` | no |
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

## PutLaunchAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | yes |
| `actionCode` | `string` | yes |
| `order` | `integer` | yes |
| `actionId` | `string` | yes |
| `optional` | `boolean` | yes |
| `active` | `boolean` | yes |
| `name` | `string` | yes |
| `actionVersion` | `string` | yes |
| `category` | `string` | yes |
| `parameters` | `Map<LaunchActionParameter>` | no |
| `description` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceId` | `string` | no |
| `actionId` | `string` | no |
| `actionCode` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | no |
| `active` | `boolean` | no |
| `order` | `integer` | no |
| `actionVersion` | `string` | no |
| `optional` | `boolean` | no |
| `parameters` | `Map<LaunchActionParameter>` | no |
| `description` | `string` | no |
| `category` | `string` | no |

## ReorderRecoveryPlanSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |
| `orderedStepArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanSteps` | `List<RecoveryPlanStep>` | yes |

## RetryDataReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `arn` | `string` | no |
| `tags` | `Map<string>` | no |
| `recoveryInstanceId` | `string` | no |
| `lastLaunchResult` | `string` | no |
| `dataReplicationInfo` | `DataReplicationInfo` | no |
| `lifeCycle` | `LifeCycle` | no |
| `sourceProperties` | `SourceProperties` | no |
| `stagingArea` | `StagingArea` | no |
| `sourceCloudProperties` | `SourceCloudProperties` | no |
| `replicationDirection` | `string` | no |
| `reversedDirectionSourceServerArn` | `string` | no |
| `sourceNetworkID` | `string` | no |
| `agentVersion` | `string` | no |

## RetryRecoveryPlanExecutionStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStepArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStep` | `RecoveryPlanExecutionStep` | yes |

## ReverseReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `reversedDirectionSourceServerArn` | `string` | no |

## StartFailbackLaunch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceIDs` | `List<string>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StartRecovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServers` | `List<StartRecoveryRequestSourceServer>` | yes |
| `isDrill` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StartRecoveryPlanExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |
| `mode` | `string` | yes |
| `clientToken` | `string` | no |
| `sourceServers` | `List<RecoveryPlanExecutionSourceServer>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecution` | `RecoveryPlanExecution` | yes |

## StartReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServer` | `SourceServer` | no |

## StartSourceNetworkRecovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworks` | `List<StartSourceNetworkRecoveryRequestNetworkEntry>` | yes |
| `deployAsNew` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StartSourceNetworkReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetwork` | `SourceNetwork` | no |

## StopFailback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServer` | `SourceServer` | no |

## StopSourceNetworkReplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetworkID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceNetwork` | `SourceNetwork` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateRecoveryInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceIDs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFailbackReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryInstanceID` | `string` | yes |
| `name` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `usePrivateIP` | `boolean` | no |
| `internetProtocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `name` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `postLaunchEnabled` | `boolean` | no |
| `launchIntoInstanceProperties` | `LaunchIntoInstanceProperties` | no |
| `recoveryMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `ec2LaunchTemplateID` | `string` | no |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `postLaunchEnabled` | `boolean` | no |
| `launchIntoInstanceProperties` | `LaunchIntoInstanceProperties` | no |
| `recoveryMode` | `string` | no |

## UpdateLaunchConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplateID` | `string` | yes |
| `launchDisposition` | `string` | no |
| `targetInstanceTypeRightSizingMethod` | `string` | no |
| `copyPrivateIp` | `boolean` | no |
| `copyTags` | `boolean` | no |
| `licensing` | `Licensing` | no |
| `exportBucketArn` | `string` | no |
| `postLaunchEnabled` | `boolean` | no |
| `launchIntoSourceInstance` | `boolean` | no |
| `recoveryMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `launchConfigurationTemplate` | `LaunchConfigurationTemplate` | no |

## UpdateRecoveryPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanArn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlan` | `RecoveryPlan` | yes |

## UpdateRecoveryPlanExecutionStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStepArn` | `string` | yes |
| `status` | `string` | no |
| `servers` | `List<RecoveryPlanServer>` | no |
| `waitDurationMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanExecutionStep` | `RecoveryPlanExecutionStep` | yes |

## UpdateRecoveryPlanStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStepArn` | `string` | yes |
| `stepName` | `string` | no |
| `configuration` | `RecoveryPlanStepConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recoveryPlanStep` | `RecoveryPlanStep` | yes |

## UpdateReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | yes |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceServerID` | `string` | no |
| `name` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `replicatedDisks` | `List<ReplicationConfigurationReplicatedDisk>` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

## UpdateReplicationConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfigurationTemplateID` | `string` | yes |
| `arn` | `string` | no |
| `stagingAreaSubnetId` | `string` | no |
| `associateDefaultSecurityGroup` | `boolean` | no |
| `replicationServersSecurityGroupsIDs` | `List<string>` | no |
| `replicationServerInstanceType` | `string` | no |
| `useDedicatedReplicationServer` | `boolean` | no |
| `defaultLargeStagingDiskType` | `string` | no |
| `ebsEncryption` | `string` | no |
| `ebsEncryptionKeyArn` | `string` | no |
| `bandwidthThrottling` | `long` | no |
| `dataPlaneRouting` | `string` | no |
| `createPublicIP` | `boolean` | no |
| `stagingAreaTags` | `Map<string>` | no |
| `tags` | `Map<string>` | no |
| `pitPolicy` | `List<PITPolicyRule>` | no |
| `autoReplicateNewDisks` | `boolean` | no |
| `internetProtocol` | `string` | no |

