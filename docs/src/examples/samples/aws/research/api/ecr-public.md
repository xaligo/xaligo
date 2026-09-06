# Amazon Elastic Container Registry Public

API version: 2020-10-30. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ecr-public/2020-10-30/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchCheckLayerAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `layerDigests` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `layers` | `List<Layer>` | no |
| `failures` | `List<LayerFailure>` | no |

## BatchDeleteImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageIds` | `List<ImageIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIds` | `List<ImageIdentifier>` | no |
| `failures` | `List<ImageFailure>` | no |

## CompleteLayerUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `uploadId` | `string` | yes |
| `layerDigests` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `uploadId` | `string` | no |
| `layerDigest` | `string` | no |

## CreateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `catalogData` | `RepositoryCatalogDataInput` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | no |
| `catalogData` | `RepositoryCatalogData` | no |

## DeleteRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | no |

## DeleteRepositoryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `policyText` | `string` | no |

## DescribeImageTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageTagDetails` | `List<ImageTagDetail>` | no |
| `nextToken` | `string` | no |

## DescribeImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageIds` | `List<ImageIdentifier>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageDetails` | `List<ImageDetail>` | no |
| `nextToken` | `string` | no |

## DescribeRegistries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registries` | `List<Registry>` | yes |
| `nextToken` | `string` | no |

## DescribeRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryNames` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositories` | `List<Repository>` | no |
| `nextToken` | `string` | no |

## GetAuthorizationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizationData` | `AuthorizationData` | no |

## GetRegistryCatalogData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryCatalogData` | `RegistryCatalogData` | yes |

## GetRepositoryCatalogData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalogData` | `RepositoryCatalogData` | no |

## GetRepositoryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `policyText` | `string` | no |

## InitiateLayerUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `uploadId` | `string` | no |
| `partSize` | `long` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## PutImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageManifest` | `string` | yes |
| `imageManifestMediaType` | `string` | no |
| `imageTag` | `string` | no |
| `imageDigest` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `image` | `Image` | no |

## PutRegistryCatalogData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `displayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryCatalogData` | `RegistryCatalogData` | yes |

## PutRepositoryCatalogData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `catalogData` | `RepositoryCatalogDataInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `catalogData` | `RepositoryCatalogData` | no |

## SetRepositoryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `policyText` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `policyText` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UploadLayerPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `uploadId` | `string` | yes |
| `partFirstByte` | `long` | yes |
| `partLastByte` | `long` | yes |
| `layerPartBlob` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `uploadId` | `string` | no |
| `lastByteReceived` | `long` | no |

