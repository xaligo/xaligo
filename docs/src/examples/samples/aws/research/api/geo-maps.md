# Amazon Location Service Maps V2

API version: 2020-11-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/geo-maps/2020-11-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetGlyphs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FontStack` | `string` | yes |
| `FontUnicodeRange` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `ETag` | `string` | no |

## GetSprites

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileName` | `string` | yes |
| `Style` | `string` | yes |
| `ColorScheme` | `string` | yes |
| `Variant` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `ETag` | `string` | no |

## GetStaticMap

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BoundingBox` | `string` | no |
| `BoundedPositions` | `string` | no |
| `Center` | `string` | no |
| `ColorScheme` | `string` | no |
| `CompactOverlay` | `string` | no |
| `CropLabels` | `boolean` | no |
| `GeoJsonOverlay` | `string` | no |
| `Height` | `integer` | yes |
| `Key` | `string` | no |
| `LabelSize` | `string` | no |
| `Language` | `string` | no |
| `Padding` | `integer` | no |
| `PoliticalView` | `string` | no |
| `PointsOfInterests` | `string` | no |
| `Radius` | `long` | no |
| `FileName` | `string` | yes |
| `ScaleBarUnit` | `string` | no |
| `Style` | `string` | no |
| `Width` | `integer` | yes |
| `Zoom` | `float` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `ETag` | `string` | no |
| `PricingBucket` | `string` | yes |

## GetStyleDescriptor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Style` | `string` | yes |
| `ColorScheme` | `string` | no |
| `PoliticalView` | `string` | no |
| `Terrain` | `string` | no |
| `ContourDensity` | `string` | no |
| `Traffic` | `string` | no |
| `TravelModes` | `List<string>` | no |
| `Buildings` | `string` | no |
| `PoiDensity` | `string` | no |
| `PoiCategories` | `List<string>` | no |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `ETag` | `string` | no |

## GetTile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AdditionalFeatures` | `List<string>` | no |
| `Tileset` | `string` | yes |
| `Z` | `string` | yes |
| `X` | `string` | yes |
| `Y` | `string` | yes |
| `Key` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blob` | `blob` | no |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `ETag` | `string` | no |
| `PricingBucket` | `string` | yes |

