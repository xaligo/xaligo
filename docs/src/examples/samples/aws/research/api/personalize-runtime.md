# Amazon Personalize Runtime

API version: 2018-05-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/personalize-runtime/2018-05-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetActionRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | no |
| `userId` | `string` | no |
| `numResults` | `integer` | no |
| `filterArn` | `string` | no |
| `filterValues` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionList` | `List<PredictedAction>` | no |
| `recommendationId` | `string` | no |

## GetPersonalizedRanking

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | yes |
| `inputList` | `List<string>` | yes |
| `userId` | `string` | yes |
| `context` | `Map<string>` | no |
| `filterArn` | `string` | no |
| `filterValues` | `Map<string>` | no |
| `metadataColumns` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `personalizedRanking` | `List<PredictedItem>` | no |
| `recommendationId` | `string` | no |

## GetRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `campaignArn` | `string` | no |
| `itemId` | `string` | no |
| `userId` | `string` | no |
| `numResults` | `integer` | no |
| `context` | `Map<string>` | no |
| `filterArn` | `string` | no |
| `filterValues` | `Map<string>` | no |
| `recommenderArn` | `string` | no |
| `promotions` | `List<Promotion>` | no |
| `metadataColumns` | `Map<List<string>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `itemList` | `List<PredictedItem>` | no |
| `recommendationId` | `string` | no |

