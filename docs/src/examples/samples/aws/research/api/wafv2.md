# AWS WAFV2

API version: 2019-07-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/wafv2/2019-07-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLArn` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `Rules` | `List<Rule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Capacity` | `long` | no |

## CreateAPIKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `TokenDomains` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `APIKey` | `string` | no |

## CreateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Description` | `string` | no |
| `IPAddressVersion` | `string` | yes |
| `Addresses` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `IPSetSummary` | no |

## CreateRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Description` | `string` | no |
| `RegularExpressionList` | `List<Regex>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `RegexPatternSetSummary` | no |

## CreateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Capacity` | `long` | yes |
| `Description` | `string` | no |
| `Rules` | `List<Rule>` | no |
| `VisibilityConfig` | `VisibilityConfig` | yes |
| `Tags` | `List<Tag>` | no |
| `CustomResponseBodies` | `Map<CustomResponseBody>` | no |
| `MonetizationConfig` | `MonetizationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `RuleGroupSummary` | no |

## CreateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `DefaultAction` | `DefaultAction` | yes |
| `Description` | `string` | no |
| `Rules` | `List<Rule>` | no |
| `VisibilityConfig` | `VisibilityConfig` | yes |
| `DataProtectionConfig` | `DataProtectionConfig` | no |
| `Tags` | `List<Tag>` | no |
| `CustomResponseBodies` | `Map<CustomResponseBody>` | no |
| `CaptchaConfig` | `CaptchaConfig` | no |
| `ChallengeConfig` | `ChallengeConfig` | no |
| `TokenDomains` | `List<string>` | no |
| `AssociationConfig` | `AssociationConfig` | no |
| `OnSourceDDoSProtectionConfig` | `OnSourceDDoSProtectionConfig` | no |
| `ApplicationConfig` | `ApplicationConfig` | no |
| `MonetizationConfig` | `MonetizationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summary` | `WebACLSummary` | no |

## DeleteAPIKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `APIKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFirewallManagerRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLArn` | `string` | yes |
| `WebACLLockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextWebACLLockToken` | `string` | no |

## DeleteIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `LogType` | `string` | no |
| `LogScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePermissionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAllManagedProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedProducts` | `List<ManagedProductDescriptor>` | no |

## DescribeManagedProductsByVendor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VendorName` | `string` | yes |
| `Scope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedProducts` | `List<ManagedProductDescriptor>` | no |

## DescribeManagedRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VendorName` | `string` | yes |
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `VersionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VersionName` | `string` | no |
| `SnsTopicArn` | `string` | no |
| `Capacity` | `long` | no |
| `Rules` | `List<RuleSummary>` | no |
| `LabelNamespace` | `string` | no |
| `AvailableLabels` | `List<LabelSummary>` | no |
| `ConsumedLabels` | `List<LabelSummary>` | no |

## DisassociateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GenerateMobileSdkReleaseUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Platform` | `string` | yes |
| `ReleaseVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |

## GetDecryptedAPIKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `APIKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TokenDomains` | `List<string>` | no |
| `CreationTimestamp` | `timestamp` | no |

## GetIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSet` | `IPSet` | no |
| `LockToken` | `string` | no |

## GetLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `LogType` | `string` | no |
| `LogScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

## GetManagedRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedRuleSet` | `ManagedRuleSet` | no |
| `LockToken` | `string` | no |

## GetMobileSdkRelease

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Platform` | `string` | yes |
| `ReleaseVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MobileSdkRelease` | `MobileSdkRelease` | no |

## GetPermissionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetRateBasedStatementManagedKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `WebACLName` | `string` | yes |
| `WebACLId` | `string` | yes |
| `RuleGroupRuleName` | `string` | no |
| `RuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedKeysIPV4` | `RateBasedStatementManagedKeysIPSet` | no |
| `ManagedKeysIPV6` | `RateBasedStatementManagedKeysIPSet` | no |

## GetRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSet` | `RegexPatternSet` | no |
| `LockToken` | `string` | no |

## GetRevenueStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticType` | `string` | yes |
| `TimeWindow` | `TimeWindow` | yes |
| `Scope` | `string` | yes |
| `Currency` | `string` | yes |
| `GroupBy` | `string` | no |
| `Filters` | `List<MonetizationFilter>` | no |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceStatistics` | `List<SourceStatistics>` | no |
| `RevenuePathStatistics` | `List<RevenuePathStatistics>` | no |
| `NextMarker` | `string` | no |

## GetRevenueStatisticsSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeWindow` | `TimeWindow` | yes |
| `Scope` | `string` | yes |
| `Currency` | `string` | yes |
| `Filters` | `List<MonetizationFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevenueBreakdown` | `RevenueBreakdown` | no |

## GetRevenueStatisticsTimeSeries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatisticType` | `string` | yes |
| `TimeWindow` | `TimeWindow` | yes |
| `Scope` | `string` | yes |
| `Interval` | `string` | yes |
| `Currency` | `string` | yes |
| `GroupBy` | `string` | no |
| `Filters` | `List<MonetizationFilter>` | no |
| `Limit` | `integer` | no |
| `NextMarker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataPoints` | `List<DataPointEntry>` | no |
| `NextMarker` | `string` | no |

## GetRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Scope` | `string` | no |
| `Id` | `string` | no |
| `ARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroup` | `RuleGroup` | no |
| `LockToken` | `string` | no |

## GetSampledRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAclArn` | `string` | yes |
| `RuleMetricName` | `string` | yes |
| `Scope` | `string` | yes |
| `TimeWindow` | `TimeWindow` | yes |
| `MaxItems` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SampledRequests` | `List<SampledHTTPRequest>` | no |
| `PopulationSize` | `long` | no |
| `TimeWindow` | `TimeWindow` | no |

## GetTopPathStatisticsByTraffic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAclArn` | `string` | yes |
| `Scope` | `string` | yes |
| `UriPathPrefix` | `string` | no |
| `TimeWindow` | `TimeWindow` | yes |
| `BotCategory` | `string` | no |
| `BotOrganization` | `string` | no |
| `BotName` | `string` | no |
| `Limit` | `integer` | yes |
| `NumberOfTopTrafficBotsPerPath` | `integer` | yes |
| `NextMarker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathStatistics` | `List<PathStatistics>` | yes |
| `TotalRequestCount` | `long` | yes |
| `NextMarker` | `string` | no |
| `TopCategories` | `List<PathStatistics>` | no |

## GetWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Scope` | `string` | no |
| `Id` | `string` | no |
| `ARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACL` | `WebACL` | no |
| `LockToken` | `string` | no |
| `ApplicationIntegrationURL` | `string` | no |

## GetWebACLForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACL` | `WebACL` | no |

## ListAPIKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `APIKeySummaries` | `List<APIKeySummary>` | no |
| `ApplicationIntegrationURL` | `string` | no |

## ListAvailableManagedRuleGroupVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VendorName` | `string` | yes |
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Versions` | `List<ManagedRuleGroupVersion>` | no |
| `CurrentDefaultVersion` | `string` | no |

## ListAvailableManagedRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ManagedRuleGroups` | `List<ManagedRuleGroupSummary>` | no |

## ListIPSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `IPSets` | `List<IPSetSummary>` | no |

## ListLoggingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |
| `LogScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfigurations` | `List<LoggingConfiguration>` | no |
| `NextMarker` | `string` | no |

## ListManagedRuleSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ManagedRuleSets` | `List<ManagedRuleSetSummary>` | no |

## ListMobileSdkReleases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Platform` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReleaseSummaries` | `List<ReleaseSummary>` | no |
| `NextMarker` | `string` | no |

## ListRegexPatternSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RegexPatternSets` | `List<RegexPatternSetSummary>` | no |

## ListResourcesForWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLArn` | `string` | yes |
| `ResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | no |

## ListRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RuleGroups` | `List<RuleGroupSummary>` | no |

## ListSettlementRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimeWindow` | `TimeWindow` | yes |
| `Scope` | `string` | yes |
| `Currency` | `string` | yes |
| `Filters` | `List<MonetizationFilter>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `Limit` | `integer` | no |
| `NextMarker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Settlements` | `List<SettlementRecord>` | no |
| `NextMarker` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `TagInfoForResource` | `TagInfoForResource` | no |

## ListWebACLs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Scope` | `string` | yes |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `WebACLs` | `List<WebACLSummary>` | no |

## PutLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

## PutManagedRuleSetVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |
| `RecommendedVersion` | `string` | no |
| `VersionsToPublish` | `Map<VersionToPublish>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextLockToken` | `string` | no |

## PutPermissionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

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


## UpdateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `Addresses` | `List<string>` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextLockToken` | `string` | no |

## UpdateManagedRuleSetVersionExpiryDate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `LockToken` | `string` | yes |
| `VersionToExpire` | `string` | yes |
| `ExpiryTimestamp` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExpiringVersion` | `string` | no |
| `ExpiryTimestamp` | `timestamp` | no |
| `NextLockToken` | `string` | no |

## UpdateRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `RegularExpressionList` | `List<Regex>` | yes |
| `LockToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextLockToken` | `string` | no |

## UpdateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<Rule>` | no |
| `VisibilityConfig` | `VisibilityConfig` | yes |
| `LockToken` | `string` | yes |
| `CustomResponseBodies` | `Map<CustomResponseBody>` | no |
| `MonetizationConfig` | `MonetizationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextLockToken` | `string` | no |

## UpdateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Scope` | `string` | yes |
| `Id` | `string` | yes |
| `DefaultAction` | `DefaultAction` | yes |
| `Description` | `string` | no |
| `Rules` | `List<Rule>` | no |
| `VisibilityConfig` | `VisibilityConfig` | yes |
| `DataProtectionConfig` | `DataProtectionConfig` | no |
| `LockToken` | `string` | yes |
| `CustomResponseBodies` | `Map<CustomResponseBody>` | no |
| `CaptchaConfig` | `CaptchaConfig` | no |
| `ChallengeConfig` | `ChallengeConfig` | no |
| `TokenDomains` | `List<string>` | no |
| `AssociationConfig` | `AssociationConfig` | no |
| `OnSourceDDoSProtectionConfig` | `OnSourceDDoSProtectionConfig` | no |
| `ApplicationConfig` | `ApplicationConfig` | no |
| `MonetizationConfig` | `MonetizationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextLockToken` | `string` | no |

