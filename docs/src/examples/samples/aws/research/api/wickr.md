# AWS Wickr Admin API

API version: 2024-02-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/wickr/2024-02-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `users` | `List<BatchCreateUserRequestItem>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<User>` | no |
| `failed` | `List<BatchUserErrorResponseItem>` | no |

## BatchDeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<BatchUserSuccessResponseItem>` | no |
| `failed` | `List<BatchUserErrorResponseItem>` | no |

## BatchLookupUserUname

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `unames` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<BatchUnameSuccessResponseItem>` | no |
| `failed` | `List<BatchUnameErrorResponseItem>` | no |

## BatchReinviteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<BatchUserSuccessResponseItem>` | no |
| `failed` | `List<BatchUserErrorResponseItem>` | no |

## BatchResetDevicesForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userId` | `string` | yes |
| `appIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<BatchDeviceSuccessResponseItem>` | no |
| `failed` | `List<BatchDeviceErrorResponseItem>` | no |

## BatchToggleUserSuspendStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `suspend` | `boolean` | yes |
| `userIds` | `List<string>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `successful` | `List<BatchUserSuccessResponseItem>` | no |
| `failed` | `List<BatchUserErrorResponseItem>` | no |

## CreateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `username` | `string` | yes |
| `displayName` | `string` | no |
| `groupId` | `string` | yes |
| `challenge` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `botId` | `string` | yes |
| `networkId` | `string` | no |
| `username` | `string` | no |
| `displayName` | `string` | no |
| `groupId` | `string` | no |

## CreateDataRetentionBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## CreateDataRetentionBotChallenge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `challenge` | `string` | yes |

## CreateNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkName` | `string` | yes |
| `accessLevel` | `string` | yes |
| `enablePremiumFreeTrial` | `boolean` | no |
| `encryptionKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | no |
| `networkName` | `string` | no |
| `encryptionKeyArn` | `string` | no |

## CreateSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `name` | `string` | yes |
| `securityGroupSettings` | `SecurityGroupSettingsRequest` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityGroup` | `SecurityGroup` | yes |

## DeleteBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## DeleteDataRetentionBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## DeleteNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## DeleteSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `groupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `networkId` | `string` | no |
| `groupId` | `string` | no |

## GetBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `botId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botId` | `string` | no |
| `displayName` | `string` | no |
| `username` | `string` | no |
| `uname` | `string` | no |
| `pubkey` | `string` | no |
| `status` | `integer` | no |
| `groupId` | `string` | no |
| `hasChallenge` | `boolean` | no |
| `suspended` | `boolean` | no |
| `lastLogin` | `string` | no |

## GetBotsCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pending` | `integer` | yes |
| `active` | `integer` | yes |
| `total` | `integer` | yes |

## GetDataRetentionBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `botName` | `string` | no |
| `botExists` | `boolean` | no |
| `isBotActive` | `boolean` | no |
| `isDataRetentionBotRegistered` | `boolean` | no |
| `isDataRetentionServiceEnabled` | `boolean` | no |
| `isPubkeyMsgAcked` | `boolean` | no |

## GetGuestUserHistoryCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `history` | `List<GuestUserHistoryCount>` | yes |

## GetNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `networkName` | `string` | yes |
| `accessLevel` | `string` | yes |
| `awsAccountId` | `string` | yes |
| `networkArn` | `string` | yes |
| `standing` | `integer` | no |
| `freeTrialExpiration` | `string` | no |
| `migrationState` | `integer` | no |
| `encryptionKeyArn` | `string` | no |

## GetNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<Setting>` | yes |

## GetOidcInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `clientId` | `string` | no |
| `code` | `string` | no |
| `grantType` | `string` | no |
| `redirectUri` | `string` | no |
| `url` | `string` | no |
| `clientSecret` | `string` | no |
| `codeVerifier` | `string` | no |
| `certificate` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `openidConnectInfo` | `OidcConfigInfo` | no |
| `tokenInfo` | `OidcTokenInfo` | no |

## GetOpentdfConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `domain` | `string` | yes |
| `clientSecret` | `string` | yes |
| `provider` | `string` | yes |

## GetSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `groupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityGroup` | `SecurityGroup` | yes |

## GetUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userId` | `string` | yes |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `username` | `string` | no |
| `isAdmin` | `boolean` | no |
| `suspended` | `boolean` | no |
| `status` | `integer` | no |
| `lastActivity` | `integer` | no |
| `lastLogin` | `integer` | no |
| `securityGroupIds` | `List<string>` | no |

## GetUsersCount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pending` | `integer` | yes |
| `active` | `integer` | yes |
| `rejected` | `integer` | yes |
| `remaining` | `integer` | yes |
| `total` | `integer` | yes |

## ListBlockedGuestUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `maxResults` | `integer` | no |
| `sortDirection` | `string` | no |
| `sortFields` | `string` | no |
| `username` | `string` | no |
| `admin` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `blocklist` | `List<BlockedGuestUser>` | yes |

## ListBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |
| `displayName` | `string` | no |
| `username` | `string` | no |
| `status` | `integer` | no |
| `groupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bots` | `List<Bot>` | yes |
| `nextToken` | `string` | no |

## ListDevicesForUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `devices` | `List<BasicDeviceObject>` | yes |

## ListGuestUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `maxResults` | `integer` | no |
| `sortDirection` | `string` | no |
| `sortFields` | `string` | no |
| `username` | `string` | no |
| `billingPeriod` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `guestlist` | `List<GuestUser>` | yes |

## ListNetworks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networks` | `List<Network>` | yes |
| `nextToken` | `string` | no |

## ListSecurityGroupUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `groupId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `users` | `List<User>` | yes |
| `nextToken` | `string` | no |

## ListSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityGroups` | `List<SecurityGroup>` | no |
| `nextToken` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `sortFields` | `string` | no |
| `sortDirection` | `string` | no |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `username` | `string` | no |
| `status` | `integer` | no |
| `groupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `users` | `List<User>` | no |

## RegisterOidcConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `companyId` | `string` | yes |
| `customUsername` | `string` | no |
| `extraAuthParams` | `string` | no |
| `issuer` | `string` | yes |
| `scopes` | `string` | yes |
| `secret` | `string` | no |
| `ssoTokenBufferMinutes` | `integer` | no |
| `userId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationName` | `string` | no |
| `clientId` | `string` | no |
| `companyId` | `string` | yes |
| `scopes` | `string` | yes |
| `issuer` | `string` | yes |
| `clientSecret` | `string` | no |
| `secret` | `string` | no |
| `redirectUrl` | `string` | no |
| `userId` | `string` | no |
| `customUsername` | `string` | no |
| `caCertificate` | `string` | no |
| `applicationId` | `integer` | no |
| `ssoTokenBufferMinutes` | `integer` | no |
| `extraAuthParams` | `string` | no |

## RegisterOidcConfigTest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `extraAuthParams` | `string` | no |
| `issuer` | `string` | yes |
| `scopes` | `string` | yes |
| `certificate` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenEndpoint` | `string` | no |
| `userinfoEndpoint` | `string` | no |
| `responseTypesSupported` | `List<string>` | no |
| `scopesSupported` | `List<string>` | no |
| `issuer` | `string` | no |
| `authorizationEndpoint` | `string` | no |
| `endSessionEndpoint` | `string` | no |
| `logoutEndpoint` | `string` | no |
| `grantTypesSupported` | `List<string>` | no |
| `revocationEndpoint` | `string` | no |
| `tokenEndpointAuthMethodsSupported` | `List<string>` | no |
| `microsoftMultiRefreshToken` | `boolean` | no |

## RegisterOpentdfConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `clientId` | `string` | yes |
| `clientSecret` | `string` | yes |
| `domain` | `string` | yes |
| `provider` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `domain` | `string` | yes |
| `clientSecret` | `string` | yes |
| `provider` | `string` | yes |

## UpdateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `botId` | `string` | yes |
| `displayName` | `string` | no |
| `groupId` | `string` | no |
| `challenge` | `string` | no |
| `suspend` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## UpdateDataRetention

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `actionType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## UpdateGuestUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `usernameHash` | `string` | yes |
| `block` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## UpdateNetwork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `networkName` | `string` | yes |
| `clientToken` | `string` | no |
| `encryptionKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |

## UpdateNetworkSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `settings` | `NetworkSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `settings` | `List<Setting>` | yes |

## UpdateSecurityGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `groupId` | `string` | yes |
| `name` | `string` | no |
| `securityGroupSettings` | `SecurityGroupSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityGroup` | `SecurityGroup` | yes |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `networkId` | `string` | yes |
| `userId` | `string` | yes |
| `userDetails` | `UpdateUserDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `userId` | `string` | yes |
| `networkId` | `string` | yes |
| `securityGroupIds` | `List<string>` | no |
| `firstName` | `string` | no |
| `lastName` | `string` | no |
| `middleName` | `string` | no |
| `suspended` | `boolean` | yes |
| `modified` | `integer` | no |
| `status` | `integer` | no |
| `inviteCode` | `string` | no |
| `inviteExpiration` | `integer` | no |
| `codeValidation` | `boolean` | no |

