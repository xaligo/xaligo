# Amazon Elastic Kubernetes Service

API version: 2017-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/eks/2017-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `certificateAuthorityId` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |
| `certificateAuthority` | `CertificateAuthoritySummary` | no |

## AssociateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |
| `policyArn` | `string` | yes |
| `accessScope` | `AccessScope` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | no |
| `principalArn` | `string` | no |
| `associatedAccessPolicy` | `AssociatedAccessPolicy` | no |

## AssociateEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `encryptionConfig` | `List<EncryptionConfig>` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## AssociateIdentityProviderConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `oidc` | `OidcIdentityProviderConfigRequest` | yes |
| `tags` | `Map<string>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |
| `tags` | `Map<string>` | no |

## CancelUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `updateId` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## CreateAccessEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |
| `kubernetesGroups` | `List<string>` | no |
| `tags` | `Map<string>` | no |
| `clientRequestToken` | `string` | no |
| `username` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessEntry` | `AccessEntry` | no |

## CreateAddon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `addonName` | `string` | yes |
| `addonVersion` | `string` | no |
| `serviceAccountRoleArn` | `string` | no |
| `resolveConflicts` | `string` | no |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `configurationValues` | `string` | no |
| `podIdentityAssociations` | `List<AddonPodIdentityAssociations>` | no |
| `namespaceConfig` | `AddonNamespaceConfigRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addon` | `Addon` | no |

## CreateCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityName` | `string` | yes |
| `clusterName` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `type` | `string` | yes |
| `roleArn` | `string` | yes |
| `configuration` | `CapabilityConfigurationRequest` | no |
| `tags` | `Map<string>` | no |
| `deletePropagationPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capability` | `Capability` | no |

## CreateCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |
| `certificateAuthority` | `CertificateAuthoritySummary` | no |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | no |
| `roleArn` | `string` | yes |
| `resourcesVpcConfig` | `VpcConfigRequest` | yes |
| `kubernetesNetworkConfig` | `KubernetesNetworkConfigRequest` | no |
| `logging` | `Logging` | no |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `encryptionConfig` | `List<EncryptionConfig>` | no |
| `outpostConfig` | `OutpostConfigRequest` | no |
| `accessConfig` | `CreateAccessConfigRequest` | no |
| `bootstrapSelfManagedAddons` | `boolean` | no |
| `upgradePolicy` | `UpgradePolicyRequest` | no |
| `zonalShiftConfig` | `ZonalShiftConfigRequest` | no |
| `remoteNetworkConfig` | `RemoteNetworkConfigRequest` | no |
| `computeConfig` | `ComputeConfigRequest` | no |
| `storageConfig` | `StorageConfigRequest` | no |
| `deletionProtection` | `boolean` | no |
| `controlPlaneScalingConfig` | `ControlPlaneScalingConfig` | no |
| `kubeApiServerConfig` | `KubeApiServerConfigRequest` | no |
| `kubeSchedulerConfig` | `KubeSchedulerConfigRequest` | no |
| `kubeControllerManagerConfig` | `KubeControllerManagerConfigRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## CreateEksAnywhereSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `term` | `EksAnywhereSubscriptionTerm` | yes |
| `licenseQuantity` | `integer` | no |
| `licenseType` | `string` | no |
| `autoRenew` | `boolean` | no |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `EksAnywhereSubscription` | no |

## CreateFargateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fargateProfileName` | `string` | yes |
| `clusterName` | `string` | yes |
| `podExecutionRoleArn` | `string` | yes |
| `subnets` | `List<string>` | no |
| `selectors` | `List<FargateProfileSelector>` | no |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fargateProfile` | `FargateProfile` | no |

## CreateNodegroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nodegroupName` | `string` | yes |
| `scalingConfig` | `NodegroupScalingConfig` | no |
| `diskSize` | `integer` | no |
| `subnets` | `List<string>` | yes |
| `instanceTypes` | `List<string>` | no |
| `amiType` | `string` | no |
| `remoteAccess` | `RemoteAccessConfig` | no |
| `nodeRole` | `string` | yes |
| `labels` | `Map<string>` | no |
| `taints` | `List<Taint>` | no |
| `tags` | `Map<string>` | no |
| `clientRequestToken` | `string` | no |
| `launchTemplate` | `LaunchTemplateSpecification` | no |
| `updateConfig` | `NodegroupUpdateConfig` | no |
| `nodeRepairConfig` | `NodeRepairConfig` | no |
| `capacityType` | `string` | no |
| `version` | `string` | no |
| `releaseVersion` | `string` | no |
| `warmPoolConfig` | `WarmPoolConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodegroup` | `Nodegroup` | no |

## CreatePodIdentityAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `namespace` | `string` | yes |
| `serviceAccount` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableSessionTags` | `boolean` | no |
| `targetRoleArn` | `string` | no |
| `policy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `PodIdentityAssociation` | no |

## DeleteAccessEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAddon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `addonName` | `string` | yes |
| `preserve` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addon` | `Addon` | no |

## DeleteCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `capabilityName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capability` | `Capability` | no |

## DeleteCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `certificateAuthorityId` | `string` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |
| `certificateAuthority` | `CertificateAuthoritySummary` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## DeleteEksAnywhereSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `EksAnywhereSubscription` | no |

## DeleteFargateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `fargateProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fargateProfile` | `FargateProfile` | no |

## DeleteNodegroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nodegroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodegroup` | `Nodegroup` | no |

## DeletePodIdentityAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `PodIdentityAssociation` | no |

## DeregisterCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## DescribeAccessEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessEntry` | `AccessEntry` | no |

## DescribeAddon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `addonName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addon` | `Addon` | no |

## DescribeAddonConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addonName` | `string` | yes |
| `addonVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addonName` | `string` | no |
| `addonVersion` | `string` | no |
| `configurationSchema` | `string` | no |
| `podIdentityConfiguration` | `List<AddonPodIdentityConfiguration>` | no |

## DescribeAddonVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `kubernetesVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `addonName` | `string` | no |
| `types` | `List<string>` | no |
| `publishers` | `List<string>` | no |
| `owners` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addons` | `List<AddonInfo>` | no |
| `nextToken` | `string` | no |

## DescribeCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `capabilityName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capability` | `Capability` | no |

## DescribeCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `certificateAuthorityId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateAuthority` | `CertificateAuthority` | no |

## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## DescribeClusterVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterType` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `defaultOnly` | `boolean` | no |
| `includeAll` | `boolean` | no |
| `clusterVersions` | `List<string>` | no |
| `status` | `string` | no |
| `versionStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `clusterVersions` | `List<ClusterVersionInformation>` | no |

## DescribeEksAnywhereSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `EksAnywhereSubscription` | no |

## DescribeFargateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `fargateProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fargateProfile` | `FargateProfile` | no |

## DescribeIdentityProviderConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `identityProviderConfig` | `IdentityProviderConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderConfig` | `IdentityProviderConfigResponse` | no |

## DescribeInsight

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `insight` | `Insight` | no |

## DescribeInsightsRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `status` | `string` | no |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |

## DescribeNodegroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nodegroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodegroup` | `Nodegroup` | no |

## DescribePodIdentityAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `PodIdentityAssociation` | no |

## DescribeUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `updateId` | `string` | yes |
| `nodegroupName` | `string` | no |
| `addonName` | `string` | no |
| `capabilityName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## DisassociateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateIdentityProviderConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `identityProviderConfig` | `IdentityProviderConfig` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## ListAccessEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `associatedPolicyArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessEntries` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicies` | `List<AccessPolicy>` | no |
| `nextToken` | `string` | no |

## ListAddons

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `addons` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListAssociatedAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | no |
| `principalArn` | `string` | no |
| `nextToken` | `string` | no |
| `associatedAccessPolicies` | `List<AssociatedAccessPolicy>` | no |

## ListCapabilities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilities` | `List<CapabilitySummary>` | no |
| `nextToken` | `string` | no |

## ListCertificateAuthorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateAuthorities` | `List<CertificateAuthoritySummary>` | no |
| `nextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusters` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListEksAnywhereSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `includeStatus` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriptions` | `List<EksAnywhereSubscription>` | no |
| `nextToken` | `string` | no |

## ListFargateProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `fargateProfileNames` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListIdentityProviderConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identityProviderConfigs` | `List<IdentityProviderConfig>` | no |
| `nextToken` | `string` | no |

## ListInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `filter` | `InsightsFilter` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `insights` | `List<InsightSummary>` | no |
| `nextToken` | `string` | no |

## ListNodegroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodegroups` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListPodIdentityAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `namespace` | `string` | no |
| `serviceAccount` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associations` | `List<PodIdentityAssociationSummary>` | no |
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

## ListUpdates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `nodegroupName` | `string` | no |
| `addonName` | `string` | no |
| `capabilityName` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateIds` | `List<string>` | no |
| `nextToken` | `string` | no |

## RegisterCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `connectorConfig` | `ConnectorConfigRequest` | yes |
| `clientRequestToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cluster` | `Cluster` | no |

## StartInsightsRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `message` | `string` | no |
| `status` | `string` | no |

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


## UpdateAccessEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `principalArn` | `string` | yes |
| `kubernetesGroups` | `List<string>` | no |
| `clientRequestToken` | `string` | no |
| `username` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessEntry` | `AccessEntry` | no |

## UpdateAddon

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `addonName` | `string` | yes |
| `addonVersion` | `string` | no |
| `serviceAccountRoleArn` | `string` | no |
| `resolveConflicts` | `string` | no |
| `clientRequestToken` | `string` | no |
| `configurationValues` | `string` | no |
| `podIdentityAssociations` | `List<AddonPodIdentityAssociations>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdateCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `capabilityName` | `string` | yes |
| `roleArn` | `string` | no |
| `configuration` | `UpdateCapabilityConfiguration` | no |
| `clientRequestToken` | `string` | no |
| `deletePropagationPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdateClusterConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `resourcesVpcConfig` | `VpcConfigRequest` | no |
| `logging` | `Logging` | no |
| `clientRequestToken` | `string` | no |
| `accessConfig` | `UpdateAccessConfigRequest` | no |
| `upgradePolicy` | `UpgradePolicyRequest` | no |
| `zonalShiftConfig` | `ZonalShiftConfigRequest` | no |
| `computeConfig` | `ComputeConfigRequest` | no |
| `kubernetesNetworkConfig` | `KubernetesNetworkConfigRequest` | no |
| `storageConfig` | `StorageConfigRequest` | no |
| `remoteNetworkConfig` | `RemoteNetworkConfigRequest` | no |
| `deletionProtection` | `boolean` | no |
| `controlPlaneScalingConfig` | `ControlPlaneScalingConfig` | no |
| `kubeApiServerConfig` | `KubeApiServerConfigRequest` | no |
| `kubeSchedulerConfig` | `KubeSchedulerConfigRequest` | no |
| `kubeControllerManagerConfig` | `KubeControllerManagerConfigRequest` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdateClusterVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `version` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `force` | `boolean` | no |
| `rollbackConfig` | `RollbackConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdateEksAnywhereSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `autoRenew` | `boolean` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscription` | `EksAnywhereSubscription` | no |

## UpdateNodegroupConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nodegroupName` | `string` | yes |
| `labels` | `UpdateLabelsPayload` | no |
| `taints` | `UpdateTaintsPayload` | no |
| `scalingConfig` | `NodegroupScalingConfig` | no |
| `updateConfig` | `NodegroupUpdateConfig` | no |
| `nodeRepairConfig` | `NodeRepairConfig` | no |
| `warmPoolConfig` | `WarmPoolConfig` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdateNodegroupVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `nodegroupName` | `string` | yes |
| `version` | `string` | no |
| `releaseVersion` | `string` | no |
| `launchTemplate` | `LaunchTemplateSpecification` | no |
| `force` | `boolean` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `update` | `Update` | no |

## UpdatePodIdentityAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clusterName` | `string` | yes |
| `associationId` | `string` | yes |
| `roleArn` | `string` | no |
| `clientRequestToken` | `string` | no |
| `disableSessionTags` | `boolean` | no |
| `targetRoleArn` | `string` | no |
| `policy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `association` | `PodIdentityAssociation` | no |

