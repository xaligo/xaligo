# Amazon Elastic Compute Cloud

API version: 2016-11-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ec2/2016-11-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAddressTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressTransfer` | `AddressTransfer` | no |

## AcceptCapacityReservationBillingOwnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityReservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## AcceptReservedInstancesExchangeQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ReservedInstanceIds` | `List<string>` | yes |
| `TargetConfigurations` | `List<TargetConfigurationRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExchangeId` | `string` | no |

## AcceptTransitGatewayClientVpnAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayClientVpnAttachment` | `TransitGatewayClientVpnAttachment` | no |

## AcceptTransitGatewayMulticastDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | no |
| `TransitGatewayAttachmentId` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `TransitGatewayMulticastDomainAssociations` | no |

## AcceptTransitGatewayPeeringAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeeringAttachment` | `TransitGatewayPeeringAttachment` | no |

## AcceptTransitGatewayVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachment` | `TransitGatewayVpcAttachment` | no |

## AcceptVpcEndpointConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `VpcEndpointIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## AcceptVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringConnection` | `VpcPeeringConnection` | no |

## AdvertiseByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `Asn` | `string` | no |
| `DryRun` | `boolean` | no |
| `NetworkBorderGroup` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## AllocateAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | no |
| `Address` | `string` | no |
| `PublicIpv4Pool` | `string` | no |
| `NetworkBorderGroup` | `string` | no |
| `CustomerOwnedIpv4Pool` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `IpamPoolId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | no |
| `PublicIpv4Pool` | `string` | no |
| `NetworkBorderGroup` | `string` | no |
| `Domain` | `string` | no |
| `CustomerOwnedIp` | `string` | no |
| `CustomerOwnedIpv4Pool` | `string` | no |
| `CarrierIp` | `string` | no |
| `PublicIp` | `string` | no |

## AllocateHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceFamily` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `HostRecovery` | `string` | no |
| `OutpostArn` | `string` | no |
| `HostMaintenance` | `string` | no |
| `AssetIds` | `List<string>` | no |
| `AvailabilityZoneId` | `string` | no |
| `CpuOptions` | `HostCpuOptionsRequest` | no |
| `AutoPlacement` | `string` | no |
| `ClientToken` | `string` | no |
| `InstanceType` | `string` | no |
| `Quantity` | `integer` | no |
| `AvailabilityZone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostIds` | `List<string>` | no |

## AllocateIpamPoolCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Cidr` | `string` | no |
| `NetmaskLength` | `integer` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `PreviewNextCidr` | `boolean` | no |
| `AllowedCidrs` | `List<string>` | no |
| `DisallowedCidrs` | `List<string>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolAllocation` | `IpamPoolAllocation` | no |

## ApplySecurityGroupsToClientVpnTargetNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `VpcId` | `string` | yes |
| `SecurityGroupIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupIds` | `List<string>` | no |

## AssignIpv6Addresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6PrefixCount` | `integer` | no |
| `Ipv6Prefixes` | `List<string>` | no |
| `NetworkInterfaceId` | `string` | yes |
| `Ipv6Addresses` | `List<string>` | no |
| `Ipv6AddressCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignedIpv6Addresses` | `List<string>` | no |
| `AssignedIpv6Prefixes` | `List<string>` | no |
| `NetworkInterfaceId` | `string` | no |

## AssignPrivateIpAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv4Prefixes` | `List<string>` | no |
| `Ipv4PrefixCount` | `integer` | no |
| `NetworkInterfaceId` | `string` | yes |
| `PrivateIpAddresses` | `List<string>` | no |
| `SecondaryPrivateIpAddressCount` | `integer` | no |
| `AllowReassignment` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | no |
| `AssignedPrivateIpAddresses` | `List<AssignedPrivateIpAddress>` | no |
| `AssignedIpv4Prefixes` | `List<Ipv4PrefixSpecification>` | no |

## AssignPrivateNatGatewayAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `PrivateIpAddresses` | `List<string>` | no |
| `PrivateIpAddressCount` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |
| `NatGatewayAddresses` | `List<NatGatewayAddress>` | no |

## AssociateAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | no |
| `InstanceId` | `string` | no |
| `PublicIp` | `string` | no |
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | no |
| `PrivateIpAddress` | `string` | no |
| `AllowReassociation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |

## AssociateApplicationStatusCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckId` | `string` | yes |
| `TargetTagAssociations` | `List<CustomTagKeyValueRequestPair>` | no |
| `InstanceIds` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulResults` | `List<SuccessfulAssociationResponseObject>` | no |
| `UnsuccessfulResults` | `List<UnsuccessfulAssociationResponseObject>` | no |

## AssociateCapacityReservationBillingOwner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityReservationId` | `string` | yes |
| `UnusedReservationBillingOwnerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## AssociateClientVpnTargetNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `SubnetId` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |
| `Status` | `AssociationStatus` | no |

## AssociateDhcpOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DhcpOptionsId` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateEnclaveCertificateIamRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateS3BucketName` | `string` | no |
| `CertificateS3ObjectKey` | `string` | no |
| `EncryptionKmsKeyId` | `string` | no |

## AssociateIamInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfile` | `IamInstanceProfileSpecification` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfileAssociation` | `IamInstanceProfileAssociation` | no |

## AssociateInstanceEventWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceEventWindowId` | `string` | yes |
| `AssociationTarget` | `InstanceEventWindowAssociationRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindow` | `InstanceEventWindow` | no |

## AssociateIpamByoasn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Asn` | `string` | yes |
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AsnAssociation` | `AsnAssociation` | no |

## AssociateIpamResourceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `IpamResourceDiscoveryId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscoveryAssociation` | `IpamResourceDiscoveryAssociation` | no |

## AssociateNatGatewayAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `AllocationIds` | `List<string>` | yes |
| `PrivateIpAddresses` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |
| `NatGatewayAddresses` | `List<NatGatewayAddress>` | no |

## AssociateRouteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerAssociation` | `RouteServerAssociation` | no |

## AssociateRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayId` | `string` | no |
| `PublicIpv4Pool` | `string` | no |
| `DryRun` | `boolean` | no |
| `SubnetId` | `string` | no |
| `RouteTableId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |
| `AssociationState` | `RouteTableAssociationState` | no |

## AssociateSecurityGroupVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |

## AssociateSubnetCidrBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6IpamPoolId` | `string` | no |
| `Ipv6NetmaskLength` | `integer` | no |
| `SubnetId` | `string` | yes |
| `Ipv6CidrBlock` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6CidrBlockAssociation` | `SubnetIpv6CidrBlockAssociation` | no |
| `SubnetId` | `string` | no |

## AssociateTransitGatewayMulticastDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `TransitGatewayMulticastDomainAssociations` | no |

## AssociateTransitGatewayPolicyTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `TransitGatewayPolicyTableAssociation` | no |

## AssociateTransitGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `TransitGatewayAssociation` | no |

## AssociateTrunkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BranchInterfaceId` | `string` | yes |
| `TrunkInterfaceId` | `string` | yes |
| `VlanId` | `integer` | no |
| `GreKey` | `integer` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InterfaceAssociation` | `TrunkInterfaceAssociation` | no |
| `ClientToken` | `string` | no |

## AssociateVpcCidrBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CidrBlock` | `string` | no |
| `Ipv6CidrBlockNetworkBorderGroup` | `string` | no |
| `Ipv6Pool` | `string` | no |
| `Ipv6CidrBlock` | `string` | no |
| `Ipv4IpamPoolId` | `string` | no |
| `Ipv4NetmaskLength` | `integer` | no |
| `Ipv6IpamPoolId` | `string` | no |
| `Ipv6NetmaskLength` | `integer` | no |
| `VpcId` | `string` | yes |
| `AmazonProvidedIpv6CidrBlock` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6CidrBlockAssociation` | `VpcIpv6CidrBlockAssociation` | no |
| `CidrBlockAssociation` | `VpcCidrBlockAssociation` | no |
| `VpcId` | `string` | no |

## AttachClassicLinkVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `VpcId` | `string` | yes |
| `Groups` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## AttachImageWatermark

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `WatermarkName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WatermarkKey` | `string` | no |

## AttachInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InternetGatewayId` | `string` | yes |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkCardIndex` | `integer` | no |
| `EnaSrdSpecification` | `EnaSrdSpecification` | no |
| `EnaQueueCount` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `DeviceIndex` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentId` | `string` | no |
| `NetworkCardIndex` | `integer` | no |

## AttachVerifiedAccessTrustProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `VerifiedAccessTrustProviderId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProvider` | `VerifiedAccessTrustProvider` | no |
| `VerifiedAccessInstance` | `VerifiedAccessInstance` | no |

## AttachVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `string` | yes |
| `InstanceId` | `string` | yes |
| `VolumeId` | `string` | yes |
| `EbsCardIndex` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeleteOnTermination` | `boolean` | no |
| `AssociatedResource` | `string` | no |
| `InstanceOwningService` | `string` | no |
| `EbsCardIndex` | `integer` | no |
| `VolumeId` | `string` | no |
| `InstanceId` | `string` | no |
| `Device` | `string` | no |
| `State` | `string` | no |
| `AttachTime` | `timestamp` | no |

## AttachVpnGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `VpnGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcAttachment` | `VpcAttachment` | no |

## AuthorizeClientVpnIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `TargetNetworkCidr` | `string` | yes |
| `AccessGroupId` | `string` | no |
| `AuthorizeAllGroups` | `boolean` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ClientVpnAuthorizationRuleStatus` | no |

## AuthorizeSecurityGroupEgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `GroupId` | `string` | yes |
| `SourceSecurityGroupName` | `string` | no |
| `SourceSecurityGroupOwnerId` | `string` | no |
| `IpProtocol` | `string` | no |
| `FromPort` | `integer` | no |
| `ToPort` | `integer` | no |
| `CidrIp` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `SecurityGroupRules` | `List<SecurityGroupRule>` | no |

## AuthorizeSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CidrIp` | `string` | no |
| `FromPort` | `integer` | no |
| `GroupId` | `string` | no |
| `GroupName` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |
| `IpProtocol` | `string` | no |
| `SourceSecurityGroupName` | `string` | no |
| `SourceSecurityGroupOwnerId` | `string` | no |
| `ToPort` | `integer` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `SecurityGroupRules` | `List<SecurityGroupRule>` | no |

## BatchModifyIpamRoutingPolicyRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `DeltaJson` | `string` | yes |
| `Force` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrationDelta` | `IpamRoutingPolicyRegistrationDelta` | no |

## BundleInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Storage` | `Storage` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleTask` | `BundleTask` | no |

## CancelBundleTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleTask` | `BundleTask` | no |

## CancelCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `ApplyCancellationCharges` | `string` | no |
| `QuoteId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## CancelCapacityReservationFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityReservationFleetIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulFleetCancellations` | `List<CapacityReservationFleetCancellationState>` | no |
| `FailedFleetCancellations` | `List<FailedCapacityReservationFleetCancellationResult>` | no |

## CancelConversionTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ConversionTaskId` | `string` | yes |
| `ReasonMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelDeclarativePoliciesReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## CancelExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelImageLaunchPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## CancelImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CancelReason` | `string` | no |
| `DryRun` | `boolean` | no |
| `ImportTaskId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportTaskId` | `string` | no |
| `PreviousState` | `string` | no |
| `State` | `string` | no |

## CancelReservedInstancesListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesListingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesListings` | `List<ReservedInstancesListing>` | no |

## CancelSpotFleetRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotFleetRequestIds` | `List<string>` | yes |
| `TerminateInstances` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulFleetRequests` | `List<CancelSpotFleetRequestsSuccessItem>` | no |
| `UnsuccessfulFleetRequests` | `List<CancelSpotFleetRequestsErrorItem>` | no |

## CancelSpotInstanceRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotInstanceRequestIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CancelledSpotInstanceRequests` | `List<CancelledSpotInstanceRequest>` | no |

## ConfirmProductInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ProductCode` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `OwnerId` | `string` | no |

## CopyFpgaImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SourceFpgaImageId` | `string` | yes |
| `Description` | `string` | no |
| `Name` | `string` | no |
| `SourceRegion` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FpgaImageId` | `string` | no |

## CopyImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `Name` | `string` | yes |
| `SourceImageId` | `string` | yes |
| `SourceRegion` | `string` | yes |
| `DestinationOutpostArn` | `string` | no |
| `CopyImageTags` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `SnapshotCopyCompletionDurationMinutes` | `long` | no |
| `DestinationAvailabilityZone` | `string` | no |
| `DestinationAvailabilityZoneId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## CopySnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DestinationOutpostArn` | `string` | no |
| `DestinationRegion` | `string` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `PresignedUrl` | `string` | no |
| `SourceRegion` | `string` | yes |
| `SourceSnapshotId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `CompletionDurationMinutes` | `integer` | no |
| `DestinationAvailabilityZone` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `SnapshotId` | `string` | no |

## CopyVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceVolumeId` | `string` | yes |
| `Iops` | `integer` | no |
| `Size` | `integer` | no |
| `VolumeType` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MultiAttachEnabled` | `boolean` | no |
| `Throughput` | `integer` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volumes` | `List<Volume>` | no |

## CreateApplicationStatusCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheckPaths` | `List<HealthCheckPathRequestObject>` | no |
| `Aggregation` | `string` | no |
| `Protocol` | `string` | yes |
| `Port` | `integer` | yes |
| `Path` | `string` | no |
| `DeviceIndex` | `integer` | no |
| `IpVersion` | `string` | no |
| `IpScope` | `string` | no |
| `Interval` | `integer` | no |
| `Timeout` | `integer` | no |
| `FailureThreshold` | `integer` | no |
| `SuccessThreshold` | `integer` | no |
| `StatusCodeMatcher` | `string` | no |
| `InitializationGracePeriodSeconds` | `integer` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheck` | `ApplicationStatusCheckResponseObject` | no |

## CreateCapacityManagerDataExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3BucketName` | `string` | yes |
| `S3BucketPrefix` | `string` | no |
| `Schedule` | `string` | yes |
| `OutputFormat` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerDataExportId` | `string` | no |

## CreateCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceType` | `string` | yes |
| `InstancePlatform` | `string` | yes |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `Tenancy` | `string` | no |
| `InstanceCount` | `integer` | yes |
| `EbsOptimized` | `boolean` | no |
| `EphemeralStorage` | `boolean` | no |
| `EndDate` | `timestamp` | no |
| `EndDateType` | `string` | no |
| `InstanceMatchCriteria` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `OutpostArn` | `string` | no |
| `PlacementGroupArn` | `string` | no |
| `StartDate` | `timestamp` | no |
| `CommitmentDuration` | `long` | no |
| `DeliveryPreference` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservation` | `CapacityReservation` | no |

## CreateCapacityReservationBySplitting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `SourceCapacityReservationId` | `string` | yes |
| `InstanceCount` | `integer` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceCapacityReservation` | `CapacityReservation` | no |
| `DestinationCapacityReservation` | `CapacityReservation` | no |
| `InstanceCount` | `integer` | no |

## CreateCapacityReservationCancellationQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationCancellationQuote` | `CapacityReservationCancellationQuote` | no |

## CreateCapacityReservationFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationStrategy` | `string` | no |
| `ClientToken` | `string` | no |
| `InstanceTypeSpecifications` | `List<ReservationFleetInstanceSpecification>` | no |
| `Tenancy` | `string` | no |
| `TotalTargetCapacity` | `integer` | yes |
| `EndDate` | `timestamp` | no |
| `InstanceMatchCriteria` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationFleetId` | `string` | no |
| `State` | `string` | no |
| `TotalTargetCapacity` | `integer` | no |
| `TotalFulfilledCapacity` | `double` | no |
| `InstanceMatchCriteria` | `string` | no |
| `AllocationStrategy` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `EndDate` | `timestamp` | no |
| `Tenancy` | `string` | no |
| `FleetCapacityReservations` | `List<FleetCapacityReservation>` | no |
| `Tags` | `List<Tag>` | no |

## CreateCarrierGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CarrierGateway` | `CarrierGateway` | no |

## CreateClientVpnEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientCidrBlock` | `string` | no |
| `ServerCertificateArn` | `string` | yes |
| `AuthenticationOptions` | `List<ClientVpnAuthenticationRequest>` | yes |
| `ConnectionLogOptions` | `ConnectionLogOptions` | yes |
| `DnsServers` | `List<string>` | no |
| `TransportProtocol` | `string` | no |
| `VpnPort` | `integer` | no |
| `Description` | `string` | no |
| `SplitTunnel` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `VpcId` | `string` | no |
| `SelfServicePortal` | `string` | no |
| `ClientConnectOptions` | `ClientConnectOptions` | no |
| `SessionTimeoutHours` | `integer` | no |
| `ClientLoginBannerOptions` | `ClientLoginBannerOptions` | no |
| `ClientRouteEnforcementOptions` | `ClientRouteEnforcementOptions` | no |
| `DisconnectOnSessionTimeout` | `boolean` | no |
| `EndpointIpAddressType` | `string` | no |
| `TrafficIpAddressType` | `string` | no |
| `TransitGatewayConfiguration` | `TransitGatewayConfigurationInputStructure` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | no |
| `Status` | `ClientVpnEndpointStatus` | no |
| `DnsName` | `string` | no |

## CreateClientVpnRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `DestinationCidrBlock` | `string` | yes |
| `TargetVpcSubnetId` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ClientVpnRouteStatus` | no |

## CreateCoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `CoipPoolId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipCidr` | `CoipCidr` | no |

## CreateCoipPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipPool` | `CoipPool` | no |

## CreateCustomerGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BgpAsn` | `integer` | no |
| `PublicIp` | `string` | no |
| `CertificateArn` | `string` | no |
| `Type` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DeviceName` | `string` | no |
| `IpAddress` | `string` | no |
| `BgpAsnExtended` | `long` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGateway` | `CustomerGateway` | no |

## CreateDefaultSubnet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `DryRun` | `boolean` | no |
| `Ipv6Native` | `boolean` | no |
| `AvailabilityZoneId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subnet` | `Subnet` | no |

## CreateDefaultVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Vpc` | `Vpc` | no |

## CreateDelegateMacVolumeOwnershipTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `MacCredentials` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MacModificationTask` | `MacModificationTask` | no |

## CreateDhcpOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DhcpConfigurations` | `List<NewDhcpConfiguration>` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DhcpOptions` | `DhcpOptions` | no |

## CreateEgressOnlyInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `EgressOnlyInternetGateway` | `EgressOnlyInternetGateway` | no |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `SpotOptions` | `SpotOptionsRequest` | no |
| `OnDemandOptions` | `OnDemandOptionsRequest` | no |
| `ReservedCapacityOptions` | `ReservedCapacityOptionsRequest` | no |
| `ExcessCapacityTerminationPolicy` | `string` | no |
| `LaunchTemplateConfigs` | `List<FleetLaunchTemplateConfigRequest>` | yes |
| `TargetCapacitySpecification` | `TargetCapacitySpecificationRequest` | yes |
| `TerminateInstancesWithExpiration` | `boolean` | no |
| `Type` | `string` | no |
| `ValidFrom` | `timestamp` | no |
| `ValidUntil` | `timestamp` | no |
| `ReplaceUnhealthyInstances` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `Context` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetId` | `string` | no |
| `Errors` | `List<CreateFleetError>` | no |
| `Instances` | `List<CreateFleetInstance>` | no |

## CreateFlowLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `DeliverLogsPermissionArn` | `string` | no |
| `DeliverCrossAccountRole` | `string` | no |
| `LogGroupName` | `string` | no |
| `ResourceIds` | `List<string>` | yes |
| `ResourceType` | `string` | yes |
| `TrafficType` | `string` | no |
| `LogDestinationType` | `string` | no |
| `LogDestination` | `string` | no |
| `LogFormat` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MaxAggregationInterval` | `integer` | no |
| `DestinationOptions` | `DestinationOptionsRequest` | no |
| `TagFieldSpecifications` | `List<TagFieldSpecificationRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `FlowLogIds` | `List<string>` | no |
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## CreateFpgaImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InputStorageLocation` | `StorageLocation` | yes |
| `LogsStorageLocation` | `StorageLocation` | no |
| `Description` | `string` | no |
| `Name` | `string` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FpgaImageId` | `string` | no |
| `FpgaImageGlobalId` | `string` | no |

## CreateImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `SnapshotLocation` | `string` | no |
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `NoReboot` | `boolean` | no |
| `BlockDeviceMappings` | `List<BlockDeviceMapping>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## CreateImageUsageReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `ResourceTypes` | `List<ImageUsageResourceTypeRequest>` | yes |
| `AccountIds` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | no |

## CreateInstanceConnectEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SubnetId` | `string` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `PreserveClientIp` | `boolean` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceConnectEndpoint` | `Ec2InstanceConnectEndpoint` | no |
| `ClientToken` | `string` | no |

## CreateInstanceEventWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Name` | `string` | no |
| `TimeRanges` | `List<InstanceEventWindowTimeRangeRequest>` | no |
| `CronExpression` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindow` | `InstanceEventWindow` | no |

## CreateInstanceExportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `Description` | `string` | no |
| `InstanceId` | `string` | yes |
| `TargetEnvironment` | `string` | yes |
| `ExportToS3Task` | `ExportToS3TaskSpecification` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTask` | `ExportTask` | no |

## CreateInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InternetGateway` | `InternetGateway` | no |

## CreateInterruptibleCapacityReservationAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `InstanceCount` | `integer` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ZeroSizePreference` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceCapacityReservationId` | `string` | no |
| `TargetInstanceCount` | `integer` | no |
| `Status` | `string` | no |
| `InterruptionType` | `string` | no |

## CreateIpam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Description` | `string` | no |
| `OperatingRegions` | `List<AddIpamOperatingRegion>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `Tier` | `string` | no |
| `EnablePrivateGua` | `boolean` | no |
| `MeteredAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipam` | `Ipam` | no |

## CreateIpamExternalResourceVerificationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamExternalResourceVerificationToken` | `IpamExternalResourceVerificationToken` | no |

## CreateIpamInternetRegistryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Rir` | `string` | yes |
| `OrganizationHandle` | `string` | yes |
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamInternetRegistryAssociation` | `IpamInternetRegistryAssociation` | no |

## CreateIpamPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `IpamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicy` | `IpamPolicy` | no |

## CreateIpamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamScopeId` | `string` | yes |
| `Locale` | `string` | no |
| `SourceIpamPoolId` | `string` | no |
| `Description` | `string` | no |
| `AddressFamily` | `string` | yes |
| `AutoImport` | `boolean` | no |
| `PubliclyAdvertisable` | `boolean` | no |
| `AllocationMinNetmaskLength` | `integer` | no |
| `AllocationMaxNetmaskLength` | `integer` | no |
| `AllocationDefaultNetmaskLength` | `integer` | no |
| `AllocationResourceTags` | `List<RequestIpamResourceTag>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `AwsService` | `string` | no |
| `PublicIpSource` | `string` | no |
| `SourceResource` | `IpamPoolSourceResourceRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPool` | `IpamPool` | no |

## CreateIpamPrefixListResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Description` | `string` | no |
| `AddressFamily` | `string` | yes |
| `Rules` | `List<IpamPrefixListResolverRuleRequest>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolver` | `IpamPrefixListResolver` | no |

## CreateIpamPrefixListResolverTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |
| `PrefixListId` | `string` | yes |
| `PrefixListRegion` | `string` | yes |
| `DesiredVersion` | `long` | no |
| `TrackLatestVersion` | `boolean` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolverTarget` | `IpamPrefixListResolverTarget` | no |

## CreateIpamResourceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Description` | `string` | no |
| `OperatingRegions` | `List<AddIpamOperatingRegion>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscovery` | `IpamResourceDiscovery` | no |

## CreateIpamRoutingPolicyRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `Cidr` | `string` | yes |
| `Asns` | `List<string>` | yes |
| `PermitMoreSpecificAnnouncements` | `boolean` | no |
| `MaxLength` | `integer` | no |
| `Description` | `string` | no |
| `Force` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrationDelta` | `IpamRoutingPolicyRegistrationDelta` | no |

## CreateIpamScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `ExternalAuthorityConfiguration` | `ExternalAuthorityConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamScope` | `IpamScope` | no |

## CreateKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | yes |
| `KeyType` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `KeyFormat` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyPairId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `KeyName` | `string` | no |
| `KeyFingerprint` | `string` | no |
| `KeyMaterial` | `string` | no |

## CreateLaunchTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `LaunchTemplateName` | `string` | yes |
| `VersionDescription` | `string` | no |
| `LaunchTemplateData` | `RequestLaunchTemplateData` | yes |
| `Operator` | `OperatorRequest` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplate` | `LaunchTemplate` | no |
| `Warning` | `ValidationWarning` | no |

## CreateLaunchTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `LaunchTemplateId` | `string` | no |
| `LaunchTemplateName` | `string` | no |
| `SourceVersion` | `string` | no |
| `VersionDescription` | `string` | no |
| `LaunchTemplateData` | `RequestLaunchTemplateData` | yes |
| `ResolveAlias` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplateVersion` | `LaunchTemplateVersion` | no |
| `Warning` | `ValidationWarning` | no |

## CreateLocalGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | no |
| `LocalGatewayRouteTableId` | `string` | yes |
| `LocalGatewayVirtualInterfaceGroupId` | `string` | no |
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | no |
| `DestinationPrefixListId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `LocalGatewayRoute` | no |

## CreateLocalGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayId` | `string` | yes |
| `Mode` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTable` | `LocalGatewayRouteTable` | no |

## CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableId` | `string` | yes |
| `LocalGatewayVirtualInterfaceGroupId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociation` | `LocalGatewayRouteTableVirtualInterfaceGroupAssociation` | no |

## CreateLocalGatewayRouteTableVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableId` | `string` | yes |
| `VpcId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVpcAssociation` | `LocalGatewayRouteTableVpcAssociation` | no |

## CreateLocalGatewayVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroupId` | `string` | yes |
| `OutpostLagId` | `string` | yes |
| `Vlan` | `integer` | yes |
| `LocalAddress` | `string` | yes |
| `PeerAddress` | `string` | yes |
| `PeerBgpAsn` | `integer` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `PeerBgpAsnExtended` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterface` | `LocalGatewayVirtualInterface` | no |

## CreateLocalGatewayVirtualInterfaceGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayId` | `string` | yes |
| `LocalBgpAsn` | `integer` | no |
| `LocalBgpAsnExtended` | `long` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroup` | `LocalGatewayVirtualInterfaceGroup` | no |

## CreateMacSystemIntegrityProtectionModificationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `MacCredentials` | `string` | no |
| `MacSystemIntegrityProtectionConfiguration` | `MacSystemIntegrityProtectionConfigurationRequest` | no |
| `MacSystemIntegrityProtectionStatus` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MacModificationTask` | `MacModificationTask` | no |

## CreateManagedPrefixList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListName` | `string` | yes |
| `Entries` | `List<AddPrefixListEntry>` | no |
| `MaxEntries` | `integer` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `AddressFamily` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixList` | `ManagedPrefixList` | no |

## CreateNatGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityMode` | `string` | no |
| `AllocationId` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SubnetId` | `string` | no |
| `VpcId` | `string` | no |
| `AvailabilityZoneAddresses` | `List<AvailabilityZoneAddress>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ConnectivityType` | `string` | no |
| `PrivateIpAddress` | `string` | no |
| `SecondaryAllocationIds` | `List<string>` | no |
| `SecondaryPrivateIpAddresses` | `List<string>` | no |
| `SecondaryPrivateIpAddressCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `NatGateway` | `NatGateway` | no |

## CreateNetworkAcl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkAcl` | `NetworkAcl` | no |
| `ClientToken` | `string` | no |

## CreateNetworkAclEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkAclId` | `string` | yes |
| `RuleNumber` | `integer` | yes |
| `Protocol` | `string` | yes |
| `RuleAction` | `string` | yes |
| `Egress` | `boolean` | yes |
| `CidrBlock` | `string` | no |
| `Ipv6CidrBlock` | `string` | no |
| `IcmpTypeCode` | `IcmpTypeCode` | no |
| `PortRange` | `PortRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateNetworkInsightsAccessScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MatchPaths` | `List<AccessScopePathRequest>` | no |
| `ExcludePaths` | `List<AccessScopePathRequest>` | no |
| `ClientToken` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScope` | `NetworkInsightsAccessScope` | no |
| `NetworkInsightsAccessScopeContent` | `NetworkInsightsAccessScopeContent` | no |

## CreateNetworkInsightsPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceIp` | `string` | no |
| `DestinationIp` | `string` | no |
| `Source` | `string` | yes |
| `Destination` | `string` | no |
| `Protocol` | `string` | yes |
| `DestinationPort` | `integer` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | yes |
| `FilterAtSource` | `PathRequestFilter` | no |
| `FilterAtDestination` | `PathRequestFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsPath` | `NetworkInsightsPath` | no |

## CreateNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv4Prefixes` | `List<Ipv4PrefixSpecificationRequest>` | no |
| `Ipv4PrefixCount` | `integer` | no |
| `Ipv6Prefixes` | `List<Ipv6PrefixSpecificationRequest>` | no |
| `Ipv6PrefixCount` | `integer` | no |
| `InterfaceType` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `EnablePrimaryIpv6` | `boolean` | no |
| `ConnectionTrackingSpecification` | `ConnectionTrackingSpecificationRequest` | no |
| `Operator` | `OperatorRequest` | no |
| `SubnetId` | `string` | yes |
| `Description` | `string` | no |
| `PrivateIpAddress` | `string` | no |
| `Groups` | `List<string>` | no |
| `PrivateIpAddresses` | `List<PrivateIpAddressSpecification>` | no |
| `SecondaryPrivateIpAddressCount` | `integer` | no |
| `Ipv6Addresses` | `List<InstanceIpv6Address>` | no |
| `Ipv6AddressCount` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterface` | `NetworkInterface` | no |
| `ClientToken` | `string` | no |

## CreateNetworkInterfacePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | yes |
| `AwsAccountId` | `string` | no |
| `AwsService` | `string` | no |
| `Permission` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InterfacePermission` | `NetworkInterfacePermission` | no |

## CreatePlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartitionCount` | `integer` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `SpreadLevel` | `string` | no |
| `LinkedGroupId` | `string` | no |
| `Operator` | `OperatorRequest` | no |
| `ParentGroupId` | `string` | no |
| `DryRun` | `boolean` | no |
| `GroupName` | `string` | no |
| `Strategy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlacementGroup` | `PlacementGroup` | no |

## CreatePublicIpv4Pool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `NetworkBorderGroup` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | no |

## CreateReplaceRootVolumeTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SnapshotId` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ImageId` | `string` | no |
| `DeleteReplacedRootVolume` | `boolean` | no |
| `VolumeInitializationRate` | `long` | no |
| `VolumeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplaceRootVolumeTask` | `ReplaceRootVolumeTask` | no |

## CreateReservedInstancesListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesId` | `string` | yes |
| `InstanceCount` | `integer` | yes |
| `PriceSchedules` | `List<PriceScheduleSpecification>` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesListings` | `List<ReservedInstancesListing>` | no |

## CreateRestoreImageTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bucket` | `string` | yes |
| `ObjectKey` | `string` | yes |
| `Name` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## CreateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPrefixListId` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `TransitGatewayId` | `string` | no |
| `LocalGatewayId` | `string` | no |
| `CarrierGatewayId` | `string` | no |
| `CoreNetworkArn` | `string` | no |
| `OdbNetworkArn` | `string` | no |
| `DryRun` | `boolean` | no |
| `RouteTableId` | `string` | yes |
| `DestinationCidrBlock` | `string` | no |
| `GatewayId` | `string` | no |
| `DestinationIpv6CidrBlock` | `string` | no |
| `EgressOnlyInternetGatewayId` | `string` | no |
| `InstanceId` | `string` | no |
| `NetworkInterfaceId` | `string` | no |
| `VpcPeeringConnectionId` | `string` | no |
| `NatGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## CreateRouteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmazonSideAsn` | `long` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `PersistRoutes` | `string` | no |
| `PersistRoutesDuration` | `long` | no |
| `SnsNotificationsEnabled` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServer` | `RouteServer` | no |

## CreateRouteServerEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `SubnetId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpoint` | `RouteServerEndpoint` | no |

## CreateRouteServerPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpointId` | `string` | yes |
| `PeerAddress` | `string` | yes |
| `BgpOptions` | `RouteServerBgpOptionsRequest` | yes |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPeer` | `RouteServerPeer` | no |

## CreateRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteTable` | `RouteTable` | no |
| `ClientToken` | `string` | no |

## CreateSecondaryNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `Ipv4CidrBlock` | `string` | yes |
| `NetworkType` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondaryNetwork` | `SecondaryNetwork` | no |
| `ClientToken` | `string` | no |

## CreateSecondarySubnet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `DryRun` | `boolean` | no |
| `Ipv4CidrBlock` | `string` | yes |
| `SecondaryNetworkId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondarySubnet` | `SecondarySubnet` | no |
| `ClientToken` | `string` | no |

## CreateSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | yes |
| `GroupName` | `string` | yes |
| `VpcId` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SecurityGroupArn` | `string` | no |

## CreateSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `OutpostArn` | `string` | no |
| `VolumeId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `Location` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerAlias` | `string` | no |
| `OutpostArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `StorageTier` | `string` | no |
| `RestoreExpiryTime` | `timestamp` | no |
| `SseType` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `TransferType` | `string` | no |
| `CompletionDurationMinutes` | `integer` | no |
| `CompletionTime` | `timestamp` | no |
| `FullSnapshotSizeInBytes` | `long` | no |
| `SnapshotId` | `string` | no |
| `VolumeId` | `string` | no |
| `State` | `string` | no |
| `StateMessage` | `string` | no |
| `StartTime` | `timestamp` | no |
| `Progress` | `string` | no |
| `OwnerId` | `string` | no |
| `Description` | `string` | no |
| `VolumeSize` | `integer` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `DataEncryptionKeyId` | `string` | no |

## CreateSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `InstanceSpecification` | `InstanceSpecification` | yes |
| `OutpostArn` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `CopyTagsFromSource` | `string` | no |
| `Location` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshots` | `List<SnapshotInfo>` | no |

## CreateSpotDatafeedSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Bucket` | `string` | yes |
| `Prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotDatafeedSubscription` | `SpotDatafeedSubscription` | no |

## CreateStoreImageTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `Bucket` | `string` | yes |
| `S3ObjectTags` | `List<S3ObjectTag>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectKey` | `string` | no |

## CreateSubnet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `CidrBlock` | `string` | no |
| `Ipv6CidrBlock` | `string` | no |
| `OutpostArn` | `string` | no |
| `VpcId` | `string` | yes |
| `Ipv6Native` | `boolean` | no |
| `Ipv4IpamPoolId` | `string` | no |
| `Ipv4NetmaskLength` | `integer` | no |
| `Ipv6IpamPoolId` | `string` | no |
| `Ipv6NetmaskLength` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subnet` | `Subnet` | no |

## CreateSubnetCidrReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetId` | `string` | yes |
| `Cidr` | `string` | yes |
| `ReservationType` | `string` | yes |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetCidrReservation` | `SubnetCidrReservation` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Resources` | `List<string>` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTrafficMirrorFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilter` | `TrafficMirrorFilter` | no |
| `ClientToken` | `string` | no |

## CreateTrafficMirrorFilterRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterId` | `string` | yes |
| `TrafficDirection` | `string` | yes |
| `RuleNumber` | `integer` | yes |
| `RuleAction` | `string` | yes |
| `DestinationPortRange` | `TrafficMirrorPortRangeRequest` | no |
| `SourcePortRange` | `TrafficMirrorPortRangeRequest` | no |
| `Protocol` | `integer` | no |
| `DestinationCidrBlock` | `string` | yes |
| `SourceCidrBlock` | `string` | yes |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRule` | `TrafficMirrorFilterRule` | no |
| `ClientToken` | `string` | no |

## CreateTrafficMirrorSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | yes |
| `TrafficMirrorTargetId` | `string` | yes |
| `TrafficMirrorFilterId` | `string` | yes |
| `PacketLength` | `integer` | no |
| `SessionNumber` | `integer` | yes |
| `VirtualNetworkId` | `integer` | no |
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSession` | `TrafficMirrorSession` | no |
| `ClientToken` | `string` | no |

## CreateTrafficMirrorTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | no |
| `NetworkLoadBalancerArn` | `string` | no |
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `GatewayLoadBalancerEndpointId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorTarget` | `TrafficMirrorTarget` | no |
| `ClientToken` | `string` | no |

## CreateTransitGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Options` | `TransitGatewayRequestOptions` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGateway` | `TransitGateway` | no |

## CreateTransitGatewayConnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransportTransitGatewayAttachmentId` | `string` | yes |
| `Options` | `CreateTransitGatewayConnectRequestOptions` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnect` | `TransitGatewayConnect` | no |

## CreateTransitGatewayConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `TransitGatewayAddress` | `string` | no |
| `PeerAddress` | `string` | yes |
| `BgpOptions` | `TransitGatewayConnectRequestBgpOptions` | no |
| `InsideCidrBlocks` | `List<string>` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeer` | `TransitGatewayConnectPeer` | no |

## CreateTransitGatewayMeteringPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `MiddleboxAttachmentIds` | `List<string>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicy` | `TransitGatewayMeteringPolicy` | no |

## CreateTransitGatewayMeteringPolicyEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyId` | `string` | yes |
| `PolicyRuleNumber` | `integer` | yes |
| `SourceTransitGatewayAttachmentId` | `string` | no |
| `SourceTransitGatewayAttachmentType` | `string` | no |
| `SourceCidrBlock` | `string` | no |
| `SourcePortRange` | `string` | no |
| `DestinationTransitGatewayAttachmentId` | `string` | no |
| `DestinationTransitGatewayAttachmentType` | `string` | no |
| `DestinationCidrBlock` | `string` | no |
| `DestinationPortRange` | `string` | no |
| `Protocol` | `string` | no |
| `MeteredAccount` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyEntry` | `TransitGatewayMeteringPolicyEntry` | no |

## CreateTransitGatewayMulticastDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `Options` | `CreateTransitGatewayMulticastDomainRequestOptions` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomain` | `TransitGatewayMulticastDomain` | no |

## CreateTransitGatewayPeeringAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `PeerTransitGatewayId` | `string` | yes |
| `PeerAccountId` | `string` | yes |
| `PeerRegion` | `string` | yes |
| `Options` | `CreateTransitGatewayPeeringAttachmentRequestOptions` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeeringAttachment` | `TransitGatewayPeeringAttachment` | no |

## CreateTransitGatewayPolicyTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTable` | `TransitGatewayPolicyTable` | no |

## CreateTransitGatewayPolicyTableEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `PolicyRuleNumber` | `string` | yes |
| `PolicyRule` | `TransitGatewayRequestPolicyRule` | no |
| `TargetRouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableEntry` | `TransitGatewayPolicyTableEntry` | no |

## CreateTransitGatewayPrefixListReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `PrefixListId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `Blackhole` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPrefixListReference` | `TransitGatewayPrefixListReference` | no |

## CreateTransitGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | yes |
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `Blackhole` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `TransitGatewayRoute` | no |

## CreateTransitGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTable` | `TransitGatewayRouteTable` | no |

## CreateTransitGatewayRouteTableAnnouncement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `PeeringAttachmentId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAnnouncement` | `TransitGatewayRouteTableAnnouncement` | no |

## CreateTransitGatewayVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `VpcId` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `Options` | `CreateTransitGatewayVpcAttachmentRequestOptions` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachment` | `TransitGatewayVpcAttachment` | no |

## CreateVerifiedAccessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupId` | `string` | yes |
| `EndpointType` | `string` | yes |
| `AttachmentType` | `string` | yes |
| `DomainCertificateArn` | `string` | no |
| `ApplicationDomain` | `string` | no |
| `EndpointDomainPrefix` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `LoadBalancerOptions` | `CreateVerifiedAccessEndpointLoadBalancerOptions` | no |
| `NetworkInterfaceOptions` | `CreateVerifiedAccessEndpointEniOptions` | no |
| `Description` | `string` | no |
| `PolicyDocument` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |
| `RdsOptions` | `CreateVerifiedAccessEndpointRdsOptions` | no |
| `CidrOptions` | `CreateVerifiedAccessEndpointCidrOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpoint` | `VerifiedAccessEndpoint` | no |

## CreateVerifiedAccessGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `Description` | `string` | no |
| `PolicyDocument` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroup` | `VerifiedAccessGroup` | no |

## CreateVerifiedAccessInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `FIPSEnabled` | `boolean` | no |
| `CidrEndpointsCustomSubDomain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstance` | `VerifiedAccessInstance` | no |

## CreateVerifiedAccessTrustProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustProviderType` | `string` | yes |
| `UserTrustProviderType` | `string` | no |
| `DeviceTrustProviderType` | `string` | no |
| `OidcOptions` | `CreateVerifiedAccessTrustProviderOidcOptions` | no |
| `DeviceOptions` | `CreateVerifiedAccessTrustProviderDeviceOptions` | no |
| `PolicyReferenceName` | `string` | yes |
| `Description` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |
| `NativeApplicationOidcOptions` | `CreateVerifiedAccessNativeApplicationOidcOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProvider` | `VerifiedAccessTrustProvider` | no |

## CreateVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `AvailabilityZoneId` | `string` | no |
| `Encrypted` | `boolean` | no |
| `Iops` | `integer` | no |
| `KmsKeyId` | `string` | no |
| `OutpostArn` | `string` | no |
| `Size` | `integer` | no |
| `SnapshotId` | `string` | no |
| `VolumeType` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `MultiAttachEnabled` | `boolean` | no |
| `Throughput` | `integer` | no |
| `ClientToken` | `string` | no |
| `VolumeInitializationRate` | `integer` | no |
| `Operator` | `OperatorRequest` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZoneId` | `string` | no |
| `OutpostArn` | `string` | no |
| `SourceVolumeId` | `string` | no |
| `Iops` | `integer` | no |
| `Tags` | `List<Tag>` | no |
| `VolumeType` | `string` | no |
| `FastRestored` | `boolean` | no |
| `MultiAttachEnabled` | `boolean` | no |
| `Throughput` | `integer` | no |
| `SseType` | `string` | no |
| `Operator` | `OperatorResponse` | no |
| `VolumeInitializationRate` | `integer` | no |
| `VolumeId` | `string` | no |
| `Size` | `integer` | no |
| `SnapshotId` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `State` | `string` | no |
| `CreateTime` | `timestamp` | no |
| `Attachments` | `List<VolumeAttachment>` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |

## CreateVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CidrBlock` | `string` | no |
| `Ipv6Pool` | `string` | no |
| `Ipv6CidrBlock` | `string` | no |
| `Ipv4IpamPoolId` | `string` | no |
| `Ipv4NetmaskLength` | `integer` | no |
| `Ipv6IpamPoolId` | `string` | no |
| `Ipv6NetmaskLength` | `integer` | no |
| `Ipv6CidrBlockNetworkBorderGroup` | `string` | no |
| `VpcEncryptionControl` | `VpcEncryptionControlConfiguration` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `InstanceTenancy` | `string` | no |
| `AmazonProvidedIpv6CidrBlock` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Vpc` | `Vpc` | no |

## CreateVpcBlockPublicAccessExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SubnetId` | `string` | no |
| `VpcId` | `string` | no |
| `InternetGatewayExclusionMode` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessExclusion` | `VpcBlockPublicAccessExclusion` | no |

## CreateVpcEncryptionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEncryptionControl` | `VpcEncryptionControl` | no |

## CreateVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEndpointType` | `string` | no |
| `VpcId` | `string` | yes |
| `ServiceName` | `string` | no |
| `PolicyDocument` | `string` | no |
| `RouteTableIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `DnsOptions` | `DnsOptionsSpecification` | no |
| `ClientToken` | `string` | no |
| `PrivateDnsEnabled` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `SubnetConfigurations` | `List<SubnetConfiguration>` | no |
| `ServiceNetworkArn` | `string` | no |
| `ResourceConfigurationArn` | `string` | no |
| `ServiceRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpoint` | `VpcEndpoint` | no |
| `ClientToken` | `string` | no |

## CreateVpcEndpointConnectionNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `ConnectionNotificationArn` | `string` | yes |
| `ConnectionEvents` | `List<string>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionNotification` | `ConnectionNotification` | no |
| `ClientToken` | `string` | no |

## CreateVpcEndpointServiceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AcceptanceRequired` | `boolean` | no |
| `PrivateDnsName` | `string` | no |
| `NetworkLoadBalancerArns` | `List<string>` | no |
| `GatewayLoadBalancerArns` | `List<string>` | no |
| `SupportedIpAddressTypes` | `List<string>` | no |
| `SupportedRegions` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceConfiguration` | `ServiceConfiguration` | no |
| `ClientToken` | `string` | no |

## CreateVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PeerRegion` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |
| `PeerVpcId` | `string` | no |
| `PeerOwnerId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringConnection` | `VpcPeeringConnection` | no |

## CreateVpnConcentrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `TransitGatewayId` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConcentrator` | `VpnConcentrator` | no |

## CreateVpnConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayId` | `string` | yes |
| `Type` | `string` | yes |
| `VpnGatewayId` | `string` | no |
| `TransitGatewayId` | `string` | no |
| `VpnConcentratorId` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `PreSharedKeyStorage` | `string` | no |
| `DryRun` | `boolean` | no |
| `Options` | `VpnConnectionOptionsSpecification` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnection` | `VpnConnection` | no |

## CreateVpnConnectionRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | yes |
| `VpnConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateVpnGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `Type` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `AmazonSideAsn` | `long` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnGateway` | `VpnGateway` | no |

## DeleteApplicationStatusCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheck` | `ApplicationStatusCheckResponseObject` | no |

## DeleteCapacityManagerDataExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerDataExportId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerDataExportId` | `string` | no |

## DeleteCarrierGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CarrierGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CarrierGateway` | `CarrierGateway` | no |

## DeleteClientVpnEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ClientVpnEndpointStatus` | no |

## DeleteClientVpnRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `TargetVpcSubnetId` | `string` | no |
| `DestinationCidrBlock` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ClientVpnRouteStatus` | no |

## DeleteCoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `CoipPoolId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipCidr` | `CoipCidr` | no |

## DeleteCoipPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipPoolId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipPool` | `CoipPool` | no |

## DeleteCustomerGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDhcpOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DhcpOptionsId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEgressOnlyInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `EgressOnlyInternetGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnCode` | `boolean` | no |

## DeleteFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FleetIds` | `List<string>` | yes |
| `TerminateInstances` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulFleetDeletions` | `List<DeleteFleetSuccessItem>` | no |
| `UnsuccessfulFleetDeletions` | `List<DeleteFleetErrorItem>` | no |

## DeleteFlowLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FlowLogIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## DeleteFpgaImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FpgaImageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DeleteImageUsageReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DeleteInstanceConnectEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceConnectEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceConnectEndpoint` | `Ec2InstanceConnectEndpoint` | no |

## DeleteInstanceEventWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ForceDelete` | `boolean` | no |
| `InstanceEventWindowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindowState` | `InstanceEventWindowStateChange` | no |

## DeleteInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InternetGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIpam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Cascade` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipam` | `Ipam` | no |

## DeleteIpamExternalResourceVerificationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamExternalResourceVerificationTokenId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamExternalResourceVerificationToken` | `IpamExternalResourceVerificationToken` | no |

## DeleteIpamInternetRegistryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamInternetRegistryAssociation` | `IpamInternetRegistryAssociation` | no |

## DeleteIpamPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicy` | `IpamPolicy` | no |

## DeleteIpamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Cascade` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPool` | `IpamPool` | no |

## DeleteIpamPrefixListResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolver` | `IpamPrefixListResolver` | no |

## DeleteIpamPrefixListResolverTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverTargetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolverTarget` | `IpamPrefixListResolverTarget` | no |

## DeleteIpamResourceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscovery` | `IpamResourceDiscovery` | no |

## DeleteIpamRoutingPolicyRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `Cidr` | `string` | yes |
| `Force` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrationDelta` | `IpamRoutingPolicyRegistrationDelta` | no |

## DeleteIpamScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamScopeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamScope` | `IpamScope` | no |

## DeleteKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | no |
| `KeyPairId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `KeyPairId` | `string` | no |

## DeleteLaunchTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `LaunchTemplateId` | `string` | no |
| `LaunchTemplateName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplate` | `LaunchTemplate` | no |

## DeleteLaunchTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `LaunchTemplateId` | `string` | no |
| `LaunchTemplateName` | `string` | no |
| `Versions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfullyDeletedLaunchTemplateVersions` | `List<DeleteLaunchTemplateVersionsResponseSuccessItem>` | no |
| `UnsuccessfullyDeletedLaunchTemplateVersions` | `List<DeleteLaunchTemplateVersionsResponseErrorItem>` | no |

## DeleteLocalGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | no |
| `LocalGatewayRouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `DestinationPrefixListId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `LocalGatewayRoute` | no |

## DeleteLocalGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTable` | `LocalGatewayRouteTable` | no |

## DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociation` | `LocalGatewayRouteTableVirtualInterfaceGroupAssociation` | no |

## DeleteLocalGatewayRouteTableVpcAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVpcAssociationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVpcAssociation` | `LocalGatewayRouteTableVpcAssociation` | no |

## DeleteLocalGatewayVirtualInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterface` | `LocalGatewayVirtualInterface` | no |

## DeleteLocalGatewayVirtualInterfaceGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroupId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroup` | `LocalGatewayVirtualInterfaceGroup` | no |

## DeleteManagedPrefixList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixList` | `ManagedPrefixList` | no |

## DeleteNatGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NatGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |

## DeleteNetworkAcl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkAclId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkAclEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkAclId` | `string` | yes |
| `RuleNumber` | `integer` | yes |
| `Egress` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkInsightsAccessScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInsightsAccessScopeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeId` | `string` | no |

## DeleteNetworkInsightsAccessScopeAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysisId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysisId` | `string` | no |

## DeleteNetworkInsightsAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInsightsAnalysisId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAnalysisId` | `string` | no |

## DeleteNetworkInsightsPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInsightsPathId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsPathId` | `string` | no |

## DeleteNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkInterfacePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfacePermissionId` | `string` | yes |
| `Force` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DeletePlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `GroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePublicIpv4Pool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PoolId` | `string` | yes |
| `NetworkBorderGroup` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## DeleteQueuedReservedInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ReservedInstancesIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulQueuedPurchaseDeletions` | `List<SuccessfulQueuedPurchaseDeletion>` | no |
| `FailedQueuedPurchaseDeletions` | `List<FailedQueuedPurchaseDeletion>` | no |

## DeleteRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPrefixListId` | `string` | no |
| `DryRun` | `boolean` | no |
| `RouteTableId` | `string` | yes |
| `DestinationCidrBlock` | `string` | no |
| `DestinationIpv6CidrBlock` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRouteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServer` | `RouteServer` | no |

## DeleteRouteServerEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpoint` | `RouteServerEndpoint` | no |

## DeleteRouteServerPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPeerId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPeer` | `RouteServerPeer` | no |

## DeleteRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `RouteTableId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecondaryNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SecondaryNetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondaryNetwork` | `SecondaryNetwork` | no |
| `ClientToken` | `string` | no |

## DeleteSecondarySubnet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SecondarySubnetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondarySubnet` | `SecondarySubnet` | no |
| `ClientToken` | `string` | no |

## DeleteSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | no |
| `GroupName` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `GroupId` | `string` | no |

## DeleteSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSpotDatafeedSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubnet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubnetCidrReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetCidrReservationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletedSubnetCidrReservation` | `SubnetCidrReservation` | no |

## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Resources` | `List<string>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrafficMirrorFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterId` | `string` | no |

## DeleteTrafficMirrorFilterRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRuleId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRuleId` | `string` | no |

## DeleteTrafficMirrorSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSessionId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSessionId` | `string` | no |

## DeleteTrafficMirrorTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorTargetId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorTargetId` | `string` | no |

## DeleteTransitGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGateway` | `TransitGateway` | no |

## DeleteTransitGatewayClientVpnAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayClientVpnAttachment` | `TransitGatewayClientVpnAttachment` | no |

## DeleteTransitGatewayConnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnect` | `TransitGatewayConnect` | no |

## DeleteTransitGatewayConnectPeer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeerId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeer` | `TransitGatewayConnectPeer` | no |

## DeleteTransitGatewayMeteringPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicy` | `TransitGatewayMeteringPolicy` | no |

## DeleteTransitGatewayMeteringPolicyEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyId` | `string` | yes |
| `PolicyRuleNumber` | `integer` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyEntry` | `TransitGatewayMeteringPolicyEntry` | no |

## DeleteTransitGatewayMulticastDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomain` | `TransitGatewayMulticastDomain` | no |

## DeleteTransitGatewayPeeringAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeeringAttachment` | `TransitGatewayPeeringAttachment` | no |

## DeleteTransitGatewayPolicyTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTable` | `TransitGatewayPolicyTable` | no |

## DeleteTransitGatewayPolicyTableEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `PolicyRuleNumber` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableEntry` | `TransitGatewayPolicyTableEntry` | no |

## DeleteTransitGatewayPrefixListReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `PrefixListId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPrefixListReference` | `TransitGatewayPrefixListReference` | no |

## DeleteTransitGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `DestinationCidrBlock` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `TransitGatewayRoute` | no |

## DeleteTransitGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTable` | `TransitGatewayRouteTable` | no |

## DeleteTransitGatewayRouteTableAnnouncement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAnnouncementId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAnnouncement` | `TransitGatewayRouteTableAnnouncement` | no |

## DeleteTransitGatewayVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachment` | `TransitGatewayVpcAttachment` | no |

## DeleteVerifiedAccessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpoint` | `VerifiedAccessEndpoint` | no |

## DeleteVerifiedAccessGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroup` | `VerifiedAccessGroup` | no |

## DeleteVerifiedAccessInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstance` | `VerifiedAccessInstance` | no |

## DeleteVerifiedAccessTrustProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProviderId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProvider` | `VerifiedAccessTrustProvider` | no |

## DeleteVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcBlockPublicAccessExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ExclusionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessExclusion` | `VpcBlockPublicAccessExclusion` | no |

## DeleteVpcEncryptionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEncryptionControlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEncryptionControl` | `VpcEncryptionControl` | no |

## DeleteVpcEndpointConnectionNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ConnectionNotificationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## DeleteVpcEndpointServiceConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## DeleteVpcEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEndpointIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## DeleteVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DeleteVpnConcentrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConcentratorId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DeleteVpnConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpnConnectionRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | yes |
| `VpnConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpnGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprovisionByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## DeprovisionIpamByoasn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Asn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Byoasn` | `Byoasn` | no |

## DeprovisionIpamPoolCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Cidr` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolCidr` | `IpamPoolCidr` | no |

## DeprovisionPublicIpv4PoolCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PoolId` | `string` | yes |
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | no |
| `DeprovisionedAddresses` | `List<string>` | no |

## DeregisterImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DeleteAssociatedSnapshots` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `DeleteSnapshotResults` | `List<DeleteSnapshotReturnCode>` | no |

## DeregisterInstanceEventNotificationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceTagAttribute` | `DeregisterInstanceTagAttributeRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTagAttribute` | `InstanceTagNotificationAttribute` | no |

## DeregisterTransitGatewayMulticastGroupMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | no |
| `GroupIpAddress` | `string` | no |
| `NetworkInterfaceIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeregisteredMulticastGroupMembers` | `TransitGatewayMulticastDeregisteredGroupMembers` | no |

## DeregisterTransitGatewayMulticastGroupSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | no |
| `GroupIpAddress` | `string` | no |
| `NetworkInterfaceIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeregisteredMulticastGroupSources` | `TransitGatewayMulticastDeregisteredGroupSources` | no |

## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAttributes` | `List<AccountAttribute>` | no |

## DescribeAccountVpcEncryptionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountVpcEncryptionControl` | `AccountVpcEncryptionControl` | no |

## DescribeAddressTransfers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressTransfers` | `List<AddressTransfer>` | no |
| `NextToken` | `string` | no |

## DescribeAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicIps` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `AllocationIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Addresses` | `List<Address>` | no |

## DescribeAddressesAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationIds` | `List<string>` | no |
| `Attribute` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Addresses` | `List<AddressAttribute>` | no |
| `NextToken` | `string` | no |

## DescribeAggregateIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UseLongIdsAggregated` | `boolean` | no |
| `Statuses` | `List<IdFormat>` | no |

## DescribeApplicationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatuses` | `ApplicationStatusesResponseType` | no |
| `NextToken` | `string` | no |

## DescribeApplicationStatusCheckAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<ApplicationStatusCheckAssociationObject>` | no |
| `NextToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

## DescribeApplicationStatusChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeAll` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusChecks` | `List<ApplicationStatusCheckResponseObject>` | no |
| `NextToken` | `string` | no |

## DescribeAvailabilityZones

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ZoneNames` | `List<string>` | no |
| `ZoneIds` | `List<string>` | no |
| `AllAvailabilityZones` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<AvailabilityZone>` | no |

## DescribeAwsNetworkPerformanceMetricSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Subscriptions` | `List<Subscription>` | no |

## DescribeBundleTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleTasks` | `List<BundleTask>` | no |

## DescribeByoipCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidrs` | `List<ByoipCidr>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityBlockExtensionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockExtensions` | `List<CapacityBlockExtension>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityBlockExtensionOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityBlockExtensionDurationHours` | `integer` | yes |
| `CapacityReservationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockExtensionOfferings` | `List<CapacityBlockExtensionOffering>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityBlockOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceType` | `string` | no |
| `InstanceCount` | `integer` | no |
| `StartDateRange` | `timestamp` | no |
| `EndDateRange` | `timestamp` | no |
| `CapacityDurationHours` | `integer` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `UltraserverType` | `string` | no |
| `UltraserverCount` | `integer` | no |
| `AllAvailabilityZones` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockOfferings` | `List<CapacityBlockOffering>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityBlockStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockStatuses` | `List<CapacityBlockStatus>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlocks` | `List<CapacityBlock>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityManagerDataExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerDataExportIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerDataExports` | `List<CapacityManagerDataExportResponse>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityReservationBillingRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationIds` | `List<string>` | no |
| `Role` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CapacityReservationBillingRequests` | `List<CapacityReservationBillingRequest>` | no |

## DescribeCapacityReservationCancellationQuotes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationCancellationQuoteIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationCancellationQuotes` | `List<CapacityReservationCancellationQuote>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityReservationFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationFleetIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationFleets` | `List<CapacityReservationFleet>` | no |
| `NextToken` | `string` | no |

