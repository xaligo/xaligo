# Amazon EventBridge

API version: 2015-10-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/events/2015-10-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelReplay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayArn` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |

## CreateApiDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ConnectionArn` | `string` | yes |
| `InvocationEndpoint` | `string` | yes |
| `HttpMethod` | `string` | yes |
| `InvocationRateLimitPerSecond` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiDestinationArn` | `string` | no |
| `ApiDestinationState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## CreateArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveName` | `string` | yes |
| `EventSourceArn` | `string` | yes |
| `Description` | `string` | no |
| `EventPattern` | `string` | no |
| `RetentionDays` | `integer` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveArn` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `CreationTime` | `timestamp` | no |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `AuthorizationType` | `string` | yes |
| `AuthParameters` | `CreateConnectionAuthRequestParameters` | yes |
| `InvocationConnectivityParameters` | `ConnectivityResourceParameters` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `ConnectionState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## CreateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `RoutingConfig` | `RoutingConfig` | yes |
| `ReplicationConfig` | `ReplicationConfig` | no |
| `EventBuses` | `List<EndpointEventBus>` | yes |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Arn` | `string` | no |
| `RoutingConfig` | `RoutingConfig` | no |
| `ReplicationConfig` | `ReplicationConfig` | no |
| `EventBuses` | `List<EndpointEventBus>` | no |
| `RoleArn` | `string` | no |
| `State` | `string` | no |

## CreateEventBus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventSourceName` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `LogConfig` | `LogConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventBusArn` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `LogConfig` | `LogConfig` | no |

## CreatePartnerEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Account` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSourceArn` | `string` | no |

## DeactivateEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeauthorizeConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `ConnectionState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastAuthorizedTime` | `timestamp` | no |

## DeleteApiDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `ConnectionState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastAuthorizedTime` | `timestamp` | no |

## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventBus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePartnerEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Account` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventBusName` | `string` | no |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeApiDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiDestinationArn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ApiDestinationState` | `string` | no |
| `ConnectionArn` | `string` | no |
| `InvocationEndpoint` | `string` | no |
| `HttpMethod` | `string` | no |
| `InvocationRateLimitPerSecond` | `integer` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribeArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveArn` | `string` | no |
| `ArchiveName` | `string` | no |
| `EventSourceArn` | `string` | no |
| `Description` | `string` | no |
| `EventPattern` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `RetentionDays` | `integer` | no |
| `SizeBytes` | `long` | no |
| `EventCount` | `long` | no |
| `CreationTime` | `timestamp` | no |

## DescribeConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `InvocationConnectivityParameters` | `DescribeConnectionConnectivityParameters` | no |
| `ConnectionState` | `string` | no |
| `StateReason` | `string` | no |
| `AuthorizationType` | `string` | no |
| `SecretArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `AuthParameters` | `ConnectionAuthResponseParameters` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastAuthorizedTime` | `timestamp` | no |

## DescribeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `HomeRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Arn` | `string` | no |
| `RoutingConfig` | `RoutingConfig` | no |
| `ReplicationConfig` | `ReplicationConfig` | no |
| `EventBuses` | `List<EndpointEventBus>` | no |
| `RoleArn` | `string` | no |
| `EndpointId` | `string` | no |
| `EndpointUrl` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribeEventBus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Policy` | `string` | no |
| `LogConfig` | `LogConfig` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribeEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedBy` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `ExpirationTime` | `timestamp` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## DescribePartnerEventSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |

## DescribeReplay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayName` | `string` | no |
| `ReplayArn` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `EventSourceArn` | `string` | no |
| `Destination` | `ReplayDestination` | no |
| `EventStartTime` | `timestamp` | no |
| `EventEndTime` | `timestamp` | no |
| `EventLastReplayedTime` | `timestamp` | no |
| `ReplayStartTime` | `timestamp` | no |
| `ReplayEndTime` | `timestamp` | no |

## DescribeRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventBusName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Arn` | `string` | no |
| `EventPattern` | `string` | no |
| `ScheduleExpression` | `string` | no |
| `State` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `ManagedBy` | `string` | no |
| `EventBusName` | `string` | no |
| `CreatedBy` | `string` | no |

## DisableRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventBusName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `EventBusName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListApiDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `ConnectionArn` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiDestinations` | `List<ApiDestination>` | no |
| `NextToken` | `string` | no |

