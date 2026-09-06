# Amazon Route 53 Resolver

API version: 2018-04-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53resolver/2018-04-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateFirewallRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `FirewallRuleGroupId` | `string` | yes |
| `VpcId` | `string` | yes |
| `Priority` | `integer` | yes |
| `Name` | `string` | yes |
| `MutationProtection` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociation` | `FirewallRuleGroupAssociation` | no |

## AssociateResolverEndpointIpAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |
| `IpAddress` | `IpAddressUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## AssociateResolverQueryLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigAssociation` | `ResolverQueryLogConfigAssociation` | no |

## AssociateResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleId` | `string` | yes |
| `Name` | `string` | no |
| `VPCId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleAssociation` | `ResolverRuleAssociation` | no |

## BatchCreateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreateFirewallRuleEntries` | `List<CreateFirewallRuleEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedFirewallRules` | `List<FirewallRule>` | no |
| `CreateErrors` | `List<BatchCreateFirewallRuleError>` | no |

## BatchDeleteFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeleteFirewallRuleEntries` | `List<DeleteFirewallRuleEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletedFirewallRules` | `List<FirewallRule>` | no |
| `DeleteErrors` | `List<BatchDeleteFirewallRuleError>` | no |

## BatchUpdateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateFirewallRuleEntries` | `List<UpdateFirewallRuleEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdatedFirewallRules` | `List<FirewallRule>` | no |
| `UpdateErrors` | `List<BatchUpdateFirewallRuleError>` | no |

## CreateFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainList` | `FirewallDomainList` | no |

## CreateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `FirewallRuleGroupId` | `string` | yes |
| `FirewallDomainListId` | `string` | no |
| `Priority` | `integer` | yes |
| `Action` | `string` | yes |
| `BlockResponse` | `string` | no |
| `BlockOverrideDomain` | `string` | no |
| `BlockOverrideDnsType` | `string` | no |
| `BlockOverrideTtl` | `integer` | no |
| `Name` | `string` | yes |
| `FirewallDomainRedirectionAction` | `string` | no |
| `Qtype` | `string` | no |
| `DnsThreatProtection` | `string` | no |
| `ConfidenceThreshold` | `string` | no |
| `FirewallRuleType` | `FirewallRuleType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRule` | `FirewallRule` | no |

## CreateFirewallRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroup` | `FirewallRuleGroup` | no |

## CreateOutpostResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `Name` | `string` | yes |
| `InstanceCount` | `integer` | no |
| `PreferredInstanceType` | `string` | yes |
| `OutpostArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostResolver` | `OutpostResolver` | no |

## CreateResolverEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `Name` | `string` | no |
| `SecurityGroupIds` | `List<string>` | yes |
| `Direction` | `string` | yes |
| `IpAddresses` | `List<IpAddressRequest>` | yes |
| `OutpostArn` | `string` | no |
| `PreferredInstanceType` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ResolverEndpointType` | `string` | no |
| `Protocols` | `List<string>` | no |
| `RniEnhancedMetricsEnabled` | `boolean` | no |
| `TargetNameServerMetricsEnabled` | `boolean` | no |
| `Dns64Enabled` | `boolean` | no |
| `Ipv6InternetAccessEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## CreateResolverQueryLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DestinationArn` | `string` | yes |
| `CreatorRequestId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfig` | `ResolverQueryLogConfig` | no |

## CreateResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatorRequestId` | `string` | yes |
| `Name` | `string` | no |
| `RuleType` | `string` | yes |
| `DomainName` | `string` | no |
| `TargetIps` | `List<TargetAddress>` | no |
| `ResolverEndpointId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `DelegationRecord` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRule` | `ResolverRule` | no |

## DeleteFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainList` | `FirewallDomainList` | no |

## DeleteFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | yes |
| `FirewallDomainListId` | `string` | no |
| `FirewallThreatProtectionId` | `string` | no |
| `Qtype` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRule` | `FirewallRule` | no |

## DeleteFirewallRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroup` | `FirewallRuleGroup` | no |

## DeleteOutpostResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostResolver` | `OutpostResolver` | no |

## DeleteResolverEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## DeleteResolverQueryLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfig` | `ResolverQueryLogConfig` | no |

## DeleteResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRule` | `ResolverRule` | no |

## DisassociateFirewallRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociation` | `FirewallRuleGroupAssociation` | no |

## DisassociateResolverEndpointIpAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |
| `IpAddress` | `IpAddressUpdate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## DisassociateResolverQueryLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigAssociation` | `ResolverQueryLogConfigAssociation` | no |

## DisassociateResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VPCId` | `string` | yes |
| `ResolverRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleAssociation` | `ResolverRuleAssociation` | no |

## GetFirewallConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallConfig` | `FirewallConfig` | no |

## GetFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainList` | `FirewallDomainList` | no |

## GetFirewallRuleGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroup` | `FirewallRuleGroup` | no |

## GetFirewallRuleGroupAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociation` | `FirewallRuleGroupAssociation` | no |

## GetFirewallRuleGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupPolicy` | `string` | no |

## GetOutpostResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostResolver` | `OutpostResolver` | no |

## GetResolverConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverConfig` | `ResolverConfig` | no |

## GetResolverDnssecConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverDNSSECConfig` | `ResolverDnssecConfig` | no |

## GetResolverEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## GetResolverQueryLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfig` | `ResolverQueryLogConfig` | no |

## GetResolverQueryLogConfigAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigAssociation` | `ResolverQueryLogConfigAssociation` | no |

## GetResolverQueryLogConfigPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverQueryLogConfigPolicy` | `string` | no |

## GetResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRule` | `ResolverRule` | no |

## GetResolverRuleAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleAssociation` | `ResolverRuleAssociation` | no |

## GetResolverRulePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRulePolicy` | `string` | no |

## ImportFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainListId` | `string` | yes |
| `Operation` | `string` | yes |
| `DomainFileUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |

## ListFirewallConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallConfigs` | `List<FirewallConfig>` | no |

## ListFirewallDomainLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallDomainLists` | `List<FirewallDomainListMetadata>` | no |

## ListFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainListId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Domains` | `List<string>` | no |

## ListFirewallRuleGroupAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | no |
| `VpcId` | `string` | no |
| `Priority` | `integer` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallRuleGroupAssociations` | `List<FirewallRuleGroupAssociation>` | no |

## ListFirewallRuleGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallRuleGroups` | `List<FirewallRuleGroupMetadata>` | no |

## ListFirewallRuleTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleTypes` | `List<FirewallRuleTypeDefinition>` | no |
| `NextToken` | `string` | no |

## ListFirewallRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | yes |
| `Priority` | `integer` | no |
| `Action` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FirewallRules` | `List<FirewallRule>` | no |

## ListOutpostResolvers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostResolvers` | `List<OutpostResolver>` | no |
| `NextToken` | `string` | no |

## ListResolverConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ResolverConfigs` | `List<ResolverConfig>` | no |

## ListResolverDnssecConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ResolverDnssecConfigs` | `List<ResolverDnssecConfig>` | no |

## ListResolverEndpointIpAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IpAddresses` | `List<IpAddressResponse>` | no |

## ListResolverEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResolverEndpoints` | `List<ResolverEndpoint>` | no |

## ListResolverQueryLogConfigAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TotalCount` | `integer` | no |
| `TotalFilteredCount` | `integer` | no |
| `ResolverQueryLogConfigAssociations` | `List<ResolverQueryLogConfigAssociation>` | no |

## ListResolverQueryLogConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TotalCount` | `integer` | no |
| `TotalFilteredCount` | `integer` | no |
| `ResolverQueryLogConfigs` | `List<ResolverQueryLogConfig>` | no |

## ListResolverRuleAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResolverRuleAssociations` | `List<ResolverRuleAssociation>` | no |

## ListResolverRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResolverRules` | `List<ResolverRule>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## PutFirewallRuleGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `FirewallRuleGroupPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## PutResolverQueryLogConfigPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ResolverQueryLogConfigPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## PutResolverRulePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ResolverRulePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

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


## UpdateFirewallConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `FirewallFailOpen` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallConfig` | `FirewallConfig` | no |

## UpdateFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallDomainListId` | `string` | yes |
| `Operation` | `string` | yes |
| `Domains` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |

## UpdateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupId` | `string` | yes |
| `FirewallDomainListId` | `string` | no |
| `FirewallThreatProtectionId` | `string` | no |
| `Priority` | `integer` | no |
| `Action` | `string` | no |
| `BlockResponse` | `string` | no |
| `BlockOverrideDomain` | `string` | no |
| `BlockOverrideDnsType` | `string` | no |
| `BlockOverrideTtl` | `integer` | no |
| `Name` | `string` | no |
| `FirewallDomainRedirectionAction` | `string` | no |
| `Qtype` | `string` | no |
| `DnsThreatProtection` | `string` | no |
| `ConfidenceThreshold` | `string` | no |
| `FirewallRuleType` | `FirewallRuleType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRule` | `FirewallRule` | no |

## UpdateFirewallRuleGroupAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociationId` | `string` | yes |
| `Priority` | `integer` | no |
| `MutationProtection` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FirewallRuleGroupAssociation` | `FirewallRuleGroupAssociation` | no |

## UpdateOutpostResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `InstanceCount` | `integer` | no |
| `PreferredInstanceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostResolver` | `OutpostResolver` | no |

## UpdateResolverConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `AutodefinedReverseFlag` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverConfig` | `ResolverConfig` | no |

## UpdateResolverDnssecConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Validation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverDNSSECConfig` | `ResolverDnssecConfig` | no |

## UpdateResolverEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpointId` | `string` | yes |
| `Name` | `string` | no |
| `ResolverEndpointType` | `string` | no |
| `UpdateIpAddresses` | `List<UpdateIpAddress>` | no |
| `Protocols` | `List<string>` | no |
| `RniEnhancedMetricsEnabled` | `boolean` | no |
| `TargetNameServerMetricsEnabled` | `boolean` | no |
| `Dns64Enabled` | `boolean` | no |
| `Ipv6InternetAccessEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverEndpoint` | `ResolverEndpoint` | no |

## UpdateResolverRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRuleId` | `string` | yes |
| `Config` | `ResolverRuleConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResolverRule` | `ResolverRule` | no |

