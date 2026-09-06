# Amazon Managed Grafana

API version: 2020-08-18. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/grafana/2020-08-18/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `licenseType` | `string` | yes |
| `grafanaToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## CreateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountAccessType` | `string` | yes |
| `clientToken` | `string` | no |
| `organizationRoleName` | `string` | no |
| `permissionType` | `string` | yes |
| `stackSetName` | `string` | no |
| `workspaceDataSources` | `List<string>` | no |
| `workspaceDescription` | `string` | no |
| `workspaceName` | `string` | no |
| `workspaceNotificationDestinations` | `List<string>` | no |
| `workspaceOrganizationalUnits` | `List<string>` | no |
| `workspaceRoleArn` | `string` | no |
| `authenticationProviders` | `List<string>` | yes |
| `tags` | `Map<string>` | no |
| `vpcConfiguration` | `VpcConfiguration` | no |
| `configuration` | `string` | no |
| `networkAccessControl` | `NetworkAccessConfiguration` | no |
| `grafanaVersion` | `string` | no |
| `ipAddressType` | `string` | no |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## CreateWorkspaceApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyName` | `string` | yes |
| `keyRole` | `string` | yes |
| `secondsToLive` | `integer` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyName` | `string` | yes |
| `key` | `string` | yes |
| `workspaceId` | `string` | yes |

## CreateWorkspaceServiceAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `grafanaRole` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `grafanaRole` | `string` | yes |
| `workspaceId` | `string` | yes |

## CreateWorkspaceServiceAccountToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `secondsToLive` | `integer` | yes |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceAccountToken` | `ServiceAccountTokenSummaryWithKey` | yes |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

## DeleteWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## DeleteWorkspaceApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyName` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `keyName` | `string` | yes |
| `workspaceId` | `string` | yes |

## DeleteWorkspaceServiceAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

## DeleteWorkspaceServiceAccountToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenId` | `string` | yes |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenId` | `string` | yes |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

## DescribeWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## DescribeWorkspaceAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authentication` | `AuthenticationDescription` | yes |

## DescribeWorkspaceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `string` | yes |
| `grafanaVersion` | `string` | no |

## DisassociateLicense

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `licenseType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## ListPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `userType` | `string` | no |
| `userId` | `string` | no |
| `groupId` | `string` | no |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `permissions` | `List<PermissionEntry>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `workspaceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `grafanaVersions` | `List<string>` | no |

## ListWorkspaceServiceAccountTokens

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serviceAccountTokens` | `List<ServiceAccountTokenSummary>` | yes |
| `serviceAccountId` | `string` | yes |
| `workspaceId` | `string` | yes |

## ListWorkspaceServiceAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serviceAccounts` | `List<ServiceAccountSummary>` | yes |
| `workspaceId` | `string` | yes |

## ListWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaces` | `List<WorkspaceSummary>` | yes |
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


## UpdatePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateInstructionBatch` | `List<UpdateInstruction>` | yes |
| `workspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<UpdateError>` | yes |

## UpdateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountAccessType` | `string` | no |
| `organizationRoleName` | `string` | no |
| `permissionType` | `string` | no |
| `stackSetName` | `string` | no |
| `workspaceDataSources` | `List<string>` | no |
| `workspaceDescription` | `string` | no |
| `workspaceId` | `string` | yes |
| `workspaceName` | `string` | no |
| `workspaceNotificationDestinations` | `List<string>` | no |
| `workspaceOrganizationalUnits` | `List<string>` | no |
| `workspaceRoleArn` | `string` | no |
| `vpcConfiguration` | `VpcConfiguration` | no |
| `removeVpcConfiguration` | `boolean` | no |
| `networkAccessControl` | `NetworkAccessConfiguration` | no |
| `removeNetworkAccessConfiguration` | `boolean` | no |
| `ipAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspace` | `WorkspaceDescription` | yes |

## UpdateWorkspaceAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workspaceId` | `string` | yes |
| `authenticationProviders` | `List<string>` | yes |
| `samlConfiguration` | `SamlConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authentication` | `AuthenticationDescription` | yes |

## UpdateWorkspaceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `string` | yes |
| `workspaceId` | `string` | yes |
| `grafanaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


