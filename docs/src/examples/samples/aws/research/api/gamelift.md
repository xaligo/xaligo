# Amazon GameLift

API version: 2015-10-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/gamelift/2015-10-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptMatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketId` | `string` | yes |
| `PlayerIds` | `List<string>` | yes |
| `AcceptanceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ClaimGameServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `GameServerId` | `string` | no |
| `GameServerData` | `string` | no |
| `FilterOption` | `ClaimFilterOption` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServer` | `GameServer` | no |

## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `RoutingStrategy` | `RoutingStrategy` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | no |

## CreateBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Version` | `string` | no |
| `StorageLocation` | `S3Location` | no |
| `OperatingSystem` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ServerSdkVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Build` | `Build` | no |
| `UploadCredentials` | `AwsCredentials` | no |
| `StorageLocation` | `S3Location` | no |

## CreateContainerFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetRoleArn` | `string` | yes |
| `Description` | `string` | no |
| `GameServerContainerGroupDefinitionName` | `string` | no |
| `PerInstanceContainerGroupDefinitionName` | `string` | no |
| `InstanceConnectionPortRange` | `ConnectionPortRange` | no |
| `InstanceInboundPermissions` | `List<IpPermission>` | no |
| `GameServerContainerGroupsPerInstance` | `integer` | no |
| `InstanceType` | `string` | no |
| `BillingType` | `string` | no |
| `Locations` | `List<LocationConfiguration>` | no |
| `MetricGroups` | `List<string>` | no |
| `NewGameSessionProtectionPolicy` | `string` | no |
| `GameSessionCreationLimitPolicy` | `GameSessionCreationLimitPolicy` | no |
| `LogConfiguration` | `LogConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `PlayerGatewayMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerFleet` | `ContainerFleet` | no |

## CreateContainerGroupDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ContainerGroupType` | `string` | no |
| `TotalMemoryLimitMebibytes` | `integer` | yes |
| `TotalVcpuLimit` | `double` | yes |
| `GameServerContainerDefinition` | `GameServerContainerDefinitionInput` | no |
| `SupportContainerDefinitions` | `List<SupportContainerDefinitionInput>` | no |
| `OperatingSystem` | `string` | yes |
| `VersionDescription` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinition` | `ContainerGroupDefinition` | no |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `BuildId` | `string` | no |
| `ScriptId` | `string` | no |
| `ServerLaunchPath` | `string` | no |
| `ServerLaunchParameters` | `string` | no |
| `LogPaths` | `List<string>` | no |
| `EC2InstanceType` | `string` | no |
| `EC2InboundPermissions` | `List<IpPermission>` | no |
| `NewGameSessionProtectionPolicy` | `string` | no |
| `RuntimeConfiguration` | `RuntimeConfiguration` | no |
| `ResourceCreationLimitPolicy` | `ResourceCreationLimitPolicy` | no |
| `MetricGroups` | `List<string>` | no |
| `PeerVpcAwsAccountId` | `string` | no |
| `PeerVpcId` | `string` | no |
| `FleetType` | `string` | no |
| `InstanceRoleArn` | `string` | no |
| `CertificateConfiguration` | `CertificateConfiguration` | no |
| `Locations` | `List<LocationConfiguration>` | no |
| `Tags` | `List<Tag>` | no |
| `ComputeType` | `string` | no |
| `AnywhereConfiguration` | `AnywhereConfiguration` | no |
| `InstanceRoleCredentialsProvider` | `string` | no |
| `PlayerGatewayMode` | `string` | no |
| `PlayerGatewayConfiguration` | `PlayerGatewayConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetAttributes` | `FleetAttributes` | no |
| `LocationStates` | `List<LocationState>` | no |

## CreateFleetLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Locations` | `List<LocationConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `LocationStates` | `List<LocationState>` | no |

## CreateGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `MinSize` | `integer` | yes |
| `MaxSize` | `integer` | yes |
| `LaunchTemplate` | `LaunchTemplateSpecification` | yes |
| `InstanceDefinitions` | `List<InstanceDefinition>` | yes |
| `AutoScalingPolicy` | `GameServerGroupAutoScalingPolicy` | no |
| `BalancingStrategy` | `string` | no |
| `GameServerProtectionPolicy` | `string` | no |
| `VpcSubnets` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## CreateGameSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `AliasId` | `string` | no |
| `MaximumPlayerSessionCount` | `integer` | yes |
| `Name` | `string` | no |
| `GameProperties` | `List<GameProperty>` | no |
| `CreatorId` | `string` | no |
| `GameSessionId` | `string` | no |
| `IdempotencyToken` | `string` | no |
| `GameSessionData` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSession` | `GameSession` | no |

## CreateGameSessionQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TimeoutInSeconds` | `integer` | no |
| `PlayerLatencyPolicies` | `List<PlayerLatencyPolicy>` | no |
| `Destinations` | `List<GameSessionQueueDestination>` | no |
| `FilterConfiguration` | `FilterConfiguration` | no |
| `PriorityConfiguration` | `PriorityConfiguration` | no |
| `CustomEventData` | `string` | no |
| `NotificationTarget` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionQueue` | `GameSessionQueue` | no |

## CreateLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `LocationModel` | no |

## CreateMatchmakingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `GameSessionQueueArns` | `List<string>` | no |
| `RequestTimeoutSeconds` | `integer` | yes |
| `AcceptanceTimeoutSeconds` | `integer` | no |
| `AcceptanceRequired` | `boolean` | yes |
| `RuleSetName` | `string` | yes |
| `NotificationTarget` | `string` | no |
| `AdditionalPlayerCount` | `integer` | no |
| `CustomEventData` | `string` | no |
| `GameProperties` | `List<GameProperty>` | no |
| `GameSessionData` | `string` | no |
| `BackfillMode` | `string` | no |
| `FlexMatchMode` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `MatchmakingConfiguration` | no |

## CreateMatchmakingRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RuleSetBody` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSet` | `MatchmakingRuleSet` | yes |

## CreatePlayerSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |
| `PlayerId` | `string` | yes |
| `PlayerData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlayerSession` | `PlayerSession` | no |

## CreatePlayerSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |
| `PlayerIds` | `List<string>` | yes |
| `PlayerDataMap` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlayerSessions` | `List<PlayerSession>` | no |

## CreateScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Version` | `string` | no |
| `StorageLocation` | `S3Location` | no |
| `ZipFile` | `blob` | no |
| `Tags` | `List<Tag>` | no |
| `NodeJsVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Script` | `Script` | no |

## CreateVpcPeeringAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameLiftAwsAccountId` | `string` | yes |
| `PeerVpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringAuthorization` | `VpcPeeringAuthorization` | no |

## CreateVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `PeerVpcAwsAccountId` | `string` | yes |
| `PeerVpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BuildId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContainerFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContainerGroupDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `VersionNumber` | `integer` | no |
| `VersionCountToRetain` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleetLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Locations` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `LocationStates` | `List<LocationState>` | no |

## DeleteGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `DeleteOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## DeleteGameSessionQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMatchmakingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMatchmakingRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `FleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScriptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcPeeringAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameLiftAwsAccountId` | `string` | yes |
| `PeerVpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `VpcPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterCompute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ComputeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterGameServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `GameServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | no |

## DescribeBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BuildId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Build` | `Build` | no |

## DescribeCompute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ComputeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Compute` | `Compute` | no |

## DescribeContainerFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerFleet` | `ContainerFleet` | no |

## DescribeContainerGroupDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `VersionNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinition` | `ContainerGroupDefinition` | no |

## DescribeContainerGroupPortMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ContainerGroupType` | `string` | yes |
| `ComputeName` | `string` | no |
| `InstanceId` | `string` | no |
| `ContainerName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `Location` | `string` | no |
| `ContainerGroupDefinitionArn` | `string` | no |
| `ContainerGroupType` | `string` | no |
| `ComputeName` | `string` | no |
| `InstanceId` | `string` | no |
| `ContainerGroupPortMappings` | `List<ContainerGroupPortMapping>` | no |

## DescribeEC2InstanceLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EC2InstanceType` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EC2InstanceLimits` | `List<EC2InstanceLimit>` | no |

## DescribeFleetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetIds` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetAttributes` | `List<FleetAttributes>` | no |
| `NextToken` | `string` | no |

## DescribeFleetCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetIds` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetCapacity` | `List<FleetCapacity>` | no |
| `NextToken` | `string` | no |

## DescribeFleetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `DeploymentId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetDeployment` | `FleetDeployment` | no |
| `LocationalDeployments` | `Map<LocationalDeployment>` | no |

## DescribeFleetEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<Event>` | no |
| `NextToken` | `string` | no |

## DescribeFleetLocationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Locations` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `LocationAttributes` | `List<LocationAttributes>` | no |
| `NextToken` | `string` | no |

## DescribeFleetLocationCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Location` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetCapacity` | `FleetCapacity` | no |

## DescribeFleetLocationUtilization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Location` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetUtilization` | `FleetUtilization` | no |

## DescribeFleetPortSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `InboundPermissions` | `List<IpPermission>` | no |
| `UpdateStatus` | `string` | no |
| `Location` | `string` | no |

## DescribeFleetUtilization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetIds` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetUtilization` | `List<FleetUtilization>` | no |
| `NextToken` | `string` | no |

## DescribeGameServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `GameServerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServer` | `GameServer` | no |

## DescribeGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## DescribeGameServerInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `InstanceIds` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerInstances` | `List<GameServerInstance>` | no |
| `NextToken` | `string` | no |

## DescribeGameSessionDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `GameSessionId` | `string` | no |
| `AliasId` | `string` | no |
| `Location` | `string` | no |
| `StatusFilter` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionDetails` | `List<GameSessionDetail>` | no |
| `NextToken` | `string` | no |

## DescribeGameSessionPlacement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlacementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionPlacement` | `GameSessionPlacement` | no |

## DescribeGameSessionQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionQueues` | `List<GameSessionQueue>` | no |
| `NextToken` | `string` | no |

## DescribeGameSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `GameSessionId` | `string` | no |
| `AliasId` | `string` | no |
| `Location` | `string` | no |
| `StatusFilter` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessions` | `List<GameSession>` | no |
| `NextToken` | `string` | no |

## DescribeInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `InstanceId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<Instance>` | no |
| `NextToken` | `string` | no |

## DescribeMatchmaking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketList` | `List<MatchmakingTicket>` | no |

## DescribeMatchmakingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `RuleSetName` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configurations` | `List<MatchmakingConfiguration>` | no |
| `NextToken` | `string` | no |

## DescribeMatchmakingRuleSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSets` | `List<MatchmakingRuleSet>` | yes |
| `NextToken` | `string` | no |

## DescribePlayerSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | no |
| `PlayerId` | `string` | no |
| `PlayerSessionId` | `string` | no |
| `PlayerSessionStatusFilter` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlayerSessions` | `List<PlayerSession>` | no |
| `NextToken` | `string` | no |

## DescribeRuntimeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuntimeConfiguration` | `RuntimeConfiguration` | no |

## DescribeScalingPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `StatusFilter` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPolicies` | `List<ScalingPolicy>` | no |
| `NextToken` | `string` | no |

## DescribeScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScriptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Script` | `Script` | no |

## DescribeVpcPeeringAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringAuthorizations` | `List<VpcPeeringAuthorization>` | no |

## DescribeVpcPeeringConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringConnections` | `List<VpcPeeringConnection>` | no |

## GetComputeAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ComputeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `ComputeName` | `string` | no |
| `ComputeArn` | `string` | no |
| `Credentials` | `AwsCredentials` | no |
| `Target` | `string` | no |
| `ContainerIdentifiers` | `List<ContainerIdentifier>` | no |

## GetComputeAuthToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ComputeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `ComputeName` | `string` | no |
| `ComputeArn` | `string` | no |
| `AuthToken` | `string` | no |
| `ExpirationTimestamp` | `timestamp` | no |

## GetGameSessionLogUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PreSignedUrl` | `string` | no |

## GetInstanceAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceAccess` | `InstanceAccess` | no |

## GetPlayerConnectionDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |
| `PlayerIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | no |
| `PlayerConnectionDetails` | `List<PlayerConnectionDetail>` | no |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingStrategyType` | `string` | no |
| `Name` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Aliases` | `List<Alias>` | no |
| `NextToken` | `string` | no |

## ListBuilds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Builds` | `List<Build>` | no |
| `NextToken` | `string` | no |

