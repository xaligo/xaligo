# AWS IoT FleetWise

API version: 2021-06-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iotfleetwise/2021-06-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateVehicleFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchCreateVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicles` | `List<CreateVehicleRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicles` | `List<CreateVehicleResponseItem>` | no |
| `errors` | `List<CreateVehicleError>` | no |

## BatchUpdateVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicles` | `List<UpdateVehicleRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicles` | `List<UpdateVehicleResponseItem>` | no |
| `errors` | `List<UpdateVehicleError>` | no |

## CreateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | yes |
| `targetArn` | `string` | yes |
| `startTime` | `timestamp` | no |
| `expiryTime` | `timestamp` | no |
| `postTriggerCollectionDuration` | `long` | no |
| `diagnosticsMode` | `string` | no |
| `spoolingMode` | `string` | no |
| `compression` | `string` | no |
| `priority` | `integer` | no |
| `signalsToCollect` | `List<SignalInformation>` | no |
| `collectionScheme` | `CollectionScheme` | yes |
| `dataExtraDimensions` | `List<string>` | no |
| `tags` | `List<Tag>` | no |
| `dataDestinationConfigs` | `List<DataDestinationConfig>` | no |
| `dataPartitions` | `List<DataPartition>` | no |
| `signalsToFetch` | `List<SignalFetchInformation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |

## CreateDecoderManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `modelManifestArn` | `string` | yes |
| `signalDecoders` | `List<SignalDecoder>` | no |
| `networkInterfaces` | `List<NetworkInterface>` | no |
| `defaultForUnmappedSignals` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |

## CreateModelManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `nodes` | `List<string>` | yes |
| `signalCatalogArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## CreateSignalCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `nodes` | `List<Node>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## CreateStateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | yes |
| `stateTemplateProperties` | `List<string>` | yes |
| `dataExtraDimensions` | `List<string>` | no |
| `metadataExtraDimensions` | `List<string>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `id` | `string` | no |

## CreateVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `modelManifestArn` | `string` | yes |
| `decoderManifestArn` | `string` | yes |
| `attributes` | `Map<string>` | no |
| `associationBehavior` | `string` | no |
| `tags` | `List<Tag>` | no |
| `stateTemplates` | `List<StateTemplateAssociation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | no |
| `arn` | `string` | no |
| `thingArn` | `string` | no |

## DeleteCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |

## DeleteDecoderManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## DeleteFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |

## DeleteModelManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## DeleteSignalCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## DeleteStateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `id` | `string` | no |

## DeleteVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `arn` | `string` | yes |

## DisassociateVehicleFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | no |
| `targetArn` | `string` | no |
| `status` | `string` | no |
| `startTime` | `timestamp` | no |
| `expiryTime` | `timestamp` | no |
| `postTriggerCollectionDuration` | `long` | no |
| `diagnosticsMode` | `string` | no |
| `spoolingMode` | `string` | no |
| `compression` | `string` | no |
| `priority` | `integer` | no |
| `signalsToCollect` | `List<SignalInformation>` | no |
| `collectionScheme` | `CollectionScheme` | no |
| `dataExtraDimensions` | `List<string>` | no |
| `creationTime` | `timestamp` | no |
| `lastModificationTime` | `timestamp` | no |
| `dataDestinationConfigs` | `List<DataDestinationConfig>` | no |
| `dataPartitions` | `List<DataPartition>` | no |
| `signalsToFetch` | `List<SignalFetchInformation>` | no |

## GetDecoderManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `modelManifestArn` | `string` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |
| `message` | `string` | no |

## GetEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKeyId` | `string` | no |
| `encryptionStatus` | `string` | yes |
| `encryptionType` | `string` | yes |
| `errorMessage` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastModificationTime` | `timestamp` | no |

## GetFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |

## GetLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudWatchLogDelivery` | `CloudWatchLogDeliveryOptions` | yes |

## GetModelManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | no |
| `status` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |

## GetRegisterAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customerAccountId` | `string` | yes |
| `accountStatus` | `string` | yes |
| `timestreamRegistrationResponse` | `TimestreamRegistrationResponse` | no |
| `iamRegistrationResponse` | `IamRegistrationResponse` | yes |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |

## GetSignalCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `nodeCounts` | `NodeCounts` | no |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |

## GetStateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `description` | `string` | no |
| `signalCatalogArn` | `string` | no |
| `stateTemplateProperties` | `List<string>` | no |
| `dataExtraDimensions` | `List<string>` | no |
| `metadataExtraDimensions` | `List<string>` | no |
| `creationTime` | `timestamp` | no |
| `lastModificationTime` | `timestamp` | no |
| `id` | `string` | no |

## GetVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | no |
| `arn` | `string` | no |
| `modelManifestArn` | `string` | no |
| `decoderManifestArn` | `string` | no |
| `attributes` | `Map<string>` | no |
| `stateTemplates` | `List<StateTemplateAssociation>` | no |
| `creationTime` | `timestamp` | no |
| `lastModificationTime` | `timestamp` | no |

## GetVehicleStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `vehicleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaigns` | `List<VehicleStatus>` | no |
| `nextToken` | `string` | no |

## ImportDecoderManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `networkFileDefinitions` | `List<NetworkFileDefinition>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## ImportSignalCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `vss` | `FormattedVss` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## ListCampaigns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `status` | `string` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignSummaries` | `List<CampaignSummary>` | no |
| `nextToken` | `string` | no |

## ListDecoderManifestNetworkInterfaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkInterfaces` | `List<NetworkInterface>` | no |
| `nextToken` | `string` | no |

## ListDecoderManifestSignals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signalDecoders` | `List<SignalDecoder>` | no |
| `nextToken` | `string` | no |

## ListDecoderManifests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelManifestArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<DecoderManifestSummary>` | no |
| `nextToken` | `string` | no |

## ListFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetSummaries` | `List<FleetSummary>` | no |
| `nextToken` | `string` | no |

## ListFleetsForVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleets` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListModelManifestNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodes` | `List<Node>` | no |
| `nextToken` | `string` | no |

## ListModelManifests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signalCatalogArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<ModelManifestSummary>` | no |
| `nextToken` | `string` | no |

## ListSignalCatalogNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `signalNodeType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodes` | `List<Node>` | no |
| `nextToken` | `string` | no |

## ListSignalCatalogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<SignalCatalogSummary>` | no |
| `nextToken` | `string` | no |

## ListStateTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `summaries` | `List<StateTemplateSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListVehicles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelManifestArn` | `string` | no |
| `attributeNames` | `List<string>` | no |
| `attributeValues` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `listResponseScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleSummaries` | `List<VehicleSummary>` | no |
| `nextToken` | `string` | no |

## ListVehiclesInFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicles` | `List<string>` | no |
| `nextToken` | `string` | no |

## PutEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKeyId` | `string` | no |
| `encryptionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kmsKeyId` | `string` | no |
| `encryptionStatus` | `string` | yes |
| `encryptionType` | `string` | yes |

## PutLoggingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudWatchLogDelivery` | `CloudWatchLogDeliveryOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `timestreamResources` | `TimestreamResources` | no |
| `iamResources` | `IamResources` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registerAccountStatus` | `string` | yes |
| `timestreamResources` | `TimestreamResources` | no |
| `iamResources` | `IamResources` | yes |
| `creationTime` | `timestamp` | yes |
| `lastModificationTime` | `timestamp` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `dataExtraDimensions` | `List<string>` | no |
| `action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |

## UpdateDecoderManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `signalDecodersToAdd` | `List<SignalDecoder>` | no |
| `signalDecodersToUpdate` | `List<SignalDecoder>` | no |
| `signalDecodersToRemove` | `List<string>` | no |
| `networkInterfacesToAdd` | `List<NetworkInterface>` | no |
| `networkInterfacesToUpdate` | `List<NetworkInterface>` | no |
| `networkInterfacesToRemove` | `List<string>` | no |
| `status` | `string` | no |
| `defaultForUnmappedSignals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## UpdateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fleetId` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |

## UpdateModelManifest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `nodesToAdd` | `List<string>` | no |
| `nodesToRemove` | `List<string>` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## UpdateSignalCatalog

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `nodesToAdd` | `List<Node>` | no |
| `nodesToUpdate` | `List<Node>` | no |
| `nodesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |

## UpdateStateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `stateTemplatePropertiesToAdd` | `List<string>` | no |
| `stateTemplatePropertiesToRemove` | `List<string>` | no |
| `dataExtraDimensions` | `List<string>` | no |
| `metadataExtraDimensions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `arn` | `string` | no |
| `id` | `string` | no |

## UpdateVehicle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | yes |
| `modelManifestArn` | `string` | no |
| `decoderManifestArn` | `string` | no |
| `attributes` | `Map<string>` | no |
| `attributeUpdateMode` | `string` | no |
| `stateTemplatesToAdd` | `List<StateTemplateAssociation>` | no |
| `stateTemplatesToRemove` | `List<string>` | no |
| `stateTemplatesToUpdate` | `List<StateTemplateAssociation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vehicleName` | `string` | no |
| `arn` | `string` | no |

