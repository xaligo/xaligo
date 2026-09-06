# AWS Auto Scaling Plans

API version: 2018-01-06. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/autoscaling-plans/2018-01-06/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateScalingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanName` | `string` | yes |
| `ApplicationSource` | `ApplicationSource` | yes |
| `ScalingInstructions` | `List<ScalingInstruction>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanVersion` | `long` | yes |

## DeleteScalingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanName` | `string` | yes |
| `ScalingPlanVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeScalingPlanResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanName` | `string` | yes |
| `ScalingPlanVersion` | `long` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanResources` | `List<ScalingPlanResource>` | no |
| `NextToken` | `string` | no |

## DescribeScalingPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanNames` | `List<string>` | no |
| `ScalingPlanVersion` | `long` | no |
| `ApplicationSources` | `List<ApplicationSource>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlans` | `List<ScalingPlan>` | no |
| `NextToken` | `string` | no |

## GetScalingPlanResourceForecastData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanName` | `string` | yes |
| `ScalingPlanVersion` | `long` | yes |
| `ServiceNamespace` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ScalableDimension` | `string` | yes |
| `ForecastDataType` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Datapoints` | `List<Datapoint>` | yes |

## UpdateScalingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScalingPlanName` | `string` | yes |
| `ScalingPlanVersion` | `long` | yes |
| `ApplicationSource` | `ApplicationSource` | no |
| `ScalingInstructions` | `List<ScalingInstruction>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


