# Account Access

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/account-access/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identitySource` | `IdentitySource` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |

## CreateEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `entitlement` | `Entitlement` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entitlementId` | `string` | yes |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `entitlementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identitySource` | `IdentitySourceDetails` | yes |
| `status` | `string` | yes |
| `tenantId` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |
| `error` | `ErrorDetails` | no |

## GetEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `entitlementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `entitlementId` | `string` | yes |
| `entitlement` | `EntitlementDetails` | yes |
| `createdAt` | `timestamp` | yes |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applications` | `List<ApplicationSummary>` | yes |
| `nextToken` | `string` | no |

## ListEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `filter` | `EntitlementFilter` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `entitlements` | `List<EntitlementsListMember>` | yes |
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


