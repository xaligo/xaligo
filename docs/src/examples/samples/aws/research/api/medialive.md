# AWS Elemental MediaLive

API version: 2017-10-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/medialive/2017-10-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptInputDeviceTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchDelete

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelIds` | `List<string>` | no |
| `InputIds` | `List<string>` | no |
| `InputSecurityGroupIds` | `List<string>` | no |
| `MultiplexIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failed` | `List<BatchFailedResultModel>` | no |
| `Successful` | `List<BatchSuccessfulResultModel>` | no |

## BatchStart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelIds` | `List<string>` | no |
| `MultiplexIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failed` | `List<BatchFailedResultModel>` | no |
| `Successful` | `List<BatchSuccessfulResultModel>` | no |

## BatchStop

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelIds` | `List<string>` | no |
| `MultiplexIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failed` | `List<BatchFailedResultModel>` | no |
| `Successful` | `List<BatchSuccessfulResultModel>` | no |

## BatchUpdateSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `Creates` | `BatchScheduleActionCreateRequest` | no |
| `Deletes` | `BatchScheduleActionDeleteRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Creates` | `BatchScheduleActionCreateResult` | no |
| `Deletes` | `BatchScheduleActionDeleteResult` | no |

## CancelInputDeviceTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ClaimDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceCreateSettings` | no |
| `Name` | `string` | no |
| `RequestId` | `string` | no |
| `Reserved` | `string` | no |
| `RoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettings` | no |
| `AnywhereSettings` | `AnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionRequest` | no |
| `DryRun` | `boolean` | no |
| `LinkedChannelSettings` | `LinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `InferenceSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `Channel` | no |

## CreateInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destinations` | `List<InputDestinationRequest>` | no |
| `InputDevices` | `List<InputDeviceSettings>` | no |
| `InputSecurityGroups` | `List<string>` | no |
| `MediaConnectFlows` | `List<MediaConnectFlowRequest>` | no |
| `Name` | `string` | no |
| `RequestId` | `string` | no |
| `RoleArn` | `string` | no |
| `Sources` | `List<InputSourceRequest>` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `Vpc` | `InputVpcRequest` | no |
| `SrtSettings` | `SrtSettingsRequest` | no |
| `InputNetworkLocation` | `string` | no |
| `MulticastSettings` | `MulticastSettingsCreateRequest` | no |
| `Smpte2110ReceiverGroupSettings` | `Smpte2110ReceiverGroupSettings` | no |
| `SdiSources` | `List<string>` | no |
| `RouterSettings` | `RouterSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Input` | `Input` | no |

## CreateInputSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |
| `WhitelistRules` | `List<InputWhitelistRuleCidr>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroup` | `InputSecurityGroup` | no |

## CreateMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | yes |
| `MultiplexSettings` | `MultiplexSettings` | yes |
| `Name` | `string` | yes |
| `RequestId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Multiplex` | `Multiplex` | no |

## CreateMultiplexProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |
| `MultiplexProgramSettings` | `MultiplexProgramSettings` | yes |
| `ProgramName` | `string` | yes |
| `RequestId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexProgram` | `MultiplexProgram` | no |

## CreatePartnerInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputId` | `string` | yes |
| `RequestId` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Input` | `Input` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EgressEndpoints` | `List<ChannelEgressEndpoint>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `Id` | `string` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceStatus` | no |
| `Name` | `string` | no |
| `PipelineDetails` | `List<PipelineDetail>` | no |
| `PipelinesRunningCount` | `integer` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettingsDescription` | no |
| `AnywhereSettings` | `DescribeAnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionResponse` | no |
| `LinkedChannelSettings` | `DescribeLinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `DescribeInferenceSettings` | no |

## DeleteInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInputSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSecurityGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AvailabilityZones` | `List<string>` | no |
| `Destinations` | `List<MultiplexOutputDestination>` | no |
| `Id` | `string` | no |
| `MultiplexSettings` | `MultiplexSettings` | no |
| `Name` | `string` | no |
| `PipelinesRunningCount` | `integer` | no |
| `ProgramCount` | `integer` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## DeleteMultiplexProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |
| `ProgramName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | no |
| `MultiplexProgramSettings` | `MultiplexProgramSettings` | no |
| `PacketIdentifiersMap` | `MultiplexProgramPacketIdentifiersMap` | no |
| `PipelineDetails` | `List<MultiplexProgramPipelineDetail>` | no |
| `ProgramName` | `string` | no |

## DeleteReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Count` | `integer` | no |
| `CurrencyCode` | `string` | no |
| `Duration` | `integer` | no |
| `DurationUnits` | `string` | no |
| `End` | `string` | no |
| `FixedPrice` | `double` | no |
| `Name` | `string` | no |
| `OfferingDescription` | `string` | no |
| `OfferingId` | `string` | no |
| `OfferingType` | `string` | no |
| `Region` | `string` | no |
| `RenewalSettings` | `RenewalSettings` | no |
| `ReservationId` | `string` | no |
| `ResourceSpecification` | `ReservationResourceSpecification` | no |
| `Start` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UsagePrice` | `double` | no |

## DeleteSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountConfiguration` | `AccountConfiguration` | no |

## DescribeChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EgressEndpoints` | `List<ChannelEgressEndpoint>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `Id` | `string` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceStatus` | no |
| `Name` | `string` | no |
| `PipelineDetails` | `List<PipelineDetail>` | no |
| `PipelinesRunningCount` | `integer` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettingsDescription` | no |
| `AnywhereSettings` | `DescribeAnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionResponse` | no |
| `LinkedChannelSettings` | `DescribeLinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `DescribeInferenceSettings` | no |

## DescribeInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AttachedChannels` | `List<string>` | no |
| `Destinations` | `List<InputDestination>` | no |
| `Id` | `string` | no |
| `InputClass` | `string` | no |
| `InputDevices` | `List<InputDeviceSettings>` | no |
| `InputPartnerIds` | `List<string>` | no |
| `InputSourceType` | `string` | no |
| `MediaConnectFlows` | `List<MediaConnectFlow>` | no |
| `Name` | `string` | no |
| `RoleArn` | `string` | no |
| `SecurityGroups` | `List<string>` | no |
| `Sources` | `List<InputSource>` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `SrtSettings` | `SrtSettings` | no |
| `InputNetworkLocation` | `string` | no |
| `MulticastSettings` | `MulticastSettings` | no |
| `Smpte2110ReceiverGroupSettings` | `Smpte2110ReceiverGroupSettings` | no |
| `SdiSources` | `List<string>` | no |
| `RouterSettings` | `RouterInputSettings` | no |

## DescribeInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ConnectionState` | `string` | no |
| `DeviceSettingsSyncState` | `string` | no |
| `DeviceUpdateStatus` | `string` | no |
| `HdDeviceSettings` | `InputDeviceHdSettings` | no |
| `Id` | `string` | no |
| `MacAddress` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `InputDeviceNetworkSettings` | no |
| `SerialNumber` | `string` | no |
| `Type` | `string` | no |
| `UhdDeviceSettings` | `InputDeviceUhdSettings` | no |
| `Tags` | `Map<string>` | no |
| `AvailabilityZone` | `string` | no |
| `MedialiveInputArns` | `List<string>` | no |
| `OutputType` | `string` | no |

## DescribeInputDeviceThumbnail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |
| `Accept` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `blob` | no |
| `ContentType` | `string` | no |
| `ContentLength` | `long` | no |
| `ETag` | `string` | no |
| `LastModified` | `timestamp` | no |

## DescribeInputSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSecurityGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Inputs` | `List<string>` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `WhitelistRules` | `List<InputWhitelistRule>` | no |
| `Channels` | `List<string>` | no |

## DescribeMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AvailabilityZones` | `List<string>` | no |
| `Destinations` | `List<MultiplexOutputDestination>` | no |
| `Id` | `string` | no |
| `MultiplexSettings` | `MultiplexSettings` | no |
| `Name` | `string` | no |
| `PipelinesRunningCount` | `integer` | no |
| `ProgramCount` | `integer` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeMultiplexProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |
| `ProgramName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | no |
| `MultiplexProgramSettings` | `MultiplexProgramSettings` | no |
| `PacketIdentifiersMap` | `MultiplexProgramPacketIdentifiersMap` | no |
| `PipelineDetails` | `List<MultiplexProgramPipelineDetail>` | no |
| `ProgramName` | `string` | no |

## DescribeOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OfferingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CurrencyCode` | `string` | no |
| `Duration` | `integer` | no |
| `DurationUnits` | `string` | no |
| `FixedPrice` | `double` | no |
| `OfferingDescription` | `string` | no |
| `OfferingId` | `string` | no |
| `OfferingType` | `string` | no |
| `Region` | `string` | no |
| `ResourceSpecification` | `ReservationResourceSpecification` | no |
| `UsagePrice` | `double` | no |

## DescribeReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Count` | `integer` | no |
| `CurrencyCode` | `string` | no |
| `Duration` | `integer` | no |
| `DurationUnits` | `string` | no |
| `End` | `string` | no |
| `FixedPrice` | `double` | no |
| `Name` | `string` | no |
| `OfferingDescription` | `string` | no |
| `OfferingId` | `string` | no |
| `OfferingType` | `string` | no |
| `Region` | `string` | no |
| `RenewalSettings` | `RenewalSettings` | no |
| `ReservationId` | `string` | no |
| `ResourceSpecification` | `ReservationResourceSpecification` | no |
| `Start` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UsagePrice` | `double` | no |

## DescribeSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ScheduleActions` | `List<ScheduleAction>` | no |

## DescribeThumbnails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `PipelineId` | `string` | yes |
| `ThumbnailType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThumbnailDetails` | `List<ThumbnailDetail>` | no |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<ChannelSummary>` | no |
| `NextToken` | `string` | no |

## ListInputDeviceTransfers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TransferType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceTransfers` | `List<TransferringInputDeviceSummary>` | no |
| `NextToken` | `string` | no |

## ListInputDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDevices` | `List<InputDeviceSummary>` | no |
| `NextToken` | `string` | no |

## ListInputSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSecurityGroups` | `List<InputSecurityGroup>` | no |
| `NextToken` | `string` | no |

## ListInputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Inputs` | `List<Input>` | no |
| `NextToken` | `string` | no |

## ListMultiplexPrograms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `MultiplexId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexPrograms` | `List<MultiplexProgramSummary>` | no |
| `NextToken` | `string` | no |

## ListMultiplexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Multiplexes` | `List<MultiplexSummary>` | no |
| `NextToken` | `string` | no |

## ListOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelClass` | `string` | no |
| `ChannelConfiguration` | `string` | no |
| `Codec` | `string` | no |
| `Duration` | `string` | no |
| `MaxResults` | `integer` | no |
| `MaximumBitrate` | `string` | no |
| `MaximumFramerate` | `string` | no |
| `NextToken` | `string` | no |
| `Resolution` | `string` | no |
| `ResourceType` | `string` | no |
| `SpecialFeature` | `string` | no |
| `VideoQuality` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Offerings` | `List<Offering>` | no |

## ListReservations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelClass` | `string` | no |
| `Codec` | `string` | no |
| `MaxResults` | `integer` | no |
| `MaximumBitrate` | `string` | no |
| `MaximumFramerate` | `string` | no |
| `NextToken` | `string` | no |
| `Resolution` | `string` | no |
| `ResourceType` | `string` | no |
| `SpecialFeature` | `string` | no |
| `VideoQuality` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Reservations` | `List<Reservation>` | no |

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
| `Count` | `integer` | yes |
| `Name` | `string` | no |
| `OfferingId` | `string` | yes |
| `RenewalSettings` | `RenewalSettings` | no |
| `RequestId` | `string` | no |
| `Start` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Reservation` | `Reservation` | no |

## RebootInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Force` | `string` | no |
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectInputDeviceTransfer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EgressEndpoints` | `List<ChannelEgressEndpoint>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `Id` | `string` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceStatus` | no |
| `Name` | `string` | no |
| `PipelineDetails` | `List<PipelineDetail>` | no |
| `PipelinesRunningCount` | `integer` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettingsDescription` | no |
| `AnywhereSettings` | `DescribeAnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionResponse` | no |
| `LinkedChannelSettings` | `DescribeLinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `DescribeInferenceSettings` | no |

## StartInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartInputDeviceMaintenanceWindow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AvailabilityZones` | `List<string>` | no |
| `Destinations` | `List<MultiplexOutputDestination>` | no |
| `Id` | `string` | no |
| `MultiplexSettings` | `MultiplexSettings` | no |
| `Name` | `string` | no |
| `PipelinesRunningCount` | `integer` | no |
| `ProgramCount` | `integer` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## StopChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EgressEndpoints` | `List<ChannelEgressEndpoint>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `Id` | `string` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceStatus` | no |
| `Name` | `string` | no |
| `PipelineDetails` | `List<PipelineDetail>` | no |
| `PipelinesRunningCount` | `integer` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettingsDescription` | no |
| `AnywhereSettings` | `DescribeAnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionResponse` | no |
| `LinkedChannelSettings` | `DescribeLinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `DescribeInferenceSettings` | no |

## StopInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AvailabilityZones` | `List<string>` | no |
| `Destinations` | `List<MultiplexOutputDestination>` | no |
| `Id` | `string` | no |
| `MultiplexSettings` | `MultiplexSettings` | no |
| `Name` | `string` | no |
| `PipelinesRunningCount` | `integer` | no |
| `ProgramCount` | `integer` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## TransferInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputDeviceId` | `string` | yes |
| `TargetCustomerId` | `string` | no |
| `TargetRegion` | `string` | no |
| `TransferMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountConfiguration` | `AccountConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountConfiguration` | `AccountConfiguration` | no |

## UpdateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelId` | `string` | yes |
| `Destinations` | `List<OutputDestination>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceUpdateSettings` | no |
| `Name` | `string` | no |
| `RoleArn` | `string` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionRequest` | no |
| `DryRun` | `boolean` | no |
| `AnywhereSettings` | `AnywhereSettings` | no |
| `LinkedChannelSettings` | `LinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `InferenceSettings` | no |
| `SpecialRouterSettings` | `SpecialRouterSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `Channel` | no |

## UpdateChannelClass

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelClass` | `string` | yes |
| `ChannelId` | `string` | yes |
| `Destinations` | `List<OutputDestination>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `Channel` | no |

## UpdateInput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Destinations` | `List<InputDestinationRequest>` | no |
| `InputDevices` | `List<InputDeviceRequest>` | no |
| `InputId` | `string` | yes |
| `InputSecurityGroups` | `List<string>` | no |
| `MediaConnectFlows` | `List<MediaConnectFlowRequest>` | no |
| `Name` | `string` | no |
| `RoleArn` | `string` | no |
| `Sources` | `List<InputSourceRequest>` | no |
| `SrtSettings` | `SrtSettingsRequest` | no |
| `MulticastSettings` | `MulticastSettingsUpdateRequest` | no |
| `Smpte2110ReceiverGroupSettings` | `Smpte2110ReceiverGroupSettings` | no |
| `SdiSources` | `List<string>` | no |
| `SpecialRouterSettings` | `SpecialRouterSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Input` | `Input` | no |

## UpdateInputDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HdDeviceSettings` | `InputDeviceConfigurableSettings` | no |
| `InputDeviceId` | `string` | yes |
| `Name` | `string` | no |
| `UhdDeviceSettings` | `InputDeviceConfigurableSettings` | no |
| `AvailabilityZone` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ConnectionState` | `string` | no |
| `DeviceSettingsSyncState` | `string` | no |
| `DeviceUpdateStatus` | `string` | no |
| `HdDeviceSettings` | `InputDeviceHdSettings` | no |
| `Id` | `string` | no |
| `MacAddress` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `InputDeviceNetworkSettings` | no |
| `SerialNumber` | `string` | no |
| `Type` | `string` | no |
| `UhdDeviceSettings` | `InputDeviceUhdSettings` | no |
| `Tags` | `Map<string>` | no |
| `AvailabilityZone` | `string` | no |
| `MedialiveInputArns` | `List<string>` | no |
| `OutputType` | `string` | no |

## UpdateInputSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InputSecurityGroupId` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `WhitelistRules` | `List<InputWhitelistRuleCidr>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroup` | `InputSecurityGroup` | no |

