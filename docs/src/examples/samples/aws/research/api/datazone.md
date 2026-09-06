# Amazon DataZone

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/datazone/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptPredictions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |
| `acceptRule` | `AcceptRule` | no |
| `acceptChoices` | `List<AcceptChoice>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `assetId` | `string` | yes |
| `revision` | `string` | yes |

## AcceptSubscriptionRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `decisionComment` | `string` | no |
| `assetScopes` | `List<AcceptedAssetScope>` | no |
| `assetPermissions` | `List<AssetPermission>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `requestReason` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipal>` | yes |
| `subscribedListings` | `List<SubscribedListing>` | yes |
| `reviewerId` | `string` | no |
| `decisionComment` | `string` | no |
| `existingSubscriptionId` | `string` | no |
| `metadataForms` | `List<FormOutput>` | no |

## AddEntityOwner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `owner` | `OwnerProperties` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AddPolicyGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `policyType` | `string` | yes |
| `principal` | `PolicyGrantPrincipal` | yes |
| `detail` | `PolicyGrantDetail` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `grantId` | `string` | no |

## AssociateEnvironmentRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `environmentRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateGovernedTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `governedGlossaryTerms` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetAttributesMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityRevision` | `string` | no |
| `attributeIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `List<BatchGetAttributeOutput>` | no |
| `errors` | `List<AttributeError>` | yes |

## BatchPutAttributesMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `attributes` | `List<AttributeInput>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<AttributeError>` | no |
| `attributes` | `List<BatchPutAttributeOutput>` | no |

## CancelMetadataGenerationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CancelSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `subscribedPrincipal` | `SubscribedPrincipal` | yes |
| `subscribedListing` | `SubscribedListing` | yes |
| `subscriptionRequestId` | `string` | no |
| `retainPermissions` | `boolean` | no |

## CreateAccountPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `resolutionStrategy` | `string` | yes |
| `accountSource` | `AccountSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `name` | `string` | no |
| `id` | `string` | no |
| `description` | `string` | no |
| `resolutionStrategy` | `string` | no |
| `accountSource` | `AccountSource` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `domainUnitId` | `string` | no |

## CreateAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `domainIdentifier` | `string` | yes |
| `externalIdentifier` | `string` | no |
| `typeIdentifier` | `string` | yes |
| `typeRevision` | `string` | no |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `formsInput` | `List<FormInput>` | no |
| `owningProjectIdentifier` | `string` | yes |
| `predictionConfiguration` | `PredictionConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `typeIdentifier` | `string` | yes |
| `typeRevision` | `string` | yes |
| `externalIdentifier` | `string` | no |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `governedGlossaryTerms` | `List<string>` | no |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `listing` | `AssetListingDetails` | no |
| `formsOutput` | `List<FormOutput>` | yes |
| `readOnlyFormsOutput` | `List<FormOutput>` | no |
| `latestTimeSeriesDataPointFormsOutput` | `List<TimeSeriesDataPointSummaryFormOutput>` | no |
| `predictionConfiguration` | `PredictionConfiguration` | no |

## CreateAssetFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `configuration` | `AssetFilterConfiguration` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `assetId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `configuration` | `AssetFilterConfiguration` | yes |
| `createdAt` | `timestamp` | no |
| `errorMessage` | `string` | no |
| `effectiveColumnNames` | `List<string>` | no |
| `effectiveRowFilter` | `string` | no |

## CreateAssetRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `typeRevision` | `string` | no |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `formsInput` | `List<FormInput>` | no |
| `predictionConfiguration` | `PredictionConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `typeIdentifier` | `string` | yes |
| `typeRevision` | `string` | yes |
| `externalIdentifier` | `string` | no |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `governedGlossaryTerms` | `List<string>` | no |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `listing` | `AssetListingDetails` | no |
| `formsOutput` | `List<FormOutput>` | yes |
| `readOnlyFormsOutput` | `List<FormOutput>` | no |
| `latestTimeSeriesDataPointFormsOutput` | `List<TimeSeriesDataPointSummaryFormOutput>` | no |
| `predictionConfiguration` | `PredictionConfiguration` | no |

## CreateAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `formsInput` | `Map<FormEntryInput>` | yes |
| `owningProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `formsOutput` | `Map<FormEntryOutput>` | yes |
| `owningProjectId` | `string` | no |
| `originDomainId` | `string` | no |
| `originProjectId` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## CreateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `awsLocation` | `AwsLocation` | no |
| `clientToken` | `string` | no |
| `configurations` | `List<Configuration>` | no |
| `description` | `string` | no |
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | no |
| `name` | `string` | yes |
| `props` | `ConnectionPropertiesInput` | no |
| `enableTrustedIdentityPropagation` | `boolean` | no |
| `scope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionId` | `string` | yes |
| `configurations` | `List<Configuration>` | no |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `domainUnitId` | `string` | yes |
| `environmentId` | `string` | no |
| `name` | `string` | yes |
| `physicalEndpoints` | `List<PhysicalEndpoint>` | yes |
| `projectId` | `string` | no |
| `props` | `ConnectionPropertiesOutput` | no |
| `type` | `string` | yes |
| `scope` | `string` | no |

## CreateDataProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `formsInput` | `List<FormInput>` | no |
| `items` | `List<DataProductItem>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `revision` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `items` | `List<DataProductItem>` | no |
| `formsOutput` | `List<FormOutput>` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |

## CreateDataProductRevision

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `items` | `List<DataProductItem>` | no |
| `formsInput` | `List<FormInput>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `revision` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `items` | `List<DataProductItem>` | no |
| `formsOutput` | `List<FormOutput>` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `domainIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | no |
| `connectionIdentifier` | `string` | no |
| `type` | `string` | yes |
| `configuration` | `DataSourceConfigurationInput` | no |
| `recommendation` | `RecommendationConfiguration` | no |
| `enableSetting` | `string` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsInput` | `List<FormInput>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | no |
| `connectionId` | `string` | no |
| `configuration` | `DataSourceConfigurationOutput` | no |
| `recommendation` | `RecommendationConfiguration` | no |
| `enableSetting` | `string` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsOutput` | `List<FormOutput>` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `lastRunStatus` | `string` | no |
| `lastRunAt` | `timestamp` | no |
| `lastRunErrorMessage` | `DataSourceErrorMessage` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `singleSignOn` | `SingleSignOn` | no |
| `domainExecutionRole` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |
| `tags` | `Map<string>` | no |
| `domainVersion` | `string` | no |
| `serviceRole` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `rootDomainUnitId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `singleSignOn` | `SingleSignOn` | no |
| `domainExecutionRole` | `string` | no |
| `arn` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |
| `status` | `string` | no |
| `portalUrl` | `string` | no |
| `tags` | `Map<string>` | no |
| `domainVersion` | `string` | no |
| `serviceRole` | `string` | no |

## CreateDomainUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `parentDomainUnitIdentifier` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `parentDomainUnitId` | `string` | no |
| `description` | `string` | no |
| `owners` | `List<DomainUnitOwnerProperties>` | yes |
| `ancestorDomainUnitIds` | `List<string>` | yes |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectIdentifier` | `string` | yes |
| `domainIdentifier` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | yes |
| `environmentProfileIdentifier` | `string` | no |
| `userParameters` | `List<EnvironmentParameter>` | no |
| `glossaryTerms` | `List<string>` | no |
| `environmentAccountIdentifier` | `string` | no |
| `environmentAccountRegion` | `string` | no |
| `environmentBlueprintIdentifier` | `string` | no |
| `deploymentOrder` | `integer` | no |
| `environmentConfigurationId` | `string` | no |
| `environmentConfigurationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `id` | `string` | no |
| `domainId` | `string` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentProfileId` | `string` | no |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `provider` | `string` | yes |
| `provisionedResources` | `List<Resource>` | no |
| `status` | `string` | no |
| `environmentActions` | `List<ConfigurableEnvironmentAction>` | no |
| `glossaryTerms` | `List<string>` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `lastDeployment` | `Deployment` | no |
| `provisioningProperties` | `ProvisioningProperties` | no |
| `deploymentProperties` | `DeploymentProperties` | no |
| `environmentBlueprintId` | `string` | no |
| `environmentConfigurationId` | `string` | no |
| `environmentConfigurationName` | `string` | no |

## CreateEnvironmentAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `parameters` | `ActionParameters` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `environmentId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `parameters` | `ActionParameters` | yes |
| `description` | `string` | no |

## CreateEnvironmentBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `provisioningProperties` | `ProvisioningProperties` | yes |
| `userParameters` | `List<CustomParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `provider` | `string` | yes |
| `provisioningProperties` | `ProvisioningProperties` | yes |
| `deploymentProperties` | `DeploymentProperties` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `glossaryTerms` | `List<string>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## CreateEnvironmentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentBlueprintIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `userParameters` | `List<EnvironmentParameter>` | no |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentBlueprintId` | `string` | yes |
| `projectId` | `string` | no |
| `userParameters` | `List<CustomParameter>` | no |

## CreateFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `model` | `Model` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `status` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `owningProjectId` | `string` | no |
| `originDomainId` | `string` | no |
| `originProjectId` | `string` | no |

## CreateGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `usageRestrictions` | `List<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `usageRestrictions` | `List<string>` | no |

## CreateGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `glossaryIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | no |
| `shortDescription` | `string` | no |
| `longDescription` | `string` | no |
| `termRelations` | `TermRelations` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `glossaryId` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `shortDescription` | `string` | no |
| `longDescription` | `string` | no |
| `termRelations` | `TermRelations` | no |
| `usageRestrictions` | `List<string>` | no |

## CreateGroupProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `groupIdentifier` | `string` | no |
| `rolePrincipalArn` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `groupName` | `string` | no |
| `rolePrincipalArn` | `string` | no |
| `rolePrincipalId` | `string` | no |

## CreateListingChangeSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityRevision` | `string` | no |
| `action` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `listingId` | `string` | yes |
| `listingRevision` | `string` | yes |
| `status` | `string` | yes |

## CreateNotebook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `cellOrder` | `List<CellInformation>` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `lockedBy` | `string` | no |
| `lockedAt` | `timestamp` | no |
| `lockExpiresAt` | `timestamp` | no |
| `computeId` | `string` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `error` | `NotebookError` | no |
| `gitMetadata` | `GitMetadata` | no |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `resourceTags` | `Map<string>` | no |
| `glossaryTerms` | `List<string>` | no |
| `domainUnitId` | `string` | no |
| `projectProfileId` | `string` | no |
| `userParameters` | `List<EnvironmentConfigurationUserParameter>` | no |
| `projectCategory` | `string` | no |
| `projectExecutionRole` | `string` | no |
| `membershipAssignments` | `List<ProjectMembershipAssignment>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `projectStatus` | `string` | no |
| `failureReasons` | `List<ProjectDeletionError>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `resourceTags` | `List<ResourceTag>` | no |
| `glossaryTerms` | `List<string>` | no |
| `domainUnitId` | `string` | no |
| `projectProfileId` | `string` | no |
| `userParameters` | `List<EnvironmentConfigurationUserParameter>` | no |
| `environmentDeploymentDetails` | `EnvironmentDeploymentDetails` | no |
| `projectCategory` | `string` | no |

## CreateProjectMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `member` | `Member` | yes |
| `designation` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateProjectProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `projectResourceTags` | `List<ResourceTagParameter>` | no |
| `allowCustomProjectResourceTags` | `boolean` | no |
| `projectResourceTagsDescription` | `string` | no |
| `environmentConfigurations` | `List<EnvironmentConfiguration>` | no |
| `domainUnitIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `projectResourceTags` | `List<ResourceTagParameter>` | no |
| `allowCustomProjectResourceTags` | `boolean` | no |
| `projectResourceTagsDescription` | `string` | no |
| `environmentConfigurations` | `List<EnvironmentConfiguration>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `domainUnitId` | `string` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `target` | `RuleTarget` | yes |
| `action` | `string` | yes |
| `scope` | `RuleScope` | yes |
| `detail` | `RuleDetail` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `name` | `string` | yes |
| `ruleType` | `string` | yes |
| `target` | `RuleTarget` | yes |
| `action` | `string` | yes |
| `scope` | `RuleScope` | yes |
| `detail` | `RuleDetail` | yes |
| `targetType` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |

## CreateSubscriptionGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `subscriptionTargetIdentifier` | `string` | no |
| `grantedEntity` | `GrantedEntityInput` | yes |
| `assetTargetNames` | `List<AssetTargetNameMap>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `environmentId` | `string` | no |
| `subscriptionTargetId` | `string` | yes |
| `grantedEntity` | `GrantedEntity` | yes |
| `status` | `string` | yes |
| `assets` | `List<SubscribedAsset>` | no |
| `subscriptionId` | `string` | no |

## CreateSubscriptionRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipalInput>` | yes |
| `subscribedListings` | `List<SubscribedListingInput>` | yes |
| `requestReason` | `string` | yes |
| `clientToken` | `string` | no |
| `metadataForms` | `List<FormInput>` | no |
| `assetPermissions` | `List<AssetPermission>` | no |
| `assetScopes` | `List<AcceptedAssetScope>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `requestReason` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipal>` | yes |
| `subscribedListings` | `List<SubscribedListing>` | yes |
| `reviewerId` | `string` | no |
| `decisionComment` | `string` | no |
| `existingSubscriptionId` | `string` | no |
| `metadataForms` | `List<FormOutput>` | no |

## CreateSubscriptionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `subscriptionTargetConfig` | `List<SubscriptionTargetForm>` | yes |
| `authorizedPrincipals` | `List<string>` | yes |
| `manageAccessRole` | `string` | yes |
| `applicableAssetTypes` | `List<string>` | yes |
| `provider` | `string` | no |
| `clientToken` | `string` | no |
| `subscriptionGrantCreationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `authorizedPrincipals` | `List<string>` | yes |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `manageAccessRole` | `string` | no |
| `applicableAssetTypes` | `List<string>` | yes |
| `subscriptionTargetConfig` | `List<SubscriptionTargetForm>` | yes |
| `provider` | `string` | yes |
| `subscriptionGrantCreationMode` | `string` | no |

## CreateUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `userIdentifier` | `string` | yes |
| `userType` | `string` | no |
| `sessionName` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `details` | `UserProfileDetails` | no |

## DeleteAccountPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssetFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## DeleteDataExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |
| `retainPermissionsOnRevokeFailure` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | no |
| `connectionId` | `string` | no |
| `configuration` | `DataSourceConfigurationOutput` | no |
| `enableSetting` | `string` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsOutput` | `List<FormOutput>` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `lastRunStatus` | `string` | no |
| `lastRunAt` | `timestamp` | no |
| `lastRunErrorMessage` | `DataSourceErrorMessage` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `selfGrantStatus` | `SelfGrantStatusOutput` | no |
| `retainPermissionsOnRevokeFailure` | `boolean` | no |

## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |
| `skipDeletionCheck` | `boolean` | no |
| `cascadeDelete` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteDomainUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentBlueprintConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentBlueprintIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnvironmentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `formTypeIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteLineageEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `domainId` | `string` | no |
| `processingStatus` | `string` | no |

## DeleteListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotebook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `skipDeletionCheck` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProjectMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `member` | `Member` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProjectProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriptionGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `environmentId` | `string` | no |
| `subscriptionTargetId` | `string` | yes |
| `grantedEntity` | `GrantedEntity` | yes |
| `status` | `string` | yes |
| `assets` | `List<SubscribedAsset>` | no |
| `subscriptionId` | `string` | no |

## DeleteSubscriptionRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriptionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTimeSeriesDataPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `formName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateEnvironmentRole

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `environmentRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateGovernedTerms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `governedGlossaryTerms` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `name` | `string` | no |
| `id` | `string` | no |
| `description` | `string` | no |
| `resolutionStrategy` | `string` | no |
| `accountSource` | `AccountSource` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `domainUnitId` | `string` | no |

## GetAsset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `typeIdentifier` | `string` | yes |
| `typeRevision` | `string` | yes |
| `externalIdentifier` | `string` | no |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `governedGlossaryTerms` | `List<string>` | no |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `listing` | `AssetListingDetails` | no |
| `formsOutput` | `List<FormOutput>` | yes |
| `readOnlyFormsOutput` | `List<FormOutput>` | no |
| `latestTimeSeriesDataPointFormsOutput` | `List<TimeSeriesDataPointSummaryFormOutput>` | no |

## GetAssetFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `assetId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `configuration` | `AssetFilterConfiguration` | yes |
| `createdAt` | `timestamp` | no |
| `errorMessage` | `string` | no |
| `effectiveColumnNames` | `List<string>` | no |
| `effectiveRowFilter` | `string` | no |

## GetAssetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `revision` | `string` | yes |
| `description` | `string` | no |
| `formsOutput` | `Map<FormEntryOutput>` | yes |
| `owningProjectId` | `string` | yes |
| `originDomainId` | `string` | no |
| `originProjectId` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## GetConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `withSecret` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionCredentials` | `ConnectionCredentials` | no |
| `configurations` | `List<Configuration>` | no |
| `connectionId` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `domainUnitId` | `string` | yes |
| `environmentId` | `string` | no |
| `environmentUserRole` | `string` | no |
| `name` | `string` | yes |
| `physicalEndpoints` | `List<PhysicalEndpoint>` | yes |
| `projectId` | `string` | no |
| `props` | `ConnectionPropertiesOutput` | no |
| `type` | `string` | yes |
| `scope` | `string` | no |

## GetDataExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isExportEnabled` | `boolean` | no |
| `status` | `string` | no |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `s3TableBucketArn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetDataProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `revision` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `items` | `List<DataProductItem>` | no |
| `formsOutput` | `List<FormOutput>` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `firstRevisionCreatedAt` | `timestamp` | no |
| `firstRevisionCreatedBy` | `string` | no |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | no |
| `connectionId` | `string` | no |
| `configuration` | `DataSourceConfigurationOutput` | no |
| `recommendation` | `RecommendationConfiguration` | no |
| `enableSetting` | `string` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsOutput` | `List<FormOutput>` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `lastRunStatus` | `string` | no |
| `lastRunAt` | `timestamp` | no |
| `lastRunErrorMessage` | `DataSourceErrorMessage` | no |
| `lastRunAssetCount` | `integer` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `selfGrantStatus` | `SelfGrantStatusOutput` | no |

## GetDataSourceRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `id` | `string` | yes |
| `projectId` | `string` | yes |
| `status` | `string` | yes |
| `type` | `string` | yes |
| `dataSourceConfigurationSnapshot` | `string` | no |
| `runStatisticsForAssets` | `RunStatisticsForAssets` | no |
| `lineageSummary` | `DataSourceRunLineageSummary` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `startedAt` | `timestamp` | no |
| `stoppedAt` | `timestamp` | no |

## GetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `rootDomainUnitId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `singleSignOn` | `SingleSignOn` | no |
| `domainExecutionRole` | `string` | yes |
| `arn` | `string` | no |
| `kmsKeyIdentifier` | `string` | no |
| `status` | `string` | yes |
| `portalUrl` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `domainVersion` | `string` | no |
| `serviceRole` | `string` | no |
| `failureReasons` | `List<FailureReason>` | no |
| `deleteProgress` | `DeleteProgress` | no |

## GetDomainUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `parentDomainUnitId` | `string` | no |
| `description` | `string` | no |
| `owners` | `List<DomainUnitOwnerProperties>` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `lastUpdatedBy` | `string` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `id` | `string` | no |
| `domainId` | `string` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentProfileId` | `string` | no |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `provider` | `string` | yes |
| `provisionedResources` | `List<Resource>` | no |
| `status` | `string` | no |
| `environmentActions` | `List<ConfigurableEnvironmentAction>` | no |
| `glossaryTerms` | `List<string>` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `lastDeployment` | `Deployment` | no |
| `provisioningProperties` | `ProvisioningProperties` | no |
| `deploymentProperties` | `DeploymentProperties` | no |
| `environmentBlueprintId` | `string` | no |
| `environmentConfigurationId` | `string` | no |
| `environmentConfigurationName` | `string` | no |

## GetEnvironmentAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `environmentId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `parameters` | `ActionParameters` | yes |
| `description` | `string` | no |

## GetEnvironmentBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `provider` | `string` | yes |
| `provisioningProperties` | `ProvisioningProperties` | yes |
| `deploymentProperties` | `DeploymentProperties` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `glossaryTerms` | `List<string>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetEnvironmentBlueprintConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentBlueprintIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `environmentBlueprintId` | `string` | yes |
| `provisioningRoleArn` | `string` | no |
| `environmentRolePermissionBoundary` | `string` | no |
| `manageAccessRoleArn` | `string` | no |
| `enabledRegions` | `List<string>` | no |
| `regionalParameters` | `Map<Map<string>>` | no |
| `allowUserProvidedConfigurations` | `boolean` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `resourceConfigurations` | `List<ResourceConfiguration>` | no |
| `provisioningConfigurations` | `List<ProvisioningConfiguration>` | no |

## GetEnvironmentCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessKeyId` | `string` | no |
| `secretAccessKey` | `string` | no |
| `sessionToken` | `string` | no |
| `expiration` | `timestamp` | no |

## GetEnvironmentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentBlueprintId` | `string` | yes |
| `projectId` | `string` | no |
| `userParameters` | `List<CustomParameter>` | no |

## GetFormType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `formTypeIdentifier` | `string` | yes |
| `revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `revision` | `string` | yes |
| `model` | `Model` | yes |
| `owningProjectId` | `string` | no |
| `originDomainId` | `string` | no |
| `originProjectId` | `string` | no |
| `status` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `description` | `string` | no |
| `imports` | `List<Import>` | no |

## GetGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `usageRestrictions` | `List<string>` | no |

## GetGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `glossaryId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `shortDescription` | `string` | no |
| `longDescription` | `string` | no |
| `termRelations` | `TermRelations` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `usageRestrictions` | `List<string>` | no |

## GetGroupProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `groupIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `groupName` | `string` | no |
| `rolePrincipalArn` | `string` | no |
| `rolePrincipalId` | `string` | no |

## GetIamPortalLoginUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `authCodeUrl` | `string` | no |
| `userProfileId` | `string` | yes |

## GetJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `jobId` | `string` | no |
| `jobType` | `string` | no |
| `runMode` | `string` | no |
| `details` | `JobRunDetails` | no |
| `status` | `string` | no |
| `error` | `JobRunError` | no |
| `createdBy` | `string` | no |
| `createdAt` | `timestamp` | no |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |

## GetLineageEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `event` | `blob` | no |
| `createdBy` | `string` | no |
| `processingStatus` | `string` | no |
| `eventTime` | `timestamp` | no |
| `createdAt` | `timestamp` | no |

## GetLineageNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `eventTimestamp` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `id` | `string` | yes |
| `typeName` | `string` | yes |
| `typeRevision` | `string` | no |
| `sourceIdentifier` | `string` | no |
| `eventTimestamp` | `timestamp` | no |
| `formsOutput` | `List<FormOutput>` | no |
| `upstreamNodes` | `List<LineageNodeReference>` | no |
| `downstreamNodes` | `List<LineageNodeReference>` | no |

## GetListing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `listingRevision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `listingRevision` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedBy` | `string` | no |
| `item` | `ListingItem` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |

## GetMetadataGenerationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `target` | `MetadataGenerationRunTarget` | no |
| `status` | `string` | no |
| `type` | `string` | no |
| `types` | `List<string>` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `owningProjectId` | `string` | yes |
| `typeStats` | `List<MetadataGenerationRunTypeStat>` | no |

## GetNotebook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `cellOrder` | `List<CellInformation>` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `lockedBy` | `string` | no |
| `lockedAt` | `timestamp` | no |
| `lockExpiresAt` | `timestamp` | no |
| `computeId` | `string` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `error` | `NotebookError` | no |
| `gitMetadata` | `GitMetadata` | no |

## GetNotebookExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `notebookId` | `string` | yes |
| `fileFormat` | `string` | yes |
| `status` | `string` | yes |
| `outputLocation` | `OutputLocation` | no |
| `error` | `NotebookExportError` | no |
| `completedAt` | `timestamp` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |

## GetNotebookRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `notebookId` | `string` | yes |
| `scheduleId` | `string` | no |
| `status` | `string` | yes |
| `cellOrder` | `List<CellInformation>` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `computeConfiguration` | `ComputeConfig` | no |
| `networkConfiguration` | `NetworkConfig` | no |
| `timeoutConfiguration` | `TimeoutConfig` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `storageConfiguration` | `StorageConfig` | no |
| `triggerSource` | `TriggerSource` | no |
| `error` | `NotebookRunError` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `startedAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |

## GetProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `projectStatus` | `string` | no |
| `failureReasons` | `List<ProjectDeletionError>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `resourceTags` | `List<ResourceTag>` | no |
| `glossaryTerms` | `List<string>` | no |
| `domainUnitId` | `string` | no |
| `projectProfileId` | `string` | no |
| `userParameters` | `List<EnvironmentConfigurationUserParameter>` | no |
| `environmentDeploymentDetails` | `EnvironmentDeploymentDetails` | no |
| `projectCategory` | `string` | no |

## GetProjectProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `projectResourceTags` | `List<ResourceTagParameter>` | no |
| `allowCustomProjectResourceTags` | `boolean` | no |
| `projectResourceTagsDescription` | `string` | no |
| `environmentConfigurations` | `List<EnvironmentConfiguration>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `domainUnitId` | `string` | no |

## GetRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `revision` | `string` | yes |
| `name` | `string` | yes |
| `ruleType` | `string` | yes |
| `target` | `RuleTarget` | yes |
| `action` | `string` | yes |
| `scope` | `RuleScope` | yes |
| `detail` | `RuleDetail` | yes |
| `targetType` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `lastUpdatedBy` | `string` | yes |

## GetSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `subscribedPrincipal` | `SubscribedPrincipal` | yes |
| `subscribedListing` | `SubscribedListing` | yes |
| `subscriptionRequestId` | `string` | no |
| `retainPermissions` | `boolean` | no |

## GetSubscriptionGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `environmentId` | `string` | no |
| `subscriptionTargetId` | `string` | yes |
| `grantedEntity` | `GrantedEntity` | yes |
| `status` | `string` | yes |
| `assets` | `List<SubscribedAsset>` | no |
| `subscriptionId` | `string` | no |

## GetSubscriptionRequestDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `requestReason` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipal>` | yes |
| `subscribedListings` | `List<SubscribedListing>` | yes |
| `reviewerId` | `string` | no |
| `decisionComment` | `string` | no |
| `existingSubscriptionId` | `string` | no |
| `metadataForms` | `List<FormOutput>` | no |

## GetSubscriptionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `authorizedPrincipals` | `List<string>` | yes |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `manageAccessRole` | `string` | no |
| `applicableAssetTypes` | `List<string>` | yes |
| `subscriptionTargetConfig` | `List<SubscriptionTargetForm>` | yes |
| `provider` | `string` | yes |
| `subscriptionGrantCreationMode` | `string` | no |

## GetTimeSeriesDataPoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `identifier` | `string` | yes |
| `formName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `entityId` | `string` | no |
| `entityType` | `string` | no |
| `formName` | `string` | no |
| `form` | `TimeSeriesDataPointFormOutput` | no |

## GetUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `userIdentifier` | `string` | yes |
| `type` | `string` | no |
| `sessionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `details` | `UserProfileDetails` | no |

## ListAccountPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AccountPoolSummary>` | no |
| `nextToken` | `string` | no |

## ListAccountsInAccountPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AccountInfo>` | no |
| `nextToken` | `string` | no |

## ListAssetFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AssetFilterSummary>` | yes |
| `nextToken` | `string` | no |

## ListAssetRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<AssetRevision>` | no |
| `nextToken` | `string` | no |

## ListConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `name` | `string` | no |
| `environmentIdentifier` | `string` | no |
| `projectIdentifier` | `string` | no |
| `type` | `string` | no |
| `scope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ConnectionSummary>` | yes |
| `nextToken` | `string` | no |

## ListDataProductRevisions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DataProductRevision>` | yes |
| `nextToken` | `string` | no |

## ListDataSourceRunActivities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DataSourceRunActivity>` | yes |
| `nextToken` | `string` | no |

## ListDataSourceRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `dataSourceIdentifier` | `string` | yes |
| `status` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DataSourceRunSummary>` | yes |
| `nextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | no |
| `connectionIdentifier` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DataSourceSummary>` | yes |
| `nextToken` | `string` | no |

## ListDomainUnitsForParent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `parentDomainUnitIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DomainUnitSummary>` | yes |
| `nextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<DomainSummary>` | yes |
| `nextToken` | `string` | no |

## ListEntityOwners

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owners` | `List<OwnerPropertiesOutput>` | yes |
| `nextToken` | `string` | no |

## ListEnvironmentActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<EnvironmentActionSummary>` | no |
| `nextToken` | `string` | no |

## ListEnvironmentBlueprintConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<EnvironmentBlueprintConfigurationItem>` | no |
| `nextToken` | `string` | no |

## ListEnvironmentBlueprints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `name` | `string` | no |
| `managed` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<EnvironmentBlueprintSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnvironmentProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `environmentBlueprintIdentifier` | `string` | no |
| `projectIdentifier` | `string` | no |
| `name` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<EnvironmentProfileSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `awsAccountId` | `string` | no |
| `status` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `projectIdentifier` | `string` | yes |
| `environmentProfileIdentifier` | `string` | no |
| `environmentBlueprintIdentifier` | `string` | no |
| `provider` | `string` | no |
| `name` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<EnvironmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListJobRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `jobIdentifier` | `string` | yes |
| `status` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<JobRunSummary>` | no |
| `nextToken` | `string` | no |

## ListLineageEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `timestampAfter` | `timestamp` | no |
| `timestampBefore` | `timestamp` | no |
| `processingStatus` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<LineageEventSummary>` | no |
| `nextToken` | `string` | no |

## ListLineageNodeHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `identifier` | `string` | yes |
| `direction` | `string` | no |
| `eventTimestampGTE` | `timestamp` | no |
| `eventTimestampLTE` | `timestamp` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nodes` | `List<LineageNodeSummary>` | no |
| `nextToken` | `string` | no |

## ListMetadataGenerationRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `targetIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<MetadataGenerationRunItem>` | no |
| `nextToken` | `string` | no |

## ListNotebookRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `notebookIdentifier` | `string` | no |
| `status` | `string` | no |
| `scheduleIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NotebookRunSummary>` | no |
| `nextToken` | `string` | no |

## ListNotebooks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `sortOrder` | `string` | no |
| `sortBy` | `string` | no |
| `status` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<NotebookSummary>` | no |
| `nextToken` | `string` | no |

## ListNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `type` | `string` | yes |
| `afterTimestamp` | `timestamp` | no |
| `beforeTimestamp` | `timestamp` | no |
| `subjects` | `List<string>` | no |
| `taskStatus` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notifications` | `List<NotificationOutput>` | no |
| `nextToken` | `string` | no |

## ListPolicyGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `policyType` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `grantList` | `List<PolicyGrantMember>` | yes |
| `nextToken` | `string` | no |

## ListProjectMemberships

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `projectIdentifier` | `string` | yes |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `members` | `List<ProjectMember>` | yes |
| `nextToken` | `string` | no |

## ListProjectProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `name` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ProjectProfileSummary>` | no |
| `nextToken` | `string` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `userIdentifier` | `string` | no |
| `groupIdentifier` | `string` | no |
| `name` | `string` | no |
| `projectCategory` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ProjectSummary>` | no |
| `nextToken` | `string` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `targetType` | `string` | yes |
| `targetIdentifier` | `string` | yes |
| `ruleType` | `string` | no |
| `action` | `string` | no |
| `projectIds` | `List<string>` | no |
| `assetTypes` | `List<string>` | no |
| `dataProduct` | `boolean` | no |
| `includeCascaded` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<RuleSummary>` | yes |
| `nextToken` | `string` | no |

