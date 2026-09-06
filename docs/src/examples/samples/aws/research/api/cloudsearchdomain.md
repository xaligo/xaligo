# Amazon CloudSearch Domain

API version: 2013-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudsearchdomain/2013-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## Search

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cursor` | `string` | no |
| `expr` | `string` | no |
| `facet` | `string` | no |
| `filterQuery` | `string` | no |
| `highlight` | `string` | no |
| `partial` | `boolean` | no |
| `query` | `string` | yes |
| `queryOptions` | `string` | no |
| `queryParser` | `string` | no |
| `return` | `string` | no |
| `size` | `long` | no |
| `sort` | `string` | no |
| `start` | `long` | no |
| `stats` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `SearchStatus` | no |
| `hits` | `Hits` | no |
| `facets` | `Map<BucketInfo>` | no |
| `stats` | `Map<FieldStats>` | no |

## Suggest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `query` | `string` | yes |
| `suggester` | `string` | yes |
| `size` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `SuggestStatus` | no |
| `suggest` | `SuggestModel` | no |

## UploadDocuments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `documents` | `blob` | yes |
| `contentType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `adds` | `long` | no |
| `deletes` | `long` | no |
| `warnings` | `List<DocumentServiceWarning>` | no |

