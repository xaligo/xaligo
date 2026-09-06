# Amazon AppConfig

API version: 2019-10-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appconfig/2019-10-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## CreateConfigurationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `LocationUri` | `string` | yes |
| `RetrievalRoleArn` | `string` | no |
| `Validators` | `List<Validator>` | no |
| `Tags` | `Map<string>` | no |
| `Type` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LocationUri` | `string` | no |
| `RetrievalRoleArn` | `string` | no |
| `Validators` | `List<Validator>` | no |
| `Type` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |

## CreateDeploymentStrategy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | yes |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `GrowthFactor` | `float` | yes |
| `GrowthType` | `string` | no |
| `ReplicateTo` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `ReplicateTo` | `string` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Monitors` | `List<Monitor>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |
| `Monitors` | `List<Monitor>` | no |

## CreateExperimentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `Name` | `string` | yes |
| `ConfigurationProfileIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `FlagKey` | `string` | yes |
| `Treatments` | `List<TreatmentInput>` | yes |
| `Control` | `TreatmentInput` | yes |
| `AudienceRule` | `string` | yes |
| `Hypothesis` | `string` | no |
| `AudienceDescription` | `string` | no |
| `LaunchCriteria` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Hypothesis` | `string` | no |
| `Status` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `FlagKey` | `string` | no |
| `AudienceRule` | `string` | no |
| `AudienceDescription` | `string` | no |
| `LaunchCriteria` | `string` | no |
| `Treatments` | `List<Treatment>` | no |
| `Control` | `Treatment` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `KmsKeyIdentifier` | `string` | no |

## CreateExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Actions` | `Map<List<Action>>` | yes |
| `Parameters` | `Map<Parameter>` | no |
| `Tags` | `Map<string>` | no |
| `LatestVersionNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `VersionNumber` | `integer` | no |
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `Actions` | `Map<List<Action>>` | no |
| `Parameters` | `Map<Parameter>` | no |

## CreateExtensionAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionIdentifier` | `string` | yes |
| `ExtensionVersionNumber` | `integer` | no |
| `ResourceIdentifier` | `string` | yes |
| `Parameters` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ExtensionArn` | `string` | no |
| `ResourceArn` | `string` | no |
| `Arn` | `string` | no |
| `Parameters` | `Map<string>` | no |
| `ExtensionVersionNumber` | `integer` | no |

## CreateHostedConfigurationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `Description` | `string` | no |
| `Content` | `blob` | yes |
| `ContentType` | `string` | yes |
| `LatestVersionNumber` | `integer` | no |
| `VersionLabel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `VersionNumber` | `integer` | no |
| `Description` | `string` | no |
| `Content` | `blob` | no |
| `ContentType` | `string` | no |
| `VersionLabel` | `string` | no |
| `KmsKeyArn` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `DeletionProtectionCheck` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeploymentStrategy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentStrategyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentId` | `string` | yes |
| `ApplicationId` | `string` | yes |
| `DeletionProtectionCheck` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExperimentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `DeleteType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionIdentifier` | `string` | yes |
| `VersionNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExtensionAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHostedConfigurationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `VersionNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionProtection` | `DeletionProtectionSettings` | no |
| `VendedMetrics` | `VendedMetricsSettings` | no |

## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## GetConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `string` | yes |
| `Environment` | `string` | yes |
| `Configuration` | `string` | yes |
| `ClientId` | `string` | yes |
| `ClientConfigurationVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Content` | `blob` | no |
| `ConfigurationVersion` | `string` | no |
| `ContentType` | `string` | no |

## GetConfigurationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LocationUri` | `string` | no |
| `RetrievalRoleArn` | `string` | no |
| `Validators` | `List<Validator>` | no |
| `Type` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |
| `DeploymentNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `DeploymentStrategyId` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `DeploymentNumber` | `integer` | no |
| `ConfigurationName` | `string` | no |
| `ConfigurationLocationUri` | `string` | no |
| `ConfigurationVersion` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `State` | `string` | no |
| `EventLog` | `List<DeploymentEvent>` | no |
| `PercentageComplete` | `float` | no |
| `StartedAt` | `timestamp` | no |
| `CompletedAt` | `timestamp` | no |
| `AppliedExtensions` | `List<AppliedExtension>` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `VersionLabel` | `string` | no |

## GetDeploymentStrategy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentStrategyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `ReplicateTo` | `string` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |
| `Monitors` | `List<Monitor>` | no |

## GetExperimentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Hypothesis` | `string` | no |
| `Status` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `FlagKey` | `string` | no |
| `AudienceRule` | `string` | no |
| `AudienceDescription` | `string` | no |
| `LaunchCriteria` | `string` | no |
| `Treatments` | `List<Treatment>` | no |
| `Control` | `Treatment` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `KmsKeyIdentifier` | `string` | no |

## GetExperimentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Run` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ExperimentDefinitionId` | `string` | no |
| `Run` | `integer` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `Result` | `ExperimentRunResult` | no |
| `StartedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `ExperimentDefinitionSnapshot` | `ExperimentDefinitionSnapshot` | no |

## GetExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionIdentifier` | `string` | yes |
| `VersionNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `VersionNumber` | `integer` | no |
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `Actions` | `Map<List<Action>>` | no |
| `Parameters` | `Map<Parameter>` | no |

## GetExtensionAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ExtensionArn` | `string` | no |
| `ResourceArn` | `string` | no |
| `Arn` | `string` | no |
| `Parameters` | `Map<string>` | no |
| `ExtensionVersionNumber` | `integer` | no |

## GetHostedConfigurationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `VersionNumber` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `VersionNumber` | `integer` | no |
| `Description` | `string` | no |
| `Content` | `blob` | no |
| `ContentType` | `string` | no |
| `VersionLabel` | `string` | no |
| `KmsKeyArn` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Application>` | no |
| `NextToken` | `string` | no |

## ListConfigurationProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ConfigurationProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListDeploymentStrategies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DeploymentStrategy>` | no |
| `NextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DeploymentSummary>` | no |
| `NextToken` | `string` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Environment>` | no |
| `NextToken` | `string` | no |

## ListExperimentDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | no |
| `ConfigurationProfileIdentifier` | `string` | no |
| `EnvironmentIdentifier` | `string` | no |
| `Status` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ExperimentDefinitionSummary>` | no |
| `NextToken` | `string` | no |

## ListExperimentRunEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Run` | `integer` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ExperimentRunEvent>` | no |
| `NextToken` | `string` | no |

## ListExperimentRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ExperimentRunSummary>` | no |
| `NextToken` | `string` | no |

## ListExtensionAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | no |
| `ExtensionIdentifier` | `string` | no |
| `ExtensionVersionNumber` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ExtensionAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListExtensions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ExtensionSummary>` | no |
| `NextToken` | `string` | no |

## ListHostedConfigurationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `VersionLabel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<HostedConfigurationVersionSummary>` | no |
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

## StartDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |
| `DeploymentStrategyId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `ConfigurationVersion` | `string` | yes |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |
| `KmsKeyIdentifier` | `string` | no |
| `DynamicExtensionParameters` | `Map<string>` | no |
| `LatestDeploymentNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `DeploymentStrategyId` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `DeploymentNumber` | `integer` | no |
| `ConfigurationName` | `string` | no |
| `ConfigurationLocationUri` | `string` | no |
| `ConfigurationVersion` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `State` | `string` | no |
| `EventLog` | `List<DeploymentEvent>` | no |
| `PercentageComplete` | `float` | no |
| `StartedAt` | `timestamp` | no |
| `CompletedAt` | `timestamp` | no |
| `AppliedExtensions` | `List<AppliedExtension>` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `VersionLabel` | `string` | no |

## StartExperimentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Description` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `Tags` | `Map<string>` | no |
| `DeploymentParameters` | `DeploymentParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ExperimentDefinitionId` | `string` | no |
| `Run` | `integer` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `Result` | `ExperimentRunResult` | no |
| `StartedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `ExperimentDefinitionSnapshot` | `ExperimentDefinitionSnapshot` | no |

## StopDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |
| `DeploymentNumber` | `integer` | yes |
| `AllowRevert` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `DeploymentStrategyId` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `DeploymentNumber` | `integer` | no |
| `ConfigurationName` | `string` | no |
| `ConfigurationLocationUri` | `string` | no |
| `ConfigurationVersion` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `State` | `string` | no |
| `EventLog` | `List<DeploymentEvent>` | no |
| `PercentageComplete` | `float` | no |
| `StartedAt` | `timestamp` | no |
| `CompletedAt` | `timestamp` | no |
| `AppliedExtensions` | `List<AppliedExtension>` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |
| `VersionLabel` | `string` | no |

## StopExperimentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Run` | `integer` | yes |
| `Result` | `ExperimentRunResult` | no |
| `DeploymentParameters` | `DeploymentParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ExperimentDefinitionId` | `string` | no |
| `Run` | `integer` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `Result` | `ExperimentRunResult` | no |
| `StartedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `ExperimentDefinitionSnapshot` | `ExperimentDefinitionSnapshot` | no |

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


## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionProtection` | `DeletionProtectionSettings` | no |
| `VendedMetrics` | `VendedMetricsSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeletionProtection` | `DeletionProtectionSettings` | no |
| `VendedMetrics` | `VendedMetricsSettings` | no |

## UpdateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |

## UpdateConfigurationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `RetrievalRoleArn` | `string` | no |
| `Validators` | `List<Validator>` | no |
| `KmsKeyIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `LocationUri` | `string` | no |
| `RetrievalRoleArn` | `string` | no |
| `Validators` | `List<Validator>` | no |
| `Type` | `string` | no |
| `KmsKeyArn` | `string` | no |
| `KmsKeyIdentifier` | `string` | no |

## UpdateDeploymentStrategy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeploymentStrategyId` | `string` | yes |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `GrowthFactor` | `float` | no |
| `GrowthType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `DeploymentDurationInMinutes` | `integer` | no |
| `GrowthType` | `string` | no |
| `GrowthFactor` | `float` | no |
| `FinalBakeTimeInMinutes` | `integer` | no |
| `ReplicateTo` | `string` | no |

## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `EnvironmentId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Monitors` | `List<Monitor>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |
| `Monitors` | `List<Monitor>` | no |

## UpdateExperimentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Treatments` | `List<TreatmentInput>` | no |
| `Control` | `TreatmentInput` | no |
| `Hypothesis` | `string` | no |
| `AudienceRule` | `string` | no |
| `AudienceDescription` | `string` | no |
| `LaunchCriteria` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Id` | `string` | no |
| `Name` | `string` | no |
| `Hypothesis` | `string` | no |
| `Status` | `string` | no |
| `ConfigurationProfileId` | `string` | no |
| `EnvironmentId` | `string` | no |
| `FlagKey` | `string` | no |
| `AudienceRule` | `string` | no |
| `AudienceDescription` | `string` | no |
| `LaunchCriteria` | `string` | no |
| `Treatments` | `List<Treatment>` | no |
| `Control` | `Treatment` | no |
| `CreatedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `KmsKeyIdentifier` | `string` | no |

## UpdateExperimentRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ExperimentDefinitionIdentifier` | `string` | yes |
| `Run` | `integer` | yes |
| `Description` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `DeploymentParameters` | `DeploymentParameters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ExperimentDefinitionId` | `string` | no |
| `Run` | `integer` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `ExposurePercentage` | `float` | no |
| `TreatmentOverrides` | `TreatmentOverrides` | no |
| `Result` | `ExperimentRunResult` | no |
| `StartedAt` | `timestamp` | no |
| `UpdatedAt` | `timestamp` | no |
| `EndedAt` | `timestamp` | no |
| `ExperimentDefinitionSnapshot` | `ExperimentDefinitionSnapshot` | no |

## UpdateExtension

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionIdentifier` | `string` | yes |
| `Description` | `string` | no |
| `Actions` | `Map<List<Action>>` | no |
| `Parameters` | `Map<Parameter>` | no |
| `VersionNumber` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Name` | `string` | no |
| `VersionNumber` | `integer` | no |
| `Arn` | `string` | no |
| `Description` | `string` | no |
| `Actions` | `Map<List<Action>>` | no |
| `Parameters` | `Map<Parameter>` | no |

## UpdateExtensionAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtensionAssociationId` | `string` | yes |
| `Parameters` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `ExtensionArn` | `string` | no |
| `ResourceArn` | `string` | no |
| `Arn` | `string` | no |
| `Parameters` | `Map<string>` | no |
| `ExtensionVersionNumber` | `integer` | no |

## ValidateConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationProfileId` | `string` | yes |
| `ConfigurationVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


