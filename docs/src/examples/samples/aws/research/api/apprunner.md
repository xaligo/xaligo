# AWS App Runner

API version: 2020-05-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/apprunner/2020-05-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateCustomDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `DomainName` | `string` | yes |
| `EnableWWWSubdomain` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DNSTarget` | `string` | yes |
| `ServiceArn` | `string` | yes |
| `CustomDomain` | `CustomDomain` | yes |
| `VpcDNSTargets` | `List<VpcDNSTarget>` | yes |

## CreateAutoScalingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationName` | `string` | yes |
| `MaxConcurrency` | `integer` | no |
| `MinSize` | `integer` | no |
| `MaxSize` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfiguration` | `AutoScalingConfiguration` | yes |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | yes |
| `ProviderType` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | yes |

## CreateObservabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfigurationName` | `string` | yes |
| `TraceConfiguration` | `TraceConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfiguration` | `ObservabilityConfiguration` | yes |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceName` | `string` | yes |
| `SourceConfiguration` | `SourceConfiguration` | yes |
| `InstanceConfiguration` | `InstanceConfiguration` | no |
| `Tags` | `List<Tag>` | no |
| `EncryptionConfiguration` | `EncryptionConfiguration` | no |
| `HealthCheckConfiguration` | `HealthCheckConfiguration` | no |
| `AutoScalingConfigurationArn` | `string` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `ObservabilityConfiguration` | `ServiceObservabilityConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `OperationId` | `string` | yes |

## CreateVpcConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectorName` | `string` | yes |
| `Subnets` | `List<string>` | yes |
| `SecurityGroups` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnector` | `VpcConnector` | yes |

## CreateVpcIngressConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `VpcIngressConnectionName` | `string` | yes |
| `IngressVpcConfiguration` | `IngressVpcConfiguration` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnection` | `VpcIngressConnection` | yes |

## DeleteAutoScalingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationArn` | `string` | yes |
| `DeleteAllRevisions` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfiguration` | `AutoScalingConfiguration` | yes |

## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `Connection` | no |

## DeleteObservabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfiguration` | `ObservabilityConfiguration` | yes |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `OperationId` | `string` | yes |

## DeleteVpcConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnector` | `VpcConnector` | yes |

## DeleteVpcIngressConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnection` | `VpcIngressConnection` | yes |

## DescribeAutoScalingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfiguration` | `AutoScalingConfiguration` | yes |

## DescribeCustomDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DNSTarget` | `string` | yes |
| `ServiceArn` | `string` | yes |
| `CustomDomains` | `List<CustomDomain>` | yes |
| `VpcDNSTargets` | `List<VpcDNSTarget>` | yes |
| `NextToken` | `string` | no |

## DescribeObservabilityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfiguration` | `ObservabilityConfiguration` | yes |

## DescribeService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |

## DescribeVpcConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnector` | `VpcConnector` | yes |

## DescribeVpcIngressConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnectionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnection` | `VpcIngressConnection` | yes |

## DisassociateCustomDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DNSTarget` | `string` | yes |
| `ServiceArn` | `string` | yes |
| `CustomDomain` | `CustomDomain` | yes |
| `VpcDNSTargets` | `List<VpcDNSTarget>` | yes |

## ListAutoScalingConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationName` | `string` | no |
| `LatestOnly` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationSummaryList` | `List<AutoScalingConfigurationSummary>` | yes |
| `NextToken` | `string` | no |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionSummaryList` | `List<ConnectionSummary>` | yes |
| `NextToken` | `string` | no |

## ListObservabilityConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfigurationName` | `string` | no |
| `LatestOnly` | `boolean` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObservabilityConfigurationSummaryList` | `List<ObservabilityConfigurationSummary>` | yes |
| `NextToken` | `string` | no |

## ListOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationSummaryList` | `List<OperationSummary>` | no |
| `NextToken` | `string` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSummaryList` | `List<ServiceSummary>` | yes |
| `NextToken` | `string` | no |

## ListServicesForAutoScalingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArnList` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListVpcConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcConnectors` | `List<VpcConnector>` | yes |
| `NextToken` | `string` | no |

## ListVpcIngressConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ListVpcIngressConnectionsFilter` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnectionSummaryList` | `List<VpcIngressConnectionSummary>` | yes |
| `NextToken` | `string` | no |

## PauseService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `OperationId` | `string` | no |

## ResumeService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `OperationId` | `string` | no |

## StartDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDefaultAutoScalingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoScalingConfiguration` | `AutoScalingConfiguration` | yes |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceArn` | `string` | yes |
| `SourceConfiguration` | `SourceConfiguration` | no |
| `InstanceConfiguration` | `InstanceConfiguration` | no |
| `AutoScalingConfigurationArn` | `string` | no |
| `HealthCheckConfiguration` | `HealthCheckConfiguration` | no |
| `NetworkConfiguration` | `NetworkConfiguration` | no |
| `ObservabilityConfiguration` | `ServiceObservabilityConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Service` | `Service` | yes |
| `OperationId` | `string` | yes |

## UpdateVpcIngressConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnectionArn` | `string` | yes |
| `IngressVpcConfiguration` | `IngressVpcConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcIngressConnection` | `VpcIngressConnection` | yes |

