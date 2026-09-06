# AWS Lambda

API version: 2015-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/lambda/2015-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddLayerVersionPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `StatementId` | `string` | yes |
| `Action` | `string` | yes |
| `Principal` | `string` | yes |
| `OrganizationId` | `string` | no |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statement` | `string` | no |
| `RevisionId` | `string` | no |

## AddPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `StatementId` | `string` | yes |
| `Action` | `string` | yes |
| `Principal` | `string` | yes |
| `SourceArn` | `string` | no |
| `FunctionUrlAuthType` | `string` | no |
| `InvokedViaFunctionUrl` | `boolean` | no |
| `SourceAccount` | `string` | no |
| `EventSourceToken` | `string` | no |
| `Qualifier` | `string` | no |
| `RevisionId` | `string` | no |
| `PrincipalOrgID` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statement` | `string` | no |

## CheckpointDurableExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `CheckpointToken` | `string` | yes |
| `Updates` | `List<OperationUpdate>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CheckpointToken` | `string` | no |
| `NewExecutionState` | `CheckpointUpdatedExecutionState` | yes |

## CreateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Name` | `string` | yes |
| `FunctionVersion` | `string` | yes |
| `Description` | `string` | no |
| `RoutingConfig` | `AliasRoutingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasArn` | `string` | no |
| `Name` | `string` | no |
| `FunctionVersion` | `string` | no |
| `Description` | `string` | no |
| `RoutingConfig` | `AliasRoutingConfiguration` | no |
| `RevisionId` | `string` | no |

## CreateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderName` | `string` | yes |
| `VpcConfig` | `CapacityProviderVpcConfig` | yes |
| `PermissionsConfig` | `CapacityProviderPermissionsConfig` | yes |
| `InstanceRequirements` | `InstanceRequirements` | no |
| `CapacityProviderScalingConfig` | `CapacityProviderScalingConfig` | no |
| `KmsKeyArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `PropagateTags` | `PropagateTags` | no |
| `TelemetryConfig` | `CapacityProviderTelemetryConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProvider` | `CapacityProvider` | yes |

## CreateCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `AllowedPublishers` | `AllowedPublishers` | yes |
| `CodeSigningPolicies` | `CodeSigningPolicies` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfig` | `CodeSigningConfig` | yes |

## CreateEventSourceMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSourceArn` | `string` | no |
| `FunctionName` | `string` | yes |
| `Enabled` | `boolean` | no |
| `BatchSize` | `integer` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `StartingPosition` | `string` | no |
| `StartingPositionTimestamp` | `timestamp` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `Tags` | `Map<string>` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `Topics` | `List<string>` | no |
| `Queues` | `List<string>` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `SelfManagedEventSource` | `SelfManagedEventSource` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | no |
| `StartingPosition` | `string` | no |
| `StartingPositionTimestamp` | `timestamp` | no |
| `BatchSize` | `integer` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `EventSourceArn` | `string` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `FilterCriteriaError` | `FilterCriteriaError` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `FunctionArn` | `string` | no |
| `LastModified` | `timestamp` | no |
| `LastProcessingResult` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `Topics` | `List<string>` | no |
| `Queues` | `List<string>` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `SelfManagedEventSource` | `SelfManagedEventSource` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `EventSourceMappingArn` | `string` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

## CreateFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Runtime` | `string` | no |
| `Role` | `string` | yes |
| `Handler` | `string` | no |
| `Code` | `FunctionCode` | yes |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `Publish` | `boolean` | no |
| `PublishTo` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `PackageType` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `Environment` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfig` | no |
| `Tags` | `Map<string>` | no |
| `Layers` | `List<string>` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `CodeSigningConfigArn` | `string` | no |
| `ImageConfig` | `ImageConfig` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStart` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `DurableConfig` | `DurableConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | no |
| `FunctionArn` | `string` | no |
| `Runtime` | `string` | no |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `CodeSize` | `long` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `LastModified` | `string` | no |
| `CodeSha256` | `string` | no |
| `Version` | `string` | no |
| `VpcConfig` | `VpcConfigResponse` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `EnvironmentResponse` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfigResponse` | no |
| `MasterArn` | `string` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<Layer>` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `SigningProfileVersionArn` | `string` | no |
| `SigningJobArn` | `string` | no |
| `PackageType` | `string` | no |
| `ImageConfigResponse` | `ImageConfigResponse` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStartResponse` | no |
| `RuntimeVersionConfig` | `RuntimeVersionConfig` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `ConfigSha256` | `string` | no |
| `DurableConfig` | `DurableConfig` | no |

## CreateFunctionUrlConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `AuthType` | `string` | yes |
| `Cors` | `Cors` | no |
| `InvokeMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionUrl` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `AuthType` | `string` | yes |
| `Cors` | `Cors` | no |
| `CreationTime` | `string` | yes |
| `InvokeMode` | `string` | no |

## DeleteAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProvider` | `CapacityProvider` | yes |

## DeleteCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventSourceMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | no |
| `StartingPosition` | `string` | no |
| `StartingPositionTimestamp` | `timestamp` | no |
| `BatchSize` | `integer` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `EventSourceArn` | `string` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `FilterCriteriaError` | `FilterCriteriaError` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `FunctionArn` | `string` | no |
| `LastModified` | `timestamp` | no |
| `LastProcessingResult` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `Topics` | `List<string>` | no |
| `Queues` | `List<string>` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `SelfManagedEventSource` | `SelfManagedEventSource` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `EventSourceMappingArn` | `string` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

## DeleteFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusCode` | `integer` | no |

## DeleteFunctionCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunctionConcurrency

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunctionEventInvokeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunctionUrlConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLayerVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `VersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisionedConcurrencyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLimit` | `AccountLimit` | no |
| `AccountUsage` | `AccountUsage` | no |

## GetAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasArn` | `string` | no |
| `Name` | `string` | no |
| `FunctionVersion` | `string` | no |
| `Description` | `string` | no |
| `RoutingConfig` | `AliasRoutingConfiguration` | no |
| `RevisionId` | `string` | no |

## GetCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProvider` | `CapacityProvider` | yes |

## GetCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfig` | `CodeSigningConfig` | yes |

## GetDurableExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `IncludeExecutionData` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `DurableExecutionName` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `InputPayload` | `string` | no |
| `Result` | `string` | no |
| `Error` | `ErrorObject` | no |
| `StartTimestamp` | `timestamp` | yes |
| `Status` | `string` | yes |
| `EndTimestamp` | `timestamp` | no |
| `Version` | `string` | no |
| `TraceHeader` | `TraceHeader` | no |
| `ExecutionDataIncluded` | `boolean` | no |
| `DurableConfig` | `DurableConfig` | no |

## GetDurableExecutionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `IncludeExecutionData` | `boolean` | no |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |
| `ReverseOrder` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<Event>` | yes |
| `NextMarker` | `string` | no |

## GetDurableExecutionState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `CheckpointToken` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operations` | `List<Operation>` | yes |
| `NextMarker` | `string` | no |

## GetEventSourceMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | no |
| `StartingPosition` | `string` | no |
| `StartingPositionTimestamp` | `timestamp` | no |
| `BatchSize` | `integer` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `EventSourceArn` | `string` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `FilterCriteriaError` | `FilterCriteriaError` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `FunctionArn` | `string` | no |
| `LastModified` | `timestamp` | no |
| `LastProcessingResult` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `Topics` | `List<string>` | no |
| `Queues` | `List<string>` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `SelfManagedEventSource` | `SelfManagedEventSource` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `EventSourceMappingArn` | `string` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

## GetFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configuration` | `FunctionConfiguration` | no |
| `Code` | `FunctionCodeLocation` | no |
| `Tags` | `Map<string>` | no |
| `TagsError` | `TagsError` | no |
| `Concurrency` | `Concurrency` | no |

## GetFunctionCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |
| `FunctionName` | `string` | yes |

