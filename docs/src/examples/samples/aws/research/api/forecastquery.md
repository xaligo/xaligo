# Amazon Forecast Query Service

API version: 2018-06-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/forecastquery/2018-06-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## QueryForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ForecastArn` | `string` | yes |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Filters` | `Map<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Forecast` | `Forecast` | no |

## QueryWhatIfForecast

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WhatIfForecastArn` | `string` | yes |
| `StartDate` | `string` | no |
| `EndDate` | `string` | no |
| `Filters` | `Map<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Forecast` | `Forecast` | no |

