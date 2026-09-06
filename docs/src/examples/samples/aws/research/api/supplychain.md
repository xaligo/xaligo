# AWS Supply Chain

API version: 2024-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/supplychain/2024-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBillOfMaterialsImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `s3uri` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobId` | `string` | yes |

## CreateDataIntegrationFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |
| `sources` | `List<DataIntegrationFlowSource>` | yes |
| `transformation` | `DataIntegrationFlowTransformation` | yes |
| `target` | `DataIntegrationFlowTarget` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

## CreateDataLakeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `schema` | `DataLakeDatasetSchema` | no |
| `description` | `string` | no |
| `partitionSpec` | `DataLakeDatasetPartitionSpec` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataset` | `DataLakeDataset` | yes |

## CreateDataLakeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `DataLakeNamespace` | yes |

## CreateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceName` | `string` | no |
| `instanceDescription` | `string` | no |
| `kmsKeyArn` | `string` | no |
| `webAppDnsDomain` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instance` | `Instance` | yes |

## DeleteDataIntegrationFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

## DeleteDataLakeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

## DeleteDataLakeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

## DeleteInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instance` | `Instance` | yes |

## GetBillOfMaterialsImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `jobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `job` | `BillOfMaterialsImportJob` | yes |

## GetDataIntegrationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `eventId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `event` | `DataIntegrationEvent` | yes |

## GetDataIntegrationFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flow` | `DataIntegrationFlow` | yes |

## GetDataIntegrationFlowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `flowName` | `string` | yes |
| `executionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecution` | `DataIntegrationFlowExecution` | yes |

## GetDataLakeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataset` | `DataLakeDataset` | yes |

## GetDataLakeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `DataLakeNamespace` | yes |

## GetInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instance` | `Instance` | yes |

## ListDataIntegrationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `eventType` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `events` | `List<DataIntegrationEvent>` | yes |
| `nextToken` | `string` | no |

## ListDataIntegrationFlowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `flowName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flowExecutions` | `List<DataIntegrationFlowExecution>` | yes |
| `nextToken` | `string` | no |

## ListDataIntegrationFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flows` | `List<DataIntegrationFlow>` | yes |
| `nextToken` | `string` | no |

## ListDataLakeDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasets` | `List<DataLakeDataset>` | yes |
| `nextToken` | `string` | no |

## ListDataLakeNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespaces` | `List<DataLakeNamespace>` | yes |
| `nextToken` | `string` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `instanceNameFilter` | `List<string>` | no |
| `instanceStateFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instances` | `List<Instance>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | yes |

## SendDataIntegrationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `eventType` | `string` | yes |
| `data` | `string` | yes |
| `eventGroupId` | `string` | yes |
| `eventTimestamp` | `timestamp` | no |
| `clientToken` | `string` | no |
| `datasetTarget` | `DataIntegrationEventDatasetTargetConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventId` | `string` | yes |

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


## UpdateDataIntegrationFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |
| `sources` | `List<DataIntegrationFlowSource>` | no |
| `transformation` | `DataIntegrationFlowTransformation` | no |
| `target` | `DataIntegrationFlowTarget` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `flow` | `DataIntegrationFlow` | yes |

## UpdateDataLakeDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `namespace` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataset` | `DataLakeDataset` | yes |

## UpdateDataLakeNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `namespace` | `DataLakeNamespace` | yes |

## UpdateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instanceId` | `string` | yes |
| `instanceName` | `string` | no |
| `instanceDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `instance` | `Instance` | yes |

