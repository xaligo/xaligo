# Amazon EMR

API version: 2009-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/emr/2009-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddInstanceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `InstanceFleet` | `InstanceFleetConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | no |
| `InstanceFleetId` | `string` | no |
| `ClusterArn` | `string` | no |

## AddInstanceGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceGroups` | `List<InstanceGroupConfig>` | yes |
| `JobFlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowId` | `string` | no |
| `InstanceGroupIds` | `List<string>` | no |
| `ClusterArn` | `string` | no |

## AddJobFlowSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowId` | `string` | yes |
| `Steps` | `List<StepConfig>` | yes |
| `ExecutionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StepIds` | `List<string>` | no |

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |
| `ClusterId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `StepIds` | `List<string>` | yes |
| `StepCancellationOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CancelStepsInfoList` | `List<CancelStepsInfo>` | no |

## CreatePersistentAppUI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetResourceArn` | `string` | yes |
| `EMRContainersConfig` | `EMRContainersConfig` | no |
| `Tags` | `List<Tag>` | no |
| `XReferer` | `string` | no |
| `ProfilerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersistentAppUIId` | `string` | no |
| `RuntimeRoleEnabledCluster` | `boolean` | no |

## CreateSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SecurityConfiguration` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CreationDateTime` | `timestamp` | yes |

## CreateStudio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `AuthMode` | `string` | yes |
| `VpcId` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `ServiceRole` | `string` | yes |
| `UserRole` | `string` | no |
| `WorkspaceSecurityGroupId` | `string` | yes |
| `EngineSecurityGroupId` | `string` | yes |
| `DefaultS3Location` | `string` | yes |
| `IdpAuthUrl` | `string` | no |
| `IdpRelayStateParameterName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `TrustedIdentityPropagationEnabled` | `boolean` | no |
| `IdcUserAssignment` | `string` | no |
| `IdcInstanceArn` | `string` | no |
| `EncryptionKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | no |
| `Url` | `string` | no |

## CreateStudioSessionMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |
| `IdentityId` | `string` | no |
| `IdentityName` | `string` | no |
| `IdentityType` | `string` | yes |
| `SessionPolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStudio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStudioSessionMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |
| `IdentityId` | `string` | no |
| `IdentityName` | `string` | no |
| `IdentityType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cluster` | `Cluster` | no |

## DescribeJobFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `JobFlowIds` | `List<string>` | no |
| `JobFlowStates` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlows` | `List<JobFlowDetail>` | no |

## DescribeNotebookExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookExecution` | `NotebookExecution` | no |

## DescribePersistentAppUI

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersistentAppUIId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersistentAppUI` | `PersistentAppUI` | no |

## DescribeReleaseLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReleaseLabel` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReleaseLabel` | `string` | no |
| `Applications` | `List<SimplifiedApplication>` | no |
| `NextToken` | `string` | no |
| `AvailableOSReleases` | `List<OSRelease>` | no |

## DescribeSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `SecurityConfiguration` | `string` | no |
| `CreationDateTime` | `timestamp` | no |

## DescribeStep

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `StepId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Step` | `Step` | no |

## DescribeStudio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Studio` | `Studio` | no |

## GetAutoTerminationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoTerminationPolicy` | `AutoTerminationPolicy` | no |

## GetBlockPublicAccessConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockPublicAccessConfiguration` | `BlockPublicAccessConfiguration` | yes |
| `BlockPublicAccessConfigurationMetadata` | `BlockPublicAccessConfigurationMetadata` | yes |

## GetClusterSessionCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `ExecutionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `ExpiresAt` | `timestamp` | no |

## GetManagedScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedScalingPolicy` | `ManagedScalingPolicy` | no |

## GetOnClusterAppUIPresignedURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `OnClusterAppUIType` | `string` | no |
| `ApplicationId` | `string` | no |
| `DryRun` | `boolean` | no |
| `ExecutionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PresignedURLReady` | `boolean` | no |
| `PresignedURL` | `string` | no |

## GetPersistentAppUIPresignedURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PersistentAppUIId` | `string` | yes |
| `PersistentAppUIType` | `string` | no |
| `ApplicationId` | `string` | no |
| `AuthProxyCall` | `boolean` | no |
| `ExecutionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PresignedURLReady` | `boolean` | no |
| `PresignedURL` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Session` | `Session` | yes |

## GetSessionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoint` | `string` | yes |
| `AuthToken` | `string` | no |
| `AuthTokenExpirationTime` | `timestamp` | no |
| `Credentials` | `Credentials` | no |

## GetStudioSessionMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |
| `IdentityId` | `string` | no |
| `IdentityName` | `string` | no |
| `IdentityType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionMapping` | `SessionMappingDetail` | no |

## ListBootstrapActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BootstrapActions` | `List<Command>` | no |
| `Marker` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `ClusterStates` | `List<string>` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Clusters` | `List<ClusterSummary>` | no |
| `Marker` | `string` | no |

## ListInstanceFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceFleets` | `List<InstanceFleet>` | no |
| `Marker` | `string` | no |

## ListInstanceGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceGroups` | `List<InstanceGroup>` | no |
| `Marker` | `string` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `InstanceGroupId` | `string` | no |
| `InstanceGroupTypes` | `List<string>` | no |
| `InstanceFleetId` | `string` | no |
| `InstanceFleetType` | `string` | no |
| `InstanceStates` | `List<string>` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<Instance>` | no |
| `Marker` | `string` | no |

## ListNotebookExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EditorId` | `string` | no |
| `Status` | `string` | no |
| `From` | `timestamp` | no |
| `To` | `timestamp` | no |
| `Marker` | `string` | no |
| `ExecutionEngineId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookExecutions` | `List<NotebookExecutionSummary>` | no |
| `Marker` | `string` | no |

## ListReleaseLabels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ReleaseLabelFilter` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReleaseLabels` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListSecurityConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityConfigurations` | `List<SecurityConfigurationSummary>` | no |
| `Marker` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SessionStates` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sessions` | `List<Session>` | no |
| `NextToken` | `string` | no |

## ListSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `StepStates` | `List<string>` | no |
| `StepIds` | `List<string>` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Steps` | `List<StepSummary>` | no |
| `Marker` | `string` | no |

## ListStudioSessionMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | no |
| `IdentityType` | `string` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionMappings` | `List<SessionMappingSummary>` | no |
| `Marker` | `string` | no |

## ListStudios

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Studios` | `List<StudioSummary>` | no |
| `Marker` | `string` | no |

## ListSupportedInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReleaseLabel` | `string` | yes |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SupportedInstanceTypes` | `List<SupportedInstanceType>` | no |
| `Marker` | `string` | no |

## ModifyCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `StepConcurrencyLevel` | `integer` | no |
| `ExtendedSupport` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StepConcurrencyLevel` | `integer` | no |
| `ExtendedSupport` | `boolean` | no |

## ModifyInstanceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `InstanceFleet` | `InstanceFleetModifyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyInstanceGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | no |
| `InstanceGroups` | `List<InstanceGroupModifyConfig>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAutoScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `InstanceGroupId` | `string` | yes |
| `AutoScalingPolicy` | `AutoScalingPolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | no |
| `InstanceGroupId` | `string` | no |
| `AutoScalingPolicy` | `AutoScalingPolicyDescription` | no |
| `ClusterArn` | `string` | no |

## PutAutoTerminationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `AutoTerminationPolicy` | `AutoTerminationPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutBlockPublicAccessConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockPublicAccessConfiguration` | `BlockPublicAccessConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutManagedScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `ManagedScalingPolicy` | `ManagedScalingPolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveAutoScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `InstanceGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveAutoTerminationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveManagedScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |
| `ClusterId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RunJobFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `LogUri` | `string` | no |
| `LogEncryptionKmsKeyId` | `string` | no |
| `AdditionalInfo` | `string` | no |
| `AmiVersion` | `string` | no |
| `ReleaseLabel` | `string` | no |
| `Instances` | `JobFlowInstancesConfig` | yes |
| `Steps` | `List<StepConfig>` | no |
| `StepExecutionRoleArn` | `string` | no |
| `BootstrapActions` | `List<BootstrapActionConfig>` | no |
| `SupportedProducts` | `List<string>` | no |
| `NewSupportedProducts` | `List<SupportedProductConfig>` | no |
| `Applications` | `List<Application>` | no |
| `Configurations` | `List<Configuration>` | no |
| `VisibleToAllUsers` | `boolean` | no |
| `JobFlowRole` | `string` | no |
| `ServiceRole` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SecurityConfiguration` | `string` | no |
| `AutoScalingRole` | `string` | no |
| `ScaleDownBehavior` | `string` | no |
| `CustomAmiId` | `string` | no |
| `EbsRootVolumeSize` | `integer` | no |
| `RepoUpgradeOnBoot` | `string` | no |
| `KerberosAttributes` | `KerberosAttributes` | no |
| `StepConcurrencyLevel` | `integer` | no |
| `ManagedScalingPolicy` | `ManagedScalingPolicy` | no |
| `PlacementGroupConfigs` | `List<PlacementGroupConfig>` | no |
| `AutoTerminationPolicy` | `AutoTerminationPolicy` | no |
| `OSReleaseLabel` | `string` | no |
| `EbsRootVolumeIops` | `integer` | no |
| `EbsRootVolumeThroughput` | `integer` | no |
| `ExtendedSupport` | `boolean` | no |
| `MonitoringConfiguration` | `MonitoringConfiguration` | no |
| `SessionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowId` | `string` | no |
| `ClusterArn` | `string` | no |

## SetKeepJobFlowAliveWhenNoSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowIds` | `List<string>` | yes |
| `KeepJobFlowAliveWhenNoSteps` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetTerminationProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowIds` | `List<string>` | yes |
| `TerminationProtected` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetUnhealthyNodeReplacement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowIds` | `List<string>` | yes |
| `UnhealthyNodeReplacement` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetVisibleToAllUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowIds` | `List<string>` | yes |
| `VisibleToAllUsers` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartNotebookExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EditorId` | `string` | no |
| `RelativePath` | `string` | no |
| `NotebookExecutionName` | `string` | no |
| `NotebookParams` | `string` | no |
| `ExecutionEngine` | `ExecutionEngineConfig` | yes |
| `ServiceRole` | `string` | yes |
| `NotebookInstanceSecurityGroupId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `NotebookS3Location` | `NotebookS3LocationFromInput` | no |
| `OutputNotebookS3Location` | `OutputNotebookS3LocationFromInput` | no |
| `OutputNotebookFormat` | `string` | no |
| `EnvironmentVariables` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookExecutionId` | `string` | no |

## StartSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ClusterId` | `string` | yes |
| `ExecutionRoleArn` | `string` | no |
| `EngineConfigurations` | `List<Configuration>` | no |
| `MonitoringConfiguration` | `SessionMonitoringConfiguration` | no |
| `SessionIdleTimeoutInMinutes` | `long` | no |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `ClusterId` | `string` | no |
| `Arn` | `string` | no |
| `AccountId` | `string` | no |
| `State` | `string` | no |

## StopNotebookExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateJobFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobFlowIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `SessionId` | `string` | yes |
| `State` | `string` | yes |

## UpdateStudio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `DefaultS3Location` | `string` | no |
| `EncryptionKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStudioSessionMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioId` | `string` | yes |
| `IdentityId` | `string` | no |
| `IdentityName` | `string` | no |
| `IdentityType` | `string` | yes |
| `SessionPolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