## ListSubscriptionGrants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentId` | `string` | no |
| `subscriptionTargetId` | `string` | no |
| `subscribedListingId` | `string` | no |
| `subscriptionId` | `string` | no |
| `owningProjectId` | `string` | no |
| `owningIamPrincipalArn` | `string` | no |
| `owningUserId` | `string` | no |
| `owningGroupId` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SubscriptionGrantSummary>` | yes |
| `nextToken` | `string` | no |

## ListSubscriptionRequests

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `status` | `string` | no |
| `subscribedListingId` | `string` | no |
| `owningProjectId` | `string` | no |
| `owningIamPrincipalArn` | `string` | no |
| `approverProjectId` | `string` | no |
| `owningUserId` | `string` | no |
| `owningGroupId` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SubscriptionRequestSummary>` | yes |
| `nextToken` | `string` | no |

## ListSubscriptionTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SubscriptionTargetSummary>` | yes |
| `nextToken` | `string` | no |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `subscriptionRequestIdentifier` | `string` | no |
| `status` | `string` | no |
| `subscribedListingId` | `string` | no |
| `owningProjectId` | `string` | no |
| `owningIamPrincipalArn` | `string` | no |
| `owningUserId` | `string` | no |
| `owningGroupId` | `string` | no |
| `approverProjectId` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SubscriptionSummary>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTimeSeriesDataPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `formName` | `string` | yes |
| `startedAt` | `timestamp` | no |
| `endedAt` | `timestamp` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TimeSeriesDataPointSummaryFormOutput>` | no |
| `nextToken` | `string` | no |

## PostLineageEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `event` | `blob` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `domainId` | `string` | no |

## PostTimeSeriesDataPoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `forms` | `List<TimeSeriesDataPointFormInput>` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `entityId` | `string` | no |
| `entityType` | `string` | no |
| `forms` | `List<TimeSeriesDataPointFormOutput>` | no |

## PutDataExportConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `enableExport` | `boolean` | yes |
| `encryptionConfiguration` | `EncryptionConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEnvironmentBlueprintConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentBlueprintIdentifier` | `string` | yes |
| `provisioningRoleArn` | `string` | no |
| `manageAccessRoleArn` | `string` | no |
| `environmentRolePermissionBoundary` | `string` | no |
| `enabledRegions` | `List<string>` | yes |
| `regionalParameters` | `Map<Map<string>>` | no |
| `resourceConfigurations` | `List<PutResourceConfiguration>` | no |
| `allowUserProvidedConfigurations` | `boolean` | no |
| `globalParameters` | `Map<string>` | no |
| `provisioningConfigurations` | `List<ProvisioningConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `environmentBlueprintId` | `string` | yes |
| `provisioningRoleArn` | `string` | no |
| `environmentRolePermissionBoundary` | `string` | no |
| `manageAccessRoleArn` | `string` | no |
| `enabledRegions` | `List<string>` | no |
| `regionalParameters` | `Map<Map<string>>` | no |
| `allowUserProvidedConfigurations` | `boolean` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `resourceConfigurations` | `List<ResourceConfiguration>` | no |
| `provisioningConfigurations` | `List<ProvisioningConfiguration>` | no |

## QueryGraph

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `match` | `List<MatchClause>` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `additionalAttributes` | `AdditionalAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<ResultItem>` | no |
| `nextToken` | `string` | no |

## RejectPredictions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `revision` | `string` | no |
| `rejectRule` | `RejectRule` | no |
| `rejectChoices` | `List<RejectChoice>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `assetId` | `string` | yes |
| `assetRevision` | `string` | yes |

## RejectSubscriptionRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `decisionComment` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `requestReason` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipal>` | yes |
| `subscribedListings` | `List<SubscribedListing>` | yes |
| `reviewerId` | `string` | no |
| `decisionComment` | `string` | no |
| `existingSubscriptionId` | `string` | no |
| `metadataForms` | `List<FormOutput>` | no |

## RemoveEntityOwner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `owner` | `OwnerProperties` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemovePolicyGrant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `entityType` | `string` | yes |
| `entityIdentifier` | `string` | yes |
| `policyType` | `string` | yes |
| `principal` | `PolicyGrantPrincipal` | yes |
| `grantIdentifier` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RevokeSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `retainPermissions` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `subscribedPrincipal` | `SubscribedPrincipal` | yes |
| `subscribedListing` | `SubscribedListing` | yes |
| `subscriptionRequestId` | `string` | no |
| `retainPermissions` | `boolean` | no |

## Search

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchScope` | `string` | yes |
| `searchText` | `string` | no |
| `searchIn` | `List<SearchInItem>` | no |
| `filters` | `FilterClause` | no |
| `sort` | `SearchSort` | no |
| `additionalAttributes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SearchInventoryResultItem>` | no |
| `nextToken` | `string` | no |
| `totalMatchCount` | `integer` | no |

## SearchGroupProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `groupType` | `string` | yes |
| `searchText` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<GroupProfileSummary>` | no |
| `nextToken` | `string` | no |

## SearchListings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `searchText` | `string` | no |
| `searchIn` | `List<SearchInItem>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `filters` | `FilterClause` | no |
| `aggregations` | `List<AggregationListItem>` | no |
| `sort` | `SearchSort` | no |
| `additionalAttributes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SearchResultItem>` | no |
| `nextToken` | `string` | no |
| `totalMatchCount` | `integer` | no |
| `aggregates` | `List<AggregationOutput>` | no |

## SearchTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `searchScope` | `string` | yes |
| `searchText` | `string` | no |
| `searchIn` | `List<SearchInItem>` | no |
| `filters` | `FilterClause` | no |
| `sort` | `SearchSort` | no |
| `managed` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<SearchTypesResultItem>` | no |
| `nextToken` | `string` | no |
| `totalMatchCount` | `integer` | no |

## SearchUserProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `userType` | `string` | yes |
| `searchText` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<UserProfileSummary>` | no |
| `nextToken` | `string` | no |

## StartDataSourceRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `dataSourceIdentifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `dataSourceId` | `string` | yes |
| `id` | `string` | yes |
| `projectId` | `string` | yes |
| `status` | `string` | yes |
| `type` | `string` | yes |
| `dataSourceConfigurationSnapshot` | `string` | no |
| `runStatisticsForAssets` | `RunStatisticsForAssets` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `startedAt` | `timestamp` | no |
| `stoppedAt` | `timestamp` | no |

## StartMetadataGenerationRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `type` | `string` | no |
| `types` | `List<string>` | no |
| `target` | `MetadataGenerationRunTarget` | yes |
| `clientToken` | `string` | no |
| `owningProjectIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `types` | `List<string>` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `owningProjectId` | `string` | no |

## StartNotebookExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `notebookIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `fileFormat` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `notebookId` | `string` | yes |
| `fileFormat` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |

## StartNotebookImport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `sourceLocation` | `SourceLocation` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notebookId` | `string` | no |
| `status` | `string` | no |
| `domainId` | `string` | no |
| `owningProjectId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `sourceLocation` | `SourceLocation` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |

## StartNotebookRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `notebookIdentifier` | `string` | yes |
| `scheduleIdentifier` | `string` | no |
| `computeConfiguration` | `ComputeConfig` | no |
| `networkConfiguration` | `NetworkConfig` | no |
| `timeoutConfiguration` | `TimeoutConfig` | no |
| `triggerSource` | `TriggerSource` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `notebookId` | `string` | yes |
| `scheduleId` | `string` | no |
| `status` | `string` | yes |
| `cellOrder` | `List<CellInformation>` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `computeConfiguration` | `ComputeConfig` | no |
| `networkConfiguration` | `NetworkConfig` | no |
| `timeoutConfiguration` | `TimeoutConfig` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `storageConfiguration` | `StorageConfig` | no |
| `triggerSource` | `TriggerSource` | no |
| `error` | `NotebookRunError` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `startedAt` | `timestamp` | no |
| `completedAt` | `timestamp` | no |

## StartNotebookSync

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `owningProjectIdentifier` | `string` | yes |
| `sourceLocation` | `SourceLocation` | yes |
| `gitMetadata` | `GitMetadata` | no |
| `notebookId` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notebookId` | `string` | no |
| `status` | `string` | no |
| `domainId` | `string` | no |
| `owningProjectId` | `string` | no |
| `sourceLocation` | `SourceLocation` | no |
| `gitMetadata` | `GitMetadata` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |

## StopNotebookRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `status` | `string` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccountPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `resolutionStrategy` | `string` | no |
| `accountSource` | `AccountSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `name` | `string` | no |
| `id` | `string` | no |
| `description` | `string` | no |
| `resolutionStrategy` | `string` | no |
| `accountSource` | `AccountSource` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `domainUnitId` | `string` | no |

## UpdateAssetFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `configuration` | `AssetFilterConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `assetId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `configuration` | `AssetFilterConfiguration` | yes |
| `createdAt` | `timestamp` | no |
| `errorMessage` | `string` | no |
| `effectiveColumnNames` | `List<string>` | no |
| `effectiveRowFilter` | `string` | no |

## UpdateConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<Configuration>` | no |
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `awsLocation` | `AwsLocation` | no |
| `props` | `ConnectionPropertiesPatch` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<Configuration>` | no |
| `connectionId` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `domainUnitId` | `string` | yes |
| `environmentId` | `string` | no |
| `name` | `string` | yes |
| `physicalEndpoints` | `List<PhysicalEndpoint>` | yes |
| `projectId` | `string` | no |
| `props` | `ConnectionPropertiesOutput` | no |
| `type` | `string` | yes |
| `scope` | `string` | no |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `enableSetting` | `string` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsInput` | `List<FormInput>` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `configuration` | `DataSourceConfigurationInput` | no |
| `recommendation` | `RecommendationConfiguration` | no |
| `retainPermissionsOnRevokeFailure` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `status` | `string` | no |
| `type` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | no |
| `connectionId` | `string` | no |
| `configuration` | `DataSourceConfigurationOutput` | no |
| `recommendation` | `RecommendationConfiguration` | no |
| `enableSetting` | `string` | no |
| `publishOnImport` | `boolean` | no |
| `assetFormsOutput` | `List<FormOutput>` | no |
| `schedule` | `ScheduleConfiguration` | no |
| `lastRunStatus` | `string` | no |
| `lastRunAt` | `timestamp` | no |
| `lastRunErrorMessage` | `DataSourceErrorMessage` | no |
| `errorMessage` | `DataSourceErrorMessage` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `selfGrantStatus` | `SelfGrantStatusOutput` | no |
| `retainPermissionsOnRevokeFailure` | `boolean` | no |

## UpdateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `singleSignOn` | `SingleSignOn` | no |
| `domainExecutionRole` | `string` | no |
| `serviceRole` | `string` | no |
| `name` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `rootDomainUnitId` | `string` | no |
| `description` | `string` | no |
| `singleSignOn` | `SingleSignOn` | no |
| `domainExecutionRole` | `string` | no |
| `serviceRole` | `string` | no |
| `name` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |

## UpdateDomainUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `name` | `string` | yes |
| `owners` | `List<DomainUnitOwnerProperties>` | yes |
| `description` | `string` | no |
| `parentDomainUnitId` | `string` | no |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `lastUpdatedBy` | `string` | no |

## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `glossaryTerms` | `List<string>` | no |
| `blueprintVersion` | `string` | no |
| `userParameters` | `List<EnvironmentParameter>` | no |
| `environmentConfigurationName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `projectId` | `string` | yes |
| `id` | `string` | no |
| `domainId` | `string` | yes |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentProfileId` | `string` | no |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `provider` | `string` | yes |
| `provisionedResources` | `List<Resource>` | no |
| `status` | `string` | no |
| `environmentActions` | `List<ConfigurableEnvironmentAction>` | no |
| `glossaryTerms` | `List<string>` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `lastDeployment` | `Deployment` | no |
| `provisioningProperties` | `ProvisioningProperties` | no |
| `deploymentProperties` | `DeploymentProperties` | no |
| `environmentBlueprintId` | `string` | no |
| `environmentConfigurationId` | `string` | no |
| `environmentConfigurationName` | `string` | no |

## UpdateEnvironmentAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `parameters` | `ActionParameters` | no |
| `name` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `environmentId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `parameters` | `ActionParameters` | yes |
| `description` | `string` | no |

## UpdateEnvironmentBlueprint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `provisioningProperties` | `ProvisioningProperties` | no |
| `userParameters` | `List<CustomParameter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `provider` | `string` | yes |
| `provisioningProperties` | `ProvisioningProperties` | yes |
| `deploymentProperties` | `DeploymentProperties` | no |
| `userParameters` | `List<CustomParameter>` | no |
| `glossaryTerms` | `List<string>` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## UpdateEnvironmentProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `userParameters` | `List<EnvironmentParameter>` | no |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `awsAccountId` | `string` | no |
| `awsAccountRegion` | `string` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `environmentBlueprintId` | `string` | yes |
| `projectId` | `string` | no |
| `userParameters` | `List<CustomParameter>` | no |

## UpdateGlossary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `usageRestrictions` | `List<string>` | no |

## UpdateGlossaryTerm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `glossaryIdentifier` | `string` | no |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `shortDescription` | `string` | no |
| `longDescription` | `string` | no |
| `termRelations` | `TermRelations` | no |
| `status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `domainId` | `string` | yes |
| `glossaryId` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `shortDescription` | `string` | no |
| `longDescription` | `string` | no |
| `termRelations` | `TermRelations` | no |
| `usageRestrictions` | `List<string>` | no |

## UpdateGroupProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `groupIdentifier` | `string` | yes |
| `status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `status` | `string` | no |
| `groupName` | `string` | no |
| `rolePrincipalArn` | `string` | no |
| `rolePrincipalId` | `string` | no |

## UpdateNotebook

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `name` | `string` | no |
| `cellOrder` | `List<CellInformation>` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `name` | `string` | yes |
| `owningProjectId` | `string` | yes |
| `domainId` | `string` | yes |
| `cellOrder` | `List<CellInformation>` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `createdBy` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |
| `lockedBy` | `string` | no |
| `lockedAt` | `timestamp` | no |
| `lockExpiresAt` | `timestamp` | no |
| `computeId` | `string` | no |
| `metadata` | `Map<string>` | no |
| `parameters` | `Map<string>` | no |
| `environmentConfiguration` | `EnvironmentConfig` | no |
| `error` | `NotebookError` | no |
| `gitMetadata` | `GitMetadata` | no |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `resourceTags` | `Map<string>` | no |
| `glossaryTerms` | `List<string>` | no |
| `domainUnitId` | `string` | no |
| `environmentDeploymentDetails` | `EnvironmentDeploymentDetails` | no |
| `userParameters` | `List<EnvironmentConfigurationUserParameter>` | no |
| `projectProfileVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `projectStatus` | `string` | no |
| `failureReasons` | `List<ProjectDeletionError>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `resourceTags` | `List<ResourceTag>` | no |
| `glossaryTerms` | `List<string>` | no |
| `domainUnitId` | `string` | no |
| `projectProfileId` | `string` | no |
| `userParameters` | `List<EnvironmentConfigurationUserParameter>` | no |
| `environmentDeploymentDetails` | `EnvironmentDeploymentDetails` | no |
| `projectCategory` | `string` | no |

## UpdateProjectProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `status` | `string` | no |
| `projectResourceTags` | `List<ResourceTagParameter>` | no |
| `allowCustomProjectResourceTags` | `boolean` | no |
| `projectResourceTagsDescription` | `string` | no |
| `environmentConfigurations` | `List<EnvironmentConfiguration>` | no |
| `domainUnitIdentifier` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | yes |
| `id` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `projectResourceTags` | `List<ResourceTagParameter>` | no |
| `allowCustomProjectResourceTags` | `boolean` | no |
| `projectResourceTagsDescription` | `string` | no |
| `environmentConfigurations` | `List<EnvironmentConfiguration>` | no |
| `createdBy` | `string` | yes |
| `createdAt` | `timestamp` | no |
| `lastUpdatedAt` | `timestamp` | no |
| `domainUnitId` | `string` | no |

## UpdateRootDomainUnitOwner

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `currentOwner` | `string` | yes |
| `newOwner` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `scope` | `RuleScope` | no |
| `detail` | `RuleDetail` | no |
| `includeChildDomainUnits` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `identifier` | `string` | yes |
| `revision` | `string` | yes |
| `name` | `string` | yes |
| `ruleType` | `string` | yes |
| `target` | `RuleTarget` | yes |
| `action` | `string` | yes |
| `scope` | `RuleScope` | yes |
| `detail` | `RuleDetail` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `createdBy` | `string` | yes |
| `lastUpdatedBy` | `string` | yes |

## UpdateSubscriptionGrantStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `assetIdentifier` | `string` | yes |
| `status` | `string` | yes |
| `failureCause` | `FailureCause` | no |
| `targetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `environmentId` | `string` | no |
| `subscriptionTargetId` | `string` | yes |
| `grantedEntity` | `GrantedEntity` | yes |
| `status` | `string` | yes |
| `assets` | `List<SubscribedAsset>` | no |
| `subscriptionId` | `string` | no |

## UpdateSubscriptionRequest

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `requestReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `domainId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `requestReason` | `string` | yes |
| `subscribedPrincipals` | `List<SubscribedPrincipal>` | yes |
| `subscribedListings` | `List<SubscribedListing>` | yes |
| `reviewerId` | `string` | no |
| `decisionComment` | `string` | no |
| `existingSubscriptionId` | `string` | no |
| `metadataForms` | `List<FormOutput>` | no |

## UpdateSubscriptionTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `environmentIdentifier` | `string` | yes |
| `identifier` | `string` | yes |
| `name` | `string` | no |
| `authorizedPrincipals` | `List<string>` | no |
| `applicableAssetTypes` | `List<string>` | no |
| `subscriptionTargetConfig` | `List<SubscriptionTargetForm>` | no |
| `manageAccessRole` | `string` | no |
| `provider` | `string` | no |
| `subscriptionGrantCreationMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `authorizedPrincipals` | `List<string>` | yes |
| `domainId` | `string` | yes |
| `projectId` | `string` | yes |
| `environmentId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `createdBy` | `string` | yes |
| `updatedBy` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | no |
| `manageAccessRole` | `string` | no |
| `applicableAssetTypes` | `List<string>` | yes |
| `subscriptionTargetConfig` | `List<SubscriptionTargetForm>` | yes |
| `provider` | `string` | yes |
| `subscriptionGrantCreationMode` | `string` | no |

## UpdateUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainIdentifier` | `string` | yes |
| `userIdentifier` | `string` | yes |
| `type` | `string` | no |
| `status` | `string` | yes |
| `sessionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainId` | `string` | no |
| `id` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `details` | `UserProfileDetails` | no |

