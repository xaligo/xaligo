# Amazon AppStream

API version: 2016-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appstream/2016-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateAppBlockBuilderAppBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockArn` | `string` | yes |
| `AppBlockBuilderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilderAppBlockAssociation` | `AppBlockBuilderAppBlockAssociation` | no |

## AssociateApplicationFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | yes |
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationFleetAssociation` | `ApplicationFleetAssociation` | no |

## AssociateApplicationToEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `EntitlementName` | `string` | yes |
| `ApplicationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | yes |
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSoftwareToImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilderName` | `string` | yes |
| `SoftwareNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchAssociateUserStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserStackAssociations` | `List<UserStackAssociation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<UserStackAssociationError>` | no |

## BatchDisassociateUserStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserStackAssociations` | `List<UserStackAssociation>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<UserStackAssociationError>` | no |

## CopyImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceImageName` | `string` | yes |
| `DestinationImageName` | `string` | yes |
| `DestinationRegion` | `string` | yes |
| `DestinationImageDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DestinationImageName` | `string` | no |

## CreateAppBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `SourceS3Location` | `S3Location` | yes |
| `SetupScriptDetails` | `ScriptDetails` | no |
| `Tags` | `Map<string>` | no |
| `PostSetupScriptDetails` | `ScriptDetails` | no |
| `PackagingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlock` | `AppBlock` | no |

## CreateAppBlockBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Platform` | `string` | yes |
| `InstanceType` | `string` | yes |
| `VpcConfig` | `VpcConfig` | yes |
| `EnableDefaultInternetAccess` | `boolean` | no |
| `IamRoleArn` | `string` | no |
| `AccessEndpoints` | `List<AccessEndpoint>` | no |
| `DisableIMDSV1` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilder` | `AppBlockBuilder` | no |

## CreateAppBlockBuilderStreamingURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilderName` | `string` | yes |
| `Validity` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingURL` | `string` | no |
| `Expires` | `timestamp` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `IconS3Location` | `S3Location` | yes |
| `LaunchPath` | `string` | yes |
| `WorkingDirectory` | `string` | no |
| `LaunchParameters` | `string` | no |
| `Platforms` | `List<string>` | yes |
| `InstanceFamilies` | `List<string>` | yes |
| `AppBlockArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `Application` | no |

## CreateDirectoryConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryName` | `string` | yes |
| `OrganizationalUnitDistinguishedNames` | `List<string>` | yes |
| `ServiceAccountCredentials` | `ServiceAccountCredentials` | no |
| `CertificateBasedAuthProperties` | `CertificateBasedAuthProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryConfig` | `DirectoryConfig` | no |

## CreateEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `StackName` | `string` | yes |
| `Description` | `string` | no |
| `AppVisibility` | `string` | yes |
| `Attributes` | `List<EntitlementAttribute>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlement` | `Entitlement` | no |

## CreateExportImageTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `AmiName` | `string` | yes |
| `IamRoleArn` | `string` | yes |
| `TagSpecifications` | `Map<string>` | no |
| `AmiDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportImageTask` | `ExportImageTask` | no |

## CreateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ImageName` | `string` | no |
| `ImageArn` | `string` | no |
| `InstanceType` | `string` | yes |
| `FleetType` | `string` | no |
| `ComputeCapacity` | `ComputeCapacity` | no |
| `VpcConfig` | `VpcConfig` | no |
| `MaxUserDurationInSeconds` | `integer` | no |
| `DisconnectTimeoutInSeconds` | `integer` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `EnableDefaultInternetAccess` | `boolean` | no |
| `DomainJoinInfo` | `DomainJoinInfo` | no |
| `Tags` | `Map<string>` | no |
| `IdleDisconnectTimeoutInSeconds` | `integer` | no |
| `IamRoleArn` | `string` | no |
| `StreamView` | `string` | no |
| `Platform` | `string` | no |
| `MaxConcurrentSessions` | `integer` | no |
| `UsbDeviceFilterStrings` | `List<string>` | no |
| `SessionScriptS3Location` | `S3Location` | no |
| `MaxSessionsPerInstance` | `integer` | no |
| `RootVolumeConfig` | `VolumeConfig` | no |
| `DisableIMDSV1` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fleet` | `Fleet` | no |

## CreateImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ImageName` | `string` | no |
| `ImageArn` | `string` | no |
| `InstanceType` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `IamRoleArn` | `string` | no |
| `EnableDefaultInternetAccess` | `boolean` | no |
| `DomainJoinInfo` | `DomainJoinInfo` | no |
| `AppstreamAgentVersion` | `string` | no |
| `Tags` | `Map<string>` | no |
| `AccessEndpoints` | `List<AccessEndpoint>` | no |
| `RootVolumeConfig` | `VolumeConfig` | no |
| `SoftwaresToInstall` | `List<string>` | no |
| `SoftwaresToUninstall` | `List<string>` | no |
| `DisableIMDSV1` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilder` | `ImageBuilder` | no |

## CreateImageBuilderStreamingURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Validity` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingURL` | `string` | no |
| `Expires` | `timestamp` | no |

## CreateImportedImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SourceAmiId` | `string` | no |
| `WorkspaceImageId` | `string` | no |
| `IamRoleArn` | `string` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `RuntimeValidationConfig` | `RuntimeValidationConfig` | no |
| `AgentSoftwareVersion` | `string` | no |
| `AppCatalogConfig` | `List<ApplicationConfig>` | no |
| `DryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | no |

## CreateStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `StorageConnectors` | `List<StorageConnector>` | no |
| `RedirectURL` | `string` | no |
| `FeedbackURL` | `string` | no |
| `UserSettings` | `List<UserSetting>` | no |
| `ApplicationSettings` | `ApplicationSettings` | no |
| `Tags` | `Map<string>` | no |
| `AccessEndpoints` | `List<AccessEndpoint>` | no |
| `EmbedHostDomains` | `List<string>` | no |
| `StreamingExperienceSettings` | `StreamingExperienceSettings` | no |
| `ContentRedirection` | `ContentRedirection` | no |
| `AgentAccessConfig` | `AgentAccessConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Stack` | `Stack` | no |

## CreateStreamingURL

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `FleetName` | `string` | yes |
| `UserId` | `string` | yes |
| `ApplicationId` | `string` | no |
| `Validity` | `long` | no |
| `SessionContext` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingURL` | `string` | no |
| `Expires` | `timestamp` | no |

## CreateThemeForStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `FooterLinks` | `List<ThemeFooterLink>` | no |
| `TitleText` | `string` | yes |
| `ThemeStyling` | `string` | yes |
| `OrganizationLogoS3Location` | `S3Location` | yes |
| `FaviconS3Location` | `S3Location` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Theme` | `Theme` | no |

## CreateUpdatedImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `existingImageName` | `string` | yes |
| `newImageName` | `string` | yes |
| `newImageDescription` | `string` | no |
| `newImageDisplayName` | `string` | no |
| `newImageTags` | `Map<string>` | no |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `image` | `Image` | no |
| `canUpdateImage` | `boolean` | no |

## CreateUsageReportSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3BucketName` | `string` | no |
| `Schedule` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `MessageAction` | `string` | no |
| `FirstName` | `string` | no |
| `LastName` | `string` | no |
| `AuthenticationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppBlockBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDirectoryConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Image` | `Image` | no |

## DeleteImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilder` | `ImageBuilder` | no |

## DeleteImagePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SharedAccountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteThemeForStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUsageReportSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AuthenticationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAppBlockBuilderAppBlockAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockArn` | `string` | no |
| `AppBlockBuilderName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilderAppBlockAssociations` | `List<AppBlockBuilderAppBlockAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeAppBlockBuilders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilders` | `List<AppBlockBuilder>` | no |
| `NextToken` | `string` | no |

## DescribeAppBlocks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arns` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlocks` | `List<AppBlock>` | no |
| `NextToken` | `string` | no |

## DescribeAppLicenseUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BillingPeriod` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppLicenseUsages` | `List<AdminAppLicenseUsageRecord>` | no |
| `NextToken` | `string` | no |

## DescribeApplicationFleetAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | no |
| `ApplicationArn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationFleetAssociations` | `List<ApplicationFleetAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arns` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<Application>` | no |
| `NextToken` | `string` | no |

## DescribeDirectoryConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryNames` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryConfigs` | `List<DirectoryConfig>` | no |
| `NextToken` | `string` | no |

## DescribeEntitlements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `StackName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlements` | `List<Entitlement>` | no |
| `NextToken` | `string` | no |

## DescribeFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fleets` | `List<Fleet>` | no |
| `NextToken` | `string` | no |

## DescribeImageBuilders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilders` | `List<ImageBuilder>` | no |
| `NextToken` | `string` | no |

## DescribeImagePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `MaxResults` | `integer` | no |
| `SharedAwsAccountIds` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `SharedImagePermissionsList` | `List<SharedImagePermissions>` | no |
| `NextToken` | `string` | no |

## DescribeImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `Arns` | `List<string>` | no |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Images` | `List<Image>` | no |
| `NextToken` | `string` | no |

## DescribeSessions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `FleetName` | `string` | yes |
| `UserId` | `string` | no |
| `NextToken` | `string` | no |
| `Limit` | `integer` | no |
| `AuthenticationType` | `string` | no |
| `InstanceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Sessions` | `List<Session>` | no |
| `NextToken` | `string` | no |

## DescribeSoftwareAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedResource` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedResource` | `string` | no |
| `SoftwareAssociations` | `List<SoftwareAssociations>` | no |
| `NextToken` | `string` | no |

## DescribeStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Stacks` | `List<Stack>` | no |
| `NextToken` | `string` | no |

## DescribeThemeForStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Theme` | `Theme` | no |

## DescribeUsageReportSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UsageReportSubscriptions` | `List<UsageReportSubscription>` | no |
| `NextToken` | `string` | no |

## DescribeUserStackAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | no |
| `UserName` | `string` | no |
| `AuthenticationType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserStackAssociations` | `List<UserStackAssociation>` | no |
| `NextToken` | `string` | no |

## DescribeUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationType` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<User>` | no |
| `NextToken` | `string` | no |

## DisableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AuthenticationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateAppBlockBuilderAppBlock

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockArn` | `string` | yes |
| `AppBlockBuilderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApplicationFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | yes |
| `ApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApplicationFromEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `EntitlementName` | `string` | yes |
| `ApplicationIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | yes |
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSoftwareFromImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilderName` | `string` | yes |
| `SoftwareNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DrainSessionInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserName` | `string` | yes |
| `AuthenticationType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExpireSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetExportImageTask

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportImageTask` | `ExportImageTask` | no |

## ListAssociatedFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAssociatedStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FleetName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Names` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListEntitledApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `EntitlementName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EntitledApplications` | `List<EntitledApplication>` | no |
| `NextToken` | `string` | no |

## ListExportImageTasks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportImageTasks` | `List<ExportImageTask>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## StartAppBlockBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilder` | `AppBlockBuilder` | no |

## StartFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AppstreamAgentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilder` | `ImageBuilder` | no |

## StartSoftwareDeploymentToImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilderName` | `string` | yes |
| `RetryFailedDeployments` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopAppBlockBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilder` | `AppBlockBuilder` | no |

## StopFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopImageBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageBuilder` | `ImageBuilder` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | yes |

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


## UpdateAppBlockBuilder

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `Platform` | `string` | no |
| `InstanceType` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `EnableDefaultInternetAccess` | `boolean` | no |
| `IamRoleArn` | `string` | no |
| `AccessEndpoints` | `List<AccessEndpoint>` | no |
| `AttributesToDelete` | `List<string>` | no |
| `DisableIMDSV1` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppBlockBuilder` | `AppBlockBuilder` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `IconS3Location` | `S3Location` | no |
| `LaunchPath` | `string` | no |
| `WorkingDirectory` | `string` | no |
| `LaunchParameters` | `string` | no |
| `AppBlockArn` | `string` | no |
| `AttributesToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `Application` | no |

## UpdateDirectoryConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryName` | `string` | yes |
| `OrganizationalUnitDistinguishedNames` | `List<string>` | no |
| `ServiceAccountCredentials` | `ServiceAccountCredentials` | no |
| `CertificateBasedAuthProperties` | `CertificateBasedAuthProperties` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryConfig` | `DirectoryConfig` | no |

## UpdateEntitlement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `StackName` | `string` | yes |
| `Description` | `string` | no |
| `AppVisibility` | `string` | no |
| `Attributes` | `List<EntitlementAttribute>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Entitlement` | `Entitlement` | no |

## UpdateFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | no |
| `ImageArn` | `string` | no |
| `Name` | `string` | no |
| `InstanceType` | `string` | no |
| `ComputeCapacity` | `ComputeCapacity` | no |
| `VpcConfig` | `VpcConfig` | no |
| `MaxUserDurationInSeconds` | `integer` | no |
| `DisconnectTimeoutInSeconds` | `integer` | no |
| `DeleteVpcConfig` | `boolean` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `EnableDefaultInternetAccess` | `boolean` | no |
| `DomainJoinInfo` | `DomainJoinInfo` | no |
| `IdleDisconnectTimeoutInSeconds` | `integer` | no |
| `AttributesToDelete` | `List<string>` | no |
| `IamRoleArn` | `string` | no |
| `StreamView` | `string` | no |
| `Platform` | `string` | no |
| `MaxConcurrentSessions` | `integer` | no |
| `UsbDeviceFilterStrings` | `List<string>` | no |
| `SessionScriptS3Location` | `S3Location` | no |
| `MaxSessionsPerInstance` | `integer` | no |
| `RootVolumeConfig` | `VolumeConfig` | no |
| `DisableIMDSV1` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Fleet` | `Fleet` | no |

## UpdateImagePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SharedAccountId` | `string` | yes |
| `ImagePermissions` | `ImagePermissions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `StorageConnectors` | `List<StorageConnector>` | no |
| `DeleteStorageConnectors` | `boolean` | no |
| `RedirectURL` | `string` | no |
| `FeedbackURL` | `string` | no |
| `AttributesToDelete` | `List<string>` | no |
| `UserSettings` | `List<UserSetting>` | no |
| `ApplicationSettings` | `ApplicationSettings` | no |
| `AccessEndpoints` | `List<AccessEndpoint>` | no |
| `EmbedHostDomains` | `List<string>` | no |
| `StreamingExperienceSettings` | `StreamingExperienceSettings` | no |
| `ContentRedirection` | `ContentRedirection` | no |
| `AgentAccessConfig` | `AgentAccessConfigForUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Stack` | `Stack` | no |

## UpdateThemeForStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `FooterLinks` | `List<ThemeFooterLink>` | no |
| `TitleText` | `string` | no |
| `ThemeStyling` | `string` | no |
| `OrganizationLogoS3Location` | `S3Location` | no |
| `FaviconS3Location` | `S3Location` | no |
| `State` | `string` | no |
| `AttributesToDelete` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Theme` | `Theme` | no |

