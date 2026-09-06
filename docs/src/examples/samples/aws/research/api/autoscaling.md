# Auto Scaling

API version: 2011-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/autoscaling/2011-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AttachInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachLoadBalancerTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TargetGroupARNs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `LoadBalancerNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachTrafficSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TrafficSources` | `List<TrafficSourceIdentifier>` | yes |
| `SkipZonalShiftValidation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDeleteScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScheduledActionNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedScheduledActions` | `List<FailedScheduledUpdateGroupActionRequest>` | no |

## BatchPutScheduledUpdateGroupAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScheduledUpdateGroupActions` | `List<ScheduledUpdateGroupActionRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedScheduledUpdateGroupActions` | `List<FailedScheduledUpdateGroupActionRequest>` | no |

## CancelInstanceRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `WaitForTransitioningInstances` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceRefreshId` | `string` | no |

## CompleteLifecycleAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHookName` | `string` | yes |
| `AutoScalingGroupName` | `string` | yes |
| `LifecycleActionToken` | `string` | no |
| `LifecycleActionResult` | `string` | yes |
| `InstanceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAutoScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `LaunchConfigurationName` | `string` | no |
| `LaunchTemplate` | `LaunchTemplateSpecification` | no |
| `MixedInstancesPolicy` | `MixedInstancesPolicy` | no |
| `InstanceId` | `string` | no |
| `MinSize` | `integer` | yes |
| `MaxSize` | `integer` | yes |
| `DesiredCapacity` | `integer` | no |
| `DefaultCooldown` | `integer` | no |
| `AvailabilityZones` | `List<string>` | no |
| `AvailabilityZoneIds` | `List<string>` | no |
| `LoadBalancerNames` | `List<string>` | no |
| `TargetGroupARNs` | `List<string>` | no |
| `HealthCheckType` | `string` | no |
| `HealthCheckGracePeriod` | `integer` | no |
| `PlacementGroup` | `string` | no |
| `VPCZoneIdentifier` | `string` | no |
| `TerminationPolicies` | `List<string>` | no |
| `NewInstancesProtectedFromScaleIn` | `boolean` | no |
| `CapacityRebalance` | `boolean` | no |
| `LifecycleHookSpecificationList` | `List<LifecycleHookSpecification>` | no |
| `DeletionProtection` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ServiceLinkedRoleARN` | `string` | no |
| `MaxInstanceLifetime` | `integer` | no |
| `Context` | `string` | no |
| `DesiredCapacityType` | `string` | no |
| `DefaultInstanceWarmup` | `integer` | no |
| `TrafficSources` | `List<TrafficSourceIdentifier>` | no |
| `InstanceMaintenancePolicy` | `InstanceMaintenancePolicy` | no |
| `AvailabilityZoneDistribution` | `AvailabilityZoneDistribution` | no |
| `AvailabilityZoneImpairmentPolicy` | `AvailabilityZoneImpairmentPolicy` | no |
| `SkipZonalShiftValidation` | `boolean` | no |
| `CapacityReservationSpecification` | `CapacityReservationSpecification` | no |
| `InstanceLifecyclePolicy` | `InstanceLifecyclePolicy` | no |
| `Operator` | `Operator` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchConfigurationName` | `string` | yes |
| `ImageId` | `string` | no |
| `KeyName` | `string` | no |
| `SecurityGroups` | `List<string>` | no |
| `ClassicLinkVPCId` | `string` | no |
| `ClassicLinkVPCSecurityGroups` | `List<string>` | no |
| `UserData` | `string` | no |
| `InstanceId` | `string` | no |
| `InstanceType` | `string` | no |
| `KernelId` | `string` | no |
| `RamdiskId` | `string` | no |
| `BlockDeviceMappings` | `List<BlockDeviceMapping>` | no |
| `InstanceMonitoring` | `InstanceMonitoring` | no |
| `SpotPrice` | `string` | no |
| `IamInstanceProfile` | `string` | no |
| `EbsOptimized` | `boolean` | no |
| `AssociatePublicIpAddress` | `boolean` | no |
| `PlacementTenancy` | `string` | no |
| `MetadataOptions` | `InstanceMetadataOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateOrUpdateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAutoScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLaunchConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLifecycleHook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHookName` | `string` | yes |
| `AutoScalingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TopicARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | no |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScheduledActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWarmPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ForceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxNumberOfAutoScalingGroups` | `integer` | no |
| `MaxNumberOfLaunchConfigurations` | `integer` | no |
| `NumberOfAutoScalingGroups` | `integer` | no |
| `NumberOfLaunchConfigurations` | `integer` | no |

## DescribeAdjustmentTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdjustmentTypes` | `List<AdjustmentType>` | no |

## DescribeAutoScalingGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupNames` | `List<string>` | no |
| `IncludeInstances` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroups` | `List<AutoScalingGroup>` | yes |
| `NextToken` | `string` | no |

## DescribeAutoScalingInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingInstances` | `List<AutoScalingInstanceDetails>` | no |
| `NextToken` | `string` | no |

## DescribeAutoScalingNotificationTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingNotificationTypes` | `List<string>` | no |

## DescribeInstanceRefreshes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `InstanceRefreshIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceRefreshes` | `List<InstanceRefresh>` | no |
| `NextToken` | `string` | no |

## DescribeLaunchConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchConfigurationNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchConfigurations` | `List<LaunchConfiguration>` | yes |
| `NextToken` | `string` | no |

## DescribeLifecycleHookTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHookTypes` | `List<string>` | no |

## DescribeLifecycleHooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `LifecycleHookNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHooks` | `List<LifecycleHook>` | no |

## DescribeLoadBalancerTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerTargetGroups` | `List<LoadBalancerTargetGroupState>` | no |
| `NextToken` | `string` | no |

## DescribeLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancers` | `List<LoadBalancerState>` | no |
| `NextToken` | `string` | no |

## DescribeMetricCollectionTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metrics` | `List<MetricCollectionType>` | no |
| `Granularities` | `List<MetricGranularityType>` | no |

## DescribeNotificationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationConfigurations` | `List<NotificationConfiguration>` | yes |
| `NextToken` | `string` | no |

## DescribePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | no |
| `PolicyNames` | `List<string>` | no |
| `PolicyTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPolicies` | `List<ScalingPolicy>` | no |
| `NextToken` | `string` | no |

## DescribeScalingActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivityIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | no |
| `IncludeDeletedGroups` | `boolean` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Activities` | `List<Activity>` | yes |
| `NextToken` | `string` | no |

## DescribeScalingProcessTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Processes` | `List<ProcessType>` | no |

## DescribeScheduledActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | no |
| `ScheduledActionNames` | `List<string>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledUpdateGroupActions` | `List<ScheduledUpdateGroupAction>` | no |
| `NextToken` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<TagDescription>` | no |
| `NextToken` | `string` | no |

## DescribeTerminationPolicyTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminationPolicyTypes` | `List<string>` | no |

## DescribeTrafficSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TrafficSourceType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxRecords` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficSources` | `List<TrafficSourceState>` | no |
| `NextToken` | `string` | no |

## DescribeWarmPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WarmPoolConfiguration` | `WarmPoolConfiguration` | no |
| `Instances` | `List<Instance>` | no |
| `NextToken` | `string` | no |

## DetachInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | yes |
| `ShouldDecrementDesiredCapacity` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Activities` | `List<Activity>` | no |

## DetachLoadBalancerTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TargetGroupARNs` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `LoadBalancerNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachTrafficSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TrafficSources` | `List<TrafficSourceIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableMetricsCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `Metrics` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableMetricsCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `Metrics` | `List<string>` | no |
| `Granularity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnterStandby

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | yes |
| `ShouldDecrementDesiredCapacity` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Activities` | `List<Activity>` | no |

## ExecutePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | no |
| `PolicyName` | `string` | yes |
| `HonorCooldown` | `boolean` | no |
| `MetricValue` | `double` | no |
| `BreachThreshold` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExitStandby

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Activities` | `List<Activity>` | no |

## GetPredictiveScalingForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadForecast` | `List<LoadForecast>` | yes |
| `CapacityForecast` | `CapacityForecast` | yes |
| `UpdateTime` | `timestamp` | yes |

## LaunchInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `RequestedCapacity` | `integer` | yes |
| `ClientToken` | `string` | yes |
| `AvailabilityZones` | `List<string>` | no |
| `AvailabilityZoneIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | no |
| `RetryStrategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | no |
| `ClientToken` | `string` | no |
| `Instances` | `List<InstanceCollection>` | no |
| `Errors` | `List<LaunchInstancesError>` | no |

