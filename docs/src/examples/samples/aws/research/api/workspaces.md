# Amazon WorkSpaces

API version: 2015-04-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/workspaces/2015-04-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptAccountLinkInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLink` | `AccountLink` | no |

## AssociateConnectionAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionIdentifier` | `string` | no |

## AssociateIpGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `GroupIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWorkspaceApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `WorkspaceResourceAssociation` | no |

## AuthorizeIpRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `UserRules` | `List<IpRuleItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CopyWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `SourceImageId` | `string` | yes |
| `SourceRegion` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## CreateAccountLinkInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetAccountId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLink` | `AccountLink` | no |

## CreateConnectClientAddIn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Name` | `string` | yes |
| `URL` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddInId` | `string` | no |

## CreateConnectionAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionString` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | no |

## CreateIpGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupName` | `string` | yes |
| `GroupDesc` | `string` | no |
| `UserRules` | `List<IpRuleItem>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | no |

## CreateStandbyWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PrimaryRegion` | `string` | yes |
| `StandbyWorkspaces` | `List<StandbyWorkspace>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedStandbyRequests` | `List<FailedCreateStandbyWorkspacesRequest>` | no |
| `PendingStandbyRequests` | `List<PendingCreateStandbyWorkspacesRequest>` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateUpdatedWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `SourceImageId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## CreateWorkspaceBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleName` | `string` | yes |
| `BundleDescription` | `string` | yes |
| `ImageId` | `string` | yes |
| `ComputeType` | `ComputeType` | yes |
| `UserStorage` | `UserStorage` | yes |
| `RootStorage` | `RootStorage` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceBundle` | `WorkspaceBundle` | no |

## CreateWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `OperatingSystem` | `OperatingSystem` | no |
| `State` | `string` | no |
| `RequiredTenancy` | `string` | no |
| `Created` | `timestamp` | no |
| `OwnerAccountId` | `string` | no |

## CreateWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workspaces` | `List<WorkspaceRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedCreateWorkspaceRequest>` | no |
| `PendingRequests` | `List<Workspace>` | no |

## CreateWorkspacesPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `Description` | `string` | yes |
| `BundleId` | `string` | yes |
| `DirectoryId` | `string` | yes |
| `Capacity` | `Capacity` | yes |
| `Tags` | `List<Tag>` | no |
| `ApplicationSettings` | `ApplicationSettingsRequest` | no |
| `TimeoutSettings` | `TimeoutSettings` | no |
| `RunningMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspacesPool` | `WorkspacesPool` | no |

## DeleteAccountLinkInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLink` | `AccountLink` | no |

## DeleteClientBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `Platforms` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectClientAddIn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddInId` | `string` | yes |
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnectionAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIpGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspaceBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeployWorkspaceApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `Force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deployment` | `WorkSpaceApplicationDeployment` | no |

## DeregisterWorkspaceDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedTenancySupport` | `string` | no |
| `DedicatedTenancyManagementCidrRange` | `string` | no |
| `DedicatedTenancyAccountType` | `string` | no |
| `Message` | `string` | no |

## DescribeAccountModifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountModifications` | `List<AccountModification>` | no |
| `NextToken` | `string` | no |

## DescribeApplicationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ApplicationId` | `string` | yes |
| `AssociatedResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<ApplicationResourceAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIds` | `List<string>` | no |
| `ComputeTypeNames` | `List<string>` | no |
| `LicenseType` | `string` | no |
| `OperatingSystemNames` | `List<string>` | no |
| `Owner` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<WorkSpaceApplication>` | no |
| `NextToken` | `string` | no |

## DescribeBundleAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleId` | `string` | yes |
| `AssociatedResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<BundleResourceAssociation>` | no |

## DescribeClientBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceTypeWindows` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeOsx` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeAndroid` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeIos` | `IosClientBrandingAttributes` | no |
| `DeviceTypeLinux` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeWeb` | `DefaultClientBrandingAttributes` | no |

## DescribeClientProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientPropertiesList` | `List<ClientPropertiesResult>` | no |

