# Amazon CloudWatch Application Signals

API version: 2024-04-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/application-signals/2024-04-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteInstrumentationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionTarget` | `BatchDeleteDeletionTarget` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletedCount` | `integer` | yes |
| `SuccessfulDeletions` | `List<BatchDeleteSuccessfulDeletion>` | yes |
| `Errors` | `List<BatchDeleteError>` | yes |

## BatchGetServiceLevelObjectiveBudgetReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Timestamp` | `timestamp` | yes |
| `SloIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Timestamp` | `timestamp` | yes |
| `Reports` | `List<ServiceLevelObjectiveBudgetReport>` | yes |
| `Errors` | `List<ServiceLevelObjectiveBudgetReportError>` | yes |

## BatchUpdateExclusionWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SloIds` | `List<string>` | yes |
| `AddExclusionWindows` | `List<ExclusionWindow>` | no |
| `RemoveExclusionWindows` | `List<ExclusionWindow>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SloIds` | `List<string>` | yes |
| `Errors` | `List<BatchUpdateExclusionWindowsError>` | yes |

## CreateInstrumentationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstrumentationType` | `string` | yes |
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `Location` | `Location` | yes |
| `Description` | `string` | no |
| `ExpiresAt` | `timestamp` | no |
| `AttributeFilters` | `List<Map<string>>` | no |
| `CaptureConfiguration` | `CaptureConfiguration` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstrumentationType` | `string` | yes |
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `Location` | `Location` | yes |
| `LocationHash` | `string` | yes |
| `Description` | `string` | no |
| `ExpiresAt` | `timestamp` | no |
| `AttributeFilters` | `List<Map<string>>` | no |
| `CaptureConfiguration` | `CaptureConfiguration` | yes |
| `CreatedAt` | `timestamp` | yes |
| `ARN` | `string` | yes |

## CreateServiceLevelObjective

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `SliConfig` | `ServiceLevelIndicatorConfig` | no |
| `RequestBasedSliConfig` | `RequestBasedServiceLevelIndicatorConfig` | no |
| `Goal` | `Goal` | no |
| `Tags` | `List<Tag>` | no |
| `BurnRateConfigurations` | `List<BurnRateConfiguration>` | no |
| `CreateRecommendedSlo` | `boolean` | no |
| `AutoInvestigationEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Slo` | `ServiceLevelObjective` | yes |

## DeleteGroupingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstrumentationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstrumentationType` | `string` | yes |
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `LocationIdentifier` | `LocationIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionStatus` | `string` | yes |

## DeleteServiceLevelObjective

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetInstrumentationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstrumentationType` | `string` | yes |
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `LocationIdentifier` | `LocationIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `InstrumentationConfiguration` | yes |

## GetInstrumentationConfigurationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstrumentationType` | `string` | yes |
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `LocationIdentifier` | `LocationIdentifier` | yes |
| `Status` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `SignalType` | `string` | yes |
| `Location` | `Location` | yes |
| `Status` | `string` | yes |
| `Events` | `List<InstrumentationStatusEvent>` | yes |
| `NextToken` | `string` | no |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `KeyAttributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `LogGroupReferences` | `List<Map<string>>` | no |

## GetServiceLevelObjective

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Slo` | `ServiceLevelObjective` | yes |

## ListAuditFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Auditors` | `List<string>` | no |
| `AuditTargets` | `List<AuditTarget>` | yes |
| `DetailLevel` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `AuditFindings` | `List<AuditFinding>` | yes |
| `NextToken` | `string` | no |

## ListEntityEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entity` | `Map<string>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ChangeEvents` | `List<ChangeEvent>` | yes |
| `NextToken` | `string` | no |

## ListGroupingAttributeDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AwsAccountId` | `string` | no |
| `IncludeLinkedAccounts` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupingAttributeDefinitions` | `List<GroupingAttributeDefinition>` | yes |
| `UpdatedAt` | `timestamp` | no |
| `NextToken` | `string` | no |

## ListInstrumentationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `InstrumentationType` | `string` | yes |
| `SyncedAt` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `Changed` | `boolean` | yes |
| `LatestConfigurations` | `List<InstrumentationConfigurationWithoutServiceEnv>` | no |
| `SyncedAt` | `timestamp` | yes |
| `SyncInterval` | `integer` | yes |
| `NextToken` | `string` | no |

## ListServiceDependencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `KeyAttributes` | `Map<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ServiceDependencies` | `List<ServiceDependency>` | yes |
| `NextToken` | `string` | no |

## ListServiceDependents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `KeyAttributes` | `Map<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ServiceDependents` | `List<ServiceDependent>` | yes |
| `NextToken` | `string` | no |

## ListServiceLevelObjectiveExclusionWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExclusionWindows` | `List<ExclusionWindow>` | yes |
| `NextToken` | `string` | no |

## ListServiceLevelObjectives

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyAttributes` | `Map<string>` | no |
| `OperationName` | `string` | no |
| `DependencyConfig` | `DependencyConfig` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `MetricSourceTypes` | `List<string>` | no |
| `IncludeLinkedAccounts` | `boolean` | no |
| `SloOwnerAwsAccountId` | `string` | no |
| `MetricSource` | `MetricSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SloSummaries` | `List<ServiceLevelObjectiveSummary>` | no |
| `NextToken` | `string` | no |

## ListServiceOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `KeyAttributes` | `Map<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ServiceOperations` | `List<ServiceOperation>` | yes |
| `NextToken` | `string` | no |

## ListServiceStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeLinkedAccounts` | `boolean` | no |
| `AwsAccountId` | `string` | no |
| `AttributeFilters` | `List<AttributeFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ServiceStates` | `List<ServiceState>` | yes |
| `NextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeLinkedAccounts` | `boolean` | no |
| `AwsAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `ServiceSummaries` | `List<ServiceSummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutGroupingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupingAttributeDefinitions` | `List<GroupingAttributeDefinition>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupingConfiguration` | `GroupingConfiguration` | yes |

## ReportInstrumentationConfigurationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `Configurations` | `List<InstrumentationConfigurationStatusReport>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `string` | yes |
| `Environment` | `string` | yes |
| `UnprocessedStatusEvents` | `List<UnprocessedStatusEvent>` | yes |

## StartDiscovery

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


## UpdateServiceLevelObjective

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `SliConfig` | `ServiceLevelIndicatorConfig` | no |
| `RequestBasedSliConfig` | `RequestBasedServiceLevelIndicatorConfig` | no |
| `Goal` | `Goal` | no |
| `BurnRateConfigurations` | `List<BurnRateConfiguration>` | no |
| `AutoInvestigationEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Slo` | `ServiceLevelObjective` | yes |

