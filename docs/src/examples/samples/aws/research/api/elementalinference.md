# AWS Elemental Inference

API version: 2018-11-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/elementalinference/2018-11-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `associatedResourceName` | `string` | yes |
| `outputs` | `List<CreateOutput>` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |

## CreateDictionary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `language` | `string` | yes |
| `entries` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `language` | `string` | yes |
| `status` | `string` | yes |
| `references` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## CreateFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `accessRoleArn` | `string` | no |
| `outputs` | `List<CreateOutput>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `id` | `string` | yes |
| `dataEndpoints` | `List<string>` | yes |
| `outputs` | `List<GetOutput>` | yes |
| `accessRoleArn` | `string` | no |
| `status` | `string` | yes |
| `association` | `FeedAssociation` | no |
| `tags` | `Map<string>` | no |

## DeleteDictionary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |

## DeleteFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | yes |

## DisassociateFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `associatedResourceName` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `id` | `string` | yes |

## ExportDictionaryEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entries` | `string` | no |

## GetDictionary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `language` | `string` | yes |
| `status` | `string` | yes |
| `references` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## GetFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `id` | `string` | yes |
| `dataEndpoints` | `List<string>` | yes |
| `outputs` | `List<GetOutput>` | yes |
| `accessRoleArn` | `string` | no |
| `status` | `string` | yes |
| `association` | `FeedAssociation` | no |
| `tags` | `Map<string>` | no |

## GetFixture

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fixtureId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fixtureId` | `string` | yes |
| `name` | `string` | yes |
| `fixtureGroup` | `string` | no |
| `scheduledStart` | `timestamp` | no |
| `status` | `string` | yes |
| `competitors` | `List<Competitor>` | yes |

## ListDictionaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dictionaries` | `List<DictionarySummary>` | yes |
| `nextToken` | `string` | no |

## ListFeeds

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `feeds` | `List<FeedSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## SearchFixtures

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sport` | `string` | yes |
| `startDate` | `string` | yes |
| `endDate` | `string` | no |
| `filters` | `List<SearchFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fixtures` | `List<FixtureSummary>` | yes |
| `nextToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDictionary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | no |
| `language` | `string` | no |
| `entries` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `id` | `string` | yes |
| `language` | `string` | yes |
| `status` | `string` | yes |
| `references` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## UpdateFeed

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `accessRoleArn` | `string` | no |
| `id` | `string` | yes |
| `outputs` | `List<UpdateOutput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `id` | `string` | yes |
| `dataEndpoints` | `List<string>` | yes |
| `outputs` | `List<GetOutput>` | yes |
| `accessRoleArn` | `string` | no |
| `status` | `string` | yes |
| `association` | `FeedAssociation` | no |
| `tags` | `Map<string>` | no |