## DescribeConnectClientAddIns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddIns` | `List<ConnectClientAddIn>` | no |
| `NextToken` | `string` | no |

## DescribeConnectionAliasPermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | no |
| `ConnectionAliasPermissions` | `List<ConnectionAliasPermission>` | no |
| `NextToken` | `string` | no |

## DescribeConnectionAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasIds` | `List<string>` | no |
| `ResourceId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionAliases` | `List<ConnectionAlias>` | no |
| `NextToken` | `string` | no |

## DescribeCustomWorkspaceImageImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `InfrastructureConfigurationArn` | `string` | no |
| `State` | `string` | no |
| `StateMessage` | `string` | no |
| `ProgressPercentage` | `integer` | no |
| `Created` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `ImageSource` | `ImageSourceIdentifier` | no |
| `ImageBuilderInstanceId` | `string` | no |
| `ErrorDetails` | `List<CustomWorkspaceImageImportErrorDetails>` | no |

## DescribeImageAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `AssociatedResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<ImageResourceAssociation>` | no |

## DescribeIpGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Result` | `List<WorkspacesIpGroup>` | no |
| `NextToken` | `string` | no |

## DescribeTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TagList` | `List<Tag>` | no |

## DescribeWorkspaceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `AssociatedResourceTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<WorkspaceResourceAssociation>` | no |

## DescribeWorkspaceBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleIds` | `List<string>` | no |
| `Owner` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Bundles` | `List<WorkspaceBundle>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspaceDirectories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryIds` | `List<string>` | no |
| `WorkspaceDirectoryNames` | `List<string>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<DescribeWorkspaceDirectoriesFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Directories` | `List<WorkspaceDirectory>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspaceImagePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `ImagePermissions` | `List<ImagePermission>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspaceImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageIds` | `List<string>` | no |
| `ImageType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Images` | `List<WorkspaceImage>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspaceSnapshots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RebuildSnapshots` | `List<Snapshot>` | no |
| `RestoreSnapshots` | `List<Snapshot>` | no |

## DescribeWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceIds` | `List<string>` | no |
| `DirectoryId` | `string` | no |
| `UserName` | `string` | no |
| `BundleId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |
| `WorkspaceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workspaces` | `List<Workspace>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspacesConnectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceIds` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspacesConnectionStatus` | `List<WorkspaceConnectionStatus>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspacesPoolSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `UserId` | `string` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sessions` | `List<WorkspacesPoolSession>` | no |
| `NextToken` | `string` | no |

## DescribeWorkspacesPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolIds` | `List<string>` | no |
| `Filters` | `List<DescribeWorkspacesPoolsFilter>` | no |
| `Limit` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspacesPools` | `List<WorkspacesPool>` | no |
| `NextToken` | `string` | no |

## DisassociateConnectionAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateIpGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `GroupIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWorkspaceApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Association` | `WorkspaceResourceAssociation` | no |

## GetAccountLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkId` | `string` | no |
| `LinkedAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLink` | `AccountLink` | no |

## ImportClientBranding

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `DeviceTypeWindows` | `DefaultImportClientBrandingAttributes` | no |
| `DeviceTypeOsx` | `DefaultImportClientBrandingAttributes` | no |
| `DeviceTypeAndroid` | `DefaultImportClientBrandingAttributes` | no |
| `DeviceTypeIos` | `IosImportClientBrandingAttributes` | no |
| `DeviceTypeLinux` | `DefaultImportClientBrandingAttributes` | no |
| `DeviceTypeWeb` | `DefaultImportClientBrandingAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceTypeWindows` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeOsx` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeAndroid` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeIos` | `IosClientBrandingAttributes` | no |
| `DeviceTypeLinux` | `DefaultClientBrandingAttributes` | no |
| `DeviceTypeWeb` | `DefaultClientBrandingAttributes` | no |

## ImportCustomWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `ImageDescription` | `string` | yes |
| `ComputeType` | `string` | yes |
| `Protocol` | `string` | yes |
| `ImageSource` | `ImageSourceIdentifier` | yes |
| `InfrastructureConfigurationArn` | `string` | yes |
| `Platform` | `string` | yes |
| `OsVersion` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |
| `State` | `string` | no |

## ImportWorkspaceImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ec2ImageId` | `string` | yes |
| `IngestionProcess` | `string` | yes |
| `ImageName` | `string` | yes |
| `ImageDescription` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `Applications` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | no |

## ListAccountLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkStatusFilter` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLinks` | `List<AccountLink>` | no |
| `NextToken` | `string` | no |

## ListAvailableManagementCidrRanges

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagementCidrRangeConstraint` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagementCidrRanges` | `List<string>` | no |
| `NextToken` | `string` | no |

## MigrateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceWorkspaceId` | `string` | yes |
| `BundleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceWorkspaceId` | `string` | no |
| `TargetWorkspaceId` | `string` | no |

## ModifyAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedTenancySupport` | `string` | no |
| `DedicatedTenancyManagementCidrRange` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## ModifyCertificateBasedAuthProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `CertificateBasedAuthProperties` | `CertificateBasedAuthProperties` | no |
| `PropertiesToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyClientProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `ClientProperties` | `ClientProperties` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyEndpointEncryptionMode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | yes |
| `EndpointEncryptionMode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifySamlProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `SamlProperties` | `SamlProperties` | no |
| `PropertiesToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifySelfservicePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `SelfservicePermissions` | `SelfservicePermissions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyStreamingProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `StreamingProperties` | `StreamingProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyWorkspaceAccessProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `WorkspaceAccessProperties` | `WorkspaceAccessProperties` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyWorkspaceCreationProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | yes |
| `WorkspaceCreationProperties` | `WorkspaceCreationProperties` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyWorkspaceProperties

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `WorkspaceProperties` | `WorkspaceProperties` | no |
| `DataReplication` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ModifyWorkspaceState

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `WorkspaceState` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RebootWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RebootWorkspaceRequests` | `List<RebootRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedWorkspaceChangeRequest>` | no |

## RebuildWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RebuildWorkspaceRequests` | `List<RebuildRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedWorkspaceChangeRequest>` | no |

## RegisterWorkspaceDirectory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `EnableSelfService` | `boolean` | no |
| `Tenancy` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `WorkspaceDirectoryName` | `string` | no |
| `WorkspaceDirectoryDescription` | `string` | no |
| `UserIdentityType` | `string` | no |
| `IdcInstanceArn` | `string` | no |
| `MicrosoftEntraConfig` | `MicrosoftEntraConfig` | no |
| `WorkspaceType` | `string` | no |
| `ActiveDirectoryConfig` | `ActiveDirectoryConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryId` | `string` | no |
| `State` | `string` | no |

## RejectAccountLinkInvitation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LinkId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLink` | `AccountLink` | no |

## RestoreWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeIpRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `UserRules` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartWorkspaceRequests` | `List<StartRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedWorkspaceChangeRequest>` | no |

## StartWorkspacesPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StopWorkspaceRequests` | `List<StopRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedWorkspaceChangeRequest>` | no |

## StopWorkspacesPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TerminateWorkspaceRequests` | `List<TerminateRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FailedRequests` | `List<FailedWorkspaceChangeRequest>` | no |

## TerminateWorkspacesPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateWorkspacesPoolSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnectClientAddIn

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AddInId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `Name` | `string` | no |
| `URL` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConnectionAliasPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AliasId` | `string` | yes |
| `ConnectionAliasPermission` | `ConnectionAliasPermission` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRulesOfIpGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `UserRules` | `List<IpRuleItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BundleId` | `string` | no |
| `ImageId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceImagePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageId` | `string` | yes |
| `AllowCopyImage` | `boolean` | yes |
| `SharedAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspacesPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolId` | `string` | yes |
| `Description` | `string` | no |
| `BundleId` | `string` | no |
| `DirectoryId` | `string` | no |
| `Capacity` | `Capacity` | no |
| `ApplicationSettings` | `ApplicationSettingsRequest` | no |
| `TimeoutSettings` | `TimeoutSettings` | no |
| `RunningMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspacesPool` | `WorkspacesPool` | no |

