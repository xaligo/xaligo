# Managed Streaming for Kafka Connect

API version: 2021-09-14. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/kafkaconnect/2021-09-14/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacity` | `Capacity` | yes |
| `connectorConfiguration` | `Map<string>` | yes |
| `connectorDescription` | `string` | no |
| `connectorName` | `string` | yes |
| `kafkaCluster` | `KafkaCluster` | yes |
| `kafkaClusterClientAuthentication` | `KafkaClusterClientAuthentication` | yes |
| `kafkaClusterEncryptionInTransit` | `KafkaClusterEncryptionInTransit` | yes |
| `kafkaConnectVersion` | `string` | yes |
| `logDelivery` | `LogDelivery` | no |
| `networkType` | `string` | no |
| `plugins` | `List<Plugin>` | yes |
| `serviceExecutionRoleArn` | `string` | yes |
| `workerConfiguration` | `WorkerConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |
| `connectorName` | `string` | no |
| `connectorState` | `string` | no |

## CreateCustomPlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | yes |
| `description` | `string` | no |
| `location` | `CustomPluginLocation` | yes |
| `name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customPluginArn` | `string` | no |
| `customPluginState` | `string` | no |
| `name` | `string` | no |
| `revision` | `long` | no |

## CreateWorkerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `name` | `string` | yes |
| `propertiesFileContent` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTime` | `timestamp` | no |
| `latestRevision` | `WorkerConfigurationRevisionSummary` | no |
| `name` | `string` | no |
| `workerConfigurationArn` | `string` | no |
| `workerConfigurationState` | `string` | no |

## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |
| `currentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |
| `connectorState` | `string` | no |

## DeleteCustomPlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customPluginArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customPluginArn` | `string` | no |
| `customPluginState` | `string` | no |

## DeleteWorkerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workerConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workerConfigurationArn` | `string` | no |
| `workerConfigurationState` | `string` | no |

## DescribeConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacity` | `CapacityDescription` | no |
| `connectorArn` | `string` | no |
| `connectorConfiguration` | `Map<string>` | no |
| `connectorDescription` | `string` | no |
| `connectorName` | `string` | no |
| `connectorState` | `string` | no |
| `creationTime` | `timestamp` | no |
| `currentVersion` | `string` | no |
| `kafkaCluster` | `KafkaClusterDescription` | no |
| `kafkaClusterClientAuthentication` | `KafkaClusterClientAuthenticationDescription` | no |
| `kafkaClusterEncryptionInTransit` | `KafkaClusterEncryptionInTransitDescription` | no |
| `kafkaConnectVersion` | `string` | no |
| `logDelivery` | `LogDeliveryDescription` | no |
| `networkType` | `string` | no |
| `plugins` | `List<PluginDescription>` | no |
| `serviceExecutionRoleArn` | `string` | no |
| `workerConfiguration` | `WorkerConfigurationDescription` | no |
| `stateDescription` | `StateDescription` | no |

## DescribeConnectorOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorOperationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |
| `connectorOperationArn` | `string` | no |
| `connectorOperationState` | `string` | no |
| `connectorOperationType` | `string` | no |
| `operationSteps` | `List<ConnectorOperationStep>` | no |
| `originWorkerSetting` | `WorkerSetting` | no |
| `originConnectorConfiguration` | `Map<string>` | no |
| `targetWorkerSetting` | `WorkerSetting` | no |
| `targetConnectorConfiguration` | `Map<string>` | no |
| `errorInfo` | `StateDescription` | no |
| `creationTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |

## DescribeCustomPlugin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customPluginArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTime` | `timestamp` | no |
| `customPluginArn` | `string` | no |
| `customPluginState` | `string` | no |
| `description` | `string` | no |
| `latestRevision` | `CustomPluginRevisionSummary` | no |
| `name` | `string` | no |
| `stateDescription` | `StateDescription` | no |

## DescribeWorkerConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workerConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTime` | `timestamp` | no |
| `description` | `string` | no |
| `latestRevision` | `WorkerConfigurationRevisionDescription` | no |
| `name` | `string` | no |
| `workerConfigurationArn` | `string` | no |
| `workerConfigurationState` | `string` | no |

## ListConnectorOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorOperations` | `List<ConnectorOperationSummary>` | no |
| `nextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorNamePrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectors` | `List<ConnectorSummary>` | no |
| `nextToken` | `string` | no |

## ListCustomPlugins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `namePrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customPlugins` | `List<CustomPluginSummary>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListWorkerConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `namePrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `workerConfigurations` | `List<WorkerConfigurationSummary>` | no |

## RestartConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | yes |
| `onlyFailedTasks` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |
| `connectorOperationArn` | `string` | no |

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


## UpdateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacity` | `CapacityUpdate` | no |
| `connectorConfiguration` | `Map<string>` | no |
| `connectorArn` | `string` | yes |
| `currentVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |
| `connectorState` | `string` | no |
| `connectorOperationArn` | `string` | no |

