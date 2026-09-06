# AWS Resource Access Manager

API version: 2018-01-04. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ram/2018-01-04/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptResourceShareInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitationArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitation` | `ResourceShareInvitation` | no |
| `clientToken` | `string` | no |

## AssociateResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `resourceArns` | `List<string>` | no |
| `principals` | `List<string>` | no |
| `clientToken` | `string` | no |
| `sources` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareAssociations` | `List<ResourceShareAssociation>` | no |
| `clientToken` | `string` | no |

## AssociateResourceSharePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `permissionArn` | `string` | yes |
| `replace` | `boolean` | no |
| `clientToken` | `string` | no |
| `permissionVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |

## CreatePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `resourceType` | `string` | yes |
| `policyTemplate` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permission` | `ResourceSharePermissionSummary` | no |
| `clientToken` | `string` | no |

## CreatePermissionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `policyTemplate` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permission` | `ResourceSharePermissionDetail` | no |
| `clientToken` | `string` | no |

## CreateResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `resourceArns` | `List<string>` | no |
| `principals` | `List<string>` | no |
| `tags` | `List<Tag>` | no |
| `allowExternalPrincipals` | `boolean` | no |
| `clientToken` | `string` | no |
| `permissionArns` | `List<string>` | no |
| `sources` | `List<string>` | no |
| `resourceShareConfiguration` | `ResourceShareConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShare` | `ResourceShare` | no |
| `clientToken` | `string` | no |

## DeletePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |
| `permissionStatus` | `string` | no |

## DeletePermissionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `permissionVersion` | `integer` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |
| `permissionStatus` | `string` | no |

## DeleteResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |

## DisassociateResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `resourceArns` | `List<string>` | no |
| `principals` | `List<string>` | no |
| `clientToken` | `string` | no |
| `sources` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareAssociations` | `List<ResourceShareAssociation>` | no |
| `clientToken` | `string` | no |

## DisassociateResourceSharePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `permissionArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |

## EnableSharingWithAwsOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |

## GetPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `permissionVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permission` | `ResourceSharePermissionDetail` | no |

## GetResourcePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArns` | `List<string>` | yes |
| `principal` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<string>` | no |
| `nextToken` | `string` | no |

## GetResourceShareAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationType` | `string` | yes |
| `resourceShareArns` | `List<string>` | no |
| `resourceArn` | `string` | no |
| `principal` | `string` | no |
| `associationStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareAssociations` | `List<ResourceShareAssociation>` | no |
| `nextToken` | `string` | no |

## GetResourceShareInvitations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitationArns` | `List<string>` | no |
| `resourceShareArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitations` | `List<ResourceShareInvitation>` | no |
| `nextToken` | `string` | no |

## GetResourceShares

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArns` | `List<string>` | no |
| `resourceShareStatus` | `string` | no |
| `resourceOwner` | `string` | yes |
| `name` | `string` | no |
| `tagFilters` | `List<TagFilter>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `permissionArn` | `string` | no |
| `permissionVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShares` | `List<ResourceShare>` | no |
| `nextToken` | `string` | no |

## ListPendingInvitationResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitationArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `resourceRegionScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resources` | `List<Resource>` | no |
| `nextToken` | `string` | no |

## ListPermissionAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | no |
| `permissionVersion` | `integer` | no |
| `associationStatus` | `string` | no |
| `resourceType` | `string` | no |
| `featureSet` | `string` | no |
| `defaultVersion` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissions` | `List<AssociatedPermission>` | no |
| `nextToken` | `string` | no |

## ListPermissionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissions` | `List<ResourceSharePermissionSummary>` | no |
| `nextToken` | `string` | no |

## ListPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `permissionType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissions` | `List<ResourceSharePermissionSummary>` | no |
| `nextToken` | `string` | no |

## ListPrincipals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceOwner` | `string` | yes |
| `resourceArn` | `string` | no |
| `principals` | `List<string>` | no |
| `resourceType` | `string` | no |
| `resourceShareArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principals` | `List<Principal>` | no |
| `nextToken` | `string` | no |

## ListReplacePermissionAssociationsWork

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workIds` | `List<string>` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replacePermissionAssociationsWorks` | `List<ReplacePermissionAssociationsWork>` | no |
| `nextToken` | `string` | no |

## ListResourceSharePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissions` | `List<ResourceSharePermissionSummary>` | no |
| `nextToken` | `string` | no |

## ListResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `resourceRegionScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceTypes` | `List<ServiceNameAndResourceType>` | no |
| `nextToken` | `string` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceOwner` | `string` | yes |
| `principal` | `string` | no |
| `resourceType` | `string` | no |
| `resourceArns` | `List<string>` | no |
| `resourceShareArns` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `resourceRegionScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resources` | `List<Resource>` | no |
| `nextToken` | `string` | no |

## ListSourceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArns` | `List<string>` | no |
| `sourceId` | `string` | no |
| `sourceType` | `string` | no |
| `associationStatus` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceAssociations` | `List<AssociatedSource>` | no |
| `nextToken` | `string` | no |

## PromotePermissionCreatedFromPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permission` | `ResourceSharePermissionSummary` | no |
| `clientToken` | `string` | no |

## PromoteResourceShareCreatedFromPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |

## RejectResourceShareInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitationArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareInvitation` | `ResourceShareInvitation` | no |
| `clientToken` | `string` | no |

## ReplacePermissionAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fromPermissionArn` | `string` | yes |
| `fromPermissionVersion` | `integer` | no |
| `toPermissionArn` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replacePermissionAssociationsWork` | `ReplacePermissionAssociationsWork` | no |
| `clientToken` | `string` | no |

## SetDefaultPermissionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `permissionArn` | `string` | yes |
| `permissionVersion` | `integer` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `returnValue` | `boolean` | no |
| `clientToken` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | no |
| `tags` | `List<Tag>` | yes |
| `resourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | no |
| `tagKeys` | `List<string>` | yes |
| `resourceArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceShare

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShareArn` | `string` | yes |
| `name` | `string` | no |
| `allowExternalPrincipals` | `boolean` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceShare` | `ResourceShare` | no |
| `clientToken` | `string` | no |

