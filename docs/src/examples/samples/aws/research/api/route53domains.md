# Amazon Route 53 Domains

API version: 2014-05-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/route53domains/2014-05-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptDomainTransferFromAnotherAwsAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Password` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## AssociateDelegationSignerToDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SigningAttributes` | `DnssecSigningAttributes` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CancelDomainTransferToAnotherAwsAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CheckDomainAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IdnLangCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Availability` | `string` | no |

## CheckDomainTransferability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AuthCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Transferability` | `DomainTransferability` | no |
| `Message` | `string` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## DeleteTagsForDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `TagsToDelete` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableDomainAutoRenew

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableDomainTransferLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## DisassociateDelegationSignerFromDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## EnableDomainAutoRenew

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableDomainTransferLock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## GetContactReachabilityStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `status` | `string` | no |

## GetDomainDetail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `Nameservers` | `List<Nameserver>` | no |
| `AutoRenew` | `boolean` | no |
| `AdminContact` | `ContactDetail` | no |
| `RegistrantContact` | `ContactDetail` | no |
| `TechContact` | `ContactDetail` | no |
| `AdminPrivacy` | `boolean` | no |
| `RegistrantPrivacy` | `boolean` | no |
| `TechPrivacy` | `boolean` | no |
| `RegistrarName` | `string` | no |
| `WhoIsServer` | `string` | no |
| `RegistrarUrl` | `string` | no |
| `AbuseContactEmail` | `string` | no |
| `AbuseContactPhone` | `string` | no |
| `RegistryDomainId` | `string` | no |
| `CreationDate` | `timestamp` | no |
| `UpdatedDate` | `timestamp` | no |
| `ExpirationDate` | `timestamp` | no |
| `Reseller` | `string` | no |
| `DnsSec` | `string` | no |
| `StatusList` | `List<string>` | no |
| `DnssecKeys` | `List<DnssecKey>` | no |
| `BillingContact` | `ContactDetail` | no |
| `BillingPrivacy` | `boolean` | no |

## GetDomainSuggestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SuggestionCount` | `integer` | yes |
| `OnlyAvailable` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuggestionsList` | `List<DomainSuggestion>` | no |

## GetOperationDetail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `DomainName` | `string` | no |
| `Type` | `string` | no |
| `SubmittedDate` | `timestamp` | no |
| `LastUpdatedDate` | `timestamp` | no |
| `StatusFlag` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterConditions` | `List<FilterCondition>` | no |
| `SortCondition` | `SortCondition` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domains` | `List<DomainSummary>` | no |
| `NextPageMarker` | `string` | no |

## ListOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubmittedSince` | `timestamp` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |
| `Status` | `List<string>` | no |
| `Type` | `List<string>` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operations` | `List<OperationSummary>` | no |
| `NextPageMarker` | `string` | no |

## ListPrices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tld` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Prices` | `List<DomainPrice>` | no |
| `NextPageMarker` | `string` | no |

## ListTagsForDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## PushDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Target` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IdnLangCode` | `string` | no |
| `DurationInYears` | `integer` | yes |
| `AutoRenew` | `boolean` | no |
| `AdminContact` | `ContactDetail` | yes |
| `RegistrantContact` | `ContactDetail` | yes |
| `TechContact` | `ContactDetail` | yes |
| `PrivacyProtectAdminContact` | `boolean` | no |
| `PrivacyProtectRegistrantContact` | `boolean` | no |
| `PrivacyProtectTechContact` | `boolean` | no |
| `BillingContact` | `ContactDetail` | no |
| `PrivacyProtectBillingContact` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## RejectDomainTransferFromAnotherAwsAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## RenewDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DurationInYears` | `integer` | no |
| `CurrentExpiryYear` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## ResendContactReachabilityEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `emailAddress` | `string` | no |
| `isAlreadyVerified` | `boolean` | no |

## ResendOperationAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RetrieveDomainAuthCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthCode` | `string` | no |

## TransferDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IdnLangCode` | `string` | no |
| `DurationInYears` | `integer` | no |
| `Nameservers` | `List<Nameserver>` | no |
| `AuthCode` | `string` | no |
| `AutoRenew` | `boolean` | no |
| `AdminContact` | `ContactDetail` | yes |
| `RegistrantContact` | `ContactDetail` | yes |
| `TechContact` | `ContactDetail` | yes |
| `PrivacyProtectAdminContact` | `boolean` | no |
| `PrivacyProtectRegistrantContact` | `boolean` | no |
| `PrivacyProtectTechContact` | `boolean` | no |
| `BillingContact` | `ContactDetail` | no |
| `PrivacyProtectBillingContact` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## TransferDomainToAnotherAwsAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |
| `Password` | `string` | no |

## UpdateDomainContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AdminContact` | `ContactDetail` | no |
| `RegistrantContact` | `ContactDetail` | no |
| `TechContact` | `ContactDetail` | no |
| `Consent` | `Consent` | no |
| `BillingContact` | `ContactDetail` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateDomainContactPrivacy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AdminPrivacy` | `boolean` | no |
| `RegistrantPrivacy` | `boolean` | no |
| `TechPrivacy` | `boolean` | no |
| `BillingPrivacy` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateDomainNameservers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `FIAuthKey` | `string` | no |
| `Nameservers` | `List<Nameserver>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateTagsForDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `TagsToUpdate` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ViewBilling

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Start` | `timestamp` | no |
| `End` | `timestamp` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextPageMarker` | `string` | no |
| `BillingRecords` | `List<BillingRecord>` | no |

