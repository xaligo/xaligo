# AWS SecurityHub

API version: 2018-10-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/securityhub/2018-10-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAdministratorInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdministratorId` | `string` | yes |
| `InvitationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AcceptInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MasterId` | `string` | yes |
| `InvitationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDeleteAutomationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationRulesArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedAutomationRules` | `List<string>` | no |
| `UnprocessedAutomationRules` | `List<UnprocessedAutomationRule>` | no |

## BatchDisableStandards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptionArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptions` | `List<StandardsSubscription>` | no |

## BatchEnableStandards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptionRequests` | `List<StandardsSubscriptionRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptions` | `List<StandardsSubscription>` | no |

## BatchGetAutomationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationRulesArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<AutomationRulesConfig>` | no |
| `UnprocessedAutomationRules` | `List<UnprocessedAutomationRule>` | no |

## BatchGetConfigurationPolicyAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyAssociationIdentifiers` | `List<ConfigurationPolicyAssociation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyAssociations` | `List<ConfigurationPolicyAssociationSummary>` | no |
| `UnprocessedConfigurationPolicyAssociations` | `List<UnprocessedConfigurationPolicyAssociation>` | no |

## BatchGetSecurityControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControls` | `List<SecurityControl>` | yes |
| `UnprocessedIds` | `List<UnprocessedSecurityControl>` | no |

## BatchGetStandardsControlAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsControlAssociationIds` | `List<StandardsControlAssociationId>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsControlAssociationDetails` | `List<StandardsControlAssociationDetail>` | yes |
| `UnprocessedAssociations` | `List<UnprocessedStandardsControlAssociation>` | no |

## BatchImportFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Findings` | `List<AwsSecurityFinding>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedCount` | `integer` | yes |
| `SuccessCount` | `integer` | yes |
| `FailedFindings` | `List<ImportFindingsError>` | no |

## BatchUpdateAutomationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateAutomationRulesRequestItems` | `List<UpdateAutomationRulesRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedAutomationRules` | `List<string>` | no |
| `UnprocessedAutomationRules` | `List<UnprocessedAutomationRule>` | no |

## BatchUpdateFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingIdentifiers` | `List<AwsSecurityFindingIdentifier>` | yes |
| `Note` | `NoteUpdate` | no |
| `Severity` | `SeverityUpdate` | no |
| `VerificationState` | `string` | no |
| `Confidence` | `integer` | no |
| `Criticality` | `integer` | no |
| `Types` | `List<string>` | no |
| `UserDefinedFields` | `Map<string>` | no |
| `Workflow` | `WorkflowUpdate` | no |
| `RelatedFindings` | `List<RelatedFinding>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedFindings` | `List<AwsSecurityFindingIdentifier>` | yes |
| `UnprocessedFindings` | `List<BatchUpdateFindingsUnprocessedFinding>` | yes |

## BatchUpdateFindingsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetadataUids` | `List<string>` | no |
| `FindingIdentifiers` | `List<OcsfFindingIdentifier>` | no |
| `Comment` | `string` | no |
| `SeverityId` | `integer` | no |
| `StatusId` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessedFindings` | `List<BatchUpdateFindingsV2ProcessedFinding>` | yes |
| `UnprocessedFindings` | `List<BatchUpdateFindingsV2UnprocessedFinding>` | yes |

## BatchUpdateStandardsControlAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsControlAssociationUpdates` | `List<StandardsControlAssociationUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAssociationUpdates` | `List<UnprocessedStandardsControlAssociationUpdate>` | no |

## CreateActionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargetArn` | `string` | yes |

## CreateAggregatorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionLinkingMode` | `string` | yes |
| `LinkedRegions` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | no |
| `AggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `LinkedRegions` | `List<string>` | no |

## CreateAutomationRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |
| `RuleStatus` | `string` | no |
| `RuleOrder` | `integer` | yes |
| `RuleName` | `string` | yes |
| `Description` | `string` | yes |
| `IsTerminal` | `boolean` | no |
| `Criteria` | `AutomationRulesFindingFilters` | yes |
| `Actions` | `List<AutomationRulesAction>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## CreateAutomationRuleV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleName` | `string` | yes |
| `RuleStatus` | `string` | no |
| `Description` | `string` | yes |
| `RuleOrder` | `float` | yes |
| `Criteria` | `Criteria` | yes |
| `Actions` | `List<AutomationRulesActionV2>` | yes |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |
| `RuleId` | `string` | no |

## CreateConfigurationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ConfigurationPolicy` | `Policy` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ConfigurationPolicy` | `Policy` | no |

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Provider` | `CspmProviderConfiguration` | yes |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `ConnectorId` | `string` | yes |
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

## CreateConnectorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Provider` | `ProviderConfiguration` | yes |
| `KmsKeyArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `ConnectorId` | `string` | yes |
| `AuthUrl` | `string` | no |
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

## CreateFindingAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionLinkingMode` | `string` | yes |
| `Regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | no |
| `FindingAggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `Regions` | `List<string>` | no |

## CreateInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Filters` | `AwsSecurityFindingFilters` | yes |
| `GroupByAttribute` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArn` | `string` | yes |

## CreateMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountDetails` | `List<AccountDetails>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<Result>` | no |

## CreateTicketV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `FindingMetadataUid` | `string` | yes |
| `ClientToken` | `string` | no |
| `Mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TicketId` | `string` | yes |
| `TicketSrcUrl` | `string` | no |

## DeclineInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<Result>` | no |

## DeleteActionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargetArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargetArn` | `string` | yes |

## DeleteAggregatorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAutomationRuleV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnablementStatus` | `string` | no |

## DeleteConnectorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnablementStatus` | `string` | no |

## DeleteFindingAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArn` | `string` | yes |

## DeleteInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<Result>` | no |

## DeleteMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<Result>` | no |

## DescribeActionTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargetArns` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargets` | `List<ActionTarget>` | yes |
| `NextToken` | `string` | no |

## DescribeHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | no |
| `SubscribedAt` | `string` | no |
| `AutoEnableControls` | `boolean` | no |
| `ControlFindingGenerator` | `string` | no |

## DescribeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnable` | `boolean` | no |
| `MemberAccountLimitReached` | `boolean` | no |
| `AutoEnableStandards` | `string` | no |
| `OrganizationConfiguration` | `OrganizationConfiguration` | no |

## DescribeProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ProductArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Products` | `List<Product>` | yes |
| `NextToken` | `string` | no |

## DescribeProductsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductsV2` | `List<ProductV2>` | yes |
| `NextToken` | `string` | no |

## DescribeSecurityHubV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubV2Arn` | `string` | no |
| `SubscribedAt` | `string` | no |
| `Features` | `Map<FeatureDetail>` | no |

## DescribeStandards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Providers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Standards` | `List<Standard>` | no |
| `NextToken` | `string` | no |

## DescribeStandardsControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptionArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Controls` | `List<StandardsControl>` | no |
| `NextToken` | `string` | no |

## DisableImportFindingsForProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductSubscriptionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | yes |
| `Feature` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableSecurityHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableSecurityHubFeatureV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableSecurityHubV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableImportFindingsForProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductSubscriptionArn` | `string` | no |

## EnableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | yes |
| `Feature` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | no |
| `Feature` | `string` | no |

## EnableSecurityHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |
| `EnableDefaultStandards` | `boolean` | no |
| `ControlFindingGenerator` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableSecurityHubFeatureV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableSecurityHubV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubV2Arn` | `string` | no |

## GenerateRecommendedPolicyV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetadataUid` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Administrator` | `Invitation` | no |

## GetAggregatorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | no |
| `AggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `LinkedRegions` | `List<string>` | no |

## GetAutomationRuleV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |
| `RuleId` | `string` | no |
| `RuleOrder` | `float` | no |
| `RuleName` | `string` | no |
| `RuleStatus` | `string` | no |
| `Description` | `string` | no |
| `Criteria` | `Criteria` | no |
| `Actions` | `List<AutomationRulesActionV2>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## GetConfigurationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ConfigurationPolicy` | `Policy` | no |

## GetConfigurationPolicyAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `Target` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyId` | `string` | no |
| `TargetId` | `string` | no |
| `TargetType` | `string` | no |
| `AssociationType` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `AssociationStatus` | `string` | no |
| `AssociationStatusMessage` | `string` | no |

## GetConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | no |
| `ConnectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Health` | `CspmHealthCheck` | yes |
| `ProviderDetail` | `CspmProviderDetail` | yes |
| `CreatedBy` | `string` | no |
| `EnablementStatus` | `string` | no |

## GetConnectorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | no |
| `ConnectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Health` | `HealthCheck` | yes |
| `ProviderDetail` | `ProviderDetail` | yes |
| `EnablementStatus` | `string` | no |
| `EnablementStatusReason` | `string` | no |

## GetEnabledStandards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptionArns` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Providers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsSubscriptions` | `List<StandardsSubscription>` | no |
| `NextToken` | `string` | no |

## GetFindingAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | no |
| `FindingAggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `Regions` | `List<string>` | no |

## GetFindingHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingIdentifier` | `AwsSecurityFindingIdentifier` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Records` | `List<FindingHistoryRecord>` | no |
| `NextToken` | `string` | no |

## GetFindingStatisticsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupByRules` | `List<GroupByRule>` | yes |
| `Scopes` | `FindingScopes` | no |
| `SortOrder` | `string` | no |
| `MaxStatisticResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupByResults` | `List<GroupByResult>` | no |

## GetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `AwsSecurityFindingFilters` | no |
| `SortCriteria` | `List<SortCriterion>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Findings` | `List<AwsSecurityFinding>` | yes |
| `NextToken` | `string` | no |

## GetFindingsTrendsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `FindingsTrendsFilters` | no |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Granularity` | `string` | yes |
| `TrendsMetrics` | `List<TrendsMetricsResult>` | yes |
| `NextToken` | `string` | no |

## GetFindingsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `OcsfFindingFilters` | no |
| `Scopes` | `FindingScopes` | no |
| `SortCriteria` | `List<SortCriterion>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Findings` | `List<OcsfFinding>` | no |
| `NextToken` | `string` | no |

## GetInsightResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightResults` | `InsightResults` | yes |

## GetInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArns` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Insights` | `List<Insight>` | yes |
| `NextToken` | `string` | no |

## GetInvitationsCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvitationsCount` | `integer` | no |

## GetMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Master` | `Invitation` | no |

## GetMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<Member>` | no |
| `UnprocessedAccounts` | `List<Result>` | no |

## GetRecommendedPolicyV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetadataUid` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecommendationType` | `string` | no |
| `RecommendationSteps` | `List<RecommendationStep>` | no |
| `Error` | `RecommendationError` | no |
| `Status` | `string` | no |
| `ResourceArn` | `string` | no |

## GetResourcesStatisticsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupByRules` | `List<ResourceGroupByRule>` | yes |
| `Scopes` | `ResourceScopes` | no |
| `SortOrder` | `string` | no |
| `MaxStatisticResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupByResults` | `List<GroupByResult>` | yes |

## GetResourcesTrendsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ResourcesTrendsFilters` | no |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Granularity` | `string` | yes |
| `TrendsMetrics` | `List<ResourcesTrendsMetricsResult>` | yes |
| `NextToken` | `string` | no |

## GetResourcesV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `ResourcesFilters` | no |
| `Scopes` | `ResourceScopes` | no |
| `SortCriteria` | `List<SortCriterion>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<ResourceResult>` | yes |
| `NextToken` | `string` | no |

## GetSecurityControlDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlDefinition` | `SecurityControlDefinition` | yes |

## InviteMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<Result>` | no |

## ListAggregatorsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorsV2` | `List<AggregatorV2>` | no |
| `NextToken` | `string` | no |

## ListAutomationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutomationRulesMetadata` | `List<AutomationRulesMetadata>` | no |
| `NextToken` | `string` | no |

## ListAutomationRulesV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<AutomationRulesMetadataV2>` | no |
| `NextToken` | `string` | no |

## ListConfigurationPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicySummaries` | `List<ConfigurationPolicySummary>` | no |
| `NextToken` | `string` | no |

## ListConfigurationPolicyAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `AssociationFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyAssociationSummaries` | `List<ConfigurationPolicyAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ProviderName` | `string` | no |
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Connectors` | `List<CspmConnectorSummary>` | yes |

## ListConnectorsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ProviderName` | `string` | no |
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Connectors` | `List<ConnectorSummary>` | yes |

## ListEnabledProductsForImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductSubscriptions` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListFindingAggregators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregators` | `List<FindingAggregator>` | no |
| `NextToken` | `string` | no |

## ListFreeTrialStatusesV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | no |
| `Statuses` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountFreeTrialStatuses` | `List<AccountFreeTrialStatus>` | yes |
| `NextToken` | `string` | no |

## ListInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invitations` | `List<Invitation>` | no |
| `NextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OnlyAssociated` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<Member>` | no |
| `NextToken` | `string` | no |

## ListOrganizationAdminAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Feature` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccounts` | `List<AdminAccount>` | no |
| `NextToken` | `string` | no |
| `Feature` | `string` | no |

## ListSecurityControlDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Providers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlDefinitions` | `List<SecurityControlDefinition>` | yes |
| `NextToken` | `string` | no |

## ListStandardsControlAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsControlAssociationSummaries` | `List<StandardsControlAssociationSummary>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## RegisterConnectorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthCode` | `string` | yes |
| `AuthState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | no |
| `ConnectorId` | `string` | yes |

## StartConfigurationPolicyAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyIdentifier` | `string` | yes |
| `Target` | `Target` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationPolicyId` | `string` | no |
| `TargetId` | `string` | no |
| `TargetType` | `string` | no |
| `AssociationType` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `AssociationStatus` | `string` | no |
| `AssociationStatusMessage` | `string` | no |

## StartConfigurationPolicyDisassociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Target` | `Target` | no |
| `ConfigurationPolicyIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateActionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionTargetArn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAggregatorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | yes |
| `RegionLinkingMode` | `string` | yes |
| `LinkedRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AggregatorV2Arn` | `string` | no |
| `AggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `LinkedRegions` | `List<string>` | no |

## UpdateAutomationRuleV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `RuleStatus` | `string` | no |
| `RuleOrder` | `float` | no |
| `Description` | `string` | no |
| `RuleName` | `string` | no |
| `Criteria` | `Criteria` | no |
| `Actions` | `List<AutomationRulesActionV2>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `UpdatedReason` | `string` | no |
| `ConfigurationPolicy` | `Policy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `UpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `ConfigurationPolicy` | `Policy` | no |

## UpdateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `Description` | `string` | no |
| `Provider` | `CspmProviderUpdateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

## UpdateConnectorV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `Description` | `string` | no |
| `Provider` | `ProviderUpdateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorStatus` | `string` | no |
| `EnablementStatus` | `string` | no |

## UpdateFindingAggregator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | yes |
| `RegionLinkingMode` | `string` | yes |
| `Regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingAggregatorArn` | `string` | no |
| `FindingAggregationRegion` | `string` | no |
| `RegionLinkingMode` | `string` | no |
| `Regions` | `List<string>` | no |

## UpdateFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `AwsSecurityFindingFilters` | yes |
| `Note` | `NoteUpdate` | no |
| `RecordState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InsightArn` | `string` | yes |
| `Name` | `string` | no |
| `Filters` | `AwsSecurityFindingFilters` | no |
| `GroupByAttribute` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnable` | `boolean` | yes |
| `AutoEnableStandards` | `string` | no |
| `OrganizationConfiguration` | `OrganizationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSecurityControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityControlId` | `string` | yes |
| `Parameters` | `Map<ParameterConfiguration>` | yes |
| `LastUpdateReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSecurityHubConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnableControls` | `boolean` | no |
| `ControlFindingGenerator` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStandardsControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardsControlArn` | `string` | yes |
| `ControlStatus` | `string` | no |
| `DisabledReason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


