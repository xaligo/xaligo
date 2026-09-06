# AWS IoT Wireless

API version: 2020-11-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotwireless/2020-11-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAwsAccountWithPartnerAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sidewalk` | `SidewalkAccountInfo` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sidewalk` | `SidewalkAccountInfo` | no |
| `Arn` | `string` | no |

## AssociateMulticastGroupWithFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `MulticastGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWirelessDeviceWithFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WirelessDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWirelessDeviceWithMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WirelessDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWirelessDeviceWithThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `ThingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWirelessGatewayWithCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IotCertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IotCertificateId` | `string` | no |

## AssociateWirelessGatewayWithThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `ThingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelMulticastGroupSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ExpressionType` | `string` | yes |
| `Expression` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |

## CreateDeviceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `LoRaWAN` | `LoRaWANDeviceProfile` | no |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |
| `Sidewalk` | `SidewalkCreateDeviceProfile` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `LoRaWAN` | `LoRaWANFuotaTask` | no |
| `FirmwareUpdateImage` | `string` | yes |
| `FirmwareUpdateRole` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `RedundancyPercent` | `integer` | no |
| `FragmentSizeBytes` | `integer` | no |
| `FragmentIntervalMS` | `integer` | no |
| `Descriptor` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `LoRaWAN` | `LoRaWANMulticast` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateNetworkAnalyzerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TraceContent` | `TraceContent` | no |
| `WirelessDevices` | `List<string>` | no |
| `WirelessGateways` | `List<string>` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |
| `MulticastGroups` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |

## CreateServiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `LoRaWAN` | `LoRaWANServiceProfile` | no |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DestinationName` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `LoRaWAN` | `LoRaWANDevice` | no |
| `Tags` | `List<Tag>` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkCreateWirelessDevice` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateWirelessGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANGateway` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateWirelessGatewayTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WirelessGatewayTaskDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessGatewayTaskDefinitionId` | `string` | no |
| `Status` | `string` | no |

## CreateWirelessGatewayTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoCreateTasks` | `boolean` | yes |
| `Name` | `string` | no |
| `Update` | `UpdateWirelessGatewayTaskCreate` | no |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## DeleteDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeviceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkAnalyzerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueuedMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `MessageId` | `string` | yes |
| `WirelessDeviceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWirelessGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWirelessGatewayTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWirelessGatewayTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `WirelessDeviceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAwsAccountFromPartnerAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerAccountId` | `string` | yes |
| `PartnerType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMulticastGroupFromFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `MulticastGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWirelessDeviceFromFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WirelessDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWirelessDeviceFromMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WirelessDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWirelessDeviceFromThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWirelessGatewayFromCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWirelessGatewayFromThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Expression` | `string` | no |
| `ExpressionType` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |

## GetDeviceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Id` | `string` | no |
| `LoRaWAN` | `LoRaWANDeviceProfile` | no |
| `Sidewalk` | `SidewalkGetDeviceProfile` | no |

## GetEventConfigurationByResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceRegistrationState` | `DeviceRegistrationStateResourceTypeEventConfiguration` | no |
| `Proximity` | `ProximityResourceTypeEventConfiguration` | no |
| `Join` | `JoinResourceTypeEventConfiguration` | no |
| `ConnectionStatus` | `ConnectionStatusResourceTypeEventConfiguration` | no |
| `MessageDeliveryStatus` | `MessageDeliveryStatusResourceTypeEventConfiguration` | no |

## GetFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANFuotaTaskGetInfo` | no |
| `FirmwareUpdateImage` | `string` | no |
| `FirmwareUpdateRole` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `RedundancyPercent` | `integer` | no |
| `FragmentSizeBytes` | `integer` | no |
| `FragmentIntervalMS` | `integer` | no |
| `Descriptor` | `string` | no |

## GetLogLevelsByResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultLogLevel` | `string` | no |
| `WirelessGatewayLogOptions` | `List<WirelessGatewayLogOption>` | no |
| `WirelessDeviceLogOptions` | `List<WirelessDeviceLogOption>` | no |
| `FuotaTaskLogOptions` | `List<FuotaTaskLogOption>` | no |

## GetMetricConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryMetric` | `SummaryMetricConfiguration` | no |

## GetMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryMetricQueries` | `List<SummaryMetricQuery>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryMetricQueryResults` | `List<SummaryMetricQueryResult>` | no |

## GetMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `LoRaWAN` | `LoRaWANMulticastGet` | no |
| `CreatedAt` | `timestamp` | no |

## GetMulticastGroupSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoRaWAN` | `LoRaWANMulticastSession` | no |

## GetNetworkAnalyzerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceContent` | `TraceContent` | no |
| `WirelessDevices` | `List<string>` | no |
| `WirelessGateways` | `List<string>` | no |
| `Description` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `MulticastGroups` | `List<string>` | no |

## GetPartnerAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerAccountId` | `string` | yes |
| `PartnerType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sidewalk` | `SidewalkAccountInfoWithFingerprint` | no |
| `AccountLinked` | `boolean` | no |

## GetPosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Position` | `List<float>` | no |
| `Accuracy` | `Accuracy` | no |
| `SolverType` | `string` | no |
| `SolverProvider` | `string` | no |
| `SolverVersion` | `string` | no |
| `Timestamp` | `string` | no |

## GetPositionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Solvers` | `PositionSolverDetails` | no |
| `Destination` | `string` | no |

## GetPositionEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WiFiAccessPoints` | `List<WiFiAccessPoint>` | no |
| `CellTowers` | `CellTowers` | no |
| `Ip` | `Ip` | no |
| `Gnss` | `Gnss` | no |
| `Timestamp` | `timestamp` | no |
| `AdvancedConfiguration` | `AdvancedConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoJsonPayload` | `blob` | no |

## GetResourceEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IdentifierType` | `string` | yes |
| `PartnerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceRegistrationState` | `DeviceRegistrationStateEventConfiguration` | no |
| `Proximity` | `ProximityEventConfiguration` | no |
| `Join` | `JoinEventConfiguration` | no |
| `ConnectionStatus` | `ConnectionStatusEventConfiguration` | no |
| `MessageDeliveryStatus` | `MessageDeliveryStatusEventConfiguration` | no |

## GetResourceLogLevel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogLevel` | `string` | no |

## GetResourcePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeoJsonPayload` | `blob` | no |

## GetServiceEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceType` | `string` | no |
| `ServiceEndpoint` | `string` | no |
| `ServerTrust` | `string` | no |

## GetServiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Id` | `string` | no |
| `LoRaWAN` | `LoRaWANGetServiceProfileInfo` | no |

## GetWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IdentifierType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DestinationName` | `string` | no |
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `ThingName` | `string` | no |
| `ThingArn` | `string` | no |
| `LoRaWAN` | `LoRaWANDevice` | no |
| `Sidewalk` | `SidewalkDevice` | no |
| `Positioning` | `string` | no |

## GetWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `DestinationName` | `string` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkGetStartImportInfo` | no |
| `CreationTime` | `timestamp` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `InitializedImportedDeviceCount` | `long` | no |
| `PendingImportedDeviceCount` | `long` | no |
| `OnboardedImportedDeviceCount` | `long` | no |
| `FailedImportedDeviceCount` | `long` | no |

## GetWirelessDeviceStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessDeviceId` | `string` | no |
| `LastUplinkReceivedAt` | `string` | no |
| `LoRaWAN` | `LoRaWANDeviceMetadata` | no |
| `Sidewalk` | `SidewalkDeviceMetadata` | no |

## GetWirelessGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IdentifierType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Id` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANGateway` | no |
| `Arn` | `string` | no |
| `ThingName` | `string` | no |
| `ThingArn` | `string` | no |

## GetWirelessGatewayCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IotCertificateId` | `string` | no |
| `LoRaWANNetworkServerCertificateId` | `string` | no |

## GetWirelessGatewayFirmwareInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoRaWAN` | `LoRaWANGatewayCurrentVersion` | no |

## GetWirelessGatewayStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessGatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessGatewayId` | `string` | no |
| `LastUplinkReceivedAt` | `string` | no |
| `ConnectionStatus` | `string` | no |

## GetWirelessGatewayTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WirelessGatewayId` | `string` | no |
| `WirelessGatewayTaskDefinitionId` | `string` | no |
| `LastUplinkReceivedAt` | `string` | no |
| `TaskCreatedAt` | `string` | no |
| `Status` | `string` | no |

## GetWirelessGatewayTaskDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoCreateTasks` | `boolean` | no |
| `Name` | `string` | no |
| `Update` | `UpdateWirelessGatewayTaskCreate` | no |
| `Arn` | `string` | no |

## ListDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DestinationList` | `List<Destinations>` | no |

## ListDeviceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DeviceProfileType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DeviceProfileList` | `List<DeviceProfile>` | no |

## ListDevicesForWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DestinationName` | `string` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkListDevicesForImportInfo` | no |
| `ImportedWirelessDeviceList` | `List<ImportedWirelessDevice>` | no |

## ListEventConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `EventConfigurationsList` | `List<EventConfigurationItem>` | no |

## ListFuotaTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `FuotaTaskList` | `List<FuotaTask>` | no |

## ListMulticastGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MulticastGroupList` | `List<MulticastGroup>` | no |

## ListMulticastGroupsByFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MulticastGroupList` | `List<MulticastGroupByFuotaTask>` | no |

## ListNetworkAnalyzerConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NetworkAnalyzerConfigurationList` | `List<NetworkAnalyzerConfigurations>` | no |

## ListPartnerAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Sidewalk` | `List<SidewalkAccountInfoWithFingerprint>` | no |

## ListPositionConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PositionConfigurationList` | `List<PositionConfigurationItem>` | no |
| `NextToken` | `string` | no |

