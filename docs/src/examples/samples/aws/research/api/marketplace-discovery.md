# AWS Marketplace Discovery

API version: 2026-02-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-discovery/2026-02-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `listingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associatedEntities` | `List<ListingAssociatedEntity>` | yes |
| `badges` | `List<ListingBadge>` | yes |
| `catalog` | `string` | yes |
| `categories` | `List<Category>` | yes |
| `fulfillmentOptionSummaries` | `List<FulfillmentOptionSummary>` | yes |
| `highlights` | `List<string>` | yes |
| `integrationGuide` | `string` | no |
| `listingId` | `string` | yes |
| `listingName` | `string` | yes |
| `logoThumbnailUrl` | `string` | yes |
| `longDescription` | `string` | yes |
| `pricingModels` | `List<PricingModel>` | yes |
| `pricingUnits` | `List<PricingUnit>` | yes |
| `promotionalMedia` | `List<PromotionalMedia>` | yes |
| `publisher` | `SellerInformation` | yes |
| `resources` | `List<Resource>` | yes |
| `reviewSummary` | `ReviewSummary` | no |
| `sellerEngagements` | `List<SellerEngagement>` | yes |
| `shortDescription` | `string` | yes |
| `useCases` | `List<UseCaseEntry>` | yes |

## GetOffer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerId` | `string` | yes |
| `catalog` | `string` | yes |
| `offerName` | `string` | no |
| `expirationTime` | `timestamp` | no |
| `availableFromTime` | `timestamp` | no |
| `sellerOfRecord` | `SellerInformation` | yes |
| `associatedEntities` | `List<OfferAssociatedEntity>` | yes |
| `agreementProposalId` | `string` | yes |
| `replacementAgreementId` | `string` | no |
| `pricingModel` | `PricingModel` | yes |
| `badges` | `List<PurchaseOptionBadge>` | yes |

## GetOfferSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerSetId` | `string` | yes |
| `catalog` | `string` | yes |
| `offerSetName` | `string` | no |
| `availableFromTime` | `timestamp` | no |
| `expirationTime` | `timestamp` | no |
| `buyerNotes` | `string` | no |
| `sellerOfRecord` | `SellerInformation` | yes |
| `badges` | `List<PurchaseOptionBadge>` | yes |
| `associatedEntities` | `List<OfferSetAssociatedEntity>` | yes |

## GetOfferTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerTerms` | `List<OfferTerm>` | yes |
| `nextToken` | `string` | no |

## GetProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `productId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `productId` | `string` | yes |
| `catalog` | `string` | yes |
| `productName` | `string` | yes |
| `manufacturer` | `SellerInformation` | yes |
| `deployedOnAws` | `string` | yes |
| `shortDescription` | `string` | yes |
| `longDescription` | `string` | yes |
| `logoThumbnailUrl` | `string` | yes |
| `fulfillmentOptionSummaries` | `List<FulfillmentOptionSummary>` | yes |
| `categories` | `List<Category>` | yes |
| `highlights` | `List<string>` | yes |
| `promotionalMedia` | `List<PromotionalMedia>` | yes |
| `resources` | `List<Resource>` | yes |
| `sellerEngagements` | `List<SellerEngagement>` | yes |

## ListFulfillmentOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `productId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fulfillmentOptions` | `List<FulfillmentOption>` | yes |
| `nextToken` | `string` | no |

## ListPurchaseOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<PurchaseOptionFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `purchaseOptions` | `List<PurchaseOptionSummary>` | no |
| `nextToken` | `string` | no |

## SearchFacets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchText` | `string` | no |
| `filters` | `List<SearchFilter>` | no |
| `facetTypes` | `List<string>` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `totalResults` | `long` | yes |
| `listingFacets` | `Map<List<ListingFacet>>` | yes |
| `nextToken` | `string` | no |

## SearchListings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `searchText` | `string` | no |
| `filters` | `List<SearchFilter>` | no |
| `maxResults` | `integer` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `totalResults` | `long` | yes |
| `listingSummaries` | `List<ListingSummary>` | yes |
| `nextToken` | `string` | no |