## PutLifecycleHook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHookName` | `string` | yes |
| `AutoScalingGroupName` | `string` | yes |
| `LifecycleTransition` | `string` | no |
| `RoleARN` | `string` | no |
| `NotificationTargetARN` | `string` | no |
| `NotificationMetadata` | `string` | no |
| `HeartbeatTimeout` | `integer` | no |
| `DefaultResult` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `TopicARN` | `string` | yes |
| `NotificationTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyType` | `string` | no |
| `AdjustmentType` | `string` | no |
| `MinAdjustmentStep` | `integer` | no |
| `MinAdjustmentMagnitude` | `integer` | no |
| `ScalingAdjustment` | `integer` | no |
| `Cooldown` | `integer` | no |
| `MetricAggregationType` | `string` | no |
| `StepAdjustments` | `List<StepAdjustment>` | no |
| `EstimatedInstanceWarmup` | `integer` | no |
| `TargetTrackingConfiguration` | `TargetTrackingConfiguration` | no |
| `Enabled` | `boolean` | no |
| `PredictiveScalingConfiguration` | `PredictiveScalingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyARN` | `string` | no |
| `Alarms` | `List<Alarm>` | no |

## PutScheduledUpdateGroupAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScheduledActionName` | `string` | yes |
| `Time` | `timestamp` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Recurrence` | `string` | no |
| `MinSize` | `integer` | no |
| `MaxSize` | `integer` | no |
| `DesiredCapacity` | `integer` | no |
| `TimeZone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutWarmPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `MaxGroupPreparedCapacity` | `integer` | no |
| `MinSize` | `integer` | no |
| `PoolState` | `string` | no |
| `InstanceReusePolicy` | `InstanceReusePolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RecordLifecycleActionHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecycleHookName` | `string` | yes |
| `AutoScalingGroupName` | `string` | yes |
| `LifecycleActionToken` | `string` | no |
| `InstanceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResumeProcesses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScalingProcesses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RollbackInstanceRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceRefreshId` | `string` | no |

## SetDesiredCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `DesiredCapacity` | `integer` | yes |
| `HonorCooldown` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetInstanceHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HealthStatus` | `string` | yes |
| `ShouldRespectGracePeriod` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetInstanceProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `AutoScalingGroupName` | `string` | yes |
| `ProtectedFromScaleIn` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartInstanceRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `Strategy` | `string` | no |
| `DesiredConfiguration` | `DesiredConfiguration` | no |
| `Preferences` | `RefreshPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceRefreshId` | `string` | no |

## SuspendProcesses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `ScalingProcesses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateInstanceInAutoScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `InstanceIds` | `List<string>` | no |
| `AutoScalingGroupName` | `string` | no |
| `ShouldDecrementDesiredCapacity` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Activity` | `Activity` | no |
| `Activities` | `List<Activity>` | no |

## UpdateAutoScalingGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingGroupName` | `string` | yes |
| `LaunchConfigurationName` | `string` | no |
| `LaunchTemplate` | `LaunchTemplateSpecification` | no |
| `MixedInstancesPolicy` | `MixedInstancesPolicy` | no |
| `MinSize` | `integer` | no |
| `MaxSize` | `integer` | no |
| `DesiredCapacity` | `integer` | no |
| `DefaultCooldown` | `integer` | no |
| `AvailabilityZones` | `List<string>` | no |
| `AvailabilityZoneIds` | `List<string>` | no |
| `HealthCheckType` | `string` | no |
| `HealthCheckGracePeriod` | `integer` | no |
| `PlacementGroup` | `string` | no |
| `VPCZoneIdentifier` | `string` | no |
| `TerminationPolicies` | `List<string>` | no |
| `NewInstancesProtectedFromScaleIn` | `boolean` | no |
| `ServiceLinkedRoleARN` | `string` | no |
| `MaxInstanceLifetime` | `integer` | no |
| `CapacityRebalance` | `boolean` | no |
| `Context` | `string` | no |
| `DesiredCapacityType` | `string` | no |
| `DefaultInstanceWarmup` | `integer` | no |
| `InstanceMaintenancePolicy` | `InstanceMaintenancePolicy` | no |
| `AvailabilityZoneDistribution` | `AvailabilityZoneDistribution` | no |
| `AvailabilityZoneImpairmentPolicy` | `AvailabilityZoneImpairmentPolicy` | no |
| `SkipZonalShiftValidation` | `boolean` | no |
| `CapacityReservationSpecification` | `CapacityReservationSpecification` | no |
| `InstanceLifecyclePolicy` | `InstanceLifecyclePolicy` | no |
| `DeletionProtection` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


