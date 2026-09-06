# Amazon CloudDirectory

API version: 2017-01-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/clouddirectory/2017-01-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddFacetToObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `SchemaFacet` | `SchemaFacet` | yes |
| `ObjectAttributeList` | `List<AttributeKeyAndValue>` | no |
| `ObjectReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApplySchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublishedSchemaArn` | `string` | yes |
| `DirectoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppliedSchemaArn` | `string` | no |
| `DirectoryArn` | `string` | no |

## AttachObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ParentReference` | `ObjectReference` | yes |
| `ChildReference` | `ObjectReference` | yes |
| `LinkName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedObjectIdentifier` | `string` | no |

## AttachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `PolicyReference` | `ObjectReference` | yes |
| `ObjectReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AttachToIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `IndexReference` | `ObjectReference` | yes |
| `TargetReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedObjectIdentifier` | `string` | no |

## AttachTypedLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `SourceObjectReference` | `ObjectReference` | yes |
| `TargetObjectReference` | `ObjectReference` | yes |
| `TypedLinkFacet` | `TypedLinkSchemaAndFacetName` | yes |
| `Attributes` | `List<AttributeNameAndValue>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypedLinkSpecifier` | `TypedLinkSpecifier` | no |

## BatchRead

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `Operations` | `List<BatchReadOperation>` | yes |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Responses` | `List<BatchReadOperationResponse>` | no |

## BatchWrite

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `Operations` | `List<BatchWriteOperation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Responses` | `List<BatchWriteOperationResponse>` | no |

## CreateDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SchemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `Name` | `string` | yes |
| `ObjectIdentifier` | `string` | yes |
| `AppliedSchemaArn` | `string` | yes |

## CreateFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |
| `Attributes` | `List<FacetAttribute>` | no |
| `ObjectType` | `string` | no |
| `FacetStyle` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `OrderedIndexedAttributeList` | `List<AttributeKey>` | yes |
| `IsUnique` | `boolean` | yes |
| `ParentReference` | `ObjectReference` | no |
| `LinkName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectIdentifier` | `string` | no |

## CreateObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `SchemaFacets` | `List<SchemaFacet>` | yes |
| `ObjectAttributeList` | `List<AttributeKeyAndValue>` | no |
| `ParentReference` | `ObjectReference` | no |
| `LinkName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectIdentifier` | `string` | no |

## CreateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |

## CreateTypedLinkFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Facet` | `TypedLinkFacet` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

## DeleteFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |

## DeleteTypedLinkFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachFromIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `IndexReference` | `ObjectReference` | yes |
| `TargetReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetachedObjectIdentifier` | `string` | no |

## DetachObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ParentReference` | `ObjectReference` | yes |
| `LinkName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetachedObjectIdentifier` | `string` | no |

## DetachPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `PolicyReference` | `ObjectReference` | yes |
| `ObjectReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetachTypedLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `TypedLinkSpecifier` | `TypedLinkSpecifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisableDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

## EnableDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

## GetAppliedSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppliedSchemaArn` | `string` | no |

## GetDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Directory` | `Directory` | yes |

## GetFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Facet` | `Facet` | no |

## GetLinkAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `TypedLinkSpecifier` | `TypedLinkSpecifier` | yes |
| `AttributeNames` | `List<string>` | yes |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<AttributeKeyAndValue>` | no |

## GetObjectAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `ConsistencyLevel` | `string` | no |
| `SchemaFacet` | `SchemaFacet` | yes |
| `AttributeNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<AttributeKeyAndValue>` | no |

## GetObjectInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaFacets` | `List<SchemaFacet>` | no |
| `ObjectIdentifier` | `string` | no |

## GetSchemaAsJson

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Document` | `string` | no |

## GetTypedLinkFacetInformation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityAttributeOrder` | `List<string>` | no |

