# AWS Migration Hub

API version: 2017-05-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mgh/2017-05-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateCreatedArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `CreatedArtifact` | `CreatedArtifact` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateDiscoveredResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `DiscoveredResource` | `DiscoveredResource` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSourceResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `SourceResource` | `SourceResource` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateProgressUpdateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStreamName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProgressUpdateStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStreamName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeApplicationState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStatus` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |

## DescribeMigrationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MigrationTask` | `MigrationTask` | no |

## DisassociateCreatedArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `CreatedArtifactName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateDiscoveredResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `ConfigurationId` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSourceResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `SourceResourceName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ImportMigrationTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListApplicationStates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationStateList` | `List<ApplicationState>` | no |
| `NextToken` | `string` | no |

## ListCreatedArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `CreatedArtifactList` | `List<CreatedArtifact>` | no |

## ListDiscoveredResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DiscoveredResourceList` | `List<DiscoveredResource>` | no |

## ListMigrationTaskUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MigrationTaskUpdateList` | `List<MigrationTaskUpdate>` | no |

## ListMigrationTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ResourceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MigrationTaskSummaryList` | `List<MigrationTaskSummary>` | no |

## ListProgressUpdateStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStreamSummaryList` | `List<ProgressUpdateStreamSummary>` | no |
| `NextToken` | `string` | no |

## ListSourceResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SourceResourceList` | `List<SourceResource>` | no |

## NotifyApplicationState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Status` | `string` | yes |
| `UpdateDateTime` | `timestamp` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## NotifyMigrationTaskState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `Task` | `Task` | yes |
| `UpdateDateTime` | `timestamp` | yes |
| `NextUpdateSeconds` | `integer` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutResourceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressUpdateStream` | `string` | yes |
| `MigrationTaskName` | `string` | yes |
| `ResourceAttributeList` | `List<ResourceAttribute>` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


