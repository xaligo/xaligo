# Amazon AppIntegrations Service

API version: 2020-07-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appintegrations/2020-07-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Namespace` | `string` | yes |
| `Description` | `string` | no |
| `ApplicationSourceConfig` | `ApplicationSourceConfig` | yes |
| `Subscriptions` | `List<Subscription>` | no |
| `Publications` | `List<Publication>` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Permissions` | `List<string>` | no |
| `IsService` | `boolean` | no |
| `InitializationTimeout` | `integer` | no |
| `ApplicationConfig` | `ApplicationConfig` | no |
| `IframeConfig` | `IframeConfig` | no |
| `ApplicationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |

## CreateDataIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `KmsKey` | `string` | yes |
| `SourceURI` | `string` | no |
| `ScheduleConfig` | `ScheduleConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `FileConfiguration` | `FileConfiguration` | no |
| `ObjectConfiguration` | `Map<Map<List<string>>>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `KmsKey` | `string` | no |
| `SourceURI` | `string` | no |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `FileConfiguration` | `FileConfiguration` | no |
| `ObjectConfiguration` | `Map<Map<List<string>>>` | no |

## CreateDataIntegrationAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationIdentifier` | `string` | yes |
| `ClientId` | `string` | no |
| `ObjectConfiguration` | `Map<Map<List<string>>>` | no |
| `DestinationURI` | `string` | no |
| `ClientAssociationMetadata` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `ExecutionConfiguration` | `ExecutionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationAssociationId` | `string` | no |
| `DataIntegrationArn` | `string` | no |

## CreateEventIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `EventFilter` | `EventFilter` | yes |
| `EventBridgeBus` | `string` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventIntegrationArn` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Namespace` | `string` | no |
| `Description` | `string` | no |
| `ApplicationSourceConfig` | `ApplicationSourceConfig` | no |
| `Subscriptions` | `List<Subscription>` | no |
| `Publications` | `List<Publication>` | no |
| `CreatedTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `Tags` | `Map<string>` | no |
| `Permissions` | `List<string>` | no |
| `IsService` | `boolean` | no |
| `InitializationTimeout` | `integer` | no |
| `ApplicationConfig` | `ApplicationConfig` | no |
| `IframeConfig` | `IframeConfig` | no |
| `ApplicationType` | `string` | no |

## GetDataIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `KmsKey` | `string` | no |
| `SourceURI` | `string` | no |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |
| `Tags` | `Map<string>` | no |
| `FileConfiguration` | `FileConfiguration` | no |
| `ObjectConfiguration` | `Map<Map<List<string>>>` | no |

## GetEventIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `EventIntegrationArn` | `string` | no |
| `EventBridgeBus` | `string` | no |
| `EventFilter` | `EventFilter` | no |
| `Tags` | `Map<string>` | no |

## ListApplicationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationAssociations` | `List<ApplicationAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ApplicationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<ApplicationSummary>` | no |
| `NextToken` | `string` | no |

## ListDataIntegrationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationIdentifier` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationAssociations` | `List<DataIntegrationAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListDataIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrations` | `List<DataIntegrationSummary>` | no |
| `NextToken` | `string` | no |

## ListEventIntegrationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventIntegrationName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventIntegrationAssociations` | `List<EventIntegrationAssociation>` | no |
| `NextToken` | `string` | no |

## ListEventIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventIntegrations` | `List<EventIntegration>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

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
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ApplicationSourceConfig` | `ApplicationSourceConfig` | no |
| `Subscriptions` | `List<Subscription>` | no |
| `Publications` | `List<Publication>` | no |
| `Permissions` | `List<string>` | no |
| `IsService` | `boolean` | no |
| `InitializationTimeout` | `integer` | no |
| `ApplicationConfig` | `ApplicationConfig` | no |
| `IframeConfig` | `IframeConfig` | no |
| `ApplicationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataIntegrationAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataIntegrationIdentifier` | `string` | yes |
| `DataIntegrationAssociationIdentifier` | `string` | yes |
| `ExecutionConfiguration` | `ExecutionConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