## ListAppliedSchemaArns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `SchemaArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAttachedIndices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `TargetReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexAttachments` | `List<IndexAttachment>` | no |
| `NextToken` | `string` | no |

## ListDevelopmentSchemaArns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListDirectories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `state` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Directories` | `List<Directory>` | yes |
| `NextToken` | `string` | no |

## ListFacetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<FacetAttribute>` | no |
| `NextToken` | `string` | no |

## ListFacetNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FacetNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListIncomingTypedLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `FilterAttributeRanges` | `List<TypedLinkAttributeRange>` | no |
| `FilterTypedLink` | `TypedLinkSchemaAndFacetName` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkSpecifiers` | `List<TypedLinkSpecifier>` | no |
| `NextToken` | `string` | no |

## ListIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `RangesOnIndexedValues` | `List<ObjectAttributeRange>` | no |
| `IndexReference` | `ObjectReference` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexAttachments` | `List<IndexAttachment>` | no |
| `NextToken` | `string` | no |

## ListManagedSchemaArns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListObjectAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |
| `FacetFilter` | `SchemaFacet` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<AttributeKeyAndValue>` | no |
| `NextToken` | `string` | no |

## ListObjectChildren

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Children` | `Map<string>` | no |
| `NextToken` | `string` | no |

## ListObjectParentPaths

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PathToObjectIdentifiersList` | `List<PathToObjectIdentifiers>` | no |
| `NextToken` | `string` | no |

## ListObjectParents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |
| `IncludeAllLinksToEachParent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parents` | `Map<string>` | no |
| `NextToken` | `string` | no |
| `ParentLinks` | `List<ObjectIdentifierAndLinkNameTuple>` | no |

## ListObjectPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedPolicyIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListOutgoingTypedLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `FilterAttributeRanges` | `List<TypedLinkAttributeRange>` | no |
| `FilterTypedLink` | `TypedLinkSchemaAndFacetName` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypedLinkSpecifiers` | `List<TypedLinkSpecifier>` | no |
| `NextToken` | `string` | no |

## ListPolicyAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `PolicyReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ConsistencyLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectIdentifiers` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPublishedSchemaArns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArns` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ListTypedLinkFacetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<TypedLinkAttributeDefinition>` | no |
| `NextToken` | `string` | no |

## ListTypedLinkFacetNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FacetNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## LookupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyToPathList` | `List<PolicyToPath>` | no |
| `NextToken` | `string` | no |

## PublishSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevelopmentSchemaArn` | `string` | yes |
| `Version` | `string` | yes |
| `MinorVersion` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublishedSchemaArn` | `string` | no |

## PutSchemaFromJson

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Document` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## RemoveFacetFromObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `SchemaFacet` | `SchemaFacet` | yes |
| `ObjectReference` | `ObjectReference` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |
| `AttributeUpdates` | `List<FacetAttributeUpdate>` | no |
| `ObjectType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLinkAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `TypedLinkSpecifier` | `TypedLinkSpecifier` | yes |
| `AttributeUpdates` | `List<LinkAttributeUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateObjectAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryArn` | `string` | yes |
| `ObjectReference` | `ObjectReference` | yes |
| `AttributeUpdates` | `List<ObjectAttributeUpdate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectIdentifier` | `string` | no |

## UpdateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | no |

## UpdateTypedLinkFacet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SchemaArn` | `string` | yes |
| `Name` | `string` | yes |
| `AttributeUpdates` | `List<TypedLinkFacetAttributeUpdate>` | yes |
| `IdentityAttributeOrder` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpgradeAppliedSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublishedSchemaArn` | `string` | yes |
| `DirectoryArn` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpgradedSchemaArn` | `string` | no |
| `DirectoryArn` | `string` | no |

## UpgradePublishedSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DevelopmentSchemaArn` | `string` | yes |
| `PublishedSchemaArn` | `string` | yes |
| `MinorVersion` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpgradedSchemaArn` | `string` | no |

