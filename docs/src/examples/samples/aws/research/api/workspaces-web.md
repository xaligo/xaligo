# Amazon WorkSpaces Web

API version: 2020-07-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workspaces-web/2020-07-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `browserSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `browserSettingsArn` | `string` | yes |

## AssociateDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `dataProtectionSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `dataProtectionSettingsArn` | `string` | yes |

## AssociateIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `ipAccessSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `ipAccessSettingsArn` | `string` | yes |

## AssociateNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `networkSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `networkSettingsArn` | `string` | yes |

## AssociateSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `sessionLoggerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `sessionLoggerArn` | `string` | yes |

## AssociateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `trustStoreArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `trustStoreArn` | `string` | yes |

## AssociateUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `userAccessLoggingSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `userAccessLoggingSettingsArn` | `string` | yes |

## AssociateUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `userSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `userSettingsArn` | `string` | yes |

## CreateBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `browserPolicy` | `string` | no |
| `clientToken` | `string` | no |
| `webContentFilteringPolicy` | `WebContentFilteringPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettingsArn` | `string` | yes |

## CreateDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `inlineRedactionConfiguration` | `InlineRedactionConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettingsArn` | `string` | yes |

## CreateIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `identityProviderName` | `string` | yes |
| `identityProviderType` | `string` | yes |
| `identityProviderDetails` | `Map<string>` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderArn` | `string` | yes |

## CreateIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `ipRules` | `List<IpRule>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettingsArn` | `string` | yes |

## CreateNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettingsArn` | `string` | yes |

## CreatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |
| `tags` | `List<Tag>` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `clientToken` | `string` | no |
| `authenticationType` | `string` | no |
| `instanceType` | `string` | no |
| `maxConcurrentSessions` | `integer` | no |
| `portalCustomDomain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `portalEndpoint` | `string` | yes |

## CreateSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventFilter` | `EventFilter` | yes |
| `logConfiguration` | `LogConfiguration` | yes |
| `displayName` | `string` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLoggerArn` | `string` | yes |

## CreateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateList` | `List<blob>` | yes |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |

## CreateUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kinesisStreamArn` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettingsArn` | `string` | yes |

## CreateUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `copyAllowed` | `string` | yes |
| `pasteAllowed` | `string` | yes |
| `downloadAllowed` | `string` | yes |
| `uploadAllowed` | `string` | yes |
| `printAllowed` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `disconnectTimeoutInMinutes` | `integer` | no |
| `idleDisconnectTimeoutInMinutes` | `integer` | no |
| `clientToken` | `string` | no |
| `cookieSynchronizationConfiguration` | `CookieSynchronizationConfiguration` | no |
| `customerManagedKey` | `string` | no |
| `additionalEncryptionContext` | `Map<string>` | no |
| `deepLinkAllowed` | `string` | no |
| `toolbarConfiguration` | `ToolbarConfiguration` | no |
| `brandingConfigurationInput` | `BrandingConfigurationCreateInput` | no |
| `webAuthnAllowed` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettingsArn` | `string` | yes |

## DeleteBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLoggerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExpireSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettings` | `BrowserSettings` | no |

## GetDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettings` | `DataProtectionSettings` | no |

## GetIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProvider` | `IdentityProvider` | no |

## GetIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettings` | `IpAccessSettings` | no |

## GetNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettings` | `NetworkSettings` | no |

## GetPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portal` | `Portal` | no |

## GetPortalServiceProviderMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `serviceProviderSamlMetadata` | `string` | no |

## GetSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `sessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `session` | `Session` | no |

## GetSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLoggerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLogger` | `SessionLogger` | no |

## GetTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStore` | `TrustStore` | no |

## GetTrustStoreCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |
| `thumbprint` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |
| `certificate` | `Certificate` | no |

## GetUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettings` | `UserAccessLoggingSettings` | no |

## GetUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettingsArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettings` | `UserSettings` | no |

## ListBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettings` | `List<BrowserSettingsSummary>` | no |
| `nextToken` | `string` | no |

## ListDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettings` | `List<DataProtectionSettingsSummary>` | no |
| `nextToken` | `string` | no |

## ListIdentityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `portalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `identityProviders` | `List<IdentityProviderSummary>` | no |

## ListIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettings` | `List<IpAccessSettingsSummary>` | no |
| `nextToken` | `string` | no |

## ListNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettings` | `List<NetworkSettingsSummary>` | no |
| `nextToken` | `string` | no |

## ListPortals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portals` | `List<PortalSummary>` | no |
| `nextToken` | `string` | no |

## ListSessionLoggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLoggers` | `List<SessionLoggerSummary>` | no |
| `nextToken` | `string` | no |

## ListSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalId` | `string` | yes |
| `username` | `string` | no |
| `sessionId` | `string` | no |
| `sortBy` | `string` | no |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessions` | `List<SessionSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListTrustStoreCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateList` | `List<CertificateSummary>` | no |
| `trustStoreArn` | `string` | yes |
| `nextToken` | `string` | no |

## ListTrustStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStores` | `List<TrustStoreSummary>` | no |
| `nextToken` | `string` | no |

## ListUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettings` | `List<UserAccessLoggingSettingsSummary>` | no |
| `nextToken` | `string` | no |

## ListUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettings` | `List<UserSettingsSummary>` | no |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBrowserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettingsArn` | `string` | yes |
| `browserPolicy` | `string` | no |
| `clientToken` | `string` | no |
| `webContentFilteringPolicy` | `WebContentFilteringPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSettings` | `BrowserSettings` | yes |

## UpdateDataProtectionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettingsArn` | `string` | yes |
| `inlineRedactionConfiguration` | `InlineRedactionConfiguration` | no |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataProtectionSettings` | `DataProtectionSettings` | yes |

## UpdateIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderArn` | `string` | yes |
| `identityProviderName` | `string` | no |
| `identityProviderType` | `string` | no |
| `identityProviderDetails` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProvider` | `IdentityProvider` | yes |

## UpdateIpAccessSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettingsArn` | `string` | yes |
| `displayName` | `string` | no |
| `description` | `string` | no |
| `ipRules` | `List<IpRule>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ipAccessSettings` | `IpAccessSettings` | yes |

## UpdateNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettingsArn` | `string` | yes |
| `vpcId` | `string` | no |
| `subnetIds` | `List<string>` | no |
| `securityGroupIds` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkSettings` | `NetworkSettings` | yes |

## UpdatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portalArn` | `string` | yes |
| `displayName` | `string` | no |
| `authenticationType` | `string` | no |
| `instanceType` | `string` | no |
| `maxConcurrentSessions` | `integer` | no |
| `portalCustomDomain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `portal` | `Portal` | no |

## UpdateSessionLogger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLoggerArn` | `string` | yes |
| `eventFilter` | `EventFilter` | no |
| `logConfiguration` | `LogConfiguration` | no |
| `displayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sessionLogger` | `SessionLogger` | yes |

## UpdateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |
| `certificatesToAdd` | `List<blob>` | no |
| `certificatesToDelete` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustStoreArn` | `string` | yes |

## UpdateUserAccessLoggingSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettingsArn` | `string` | yes |
| `kinesisStreamArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userAccessLoggingSettings` | `UserAccessLoggingSettings` | yes |

## UpdateUserSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettingsArn` | `string` | yes |
| `copyAllowed` | `string` | no |
| `pasteAllowed` | `string` | no |
| `downloadAllowed` | `string` | no |
| `uploadAllowed` | `string` | no |
| `printAllowed` | `string` | no |
| `disconnectTimeoutInMinutes` | `integer` | no |
| `idleDisconnectTimeoutInMinutes` | `integer` | no |
| `clientToken` | `string` | no |
| `cookieSynchronizationConfiguration` | `CookieSynchronizationConfiguration` | no |
| `deepLinkAllowed` | `string` | no |
| `toolbarConfiguration` | `ToolbarConfiguration` | no |
| `brandingConfigurationInput` | `BrandingConfigurationUpdateInput` | no |
| `webAuthnAllowed` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userSettings` | `UserSettings` | yes |

