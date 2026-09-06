# Amazon CloudWatch

API version: 2010-08-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudwatch/2010-08-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateDatasetKmsKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetIdentifier` | `string` | yes |
| `KmsKeyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAlarmMuteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmMuteRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAlarms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyDetectorId` | `string` | no |
| `Namespace` | `string` | no |
| `MetricName` | `string` | no |
| `Dimensions` | `List<Dimension>` | no |
| `Stat` | `string` | no |
| `SingleMetricAnomalyDetector` | `SingleMetricAnomalyDetector` | no |
| `MetricMathAnomalyDetector` | `MetricMathAnomalyDetector` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<PartialFailure>` | no |

## DeleteMetricStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAlarmContributors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmContributors` | `List<AlarmContributor>` | yes |
| `NextToken` | `string` | no |

## DescribeAlarmHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | no |
| `AlarmContributorId` | `string` | no |
| `AlarmTypes` | `List<string>` | no |
| `HistoryItemType` | `string` | no |
| `StartDate` | `timestamp` | no |
| `EndDate` | `timestamp` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |
| `ScanBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmHistoryItems` | `List<AlarmHistoryItem>` | no |
| `NextToken` | `string` | no |

## DescribeAlarms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmNames` | `List<string>` | no |
| `AlarmNamePrefix` | `string` | no |
| `AlarmTypes` | `List<string>` | no |
| `ChildrenOfAlarmName` | `string` | no |
| `ParentsOfAlarmName` | `string` | no |
| `StateValue` | `string` | no |
| `ActionPrefix` | `string` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompositeAlarms` | `List<CompositeAlarm>` | no |
| `MetricAlarms` | `List<MetricAlarm>` | no |
| `LogAlarms` | `List<LogAlarm>` | no |
| `NextToken` | `string` | no |

## DescribeAlarmsForMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricName` | `string` | yes |
| `Namespace` | `string` | yes |
| `Statistic` | `string` | no |
| `ExtendedStatistic` | `string` | no |
| `Dimensions` | `List<Dimension>` | no |
| `Period` | `integer` | no |
| `Unit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricAlarms` | `List<MetricAlarm>` | no |

## DescribeAnomalyDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyDetectorIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Namespace` | `string` | no |
| `MetricName` | `string` | no |
| `Dimensions` | `List<Dimension>` | no |
| `AnomalyDetectorTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyDetectors` | `List<AnomalyDetector>` | no |
| `NextToken` | `string` | no |

## DescribeInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `InsightRules` | `List<InsightRule>` | no |

## DisableAlarmActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<PartialFailure>` | no |

## DisassociateDatasetKmsKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableAlarmActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<PartialFailure>` | no |

## GetAlarmMuteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmMuteRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `AlarmMuteRuleArn` | `string` | no |
| `Description` | `string` | no |
| `Rule` | `Rule` | no |
| `MuteTargets` | `MuteTargets` | no |
| `StartDate` | `timestamp` | no |
| `ExpireDate` | `timestamp` | no |
| `Status` | `string` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |
| `MuteType` | `string` | no |

## GetDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardArn` | `string` | no |
| `DashboardBody` | `string` | no |
| `DashboardName` | `string` | no |

## GetDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DatasetId` | `string` | yes |
| `Arn` | `string` | yes |
| `KmsKeyArn` | `string` | no |

## GetInsightRuleReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Period` | `integer` | yes |
| `MaxContributorCount` | `integer` | no |
| `Metrics` | `List<string>` | no |
| `OrderBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyLabels` | `List<string>` | no |
| `AggregationStatistic` | `string` | no |
| `AggregateValue` | `double` | no |
| `ApproximateUniqueCount` | `long` | no |
| `Contributors` | `List<InsightRuleContributor>` | no |
| `MetricDatapoints` | `List<InsightRuleMetricDatapoint>` | no |

## GetMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDataQueries` | `List<MetricDataQuery>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `ScanBy` | `string` | no |
| `MaxDatapoints` | `integer` | no |
| `LabelOptions` | `LabelOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDataResults` | `List<MetricDataResult>` | no |
| `NextToken` | `string` | no |
| `Messages` | `List<MessageData>` | no |

## GetMetricStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `string` | yes |
| `MetricName` | `string` | yes |
| `Dimensions` | `List<Dimension>` | no |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Period` | `integer` | yes |
| `Statistics` | `List<string>` | no |
| `ExtendedStatistics` | `List<string>` | no |
| `Unit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Label` | `string` | no |
| `Datapoints` | `List<Datapoint>` | no |

## GetMetricStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `IncludeFilters` | `List<MetricStreamFilter>` | no |
| `ExcludeFilters` | `List<MetricStreamFilter>` | no |
| `FirehoseArn` | `string` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `LastUpdateDate` | `timestamp` | no |
| `OutputFormat` | `string` | no |
| `StatisticsConfigurations` | `List<MetricStreamStatisticsConfiguration>` | no |
| `IncludeLinkedAccountsMetrics` | `boolean` | no |

## GetMetricWidgetImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricWidget` | `string` | yes |
| `OutputFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricWidgetImage` | `blob` | no |

## GetOTelEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## ListAlarmMuteRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | no |
| `Statuses` | `List<string>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmMuteRuleSummaries` | `List<AlarmMuteRuleSummary>` | no |
| `NextToken` | `string` | no |

## ListDashboards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardNamePrefix` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardEntries` | `List<DashboardEntry>` | no |
| `NextToken` | `string` | no |

## ListManagedInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedRules` | `List<ManagedRuleDescription>` | no |
| `NextToken` | `string` | no |

## ListMetricStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Entries` | `List<MetricStreamEntry>` | no |

## ListMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `string` | no |
| `MetricName` | `string` | no |
| `Dimensions` | `List<DimensionFilter>` | no |
| `NextToken` | `string` | no |
| `RecentlyActive` | `string` | no |
| `IncludeLinkedAccounts` | `boolean` | no |
| `OwningAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metrics` | `List<Metric>` | no |
| `NextToken` | `string` | no |
| `OwningAccounts` | `List<string>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutAlarmMuteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Rule` | `Rule` | yes |
| `MuteTargets` | `MuteTargets` | no |
| `Tags` | `List<Tag>` | no |
| `StartDate` | `timestamp` | no |
| `ExpireDate` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAnomalyDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `string` | no |
| `MetricName` | `string` | no |
| `Dimensions` | `List<Dimension>` | no |
| `Stat` | `string` | no |
| `Configuration` | `AnomalyDetectorConfiguration` | no |
| `MetricCharacteristics` | `MetricCharacteristics` | no |
| `SingleMetricAnomalyDetector` | `SingleMetricAnomalyDetector` | no |
| `MetricMathAnomalyDetector` | `MetricMathAnomalyDetector` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnomalyDetectorId` | `string` | no |

## PutCompositeAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionsEnabled` | `boolean` | no |
| `AlarmActions` | `List<string>` | no |
| `AlarmDescription` | `string` | no |
| `AlarmName` | `string` | yes |
| `AlarmRule` | `string` | yes |
| `InsufficientDataActions` | `List<string>` | no |
| `OKActions` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `ActionsSuppressor` | `string` | no |
| `ActionsSuppressorWaitPeriod` | `integer` | no |
| `ActionsSuppressorExtensionPeriod` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDashboard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardName` | `string` | yes |
| `DashboardBody` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardValidationMessages` | `List<DashboardValidationMessage>` | no |

## PutInsightRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `RuleState` | `string` | no |
| `RuleDefinition` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ApplyOnTransformedLogs` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutLogAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | yes |
| `AlarmDescription` | `string` | no |
| `ScheduledQueryConfiguration` | `ScheduledQueryConfiguration` | yes |
| `ActionLogLineCount` | `integer` | no |
| `ActionLogLineRoleArn` | `string` | no |
| `ActionsEnabled` | `boolean` | no |
| `OKActions` | `List<string>` | no |
| `AlarmActions` | `List<string>` | no |
| `InsufficientDataActions` | `List<string>` | no |
| `QueryResultsToEvaluate` | `integer` | yes |
| `QueryResultsToAlarm` | `integer` | yes |
| `Threshold` | `double` | yes |
| `ComparisonOperator` | `string` | yes |
| `TreatMissingData` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `WarmUpConfiguration` | `WarmUpConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutManagedInsightRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedRules` | `List<ManagedRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failures` | `List<PartialFailure>` | no |

## PutMetricAlarm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | yes |
| `AlarmDescription` | `string` | no |
| `ActionsEnabled` | `boolean` | no |
| `OKActions` | `List<string>` | no |
| `AlarmActions` | `List<string>` | no |
| `InsufficientDataActions` | `List<string>` | no |
| `MetricName` | `string` | no |
| `Namespace` | `string` | no |
| `Statistic` | `string` | no |
| `ExtendedStatistic` | `string` | no |
| `Dimensions` | `List<Dimension>` | no |
| `Period` | `integer` | no |
| `Unit` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `DatapointsToAlarm` | `integer` | no |
| `Threshold` | `double` | no |
| `ComparisonOperator` | `string` | no |
| `TreatMissingData` | `string` | no |
| `EvaluateLowSampleCountPercentile` | `string` | no |
| `Metrics` | `List<MetricDataQuery>` | no |
| `Tags` | `List<Tag>` | no |
| `ThresholdMetricId` | `string` | no |
| `EvaluationWindow` | `EvaluationWindow` | no |
| `WarmUpConfiguration` | `WarmUpConfiguration` | no |
| `EvaluationCriteria` | `EvaluationCriteria` | no |
| `EvaluationInterval` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `string` | yes |
| `MetricData` | `List<MetricDatum>` | no |
| `EntityMetricData` | `List<EntityMetricData>` | no |
| `StrictEntityValidation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMetricStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IncludeFilters` | `List<MetricStreamFilter>` | no |
| `ExcludeFilters` | `List<MetricStreamFilter>` | no |
| `FirehoseArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `OutputFormat` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `StatisticsConfigurations` | `List<MetricStreamStatisticsConfiguration>` | no |
| `IncludeLinkedAccountsMetrics` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## SetAlarmState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlarmName` | `string` | yes |
| `StateValue` | `string` | yes |
| `StateReason` | `string` | yes |
| `StateReasonData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMetricStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartOTelEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopMetricStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopOTelEnrichment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


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


