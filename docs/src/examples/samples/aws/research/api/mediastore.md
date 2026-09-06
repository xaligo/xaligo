# AWS Elemental MediaStore

API version: 2017-09-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mediastore/2017-09-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateContainer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Container` | `Container` | yes |

## DeleteContainer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContainerPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCorsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMetricPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeContainer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Container` | `Container` | no |

## GetContainerPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | yes |

## GetCorsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CorsPolicy` | `List<CorsRule>` | yes |

## GetLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LifecyclePolicy` | `string` | yes |

## GetMetricPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricPolicy` | `MetricPolicy` | yes |

## ListContainers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Containers` | `List<Container>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutContainerPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutCorsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |
| `CorsPolicy` | `List<CorsRule>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |
| `LifecyclePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutMetricPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |
| `MetricPolicy` | `MetricPolicy` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAccessLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopAccessLogging

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContainerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


