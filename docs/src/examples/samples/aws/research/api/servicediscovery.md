# AWS Cloud Map

API version: 2017-03-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/servicediscovery/2017-03-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateHttpNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CreatorRequestId` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CreatePrivateDnsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CreatorRequestId` | `string` | no |
| `Description` | `string` | no |
| `Vpc` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `Properties` | `PrivateDnsNamespaceProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CreatePublicDnsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CreatorRequestId` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Properties` | `PublicDnsNamespaceProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `NamespaceId` | `string` | no |
| `CreatorRequestId` | `string` | no |
| `Description` | `string` | no |
| `DnsConfig` | `DnsConfig` | no |
| `HealthCheckConfig` | `HealthCheckConfig` | no |
| `HealthCheckCustomConfig` | `HealthCheckCustomConfig` | no |
| `Tags` | `List<Tag>` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | no |

## DeleteNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServiceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `Attributes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## DiscoverInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamespaceName` | `string` | yes |
| `ServiceName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `QueryParameters` | `Map<string>` | no |
| `OptionalParameters` | `Map<string>` | no |
| `HealthStatus` | `string` | no |
| `OwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<HttpInstanceSummary>` | no |
| `InstancesRevision` | `long` | no |

## DiscoverInstancesRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamespaceName` | `string` | yes |
| `ServiceName` | `string` | yes |
| `OwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstancesRevision` | `long` | no |

## GetInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceOwner` | `string` | no |
| `Instance` | `Instance` | no |

## GetInstancesHealthStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `Instances` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `Map<string>` | no |
| `NextToken` | `string` | no |

## GetNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespace` | `Namespace` | no |

## GetOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |
| `OwnerAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operation` | `Operation` | no |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | no |

## GetServiceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceAttributes` | `ServiceAttributes` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceOwner` | `string` | no |
| `Instances` | `List<InstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<NamespaceFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Namespaces` | `List<NamespaceSummary>` | no |
| `NextToken` | `string` | no |

## ListOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<OperationFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operations` | `List<OperationSummary>` | no |
| `NextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<ServiceFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Services` | `List<ServiceSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## RegisterInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `CreatorRequestId` | `string` | no |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHttpNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `UpdaterRequestId` | `string` | no |
| `Namespace` | `HttpNamespaceChange` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateInstanceCustomHealthStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePrivateDnsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `UpdaterRequestId` | `string` | no |
| `Namespace` | `PrivateDnsNamespaceChange` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdatePublicDnsNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `UpdaterRequestId` | `string` | no |
| `Namespace` | `PublicDnsNamespaceChange` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Service` | `ServiceChange` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateServiceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceId` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


