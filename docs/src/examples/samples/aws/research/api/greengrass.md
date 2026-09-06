# AWS Greengrass

API version: 2017-06-07. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/greengrass/2017-06-07/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateRoleToGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedAt` | `string` | no |

## AssociateServiceRoleToAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedAt` | `string` | no |

## CreateConnectorDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `ConnectorDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateConnectorDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `ConnectorDefinitionId` | `string` | yes |
| `Connectors` | `List<Connector>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateCoreDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `CoreDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateCoreDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `CoreDefinitionId` | `string` | yes |
| `Cores` | `List<Core>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `DeploymentId` | `string` | no |
| `DeploymentType` | `string` | yes |
| `GroupId` | `string` | yes |
| `GroupVersionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentArn` | `string` | no |
| `DeploymentId` | `string` | no |

## CreateDeviceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `DeviceDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateDeviceDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `DeviceDefinitionId` | `string` | yes |
| `Devices` | `List<Device>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateFunctionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `FunctionDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateFunctionDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `DefaultConfig` | `FunctionDefaultConfig` | no |
| `FunctionDefinitionId` | `string` | yes |
| `Functions` | `List<Function>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `GroupVersion` | no |
| `Name` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateGroupCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupCertificateAuthorityArn` | `string` | no |

## CreateGroupVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `ConnectorDefinitionVersionArn` | `string` | no |
| `CoreDefinitionVersionArn` | `string` | no |
| `DeviceDefinitionVersionArn` | `string` | no |
| `FunctionDefinitionVersionArn` | `string` | no |
| `GroupId` | `string` | yes |
| `LoggerDefinitionVersionArn` | `string` | no |
| `ResourceDefinitionVersionArn` | `string` | no |
| `SubscriptionDefinitionVersionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateLoggerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `LoggerDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateLoggerDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `LoggerDefinitionId` | `string` | yes |
| `Loggers` | `List<Logger>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateResourceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `ResourceDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateResourceDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `ResourceDefinitionId` | `string` | yes |
| `Resources` | `List<Resource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## CreateSoftwareUpdateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `S3UrlSignerRole` | `string` | yes |
| `SoftwareToUpdate` | `string` | yes |
| `UpdateAgentLogLevel` | `string` | no |
| `UpdateTargets` | `List<string>` | yes |
| `UpdateTargetsArchitecture` | `string` | yes |
| `UpdateTargetsOperatingSystem` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IotJobArn` | `string` | no |
| `IotJobId` | `string` | no |
| `PlatformSoftwareVersion` | `string` | no |

## CreateSubscriptionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `InitialVersion` | `SubscriptionDefinitionVersion` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |

## CreateSubscriptionDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `SubscriptionDefinitionId` | `string` | yes |
| `Subscriptions` | `List<Subscription>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## DeleteConnectorDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCoreDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeviceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunctionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLoggerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggerDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriptionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateRoleFromGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisassociatedAt` | `string` | no |

## DisassociateServiceRoleFromAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisassociatedAt` | `string` | no |

## GetAssociatedRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedAt` | `string` | no |
| `RoleArn` | `string` | no |

## GetBulkDeploymentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeploymentMetrics` | `BulkDeploymentMetrics` | no |
| `BulkDeploymentStatus` | `string` | no |
| `CreatedAt` | `string` | no |
| `ErrorDetails` | `List<ErrorDetail>` | no |
| `ErrorMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## GetConnectivityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectivityInfo` | `List<ConnectivityInfo>` | no |
| `Message` | `string` | no |

## GetConnectorDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetConnectorDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDefinitionId` | `string` | yes |
| `ConnectorDefinitionVersionId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `ConnectorDefinitionVersion` | no |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `Version` | `string` | no |

## GetCoreDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetCoreDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreDefinitionId` | `string` | yes |
| `CoreDefinitionVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `CoreDefinitionVersion` | no |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `Version` | `string` | no |

## GetDeploymentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentStatus` | `string` | no |
| `DeploymentType` | `string` | no |
| `ErrorDetails` | `List<ErrorDetail>` | no |
| `ErrorMessage` | `string` | no |
| `UpdatedAt` | `string` | no |

## GetDeviceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetDeviceDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDefinitionId` | `string` | yes |
| `DeviceDefinitionVersionId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `DeviceDefinitionVersion` | no |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `Version` | `string` | no |

## GetFunctionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetFunctionDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionDefinitionId` | `string` | yes |
| `FunctionDefinitionVersionId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `FunctionDefinitionVersion` | no |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `Version` | `string` | no |

## GetGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetGroupCertificateAuthority

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityId` | `string` | yes |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupCertificateAuthorityArn` | `string` | no |
| `GroupCertificateAuthorityId` | `string` | no |
| `PemEncodedCertificate` | `string` | no |

## GetGroupCertificateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityExpiryInMilliseconds` | `string` | no |
| `CertificateExpiryInMilliseconds` | `string` | no |
| `GroupId` | `string` | no |

## GetGroupVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `GroupVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `GroupVersion` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## GetLoggerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggerDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetLoggerDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggerDefinitionId` | `string` | yes |
| `LoggerDefinitionVersionId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `LoggerDefinitionVersion` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## GetResourceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetResourceDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceDefinitionId` | `string` | yes |
| `ResourceDefinitionVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `ResourceDefinitionVersion` | no |
| `Id` | `string` | no |
| `Version` | `string` | no |

## GetServiceRoleForAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedAt` | `string` | no |
| `RoleArn` | `string` | no |

## GetSubscriptionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Id` | `string` | no |
| `LastUpdatedTimestamp` | `string` | no |
| `LatestVersion` | `string` | no |
| `LatestVersionArn` | `string` | no |
| `Name` | `string` | no |
| `tags` | `Map<string>` | no |

## GetSubscriptionDefinitionVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SubscriptionDefinitionId` | `string` | yes |
| `SubscriptionDefinitionVersionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreationTimestamp` | `string` | no |
| `Definition` | `SubscriptionDefinitionVersion` | no |
| `Id` | `string` | no |
| `NextToken` | `string` | no |
| `Version` | `string` | no |

## GetThingRuntimeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuntimeConfiguration` | `RuntimeConfiguration` | no |

## ListBulkDeploymentDetailedReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeploymentId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deployments` | `List<BulkDeploymentResult>` | no |
| `NextToken` | `string` | no |

## ListBulkDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeployments` | `List<BulkDeployment>` | no |
| `NextToken` | `string` | no |

## ListConnectorDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDefinitionId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListConnectorDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListCoreDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreDefinitionId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListCoreDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deployments` | `List<Deployment>` | no |
| `NextToken` | `string` | no |

## ListDeviceDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDefinitionId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListDeviceDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListFunctionDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionDefinitionId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListFunctionDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListGroupCertificateAuthorities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupCertificateAuthorities` | `List<GroupCertificateAuthorityProperties>` | no |

## ListGroupVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Groups` | `List<GroupInformation>` | no |
| `NextToken` | `string` | no |

## ListLoggerDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggerDefinitionId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListLoggerDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListResourceDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `ResourceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListResourceDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListSubscriptionDefinitionVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `SubscriptionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Versions` | `List<VersionInformation>` | no |

## ListSubscriptionDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definitions` | `List<DefinitionInformation>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ResetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `Force` | `boolean` | no |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentArn` | `string` | no |
| `DeploymentId` | `string` | no |

## StartBulkDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AmznClientToken` | `string` | no |
| `ExecutionRoleArn` | `string` | yes |
| `InputFileUri` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeploymentArn` | `string` | no |
| `BulkDeploymentId` | `string` | no |

## StopBulkDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkDeploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `tags` | `Map<string>` | no |

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


## UpdateConnectivityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectivityInfo` | `List<ConnectivityInfo>` | no |
| `ThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |
| `Version` | `string` | no |

## UpdateConnectorDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorDefinitionId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCoreDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CoreDefinitionId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDeviceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDefinitionId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateFunctionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FunctionDefinitionId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateGroupCertificateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateExpiryInMilliseconds` | `string` | no |
| `GroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityExpiryInMilliseconds` | `string` | no |
| `CertificateExpiryInMilliseconds` | `string` | no |
| `GroupId` | `string` | no |

## UpdateLoggerDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LoggerDefinitionId` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateResourceDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `ResourceDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSubscriptionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `SubscriptionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateThingRuntimeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelemetryConfiguration` | `TelemetryConfigurationUpdate` | no |
| `ThingName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


