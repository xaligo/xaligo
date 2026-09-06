# Elastic Load Balancing

API version: 2015-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/elbv2/2015-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddListenerCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Certificates` | `List<Certificate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificates` | `List<Certificate>` | no |

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddTrustStoreRevocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `RevocationContents` | `List<RevocationContent>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreRevocations` | `List<TrustStoreRevocation>` | no |

## CreateListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `Protocol` | `string` | no |
| `Port` | `integer` | no |
| `SslPolicy` | `string` | no |
| `Certificates` | `List<Certificate>` | no |
| `DefaultActions` | `List<Action>` | yes |
| `AlpnPolicy` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `MutualAuthentication` | `MutualAuthenticationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listeners` | `List<Listener>` | no |

## CreateLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Subnets` | `List<string>` | no |
| `SubnetMappings` | `List<SubnetMapping>` | no |
| `SecurityGroups` | `List<string>` | no |
| `Scheme` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `Type` | `string` | no |
| `IpAddressType` | `string` | no |
| `CustomerOwnedIpv4Pool` | `string` | no |
| `EnablePrefixForIpv6SourceNat` | `string` | no |
| `IpamPools` | `IpamPools` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancers` | `List<LoadBalancer>` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Conditions` | `List<RuleCondition>` | yes |
| `Priority` | `integer` | yes |
| `Actions` | `List<Action>` | yes |
| `Tags` | `List<Tag>` | no |
| `Transforms` | `List<RuleTransform>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<Rule>` | no |

## CreateTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Protocol` | `string` | no |
| `ProtocolVersion` | `string` | no |
| `Port` | `integer` | no |
| `VpcId` | `string` | no |
| `HealthCheckProtocol` | `string` | no |
| `HealthCheckPort` | `string` | no |
| `HealthCheckEnabled` | `boolean` | no |
| `HealthCheckPath` | `string` | no |
| `HealthCheckIntervalSeconds` | `integer` | no |
| `HealthCheckTimeoutSeconds` | `integer` | no |
| `HealthyThresholdCount` | `integer` | no |
| `UnhealthyThresholdCount` | `integer` | no |
| `Matcher` | `Matcher` | no |
| `TargetType` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `IpAddressType` | `string` | no |
| `TargetControlPort` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroups` | `List<TargetGroup>` | no |

## CreateTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `CaCertificatesBundleS3Bucket` | `string` | yes |
| `CaCertificatesBundleS3Key` | `string` | yes |
| `CaCertificatesBundleS3ObjectVersion` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStores` | `List<TrustStore>` | no |

## DeleteListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSharedTrustStoreAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |
| `Targets` | `List<TargetDescription>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Limits` | `List<Limit>` | no |
| `NextMarker` | `string` | no |

## DescribeCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModifiedTime` | `timestamp` | no |
| `DecreaseRequestsRemaining` | `integer` | no |
| `MinimumLoadBalancerCapacity` | `MinimumLoadBalancerCapacity` | no |
| `CapacityReservationState` | `List<ZonalCapacityReservationState>` | no |

## DescribeListenerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<ListenerAttribute>` | no |

## DescribeListenerCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificates` | `List<Certificate>` | no |
| `NextMarker` | `string` | no |

## DescribeListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | no |
| `ListenerArns` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listeners` | `List<Listener>` | no |
| `NextMarker` | `string` | no |

## DescribeLoadBalancerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<LoadBalancerAttribute>` | no |

## DescribeLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArns` | `List<string>` | no |
| `Names` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancers` | `List<LoadBalancer>` | no |
| `NextMarker` | `string` | no |

## DescribeRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | no |
| `RuleArns` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<Rule>` | no |
| `NextMarker` | `string` | no |

## DescribeSSLPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |
| `LoadBalancerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SslPolicies` | `List<SslPolicy>` | no |
| `NextMarker` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagDescriptions` | `List<TagDescription>` | no |

## DescribeTargetGroupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<TargetGroupAttribute>` | no |

## DescribeTargetGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | no |
| `TargetGroupArns` | `List<string>` | no |
| `Names` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroups` | `List<TargetGroup>` | no |
| `NextMarker` | `string` | no |

## DescribeTargetHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |
| `Targets` | `List<TargetDescription>` | no |
| `Include` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetHealthDescriptions` | `List<TargetHealthDescription>` | no |

## DescribeTrustStoreAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreAssociations` | `List<TrustStoreAssociation>` | no |
| `NextMarker` | `string` | no |

## DescribeTrustStoreRevocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `RevocationIds` | `List<long>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreRevocations` | `List<DescribeTrustStoreRevocation>` | no |
| `NextMarker` | `string` | no |

## DescribeTrustStores

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArns` | `List<string>` | no |
| `Names` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStores` | `List<TrustStore>` | no |
| `NextMarker` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetTrustStoreCaCertificatesBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `string` | no |

## GetTrustStoreRevocationContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `RevocationId` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Location` | `string` | no |

## ModifyCapacityReservation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `MinimumLoadBalancerCapacity` | `MinimumLoadBalancerCapacity` | no |
| `ResetCapacityReservation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastModifiedTime` | `timestamp` | no |
| `DecreaseRequestsRemaining` | `integer` | no |
| `MinimumLoadBalancerCapacity` | `MinimumLoadBalancerCapacity` | no |
| `CapacityReservationState` | `List<ZonalCapacityReservationState>` | no |

## ModifyIpPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `IpamPools` | `IpamPools` | no |
| `RemoveIpamPools` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpamPools` | `IpamPools` | no |

## ModifyListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Port` | `integer` | no |
| `Protocol` | `string` | no |
| `SslPolicy` | `string` | no |
| `Certificates` | `List<Certificate>` | no |
| `DefaultActions` | `List<Action>` | no |
| `AlpnPolicy` | `List<string>` | no |
| `MutualAuthentication` | `MutualAuthenticationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Listeners` | `List<Listener>` | no |

## ModifyListenerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Attributes` | `List<ListenerAttribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<ListenerAttribute>` | no |

## ModifyLoadBalancerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `Attributes` | `List<LoadBalancerAttribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<LoadBalancerAttribute>` | no |

## ModifyRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | yes |
| `Conditions` | `List<RuleCondition>` | no |
| `Actions` | `List<Action>` | no |
| `Transforms` | `List<RuleTransform>` | no |
| `ResetTransforms` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<Rule>` | no |

## ModifyTargetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |
| `HealthCheckProtocol` | `string` | no |
| `HealthCheckPort` | `string` | no |
| `HealthCheckPath` | `string` | no |
| `HealthCheckEnabled` | `boolean` | no |
| `HealthCheckIntervalSeconds` | `integer` | no |
| `HealthCheckTimeoutSeconds` | `integer` | no |
| `HealthyThresholdCount` | `integer` | no |
| `UnhealthyThresholdCount` | `integer` | no |
| `Matcher` | `Matcher` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroups` | `List<TargetGroup>` | no |

## ModifyTargetGroupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |
| `Attributes` | `List<TargetGroupAttribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<TargetGroupAttribute>` | no |

## ModifyTrustStore

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `CaCertificatesBundleS3Bucket` | `string` | yes |
| `CaCertificatesBundleS3Key` | `string` | yes |
| `CaCertificatesBundleS3ObjectVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStores` | `List<TrustStore>` | no |

## RegisterTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetGroupArn` | `string` | yes |
| `Targets` | `List<TargetDescription>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveListenerCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ListenerArn` | `string` | yes |
| `Certificates` | `List<Certificate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArns` | `List<string>` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemoveTrustStoreRevocations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrustStoreArn` | `string` | yes |
| `RevocationIds` | `List<long>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIpAddressType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `IpAddressType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IpAddressType` | `string` | no |

## SetRulePriorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RulePriorities` | `List<RulePriorityPair>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<Rule>` | no |

## SetSecurityGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `SecurityGroups` | `List<string>` | yes |
| `EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroupIds` | `List<string>` | no |
| `EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic` | `string` | no |

## SetSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerArn` | `string` | yes |
| `Subnets` | `List<string>` | no |
| `SubnetMappings` | `List<SubnetMapping>` | no |
| `IpAddressType` | `string` | no |
| `EnablePrefixForIpv6SourceNat` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<AvailabilityZone>` | no |
| `IpAddressType` | `string` | no |
| `EnablePrefixForIpv6SourceNat` | `string` | no |

