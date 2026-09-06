# AWS CloudFormation

API version: 2010-05-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/cloudformation/2010-05-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ActivateType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `PublicTypeArn` | `string` | no |
| `PublisherId` | `string` | no |
| `TypeName` | `string` | no |
| `TypeNameAlias` | `string` | no |
| `AutoUpdate` | `boolean` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `ExecutionRoleArn` | `string` | no |
| `VersionBump` | `string` | no |
| `MajorVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## BatchDescribeTypeConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeConfigurationIdentifiers` | `List<TypeConfigurationIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchDescribeTypeConfigurationsError>` | no |
| `UnprocessedTypeConfigurations` | `List<TypeConfigurationIdentifier>` | no |
| `TypeConfigurations` | `List<TypeConfigurationDetails>` | no |

## CancelUpdateStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ContinueUpdateRollback

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `RoleARN` | `string` | no |
| `ResourcesToSkip` | `List<string>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `UsePreviousTemplate` | `boolean` | no |
| `Parameters` | `List<Parameter>` | no |
| `Capabilities` | `List<string>` | no |
| `ResourceTypes` | `List<string>` | no |
| `RoleARN` | `string` | no |
| `RollbackConfiguration` | `RollbackConfiguration` | no |
| `NotificationARNs` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `ChangeSetName` | `string` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `ChangeSetType` | `string` | no |
| `ResourcesToImport` | `List<ResourceToImport>` | no |
| `IncludeNestedStacks` | `boolean` | no |
| `OnStackFailure` | `string` | no |
| `ImportExistingResources` | `boolean` | no |
| `DeploymentMode` | `string` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |
| `DisableValidation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `StackId` | `string` | no |

## CreateGeneratedTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<ResourceDefinition>` | no |
| `GeneratedTemplateName` | `string` | yes |
| `StackName` | `string` | no |
| `TemplateConfiguration` | `TemplateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateId` | `string` | no |

## CreateStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `Parameters` | `List<Parameter>` | no |
| `DisableRollback` | `boolean` | no |
| `RollbackConfiguration` | `RollbackConfiguration` | no |
| `TimeoutInMinutes` | `integer` | no |
| `NotificationARNs` | `List<string>` | no |
| `Capabilities` | `List<string>` | no |
| `ResourceTypes` | `List<string>` | no |
| `RoleARN` | `string` | no |
| `OnFailure` | `string` | no |
| `StackPolicyBody` | `string` | no |
| `StackPolicyURL` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ClientRequestToken` | `string` | no |
| `EnableTerminationProtection` | `boolean` | no |
| `RetainExceptOnCreate` | `boolean` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |
| `DisableValidation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackId` | `string` | no |
| `OperationId` | `string` | no |

## CreateStackInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `Accounts` | `List<string>` | no |
| `DeploymentTargets` | `DeploymentTargets` | no |
| `Regions` | `List<string>` | yes |
| `ParameterOverrides` | `List<Parameter>` | no |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `OperationId` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## CreateStackRefactor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `EnableStackCreation` | `boolean` | no |
| `ResourceMappings` | `List<ResourceMapping>` | no |
| `StackDefinitions` | `List<StackDefinition>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorId` | `string` | yes |

## CreateStackSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `Description` | `string` | no |
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `StackId` | `string` | no |
| `Parameters` | `List<Parameter>` | no |
| `Capabilities` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `AdministrationRoleARN` | `string` | no |
| `ExecutionRoleName` | `string` | no |
| `PermissionModel` | `string` | no |
| `AutoDeployment` | `AutoDeployment` | no |
| `CallAs` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `ManagedExecution` | `ManagedExecution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetId` | `string` | no |

## DeactivateOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeactivateType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeName` | `string` | no |
| `Type` | `string` | no |
| `Arn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetName` | `string` | yes |
| `StackName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGeneratedTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `RetainResources` | `List<string>` | no |
| `RoleARN` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `DeletionMode` | `string` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStackInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `Accounts` | `List<string>` | no |
| `DeploymentTargets` | `DeploymentTargets` | no |
| `Regions` | `List<string>` | yes |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `RetainStacks` | `boolean` | yes |
| `OperationId` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## DeleteStackSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `VersionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAccountLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountLimits` | `List<AccountLimit>` | no |
| `NextToken` | `string` | no |

## DescribeChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetName` | `string` | yes |
| `StackName` | `string` | no |
| `NextToken` | `string` | no |
| `IncludePropertyValues` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetName` | `string` | no |
| `ChangeSetId` | `string` | no |
| `StackId` | `string` | no |
| `StackName` | `string` | no |
| `Description` | `string` | no |
| `Parameters` | `List<Parameter>` | no |
| `CreationTime` | `timestamp` | no |
| `ExecutionStatus` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `StackDriftStatus` | `string` | no |
| `NotificationARNs` | `List<string>` | no |
| `RollbackConfiguration` | `RollbackConfiguration` | no |
| `Capabilities` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `Changes` | `List<Change>` | no |
| `NextToken` | `string` | no |
| `IncludeNestedStacks` | `boolean` | no |
| `ParentChangeSetId` | `string` | no |
| `RootChangeSetId` | `string` | no |
| `OnStackFailure` | `string` | no |
| `ImportExistingResources` | `boolean` | no |
| `DeploymentMode` | `string` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |

## DescribeChangeSetHooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetName` | `string` | yes |
| `StackName` | `string` | no |
| `NextToken` | `string` | no |
| `LogicalResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetId` | `string` | no |
| `ChangeSetName` | `string` | no |
| `Hooks` | `List<ChangeSetHook>` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `StackId` | `string` | no |
| `StackName` | `string` | no |

## DescribeEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | no |
| `ChangeSetName` | `string` | no |
| `OperationId` | `string` | no |
| `Filters` | `EventFilter` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationEvents` | `List<OperationEvent>` | no |
| `NextToken` | `string` | no |

## DescribeGeneratedTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateId` | `string` | no |
| `GeneratedTemplateName` | `string` | no |
| `Resources` | `List<ResourceDetail>` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Progress` | `TemplateProgress` | no |
| `StackId` | `string` | no |
| `TemplateConfiguration` | `TemplateConfiguration` | no |
| `TotalWarnings` | `integer` | no |

## DescribeOrganizationsAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DescribePublisher

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublisherId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublisherId` | `string` | no |
| `PublisherStatus` | `string` | no |
| `IdentityProvider` | `string` | no |
| `PublisherProfile` | `string` | no |

## DescribeResourceScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanId` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `PercentageCompleted` | `double` | no |
| `ResourceTypes` | `List<string>` | no |
| `ResourcesScanned` | `integer` | no |
| `ResourcesRead` | `integer` | no |
| `ScanFilters` | `List<ScanFilter>` | no |

## DescribeStackDriftDetectionStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackDriftDetectionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackId` | `string` | yes |
| `StackDriftDetectionId` | `string` | yes |
| `StackDriftStatus` | `string` | no |
| `DetectionStatus` | `string` | yes |
| `DetectionStatusReason` | `string` | no |
| `DriftedStackResourceCount` | `integer` | no |
| `Timestamp` | `timestamp` | yes |

## DescribeStackEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackEvents` | `List<StackEvent>` | no |
| `NextToken` | `string` | no |

## DescribeStackInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `StackInstanceAccount` | `string` | yes |
| `StackInstanceRegion` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackInstance` | `StackInstance` | no |

## DescribeStackRefactor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `StackRefactorId` | `string` | no |
| `StackIds` | `List<string>` | no |
| `ExecutionStatus` | `string` | no |
| `ExecutionStatusReason` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |

## DescribeStackResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `LogicalResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackResourceDetail` | `StackResourceDetail` | no |

## DescribeStackResourceDrifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `StackResourceDriftStatusFilters` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackResourceDrifts` | `List<StackResourceDrift>` | yes |
| `NextToken` | `string` | no |

## DescribeStackResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | no |
| `LogicalResourceId` | `string` | no |
| `PhysicalResourceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackResources` | `List<StackResource>` | no |

## DescribeStackSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSet` | `StackSet` | no |

## DescribeStackSetOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `OperationId` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetOperation` | `StackSetOperation` | no |

## DescribeStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Stacks` | `List<Stack>` | no |
| `NextToken` | `string` | no |

## DescribeType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `Arn` | `string` | no |
| `VersionId` | `string` | no |
| `PublisherId` | `string` | no |
| `PublicVersionNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `DefaultVersionId` | `string` | no |
| `IsDefaultVersion` | `boolean` | no |
| `TypeTestsStatus` | `string` | no |
| `TypeTestsStatusDescription` | `string` | no |
| `Description` | `string` | no |
| `Schema` | `string` | no |
| `ProvisioningType` | `string` | no |
| `DeprecatedStatus` | `string` | no |
| `LoggingConfig` | `LoggingConfig` | no |
| `RequiredActivatedTypes` | `List<RequiredActivatedType>` | no |
| `ExecutionRoleArn` | `string` | no |
| `Visibility` | `string` | no |
| `SourceUrl` | `string` | no |
| `DocumentationUrl` | `string` | no |
| `LastUpdated` | `timestamp` | no |
| `TimeCreated` | `timestamp` | no |
| `ConfigurationSchema` | `string` | no |
| `PublisherId` | `string` | no |
| `OriginalTypeName` | `string` | no |
| `OriginalTypeArn` | `string` | no |
| `PublicVersionNumber` | `string` | no |
| `LatestPublicVersion` | `string` | no |
| `IsActivated` | `boolean` | no |
| `AutoUpdate` | `boolean` | no |

## DescribeTypeRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProgressStatus` | `string` | no |
| `Description` | `string` | no |
| `TypeArn` | `string` | no |
| `TypeVersionArn` | `string` | no |

## DetectStackDrift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `LogicalResourceIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackDriftDetectionId` | `string` | yes |

## DetectStackResourceDrift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `LogicalResourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackResourceDrift` | `StackResourceDrift` | yes |

## DetectStackSetDrift

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `OperationId` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## EstimateTemplateCost

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `Parameters` | `List<Parameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |

## ExecuteChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChangeSetName` | `string` | yes |
| `StackName` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `DisableRollback` | `boolean` | no |
| `RetainExceptOnCreate` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExecuteStackRefactor

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetGeneratedTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Format` | `string` | no |
| `GeneratedTemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `TemplateBody` | `string` | no |

## GetHookResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HookResultId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HookResultId` | `string` | no |
| `InvocationPoint` | `string` | no |
| `FailureMode` | `string` | no |
| `TypeName` | `string` | no |
| `OriginalTypeName` | `string` | no |
| `TypeVersionId` | `string` | no |
| `TypeConfigurationVersionId` | `string` | no |
| `TypeArn` | `string` | no |
| `Status` | `string` | no |
| `HookStatusReason` | `string` | no |
| `InvokedAt` | `timestamp` | no |
| `Target` | `HookTarget` | no |
| `Annotations` | `List<Annotation>` | no |

## GetStackPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackPolicyBody` | `string` | no |

## GetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | no |
| `ChangeSetName` | `string` | no |
| `TemplateStage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateBody` | `string` | no |
| `StagesAvailable` | `List<string>` | no |

## GetTemplateSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `StackName` | `string` | no |
| `StackSetName` | `string` | no |
| `CallAs` | `string` | no |
| `TemplateSummaryConfig` | `TemplateSummaryConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<ParameterDeclaration>` | no |
| `Description` | `string` | no |
| `Capabilities` | `List<string>` | no |
| `CapabilitiesReason` | `string` | no |
| `ResourceTypes` | `List<string>` | no |
| `Version` | `string` | no |
| `Metadata` | `string` | no |
| `DeclaredTransforms` | `List<string>` | no |
| `ResourceIdentifierSummaries` | `List<ResourceIdentifierSummary>` | no |
| `Warnings` | `Warnings` | no |

## ImportStacksToStackSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `StackIds` | `List<string>` | no |
| `StackIdsUrl` | `string` | no |
| `OrganizationalUnitIds` | `List<string>` | no |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `OperationId` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## ListChangeSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<ChangeSetSummary>` | no |
| `NextToken` | `string` | no |

## ListExports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Exports` | `List<Export>` | no |
| `NextToken` | `string` | no |

## ListGeneratedTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<TemplateSummary>` | no |
| `NextToken` | `string` | no |

## ListHookResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetType` | `string` | no |
| `TargetId` | `string` | no |
| `TypeArn` | `string` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetType` | `string` | no |
| `TargetId` | `string` | no |
| `HookResults` | `List<HookResultSummary>` | no |
| `NextToken` | `string` | no |

## ListImports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Imports` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListResourceScanRelatedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanId` | `string` | yes |
| `Resources` | `List<ScannedResourceIdentifier>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RelatedResources` | `List<ScannedResource>` | no |
| `NextToken` | `string` | no |

## ListResourceScanResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanId` | `string` | yes |
| `ResourceIdentifier` | `string` | no |
| `ResourceTypePrefix` | `string` | no |
| `TagKey` | `string` | no |
| `TagValue` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resources` | `List<ScannedResource>` | no |
| `NextToken` | `string` | no |

## ListResourceScans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ScanTypeFilter` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanSummaries` | `List<ResourceScanSummary>` | no |
| `NextToken` | `string` | no |

## ListStackInstanceResourceDrifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StackInstanceResourceDriftStatuses` | `List<string>` | no |
| `StackInstanceAccount` | `string` | yes |
| `StackInstanceRegion` | `string` | yes |
| `OperationId` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackInstanceResourceDriftsSummary>` | no |
| `NextToken` | `string` | no |

## ListStackInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<StackInstanceFilter>` | no |
| `StackInstanceAccount` | `string` | no |
| `StackInstanceRegion` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackInstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListStackRefactorActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorActions` | `List<StackRefactorAction>` | yes |
| `NextToken` | `string` | no |

## ListStackRefactors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionStatusFilter` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackRefactorSummaries` | `List<StackRefactorSummary>` | yes |
| `NextToken` | `string` | no |

## ListStackResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackResourceSummaries` | `List<StackResourceSummary>` | no |
| `NextToken` | `string` | no |

## ListStackSetAutoDeploymentTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackSetAutoDeploymentTargetSummary>` | no |
| `NextToken` | `string` | no |

## ListStackSetOperationResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `OperationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CallAs` | `string` | no |
| `Filters` | `List<OperationResultFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackSetOperationResultSummary>` | no |
| `NextToken` | `string` | no |

## ListStackSetOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackSetOperationSummary>` | no |
| `NextToken` | `string` | no |

## ListStackSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<StackSetSummary>` | no |
| `NextToken` | `string` | no |

## ListStacks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StackStatusFilter` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSummaries` | `List<StackSummary>` | no |
| `NextToken` | `string` | no |

## ListTypeRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `TypeArn` | `string` | no |
| `RegistrationStatusFilter` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationTokenList` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListTypeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `Arn` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `DeprecatedStatus` | `string` | no |
| `PublisherId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeVersionSummaries` | `List<TypeVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Visibility` | `string` | no |
| `ProvisioningType` | `string` | no |
| `DeprecatedStatus` | `string` | no |
| `Type` | `string` | no |
| `Filters` | `TypeFilters` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeSummaries` | `List<TypeSummary>` | no |
| `NextToken` | `string` | no |

## PublishType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `Arn` | `string` | no |
| `TypeName` | `string` | no |
| `PublicVersionNumber` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublicTypeArn` | `string` | no |

## RecordHandlerProgress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BearerToken` | `string` | yes |
| `OperationStatus` | `string` | yes |
| `CurrentOperationStatus` | `string` | no |
| `StatusMessage` | `string` | no |
| `ErrorCode` | `string` | no |
| `ResourceModel` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterPublisher

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AcceptTermsAndConditions` | `boolean` | no |
| `ConnectionArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PublisherId` | `string` | no |

## RegisterType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Type` | `string` | no |
| `TypeName` | `string` | yes |
| `SchemaHandlerPackage` | `string` | yes |
| `LoggingConfig` | `LoggingConfig` | no |
| `ExecutionRoleArn` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationToken` | `string` | no |

## RollbackStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `RoleARN` | `string` | no |
| `ClientRequestToken` | `string` | no |
| `RetainExceptOnCreate` | `boolean` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackId` | `string` | no |
| `OperationId` | `string` | no |

## SetStackPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `StackPolicyBody` | `string` | no |
| `StackPolicyURL` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetTypeConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeArn` | `string` | no |
| `Configuration` | `string` | yes |
| `ConfigurationAlias` | `string` | no |
| `TypeName` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationArn` | `string` | no |

## SetTypeDefaultVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `VersionId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SignalResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `LogicalResourceId` | `string` | yes |
| `UniqueId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartResourceScan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientRequestToken` | `string` | no |
| `ScanFilters` | `List<ScanFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceScanId` | `string` | no |

## StopStackSetOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `OperationId` | `string` | yes |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Type` | `string` | no |
| `TypeName` | `string` | no |
| `VersionId` | `string` | no |
| `LogDeliveryBucket` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TypeVersionArn` | `string` | no |

## UpdateGeneratedTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateName` | `string` | yes |
| `NewGeneratedTemplateName` | `string` | no |
| `AddResources` | `List<ResourceDefinition>` | no |
| `RemoveResources` | `List<string>` | no |
| `RefreshAllResources` | `boolean` | no |
| `TemplateConfiguration` | `TemplateConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GeneratedTemplateId` | `string` | no |

## UpdateStack

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackName` | `string` | yes |
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `UsePreviousTemplate` | `boolean` | no |
| `StackPolicyDuringUpdateBody` | `string` | no |
| `StackPolicyDuringUpdateURL` | `string` | no |
| `Parameters` | `List<Parameter>` | no |
| `Capabilities` | `List<string>` | no |
| `ResourceTypes` | `List<string>` | no |
| `RoleARN` | `string` | no |
| `RollbackConfiguration` | `RollbackConfiguration` | no |
| `StackPolicyBody` | `string` | no |
| `StackPolicyURL` | `string` | no |
| `NotificationARNs` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `DisableRollback` | `boolean` | no |
| `ClientRequestToken` | `string` | no |
| `RetainExceptOnCreate` | `boolean` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |
| `DisableValidation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackId` | `string` | no |
| `OperationId` | `string` | no |

## UpdateStackInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `Accounts` | `List<string>` | no |
| `DeploymentTargets` | `DeploymentTargets` | no |
| `Regions` | `List<string>` | yes |
| `ParameterOverrides` | `List<Parameter>` | no |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `OperationId` | `string` | no |
| `CallAs` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateStackSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackSetName` | `string` | yes |
| `Description` | `string` | no |
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |
| `UsePreviousTemplate` | `boolean` | no |
| `Parameters` | `List<Parameter>` | no |
| `Capabilities` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |
| `OperationPreferences` | `StackSetOperationPreferences` | no |
| `AdministrationRoleARN` | `string` | no |
| `ExecutionRoleName` | `string` | no |
| `DeploymentTargets` | `DeploymentTargets` | no |
| `PermissionModel` | `string` | no |
| `AutoDeployment` | `AutoDeployment` | no |
| `OperationId` | `string` | no |
| `Accounts` | `List<string>` | no |
| `Regions` | `List<string>` | no |
| `CallAs` | `string` | no |
| `ManagedExecution` | `ManagedExecution` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## UpdateTerminationProtection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnableTerminationProtection` | `boolean` | yes |
| `StackName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StackId` | `string` | no |

## ValidateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateBody` | `string` | no |
| `TemplateURL` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Parameters` | `List<TemplateParameter>` | no |
| `Description` | `string` | no |
| `Capabilities` | `List<string>` | no |
| `CapabilitiesReason` | `string` | no |
| `DeclaredTransforms` | `List<string>` | no |

