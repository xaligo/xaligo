# Elastic Load Balancing

API version: 2012-06-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/elb/2012-06-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerNames` | `List<string>` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApplySecurityGroupsToLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `SecurityGroups` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityGroups` | `List<string>` | no |

## AttachLoadBalancerToSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Subnets` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subnets` | `List<string>` | no |

## ConfigureHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `HealthCheck` | `HealthCheck` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HealthCheck` | `HealthCheck` | no |

## CreateAppCookieStickinessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `CookieName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLBCookieStickinessPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `CookieExpirationPeriod` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Listeners` | `List<Listener>` | yes |
| `AvailabilityZones` | `List<string>` | no |
| `Subnets` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `Scheme` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DNSName` | `string` | no |

## CreateLoadBalancerListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Listeners` | `List<Listener>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateLoadBalancerPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `PolicyName` | `string` | yes |
| `PolicyTypeName` | `string` | yes |
| `PolicyAttributes` | `List<PolicyAttribute>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoadBalancerListeners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `LoadBalancerPorts` | `List<integer>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoadBalancerPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterInstancesFromLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Instances` | `List<Instance>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<Instance>` | no |

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

## DescribeInstanceHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Instances` | `List<Instance>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceStates` | `List<InstanceState>` | no |

## DescribeLoadBalancerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerAttributes` | `LoadBalancerAttributes` | no |

## DescribeLoadBalancerPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | no |
| `PolicyNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDescriptions` | `List<PolicyDescription>` | no |

## DescribeLoadBalancerPolicyTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyTypeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyTypeDescriptions` | `List<PolicyTypeDescription>` | no |

## DescribeLoadBalancers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerNames` | `List<string>` | no |
| `Marker` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerDescriptions` | `List<LoadBalancerDescription>` | no |
| `NextMarker` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagDescriptions` | `List<TagDescription>` | no |

## DetachLoadBalancerFromSubnets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Subnets` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subnets` | `List<string>` | no |

## DisableAvailabilityZonesForLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `AvailabilityZones` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |

## EnableAvailabilityZonesForLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `AvailabilityZones` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AvailabilityZones` | `List<string>` | no |

## ModifyLoadBalancerAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `LoadBalancerAttributes` | `LoadBalancerAttributes` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | no |
| `LoadBalancerAttributes` | `LoadBalancerAttributes` | no |

## RegisterInstancesWithLoadBalancer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `Instances` | `List<Instance>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<Instance>` | no |

## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerNames` | `List<string>` | yes |
| `Tags` | `List<TagKeyOnly>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetLoadBalancerListenerSSLCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `LoadBalancerPort` | `integer` | yes |
| `SSLCertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetLoadBalancerPoliciesForBackendServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `InstancePort` | `integer` | yes |
| `PolicyNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetLoadBalancerPoliciesOfListener

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoadBalancerName` | `string` | yes |
| `LoadBalancerPort` | `integer` | yes |
| `PolicyNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


