# Partner Central Revenue Measurement API

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/partnercentral-revenue-measurement/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateMarketplaceRevenueShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | no |
| `ProductId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductId` | `string` | yes |
| `Arn` | `string` | yes |
| `Catalog` | `string` | no |
| `ProductCode` | `string` | no |
| `ProductName` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `Revision` | `integer` | no |

## CreateMarketplaceRevenueShareAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductId` | `string` | yes |
| `ClientToken` | `string` | no |
| `EffectiveFrom` | `string` | yes |
| `EffectiveUntil` | `string` | no |
| `RevenueSharePercent` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MarketplaceRevenueShareAllocationId` | `string` | yes |
| `ProductId` | `string` | yes |
| `ProductName` | `string` | no |
| `Arn` | `string` | yes |
| `EffectiveFrom` | `string` | yes |
| `EffectiveUntil` | `string` | no |
| `RevenueSharePercent` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `LatestMarketplaceRevenueShareRevision` | `string` | no |

## CreateRevenueAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ClientToken` | `string` | no |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `TenancyModel` | `string` | yes |
| `ProductIdentifier` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `TenancyModel` | `string` | yes |
| `MarketplaceProduct` | `MarketplaceProductSummary` | no |
| `Revision` | `string` | no |

## GetMarketplaceRevenueShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductId` | `string` | yes |
| `Revision` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductId` | `string` | yes |
| `Arn` | `string` | yes |
| `Catalog` | `string` | yes |
| `ProductCode` | `string` | no |
| `ProductName` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `Revision` | `integer` | no |
| `LatestRevision` | `integer` | no |
| `TotalActiveMarketplaceRevenueShareAllocationCount` | `integer` | no |
| `TotalMarketplaceRevenueShareAllocationCount` | `integer` | no |

## GetMarketplaceRevenueShareAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductId` | `string` | yes |
| `MarketplaceRevenueShareAllocationId` | `string` | yes |
| `MarketplaceRevenueShareRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MarketplaceRevenueShareAllocationId` | `string` | yes |
| `ProductId` | `string` | yes |
| `ProductName` | `string` | no |
| `Arn` | `string` | yes |
| `EffectiveFrom` | `string` | yes |
| `EffectiveUntil` | `string` | no |
| `RevenueSharePercent` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `LatestMarketplaceRevenueShareRevision` | `string` | no |

## GetRevenueAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `Revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Id` | `string` | yes |
| `Catalog` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `TenancyModel` | `string` | yes |
| `MarketplaceProduct` | `MarketplaceProductSummary` | no |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `Revision` | `string` | no |
| `LatestRevision` | `string` | no |
| `EffectiveFrom` | `string` | no |
| `EffectiveUntil` | `string` | no |
| `TotalActiveRevenueAttributionAllocationCount` | `integer` | no |
| `TotalRevenueAttributionAllocationCount` | `integer` | no |

## GetRevenueAttributionAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `RevenueAttributionIdentifier` | `string` | yes |
| `RevenueAttributionAllocationId` | `string` | yes |
| `RevenueAttributionRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevenueAttributionAllocationId` | `string` | yes |
| `RevenueAttributionIdentifier` | `string` | yes |
| `EntityType` | `string` | yes |
| `EntityIdentifier` | `string` | yes |
| `EntityName` | `string` | no |
| `CustomerAwsAccountId` | `string` | yes |
| `RevenueSharePercent` | `string` | yes |
| `EffectiveFrom` | `string` | yes |
| `EffectiveUntil` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedDate` | `timestamp` | yes |
| `LastModifiedDate` | `timestamp` | yes |
| `RevenueAttributionRevision` | `string` | yes |
| `RevenueAttributionLatestRevision` | `string` | yes |

## GetRevenueAttributionAllocationsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `RevenueAttributionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `Catalog` | `string` | yes |
| `RevenueAttributionArn` | `string` | yes |
| `StartedAt` | `timestamp` | yes |
| `EndedAt` | `timestamp` | no |
| `TotalRevenueAttributionAllocationRecords` | `integer` | yes |
| `Description` | `string` | no |
| `RevenueAttributionLatestRevision` | `string` | no |
| `ErrorDetailList` | `List<RevenueAttributionAllocationErrorDetail>` | no |

## ListMarketplaceRevenueShareAllocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductId` | `string` | yes |
| `Status` | `string` | no |
| `AfterEffectiveFrom` | `string` | no |
| `BeforeEffectiveFrom` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `MarketplaceRevenueShareRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MarketplaceRevenueShareAllocationSummaries` | `List<MarketplaceRevenueShareAllocationSummary>` | yes |
| `NextToken` | `string` | no |

## ListMarketplaceRevenueShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductIds` | `List<string>` | no |
| `ProductCodes` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MarketplaceRevenueShareSummaries` | `List<MarketplaceRevenueShareSummary>` | yes |
| `NextToken` | `string` | no |

## ListRevenueAttributionAllocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `RevenueAttributionIdentifier` | `string` | yes |
| `EntityTypeFilters` | `List<string>` | no |
| `EntityIdentifierFilters` | `List<string>` | no |
| `CustomerAwsAccountIdFilters` | `List<string>` | no |
| `StatusFilter` | `string` | no |
| `AfterEffectiveFrom` | `string` | no |
| `BeforeEffectiveFrom` | `string` | no |
| `AfterEffectiveUntil` | `string` | no |
| `BeforeEffectiveUntil` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `RevenueAttributionRevision` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevenueAttributionAllocationSummaries` | `List<RevenueAttributionAllocationSummary>` | yes |
| `NextToken` | `string` | no |

## ListRevenueAttributions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifiers` | `List<string>` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RevenueAttributionSummaries` | `List<AttributionSummary>` | no |
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

## StartRevenueAttributionAllocationsTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `RevenueAttributionIdentifier` | `string` | yes |
| `RevenueAttributionRevision` | `string` | yes |
| `RevenueShareAllocations` | `List<RevenueShareAllocation>` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |
| `Status` | `string` | yes |
| `Catalog` | `string` | yes |
| `RevenueAttributionArn` | `string` | yes |
| `StartedAt` | `timestamp` | yes |
| `TotalRevenueAttributionAllocationRecords` | `integer` | yes |

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


## UpdateMarketplaceRevenueShareAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ProductId` | `string` | yes |
| `MarketplaceRevenueShareAllocationId` | `string` | yes |
| `MarketplaceRevenueShareRevision` | `string` | yes |
| `ClientToken` | `string` | no |
| `EffectiveFrom` | `string` | no |
| `EffectiveUntil` | `string` | no |
| `RevenueSharePercent` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MarketplaceRevenueShareAllocationId` | `string` | yes |
| `ProductId` | `string` | yes |
| `ProductName` | `string` | no |
| `Arn` | `string` | yes |
| `EffectiveFrom` | `string` | yes |
| `EffectiveUntil` | `string` | no |
| `RevenueSharePercent` | `string` | yes |
| `Status` | `string` | yes |
| `CreatedDate` | `timestamp` | no |
| `LastModifiedDate` | `timestamp` | no |
| `LatestMarketplaceRevenueShareRevision` | `string` | no |

## UpdateRevenueAttribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `Identifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Revision` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `Description` | `string` | no |
| `LastModifiedDate` | `timestamp` | yes |
| `LatestRevision` | `string` | no |

