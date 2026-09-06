# Application Auto Scaling

API version: 2016-02-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/application-autoscaling/2016-02-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyName` | `string` | yes |
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ScheduledActionName` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterScalableTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeScalableTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ResourceIds` | `List<string>` | no |
| `ScalableDimension` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalableTargets` | `List<ScalableTarget>` | no |
| `NextToken` | `string` | no |

## DescribeScalingActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | no |
| `ScalableDimension` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeNotScaledActivities` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingActivities` | `List<ScalingActivity>` | no |
| `NextToken` | `string` | no |

## DescribeScalingPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | no |
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | no |
| `ScalableDimension` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPolicies` | `List<ScalingPolicy>` | no |
| `NextToken` | `string` | no |

## DescribeScheduledActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActionNames` | `List<string>` | no |
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | no |
| `ScalableDimension` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActions` | `List<ScheduledAction>` | no |
| `NextToken` | `string` | no |

## GetPredictiveScalingForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |
| `PolicyName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadForecast` | `List<LoadForecast>` | no |
| `CapacityForecast` | `CapacityForecast` | no |
| `UpdateTime` | `timestamp` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PutScalingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyName` | `string` | yes |
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |
| `PolicyType` | `string` | no |
| `StepScalingPolicyConfiguration` | `StepScalingPolicyConfiguration` | no |
| `TargetTrackingScalingPolicyConfiguration` | `TargetTrackingScalingPolicyConfiguration` | no |
| `PredictiveScalingPolicyConfiguration` | `PredictiveScalingPolicyConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyARN` | `string` | yes |
| `Alarms` | `List<Alarm>` | no |

## PutScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `Schedule` | `string` | no |
| `Timezone` | `string` | no |
| `ScheduledActionName` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `ScalableTargetAction` | `ScalableTargetAction` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterScalableTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |
| `MinCapacity` | `integer` | no |
| `MaxCapacity` | `integer` | no |
| `RoleARN` | `string` | no |
| `SuspendedState` | `SuspendedState` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalableTargetARN` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


