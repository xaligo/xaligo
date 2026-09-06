# CloudWatch Observability Access Manager

API version: 2022-06-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/oam/2022-06-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelTemplate` | `string` | yes |
| `LinkConfiguration` | `LinkConfiguration` | no |
| `ResourceTypes` | `List<string>` | yes |
| `SinkIdentifier` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Label` | `string` | no |
| `LabelTemplate` | `string` | no |
| `LinkConfiguration` | `LinkConfiguration` | no |
| `ResourceTypes` | `List<string>` | no |
| `SinkArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateSink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## DeleteLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IncludeTags` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Label` | `string` | no |
| `LabelTemplate` | `string` | no |
| `LinkConfiguration` | `LinkConfiguration` | no |
| `ResourceTypes` | `List<string>` | no |
| `SinkArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetSink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IncludeTags` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetSinkPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SinkIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `SinkArn` | `string` | no |
| `SinkId` | `string` | no |

## ListAttachedLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SinkIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListAttachedLinksItem>` | yes |
| `NextToken` | `string` | no |

## ListLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListLinksItem>` | yes |
| `NextToken` | `string` | no |

## ListSinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListSinksItem>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PutSinkPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | yes |
| `SinkIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `SinkArn` | `string` | no |
| `SinkId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `IncludeTags` | `boolean` | no |
| `LinkConfiguration` | `LinkConfiguration` | no |
| `ResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Label` | `string` | no |
| `LabelTemplate` | `string` | no |
| `LinkConfiguration` | `LinkConfiguration` | no |
| `ResourceTypes` | `List<string>` | no |
| `SinkArn` | `string` | no |
| `Tags` | `Map<string>` | no |

