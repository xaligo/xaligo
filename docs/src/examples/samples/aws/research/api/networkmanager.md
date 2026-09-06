# AWS Network Manager

API version: 2019-07-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/networkmanager/2019-07-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attachment` | `Attachment` | no |

## AssociateConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectPeerId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `LinkId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeerAssociation` | `ConnectPeerAssociation` | no |

## AssociateCustomerGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayArn` | `string` | yes |
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `LinkId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayAssociation` | `CustomerGatewayAssociation` | no |

## AssociateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `LinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkAssociation` | `LinkAssociation` | no |

## AssociateTransitGatewayConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayConnectPeerArn` | `string` | yes |
| `DeviceId` | `string` | yes |
| `LinkId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeerAssociation` | `TransitGatewayConnectPeerAssociation` | no |

## CreateConnectAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `EdgeLocation` | `string` | yes |
| `TransportAttachmentId` | `string` | yes |
| `RoutingPolicyLabel` | `string` | no |
| `Options` | `ConnectAttachmentOptions` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectAttachment` | `ConnectAttachment` | no |

## CreateConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectAttachmentId` | `string` | yes |
| `CoreNetworkAddress` | `string` | no |
| `PeerAddress` | `string` | yes |
| `BgpOptions` | `BgpOptions` | no |
| `InsideCidrBlocks` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |
| `SubnetArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeer` | `ConnectPeer` | no |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `ConnectedDeviceId` | `string` | yes |
| `LinkId` | `string` | no |
| `ConnectedLinkId` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## CreateCoreNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `PolicyDocument` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetwork` | `CoreNetwork` | no |

## CreateCoreNetworkPrefixListAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PrefixListArn` | `string` | yes |
| `PrefixListAlias` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `PrefixListArn` | `string` | no |
| `PrefixListAlias` | `string` | no |

## CreateDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `AWSLocation` | `AWSLocation` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `Vendor` | `string` | no |
| `Model` | `string` | no |
| `SerialNumber` | `string` | no |
| `Location` | `Location` | no |
| `SiteId` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `Device` | no |

## CreateDirectConnectGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `DirectConnectGatewayArn` | `string` | yes |
| `RoutingPolicyLabel` | `string` | no |
| `EdgeLocations` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectConnectGatewayAttachment` | `DirectConnectGatewayAttachment` | no |

## CreateGlobalNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetwork` | `GlobalNetwork` | no |

## CreateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `Bandwidth` | `Bandwidth` | yes |
| `Provider` | `string` | no |
| `SiteId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Link` | `Link` | no |

## CreateSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `Description` | `string` | no |
| `Location` | `Location` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## CreateSiteToSiteVpnAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `VpnConnectionArn` | `string` | yes |
| `RoutingPolicyLabel` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteToSiteVpnAttachment` | `SiteToSiteVpnAttachment` | no |

## CreateTransitGatewayPeering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `TransitGatewayArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeering` | `TransitGatewayPeering` | no |

## CreateTransitGatewayRouteTableAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PeeringId` | `string` | yes |
| `TransitGatewayRouteTableArn` | `string` | yes |
| `RoutingPolicyLabel` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAttachment` | `TransitGatewayRouteTableAttachment` | no |

## CreateVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `VpcArn` | `string` | yes |
| `SubnetArns` | `List<string>` | yes |
| `Options` | `VpcOptions` | no |
| `RoutingPolicyLabel` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcAttachment` | `VpcAttachment` | no |

## DeleteAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attachment` | `Attachment` | no |

## DeleteConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeer` | `ConnectPeer` | no |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## DeleteCoreNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetwork` | `CoreNetwork` | no |

## DeleteCoreNetworkPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkPolicy` | `CoreNetworkPolicy` | no |

## DeleteCoreNetworkPrefixListAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PrefixListArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `PrefixListArn` | `string` | no |

## DeleteDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `Device` | no |

## DeleteGlobalNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetwork` | `GlobalNetwork` | no |

## DeleteLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `LinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Link` | `Link` | no |

## DeletePeering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PeeringId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Peering` | `Peering` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `SiteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## DeregisterTransitGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRegistration` | `TransitGatewayRegistration` | no |

## DescribeGlobalNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworks` | `List<GlobalNetwork>` | no |
| `NextToken` | `string` | no |

## DisassociateConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectPeerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeerAssociation` | `ConnectPeerAssociation` | no |

## DisassociateCustomerGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `CustomerGatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayAssociation` | `CustomerGatewayAssociation` | no |

## DisassociateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `LinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkAssociation` | `LinkAssociation` | no |

## DisassociateTransitGatewayConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayConnectPeerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeerAssociation` | `TransitGatewayConnectPeerAssociation` | no |

## ExecuteCoreNetworkChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConnectAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectAttachment` | `ConnectAttachment` | no |

## GetConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeer` | `ConnectPeer` | no |

## GetConnectPeerAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectPeerIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeerAssociations` | `List<ConnectPeerAssociation>` | no |
| `NextToken` | `string` | no |

## GetConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectionIds` | `List<string>` | no |
| `DeviceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<Connection>` | no |
| `NextToken` | `string` | no |

## GetCoreNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetwork` | `CoreNetwork` | no |

## GetCoreNetworkChangeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkChangeEvents` | `List<CoreNetworkChangeEvent>` | no |
| `NextToken` | `string` | no |

## GetCoreNetworkChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkChanges` | `List<CoreNetworkChange>` | no |
| `NextToken` | `string` | no |

## GetCoreNetworkPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | no |
| `Alias` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkPolicy` | `CoreNetworkPolicy` | no |

## GetCustomerGatewayAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `CustomerGatewayArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayAssociations` | `List<CustomerGatewayAssociation>` | no |
| `NextToken` | `string` | no |

## GetDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceIds` | `List<string>` | no |
| `SiteId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Devices` | `List<Device>` | no |
| `NextToken` | `string` | no |

## GetDirectConnectGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectConnectGatewayAttachment` | `DirectConnectGatewayAttachment` | no |

## GetLinkAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | no |
| `LinkId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkAssociations` | `List<LinkAssociation>` | no |
| `NextToken` | `string` | no |

## GetLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `LinkIds` | `List<string>` | no |
| `SiteId` | `string` | no |
| `Type` | `string` | no |
| `Provider` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Links` | `List<Link>` | no |
| `NextToken` | `string` | no |

## GetNetworkResourceCounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ResourceType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkResourceCounts` | `List<NetworkResourceCount>` | no |
| `NextToken` | `string` | no |

## GetNetworkResourceRelationships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `CoreNetworkId` | `string` | no |
| `RegisteredGatewayArn` | `string` | no |
| `AwsRegion` | `string` | no |
| `AccountId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Relationships` | `List<Relationship>` | no |
| `NextToken` | `string` | no |

## GetNetworkResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `CoreNetworkId` | `string` | no |
| `RegisteredGatewayArn` | `string` | no |
| `AwsRegion` | `string` | no |
| `AccountId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkResources` | `List<NetworkResource>` | no |
| `NextToken` | `string` | no |

## GetNetworkRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `RouteTableIdentifier` | `RouteTableIdentifier` | yes |
| `ExactCidrMatches` | `List<string>` | no |
| `LongestPrefixMatches` | `List<string>` | no |
| `SubnetOfMatches` | `List<string>` | no |
| `SupernetOfMatches` | `List<string>` | no |
| `PrefixListIds` | `List<string>` | no |
| `States` | `List<string>` | no |
| `Types` | `List<string>` | no |
| `DestinationFilters` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteTableArn` | `string` | no |
| `CoreNetworkSegmentEdge` | `CoreNetworkSegmentEdgeIdentifier` | no |
| `RouteTableType` | `string` | no |
| `RouteTableTimestamp` | `timestamp` | no |
| `NetworkRoutes` | `List<NetworkRoute>` | no |

## GetNetworkTelemetry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `CoreNetworkId` | `string` | no |
| `RegisteredGatewayArn` | `string` | no |
| `AwsRegion` | `string` | no |
| `AccountId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkTelemetry` | `List<NetworkTelemetry>` | no |
| `NextToken` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | no |

## GetRouteAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `RouteAnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteAnalysis` | `RouteAnalysis` | no |

## GetSiteToSiteVpnAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SiteToSiteVpnAttachment` | `SiteToSiteVpnAttachment` | no |

## GetSites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `SiteIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sites` | `List<Site>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayConnectPeerAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayConnectPeerArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeerAssociations` | `List<TransitGatewayConnectPeerAssociation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayPeering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PeeringId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeering` | `TransitGatewayPeering` | no |

## GetTransitGatewayRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRegistrations` | `List<TransitGatewayRegistration>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayRouteTableAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAttachment` | `TransitGatewayRouteTableAttachment` | no |

## GetVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcAttachment` | `VpcAttachment` | no |

## ListAttachmentRoutingPolicyAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `AttachmentId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentRoutingPolicyAssociations` | `List<AttachmentRoutingPolicyAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `AttachmentType` | `string` | no |
| `EdgeLocation` | `string` | no |
| `State` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attachments` | `List<Attachment>` | no |
| `NextToken` | `string` | no |

## ListConnectPeers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `ConnectAttachmentId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectPeers` | `List<ConnectPeerSummary>` | no |
| `NextToken` | `string` | no |

## ListCoreNetworkPolicyVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkPolicyVersions` | `List<CoreNetworkPolicyVersion>` | no |
| `NextToken` | `string` | no |

## ListCoreNetworkPrefixListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PrefixListArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixListAssociations` | `List<PrefixListAssociation>` | no |
| `NextToken` | `string` | no |

## ListCoreNetworkRoutingInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `SegmentName` | `string` | yes |
| `EdgeLocation` | `string` | yes |
| `NextHopFilters` | `Map<List<string>>` | no |
| `LocalPreferenceMatches` | `List<string>` | no |
| `ExactAsPathMatches` | `List<string>` | no |
| `MedMatches` | `List<string>` | no |
| `CommunityMatches` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkRoutingInformation` | `List<CoreNetworkRoutingInformation>` | no |
| `NextToken` | `string` | no |

## ListCoreNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworks` | `List<CoreNetworkSummary>` | no |
| `NextToken` | `string` | no |

## ListOrganizationServiceAccessStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationStatus` | `OrganizationStatus` | no |
| `NextToken` | `string` | no |

## ListPeerings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `PeeringType` | `string` | no |
| `EdgeLocation` | `string` | no |
| `State` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Peerings` | `List<Peering>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## PutAttachmentRoutingPolicyLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `AttachmentId` | `string` | yes |
| `RoutingPolicyLabel` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `AttachmentId` | `string` | no |
| `RoutingPolicyLabel` | `string` | no |

## PutCoreNetworkPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyDocument` | `string` | yes |
| `Description` | `string` | no |
| `LatestVersionId` | `integer` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkPolicy` | `CoreNetworkPolicy` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterTransitGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `TransitGatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRegistration` | `TransitGatewayRegistration` | no |

## RejectAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attachment` | `Attachment` | no |

## RemoveAttachmentRoutingPolicyLabel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `AttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | no |
| `AttachmentId` | `string` | no |
| `RoutingPolicyLabel` | `string` | no |

## RestoreCoreNetworkPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `PolicyVersionId` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkPolicy` | `CoreNetworkPolicy` | no |

## StartOrganizationServiceAccessUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationStatus` | `OrganizationStatus` | no |

## StartRouteAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `Source` | `RouteAnalysisEndpointOptionsSpecification` | yes |
| `Destination` | `RouteAnalysisEndpointOptionsSpecification` | yes |
| `IncludeReturnPath` | `boolean` | no |
| `UseMiddleboxes` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteAnalysis` | `RouteAnalysis` | no |

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


## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ConnectionId` | `string` | yes |
| `LinkId` | `string` | no |
| `ConnectedLinkId` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## UpdateCoreNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetworkId` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreNetwork` | `CoreNetwork` | no |

## UpdateDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `DeviceId` | `string` | yes |
| `AWSLocation` | `AWSLocation` | no |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `Vendor` | `string` | no |
| `Model` | `string` | no |
| `SerialNumber` | `string` | no |
| `Location` | `Location` | no |
| `SiteId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `Device` | no |

## UpdateDirectConnectGatewayAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |
| `EdgeLocations` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectConnectGatewayAttachment` | `DirectConnectGatewayAttachment` | no |

## UpdateGlobalNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetwork` | `GlobalNetwork` | no |

## UpdateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `LinkId` | `string` | yes |
| `Description` | `string` | no |
| `Type` | `string` | no |
| `Bandwidth` | `Bandwidth` | no |
| `Provider` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Link` | `Link` | no |

## UpdateNetworkResourceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `Metadata` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Metadata` | `Map<string>` | no |

## UpdateSite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GlobalNetworkId` | `string` | yes |
| `SiteId` | `string` | yes |
| `Description` | `string` | no |
| `Location` | `Location` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Site` | `Site` | no |

## UpdateVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | yes |
| `AddSubnetArns` | `List<string>` | no |
| `RemoveSubnetArns` | `List<string>` | no |
| `Options` | `VpcOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcAttachment` | `VpcAttachment` | no |

