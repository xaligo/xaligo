# AWS Sustainability

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sustainability/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetEstimatedCarbonEmissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `TimePeriod` | yes |
| `GroupBy` | `List<string>` | no |
| `FilterBy` | `FilterExpression` | no |
| `EmissionsTypes` | `List<string>` | no |
| `Granularity` | `string` | no |
| `GranularityConfiguration` | `GranularityConfiguration` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<EstimatedCarbonEmissions>` | yes |
| `NextToken` | `string` | no |

## GetEstimatedCarbonEmissionsDimensionValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `TimePeriod` | yes |
| `Dimensions` | `List<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<DimensionEntry>` | no |
| `NextToken` | `string` | no |

## GetEstimatedWaterAllocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `TimePeriod` | yes |
| `GroupBy` | `List<string>` | no |
| `FilterBy` | `FilterExpression` | no |
| `AllocationTypes` | `List<string>` | no |
| `Granularity` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<EstimatedWaterAllocation>` | yes |
| `NextToken` | `string` | no |

## GetEstimatedWaterAllocationDimensionValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TimePeriod` | `TimePeriod` | yes |
| `Dimensions` | `List<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<DimensionEntry>` | yes |
| `NextToken` | `string` | no |

