# Amazon Pinpoint SMS Voice V2

API version: 2022-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pinpoint-sms-voice-v2/2022-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateOriginationIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `OriginationIdentity` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `OriginationIdentityArn` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `IsoCountryCode` | `string` | no |

## AssociateProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | yes |
| `ConfigurationSetName` | `string` | yes |
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |

## CarrierLookup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `E164PhoneNumber` | `string` | yes |
| `DialingCountryCode` | `string` | no |
| `IsoCountryCode` | `string` | no |
| `Country` | `string` | no |
| `MCC` | `string` | no |
| `MNC` | `string` | no |
| `Carrier` | `string` | no |
| `PhoneNumberType` | `string` | yes |

## CreateConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | no |

## CreateEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |
| `MatchingEventTypes` | `List<string>` | yes |
| `CloudWatchLogsDestination` | `CloudWatchLogsDestination` | no |
| `KinesisFirehoseDestination` | `KinesisFirehoseDestination` | no |
| `SnsDestination` | `SnsDestination` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `EventDestination` | `EventDestination` | no |

## CreateNotifyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayName` | `string` | yes |
| `UseCase` | `string` | yes |
| `DefaultTemplateId` | `string` | no |
| `PoolId` | `string` | no |
| `EnabledCountries` | `List<string>` | no |
| `EnabledChannels` | `List<string>` | yes |
| `DeletionProtectionEnabled` | `boolean` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationArn` | `string` | yes |
| `NotifyConfigurationId` | `string` | yes |
| `DisplayName` | `string` | yes |
| `UseCase` | `string` | yes |
| `DefaultTemplateId` | `string` | no |
| `PoolId` | `string` | no |
| `EnabledCountries` | `List<string>` | no |
| `EnabledChannels` | `List<string>` | yes |
| `Tier` | `string` | yes |
| `TierUpgradeStatus` | `string` | yes |
| `Status` | `string` | yes |
| `RejectionReason` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## CreateOptOutList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListArn` | `string` | no |
| `OptOutListName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | no |

## CreatePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentity` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `MessageType` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `Status` | `string` | no |
| `MessageType` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `SharedRoutesEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | no |

## CreateProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `AccountDefault` | `boolean` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `Tags` | `List<Tag>` | no |

## CreateRcsAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionProtectionEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentArn` | `string` | yes |
| `RcsAgentId` | `string` | yes |
| `Status` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `OptOutListName` | `string` | no |
| `CreatedTimestamp` | `timestamp` | yes |
| `SelfManagedOptOutsEnabled` | `boolean` | yes |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `TwoWayEnabled` | `boolean` | yes |
| `TwoWayMediaS3BucketName` | `string` | no |
| `TwoWayMediaS3KeyPrefix` | `string` | no |
| `TwoWayMediaS3Role` | `string` | no |
| `TwoWayRcsEventsEnabled` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

## CreateRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationType` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `RegistrationType` | `string` | yes |
| `RegistrationStatus` | `string` | yes |
| `CurrentVersionNumber` | `long` | yes |
| `AdditionalAttributes` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## CreateRegistrationAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `RegistrationType` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `PhoneNumber` | `string` | no |

## CreateRegistrationAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachmentBody` | `blob` | no |
| `AttachmentUrl` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationAttachmentArn` | `string` | yes |
| `RegistrationAttachmentId` | `string` | yes |
| `AttachmentStatus` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## CreateRegistrationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `RegistrationVersionStatus` | `string` | yes |
| `RegistrationVersionStatusHistory` | `RegistrationVersionStatusHistory` | yes |

## CreateVerifiedDestinationNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPhoneNumber` | `string` | yes |
| `RcsAgentId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberArn` | `string` | yes |
| `VerifiedDestinationNumberId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `Status` | `string` | yes |
| `RcsAgentId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## DeleteAccountDefaultProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultProtectConfigurationArn` | `string` | yes |
| `DefaultProtectConfigurationId` | `string` | yes |

## DeleteConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `EventDestinations` | `List<EventDestination>` | no |
| `DefaultMessageType` | `string` | no |
| `DefaultSenderId` | `string` | no |
| `DefaultMessageFeedbackEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |

## DeleteDefaultMessageType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MessageType` | `string` | no |

