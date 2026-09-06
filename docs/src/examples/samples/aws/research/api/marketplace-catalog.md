# AWS Marketplace Catalog Service

API version: 2018-09-17. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/marketplace-catalog/2018-09-17/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDescribeEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityRequestList` | `List<EntityRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityDetails` | `Map<EntityDetail>` | no |
| `Errors` | `Map<BatchDescribeErrorDetail>` | no |

## CancelChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ChangeSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetId` | `string` | no |
| `ChangeSetArn` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAssessment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `AssessmentIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentArn` | `string` | no |
| `AssessmentId` | `string` | no |
| `FrameworkId` | `string` | no |
| `AssessmentTargetSummary` | `AssessmentTargetSummary` | no |
| `FrameworkSummary` | `FrameworkSummary` | no |
| `AssessmentResult` | `string` | no |
| `CreatedAt` | `string` | no |
| `ExpiresAt` | `string` | no |
| `ControlAssessments` | `List<ControlAssessment>` | no |
| `NextToken` | `string` | no |

## DescribeChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ChangeSetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetId` | `string` | no |
| `ChangeSetArn` | `string` | no |
| `ChangeSetName` | `string` | no |
| `Intent` | `string` | no |
| `StartTime` | `string` | no |
| `EndTime` | `string` | no |
| `Status` | `string` | no |
| `FailureCode` | `string` | no |
| `FailureDescription` | `string` | no |
| `ChangeSet` | `List<ChangeSummary>` | no |

## DescribeEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EntityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntityType` | `string` | no |
| `EntityIdentifier` | `string` | no |
| `EntityArn` | `string` | no |
| `LastModifiedDate` | `string` | no |
| `Details` | `string` | no |
| `DetailsDocument` | `JsonDocumentType` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## ListAssessments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `FrameworkId` | `string` | no |
| `AssessmentTargetFilter` | `AssessmentTargetFilter` | no |
| `FrameworkFilters` | `FrameworkFilters` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssessmentSummaryList` | `List<AssessmentSummary>` | no |
| `NextToken` | `string` | no |

## ListChangeSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `FilterList` | `List<Filter>` | no |
| `Sort` | `Sort` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetSummaryList` | `List<ChangeSetSummaryListItem>` | no |
| `NextToken` | `string` | no |

## ListEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `EntityType` | `string` | yes |
| `FilterList` | `List<Filter>` | no |
| `Sort` | `Sort` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `OwnershipType` | `string` | no |
| `EntityTypeFilters` | `EntityTypeFilters` | no |
| `EntityTypeSort` | `EntityTypeSort` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitySummaryList` | `List<EntitySummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Catalog` | `string` | yes |
| `ChangeSet` | `List<Change>` | yes |
| `ChangeSetName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `ChangeSetTags` | `List<Tag>` | no |
| `Intent` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetId` | `string` | no |
| `ChangeSetArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


