# Amazon Elastic Container Registry

API version: 2015-09-21. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ecr/2015-09-21/service-2.json).

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

## BatchGetImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageIds` | `List<ImageIdentifier>` | yes |
| `acceptedMediaTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `images` | `List<Image>` | no |
| `failures` | `List<ImageFailure>` | no |

## BatchGetRepositoryScanningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanningConfigurations` | `List<RepositoryScanningConfiguration>` | no |
| `failures` | `List<RepositoryScanningConfigurationFailure>` | no |

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

## CreatePullThroughCacheRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | yes |
| `upstreamRegistryUrl` | `string` | yes |
| `registryId` | `string` | no |
| `upstreamRegistry` | `string` | no |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |
| `upstreamRepositoryPrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | no |
| `upstreamRegistryUrl` | `string` | no |
| `createdAt` | `timestamp` | no |
| `registryId` | `string` | no |
| `upstreamRegistry` | `string` | no |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |
| `upstreamRepositoryPrefix` | `string` | no |

## CreateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `imageTagMutability` | `string` | no |
| `imageTagMutabilityExclusionFilters` | `List<ImageTagMutabilityExclusionFilter>` | no |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | no |

## CreateRepositoryCreationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `prefix` | `string` | yes |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfigurationForRepositoryCreationTemplate` | no |
| `resourceTags` | `List<Tag>` | no |
| `imageTagMutability` | `string` | no |
| `imageTagMutabilityExclusionFilters` | `List<ImageTagMutabilityExclusionFilter>` | no |
| `repositoryPolicy` | `string` | no |
| `lifecyclePolicy` | `string` | no |
| `appliedFor` | `List<string>` | yes |
| `customRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryCreationTemplate` | `RepositoryCreationTemplate` | no |

## DeleteLifecyclePolicy

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
| `lifecyclePolicyText` | `string` | no |
| `lastEvaluatedAt` | `timestamp` | no |

## DeletePullThroughCacheRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | yes |
| `registryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | no |
| `upstreamRegistryUrl` | `string` | no |
| `createdAt` | `timestamp` | no |
| `registryId` | `string` | no |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |
| `upstreamRepositoryPrefix` | `string` | no |

## DeleteRegistryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `policyText` | `string` | no |

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

## DeleteRepositoryCreationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `prefix` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryCreationTemplate` | `RepositoryCreationTemplate` | no |

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

## DeleteSigningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `signingConfiguration` | `SigningConfiguration` | no |

## DeregisterPullTimeUpdateExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalArn` | `string` | no |

## DescribeImageReplicationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `imageId` | `ImageIdentifier` | yes |
| `registryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | no |
| `imageId` | `ImageIdentifier` | no |
| `replicationStatuses` | `List<ImageReplicationStatus>` | no |

## DescribeImageScanFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageId` | `ImageIdentifier` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `imageId` | `ImageIdentifier` | no |
| `imageScanStatus` | `ImageScanStatus` | no |
| `imageScanFindings` | `ImageScanFindings` | no |
| `nextToken` | `string` | no |

## DescribeImageSigningStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | yes |
| `imageId` | `ImageIdentifier` | yes |
| `registryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryName` | `string` | no |
| `imageId` | `ImageIdentifier` | no |
| `registryId` | `string` | no |
| `signingStatuses` | `List<ImageSigningStatus>` | no |

## DescribeImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageIds` | `List<ImageIdentifier>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `DescribeImagesFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageDetails` | `List<ImageDetail>` | no |
| `nextToken` | `string` | no |

## DescribePullThroughCacheRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `ecrRepositoryPrefixes` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullThroughCacheRules` | `List<PullThroughCacheRule>` | no |
| `nextToken` | `string` | no |

## DescribeRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `replicationConfiguration` | `ReplicationConfiguration` | no |

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

## DescribeRepositoryCreationTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `prefixes` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryCreationTemplates` | `List<RepositoryCreationTemplate>` | no |
| `nextToken` | `string` | no |

## GetAccountSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `value` | `string` | no |

## GetAuthorizationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizationData` | `List<AuthorizationData>` | no |

