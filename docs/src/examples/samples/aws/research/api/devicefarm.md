# AWS Device Farm

API version: 2015-06-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/devicefarm/2015-06-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDevicePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `rules` | `List<Rule>` | yes |
| `maxDevices` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devicePool` | `DevicePool` | no |

## CreateInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `packageCleanup` | `boolean` | no |
| `excludeAppPackagesFromCleanup` | `List<string>` | no |
| `rebootAfterUse` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceProfile` | `InstanceProfile` | no |

## CreateNetworkProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | no |
| `uplinkBandwidthBits` | `long` | no |
| `downlinkBandwidthBits` | `long` | no |
| `uplinkDelayMs` | `long` | no |
| `downlinkDelayMs` | `long` | no |
| `uplinkJitterMs` | `long` | no |
| `downlinkJitterMs` | `long` | no |
| `uplinkLossPercent` | `integer` | no |
| `downlinkLossPercent` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkProfile` | `NetworkProfile` | no |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `defaultJobTimeoutMinutes` | `integer` | no |
| `vpcConfig` | `VpcConfig` | no |
| `environmentVariables` | `List<EnvironmentVariable>` | no |
| `executionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `Project` | no |

## CreateRemoteAccessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `deviceArn` | `string` | yes |
| `appArn` | `string` | no |
| `instanceArn` | `string` | no |
| `name` | `string` | no |
| `configuration` | `CreateRemoteAccessSessionConfiguration` | no |
| `interactionMode` | `string` | no |
| `skipAppResign` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `remoteAccessSession` | `RemoteAccessSession` | no |

## CreateTestGridProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `vpcConfig` | `TestGridVpcConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridProject` | `TestGridProject` | no |

## CreateTestGridUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `expiresInSeconds` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `url` | `string` | no |
| `expires` | `timestamp` | no |

## CreateUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `contentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `upload` | `Upload` | no |

## CreateVPCEConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpceConfigurationName` | `string` | yes |
| `vpceServiceName` | `string` | yes |
| `serviceDnsName` | `string` | yes |
| `vpceConfigurationDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpceConfiguration` | `VPCEConfiguration` | no |

## DeleteDevicePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkProfile

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
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRemoteAccessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTestGridProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVPCEConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettings` | `AccountSettings` | no |

## GetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `device` | `Device` | no |

## GetDeviceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceInstance` | `DeviceInstance` | no |

## GetDevicePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devicePool` | `DevicePool` | no |

## GetDevicePoolCompatibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devicePoolArn` | `string` | yes |
| `appArn` | `string` | no |
| `testType` | `string` | no |
| `test` | `ScheduleRunTest` | no |
| `configuration` | `ScheduleRunConfiguration` | no |
| `projectArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `compatibleDevices` | `List<DevicePoolCompatibilityResult>` | no |
| `incompatibleDevices` | `List<DevicePoolCompatibilityResult>` | no |

## GetInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceProfile` | `InstanceProfile` | no |

## GetJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## GetNetworkProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkProfile` | `NetworkProfile` | no |

## GetOfferingStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `current` | `Map<OfferingStatus>` | no |
| `nextPeriod` | `Map<OfferingStatus>` | no |
| `nextToken` | `string` | no |

## GetProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `Project` | no |

## GetRemoteAccessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `remoteAccessSession` | `RemoteAccessSession` | no |

## GetRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `run` | `Run` | no |

## GetSuite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suite` | `Suite` | no |

## GetTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `test` | `Test` | no |

## GetTestGridProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridProject` | `TestGridProject` | no |

## GetTestGridSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | no |
| `sessionId` | `string` | no |
| `sessionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridSession` | `TestGridSession` | no |

## GetUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `upload` | `Upload` | no |

## GetVPCEConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpceConfiguration` | `VPCEConfiguration` | no |

## InstallToRemoteAccessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `remoteAccessSessionArn` | `string` | yes |
| `appArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `appUpload` | `Upload` | no |

## ListArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `type` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifacts` | `List<Artifact>` | no |
| `nextToken` | `string` | no |

## ListDeviceInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceInstances` | `List<DeviceInstance>` | no |
| `nextToken` | `string` | no |

## ListDevicePools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `type` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devicePools` | `List<DevicePool>` | no |
| `nextToken` | `string` | no |

## ListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `nextToken` | `string` | no |
| `filters` | `List<DeviceFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devices` | `List<Device>` | no |
| `nextToken` | `string` | no |

## ListInstanceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceProfiles` | `List<InstanceProfile>` | no |
| `nextToken` | `string` | no |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobs` | `List<Job>` | no |
| `nextToken` | `string` | no |

## ListNetworkProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `type` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkProfiles` | `List<NetworkProfile>` | no |
| `nextToken` | `string` | no |