## DeleteDefaultSenderId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `SenderId` | `string` | no |

## DeleteEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `EventDestination` | `EventDestination` | no |

## DeleteKeyword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentity` | `string` | yes |
| `Keyword` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentityArn` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `Keyword` | `string` | no |
| `KeywordMessage` | `string` | no |
| `KeywordAction` | `string` | no |

## DeleteMediaMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## DeleteNotifyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationArn` | `string` | yes |
| `NotifyConfigurationId` | `string` | yes |
| `DisplayName` | `string` | yes |
| `UseCase` | `string` | yes |
| `DefaultTemplateId` | `string` | no |
| `PoolId` | `string` | no |
| `EnabledCountries` | `List<string>` | no |
| `EnabledChannels` | `List<string>` | yes |
| `Tier` | `string` | yes |
| `TierUpgradeStatus` | `string` | yes |
| `Status` | `string` | yes |
| `RejectionReason` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `CreatedTimestamp` | `timestamp` | yes |

## DeleteNotifyMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## DeleteOptOutList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListArn` | `string` | no |
| `OptOutListName` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## DeleteOptedOutNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListName` | `string` | yes |
| `OptedOutNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListArn` | `string` | no |
| `OptOutListName` | `string` | no |
| `OptedOutNumber` | `string` | no |
| `OptedOutTimestamp` | `timestamp` | no |
| `EndUserOptedOut` | `boolean` | no |

## DeletePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `Status` | `string` | no |
| `MessageType` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `SharedRoutesEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |

## DeleteProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `AccountDefault` | `boolean` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |

## DeleteProtectConfigurationRuleSetNumberOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `Action` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `ExpirationTimestamp` | `timestamp` | no |

## DeleteRcsAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentArn` | `string` | yes |
| `RcsAgentId` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `OptOutListName` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | yes |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `TwoWayEnabled` | `boolean` | yes |
| `TwoWayRcsEventsEnabled` | `List<string>` | no |

## DeleteRcsMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## DeleteRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `RegistrationType` | `string` | yes |
| `RegistrationStatus` | `string` | yes |
| `CurrentVersionNumber` | `long` | yes |
| `ApprovedVersionNumber` | `long` | no |
| `LatestDeniedVersionNumber` | `long` | no |
| `AdditionalAttributes` | `Map<string>` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## DeleteRegistrationAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationAttachmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationAttachmentArn` | `string` | yes |
| `RegistrationAttachmentId` | `string` | yes |
| `AttachmentStatus` | `string` | yes |
| `AttachmentUploadErrorReason` | `string` | no |
| `CreatedTimestamp` | `timestamp` | yes |

## DeleteRegistrationFieldValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `FieldPath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `FieldPath` | `string` | yes |
| `SelectChoices` | `List<string>` | no |
| `TextValue` | `string` | no |
| `RegistrationAttachmentId` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## DeleteTextMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## DeleteVerifiedDestinationNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberArn` | `string` | yes |
| `VerifiedDestinationNumberId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |

## DeleteVoiceMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountAttributes` | `List<AccountAttribute>` | no |
| `NextToken` | `string` | no |

## DescribeAccountLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLimits` | `List<AccountLimit>` | no |
| `NextToken` | `string` | no |

## DescribeConfigurationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetNames` | `List<string>` | no |
| `Filters` | `List<ConfigurationSetFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSets` | `List<ConfigurationSetInformation>` | no |
| `NextToken` | `string` | no |

## DescribeKeywords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentity` | `string` | yes |
| `Keywords` | `List<string>` | no |
| `Filters` | `List<KeywordFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentityArn` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `Keywords` | `List<KeywordInformation>` | no |
| `NextToken` | `string` | no |

## DescribeNotifyConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationIds` | `List<string>` | no |
| `Filters` | `List<NotifyConfigurationFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurations` | `List<NotifyConfigurationInformation>` | no |
| `NextToken` | `string` | no |

## DescribeNotifyTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateIds` | `List<string>` | no |
| `Filters` | `List<NotifyTemplateFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyTemplates` | `List<NotifyTemplateInformation>` | no |
| `NextToken` | `string` | no |

## DescribeOptOutLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListNames` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Owner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutLists` | `List<OptOutListInformation>` | no |
| `NextToken` | `string` | no |

## DescribeOptedOutNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListName` | `string` | yes |
| `OptedOutNumbers` | `List<string>` | no |
| `Filters` | `List<OptedOutFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListArn` | `string` | no |
| `OptOutListName` | `string` | no |
| `OptedOutNumbers` | `List<OptedOutNumberInformation>` | no |
| `NextToken` | `string` | no |

## DescribePhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberIds` | `List<string>` | no |
| `Filters` | `List<PhoneNumberFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Owner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumbers` | `List<PhoneNumberInformation>` | no |
| `NextToken` | `string` | no |

## DescribePools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolIds` | `List<string>` | no |
| `Filters` | `List<PoolFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Owner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pools` | `List<PoolInformation>` | no |
| `NextToken` | `string` | no |

## DescribeProtectConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationIds` | `List<string>` | no |
| `Filters` | `List<ProtectConfigurationFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurations` | `List<ProtectConfigurationInformation>` | no |
| `NextToken` | `string` | no |

## DescribeRcsAgentCountryLaunchStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentId` | `string` | yes |
| `IsoCountryCodes` | `List<string>` | no |
| `Filters` | `List<CountryLaunchStatusFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentId` | `string` | yes |
| `RcsAgentArn` | `string` | yes |
| `CountryLaunchStatus` | `List<CountryLaunchStatusInformation>` | no |
| `NextToken` | `string` | no |

## DescribeRcsAgents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentIds` | `List<string>` | no |
| `Owner` | `string` | no |
| `Filters` | `List<RcsAgentFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgents` | `List<RcsAgentInformation>` | no |
| `NextToken` | `string` | no |

## DescribeRegistrationAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationAttachmentIds` | `List<string>` | no |
| `Filters` | `List<RegistrationAttachmentFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationAttachments` | `List<RegistrationAttachmentsInformation>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrationFieldDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationType` | `string` | yes |
| `SectionPath` | `string` | no |
| `FieldPaths` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationType` | `string` | yes |
| `RegistrationFieldDefinitions` | `List<RegistrationFieldDefinition>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrationFieldValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | no |
| `SectionPath` | `string` | no |
| `FieldPaths` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `RegistrationFieldValues` | `List<RegistrationFieldValueInformation>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrationSectionDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationType` | `string` | yes |
| `SectionPaths` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationType` | `string` | yes |
| `RegistrationSectionDefinitions` | `List<RegistrationSectionDefinition>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrationTypeDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationTypes` | `List<string>` | no |
| `Filters` | `List<RegistrationTypeFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationTypeDefinitions` | `List<RegistrationTypeDefinition>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `VersionNumbers` | `List<long>` | no |
| `Filters` | `List<RegistrationVersionFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `RegistrationVersions` | `List<RegistrationVersionInformation>` | yes |
| `NextToken` | `string` | no |

## DescribeRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationIds` | `List<string>` | no |
| `Filters` | `List<RegistrationFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Registrations` | `List<RegistrationInformation>` | yes |
| `NextToken` | `string` | no |

## DescribeSenderIds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderIds` | `List<SenderIdAndCountry>` | no |
| `Filters` | `List<SenderIdFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Owner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderIds` | `List<SenderIdInformation>` | no |
| `NextToken` | `string` | no |

## DescribeSpendLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpendLimits` | `List<SpendLimit>` | no |
| `NextToken` | `string` | no |

## DescribeVerifiedDestinationNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberIds` | `List<string>` | no |
| `DestinationPhoneNumbers` | `List<string>` | no |
| `Filters` | `List<VerifiedDestinationNumberFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumbers` | `List<VerifiedDestinationNumberInformation>` | yes |
| `NextToken` | `string` | no |

## DisassociateOriginationIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `OriginationIdentity` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `OriginationIdentityArn` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `IsoCountryCode` | `string` | no |

## DisassociateProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | yes |
| `ConfigurationSetName` | `string` | yes |
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |

## DiscardRegistrationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `RegistrationVersionStatus` | `string` | yes |
| `RegistrationVersionStatusHistory` | `RegistrationVersionStatusHistory` | yes |

## GetProtectConfigurationCountryRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `NumberCapability` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `NumberCapability` | `string` | yes |
| `CountryRuleSet` | `Map<ProtectConfigurationCountryRuleSetInformation>` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## ListNotifyCountries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channels` | `List<string>` | no |
| `UseCases` | `List<string>` | no |
| `Tier` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyCountries` | `List<NotifyCountryInformation>` | no |
| `NextToken` | `string` | no |

## ListPoolOriginationIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `Filters` | `List<PoolOriginationIdentitiesFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `OriginationIdentities` | `List<OriginationIdentityMetadata>` | no |
| `NextToken` | `string` | no |

## ListProtectConfigurationRuleSetNumberOverrides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `Filters` | `List<ProtectConfigurationRuleSetNumberOverrideFilterItem>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `RuleSetNumberOverrides` | `List<ProtectConfigurationRuleSetNumberOverride>` | no |
| `NextToken` | `string` | no |

## ListRegistrationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `Filters` | `List<RegistrationAssociationFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `RegistrationType` | `string` | yes |
| `RegistrationAssociations` | `List<RegistrationAssociationMetadata>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## PutKeyword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentity` | `string` | yes |
| `Keyword` | `string` | yes |
| `KeywordMessage` | `string` | yes |
| `KeywordAction` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginationIdentityArn` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `Keyword` | `string` | no |
| `KeywordMessage` | `string` | no |
| `KeywordAction` | `string` | no |

## PutMessageFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |
| `MessageFeedbackStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |
| `MessageFeedbackStatus` | `string` | yes |

## PutOptedOutNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListName` | `string` | yes |
| `OptedOutNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptOutListArn` | `string` | no |
| `OptOutListName` | `string` | no |
| `OptedOutNumber` | `string` | no |
| `OptedOutTimestamp` | `timestamp` | no |
| `EndUserOptedOut` | `boolean` | no |

## PutProtectConfigurationRuleSetNumberOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ProtectConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `Action` | `string` | yes |
| `ExpirationTimestamp` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `Action` | `string` | yes |
| `IsoCountryCode` | `string` | no |
| `ExpirationTimestamp` | `timestamp` | no |

## PutRegistrationFieldValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `FieldPath` | `string` | yes |
| `SelectChoices` | `List<string>` | no |
| `TextValue` | `string` | no |
| `RegistrationAttachmentId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `FieldPath` | `string` | yes |
| `SelectChoices` | `List<string>` | no |
| `TextValue` | `string` | no |
| `RegistrationAttachmentId` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## ReleasePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberArn` | `string` | no |
| `PhoneNumberId` | `string` | no |
| `PhoneNumber` | `string` | no |
| `Status` | `string` | no |
| `IsoCountryCode` | `string` | no |
| `MessageType` | `string` | no |
| `NumberCapabilities` | `List<string>` | no |
| `NumberType` | `string` | no |
| `MonthlyLeasingPrice` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `RegistrationId` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## ReleaseSenderId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderIdArn` | `string` | yes |
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |
| `MessageTypes` | `List<string>` | yes |
| `MonthlyLeasingPrice` | `string` | yes |
| `Registered` | `boolean` | yes |
| `RegistrationId` | `string` | no |

## RequestPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IsoCountryCode` | `string` | yes |
| `MessageType` | `string` | yes |
| `NumberCapabilities` | `List<string>` | yes |
| `NumberType` | `string` | yes |
| `OptOutListName` | `string` | no |
| `PoolId` | `string` | no |
| `RegistrationId` | `string` | no |
| `InternationalSendingEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberArn` | `string` | no |
| `PhoneNumberId` | `string` | no |
| `PhoneNumber` | `string` | no |
| `Status` | `string` | no |
| `IsoCountryCode` | `string` | no |
| `MessageType` | `string` | no |
| `NumberCapabilities` | `List<string>` | no |
| `NumberType` | `string` | no |
| `MonthlyLeasingPrice` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `InternationalSendingEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `PoolId` | `string` | no |
| `RegistrationId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `CreatedTimestamp` | `timestamp` | no |

