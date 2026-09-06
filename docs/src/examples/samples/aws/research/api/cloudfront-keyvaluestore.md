# Amazon CloudFront KeyValueStore

API version: 2022-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudfront-keyvaluestore/2022-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KvsARN` | `string` | yes |
| `Key` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ItemCount` | `integer` | yes |
| `TotalSizeInBytes` | `long` | yes |
| `ETag` | `string` | yes |

## DescribeKeyValueStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KvsARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ItemCount` | `integer` | yes |
| `TotalSizeInBytes` | `long` | yes |
| `KvsARN` | `string` | yes |
| `Created` | `timestamp` | yes |
| `ETag` | `string` | yes |
| `LastModified` | `timestamp` | no |
| `Status` | `string` | no |
| `FailureReason` | `string` | no |

## GetKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KvsARN` | `string` | yes |
| `Key` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | yes |
| `Value` | `string` | yes |
| `ItemCount` | `integer` | yes |
| `TotalSizeInBytes` | `long` | yes |

## ListKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KvsARN` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Items` | `List<ListKeysResponseListItem>` | no |

## PutKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Key` | `string` | yes |
| `Value` | `string` | yes |
| `KvsARN` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ItemCount` | `integer` | yes |
| `TotalSizeInBytes` | `long` | yes |
| `ETag` | `string` | yes |

## UpdateKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KvsARN` | `string` | yes |
| `IfMatch` | `string` | yes |
| `Puts` | `List<PutKeyRequestListItem>` | no |
| `Deletes` | `List<DeleteKeyRequestListItem>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ItemCount` | `integer` | yes |
| `TotalSizeInBytes` | `long` | yes |
| `ETag` | `string` | yes |

