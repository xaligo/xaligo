# Amazon Appflow

API version: 2020-08-23. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appflow/2020-08-23/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelFlowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `executionIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `invalidExecutions` | `List<string>` | no |

## CreateConnectorProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileName` | `string` | yes |
| `kmsArn` | `string` | no |
| `connectorType` | `string` | yes |
| `connectorLabel` | `string` | no |
| `connectionMode` | `string` | yes |
| `connectorProfileConfig` | `ConnectorProfileConfig` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileArn` | `string` | no |

## CreateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `description` | `string` | no |
| `kmsArn` | `string` | no |
| `triggerConfig` | `TriggerConfig` | yes |
| `sourceFlowConfig` | `SourceFlowConfig` | yes |
| `destinationFlowConfigList` | `List<DestinationFlowConfig>` | yes |
| `tasks` | `List<Task>` | yes |
| `tags` | `Map<string>` | no |
| `metadataCatalogConfig` | `MetadataCatalogConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowArn` | `string` | no |
| `flowStatus` | `string` | no |

## DeleteConnectorProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileName` | `string` | yes |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `forceDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorType` | `string` | yes |
| `connectorLabel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorConfiguration` | `ConnectorConfiguration` | no |

## DescribeConnectorEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorEntityName` | `string` | yes |
| `connectorType` | `string` | no |
| `connectorProfileName` | `string` | no |
| `apiVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorEntityFields` | `List<ConnectorEntityField>` | yes |

## DescribeConnectorProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileNames` | `List<string>` | no |
| `connectorType` | `string` | no |
| `connectorLabel` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileDetails` | `List<ConnectorProfile>` | no |
| `nextToken` | `string` | no |

## DescribeConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorTypes` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorConfigurations` | `Map<ConnectorConfiguration>` | no |
| `connectors` | `List<ConnectorDetail>` | no |
| `nextToken` | `string` | no |

## DescribeFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowArn` | `string` | no |
| `description` | `string` | no |
| `flowName` | `string` | no |
| `kmsArn` | `string` | no |
| `flowStatus` | `string` | no |
| `flowStatusMessage` | `string` | no |
| `sourceFlowConfig` | `SourceFlowConfig` | no |
| `destinationFlowConfigList` | `List<DestinationFlowConfig>` | no |
| `lastRunExecutionDetails` | `ExecutionDetails` | no |
| `triggerConfig` | `TriggerConfig` | no |
| `tasks` | `List<Task>` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `lastUpdatedBy` | `string` | no |
| `tags` | `Map<string>` | no |
| `metadataCatalogConfig` | `MetadataCatalogConfig` | no |
| `lastRunMetadataCatalogDetails` | `List<MetadataCatalogDetail>` | no |
| `schemaVersion` | `long` | no |

## DescribeFlowExecutionRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecutions` | `List<ExecutionRecord>` | no |
| `nextToken` | `string` | no |

## ListConnectorEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileName` | `string` | no |
| `connectorType` | `string` | no |
| `entitiesPath` | `string` | no |
| `apiVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorEntityMap` | `Map<List<ConnectorEntity>>` | yes |
| `nextToken` | `string` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectors` | `List<ConnectorDetail>` | no |
| `nextToken` | `string` | no |

## ListFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flows` | `List<FlowDefinition>` | no |
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

## RegisterConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorLabel` | `string` | no |
| `description` | `string` | no |
| `connectorProvisioningType` | `string` | no |
| `connectorProvisioningConfig` | `ConnectorProvisioningConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |

## ResetConnectorMetadataCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileName` | `string` | no |
| `connectorType` | `string` | no |
| `connectorEntityName` | `string` | no |
| `entitiesPath` | `string` | no |
| `apiVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowArn` | `string` | no |
| `flowStatus` | `string` | no |
| `executionId` | `string` | no |

## StopFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowArn` | `string` | no |
| `flowStatus` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UnregisterConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorLabel` | `string` | yes |
| `forceDelete` | `boolean` | no |

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


## UpdateConnectorProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileName` | `string` | yes |
| `connectionMode` | `string` | yes |
| `connectorProfileConfig` | `ConnectorProfileConfig` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorProfileArn` | `string` | no |

## UpdateConnectorRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorLabel` | `string` | yes |
| `description` | `string` | no |
| `connectorProvisioningConfig` | `ConnectorProvisioningConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectorArn` | `string` | no |

## UpdateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowName` | `string` | yes |
| `description` | `string` | no |
| `triggerConfig` | `TriggerConfig` | yes |
| `sourceFlowConfig` | `SourceFlowConfig` | yes |
| `destinationFlowConfigList` | `List<DestinationFlowConfig>` | yes |
| `tasks` | `List<Task>` | yes |
| `metadataCatalogConfig` | `MetadataCatalogConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowStatus` | `string` | no |