## RequestSenderId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |
| `MessageTypes` | `List<string>` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderIdArn` | `string` | yes |
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |
| `MessageTypes` | `List<string>` | yes |
| `MonthlyLeasingPrice` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `Registered` | `boolean` | yes |
| `Tags` | `List<Tag>` | no |

## SendDestinationNumberVerificationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberId` | `string` | yes |
| `VerificationChannel` | `string` | yes |
| `LanguageCode` | `string` | no |
| `OriginationIdentity` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `Context` | `Map<string>` | no |
| `DestinationCountryParameters` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |

## SendMediaMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPhoneNumber` | `string` | yes |
| `OriginationIdentity` | `string` | yes |
| `MessageBody` | `string` | no |
| `MediaUrls` | `List<string>` | no |
| `ConfigurationSetName` | `string` | no |
| `MaxPrice` | `string` | no |
| `TimeToLive` | `integer` | no |
| `Context` | `Map<string>` | no |
| `DryRun` | `boolean` | no |
| `ProtectConfigurationId` | `string` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendNotifyTextMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `TemplateId` | `string` | no |
| `TemplateVariables` | `Map<string>` | yes |
| `TimeToLive` | `integer` | no |
| `Context` | `Map<string>` | no |
| `ConfigurationSetName` | `string` | no |
| `DryRun` | `boolean` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |
| `TemplateId` | `string` | no |
| `ResolvedMessageBody` | `string` | no |

## SendNotifyVoiceMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `TemplateId` | `string` | no |
| `TemplateVariables` | `Map<string>` | yes |
| `VoiceId` | `string` | no |
| `TimeToLive` | `integer` | no |
| `Context` | `Map<string>` | no |
| `ConfigurationSetName` | `string` | no |
| `DryRun` | `boolean` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |
| `TemplateId` | `string` | no |
| `ResolvedMessageBody` | `string` | no |

## SendRcsMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPhoneNumber` | `string` | yes |
| `OriginationIdentity` | `string` | yes |
| `RcsMessageContent` | `RcsMessageContent` | no |
| `TimeToLive` | `integer` | no |
| `MessageTrafficType` | `string` | no |
| `FallbackConfiguration` | `RcsFallbackConfiguration` | no |
| `ProtectConfigurationId` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MaxPrice` | `string` | no |
| `DryRun` | `boolean` | no |
| `Context` | `Map<string>` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendTextMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPhoneNumber` | `string` | yes |
| `OriginationIdentity` | `string` | no |
| `MessageBody` | `string` | no |
| `MessageType` | `string` | no |
| `Keyword` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MaxPrice` | `string` | no |
| `TimeToLive` | `integer` | no |
| `Context` | `Map<string>` | no |
| `DestinationCountryParameters` | `Map<string>` | no |
| `DryRun` | `boolean` | no |
| `ProtectConfigurationId` | `string` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendVoiceMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationPhoneNumber` | `string` | yes |
| `OriginationIdentity` | `string` | yes |
| `MessageBody` | `string` | no |
| `MessageBodyTextType` | `string` | no |
| `VoiceId` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MaxPricePerMinute` | `string` | no |
| `TimeToLive` | `integer` | no |
| `Context` | `Map<string>` | no |
| `DryRun` | `boolean` | no |
| `ProtectConfigurationId` | `string` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SetAccountDefaultProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultProtectConfigurationArn` | `string` | yes |
| `DefaultProtectConfigurationId` | `string` | yes |

## SetDefaultMessageFeedbackEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `MessageFeedbackEnabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MessageFeedbackEnabled` | `boolean` | no |

## SetDefaultMessageType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `MessageType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `MessageType` | `string` | no |

## SetDefaultSenderId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `SenderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `SenderId` | `string` | no |

## SetMediaMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## SetNotifyMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## SetRcsMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## SetTextMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## SetVoiceMessageSpendLimitOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonthlyLimit` | `long` | no |

## SubmitRegistrationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |
| `AwsReview` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationArn` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `RegistrationVersionStatus` | `string` | yes |
| `RegistrationVersionStatusHistory` | `RegistrationVersionStatusHistory` | yes |
| `AwsReview` | `boolean` | yes |

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


## UpdateEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |
| `Enabled` | `boolean` | no |
| `MatchingEventTypes` | `List<string>` | no |
| `CloudWatchLogsDestination` | `CloudWatchLogsDestination` | no |
| `KinesisFirehoseDestination` | `KinesisFirehoseDestination` | no |
| `SnsDestination` | `SnsDestination` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `EventDestination` | `EventDestination` | no |

## UpdateNotifyConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationId` | `string` | yes |
| `DefaultTemplateId` | `string` | no |
| `PoolId` | `string` | no |
| `EnabledCountries` | `List<string>` | no |
| `EnabledChannels` | `List<string>` | no |
| `DeletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotifyConfigurationArn` | `string` | yes |
| `NotifyConfigurationId` | `string` | yes |
| `DisplayName` | `string` | yes |
| `UseCase` | `string` | yes |
| `DefaultTemplateId` | `string` | no |
| `PoolId` | `string` | no |
| `EnabledCountries` | `List<string>` | no |
| `EnabledChannels` | `List<string>` | yes |
| `Tier` | `string` | yes |
| `TierUpgradeStatus` | `string` | yes |
| `Status` | `string` | yes |
| `RejectionReason` | `string` | no |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `CreatedTimestamp` | `timestamp` | yes |

## UpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `InternationalSendingEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberArn` | `string` | no |
| `PhoneNumberId` | `string` | no |
| `PhoneNumber` | `string` | no |
| `Status` | `string` | no |
| `IsoCountryCode` | `string` | no |
| `MessageType` | `string` | no |
| `NumberCapabilities` | `List<string>` | no |
| `NumberType` | `string` | no |
| `MonthlyLeasingPrice` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `InternationalSendingEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `RegistrationId` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |

## UpdatePool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `SharedRoutesEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolArn` | `string` | no |
| `PoolId` | `string` | no |
| `Status` | `string` | no |
| `MessageType` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `SharedRoutesEnabled` | `boolean` | no |
| `DeletionProtectionEnabled` | `boolean` | no |
| `CreatedTimestamp` | `timestamp` | no |

## UpdateProtectConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `AccountDefault` | `boolean` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |

## UpdateProtectConfigurationCountryRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationId` | `string` | yes |
| `NumberCapability` | `string` | yes |
| `CountryRuleSetUpdates` | `Map<ProtectConfigurationCountryRuleSetInformation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProtectConfigurationArn` | `string` | yes |
| `ProtectConfigurationId` | `string` | yes |
| `NumberCapability` | `string` | yes |
| `CountryRuleSet` | `Map<ProtectConfigurationCountryRuleSetInformation>` | yes |

## UpdateRcsAgent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentId` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | no |
| `OptOutListName` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | no |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `TwoWayEnabled` | `boolean` | no |
| `TwoWayMediaS3BucketName` | `string` | no |
| `TwoWayMediaS3KeyPrefix` | `string` | no |
| `TwoWayMediaS3Role` | `string` | no |
| `TwoWayRcsEventsEnabled` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RcsAgentArn` | `string` | yes |
| `RcsAgentId` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `OptOutListName` | `string` | no |
| `SelfManagedOptOutsEnabled` | `boolean` | yes |
| `TwoWayChannelArn` | `string` | no |
| `TwoWayChannelRole` | `string` | no |
| `TwoWayEnabled` | `boolean` | yes |
| `TwoWayMediaS3BucketName` | `string` | no |
| `TwoWayMediaS3KeyPrefix` | `string` | no |
| `TwoWayMediaS3Role` | `string` | no |
| `TwoWayRcsEventsEnabled` | `List<string>` | no |

## UpdateSenderId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SenderIdArn` | `string` | yes |
| `SenderId` | `string` | yes |
| `IsoCountryCode` | `string` | yes |
| `MessageTypes` | `List<string>` | yes |
| `MonthlyLeasingPrice` | `string` | yes |
| `DeletionProtectionEnabled` | `boolean` | yes |
| `Registered` | `boolean` | yes |
| `RegistrationId` | `string` | no |

## VerifyDestinationNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberId` | `string` | yes |
| `VerificationCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedDestinationNumberArn` | `string` | yes |
| `VerifiedDestinationNumberId` | `string` | yes |
| `DestinationPhoneNumber` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedTimestamp` | `timestamp` | yes |

