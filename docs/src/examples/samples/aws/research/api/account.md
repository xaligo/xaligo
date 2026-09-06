# AWS Account

API version: 2021-02-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/account/2021-02-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptPrimaryEmailUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `PrimaryEmail` | `string` | yes |
| `Otp` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DeleteAlternateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlternateContactType` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `AccountName` | `string` | no |
| `AccountCreatedDate` | `timestamp` | no |
| `AccountState` | `string` | no |

## GetAlternateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlternateContactType` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlternateContact` | `AlternateContact` | no |

## GetContactInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactInformation` | `ContactInformation` | no |

## GetGovCloudAccountInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StandardAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GovCloudAccountId` | `string` | yes |
| `AccountState` | `string` | yes |

## GetPrimaryEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrimaryEmail` | `string` | no |

## GetPrimaryEmailUpdateStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |
| `UpdatedAt` | `timestamp` | no |

## GetRegionOptStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `RegionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionName` | `string` | no |
| `RegionOptStatus` | `string` | no |

## ListRegions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RegionOptStatusContains` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Regions` | `List<Region>` | no |

## PutAccountName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountName` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAlternateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Title` | `string` | yes |
| `EmailAddress` | `string` | yes |
| `PhoneNumber` | `string` | yes |
| `AlternateContactType` | `string` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutContactInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactInformation` | `ContactInformation` | yes |
| `AccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartPrimaryEmailUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountId` | `string` | yes |
| `PrimaryEmail` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

