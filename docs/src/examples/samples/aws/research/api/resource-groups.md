# AWS Resource Groups

API version: 2017-11-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/resource-groups/2017-11-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelTagSyncTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ResourceQuery` | `ResourceQuery` | no |
| `Tags` | `Map<string>` | no |
| `Configuration` | `List<GroupConfigurationItem>` | no |
| `Criticality` | `integer` | no |
| `Owner` | `string` | no |
| `DisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |
| `ResourceQuery` | `ResourceQuery` | no |
| `Tags` | `Map<string>` | no |
| `GroupConfiguration` | `GroupConfiguration` | no |

## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountSettings` | `AccountSettings` | no |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## GetGroupConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupConfiguration` | `GroupConfiguration` | no |

## GetGroupQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupQuery` | `GroupQuery` | no |

## GetTagSyncTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupArn` | `string` | no |
| `GroupName` | `string` | no |
| `TaskArn` | `string` | no |
| `TagKey` | `string` | no |
| `TagValue` | `string` | no |
| `ResourceQuery` | `ResourceQuery` | no |
| `RoleArn` | `string` | no |
| `Status` | `string` | no |
| `ErrorMessage` | `string` | no |
| `CreatedAt` | `timestamp` | no |

## GetTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Tags` | `Map<string>` | no |

## GroupResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Succeeded` | `List<string>` | no |
| `Failed` | `List<FailedResource>` | no |
| `Pending` | `List<PendingResource>` | no |

## ListGroupResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |
| `Filters` | `List<ResourceFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<ListGroupResourcesItem>` | no |
| `ResourceIdentifiers` | `List<ResourceIdentifier>` | no |
| `NextToken` | `string` | no |
| `QueryErrors` | `List<QueryError>` | no |

## ListGroupingStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | yes |
| `MaxResults` | `integer` | no |
| `Filters` | `List<ListGroupingStatusesFilter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | no |
| `GroupingStatuses` | `List<GroupingStatusesItem>` | no |
| `NextToken` | `string` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<GroupFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIdentifiers` | `List<GroupIdentifier>` | no |
| `Groups` | `List<Group>` | no |
| `NextToken` | `string` | no |

## ListTagSyncTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ListTagSyncTasksFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagSyncTasks` | `List<TagSyncTaskItem>` | no |
| `NextToken` | `string` | no |

## PutGroupConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | no |
| `Configuration` | `List<GroupConfigurationItem>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceQuery` | `ResourceQuery` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifiers` | `List<ResourceIdentifier>` | no |
| `NextToken` | `string` | no |
| `QueryErrors` | `List<QueryError>` | no |

## StartTagSyncTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | yes |
| `TagKey` | `string` | no |
| `TagValue` | `string` | no |
| `ResourceQuery` | `ResourceQuery` | no |
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupArn` | `string` | no |
| `GroupName` | `string` | no |
| `TaskArn` | `string` | no |
| `TagKey` | `string` | no |
| `TagValue` | `string` | no |
| `ResourceQuery` | `ResourceQuery` | no |
| `RoleArn` | `string` | no |

## Tag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Tags` | `Map<string>` | no |

## UngroupResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Succeeded` | `List<string>` | no |
| `Failed` | `List<FailedResource>` | no |
| `Pending` | `List<PendingResource>` | no |

## Untag

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Keys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Keys` | `List<string>` | no |

## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupLifecycleEventsDesiredStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountSettings` | `AccountSettings` | no |

## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |
| `Description` | `string` | no |
| `Criticality` | `integer` | no |
| `Owner` | `string` | no |
| `DisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Group` | `Group` | no |

## UpdateGroupQuery

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | no |
| `Group` | `string` | no |
| `ResourceQuery` | `ResourceQuery` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupQuery` | `GroupQuery` | no |

