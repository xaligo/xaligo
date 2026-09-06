# AWS Migration Hub Config

API version: 2019-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/migrationhub-config/2019-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateHomeRegionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegion` | `string` | yes |
| `Target` | `Target` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegionControl` | `HomeRegionControl` | no |

## DeleteHomeRegionControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeHomeRegionControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ControlId` | `string` | no |
| `HomeRegion` | `string` | no |
| `Target` | `Target` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegionControls` | `List<HomeRegionControl>` | no |
| `NextToken` | `string` | no |

## GetHomeRegion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegion` | `string` | no |

