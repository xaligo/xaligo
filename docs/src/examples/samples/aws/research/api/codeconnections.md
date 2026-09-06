# AWS CodeConnections

API version: 2023-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codeconnections/2023-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProviderType` | `string` | no |
| `ConnectionName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `HostArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

## CreateHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ProviderType` | `string` | yes |
| `ProviderEndpoint` | `string` | yes |
| `VpcConfiguration` | `VpcConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

## CreateRepositoryLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | yes |
| `OwnerId` | `string` | yes |
| `RepositoryName` | `string` | yes |
| `EncryptionKeyArn` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkInfo` | `RepositoryLinkInfo` | yes |

## CreateSyncConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Branch` | `string` | yes |
| `ConfigFile` | `string` | yes |
| `RepositoryLinkId` | `string` | yes |
| `ResourceName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `SyncType` | `string` | yes |
| `PublishDeploymentStatus` | `string` | no |
| `TriggerResourceUpdateOn` | `string` | no |
| `PullRequestComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncConfiguration` | `SyncConfiguration` | yes |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRepositoryLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSyncConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncType` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## GetHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Status` | `string` | no |
| `ProviderType` | `string` | no |
| `ProviderEndpoint` | `string` | no |
| `VpcConfiguration` | `VpcConfiguration` | no |

## GetRepositoryLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkInfo` | `RepositoryLinkInfo` | yes |

## GetRepositorySyncStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Branch` | `string` | yes |
| `RepositoryLinkId` | `string` | yes |
| `SyncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LatestSync` | `RepositorySyncAttempt` | yes |

## GetResourceSyncStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `SyncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DesiredState` | `Revision` | no |
| `LatestSuccessfulSync` | `ResourceSyncAttempt` | no |
| `LatestSync` | `ResourceSyncAttempt` | yes |

## GetSyncBlockerSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncType` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncBlockerSummary` | `SyncBlockerSummary` | yes |

## GetSyncConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncType` | `string` | yes |
| `ResourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncConfiguration` | `SyncConfiguration` | yes |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProviderTypeFilter` | `string` | no |
| `HostArnFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<Connection>` | no |
| `NextToken` | `string` | no |

## ListHosts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Hosts` | `List<Host>` | no |
| `NextToken` | `string` | no |

## ListRepositoryLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinks` | `List<RepositoryLinkInfo>` | yes |
| `NextToken` | `string` | no |

## ListRepositorySyncDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkId` | `string` | yes |
| `SyncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositorySyncDefinitions` | `List<RepositorySyncDefinition>` | yes |
| `NextToken` | `string` | no |

## ListSyncConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RepositoryLinkId` | `string` | yes |
| `SyncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncConfigurations` | `List<SyncConfiguration>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

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


## UpdateHost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HostArn` | `string` | yes |
| `ProviderEndpoint` | `string` | no |
| `VpcConfiguration` | `VpcConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRepositoryLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `EncryptionKeyArn` | `string` | no |
| `RepositoryLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RepositoryLinkInfo` | `RepositoryLinkInfo` | yes |

## UpdateSyncBlocker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `SyncType` | `string` | yes |
| `ResourceName` | `string` | yes |
| `ResolvedReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceName` | `string` | yes |
| `ParentResourceName` | `string` | no |
| `SyncBlocker` | `SyncBlocker` | yes |

## UpdateSyncConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Branch` | `string` | no |
| `ConfigFile` | `string` | no |
| `RepositoryLinkId` | `string` | no |
| `ResourceName` | `string` | yes |
| `RoleArn` | `string` | no |
| `SyncType` | `string` | yes |
| `PublishDeploymentStatus` | `string` | no |
| `TriggerResourceUpdateOn` | `string` | no |
| `PullRequestComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SyncConfiguration` | `SyncConfiguration` | yes |

