# AWS Elemental MediaStore Data Plane

API version: 2017-09-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediastore-data/2017-09-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ETag` | `string` | no |
| `ContentType` | `string` | no |
| `ContentLength` | `long` | no |
| `CacheControl` | `string` | no |
| `LastModified` | `timestamp` | no |

## GetObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | yes |
| `Range` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `blob` | no |
| `CacheControl` | `string` | no |
| `ContentRange` | `string` | no |
| `ContentLength` | `long` | no |
| `ContentType` | `string` | no |
| `ETag` | `string` | no |
| `LastModified` | `timestamp` | no |
| `StatusCode` | `integer` | yes |

## ListItems

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Path` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Item>` | no |
| `NextToken` | `string` | no |

## PutObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `blob` | yes |
| `Path` | `string` | yes |
| `ContentType` | `string` | no |
| `CacheControl` | `string` | no |
| `StorageClass` | `string` | no |
| `UploadAvailability` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentSHA256` | `string` | no |
| `ETag` | `string` | no |
| `StorageClass` | `string` | no |