## DescribeCapacityReservationTopology

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CapacityReservationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CapacityReservations` | `List<CapacityReservationTopology>` | no |

## DescribeCapacityReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CapacityReservations` | `List<CapacityReservation>` | no |

## DescribeCarrierGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CarrierGatewayIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CarrierGateways` | `List<CarrierGateway>` | no |
| `NextToken` | `string` | no |

## DescribeClassicLinkInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<ClassicLinkInstance>` | no |
| `NextToken` | `string` | no |

## DescribeClientVpnAuthorizationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizationRules` | `List<AuthorizationRule>` | no |
| `NextToken` | `string` | no |

## DescribeClientVpnConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<ClientVpnConnection>` | no |
| `NextToken` | `string` | no |

## DescribeClientVpnEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpoints` | `List<ClientVpnEndpoint>` | no |
| `NextToken` | `string` | no |

## DescribeClientVpnRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Routes` | `List<ClientVpnRoute>` | no |
| `NextToken` | `string` | no |

## DescribeClientVpnTargetNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `AssociationIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnTargetNetworks` | `List<TargetNetwork>` | no |
| `NextToken` | `string` | no |

## DescribeCoipPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipPools` | `List<CoipPool>` | no |
| `NextToken` | `string` | no |

## DescribeConversionTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ConversionTaskIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConversionTasks` | `List<ConversionTask>` | no |

## DescribeCustomerGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGatewayIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomerGateways` | `List<CustomerGateway>` | no |

## DescribeDeclarativePoliciesReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ReportIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Reports` | `List<DeclarativePoliciesReport>` | no |

## DescribeDhcpOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DhcpOptionsIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DhcpOptions` | `List<DhcpOptions>` | no |

## DescribeEgressOnlyInternetGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `EgressOnlyInternetGatewayIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EgressOnlyInternetGateways` | `List<EgressOnlyInternetGateway>` | no |
| `NextToken` | `string` | no |

## DescribeElasticGpus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ElasticGpuIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ElasticGpuSet` | `List<ElasticGpus>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

## DescribeExportImageTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `ExportImageTaskIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportImageTasks` | `List<ExportImageTask>` | no |
| `NextToken` | `string` | no |

## DescribeExportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `ExportTaskIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportTasks` | `List<ExportTask>` | no |

## DescribeFastLaunchImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FastLaunchImages` | `List<DescribeFastLaunchImagesSuccessItem>` | no |
| `NextToken` | `string` | no |

## DescribeFastSnapshotRestores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FastSnapshotRestores` | `List<DescribeFastSnapshotRestoreSuccessItem>` | no |
| `NextToken` | `string` | no |

## DescribeFleetHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `EventType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `FleetId` | `string` | yes |
| `StartTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HistoryRecords` | `List<HistoryRecordEntry>` | no |
| `LastEvaluatedTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `FleetId` | `string` | no |
| `StartTime` | `timestamp` | no |

## DescribeFleetInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `FleetId` | `string` | yes |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveInstances` | `List<ActiveInstance>` | no |
| `NextToken` | `string` | no |
| `FleetId` | `string` | no |

## DescribeFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `FleetIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Fleets` | `List<FleetData>` | no |

## DescribeFlowLogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filter` | `List<Filter>` | no |
| `FlowLogIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowLogs` | `List<FlowLog>` | no |
| `NextToken` | `string` | no |

## DescribeFpgaImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FpgaImageId` | `string` | yes |
| `Attribute` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FpgaImageAttribute` | `FpgaImageAttribute` | no |

## DescribeFpgaImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FpgaImageIds` | `List<string>` | no |
| `Owners` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FpgaImages` | `List<FpgaImage>` | no |
| `NextToken` | `string` | no |

## DescribeHostReservationOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `List<Filter>` | no |
| `MaxDuration` | `integer` | no |
| `MaxResults` | `integer` | no |
| `MinDuration` | `integer` | no |
| `NextToken` | `string` | no |
| `OfferingId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `OfferingSet` | `List<HostOffering>` | no |

## DescribeHostReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `List<Filter>` | no |
| `HostReservationIdSet` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostReservationSet` | `List<HostReservation>` | no |
| `NextToken` | `string` | no |

## DescribeHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Hosts` | `List<Host>` | no |
| `NextToken` | `string` | no |

## DescribeIamInstanceProfileAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfileAssociations` | `List<IamInstanceProfileAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statuses` | `List<IdFormat>` | no |

## DescribeIdentityIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | no |
| `PrincipalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statuses` | `List<IdFormat>` | no |

## DescribeImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `AttributeValue` | no |
| `KernelId` | `AttributeValue` | no |
| `RamdiskId` | `AttributeValue` | no |
| `SriovNetSupport` | `AttributeValue` | no |
| `BootMode` | `AttributeValue` | no |
| `TpmSupport` | `AttributeValue` | no |
| `UefiData` | `AttributeValue` | no |
| `LastLaunchedTime` | `AttributeValue` | no |
| `ImdsSupport` | `AttributeValue` | no |
| `DeregistrationProtection` | `AttributeValue` | no |
| `ImageId` | `string` | no |
| `LaunchPermissions` | `List<LaunchPermission>` | no |
| `ProductCodes` | `List<ProductCode>` | no |
| `BlockDeviceMappings` | `List<BlockDeviceMapping>` | no |

## DescribeImageReferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | yes |
| `IncludeAllResourceTypes` | `boolean` | no |
| `ResourceTypes` | `List<ResourceTypeRequest>` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ImageReferences` | `List<ImageReference>` | no |

## DescribeImageUsageReportEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `ReportIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ImageUsageReportEntries` | `List<ImageUsageReportEntry>` | no |

## DescribeImageUsageReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `ReportIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ImageUsageReports` | `List<ImageUsageReport>` | no |

## DescribeImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutableUsers` | `List<string>` | no |
| `ImageIds` | `List<string>` | no |
| `Owners` | `List<string>` | no |
| `IncludeDeprecated` | `boolean` | no |
| `IncludeDisabled` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Images` | `List<Image>` | no |

## DescribeImportImageTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `ImportTaskIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportImageTasks` | `List<ImportImageTask>` | no |
| `NextToken` | `string` | no |

## DescribeImportSnapshotTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `ImportTaskIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportSnapshotTasks` | `List<ImportSnapshotTask>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `Attribute` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockDeviceMappings` | `List<InstanceBlockDeviceMapping>` | no |
| `DisableApiTermination` | `AttributeBooleanValue` | no |
| `EnaSupport` | `AttributeBooleanValue` | no |
| `EnclaveOptions` | `EnclaveOptions` | no |
| `EbsOptimized` | `AttributeBooleanValue` | no |
| `InstanceId` | `string` | no |
| `InstanceInitiatedShutdownBehavior` | `AttributeValue` | no |
| `InstanceType` | `AttributeValue` | no |
| `KernelId` | `AttributeValue` | no |
| `ProductCodes` | `List<ProductCode>` | no |
| `RamdiskId` | `AttributeValue` | no |
| `RootDeviceName` | `AttributeValue` | no |
| `SourceDestCheck` | `AttributeBooleanValue` | no |
| `SriovNetSupport` | `AttributeValue` | no |
| `UserData` | `AttributeValue` | no |
| `DisableApiStop` | `AttributeBooleanValue` | no |
| `Groups` | `List<GroupIdentifier>` | no |

## DescribeInstanceConnectEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `InstanceConnectEndpointIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceConnectEndpoints` | `List<Ec2InstanceConnectEndpoint>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceCreditSpecifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `InstanceIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceCreditSpecifications` | `List<InstanceCreditSpecification>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceEventNotificationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTagAttribute` | `InstanceTagNotificationAttribute` | no |

## DescribeInstanceEventWindows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceEventWindowIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindows` | `List<InstanceEventWindow>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceImageMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `InstanceIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceImageMetadata` | `List<InstanceImageMetadata>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceSqlHaHistoryStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<RegisteredInstance>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceSqlHaStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<RegisteredInstance>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeManagedResources` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `IncludeAllInstances` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceStatuses` | `List<InstanceStatus>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceTopology

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `InstanceIds` | `List<string>` | no |
| `GroupNames` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<InstanceTopology>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceTypeOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `LocationType` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypeOfferings` | `List<InstanceTypeOffering>` | no |
| `NextToken` | `string` | no |

## DescribeInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceTypes` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeUnsupportedInRegion` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<InstanceTypeInfo>` | no |
| `NextToken` | `string` | no |

## DescribeInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `IncludeManagedResources` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Reservations` | `List<Reservation>` | no |

## DescribeInternetGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `InternetGatewayIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InternetGateways` | `List<InternetGateway>` | no |
| `NextToken` | `string` | no |

## DescribeIpamByoasn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Byoasns` | `List<Byoasn>` | no |
| `NextToken` | `string` | no |

## DescribeIpamExternalResourceVerificationTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IpamExternalResourceVerificationTokenIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamExternalResourceVerificationTokens` | `List<IpamExternalResourceVerificationToken>` | no |

## DescribeIpamInternetRegistryAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamInternetRegistryAssociations` | `List<IpamInternetRegistryAssociation>` | no |

## DescribeIpamPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamPolicyIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamPolicies` | `List<IpamPolicy>` | no |

## DescribeIpamPoolAllocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolAllocationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolAllocations` | `List<IpamPoolAllocation>` | no |
| `NextToken` | `string` | no |

## DescribeIpamPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamPoolIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamPools` | `List<IpamPool>` | no |

## DescribeIpamPrefixListResolverTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamPrefixListResolverTargetIds` | `List<string>` | no |
| `IpamPrefixListResolverId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamPrefixListResolverTargets` | `List<IpamPrefixListResolverTarget>` | no |

## DescribeIpamPrefixListResolvers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamPrefixListResolverIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamPrefixListResolvers` | `List<IpamPrefixListResolver>` | no |

## DescribeIpamResourceDiscoveries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscoveries` | `List<IpamResourceDiscovery>` | no |
| `NextToken` | `string` | no |

## DescribeIpamResourceDiscoveryAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryAssociationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscoveryAssociations` | `List<IpamResourceDiscoveryAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeIpamScopes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamScopeIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamScopes` | `List<IpamScope>` | no |

## DescribeIpams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Ipams` | `List<Ipam>` | no |

## DescribeIpv6Pools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6Pools` | `List<Ipv6Pool>` | no |
| `NextToken` | `string` | no |

## DescribeKeyPairs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyNames` | `List<string>` | no |
| `KeyPairIds` | `List<string>` | no |
| `IncludePublicKey` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyPairs` | `List<KeyPairInfo>` | no |

## DescribeLaunchTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `LaunchTemplateId` | `string` | no |
| `LaunchTemplateName` | `string` | no |
| `Versions` | `List<string>` | no |
| `MinVersion` | `string` | no |
| `MaxVersion` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `ResolveAlias` | `boolean` | no |
| `IncludeManagedResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplateVersions` | `List<LaunchTemplateVersion>` | no |
| `NextToken` | `string` | no |

## DescribeLaunchTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `LaunchTemplateIds` | `List<string>` | no |
| `LaunchTemplateNames` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeManagedResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplates` | `List<LaunchTemplate>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVirtualInterfaceGroupAssociations` | `List<LocalGatewayRouteTableVirtualInterfaceGroupAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGatewayRouteTableVpcAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVpcAssociationIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableVpcAssociations` | `List<LocalGatewayRouteTableVpcAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGatewayRouteTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTables` | `List<LocalGatewayRouteTable>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGatewayVirtualInterfaceGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroupIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceGroups` | `List<LocalGatewayVirtualInterfaceGroup>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGatewayVirtualInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayVirtualInterfaces` | `List<LocalGatewayVirtualInterface>` | no |
| `NextToken` | `string` | no |

## DescribeLocalGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGateways` | `List<LocalGateway>` | no |
| `NextToken` | `string` | no |

## DescribeLockedSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SnapshotIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshots` | `List<LockedSnapshotsInfo>` | no |
| `NextToken` | `string` | no |

## DescribeMacHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `HostIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MacHosts` | `List<MacHost>` | no |
| `NextToken` | `string` | no |

## DescribeMacModificationTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MacModificationTaskIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MacModificationTasks` | `List<MacModificationTask>` | no |
| `NextToken` | `string` | no |

## DescribeManagedPrefixLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PrefixListIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PrefixLists` | `List<ManagedPrefixList>` | no |

## DescribeMovingAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PublicIps` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MovingAddressStatuses` | `List<MovingAddressStatus>` | no |
| `NextToken` | `string` | no |

## DescribeNatGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filter` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NatGatewayIds` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGateways` | `List<NatGateway>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkAcls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NetworkAclIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkAcls` | `List<NetworkAcl>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInsightsAccessScopeAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysisIds` | `List<string>` | no |
| `NetworkInsightsAccessScopeId` | `string` | no |
| `AnalysisStartTimeBegin` | `timestamp` | no |
| `AnalysisStartTimeEnd` | `timestamp` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalyses` | `List<NetworkInsightsAccessScopeAnalysis>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInsightsAccessScopes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopes` | `List<NetworkInsightsAccessScope>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInsightsAnalyses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAnalysisIds` | `List<string>` | no |
| `NetworkInsightsPathId` | `string` | no |
| `AnalysisStartTime` | `timestamp` | no |
| `AnalysisEndTime` | `timestamp` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAnalyses` | `List<NetworkInsightsAnalysis>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInsightsPaths

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsPathIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsPaths` | `List<NetworkInsightsPath>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInterfaceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | yes |
| `Attribute` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attachment` | `NetworkInterfaceAttachment` | no |
| `Description` | `AttributeValue` | no |
| `Groups` | `List<GroupIdentifier>` | no |
| `NetworkInterfaceId` | `string` | no |
| `SourceDestCheck` | `AttributeBooleanValue` | no |
| `AssociatePublicIpAddress` | `boolean` | no |

## DescribeNetworkInterfacePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfacePermissionIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfacePermissions` | `List<NetworkInterfacePermission>` | no |
| `NextToken` | `string` | no |

## DescribeNetworkInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeManagedResources` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `NetworkInterfaceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaces` | `List<NetworkInterface>` | no |
| `NextToken` | `string` | no |

## DescribeOutpostLags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostLagIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OutpostLags` | `List<OutpostLag>` | no |
| `NextToken` | `string` | no |

## DescribePlacementGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `GroupNames` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlacementGroups` | `List<PlacementGroup>` | no |

## DescribePrefixLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PrefixListIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PrefixLists` | `List<PrefixList>` | no |

## DescribePrincipalIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Resources` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Principals` | `List<PrincipalIdFormat>` | no |
| `NextToken` | `string` | no |

## DescribePublicIpv4Pools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicIpv4Pools` | `List<PublicIpv4Pool>` | no |
| `NextToken` | `string` | no |

## DescribeRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionNames` | `List<string>` | no |
| `AllRegions` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<Region>` | no |

## DescribeReplaceRootVolumeTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplaceRootVolumeTaskIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplaceRootVolumeTasks` | `List<ReplaceRootVolumeTask>` | no |
| `NextToken` | `string` | no |

## DescribeReservedInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OfferingClass` | `string` | no |
| `ReservedInstancesIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `OfferingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstances` | `List<ReservedInstances>` | no |

## DescribeReservedInstancesListings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesId` | `string` | no |
| `ReservedInstancesListingId` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesListings` | `List<ReservedInstancesListing>` | no |

## DescribeReservedInstancesModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesModificationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedInstancesModifications` | `List<ReservedInstancesModification>` | no |

## DescribeReservedInstancesOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `IncludeMarketplace` | `boolean` | no |
| `InstanceType` | `string` | no |
| `MaxDuration` | `long` | no |
| `MaxInstanceCount` | `integer` | no |
| `MinDuration` | `long` | no |
| `OfferingClass` | `string` | no |
| `ProductDescription` | `string` | no |
| `ReservedInstancesOfferingIds` | `List<string>` | no |
| `AvailabilityZoneId` | `string` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `InstanceTenancy` | `string` | no |
| `OfferingType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedInstancesOfferings` | `List<ReservedInstancesOffering>` | no |

## DescribeRouteServerEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpointIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerEndpoints` | `List<RouteServerEndpoint>` | no |
| `NextToken` | `string` | no |

## DescribeRouteServerPeers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPeerIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPeers` | `List<RouteServerPeer>` | no |
| `NextToken` | `string` | no |

## DescribeRouteServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServers` | `List<RouteServer>` | no |
| `NextToken` | `string` | no |

## DescribeRouteTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `RouteTableIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteTables` | `List<RouteTable>` | no |
| `NextToken` | `string` | no |

## DescribeScheduledInstanceAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `FirstSlotStartTimeRange` | `SlotDateTimeRangeRequest` | yes |
| `MaxResults` | `integer` | no |
| `MaxSlotDurationInHours` | `integer` | no |
| `MinSlotDurationInHours` | `integer` | no |
| `NextToken` | `string` | no |
| `Recurrence` | `ScheduledInstanceRecurrenceRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ScheduledInstanceAvailabilitySet` | `List<ScheduledInstanceAvailability>` | no |

## DescribeScheduledInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ScheduledInstanceIds` | `List<string>` | no |
| `SlotStartTimeRange` | `SlotStartTimeRangeRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ScheduledInstanceSet` | `List<ScheduledInstance>` | no |

## DescribeSecondaryInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SecondaryInterfaceIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondaryInterfaces` | `List<SecondaryInterface>` | no |
| `NextToken` | `string` | no |

## DescribeSecondaryNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SecondaryNetworkIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondaryNetworks` | `List<SecondaryNetwork>` | no |
| `NextToken` | `string` | no |

## DescribeSecondarySubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SecondarySubnetIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecondarySubnets` | `List<SecondarySubnet>` | no |
| `NextToken` | `string` | no |

## DescribeSecurityGroupReferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `GroupId` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupReferenceSet` | `List<SecurityGroupReference>` | no |

## DescribeSecurityGroupRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `SecurityGroupRuleIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupRules` | `List<SecurityGroupRule>` | no |
| `NextToken` | `string` | no |

## DescribeSecurityGroupVpcAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupVpcAssociations` | `List<SecurityGroupVpcAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIds` | `List<string>` | no |
| `GroupNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SecurityGroups` | `List<SecurityGroup>` | no |

## DescribeServiceLinkVirtualInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceLinkVirtualInterfaceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceLinkVirtualInterfaces` | `List<ServiceLinkVirtualInterface>` | no |
| `NextToken` | `string` | no |

## DescribeSnapshotAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductCodes` | `List<ProductCode>` | no |
| `SnapshotId` | `string` | no |
| `CreateVolumePermissions` | `List<CreateVolumePermission>` | no |

## DescribeSnapshotTierStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotTierStatuses` | `List<SnapshotTierStatus>` | no |
| `NextToken` | `string` | no |

## DescribeSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `OwnerIds` | `List<string>` | no |
| `RestorableByUserIds` | `List<string>` | no |
| `SnapshotIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Snapshots` | `List<Snapshot>` | no |

## DescribeSpotDatafeedSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotDatafeedSubscription` | `SpotDatafeedSubscription` | no |

## DescribeSpotFleetInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotFleetRequestId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveInstances` | `List<ActiveInstance>` | no |
| `NextToken` | `string` | no |
| `SpotFleetRequestId` | `string` | no |

## DescribeSpotFleetRequestHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotFleetRequestId` | `string` | yes |
| `EventType` | `string` | no |
| `StartTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HistoryRecords` | `List<HistoryRecord>` | no |
| `LastEvaluatedTime` | `timestamp` | no |
| `NextToken` | `string` | no |
| `SpotFleetRequestId` | `string` | no |
| `StartTime` | `timestamp` | no |

## DescribeSpotFleetRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotFleetRequestIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SpotFleetRequestConfigs` | `List<SpotFleetRequestConfig>` | no |

## DescribeSpotInstanceRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `SpotInstanceRequestIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotInstanceRequests` | `List<SpotInstanceRequest>` | no |
| `NextToken` | `string` | no |

## DescribeSpotPriceHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZoneId` | `string` | no |
| `DryRun` | `boolean` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `InstanceTypes` | `List<string>` | no |
| `ProductDescriptions` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `AvailabilityZone` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SpotPriceHistory` | `List<SpotPrice>` | no |

## DescribeStaleSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StaleSecurityGroupSet` | `List<StaleSecurityGroup>` | no |

## DescribeStoreImageTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StoreImageTaskResults` | `List<StoreImageTaskResult>` | no |
| `NextToken` | `string` | no |

## DescribeSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `SubnetIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Subnets` | `List<Subnet>` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Tags` | `List<TagDescription>` | no |

## DescribeTrafficMirrorFilterRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRuleIds` | `List<string>` | no |
| `TrafficMirrorFilterId` | `string` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRules` | `List<TrafficMirrorFilterRule>` | no |
| `NextToken` | `string` | no |

## DescribeTrafficMirrorFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilters` | `List<TrafficMirrorFilter>` | no |
| `NextToken` | `string` | no |

## DescribeTrafficMirrorSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSessionIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSessions` | `List<TrafficMirrorSession>` | no |
| `NextToken` | `string` | no |

## DescribeTrafficMirrorTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorTargetIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorTargets` | `List<TrafficMirrorTarget>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachments` | `List<TransitGatewayAttachment>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayConnectPeers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeerIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnectPeers` | `List<TransitGatewayConnectPeer>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayConnects` | `List<TransitGatewayConnect>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayMeteringPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicies` | `List<TransitGatewayMeteringPolicy>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayMulticastDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomains` | `List<TransitGatewayMulticastDomain>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayPeeringAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeeringAttachments` | `List<TransitGatewayPeeringAttachment>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayPolicyTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTables` | `List<TransitGatewayPolicyTable>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayRouteTableAnnouncements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAnnouncementIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableAnnouncements` | `List<TransitGatewayRouteTableAnnouncement>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayRouteTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTables` | `List<TransitGatewayRouteTable>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGatewayVpcAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachments` | `List<TransitGatewayVpcAttachment>` | no |
| `NextToken` | `string` | no |

## DescribeTransitGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGateways` | `List<TransitGateway>` | no |
| `NextToken` | `string` | no |

## DescribeTrunkInterfaceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InterfaceAssociations` | `List<TrunkInterfaceAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedAccessEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointIds` | `List<string>` | no |
| `VerifiedAccessInstanceId` | `string` | no |
| `VerifiedAccessGroupId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpoints` | `List<VerifiedAccessEndpoint>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedAccessGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupIds` | `List<string>` | no |
| `VerifiedAccessInstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroups` | `List<VerifiedAccessGroup>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedAccessInstanceLoggingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfigurations` | `List<VerifiedAccessInstanceLoggingConfiguration>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedAccessInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstances` | `List<VerifiedAccessInstance>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedAccessTrustProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProviderIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProviders` | `List<VerifiedAccessTrustProvider>` | no |
| `NextToken` | `string` | no |

## DescribeVolumeAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `VolumeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnableIO` | `AttributeBooleanValue` | no |
| `ProductCodes` | `List<ProductCode>` | no |
| `VolumeId` | `string` | no |

## DescribeVolumeStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `VolumeIds` | `List<string>` | no |
| `IncludeManagedResources` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VolumeStatuses` | `List<VolumeStatusItem>` | no |

## DescribeVolumes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeIds` | `List<string>` | no |
| `IncludeManagedResources` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Volumes` | `List<Volume>` | no |

## DescribeVolumesModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VolumeIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeManagedResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `VolumesModifications` | `List<VolumeModification>` | no |

## DescribeVpcAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableDnsHostnames` | `AttributeBooleanValue` | no |
| `EnableDnsSupport` | `AttributeBooleanValue` | no |
| `EnableNetworkAddressUsageMetrics` | `AttributeBooleanValue` | no |
| `VpcId` | `string` | no |

## DescribeVpcBlockPublicAccessExclusions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `ExclusionIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessExclusions` | `List<VpcBlockPublicAccessExclusion>` | no |
| `NextToken` | `string` | no |

## DescribeVpcBlockPublicAccessOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessOptions` | `VpcBlockPublicAccessOptions` | no |

## DescribeVpcClassicLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Vpcs` | `List<VpcClassicLink>` | no |

## DescribeVpcClassicLinkDnsSupport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Vpcs` | `List<ClassicLinkDnsSupport>` | no |

## DescribeVpcEncryptionControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `VpcEncryptionControlIds` | `List<string>` | no |
| `VpcIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEncryptionControls` | `List<VpcEncryptionControl>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEndpointIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointAssociations` | `List<VpcEndpointAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointConnectionNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ConnectionNotificationId` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionNotificationSet` | `List<ConnectionNotification>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointConnections` | `List<VpcEndpointConnection>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointServiceConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceConfigurations` | `List<ServiceConfiguration>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointServicePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowedPrincipals` | `List<AllowedPrincipal>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpointServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceNames` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServiceRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceNames` | `List<string>` | no |
| `ServiceDetails` | `List<ServiceDetail>` | no |
| `NextToken` | `string` | no |

## DescribeVpcEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEndpointIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpoints` | `List<VpcEndpoint>` | no |
| `NextToken` | `string` | no |

## DescribeVpcPeeringConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `VpcPeeringConnectionIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcPeeringConnections` | `List<VpcPeeringConnection>` | no |
| `NextToken` | `string` | no |

## DescribeVpcs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `VpcIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Vpcs` | `List<Vpc>` | no |

## DescribeVpnConcentrators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConcentratorIds` | `List<string>` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConcentrators` | `List<VpnConcentrator>` | no |
| `NextToken` | `string` | no |

## DescribeVpnConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `VpnConnectionIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnections` | `List<VpnConnection>` | no |

## DescribeVpnGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `VpnGatewayIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnGateways` | `List<VpnGateway>` | no |

## DetachClassicLinkVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DetachImageWatermark

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `WatermarkKey` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DetachInternetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InternetGatewayId` | `string` | yes |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AttachmentId` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachVerifiedAccessTrustProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `VerifiedAccessTrustProviderId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProvider` | `VerifiedAccessTrustProvider` | no |
| `VerifiedAccessInstance` | `VerifiedAccessInstance` | no |

## DetachVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `string` | no |
| `Force` | `boolean` | no |
| `InstanceId` | `string` | no |
| `VolumeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeleteOnTermination` | `boolean` | no |
| `AssociatedResource` | `string` | no |
| `InstanceOwningService` | `string` | no |
| `EbsCardIndex` | `integer` | no |
| `VolumeId` | `string` | no |
| `InstanceId` | `string` | no |
| `Device` | `string` | no |
| `State` | `string` | no |
| `AttachTime` | `timestamp` | no |

## DetachVpnGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `VpnGatewayId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableAddressTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressTransfer` | `AddressTransfer` | no |

## DisableAllowedImagesSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowedImagesSettingsState` | `string` | no |

## DisableApplicationStatusCheckSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulResults` | `List<SuccessfulSuppressionResponseObject>` | no |
| `UnsuccessfulResults` | `List<UnsuccessfulSuppressionResponseObject>` | no |

## DisableAwsNetworkPerformanceMetricSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `string` | no |
| `Destination` | `string` | no |
| `Metric` | `string` | no |
| `Statistic` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Output` | `boolean` | no |

## DisableCapacityManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerStatus` | `string` | no |
| `OrganizationsAccess` | `boolean` | no |

## DisableEbsEncryptionByDefault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EbsEncryptionByDefault` | `boolean` | no |

## DisableFastLaunch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `Force` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `ResourceType` | `string` | no |
| `SnapshotConfiguration` | `FastLaunchSnapshotConfigurationResponse` | no |
| `LaunchTemplate` | `FastLaunchLaunchTemplateSpecificationResponse` | no |
| `MaxParallelLaunches` | `integer` | no |
| `OwnerId` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `StateTransitionTime` | `timestamp` | no |

## DisableFastSnapshotRestores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `AvailabilityZoneIds` | `List<string>` | no |
| `SourceSnapshotIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<DisableFastSnapshotRestoreSuccessItem>` | no |
| `Unsuccessful` | `List<DisableFastSnapshotRestoreErrorItem>` | no |

## DisableImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisableImageBlockPublicAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBlockPublicAccessState` | `string` | no |

## DisableImageDeprecation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisableImageDeregistrationProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `string` | no |

## DisableInstanceSqlHaStandbyDetections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<RegisteredInstance>` | no |

## DisableIpamOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `DelegatedAdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | no |

## DisableIpamPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPolicyId` | `string` | yes |
| `OrganizationTargetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisableRouteServerPropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `RouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPropagation` | `RouteServerPropagation` | no |

## DisableSerialConsoleAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialConsoleAccessEnabled` | `boolean` | no |

## DisableSnapshotBlockPublicAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |

## DisableTransitGatewayRouteTablePropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `DryRun` | `boolean` | no |
| `TransitGatewayRouteTableAnnouncementId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Propagation` | `TransitGatewayPropagation` | no |

## DisableVgwRoutePropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayId` | `string` | yes |
| `RouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableVpcClassicLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisableVpcClassicLinkDnsSupport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisassociateAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |
| `PublicIp` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApplicationStatusCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckId` | `string` | yes |
| `TargetTagAssociations` | `List<CustomTagKeyValueRequestPair>` | no |
| `InstanceIds` | `List<string>` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulResults` | `List<SuccessfulAssociationResponseObject>` | no |
| `UnsuccessfulResults` | `List<UnsuccessfulAssociationResponseObject>` | no |

## DisassociateCapacityReservationBillingOwner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityReservationId` | `string` | yes |
| `UnusedReservationBillingOwnerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisassociateClientVpnTargetNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |
| `Status` | `AssociationStatus` | no |

## DisassociateEnclaveCertificateIamRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## DisassociateIamInstanceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfileAssociation` | `IamInstanceProfileAssociation` | no |

## DisassociateInstanceEventWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceEventWindowId` | `string` | yes |
| `AssociationTarget` | `InstanceEventWindowDisassociationRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindow` | `InstanceEventWindow` | no |

## DisassociateIpamByoasn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Asn` | `string` | yes |
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AsnAssociation` | `AsnAssociation` | no |

## DisassociateIpamResourceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscoveryAssociation` | `IpamResourceDiscoveryAssociation` | no |

## DisassociateNatGatewayAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `AssociationIds` | `List<string>` | yes |
| `MaxDrainDurationSeconds` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |
| `NatGatewayAddresses` | `List<NatGatewayAddress>` | no |

## DisassociateRouteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerAssociation` | `RouteServerAssociation` | no |

## DisassociateRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSecurityGroupVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `VpcId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |

## DisassociateSubnetCidrBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6CidrBlockAssociation` | `SubnetIpv6CidrBlockAssociation` | no |
| `SubnetId` | `string` | no |

## DisassociateTransitGatewayMulticastDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `SubnetIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `TransitGatewayMulticastDomainAssociations` | no |

## DisassociateTransitGatewayPolicyTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `TransitGatewayPolicyTableAssociation` | no |

## DisassociateTransitGatewayRouteTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `TransitGatewayAssociation` | no |

## DisassociateTrunkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `ClientToken` | `string` | no |

## DisassociateVpcCidrBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6CidrBlockAssociation` | `VpcIpv6CidrBlockAssociation` | no |
| `CidrBlockAssociation` | `VpcCidrBlockAssociation` | no |
| `VpcId` | `string` | no |

## EnableAddressTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | yes |
| `TransferAccountId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddressTransfer` | `AddressTransfer` | no |

## EnableAllowedImagesSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowedImagesSettingsState` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowedImagesSettingsState` | `string` | no |

## EnableApplicationStatusCheckSuppression

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | no |
| `DurationSeconds` | `integer` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulResults` | `List<SuccessfulSuppressionResponseObject>` | no |
| `UnsuccessfulResults` | `List<UnsuccessfulSuppressionResponseObject>` | no |

## EnableAwsNetworkPerformanceMetricSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `string` | no |
| `Destination` | `string` | no |
| `Metric` | `string` | no |
| `Statistic` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Output` | `boolean` | no |

## EnableCapacityManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationsAccess` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerStatus` | `string` | no |
| `OrganizationsAccess` | `boolean` | no |

## EnableEbsEncryptionByDefault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EbsEncryptionByDefault` | `boolean` | no |

## EnableFastLaunch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `ResourceType` | `string` | no |
| `SnapshotConfiguration` | `FastLaunchSnapshotConfigurationRequest` | no |
| `LaunchTemplate` | `FastLaunchLaunchTemplateSpecificationRequest` | no |
| `MaxParallelLaunches` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `ResourceType` | `string` | no |
| `SnapshotConfiguration` | `FastLaunchSnapshotConfigurationResponse` | no |
| `LaunchTemplate` | `FastLaunchLaunchTemplateSpecificationResponse` | no |
| `MaxParallelLaunches` | `integer` | no |
| `OwnerId` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `StateTransitionTime` | `timestamp` | no |

## EnableFastSnapshotRestores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |
| `AvailabilityZoneIds` | `List<string>` | no |
| `SourceSnapshotIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<EnableFastSnapshotRestoreSuccessItem>` | no |
| `Unsuccessful` | `List<EnableFastSnapshotRestoreErrorItem>` | no |

## EnableImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## EnableImageBlockPublicAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBlockPublicAccessState` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBlockPublicAccessState` | `string` | no |

## EnableImageDeprecation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DeprecateAt` | `timestamp` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## EnableImageDeregistrationProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `WithCooldown` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `string` | no |

## EnableInstanceSqlHaStandbyDetections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `SqlServerCredentials` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<RegisteredInstance>` | no |

## EnableIpamInternetRegistryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `RpkiVersion` | `string` | yes |
| `ServiceUri` | `string` | yes |
| `ChildHandle` | `string` | yes |
| `ParentHandle` | `string` | yes |
| `ParentBpkiTa` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamInternetRegistryAssociation` | `IpamInternetRegistryAssociation` | no |

## EnableIpamOrganizationAdminAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `DelegatedAdminAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | no |

## EnableIpamPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPolicyId` | `string` | yes |
| `OrganizationTargetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicyId` | `string` | no |

## EnableReachabilityAnalyzerOrganizationSharing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## EnableRouteServerPropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `RouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPropagation` | `RouteServerPropagation` | no |

## EnableSerialConsoleAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialConsoleAccessEnabled` | `boolean` | no |

## EnableSnapshotBlockPublicAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |

## EnableTransitGatewayRouteTablePropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `DryRun` | `boolean` | no |
| `TransitGatewayRouteTableAnnouncementId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Propagation` | `TransitGatewayPropagation` | no |

## EnableVgwRoutePropagation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayId` | `string` | yes |
| `RouteTableId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableVolumeIO

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VolumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableVpcClassicLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## EnableVpcClassicLinkDnsSupport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ExportClientVpnClientCertificateRevocationList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateRevocationList` | `string` | no |
| `Status` | `ClientCertificateRevocationListStatus` | no |

## ExportClientVpnClientConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientConfiguration` | `string` | no |

## ExportImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `DiskImageFormat` | `string` | yes |
| `DryRun` | `boolean` | no |
| `ImageId` | `string` | yes |
| `S3ExportLocation` | `ExportTaskS3LocationRequest` | yes |
| `RoleName` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DiskImageFormat` | `string` | no |
| `ExportImageTaskId` | `string` | no |
| `ImageId` | `string` | no |
| `RoleName` | `string` | no |
| `Progress` | `string` | no |
| `S3ExportLocation` | `ExportTaskS3Location` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ExportTransitGatewayRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `S3Bucket` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3Location` | `string` | no |

## ExportVerifiedAccessInstanceClientConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Version` | `string` | no |
| `VerifiedAccessInstanceId` | `string` | no |
| `Region` | `string` | no |
| `DeviceTrustProviders` | `List<string>` | no |
| `UserTrustProvider` | `VerifiedAccessInstanceUserTrustProviderClientConfiguration` | no |
| `OpenVpnConfigurations` | `List<VerifiedAccessInstanceOpenVpnClientConfiguration>` | no |

## GetActiveVpnTunnelStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnTunnelOutsideIpAddress` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActiveVpnTunnelStatus` | `ActiveVpnTunnelStatus` | no |

## GetAllowedImagesSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `ImageCriteria` | `List<ImageCriterion>` | no |
| `ManagedBy` | `string` | no |

## GetAssociatedEnclaveCertificateIamRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedRoles` | `List<AssociatedRole>` | no |

## GetAssociatedIpv6PoolCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6CidrAssociations` | `List<Ipv6CidrAssociation>` | no |
| `NextToken` | `string` | no |

## GetAwsNetworkPerformanceData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataQueries` | `List<DataQuery>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataResponses` | `List<DataResponse>` | no |
| `NextToken` | `string` | no |

## GetCapacityManagerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerStatus` | `string` | no |
| `OrganizationsAccess` | `boolean` | no |
| `DataExportCount` | `integer` | no |
| `IngestionStatus` | `string` | no |
| `IngestionStatusMessage` | `string` | no |
| `EarliestDatapointTimestamp` | `timestamp` | no |
| `LatestDatapointTimestamp` | `timestamp` | no |

## GetCapacityManagerMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricNames` | `List<string>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Period` | `integer` | yes |
| `GroupBy` | `List<string>` | no |
| `FilterBy` | `List<CapacityManagerCondition>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDataResults` | `List<MetricDataResult>` | no |
| `NextToken` | `string` | no |

## GetCapacityManagerMetricDimensions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupBy` | `List<string>` | yes |
| `FilterBy` | `List<CapacityManagerCondition>` | no |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `MetricNames` | `List<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricDimensionResults` | `List<CapacityManagerDimension>` | no |
| `NextToken` | `string` | no |

## GetCapacityManagerMonitoredTagKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerTagKeys` | `List<CapacityManagerMonitoredTagKey>` | no |
| `NextToken` | `string` | no |

## GetCapacityReservationUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CapacityReservationId` | `string` | no |
| `InstanceType` | `string` | no |
| `TotalInstanceCount` | `integer` | no |
| `AvailableInstanceCount` | `integer` | no |
| `State` | `string` | no |
| `InstanceUsages` | `List<InstanceUsage>` | no |
| `Interruptible` | `boolean` | no |
| `InterruptibleCapacityAllocation` | `InterruptibleCapacityAllocation` | no |
| `InterruptionInfo` | `InterruptionInfo` | no |

## GetCoipPoolUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoipPoolId` | `string` | no |
| `CoipAddressUsages` | `List<CoipAddressUsage>` | no |
| `LocalGatewayRouteTableId` | `string` | no |
| `NextToken` | `string` | no |

## GetConsoleOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Latest` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `Timestamp` | `timestamp` | no |
| `Output` | `string` | no |

## GetConsoleScreenshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `WakeUp` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageData` | `string` | no |
| `InstanceId` | `string` | no |

## GetDeclarativePoliciesReportSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | no |
| `S3Bucket` | `string` | no |
| `S3Prefix` | `string` | no |
| `TargetId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `NumberOfAccounts` | `integer` | no |
| `NumberOfFailedAccounts` | `integer` | no |
| `AttributeSummaries` | `List<AttributeSummary>` | no |

## GetDefaultCreditSpecification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceFamily` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceFamilyCreditSpecification` | `InstanceFamilyCreditSpecification` | no |

## GetEbsDefaultKmsKeyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |

## GetEbsEncryptionByDefault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EbsEncryptionByDefault` | `boolean` | no |
| `SseType` | `string` | no |

## GetEnabledIpamPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicyEnabled` | `boolean` | no |
| `IpamPolicyId` | `string` | no |
| `ManagedBy` | `string` | no |

## GetFlowLogsIntegrationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FlowLogId` | `string` | yes |
| `ConfigDeliveryS3DestinationArn` | `string` | yes |
| `IntegrateServices` | `IntegrateServices` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `string` | no |

## GetGroupsForCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CapacityReservationGroups` | `List<CapacityReservationGroup>` | no |

## GetHostReservationPurchasePreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostIdSet` | `List<string>` | yes |
| `OfferingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CurrencyCode` | `string` | no |
| `Purchase` | `List<Purchase>` | no |
| `TotalHourlyPrice` | `string` | no |
| `TotalUpfrontPrice` | `string` | no |

## GetImageAncestry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageAncestryEntries` | `List<ImageAncestryEntry>` | no |

## GetImageBlockPublicAccessState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBlockPublicAccessState` | `string` | no |
| `ManagedBy` | `string` | no |

## GetInstanceMetadataDefaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLevel` | `InstanceMetadataDefaultsResponse` | no |

## GetInstanceTpmEkPub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `KeyType` | `string` | yes |
| `KeyFormat` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `KeyType` | `string` | no |
| `KeyFormat` | `string` | no |
| `KeyValue` | `string` | no |

## GetInstanceTypesFromInstanceRequirements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ArchitectureTypes` | `List<string>` | yes |
| `VirtualizationTypes` | `List<string>` | yes |
| `InstanceRequirements` | `InstanceRequirementsRequest` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Context` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<InstanceTypeInfoFromInstanceRequirements>` | no |
| `NextToken` | `string` | no |

## GetInstanceUefiData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `UefiData` | `string` | no |

## GetIpamAddressHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Cidr` | `string` | yes |
| `IpamScopeId` | `string` | yes |
| `VpcId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HistoryRecords` | `List<IpamAddressHistoryRecord>` | no |
| `NextToken` | `string` | no |

## GetIpamDiscoveredAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |
| `DiscoveryRegion` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamDiscoveredAccounts` | `List<IpamDiscoveredAccount>` | no |
| `NextToken` | `string` | no |

## GetIpamDiscoveredPublicAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |
| `AddressRegion` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamDiscoveredPublicAddresses` | `List<IpamDiscoveredPublicAddress>` | no |
| `OldestSampleTime` | `timestamp` | no |
| `NextToken` | `string` | no |

## GetIpamDiscoveredResourceCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |
| `ResourceRegion` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamDiscoveredResourceCidrs` | `List<IpamDiscoveredResourceCidr>` | no |
| `NextToken` | `string` | no |

## GetIpamDiscoveredRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |
| `ResourceRegion` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamDiscoveredRoutes` | `List<IpamDiscoveredRoute>` | no |
| `NextToken` | `string` | no |

## GetIpamInternetRegistryAssociationAsns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamInternetRegistryAssociationAsns` | `List<IpamInternetRegistryAssociationAsn>` | no |

## GetIpamInternetRegistryAssociationCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamInternetRegistryAssociationCidrs` | `List<IpamInternetRegistryAssociationCidr>` | no |

## GetIpamPolicyAllocationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPolicyId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `Locale` | `string` | no |
| `ResourceType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicyDocuments` | `List<IpamPolicyDocument>` | no |
| `NextToken` | `string` | no |

## GetIpamPolicyOrganizationTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamPolicyId` | `string` | yes |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationTargets` | `List<IpamPolicyOrganizationTarget>` | no |
| `NextToken` | `string` | no |

## GetIpamPoolAllocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `IpamPoolAllocationId` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolAllocations` | `List<IpamPoolAllocation>` | no |
| `NextToken` | `string` | no |

## GetIpamPoolCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolCidrs` | `List<IpamPoolCidr>` | no |
| `NextToken` | `string` | no |

## GetIpamPrefixListResolverRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<IpamPrefixListResolverRule>` | no |
| `NextToken` | `string` | no |

## GetIpamPrefixListResolverVersionEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |
| `IpamPrefixListResolverVersion` | `long` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<IpamPrefixListResolverVersionEntry>` | no |
| `NextToken` | `string` | no |

## GetIpamPrefixListResolverVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |
| `IpamPrefixListResolverVersions` | `List<long>` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolverVersions` | `List<IpamPrefixListResolverVersion>` | no |
| `NextToken` | `string` | no |

## GetIpamResourceCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IpamScopeId` | `string` | yes |
| `IpamPoolId` | `string` | no |
| `ResourceId` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceTag` | `RequestIpamResourceTag` | no |
| `ResourceOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `IpamResourceCidrs` | `List<IpamResourceCidr>` | no |

## GetIpamRouteOriginAuthorizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `Cidr` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRouteOriginAuthorizations` | `List<IpamRouteOriginAuthorizationInfo>` | no |
| `NextToken` | `string` | no |

## GetIpamRouteProtectionFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamId` | `string` | no |
| `RouteProtectionFindings` | `List<IpamRouteProtectionFinding>` | no |
| `NextToken` | `string` | no |

## GetIpamRoutingPolicyRegistrationDeltas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `DeltaId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `ChronologicalOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrationDeltas` | `List<IpamRoutingPolicyRegistrationDelta>` | no |
| `NextToken` | `string` | no |

## GetIpamRoutingPolicyRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `Cidr` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrations` | `List<IpamRoutingPolicyRegistration>` | no |
| `NextToken` | `string` | no |

## GetLaunchTemplateData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplateData` | `ResponseLaunchTemplateData` | no |

## GetManagedPrefixListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixListAssociations` | `List<PrefixListAssociation>` | no |
| `NextToken` | `string` | no |

## GetManagedPrefixListEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListId` | `string` | yes |
| `TargetVersion` | `long` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<PrefixListEntry>` | no |
| `NextToken` | `string` | no |

## GetManagedResourceVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Visibility` | `ManagedResourceVisibilitySettings` | no |

## GetNetworkInsightsAccessScopeAnalysisFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysisId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysisId` | `string` | no |
| `AnalysisStatus` | `string` | no |
| `AnalysisFindings` | `List<AccessScopeAnalysisFinding>` | no |
| `NextToken` | `string` | no |

## GetNetworkInsightsAccessScopeContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeContent` | `NetworkInsightsAccessScopeContent` | no |

## GetPasswordData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `Timestamp` | `timestamp` | no |
| `PasswordData` | `string` | no |

## GetReservedInstancesExchangeQuote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ReservedInstanceIds` | `List<string>` | yes |
| `TargetConfigurations` | `List<TargetConfigurationRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CurrencyCode` | `string` | no |
| `IsValidExchange` | `boolean` | no |
| `OutputReservedInstancesWillExpireAt` | `timestamp` | no |
| `PaymentDue` | `string` | no |
| `ReservedInstanceValueRollup` | `ReservationValue` | no |
| `ReservedInstanceValueSet` | `List<ReservedInstanceReservationValue>` | no |
| `TargetConfigurationValueRollup` | `ReservationValue` | no |
| `TargetConfigurationValueSet` | `List<TargetReservationValue>` | no |
| `ValidationFailureReason` | `string` | no |

## GetRouteServerAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerAssociations` | `List<RouteServerAssociation>` | no |

## GetRouteServerPropagations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `RouteTableId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerPropagations` | `List<RouteServerPropagation>` | no |

## GetRouteServerRoutingDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AreRoutesPersisted` | `boolean` | no |
| `Routes` | `List<RouteServerRoute>` | no |
| `NextToken` | `string` | no |

## GetSecurityGroupsForVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SecurityGroupForVpcs` | `List<SecurityGroupForVpc>` | no |

## GetSerialConsoleAccessStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SerialConsoleAccessEnabled` | `boolean` | no |
| `ManagedBy` | `string` | no |

## GetSnapshotBlockPublicAccessState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `ManagedBy` | `string` | no |

## GetSpotPlacementScores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypes` | `List<string>` | no |
| `TargetCapacity` | `integer` | yes |
| `TargetCapacityUnitType` | `string` | no |
| `SingleAvailabilityZone` | `boolean` | no |
| `RegionNames` | `List<string>` | no |
| `InstanceRequirementsWithMetadata` | `InstanceRequirementsWithMetadataRequest` | no |
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `IncludeLocalZones` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotPlacementScores` | `List<SpotPlacementScore>` | no |
| `NextToken` | `string` | no |

## GetSubnetCidrReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `SubnetId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubnetIpv4CidrReservations` | `List<SubnetCidrReservation>` | no |
| `SubnetIpv6CidrReservations` | `List<SubnetCidrReservation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayAttachmentPropagations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentPropagations` | `List<TransitGatewayAttachmentPropagation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayMeteringPolicyEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyEntries` | `List<TransitGatewayMeteringPolicyEntry>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayMulticastDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MulticastDomainAssociations` | `List<TransitGatewayMulticastDomainAssociation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayPolicyTableAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<TransitGatewayPolicyTableAssociation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayPolicyTableEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableEntries` | `List<TransitGatewayPolicyTableEntry>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayPrefixListReferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPrefixListReferences` | `List<TransitGatewayPrefixListReference>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayRouteTableAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<TransitGatewayRouteTableAssociation>` | no |
| `NextToken` | `string` | no |

## GetTransitGatewayRouteTablePropagations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTablePropagations` | `List<TransitGatewayRouteTablePropagation>` | no |
| `NextToken` | `string` | no |

## GetVerifiedAccessEndpointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |

## GetVerifiedAccessEndpointTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointTargets` | `List<VerifiedAccessEndpointTarget>` | no |
| `NextToken` | `string` | no |

## GetVerifiedAccessGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |

## GetVpcResourcesBlockingEncryptionEnforcement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NonCompliantResources` | `List<VpcEncryptionNonCompliantResource>` | no |
| `NextToken` | `string` | no |

## GetVpnConnectionDeviceSampleConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnConnectionDeviceTypeId` | `string` | yes |
| `InternetKeyExchangeVersion` | `string` | no |
| `SampleType` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionDeviceSampleConfiguration` | `string` | no |

## GetVpnConnectionDeviceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionDeviceTypes` | `List<VpnConnectionDeviceType>` | no |
| `NextToken` | `string` | no |

## GetVpnTunnelReplacementStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnTunnelOutsideIpAddress` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | no |
| `TransitGatewayId` | `string` | no |
| `CustomerGatewayId` | `string` | no |
| `VpnGatewayId` | `string` | no |
| `VpnTunnelOutsideIpAddress` | `string` | no |
| `MaintenanceDetails` | `MaintenanceDetails` | no |

## ImportClientVpnClientCertificateRevocationList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `CertificateRevocationList` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ImportImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Architecture` | `string` | no |
| `ClientData` | `ClientData` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `DiskContainers` | `List<ImageDiskContainer>` | no |
| `DryRun` | `boolean` | no |
| `Encrypted` | `boolean` | no |
| `Hypervisor` | `string` | no |
| `KmsKeyId` | `string` | no |
| `LicenseType` | `string` | no |
| `Platform` | `string` | no |
| `RoleName` | `string` | no |
| `LicenseSpecifications` | `List<ImportImageLicenseConfigurationRequest>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `UsageOperation` | `string` | no |
| `BootMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Architecture` | `string` | no |
| `Description` | `string` | no |
| `Encrypted` | `boolean` | no |
| `Hypervisor` | `string` | no |
| `ImageId` | `string` | no |
| `ImportTaskId` | `string` | no |
| `KmsKeyId` | `string` | no |
| `LicenseType` | `string` | no |
| `Platform` | `string` | no |
| `Progress` | `string` | no |
| `SnapshotDetails` | `List<SnapshotDetail>` | no |
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `LicenseSpecifications` | `List<ImportImageLicenseConfigurationResponse>` | no |
| `Tags` | `List<Tag>` | no |
| `UsageOperation` | `string` | no |

## ImportInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Description` | `string` | no |
| `LaunchSpecification` | `ImportInstanceLaunchSpecification` | no |
| `DiskImages` | `List<DiskImage>` | no |
| `Platform` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConversionTask` | `ConversionTask` | no |

## ImportKeyPair

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `KeyName` | `string` | yes |
| `PublicKeyMaterial` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyFingerprint` | `string` | no |
| `KeyName` | `string` | no |
| `KeyPairId` | `string` | no |
| `Tags` | `List<Tag>` | no |

## ImportSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientData` | `ClientData` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `DiskContainer` | `SnapshotDiskContainer` | no |
| `DryRun` | `boolean` | no |
| `Encrypted` | `boolean` | no |
| `KmsKeyId` | `string` | no |
| `RoleName` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `ImportTaskId` | `string` | no |
| `SnapshotTaskDetail` | `SnapshotTaskDetail` | no |
| `Tags` | `List<Tag>` | no |

## ImportVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZoneId` | `string` | no |
| `DryRun` | `boolean` | no |
| `AvailabilityZone` | `string` | no |
| `Image` | `DiskImageDetail` | yes |
| `Description` | `string` | no |
| `Volume` | `VolumeDetail` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConversionTask` | `ConversionTask` | no |

## ListImagesInRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Images` | `List<ImageRecycleBinInfo>` | no |
| `NextToken` | `string` | no |

## ListSnapshotsInRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SnapshotIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Snapshots` | `List<SnapshotRecycleBinInfo>` | no |
| `NextToken` | `string` | no |

## ListVolumesInRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Volumes` | `List<VolumeRecycleBinInfo>` | no |
| `NextToken` | `string` | no |

## LockSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `LockMode` | `string` | yes |
| `CoolOffPeriod` | `integer` | no |
| `LockDuration` | `integer` | no |
| `ExpirationDate` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `LockState` | `string` | no |
| `LockDuration` | `integer` | no |
| `CoolOffPeriod` | `integer` | no |
| `CoolOffPeriodExpiresOn` | `timestamp` | no |
| `LockCreatedOn` | `timestamp` | no |
| `LockExpiresOn` | `timestamp` | no |
| `LockDurationStartTime` | `timestamp` | no |

## ModifyAccountVpcEncryptionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Mode` | `string` | no |
| `InternetGateway` | `string` | no |
| `EgressOnlyInternetGateway` | `string` | no |
| `NatGateway` | `string` | no |
| `VirtualPrivateGateway` | `string` | no |
| `VpcPeering` | `string` | no |
| `Lambda` | `string` | no |
| `VpcLattice` | `string` | no |
| `ElasticFileSystem` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountVpcEncryptionControl` | `AccountVpcEncryptionControl` | no |

## ModifyAddressAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | yes |
| `DomainName` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `AddressAttribute` | no |

## ModifyApplicationStatusCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheckId` | `string` | yes |
| `Aggregation` | `string` | no |
| `HealthCheckPaths` | `List<HealthCheckPathRequestObject>` | no |
| `Protocol` | `string` | no |
| `Port` | `integer` | no |
| `Path` | `string` | no |
| `DeviceIndex` | `integer` | no |
| `IpVersion` | `string` | no |
| `IpScope` | `string` | no |
| `Interval` | `integer` | no |
| `Timeout` | `integer` | no |
| `FailureThreshold` | `integer` | no |
| `SuccessThreshold` | `integer` | no |
| `StatusCodeMatcher` | `string` | no |
| `InitializationGracePeriodSeconds` | `integer` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatusCheck` | `ApplicationStatusCheckResponseObject` | no |

## ModifyAvailabilityZoneGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `OptInStatus` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `InstanceCount` | `integer` | no |
| `EndDate` | `timestamp` | no |
| `EndDateType` | `string` | no |
| `Accept` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `AdditionalInfo` | `string` | no |
| `InstanceMatchCriteria` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyCapacityReservationFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationFleetId` | `string` | yes |
| `TotalTargetCapacity` | `integer` | no |
| `EndDate` | `timestamp` | no |
| `DryRun` | `boolean` | no |
| `RemoveEndDate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyClientVpnEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `ServerCertificateArn` | `string` | no |
| `ConnectionLogOptions` | `ConnectionLogOptions` | no |
| `DnsServers` | `DnsServersOptionsModifyStructure` | no |
| `VpnPort` | `integer` | no |
| `Description` | `string` | no |
| `SplitTunnel` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `VpcId` | `string` | no |
| `SelfServicePortal` | `string` | no |
| `ClientConnectOptions` | `ClientConnectOptions` | no |
| `SessionTimeoutHours` | `integer` | no |
| `ClientLoginBannerOptions` | `ClientLoginBannerOptions` | no |
| `ClientRouteEnforcementOptions` | `ClientRouteEnforcementOptions` | no |
| `DisconnectOnSessionTimeout` | `boolean` | no |
| `TransitGatewayConfiguration` | `TransitGatewayConfigurationInputStructure` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyDefaultCreditSpecification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceFamily` | `string` | yes |
| `CpuCredits` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceFamilyCreditSpecification` | `InstanceFamilyCreditSpecification` | no |

## ModifyEbsDefaultKmsKeyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |

## ModifyFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ExcessCapacityTerminationPolicy` | `string` | no |
| `LaunchTemplateConfigs` | `List<FleetLaunchTemplateConfigRequest>` | no |
| `FleetId` | `string` | yes |
| `TargetCapacitySpecification` | `TargetCapacitySpecificationRequest` | no |
| `Context` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyFpgaImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FpgaImageId` | `string` | yes |
| `Attribute` | `string` | no |
| `OperationType` | `string` | no |
| `UserIds` | `List<string>` | no |
| `UserGroups` | `List<string>` | no |
| `ProductCodes` | `List<string>` | no |
| `LoadPermission` | `LoadPermissionModifications` | no |
| `Description` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FpgaImageAttribute` | `FpgaImageAttribute` | no |

## ModifyHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostRecovery` | `string` | no |
| `InstanceType` | `string` | no |
| `InstanceFamily` | `string` | no |
| `HostMaintenance` | `string` | no |
| `HostIds` | `List<string>` | yes |
| `AutoPlacement` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<string>` | no |
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## ModifyIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `UseLongIds` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyIdentityIdFormat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `UseLongIds` | `boolean` | yes |
| `PrincipalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | no |
| `Description` | `AttributeValue` | no |
| `ImageId` | `string` | yes |
| `LaunchPermission` | `LaunchPermissionModifications` | no |
| `OperationType` | `string` | no |
| `ProductCodes` | `List<string>` | no |
| `UserGroups` | `List<string>` | no |
| `UserIds` | `List<string>` | no |
| `Value` | `string` | no |
| `OrganizationArns` | `List<string>` | no |
| `OrganizationalUnitArns` | `List<string>` | no |
| `ImdsSupport` | `AttributeValue` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyInstanceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDestCheck` | `AttributeBooleanValue` | no |
| `EnclaveOptions` | `EnclaveOptionsRequest` | no |
| `DisableApiStop` | `AttributeBooleanValue` | no |
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `Attribute` | `string` | no |
| `Value` | `string` | no |
| `BlockDeviceMappings` | `List<InstanceBlockDeviceMappingSpecification>` | no |
| `DisableApiTermination` | `AttributeBooleanValue` | no |
| `InstanceType` | `AttributeValue` | no |
| `Kernel` | `AttributeValue` | no |
| `Ramdisk` | `AttributeValue` | no |
| `UserData` | `SecureBlobAttributeValue` | no |
| `InstanceInitiatedShutdownBehavior` | `AttributeValue` | no |
| `Groups` | `List<string>` | no |
| `EbsOptimized` | `AttributeBooleanValue` | no |
| `SriovNetSupport` | `AttributeValue` | no |
| `EnaSupport` | `AttributeBooleanValue` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyInstanceCapacityReservationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `CapacityReservationSpecification` | `CapacityReservationSpecification` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyInstanceConnectEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceConnectEndpointId` | `string` | yes |
| `IpAddressType` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `PreserveClientIp` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyInstanceCpuOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `CoreCount` | `integer` | no |
| `ThreadsPerCore` | `integer` | no |
| `NestedVirtualization` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `CoreCount` | `integer` | no |
| `ThreadsPerCore` | `integer` | no |
| `NestedVirtualization` | `string` | no |