## ListQueuedMessages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `WirelessDeviceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DownlinkQueueMessagesList` | `List<DownlinkQueueMessage>` | no |

## ListServiceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServiceProfileList` | `List<ServiceProfile>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListWirelessDeviceImportTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WirelessDeviceImportTaskList` | `List<WirelessDeviceImportTask>` | no |

## ListWirelessDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DestinationName` | `string` | no |
| `DeviceProfileId` | `string` | no |
| `ServiceProfileId` | `string` | no |
| `WirelessDeviceType` | `string` | no |
| `FuotaTaskId` | `string` | no |
| `MulticastGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WirelessDeviceList` | `List<WirelessDeviceStatistics>` | no |

## ListWirelessGatewayTaskDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TaskDefinitionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TaskDefinitions` | `List<UpdateWirelessGatewayTaskEntry>` | no |

## ListWirelessGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WirelessGatewayList` | `List<WirelessGatewayStatistics>` | no |

## PutPositionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `Solvers` | `PositionSolverConfigurations` | no |
| `Destination` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutResourceLogLevel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `LogLevel` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetAllResourceLogLevels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetResourceLogLevel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDataToMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `PayloadData` | `string` | yes |
| `WirelessMetadata` | `MulticastWirelessMetadata` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendDataToWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `TransmitMode` | `integer` | yes |
| `PayloadData` | `string` | yes |
| `WirelessMetadata` | `WirelessMetadata` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## StartBulkAssociateWirelessDeviceWithMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `QueryString` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartBulkDisassociateWirelessDeviceFromMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `QueryString` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `LoRaWAN` | `LoRaWANStartFuotaTask` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMulticastGroupSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `LoRaWAN` | `LoRaWANMulticastSession` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartSingleWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationName` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `DeviceName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkSingleStartImportInfo` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## StartWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationName` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkStartImportInfo` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `string` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ExpressionType` | `string` | no |
| `Expression` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventConfigurationByResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceRegistrationState` | `DeviceRegistrationStateResourceTypeEventConfiguration` | no |
| `Proximity` | `ProximityResourceTypeEventConfiguration` | no |
| `Join` | `JoinResourceTypeEventConfiguration` | no |
| `ConnectionStatus` | `ConnectionStatusResourceTypeEventConfiguration` | no |
| `MessageDeliveryStatus` | `MessageDeliveryStatusResourceTypeEventConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFuotaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANFuotaTask` | no |
| `FirmwareUpdateImage` | `string` | no |
| `FirmwareUpdateRole` | `string` | no |
| `RedundancyPercent` | `integer` | no |
| `FragmentSizeBytes` | `integer` | no |
| `FragmentIntervalMS` | `integer` | no |
| `Descriptor` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLogLevelsByResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultLogLevel` | `string` | no |
| `FuotaTaskLogOptions` | `List<FuotaTaskLogOption>` | no |
| `WirelessDeviceLogOptions` | `List<WirelessDeviceLogOption>` | no |
| `WirelessGatewayLogOptions` | `List<WirelessGatewayLogOption>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMetricConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SummaryMetric` | `SummaryMetricConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMulticastGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANMulticast` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNetworkAnalyzerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationName` | `string` | yes |
| `TraceContent` | `TraceContent` | no |
| `WirelessDevicesToAdd` | `List<string>` | no |
| `WirelessDevicesToRemove` | `List<string>` | no |
| `WirelessGatewaysToAdd` | `List<string>` | no |
| `WirelessGatewaysToRemove` | `List<string>` | no |
| `Description` | `string` | no |
| `MulticastGroupsToAdd` | `List<string>` | no |
| `MulticastGroupsToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePartnerAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sidewalk` | `SidewalkUpdateAccount` | yes |
| `PartnerAccountId` | `string` | yes |
| `PartnerType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `Position` | `List<float>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceEventConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IdentifierType` | `string` | yes |
| `PartnerType` | `string` | no |
| `DeviceRegistrationState` | `DeviceRegistrationStateEventConfiguration` | no |
| `Proximity` | `ProximityEventConfiguration` | no |
| `Join` | `JoinEventConfiguration` | no |
| `ConnectionStatus` | `ConnectionStatusEventConfiguration` | no |
| `MessageDeliveryStatus` | `MessageDeliveryStatusEventConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourcePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |
| `ResourceType` | `string` | yes |
| `GeoJsonPayload` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWirelessDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `DestinationName` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LoRaWAN` | `LoRaWANUpdateDevice` | no |
| `Positioning` | `string` | no |
| `Sidewalk` | `SidewalkUpdateWirelessDevice` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWirelessDeviceImportTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Sidewalk` | `SidewalkUpdateImportInfo` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWirelessGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `JoinEuiFilters` | `List<List<string>>` | no |
| `NetIdFilters` | `List<string>` | no |
| `MaxEirp` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


