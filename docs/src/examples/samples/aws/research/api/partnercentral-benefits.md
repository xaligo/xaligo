# Partner Central Benefits API

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/partnercentral-benefits/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AmendBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Revision` | `string` | yes |
| `Identifier` | `string` | yes |
| `AmendmentReason` | `string` | yes |
| `Amendments` | `List<Amendment>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateBenefitApplicationResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `BenefitApplicationIdentifier` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Revision` | `string` | no |

## CancelBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Identifier` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `BenefitIdentifier` | `string` | yes |
| `FulfillmentTypes` | `List<string>` | no |
| `BenefitApplicationDetails` | `Document` | no |
| `Tags` | `List<Tag>` | no |
| `AssociatedResources` | `List<string>` | no |
| `PartnerContacts` | `List<Contact>` | no |
| `FileDetails` | `List<FileInput>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Revision` | `string` | no |

## DisassociateBenefitApplicationResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `BenefitApplicationIdentifier` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Revision` | `string` | no |

## GetBenefit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Catalog` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Programs` | `List<string>` | no |
| `FulfillmentTypes` | `List<string>` | no |
| `BenefitRequestSchema` | `Document` | no |
| `Status` | `string` | no |

## GetBenefitAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Catalog` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `BenefitApplicationId` | `string` | no |
| `BenefitId` | `string` | no |
| `FulfillmentType` | `string` | no |
| `ApplicableBenefitIds` | `List<string>` | no |
| `FulfillmentDetail` | `FulfillmentDetails` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `StartsAt` | `timestamp` | no |
| `ExpiresAt` | `timestamp` | no |

## GetBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Catalog` | `string` | no |
| `BenefitId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `FulfillmentTypes` | `List<string>` | no |
| `BenefitApplicationDetails` | `Document` | no |
| `Programs` | `List<string>` | no |
| `Status` | `string` | no |
| `Stage` | `string` | no |
| `StatusReason` | `string` | no |
| `StatusReasonCode` | `string` | no |
| `StatusReasonCodes` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `Revision` | `string` | no |
| `AssociatedResources` | `List<string>` | no |
| `PartnerContacts` | `List<Contact>` | no |
| `FileDetails` | `List<FileDetail>` | no |

## ListBenefitAllocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `FulfillmentTypes` | `List<string>` | no |
| `BenefitIdentifiers` | `List<string>` | no |
| `BenefitApplicationIdentifiers` | `List<string>` | no |
| `Status` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BenefitAllocationSummaries` | `List<BenefitAllocationSummary>` | no |
| `NextToken` | `string` | no |

## ListBenefitApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Programs` | `List<string>` | no |
| `FulfillmentTypes` | `List<string>` | no |
| `BenefitIdentifiers` | `List<string>` | no |
| `Status` | `List<string>` | no |
| `Stages` | `List<string>` | no |
| `AssociatedResources` | `List<AssociatedResource>` | no |
| `AssociatedResourceArns` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BenefitApplicationSummaries` | `List<BenefitApplicationSummary>` | no |
| `NextToken` | `string` | no |

## ListBenefits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Programs` | `List<string>` | no |
| `FulfillmentTypes` | `List<string>` | no |
| `Status` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BenefitSummaries` | `List<BenefitSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## RecallBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | no |
| `Identifier` | `string` | yes |
| `Reason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubmitBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBenefitApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Identifier` | `string` | yes |
| `Revision` | `string` | yes |
| `BenefitApplicationDetails` | `Document` | no |
| `PartnerContacts` | `List<Contact>` | no |
| `FileDetails` | `List<FileInput>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Revision` | `string` | no |

