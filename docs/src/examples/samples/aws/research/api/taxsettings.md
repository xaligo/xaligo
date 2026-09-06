# Tax Settings

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/taxsettings/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteTaxRegistrationError>` | yes |

## BatchGetTaxExemptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taxExemptionDetailsMap` | `Map<TaxExemptionDetails>` | no |
| `failedAccounts` | `List<string>` | no |

## BatchPutTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |
| `taxRegistrationEntry` | `TaxRegistrationEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `errors` | `List<BatchPutTaxRegistrationError>` | yes |

## DeleteSupplementalTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetTaxExemptionTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taxExemptionTypes` | `List<TaxExemptionType>` | no |

## GetTaxInheritance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `heritageStatus` | `string` | no |

## GetTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taxRegistration` | `TaxRegistration` | no |

## GetTaxRegistrationDocument

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationS3Location` | `DestinationS3Location` | no |
| `taxDocumentMetadata` | `TaxDocumentMetadata` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `destinationFilePath` | `string` | no |
| `presignedS3Url` | `string` | no |

## ListSupplementalTaxRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taxRegistrations` | `List<SupplementalTaxRegistration>` | yes |
| `nextToken` | `string` | no |

## ListTaxExemptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `taxExemptionDetailsMap` | `Map<TaxExemptionDetails>` | no |

## ListTaxRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountDetails` | `List<AccountDetails>` | yes |
| `nextToken` | `string` | no |

## PutSupplementalTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `taxRegistrationEntry` | `SupplementalTaxRegistrationEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorityId` | `string` | yes |
| `status` | `string` | yes |

## PutTaxExemption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountIds` | `List<string>` | yes |
| `authority` | `Authority` | yes |
| `exemptionType` | `string` | yes |
| `exemptionCertificate` | `ExemptionCertificate` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `caseId` | `string` | no |

## PutTaxInheritance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `heritageStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTaxRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | no |
| `taxRegistrationEntry` | `TaxRegistrationEntry` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