## UpdateMultiplex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |
| `MultiplexSettings` | `MultiplexSettings` | no |
| `Name` | `string` | no |
| `PacketIdentifiersMapping` | `Map<MultiplexProgramPacketIdentifiersMap>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Multiplex` | `Multiplex` | no |

## UpdateMultiplexProgram

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexId` | `string` | yes |
| `MultiplexProgramSettings` | `MultiplexProgramSettings` | no |
| `ProgramName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiplexProgram` | `MultiplexProgram` | no |

## UpdateReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `RenewalSettings` | `RenewalSettings` | no |
| `ReservationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Reservation` | `Reservation` | no |

## RestartChannelPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `PipelineIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CdiInputSpecification` | `CdiInputSpecification` | no |
| `ChannelClass` | `string` | no |
| `Destinations` | `List<OutputDestination>` | no |
| `EgressEndpoints` | `List<ChannelEgressEndpoint>` | no |
| `EncoderSettings` | `EncoderSettings` | no |
| `Id` | `string` | no |
| `InputAttachments` | `List<InputAttachment>` | no |
| `InputSpecification` | `InputSpecification` | no |
| `LogLevel` | `string` | no |
| `Maintenance` | `MaintenanceStatus` | no |
| `MaintenanceStatus` | `string` | no |
| `Name` | `string` | no |
| `PipelineDetails` | `List<PipelineDetail>` | no |
| `PipelinesRunningCount` | `integer` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Vpc` | `VpcOutputSettingsDescription` | no |
| `AnywhereSettings` | `DescribeAnywhereSettings` | no |
| `ChannelEngineVersion` | `ChannelEngineVersionResponse` | no |
| `LinkedChannelSettings` | `DescribeLinkedChannelSettings` | no |
| `ChannelSecurityGroups` | `List<string>` | no |
| `InferenceSettings` | `DescribeInferenceSettings` | no |

## CreateCloudWatchAlarmTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComparisonOperator` | `string` | yes |
| `DatapointsToAlarm` | `integer` | no |
| `Description` | `string` | no |
| `EvaluationPeriods` | `integer` | yes |
| `GroupIdentifier` | `string` | yes |
| `MetricName` | `string` | yes |
| `Name` | `string` | yes |
| `Period` | `integer` | yes |
| `Statistic` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `TargetResourceType` | `string` | yes |
| `Threshold` | `double` | yes |
| `TreatMissingData` | `string` | yes |
| `RequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ComparisonOperator` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DatapointsToAlarm` | `integer` | no |
| `Description` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `MetricName` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Period` | `integer` | no |
| `Statistic` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TargetResourceType` | `string` | no |
| `Threshold` | `double` | no |
| `TreatMissingData` | `string` | no |

## CreateCloudWatchAlarmTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `RequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateEventBridgeRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `EventTargets` | `List<EventBridgeRuleTemplateTarget>` | no |
| `EventType` | `string` | yes |
| `GroupIdentifier` | `string` | yes |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `RequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `EventTargets` | `List<EventBridgeRuleTemplateTarget>` | no |
| `EventType` | `string` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateEventBridgeRuleTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `RequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateSignalMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudWatchAlarmTemplateGroupIdentifiers` | `List<string>` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | yes |
| `EventBridgeRuleTemplateGroupIdentifiers` | `List<string>` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `RequestId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CloudWatchAlarmTemplateGroupIds` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `EventBridgeRuleTemplateGroupIds` | `List<string>` | no |
| `FailedMediaResourceMap` | `Map<MediaResource>` | no |
| `Id` | `string` | no |
| `LastDiscoveredAt` | `timestamp` | no |
| `LastSuccessfulMonitorDeployment` | `SuccessfulMonitorDeployment` | no |
| `MediaResourceMap` | `Map<MediaResource>` | no |
| `ModifiedAt` | `timestamp` | no |
| `MonitorChangesPendingDeployment` | `boolean` | no |
| `MonitorDeployment` | `MonitorDeployment` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## DeleteCloudWatchAlarmTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudWatchAlarmTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventBridgeRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventBridgeRuleTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSignalMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCloudWatchAlarmTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ComparisonOperator` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DatapointsToAlarm` | `integer` | no |
| `Description` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `MetricName` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Period` | `integer` | no |
| `Statistic` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TargetResourceType` | `string` | no |
| `Threshold` | `double` | no |
| `TreatMissingData` | `string` | no |

## GetCloudWatchAlarmTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetEventBridgeRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `EventTargets` | `List<EventBridgeRuleTemplateTarget>` | no |
| `EventType` | `string` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetEventBridgeRuleTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetSignalMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CloudWatchAlarmTemplateGroupIds` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `EventBridgeRuleTemplateGroupIds` | `List<string>` | no |
| `FailedMediaResourceMap` | `Map<MediaResource>` | no |
| `Id` | `string` | no |
| `LastDiscoveredAt` | `timestamp` | no |
| `LastSuccessfulMonitorDeployment` | `SuccessfulMonitorDeployment` | no |
| `MediaResourceMap` | `Map<MediaResource>` | no |
| `ModifiedAt` | `timestamp` | no |
| `MonitorChangesPendingDeployment` | `boolean` | no |
| `MonitorDeployment` | `MonitorDeployment` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## ListCloudWatchAlarmTemplateGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Scope` | `string` | no |
| `SignalMapIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudWatchAlarmTemplateGroups` | `List<CloudWatchAlarmTemplateGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListCloudWatchAlarmTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Scope` | `string` | no |
| `SignalMapIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudWatchAlarmTemplates` | `List<CloudWatchAlarmTemplateSummary>` | no |
| `NextToken` | `string` | no |

## ListEventBridgeRuleTemplateGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SignalMapIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventBridgeRuleTemplateGroups` | `List<EventBridgeRuleTemplateGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListEventBridgeRuleTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifier` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SignalMapIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventBridgeRuleTemplates` | `List<EventBridgeRuleTemplateSummary>` | no |
| `NextToken` | `string` | no |

## ListSignalMaps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudWatchAlarmTemplateGroupIdentifier` | `string` | no |
| `EventBridgeRuleTemplateGroupIdentifier` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SignalMaps` | `List<SignalMapSummary>` | no |

## StartDeleteMonitorDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CloudWatchAlarmTemplateGroupIds` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `EventBridgeRuleTemplateGroupIds` | `List<string>` | no |
| `FailedMediaResourceMap` | `Map<MediaResource>` | no |
| `Id` | `string` | no |
| `LastDiscoveredAt` | `timestamp` | no |
| `LastSuccessfulMonitorDeployment` | `SuccessfulMonitorDeployment` | no |
| `MediaResourceMap` | `Map<MediaResource>` | no |
| `ModifiedAt` | `timestamp` | no |
| `MonitorChangesPendingDeployment` | `boolean` | no |
| `MonitorDeployment` | `MonitorDeployment` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## StartMonitorDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CloudWatchAlarmTemplateGroupIds` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `EventBridgeRuleTemplateGroupIds` | `List<string>` | no |
| `FailedMediaResourceMap` | `Map<MediaResource>` | no |
| `Id` | `string` | no |
| `LastDiscoveredAt` | `timestamp` | no |
| `LastSuccessfulMonitorDeployment` | `SuccessfulMonitorDeployment` | no |
| `MediaResourceMap` | `Map<MediaResource>` | no |
| `ModifiedAt` | `timestamp` | no |
| `MonitorChangesPendingDeployment` | `boolean` | no |
| `MonitorDeployment` | `MonitorDeployment` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## StartUpdateSignalMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudWatchAlarmTemplateGroupIdentifiers` | `List<string>` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `EventBridgeRuleTemplateGroupIdentifiers` | `List<string>` | no |
| `ForceRediscovery` | `boolean` | no |
| `Identifier` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CloudWatchAlarmTemplateGroupIds` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `DiscoveryEntryPointArn` | `string` | no |
| `ErrorMessage` | `string` | no |
| `EventBridgeRuleTemplateGroupIds` | `List<string>` | no |
| `FailedMediaResourceMap` | `Map<MediaResource>` | no |
| `Id` | `string` | no |
| `LastDiscoveredAt` | `timestamp` | no |
| `LastSuccessfulMonitorDeployment` | `SuccessfulMonitorDeployment` | no |
| `MediaResourceMap` | `Map<MediaResource>` | no |
| `ModifiedAt` | `timestamp` | no |
| `MonitorChangesPendingDeployment` | `boolean` | no |
| `MonitorDeployment` | `MonitorDeployment` | no |
| `Name` | `string` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateCloudWatchAlarmTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComparisonOperator` | `string` | no |
| `DatapointsToAlarm` | `integer` | no |
| `Description` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `GroupIdentifier` | `string` | no |
| `Identifier` | `string` | yes |
| `MetricName` | `string` | no |
| `Name` | `string` | no |
| `Period` | `integer` | no |
| `Statistic` | `string` | no |
| `TargetResourceType` | `string` | no |
| `Threshold` | `double` | no |
| `TreatMissingData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ComparisonOperator` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `DatapointsToAlarm` | `integer` | no |
| `Description` | `string` | no |
| `EvaluationPeriods` | `integer` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `MetricName` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Period` | `integer` | no |
| `Statistic` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TargetResourceType` | `string` | no |
| `Threshold` | `double` | no |
| `TreatMissingData` | `string` | no |

## UpdateCloudWatchAlarmTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateEventBridgeRuleTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `EventTargets` | `List<EventBridgeRuleTemplateTarget>` | no |
| `EventType` | `string` | no |
| `GroupIdentifier` | `string` | no |
| `Identifier` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `EventTargets` | `List<EventBridgeRuleTemplateTarget>` | no |
| `EventType` | `string` | no |
| `GroupId` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateEventBridgeRuleTemplateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `Id` | `string` | no |
| `ModifiedAt` | `timestamp` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateChannelPlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |
| `RequestId` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Channels` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |
| `State` | `string` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterType` | `string` | no |
| `InstanceRoleArn` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettingsCreateRequest` | no |
| `RequestId` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelIds` | `List<string>` | no |
| `ClusterType` | `string` | no |
| `Id` | `string` | no |
| `InstanceRoleArn` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettings` | no |
| `State` | `string` | no |

## CreateNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpPools` | `List<IpPoolCreateRequest>` | no |
| `Name` | `string` | no |
| `RequestId` | `string` | no |
| `Routes` | `List<RouteCreateRequest>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssociatedClusterIds` | `List<string>` | no |
| `Id` | `string` | no |
| `IpPools` | `List<IpPool>` | no |
| `Name` | `string` | no |
| `Routes` | `List<Route>` | no |
| `State` | `string` | no |

## CreateNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMappingCreateRequest>` | no |
| `RequestId` | `string` | no |
| `Role` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelPlacementGroups` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `ConnectionState` | `string` | no |
| `Id` | `string` | no |
| `InstanceArn` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `Role` | `string` | no |
| `State` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMapping>` | no |

## CreateNodeRegistrationScript

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `RequestId` | `string` | no |
| `Role` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NodeRegistrationScript` | `string` | no |

## DeleteChannelPlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelPlacementGroupId` | `string` | yes |
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Channels` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |
| `State` | `string` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelIds` | `List<string>` | no |
| `ClusterType` | `string` | no |
| `Id` | `string` | no |
| `InstanceRoleArn` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettings` | no |
| `State` | `string` | no |

## DeleteNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssociatedClusterIds` | `List<string>` | no |
| `Id` | `string` | no |
| `IpPools` | `List<IpPool>` | no |
| `Name` | `string` | no |
| `Routes` | `List<Route>` | no |
| `State` | `string` | no |

## DeleteNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `NodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelPlacementGroups` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `ConnectionState` | `string` | no |
| `Id` | `string` | no |
| `InstanceArn` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `Role` | `string` | no |
| `State` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMapping>` | no |

## DescribeChannelPlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelPlacementGroupId` | `string` | yes |
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Channels` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |
| `State` | `string` | no |

## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelIds` | `List<string>` | no |
| `ClusterType` | `string` | no |
| `Id` | `string` | no |
| `InstanceRoleArn` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettings` | no |
| `State` | `string` | no |

## DescribeNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NetworkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssociatedClusterIds` | `List<string>` | no |
| `Id` | `string` | no |
| `IpPools` | `List<IpPool>` | no |
| `Name` | `string` | no |
| `Routes` | `List<Route>` | no |
| `State` | `string` | no |

## DescribeNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `NodeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelPlacementGroups` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `ConnectionState` | `string` | no |
| `Id` | `string` | no |
| `InstanceArn` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `Role` | `string` | no |
| `State` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMapping>` | no |

## ListChannelPlacementGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelPlacementGroups` | `List<DescribeChannelPlacementGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Clusters` | `List<DescribeClusterSummary>` | no |
| `NextToken` | `string` | no |

## ListNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Networks` | `List<DescribeNetworkSummary>` | no |
| `NextToken` | `string` | no |

## ListNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Nodes` | `List<DescribeNodeSummary>` | no |

## UpdateChannelPlacementGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelPlacementGroupId` | `string` | yes |
| `ClusterId` | `string` | yes |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Channels` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Nodes` | `List<string>` | no |
| `State` | `string` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettingsUpdateRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelIds` | `List<string>` | no |
| `ClusterType` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `NetworkSettings` | `ClusterNetworkSettings` | no |
| `State` | `string` | no |

## UpdateNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpPools` | `List<IpPoolUpdateRequest>` | no |
| `Name` | `string` | no |
| `NetworkId` | `string` | yes |
| `Routes` | `List<RouteUpdateRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AssociatedClusterIds` | `List<string>` | no |
| `Id` | `string` | no |
| `IpPools` | `List<IpPool>` | no |
| `Name` | `string` | no |
| `Routes` | `List<Route>` | no |
| `State` | `string` | no |

## UpdateNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `Name` | `string` | no |
| `NodeId` | `string` | yes |
| `Role` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMappingUpdateRequest>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelPlacementGroups` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `ConnectionState` | `string` | no |
| `Id` | `string` | no |
| `InstanceArn` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `Role` | `string` | no |
| `State` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMapping>` | no |

## UpdateNodeState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `NodeId` | `string` | yes |
| `State` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `ChannelPlacementGroups` | `List<string>` | no |
| `ClusterId` | `string` | no |
| `ConnectionState` | `string` | no |
| `Id` | `string` | no |
| `InstanceArn` | `string` | no |
| `Name` | `string` | no |
| `NodeInterfaceMappings` | `List<NodeInterfaceMapping>` | no |
| `Role` | `string` | no |
| `State` | `string` | no |
| `SdiSourceMappings` | `List<SdiSourceMapping>` | no |

## ListVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Versions` | `List<ChannelEngineVersionResponse>` | no |

## CreateSdiSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mode` | `string` | no |
| `Name` | `string` | no |
| `RequestId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSource` | `SdiSource` | no |

## DeleteSdiSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSource` | `SdiSource` | no |

## DescribeSdiSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSource` | `SdiSource` | no |

## ListSdiSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SdiSources` | `List<SdiSourceSummary>` | no |

## UpdateSdiSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Mode` | `string` | no |
| `Name` | `string` | no |
| `SdiSourceId` | `string` | yes |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SdiSource` | `SdiSource` | no |

## ListAlerts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `StateFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alerts` | `List<ChannelAlert>` | no |
| `NextToken` | `string` | no |

## ListClusterAlerts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `StateFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alerts` | `List<ClusterAlert>` | no |
| `NextToken` | `string` | no |

## ListMultiplexAlerts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `MultiplexId` | `string` | yes |
| `NextToken` | `string` | no |
| `StateFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alerts` | `List<MultiplexAlert>` | no |
| `NextToken` | `string` | no |

