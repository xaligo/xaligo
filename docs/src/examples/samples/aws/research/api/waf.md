# AWS WAF

API version: 2015-08-24. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/waf/2015-08-24/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateByteMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByteMatchSet` | `ByteMatchSet` | no |
| `ChangeToken` | `string` | no |

## CreateGeoMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoMatchSet` | `GeoMatchSet` | no |
| `ChangeToken` | `string` | no |

## CreateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSet` | `IPSet` | no |
| `ChangeToken` | `string` | no |

## CreateRateBasedRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MetricName` | `string` | yes |
| `RateKey` | `string` | yes |
| `RateLimit` | `long` | yes |
| `ChangeToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `RateBasedRule` | no |
| `ChangeToken` | `string` | no |

## CreateRegexMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexMatchSet` | `RegexMatchSet` | no |
| `ChangeToken` | `string` | no |

## CreateRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSet` | `RegexPatternSet` | no |
| `ChangeToken` | `string` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MetricName` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `Rule` | no |
| `ChangeToken` | `string` | no |

## CreateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MetricName` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroup` | `RuleGroup` | no |
| `ChangeToken` | `string` | no |

## CreateSizeConstraintSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SizeConstraintSet` | `SizeConstraintSet` | no |
| `ChangeToken` | `string` | no |

## CreateSqlInjectionMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SqlInjectionMatchSet` | `SqlInjectionMatchSet` | no |
| `ChangeToken` | `string` | no |

## CreateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MetricName` | `string` | yes |
| `DefaultAction` | `WafAction` | yes |
| `ChangeToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACL` | `WebACL` | no |
| `ChangeToken` | `string` | no |

## CreateWebACLMigrationStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLId` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `IgnoreUnsupportedType` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3ObjectUrl` | `string` | yes |

## CreateXssMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `XssMatchSet` | `XssMatchSet` | no |
| `ChangeToken` | `string` | no |

## DeleteByteMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByteMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteGeoMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

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


## DeleteRateBasedRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteRegexMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteSizeConstraintSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SizeConstraintSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteSqlInjectionMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SqlInjectionMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## DeleteXssMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `XssMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## GetByteMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByteMatchSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByteMatchSet` | `ByteMatchSet` | no |

## GetChangeToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## GetChangeTokenStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeTokenStatus` | `string` | no |

## GetGeoMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoMatchSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoMatchSet` | `GeoMatchSet` | no |

## GetIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSet` | `IPSet` | no |

## GetLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

## GetPermissionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetRateBasedRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `RateBasedRule` | no |

## GetRateBasedRuleManagedKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `NextMarker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedKeys` | `List<string>` | no |
| `NextMarker` | `string` | no |

## GetRegexMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexMatchSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexMatchSet` | `RegexMatchSet` | no |

## GetRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSet` | `RegexPatternSet` | no |

## GetRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `Rule` | no |

## GetRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroup` | `RuleGroup` | no |

## GetSampledRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebAclId` | `string` | yes |
| `RuleId` | `string` | yes |
| `TimeWindow` | `TimeWindow` | yes |
| `MaxItems` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SampledRequests` | `List<SampledHTTPRequest>` | no |
| `PopulationSize` | `long` | no |
| `TimeWindow` | `TimeWindow` | no |

## GetSizeConstraintSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SizeConstraintSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SizeConstraintSet` | `SizeConstraintSet` | no |

## GetSqlInjectionMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SqlInjectionMatchSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SqlInjectionMatchSet` | `SqlInjectionMatchSet` | no |

## GetWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACL` | `WebACL` | no |

## GetXssMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `XssMatchSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `XssMatchSet` | `XssMatchSet` | no |

## ListActivatedRulesInRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupId` | `string` | no |
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ActivatedRules` | `List<ActivatedRule>` | no |

## ListByteMatchSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ByteMatchSets` | `List<ByteMatchSetSummary>` | no |

## ListGeoMatchSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `GeoMatchSets` | `List<GeoMatchSetSummary>` | no |

## ListIPSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
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
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfigurations` | `List<LoggingConfiguration>` | no |
| `NextMarker` | `string` | no |

## ListRateBasedRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Rules` | `List<RuleSummary>` | no |

## ListRegexMatchSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RegexMatchSets` | `List<RegexMatchSetSummary>` | no |

## ListRegexPatternSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RegexPatternSets` | `List<RegexPatternSetSummary>` | no |

## ListRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RuleGroups` | `List<RuleGroupSummary>` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Rules` | `List<RuleSummary>` | no |

## ListSizeConstraintSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `SizeConstraintSets` | `List<SizeConstraintSetSummary>` | no |

## ListSqlInjectionMatchSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `SqlInjectionMatchSets` | `List<SqlInjectionMatchSetSummary>` | no |

## ListSubscribedRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `RuleGroups` | `List<SubscribedRuleGroupSummary>` | no |

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
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `WebACLs` | `List<WebACLSummary>` | no |

## ListXssMatchSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `XssMatchSets` | `List<XssMatchSetSummary>` | no |

## PutLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

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


## UpdateByteMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByteMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<ByteMatchSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateGeoMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<GeoMatchSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateIPSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IPSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<IPSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateRateBasedRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<RuleUpdate>` | yes |
| `RateLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateRegexMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexMatchSetId` | `string` | yes |
| `Updates` | `List<RegexMatchSetUpdate>` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateRegexPatternSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegexPatternSetId` | `string` | yes |
| `Updates` | `List<RegexPatternSetUpdate>` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<RuleUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleGroupId` | `string` | yes |
| `Updates` | `List<RuleGroupUpdate>` | yes |
| `ChangeToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateSizeConstraintSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SizeConstraintSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<SizeConstraintSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateSqlInjectionMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SqlInjectionMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<SqlInjectionMatchSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<WebACLUpdate>` | no |
| `DefaultAction` | `WafAction` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

## UpdateXssMatchSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `XssMatchSetId` | `string` | yes |
| `ChangeToken` | `string` | yes |
| `Updates` | `List<XssMatchSetUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeToken` | `string` | no |

