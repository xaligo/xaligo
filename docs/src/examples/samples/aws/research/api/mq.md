# AmazonMQ

API version: 2017-11-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mq/2017-11-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationStrategy` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `BrokerName` | `string` | yes |
| `Configuration` | `ConfigurationId` | no |
| `CreatorRequestId` | `string` | no |
| `DeploymentMode` | `string` | yes |
| `EncryptionOptions` | `EncryptionOptions` | no |
| `EngineType` | `string` | yes |
| `EngineVersion` | `string` | no |
| `HostInstanceType` | `string` | yes |
| `LdapServerMetadata` | `LdapServerMetadataInput` | no |
| `Logs` | `Logs` | no |
| `MaintenanceWindowStartTime` | `WeeklyStartTime` | no |
| `PubliclyAccessible` | `boolean` | yes |
| `SecurityGroups` | `List<string>` | no |
| `StorageSize` | `integer` | no |
| `StorageType` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `Users` | `List<User>` | no |
| `DataReplicationMode` | `string` | no |
| `DataReplicationPrimaryBrokerArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerArn` | `string` | no |
| `BrokerId` | `string` | no |

## CreateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationStrategy` | `string` | no |
| `EngineType` | `string` | yes |
| `EngineVersion` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AuthenticationStrategy` | `string` | no |
| `Created` | `timestamp` | no |
| `Id` | `string` | no |
| `LatestRevision` | `ConfigurationRevision` | no |
| `Name` | `string` | no |

## CreateTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `ConsoleAccess` | `boolean` | no |
| `Groups` | `List<string>` | no |
| `Password` | `string` | yes |
| `Username` | `string` | yes |
| `ReplicationUser` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | no |

## DeleteConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | no |

## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionsRequired` | `List<ActionRequired>` | no |
| `AuthenticationStrategy` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `BrokerArn` | `string` | no |
| `BrokerId` | `string` | no |
| `BrokerInstances` | `List<BrokerInstance>` | no |
| `BrokerName` | `string` | no |
| `BrokerState` | `string` | no |
| `Configurations` | `Configurations` | no |
| `Created` | `timestamp` | no |
| `DeploymentMode` | `string` | no |
| `EncryptionOptions` | `EncryptionOptions` | no |
| `EngineType` | `string` | no |
| `EngineVersion` | `string` | no |
| `HostInstanceType` | `string` | no |
| `LdapServerMetadata` | `LdapServerMetadataOutput` | no |
| `Logs` | `LogsSummary` | no |
| `MaintenanceWindowStartTime` | `WeeklyStartTime` | no |
| `PendingAuthenticationStrategy` | `string` | no |
| `PendingEngineVersion` | `string` | no |
| `PendingHostInstanceType` | `string` | no |
| `PendingLdapServerMetadata` | `LdapServerMetadataOutput` | no |
| `PendingSecurityGroups` | `List<string>` | no |
| `PendingStorageSize` | `integer` | no |
| `PubliclyAccessible` | `boolean` | no |
| `SecurityGroups` | `List<string>` | no |
| `StorageSize` | `integer` | no |
| `StorageType` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `Users` | `List<UserSummary>` | no |
| `DataReplicationMetadata` | `DataReplicationMetadataOutput` | no |
| `DataReplicationMode` | `string` | no |
| `PendingDataReplicationMetadata` | `DataReplicationMetadataOutput` | no |
| `PendingDataReplicationMode` | `string` | no |

## DescribeBrokerEngineTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerEngineTypes` | `List<BrokerEngineType>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

## DescribeBrokerInstanceOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngineType` | `string` | no |
| `HostInstanceType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `StorageType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerInstanceOptions` | `List<BrokerInstanceOption>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

## DescribeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `AuthenticationStrategy` | `string` | no |
| `Created` | `timestamp` | no |
| `Description` | `string` | no |
| `EngineType` | `string` | no |
| `EngineVersion` | `string` | no |
| `Id` | `string` | no |
| `LatestRevision` | `ConfigurationRevision` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

## DescribeConfigurationRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |
| `ConfigurationRevision` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | no |
| `Created` | `timestamp` | no |
| `Data` | `string` | no |
| `Description` | `string` | no |

## DescribeSharedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SharedResources` | `List<SharedResource>` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `Username` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | no |
| `ConsoleAccess` | `boolean` | no |
| `Groups` | `List<string>` | no |
| `Pending` | `UserPendingChanges` | no |
| `Username` | `string` | no |
| `ReplicationUser` | `boolean` | no |

## ListBrokers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerSummaries` | `List<BrokerSummary>` | no |
| `NextToken` | `string` | no |

## ListConfigurationRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Revisions` | `List<ConfigurationRevision>` | no |

## ListConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Configurations` | `List<Configuration>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Users` | `List<UserSummary>` | no |

## Promote

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `Mode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | no |

## RebootBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateBroker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationStrategy` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `BrokerId` | `string` | yes |
| `Configuration` | `ConfigurationId` | no |
| `EngineVersion` | `string` | no |
| `HostInstanceType` | `string` | no |
| `LdapServerMetadata` | `LdapServerMetadataInput` | no |
| `Logs` | `Logs` | no |
| `MaintenanceWindowStartTime` | `WeeklyStartTime` | no |
| `ResourceShareArns` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `StorageSize` | `integer` | no |
| `DataReplicationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationStrategy` | `string` | no |
| `AutoMinorVersionUpgrade` | `boolean` | no |
| `BrokerId` | `string` | no |
| `Configuration` | `ConfigurationId` | no |
| `EngineVersion` | `string` | no |
| `HostInstanceType` | `string` | no |
| `LdapServerMetadata` | `LdapServerMetadataOutput` | no |
| `Logs` | `Logs` | no |
| `MaintenanceWindowStartTime` | `WeeklyStartTime` | no |
| `ResourceShareArns` | `List<string>` | no |
| `SecurityGroups` | `List<string>` | no |
| `DataReplicationMetadata` | `DataReplicationMetadataOutput` | no |
| `DataReplicationMode` | `string` | no |
| `PendingDataReplicationMetadata` | `DataReplicationMetadataOutput` | no |
| `PendingDataReplicationMode` | `string` | no |
| `StorageSize` | `integer` | no |

## UpdateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |
| `Data` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Created` | `timestamp` | no |
| `Id` | `string` | no |
| `LatestRevision` | `ConfigurationRevision` | no |
| `Name` | `string` | no |
| `Warnings` | `List<SanitizationWarning>` | no |

## UpdateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BrokerId` | `string` | yes |
| `ConsoleAccess` | `boolean` | no |
| `Groups` | `List<string>` | no |
| `Password` | `string` | no |
| `Username` | `string` | yes |
| `ReplicationUser` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


