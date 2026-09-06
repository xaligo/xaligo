# Amazon EC2 Container Service

API version: 2014-11-13. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ecs/2014-11-13/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ContinueServiceDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeploymentArn` | `string` | yes |
| `hookId` | `string` | yes |
| `action` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeploymentArn` | `string` | no |

## CreateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `cluster` | `string` | no |
| `autoScalingGroupProvider` | `AutoScalingGroupProvider` | no |
| `managedInstancesProvider` | `CreateManagedInstancesProviderConfiguration` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProvider` | `CapacityProvider` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | no |
| `tags` | `List<Tag>` | no |
| `settings` | `List<ClusterSetting>` | no |
| `configuration` | `ClusterConfiguration` | no |
| `capacityProviders` | `List<string>` | no |
| `defaultCapacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | no |
| `serviceConnectDefaults` | `ClusterServiceConnectDefaultsRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## CreateDaemon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonName` | `string` | yes |
| `clusterArn` | `string` | no |
| `daemonTaskDefinitionArn` | `string` | yes |
| `capacityProviderArns` | `List<string>` | yes |
| `deploymentConfiguration` | `DaemonDeploymentConfiguration` | no |
| `tags` | `List<Tag>` | no |
| `propagateTags` | `string` | no |
| `enableECSManagedTags` | `boolean` | no |
| `enableExecuteCommand` | `boolean` | no |
| `clientToken` | `string` | no |
| `critical` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `deploymentArn` | `string` | no |

## CreateExpressGatewayService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionRoleArn` | `string` | no |
| `infrastructureRoleArn` | `string` | yes |
| `serviceName` | `string` | no |
| `cluster` | `string` | no |
| `healthCheckPath` | `string` | no |
| `primaryContainer` | `ExpressGatewayContainer` | no |
| `taskRoleArn` | `string` | no |
| `networkConfiguration` | `ExpressGatewayServiceNetworkConfiguration` | no |
| `cpu` | `string` | no |
| `memory` | `string` | no |
| `scalingTarget` | `ExpressGatewayScalingTarget` | no |
| `tags` | `List<Tag>` | no |
| `taskDefinitionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `ECSExpressGatewayService` | no |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `serviceName` | `string` | yes |
| `taskDefinition` | `string` | no |
| `availabilityZoneRebalancing` | `string` | no |
| `loadBalancers` | `List<LoadBalancer>` | no |
| `serviceRegistries` | `List<ServiceRegistry>` | no |
| `desiredCount` | `integer` | no |
| `clientToken` | `string` | no |
| `launchType` | `string` | no |
| `capacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | no |
| `platformVersion` | `string` | no |
| `role` | `string` | no |
| `deploymentConfiguration` | `DeploymentConfiguration` | no |
| `placementConstraints` | `List<PlacementConstraint>` | no |
| `placementStrategy` | `List<PlacementStrategy>` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `healthCheckGracePeriodSeconds` | `integer` | no |
| `schedulingStrategy` | `string` | no |
| `deploymentController` | `DeploymentController` | no |
| `tags` | `List<Tag>` | no |
| `enableECSManagedTags` | `boolean` | no |
| `propagateTags` | `string` | no |
| `enableExecuteCommand` | `boolean` | no |
| `serviceConnectConfiguration` | `ServiceConnectConfiguration` | no |
| `volumeConfigurations` | `List<ServiceVolumeConfiguration>` | no |
| `vpcLatticeConfigurations` | `List<VpcLatticeConfiguration>` | no |
| `monitoring` | `MonitoringConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | no |

## CreateTaskSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `string` | yes |
| `cluster` | `string` | yes |
| `externalId` | `string` | no |
| `taskDefinition` | `string` | yes |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `loadBalancers` | `List<LoadBalancer>` | no |
| `serviceRegistries` | `List<ServiceRegistry>` | no |
| `launchType` | `string` | no |
| `capacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | no |
| `platformVersion` | `string` | no |
| `scale` | `Scale` | no |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSet` | `TaskSet` | no |

## DeleteAccountSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `principalArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `setting` | `Setting` | no |

## DeleteAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `attributes` | `List<Attribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `List<Attribute>` | no |

## DeleteCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProvider` | `string` | yes |
| `cluster` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProvider` | `CapacityProvider` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## DeleteDaemon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `deploymentArn` | `string` | no |

## DeleteDaemonTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinitionArn` | `string` | no |

## DeleteExpressGatewayService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `ECSExpressGatewayService` | no |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `service` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | no |

## DeleteTaskDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinitions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinitions` | `List<TaskDefinition>` | no |
| `failures` | `List<Failure>` | no |

## DeleteTaskSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `service` | `string` | yes |
| `taskSet` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSet` | `TaskSet` | no |

## DeregisterContainerInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstance` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstance` | `ContainerInstance` | no |

## DeregisterTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinition` | `TaskDefinition` | no |

## DescribeCapacityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviders` | `List<string>` | no |
| `cluster` | `string` | no |
| `include` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviders` | `List<CapacityProvider>` | no |
| `failures` | `List<Failure>` | no |
| `nextToken` | `string` | no |

## DescribeClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusters` | `List<string>` | no |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusters` | `List<Cluster>` | no |
| `failures` | `List<Failure>` | no |

## DescribeContainerInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstances` | `List<string>` | yes |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstances` | `List<ContainerInstance>` | no |
| `failures` | `List<Failure>` | no |

## DescribeDaemon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemon` | `DaemonDetail` | no |

## DescribeDaemonDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonDeploymentArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failures` | `List<Failure>` | no |
| `daemonDeployments` | `List<DaemonDeployment>` | no |

## DescribeDaemonRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonRevisionArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonRevisions` | `List<DaemonRevision>` | no |
| `failures` | `List<Failure>` | no |

## DescribeDaemonTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinition` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinition` | `DaemonTaskDefinition` | no |

## DescribeExpressGatewayService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `ECSExpressGatewayService` | no |

## DescribeServiceDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeploymentArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeployments` | `List<ServiceDeployment>` | no |
| `failures` | `List<Failure>` | no |

## DescribeServiceRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceRevisionArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceRevisions` | `List<ServiceRevision>` | no |
| `failures` | `List<Failure>` | no |

## DescribeServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `services` | `List<string>` | yes |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `services` | `List<Service>` | no |
| `failures` | `List<Failure>` | no |

## DescribeTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinition` | `string` | yes |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinition` | `TaskDefinition` | no |
| `tags` | `List<Tag>` | no |

## DescribeTaskSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `service` | `string` | yes |
| `taskSets` | `List<string>` | no |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSets` | `List<TaskSet>` | no |
| `failures` | `List<Failure>` | no |

## DescribeTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `tasks` | `List<string>` | yes |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<Task>` | no |
| `failures` | `List<Failure>` | no |

## DiscoverPollEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstance` | `string` | no |
| `cluster` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `string` | no |
| `telemetryEndpoint` | `string` | no |
| `serviceConnectEndpoint` | `string` | no |

## ExecuteCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `container` | `string` | no |
| `command` | `string` | yes |
| `interactive` | `boolean` | yes |
| `task` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | no |
| `containerArn` | `string` | no |
| `containerName` | `string` | no |
| `interactive` | `boolean` | no |
| `session` | `Session` | no |
| `taskArn` | `string` | no |

## GetTaskProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `tasks` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedTasks` | `List<ProtectedTask>` | no |
| `failures` | `List<Failure>` | no |

## ListAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `value` | `string` | no |
| `principalArn` | `string` | no |
| `effectiveSettings` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<Setting>` | no |
| `nextToken` | `string` | no |

## ListAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `targetType` | `string` | yes |
| `attributeName` | `string` | no |
| `attributeValue` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `List<Attribute>` | no |
| `nextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListContainerInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `filter` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstanceArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListDaemonDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | yes |
| `status` | `List<string>` | no |
| `createdAt` | `CreatedAt` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `daemonDeployments` | `List<DaemonDeploymentSummary>` | no |

## ListDaemonTaskDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `familyPrefix` | `string` | no |
| `family` | `string` | no |
| `revision` | `string` | no |
| `status` | `string` | no |
| `sort` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinitions` | `List<DaemonTaskDefinitionSummary>` | no |
| `nextToken` | `string` | no |

## ListDaemons

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterArn` | `string` | no |
| `capacityProviderArns` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonSummariesList` | `List<DaemonSummary>` | no |
| `nextToken` | `string` | no |

## ListServiceDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `string` | yes |
| `cluster` | `string` | no |
| `status` | `List<string>` | no |
| `createdAt` | `CreatedAt` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeployments` | `List<ServiceDeploymentBrief>` | no |
| `nextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `launchType` | `string` | no |
| `schedulingStrategy` | `string` | no |
| `resourceManagementType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListServicesByNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListTaskDefinitionFamilies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `familyPrefix` | `string` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `families` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTaskDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `familyPrefix` | `string` | no |
| `status` | `string` | no |
| `sort` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinitionArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstance` | `string` | no |
| `family` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `startedBy` | `string` | no |
| `serviceName` | `string` | no |
| `desiredStatus` | `string` | no |
| `launchType` | `string` | no |
| `daemonName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskArns` | `List<string>` | no |
| `nextToken` | `string` | no |

## PutAccountSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `value` | `string` | yes |
| `principalArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `setting` | `Setting` | no |

