# AWS Certificate Manager Private Certificate Authority

API version: 2017-08-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/acm-pca/2017-08-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityConfiguration` | `CertificateAuthorityConfiguration` | yes |
| `RevocationConfiguration` | `RevocationConfiguration` | no |
| `CertificateAuthorityType` | `string` | yes |
| `IdempotencyToken` | `string` | no |
| `KeyStorageSecurityStandard` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UsageMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | no |

## CreateCertificateAuthorityAuditReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `S3BucketName` | `string` | yes |
| `AuditReportResponseFormat` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuditReportId` | `string` | no |
| `S3Key` | `string` | no |

## CreatePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `Principal` | `string` | yes |
| `SourceAccount` | `string` | no |
| `Actions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `PermanentDeletionTimeInDays` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `Principal` | `string` | yes |
| `SourceAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthority` | `CertificateAuthority` | no |

## DescribeCertificateAuthorityAuditReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `AuditReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuditReportStatus` | `string` | no |
| `S3BucketName` | `string` | no |
| `S3Key` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## GetCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `CertificateChain` | `string` | no |

## GetCertificateAuthorityCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificate` | `string` | no |
| `CertificateChain` | `string` | no |

## GetCertificateAuthorityCsr

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Csr` | `string` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## ImportCertificateAuthorityCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `Certificate` | `blob` | yes |
| `CertificateChain` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## IssueCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiPassthrough` | `ApiPassthrough` | no |
| `CertificateAuthorityArn` | `string` | yes |
| `Csr` | `blob` | yes |
| `SigningAlgorithm` | `string` | yes |
| `TemplateArn` | `string` | no |
| `Validity` | `Validity` | yes |
| `ValidityNotBefore` | `Validity` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateArn` | `string` | no |

## ListCertificateAuthorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CertificateAuthorities` | `List<CertificateAuthority>` | no |

## ListPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Permissions` | `List<Permission>` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

## PutPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestoreCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `CertificateSerial` | `string` | yes |
| `RevocationReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `RevocationConfiguration` | `RevocationConfiguration` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


