# Amazon SimpleDB

API version: 2009-04-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sdb/2009-04-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Items` | `List<DeletableItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchPutAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Items` | `List<ReplaceableItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ItemName` | `string` | yes |
| `Attributes` | `List<Attribute>` | no |
| `Expected` | `UpdateCondition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DomainMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ItemCount` | `integer` | no |
| `ItemNamesSizeBytes` | `long` | no |
| `AttributeNameCount` | `integer` | no |
| `AttributeNamesSizeBytes` | `long` | no |
| `AttributeValueCount` | `integer` | no |
| `AttributeValuesSizeBytes` | `long` | no |
| `Timestamp` | `integer` | no |

## GetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ItemName` | `string` | yes |
| `AttributeNames` | `List<string>` | no |
| `ConsistentRead` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<Attribute>` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxNumberOfDomains` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## PutAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ItemName` | `string` | yes |
| `Attributes` | `List<ReplaceableAttribute>` | yes |
| `Expected` | `UpdateCondition` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Select

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SelectExpression` | `string` | yes |
| `NextToken` | `string` | no |
| `ConsistentRead` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Item>` | no |
| `NextToken` | `string` | no |

