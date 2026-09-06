# Amazon Recycle Bin

API version: 2021-06-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rbin/2021-06-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RetentionPeriod` | `RetentionPeriod` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ResourceType` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |
| `LockConfiguration` | `LockConfiguration` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ResourceType` | `string` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `Status` | `string` | no |
| `LockConfiguration` | `LockConfiguration` | no |
| `LockState` | `string` | no |
| `RuleArn` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `Description` | `string` | no |
| `ResourceType` | `string` | no |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `Status` | `string` | no |
| `LockConfiguration` | `LockConfiguration` | no |
| `LockState` | `string` | no |
| `LockEndTime` | `timestamp` | no |
| `RuleArn` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ResourceType` | `string` | yes |
| `ResourceTags` | `List<ResourceTag>` | no |
| `LockState` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<RuleSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## LockRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `LockConfiguration` | `LockConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `Description` | `string` | no |
| `ResourceType` | `string` | no |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `Status` | `string` | no |
| `LockConfiguration` | `LockConfiguration` | no |
| `LockState` | `string` | no |
| `RuleArn` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnlockRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `Description` | `string` | no |
| `ResourceType` | `string` | no |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `Status` | `string` | no |
| `LockConfiguration` | `LockConfiguration` | no |
| `LockState` | `string` | no |
| `LockEndTime` | `timestamp` | no |
| `RuleArn` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `Description` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | no |
| `RetentionPeriod` | `RetentionPeriod` | no |
| `Description` | `string` | no |
| `ResourceType` | `string` | no |
| `ResourceTags` | `List<ResourceTag>` | no |
| `Status` | `string` | no |
| `LockState` | `string` | no |
| `LockEndTime` | `timestamp` | no |
| `RuleArn` | `string` | no |
| `ExcludeResourceTags` | `List<ResourceTag>` | no |

