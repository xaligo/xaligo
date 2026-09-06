# AWS Single Sign-On

API version: 2019-06-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sso/2019-06-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## GetRoleCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleName` | `string` | yes |
| `accountId` | `string` | yes |
| `accessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `roleCredentials` | `RoleCredentials` | no |

## ListAccountRoles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `accessToken` | `string` | yes |
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `roleList` | `List<RoleInfo>` | no |

## ListAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `accessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `accountList` | `List<AccountInfo>` | no |

## Logout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


