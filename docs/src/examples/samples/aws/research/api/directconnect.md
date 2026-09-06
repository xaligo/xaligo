# AWS Direct Connect

API version: 2012-10-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/directconnect/2012-10-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptDirectConnectGatewayAssociationProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | yes |
| `proposalId` | `string` | yes |
| `associatedGatewayOwnerAccount` | `string` | yes |
| `overrideAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociation` | `DirectConnectGatewayAssociation` | no |

## AllocateConnectionOnInterconnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bandwidth` | `string` | yes |
| `connectionName` | `string` | yes |
| `ownerAccount` | `string` | yes |
| `interconnectId` | `string` | yes |
| `vlan` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## AllocateHostedConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `ownerAccount` | `string` | yes |
| `bandwidth` | `string` | yes |
| `connectionName` | `string` | yes |
| `vlan` | `integer` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## AllocatePrivateVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `ownerAccount` | `string` | yes |
| `newPrivateVirtualInterfaceAllocation` | `NewPrivateVirtualInterfaceAllocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

## AllocatePublicVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `ownerAccount` | `string` | yes |
| `newPublicVirtualInterfaceAllocation` | `NewPublicVirtualInterfaceAllocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

## AllocateTransitVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `ownerAccount` | `string` | yes |
| `newTransitVirtualInterfaceAllocation` | `NewTransitVirtualInterfaceAllocation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterface` | `VirtualInterface` | no |

## AssociateConnectionWithLag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `lagId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## AssociateHostedConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `parentConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## AssociateMacSecKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `secretARN` | `string` | no |
| `ckn` | `string` | no |
| `cak` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |

## AssociateVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `connectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

## ConfirmConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionState` | `string` | no |

## ConfirmCustomerAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreementName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## ConfirmPrivateVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceState` | `string` | no |

## ConfirmPublicVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceState` | `string` | no |

## ConfirmTransitVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `directConnectGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceState` | `string` | no |

## CreateBGPPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | no |
| `newBGPPeer` | `NewBGPPeer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterface` | `VirtualInterface` | no |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `location` | `string` | yes |
| `bandwidth` | `string` | yes |
| `connectionName` | `string` | yes |
| `lagId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `requestMACSec` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## CreateDirectConnectGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayName` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `amazonSideAsn` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGateway` | `DirectConnectGateway` | no |

## CreateDirectConnectGatewayAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | yes |
| `gatewayId` | `string` | no |
| `addAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |
| `virtualGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociation` | `DirectConnectGatewayAssociation` | no |

## CreateDirectConnectGatewayAssociationProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | yes |
| `directConnectGatewayOwnerAccount` | `string` | yes |
| `gatewayId` | `string` | yes |
| `addAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |
| `removeAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociationProposal` | `DirectConnectGatewayAssociationProposal` | no |

## CreateInterconnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectName` | `string` | yes |
| `bandwidth` | `string` | yes |
| `location` | `string` | yes |
| `lagId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `requestMACSec` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectId` | `string` | no |
| `interconnectName` | `string` | no |
| `interconnectState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |

## CreateLag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `numberOfConnections` | `integer` | yes |
| `location` | `string` | yes |
| `connectionsBandwidth` | `string` | yes |
| `lagName` | `string` | yes |
| `connectionId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `childConnectionTags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `requestMACSec` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionsBandwidth` | `string` | no |
| `numberOfConnections` | `integer` | no |
| `lagId` | `string` | no |
| `ownerAccount` | `string` | no |
| `lagName` | `string` | no |
| `lagState` | `string` | no |
| `location` | `string` | no |
| `region` | `string` | no |
| `minimumLinks` | `integer` | no |
| `awsDevice` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `connections` | `List<Connection>` | no |
| `allowsHostedConnections` | `boolean` | no |
| `jumboFrameCapable` | `boolean` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |

## CreatePrivateVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `newPrivateVirtualInterface` | `NewPrivateVirtualInterface` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

## CreatePublicVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `newPublicVirtualInterface` | `NewPublicVirtualInterface` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

## CreateTransitVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `newTransitVirtualInterface` | `NewTransitVirtualInterface` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterface` | `VirtualInterface` | no |

## DeleteBGPPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `customerAddress` | `string` | no |
| `bgpPeerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterface` | `VirtualInterface` | no |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## DeleteDirectConnectGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGateway` | `DirectConnectGateway` | no |

## DeleteDirectConnectGatewayAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `virtualGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociation` | `DirectConnectGatewayAssociation` | no |

## DeleteDirectConnectGatewayAssociationProposal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `proposalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociationProposal` | `DirectConnectGatewayAssociationProposal` | no |

## DeleteInterconnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectState` | `string` | no |

## DeleteLag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lagId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionsBandwidth` | `string` | no |
| `numberOfConnections` | `integer` | no |
| `lagId` | `string` | no |
| `ownerAccount` | `string` | no |
| `lagName` | `string` | no |
| `lagState` | `string` | no |
| `location` | `string` | no |
| `region` | `string` | no |
| `minimumLinks` | `integer` | no |
| `awsDevice` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `connections` | `List<Connection>` | no |
| `allowsHostedConnections` | `boolean` | no |
| `jumboFrameCapable` | `boolean` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |

## DeleteVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceState` | `string` | no |

## DescribeConnectionLoa

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `providerName` | `string` | no |
| `loaContentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loa` | `Loa` | no |

## DescribeConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connections` | `List<Connection>` | no |
| `nextToken` | `string` | no |

## DescribeConnectionsOnInterconnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connections` | `List<Connection>` | no |
| `nextToken` | `string` | no |

## DescribeCustomerMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agreements` | `List<CustomerAgreement>` | no |
| `nniPartnerType` | `string` | no |

## DescribeDirectConnectGatewayAssociationProposals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | no |
| `proposalId` | `string` | no |
| `associatedGatewayId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociationProposals` | `List<DirectConnectGatewayAssociationProposal>` | no |
| `nextToken` | `string` | no |

## DescribeDirectConnectGatewayAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationId` | `string` | no |
| `associatedGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `virtualGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociations` | `List<DirectConnectGatewayAssociation>` | no |
| `nextToken` | `string` | no |

## DescribeDirectConnectGatewayAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAttachments` | `List<DirectConnectGatewayAttachment>` | no |
| `nextToken` | `string` | no |

## DescribeDirectConnectGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGateways` | `List<DirectConnectGateway>` | no |
| `nextToken` | `string` | no |

## DescribeHostedConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connections` | `List<Connection>` | no |
| `nextToken` | `string` | no |

## DescribeInterconnectLoa

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectId` | `string` | yes |
| `providerName` | `string` | no |
| `loaContentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loa` | `Loa` | no |

## DescribeInterconnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnectId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `interconnects` | `List<Interconnect>` | no |
| `nextToken` | `string` | no |

## DescribeLags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lagId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lags` | `List<Lag>` | no |
| `nextToken` | `string` | no |

## DescribeLoa

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `providerName` | `string` | no |
| `loaContentType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loaContent` | `blob` | no |
| `loaContentType` | `string` | no |

## DescribeLocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `locations` | `List<Location>` | no |

## DescribeRouterConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `routerTypeIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customerRouterConfig` | `string` | no |
| `router` | `RouterType` | no |
| `virtualInterfaceId` | `string` | no |
| `virtualInterfaceName` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceTags` | `List<ResourceTag>` | no |

## DescribeVirtualGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualGateways` | `List<VirtualGateway>` | no |

## DescribeVirtualInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaces` | `List<VirtualInterface>` | no |
| `nextToken` | `string` | no |

## DisassociateConnectionFromLag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `lagId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## DisassociateMacSecKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `secretARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |

## ListVirtualInterfaceRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | no |
| `filters` | `RouteFilters` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | no |
| `routes` | `List<Route>` | no |
| `nextToken` | `string` | no |

## ListVirtualInterfaceTestHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testId` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `bgpPeers` | `List<string>` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceTestHistory` | `List<VirtualInterfaceTestHistory>` | no |
| `nextToken` | `string` | no |

## StartBgpFailoverTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `bgpPeers` | `List<string>` | no |
| `testDurationInMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceTest` | `VirtualInterfaceTestHistory` | no |

## StopBgpFailoverTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceTest` | `VirtualInterfaceTestHistory` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `connectionName` | `string` | no |
| `encryptionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `connectionId` | `string` | no |
| `connectionName` | `string` | no |
| `connectionState` | `string` | no |
| `region` | `string` | no |
| `location` | `string` | no |
| `bandwidth` | `string` | no |
| `vlan` | `integer` | no |
| `partnerName` | `string` | no |
| `loaIssueTime` | `timestamp` | no |
| `lagId` | `string` | no |
| `awsDevice` | `string` | no |
| `jumboFrameCapable` | `boolean` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `portEncryptionStatus` | `string` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |
| `partnerInterconnectMacSecCapable` | `boolean` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |

## UpdateDirectConnectGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayId` | `string` | yes |
| `newDirectConnectGatewayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGateway` | `DirectConnectGateway` | no |

## UpdateDirectConnectGatewayAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationId` | `string` | no |
| `addAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |
| `removeAllowedPrefixesToDirectConnectGateway` | `List<RouteFilterPrefix>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `directConnectGatewayAssociation` | `DirectConnectGatewayAssociation` | no |

## UpdateLag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lagId` | `string` | yes |
| `lagName` | `string` | no |
| `minimumLinks` | `integer` | no |
| `encryptionMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionsBandwidth` | `string` | no |
| `numberOfConnections` | `integer` | no |
| `lagId` | `string` | no |
| `ownerAccount` | `string` | no |
| `lagName` | `string` | no |
| `lagState` | `string` | no |
| `location` | `string` | no |
| `region` | `string` | no |
| `minimumLinks` | `integer` | no |
| `awsDevice` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `connections` | `List<Connection>` | no |
| `allowsHostedConnections` | `boolean` | no |
| `jumboFrameCapable` | `boolean` | no |
| `hasLogicalRedundancy` | `string` | no |
| `tags` | `List<Tag>` | no |
| `providerName` | `string` | no |
| `macSecCapable` | `boolean` | no |
| `encryptionMode` | `string` | no |
| `macSecKeys` | `List<MacSecKey>` | no |
| `prefixPoolSizeIpv4` | `integer` | no |
| `prefixPoolSizeIpv6` | `integer` | no |
| `prefixPoolUnallocatedCountIpv4` | `integer` | no |
| `prefixPoolUnallocatedCountIpv6` | `integer` | no |
| `rateLimiterStatus` | `RateLimiterStatus` | no |

## UpdateVirtualInterfaceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualInterfaceId` | `string` | yes |
| `mtu` | `integer` | no |
| `enableSiteLink` | `boolean` | no |
| `virtualInterfaceName` | `string` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ownerAccount` | `string` | no |
| `virtualInterfaceId` | `string` | no |
| `location` | `string` | no |
| `connectionId` | `string` | no |
| `virtualInterfaceType` | `string` | no |
| `virtualInterfaceName` | `string` | no |
| `vlan` | `integer` | no |
| `asn` | `integer` | no |
| `asnLong` | `long` | no |
| `amazonSideAsn` | `long` | no |
| `authKey` | `string` | no |
| `amazonAddress` | `string` | no |
| `customerAddress` | `string` | no |
| `addressFamily` | `string` | no |
| `virtualInterfaceState` | `string` | no |
| `customerRouterConfig` | `string` | no |
| `mtu` | `integer` | no |
| `jumboFrameCapable` | `boolean` | no |
| `virtualGatewayId` | `string` | no |
| `directConnectGatewayId` | `string` | no |
| `routeFilterPrefixes` | `List<RouteFilterPrefix>` | no |
| `bgpPeers` | `List<BGPPeer>` | no |
| `region` | `string` | no |
| `awsDeviceV2` | `string` | no |
| `awsLogicalDeviceId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `siteLinkEnabled` | `boolean` | no |
| `prefixPoolAllocatedCountIpv4` | `integer` | no |
| `prefixPoolAllocatedCountIpv6` | `integer` | no |
| `rateLimit` | `string` | no |

