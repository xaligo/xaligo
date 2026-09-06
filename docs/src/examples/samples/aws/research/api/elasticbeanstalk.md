# AWS Elastic Beanstalk

API version: 2010-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/elasticbeanstalk/2010-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AbortEnvironmentUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ApplyEnvironmentManagedAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `ActionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionId` | `string` | no |
| `ActionDescription` | `string` | no |
| `ActionType` | `string` | no |
| `Status` | `string` | no |

## AssociateEnvironmentOperationsRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | yes |
| `OperationsRole` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckDNSAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CNAMEPrefix` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Available` | `boolean` | no |
| `FullyQualifiedCNAME` | `string` | no |

## ComposeEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `GroupName` | `string` | no |
| `VersionLabels` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Environments` | `List<EnvironmentDescription>` | no |
| `NextToken` | `string` | no |

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Description` | `string` | no |
| `ResourceLifecycleConfig` | `ApplicationResourceLifecycleConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `ApplicationDescription` | no |

## CreateApplicationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `VersionLabel` | `string` | yes |
| `Description` | `string` | no |
| `SourceBuildInformation` | `SourceBuildInformation` | no |
| `SourceBundle` | `S3Location` | no |
| `BuildConfiguration` | `BuildConfiguration` | no |
| `AutoCreateApplication` | `boolean` | no |
| `Process` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationVersion` | `ApplicationVersionDescription` | no |

## CreateConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TemplateName` | `string` | yes |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `SourceConfiguration` | `SourceConfiguration` | no |
| `EnvironmentId` | `string` | no |
| `Description` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `ApplicationName` | `string` | no |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `EnvironmentName` | `string` | no |
| `DeploymentStatus` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateUpdated` | `timestamp` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `EnvironmentName` | `string` | no |
| `GroupName` | `string` | no |
| `Description` | `string` | no |
| `CNAMEPrefix` | `string` | no |
| `Tier` | `EnvironmentTier` | no |
| `Tags` | `List<Tag>` | no |
| `VersionLabel` | `string` | no |
| `TemplateName` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |
| `OptionsToRemove` | `List<OptionSpecification>` | no |
| `OperationsRole` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `ApplicationName` | `string` | no |
| `VersionLabel` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `EndpointURL` | `string` | no |
| `CNAME` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateUpdated` | `timestamp` | no |
| `Status` | `string` | no |
| `AbortableOperationInProgress` | `boolean` | no |
| `Health` | `string` | no |
| `HealthStatus` | `string` | no |
| `Resources` | `EnvironmentResourcesDescription` | no |
| `Tier` | `EnvironmentTier` | no |
| `EnvironmentLinks` | `List<EnvironmentLink>` | no |
| `EnvironmentArn` | `string` | no |
| `OperationsRole` | `string` | no |

## CreatePlatformVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformName` | `string` | yes |
| `PlatformVersion` | `string` | yes |
| `PlatformDefinitionBundle` | `S3Location` | yes |
| `EnvironmentName` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformSummary` | `PlatformSummary` | no |
| `Builder` | `Builder` | no |

## CreateStorageLocation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `S3Bucket` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TerminateEnvByForce` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApplicationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `VersionLabel` | `string` | yes |
| `DeleteSourceBundle` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `EnvironmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlatformVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformSummary` | `PlatformSummary` | no |

## DescribeAccountAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceQuotas` | `ResourceQuotas` | no |

## DescribeApplicationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `VersionLabels` | `List<string>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationVersions` | `List<ApplicationVersionDescription>` | no |
| `NextToken` | `string` | no |

## DescribeApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<ApplicationDescription>` | no |

## DescribeConfigurationOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `TemplateName` | `string` | no |
| `EnvironmentName` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `Options` | `List<OptionSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `Options` | `List<ConfigurationOptionDescription>` | no |

## DescribeConfigurationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TemplateName` | `string` | no |
| `EnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSettings` | `List<ConfigurationSettingsDescription>` | no |

## DescribeEnvironmentHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `AttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `HealthStatus` | `string` | no |
| `Status` | `string` | no |
| `Color` | `string` | no |
| `Causes` | `List<string>` | no |
| `ApplicationMetrics` | `ApplicationMetrics` | no |
| `InstancesHealth` | `InstanceHealthSummary` | no |
| `RefreshedAt` | `timestamp` | no |

## DescribeEnvironmentManagedActionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedActionHistoryItems` | `List<ManagedActionHistoryItem>` | no |
| `NextToken` | `string` | no |

## DescribeEnvironmentManagedActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagedActions` | `List<ManagedAction>` | no |

## DescribeEnvironmentResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentResources` | `EnvironmentResourceDescription` | no |

## DescribeEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `VersionLabel` | `string` | no |
| `EnvironmentIds` | `List<string>` | no |
| `EnvironmentNames` | `List<string>` | no |
| `IncludeDeleted` | `boolean` | no |
| `IncludedDeletedBackTo` | `timestamp` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Environments` | `List<EnvironmentDescription>` | no |
| `NextToken` | `string` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `VersionLabel` | `string` | no |
| `TemplateName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `PlatformArn` | `string` | no |
| `RequestId` | `string` | no |
| `Severity` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<EventDescription>` | no |
| `NextToken` | `string` | no |

## DescribeInstancesHealth

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `AttributeNames` | `List<string>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceHealthList` | `List<SingleInstanceHealth>` | no |
| `RefreshedAt` | `timestamp` | no |
| `NextToken` | `string` | no |

## DescribePlatformVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformDescription` | `PlatformDescription` | no |

## DisassociateEnvironmentOperationsRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListAvailableSolutionStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SolutionStacks` | `List<string>` | no |
| `SolutionStackDetails` | `List<SolutionStackDescription>` | no |

## ListPlatformBranches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<SearchFilter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformBranchSummaryList` | `List<PlatformBranchSummary>` | no |
| `NextToken` | `string` | no |

## ListPlatformVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<PlatformFilter>` | no |
| `MaxRecords` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformSummaryList` | `List<PlatformSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | no |
| `ResourceTags` | `List<Tag>` | no |

## RebuildEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RequestEnvironmentInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `InfoType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RestartAppServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RetrieveEnvironmentInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `InfoType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentInfo` | `List<EnvironmentInfoDescription>` | no |

## SwapEnvironmentCNAMEs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceEnvironmentId` | `string` | no |
| `SourceEnvironmentName` | `string` | no |
| `DestinationEnvironmentId` | `string` | no |
| `DestinationEnvironmentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TerminateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `TerminateResources` | `boolean` | no |
| `ForceTerminate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `ApplicationName` | `string` | no |
| `VersionLabel` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `EndpointURL` | `string` | no |
| `CNAME` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateUpdated` | `timestamp` | no |
| `Status` | `string` | no |
| `AbortableOperationInProgress` | `boolean` | no |
| `Health` | `string` | no |
| `HealthStatus` | `string` | no |
| `Resources` | `EnvironmentResourcesDescription` | no |
| `Tier` | `EnvironmentTier` | no |
| `EnvironmentLinks` | `List<EnvironmentLink>` | no |
| `EnvironmentArn` | `string` | no |
| `OperationsRole` | `string` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `ApplicationDescription` | no |

## UpdateApplicationResourceLifecycle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `ResourceLifecycleConfig` | `ApplicationResourceLifecycleConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `ResourceLifecycleConfig` | `ApplicationResourceLifecycleConfig` | no |

## UpdateApplicationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `VersionLabel` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationVersion` | `ApplicationVersionDescription` | no |

## UpdateConfigurationTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TemplateName` | `string` | yes |
| `Description` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |
| `OptionsToRemove` | `List<OptionSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `ApplicationName` | `string` | no |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `EnvironmentName` | `string` | no |
| `DeploymentStatus` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateUpdated` | `timestamp` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |

## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `EnvironmentName` | `string` | no |
| `GroupName` | `string` | no |
| `Description` | `string` | no |
| `Tier` | `EnvironmentTier` | no |
| `VersionLabel` | `string` | no |
| `TemplateName` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | no |
| `OptionsToRemove` | `List<OptionSpecification>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | no |
| `EnvironmentId` | `string` | no |
| `ApplicationName` | `string` | no |
| `VersionLabel` | `string` | no |
| `SolutionStackName` | `string` | no |
| `PlatformArn` | `string` | no |
| `TemplateName` | `string` | no |
| `Description` | `string` | no |
| `EndpointURL` | `string` | no |
| `CNAME` | `string` | no |
| `DateCreated` | `timestamp` | no |
| `DateUpdated` | `timestamp` | no |
| `Status` | `string` | no |
| `AbortableOperationInProgress` | `boolean` | no |
| `Health` | `string` | no |
| `HealthStatus` | `string` | no |
| `Resources` | `EnvironmentResourcesDescription` | no |
| `Tier` | `EnvironmentTier` | no |
| `EnvironmentLinks` | `List<EnvironmentLink>` | no |
| `EnvironmentArn` | `string` | no |
| `OperationsRole` | `string` | no |

## UpdateTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagsToAdd` | `List<Tag>` | no |
| `TagsToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ValidateConfigurationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationName` | `string` | yes |
| `TemplateName` | `string` | no |
| `EnvironmentName` | `string` | no |
| `OptionSettings` | `List<ConfigurationOptionSetting>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Messages` | `List<ValidationMessage>` | no |

