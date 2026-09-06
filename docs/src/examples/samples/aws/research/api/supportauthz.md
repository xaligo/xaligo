# SupportAuthZ

API version: 2026-06-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/supportauthz/2026-06-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSupportPermit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permit` | `Permit` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `signingKeyInfo` | `SigningKeyInfo` | yes |
| `supportCaseDisplayId` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `permit` | `Permit` | yes |
| `status` | `string` | yes |
| `signingKeyInfo` | `SigningKeyInfo` | yes |
| `createdAt` | `timestamp` | yes |
| `supportCaseDisplayId` | `string` | no |
| `tags` | `Map<string>` | no |

## DeleteSupportPermit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `supportPermitIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `permit` | `Permit` | yes |
| `status` | `string` | yes |
| `signingKeyInfo` | `SigningKeyInfo` | yes |
| `createdAt` | `timestamp` | yes |
| `supportCaseDisplayId` | `string` | no |

## GetAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `action` | `string` | yes |
| `service` | `string` | yes |
| `description` | `string` | yes |

## GetSupportPermit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `supportPermitIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `arn` | `string` | yes |
| `description` | `string` | no |
| `permit` | `Permit` | yes |
| `status` | `string` | yes |
| `signingKeyInfo` | `SigningKeyInfo` | yes |
| `createdAt` | `timestamp` | yes |
| `supportCaseDisplayId` | `string` | no |
| `tags` | `Map<string>` | no |

## ListActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `service` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `actionSummaries` | `List<ActionSummary>` | yes |
| `nextToken` | `string` | no |

## ListSupportPermitRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `supportCaseDisplayId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `supportPermitRequests` | `List<SupportPermitRequest>` | yes |
| `nextToken` | `string` | no |

## ListSupportPermits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `supportPermitStatuses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `supportPermits` | `List<SupportPermitSummary>` | yes |
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

## RejectSupportPermitRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestArn` | `string` | yes |

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


