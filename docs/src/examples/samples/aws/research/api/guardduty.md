# Amazon GuardDuty

API version: 2017-11-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/guardduty/2017-11-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAdministratorInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AdministratorId` | `string` | yes |
| `InvitationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AcceptInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MasterId` | `string` | yes |
| `InvitationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ArchiveFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCustomDetectionRuleAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `Mode` | `string` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleAssociation` | `AssociationDetail` | yes |

## CreateCustomDetectionRuleOrgConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `Mode` | `string` | yes |
| `IncludeAccountIds` | `List<string>` | no |
| `ExcludeAccountIds` | `List<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Enable` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `FindingPublishingFrequency` | `string` | no |
| `DataSources` | `DataSourceConfigurations` | no |
| `Tags` | `Map<string>` | no |
| `Features` | `List<DetectorFeatureConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | no |
| `UnprocessedDataSources` | `UnprocessedDataSourcesResult` | no |

## CreateFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Action` | `string` | no |
| `Rank` | `integer` | no |
| `FindingCriteria` | `FindingCriteria` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## CreateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `Activate` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ExpectedBucketOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpSetId` | `string` | yes |

## CreateInvestigation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `TriggerPrompt` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvestigationId` | `string` | yes |

## CreateMalwareProtectionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Role` | `string` | yes |
| `ProtectedResource` | `CreateProtectedResource` | yes |
| `Actions` | `MalwareProtectionPlanActions` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MalwareProtectionPlanId` | `string` | no |

## CreateMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountDetails` | `List<AccountDetail>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## CreatePublishingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `DestinationType` | `string` | yes |
| `DestinationProperties` | `DestinationProperties` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationId` | `string` | yes |

## CreateSampleFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateThreatEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `ExpectedBucketOwner` | `string` | no |
| `Activate` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThreatEntitySetId` | `string` | yes |

## CreateThreatIntelSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `Activate` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ExpectedBucketOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThreatIntelSetId` | `string` | yes |

## CreateTrustedEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `ExpectedBucketOwner` | `string` | no |
| `Activate` | `boolean` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedEntitySetId` | `string` | yes |

## DeclineInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## DeleteCustomDetectionRuleAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomDetectionRuleOrgConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `Mode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `IpSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## DeleteMalwareProtectionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MalwareProtectionPlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## DeletePublishingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `DestinationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThreatEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatEntitySetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThreatIntelSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatIntelSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrustedEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `TrustedEntitySetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeMalwareScans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `SortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scans` | `List<Scan>` | yes |
| `NextToken` | `string` | no |

## DescribeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnable` | `boolean` | no |
| `MemberAccountLimitReached` | `boolean` | yes |
| `DataSources` | `OrganizationDataSourceConfigurationsResult` | no |
| `Features` | `List<OrganizationFeatureConfigurationResult>` | no |
| `NextToken` | `string` | no |
| `AutoEnableOrganizationMembers` | `string` | no |

## DescribePublishingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `DestinationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationId` | `string` | yes |
| `DestinationType` | `string` | yes |
| `Status` | `string` | yes |
| `PublishingFailureStartTimestamp` | `long` | yes |
| `DestinationProperties` | `DestinationProperties` | yes |
| `Tags` | `Map<string>` | no |

## DisableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFromMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## EnableOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAdministratorAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Administrator` | `Administrator` | yes |

## GetCoverageStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FilterCriteria` | `CoverageFilterCriteria` | no |
| `StatisticsType` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoverageStatistics` | `CoverageStatistics` | no |

## GetCustomDetectionRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `RuleDetail` | yes |

## GetCustomDetectionRuleAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleAssociation` | `AssociationDetail` | yes |
| `Tags` | `Map<string>` | no |

## GetCustomDetectionRuleOrgConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `Mode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `DetectionRuleOrgConfiguration` | yes |

## GetDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAt` | `string` | no |
| `FindingPublishingFrequency` | `string` | no |
| `ServiceRole` | `string` | yes |
| `Status` | `string` | yes |
| `UpdatedAt` | `string` | no |
| `DataSources` | `DataSourceConfigurationsResult` | no |
| `Tags` | `Map<string>` | no |
| `Features` | `List<DetectorFeatureConfigurationResult>` | no |

## GetFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Action` | `string` | yes |
| `Rank` | `integer` | no |
| `FindingCriteria` | `FindingCriteria` | yes |
| `Tags` | `Map<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Version` | `long` | no |

## GetFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingIds` | `List<string>` | yes |
| `SortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Findings` | `List<Finding>` | yes |

## GetFindingsStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingStatisticTypes` | `List<string>` | no |
| `FindingCriteria` | `FindingCriteria` | no |
| `GroupBy` | `string` | no |
| `OrderBy` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingStatistics` | `FindingStatistics` | yes |
| `NextToken` | `string` | no |

## GetIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `IpSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `ExpectedBucketOwner` | `string` | no |

## GetInvestigation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `InvestigationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Investigation` | `Investigation` | yes |

## GetInvitationsCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvitationsCount` | `integer` | no |

## GetMalwareProtectionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MalwareProtectionPlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Role` | `string` | no |
| `ProtectedResource` | `CreateProtectedResource` | no |
| `Actions` | `MalwareProtectionPlanActions` | no |
| `CreatedAt` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusReasons` | `List<MalwareProtectionPlanStatusReason>` | no |
| `Tags` | `Map<string>` | no |

## GetMalwareScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanId` | `string` | no |
| `DetectorId` | `string` | no |
| `AdminDetectorId` | `string` | no |
| `ResourceArn` | `string` | no |
| `ResourceType` | `string` | no |
| `ScannedResourcesCount` | `integer` | no |
| `SkippedResourcesCount` | `integer` | no |
| `FailedResourcesCount` | `integer` | no |
| `ScannedResources` | `List<ScannedResource>` | no |
| `ScanConfiguration` | `ScanConfiguration` | no |
| `ScanCategory` | `string` | no |
| `ScanStatus` | `string` | no |
| `ScanStatusReason` | `string` | no |
| `ScanType` | `string` | no |
| `ScanStartedAt` | `timestamp` | no |
| `ScanCompletedAt` | `timestamp` | no |
| `ScanResultDetails` | `GetMalwareScanResultDetails` | no |

## GetMalwareScanSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanResourceCriteria` | `ScanResourceCriteria` | no |
| `EbsSnapshotPreservation` | `string` | no |

## GetMasterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Master` | `Master` | yes |

## GetMemberDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MemberDataSourceConfigurations` | `List<MemberDataSourceConfiguration>` | yes |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## GetMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Members` | `List<Member>` | yes |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## GetOrganizationStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationDetails` | `OrganizationDetails` | no |

## GetRemainingFreeTrialDays

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accounts` | `List<AccountFreeTrialInfo>` | no |
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | no |

## GetThreatEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatEntitySetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `ExpectedBucketOwner` | `string` | no |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `ErrorDetails` | `string` | no |

## GetThreatIntelSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatIntelSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `ExpectedBucketOwner` | `string` | no |

## GetTrustedEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `TrustedEntitySetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Format` | `string` | yes |
| `Location` | `string` | yes |
| `ExpectedBucketOwner` | `string` | no |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `ErrorDetails` | `string` | no |

## GetUsageStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `UsageStatisticType` | `string` | yes |
| `UsageCriteria` | `UsageCriteria` | yes |
| `Unit` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageStatistics` | `UsageStatistics` | no |
| `NextToken` | `string` | no |

## InviteMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |
| `DisableEmailNotification` | `boolean` | no |
| `Message` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## ListCoverage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `FilterCriteria` | `CoverageFilterCriteria` | no |
| `SortCriteria` | `CoverageSortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<CoverageResource>` | yes |
| `NextToken` | `string` | no |

## ListCustomDetectionRuleAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RuleId` | `string` | no |
| `Mode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleAssociations` | `List<AssociationSummary>` | yes |
| `NextToken` | `string` | no |

## ListCustomDetectionRuleOrgConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configurations` | `List<DetectionRuleOrgConfigurationSummary>` | yes |
| `NextToken` | `string` | no |

## ListCustomDetectionRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<DetectionRuleFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<RuleSummary>` | yes |
| `NextToken` | `string` | no |

## ListDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterNames` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingCriteria` | `FindingCriteria` | no |
| `SortCriteria` | `SortCriteria` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FindingIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListIPSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpSetIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListInvestigations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `SortCriteria` | `InvestigationSortCriteria` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Investigations` | `List<InvestigationSummary>` | yes |
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

## ListMalwareProtectionPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MalwareProtectionPlans` | `List<MalwareProtectionPlanSummary>` | no |
| `NextToken` | `string` | no |

## ListMalwareScans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `FilterCriteria` | `ListMalwareScansFilterCriteria` | no |
| `SortCriteria` | `SortCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scans` | `List<MalwareScan>` | yes |
| `NextToken` | `string` | no |

## ListMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `OnlyAssociated` | `string` | no |

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

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdminAccounts` | `List<AdminAccount>` | no |
| `NextToken` | `string` | no |

## ListPublishingDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destinations` | `List<Destination>` | yes |
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

## ListThreatEntitySets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThreatEntitySetIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListThreatIntelSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThreatIntelSetIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListTrustedEntitySets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustedEntitySetIds` | `List<string>` | yes |
| `NextToken` | `string` | no |

## SendObjectMalwareScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3Object` | `S3ObjectForSendObjectMalwareScan` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMalwareScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `ScanConfiguration` | `StartMalwareScanConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScanId` | `string` | no |

## StartMonitoringMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## StopMonitoringMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnarchiveFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingIds` | `List<string>` | yes |

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


## UpdateCustomDetectionRuleAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `Mode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCustomDetectionRuleOrgConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `Mode` | `string` | yes |
| `IncludeAccountIds` | `List<string>` | no |
| `ExcludeAccountIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDetector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `Enable` | `boolean` | no |
| `FindingPublishingFrequency` | `string` | no |
| `DataSources` | `DataSourceConfigurations` | no |
| `Features` | `List<DetectorFeatureConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FilterName` | `string` | yes |
| `Description` | `string` | no |
| `Action` | `string` | no |
| `Rank` | `integer` | no |
| `FindingCriteria` | `FindingCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

## UpdateFindingsFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `FindingIds` | `List<string>` | yes |
| `Feedback` | `string` | yes |
| `Comments` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `IpSetId` | `string` | yes |
| `Name` | `string` | no |
| `Location` | `string` | no |
| `Activate` | `boolean` | no |
| `ExpectedBucketOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMalwareProtectionPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MalwareProtectionPlanId` | `string` | yes |
| `Role` | `string` | no |
| `Actions` | `MalwareProtectionPlanActions` | no |
| `ProtectedResource` | `UpdateProtectedResource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMalwareScanSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ScanResourceCriteria` | `ScanResourceCriteria` | no |
| `EbsSnapshotPreservation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMemberDetectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AccountIds` | `List<string>` | yes |
| `DataSources` | `DataSourceConfigurations` | no |
| `Features` | `List<MemberFeaturesConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedAccounts` | `List<UnprocessedAccount>` | yes |

## UpdateOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `AutoEnable` | `boolean` | no |
| `DataSources` | `OrganizationDataSourceConfigurations` | no |
| `Features` | `List<OrganizationFeatureConfiguration>` | no |
| `AutoEnableOrganizationMembers` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePublishingDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `DestinationId` | `string` | yes |
| `DestinationProperties` | `DestinationProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThreatEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatEntitySetId` | `string` | yes |
| `Name` | `string` | no |
| `Location` | `string` | no |
| `ExpectedBucketOwner` | `string` | no |
| `Activate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThreatIntelSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `ThreatIntelSetId` | `string` | yes |
| `Name` | `string` | no |
| `Location` | `string` | no |
| `Activate` | `boolean` | no |
| `ExpectedBucketOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrustedEntitySet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectorId` | `string` | yes |
| `TrustedEntitySetId` | `string` | yes |
| `Name` | `string` | no |
| `Location` | `string` | no |
| `ExpectedBucketOwner` | `string` | no |
| `Activate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


