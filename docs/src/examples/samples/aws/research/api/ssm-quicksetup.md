# AWS Systems Manager QuickSetup

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm-quicksetup/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConfigurationManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationDefinitions` | `List<ConfigurationDefinitionInput>` | yes |
| `Description` | `string` | no |
| `Name` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagerArn` | `string` | yes |

## DeleteConfigurationManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Account` | `string` | no |
| `ConfigurationDefinitionId` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `Id` | `string` | no |
| `LastModifiedAt` | `timestamp` | no |
| `ManagerArn` | `string` | no |
| `Parameters` | `Map<string>` | no |
| `Region` | `string` | no |
| `StatusSummaries` | `List<StatusSummary>` | no |
| `Type` | `string` | no |
| `TypeVersion` | `string` | no |

## GetConfigurationManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ManagerArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationDefinitions` | `List<ConfigurationDefinition>` | no |
| `CreatedAt` | `timestamp` | no |
| `Description` | `string` | no |
| `LastModifiedAt` | `timestamp` | no |
| `ManagerArn` | `string` | yes |
| `Name` | `string` | no |
| `StatusSummaries` | `List<StatusSummary>` | no |
| `Tags` | `Map<string>` | no |

## GetServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServiceSettings` | `ServiceSettings` | no |

## ListConfigurationManagers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxItems` | `integer` | no |
| `StartingToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationManagersList` | `List<ConfigurationManagerSummary>` | no |
| `NextToken` | `string` | no |

## ListConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationDefinitionId` | `string` | no |
| `Filters` | `List<Filter>` | no |
| `ManagerArn` | `string` | no |
| `MaxItems` | `integer` | no |
| `StartingToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationsList` | `List<ConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListQuickSetupTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuickSetupTypeList` | `List<QuickSetupTypeOutput>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<TagEntry>` | no |

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


## UpdateConfigurationDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `LocalDeploymentAdministrationRoleArn` | `string` | no |
| `LocalDeploymentExecutionRoleName` | `string` | no |
| `ManagerArn` | `string` | yes |
| `Parameters` | `Map<string>` | no |
| `TypeVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `ManagerArn` | `string` | yes |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExplorerEnablingRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


