# AWS AI Ops

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/aiops/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateInvestigationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `roleArn` | `string` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `retentionInDays` | `long` | no |
| `tags` | `Map<string>` | no |
| `tagKeyBoundaries` | `List<string>` | no |
| `chatbotNotificationChannel` | `Map<List<string>>` | no |
| `isCloudTrailEventHistoryEnabled` | `boolean` | no |
| `crossAccountConfigurations` | `List<CrossAccountConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | no |

## DeleteInvestigationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInvestigationGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetInvestigationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdBy` | `string` | no |
| `createdAt` | `long` | no |
| `lastModifiedBy` | `string` | no |
| `lastModifiedAt` | `long` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `roleArn` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `retentionInDays` | `long` | no |
| `chatbotNotificationChannel` | `Map<List<string>>` | no |
| `tagKeyBoundaries` | `List<string>` | no |
| `isCloudTrailEventHistoryEnabled` | `boolean` | no |
| `crossAccountConfigurations` | `List<CrossAccountConfiguration>` | no |

## GetInvestigationGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `investigationGroupArn` | `string` | no |
| `policy` | `string` | no |

## ListInvestigationGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `investigationGroups` | `List<ListInvestigationGroupsModel>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutInvestigationGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `investigationGroupArn` | `string` | no |

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


## UpdateInvestigationGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `roleArn` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `tagKeyBoundaries` | `List<string>` | no |
| `chatbotNotificationChannel` | `Map<List<string>>` | no |
| `isCloudTrailEventHistoryEnabled` | `boolean` | no |
| `crossAccountConfigurations` | `List<CrossAccountConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


