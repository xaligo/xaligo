# Amazon EventBridge Pipes

API version: 2015-10-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pipes/2015-10-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreatePipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DesiredState` | `string` | no |
| `Source` | `string` | yes |
| `SourceParameters` | `PipeSourceParameters` | no |
| `Enrichment` | `string` | no |
| `EnrichmentParameters` | `PipeEnrichmentParameters` | no |
| `Target` | `string` | yes |
| `TargetParameters` | `PipeTargetParameters` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `LogConfiguration` | `PipeLogConfigurationParameters` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DeletePipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribePipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `StateReason` | `string` | no |
| `Source` | `string` | no |
| `SourceParameters` | `PipeSourceParameters` | no |
| `Enrichment` | `string` | no |
| `EnrichmentParameters` | `PipeEnrichmentParameters` | no |
| `Target` | `string` | no |
| `TargetParameters` | `PipeTargetParameters` | no |
| `RoleArn` | `string` | no |
| `Tags` | `Map<string>` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LogConfiguration` | `PipeLogConfiguration` | no |
| `KmsKeyIdentifier` | `string` | no |

## ListPipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NamePrefix` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `SourcePrefix` | `string` | no |
| `TargetPrefix` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Pipes` | `List<Pipe>` | no |
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

## StartPipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## StopPipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

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


## UpdatePipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DesiredState` | `string` | no |
| `SourceParameters` | `UpdatePipeSourceParameters` | no |
| `Enrichment` | `string` | no |
| `EnrichmentParameters` | `PipeEnrichmentParameters` | no |
| `Target` | `string` | no |
| `TargetParameters` | `PipeTargetParameters` | no |
| `RoleArn` | `string` | yes |
| `LogConfiguration` | `PipeLogConfigurationParameters` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `DesiredState` | `string` | no |
| `CurrentState` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