## ListArchives

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `EventSourceArn` | `string` | no |
| `State` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Archives` | `List<Archive>` | no |
| `NextToken` | `string` | no |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `ConnectionState` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<Connection>` | no |
| `NextToken` | `string` | no |

## ListEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `HomeRegion` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | no |
| `NextToken` | `string` | no |

## ListEventBuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventBuses` | `List<EventBus>` | no |
| `NextToken` | `string` | no |

## ListEventSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSources` | `List<EventSource>` | no |
| `NextToken` | `string` | no |

## ListPartnerEventSourceAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSourceName` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerEventSourceAccounts` | `List<PartnerEventSourceAccount>` | no |
| `NextToken` | `string` | no |

## ListPartnerEventSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | yes |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PartnerEventSources` | `List<PartnerEventSource>` | no |
| `NextToken` | `string` | no |

## ListReplays

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `State` | `string` | no |
| `EventSourceArn` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Replays` | `List<Replay>` | no |
| `NextToken` | `string` | no |

## ListRuleNamesByTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | yes |
| `EventBusName` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleNames` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `EventBusName` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<Rule>` | no |
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

## ListTargetsByRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `string` | yes |
| `EventBusName` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Targets` | `List<Target>` | no |
| `NextToken` | `string` | no |

## PutEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<PutEventsRequestEntry>` | yes |
| `EndpointId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntryCount` | `integer` | no |
| `Entries` | `List<PutEventsResultEntry>` | no |

## PutPartnerEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entries` | `List<PutPartnerEventsRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntryCount` | `integer` | no |
| `Entries` | `List<PutPartnerEventsResultEntry>` | no |

## PutPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventBusName` | `string` | no |
| `Action` | `string` | no |
| `Principal` | `string` | no |
| `StatementId` | `string` | no |
| `Condition` | `Condition` | no |
| `Policy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ScheduleExpression` | `string` | no |
| `EventPattern` | `string` | no |
| `State` | `string` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `EventBusName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | no |

## PutTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `string` | yes |
| `EventBusName` | `string` | no |
| `Targets` | `List<Target>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntryCount` | `integer` | no |
| `FailedEntries` | `List<PutTargetsResultEntry>` | no |

## RemovePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatementId` | `string` | no |
| `RemoveAllPermissions` | `boolean` | no |
| `EventBusName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `string` | yes |
| `EventBusName` | `string` | no |
| `Ids` | `List<string>` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedEntryCount` | `integer` | no |
| `FailedEntries` | `List<RemoveTargetsResultEntry>` | no |

## StartReplay

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayName` | `string` | yes |
| `Description` | `string` | no |
| `EventSourceArn` | `string` | yes |
| `EventStartTime` | `timestamp` | yes |
| `EventEndTime` | `timestamp` | yes |
| `Destination` | `ReplayDestination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReplayArn` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `ReplayStartTime` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestEventPattern

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventPattern` | `string` | yes |
| `Event` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `boolean` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApiDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ConnectionArn` | `string` | no |
| `InvocationEndpoint` | `string` | no |
| `HttpMethod` | `string` | no |
| `InvocationRateLimitPerSecond` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiDestinationArn` | `string` | no |
| `ApiDestinationState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## UpdateArchive

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveName` | `string` | yes |
| `Description` | `string` | no |
| `EventPattern` | `string` | no |
| `RetentionDays` | `integer` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArchiveArn` | `string` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `CreationTime` | `timestamp` | no |

## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `AuthorizationType` | `string` | no |
| `AuthParameters` | `UpdateConnectionAuthRequestParameters` | no |
| `InvocationConnectivityParameters` | `ConnectivityResourceParameters` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | no |
| `ConnectionState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastAuthorizedTime` | `timestamp` | no |

## UpdateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `RoutingConfig` | `RoutingConfig` | no |
| `ReplicationConfig` | `ReplicationConfig` | no |
| `EventBuses` | `List<EndpointEventBus>` | no |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Arn` | `string` | no |
| `RoutingConfig` | `RoutingConfig` | no |
| `ReplicationConfig` | `ReplicationConfig` | no |
| `EventBuses` | `List<EndpointEventBus>` | no |
| `RoleArn` | `string` | no |
| `EndpointId` | `string` | no |
| `EndpointUrl` | `string` | no |
| `State` | `string` | no |

## UpdateEventBus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `Description` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `LogConfig` | `LogConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `Description` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `LogConfig` | `LogConfig` | no |

