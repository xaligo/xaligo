# AWS Global Accelerator

API version: 2018-08-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/globalaccelerator/2018-08-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddCustomRoutingEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigurations` | `List<CustomRoutingEndpointConfiguration>` | yes |
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointDescriptions` | `List<CustomRoutingEndpointDescription>` | no |
| `EndpointGroupArn` | `string` | no |

## AddEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigurations` | `List<EndpointConfiguration>` | yes |
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointDescriptions` | `List<EndpointDescription>` | no |
| `EndpointGroupArn` | `string` | no |

## AdvertiseByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## AllowCustomRoutingTraffic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |
| `EndpointId` | `string` | yes |
| `DestinationAddresses` | `List<string>` | no |
| `DestinationPorts` | `List<integer>` | no |
| `AllowAllTrafficToEndpoint` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IpAddressType` | `string` | no |
| `IpAddresses` | `List<string>` | no |
| `Enabled` | `boolean` | no |
| `IdempotencyToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `Accelerator` | no |

## CreateCrossAccountAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Principals` | `List<string>` | no |
| `Resources` | `List<Resource>` | no |
| `IdempotencyToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAttachment` | `Attachment` | no |

## CreateCustomRoutingAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IpAddressType` | `string` | no |
| `IpAddresses` | `List<string>` | no |
| `Enabled` | `boolean` | no |
| `IdempotencyToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `CustomRoutingAccelerator` | no |

## CreateCustomRoutingEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `EndpointGroupRegion` | `string` | yes |
| `DestinationConfigurations` | `List<CustomRoutingDestinationConfiguration>` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroup` | `CustomRoutingEndpointGroup` | no |

## CreateCustomRoutingListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `PortRanges` | `List<PortRange>` | yes |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `CustomRoutingListener` | no |

## CreateEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `EndpointGroupRegion` | `string` | yes |
| `EndpointConfigurations` | `List<EndpointConfiguration>` | no |
| `TrafficDialPercentage` | `float` | no |
| `HealthCheckPort` | `integer` | no |
| `HealthCheckProtocol` | `string` | no |
| `HealthCheckPath` | `string` | no |
| `HealthCheckIntervalSeconds` | `integer` | no |
| `ThresholdCount` | `integer` | no |
| `IdempotencyToken` | `string` | yes |
| `PortOverrides` | `List<PortOverride>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroup` | `EndpointGroup` | no |

## CreateListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `PortRanges` | `List<PortRange>` | yes |
| `Protocol` | `string` | yes |
| `ClientAffinity` | `string` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `Listener` | no |

## DeleteAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCrossAccountAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomRoutingAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomRoutingEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomRoutingListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DenyCustomRoutingTraffic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |
| `EndpointId` | `string` | yes |
| `DestinationAddresses` | `List<string>` | no |
| `DestinationPorts` | `List<integer>` | no |
| `DenyAllTrafficToEndpoint` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeprovisionByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## DescribeAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `Accelerator` | no |

## DescribeAcceleratorAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorAttributes` | `AcceleratorAttributes` | no |

## DescribeCrossAccountAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAttachment` | `Attachment` | no |

## DescribeCustomRoutingAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `CustomRoutingAccelerator` | no |

## DescribeCustomRoutingAcceleratorAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorAttributes` | `CustomRoutingAcceleratorAttributes` | no |

## DescribeCustomRoutingEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroup` | `CustomRoutingEndpointGroup` | no |

## DescribeCustomRoutingListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `CustomRoutingListener` | no |

## DescribeEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroup` | `EndpointGroup` | no |

## DescribeListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `Listener` | no |

## ListAccelerators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerators` | `List<Accelerator>` | no |
| `NextToken` | `string` | no |

## ListByoipCidrs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidrs` | `List<ByoipCidr>` | no |
| `NextToken` | `string` | no |

## ListCrossAccountAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAttachments` | `List<Attachment>` | no |
| `NextToken` | `string` | no |

## ListCrossAccountResourceAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceOwnerAwsAccountIds` | `List<string>` | no |

## ListCrossAccountResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | no |
| `ResourceOwnerAwsAccountId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountResources` | `List<CrossAccountResource>` | no |
| `NextToken` | `string` | no |

## ListCustomRoutingAccelerators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerators` | `List<CustomRoutingAccelerator>` | no |
| `NextToken` | `string` | no |

## ListCustomRoutingEndpointGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroups` | `List<CustomRoutingEndpointGroup>` | no |
| `NextToken` | `string` | no |

## ListCustomRoutingListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listeners` | `List<CustomRoutingListener>` | no |
| `NextToken` | `string` | no |

## ListCustomRoutingPortMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `EndpointGroupArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortMappings` | `List<PortMapping>` | no |
| `NextToken` | `string` | no |

## ListCustomRoutingPortMappingsByDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointId` | `string` | yes |
| `DestinationAddress` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPortMappings` | `List<DestinationPortMapping>` | no |
| `NextToken` | `string` | no |

## ListEndpointGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroups` | `List<EndpointGroup>` | no |
| `NextToken` | `string` | no |

## ListListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listeners` | `List<Listener>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ProvisionByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |
| `CidrAuthorizationContext` | `CidrAuthorizationContext` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

## RemoveCustomRoutingEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointIds` | `List<string>` | yes |
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointIdentifiers` | `List<EndpointIdentifier>` | yes |
| `EndpointGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `Name` | `string` | no |
| `IpAddressType` | `string` | no |
| `IpAddresses` | `List<string>` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `Accelerator` | no |

## UpdateAcceleratorAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `FlowLogsEnabled` | `boolean` | no |
| `FlowLogsS3Bucket` | `string` | no |
| `FlowLogsS3Prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorAttributes` | `AcceleratorAttributes` | no |

## UpdateCrossAccountAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentArn` | `string` | yes |
| `Name` | `string` | no |
| `AddPrincipals` | `List<string>` | no |
| `RemovePrincipals` | `List<string>` | no |
| `AddResources` | `List<Resource>` | no |
| `RemoveResources` | `List<Resource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossAccountAttachment` | `Attachment` | no |

## UpdateCustomRoutingAccelerator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `Name` | `string` | no |
| `IpAddressType` | `string` | no |
| `IpAddresses` | `List<string>` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Accelerator` | `CustomRoutingAccelerator` | no |

## UpdateCustomRoutingAcceleratorAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorArn` | `string` | yes |
| `FlowLogsEnabled` | `boolean` | no |
| `FlowLogsS3Bucket` | `string` | no |
| `FlowLogsS3Prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceleratorAttributes` | `CustomRoutingAcceleratorAttributes` | no |

## UpdateCustomRoutingListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `PortRanges` | `List<PortRange>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `CustomRoutingListener` | no |

## UpdateEndpointGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroupArn` | `string` | yes |
| `EndpointConfigurations` | `List<EndpointConfiguration>` | no |
| `TrafficDialPercentage` | `float` | no |
| `HealthCheckPort` | `integer` | no |
| `HealthCheckProtocol` | `string` | no |
| `HealthCheckPath` | `string` | no |
| `HealthCheckIntervalSeconds` | `integer` | no |
| `ThresholdCount` | `integer` | no |
| `PortOverrides` | `List<PortOverride>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointGroup` | `EndpointGroup` | no |

## UpdateListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `PortRanges` | `List<PortRange>` | no |
| `Protocol` | `string` | no |
| `ClientAffinity` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listener` | `Listener` | no |

## WithdrawByoipCidr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Cidr` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ByoipCidr` | `ByoipCidr` | no |

