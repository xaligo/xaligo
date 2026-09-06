# AWS MediaConnect

API version: 2018-11-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediaconnect/2018-11-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddBridgeOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `Outputs` | `List<AddBridgeOutputRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `Outputs` | `List<BridgeOutput>` | no |

## AddBridgeSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `Sources` | `List<AddBridgeSourceRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `Sources` | `List<BridgeSource>` | no |

## AddFlowMediaStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `MediaStreams` | `List<AddMediaStreamRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `MediaStreams` | `List<MediaStream>` | no |

## AddFlowOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `Outputs` | `List<AddOutputRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Outputs` | `List<Output>` | no |

## AddFlowSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `Sources` | `List<SetSourceRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Sources` | `List<Source>` | no |

## AddFlowVpcInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `VpcInterfaces` | `List<VpcInterfaceRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `VpcInterfaces` | `List<VpcInterface>` | no |

## BatchGetRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterInputs` | `List<RouterInput>` | yes |
| `Errors` | `List<BatchGetRouterInputError>` | yes |

## BatchGetRouterNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterNetworkInterfaces` | `List<RouterNetworkInterface>` | yes |
| `Errors` | `List<BatchGetRouterNetworkInterfaceError>` | yes |

## BatchGetRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutputs` | `List<RouterOutput>` | yes |
| `Errors` | `List<BatchGetRouterOutputError>` | yes |

## CreateBridge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EgressGatewayBridge` | `AddEgressGatewayBridgeRequest` | no |
| `IngressGatewayBridge` | `AddIngressGatewayBridgeRequest` | no |
| `Name` | `string` | yes |
| `Outputs` | `List<AddBridgeOutputRequest>` | no |
| `PlacementArn` | `string` | yes |
| `SourceFailoverConfig` | `FailoverConfig` | no |
| `Sources` | `List<AddBridgeSourceRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bridge` | `Bridge` | no |

## CreateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZone` | `string` | no |
| `Entitlements` | `List<GrantEntitlementRequest>` | no |
| `MediaStreams` | `List<AddMediaStreamRequest>` | no |
| `Name` | `string` | yes |
| `Outputs` | `List<AddOutputRequest>` | no |
| `Source` | `SetSourceRequest` | no |
| `SourceFailoverConfig` | `FailoverConfig` | no |
| `Sources` | `List<SetSourceRequest>` | no |
| `VpcInterfaces` | `List<VpcInterfaceRequest>` | no |
| `Maintenance` | `AddMaintenance` | no |
| `SourceMonitoringConfig` | `MonitoringConfig` | no |
| `FlowSize` | `string` | no |
| `NdiConfig` | `NdiConfig` | no |
| `EncodingConfig` | `EncodingConfig` | no |
| `FlowTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Flow` | `Flow` | no |

## CreateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EgressCidrBlocks` | `List<string>` | yes |
| `Name` | `string` | yes |
| `Networks` | `List<GatewayNetwork>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateway` | `Gateway` | no |

## CreateRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Configuration` | `RouterInputConfiguration` | yes |
| `MaximumBitrate` | `long` | yes |
| `RoutingScope` | `string` | yes |
| `Tier` | `string` | yes |
| `RegionName` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `TransitEncryption` | `RouterInputTransitEncryption` | no |
| `MaintenanceConfiguration` | `MaintenanceConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `ContentQualityAnalysisConfiguration` | `RouterContentQualityAnalysisConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterInput` | `RouterInput` | yes |

## CreateRouterNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Configuration` | `RouterNetworkInterfaceConfiguration` | yes |
| `RegionName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterNetworkInterface` | `RouterNetworkInterface` | yes |

## CreateRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Configuration` | `RouterOutputConfiguration` | yes |
| `MaximumBitrate` | `long` | yes |
| `RoutingScope` | `string` | yes |
| `Tier` | `string` | yes |
| `RegionName` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `MaintenanceConfiguration` | `MaintenanceConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `FabricConfiguration` | `FabricConfiguration` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutput` | `RouterOutput` | yes |

## DeleteBridge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |

## DeleteFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Status` | `string` | no |

## DeleteGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | no |

## DeleteRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## DeleteRouterNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## DeleteRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## DeregisterGatewayInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Force` | `boolean` | no |
| `GatewayInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayInstanceArn` | `string` | no |
| `InstanceState` | `string` | no |

## DescribeBridge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bridge` | `Bridge` | no |

## DescribeFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Flow` | `Flow` | no |
| `Messages` | `Messages` | no |

## DescribeFlowSourceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Messages` | `List<MessageDetail>` | no |
| `Timestamp` | `timestamp` | no |
| `TransportMediaInfo` | `TransportMediaInfo` | no |
| `NdiInfo` | `NdiSourceMetadataInfo` | no |