## GetFunctionConcurrency

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedConcurrentExecutions` | `integer` | no |

## GetFunctionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | no |
| `FunctionArn` | `string` | no |
| `Runtime` | `string` | no |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `CodeSize` | `long` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `LastModified` | `string` | no |
| `CodeSha256` | `string` | no |
| `Version` | `string` | no |
| `VpcConfig` | `VpcConfigResponse` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `EnvironmentResponse` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfigResponse` | no |
| `MasterArn` | `string` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<Layer>` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `SigningProfileVersionArn` | `string` | no |
| `SigningJobArn` | `string` | no |
| `PackageType` | `string` | no |
| `ImageConfigResponse` | `ImageConfigResponse` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStartResponse` | no |
| `RuntimeVersionConfig` | `RuntimeVersionConfig` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `ConfigSha256` | `string` | no |
| `DurableConfig` | `DurableConfig` | no |

## GetFunctionEventInvokeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModified` | `timestamp` | no |
| `FunctionArn` | `string` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `MaximumEventAgeInSeconds` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |

## GetFunctionRecursionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecursiveLoop` | `string` | no |

## GetFunctionScalingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionArn` | `string` | no |
| `AppliedFunctionScalingConfig` | `FunctionScalingConfig` | no |
| `RequestedFunctionScalingConfig` | `FunctionScalingConfig` | no |

## GetFunctionUrlConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionUrl` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `AuthType` | `string` | yes |
| `Cors` | `Cors` | no |
| `CreationTime` | `string` | yes |
| `LastModifiedTime` | `string` | yes |
| `InvokeMode` | `string` | no |

## GetLayerVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `VersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `LayerVersionContentOutput` | no |
| `LayerArn` | `string` | no |
| `LayerVersionArn` | `string` | no |
| `Description` | `string` | no |
| `CreatedDate` | `string` | no |
| `Version` | `long` | no |
| `CompatibleArchitectures` | `List<string>` | no |
| `CompatibleRuntimes` | `List<string>` | no |
| `LicenseInfo` | `string` | no |

## GetLayerVersionByArn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `LayerVersionContentOutput` | no |
| `LayerArn` | `string` | no |
| `LayerVersionArn` | `string` | no |
| `Description` | `string` | no |
| `CreatedDate` | `string` | no |
| `Version` | `long` | no |
| `CompatibleArchitectures` | `List<string>` | no |
| `CompatibleRuntimes` | `List<string>` | no |
| `LicenseInfo` | `string` | no |

## GetLayerVersionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `VersionNumber` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## GetProvisionedConcurrencyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestedProvisionedConcurrentExecutions` | `integer` | no |
| `AvailableProvisionedConcurrentExecutions` | `integer` | no |
| `AllocatedProvisionedConcurrentExecutions` | `integer` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `LastModified` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## GetRuntimeManagementConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateRuntimeOn` | `string` | no |
| `FunctionArn` | `string` | no |
| `RuntimeVersionArn` | `string` | no |

## Invoke

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `InvocationType` | `string` | no |
| `LogType` | `string` | no |
| `ClientContext` | `string` | no |
| `DurableExecutionName` | `string` | no |
| `Payload` | `blob` | no |
| `Qualifier` | `string` | no |
| `TenantId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusCode` | `integer` | no |
| `FunctionError` | `string` | no |
| `LogResult` | `string` | no |
| `Payload` | `blob` | no |
| `ExecutedVersion` | `string` | no |
| `DurableExecutionArn` | `string` | no |

## InvokeAsync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `InvokeArgs` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `integer` | no |

## InvokeWithResponseStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `LogType` | `string` | no |
| `ClientContext` | `string` | no |
| `Qualifier` | `string` | no |
| `Payload` | `blob` | no |
| `TenantId` | `string` | no |
| `InvocationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StatusCode` | `integer` | no |
| `ExecutedVersion` | `string` | no |
| `EventStream` | `InvokeWithResponseStreamResponseEvent` | no |
| `ResponseStreamContentType` | `string` | no |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `FunctionVersion` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Aliases` | `List<AliasConfiguration>` | no |

## ListCapacityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviders` | `List<CapacityProvider>` | yes |
| `NextMarker` | `string` | no |

## ListCodeSigningConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `CodeSigningConfigs` | `List<CodeSigningConfig>` | no |

## ListDurableExecutionsByFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `DurableExecutionName` | `string` | no |
| `Statuses` | `List<string>` | no |
| `StartedAfter` | `timestamp` | no |
| `StartedBefore` | `timestamp` | no |
| `ReverseOrder` | `boolean` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutions` | `List<Execution>` | no |
| `NextMarker` | `string` | no |

## ListEventSourceMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventSourceArn` | `string` | no |
| `FunctionName` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `EventSourceMappings` | `List<EventSourceMappingConfiguration>` | no |

## ListFunctionEventInvokeConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionEventInvokeConfigs` | `List<FunctionEventInvokeConfig>` | no |
| `NextMarker` | `string` | no |

## ListFunctionUrlConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionUrlConfigs` | `List<FunctionUrlConfig>` | yes |
| `NextMarker` | `string` | no |

## ListFunctionVersionsByCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderArn` | `string` | yes |
| `FunctionVersions` | `List<FunctionVersionsByCapacityProviderListItem>` | yes |
| `NextMarker` | `string` | no |

## ListFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MasterRegion` | `string` | no |
| `FunctionVersion` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Functions` | `List<FunctionConfiguration>` | no |

## ListFunctionsByCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `FunctionArns` | `List<string>` | no |

## ListLayerVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleArchitecture` | `string` | no |
| `CompatibleRuntime` | `string` | no |
| `LayerName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `LayerVersions` | `List<LayerVersionsListItem>` | no |

## ListLayers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleArchitecture` | `string` | no |
| `CompatibleRuntime` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Layers` | `List<LayersListItem>` | no |

## ListProvisionedConcurrencyConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProvisionedConcurrencyConfigs` | `List<ProvisionedConcurrencyConfigListItem>` | no |
| `NextMarker` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListVersionsByFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `Versions` | `List<FunctionConfiguration>` | no |

## PublishLayerVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `Description` | `string` | no |
| `Content` | `LayerVersionContentInput` | yes |
| `CompatibleArchitectures` | `List<string>` | no |
| `CompatibleRuntimes` | `List<string>` | no |
| `LicenseInfo` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `LayerVersionContentOutput` | no |
| `LayerArn` | `string` | no |
| `LayerVersionArn` | `string` | no |
| `Description` | `string` | no |
| `CreatedDate` | `string` | no |
| `Version` | `long` | no |
| `CompatibleArchitectures` | `List<string>` | no |
| `CompatibleRuntimes` | `List<string>` | no |
| `LicenseInfo` | `string` | no |

## PublishVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `CodeSha256` | `string` | no |
| `Description` | `string` | no |
| `RevisionId` | `string` | no |
| `PublishTo` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | no |
| `FunctionArn` | `string` | no |
| `Runtime` | `string` | no |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `CodeSize` | `long` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `LastModified` | `string` | no |
| `CodeSha256` | `string` | no |
| `Version` | `string` | no |
| `VpcConfig` | `VpcConfigResponse` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `EnvironmentResponse` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfigResponse` | no |
| `MasterArn` | `string` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<Layer>` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `SigningProfileVersionArn` | `string` | no |
| `SigningJobArn` | `string` | no |
| `PackageType` | `string` | no |
| `ImageConfigResponse` | `ImageConfigResponse` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStartResponse` | no |
| `RuntimeVersionConfig` | `RuntimeVersionConfig` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `ConfigSha256` | `string` | no |
| `DurableConfig` | `DurableConfig` | no |

## PutFunctionCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |
| `FunctionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |
| `FunctionName` | `string` | yes |

## PutFunctionConcurrency

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `ReservedConcurrentExecutions` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedConcurrentExecutions` | `integer` | no |

## PutFunctionEventInvokeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `MaximumEventAgeInSeconds` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModified` | `timestamp` | no |
| `FunctionArn` | `string` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `MaximumEventAgeInSeconds` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |

## PutFunctionRecursionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `RecursiveLoop` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecursiveLoop` | `string` | no |

## PutFunctionScalingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | yes |
| `FunctionScalingConfig` | `FunctionScalingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionState` | `string` | no |

## PutProvisionedConcurrencyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | yes |
| `ProvisionedConcurrentExecutions` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RequestedProvisionedConcurrentExecutions` | `integer` | no |
| `AllocatedProvisionedConcurrentExecutions` | `integer` | no |
| `AvailableProvisionedConcurrentExecutions` | `integer` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `LastModified` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |
| `RevisionId` | `string` | no |

## PutRuntimeManagementConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `UpdateRuntimeOn` | `string` | yes |
| `RuntimeVersionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateRuntimeOn` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `RuntimeVersionArn` | `string` | no |

## RemoveLayerVersionPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayerName` | `string` | yes |
| `VersionNumber` | `long` | yes |
| `StatementId` | `string` | yes |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemovePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `StatementId` | `string` | yes |
| `Qualifier` | `string` | no |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDurableExecutionCallbackFailure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackId` | `string` | yes |
| `Error` | `ErrorObject` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDurableExecutionCallbackHeartbeat

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendDurableExecutionCallbackSuccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackId` | `string` | yes |
| `Result` | `blob` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopDurableExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DurableExecutionArn` | `string` | yes |
| `Error` | `ErrorObject` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StopTimestamp` | `timestamp` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Name` | `string` | yes |
| `FunctionVersion` | `string` | no |
| `Description` | `string` | no |
| `RoutingConfig` | `AliasRoutingConfiguration` | no |
| `RevisionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasArn` | `string` | no |
| `Name` | `string` | no |
| `FunctionVersion` | `string` | no |
| `Description` | `string` | no |
| `RoutingConfig` | `AliasRoutingConfiguration` | no |
| `RevisionId` | `string` | no |

## UpdateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProviderName` | `string` | yes |
| `CapacityProviderScalingConfig` | `CapacityProviderScalingConfig` | no |
| `PropagateTags` | `PropagateTags` | no |
| `TelemetryConfig` | `CapacityProviderTelemetryConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CapacityProvider` | `CapacityProvider` | yes |

## UpdateCodeSigningConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfigArn` | `string` | yes |
| `Description` | `string` | no |
| `AllowedPublishers` | `AllowedPublishers` | no |
| `CodeSigningPolicies` | `CodeSigningPolicies` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeSigningConfig` | `CodeSigningConfig` | yes |

## UpdateEventSourceMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | yes |
| `FunctionName` | `string` | no |
| `Enabled` | `boolean` | no |
| `BatchSize` | `integer` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UUID` | `string` | no |
| `StartingPosition` | `string` | no |
| `StartingPositionTimestamp` | `timestamp` | no |
| `BatchSize` | `integer` | no |
| `MaximumBatchingWindowInSeconds` | `integer` | no |
| `ParallelizationFactor` | `integer` | no |
| `EventSourceArn` | `string` | no |
| `FilterCriteria` | `FilterCriteria` | no |
| `FilterCriteriaError` | `FilterCriteriaError` | no |
| `KMSKeyArn` | `string` | no |
| `MetricsConfig` | `EventSourceMappingMetricsConfig` | no |
| `LoggingConfig` | `EventSourceMappingLoggingConfig` | no |
| `ScalingConfig` | `ScalingConfig` | no |
| `FunctionArn` | `string` | no |
| `LastModified` | `timestamp` | no |
| `LastProcessingResult` | `string` | no |
| `State` | `string` | no |
| `StateTransitionReason` | `string` | no |
| `DestinationConfig` | `DestinationConfig` | no |
| `Topics` | `List<string>` | no |
| `Queues` | `List<string>` | no |
| `SourceAccessConfigurations` | `List<SourceAccessConfiguration>` | no |
| `SelfManagedEventSource` | `SelfManagedEventSource` | no |
| `MaximumRecordAgeInSeconds` | `integer` | no |
| `BisectBatchOnFunctionError` | `boolean` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `TumblingWindowInSeconds` | `integer` | no |
| `FunctionResponseTypes` | `List<string>` | no |
| `AmazonManagedKafkaEventSourceConfig` | `AmazonManagedKafkaEventSourceConfig` | no |
| `SelfManagedKafkaEventSourceConfig` | `SelfManagedKafkaEventSourceConfig` | no |
| `DocumentDBEventSourceConfig` | `DocumentDBEventSourceConfig` | no |
| `EventSourceMappingArn` | `string` | no |
| `ProvisionedPollerConfig` | `ProvisionedPollerConfig` | no |

## UpdateFunctionCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `ZipFile` | `blob` | no |
| `S3Bucket` | `string` | no |
| `S3Key` | `string` | no |
| `S3ObjectVersion` | `string` | no |
| `S3ObjectStorageMode` | `string` | no |
| `ImageUri` | `string` | no |
| `Architectures` | `List<string>` | no |
| `Publish` | `boolean` | no |
| `PublishTo` | `string` | no |
| `DryRun` | `boolean` | no |
| `RevisionId` | `string` | no |
| `SourceKMSKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | no |
| `FunctionArn` | `string` | no |
| `Runtime` | `string` | no |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `CodeSize` | `long` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `LastModified` | `string` | no |
| `CodeSha256` | `string` | no |
| `Version` | `string` | no |
| `VpcConfig` | `VpcConfigResponse` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `EnvironmentResponse` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfigResponse` | no |
| `MasterArn` | `string` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<Layer>` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `SigningProfileVersionArn` | `string` | no |
| `SigningJobArn` | `string` | no |
| `PackageType` | `string` | no |
| `ImageConfigResponse` | `ImageConfigResponse` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStartResponse` | no |
| `RuntimeVersionConfig` | `RuntimeVersionConfig` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `ConfigSha256` | `string` | no |
| `DurableConfig` | `DurableConfig` | no |

## UpdateFunctionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Environment` | `Environment` | no |
| `Runtime` | `string` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfig` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<string>` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `ImageConfig` | `ImageConfig` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStart` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `DurableConfig` | `DurableConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | no |
| `FunctionArn` | `string` | no |
| `Runtime` | `string` | no |
| `Role` | `string` | no |
| `Handler` | `string` | no |
| `CodeSize` | `long` | no |
| `Description` | `string` | no |
| `Timeout` | `integer` | no |
| `MemorySize` | `integer` | no |
| `LastModified` | `string` | no |
| `CodeSha256` | `string` | no |
| `Version` | `string` | no |
| `VpcConfig` | `VpcConfigResponse` | no |
| `DeadLetterConfig` | `DeadLetterConfig` | no |
| `Environment` | `EnvironmentResponse` | no |
| `KMSKeyArn` | `string` | no |
| `TracingConfig` | `TracingConfigResponse` | no |
| `MasterArn` | `string` | no |
| `RevisionId` | `string` | no |
| `Layers` | `List<Layer>` | no |
| `State` | `string` | no |
| `StateReason` | `string` | no |
| `StateReasonCode` | `string` | no |
| `LastUpdateStatus` | `string` | no |
| `LastUpdateStatusReason` | `string` | no |
| `LastUpdateStatusReasonCode` | `string` | no |
| `FileSystemConfigs` | `List<FileSystemConfig>` | no |
| `SigningProfileVersionArn` | `string` | no |
| `SigningJobArn` | `string` | no |
| `PackageType` | `string` | no |
| `ImageConfigResponse` | `ImageConfigResponse` | no |
| `Architectures` | `List<string>` | no |
| `EphemeralStorage` | `EphemeralStorage` | no |
| `SnapStart` | `SnapStartResponse` | no |
| `RuntimeVersionConfig` | `RuntimeVersionConfig` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `TenancyConfig` | `TenancyConfig` | no |
| `CapacityProviderConfig` | `CapacityProviderConfig` | no |
| `ConfigSha256` | `string` | no |
| `DurableConfig` | `DurableConfig` | no |

## UpdateFunctionEventInvokeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `MaximumEventAgeInSeconds` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModified` | `timestamp` | no |
| `FunctionArn` | `string` | no |
| `MaximumRetryAttempts` | `integer` | no |
| `MaximumEventAgeInSeconds` | `integer` | no |
| `DestinationConfig` | `DestinationConfig` | no |

## UpdateFunctionUrlConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionName` | `string` | yes |
| `Qualifier` | `string` | no |
| `AuthType` | `string` | no |
| `Cors` | `Cors` | no |
| `InvokeMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionUrl` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `AuthType` | `string` | yes |
| `Cors` | `Cors` | no |
| `CreationTime` | `string` | yes |
| `LastModifiedTime` | `string` | yes |
| `InvokeMode` | `string` | no |

