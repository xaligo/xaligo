# Amazon CodeGuru Reviewer

API version: 2019-09-19. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codeguru-reviewer/2019-09-19/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Repository` | `Repository` | yes |
| `ClientRequestToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `KMSKeyDetails` | `KMSKeyDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryAssociation` | `RepositoryAssociation` | no |
| `Tags` | `Map<string>` | no |

## CreateCodeReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `RepositoryAssociationArn` | `string` | yes |
| `Type` | `CodeReviewType` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReview` | `CodeReview` | no |

## DescribeCodeReview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReviewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReview` | `CodeReview` | no |

## DescribeRecommendationFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReviewArn` | `string` | yes |
| `RecommendationId` | `string` | yes |
| `UserId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationFeedback` | `RecommendationFeedback` | no |

## DescribeRepositoryAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryAssociation` | `RepositoryAssociation` | no |
| `Tags` | `Map<string>` | no |

## DisassociateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryAssociation` | `RepositoryAssociation` | no |
| `Tags` | `Map<string>` | no |

## ListCodeReviews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProviderTypes` | `List<string>` | no |
| `States` | `List<string>` | no |
| `RepositoryNames` | `List<string>` | no |
| `Type` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReviewSummaries` | `List<CodeReviewSummary>` | no |
| `NextToken` | `string` | no |

## ListRecommendationFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CodeReviewArn` | `string` | yes |
| `UserIds` | `List<string>` | no |
| `RecommendationIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationFeedbackSummaries` | `List<RecommendationFeedbackSummary>` | no |
| `NextToken` | `string` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CodeReviewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommendationSummaries` | `List<RecommendationSummary>` | no |
| `NextToken` | `string` | no |

## ListRepositoryAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProviderTypes` | `List<string>` | no |
| `States` | `List<string>` | no |
| `Names` | `List<string>` | no |
| `Owners` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryAssociationSummaries` | `List<RepositoryAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PutRecommendationFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeReviewArn` | `string` | yes |
| `RecommendationId` | `string` | yes |
| `Reactions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


