# Amazon Elasticsearch Service

API version: 2015-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/es/2015-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptInboundCrossClusterSearchConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnection` | `InboundCrossClusterSearchConnection` | no |

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | yes |
| `TagList` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetails` | `DomainPackageDetails` | no |

## AuthorizeVpcEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Account` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedPrincipal` | `AuthorizedPrincipal` | yes |

## CancelDomainConfigChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRun` | `boolean` | no |
| `CancelledChangeIds` | `List<string>` | no |
| `CancelledChangeProperties` | `List<CancelledChangeProperty>` | no |

## CancelElasticsearchServiceSoftwareUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSoftwareOptions` | `ServiceSoftwareOptions` | no |

## CreateElasticsearchDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ElasticsearchVersion` | `string` | no |
| `ElasticsearchClusterConfig` | `ElasticsearchClusterConfig` | no |
| `EBSOptions` | `EBSOptions` | no |
| `AccessPolicies` | `string` | no |
| `SnapshotOptions` | `SnapshotOptions` | no |
| `VPCOptions` | `VPCOptions` | no |
| `CognitoOptions` | `CognitoOptions` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `NodeToNodeEncryptionOptions` | `NodeToNodeEncryptionOptions` | no |
| `AdvancedOptions` | `Map<string>` | no |
| `LogPublishingOptions` | `Map<LogPublishingOption>` | no |
| `DomainEndpointOptions` | `DomainEndpointOptions` | no |
| `AdvancedSecurityOptions` | `AdvancedSecurityOptionsInput` | no |
| `AutoTuneOptions` | `AutoTuneOptionsInput` | no |
| `TagList` | `List<Tag>` | no |
| `DeploymentStrategyOptions` | `DeploymentStrategyOptions` | no |
| `AutomatedSnapshotPauseOptions` | `AutomatedSnapshotPauseRequestOptions` | no |
| `UseCase` | `string` | no |
| `EngineMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `ElasticsearchDomainStatus` | no |

## CreateOutboundCrossClusterSearchConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDomainInfo` | `DomainInformation` | yes |
| `DestinationDomainInfo` | `DomainInformation` | yes |
| `ConnectionAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceDomainInfo` | `DomainInformation` | no |
| `DestinationDomainInfo` | `DomainInformation` | no |
| `ConnectionAlias` | `string` | no |
| `ConnectionStatus` | `OutboundCrossClusterSearchConnectionStatus` | no |
| `CrossClusterSearchConnectionId` | `string` | no |

## CreatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageName` | `string` | yes |
| `PackageType` | `string` | yes |
| `PackageDescription` | `string` | no |
| `PackageSource` | `PackageSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageDetails` | `PackageDetails` | no |

## CreateVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainArn` | `string` | yes |
| `VpcOptions` | `VPCOptions` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpoint` | `VpcEndpoint` | yes |

## DeleteElasticsearchDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `ElasticsearchDomainStatus` | no |

## DeleteElasticsearchServiceRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInboundCrossClusterSearchConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnection` | `InboundCrossClusterSearchConnection` | no |

## DeleteOutboundCrossClusterSearchConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnection` | `OutboundCrossClusterSearchConnection` | no |

## DeletePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageDetails` | `PackageDetails` | no |

## DeleteVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointSummary` | `VpcEndpointSummary` | yes |

## DescribeDomainAutoTunes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoTunes` | `List<AutoTune>` | no |
| `NextToken` | `string` | no |

## DescribeDomainChangeProgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ChangeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeProgressStatus` | `ChangeProgressStatusDetails` | no |

## DescribeElasticsearchDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `ElasticsearchDomainStatus` | yes |

## DescribeElasticsearchDomainConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainConfig` | `ElasticsearchDomainConfig` | yes |

## DescribeElasticsearchDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatusList` | `List<ElasticsearchDomainStatus>` | yes |

## DescribeElasticsearchInstanceTypeLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `InstanceType` | `string` | yes |
| `ElasticsearchVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LimitsByRole` | `Map<Limits>` | no |

## DescribeInboundCrossClusterSearchConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnections` | `List<InboundCrossClusterSearchConnection>` | no |
| `NextToken` | `string` | no |

## DescribeOutboundCrossClusterSearchConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnections` | `List<OutboundCrossClusterSearchConnection>` | no |
| `NextToken` | `string` | no |

## DescribePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<DescribePackagesFilter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageDetailsList` | `List<PackageDetails>` | no |
| `NextToken` | `string` | no |

## DescribeReservedElasticsearchInstanceOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedElasticsearchInstanceOfferingId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedElasticsearchInstanceOfferings` | `List<ReservedElasticsearchInstanceOffering>` | no |

## DescribeReservedElasticsearchInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedElasticsearchInstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedElasticsearchInstances` | `List<ReservedElasticsearchInstance>` | no |

## DescribeVpcEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpoints` | `List<VpcEndpoint>` | yes |
| `VpcEndpointErrors` | `List<VpcEndpointError>` | yes |

## DissociatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetails` | `DomainPackageDetails` | no |

## GetCompatibleElasticsearchVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleElasticsearchVersions` | `List<CompatibleVersionsMap>` | no |

## GetPackageVersionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | no |
| `PackageVersionHistoryList` | `List<PackageVersionHistory>` | no |
| `NextToken` | `string` | no |

## GetUpgradeHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpgradeHistories` | `List<UpgradeHistory>` | no |
| `NextToken` | `string` | no |

## GetUpgradeStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpgradeStep` | `string` | no |
| `StepStatus` | `string` | no |
| `UpgradeName` | `string` | no |

## ListDomainNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `List<DomainInfo>` | no |

## ListDomainsForPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetailsList` | `List<DomainPackageDetails>` | no |
| `NextToken` | `string` | no |

## ListElasticsearchInstanceTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ElasticsearchVersion` | `string` | yes |
| `DomainName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ElasticsearchInstanceTypes` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListElasticsearchVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ElasticsearchVersions` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListPackagesForDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetailsList` | `List<DomainPackageDetails>` | no |
| `NextToken` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## ListVpcEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedPrincipalList` | `List<AuthorizedPrincipal>` | yes |
| `NextToken` | `string` | yes |

## ListVpcEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointSummaryList` | `List<VpcEndpointSummary>` | yes |
| `NextToken` | `string` | yes |

## ListVpcEndpointsForDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointSummaryList` | `List<VpcEndpointSummary>` | yes |
| `NextToken` | `string` | yes |

## PurchaseReservedElasticsearchInstanceOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedElasticsearchInstanceOfferingId` | `string` | yes |
| `ReservationName` | `string` | yes |
| `InstanceCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedElasticsearchInstanceId` | `string` | no |
| `ReservationName` | `string` | no |

## RejectInboundCrossClusterSearchConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CrossClusterSearchConnection` | `InboundCrossClusterSearchConnection` | no |

## RemoveTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeVpcEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Account` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartElasticsearchServiceSoftwareUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSoftwareOptions` | `ServiceSoftwareOptions` | no |

## UpdateElasticsearchDomainConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ElasticsearchClusterConfig` | `ElasticsearchClusterConfig` | no |
| `EBSOptions` | `EBSOptions` | no |
| `SnapshotOptions` | `SnapshotOptions` | no |
| `VPCOptions` | `VPCOptions` | no |
| `CognitoOptions` | `CognitoOptions` | no |
| `AdvancedOptions` | `Map<string>` | no |
| `AccessPolicies` | `string` | no |
| `LogPublishingOptions` | `Map<LogPublishingOption>` | no |
| `DomainEndpointOptions` | `DomainEndpointOptions` | no |
| `AdvancedSecurityOptions` | `AdvancedSecurityOptionsInput` | no |
| `NodeToNodeEncryptionOptions` | `NodeToNodeEncryptionOptions` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `AutoTuneOptions` | `AutoTuneOptions` | no |
| `DryRun` | `boolean` | no |
| `DeploymentStrategyOptions` | `DeploymentStrategyOptions` | no |
| `AutomatedSnapshotPauseOptions` | `AutomatedSnapshotPauseRequestOptions` | no |
| `UseCase` | `string` | no |
| `EngineMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainConfig` | `ElasticsearchDomainConfig` | yes |
| `DryRunResults` | `DryRunResults` | no |

## UpdatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `PackageSource` | `PackageSource` | yes |
| `PackageDescription` | `string` | no |
| `CommitMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageDetails` | `PackageDetails` | no |

## UpdateVpcEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpointId` | `string` | yes |
| `VpcOptions` | `VPCOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcEndpoint` | `VpcEndpoint` | yes |

## UpgradeElasticsearchDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `TargetVersion` | `string` | yes |
| `PerformCheckOnly` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `TargetVersion` | `string` | no |
| `PerformCheckOnly` | `boolean` | no |
| `ChangeProgressDetails` | `ChangeProgressDetails` | no |

