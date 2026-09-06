# AWS SSO OIDC

API version: 2019-06-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sso-oidc/2019-06-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `clientSecret` | `string` | yes |
| `grantType` | `string` | yes |
| `deviceCode` | `string` | no |
| `code` | `string` | no |
| `refreshToken` | `string` | no |
| `scope` | `List<string>` | no |
| `redirectUri` | `string` | no |
| `codeVerifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | no |
| `tokenType` | `string` | no |
| `expiresIn` | `integer` | no |
| `refreshToken` | `string` | no |
| `idToken` | `string` | no |

## CreateTokenWithIAM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `grantType` | `string` | yes |
| `code` | `string` | no |
| `refreshToken` | `string` | no |
| `assertion` | `string` | no |
| `scope` | `List<string>` | no |
| `redirectUri` | `string` | no |
| `subjectToken` | `string` | no |
| `subjectTokenType` | `string` | no |
| `requestedTokenType` | `string` | no |
| `codeVerifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | no |
| `tokenType` | `string` | no |
| `expiresIn` | `integer` | no |
| `refreshToken` | `string` | no |
| `idToken` | `string` | no |
| `issuedTokenType` | `string` | no |
| `scope` | `List<string>` | no |
| `awsAdditionalDetails` | `AwsAdditionalDetails` | no |

## RegisterClient

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientName` | `string` | yes |
| `clientType` | `string` | yes |
| `scopes` | `List<string>` | no |
| `redirectUris` | `List<string>` | no |
| `grantTypes` | `List<string>` | no |
| `issuerUrl` | `string` | no |
| `entitledApplicationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | no |
| `clientSecret` | `string` | no |
| `clientIdIssuedAt` | `long` | no |
| `clientSecretExpiresAt` | `long` | no |
| `authorizationEndpoint` | `string` | no |
| `tokenEndpoint` | `string` | no |

## StartDeviceAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientId` | `string` | yes |
| `clientSecret` | `string` | yes |
| `startUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deviceCode` | `string` | no |
| `userCode` | `string` | no |
| `verificationUri` | `string` | no |
| `verificationUriComplete` | `string` | no |
| `expiresIn` | `integer` | no |
| `interval` | `integer` | no |