## ListCompute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Location` | `string` | no |
| `ContainerGroupDefinitionName` | `string` | no |
| `ComputeStatus` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeList` | `List<Compute>` | no |
| `NextToken` | `string` | no |

## ListContainerFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinitionName` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerFleets` | `List<ContainerFleet>` | no |
| `NextToken` | `string` | no |

## ListContainerGroupDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinitions` | `List<ContainerGroupDefinition>` | no |
| `NextToken` | `string` | no |

## ListContainerGroupDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupType` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinitions` | `List<ContainerGroupDefinition>` | no |
| `NextToken` | `string` | no |

## ListFleetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetDeployments` | `List<FleetDeployment>` | no |
| `NextToken` | `string` | no |

## ListFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BuildId` | `string` | no |
| `ScriptId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListGameServerGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroups` | `List<GameServerGroup>` | no |
| `NextToken` | `string` | no |

## ListGameServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `SortOrder` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServers` | `List<GameServer>` | no |
| `NextToken` | `string` | no |

## ListLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Locations` | `List<LocationModel>` | no |
| `NextToken` | `string` | no |

## ListScripts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scripts` | `List<Script>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `FleetId` | `string` | yes |
| `ScalingAdjustment` | `integer` | no |
| `ScalingAdjustmentType` | `string` | no |
| `Threshold` | `double` | no |
| `ComparisonOperator` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `MetricName` | `string` | yes |
| `PolicyType` | `string` | no |
| `TargetConfiguration` | `TargetConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## RegisterCompute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `ComputeName` | `string` | yes |
| `CertificatePath` | `string` | no |
| `DnsName` | `string` | no |
| `IpAddress` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Compute` | `Compute` | no |

## RegisterGameServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `GameServerId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ConnectionInfo` | `string` | no |
| `GameServerData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServer` | `GameServer` | no |

## RequestUploadCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BuildId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UploadCredentials` | `AwsCredentials` | no |
| `StorageLocation` | `S3Location` | no |

## ResolveAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |

## ResumeGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `ResumeActions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## SearchGameSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `AliasId` | `string` | no |
| `Location` | `string` | no |
| `FilterExpression` | `string` | no |
| `SortExpression` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessions` | `List<GameSession>` | no |
| `NextToken` | `string` | no |

## StartFleetActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Actions` | `List<string>` | yes |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |

## StartGameSessionPlacement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlacementId` | `string` | yes |
| `GameSessionQueueName` | `string` | yes |
| `GameProperties` | `List<GameProperty>` | no |
| `MaximumPlayerSessionCount` | `integer` | yes |
| `GameSessionName` | `string` | no |
| `PlayerLatencies` | `List<PlayerLatency>` | no |
| `DesiredPlayerSessions` | `List<DesiredPlayerSession>` | no |
| `GameSessionData` | `string` | no |
| `PriorityConfigurationOverride` | `PriorityConfigurationOverride` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionPlacement` | `GameSessionPlacement` | no |

## StartMatchBackfill

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketId` | `string` | no |
| `ConfigurationName` | `string` | yes |
| `GameSessionArn` | `string` | no |
| `Players` | `List<Player>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MatchmakingTicket` | `MatchmakingTicket` | no |

## StartMatchmaking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketId` | `string` | no |
| `ConfigurationName` | `string` | yes |
| `Players` | `List<Player>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MatchmakingTicket` | `MatchmakingTicket` | no |

## StopFleetActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Actions` | `List<string>` | yes |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |

## StopGameSessionPlacement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlacementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionPlacement` | `GameSessionPlacement` | no |

## StopMatchmaking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SuspendGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `SuspendActions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateGameSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |
| `TerminationMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSession` | `GameSession` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `RoutingStrategy` | `RoutingStrategy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `Alias` | no |

## UpdateBuild

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BuildId` | `string` | yes |
| `Name` | `string` | no |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Build` | `Build` | no |

## UpdateContainerFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `GameServerContainerGroupDefinitionName` | `string` | no |
| `PerInstanceContainerGroupDefinitionName` | `string` | no |
| `GameServerContainerGroupsPerInstance` | `integer` | no |
| `InstanceConnectionPortRange` | `ConnectionPortRange` | no |
| `InstanceInboundPermissionAuthorizations` | `List<IpPermission>` | no |
| `InstanceInboundPermissionRevocations` | `List<IpPermission>` | no |
| `DeploymentConfiguration` | `DeploymentConfiguration` | no |
| `Description` | `string` | no |
| `MetricGroups` | `List<string>` | no |
| `NewGameSessionProtectionPolicy` | `string` | no |
| `GameSessionCreationLimitPolicy` | `GameSessionCreationLimitPolicy` | no |
| `LogConfiguration` | `LogConfiguration` | no |
| `RemoveAttributes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerFleet` | `ContainerFleet` | no |

## UpdateContainerGroupDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `GameServerContainerDefinition` | `GameServerContainerDefinitionInput` | no |
| `SupportContainerDefinitions` | `List<SupportContainerDefinitionInput>` | no |
| `TotalMemoryLimitMebibytes` | `integer` | no |
| `TotalVcpuLimit` | `double` | no |
| `VersionDescription` | `string` | no |
| `SourceVersionNumber` | `integer` | no |
| `OperatingSystem` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerGroupDefinition` | `ContainerGroupDefinition` | no |

## UpdateFleetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `NewGameSessionProtectionPolicy` | `string` | no |
| `ResourceCreationLimitPolicy` | `ResourceCreationLimitPolicy` | no |
| `MetricGroups` | `List<string>` | no |
| `AnywhereConfiguration` | `AnywhereConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |

## UpdateFleetCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `DesiredInstances` | `integer` | no |
| `MinSize` | `integer` | no |
| `MaxSize` | `integer` | no |
| `Location` | `string` | no |
| `ManagedCapacityConfiguration` | `ManagedCapacityConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |
| `Location` | `string` | no |
| `ManagedCapacityConfiguration` | `ManagedCapacityConfiguration` | no |

## UpdateFleetPortSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `InboundPermissionAuthorizations` | `List<IpPermission>` | no |
| `InboundPermissionRevocations` | `List<IpPermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `FleetArn` | `string` | no |

## UpdateGameServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `GameServerId` | `string` | yes |
| `GameServerData` | `string` | no |
| `UtilizationStatus` | `string` | no |
| `HealthCheck` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServer` | `GameServer` | no |

## UpdateGameServerGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroupName` | `string` | yes |
| `RoleArn` | `string` | no |
| `InstanceDefinitions` | `List<InstanceDefinition>` | no |
| `GameServerProtectionPolicy` | `string` | no |
| `BalancingStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameServerGroup` | `GameServerGroup` | no |

## UpdateGameSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionId` | `string` | yes |
| `MaximumPlayerSessionCount` | `integer` | no |
| `Name` | `string` | no |
| `PlayerSessionCreationPolicy` | `string` | no |
| `ProtectionPolicy` | `string` | no |
| `GameProperties` | `List<GameProperty>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSession` | `GameSession` | no |

## UpdateGameSessionQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TimeoutInSeconds` | `integer` | no |
| `PlayerLatencyPolicies` | `List<PlayerLatencyPolicy>` | no |
| `Destinations` | `List<GameSessionQueueDestination>` | no |
| `FilterConfiguration` | `FilterConfiguration` | no |
| `PriorityConfiguration` | `PriorityConfiguration` | no |
| `CustomEventData` | `string` | no |
| `NotificationTarget` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GameSessionQueue` | `GameSessionQueue` | no |

## UpdateMatchmakingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `GameSessionQueueArns` | `List<string>` | no |
| `RequestTimeoutSeconds` | `integer` | no |
| `AcceptanceTimeoutSeconds` | `integer` | no |
| `AcceptanceRequired` | `boolean` | no |
| `RuleSetName` | `string` | no |
| `NotificationTarget` | `string` | no |
| `AdditionalPlayerCount` | `integer` | no |
| `CustomEventData` | `string` | no |
| `GameProperties` | `List<GameProperty>` | no |
| `GameSessionData` | `string` | no |
| `BackfillMode` | `string` | no |
| `FlexMatchMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `MatchmakingConfiguration` | no |

## UpdateRuntimeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | yes |
| `RuntimeConfiguration` | `RuntimeConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuntimeConfiguration` | `RuntimeConfiguration` | no |

## UpdateScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScriptId` | `string` | yes |
| `Name` | `string` | no |
| `Version` | `string` | no |
| `StorageLocation` | `S3Location` | no |
| `ZipFile` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Script` | `Script` | no |

## ValidateMatchmakingRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetBody` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Valid` | `boolean` | no |

