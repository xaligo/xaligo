# Amazon Location Service Places V2

API version: 2020-11-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/geo-places/2020-11-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## Autocomplete

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryText` | `string` | yes |
| `MaxResults` | `integer` | no |
| `BiasPosition` | `List<double>` | no |
| `Filter` | `AutocompleteFilter` | no |
| `PostalCodeMode` | `string` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<AutocompleteResultItem>` | no |

## Geocode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryText` | `string` | no |
| `QueryComponents` | `GeocodeQueryComponents` | no |
| `MaxResults` | `integer` | no |
| `BiasPosition` | `List<double>` | no |
| `Filter` | `GeocodeFilter` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `Key` | `string` | no |
| `PostalCodeMode` | `string` | no |
| `AddressTranslations` | `List<string>` | no |
| `AddressNamesMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<GeocodeResultItem>` | no |

## GetPlace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlaceId` | `string` | yes |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `Key` | `string` | no |
| `AddressNamesMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlaceId` | `string` | yes |
| `PlaceType` | `string` | yes |
| `Title` | `string` | yes |
| `PricingBucket` | `string` | yes |
| `Address` | `Address` | no |
| `AddressNumberCorrected` | `boolean` | no |
| `PostalCodeDetails` | `List<PostalCodeDetails>` | no |
| `Position` | `List<double>` | no |
| `MapView` | `List<double>` | no |
| `Categories` | `List<Category>` | no |
| `FoodTypes` | `List<FoodType>` | no |
| `BusinessChains` | `List<BusinessChain>` | no |
| `Contacts` | `Contacts` | no |
| `OpeningHours` | `List<OpeningHours>` | no |
| `AccessPoints` | `List<AccessPoint>` | no |
| `AccessRestrictions` | `List<AccessRestriction>` | no |
| `TimeZone` | `TimeZone` | no |
| `PoliticalView` | `string` | no |
| `Phonemes` | `PhonemeDetails` | no |
| `MainAddress` | `RelatedPlace` | no |
| `SecondaryAddresses` | `List<RelatedPlace>` | no |
| `PlaceAttributes` | `List<string>` | no |
| `EstimatedPointAddress` | `boolean` | no |
| `CrossReferences` | `List<CrossReference>` | no |

## ReverseGeocode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryPosition` | `List<double>` | yes |
| `QueryRadius` | `long` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `ReverseGeocodeFilter` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `Key` | `string` | no |
| `Heading` | `double` | no |
| `AddressNamesMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<ReverseGeocodeResultItem>` | no |

## SearchNearby

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryPosition` | `List<double>` | yes |
| `QueryRadius` | `long` | no |
| `MaxResults` | `integer` | no |
| `Filter` | `SearchNearbyFilter` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `NextToken` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<SearchNearbyResultItem>` | no |
| `NextToken` | `string` | no |

## SearchText

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryText` | `string` | no |
| `QueryId` | `string` | no |
| `MaxResults` | `integer` | no |
| `BiasPosition` | `List<double>` | no |
| `Filter` | `SearchTextFilter` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `NextToken` | `string` | no |
| `TravelMode` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<SearchTextResultItem>` | no |
| `NextToken` | `string` | no |

## Suggest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryText` | `string` | yes |
| `MaxResults` | `integer` | no |
| `MaxQueryRefinements` | `integer` | no |
| `BiasPosition` | `List<double>` | no |
| `Filter` | `SuggestFilter` | no |
| `AdditionalFeatures` | `List<string>` | no |
| `Language` | `string` | no |
| `PoliticalView` | `string` | no |
| `IntendedUse` | `string` | no |
| `TravelMode` | `string` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PricingBucket` | `string` | yes |
| `ResultItems` | `List<SuggestResultItem>` | no |
| `QueryRefinements` | `List<QueryRefinement>` | no |