## ListOfferingPromotions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringPromotions` | `List<OfferingPromotion>` | no |
| `nextToken` | `string` | no |

## ListOfferingTransactions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringTransactions` | `List<OfferingTransaction>` | no |
| `nextToken` | `string` | no |

## ListOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerings` | `List<Offering>` | no |
| `nextToken` | `string` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projects` | `List<Project>` | no |
| `nextToken` | `string` | no |

## ListRemoteAccessSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `remoteAccessSessions` | `List<RemoteAccessSession>` | no |
| `nextToken` | `string` | no |

## ListRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runs` | `List<Run>` | no |
| `nextToken` | `string` | no |

## ListSamples

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `samples` | `List<Sample>` | no |
| `nextToken` | `string` | no |

## ListSuites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `suites` | `List<Suite>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTestGridProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResult` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridProjects` | `List<TestGridProject>` | no |
| `nextToken` | `string` | no |

## ListTestGridSessionActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionArn` | `string` | yes |
| `maxResult` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actions` | `List<TestGridSessionAction>` | no |
| `nextToken` | `string` | no |

## ListTestGridSessionArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionArn` | `string` | yes |
| `type` | `string` | no |
| `maxResult` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `artifacts` | `List<TestGridSessionArtifact>` | no |
| `nextToken` | `string` | no |

## ListTestGridSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `status` | `string` | no |
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `endTimeAfter` | `timestamp` | no |
| `endTimeBefore` | `timestamp` | no |
| `maxResult` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridSessions` | `List<TestGridSession>` | no |
| `nextToken` | `string` | no |

## ListTests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tests` | `List<Test>` | no |
| `nextToken` | `string` | no |

## ListUniqueProblems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uniqueProblems` | `Map<List<UniqueProblem>>` | no |
| `nextToken` | `string` | no |

## ListUploads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `type` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploads` | `List<Upload>` | no |
| `nextToken` | `string` | no |

## ListVPCEConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpceConfigurations` | `List<VPCEConfiguration>` | no |
| `nextToken` | `string` | no |

## PurchaseOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringId` | `string` | yes |
| `quantity` | `integer` | yes |
| `offeringPromotionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringTransaction` | `OfferingTransaction` | no |

## RenewOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringId` | `string` | yes |
| `quantity` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offeringTransaction` | `OfferingTransaction` | no |

## ScheduleRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `appArn` | `string` | no |
| `devicePoolArn` | `string` | no |
| `deviceSelectionConfiguration` | `DeviceSelectionConfiguration` | no |
| `name` | `string` | no |
| `test` | `ScheduleRunTest` | yes |
| `configuration` | `ScheduleRunConfiguration` | no |
| `executionConfiguration` | `ExecutionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `run` | `Run` | no |

## StopJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `Job` | no |

## StopRemoteAccessSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `remoteAccessSession` | `RemoteAccessSession` | no |

## StopRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `run` | `Run` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDeviceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `profileArn` | `string` | no |
| `labels` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceInstance` | `DeviceInstance` | no |

## UpdateDevicePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `rules` | `List<Rule>` | no |
| `maxDevices` | `integer` | no |
| `clearMaxDevices` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `devicePool` | `DevicePool` | no |

## UpdateInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `packageCleanup` | `boolean` | no |
| `excludeAppPackagesFromCleanup` | `List<string>` | no |
| `rebootAfterUse` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceProfile` | `InstanceProfile` | no |

## UpdateNetworkProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `type` | `string` | no |
| `uplinkBandwidthBits` | `long` | no |
| `downlinkBandwidthBits` | `long` | no |
| `uplinkDelayMs` | `long` | no |
| `downlinkDelayMs` | `long` | no |
| `uplinkJitterMs` | `long` | no |
| `downlinkJitterMs` | `long` | no |
| `uplinkLossPercent` | `integer` | no |
| `downlinkLossPercent` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkProfile` | `NetworkProfile` | no |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `defaultJobTimeoutMinutes` | `integer` | no |
| `vpcConfig` | `VpcConfig` | no |
| `environmentVariables` | `List<EnvironmentVariable>` | no |
| `executionRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `project` | `Project` | no |

## UpdateTestGridProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectArn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `vpcConfig` | `TestGridVpcConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testGridProject` | `TestGridProject` | no |

## UpdateUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `contentType` | `string` | no |
| `editContent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `upload` | `Upload` | no |

## UpdateVPCEConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `vpceConfigurationName` | `string` | no |
| `vpceServiceName` | `string` | no |
| `serviceDnsName` | `string` | no |
| `vpceConfigurationDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpceConfiguration` | `VPCEConfiguration` | no |

