# Amazon CloudWatch Application Insights

API version: 2018-11-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/application-insights/2018-11-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `WorkloadConfiguration` | `WorkloadConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadConfiguration` | `WorkloadConfiguration` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `OpsCenterEnabled` | `boolean` | no |
| `CWEMonitorEnabled` | `boolean` | no |
| `OpsItemSNSTopicArn` | `string` | no |
| `SNSNotificationArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `AutoConfigEnabled` | `boolean` | no |
| `AutoCreate` | `boolean` | no |
| `GroupingType` | `string` | no |
| `AttachMissingPermission` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationInfo` | `ApplicationInfo` | no |

## CreateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `ResourceList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLogPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `PatternSetName` | `string` | yes |
| `PatternName` | `string` | yes |
| `Pattern` | `string` | yes |
| `Rank` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogPattern` | `LogPattern` | no |
| `ResourceGroupName` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLogPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `PatternSetName` | `string` | yes |
| `PatternName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationInfo` | `ApplicationInfo` | no |

## DescribeComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationComponent` | `ApplicationComponent` | no |
| `ResourceList` | `List<string>` | no |

## DescribeComponentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Monitor` | `boolean` | no |
| `Tier` | `string` | no |
| `ComponentConfiguration` | `string` | no |

## DescribeComponentConfigurationRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `Tier` | `string` | yes |
| `WorkloadName` | `string` | no |
| `RecommendationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComponentConfiguration` | `string` | no |

## DescribeLogPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `PatternSetName` | `string` | yes |
| `PatternName` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `AccountId` | `string` | no |
| `LogPattern` | `LogPattern` | no |

## DescribeObservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservationId` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Observation` | `Observation` | no |

## DescribeProblem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProblemId` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Problem` | `Problem` | no |
| `SNSNotificationArn` | `string` | no |

## DescribeProblemObservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProblemId` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelatedObservations` | `RelatedObservations` | no |

## DescribeWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `WorkloadId` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadRemarks` | `string` | no |
| `WorkloadConfiguration` | `WorkloadConfiguration` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationInfoList` | `List<ApplicationInfo>` | no |
| `NextToken` | `string` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationComponentList` | `List<ApplicationComponent>` | no |
| `NextToken` | `string` | no |

## ListConfigurationHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `EventStatus` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventList` | `List<ConfigurationEvent>` | no |
| `NextToken` | `string` | no |

## ListLogPatternSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `AccountId` | `string` | no |
| `LogPatternSets` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListLogPatterns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `PatternSetName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `AccountId` | `string` | no |
| `LogPatterns` | `List<LogPattern>` | no |
| `NextToken` | `string` | no |

## ListProblems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `ResourceGroupName` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ComponentName` | `string` | no |
| `Visibility` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProblemList` | `List<Problem>` | no |
| `NextToken` | `string` | no |
| `ResourceGroupName` | `string` | no |
| `AccountId` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListWorkloads

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadList` | `List<Workload>` | no |
| `NextToken` | `string` | no |

## RemoveWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `WorkloadId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `OpsCenterEnabled` | `boolean` | no |
| `CWEMonitorEnabled` | `boolean` | no |
| `OpsItemSNSTopicArn` | `string` | no |
| `SNSNotificationArn` | `string` | no |
| `RemoveSNSTopic` | `boolean` | no |
| `AutoConfigEnabled` | `boolean` | no |
| `AttachMissingPermission` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationInfo` | `ApplicationInfo` | no |

## UpdateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `NewComponentName` | `string` | no |
| `ResourceList` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateComponentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `Monitor` | `boolean` | no |
| `Tier` | `string` | no |
| `ComponentConfiguration` | `string` | no |
| `AutoConfigEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLogPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `PatternSetName` | `string` | yes |
| `PatternName` | `string` | yes |
| `Pattern` | `string` | no |
| `Rank` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | no |
| `LogPattern` | `LogPattern` | no |

## UpdateProblem

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProblemId` | `string` | yes |
| `UpdateStatus` | `string` | no |
| `Visibility` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceGroupName` | `string` | yes |
| `ComponentName` | `string` | yes |
| `WorkloadId` | `string` | no |
| `WorkloadConfiguration` | `WorkloadConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkloadId` | `string` | no |
| `WorkloadConfiguration` | `WorkloadConfiguration` | no |

