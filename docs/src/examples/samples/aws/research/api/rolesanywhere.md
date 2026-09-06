# IAM Roles Anywhere

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rolesanywhere/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `requireInstanceProperties` | `boolean` | no |
| `sessionPolicy` | `string` | no |
| `roleArns` | `List<string>` | yes |
| `managedPolicyArns` | `List<string>` | no |
| `durationSeconds` | `integer` | no |
| `enabled` | `boolean` | no |
| `tags` | `List<Tag>` | no |
| `acceptRoleSessionName` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## CreateTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `source` | `Source` | yes |
| `enabled` | `boolean` | no |
| `tags` | `List<Tag>` | no |
| `notificationSettings` | `List<NotificationSetting>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## DeleteAttributeMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `certificateField` | `string` | yes |
| `specifiers` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | yes |

## DeleteCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## DeleteTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## DisableCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## DisableProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## DisableTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## EnableCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## EnableProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## EnableTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## GetCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crlId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## GetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## GetSubject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subjectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subject` | `SubjectDetail` | no |

## GetTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## ImportCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `crlData` | `blob` | yes |
| `enabled` | `boolean` | no |
| `tags` | `List<Tag>` | no |
| `trustAnchorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## ListCrls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `crls` | `List<CrlDetail>` | no |

## ListProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `profiles` | `List<ProfileDetail>` | no |

## ListSubjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subjects` | `List<SubjectSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## ListTrustAnchors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `pageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `trustAnchors` | `List<TrustAnchorDetail>` | no |

## PutAttributeMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `certificateField` | `string` | yes |
| `mappingRules` | `List<MappingRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | yes |

## PutNotificationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |
| `notificationSettings` | `List<NotificationSetting>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## ResetNotificationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |
| `notificationSettingKeys` | `List<NotificationSettingKey>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateCrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crlId` | `string` | yes |
| `name` | `string` | no |
| `crlData` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `crl` | `CrlDetail` | yes |

## UpdateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `name` | `string` | no |
| `sessionPolicy` | `string` | no |
| `roleArns` | `List<string>` | no |
| `managedPolicyArns` | `List<string>` | no |
| `durationSeconds` | `integer` | no |
| `acceptRoleSessionName` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profile` | `ProfileDetail` | no |

## UpdateTrustAnchor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchorId` | `string` | yes |
| `name` | `string` | no |
| `source` | `Source` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `trustAnchor` | `TrustAnchorDetail` | yes |

