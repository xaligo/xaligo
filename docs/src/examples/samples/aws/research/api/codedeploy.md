# AWS CodeDeploy

API version: 2014-10-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codedeploy/2014-10-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToOnPremisesInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | yes |
| `instanceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetApplicationRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `revisions` | `List<RevisionLocation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `errorMessage` | `string` | no |
| `revisions` | `List<RevisionInfo>` | no |

## BatchGetApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationsInfo` | `List<ApplicationInfo>` | no |

## BatchGetDeploymentGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `deploymentGroupNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentGroupsInfo` | `List<DeploymentGroupInfo>` | no |
| `errorMessage` | `string` | no |

## BatchGetDeploymentInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `instanceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instancesSummary` | `List<InstanceSummary>` | no |
| `errorMessage` | `string` | no |

## BatchGetDeploymentTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `targetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentTargets` | `List<DeploymentTarget>` | no |

## BatchGetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentsInfo` | `List<DeploymentInfo>` | no |

## BatchGetOnPremisesInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceInfos` | `List<InstanceInfo>` | no |

## ContinueDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `deploymentWaitType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `computePlatform` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | no |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `deploymentGroupName` | `string` | no |
| `revision` | `RevisionLocation` | no |
| `deploymentConfigName` | `string` | no |
| `description` | `string` | no |
| `ignoreApplicationStopFailures` | `boolean` | no |
| `targetInstances` | `TargetInstances` | no |
| `autoRollbackConfiguration` | `AutoRollbackConfiguration` | no |
| `updateOutdatedInstancesOnly` | `boolean` | no |
| `fileExistsBehavior` | `string` | no |
| `deploymentMode` | `string` | no |
| `overrideAlarmConfiguration` | `AlarmConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |

## CreateDeploymentConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigName` | `string` | yes |
| `minimumHealthyHosts` | `MinimumHealthyHosts` | no |
| `trafficRoutingConfig` | `TrafficRoutingConfig` | no |
| `computePlatform` | `string` | no |
| `zonalConfig` | `ZonalConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigId` | `string` | no |

## CreateDeploymentGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `deploymentGroupName` | `string` | yes |
| `deploymentConfigName` | `string` | no |
| `ec2TagFilters` | `List<EC2TagFilter>` | no |
| `onPremisesInstanceTagFilters` | `List<TagFilter>` | no |
| `autoScalingGroups` | `List<string>` | no |
| `serviceRoleArn` | `string` | yes |
| `triggerConfigurations` | `List<TriggerConfig>` | no |
| `alarmConfiguration` | `AlarmConfiguration` | no |
| `autoRollbackConfiguration` | `AutoRollbackConfiguration` | no |
| `outdatedInstancesStrategy` | `string` | no |
| `deploymentStyle` | `DeploymentStyle` | no |
| `blueGreenDeploymentConfiguration` | `BlueGreenDeploymentConfiguration` | no |
| `loadBalancerInfo` | `LoadBalancerInfo` | no |
| `ec2TagSet` | `EC2TagSet` | no |
| `ecsServices` | `List<ECSService>` | no |
| `onPremisesTagSet` | `OnPremisesTagSet` | no |
| `tags` | `List<Tag>` | no |
| `terminationHookEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentGroupId` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeploymentConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeploymentGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `deploymentGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hooksNotCleanedUp` | `List<AutoScalingGroup>` | no |

## DeleteGitHubAccountToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenName` | `string` | no |

## DeleteResourcesByExternalId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `externalId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterOnPremisesInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `ApplicationInfo` | no |

## GetApplicationRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `revision` | `RevisionLocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `revision` | `RevisionLocation` | no |
| `revisionInfo` | `GenericRevisionInfo` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentInfo` | `DeploymentInfo` | no |

## GetDeploymentConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigInfo` | `DeploymentConfigInfo` | no |

## GetDeploymentGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `deploymentGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentGroupInfo` | `DeploymentGroupInfo` | no |

## GetDeploymentInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `instanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceSummary` | `InstanceSummary` | no |

## GetDeploymentTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `targetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentTarget` | `DeploymentTarget` | no |

## GetOnPremisesInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceInfo` | `InstanceInfo` | no |

## ListApplicationRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `s3Bucket` | `string` | no |
| `s3KeyPrefix` | `string` | no |
| `deployed` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `revisions` | `List<RevisionLocation>` | no |
| `nextToken` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applications` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDeploymentConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentConfigsList` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDeploymentGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `deploymentGroups` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDeploymentInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `nextToken` | `string` | no |
| `instanceStatusFilter` | `List<string>` | no |
| `instanceTypeFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instancesList` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDeploymentTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `nextToken` | `string` | no |
| `targetFilters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetIds` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `deploymentGroupName` | `string` | no |
| `externalId` | `string` | no |
| `includeOnlyStatuses` | `List<string>` | no |
| `createTimeRange` | `TimeRange` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployments` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListGitHubAccountTokenNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenNameList` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListOnPremisesInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registrationStatus` | `string` | no |
| `tagFilters` | `List<TagFilter>` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## PutLifecycleEventHookExecutionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `lifecycleEventHookExecutionId` | `string` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleEventHookExecutionId` | `string` | no |

## RegisterApplicationRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `description` | `string` | no |
| `revision` | `RevisionLocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterOnPremisesInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | yes |
| `iamSessionArn` | `string` | no |
| `iamUserArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTagsFromOnPremisesInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | yes |
| `instanceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SkipWaitTimeForInstanceTermination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | yes |
| `autoRollbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `statusMessage` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `newApplicationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDeploymentGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | yes |
| `currentDeploymentGroupName` | `string` | yes |
| `newDeploymentGroupName` | `string` | no |
| `deploymentConfigName` | `string` | no |
| `ec2TagFilters` | `List<EC2TagFilter>` | no |
| `onPremisesInstanceTagFilters` | `List<TagFilter>` | no |
| `autoScalingGroups` | `List<string>` | no |
| `serviceRoleArn` | `string` | no |
| `triggerConfigurations` | `List<TriggerConfig>` | no |
| `alarmConfiguration` | `AlarmConfiguration` | no |
| `autoRollbackConfiguration` | `AutoRollbackConfiguration` | no |
| `outdatedInstancesStrategy` | `string` | no |
| `deploymentStyle` | `DeploymentStyle` | no |
| `blueGreenDeploymentConfiguration` | `BlueGreenDeploymentConfiguration` | no |
| `loadBalancerInfo` | `LoadBalancerInfo` | no |
| `ec2TagSet` | `EC2TagSet` | no |
| `ecsServices` | `List<ECSService>` | no |
| `onPremisesTagSet` | `OnPremisesTagSet` | no |
| `terminationHookEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hooksNotCleanedUp` | `List<AutoScalingGroup>` | no |

