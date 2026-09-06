# Schemas

API version: 2019-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/schemas/2019-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `SourceArn` | `string` | yes |
| `CrossAccount` | `boolean` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DiscovererArn` | `string` | no |
| `DiscovererId` | `string` | no |
| `SourceArn` | `string` | no |
| `State` | `string` | no |
| `CrossAccount` | `boolean` | no |
| `Tags` | `Map<string>` | no |

## CreateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `RegistryName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `RegistryArn` | `string` | no |
| `RegistryName` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | yes |
| `Description` | `string` | no |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `LastModified` | `timestamp` | no |
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaVersion` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `VersionCreatedDate` | `timestamp` | no |

## DeleteDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeCodeBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Language` | `string` | yes |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationDate` | `timestamp` | no |
| `LastModified` | `timestamp` | no |
| `SchemaVersion` | `string` | no |
| `Status` | `string` | no |

## DescribeDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DiscovererArn` | `string` | no |
| `DiscovererId` | `string` | no |
| `SourceArn` | `string` | no |
| `State` | `string` | no |
| `CrossAccount` | `boolean` | no |
| `Tags` | `Map<string>` | no |

## DescribeRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `RegistryArn` | `string` | no |
| `RegistryName` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | no |
| `Description` | `string` | no |
| `LastModified` | `timestamp` | no |
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaVersion` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `VersionCreatedDate` | `timestamp` | no |

## ExportSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | no |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | no |
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaVersion` | `string` | no |
| `Type` | `string` | no |

## GetCodeBindingSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Language` | `string` | yes |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Body` | `blob` | no |

## GetDiscoveredSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<string>` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistryName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## ListDiscoverers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererIdPrefix` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `SourceArnPrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Discoverers` | `List<DiscovererSummary>` | no |
| `NextToken` | `string` | no |

## ListRegistries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `RegistryNamePrefix` | `string` | no |
| `Scope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Registries` | `List<RegistrySummary>` | no |

## ListSchemaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SchemaVersions` | `List<SchemaVersionSummary>` | no |

## ListSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `RegistryName` | `string` | yes |
| `SchemaNamePrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Schemas` | `List<SchemaSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PutCodeBinding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Language` | `string` | yes |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `SchemaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationDate` | `timestamp` | no |
| `LastModified` | `timestamp` | no |
| `SchemaVersion` | `string` | no |
| `Status` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | yes |
| `RegistryName` | `string` | no |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## SearchSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Keywords` | `string` | yes |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `RegistryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Schemas` | `List<SearchSchemaSummary>` | no |

## StartDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | no |
| `State` | `string` | no |

## StopDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DiscovererId` | `string` | no |
| `State` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateDiscoverer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DiscovererId` | `string` | yes |
| `CrossAccount` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DiscovererArn` | `string` | no |
| `DiscovererId` | `string` | no |
| `SourceArn` | `string` | no |
| `State` | `string` | no |
| `CrossAccount` | `boolean` | no |
| `Tags` | `Map<string>` | no |

## UpdateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `RegistryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `RegistryArn` | `string` | no |
| `RegistryName` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientTokenId` | `string` | no |
| `Content` | `string` | no |
| `Description` | `string` | no |
| `RegistryName` | `string` | yes |
| `SchemaName` | `string` | yes |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `LastModified` | `timestamp` | no |
| `SchemaArn` | `string` | no |
| `SchemaName` | `string` | no |
| `SchemaVersion` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `VersionCreatedDate` | `timestamp` | no |

