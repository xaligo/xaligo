# CodeArtifact

API version: 2018-09-22. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/codeartifact/2018-09-22/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateExternalConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `externalConnection` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

## CopyPackageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `sourceRepository` | `string` | yes |
| `destinationRepository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `versions` | `List<string>` | no |
| `versionRevisions` | `Map<string>` | no |
| `allowOverwrite` | `boolean` | no |
| `includeFromUpstream` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulVersions` | `Map<SuccessfulPackageVersionInfo>` | no |
| `failedVersions` | `Map<PackageVersionError>` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `encryptionKey` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `DomainDescription` | no |

## CreatePackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `contactInfo` | `string` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |

## CreateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `description` | `string` | no |
| `upstreams` | `List<UpstreamRepository>` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `DomainDescription` | no |

## DeleteDomainPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `policyRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

## DeletePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deletedPackage` | `PackageSummary` | no |

## DeletePackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |

## DeletePackageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `versions` | `List<string>` | yes |
| `expectedStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulVersions` | `Map<SuccessfulPackageVersionInfo>` | no |
| `failedVersions` | `Map<PackageVersionError>` | no |

## DeleteRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

## DeleteRepositoryPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `policyRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

## DescribeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `DomainDescription` | no |

## DescribePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `package` | `PackageDescription` | yes |

## DescribePackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |

## DescribePackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageVersion` | `PackageVersionDescription` | yes |

## DescribeRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

## DisassociateExternalConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `externalConnection` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

## DisposePackageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `versions` | `List<string>` | yes |
| `versionRevisions` | `Map<string>` | no |
| `expectedStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulVersions` | `Map<SuccessfulPackageVersionInfo>` | no |
| `failedVersions` | `Map<PackageVersionError>` | no |

## GetAssociatedPackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |
| `associationType` | `string` | no |

## GetAuthorizationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `durationSeconds` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authorizationToken` | `string` | no |
| `expiration` | `timestamp` | no |

## GetDomainPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

## GetPackageVersionAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |
| `asset` | `string` | yes |
| `packageVersionRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `asset` | `blob` | no |
| `assetName` | `string` | no |
| `packageVersion` | `string` | no |
| `packageVersionRevision` | `string` | no |

## GetPackageVersionReadme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `format` | `string` | no |
| `namespace` | `string` | no |
| `package` | `string` | no |
| `version` | `string` | no |
| `versionRevision` | `string` | no |
| `readme` | `string` | no |

## GetRepositoryEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `endpointType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryEndpoint` | `string` | no |

## GetRepositoryPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

## ListAllowedRepositoriesForGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `originRestrictionType` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `allowedRepositories` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListAssociatedPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `preview` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packages` | `List<AssociatedPackage>` | no |
| `nextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domains` | `List<DomainSummary>` | no |
| `nextToken` | `string` | no |

## ListPackageGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `prefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroups` | `List<PackageGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListPackageVersionAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `format` | `string` | no |
| `namespace` | `string` | no |
| `package` | `string` | no |
| `version` | `string` | no |
| `versionRevision` | `string` | no |
| `nextToken` | `string` | no |
| `assets` | `List<AssetSummary>` | no |

## ListPackageVersionDependencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `format` | `string` | no |
| `namespace` | `string` | no |
| `package` | `string` | no |
| `version` | `string` | no |
| `versionRevision` | `string` | no |
| `nextToken` | `string` | no |
| `dependencies` | `List<PackageDependency>` | no |

## ListPackageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `status` | `string` | no |
| `sortBy` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `originType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `defaultDisplayVersion` | `string` | no |
| `format` | `string` | no |
| `namespace` | `string` | no |
| `package` | `string` | no |
| `versions` | `List<PackageVersionSummary>` | no |
| `nextToken` | `string` | no |

## ListPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | no |
| `namespace` | `string` | no |
| `packagePrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `publish` | `string` | no |
| `upstream` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packages` | `List<PackageSummary>` | no |
| `nextToken` | `string` | no |

## ListRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositoryPrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositories` | `List<RepositorySummary>` | no |
| `nextToken` | `string` | no |

## ListRepositoriesInDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `administratorAccount` | `string` | no |
| `repositoryPrefix` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repositories` | `List<RepositorySummary>` | no |
| `nextToken` | `string` | no |

## ListSubPackageGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroups` | `List<PackageGroupSummary>` | no |
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

## PublishPackageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `packageVersion` | `string` | yes |
| `assetContent` | `blob` | yes |
| `assetName` | `string` | yes |
| `assetSHA256` | `string` | yes |
| `unfinished` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `format` | `string` | no |
| `namespace` | `string` | no |
| `package` | `string` | no |
| `version` | `string` | no |
| `versionRevision` | `string` | no |
| `status` | `string` | no |
| `asset` | `AssetSummary` | no |

## PutDomainPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `policyRevision` | `string` | no |
| `policyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

## PutPackageOriginConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `restrictions` | `PackageOriginRestrictions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `originConfiguration` | `PackageOriginConfiguration` | no |

## PutRepositoryPermissionsPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `policyRevision` | `string` | no |
| `policyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `ResourcePolicy` | no |

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


## UpdatePackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `contactInfo` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |

## UpdatePackageGroupOriginConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `packageGroup` | `string` | yes |
| `restrictions` | `Map<string>` | no |
| `addAllowedRepositories` | `List<PackageGroupAllowedRepository>` | no |
| `removeAllowedRepositories` | `List<PackageGroupAllowedRepository>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `packageGroup` | `PackageGroupDescription` | no |
| `allowedRepositoryUpdates` | `Map<Map<List<string>>>` | no |

## UpdatePackageVersionsStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `format` | `string` | yes |
| `namespace` | `string` | no |
| `package` | `string` | yes |
| `versions` | `List<string>` | yes |
| `versionRevisions` | `Map<string>` | no |
| `expectedStatus` | `string` | no |
| `targetStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `successfulVersions` | `Map<SuccessfulPackageVersionInfo>` | no |
| `failedVersions` | `Map<PackageVersionError>` | no |

## UpdateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domain` | `string` | yes |
| `domainOwner` | `string` | no |
| `repository` | `string` | yes |
| `description` | `string` | no |
| `upstreams` | `List<UpstreamRepository>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `RepositoryDescription` | no |

