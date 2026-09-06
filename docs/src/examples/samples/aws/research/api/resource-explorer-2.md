# AWS Resource Explorer

API version: 2022-07-28. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/resource-explorer-2/2022-07-28/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateDefaultView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | no |

## BatchGetView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Views` | `List<View>` | no |
| `Errors` | `List<BatchGetViewError>` | no |

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `State` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## CreateResourceExplorerSetup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionList` | `List<string>` | yes |
| `AggregatorRegions` | `List<string>` | no |
| `ViewName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

## CreateView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ViewName` | `string` | yes |
| `IncludedProperties` | `List<IncludedProperty>` | no |
| `Scope` | `string` | no |
| `Filters` | `SearchFilter` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `State` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |

## DeleteResourceExplorerSetup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegionList` | `List<string>` | no |
| `DeleteInAllRegions` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |

## DeleteView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | no |

## DisassociateDefaultView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountLevelServiceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OrgConfiguration` | `OrgConfiguration` | no |

## GetDefaultView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | no |

## GetIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `State` | `string` | no |
| `ReplicatingFrom` | `List<string>` | no |
| `ReplicatingTo` | `List<string>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetManagedView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedViewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedView` | `ManagedView` | no |

## GetResourceExplorerSetup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<RegionStatus>` | no |
| `NextToken` | `string` | no |

## GetServiceIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |

## GetServiceView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceViewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `ServiceView` | yes |

## GetView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |
| `Tags` | `Map<string>` | no |

## ListIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Regions` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Indexes` | `List<Index>` | no |
| `NextToken` | `string` | no |

## ListIndexesForMembers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountIdList` | `List<string>` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Indexes` | `List<MemberIndex>` | no |
| `NextToken` | `string` | no |

## ListManagedViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ServicePrincipal` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ManagedViews` | `List<string>` | no |

## ListResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `SearchFilter` | no |
| `MaxResults` | `integer` | no |
| `ViewArn` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<Resource>` | no |
| `NextToken` | `string` | no |
| `ViewArn` | `string` | no |

## ListServiceIndexes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Regions` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Indexes` | `List<Index>` | no |
| `NextToken` | `string` | no |

## ListServiceViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServiceViews` | `List<string>` | no |

## ListStreamingAccessForServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingAccessForServices` | `List<StreamingAccessDetails>` | yes |
| `NextToken` | `string` | no |

## ListSupportedResourceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTypes` | `List<SupportedResourceType>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Views` | `List<string>` | no |
| `NextToken` | `string` | no |

## Search

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueryString` | `string` | yes |
| `MaxResults` | `integer` | no |
| `ViewArn` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<Resource>` | no |
| `NextToken` | `string` | no |
| `ViewArn` | `string` | no |
| `Count` | `ResourceCount` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

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


## UpdateIndexType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Type` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `State` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |

## UpdateView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewArn` | `string` | yes |
| `IncludedProperties` | `List<IncludedProperty>` | no |
| `Filters` | `SearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

