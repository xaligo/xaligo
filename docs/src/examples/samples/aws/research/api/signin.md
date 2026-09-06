# AWS Sign-In Service

API version: 2023-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/signin/2023-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateOAuth2Token

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenInput` | `CreateOAuth2TokenRequestBody` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenOutput` | `CreateOAuth2TokenResponseBody` | yes |

## CreateOAuth2TokenWithIAM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `grantType` | `string` | yes |
| `resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | yes |
| `tokenType` | `string` | yes |
| `expiresIn` | `integer` | yes |

## DeleteConsoleAuthorizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | yes |
| `scope` | `string` | yes |
| `consoleAuthorizationEnabled` | `boolean` | yes |

## DeleteResourcePermissionStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statementId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConsoleAuthorizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | yes |
| `scope` | `string` | yes |
| `consoleAuthorizationEnabled` | `boolean` | yes |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signinResourceBasedPolicy` | `SigninResourceBasedPolicy` | yes |

## IntrospectOAuth2TokenWithIAM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `token` | `string` | yes |
| `tokenTypeHint` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `active` | `boolean` | yes |
| `clientId` | `string` | no |
| `userId` | `string` | no |
| `tokenType` | `string` | no |
| `exp` | `long` | no |
| `iat` | `long` | no |
| `nbf` | `long` | no |
| `sub` | `string` | no |
| `aud` | `string` | no |
| `iss` | `string` | no |
| `jti` | `string` | no |
| `accountId` | `string` | no |
| `signinSession` | `string` | no |
| `resource` | `string` | no |

## ListResourcePermissionStatements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionStatements` | `List<PermissionStatementSummary>` | yes |
| `nextToken` | `string` | no |

## PutConsoleAuthorizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetId` | `string` | yes |
| `scope` | `string` | yes |
| `consoleAuthorizationEnabled` | `boolean` | yes |

## PutResourcePermissionStatement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceVpc` | `string` | no |
| `signinSourceVpce` | `string` | no |
| `consoleSourceVpce` | `string` | no |
| `vpcSourceIp` | `string` | no |
| `sourceIp` | `string` | no |
| `requestedRegion` | `string` | no |
| `excludedPrincipal` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statementId` | `string` | yes |

## RevokeOAuth2TokenWithIAM

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `token` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


