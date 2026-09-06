# Amazon Cognito Identity

API version: 2014-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cognito-identity/2014-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateIdentityPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolName` | `string` | yes |
| `AllowUnauthenticatedIdentities` | `boolean` | yes |
| `AllowClassicFlow` | `boolean` | no |
| `SupportedLoginProviders` | `Map<string>` | no |
| `DeveloperProviderName` | `string` | no |
| `OpenIdConnectProviderARNs` | `List<string>` | no |
| `CognitoIdentityProviders` | `List<CognitoIdentityProvider>` | no |
| `SamlProviderARNs` | `List<string>` | no |
| `IdentityPoolTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityPoolName` | `string` | yes |
| `AllowUnauthenticatedIdentities` | `boolean` | yes |
| `AllowClassicFlow` | `boolean` | no |
| `SupportedLoginProviders` | `Map<string>` | no |
| `DeveloperProviderName` | `string` | no |
| `OpenIdConnectProviderARNs` | `List<string>` | no |
| `CognitoIdentityProviders` | `List<CognitoIdentityProvider>` | no |
| `SamlProviderARNs` | `List<string>` | no |
| `IdentityPoolTags` | `Map<string>` | no |

## DeleteIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityIdsToDelete` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UnprocessedIdentityIds` | `List<UnprocessedIdentityId>` | no |

## DeleteIdentityPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |
| `Logins` | `List<string>` | no |
| `CreationDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |

## DescribeIdentityPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityPoolName` | `string` | yes |
| `AllowUnauthenticatedIdentities` | `boolean` | yes |
| `AllowClassicFlow` | `boolean` | no |
| `SupportedLoginProviders` | `Map<string>` | no |
| `DeveloperProviderName` | `string` | no |
| `OpenIdConnectProviderARNs` | `List<string>` | no |
| `CognitoIdentityProviders` | `List<CognitoIdentityProvider>` | no |
| `SamlProviderARNs` | `List<string>` | no |
| `IdentityPoolTags` | `Map<string>` | no |

## GetCredentialsForIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | yes |
| `Logins` | `Map<string>` | no |
| `CustomRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |
| `Credentials` | `Credentials` | no |

## GetId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `IdentityPoolId` | `string` | yes |
| `Logins` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |

## GetIdentityPoolRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `Roles` | `Map<string>` | no |
| `RoleMappings` | `Map<RoleMapping>` | no |

## GetOpenIdToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | yes |
| `Logins` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |
| `Token` | `string` | no |

## GetOpenIdTokenForDeveloperIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | no |
| `Logins` | `Map<string>` | yes |
| `PrincipalTags` | `Map<string>` | no |
| `TokenDuration` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |
| `Token` | `string` | no |

## GetPrincipalTagAttributeMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `IdentityProviderName` | `string` | no |
| `UseDefaults` | `boolean` | no |
| `PrincipalTags` | `Map<string>` | no |

## ListIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `MaxResults` | `integer` | yes |
| `NextToken` | `string` | no |
| `HideDisabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `Identities` | `List<IdentityDescription>` | no |
| `NextToken` | `string` | no |

## ListIdentityPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPools` | `List<IdentityPoolShortDescription>` | no |
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

## LookupDeveloperIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityId` | `string` | no |
| `DeveloperUserIdentifier` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |
| `DeveloperUserIdentifierList` | `List<string>` | no |
| `NextToken` | `string` | no |

## MergeDeveloperIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceUserIdentifier` | `string` | yes |
| `DestinationUserIdentifier` | `string` | yes |
| `DeveloperProviderName` | `string` | yes |
| `IdentityPoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | no |

## SetIdentityPoolRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `Roles` | `Map<string>` | yes |
| `RoleMappings` | `Map<RoleMapping>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetPrincipalTagAttributeMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityProviderName` | `string` | yes |
| `UseDefaults` | `boolean` | no |
| `PrincipalTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | no |
| `IdentityProviderName` | `string` | no |
| `UseDefaults` | `boolean` | no |
| `PrincipalTags` | `Map<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnlinkDeveloperIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | yes |
| `IdentityPoolId` | `string` | yes |
| `DeveloperProviderName` | `string` | yes |
| `DeveloperUserIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnlinkIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityId` | `string` | yes |
| `Logins` | `Map<string>` | yes |
| `LoginsToRemove` | `List<string>` | yes |

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


## UpdateIdentityPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityPoolName` | `string` | yes |
| `AllowUnauthenticatedIdentities` | `boolean` | yes |
| `AllowClassicFlow` | `boolean` | no |
| `SupportedLoginProviders` | `Map<string>` | no |
| `DeveloperProviderName` | `string` | no |
| `OpenIdConnectProviderARNs` | `List<string>` | no |
| `CognitoIdentityProviders` | `List<CognitoIdentityProvider>` | no |
| `SamlProviderARNs` | `List<string>` | no |
| `IdentityPoolTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityPoolId` | `string` | yes |
| `IdentityPoolName` | `string` | yes |
| `AllowUnauthenticatedIdentities` | `boolean` | yes |
| `AllowClassicFlow` | `boolean` | no |
| `SupportedLoginProviders` | `Map<string>` | no |
| `DeveloperProviderName` | `string` | no |
| `OpenIdConnectProviderARNs` | `List<string>` | no |
| `CognitoIdentityProviders` | `List<CognitoIdentityProvider>` | no |
| `SamlProviderARNs` | `List<string>` | no |
| `IdentityPoolTags` | `Map<string>` | no |

