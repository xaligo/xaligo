# AWS Security Token Service

API version: 2011-06-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sts/2011-06-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssumeRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |
| `RoleSessionName` | `string` | yes |
| `PolicyArns` | `List<PolicyDescriptorType>` | no |
| `Policy` | `string` | no |
| `DurationSeconds` | `integer` | no |
| `Tags` | `List<Tag>` | no |
| `TransitiveTagKeys` | `List<string>` | no |
| `ExternalId` | `string` | no |
| `SerialNumber` | `string` | no |
| `TokenCode` | `string` | no |
| `SourceIdentity` | `string` | no |
| `ProvidedContexts` | `List<ProvidedContext>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `AssumedRoleUser` | `AssumedRoleUser` | no |
| `PackedPolicySize` | `integer` | no |
| `SourceIdentity` | `string` | no |

## AssumeRoleWithSAML

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |
| `PrincipalArn` | `string` | yes |
| `SAMLAssertion` | `string` | yes |
| `PolicyArns` | `List<PolicyDescriptorType>` | no |
| `Policy` | `string` | no |
| `DurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `AssumedRoleUser` | `AssumedRoleUser` | no |
| `PackedPolicySize` | `integer` | no |
| `Subject` | `string` | no |
| `SubjectType` | `string` | no |
| `Issuer` | `string` | no |
| `Audience` | `string` | no |
| `NameQualifier` | `string` | no |
| `SourceIdentity` | `string` | no |

## AssumeRoleWithWebIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |
| `RoleSessionName` | `string` | yes |
| `WebIdentityToken` | `string` | yes |
| `ProviderId` | `string` | no |
| `PolicyArns` | `List<PolicyDescriptorType>` | no |
| `Policy` | `string` | no |
| `DurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `SubjectFromWebIdentityToken` | `string` | no |
| `AssumedRoleUser` | `AssumedRoleUser` | no |
| `PackedPolicySize` | `integer` | no |
| `Provider` | `string` | no |
| `Audience` | `string` | no |
| `SourceIdentity` | `string` | no |

## AssumeRoot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetPrincipal` | `string` | yes |
| `TaskPolicyArn` | `PolicyDescriptorType` | yes |
| `DurationSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `SourceIdentity` | `string` | no |

## DecodeAuthorizationMessage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EncodedMessage` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DecodedMessage` | `string` | no |

## GetAccessKeyInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessKeyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `string` | no |

## GetCallerIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `Account` | `string` | no |
| `Arn` | `string` | no |

## GetDelegatedAccessToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TradeInToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `PackedPolicySize` | `integer` | no |
| `AssumedPrincipal` | `string` | no |

## GetFederationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Policy` | `string` | no |
| `PolicyArns` | `List<PolicyDescriptorType>` | no |
| `DurationSeconds` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |
| `FederatedUser` | `FederatedUser` | no |
| `PackedPolicySize` | `integer` | no |

## GetSessionToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurationSeconds` | `integer` | no |
| `SerialNumber` | `string` | no |
| `TokenCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Credentials` | `Credentials` | no |

## GetWebIdentityToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Audience` | `List<string>` | yes |
| `DurationSeconds` | `integer` | no |
| `SigningAlgorithm` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebIdentityToken` | `string` | no |
| `Expiration` | `timestamp` | no |

