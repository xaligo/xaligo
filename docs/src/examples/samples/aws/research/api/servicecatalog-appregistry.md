# AWS Service Catalog App Registry

API version: 2020-06-24. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/servicecatalog-appregistry/2020-06-24/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `attributeGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |
| `attributeGroupArn` | `string` | no |

## AssociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `resourceType` | `string` | yes |
| `resource` | `string` | yes |
| `options` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |
| `resourceArn` | `string` | no |
| `options` | `List<string>` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `Application` | no |

## CreateAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `attributes` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `AttributeGroup` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `ApplicationSummary` | no |

## DeleteAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `AttributeGroupSummary` | no |

## DisassociateAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `attributeGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |
| `attributeGroupArn` | `string` | no |

## DisassociateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `resourceType` | `string` | yes |
| `resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |
| `resourceArn` | `string` | no |

## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastUpdateTime` | `timestamp` | no |
| `associatedResourceCount` | `integer` | no |
| `tags` | `Map<string>` | no |
| `integrations` | `Integrations` | no |
| `applicationTag` | `Map<string>` | no |

## GetAssociatedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `resourceType` | `string` | yes |
| `resource` | `string` | yes |
| `nextToken` | `string` | no |
| `resourceTagStatus` | `List<string>` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resource` | `Resource` | no |
| `options` | `List<string>` | no |
| `applicationTagResult` | `ApplicationTagResult` | no |

## GetAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `attributes` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastUpdateTime` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `createdBy` | `string` | no |

## GetConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `AppRegistryConfiguration` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applications` | `List<ApplicationSummary>` | no |
| `nextToken` | `string` | no |

## ListAssociatedAttributeGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroups` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListAssociatedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resources` | `List<ResourceInfo>` | no |
| `nextToken` | `string` | no |

## ListAttributeGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroups` | `List<AttributeGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListAttributeGroupsForApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroupsDetails` | `List<AttributeGroupDetails>` | no |
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

## PutConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `AppRegistryConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SyncResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |
| `resourceArn` | `string` | no |
| `actionTaken` | `string` | no |

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


## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `application` | `Application` | no |

## UpdateAttributeGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `attributes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributeGroup` | `AttributeGroup` | no |

