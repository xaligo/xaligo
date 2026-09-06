# Amazon Route 53 Global Resolver

API version: 2022-09-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53globalresolver/2022-09-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hostedZoneId` | `string` | yes |
| `resourceArn` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `resourceArn` | `string` | yes |
| `hostedZoneId` | `string` | yes |
| `hostedZoneName` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## BatchCreateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallRules` | `List<BatchCreateFirewallRuleInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failures` | `List<BatchCreateFirewallRuleOutputItem>` | yes |
| `successes` | `List<BatchCreateFirewallRuleOutputItem>` | yes |

## BatchDeleteFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallRules` | `List<BatchDeleteFirewallRuleInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failures` | `List<BatchDeleteFirewallRuleOutputItem>` | yes |
| `successes` | `List<BatchDeleteFirewallRuleOutputItem>` | yes |

## BatchUpdateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallRules` | `List<BatchUpdateFirewallRuleInputItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failures` | `List<BatchUpdateFirewallRuleOutputItem>` | yes |
| `successes` | `List<BatchUpdateFirewallRuleOutputItem>` | yes |

## CreateAccessSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cidr` | `string` | yes |
| `clientToken` | `string` | no |
| `ipAddressType` | `string` | no |
| `name` | `string` | no |
| `dnsViewId` | `string` | yes |
| `protocol` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `cidr` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `id` | `string` | yes |
| `ipAddressType` | `string` | yes |
| `name` | `string` | no |
| `dnsViewId` | `string` | yes |
| `protocol` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `dnsViewId` | `string` | yes |
| `expiresAt` | `timestamp` | no |
| `name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `dnsViewId` | `string` | yes |
| `expiresAt` | `timestamp` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `value` | `string` | yes |

## CreateDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `globalResolverId` | `string` | yes |
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `dnssecValidation` | `string` | no |
| `ednsClientSubnet` | `string` | no |
| `firewallRulesFailOpen` | `string` | no |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `globalResolverId` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `domainCount` | `integer` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `clientToken` | `string` | no |
| `confidenceThreshold` | `string` | no |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallDomainListId` | `string` | no |
| `name` | `string` | yes |
| `priority` | `long` | no |
| `dnsViewId` | `string` | yes |
| `qType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `confidenceThreshold` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallDomainListId` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `priority` | `long` | yes |
| `dnsViewId` | `string` | yes |
| `queryType` | `string` | no |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateGlobalResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `ipAddressType` | `string` | no |
| `name` | `string` | yes |
| `observabilityRegion` | `string` | no |
| `regions` | `List<string>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `dnsName` | `string` | yes |
| `ipAddressType` | `string` | no |
| `ipv4Addresses` | `List<string>` | yes |
| `ipv6Addresses` | `List<string>` | no |
| `name` | `string` | yes |
| `observabilityRegion` | `string` | no |
| `regions` | `List<string>` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteAccessSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `cidr` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `id` | `string` | yes |
| `ipAddressType` | `string` | yes |
| `name` | `string` | no |
| `dnsViewId` | `string` | yes |
| `protocol` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessTokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | yes |
| `deletedAt` | `timestamp` | yes |

## DeleteDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dnsViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## DeleteFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |

## DeleteFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `confidenceThreshold` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallDomainListId` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `priority` | `long` | yes |
| `dnsViewId` | `string` | yes |
| `queryType` | `string` | no |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteGlobalResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `globalResolverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | yes |
| `dnsName` | `string` | yes |
| `observabilityRegion` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `regions` | `List<string>` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `ipv4Addresses` | `List<string>` | yes |
| `ipv6Addresses` | `List<string>` | no |
| `ipAddressType` | `string` | no |

## DisableDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dnsViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## DisassociateHostedZone

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hostedZoneId` | `string` | yes |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `resourceArn` | `string` | yes |
| `hostedZoneId` | `string` | yes |
| `hostedZoneName` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## EnableDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dnsViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## GetAccessSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `cidr` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `id` | `string` | yes |
| `ipAddressType` | `string` | yes |
| `name` | `string` | no |
| `dnsViewId` | `string` | yes |
| `protocol` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## GetAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessTokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `dnsViewId` | `string` | yes |
| `expiresAt` | `timestamp` | yes |
| `globalResolverId` | `string` | yes |
| `name` | `string` | no |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `value` | `string` | yes |

## GetDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dnsViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## GetFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `globalResolverId` | `string` | yes |
| `clientToken` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `domainCount` | `integer` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `statusMessage` | `string` | no |
| `updatedAt` | `timestamp` | yes |

## GetFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `firewallRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `confidenceThreshold` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallDomainListId` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `priority` | `long` | yes |
| `dnsViewId` | `string` | yes |
| `queryType` | `string` | no |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## GetGlobalResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `globalResolverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | yes |
| `dnsName` | `string` | yes |
| `observabilityRegion` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `regions` | `List<string>` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `ipv4Addresses` | `List<string>` | yes |
| `ipv6Addresses` | `List<string>` | no |
| `ipAddressType` | `string` | no |

## GetHostedZoneAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hostedZoneAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `resourceArn` | `string` | yes |
| `hostedZoneId` | `string` | yes |
| `hostedZoneName` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## GetManagedFirewallDomainList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedFirewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `managedListType` | `string` | yes |

## ImportFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainFileUrl` | `string` | yes |
| `firewallDomainListId` | `string` | yes |
| `operation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |

## ListAccessSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `accessSources` | `List<AccessSourcesItem>` | yes |

## ListAccessTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `dnsViewId` | `string` | yes |
| `filters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `accessTokens` | `List<AccessTokenItem>` | no |

## ListDNSViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `globalResolverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dnsViews` | `List<DNSViewSummary>` | yes |

## ListFirewallDomainLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `globalResolverId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `firewallDomainLists` | `List<FirewallDomainListsItem>` | yes |

## ListFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `firewallDomainListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `domains` | `List<string>` | yes |

## ListFirewallRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `dnsViewId` | `string` | yes |
| `filters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `firewallRules` | `List<FirewallRulesItem>` | yes |

## ListGlobalResolvers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `globalResolvers` | `List<GlobalResolversItem>` | yes |

## ListHostedZoneAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `hostedZoneAssociations` | `List<HostedZoneAssociationSummary>` | yes |

## ListManagedFirewallDomainLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `managedFirewallDomainListType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `managedFirewallDomainLists` | `List<ManagedFirewallDomainListsItem>` | yes |

## ListSharedDNSViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `dnsViews` | `List<SharedDNSViewSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccessSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessSourceId` | `string` | yes |
| `cidr` | `string` | no |
| `ipAddressType` | `string` | no |
| `name` | `string` | no |
| `protocol` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `cidr` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `id` | `string` | yes |
| `ipAddressType` | `string` | yes |
| `name` | `string` | no |
| `dnsViewId` | `string` | yes |
| `protocol` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessTokenId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |

## UpdateDNSView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dnsViewId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `dnssecValidation` | `string` | no |
| `ednsClientSubnet` | `string` | no |
| `firewallRulesFailOpen` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | no |
| `dnssecValidation` | `string` | yes |
| `ednsClientSubnet` | `string` | yes |
| `firewallRulesFailOpen` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `globalResolverId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## UpdateFirewallDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domains` | `List<string>` | yes |
| `firewallDomainListId` | `string` | yes |
| `operation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |

## UpdateFirewallRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | no |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `clientToken` | `string` | yes |
| `confidenceThreshold` | `string` | no |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallRuleId` | `string` | yes |
| `name` | `string` | no |
| `priority` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `blockOverrideDnsType` | `string` | no |
| `blockOverrideDomain` | `string` | no |
| `blockOverrideTtl` | `integer` | no |
| `blockResponse` | `string` | no |
| `confidenceThreshold` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `dnsAdvancedProtection` | `string` | no |
| `firewallDomainListId` | `string` | no |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `priority` | `long` | yes |
| `dnsViewId` | `string` | yes |
| `queryType` | `string` | no |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateGlobalResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `globalResolverId` | `string` | yes |
| `name` | `string` | no |
| `observabilityRegion` | `string` | no |
| `description` | `string` | no |
| `ipAddressType` | `string` | no |
| `regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `clientToken` | `string` | yes |
| `dnsName` | `string` | yes |
| `observabilityRegion` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `regions` | `List<string>` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `ipv4Addresses` | `List<string>` | yes |
| `ipv6Addresses` | `List<string>` | no |
| `ipAddressType` | `string` | no |

## UpdateHostedZoneAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `hostedZoneAssociationId` | `string` | yes |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `resourceArn` | `string` | yes |
| `hostedZoneId` | `string` | yes |
| `hostedZoneName` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