## PutAccountSettingDefault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `value` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `setting` | `Setting` | no |

## PutAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `attributes` | `List<Attribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `List<Attribute>` | no |

## PutClusterCapacityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `capacityProviders` | `List<string>` | yes |
| `defaultCapacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## RegisterContainerInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `instanceIdentityDocument` | `string` | no |
| `instanceIdentityDocumentSignature` | `string` | no |
| `totalResources` | `List<Resource>` | no |
| `versionInfo` | `VersionInfo` | no |
| `containerInstanceArn` | `string` | no |
| `attributes` | `List<Attribute>` | no |
| `platformDevices` | `List<PlatformDevice>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstance` | `ContainerInstance` | no |

## RegisterDaemonTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `family` | `string` | yes |
| `taskRoleArn` | `string` | no |
| `executionRoleArn` | `string` | no |
| `containerDefinitions` | `List<DaemonContainerDefinition>` | yes |
| `cpu` | `string` | no |
| `memory` | `string` | no |
| `volumes` | `List<DaemonVolume>` | no |
| `tags` | `List<Tag>` | no |
| `pidMode` | `string` | no |
| `ipcMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonTaskDefinitionArn` | `string` | no |

## RegisterTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `family` | `string` | yes |
| `taskRoleArn` | `string` | no |
| `executionRoleArn` | `string` | no |
| `networkMode` | `string` | no |
| `containerDefinitions` | `List<ContainerDefinition>` | yes |
| `volumes` | `List<Volume>` | no |
| `placementConstraints` | `List<TaskDefinitionPlacementConstraint>` | no |
| `requiresCompatibilities` | `List<string>` | no |
| `cpu` | `string` | no |
| `memory` | `string` | no |
| `tags` | `List<Tag>` | no |
| `pidMode` | `string` | no |
| `ipcMode` | `string` | no |
| `proxyConfiguration` | `ProxyConfiguration` | no |
| `inferenceAccelerators` | `List<InferenceAccelerator>` | no |
| `ephemeralStorage` | `EphemeralStorage` | no |
| `runtimePlatform` | `RuntimePlatform` | no |
| `enableFaultInjection` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskDefinition` | `TaskDefinition` | no |
| `tags` | `List<Tag>` | no |

## RunTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | no |
| `cluster` | `string` | no |
| `count` | `integer` | no |
| `enableECSManagedTags` | `boolean` | no |
| `enableExecuteCommand` | `boolean` | no |
| `group` | `string` | no |
| `launchType` | `string` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `overrides` | `TaskOverride` | no |
| `placementConstraints` | `List<PlacementConstraint>` | no |
| `placementStrategy` | `List<PlacementStrategy>` | no |
| `platformVersion` | `string` | no |
| `propagateTags` | `string` | no |
| `referenceId` | `string` | no |
| `startedBy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `taskDefinition` | `string` | yes |
| `clientToken` | `string` | no |
| `volumeConfigurations` | `List<TaskVolumeConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<Task>` | no |
| `failures` | `List<Failure>` | no |

## StartTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstances` | `List<string>` | yes |
| `enableECSManagedTags` | `boolean` | no |
| `enableExecuteCommand` | `boolean` | no |
| `group` | `string` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `overrides` | `TaskOverride` | no |
| `propagateTags` | `string` | no |
| `referenceId` | `string` | no |
| `startedBy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `taskDefinition` | `string` | yes |
| `volumeConfigurations` | `List<TaskVolumeConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tasks` | `List<Task>` | no |
| `failures` | `List<Failure>` | no |

## StopServiceDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeploymentArn` | `string` | yes |
| `stopType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceDeploymentArn` | `string` | no |

## StopTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `task` | `string` | yes |
| `reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `task` | `Task` | no |

## SubmitAttachmentStateChanges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `attachments` | `List<AttachmentStateChange>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acknowledgment` | `string` | no |

## SubmitContainerStateChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `task` | `string` | no |
| `containerName` | `string` | no |
| `runtimeId` | `string` | no |
| `status` | `string` | no |
| `exitCode` | `integer` | no |
| `reason` | `string` | no |
| `networkBindings` | `List<NetworkBinding>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acknowledgment` | `string` | no |

## SubmitTaskStateChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `task` | `string` | no |
| `status` | `string` | no |
| `reason` | `string` | no |
| `containers` | `List<ContainerStateChange>` | no |
| `attachments` | `List<AttachmentStateChange>` | no |
| `managedAgents` | `List<ManagedAgentStateChange>` | no |
| `pullStartedAt` | `timestamp` | no |
| `pullStoppedAt` | `timestamp` | no |
| `executionStoppedAt` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `acknowledgment` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `cluster` | `string` | no |
| `autoScalingGroupProvider` | `AutoScalingGroupProviderUpdate` | no |
| `managedInstancesProvider` | `UpdateManagedInstancesProviderConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProvider` | `CapacityProvider` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `settings` | `List<ClusterSetting>` | no |
| `configuration` | `ClusterConfiguration` | no |
| `serviceConnectDefaults` | `ClusterServiceConnectDefaultsRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## UpdateClusterSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `settings` | `List<ClusterSetting>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## UpdateContainerAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstance` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstance` | `ContainerInstance` | no |

## UpdateContainerInstancesState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `containerInstances` | `List<string>` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerInstances` | `List<ContainerInstance>` | no |
| `failures` | `List<Failure>` | no |

## UpdateDaemon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | yes |
| `daemonTaskDefinitionArn` | `string` | yes |
| `capacityProviderArns` | `List<string>` | yes |
| `deploymentConfiguration` | `DaemonDeploymentConfiguration` | no |
| `propagateTags` | `string` | no |
| `enableECSManagedTags` | `boolean` | no |
| `enableExecuteCommand` | `boolean` | no |
| `critical` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `daemonArn` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `deploymentArn` | `string` | no |

## UpdateExpressGatewayService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceArn` | `string` | yes |
| `executionRoleArn` | `string` | no |
| `healthCheckPath` | `string` | no |
| `primaryContainer` | `ExpressGatewayContainer` | no |
| `taskRoleArn` | `string` | no |
| `networkConfiguration` | `ExpressGatewayServiceNetworkConfiguration` | no |
| `cpu` | `string` | no |
| `memory` | `string` | no |
| `scalingTarget` | `ExpressGatewayScalingTarget` | no |
| `taskDefinitionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `UpdatedExpressGatewayService` | no |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | no |
| `service` | `string` | yes |
| `desiredCount` | `integer` | no |
| `taskDefinition` | `string` | no |
| `capacityProviderStrategy` | `List<CapacityProviderStrategyItem>` | no |
| `deploymentConfiguration` | `DeploymentConfiguration` | no |
| `availabilityZoneRebalancing` | `string` | no |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `placementConstraints` | `List<PlacementConstraint>` | no |
| `placementStrategy` | `List<PlacementStrategy>` | no |
| `platformVersion` | `string` | no |
| `forceNewDeployment` | `boolean` | no |
| `healthCheckGracePeriodSeconds` | `integer` | no |
| `deploymentController` | `DeploymentController` | no |
| `enableExecuteCommand` | `boolean` | no |
| `enableECSManagedTags` | `boolean` | no |
| `loadBalancers` | `List<LoadBalancer>` | no |
| `propagateTags` | `string` | no |
| `serviceRegistries` | `List<ServiceRegistry>` | no |
| `serviceConnectConfiguration` | `ServiceConnectConfiguration` | no |
| `volumeConfigurations` | `List<ServiceVolumeConfiguration>` | no |
| `vpcLatticeConfigurations` | `List<VpcLatticeConfiguration>` | no |
| `monitoring` | `MonitoringConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | no |

## UpdateServicePrimaryTaskSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `service` | `string` | yes |
| `primaryTaskSet` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSet` | `TaskSet` | no |

## UpdateTaskProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `tasks` | `List<string>` | yes |
| `protectionEnabled` | `boolean` | yes |
| `expiresInMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `protectedTasks` | `List<ProtectedTask>` | no |
| `failures` | `List<Failure>` | no |

## UpdateTaskSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `string` | yes |
| `service` | `string` | yes |
| `taskSet` | `string` | yes |
| `scale` | `Scale` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taskSet` | `TaskSet` | no |

