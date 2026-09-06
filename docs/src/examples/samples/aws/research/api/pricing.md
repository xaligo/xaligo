# AWS Price List Service

API version: 2017-10-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pricing/2017-10-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DescribeServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | no |
| `FormatVersion` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Services` | `List<Service>` | no |
| `FormatVersion` | `string` | no |
| `NextToken` | `string` | no |

## GetAttributeValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `AttributeName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttributeValues` | `List<AttributeValue>` | no |
| `NextToken` | `string` | no |

## GetPriceListFileUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PriceListArn` | `string` | yes |
| `FileFormat` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |

## GetProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `Filters` | `List<Filter>` | no |
| `FormatVersion` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FormatVersion` | `string` | no |
| `PriceList` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPriceLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceCode` | `string` | yes |
| `EffectiveDate` | `timestamp` | yes |
| `RegionCode` | `string` | no |
| `CurrencyCode` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PriceLists` | `List<PriceList>` | no |
| `NextToken` | `string` | no |

