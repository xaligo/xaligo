# Amazon CloudFront

API version: 2020-05-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudfront/2020-05-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetDistributionId` | `string` | yes |
| `Alias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateDistributionTenantWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WebACLArn` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `WebACLArn` | `string` | no |
| `ETag` | `string` | no |

## AssociateDistributionWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `WebACLArn` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `WebACLArn` | `string` | no |
| `ETag` | `string` | no |

## CopyDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrimaryDistributionId` | `string` | yes |
| `Staging` | `boolean` | no |
| `IfMatch` | `string` | no |
| `CallerReference` | `string` | yes |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateAnycastIpList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IpCount` | `integer` | yes |
| `Tags` | `Tags` | no |
| `IpAddressType` | `string` | no |
| `IpamCidrConfigs` | `List<IpamCidrConfig>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnycastIpList` | `AnycastIpList` | no |
| `ETag` | `string` | no |

## CreateCachePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicyConfig` | `CachePolicyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicy` | `CachePolicy` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateCloudFrontOriginAccessIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentityConfig` | `CloudFrontOriginAccessIdentityConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentity` | `CloudFrontOriginAccessIdentity` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ConnectionFunctionConfig` | `FunctionConfig` | yes |
| `ConnectionFunctionCode` | `blob` | yes |
| `Tags` | `Tags` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionSummary` | `ConnectionFunctionSummary` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateConnectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Ipv6Enabled` | `boolean` | no |
| `Tags` | `Tags` | no |
| `AnycastIpListId` | `string` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionGroup` | `ConnectionGroup` | no |
| `ETag` | `string` | no |

## CreateContinuousDeploymentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicyConfig` | `ContinuousDeploymentPolicyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicy` | `ContinuousDeploymentPolicy` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionConfig` | `DistributionConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `Name` | `string` | yes |
| `Domains` | `List<DomainItem>` | yes |
| `Tags` | `Tags` | no |
| `Customizations` | `Customizations` | no |
| `Parameters` | `List<Parameter>` | no |
| `ConnectionGroupId` | `string` | no |
| `ManagedCertificateRequest` | `ManagedCertificateRequest` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionTenant` | `DistributionTenant` | no |
| `ETag` | `string` | no |

## CreateDistributionWithTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionConfigWithTags` | `DistributionConfigWithTags` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateFieldLevelEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionConfig` | `FieldLevelEncryptionConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryption` | `FieldLevelEncryption` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateFieldLevelEncryptionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfileConfig` | `FieldLevelEncryptionProfileConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfile` | `FieldLevelEncryptionProfile` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `FunctionConfig` | `FunctionConfig` | yes |
| `FunctionCode` | `blob` | yes |
| `Tags` | `Tags` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionSummary` | `FunctionSummary` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateInvalidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `InvalidationBatch` | `InvalidationBatch` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `string` | no |
| `Invalidation` | `Invalidation` | no |

## CreateInvalidationForDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `InvalidationBatch` | `InvalidationBatch` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `string` | no |
| `Invalidation` | `Invalidation` | no |

## CreateKeyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroupConfig` | `KeyGroupConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroup` | `KeyGroup` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateKeyValueStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Comment` | `string` | no |
| `ImportSource` | `ImportSource` | no |
| `Tags` | `Tags` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyValueStore` | `KeyValueStore` | no |
| `ETag` | `string` | no |
| `Location` | `string` | no |

## CreateMonitoringSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `MonitoringSubscription` | `MonitoringSubscription` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringSubscription` | `MonitoringSubscription` | no |

## CreateOriginAccessControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControlConfig` | `OriginAccessControlConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControl` | `OriginAccessControl` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateOriginRequestPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicyConfig` | `OriginRequestPolicyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicy` | `OriginRequestPolicy` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreatePublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyConfig` | `PublicKeyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKey` | `PublicKey` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateRealtimeLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndPoints` | `List<EndPoint>` | yes |
| `Fields` | `List<string>` | yes |
| `Name` | `string` | yes |
| `SamplingRate` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RealtimeLogConfig` | `RealtimeLogConfig` | no |

## CreateResponseHeadersPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicyConfig` | `ResponseHeadersPolicyConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicy` | `ResponseHeadersPolicy` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateStreamingDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistributionConfig` | `StreamingDistributionConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistribution` | `StreamingDistribution` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateStreamingDistributionWithTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistributionConfigWithTags` | `StreamingDistributionConfigWithTags` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistribution` | `StreamingDistribution` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## CreateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CaCertificatesBundleSource` | `CaCertificatesBundleSource` | yes |
| `UseClientCertificateOCSPEndpoint` | `boolean` | no |
| `Tags` | `Tags` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStore` | `TrustStore` | no |
| `ETag` | `string` | no |

## CreateVpcOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOriginEndpointConfig` | `VpcOriginEndpointConfig` | yes |
| `Tags` | `Tags` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOrigin` | `VpcOrigin` | no |
| `Location` | `string` | no |
| `ETag` | `string` | no |

## DeleteAnycastIpList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCachePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCloudFrontOriginAccessIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContinuousDeploymentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFieldLevelEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFieldLevelEncryptionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKeyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteKeyValueStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMonitoringSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOriginAccessControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOriginRequestPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRealtimeLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ARN` | `string` | no |

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


## DeleteResponseHeadersPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStreamingDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOrigin` | `VpcOrigin` | no |
| `ETag` | `string` | no |

## DescribeConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionSummary` | `ConnectionFunctionSummary` | no |
| `ETag` | `string` | no |

## DescribeFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionSummary` | `FunctionSummary` | no |
| `ETag` | `string` | no |

## DescribeKeyValueStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyValueStore` | `KeyValueStore` | no |
| `ETag` | `string` | no |

## DisassociateDistributionTenantWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ETag` | `string` | no |

## DisassociateDistributionWebACL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ETag` | `string` | no |

## GetAnycastIpList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnycastIpList` | `AnycastIpList` | no |
| `ETag` | `string` | no |

## GetCachePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicy` | `CachePolicy` | no |
| `ETag` | `string` | no |

## GetCachePolicyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicyConfig` | `CachePolicyConfig` | no |
| `ETag` | `string` | no |

## GetCloudFrontOriginAccessIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentity` | `CloudFrontOriginAccessIdentity` | no |
| `ETag` | `string` | no |

## GetCloudFrontOriginAccessIdentityConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentityConfig` | `CloudFrontOriginAccessIdentityConfig` | no |
| `ETag` | `string` | no |

## GetConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionCode` | `blob` | no |
| `ETag` | `string` | no |
| `ContentType` | `string` | no |

## GetConnectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionGroup` | `ConnectionGroup` | no |
| `ETag` | `string` | no |

## GetConnectionGroupByRoutingEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingEndpoint` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionGroup` | `ConnectionGroup` | no |
| `ETag` | `string` | no |

## GetContinuousDeploymentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicy` | `ContinuousDeploymentPolicy` | no |
| `ETag` | `string` | no |

## GetContinuousDeploymentPolicyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicyConfig` | `ContinuousDeploymentPolicyConfig` | no |
| `ETag` | `string` | no |

## GetDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `ETag` | `string` | no |

## GetDistributionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionConfig` | `DistributionConfig` | no |
| `ETag` | `string` | no |

## GetDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionTenant` | `DistributionTenant` | no |
| `ETag` | `string` | no |

## GetDistributionTenantByDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionTenant` | `DistributionTenant` | no |
| `ETag` | `string` | no |

## GetFieldLevelEncryption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryption` | `FieldLevelEncryption` | no |
| `ETag` | `string` | no |

## GetFieldLevelEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionConfig` | `FieldLevelEncryptionConfig` | no |
| `ETag` | `string` | no |

## GetFieldLevelEncryptionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfile` | `FieldLevelEncryptionProfile` | no |
| `ETag` | `string` | no |

## GetFieldLevelEncryptionProfileConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfileConfig` | `FieldLevelEncryptionProfileConfig` | no |
| `ETag` | `string` | no |

## GetFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionCode` | `blob` | no |
| `ETag` | `string` | no |
| `ContentType` | `string` | no |

## GetInvalidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invalidation` | `Invalidation` | no |

## GetInvalidationForDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionTenantId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Invalidation` | `Invalidation` | no |

## GetKeyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroup` | `KeyGroup` | no |
| `ETag` | `string` | no |

## GetKeyGroupConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroupConfig` | `KeyGroupConfig` | no |
| `ETag` | `string` | no |

## GetManagedCertificateDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedCertificateDetails` | `ManagedCertificateDetails` | no |

## GetMonitoringSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringSubscription` | `MonitoringSubscription` | no |

## GetOriginAccessControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControl` | `OriginAccessControl` | no |
| `ETag` | `string` | no |

## GetOriginAccessControlConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControlConfig` | `OriginAccessControlConfig` | no |
| `ETag` | `string` | no |

## GetOriginRequestPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicy` | `OriginRequestPolicy` | no |
| `ETag` | `string` | no |

## GetOriginRequestPolicyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicyConfig` | `OriginRequestPolicyConfig` | no |
| `ETag` | `string` | no |

## GetPublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKey` | `PublicKey` | no |
| `ETag` | `string` | no |

## GetPublicKeyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyConfig` | `PublicKeyConfig` | no |
| `ETag` | `string` | no |

## GetRealtimeLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ARN` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RealtimeLogConfig` | `RealtimeLogConfig` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `PolicyDocument` | `string` | no |

## GetResponseHeadersPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicy` | `ResponseHeadersPolicy` | no |
| `ETag` | `string` | no |

## GetResponseHeadersPolicyConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicyConfig` | `ResponseHeadersPolicyConfig` | no |
| `ETag` | `string` | no |

## GetStreamingDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistribution` | `StreamingDistribution` | no |
| `ETag` | `string` | no |

## GetStreamingDistributionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistributionConfig` | `StreamingDistributionConfig` | no |
| `ETag` | `string` | no |

## GetTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStore` | `TrustStore` | no |
| `ETag` | `string` | no |

## GetVpcOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOrigin` | `VpcOrigin` | no |
| `ETag` | `string` | no |

## ListAnycastIpLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnycastIpLists` | `AnycastIpListCollection` | no |

## ListCachePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicyList` | `CachePolicyList` | no |

## ListCloudFrontOriginAccessIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentityList` | `CloudFrontOriginAccessIdentityList` | no |

## ListConflictingAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `Alias` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConflictingAliasesList` | `ConflictingAliasesList` | no |

## ListConnectionFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ConnectionFunctions` | `List<ConnectionFunctionSummary>` | no |

## ListConnectionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationFilter` | `ConnectionGroupAssociationFilter` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `ConnectionGroups` | `List<ConnectionGroupSummary>` | no |

## ListContinuousDeploymentPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicyList` | `ContinuousDeploymentPolicyList` | no |

## ListDistributionTenants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationFilter` | `DistributionTenantAssociationFilter` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `DistributionTenantList` | `List<DistributionTenantSummary>` | no |

## ListDistributionTenantsByCustomization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebACLArn` | `string` | no |
| `CertificateArn` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `DistributionTenantList` | `List<DistributionTenantSummary>` | no |

## ListDistributions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByAnycastIpListId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `AnycastIpListId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByCachePolicyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `CachePolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionIdList` | `DistributionIdList` | no |

## ListDistributionsByConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |
| `ConnectionFunctionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByConnectionMode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |
| `ConnectionMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByKeyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `KeyGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionIdList` | `DistributionIdList` | no |

## ListDistributionsByOriginRequestPolicyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `OriginRequestPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionIdList` | `DistributionIdList` | no |

## ListDistributionsByOwnedResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionIdOwnerList` | no |

## ListDistributionsByRealtimeLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `RealtimeLogConfigName` | `string` | no |
| `RealtimeLogConfigArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByResponseHeadersPolicyId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `ResponseHeadersPolicyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionIdList` | `DistributionIdList` | no |

## ListDistributionsByTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreIdentifier` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDistributionsByVpcOriginId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `VpcOriginId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionIdList` | `DistributionIdList` | no |

## ListDistributionsByWebACLId

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `WebACLId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionList` | `DistributionList` | no |

## ListDomainConflicts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `DomainControlValidationResource` | `DistributionResourceId` | yes |
| `MaxItems` | `integer` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainConflicts` | `List<DomainConflict>` | no |
| `NextMarker` | `string` | no |

## ListFieldLevelEncryptionConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionList` | `FieldLevelEncryptionList` | no |

## ListFieldLevelEncryptionProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfileList` | `FieldLevelEncryptionProfileList` | no |

## ListFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionList` | `FunctionList` | no |

## ListInvalidations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionId` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvalidationList` | `InvalidationList` | no |

## ListInvalidationsForDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InvalidationList` | `InvalidationList` | no |

## ListKeyGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroupList` | `KeyGroupList` | no |

## ListKeyValueStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyValueStoreList` | `KeyValueStoreList` | no |

## ListOriginAccessControls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControlList` | `OriginAccessControlList` | no |

## ListOriginRequestPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicyList` | `OriginRequestPolicyList` | no |

## ListPublicKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyList` | `PublicKeyList` | no |

## ListRealtimeLogConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxItems` | `string` | no |
| `Marker` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RealtimeLogConfigs` | `RealtimeLogConfigs` | no |

## ListResponseHeadersPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicyList` | `ResponseHeadersPolicyList` | no |

## ListStreamingDistributions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistributionList` | `StreamingDistributionList` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Tags` | yes |

## ListTrustStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextMarker` | `string` | no |
| `TrustStoreList` | `List<TrustStoreSummary>` | no |

## ListVpcOrigins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `MaxItems` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOriginList` | `VpcOriginList` | no |

## PublishConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionSummary` | `ConnectionFunctionSummary` | no |

## PublishFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionSummary` | `FunctionSummary` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PolicyDocument` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `Tags` | `Tags` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |
| `Stage` | `string` | no |
| `ConnectionObject` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionTestResult` | `ConnectionFunctionTestResult` | no |

## TestFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IfMatch` | `string` | yes |
| `Stage` | `string` | no |
| `EventObject` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestResult` | `TestResult` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `TagKeys` | `TagKeys` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAnycastIpList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IpAddressType` | `string` | no |
| `IpamCidrConfigs` | `List<IpamCidrConfig>` | no |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AnycastIpList` | `AnycastIpList` | no |
| `ETag` | `string` | no |

## UpdateCachePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicyConfig` | `CachePolicyConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CachePolicy` | `CachePolicy` | no |
| `ETag` | `string` | no |

## UpdateCloudFrontOriginAccessIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentityConfig` | `CloudFrontOriginAccessIdentityConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CloudFrontOriginAccessIdentity` | `CloudFrontOriginAccessIdentity` | no |
| `ETag` | `string` | no |

## UpdateConnectionFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |
| `ConnectionFunctionConfig` | `FunctionConfig` | yes |
| `ConnectionFunctionCode` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionFunctionSummary` | `ConnectionFunctionSummary` | no |
| `ETag` | `string` | no |

## UpdateConnectionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Ipv6Enabled` | `boolean` | no |
| `IfMatch` | `string` | yes |
| `AnycastIpListId` | `string` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionGroup` | `ConnectionGroup` | no |
| `ETag` | `string` | no |

## UpdateContinuousDeploymentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicyConfig` | `ContinuousDeploymentPolicyConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuousDeploymentPolicy` | `ContinuousDeploymentPolicy` | no |
| `ETag` | `string` | no |

## UpdateDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionConfig` | `DistributionConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `ETag` | `string` | no |

## UpdateDistributionTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `DistributionId` | `string` | no |
| `Domains` | `List<DomainItem>` | no |
| `Customizations` | `Customizations` | no |
| `Parameters` | `List<Parameter>` | no |
| `ConnectionGroupId` | `string` | no |
| `IfMatch` | `string` | yes |
| `ManagedCertificateRequest` | `ManagedCertificateRequest` | no |
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DistributionTenant` | `DistributionTenant` | no |
| `ETag` | `string` | no |

## UpdateDistributionWithStagingConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `StagingDistributionId` | `string` | no |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Distribution` | `Distribution` | no |
| `ETag` | `string` | no |

## UpdateDomainAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `TargetResource` | `DistributionResourceId` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | no |
| `ResourceId` | `string` | no |
| `ETag` | `string` | no |

## UpdateFieldLevelEncryptionConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionConfig` | `FieldLevelEncryptionConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryption` | `FieldLevelEncryption` | no |
| `ETag` | `string` | no |

## UpdateFieldLevelEncryptionProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfileConfig` | `FieldLevelEncryptionProfileConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FieldLevelEncryptionProfile` | `FieldLevelEncryptionProfile` | no |
| `ETag` | `string` | no |

## UpdateFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `IfMatch` | `string` | yes |
| `FunctionConfig` | `FunctionConfig` | yes |
| `FunctionCode` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionSummary` | `FunctionSummary` | no |
| `ETag` | `string` | no |

## UpdateKeyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroupConfig` | `KeyGroupConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyGroup` | `KeyGroup` | no |
| `ETag` | `string` | no |

## UpdateKeyValueStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Comment` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyValueStore` | `KeyValueStore` | no |
| `ETag` | `string` | no |

## UpdateOriginAccessControl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControlConfig` | `OriginAccessControlConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginAccessControl` | `OriginAccessControl` | no |
| `ETag` | `string` | no |

## UpdateOriginRequestPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicyConfig` | `OriginRequestPolicyConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginRequestPolicy` | `OriginRequestPolicy` | no |
| `ETag` | `string` | no |

## UpdatePublicKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKeyConfig` | `PublicKeyConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicKey` | `PublicKey` | no |
| `ETag` | `string` | no |

## UpdateRealtimeLogConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndPoints` | `List<EndPoint>` | no |
| `Fields` | `List<string>` | no |
| `Name` | `string` | no |
| `ARN` | `string` | no |
| `SamplingRate` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RealtimeLogConfig` | `RealtimeLogConfig` | no |

## UpdateResponseHeadersPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicyConfig` | `ResponseHeadersPolicyConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResponseHeadersPolicy` | `ResponseHeadersPolicy` | no |
| `ETag` | `string` | no |

## UpdateStreamingDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistributionConfig` | `StreamingDistributionConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingDistribution` | `StreamingDistribution` | no |
| `ETag` | `string` | no |

## UpdateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `CaCertificatesBundleSource` | `CaCertificatesBundleSource` | no |
| `UseClientCertificateOCSPEndpoint` | `boolean` | no |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStore` | `TrustStore` | no |
| `ETag` | `string` | no |

## UpdateVpcOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOriginEndpointConfig` | `VpcOriginEndpointConfig` | yes |
| `Id` | `string` | yes |
| `IfMatch` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcOrigin` | `VpcOrigin` | no |
| `ETag` | `string` | no |

## VerifyDnsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | no |
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DnsConfigurationList` | `List<DnsConfiguration>` | no |