## ModifyInstanceCreditSpecification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `InstanceCreditSpecifications` | `List<InstanceCreditSpecificationRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulInstanceCreditSpecifications` | `List<SuccessfulInstanceCreditSpecificationItem>` | no |
| `UnsuccessfulInstanceCreditSpecifications` | `List<UnsuccessfulInstanceCreditSpecificationItem>` | no |

## ModifyInstanceEventStartTime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `InstanceEventId` | `string` | yes |
| `NotBefore` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Event` | `InstanceStatusEvent` | no |

## ModifyInstanceEventWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Name` | `string` | no |
| `InstanceEventWindowId` | `string` | yes |
| `TimeRanges` | `List<InstanceEventWindowTimeRangeRequest>` | no |
| `CronExpression` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceEventWindow` | `InstanceEventWindow` | no |

## ModifyInstanceMaintenanceOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AutoRecovery` | `string` | no |
| `RebootMigration` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `AutoRecovery` | `string` | no |
| `RebootMigration` | `string` | no |

## ModifyInstanceMetadataDefaults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HttpTokens` | `string` | no |
| `HttpPutResponseHopLimit` | `integer` | no |
| `HttpEndpoint` | `string` | no |
| `InstanceMetadataTags` | `string` | no |
| `DryRun` | `boolean` | no |
| `HttpTokensEnforced` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyInstanceMetadataOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HttpTokens` | `string` | no |
| `HttpPutResponseHopLimit` | `integer` | no |
| `HttpEndpoint` | `string` | no |
| `DryRun` | `boolean` | no |
| `HttpProtocolIpv6` | `string` | no |
| `InstanceMetadataTags` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `InstanceMetadataOptions` | `InstanceMetadataOptionsResponse` | no |

## ModifyInstanceNetworkPerformanceOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `BandwidthWeighting` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `BandwidthWeighting` | `string` | no |

## ModifyInstancePlacement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `PartitionNumber` | `integer` | no |
| `HostResourceGroupArn` | `string` | no |
| `GroupId` | `string` | no |
| `InstanceId` | `string` | yes |
| `Tenancy` | `string` | no |
| `Affinity` | `string` | no |
| `HostId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyIpam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Description` | `string` | no |
| `AddOperatingRegions` | `List<AddIpamOperatingRegion>` | no |
| `RemoveOperatingRegions` | `List<RemoveIpamOperatingRegion>` | no |
| `Tier` | `string` | no |
| `EnablePrivateGua` | `boolean` | no |
| `MeteredAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipam` | `Ipam` | no |

## ModifyIpamPolicyAllocationRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPolicyId` | `string` | yes |
| `Locale` | `string` | yes |
| `ResourceType` | `string` | yes |
| `AllocationRules` | `List<IpamPolicyAllocationRuleRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPolicyDocument` | `IpamPolicyDocument` | no |

## ModifyIpamPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Description` | `string` | no |
| `AutoImport` | `boolean` | no |
| `AllocationMinNetmaskLength` | `integer` | no |
| `AllocationMaxNetmaskLength` | `integer` | no |
| `AllocationDefaultNetmaskLength` | `integer` | no |
| `ClearAllocationDefaultNetmaskLength` | `boolean` | no |
| `AddAllocationResourceTags` | `List<RequestIpamResourceTag>` | no |
| `RemoveAllocationResourceTags` | `List<RequestIpamResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPool` | `IpamPool` | no |

## ModifyIpamPoolAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolAllocationId` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolAllocation` | `IpamPoolAllocation` | no |

## ModifyIpamPrefixListResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverId` | `string` | yes |
| `Description` | `string` | no |
| `Rules` | `List<IpamPrefixListResolverRuleRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolver` | `IpamPrefixListResolver` | no |

## ModifyIpamPrefixListResolverTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPrefixListResolverTargetId` | `string` | yes |
| `DesiredVersion` | `long` | no |
| `TrackLatestVersion` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPrefixListResolverTarget` | `IpamPrefixListResolverTarget` | no |

## ModifyIpamResourceCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ResourceId` | `string` | yes |
| `ResourceCidr` | `string` | yes |
| `ResourceRegion` | `string` | yes |
| `CurrentIpamScopeId` | `string` | yes |
| `DestinationIpamScopeId` | `string` | no |
| `Monitored` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceCidr` | `IpamResourceCidr` | no |

## ModifyIpamResourceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamResourceDiscoveryId` | `string` | yes |
| `Description` | `string` | no |
| `AddOperatingRegions` | `List<AddIpamOperatingRegion>` | no |
| `RemoveOperatingRegions` | `List<RemoveIpamOperatingRegion>` | no |
| `AddOrganizationalUnitExclusions` | `List<AddIpamOrganizationalUnitExclusion>` | no |
| `RemoveOrganizationalUnitExclusions` | `List<RemoveIpamOrganizationalUnitExclusion>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamResourceDiscovery` | `IpamResourceDiscovery` | no |

## ModifyIpamRoutingPolicyRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamInternetRegistryAssociationId` | `string` | yes |
| `Cidr` | `string` | yes |
| `Asns` | `List<string>` | yes |
| `PermitMoreSpecificAnnouncements` | `boolean` | no |
| `MaxLength` | `integer` | no |
| `Description` | `string` | no |
| `Force` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamRoutingPolicyRegistrationDelta` | `IpamRoutingPolicyRegistrationDelta` | no |

## ModifyIpamScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamScopeId` | `string` | yes |
| `Description` | `string` | no |
| `ExternalAuthorityConfiguration` | `ExternalAuthorityConfiguration` | no |
| `RemoveExternalAuthorityConfiguration` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamScope` | `IpamScope` | no |

## ModifyLaunchTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `LaunchTemplateId` | `string` | no |
| `LaunchTemplateName` | `string` | no |
| `DefaultVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplate` | `LaunchTemplate` | no |

## ModifyLocalGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | no |
| `LocalGatewayRouteTableId` | `string` | yes |
| `LocalGatewayVirtualInterfaceGroupId` | `string` | no |
| `NetworkInterfaceId` | `string` | no |
| `DryRun` | `boolean` | no |
| `DestinationPrefixListId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `LocalGatewayRoute` | no |

## ModifyManagedPrefixList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListId` | `string` | yes |
| `CurrentVersion` | `long` | no |
| `PrefixListName` | `string` | no |
| `AddEntries` | `List<AddPrefixListEntry>` | no |
| `RemoveEntries` | `List<RemovePrefixListEntry>` | no |
| `MaxEntries` | `integer` | no |
| `IpamPrefixListResolverSyncEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixList` | `ManagedPrefixList` | no |

## ModifyManagedResourceVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `DefaultVisibility` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Visibility` | `ManagedResourceVisibilitySettings` | no |

## ModifyNetworkInterfaceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnaSrdSpecification` | `EnaSrdSpecification` | no |
| `EnablePrimaryIpv6` | `boolean` | no |
| `ConnectionTrackingSpecification` | `ConnectionTrackingSpecificationRequest` | no |
| `AssociatePublicIpAddress` | `boolean` | no |
| `AssociatedSubnetIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | yes |
| `Description` | `AttributeValue` | no |
| `SourceDestCheck` | `AttributeBooleanValue` | no |
| `Groups` | `List<string>` | no |
| `Attachment` | `NetworkInterfaceAttachmentChanges` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyPrivateDnsNameOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `PrivateDnsHostnameType` | `string` | no |
| `EnableResourceNameDnsARecord` | `boolean` | no |
| `EnableResourceNameDnsAAAARecord` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyPublicIpDnsNameOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | yes |
| `HostnameType` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `boolean` | no |

## ModifyReservedInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesIds` | `List<string>` | yes |
| `ClientToken` | `string` | no |
| `TargetConfigurations` | `List<ReservedInstancesConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesModificationId` | `string` | no |

## ModifyRouteServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServerId` | `string` | yes |
| `PersistRoutes` | `string` | no |
| `PersistRoutesDuration` | `long` | no |
| `SnsNotificationsEnabled` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouteServer` | `RouteServer` | no |

## ModifySecurityGroupRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `SecurityGroupRules` | `List<SecurityGroupRuleUpdate>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifySnapshotAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | no |
| `CreateVolumePermission` | `CreateVolumePermissionModifications` | no |
| `GroupNames` | `List<string>` | no |
| `OperationType` | `string` | no |
| `SnapshotId` | `string` | yes |
| `UserIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifySnapshotTier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `StorageTier` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `TieringStartTime` | `timestamp` | no |

## ModifySpotFleetRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchTemplateConfigs` | `List<LaunchTemplateConfig>` | no |
| `OnDemandTargetCapacity` | `integer` | no |
| `Context` | `string` | no |
| `SpotFleetRequestId` | `string` | yes |
| `TargetCapacity` | `integer` | no |
| `ExcessCapacityTerminationPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifySubnetAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssignIpv6AddressOnCreation` | `AttributeBooleanValue` | no |
| `MapPublicIpOnLaunch` | `AttributeBooleanValue` | no |
| `SubnetId` | `string` | yes |
| `MapCustomerOwnedIpOnLaunch` | `AttributeBooleanValue` | no |
| `CustomerOwnedIpv4Pool` | `string` | no |
| `EnableDns64` | `AttributeBooleanValue` | no |
| `PrivateDnsHostnameTypeOnLaunch` | `string` | no |
| `EnableResourceNameDnsARecordOnLaunch` | `AttributeBooleanValue` | no |
| `EnableResourceNameDnsAAAARecordOnLaunch` | `AttributeBooleanValue` | no |
| `EnableLniAtDeviceIndex` | `integer` | no |
| `DisableLniAtDeviceIndex` | `AttributeBooleanValue` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyTrafficMirrorFilterNetworkServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterId` | `string` | yes |
| `AddNetworkServices` | `List<string>` | no |
| `RemoveNetworkServices` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilter` | `TrafficMirrorFilter` | no |

## ModifyTrafficMirrorFilterRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRuleId` | `string` | yes |
| `TrafficDirection` | `string` | no |
| `RuleNumber` | `integer` | no |
| `RuleAction` | `string` | no |
| `DestinationPortRange` | `TrafficMirrorPortRangeRequest` | no |
| `SourcePortRange` | `TrafficMirrorPortRangeRequest` | no |
| `Protocol` | `integer` | no |
| `DestinationCidrBlock` | `string` | no |
| `SourceCidrBlock` | `string` | no |
| `Description` | `string` | no |
| `RemoveFields` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorFilterRule` | `TrafficMirrorFilterRule` | no |

## ModifyTrafficMirrorSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSessionId` | `string` | yes |
| `TrafficMirrorTargetId` | `string` | no |
| `TrafficMirrorFilterId` | `string` | no |
| `PacketLength` | `integer` | no |
| `SessionNumber` | `integer` | no |
| `VirtualNetworkId` | `integer` | no |
| `Description` | `string` | no |
| `RemoveFields` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficMirrorSession` | `TrafficMirrorSession` | no |

## ModifyTransitGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayId` | `string` | yes |
| `Description` | `string` | no |
| `Options` | `ModifyTransitGatewayOptions` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGateway` | `TransitGateway` | no |

## ModifyTransitGatewayMeteringPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicyId` | `string` | yes |
| `AddMiddleboxAttachmentIds` | `List<string>` | no |
| `RemoveMiddleboxAttachmentIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMeteringPolicy` | `TransitGatewayMeteringPolicy` | no |

## ModifyTransitGatewayPolicyTableEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableId` | `string` | yes |
| `PolicyRuleNumber` | `string` | yes |
| `PolicyRule` | `TransitGatewayRequestPolicyRule` | no |
| `TargetRouteTableId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPolicyTableEntry` | `TransitGatewayPolicyTableEntry` | no |

## ModifyTransitGatewayPrefixListReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `PrefixListId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `Blackhole` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPrefixListReference` | `TransitGatewayPrefixListReference` | no |

## ModifyTransitGatewayVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `AddSubnetIds` | `List<string>` | no |
| `RemoveSubnetIds` | `List<string>` | no |
| `Options` | `ModifyTransitGatewayVpcAttachmentRequestOptions` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachment` | `TransitGatewayVpcAttachment` | no |

## ModifyVerifiedAccessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointId` | `string` | yes |
| `VerifiedAccessGroupId` | `string` | no |
| `LoadBalancerOptions` | `ModifyVerifiedAccessEndpointLoadBalancerOptions` | no |
| `NetworkInterfaceOptions` | `ModifyVerifiedAccessEndpointEniOptions` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `RdsOptions` | `ModifyVerifiedAccessEndpointRdsOptions` | no |
| `CidrOptions` | `ModifyVerifiedAccessEndpointCidrOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpoint` | `VerifiedAccessEndpoint` | no |

## ModifyVerifiedAccessEndpointPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessEndpointId` | `string` | yes |
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationResponse` | no |

## ModifyVerifiedAccessGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupId` | `string` | yes |
| `VerifiedAccessInstanceId` | `string` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroup` | `VerifiedAccessGroup` | no |

## ModifyVerifiedAccessGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessGroupId` | `string` | yes |
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyEnabled` | `boolean` | no |
| `PolicyDocument` | `string` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationResponse` | no |

## ModifyVerifiedAccessInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `CidrEndpointsCustomSubDomain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstance` | `VerifiedAccessInstance` | no |

## ModifyVerifiedAccessInstanceLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessInstanceId` | `string` | yes |
| `AccessLogs` | `VerifiedAccessLogOptions` | yes |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `VerifiedAccessInstanceLoggingConfiguration` | no |

## ModifyVerifiedAccessTrustProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProviderId` | `string` | yes |
| `OidcOptions` | `ModifyVerifiedAccessTrustProviderOidcOptions` | no |
| `DeviceOptions` | `ModifyVerifiedAccessTrustProviderDeviceOptions` | no |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `SseSpecification` | `VerifiedAccessSseSpecificationRequest` | no |
| `NativeApplicationOidcOptions` | `ModifyVerifiedAccessNativeApplicationOidcOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedAccessTrustProvider` | `VerifiedAccessTrustProvider` | no |

## ModifyVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VolumeId` | `string` | yes |
| `Size` | `integer` | no |
| `VolumeType` | `string` | no |
| `Iops` | `integer` | no |
| `Throughput` | `integer` | no |
| `MultiAttachEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeModification` | `VolumeModification` | no |

## ModifyVolumeAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoEnableIO` | `AttributeBooleanValue` | no |
| `VolumeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyVpcAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableDnsHostnames` | `AttributeBooleanValue` | no |
| `EnableDnsSupport` | `AttributeBooleanValue` | no |
| `VpcId` | `string` | yes |
| `EnableNetworkAddressUsageMetrics` | `AttributeBooleanValue` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyVpcBlockPublicAccessExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ExclusionId` | `string` | yes |
| `InternetGatewayExclusionMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessExclusion` | `VpcBlockPublicAccessExclusion` | no |

## ModifyVpcBlockPublicAccessOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InternetGatewayBlockMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcBlockPublicAccessOptions` | `VpcBlockPublicAccessOptions` | no |

## ModifyVpcEncryptionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEncryptionControlId` | `string` | yes |
| `Mode` | `string` | no |
| `InternetGatewayExclusion` | `string` | no |
| `EgressOnlyInternetGatewayExclusion` | `string` | no |
| `NatGatewayExclusion` | `string` | no |
| `VirtualPrivateGatewayExclusion` | `string` | no |
| `VpcPeeringExclusion` | `string` | no |
| `LambdaExclusion` | `string` | no |
| `VpcLatticeExclusion` | `string` | no |
| `ElasticFileSystemExclusion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEncryptionControl` | `VpcEncryptionControl` | no |

## ModifyVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcEndpointId` | `string` | yes |
| `ResetPolicy` | `boolean` | no |
| `PolicyDocument` | `string` | no |
| `AddRouteTableIds` | `List<string>` | no |
| `RemoveRouteTableIds` | `List<string>` | no |
| `AddSubnetIds` | `List<string>` | no |
| `RemoveSubnetIds` | `List<string>` | no |
| `AddSecurityGroupIds` | `List<string>` | no |
| `RemoveSecurityGroupIds` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `DnsOptions` | `DnsOptionsSpecification` | no |
| `PrivateDnsEnabled` | `boolean` | no |
| `SubnetConfigurations` | `List<SubnetConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyVpcEndpointConnectionNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ConnectionNotificationId` | `string` | yes |
| `ConnectionNotificationArn` | `string` | no |
| `ConnectionEvents` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## ModifyVpcEndpointPayerResponsibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | no |
| `VpcEndpointId` | `string` | yes |
| `PayerResponsibility` | `string` | yes |
| `Scope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointId` | `string` | no |
| `PayerResponsibilities` | `List<PayerResponsibilityEntry>` | no |

## ModifyVpcEndpointServiceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `PrivateDnsName` | `string` | no |
| `RemovePrivateDnsName` | `boolean` | no |
| `AcceptanceRequired` | `boolean` | no |
| `AddNetworkLoadBalancerArns` | `List<string>` | no |
| `RemoveNetworkLoadBalancerArns` | `List<string>` | no |
| `AddGatewayLoadBalancerArns` | `List<string>` | no |
| `RemoveGatewayLoadBalancerArns` | `List<string>` | no |
| `AddSupportedIpAddressTypes` | `List<string>` | no |
| `RemoveSupportedIpAddressTypes` | `List<string>` | no |
| `AddSupportedRegions` | `List<string>` | no |
| `RemoveSupportedRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ModifyVpcEndpointServicePayerResponsibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `PayerResponsibility` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## ModifyVpcEndpointServicePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `AddAllowedPrincipals` | `List<string>` | no |
| `RemoveAllowedPrincipals` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddedPrincipals` | `List<AddedPrincipal>` | no |
| `ReturnValue` | `boolean` | no |

## ModifyVpcPeeringConnectionOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccepterPeeringConnectionOptions` | `PeeringConnectionOptionsRequest` | no |
| `DryRun` | `boolean` | no |
| `RequesterPeeringConnectionOptions` | `PeeringConnectionOptionsRequest` | no |
| `VpcPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccepterPeeringConnectionOptions` | `PeeringConnectionOptions` | no |
| `RequesterPeeringConnectionOptions` | `PeeringConnectionOptions` | no |

## ModifyVpcTenancy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcId` | `string` | yes |
| `InstanceTenancy` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## ModifyVpnConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `TransitGatewayId` | `string` | no |
| `CustomerGatewayId` | `string` | no |
| `VpnGatewayId` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnection` | `VpnConnection` | no |

## ModifyVpnConnectionOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `LocalIpv4NetworkCidr` | `string` | no |
| `RemoteIpv4NetworkCidr` | `string` | no |
| `LocalIpv6NetworkCidr` | `string` | no |
| `RemoteIpv6NetworkCidr` | `string` | no |
| `TunnelBandwidth` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnection` | `VpnConnection` | no |

## ModifyVpnTunnelCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnTunnelOutsideIpAddress` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnection` | `VpnConnection` | no |

## ModifyVpnTunnelOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnTunnelOutsideIpAddress` | `string` | yes |
| `TunnelOptions` | `ModifyVpnTunnelOptionsSpecification` | yes |
| `DryRun` | `boolean` | no |
| `SkipTunnelReplacement` | `boolean` | no |
| `PreSharedKeyStorage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnection` | `VpnConnection` | no |

## MonitorInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceMonitorings` | `List<InstanceMonitoring>` | no |

## MoveAddressToVpc

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PublicIp` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | no |
| `Status` | `string` | no |

## MoveByoipCidrToIpam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Cidr` | `string` | yes |
| `IpamPoolId` | `string` | yes |
| `IpamPoolOwner` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## MoveCapacityReservationInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |
| `SourceCapacityReservationId` | `string` | yes |
| `DestinationCapacityReservationId` | `string` | yes |
| `InstanceCount` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceCapacityReservation` | `CapacityReservation` | no |
| `DestinationCapacityReservation` | `CapacityReservation` | no |
| `InstanceCount` | `integer` | no |

## ProvisionByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `CidrAuthorizationContext` | `CidrAuthorizationContext` | no |
| `PubliclyAdvertisable` | `boolean` | no |
| `Description` | `string` | no |
| `DryRun` | `boolean` | no |
| `PoolTagSpecifications` | `List<TagSpecification>` | no |
| `MultiRegion` | `boolean` | no |
| `NetworkBorderGroup` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## ProvisionIpamByoasn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamId` | `string` | yes |
| `Asn` | `string` | yes |
| `AsnAuthorizationContext` | `AsnAuthorizationContext` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Byoasn` | `Byoasn` | no |

## ProvisionIpamPoolCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Cidr` | `string` | no |
| `CidrAuthorizationContext` | `IpamCidrAuthorizationContext` | no |
| `NetmaskLength` | `integer` | no |
| `ClientToken` | `string` | no |
| `VerificationMethod` | `string` | no |
| `IpamExternalResourceVerificationTokenId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPoolCidr` | `IpamPoolCidr` | no |

## ProvisionPublicIpv4PoolCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `PoolId` | `string` | yes |
| `NetmaskLength` | `integer` | yes |
| `NetworkBorderGroup` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | no |
| `PoolAddressRange` | `PublicIpv4PoolRange` | no |

## PurchaseCapacityBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `CapacityBlockOfferingId` | `string` | yes |
| `InstancePlatform` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservation` | `CapacityReservation` | no |
| `CapacityBlocks` | `List<CapacityBlock>` | no |

## PurchaseCapacityBlockExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockExtensionOfferingId` | `string` | yes |
| `CapacityReservationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityBlockExtensions` | `List<CapacityBlockExtension>` | no |

## PurchaseHostReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `CurrencyCode` | `string` | no |
| `HostIdSet` | `List<string>` | yes |
| `LimitPrice` | `string` | no |
| `OfferingId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `CurrencyCode` | `string` | no |
| `Purchase` | `List<Purchase>` | no |
| `TotalHourlyPrice` | `string` | no |
| `TotalUpfrontPrice` | `string` | no |

## PurchaseReservedInstancesOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceCount` | `integer` | yes |
| `ReservedInstancesOfferingId` | `string` | yes |
| `PurchaseTime` | `timestamp` | no |
| `DryRun` | `boolean` | no |
| `LimitPrice` | `ReservedInstanceLimitPrice` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstancesId` | `string` | no |

## PurchaseScheduledInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `PurchaseRequests` | `List<PurchaseRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledInstanceSet` | `List<ScheduledInstance>` | no |

## RebootInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageLocation` | `string` | no |
| `BillingProducts` | `List<string>` | no |
| `BootMode` | `string` | no |
| `TpmSupport` | `string` | no |
| `UefiData` | `string` | no |
| `ImdsSupport` | `string` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `DryRun` | `boolean` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Architecture` | `string` | no |
| `KernelId` | `string` | no |
| `RamdiskId` | `string` | no |
| `RootDeviceName` | `string` | no |
| `BlockDeviceMappings` | `List<BlockDeviceMapping>` | no |
| `VirtualizationType` | `string` | no |
| `SriovNetSupport` | `string` | no |
| `EnaSupport` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## RegisterInstanceEventNotificationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceTagAttribute` | `RegisterInstanceTagAttributeRequest` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTagAttribute` | `InstanceTagNotificationAttribute` | no |

## RegisterTransitGatewayMulticastGroupMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `GroupIpAddress` | `string` | no |
| `NetworkInterfaceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegisteredMulticastGroupMembers` | `TransitGatewayMulticastRegisteredGroupMembers` | no |

## RegisterTransitGatewayMulticastGroupSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `GroupIpAddress` | `string` | no |
| `NetworkInterfaceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegisteredMulticastGroupSources` | `TransitGatewayMulticastRegisteredGroupSources` | no |

## RejectCapacityReservationBillingOwnership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CapacityReservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## RejectTransitGatewayClientVpnAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayClientVpnAttachment` | `TransitGatewayClientVpnAttachment` | no |

## RejectTransitGatewayMulticastDomainAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | no |
| `TransitGatewayAttachmentId` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `TransitGatewayMulticastDomainAssociations` | no |

## RejectTransitGatewayPeeringAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayPeeringAttachment` | `TransitGatewayPeeringAttachment` | no |

## RejectTransitGatewayVpcAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayAttachmentId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayVpcAttachment` | `TransitGatewayVpcAttachment` | no |

## RejectVpcEndpointConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |
| `VpcEndpointIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## RejectVpcPeeringConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `VpcPeeringConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ReleaseAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | no |
| `PublicIp` | `string` | no |
| `NetworkBorderGroup` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReleaseHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<string>` | no |
| `Unsuccessful` | `List<UnsuccessfulItem>` | no |

## ReleaseIpamPoolAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `IpamPoolId` | `string` | yes |
| `Cidr` | `string` | yes |
| `IpamPoolAllocationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | no |

## ReplaceIamInstanceProfileAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfile` | `IamInstanceProfileSpecification` | yes |
| `AssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IamInstanceProfileAssociation` | `IamInstanceProfileAssociation` | no |

## ReplaceImageCriteriaInAllowedImagesSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageCriteria` | `List<ImageCriterionRequest>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## ReplaceImageInstanceTypeSpecification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `InstanceTypeSpecification` | `InstanceTypeSpecificationRequest` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## ReplaceNetworkAclAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AssociationId` | `string` | yes |
| `NetworkAclId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NewAssociationId` | `string` | no |

## ReplaceNetworkAclEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkAclId` | `string` | yes |
| `RuleNumber` | `integer` | yes |
| `Protocol` | `string` | yes |
| `RuleAction` | `string` | yes |
| `Egress` | `boolean` | yes |
| `CidrBlock` | `string` | no |
| `Ipv6CidrBlock` | `string` | no |
| `IcmpTypeCode` | `IcmpTypeCode` | no |
| `PortRange` | `PortRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReplaceRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPrefixListId` | `string` | no |
| `VpcEndpointId` | `string` | no |
| `LocalTarget` | `boolean` | no |
| `TransitGatewayId` | `string` | no |
| `LocalGatewayId` | `string` | no |
| `CarrierGatewayId` | `string` | no |
| `CoreNetworkArn` | `string` | no |
| `OdbNetworkArn` | `string` | no |
| `DryRun` | `boolean` | no |
| `RouteTableId` | `string` | yes |
| `DestinationCidrBlock` | `string` | no |
| `GatewayId` | `string` | no |
| `DestinationIpv6CidrBlock` | `string` | no |
| `EgressOnlyInternetGatewayId` | `string` | no |
| `InstanceId` | `string` | no |
| `NetworkInterfaceId` | `string` | no |
| `VpcPeeringConnectionId` | `string` | no |
| `NatGatewayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReplaceRouteTableAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `AssociationId` | `string` | yes |
| `RouteTableId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NewAssociationId` | `string` | no |
| `AssociationState` | `RouteTableAssociationState` | no |

## ReplaceTransitGatewayRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationCidrBlock` | `string` | yes |
| `TransitGatewayRouteTableId` | `string` | yes |
| `TransitGatewayAttachmentId` | `string` | no |
| `Blackhole` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Route` | `TransitGatewayRoute` | no |

## ReplaceVpnTunnel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpnConnectionId` | `string` | yes |
| `VpnTunnelOutsideIpAddress` | `string` | yes |
| `ApplyPendingMaintenance` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ReportInstanceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Instances` | `List<string>` | yes |
| `Status` | `string` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `ReasonCodes` | `List<string>` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RequestSpotFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `SpotFleetRequestConfig` | `SpotFleetRequestConfigData` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotFleetRequestId` | `string` | no |

## RequestSpotInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LaunchSpecification` | `RequestSpotLaunchSpecification` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `InstanceInterruptionBehavior` | `string` | no |
| `DryRun` | `boolean` | no |
| `SpotPrice` | `string` | no |
| `ClientToken` | `string` | no |
| `InstanceCount` | `integer` | no |
| `Type` | `string` | no |
| `ValidFrom` | `timestamp` | no |
| `ValidUntil` | `timestamp` | no |
| `LaunchGroup` | `string` | no |
| `AvailabilityZoneGroup` | `string` | no |
| `BlockDurationMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpotInstanceRequests` | `List<SpotInstanceRequest>` | no |

## ResetAddressAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllocationId` | `string` | yes |
| `Attribute` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Address` | `AddressAttribute` | no |

## ResetEbsDefaultKmsKeyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KmsKeyId` | `string` | no |

## ResetFpgaImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `FpgaImageId` | `string` | yes |
| `Attribute` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ResetImageAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetInstanceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `InstanceId` | `string` | yes |
| `Attribute` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetNetworkInterfaceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `NetworkInterfaceId` | `string` | yes |
| `SourceDestCheck` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetSnapshotAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `string` | yes |
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestoreAddressToClassic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PublicIp` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicIp` | `string` | no |
| `Status` | `string` | no |

## RestoreImageFromRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## RestoreManagedPrefixListVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `PrefixListId` | `string` | yes |
| `PreviousVersion` | `long` | yes |
| `CurrentVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrefixList` | `ManagedPrefixList` | no |

## RestoreSnapshotFromRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `OutpostArn` | `string` | no |
| `Description` | `string` | no |
| `Encrypted` | `boolean` | no |
| `OwnerId` | `string` | no |
| `Progress` | `string` | no |
| `StartTime` | `timestamp` | no |
| `State` | `string` | no |
| `VolumeId` | `string` | no |
| `VolumeSize` | `integer` | no |
| `SseType` | `string` | no |

## RestoreSnapshotTier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `TemporaryRestoreDays` | `integer` | no |
| `PermanentRestore` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |
| `RestoreStartTime` | `timestamp` | no |
| `RestoreDuration` | `integer` | no |
| `IsPermanentRestore` | `boolean` | no |

## RestoreVolumeFromRecycleBin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VolumeId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## RevokeClientVpnIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `TargetNetworkCidr` | `string` | yes |
| `AccessGroupId` | `string` | no |
| `RevokeAllGroups` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `ClientVpnAuthorizationRuleStatus` | no |

## RevokeSecurityGroupEgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupRuleIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `GroupId` | `string` | yes |
| `SourceSecurityGroupName` | `string` | no |
| `SourceSecurityGroupOwnerId` | `string` | no |
| `IpProtocol` | `string` | no |
| `FromPort` | `integer` | no |
| `ToPort` | `integer` | no |
| `CidrIp` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `UnknownIpPermissions` | `List<IpPermission>` | no |
| `RevokedSecurityGroupRules` | `List<RevokedSecurityGroupRule>` | no |

## RevokeSecurityGroupIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CidrIp` | `string` | no |
| `FromPort` | `integer` | no |
| `GroupId` | `string` | no |
| `GroupName` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |
| `IpProtocol` | `string` | no |
| `SourceSecurityGroupName` | `string` | no |
| `SourceSecurityGroupOwnerId` | `string` | no |
| `ToPort` | `integer` | no |
| `SecurityGroupRuleIds` | `List<string>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |
| `UnknownIpPermissions` | `List<IpPermission>` | no |
| `RevokedSecurityGroupRules` | `List<RevokedSecurityGroupRule>` | no |

## RunInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlockDeviceMappings` | `List<BlockDeviceMapping>` | no |
| `ImageId` | `string` | no |
| `InstanceType` | `string` | no |
| `Ipv6AddressCount` | `integer` | no |
| `Ipv6Addresses` | `List<InstanceIpv6Address>` | no |
| `KernelId` | `string` | no |
| `KeyName` | `string` | no |
| `MaxCount` | `integer` | yes |
| `MinCount` | `integer` | yes |
| `Monitoring` | `RunInstancesMonitoringEnabled` | no |
| `Placement` | `Placement` | no |
| `RamdiskId` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `SubnetId` | `string` | no |
| `UserData` | `string` | no |
| `ElasticGpuSpecification` | `List<ElasticGpuSpecification>` | no |
| `ElasticInferenceAccelerators` | `List<ElasticInferenceAccelerator>` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `LaunchTemplate` | `LaunchTemplateSpecification` | no |
| `InstanceMarketOptions` | `InstanceMarketOptionsRequest` | no |
| `CreditSpecification` | `CreditSpecificationRequest` | no |
| `CpuOptions` | `CpuOptionsRequest` | no |
| `CapacityReservationSpecification` | `CapacityReservationSpecification` | no |
| `HibernationOptions` | `HibernationOptionsRequest` | no |
| `LicenseSpecifications` | `List<LicenseConfigurationRequest>` | no |
| `MetadataOptions` | `InstanceMetadataOptionsRequest` | no |
| `EnclaveOptions` | `EnclaveOptionsRequest` | no |
| `PrivateDnsNameOptions` | `PrivateDnsNameOptionsRequest` | no |
| `MaintenanceOptions` | `InstanceMaintenanceOptionsRequest` | no |
| `DisableApiStop` | `boolean` | no |
| `EnablePrimaryIpv6` | `boolean` | no |
| `NetworkPerformanceOptions` | `InstanceNetworkPerformanceOptionsRequest` | no |
| `Operator` | `OperatorRequest` | no |
| `SecondaryInterfaces` | `List<InstanceSecondaryInterfaceSpecificationRequest>` | no |
| `DryRun` | `boolean` | no |
| `DisableApiTermination` | `boolean` | no |
| `InstanceInitiatedShutdownBehavior` | `string` | no |
| `PrivateIpAddress` | `string` | no |
| `ClientToken` | `string` | no |
| `AdditionalInfo` | `string` | no |
| `NetworkInterfaces` | `List<InstanceNetworkInterfaceSpecification>` | no |
| `IamInstanceProfile` | `IamInstanceProfileSpecification` | no |
| `EbsOptimized` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservationId` | `string` | no |
| `OwnerId` | `string` | no |
| `RequesterId` | `string` | no |
| `Groups` | `List<GroupIdentifier>` | no |
| `Instances` | `List<Instance>` | no |

## RunScheduledInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DryRun` | `boolean` | no |
| `InstanceCount` | `integer` | no |
| `LaunchSpecification` | `ScheduledInstancesLaunchSpecification` | yes |
| `ScheduledInstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIdSet` | `List<string>` | no |

## SearchLocalGatewayRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Routes` | `List<LocalGatewayRoute>` | no |
| `NextToken` | `string` | no |

## SearchTransitGatewayMulticastGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayMulticastDomainId` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MulticastGroups` | `List<TransitGatewayMulticastGroup>` | no |
| `NextToken` | `string` | no |

## SearchTransitGatewayRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransitGatewayRouteTableId` | `string` | yes |
| `Filters` | `List<Filter>` | yes |
| `MaxResults` | `integer` | no |
| `DryRun` | `boolean` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Routes` | `List<TransitGatewayRoute>` | no |
| `AdditionalRoutesAvailable` | `boolean` | no |
| `NextToken` | `string` | no |

## SendDiagnosticInterrupt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartDeclarativePoliciesReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `S3Bucket` | `string` | yes |
| `S3Prefix` | `string` | no |
| `TargetId` | `string` | yes |
| `TagSpecifications` | `List<TagSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | no |

## StartInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `AdditionalInfo` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartingInstances` | `List<InstanceStateChange>` | no |

## StartNetworkInsightsAccessScopeAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeId` | `string` | yes |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAccessScopeAnalysis` | `NetworkInsightsAccessScopeAnalysis` | no |

## StartNetworkInsightsAnalysis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsPathId` | `string` | yes |
| `AdditionalAccounts` | `List<string>` | no |
| `FilterInArns` | `List<string>` | no |
| `FilterOutArns` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `TagSpecifications` | `List<TagSpecification>` | no |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInsightsAnalysis` | `NetworkInsightsAnalysis` | no |

## StartVpcEndpointServicePrivateDnsVerification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `ServiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReturnValue` | `boolean` | no |

## StopInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `Hibernate` | `boolean` | no |
| `SkipOsShutdown` | `boolean` | no |
| `DryRun` | `boolean` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StoppingInstances` | `List<InstanceStateChange>` | no |

## TerminateClientVpnConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | yes |
| `ConnectionId` | `string` | no |
| `Username` | `string` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientVpnEndpointId` | `string` | no |
| `Username` | `string` | no |
| `ConnectionStatuses` | `List<TerminateConnectionStatus>` | no |

## TerminateInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `Force` | `boolean` | no |
| `SkipOsShutdown` | `boolean` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminatingInstances` | `List<InstanceStateChange>` | no |

## UnassignIpv6Addresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv6Prefixes` | `List<string>` | no |
| `NetworkInterfaceId` | `string` | yes |
| `Ipv6Addresses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkInterfaceId` | `string` | no |
| `UnassignedIpv6Addresses` | `List<string>` | no |
| `UnassignedIpv6Prefixes` | `List<string>` | no |

## UnassignPrivateIpAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ipv4Prefixes` | `List<string>` | no |
| `NetworkInterfaceId` | `string` | yes |
| `PrivateIpAddresses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnassignPrivateNatGatewayAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | yes |
| `PrivateIpAddresses` | `List<string>` | yes |
| `MaxDrainDurationSeconds` | `integer` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NatGatewayId` | `string` | no |
| `NatGatewayAddresses` | `List<NatGatewayAddress>` | no |

## UnlockSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | no |

## UnmonitorInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceMonitorings` | `List<InstanceMonitoring>` | no |

## UpdateCapacityManagerMonitoredTagKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivateTagKeys` | `List<string>` | no |
| `DeactivateTagKeys` | `List<string>` | no |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerTagKeys` | `List<CapacityManagerMonitoredTagKey>` | no |

## UpdateCapacityManagerOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrganizationsAccess` | `boolean` | yes |
| `DryRun` | `boolean` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityManagerStatus` | `string` | no |
| `OrganizationsAccess` | `boolean` | no |

## UpdateInterruptibleCapacityReservationAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityReservationId` | `string` | yes |
| `TargetInstanceCount` | `integer` | no |
| `DryRun` | `boolean` | no |
| `ZeroSizePreference` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InterruptibleCapacityReservationId` | `string` | no |
| `SourceCapacityReservationId` | `string` | no |
| `InstanceCount` | `integer` | no |
| `TargetInstanceCount` | `integer` | no |
| `Status` | `string` | no |
| `InterruptionType` | `string` | no |

## UpdateSecurityGroupRuleDescriptionsEgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `GroupId` | `string` | no |
| `GroupName` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |
| `SecurityGroupRuleDescriptions` | `List<SecurityGroupRuleDescription>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## UpdateSecurityGroupRuleDescriptionsIngress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `GroupId` | `string` | no |
| `GroupName` | `string` | no |
| `IpPermissions` | `List<IpPermission>` | no |
| `SecurityGroupRuleDescriptions` | `List<SecurityGroupRuleDescription>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Return` | `boolean` | no |

## ValidateSecurityGroupQuotasForInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupIds` | `List<string>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Valid` | `boolean` | no |

## WithdrawByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