## DescribeFlowSourceThumbnail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThumbnailDetails` | `ThumbnailDetails` | no |

## DescribeGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateway` | `Gateway` | no |

## DescribeGatewayInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GatewayInstance` | `GatewayInstance` | no |

## DescribeOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OfferingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Offering` | `Offering` | no |

## DescribeReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Reservation` | `Reservation` | no |

## GetRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterInput` | `RouterInput` | yes |

## GetRouterInputSourceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `SourceMetadataDetails` | `RouterInputSourceMetadataDetails` | yes |

## GetRouterInputThumbnail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `ThumbnailDetails` | `RouterInputThumbnailDetails` | yes |

## GetRouterNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterNetworkInterface` | `RouterNetworkInterface` | yes |

## GetRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutput` | `RouterOutput` | yes |

## GrantFlowEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlements` | `List<GrantEntitlementRequest>` | yes |
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlements` | `List<Entitlement>` | no |
| `FlowArn` | `string` | no |

## ListBridges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bridges` | `List<ListedBridge>` | no |
| `NextToken` | `string` | no |

## ListEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlements` | `List<ListedEntitlement>` | no |
| `NextToken` | `string` | no |

## ListFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Flows` | `List<ListedFlow>` | no |
| `NextToken` | `string` | no |

## ListGatewayInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<ListedGatewayInstance>` | no |
| `NextToken` | `string` | no |

## ListGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Gateways` | `List<ListedGateway>` | no |
| `NextToken` | `string` | no |

## ListOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Offerings` | `List<Offering>` | no |

## ListReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Reservations` | `List<Reservation>` | no |

## ListRouterInputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<RouterInputFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterInputs` | `List<ListedRouterInput>` | yes |
| `NextToken` | `string` | no |

## ListRouterNetworkInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<RouterNetworkInterfaceFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterNetworkInterfaces` | `List<ListedRouterNetworkInterface>` | yes |
| `NextToken` | `string` | no |

## ListRouterOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<RouterOutputFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutputs` | `List<ListedRouterOutput>` | yes |
| `NextToken` | `string` | no |

## ListTagsForGlobalResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PurchaseOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OfferingArn` | `string` | yes |
| `ReservationName` | `string` | yes |
| `Start` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Reservation` | `Reservation` | no |

## RemoveBridgeOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `OutputName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `OutputName` | `string` | no |

## RemoveBridgeSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `SourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `SourceName` | `string` | no |

## RemoveFlowMediaStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `MediaStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `MediaStreamName` | `string` | no |

## RemoveFlowOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `OutputArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `OutputArn` | `string` | no |

## RemoveFlowSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `SourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `SourceArn` | `string` | no |

## RemoveFlowVpcInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `VpcInterfaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `NonDeletedNetworkInterfaceIds` | `List<string>` | no |
| `VpcInterfaceName` | `string` | no |

## RestartRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## RestartRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## RevokeFlowEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitlementArn` | `string` | yes |
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitlementArn` | `string` | no |
| `FlowArn` | `string` | no |

## StartFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Status` | `string` | no |

## StartRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |
| `MaintenanceScheduleType` | `string` | yes |
| `MaintenanceSchedule` | `MaintenanceSchedule` | yes |

## StartRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |
| `MaintenanceScheduleType` | `string` | yes |
| `MaintenanceSchedule` | `MaintenanceSchedule` | yes |

## StopFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Status` | `string` | no |

## StopRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## StopRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `State` | `string` | yes |

## TagGlobalResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TakeRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutputArn` | `string` | yes |
| `RouterInputArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutedState` | `string` | yes |
| `RouterOutputArn` | `string` | yes |
| `RouterOutputName` | `string` | yes |
| `RouterInputArn` | `string` | no |
| `RouterInputName` | `string` | no |

## UntagGlobalResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

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


## UpdateBridge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `EgressGatewayBridge` | `UpdateEgressGatewayBridgeRequest` | no |
| `IngressGatewayBridge` | `UpdateIngressGatewayBridgeRequest` | no |
| `SourceFailoverConfig` | `UpdateFailoverConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bridge` | `Bridge` | no |

## UpdateBridgeOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `NetworkOutput` | `UpdateBridgeNetworkOutputRequest` | no |
| `OutputName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `Output` | `BridgeOutput` | no |

## UpdateBridgeSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `FlowSource` | `UpdateBridgeFlowSourceRequest` | no |
| `NetworkSource` | `UpdateBridgeNetworkSourceRequest` | no |
| `SourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `Source` | `BridgeSource` | no |

## UpdateBridgeState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | yes |
| `DesiredState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgeArn` | `string` | no |
| `DesiredState` | `string` | no |

## UpdateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | yes |
| `SourceFailoverConfig` | `UpdateFailoverConfig` | no |
| `Maintenance` | `UpdateMaintenance` | no |
| `SourceMonitoringConfig` | `MonitoringConfig` | no |
| `NdiConfig` | `NdiConfig` | no |
| `FlowSize` | `string` | no |
| `EncodingConfig` | `EncodingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Flow` | `Flow` | no |

## UpdateFlowEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Encryption` | `UpdateEncryption` | no |
| `EntitlementArn` | `string` | yes |
| `EntitlementStatus` | `string` | no |
| `FlowArn` | `string` | yes |
| `Subscribers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlement` | `Entitlement` | no |
| `FlowArn` | `string` | no |

## UpdateFlowMediaStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `MediaStreamAttributesRequest` | no |
| `ClockRate` | `integer` | no |
| `Description` | `string` | no |
| `FlowArn` | `string` | yes |
| `MediaStreamName` | `string` | yes |
| `MediaStreamType` | `string` | no |
| `VideoFormat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `MediaStream` | `MediaStream` | no |

## UpdateFlowOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CidrAllowList` | `List<string>` | no |
| `Description` | `string` | no |
| `Destination` | `string` | no |
| `Encryption` | `UpdateEncryption` | no |
| `FlowArn` | `string` | yes |
| `MaxLatency` | `integer` | no |
| `MediaStreamOutputConfigurations` | `List<MediaStreamOutputConfigurationRequest>` | no |
| `MinLatency` | `integer` | no |
| `OutputArn` | `string` | yes |
| `Port` | `integer` | no |
| `Protocol` | `string` | no |
| `RemoteId` | `string` | no |
| `SenderControlPort` | `integer` | no |
| `SenderIpAddress` | `string` | no |
| `SmoothingLatency` | `integer` | no |
| `StreamId` | `string` | no |
| `VpcInterfaceAttachment` | `VpcInterfaceAttachment` | no |
| `OutputStatus` | `string` | no |
| `NdiProgramName` | `string` | no |
| `NdiSpeedHqQuality` | `integer` | no |
| `RouterIntegrationState` | `string` | no |
| `RouterIntegrationTransitEncryption` | `FlowTransitEncryption` | no |
| `NdiOutputTimecodeSource` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Output` | `Output` | no |

## UpdateFlowSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Decryption` | `UpdateEncryption` | no |
| `Description` | `string` | no |
| `EntitlementArn` | `string` | no |
| `FlowArn` | `string` | yes |
| `IngestPort` | `integer` | no |
| `MaxBitrate` | `integer` | no |
| `MaxLatency` | `integer` | no |
| `MaxSyncBuffer` | `integer` | no |
| `MediaStreamSourceConfigurations` | `List<MediaStreamSourceConfigurationRequest>` | no |
| `MinLatency` | `integer` | no |
| `Protocol` | `string` | no |
| `SenderControlPort` | `integer` | no |
| `SenderIpAddress` | `string` | no |
| `SourceArn` | `string` | yes |
| `SourceListenerAddress` | `string` | no |
| `SourceListenerPort` | `integer` | no |
| `StreamId` | `string` | no |
| `VpcInterfaceName` | `string` | no |
| `WhitelistCidr` | `string` | no |
| `GatewayBridgeSource` | `UpdateGatewayBridgeSourceRequest` | no |
| `NdiSourceSettings` | `NdiSourceSettings` | no |
| `RouterIntegrationState` | `string` | no |
| `RouterIntegrationTransitDecryption` | `FlowTransitEncryption` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowArn` | `string` | no |
| `Source` | `Source` | no |

## UpdateGatewayInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgePlacement` | `string` | no |
| `GatewayInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BridgePlacement` | `string` | no |
| `GatewayInstanceArn` | `string` | no |

## UpdateRouterInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Configuration` | `RouterInputConfiguration` | no |
| `MaximumBitrate` | `long` | no |
| `RoutingScope` | `string` | no |
| `Tier` | `string` | no |
| `TransitEncryption` | `RouterInputTransitEncryption` | no |
| `MaintenanceConfiguration` | `MaintenanceConfiguration` | no |
| `ContentQualityAnalysisConfiguration` | `RouterContentQualityAnalysisConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterInput` | `RouterInput` | yes |

## UpdateRouterNetworkInterface

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Configuration` | `RouterNetworkInterfaceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterNetworkInterface` | `RouterNetworkInterface` | yes |

## UpdateRouterOutput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Configuration` | `RouterOutputConfiguration` | no |
| `MaximumBitrate` | `long` | no |
| `RoutingScope` | `string` | no |
| `Tier` | `string` | no |
| `MaintenanceConfiguration` | `MaintenanceConfiguration` | no |
| `FabricConfiguration` | `FabricConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RouterOutput` | `RouterOutput` | yes |