## GetDownloadUrlForLayer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `layerDigest` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `downloadUrl` | `string` | no |
| `layerDigest` | `string` | no |

## GetLifecyclePolicy

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
| `lifecyclePolicyText` | `string` | no |
| `lastEvaluatedAt` | `timestamp` | no |

## GetLifecyclePolicyPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageIds` | `List<ImageIdentifier>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `LifecyclePolicyPreviewFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `lifecyclePolicyText` | `string` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `previewResults` | `List<LifecyclePolicyPreviewResult>` | no |
| `summary` | `LifecyclePolicyPreviewSummary` | no |

## GetRegistryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `policyText` | `string` | no |

## GetRegistryScanningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `scanningConfiguration` | `RegistryScanningConfiguration` | no |

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

## GetSigningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `signingConfiguration` | `SigningConfiguration` | no |

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

## ListImageReferrers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `subjectId` | `SubjectIdentifier` | yes |
| `filter` | `ListImageReferrersFilter` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `referrers` | `List<ImageReferrer>` | no |
| `nextToken` | `string` | no |

## ListImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `ListImagesFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageIds` | `List<ImageIdentifier>` | no |
| `nextToken` | `string` | no |

## ListPullTimeUpdateExclusions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pullTimeUpdateExclusions` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## PutAccountSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `value` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `value` | `string` | no |

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

## PutImageScanningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | no |

## PutImageTagMutability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageTagMutability` | `string` | yes |
| `imageTagMutabilityExclusionFilters` | `List<ImageTagMutabilityExclusionFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `imageTagMutability` | `string` | no |
| `imageTagMutabilityExclusionFilters` | `List<ImageTagMutabilityExclusionFilter>` | no |

## PutLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `lifecyclePolicyText` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `lifecyclePolicyText` | `string` | no |

## PutRegistryPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyText` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `policyText` | `string` | no |

## PutRegistryScanningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `scanType` | `string` | no |
| `rules` | `List<RegistryScanningRule>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryScanningConfiguration` | `RegistryScanningConfiguration` | no |

## PutReplicationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfiguration` | `ReplicationConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `replicationConfiguration` | `ReplicationConfiguration` | no |

## PutSigningConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signingConfiguration` | `SigningConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `signingConfiguration` | `SigningConfiguration` | no |

## RegisterPullTimeUpdateExclusion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `principalArn` | `string` | no |
| `createdAt` | `timestamp` | no |

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

## StartImageScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageId` | `ImageIdentifier` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `imageId` | `ImageIdentifier` | no |
| `imageScanStatus` | `ImageScanStatus` | no |

## StartLifecyclePolicyPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `lifecyclePolicyText` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `lifecyclePolicyText` | `string` | no |
| `status` | `string` | no |

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


## UpdateImageStorageClass

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | yes |
| `imageId` | `ImageIdentifier` | yes |
| `targetStorageClass` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryName` | `string` | no |
| `imageId` | `ImageIdentifier` | no |
| `imageStatus` | `string` | no |

## UpdatePullThroughCacheRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `ecrRepositoryPrefix` | `string` | yes |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | no |
| `registryId` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |
| `upstreamRepositoryPrefix` | `string` | no |

## UpdateRepositoryCreationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `prefix` | `string` | yes |
| `description` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfigurationForRepositoryCreationTemplate` | no |
| `resourceTags` | `List<Tag>` | no |
| `imageTagMutability` | `string` | no |
| `imageTagMutabilityExclusionFilters` | `List<ImageTagMutabilityExclusionFilter>` | no |
| `repositoryPolicy` | `string` | no |
| `lifecyclePolicy` | `string` | no |
| `appliedFor` | `List<string>` | no |
| `customRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | no |
| `repositoryCreationTemplate` | `RepositoryCreationTemplate` | no |

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

## ValidatePullThroughCacheRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | yes |
| `registryId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ecrRepositoryPrefix` | `string` | no |
| `registryId` | `string` | no |
| `upstreamRegistryUrl` | `string` | no |
| `credentialArn` | `string` | no |
| `customRoleArn` | `string` | no |
| `upstreamRepositoryPrefix` | `string` | no |
| `isValid` | `boolean` | no |
| `failure` | `string` | no |

