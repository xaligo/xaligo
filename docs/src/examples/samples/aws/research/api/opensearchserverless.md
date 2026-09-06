# OpenSearch Service Serverless

API version: 2021-11-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/opensearchserverless/2021-11-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchGetCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `names` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collectionDetails` | `List<CollectionDetail>` | no |
| `collectionErrorDetails` | `List<CollectionErrorDetail>` | no |

## BatchGetCollectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `names` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collectionGroupDetails` | `List<CollectionGroupDetail>` | no |
| `collectionGroupErrorDetails` | `List<CollectionGroupErrorDetail>` | no |

## BatchGetEffectiveLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceIdentifiers` | `List<LifecyclePolicyResourceIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `effectiveLifecyclePolicyDetails` | `List<EffectiveLifecyclePolicyDetail>` | no |
| `effectiveLifecyclePolicyErrorDetails` | `List<EffectiveLifecyclePolicyErrorDetail>` | no |

## BatchGetLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifiers` | `List<LifecyclePolicyIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyDetails` | `List<LifecyclePolicyDetail>` | no |
| `lifecyclePolicyErrorDetails` | `List<LifecyclePolicyErrorDetail>` | no |

## BatchGetVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcEndpointDetails` | `List<VpcEndpointDetail>` | no |
| `vpcEndpointErrorDetails` | `List<VpcEndpointErrorDetail>` | no |

## CreateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyDetail` | `AccessPolicyDetail` | no |

## CreateCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `type` | `string` | no |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `standbyReplicas` | `string` | no |
| `vectorOptions` | `VectorOptions` | no |
| `collectionGroupName` | `string` | no |
| `encryptionConfig` | `EncryptionConfig` | no |
| `deletionProtection` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createCollectionDetail` | `CreateCollectionDetail` | no |

## CreateCollectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `standbyReplicas` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `capacityLimits` | `CollectionGroupCapacityLimits` | no |
| `generation` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createCollectionGroupDetail` | `CreateCollectionGroupDetail` | no |

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `indexName` | `string` | yes |
| `indexSchema` | `IndexSchema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyDetail` | `LifecyclePolicyDetail` | no |

## CreateSecurityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `samlOptions` | `SamlConfigOptions` | no |
| `iamIdentityCenterOptions` | `CreateIamIdentityCenterConfigOptions` | no |
| `iamFederationOptions` | `IamFederationConfigOptions` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfigDetail` | `SecurityConfigDetail` | no |

## CreateSecurityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityPolicyDetail` | `SecurityPolicyDetail` | no |

## CreateVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createVpcEndpointDetail` | `CreateVpcEndpointDetail` | no |

## DeleteAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleteCollectionDetail` | `DeleteCollectionDetail` | no |

## DeleteCollectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `indexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecurityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecurityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deleteVpcEndpointDetail` | `DeleteVpcEndpointDetail` | no |

## GetAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyDetail` | `AccessPolicyDetail` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettingsDetail` | `AccountSettingsDetail` | no |

## GetIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `indexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `indexSchema` | `IndexSchema` | no |

## GetPoliciesStats

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessPolicyStats` | `AccessPolicyStats` | no |
| `SecurityPolicyStats` | `SecurityPolicyStats` | no |
| `SecurityConfigStats` | `SecurityConfigStats` | no |
| `LifecyclePolicyStats` | `LifecyclePolicyStats` | no |
| `TotalPolicyCount` | `long` | no |

## GetSecurityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfigDetail` | `SecurityConfigDetail` | no |

## GetSecurityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityPolicyDetail` | `SecurityPolicyDetail` | no |

## ListAccessPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `resource` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicySummaries` | `List<AccessPolicySummary>` | no |
| `nextToken` | `string` | no |

## ListCollectionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collectionGroupSummaries` | `List<CollectionGroupSummary>` | no |
| `nextToken` | `string` | no |

## ListCollections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collectionFilters` | `CollectionFilters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `collectionSummaries` | `List<CollectionSummary>` | no |
| `nextToken` | `string` | no |

## ListLifecyclePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `resources` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicySummaries` | `List<LifecyclePolicySummary>` | no |
| `nextToken` | `string` | no |

## ListSecurityConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfigSummaries` | `List<SecurityConfigSummary>` | no |
| `nextToken` | `string` | no |

## ListSecurityPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `resource` | `List<string>` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityPolicySummaries` | `List<SecurityPolicySummary>` | no |
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

## ListVpcEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcEndpointFilters` | `VpcEndpointFilters` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcEndpointSummaries` | `List<VpcEndpointSummary>` | no |
| `nextToken` | `string` | no |

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


## UpdateAccessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `policyVersion` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessPolicyDetail` | `AccessPolicyDetail` | no |

## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityLimits` | `CapacityLimits` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettingsDetail` | `AccountSettingsDetail` | no |

## UpdateCollection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `description` | `string` | no |
| `vectorOptions` | `VectorOptions` | no |
| `deletionProtection` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateCollectionDetail` | `UpdateCollectionDetail` | no |

## UpdateCollectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `description` | `string` | no |
| `capacityLimits` | `CollectionGroupCapacityLimits` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `updateCollectionGroupDetail` | `UpdateCollectionGroupDetail` | no |

## UpdateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `indexName` | `string` | yes |
| `indexSchema` | `IndexSchema` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `policyVersion` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyDetail` | `LifecyclePolicyDetail` | no |

## UpdateSecurityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `configVersion` | `string` | yes |
| `description` | `string` | no |
| `samlOptions` | `SamlConfigOptions` | no |
| `iamIdentityCenterOptionsUpdates` | `UpdateIamIdentityCenterConfigOptions` | no |
| `iamFederationOptions` | `IamFederationConfigOptions` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfigDetail` | `SecurityConfigDetail` | no |

## UpdateSecurityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | yes |
| `name` | `string` | yes |
| `policyVersion` | `string` | yes |
| `description` | `string` | no |
| `policy` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityPolicyDetail` | `SecurityPolicyDetail` | no |

## UpdateVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `addSubnetIds` | `List<string>` | no |
| `removeSubnetIds` | `List<string>` | no |
| `addSecurityGroupIds` | `List<string>` | no |
| `removeSecurityGroupIds` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpdateVpcEndpointDetail` | `UpdateVpcEndpointDetail` | no |

