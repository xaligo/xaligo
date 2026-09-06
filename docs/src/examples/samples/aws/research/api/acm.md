# AWS Certificate Manager

API version: 2015-12-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/acm/2015-12-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTagsToCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAcmeDomainValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdempotencyToken` | `string` | no |
| `AcmeEndpointArn` | `string` | yes |
| `DomainName` | `string` | yes |
| `PrevalidationOptions` | `PrevalidationOptions` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidationArn` | `string` | yes |

## CreateAcmeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdempotencyToken` | `string` | no |
| `AuthorizationBehavior` | `string` | yes |
| `Contact` | `string` | no |
| `CertificateAuthority` | `CertificateAuthority` | yes |
| `Tags` | `List<Tag>` | no |
| `CertificateTags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | no |

## CreateAcmeExternalAccountBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdempotencyToken` | `string` | no |
| `AcmeEndpointArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Expiration` | `Expiration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExternalAccountBinding` | `AcmeExternalAccountBinding` | no |

## DeleteAcmeDomainValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAcmeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAcmeExternalAccountBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeExternalAccountBindingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAcmeAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | yes |
| `AccountUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeAccount` | `AcmeAccount` | no |

## DescribeAcmeDomainValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidation` | `AcmeDomainValidation` | no |

## DescribeAcmeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpoint` | `AcmeEndpoint` | no |

## DescribeAcmeExternalAccountBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeExternalAccountBindingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExternalAccountBinding` | `AcmeExternalAccountBinding` | no |

## DescribeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `CertificateDetail` | no |

## ExportCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `Passphrase` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `CertificateChain` | `string` | no |
| `PrivateKey` | `string` | no |

## GetAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExpiryEvents` | `ExpiryEventsConfiguration` | no |

## GetAcmeExternalAccountBindingCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeExternalAccountBindingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyId` | `string` | no |
| `MacKey` | `string` | no |

## GetCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `CertificateChain` | `string` | no |

## ImportCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | no |
| `Certificate` | `blob` | yes |
| `PrivateKey` | `blob` | yes |
| `CertificateChain` | `blob` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | no |

## ListAcmeAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AcmeEndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeAccounts` | `List<AcmeAccountSummary>` | no |
| `NextToken` | `string` | no |

## ListAcmeDomainValidations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AcmeEndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidations` | `List<AcmeDomainValidationSummary>` | no |
| `NextToken` | `string` | no |

## ListAcmeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpoints` | `List<AcmeEndpointSummary>` | no |
| `NextToken` | `string` | no |

## ListAcmeExternalAccountBindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AcmeEndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExternalAccountBindings` | `List<AcmeExternalAccountBindingSummary>` | no |
| `NextToken` | `string` | no |

## ListCertificateDomainValidations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainValidationSummaryList` | `List<DomainValidationSummary>` | no |
| `NextToken` | `string` | no |

## ListCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateStatuses` | `List<string>` | no |
| `CertificateKeyPairOrigins` | `List<string>` | no |
| `Includes` | `Filters` | no |
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CertificateSummaryList` | `List<CertificateSummary>` | no |

## ListTagsForCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutAccountConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExpiryEvents` | `ExpiryEventsConfiguration` | no |
| `IdempotencyToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTagsFromCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RenewCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RequestCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ValidationMethod` | `string` | no |
| `SubjectAlternativeNames` | `List<string>` | no |
| `IdempotencyToken` | `string` | no |
| `DomainValidationOptions` | `List<DomainValidationOption>` | no |
| `Options` | `CertificateOptions` | no |
| `CertificateAuthorityArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `KeyAlgorithm` | `string` | no |
| `ManagedBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | no |

## ResendValidationEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `Domain` | `string` | yes |
| `ValidationDomain` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeAcmeAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | yes |
| `AccountUrl` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeAcmeExternalAccountBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeExternalAccountBindingArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `RevocationReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | no |

## SearchCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterStatement` | `CertificateFilterStatement` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<CertificateSearchResult>` | no |
| `NextToken` | `string` | no |

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


## UpdateAcmeDomainValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeDomainValidationArn` | `string` | yes |
| `PrevalidationOptions` | `PrevalidationOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAcmeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcmeEndpointArn` | `string` | yes |
| `AuthorizationBehavior` | `string` | no |
| `Contact` | `string` | no |
| `CertificateAuthority` | `CertificateAuthority` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCertificateOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | yes |
| `Options` | `CertificateOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


