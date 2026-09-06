# AWS Config

API version: 2014-11-12. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/config/2014-11-12/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderArn` | `string` | yes |
| `ResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorder` | `ConfigurationRecorder` | yes |

## BatchGetAggregateResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `ResourceIdentifiers` | `List<AggregateResourceIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaseConfigurationItems` | `List<BaseConfigurationItem>` | no |
| `UnprocessedResourceIdentifiers` | `List<AggregateResourceIdentifier>` | no |

## BatchGetResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceKeys` | `List<ResourceKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `baseConfigurationItems` | `List<BaseConfigurationItem>` | no |
| `unprocessedResourceKeys` | `List<ResourceKey>` | no |

## DeleteAggregationAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedAccountId` | `string` | yes |
| `AuthorizedAwsRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConformancePack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeliveryChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEvaluationResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOrganizationConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOrganizationConformancePack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePendingAggregationRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequesterAccountId` | `string` | yes |
| `RequesterAwsRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRemediationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRemediationExceptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceKeys` | `List<RemediationExceptionResourceKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedBatches` | `List<FailedDeleteRemediationExceptionsBatch>` | no |

## DeleteResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRetentionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceLinkedConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | no |
| `Arn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |

## DeleteStoredQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeliverConfigSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deliveryChannelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configSnapshotId` | `string` | no |

## DescribeAggregateComplianceByConfigRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `Filters` | `ConfigRuleComplianceFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregateComplianceByConfigRules` | `List<AggregateComplianceByConfigRule>` | no |
| `NextToken` | `string` | no |

## DescribeAggregateComplianceByConformancePacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `Filters` | `AggregateConformancePackComplianceFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregateComplianceByConformancePacks` | `List<AggregateComplianceByConformancePack>` | no |
| `NextToken` | `string` | no |

## DescribeAggregationAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregationAuthorizations` | `List<AggregationAuthorization>` | no |
| `NextToken` | `string` | no |

## DescribeComplianceByConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleNames` | `List<string>` | no |
| `ComplianceTypes` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceByConfigRules` | `List<ComplianceByConfigRule>` | no |
| `NextToken` | `string` | no |

## DescribeComplianceByResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | no |
| `ResourceId` | `string` | no |
| `ComplianceTypes` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceByResources` | `List<ComplianceByResource>` | no |
| `NextToken` | `string` | no |

## DescribeConfigRuleEvaluationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRulesEvaluationStatus` | `List<ConfigRuleEvaluationStatus>` | no |
| `NextToken` | `string` | no |

## DescribeConfigRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleNames` | `List<string>` | no |
| `Filters` | `DescribeConfigRulesFilters` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRules` | `List<ConfigRule>` | no |
| `NextToken` | `string` | no |

## DescribeConfigurationAggregatorSourcesStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `UpdateStatus` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatedSourceStatusList` | `List<AggregatedSourceStatus>` | no |
| `NextToken` | `string` | no |

## DescribeConfigurationAggregators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregators` | `List<ConfigurationAggregator>` | no |
| `NextToken` | `string` | no |

## DescribeConfigurationRecorderStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderNames` | `List<string>` | no |
| `ServicePrincipal` | `string` | no |
| `Arn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecordersStatus` | `List<ConfigurationRecorderStatus>` | no |

## DescribeConfigurationRecorders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderNames` | `List<string>` | no |
| `ServicePrincipal` | `string` | no |
| `Arn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorders` | `List<ConfigurationRecorder>` | no |

## DescribeConformancePackCompliance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |
| `Filters` | `ConformancePackComplianceFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |
| `ConformancePackRuleComplianceList` | `List<ConformancePackRuleCompliance>` | yes |
| `NextToken` | `string` | no |

## DescribeConformancePackStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackStatusDetails` | `List<ConformancePackStatusDetail>` | no |
| `NextToken` | `string` | no |

## DescribeConformancePacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackDetails` | `List<ConformancePackDetail>` | no |
| `NextToken` | `string` | no |

## DescribeDeliveryChannelStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannelNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannelsStatus` | `List<DeliveryChannelStatus>` | no |

## DescribeDeliveryChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannelNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannels` | `List<DeliveryChannel>` | no |

## DescribeOrganizationConfigRuleStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleStatuses` | `List<OrganizationConfigRuleStatus>` | no |
| `NextToken` | `string` | no |

## DescribeOrganizationConfigRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRules` | `List<OrganizationConfigRule>` | no |
| `NextToken` | `string` | no |

## DescribeOrganizationConformancePackStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackStatuses` | `List<OrganizationConformancePackStatus>` | no |
| `NextToken` | `string` | no |

## DescribeOrganizationConformancePacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePacks` | `List<OrganizationConformancePack>` | no |
| `NextToken` | `string` | no |

## DescribePendingAggregationRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PendingAggregationRequests` | `List<PendingAggregationRequest>` | no |
| `NextToken` | `string` | no |

## DescribeRemediationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RemediationConfigurations` | `List<RemediationConfiguration>` | no |

## DescribeRemediationExceptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceKeys` | `List<RemediationExceptionResourceKey>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RemediationExceptions` | `List<RemediationException>` | no |
| `NextToken` | `string` | no |

## DescribeRemediationExecutionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceKeys` | `List<ResourceKey>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RemediationExecutionStatuses` | `List<RemediationExecutionStatus>` | no |
| `NextToken` | `string` | no |

## DescribeRetentionConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionConfigurationNames` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionConfigurations` | `List<RetentionConfiguration>` | no |
| `NextToken` | `string` | no |

## DisassociateResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderArn` | `string` | yes |
| `ResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorder` | `ConfigurationRecorder` | yes |

## GetAggregateComplianceDetailsByConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `ConfigRuleName` | `string` | yes |
| `AccountId` | `string` | yes |
| `AwsRegion` | `string` | yes |
| `ComplianceType` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregateEvaluationResults` | `List<AggregateEvaluationResult>` | no |
| `NextToken` | `string` | no |

## GetAggregateConfigRuleComplianceSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `Filters` | `ConfigRuleComplianceSummaryFilters` | no |
| `GroupByKey` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupByKey` | `string` | no |
| `AggregateComplianceCounts` | `List<AggregateComplianceCount>` | no |
| `NextToken` | `string` | no |

## GetAggregateConformancePackComplianceSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `Filters` | `AggregateConformancePackComplianceSummaryFilters` | no |
| `GroupByKey` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregateConformancePackComplianceSummaries` | `List<AggregateConformancePackComplianceSummary>` | no |
| `GroupByKey` | `string` | no |
| `NextToken` | `string` | no |

## GetAggregateDiscoveredResourceCounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `Filters` | `ResourceCountFilters` | no |
| `GroupByKey` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TotalDiscoveredResources` | `long` | yes |
| `GroupByKey` | `string` | no |
| `GroupedResourceCounts` | `List<GroupedResourceCount>` | no |
| `NextToken` | `string` | no |

## GetAggregateResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `ResourceIdentifier` | `AggregateResourceIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationItem` | `ConfigurationItem` | no |

## GetComplianceDetailsByConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ComplianceTypes` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationResults` | `List<EvaluationResult>` | no |
| `NextToken` | `string` | no |

## GetComplianceDetailsByResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | no |
| `ResourceId` | `string` | no |
| `ComplianceTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `ResourceEvaluationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationResults` | `List<EvaluationResult>` | no |
| `NextToken` | `string` | no |

## GetComplianceSummaryByConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceSummary` | `ComplianceSummary` | no |

## GetComplianceSummaryByResourceType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComplianceSummariesByResourceType` | `List<ComplianceSummaryByResourceType>` | no |

## GetConformancePackComplianceDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |
| `Filters` | `ConformancePackEvaluationFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |
| `ConformancePackRuleEvaluationResults` | `List<ConformancePackEvaluationResult>` | no |
| `NextToken` | `string` | no |

## GetConformancePackComplianceSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackNames` | `List<string>` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackComplianceSummaryList` | `List<ConformancePackComplianceSummary>` | no |
| `NextToken` | `string` | no |

## GetConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connector` | `Connector` | yes |

## GetCustomRulePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyText` | `string` | no |

## GetDiscoveredResourceCounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceTypes` | `List<string>` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `totalDiscoveredResources` | `long` | no |
| `resourceCounts` | `List<ResourceCount>` | no |
| `nextToken` | `string` | no |

## GetOrganizationConfigRuleDetailedStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleName` | `string` | yes |
| `Filters` | `StatusDetailFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleDetailedStatus` | `List<MemberAccountStatus>` | no |
| `NextToken` | `string` | no |

## GetOrganizationConformancePackDetailedStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackName` | `string` | yes |
| `Filters` | `OrganizationResourceDetailedStatusFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackDetailedStatuses` | `List<OrganizationConformancePackDetailedStatus>` | no |
| `NextToken` | `string` | no |

## GetOrganizationCustomRulePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyText` | `string` | no |

## GetResourceConfigHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `resourceId` | `string` | yes |
| `laterTime` | `timestamp` | no |
| `earlierTime` | `timestamp` | no |
| `chronologicalOrder` | `string` | no |
| `limit` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationItems` | `List<ConfigurationItem>` | no |
| `nextToken` | `string` | no |

## GetResourceEvaluationSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceEvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceEvaluationId` | `string` | no |
| `EvaluationMode` | `string` | no |
| `EvaluationStatus` | `EvaluationStatus` | no |
| `EvaluationStartTimestamp` | `timestamp` | no |
| `Compliance` | `string` | no |
| `EvaluationContext` | `EvaluationContext` | no |
| `ResourceDetails` | `ResourceDetails` | no |

## GetStoredQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StoredQuery` | `StoredQuery` | no |

## ListAggregateDiscoveredResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `ResourceType` | `string` | yes |
| `Filters` | `ResourceFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifiers` | `List<AggregateResourceIdentifier>` | no |
| `NextToken` | `string` | no |

## ListConfigurationRecorders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ConfigurationRecorderFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderSummaries` | `List<ConfigurationRecorderSummary>` | yes |
| `NextToken` | `string` | no |

