# Amazon Cognito Identity Provider

API version: 2016-04-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cognito-idp/2016-04-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddCustomAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `CustomAttributes` | `List<SchemaAttributeType>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddUserPoolClientSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientSecret` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientSecretDescriptor` | `ClientSecretDescriptorType` | no |

## AdminAddUserToGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `GroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminConfirmSignUp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminCreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `UserAttributes` | `List<AttributeType>` | no |
| `ValidationData` | `List<AttributeType>` | no |
| `TemporaryPassword` | `string` | no |
| `ForceAliasCreation` | `boolean` | no |
| `MessageAction` | `string` | no |
| `DesiredDeliveryMediums` | `List<string>` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `UserType` | no |

## AdminDeleteSoftwareToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminDeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminDeleteUserAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `UserAttributeNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminDisableProviderForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `User` | `ProviderUserIdentifierType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminDisableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminEnableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminForgetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `DeviceKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminGetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceKey` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `DeviceType` | yes |

## AdminGetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `UserAttributes` | `List<AttributeType>` | no |
| `UserCreateDate` | `timestamp` | no |
| `UserLastModifiedDate` | `timestamp` | no |
| `Enabled` | `boolean` | no |
| `UserStatus` | `string` | no |
| `MFAOptions` | `List<MFAOptionType>` | no |
| `PreferredMfaSetting` | `string` | no |
| `UserMFASettingList` | `List<string>` | no |

## AdminGetUserAuthFactors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `PreferredMfaSetting` | `string` | no |
| `UserMFASettingList` | `List<string>` | no |
| `ConfiguredUserAuthFactors` | `List<string>` | no |

## AdminInitiateAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `AuthFlow` | `string` | yes |
| `AuthParameters` | `Map<string>` | no |
| `ClientMetadata` | `Map<string>` | no |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `ContextData` | `ContextDataType` | no |
| `Session` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeName` | `string` | no |
| `Session` | `string` | no |
| `ChallengeParameters` | `Map<string>` | no |
| `AuthenticationResult` | `AuthenticationResultType` | no |
| `AvailableChallenges` | `List<string>` | no |

## AdminLinkProviderForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `DestinationUser` | `ProviderUserIdentifierType` | yes |
| `SourceUser` | `ProviderUserIdentifierType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `Limit` | `integer` | no |
| `PaginationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Devices` | `List<DeviceType>` | no |
| `PaginationToken` | `string` | no |

## AdminListGroupsForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupType>` | no |
| `NextToken` | `string` | no |

## AdminListUserAuthEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthEvents` | `List<AuthEventType>` | no |
| `NextToken` | `string` | no |

## AdminRemoveUserFromGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `GroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminResetUserPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminRespondToAuthChallenge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `ChallengeName` | `string` | yes |
| `ChallengeResponses` | `Map<string>` | no |
| `Session` | `string` | no |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `ContextData` | `ContextDataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeName` | `string` | no |
| `Session` | `string` | no |
| `ChallengeParameters` | `Map<string>` | no |
| `AuthenticationResult` | `AuthenticationResultType` | no |

## AdminSetUserMFAPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSMfaSettings` | `SMSMfaSettingsType` | no |
| `SoftwareTokenMfaSettings` | `SoftwareTokenMfaSettingsType` | no |
| `EmailMfaSettings` | `EmailMfaSettingsType` | no |
| `WebAuthnMfaSettings` | `WebAuthnMfaSettingsType` | no |
| `Username` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminSetUserPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `Password` | `string` | yes |
| `Permanent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminSetUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `MFAOptions` | `List<MFAOptionType>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminUpdateAuthEventFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `EventId` | `string` | yes |
| `FeedbackValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminUpdateDeviceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `DeviceKey` | `string` | yes |
| `DeviceRememberedStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminUpdateUserAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `UserAttributes` | `List<AttributeType>` | yes |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AdminUserGlobalSignOut

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSoftwareToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | no |
| `Session` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretCode` | `string` | no |
| `Session` | `string` | no |

## ChangePassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PreviousPassword` | `string` | no |
| `ProposedPassword` | `string` | yes |
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CompleteWebAuthnRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `Credential` | `Document` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ConfirmDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `DeviceKey` | `string` | yes |
| `DeviceSecretVerifierConfig` | `DeviceSecretVerifierConfigType` | no |
| `DeviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserConfirmationNecessary` | `boolean` | no |

## ConfirmForgotPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `SecretHash` | `string` | no |
| `Username` | `string` | yes |
| `ConfirmationCode` | `string` | yes |
| `Password` | `string` | yes |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `UserContextData` | `UserContextDataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ConfirmSignUp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `SecretHash` | `string` | no |
| `Username` | `string` | yes |
| `ConfirmationCode` | `string` | yes |
| `ForceAliasCreation` | `boolean` | no |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `UserContextData` | `UserContextDataType` | no |
| `ClientMetadata` | `Map<string>` | no |
| `Session` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Session` | `string` | no |

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `Precedence` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `GroupType` | no |

## CreateIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ProviderName` | `string` | yes |
| `ProviderType` | `string` | yes |
| `ProviderDetails` | `Map<string>` | yes |
| `AttributeMapping` | `Map<string>` | no |
| `IdpIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProviderType` | yes |

## CreateManagedLoginBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `UseCognitoProvidedValues` | `boolean` | no |
| `Settings` | `Document` | no |
| `Assets` | `List<AssetType>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginBranding` | `ManagedLoginBrandingType` | no |

## CreateResourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Identifier` | `string` | yes |
| `Name` | `string` | yes |
| `Scopes` | `List<ResourceServerScopeType>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceServer` | `ResourceServerType` | yes |

## CreateTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `TermsName` | `string` | yes |
| `TermsSource` | `string` | yes |
| `Enforcement` | `string` | yes |
| `Links` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Terms` | `TermsType` | no |

## CreateUserImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `CloudWatchLogsRoleArn` | `string` | yes |
| `PasswordHashingAlgorithm` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserImportJob` | `UserImportJobType` | no |

## CreateUserPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `Policies` | `UserPoolPolicyType` | no |
| `DeletionProtection` | `string` | no |
| `LambdaConfig` | `LambdaConfigType` | no |
| `AutoVerifiedAttributes` | `List<string>` | no |
| `AliasAttributes` | `List<string>` | no |
| `UsernameAttributes` | `List<string>` | no |
| `SmsVerificationMessage` | `string` | no |
| `EmailVerificationMessage` | `string` | no |
| `EmailVerificationSubject` | `string` | no |
| `VerificationMessageTemplate` | `VerificationMessageTemplateType` | no |
| `SmsAuthenticationMessage` | `string` | no |
| `MfaConfiguration` | `string` | no |
| `UserAttributeUpdateSettings` | `UserAttributeUpdateSettingsType` | no |
| `DeviceConfiguration` | `DeviceConfigurationType` | no |
| `EmailConfiguration` | `EmailConfigurationType` | no |
| `SmsConfiguration` | `SmsConfigurationType` | no |
| `UserPoolTags` | `Map<string>` | no |
| `AdminCreateUserConfig` | `AdminCreateUserConfigType` | no |
| `Schema` | `List<SchemaAttributeType>` | no |
| `UserPoolAddOns` | `UserPoolAddOnsType` | no |
| `UsernameConfiguration` | `UsernameConfigurationType` | no |
| `AccountRecoverySetting` | `AccountRecoverySettingType` | no |
| `UserPoolTier` | `string` | no |
| `KeyConfiguration` | `KeyConfigurationType` | no |
| `IssuerConfiguration` | `IssuerConfigurationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPool` | `UserPoolType` | no |

## CreateUserPoolClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientName` | `string` | yes |
| `GenerateSecret` | `boolean` | no |
| `ClientSecret` | `string` | no |
| `RefreshTokenValidity` | `integer` | no |
| `AccessTokenValidity` | `integer` | no |
| `IdTokenValidity` | `integer` | no |
| `TokenValidityUnits` | `TokenValidityUnitsType` | no |
| `ReadAttributes` | `List<string>` | no |
| `WriteAttributes` | `List<string>` | no |
| `ExplicitAuthFlows` | `List<string>` | no |
| `SupportedIdentityProviders` | `List<string>` | no |
| `CallbackURLs` | `List<string>` | no |
| `LogoutURLs` | `List<string>` | no |
| `DefaultRedirectURI` | `string` | no |
| `AllowedOAuthFlows` | `List<string>` | no |
| `AllowedOAuthScopes` | `List<string>` | no |
| `AllowedOAuthFlowsUserPoolClient` | `boolean` | no |
| `AnalyticsConfiguration` | `AnalyticsConfigurationType` | no |
| `PreventUserExistenceErrors` | `string` | no |
| `EnableTokenRevocation` | `boolean` | no |
| `EnablePropagateAdditionalUserContextData` | `boolean` | no |
| `AuthSessionValidity` | `integer` | no |
| `RefreshTokenRotation` | `RefreshTokenRotationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolClient` | `UserPoolClientType` | no |

## CreateUserPoolDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `ManagedLoginVersion` | `integer` | no |
| `CustomDomainConfig` | `CustomDomainConfigType` | no |
| `Routing` | `RoutingType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginVersion` | `integer` | no |
| `CloudFrontDomain` | `string` | no |
| `Routing` | `RoutingType` | no |

## CreateUserPoolReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `RegionName` | `string` | yes |
| `UserPoolTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolReplica` | `UserPoolReplicaType` | no |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteManagedLoginBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginBrandingId` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TermsId` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserAttributeNames` | `List<string>` | yes |
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPoolClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPoolClientSecret

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientSecretId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPoolDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserPoolReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolReplica` | `UserPoolReplicaType` | no |

## DeleteWebAuthnCredential

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `CredentialId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProviderType` | yes |

## DescribeManagedLoginBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ManagedLoginBrandingId` | `string` | yes |
| `ReturnMergedResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginBranding` | `ManagedLoginBrandingType` | no |

## DescribeManagedLoginBrandingByClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `ReturnMergedResources` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginBranding` | `ManagedLoginBrandingType` | no |

## DescribeResourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceServer` | `ResourceServerType` | yes |

## DescribeRiskConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RiskConfiguration` | `RiskConfigurationType` | yes |

## DescribeTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TermsId` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Terms` | `TermsType` | no |

## DescribeTermsByClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `TermsName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Terms` | `TermsType` | no |

## DescribeUserImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserImportJob` | `UserImportJobType` | no |

## DescribeUserPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPool` | `UserPoolType` | no |

## DescribeUserPoolClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolClient` | `UserPoolClientType` | no |

## DescribeUserPoolDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainDescription` | `DomainDescriptionType` | no |

## ForgetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | no |
| `DeviceKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ForgotPassword

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `SecretHash` | `string` | no |
| `UserContextData` | `UserContextDataType` | no |
| `Username` | `string` | yes |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeDeliveryDetails` | `CodeDeliveryDetailsType` | no |

## GetCSVHeader

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | no |
| `CSVHeader` | `List<string>` | no |

## GetClientToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `Secret` | `string` | yes |
| `Scopes` | `List<string>` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientAuthenticationResult` | `ClientAuthenticationResultType` | no |

## GetDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceKey` | `string` | yes |
| `AccessToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Device` | `DeviceType` | yes |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `GroupType` | no |

## GetIdentityProviderByIdentifier

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `IdpIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProviderType` | yes |

## GetLogDeliveryConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogDeliveryConfiguration` | `LogDeliveryConfigurationType` | no |

## GetProvisionedLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LimitDefinition` | `LimitDefinitionType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `LimitType` | yes |

## GetSigningCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |

## GetTokensFromRefreshToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RefreshToken` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientSecret` | `string` | no |
| `DeviceKey` | `string` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationResult` | `AuthenticationResultType` | no |

## GetUICustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UICustomization` | `UICustomizationType` | yes |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `UserAttributes` | `List<AttributeType>` | yes |
| `MFAOptions` | `List<MFAOptionType>` | no |
| `PreferredMfaSetting` | `string` | no |
| `UserMFASettingList` | `List<string>` | no |

## GetUserAttributeVerificationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `AttributeName` | `string` | yes |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeDeliveryDetails` | `CodeDeliveryDetailsType` | no |

## GetUserAuthFactors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `PreferredMfaSetting` | `string` | no |
| `UserMFASettingList` | `List<string>` | no |
| `ConfiguredUserAuthFactors` | `List<string>` | no |

## GetUserPoolMfaConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SmsMfaConfiguration` | `SmsMfaConfigType` | no |
| `SoftwareTokenMfaConfiguration` | `SoftwareTokenMfaConfigType` | no |
| `EmailMfaConfiguration` | `EmailMfaConfigType` | no |
| `MfaConfiguration` | `string` | no |
| `WebAuthnConfiguration` | `WebAuthnConfigurationType` | no |

## GlobalSignOut

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## InitiateAuth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthFlow` | `string` | yes |
| `AuthParameters` | `Map<string>` | no |
| `ClientMetadata` | `Map<string>` | no |
| `ClientId` | `string` | yes |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `UserContextData` | `UserContextDataType` | no |
| `Session` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeName` | `string` | no |
| `Session` | `string` | no |
| `ChallengeParameters` | `Map<string>` | no |
| `AuthenticationResult` | `AuthenticationResultType` | no |
| `AvailableChallenges` | `List<string>` | no |

## ListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `Limit` | `integer` | no |
| `PaginationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Devices` | `List<DeviceType>` | no |
| `PaginationToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupType>` | no |
| `NextToken` | `string` | no |

## ListIdentityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Providers` | `List<ProviderDescription>` | yes |
| `NextToken` | `string` | no |

## ListResourceServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceServers` | `List<ResourceServerType>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Terms` | `List<TermsDescriptionType>` | yes |
| `NextToken` | `string` | no |

## ListUserImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `MaxResults` | `integer` | yes |
| `PaginationToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserImportJobs` | `List<UserImportJobType>` | no |
| `PaginationToken` | `string` | no |

## ListUserPoolClientSecrets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientSecrets` | `List<ClientSecretDescriptorType>` | no |
| `NextToken` | `string` | no |

## ListUserPoolClients

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolClients` | `List<UserPoolClientDescription>` | no |
| `NextToken` | `string` | no |

## ListUserPoolReplicas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolReplicas` | `List<UserPoolReplicaType>` | no |
| `NextToken` | `string` | no |

## ListUserPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPools` | `List<UserPoolDescriptionType>` | no |
| `NextToken` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `AttributesToGet` | `List<string>` | no |
| `Limit` | `integer` | no |
| `PaginationToken` | `string` | no |
| `Filter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<UserType>` | no |
| `PaginationToken` | `string` | no |

## ListUsersInGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `GroupName` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<UserType>` | no |
| `NextToken` | `string` | no |

## ListWebAuthnCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `List<WebAuthnCredentialDescription>` | yes |
| `NextToken` | `string` | no |

## ResendConfirmationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `SecretHash` | `string` | no |
| `UserContextData` | `UserContextDataType` | no |
| `Username` | `string` | yes |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeDeliveryDetails` | `CodeDeliveryDetailsType` | no |

## RespondToAuthChallenge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `ChallengeName` | `string` | yes |
| `Session` | `string` | no |
| `ChallengeResponses` | `Map<string>` | no |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `UserContextData` | `UserContextDataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChallengeName` | `string` | no |
| `Session` | `string` | no |
| `ChallengeParameters` | `Map<string>` | no |
| `AuthenticationResult` | `AuthenticationResultType` | no |

## RevokeToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Token` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientSecret` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetLogDeliveryConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `LogConfigurations` | `List<LogConfigurationType>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LogDeliveryConfiguration` | `LogDeliveryConfigurationType` | no |

## SetRiskConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | no |
| `CompromisedCredentialsRiskConfiguration` | `CompromisedCredentialsRiskConfigurationType` | no |
| `AccountTakeoverRiskConfiguration` | `AccountTakeoverRiskConfigurationType` | no |
| `RiskExceptionConfiguration` | `RiskExceptionConfigurationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RiskConfiguration` | `RiskConfigurationType` | yes |

## SetUICustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | no |
| `CSS` | `string` | no |
| `ImageFile` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UICustomization` | `UICustomizationType` | yes |

## SetUserMFAPreference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SMSMfaSettings` | `SMSMfaSettingsType` | no |
| `SoftwareTokenMfaSettings` | `SoftwareTokenMfaSettingsType` | no |
| `EmailMfaSettings` | `EmailMfaSettingsType` | no |
| `WebAuthnMfaSettings` | `WebAuthnMfaSettingsType` | no |
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetUserPoolMfaConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `SmsMfaConfiguration` | `SmsMfaConfigType` | no |
| `SoftwareTokenMfaConfiguration` | `SoftwareTokenMfaConfigType` | no |
| `EmailMfaConfiguration` | `EmailMfaConfigType` | no |
| `MfaConfiguration` | `string` | no |
| `WebAuthnConfiguration` | `WebAuthnConfigurationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SmsMfaConfiguration` | `SmsMfaConfigType` | no |
| `SoftwareTokenMfaConfiguration` | `SoftwareTokenMfaConfigType` | no |
| `EmailMfaConfiguration` | `EmailMfaConfigType` | no |
| `MfaConfiguration` | `string` | no |
| `WebAuthnConfiguration` | `WebAuthnConfigurationType` | no |

## SetUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `MFAOptions` | `List<MFAOptionType>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SignUp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientId` | `string` | yes |
| `SecretHash` | `string` | no |
| `Username` | `string` | yes |
| `Password` | `string` | no |
| `UserAttributes` | `List<AttributeType>` | no |
| `ValidationData` | `List<AttributeType>` | no |
| `AnalyticsMetadata` | `AnalyticsMetadataType` | no |
| `UserContextData` | `UserContextDataType` | no |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserConfirmed` | `boolean` | yes |
| `CodeDeliveryDetails` | `CodeDeliveryDetailsType` | no |
| `UserSub` | `string` | yes |
| `Session` | `string` | no |

## StartUserImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserImportJob` | `UserImportJobType` | no |

## StartWebAuthnRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CredentialCreationOptions` | `Document` | yes |

## StopUserImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserImportJob` | `UserImportJobType` | no |

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


## UpdateAuthEventFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Username` | `string` | yes |
| `EventId` | `string` | yes |
| `FeedbackToken` | `string` | yes |
| `FeedbackValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDeviceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `DeviceKey` | `string` | yes |
| `DeviceRememberedStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `Precedence` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `GroupType` | no |

## UpdateIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ProviderName` | `string` | yes |
| `ProviderDetails` | `Map<string>` | no |
| `AttributeMapping` | `Map<string>` | no |
| `IdpIdentifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProviderType` | yes |

## UpdateManagedLoginBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | no |
| `ManagedLoginBrandingId` | `string` | no |
| `UseCognitoProvidedValues` | `boolean` | no |
| `Settings` | `Document` | no |
| `Assets` | `List<AssetType>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginBranding` | `ManagedLoginBrandingType` | no |

## UpdateProvisionedLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LimitDefinition` | `LimitDefinitionType` | yes |
| `RequestedLimitValue` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `LimitType` | yes |

## UpdateResourceServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Identifier` | `string` | yes |
| `Name` | `string` | yes |
| `Scopes` | `List<ResourceServerScopeType>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceServer` | `ResourceServerType` | yes |

## UpdateTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TermsId` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `TermsName` | `string` | no |
| `TermsSource` | `string` | no |
| `Enforcement` | `string` | no |
| `Links` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Terms` | `TermsType` | no |

## UpdateUserAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserAttributes` | `List<AttributeType>` | yes |
| `AccessToken` | `string` | yes |
| `ClientMetadata` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeDeliveryDetailsList` | `List<CodeDeliveryDetailsType>` | no |

## UpdateUserPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `Policies` | `UserPoolPolicyType` | no |
| `DeletionProtection` | `string` | no |
| `LambdaConfig` | `LambdaConfigType` | no |
| `AutoVerifiedAttributes` | `List<string>` | no |
| `SmsVerificationMessage` | `string` | no |
| `EmailVerificationMessage` | `string` | no |
| `EmailVerificationSubject` | `string` | no |
| `VerificationMessageTemplate` | `VerificationMessageTemplateType` | no |
| `SmsAuthenticationMessage` | `string` | no |
| `UserAttributeUpdateSettings` | `UserAttributeUpdateSettingsType` | no |
| `MfaConfiguration` | `string` | no |
| `DeviceConfiguration` | `DeviceConfigurationType` | no |
| `EmailConfiguration` | `EmailConfigurationType` | no |
| `SmsConfiguration` | `SmsConfigurationType` | no |
| `UserPoolTags` | `Map<string>` | no |
| `AdminCreateUserConfig` | `AdminCreateUserConfigType` | no |
| `UserPoolAddOns` | `UserPoolAddOnsType` | no |
| `AccountRecoverySetting` | `AccountRecoverySettingType` | no |
| `PoolName` | `string` | no |
| `UserPoolTier` | `string` | no |
| `KeyConfiguration` | `KeyConfigurationType` | no |
| `IssuerConfiguration` | `IssuerConfigurationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserPoolClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientName` | `string` | no |
| `RefreshTokenValidity` | `integer` | no |
| `AccessTokenValidity` | `integer` | no |
| `IdTokenValidity` | `integer` | no |
| `TokenValidityUnits` | `TokenValidityUnitsType` | no |
| `ReadAttributes` | `List<string>` | no |
| `WriteAttributes` | `List<string>` | no |
| `ExplicitAuthFlows` | `List<string>` | no |
| `SupportedIdentityProviders` | `List<string>` | no |
| `CallbackURLs` | `List<string>` | no |
| `LogoutURLs` | `List<string>` | no |
| `DefaultRedirectURI` | `string` | no |
| `AllowedOAuthFlows` | `List<string>` | no |
| `AllowedOAuthScopes` | `List<string>` | no |
| `AllowedOAuthFlowsUserPoolClient` | `boolean` | no |
| `AnalyticsConfiguration` | `AnalyticsConfigurationType` | no |
| `PreventUserExistenceErrors` | `string` | no |
| `EnableTokenRevocation` | `boolean` | no |
| `EnablePropagateAdditionalUserContextData` | `boolean` | no |
| `AuthSessionValidity` | `integer` | no |
| `RefreshTokenRotation` | `RefreshTokenRotationType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolClient` | `UserPoolClientType` | no |

## UpdateUserPoolDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `UserPoolId` | `string` | yes |
| `ManagedLoginVersion` | `integer` | no |
| `CustomDomainConfig` | `CustomDomainConfigType` | no |
| `Routing` | `RoutingType` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedLoginVersion` | `integer` | no |
| `CloudFrontDomain` | `string` | no |
| `Routing` | `RoutingType` | no |

## UpdateUserPoolReplica

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolId` | `string` | yes |
| `RegionName` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserPoolReplica` | `UserPoolReplicaType` | no |

## VerifySoftwareToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | no |
| `Session` | `string` | no |
| `UserCode` | `string` | yes |
| `FriendlyDeviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `Session` | `string` | no |

## VerifyUserAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessToken` | `string` | yes |
| `AttributeName` | `string` | yes |
| `Code` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


