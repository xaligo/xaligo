# Amazon OpenSearch Service

API version: 2021-01-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/opensearch/2021-01-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptInboundConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `InboundConnection` | no |

## AddDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Name` | `string` | yes |
| `DataSourceType` | `DataSourceType` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## AddDirectQueryDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceName` | `string` | yes |
| `DataSourceType` | `DirectQueryDataSourceType` | yes |
| `Description` | `string` | no |
| `OpenSearchArns` | `List<string>` | no |
| `DataSourceAccessPolicy` | `string` | no |
| `TagList` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceArn` | `string` | no |

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
| `PrerequisitePackageIDList` | `List<string>` | no |
| `AssociationConfiguration` | `PackageAssociationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetails` | `DomainPackageDetails` | no |

## AssociatePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageList` | `List<PackageDetailsForAssociation>` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetailsList` | `List<DomainPackageDetails>` | no |

## AttachDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dataSourceArn` | `string` | yes |
| `workspaceId` | `string` | no |
| `workspaceConfiguration` | `WorkspaceConfigurationInput` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentId` | `string` | no |
| `id` | `string` | no |
| `arn` | `string` | no |
| `dataSourceArn` | `string` | no |
| `status` | `string` | no |

## AuthorizeVpcEndpointAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Account` | `string` | no |
| `Service` | `string` | no |
| `ServiceOptions` | `ServiceOptions` | no |

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
| `CancelledChangeIds` | `List<string>` | no |
| `CancelledChangeProperties` | `List<CancelledChangeProperty>` | no |
| `DryRun` | `boolean` | no |

## CancelServiceSoftwareUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSoftwareOptions` | `ServiceSoftwareOptions` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `dataSources` | `List<DataSource>` | no |
| `iamIdentityCenterOptions` | `IamIdentityCenterOptionsInput` | no |
| `appConfigs` | `List<AppConfig>` | no |
| `tagList` | `List<Tag>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `dataSources` | `List<DataSource>` | no |
| `iamIdentityCenterOptions` | `IamIdentityCenterOptions` | no |
| `appConfigs` | `List<AppConfig>` | no |
| `tagList` | `List<Tag>` | no |
| `createdAt` | `timestamp` | no |
| `kmsKeyArn` | `string` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EngineVersion` | `string` | no |
| `ClusterConfig` | `ClusterConfig` | no |
| `EBSOptions` | `EBSOptions` | no |
| `AccessPolicies` | `string` | no |
| `IPAddressType` | `string` | no |
| `SnapshotOptions` | `SnapshotOptions` | no |
| `VPCOptions` | `VPCOptions` | no |
| `CognitoOptions` | `CognitoOptions` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `NodeToNodeEncryptionOptions` | `NodeToNodeEncryptionOptions` | no |
| `AdvancedOptions` | `Map<string>` | no |
| `LogPublishingOptions` | `Map<LogPublishingOption>` | no |
| `DomainEndpointOptions` | `DomainEndpointOptions` | no |
| `AdvancedSecurityOptions` | `AdvancedSecurityOptionsInput` | no |
| `IdentityCenterOptions` | `IdentityCenterOptionsInput` | no |
| `TagList` | `List<Tag>` | no |
| `AutoTuneOptions` | `AutoTuneOptionsInput` | no |
| `OffPeakWindowOptions` | `OffPeakWindowOptions` | no |
| `SoftwareUpdateOptions` | `SoftwareUpdateOptions` | no |
| `AIMLOptions` | `AIMLOptionsInput` | no |
| `DeploymentStrategyOptions` | `DeploymentStrategyOptions` | no |
| `AutomatedSnapshotPauseOptions` | `AutomatedSnapshotPauseRequestOptions` | no |
| `UseCase` | `string` | no |
| `EngineMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `DomainStatus` | no |

## CreateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexName` | `string` | yes |
| `IndexSchema` | `IndexSchema` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## CreateOutboundConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalDomainInfo` | `DomainInformationContainer` | yes |
| `RemoteDomainInfo` | `DomainInformationContainer` | yes |
| `ConnectionAlias` | `string` | yes |
| `ConnectionMode` | `string` | no |
| `ConnectionProperties` | `ConnectionProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LocalDomainInfo` | `DomainInformationContainer` | no |
| `RemoteDomainInfo` | `DomainInformationContainer` | no |
| `ConnectionAlias` | `string` | no |
| `ConnectionStatus` | `OutboundConnectionStatus` | no |
| `ConnectionId` | `string` | no |
| `ConnectionMode` | `string` | no |
| `ConnectionProperties` | `ConnectionProperties` | no |

## CreatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageName` | `string` | yes |
| `PackageType` | `string` | yes |
| `PackageDescription` | `string` | no |
| `PackageSource` | `PackageSource` | yes |
| `PackageConfiguration` | `PackageConfiguration` | no |
| `EngineVersion` | `string` | no |
| `PackageVendingOptions` | `PackageVendingOptions` | no |
| `PackageEncryptionOptions` | `PackageEncryptionOptions` | no |

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

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteDirectQueryDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `DomainStatus` | no |

## DeleteInboundConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `InboundConnection` | no |

## DeleteIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## DeleteOutboundConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `OutboundConnection` | no |

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

## DeregisterCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `capabilityName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## DescribeDataSourceAttachment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dataSourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachmentId` | `string` | no |
| `id` | `string` | no |
| `arn` | `string` | no |
| `dataSourceArn` | `string` | no |
| `status` | `string` | no |

## DescribeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatus` | `DomainStatus` | yes |

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

## DescribeDomainConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainConfig` | `DomainConfig` | yes |

## DescribeDomainHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainState` | `string` | no |
| `AvailabilityZoneCount` | `string` | no |
| `ActiveAvailabilityZoneCount` | `string` | no |
| `StandByAvailabilityZoneCount` | `string` | no |
| `DataNodeCount` | `string` | no |
| `DedicatedMaster` | `boolean` | no |
| `MasterEligibleNodeCount` | `string` | no |
| `WarmNodeCount` | `string` | no |
| `MasterNode` | `string` | no |
| `ClusterHealth` | `string` | no |
| `TotalShards` | `string` | no |
| `TotalUnAssignedShards` | `string` | no |
| `EnvironmentInformation` | `List<EnvironmentInfo>` | no |

## DescribeDomainNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNodesStatusList` | `List<DomainNodesStatus>` | no |

## DescribeDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainStatusList` | `List<DomainStatus>` | yes |

## DescribeDryRunProgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DryRunId` | `string` | no |
| `LoadDryRunConfig` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DryRunProgressStatus` | `DryRunProgressStatus` | no |
| `DryRunConfig` | `DomainStatus` | no |
| `DryRunResults` | `DryRunResults` | no |

## DescribeInboundConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<InboundConnection>` | no |
| `NextToken` | `string` | no |

## DescribeInsightDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entity` | `InsightEntity` | yes |
| `InsightId` | `string` | yes |
| `ShowHtmlContent` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fields` | `List<InsightField>` | yes |

## DescribeInstanceTypeLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `InstanceType` | `string` | yes |
| `EngineVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LimitsByRole` | `Map<Limits>` | no |

## DescribeOutboundConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connections` | `List<OutboundConnection>` | no |
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

## DescribeReservedInstanceOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstanceOfferingId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedInstanceOfferings` | `List<ReservedInstanceOffering>` | no |

## DescribeReservedInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ReservedInstances` | `List<ReservedInstance>` | no |

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

## DetachDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dataSourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `dataSourceArn` | `string` | no |

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

## DissociatePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageList` | `List<string>` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainPackageDetailsList` | `List<DomainPackageDetails>` | no |

## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `arn` | `string` | no |
| `name` | `string` | no |
| `endpoint` | `string` | no |
| `status` | `string` | no |
| `iamIdentityCenterOptions` | `IamIdentityCenterOptions` | no |
| `dataSources` | `List<DataSource>` | no |
| `appConfigs` | `List<AppConfig>` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `kmsKeyArn` | `string` | no |

## GetCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `capabilityName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityName` | `string` | no |
| `applicationId` | `string` | no |
| `status` | `string` | no |
| `capabilityConfig` | `CapabilityExtendedResponseConfig` | no |
| `failures` | `List<CapabilityFailure>` | no |

## GetCompatibleVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompatibleVersions` | `List<CompatibleVersionsMap>` | no |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceType` | `DataSourceType` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |

## GetDefaultApplicationSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |

## GetDirectQueryDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceName` | `string` | no |
| `DataSourceType` | `DirectQueryDataSourceType` | no |
| `Description` | `string` | no |
| `OpenSearchArns` | `List<string>` | no |
| `DataSourceAccessPolicy` | `string` | no |
| `DataSourceArn` | `string` | no |

## GetDomainMaintenanceStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaintenanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `StatusMessage` | `string` | no |
| `NodeId` | `string` | no |
| `Action` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |

## GetIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IndexSchema` | `IndexSchema` | yes |

## GetMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationId` | `string` | no |
| `status` | `string` | no |
| `applicationId` | `string` | no |
| `source` | `MigrationSource` | no |
| `exportedCount` | `integer` | no |
| `importedCount` | `integer` | no |
| `error` | `MigrationError` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

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

## InsightFeedback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entity` | `InsightFeedbackEntity` | yes |
| `InsightId` | `string` | yes |
| `Thumbs` | `string` | yes |
| `FeedbackText` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `statuses` | `List<string>` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSummaries` | `List<ApplicationSummary>` | no |
| `nextToken` | `string` | no |

## ListDataSourceAttachments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attachments` | `List<DataSourceAttachmentSummary>` | no |
| `nextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSources` | `List<DataSourceDetails>` | no |

## ListDirectQueryDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DirectQueryDataSources` | `List<DirectQueryDataSource>` | no |

## ListDomainMaintenances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Action` | `string` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainMaintenances` | `List<DomainMaintenanceDetails>` | no |
| `NextToken` | `string` | no |

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

## ListInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entity` | `InsightEntity` | yes |
| `TimeRange` | `InsightTimeRange` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Insights` | `List<Insight>` | no |
| `NextToken` | `string` | no |

## ListInstanceTypeDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineVersion` | `string` | yes |
| `DomainName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `RetrieveAZs` | `boolean` | no |
| `InstanceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceTypeDetails` | `List<InstanceTypeDetails>` | no |
| `NextToken` | `string` | no |

## ListMigrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrations` | `List<MigrationSummary>` | no |
| `nextToken` | `string` | no |

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

## ListScheduledActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledActions` | `List<ScheduledAction>` | no |
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

## ListVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Versions` | `List<string>` | no |
| `NextToken` | `string` | no |

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

## PurchaseReservedInstanceOffering

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstanceOfferingId` | `string` | yes |
| `ReservationName` | `string` | yes |
| `InstanceCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedInstanceId` | `string` | no |
| `ReservationName` | `string` | no |

## PutDefaultApplicationSetting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | yes |
| `setAsDefault` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationArn` | `string` | no |

## RegisterCapability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `capabilityName` | `string` | yes |
| `capabilityConfig` | `CapabilityBaseRequestConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capabilityName` | `string` | no |
| `applicationId` | `string` | no |
| `status` | `string` | no |
| `capabilityConfig` | `CapabilityBaseResponseConfig` | no |

## RejectInboundConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connection` | `InboundConnection` | no |

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
| `Account` | `string` | no |
| `Service` | `string` | no |
| `ServiceOptions` | `ServiceOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RollbackServiceSoftwareUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RollbackServiceSoftwareOptions` | `RollbackServiceSoftwareOptions` | no |

## StartDomainMaintenance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Action` | `string` | yes |
| `NodeId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaintenanceId` | `string` | no |

## StartMigration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `applicationId` | `string` | yes |
| `migrationOptions` | `MigrationOptions` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `migrationId` | `string` | no |
| `status` | `string` | no |

## StartServiceSoftwareUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ScheduleAt` | `string` | no |
| `DesiredStartTime` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSoftwareOptions` | `ServiceSoftwareOptions` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `dataSources` | `List<DataSource>` | no |
| `appConfigs` | `List<AppConfig>` | no |
| `iamIdentityCenterOptions` | `IamIdentityCenterOptionsInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `dataSources` | `List<DataSource>` | no |
| `iamIdentityCenterOptions` | `IamIdentityCenterOptions` | no |
| `appConfigs` | `List<AppConfig>` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Name` | `string` | yes |
| `DataSourceType` | `DataSourceType` | yes |
| `Description` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## UpdateDirectQueryDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceName` | `string` | yes |
| `DataSourceType` | `DirectQueryDataSourceType` | yes |
| `Description` | `string` | no |
| `OpenSearchArns` | `List<string>` | no |
| `DataSourceAccessPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSourceArn` | `string` | no |

## UpdateDomainConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ClusterConfig` | `ClusterConfig` | no |
| `EBSOptions` | `EBSOptions` | no |
| `SnapshotOptions` | `SnapshotOptions` | no |
| `VPCOptions` | `VPCOptions` | no |
| `CognitoOptions` | `CognitoOptions` | no |
| `AdvancedOptions` | `Map<string>` | no |
| `AccessPolicies` | `string` | no |
| `IPAddressType` | `string` | no |
| `LogPublishingOptions` | `Map<LogPublishingOption>` | no |
| `EncryptionAtRestOptions` | `EncryptionAtRestOptions` | no |
| `DomainEndpointOptions` | `DomainEndpointOptions` | no |
| `NodeToNodeEncryptionOptions` | `NodeToNodeEncryptionOptions` | no |
| `AdvancedSecurityOptions` | `AdvancedSecurityOptionsInput` | no |
| `IdentityCenterOptions` | `IdentityCenterOptionsInput` | no |
| `AutoTuneOptions` | `AutoTuneOptions` | no |
| `DryRun` | `boolean` | no |
| `DryRunMode` | `string` | no |
| `OffPeakWindowOptions` | `OffPeakWindowOptions` | no |
| `SoftwareUpdateOptions` | `SoftwareUpdateOptions` | no |
| `AIMLOptions` | `AIMLOptionsInput` | no |
| `DeploymentStrategyOptions` | `DeploymentStrategyOptions` | no |
| `AutomatedSnapshotPauseOptions` | `AutomatedSnapshotPauseRequestOptions` | no |
| `UseCase` | `string` | no |
| `EngineMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainConfig` | `DomainConfig` | yes |
| `DryRunResults` | `DryRunResults` | no |
| `DryRunProgressStatus` | `DryRunProgressStatus` | no |

## UpdateIndex

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `IndexName` | `string` | yes |
| `IndexSchema` | `IndexSchema` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |

## UpdatePackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `PackageSource` | `PackageSource` | yes |
| `PackageDescription` | `string` | no |
| `CommitMessage` | `string` | no |
| `PackageConfiguration` | `PackageConfiguration` | no |
| `PackageEncryptionOptions` | `PackageEncryptionOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageDetails` | `PackageDetails` | no |

## UpdatePackageScope

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | yes |
| `Operation` | `string` | yes |
| `PackageUserList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PackageID` | `string` | no |
| `Operation` | `string` | no |
| `PackageUserList` | `List<string>` | no |

## UpdateScheduledAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ActionID` | `string` | yes |
| `ActionType` | `string` | yes |
| `ScheduleAt` | `string` | yes |
| `DesiredStartTime` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ScheduledAction` | `ScheduledAction` | no |

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

## UpgradeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `TargetVersion` | `string` | yes |
| `PerformCheckOnly` | `boolean` | no |
| `AdvancedOptions` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UpgradeId` | `string` | no |
| `DomainName` | `string` | no |
| `TargetVersion` | `string` | no |
| `PerformCheckOnly` | `boolean` | no |
| `AdvancedOptions` | `Map<string>` | no |
| `ChangeProgressDetails` | `ChangeProgressDetails` | no |

