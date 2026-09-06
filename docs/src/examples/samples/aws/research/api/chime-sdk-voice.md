# Amazon Chime SDK Voice

API version: 2022-08-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime-sdk-voice/2022-08-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociatePhoneNumbersWithVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |
| `ForceAssociate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## AssociatePhoneNumbersWithVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroupId` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |
| `ForceAssociate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## BatchDeletePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## BatchUpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdatePhoneNumberRequestItems` | `List<UpdatePhoneNumberRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## CreatePhoneNumberOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductType` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrder` | `PhoneNumberOrder` | no |

## CreateProxySession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `ParticipantPhoneNumbers` | `List<string>` | yes |
| `Name` | `string` | no |
| `ExpiryMinutes` | `integer` | no |
| `Capabilities` | `List<string>` | yes |
| `NumberSelectionBehavior` | `string` | no |
| `GeoMatchLevel` | `string` | no |
| `GeoMatchParams` | `GeoMatchParams` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxySession` | `ProxySession` | no |

## CreateSipMediaApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsRegion` | `string` | yes |
| `Name` | `string` | yes |
| `Endpoints` | `List<SipMediaApplicationEndpoint>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplication` | `SipMediaApplication` | no |

## CreateSipMediaApplicationCall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromPhoneNumber` | `string` | yes |
| `ToPhoneNumber` | `string` | yes |
| `SipMediaApplicationId` | `string` | yes |
| `SipHeaders` | `Map<string>` | no |
| `ArgumentsMap` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationCall` | `SipMediaApplicationCall` | no |

## CreateSipRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `TriggerType` | `string` | yes |
| `TriggerValue` | `string` | yes |
| `Disabled` | `boolean` | no |
| `TargetApplications` | `List<SipRuleTargetApplication>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRule` | `SipRule` | no |

## CreateVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AwsRegion` | `string` | no |
| `RequireEncryption` | `boolean` | yes |
| `Tags` | `List<Tag>` | no |
| `IntegrationType` | `string` | no |
| `NetworkType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnector` | `VoiceConnector` | no |

## CreateVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `VoiceConnectorItems` | `List<VoiceConnectorItem>` | no |
| `CallDistributionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroup` | `VoiceConnectorGroup` | no |

## CreateVoiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfile` | `VoiceProfile` | no |

## CreateVoiceProfileDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ServerSideEncryptionConfiguration` | `ServerSideEncryptionConfiguration` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomain` | `VoiceProfileDomain` | no |

## DeletePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProxySession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `ProxySessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSipMediaApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSipRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorEmergencyCallingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorExternalSystemsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorOrigination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorStreamingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorTermination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceConnectorTerminationCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Usernames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVoiceProfileDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociatePhoneNumbersFromVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## DisassociatePhoneNumbersFromVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroupId` | `string` | yes |
| `E164PhoneNumbers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberErrors` | `List<PhoneNumberError>` | no |

## GetGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnector` | `VoiceConnectorSettings` | no |

## GetPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## GetPhoneNumberOrder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrder` | `PhoneNumberOrder` | no |

## GetPhoneNumberSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallingName` | `string` | no |
| `CallingNameUpdatedTimestamp` | `timestamp` | no |

## GetProxySession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `ProxySessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxySession` | `ProxySession` | no |

## GetSipMediaApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplication` | `SipMediaApplication` | no |

## GetSipMediaApplicationAlexaSkillConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationAlexaSkillConfiguration` | `SipMediaApplicationAlexaSkillConfiguration` | no |

## GetSipMediaApplicationLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationLoggingConfiguration` | `SipMediaApplicationLoggingConfiguration` | no |

## GetSipRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRule` | `SipRule` | no |

## GetSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpeakerSearchTask` | `SpeakerSearchTask` | no |

## GetVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnector` | `VoiceConnector` | no |

## GetVoiceConnectorEmergencyCallingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmergencyCallingConfiguration` | `EmergencyCallingConfiguration` | no |

## GetVoiceConnectorExternalSystemsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExternalSystemsConfiguration` | `ExternalSystemsConfiguration` | no |

## GetVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroup` | `VoiceConnectorGroup` | no |

## GetVoiceConnectorLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

## GetVoiceConnectorOrigination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Origination` | `Origination` | no |

## GetVoiceConnectorProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxy` | `Proxy` | no |

## GetVoiceConnectorStreamingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingConfiguration` | `StreamingConfiguration` | no |

## GetVoiceConnectorTermination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Termination` | `Termination` | no |

## GetVoiceConnectorTerminationHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminationHealth` | `TerminationHealth` | no |

## GetVoiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfile` | `VoiceProfile` | no |

## GetVoiceProfileDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomain` | `VoiceProfileDomain` | no |

## GetVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `VoiceToneAnalysisTaskId` | `string` | yes |
| `IsCaller` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceToneAnalysisTask` | `VoiceToneAnalysisTask` | no |

## ListAvailableVoiceConnectorRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorRegions` | `List<string>` | no |

## ListPhoneNumberOrders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberOrders` | `List<PhoneNumberOrder>` | no |
| `NextToken` | `string` | no |

## ListPhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ProductType` | `string` | no |
| `FilterName` | `string` | no |
| `FilterValue` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumbers` | `List<PhoneNumber>` | no |
| `NextToken` | `string` | no |

## ListProxySessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxySessions` | `List<ProxySession>` | no |
| `NextToken` | `string` | no |

## ListSipMediaApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplications` | `List<SipMediaApplication>` | no |
| `NextToken` | `string` | no |

## ListSipRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRules` | `List<SipRule>` | no |
| `NextToken` | `string` | no |

## ListSupportedPhoneNumberCountries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberCountries` | `List<PhoneNumberCountry>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListVoiceConnectorGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroups` | `List<VoiceConnectorGroup>` | no |
| `NextToken` | `string` | no |

## ListVoiceConnectorTerminationCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Usernames` | `List<string>` | no |

## ListVoiceConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectors` | `List<VoiceConnector>` | no |
| `NextToken` | `string` | no |

## ListVoiceProfileDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomains` | `List<VoiceProfileDomainSummary>` | no |
| `NextToken` | `string` | no |

## ListVoiceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomainId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfiles` | `List<VoiceProfileSummary>` | no |
| `NextToken` | `string` | no |

## PutSipMediaApplicationAlexaSkillConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |
| `SipMediaApplicationAlexaSkillConfiguration` | `SipMediaApplicationAlexaSkillConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationAlexaSkillConfiguration` | `SipMediaApplicationAlexaSkillConfiguration` | no |

## PutSipMediaApplicationLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |
| `SipMediaApplicationLoggingConfiguration` | `SipMediaApplicationLoggingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationLoggingConfiguration` | `SipMediaApplicationLoggingConfiguration` | no |

## PutVoiceConnectorEmergencyCallingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `EmergencyCallingConfiguration` | `EmergencyCallingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmergencyCallingConfiguration` | `EmergencyCallingConfiguration` | no |

## PutVoiceConnectorExternalSystemsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `SessionBorderControllerTypes` | `List<string>` | no |
| `ContactCenterSystemTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExternalSystemsConfiguration` | `ExternalSystemsConfiguration` | no |

## PutVoiceConnectorLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `LoggingConfiguration` | `LoggingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggingConfiguration` | `LoggingConfiguration` | no |

## PutVoiceConnectorOrigination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Origination` | `Origination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Origination` | `Origination` | no |

## PutVoiceConnectorProxy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `DefaultSessionExpiryMinutes` | `integer` | yes |
| `PhoneNumberPoolCountries` | `List<string>` | yes |
| `FallBackPhoneNumber` | `string` | no |
| `Disabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Proxy` | `Proxy` | no |

## PutVoiceConnectorStreamingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `StreamingConfiguration` | `StreamingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingConfiguration` | `StreamingConfiguration` | no |

## PutVoiceConnectorTermination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Termination` | `Termination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Termination` | `Termination` | no |

## PutVoiceConnectorTerminationCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Credentials` | `List<Credential>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestorePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## SearchAvailablePhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AreaCode` | `string` | no |
| `City` | `string` | no |
| `Country` | `string` | no |
| `State` | `string` | no |
| `TollFreePrefix` | `string` | no |
| `PhoneNumberType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `E164PhoneNumbers` | `List<string>` | no |
| `NextToken` | `string` | no |

## StartSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `TransactionId` | `string` | yes |
| `VoiceProfileDomainId` | `string` | yes |
| `ClientRequestToken` | `string` | no |
| `CallLeg` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpeakerSearchTask` | `SpeakerSearchTask` | no |

## StartVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `TransactionId` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceToneAnalysisTask` | `VoiceToneAnalysisTask` | no |

## StopSpeakerSearchTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopVoiceToneAnalysisTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `VoiceToneAnalysisTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateGlobalSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnector` | `VoiceConnectorSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `ProductType` | `string` | no |
| `CallingName` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `PhoneNumber` | no |

## UpdatePhoneNumberSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateProxySession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `ProxySessionId` | `string` | yes |
| `Capabilities` | `List<string>` | yes |
| `ExpiryMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProxySession` | `ProxySession` | no |

## UpdateSipMediaApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |
| `Name` | `string` | no |
| `Endpoints` | `List<SipMediaApplicationEndpoint>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplication` | `SipMediaApplication` | no |

## UpdateSipMediaApplicationCall

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationId` | `string` | yes |
| `TransactionId` | `string` | yes |
| `Arguments` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipMediaApplicationCall` | `SipMediaApplicationCall` | no |

## UpdateSipRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRuleId` | `string` | yes |
| `Name` | `string` | yes |
| `Disabled` | `boolean` | no |
| `TargetApplications` | `List<SipRuleTargetApplication>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SipRule` | `SipRule` | no |

## UpdateVoiceConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorId` | `string` | yes |
| `Name` | `string` | yes |
| `RequireEncryption` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnector` | `VoiceConnector` | no |

## UpdateVoiceConnectorGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroupId` | `string` | yes |
| `Name` | `string` | yes |
| `VoiceConnectorItems` | `List<VoiceConnectorItem>` | yes |
| `CallDistributionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceConnectorGroup` | `VoiceConnectorGroup` | no |

## UpdateVoiceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileId` | `string` | yes |
| `SpeakerSearchTaskId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfile` | `VoiceProfile` | no |

## UpdateVoiceProfileDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomainId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VoiceProfileDomain` | `VoiceProfileDomain` | no |

## ValidateE911Address

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AwsAccountId` | `string` | yes |
| `StreetNumber` | `string` | yes |
| `StreetInfo` | `string` | yes |
| `City` | `string` | yes |
| `State` | `string` | yes |
| `Country` | `string` | yes |
| `PostalCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ValidationResult` | `integer` | no |
| `AddressExternalId` | `string` | no |
| `Address` | `Address` | no |
| `CandidateAddressList` | `List<CandidateAddress>` | no |

