# AWS User Experience Customization

API version: 2024-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/uxc/2024-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetAccountCustomizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountColor` | `string` | no |
| `visibleServices` | `List<string>` | no |
| `visibleRegions` | `List<string>` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `services` | `List<string>` | no |

## UpdateAccountCustomizations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountColor` | `string` | no |
| `visibleServices` | `List<string>` | no |
| `visibleRegions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountColor` | `string` | no |
| `visibleServices` | `List<string>` | no |
| `visibleRegions` | `List<string>` | no |

