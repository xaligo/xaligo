# Amazon OpenSearch Ingestion

API version: 2022-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/osis/2022-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `MinUnits` | `integer` | yes |
| `MaxUnits` | `integer` | yes |
| `PipelineConfigurationBody` | `string` | yes |
| `LogPublishingOptions` | `LogPublishingOptions` | no |
| `VpcOptions` | `VpcOptions` | no |
| `BufferOptions` | `BufferOptions` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `Tags` | `List<Tag>` | no |
| `PipelineRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `Pipeline` | no |

## CreatePipelineEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | yes |
| `VpcOptions` | `PipelineEndpointVpcOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |
| `EndpointId` | `string` | no |
| `Status` | `string` | no |
| `VpcId` | `string` | no |

## DeletePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePipelineEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `Pipeline` | no |

## GetPipelineBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlueprintName` | `string` | yes |
| `Format` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blueprint` | `PipelineBlueprint` | no |
| `Format` | `string` | no |

## GetPipelineChangeProgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeProgressStatuses` | `List<ChangeProgressStatus>` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |

## ListPipelineBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Blueprints` | `List<PipelineBlueprintSummary>` | no |

## ListPipelineEndpointConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PipelineEndpointConnections` | `List<PipelineEndpointConnection>` | no |

## ListPipelineEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PipelineEndpoints` | `List<PipelineEndpoint>` | no |

## ListPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Pipelines` | `List<PipelineSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `Policy` | `string` | no |

## RevokePipelineEndpointConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | yes |
| `EndpointIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |

## StartPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `Pipeline` | no |

## StopPipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `Pipeline` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `MinUnits` | `integer` | no |
| `MaxUnits` | `integer` | no |
| `PipelineConfigurationBody` | `string` | no |
| `LogPublishingOptions` | `LogPublishingOptions` | no |
| `BufferOptions` | `BufferOptions` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `PipelineRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipeline` | `Pipeline` | no |

## ValidatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineConfigurationBody` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isValid` | `boolean` | no |
| `Errors` | `List<ValidationMessage>` | no |

