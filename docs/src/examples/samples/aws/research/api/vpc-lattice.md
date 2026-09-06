# Amazon VPC Lattice

API version: 2022-11-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/vpc-lattice/2022-11-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchUpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `rules` | `List<RuleUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successful` | `List<RuleUpdateSuccess>` | no |
| `unsuccessful` | `List<RuleUpdateFailure>` | no |

## CreateAccessLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `resourceIdentifier` | `string` | yes |
| `destinationArn` | `string` | yes |
| `serviceNetworkLogType` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `resourceId` | `string` | yes |
| `resourceArn` | `string` | yes |
| `serviceNetworkLogType` | `string` | no |
| `destinationArn` | `string` | yes |

## CreateListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `protocol` | `string` | yes |
| `port` | `integer` | no |
| `defaultAction` | `RuleAction` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `protocol` | `string` | no |
| `port` | `integer` | no |
| `serviceArn` | `string` | no |
| `serviceId` | `string` | no |
| `defaultAction` | `RuleAction` | no |

## CreateResourceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `portRanges` | `List<string>` | no |
| `protocol` | `string` | no |
| `resourceGatewayIdentifier` | `string` | no |
| `resourceConfigurationGroupIdentifier` | `string` | no |
| `resourceConfigurationDefinition` | `ResourceConfigurationDefinition` | no |
| `allowAssociationToShareableServiceNetwork` | `boolean` | no |
| `customDomainName` | `string` | no |
| `groupDomain` | `string` | no |
| `domainVerificationIdentifier` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `resourceGatewayId` | `string` | no |
| `resourceConfigurationGroupId` | `string` | no |
| `type` | `string` | no |
| `portRanges` | `List<string>` | no |
| `protocol` | `string` | no |
| `status` | `string` | no |
| `resourceConfigurationDefinition` | `ResourceConfigurationDefinition` | no |
| `allowAssociationToShareableServiceNetwork` | `boolean` | no |
| `createdAt` | `timestamp` | no |
| `failureReason` | `string` | no |
| `customDomainName` | `string` | no |
| `domainVerificationId` | `string` | no |
| `groupDomain` | `string` | no |
| `domainVerificationArn` | `string` | no |

## CreateResourceGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `vpcIdentifier` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `ipAddressType` | `string` | no |
| `ipv4AddressesPerEni` | `integer` | no |
| `resourceConfigDnsResolution` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `vpcIdentifier` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `ipAddressType` | `string` | no |
| `ipv4AddressesPerEni` | `integer` | no |
| `resourceConfigDnsResolution` | `string` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `match` | `RuleMatch` | yes |
| `priority` | `integer` | yes |
| `action` | `RuleAction` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `match` | `RuleMatch` | no |
| `priority` | `integer` | no |
| `action` | `RuleAction` | no |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |
| `customDomainName` | `string` | no |
| `certificateArn` | `string` | no |
| `authType` | `string` | no |
| `idleTimeoutSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `customDomainName` | `string` | no |
| `certificateArn` | `string` | no |
| `status` | `string` | no |
| `authType` | `string` | no |
| `idleTimeoutSeconds` | `integer` | no |
| `dnsEntry` | `DnsEntry` | no |

## CreateServiceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `authType` | `string` | no |
| `tags` | `Map<string>` | no |
| `sharingConfig` | `SharingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `sharingConfig` | `SharingConfig` | no |
| `authType` | `string` | no |

## CreateServiceNetworkResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `resourceConfigurationIdentifier` | `string` | yes |
| `serviceNetworkIdentifier` | `string` | yes |
| `privateDnsEnabled` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `createdBy` | `string` | no |
| `privateDnsEnabled` | `boolean` | no |

## CreateServiceNetworkServiceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `serviceIdentifier` | `string` | yes |
| `serviceNetworkIdentifier` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |
| `createdBy` | `string` | no |
| `customDomainName` | `string` | no |
| `dnsEntry` | `DnsEntry` | no |

## CreateServiceNetworkVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `serviceNetworkIdentifier` | `string` | yes |
| `vpcIdentifier` | `string` | yes |
| `privateDnsEnabled` | `boolean` | no |
| `securityGroupIds` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `dnsOptions` | `DnsOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |
| `createdBy` | `string` | no |
| `securityGroupIds` | `List<string>` | no |
| `privateDnsEnabled` | `boolean` | no |
| `dnsOptions` | `DnsOptions` | no |

## CreateTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | yes |
| `config` | `TargetGroupConfig` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `config` | `TargetGroupConfig` | no |
| `status` | `string` | no |

## DeleteAccessLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessLogSubscriptionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAuthPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainVerificationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceConfigurationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceEndpointAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceEndpointAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `resourceConfigurationId` | `string` | no |
| `resourceConfigurationArn` | `string` | no |
| `vpcEndpointId` | `string` | no |

## DeleteResourceGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGatewayIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `ruleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |

## DeleteServiceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceNetworkResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkResourceAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |

## DeleteServiceNetworkServiceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkServiceAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |

## DeleteServiceNetworkVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkVpcAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |

## DeleteTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |

## DeregisterTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |
| `targets` | `List<Target>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successful` | `List<Target>` | no |
| `unsuccessful` | `List<TargetFailure>` | no |

## GetAccessLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessLogSubscriptionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `resourceId` | `string` | yes |
| `resourceArn` | `string` | yes |
| `destinationArn` | `string` | yes |
| `serviceNetworkLogType` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## GetAuthPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |
| `state` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetDomainVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainVerificationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `domainName` | `string` | yes |
| `status` | `string` | yes |
| `txtMethodConfig` | `TxtMethodConfig` | no |
| `createdAt` | `timestamp` | yes |
| `lastVerifiedTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `protocol` | `string` | no |
| `port` | `integer` | no |
| `serviceArn` | `string` | no |
| `serviceId` | `string` | no |
| `defaultAction` | `RuleAction` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetResourceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceConfigurationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `resourceGatewayId` | `string` | no |
| `resourceConfigurationGroupId` | `string` | no |
| `type` | `string` | no |
| `allowAssociationToShareableServiceNetwork` | `boolean` | no |
| `portRanges` | `List<string>` | no |
| `protocol` | `string` | no |
| `customDomainName` | `string` | no |
| `status` | `string` | no |
| `resourceConfigurationDefinition` | `ResourceConfigurationDefinition` | no |
| `createdAt` | `timestamp` | no |
| `amazonManaged` | `boolean` | no |
| `failureReason` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `domainVerificationId` | `string` | no |
| `domainVerificationArn` | `string` | no |
| `domainVerificationStatus` | `string` | no |
| `groupDomain` | `string` | no |

## GetResourceGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGatewayIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `vpcId` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `serviceManaged` | `boolean` | no |
| `managedBy` | `string` | no |
| `securityGroupIds` | `List<string>` | no |
| `ipAddressType` | `string` | no |
| `ipv4AddressesPerEni` | `integer` | no |
| `resourceConfigDnsResolution` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |

## GetRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `ruleIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `isDefault` | `boolean` | no |
| `match` | `RuleMatch` | no |
| `priority` | `integer` | no |
| `action` | `RuleAction` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `dnsEntry` | `DnsEntry` | no |
| `customDomainName` | `string` | no |
| `certificateArn` | `string` | no |
| `status` | `string` | no |
| `authType` | `string` | no |
| `idleTimeoutSeconds` | `integer` | no |
| `failureCode` | `string` | no |
| `failureMessage` | `string` | no |

## GetServiceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `arn` | `string` | no |
| `authType` | `string` | no |
| `sharingConfig` | `SharingConfig` | no |
| `numberOfAssociatedVPCs` | `long` | no |
| `numberOfAssociatedServices` | `long` | no |

## GetServiceNetworkResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkResourceAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `createdBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `resourceConfigurationId` | `string` | no |
| `resourceConfigurationArn` | `string` | no |
| `resourceConfigurationName` | `string` | no |
| `serviceNetworkId` | `string` | no |
| `serviceNetworkArn` | `string` | no |
| `serviceNetworkName` | `string` | no |
| `failureReason` | `string` | no |
| `failureCode` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `privateDnsEntry` | `DnsEntry` | no |
| `privateDnsEnabled` | `boolean` | no |
| `dnsEntry` | `DnsEntry` | no |
| `isManagedAssociation` | `boolean` | no |
| `domainVerificationStatus` | `string` | no |

## GetServiceNetworkServiceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkServiceAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |
| `createdBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `serviceId` | `string` | no |
| `serviceName` | `string` | no |
| `serviceArn` | `string` | no |
| `serviceNetworkId` | `string` | no |
| `serviceNetworkName` | `string` | no |
| `serviceNetworkArn` | `string` | no |
| `dnsEntry` | `DnsEntry` | no |
| `customDomainName` | `string` | no |
| `failureMessage` | `string` | no |
| `failureCode` | `string` | no |

## GetServiceNetworkVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkVpcAssociationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `status` | `string` | no |
| `arn` | `string` | no |
| `createdBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `serviceNetworkId` | `string` | no |
| `serviceNetworkName` | `string` | no |
| `serviceNetworkArn` | `string` | no |
| `vpcId` | `string` | no |
| `securityGroupIds` | `List<string>` | no |
| `privateDnsEnabled` | `boolean` | no |
| `failureMessage` | `string` | no |
| `failureCode` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `dnsOptions` | `DnsOptions` | no |

## GetTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `config` | `TargetGroupConfig` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `status` | `string` | no |
| `serviceArns` | `List<string>` | no |
| `failureMessage` | `string` | no |
| `failureCode` | `string` | no |

## ListAccessLogSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AccessLogSubscriptionSummary>` | yes |
| `nextToken` | `string` | no |

## ListDomainVerifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DomainVerificationSummary>` | yes |
| `nextToken` | `string` | no |

## ListListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ListenerSummary>` | yes |
| `nextToken` | `string` | no |

## ListResourceConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGatewayIdentifier` | `string` | no |
| `resourceConfigurationGroupIdentifier` | `string` | no |
| `domainVerificationIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ResourceConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListResourceEndpointAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceConfigurationIdentifier` | `string` | yes |
| `resourceEndpointAssociationIdentifier` | `string` | no |
| `vpcEndpointId` | `string` | no |
| `vpcEndpointOwner` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ResourceEndpointAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListResourceGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ResourceGatewaySummary>` | no |
| `nextToken` | `string` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RuleSummary>` | yes |
| `nextToken` | `string` | no |

## ListServiceNetworkResourceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | no |
| `resourceConfigurationIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `includeChildren` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceNetworkResourceAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListServiceNetworkServiceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | no |
| `serviceIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceNetworkServiceAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListServiceNetworkVpcAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | no |
| `vpcIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceNetworkVpcAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListServiceNetworkVpcEndpointAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceNetworkEndpointAssociation>` | yes |
| `nextToken` | `string` | no |

## ListServiceNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceNetworkSummary>` | yes |
| `nextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ServiceSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `vpcIdentifier` | `string` | no |
| `targetGroupType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TargetGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `targets` | `List<Target>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TargetSummary>` | yes |
| `nextToken` | `string` | no |

## PutAuthPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifier` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |
| `state` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |
| `targets` | `List<Target>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successful` | `List<Target>` | no |
| `unsuccessful` | `List<TargetFailure>` | no |

## StartDomainVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `domainName` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `domainName` | `string` | yes |
| `status` | `string` | yes |
| `txtMethodConfig` | `TxtMethodConfig` | no |

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


## UpdateAccessLogSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessLogSubscriptionIdentifier` | `string` | yes |
| `destinationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `resourceId` | `string` | yes |
| `resourceArn` | `string` | yes |
| `destinationArn` | `string` | yes |

## UpdateListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `defaultAction` | `RuleAction` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `protocol` | `string` | no |
| `port` | `integer` | no |
| `serviceArn` | `string` | no |
| `serviceId` | `string` | no |
| `defaultAction` | `RuleAction` | no |

## UpdateResourceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceConfigurationIdentifier` | `string` | yes |
| `resourceConfigurationDefinition` | `ResourceConfigurationDefinition` | no |
| `allowAssociationToShareableServiceNetwork` | `boolean` | no |
| `portRanges` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `resourceGatewayId` | `string` | no |
| `resourceConfigurationGroupId` | `string` | no |
| `type` | `string` | no |
| `portRanges` | `List<string>` | no |
| `allowAssociationToShareableServiceNetwork` | `boolean` | no |
| `protocol` | `string` | no |
| `status` | `string` | no |
| `resourceConfigurationDefinition` | `ResourceConfigurationDefinition` | no |

## UpdateResourceGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceGatewayIdentifier` | `string` | yes |
| `securityGroupIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `vpcId` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `ipAddressType` | `string` | no |

## UpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `listenerIdentifier` | `string` | yes |
| `ruleIdentifier` | `string` | yes |
| `match` | `RuleMatch` | no |
| `priority` | `integer` | no |
| `action` | `RuleAction` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `id` | `string` | no |
| `name` | `string` | no |
| `isDefault` | `boolean` | no |
| `match` | `RuleMatch` | no |
| `priority` | `integer` | no |
| `action` | `RuleAction` | no |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceIdentifier` | `string` | yes |
| `certificateArn` | `string` | no |
| `authType` | `string` | no |
| `idleTimeoutSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `customDomainName` | `string` | no |
| `certificateArn` | `string` | no |
| `authType` | `string` | no |
| `idleTimeoutSeconds` | `integer` | no |

## UpdateServiceNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkIdentifier` | `string` | yes |
| `authType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `authType` | `string` | no |

## UpdateServiceNetworkVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceNetworkVpcAssociationIdentifier` | `string` | yes |
| `securityGroupIds` | `List<string>` | no |
| `privateDnsEnabled` | `boolean` | no |
| `dnsOptions` | `DnsOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `status` | `string` | no |
| `createdBy` | `string` | no |
| `securityGroupIds` | `List<string>` | no |
| `privateDnsEnabled` | `boolean` | no |
| `dnsOptions` | `DnsOptions` | no |

## UpdateTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetGroupIdentifier` | `string` | yes |
| `healthCheck` | `HealthCheckConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `config` | `TargetGroupConfig` | no |
| `status` | `string` | no |

