# Managed integrations for AWS IoT Device Management

API version: 2025-03-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/iot-managed-integrations/2025-03-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectorDestinationId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |
| `GeneralAuthorization` | `GeneralAuthorizationName` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OAuthAuthorizationUrl` | `string` | yes |
| `AccountAssociationId` | `string` | yes |
| `AssociationState` | `string` | yes |
| `Arn` | `string` | no |

## CreateCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EndpointConfig` | `EndpointConfig` | yes |
| `Description` | `string` | no |
| `EndpointType` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateConnectorDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CloudConnectorId` | `string` | yes |
| `AuthType` | `string` | no |
| `AuthConfig` | `AuthConfig` | yes |
| `SecretsManager` | `SecretsManager` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateCredentialLocker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## CreateDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliveryDestinationArn` | `string` | yes |
| `DeliveryDestinationType` | `string` | yes |
| `Name` | `string` | yes |
| `RoleArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

## CreateEventLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceType` | `string` | yes |
| `ResourceId` | `string` | no |
| `EventLogLevel` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |

## CreateManagedThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Role` | `string` | yes |
| `Owner` | `string` | no |
| `CredentialLockerId` | `string` | no |
| `AuthenticationMaterial` | `string` | yes |
| `AuthenticationMaterialType` | `string` | yes |
| `WiFiSimpleSetupConfiguration` | `WiFiSimpleSetupConfiguration` | no |
| `SerialNumber` | `string` | no |
| `Brand` | `string` | no |
| `Model` | `string` | no |
| `Name` | `string` | no |
| `CapabilityReport` | `CapabilityReport` | no |
| `CapabilitySchemas` | `List<CapabilitySchemaItem>` | no |
| `Capabilities` | `string` | no |
| `ClientToken` | `string` | no |
| `Classification` | `string` | no |
| `Tags` | `Map<string>` | no |
| `MetaData` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## CreateNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | yes |
| `DestinationName` | `string` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | no |

## CreateOtaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `S3Url` | `string` | yes |
| `Protocol` | `string` | no |
| `Target` | `List<string>` | no |
| `TaskConfigurationId` | `string` | no |
| `OtaMechanism` | `string` | no |
| `OtaType` | `string` | yes |
| `OtaTargetQueryString` | `string` | no |
| `ClientToken` | `string` | no |
| `OtaSchedulingConfig` | `OtaTaskSchedulingConfig` | no |
| `OtaTaskExecutionRetryConfig` | `OtaTaskExecutionRetryConfig` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `Description` | `string` | no |

## CreateOtaTaskConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Name` | `string` | no |
| `PushConfig` | `PushConfig` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskConfigurationId` | `string` | no |

## CreateProvisioningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisioningType` | `string` | yes |
| `CaCertificate` | `string` | no |
| `ClaimCertificate` | `string` | no |
| `Name` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `ProvisioningType` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `ClaimCertificate` | `string` | no |
| `ClaimCertificatePrivateKey` | `string` | no |

## DeleteAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectorDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCredentialLocker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteManagedThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOtaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOtaTaskConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisioningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |
| `AccountAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | yes |
| `AssociationState` | `string` | yes |
| `ErrorMessage` | `string` | no |
| `ConnectorDestinationId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Arn` | `string` | no |
| `OAuthAuthorizationUrl` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `GeneralAuthorization` | `GeneralAuthorizationName` | no |

## GetCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EndpointConfig` | `EndpointConfig` | yes |
| `Description` | `string` | no |
| `EndpointType` | `string` | no |
| `Id` | `string` | no |
| `Type` | `string` | no |

## GetConnectorDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `CloudConnectorId` | `string` | no |
| `Id` | `string` | no |
| `AuthType` | `string` | no |
| `AuthConfig` | `AuthConfig` | no |
| `SecretsManager` | `SecretsManager` | no |
| `OAuthCompleteRedirectUrl` | `string` | no |

## GetCredentialLocker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetCustomEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointAddress` | `string` | yes |

## GetDefaultEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationStatus` | `ConfigurationStatus` | yes |
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |

## GetDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DeliveryDestinationArn` | `string` | no |
| `DeliveryDestinationType` | `string` | no |
| `Name` | `string` | no |
| `RoleArn` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetDeviceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `DiscoveryType` | `string` | yes |
| `Status` | `string` | yes |
| `StartedAt` | `timestamp` | yes |
| `ControllerId` | `string` | no |
| `ConnectorAssociationId` | `string` | no |
| `AccountAssociationId` | `string` | no |
| `FinishedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetEventLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceId` | `string` | no |
| `EventLogLevel` | `string` | no |

## GetHubConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubTokenTimerExpirySettingInSeconds` | `long` | no |
| `UpdatedAt` | `timestamp` | no |

## GetManagedThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Owner` | `string` | no |
| `CredentialLockerId` | `string` | no |
| `AdvertisedProductId` | `string` | no |
| `Role` | `string` | no |
| `ProvisioningStatus` | `string` | no |
| `Name` | `string` | no |
| `Model` | `string` | no |
| `Brand` | `string` | no |
| `SerialNumber` | `string` | no |
| `UniversalProductCode` | `string` | no |
| `InternationalArticleNumber` | `string` | no |
| `ConnectorPolicyId` | `string` | no |
| `ConnectorDestinationId` | `string` | no |
| `ConnectorDeviceId` | `string` | no |
| `DeviceSpecificKey` | `string` | no |
| `MacAddress` | `string` | no |
| `ParentControllerId` | `string` | no |
| `Classification` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `ActivatedAt` | `timestamp` | no |
| `HubNetworkMode` | `string` | no |
| `MetaData` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |
| `WiFiSimpleSetupConfiguration` | `WiFiSimpleSetupConfiguration` | no |

## GetManagedThingCapabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `Capabilities` | `string` | no |
| `CapabilityReport` | `CapabilityReport` | no |

## GetManagedThingCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `CertificatePem` | `string` | no |

## GetManagedThingConnectivityData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `Connected` | `boolean` | no |
| `Timestamp` | `timestamp` | no |
| `DisconnectReason` | `string` | no |

## GetManagedThingMetaData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `MetaData` | `Map<string>` | no |

## GetManagedThingState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<StateEndpoint>` | yes |

## GetNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | no |
| `DestinationName` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetOtaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |
| `TaskArn` | `string` | no |
| `Description` | `string` | no |
| `S3Url` | `string` | no |
| `Protocol` | `string` | no |
| `OtaType` | `string` | no |
| `OtaTargetQueryString` | `string` | no |
| `OtaMechanism` | `string` | no |
| `Target` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `TaskConfigurationId` | `string` | no |
| `TaskProcessingDetails` | `TaskProcessingDetails` | no |
| `OtaSchedulingConfig` | `OtaTaskSchedulingConfig` | no |
| `OtaTaskExecutionRetryConfig` | `OtaTaskExecutionRetryConfig` | no |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetOtaTaskConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskConfigurationId` | `string` | no |
| `Name` | `string` | no |
| `PushConfig` | `PushConfig` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## GetProvisioningProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `ProvisioningType` | `string` | no |
| `Id` | `string` | no |
| `Status` | `string` | no |
| `ClaimCertificate` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetRuntimeLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `RuntimeLogConfigurations` | `RuntimeLogConfigurations` | no |

## GetSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `SchemaVersionedId` | `string` | yes |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaId` | `string` | no |
| `Type` | `string` | no |
| `Description` | `string` | no |
| `Namespace` | `string` | no |
| `SemanticVersion` | `string` | no |
| `Visibility` | `string` | no |
| `Schema` | `SchemaVersionSchema` | no |

## ListAccountAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDestinationId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<AccountAssociationItem>` | no |
| `NextToken` | `string` | no |

## ListCloudConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `LambdaArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ConnectorItem>` | no |
| `NextToken` | `string` | no |

## ListConnectorDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudConnectorId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDestinationList` | `List<ConnectorDestinationSummary>` | no |
| `NextToken` | `string` | no |

## ListCredentialLockers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<CredentialLockerSummary>` | no |
| `NextToken` | `string` | no |

## ListDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationList` | `List<DestinationSummary>` | no |
| `NextToken` | `string` | no |

## ListDeviceDiscoveries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `TypeFilter` | `string` | no |
| `StatusFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DeviceDiscoverySummary>` | no |
| `NextToken` | `string` | no |

## ListDiscoveredDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DiscoveredDeviceSummary>` | no |
| `NextToken` | `string` | no |

## ListEventLogConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventLogConfigurationList` | `List<EventLogConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListManagedThingAccountAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | no |
| `AccountAssociationId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ManagedThingAssociation>` | no |
| `NextToken` | `string` | no |

## ListManagedThingSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `EndpointIdFilter` | `string` | no |
| `CapabilityIdFilter` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ManagedThingSchemaListItem>` | no |
| `NextToken` | `string` | no |

## ListManagedThings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OwnerFilter` | `string` | no |
| `CredentialLockerFilter` | `string` | no |
| `RoleFilter` | `string` | no |
| `ParentControllerIdentifierFilter` | `string` | no |
| `ConnectorPolicyIdFilter` | `string` | no |
| `ConnectorDestinationIdFilter` | `string` | no |
| `ConnectorDeviceIdFilter` | `string` | no |
| `SerialNumberFilter` | `string` | no |
| `ProvisioningStatusFilter` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ManagedThingSummary>` | no |
| `NextToken` | `string` | no |

## ListNotificationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationConfigurationList` | `List<NotificationConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListOtaTaskConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<OtaTaskConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListOtaTaskExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionSummaries` | `List<OtaTaskExecutionSummaries>` | no |
| `NextToken` | `string` | no |

## ListOtaTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tasks` | `List<OtaTaskSummary>` | no |
| `NextToken` | `string` | no |

## ListProvisioningProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ProvisioningProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListSchemaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SchemaId` | `string` | no |
| `Namespace` | `string` | no |
| `Visibility` | `string` | no |
| `SemanticVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<SchemaVersionListItem>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutDefaultEncryptionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurationStatus` | `ConfigurationStatus` | yes |
| `encryptionType` | `string` | yes |
| `kmsKeyArn` | `string` | no |

## PutHubConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubTokenTimerExpirySettingInSeconds` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubTokenTimerExpirySettingInSeconds` | `long` | no |

## PutRuntimeLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |
| `RuntimeLogConfigurations` | `RuntimeLogConfigurations` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |
| `AccountAssociationId` | `string` | yes |
| `DeviceDiscoveryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | no |
| `DeviceDiscoveryId` | `string` | no |
| `ManagedThingId` | `string` | no |

## RegisterCustomEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointAddress` | `string` | yes |

## ResetRuntimeLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendConnectorEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |
| `UserId` | `string` | no |
| `Operation` | `string` | yes |
| `OperationVersion` | `string` | no |
| `StatusCode` | `integer` | no |
| `Message` | `string` | no |
| `DeviceDiscoveryId` | `string` | no |
| `ConnectorDeviceId` | `string` | no |
| `TraceId` | `string` | no |
| `Devices` | `List<Device>` | no |
| `MatterEndpoint` | `MatterEndpoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorId` | `string` | yes |

## SendManagedThingCommand

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedThingId` | `string` | yes |
| `Endpoints` | `List<CommandEndpoint>` | yes |
| `ConnectorAssociationId` | `string` | no |
| `AccountAssociationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TraceId` | `string` | no |

## StartAccountAssociationRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OAuthAuthorizationUrl` | `string` | yes |

## StartDeviceDiscovery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscoveryType` | `string` | yes |
| `CustomProtocolDetail` | `Map<string>` | no |
| `ControllerIdentifier` | `string` | no |
| `ConnectorAssociationIdentifier` | `string` | no |
| `AccountAssociationId` | `string` | no |
| `AuthenticationMaterial` | `string` | no |
| `AuthenticationMaterialType` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ConnectorDeviceIdList` | `List<string>` | no |
| `Protocol` | `string` | no |
| `EndDeviceIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `StartedAt` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateAccountAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAssociationId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCloudConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnectorDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Description` | `string` | no |
| `Name` | `string` | no |
| `AuthType` | `string` | no |
| `AuthConfig` | `AuthConfigUpdate` | no |
| `SecretsManager` | `SecretsManager` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DeliveryDestinationArn` | `string` | no |
| `DeliveryDestinationType` | `string` | no |
| `RoleArn` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventLogConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `EventLogLevel` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateManagedThing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Owner` | `string` | no |
| `CredentialLockerId` | `string` | no |
| `SerialNumber` | `string` | no |
| `WiFiSimpleSetupConfiguration` | `WiFiSimpleSetupConfiguration` | no |
| `Brand` | `string` | no |
| `Model` | `string` | no |
| `Name` | `string` | no |
| `CapabilityReport` | `CapabilityReport` | no |
| `CapabilitySchemas` | `List<CapabilitySchemaItem>` | no |
| `Capabilities` | `string` | no |
| `Classification` | `string` | no |
| `HubNetworkMode` | `string` | no |
| `MetaData` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventType` | `string` | yes |
| `DestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateOtaTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Description` | `string` | no |
| `TaskConfigurationId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