## ListConformancePackComplianceScores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ConformancePackComplianceScoresFilters` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ConformancePackComplianceScores` | `List<ConformancePackComplianceScore>` | yes |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<ConnectorFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorSummaries` | `List<ConnectorSummary>` | yes |
| `NextToken` | `string` | no |

## ListDiscoveredResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `resourceIds` | `List<string>` | no |
| `resourceName` | `string` | no |
| `limit` | `integer` | no |
| `includeDeletedResources` | `boolean` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifiers` | `List<ResourceIdentifier>` | no |
| `nextToken` | `string` | no |

## ListResourceEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ResourceEvaluationFilters` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceEvaluations` | `List<ResourceEvaluation>` | no |
| `NextToken` | `string` | no |

## ListStoredQueries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StoredQueryMetadata` | `List<StoredQueryMetadata>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## PutAggregationAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedAccountId` | `string` | yes |
| `AuthorizedAwsRegion` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregationAuthorization` | `AggregationAuthorization` | no |

## PutConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRule` | `ConfigRule` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregatorName` | `string` | yes |
| `AccountAggregationSources` | `List<AccountAggregationSource>` | no |
| `OrganizationAggregationSource` | `OrganizationAggregationSource` | no |
| `Tags` | `List<Tag>` | no |
| `AggregatorFilters` | `AggregatorFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationAggregator` | `ConfigurationAggregator` | no |

## PutConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorder` | `ConfigurationRecorder` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConformancePack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackName` | `string` | yes |
| `TemplateS3Uri` | `string` | no |
| `TemplateBody` | `string` | no |
| `DeliveryS3Bucket` | `string` | no |
| `DeliveryS3KeyPrefix` | `string` | no |
| `ConformancePackInputParameters` | `List<ConformancePackInputParameter>` | no |
| `TemplateSSMDocumentDetails` | `TemplateSSMDocumentDetails` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConformancePackArn` | `string` | no |

## PutConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorConfiguration` | `ConnectorConfiguration` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

## PutDeliveryChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryChannel` | `DeliveryChannel` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Evaluations` | `List<Evaluation>` | no |
| `ResultToken` | `string` | yes |
| `TestMode` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEvaluations` | `List<Evaluation>` | no |

## PutExternalEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ExternalEvaluation` | `ExternalEvaluation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutOrganizationConfigRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleName` | `string` | yes |
| `OrganizationManagedRuleMetadata` | `OrganizationManagedRuleMetadata` | no |
| `OrganizationCustomRuleMetadata` | `OrganizationCustomRuleMetadata` | no |
| `ExcludedAccounts` | `List<string>` | no |
| `OrganizationCustomPolicyRuleMetadata` | `OrganizationCustomPolicyRuleMetadata` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConfigRuleArn` | `string` | no |

## PutOrganizationConformancePack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackName` | `string` | yes |
| `TemplateS3Uri` | `string` | no |
| `TemplateBody` | `string` | no |
| `DeliveryS3Bucket` | `string` | no |
| `DeliveryS3KeyPrefix` | `string` | no |
| `ConformancePackInputParameters` | `List<ConformancePackInputParameter>` | no |
| `ExcludedAccounts` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationConformancePackArn` | `string` | no |

## PutRemediationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RemediationConfigurations` | `List<RemediationConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedBatches` | `List<FailedRemediationBatch>` | no |

## PutRemediationExceptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceKeys` | `List<RemediationExceptionResourceKey>` | yes |
| `Message` | `string` | no |
| `ExpirationTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedBatches` | `List<FailedRemediationExceptionBatch>` | no |

## PutResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `SchemaVersionId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ResourceName` | `string` | no |
| `Configuration` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRetentionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionPeriodInDays` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionConfiguration` | `RetentionConfiguration` | no |

## PutServiceLinkedConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |

## PutStoredQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StoredQuery` | `StoredQuery` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryArn` | `string` | no |

## PutThirdPartyServiceLinkedConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipal` | `string` | yes |
| `ConnectorArn` | `string` | yes |
| `ScopeConfiguration` | `ScopeConfiguration` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |

## SelectAggregateResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Expression` | `string` | yes |
| `ConfigurationAggregatorName` | `string` | yes |
| `Limit` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<string>` | no |
| `QueryInfo` | `QueryInfo` | no |
| `NextToken` | `string` | no |

## SelectResourceConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Expression` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<string>` | no |
| `QueryInfo` | `QueryInfo` | no |
| `NextToken` | `string` | no |

## StartConfigRulesEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartRemediationExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigRuleName` | `string` | yes |
| `ResourceKeys` | `List<ResourceKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailureMessage` | `string` | no |
| `FailedItems` | `List<ResourceKey>` | no |

## StartResourceEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDetails` | `ResourceDetails` | yes |
| `EvaluationContext` | `EvaluationContext` | no |
| `EvaluationMode` | `string` | yes |
| `EvaluationTimeout` | `integer` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceEvaluationId` | `string` | no |

## StopConfigurationRecorder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationRecorderName` | `string` | yes |

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


