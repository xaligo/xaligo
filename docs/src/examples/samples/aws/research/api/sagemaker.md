# Amazon SageMaker Service

API version: 2017-07-24. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sagemaker/2017-07-24/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `DestinationArn` | `string` | yes |
| `AssociationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `DestinationArn` | `string` | no |

## AddTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## AssociateTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |
| `TrialName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentArn` | `string` | no |
| `TrialArn` | `string` | no |

## AttachClusterNodeVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `NodeId` | `string` | yes |
| `VolumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `NodeId` | `string` | yes |
| `VolumeId` | `string` | yes |
| `AttachTime` | `timestamp` | yes |
| `Status` | `string` | yes |
| `DeviceName` | `string` | yes |

## BatchAddClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `ClientToken` | `string` | no |
| `NodesToAdd` | `List<AddClusterNodeSpecification>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<NodeAdditionResult>` | yes |
| `Failed` | `List<BatchAddClusterNodesError>` | yes |

## BatchDeleteClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeIds` | `List<string>` | no |
| `NodeLogicalIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Failed` | `List<BatchDeleteClusterNodesError>` | no |
| `Successful` | `List<string>` | no |
| `FailedNodeLogicalIds` | `List<BatchDeleteClusterNodeLogicalIdsError>` | no |
| `SuccessfulNodeLogicalIds` | `List<string>` | no |

## BatchDescribeModelPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageArnList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageSummaries` | `Map<BatchDescribeModelPackageSummary>` | no |
| `BatchDescribeModelPackageErrorMap` | `Map<BatchDescribeModelPackageError>` | no |

## BatchRebootClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeIds` | `List<string>` | no |
| `NodeLogicalIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<string>` | no |
| `Failed` | `List<BatchRebootClusterNodesError>` | no |
| `FailedNodeLogicalIds` | `List<BatchRebootClusterNodeLogicalIdsError>` | no |
| `SuccessfulNodeLogicalIds` | `List<string>` | no |

## BatchReplaceClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeIds` | `List<string>` | no |
| `NodeLogicalIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<string>` | no |
| `Failed` | `List<BatchReplaceClusterNodesError>` | no |
| `FailedNodeLogicalIds` | `List<BatchReplaceClusterNodeLogicalIdsError>` | no |
| `SuccessfulNodeLogicalIds` | `List<string>` | no |

## CreateAIBenchmarkJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobName` | `string` | yes |
| `BenchmarkTarget` | `AIBenchmarkTarget` | yes |
| `OutputConfig` | `AIBenchmarkOutputConfig` | yes |
| `AIWorkloadConfigIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `NetworkConfig` | `AIBenchmarkNetworkConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobArn` | `string` | yes |

## CreateAIRecommendationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobName` | `string` | yes |
| `ModelSource` | `AIModelSource` | yes |
| `OutputConfig` | `AIRecommendationOutputConfig` | yes |
| `AIWorkloadConfigIdentifier` | `string` | yes |
| `PerformanceTarget` | `AIRecommendationPerformanceTarget` | yes |
| `RoleArn` | `string` | yes |
| `InferenceSpecification` | `AIRecommendationInferenceSpecification` | no |
| `OptimizeModel` | `boolean` | no |
| `ComputeSpec` | `AIRecommendationComputeSpec` | no |
| `AdapterSource` | `AIAdapterSource` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobArn` | `string` | yes |

## CreateAIWorkloadConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigName` | `string` | yes |
| `DatasetConfig` | `AIDatasetConfig` | no |
| `AIWorkloadConfigs` | `AIWorkloadConfigs` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigArn` | `string` | yes |

## CreateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionName` | `string` | yes |
| `Source` | `ActionSource` | yes |
| `ActionType` | `string` | yes |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `Properties` | `Map<string>` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionArn` | `string` | no |

## CreateAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmName` | `string` | yes |
| `AlgorithmDescription` | `string` | no |
| `TrainingSpecification` | `TrainingSpecification` | yes |
| `InferenceSpecification` | `InferenceSpecification` | no |
| `ValidationSpecification` | `AlgorithmValidationSpecification` | no |
| `CertifyForMarketplace` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmArn` | `string` | yes |

## CreateApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | no |
| `SpaceName` | `string` | no |
| `AppType` | `string` | yes |
| `AppName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ResourceSpec` | `ResourceSpec` | no |
| `RecoveryMode` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppArn` | `string` | no |

## CreateAppImageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `KernelGatewayImageConfig` | `KernelGatewayImageConfig` | no |
| `JupyterLabAppImageConfig` | `JupyterLabAppImageConfig` | no |
| `CodeEditorAppImageConfig` | `CodeEditorAppImageConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigArn` | `string` | no |

## CreateArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactName` | `string` | no |
| `Source` | `ArtifactSource` | yes |
| `ArtifactType` | `string` | yes |
| `Properties` | `Map<string>` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | no |

## CreateAutoMLJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |
| `InputDataConfig` | `List<AutoMLChannel>` | yes |
| `OutputDataConfig` | `AutoMLOutputDataConfig` | yes |
| `ProblemType` | `string` | no |
| `AutoMLJobObjective` | `AutoMLJobObjective` | no |
| `AutoMLJobConfig` | `AutoMLJobConfig` | no |
| `RoleArn` | `string` | yes |
| `GenerateCandidateDefinitionsOnly` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ModelDeployConfig` | `ModelDeployConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobArn` | `string` | yes |

## CreateAutoMLJobV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |
| `AutoMLJobInputDataConfig` | `List<AutoMLJobChannel>` | yes |
| `OutputDataConfig` | `AutoMLOutputDataConfig` | yes |
| `AutoMLProblemTypeConfig` | `AutoMLProblemTypeConfig` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `SecurityConfig` | `AutoMLSecurityConfig` | no |
| `AutoMLJobObjective` | `AutoMLJobObjective` | no |
| `ModelDeployConfig` | `ModelDeployConfig` | no |
| `DataSplitConfig` | `AutoMLDataSplitConfig` | no |
| `AutoMLComputeConfig` | `AutoMLComputeConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobArn` | `string` | yes |

## CreateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `InstanceGroups` | `List<ClusterInstanceGroupSpecification>` | no |
| `RestrictedInstanceGroups` | `List<ClusterRestrictedInstanceGroupSpecification>` | no |
| `RestrictedInstanceGroupsConfig` | `ClusterRestrictedInstanceGroupsConfig` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Tags` | `List<Tag>` | no |
| `Orchestrator` | `ClusterOrchestrator` | no |
| `NodeRecovery` | `string` | no |
| `TieredStorageConfig` | `ClusterTieredStorageConfig` | no |
| `NodeProvisioningMode` | `string` | no |
| `ClusterRole` | `string` | no |
| `AutoScaling` | `ClusterAutoScalingConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

## CreateClusterSchedulerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ClusterArn` | `string` | yes |
| `SchedulerConfig` | `SchedulerConfig` | yes |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigArn` | `string` | yes |
| `ClusterSchedulerConfigId` | `string` | yes |

## CreateCodeRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryName` | `string` | yes |
| `GitConfig` | `GitConfig` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryArn` | `string` | yes |

## CreateCompilationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `ModelPackageVersionArn` | `string` | no |
| `InputConfig` | `InputConfig` | no |
| `OutputConfig` | `OutputConfig` | yes |
| `VpcConfig` | `NeoVpcConfig` | no |
| `StoppingCondition` | `StoppingCondition` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobArn` | `string` | yes |

## CreateComputeQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ClusterArn` | `string` | yes |
| `ComputeQuotaConfig` | `ComputeQuotaConfig` | yes |
| `ComputeQuotaTarget` | `ComputeQuotaTarget` | yes |
| `ActivationState` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaArn` | `string` | yes |
| `ComputeQuotaId` | `string` | yes |

## CreateContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextName` | `string` | yes |
| `Source` | `ContextSource` | yes |
| `ContextType` | `string` | yes |
| `Description` | `string` | no |
| `Properties` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextArn` | `string` | no |

## CreateDataQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |
| `DataQualityBaselineConfig` | `DataQualityBaselineConfig` | no |
| `DataQualityAppSpecification` | `DataQualityAppSpecification` | yes |
| `DataQualityJobInput` | `DataQualityJobInput` | yes |
| `DataQualityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |

## CreateDeviceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `RoleArn` | `string` | no |
| `Description` | `string` | no |
| `OutputConfig` | `EdgeOutputConfig` | yes |
| `Tags` | `List<Tag>` | no |
| `EnableIotRoleAlias` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AuthMode` | `string` | yes |
| `DefaultUserSettings` | `UserSettings` | yes |
| `DomainSettings` | `DomainSettings` | no |
| `SubnetIds` | `List<string>` | no |
| `VpcId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `AppNetworkAccessType` | `string` | no |
| `HomeEfsFileSystemKmsKeyId` | `string` | no |
| `KmsKeyId` | `string` | no |
| `AppSecurityGroupManagement` | `string` | no |
| `HomeEfsFileSystemCreation` | `string` | no |
| `TagPropagation` | `string` | no |
| `DefaultSpaceSettings` | `DefaultSpaceSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainArn` | `string` | no |
| `DomainId` | `string` | no |
| `Url` | `string` | no |

## CreateEdgeDeploymentPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `ModelConfigs` | `List<EdgeDeploymentModelConfig>` | yes |
| `DeviceFleetName` | `string` | yes |
| `Stages` | `List<DeploymentStage>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanArn` | `string` | yes |

## CreateEdgeDeploymentStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `Stages` | `List<DeploymentStage>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateEdgePackagingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgePackagingJobName` | `string` | yes |
| `CompilationJobName` | `string` | yes |
| `ModelName` | `string` | yes |
| `ModelVersion` | `string` | yes |
| `RoleArn` | `string` | yes |
| `OutputConfig` | `EdgeOutputConfig` | yes |
| `ResourceKey` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `EndpointConfigName` | `string` | yes |
| `DeploymentConfig` | `DeploymentConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

## CreateEndpointConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigName` | `string` | yes |
| `ProductionVariants` | `List<ProductionVariant>` | yes |
| `DataCaptureConfig` | `DataCaptureConfig` | no |
| `Tags` | `List<Tag>` | no |
| `KmsKeyId` | `string` | no |
| `AsyncInferenceConfig` | `AsyncInferenceConfig` | no |
| `ExplainerConfig` | `ExplainerConfig` | no |
| `ShadowProductionVariants` | `List<ProductionVariant>` | no |
| `ExecutionRoleArn` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `EnableNetworkIsolation` | `boolean` | no |
| `MetricsConfig` | `MetricsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigArn` | `string` | yes |

## CreateExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentArn` | `string` | no |

## CreateFeatureGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `RecordIdentifierFeatureName` | `string` | yes |
| `EventTimeFeatureName` | `string` | yes |
| `FeatureDefinitions` | `List<FeatureDefinition>` | yes |
| `OnlineStoreConfig` | `OnlineStoreConfig` | no |
| `OfflineStoreConfig` | `OfflineStoreConfig` | no |
| `ThroughputConfig` | `ThroughputConfig` | no |
| `RoleArn` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupArn` | `string` | yes |

## CreateFlowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionName` | `string` | yes |
| `HumanLoopRequestSource` | `HumanLoopRequestSource` | no |
| `HumanLoopActivationConfig` | `HumanLoopActivationConfig` | no |
| `HumanLoopConfig` | `HumanLoopConfig` | no |
| `OutputConfig` | `FlowDefinitionOutputConfig` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionArn` | `string` | yes |

## CreateHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubDescription` | `string` | yes |
| `HubDisplayName` | `string` | no |
| `HubSearchKeywords` | `List<string>` | no |
| `S3StorageConfig` | `HubS3StorageConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |

## CreateHubContentPresignedUrls

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentName` | `string` | yes |
| `HubContentVersion` | `string` | no |
| `AccessConfig` | `PresignedUrlAccessConfig` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrlConfigs` | `List<AuthorizedUrl>` | yes |
| `NextToken` | `string` | no |

## CreateHubContentReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `SageMakerPublicHubContentArn` | `string` | yes |
| `HubContentName` | `string` | no |
| `MinVersion` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |
| `HubContentArn` | `string` | yes |

## CreateHumanTaskUi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiName` | `string` | yes |
| `UiTemplate` | `UiTemplate` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiArn` | `string` | yes |

## CreateHyperParameterTuningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |
| `HyperParameterTuningJobConfig` | `HyperParameterTuningJobConfig` | yes |
| `TrainingJobDefinition` | `HyperParameterTrainingJobDefinition` | no |
| `TrainingJobDefinitions` | `List<HyperParameterTrainingJobDefinition>` | no |
| `WarmStartConfig` | `HyperParameterTuningJobWarmStartConfig` | no |
| `Tags` | `List<Tag>` | no |
| `Autotune` | `Autotune` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobArn` | `string` | yes |

## CreateImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `ImageName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageArn` | `string` | no |

## CreateImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaseImage` | `string` | yes |
| `ClientToken` | `string` | yes |
| `ImageName` | `string` | yes |
| `Aliases` | `List<string>` | no |
| `VendorGuidance` | `string` | no |
| `JobType` | `string` | no |
| `MLFramework` | `string` | no |
| `ProgrammingLang` | `string` | no |
| `Processor` | `string` | no |
| `Horovod` | `boolean` | no |
| `ReleaseNotes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageVersionArn` | `string` | no |

## CreateInferenceComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |
| `EndpointName` | `string` | yes |
| `VariantName` | `string` | no |
| `Specification` | `InferenceComponentSpecification` | no |
| `Specifications` | `List<InferenceComponentSpecification>` | no |
| `RuntimeConfig` | `InferenceComponentRuntimeConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentArn` | `string` | yes |

## CreateInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Schedule` | `InferenceExperimentSchedule` | no |
| `Description` | `string` | no |
| `RoleArn` | `string` | yes |
| `EndpointName` | `string` | yes |
| `ModelVariants` | `List<ModelVariantConfig>` | yes |
| `DataStorageConfig` | `InferenceExperimentDataStorageConfig` | no |
| `ShadowModeConfig` | `ShadowModeConfig` | yes |
| `KmsKey` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperimentArn` | `string` | yes |

## CreateInferenceRecommendationsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobType` | `string` | yes |
| `RoleArn` | `string` | yes |
| `InputConfig` | `RecommendationJobInputConfig` | yes |
| `JobDescription` | `string` | no |
| `StoppingConditions` | `RecommendationJobStoppingConditions` | no |
| `OutputConfig` | `RecommendationJobOutputConfig` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |

## CreateJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `JobCategory` | `string` | yes |
| `JobConfigSchemaVersion` | `string` | yes |
| `JobConfigDocument` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobArn` | `string` | yes |

## CreateLabelingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobName` | `string` | yes |
| `LabelAttributeName` | `string` | yes |
| `InputConfig` | `LabelingJobInputConfig` | yes |
| `OutputConfig` | `LabelingJobOutputConfig` | yes |
| `RoleArn` | `string` | yes |
| `LabelCategoryConfigS3Uri` | `string` | no |
| `StoppingConditions` | `LabelingJobStoppingConditions` | no |
| `LabelingJobAlgorithmsConfig` | `LabelingJobAlgorithmsConfig` | no |
| `HumanTaskConfig` | `HumanTaskConfig` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobArn` | `string` | yes |

## CreateMlflowApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ArtifactStoreUri` | `string` | yes |
| `RoleArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `ModelRegistrationMode` | `string` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `AccountDefaultStatus` | `string` | no |
| `DefaultDomainIdList` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreateMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |
| `ArtifactStoreUri` | `string` | yes |
| `TrackingServerSize` | `string` | no |
| `MlflowVersion` | `string` | no |
| `RoleArn` | `string` | yes |
| `AutomaticModelRegistration` | `boolean` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `S3BucketOwnerAccountId` | `string` | no |
| `S3BucketOwnerVerification` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |

## CreateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `PrimaryContainer` | `ContainerDefinition` | no |
| `Containers` | `List<ContainerDefinition>` | no |
| `InferenceExecutionConfig` | `InferenceExecutionConfig` | no |
| `ExecutionRoleArn` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `VpcConfig` | `VpcConfig` | no |
| `EnableNetworkIsolation` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelArn` | `string` | yes |

## CreateModelBiasJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |
| `ModelBiasBaselineConfig` | `ModelBiasBaselineConfig` | no |
| `ModelBiasAppSpecification` | `ModelBiasAppSpecification` | yes |
| `ModelBiasJobInput` | `ModelBiasJobInput` | yes |
| `ModelBiasJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |

## CreateModelCard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |
| `SecurityConfig` | `ModelCardSecurityConfig` | no |
| `Content` | `string` | yes |
| `ModelCardStatus` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardArn` | `string` | yes |

## CreateModelCardExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |
| `ModelCardVersion` | `integer` | no |
| `ModelCardExportJobName` | `string` | yes |
| `OutputConfig` | `ModelCardExportOutputConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardExportJobArn` | `string` | yes |

## CreateModelExplainabilityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |
| `ModelExplainabilityBaselineConfig` | `ModelExplainabilityBaselineConfig` | no |
| `ModelExplainabilityAppSpecification` | `ModelExplainabilityAppSpecification` | yes |
| `ModelExplainabilityJobInput` | `ModelExplainabilityJobInput` | yes |
| `ModelExplainabilityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |

## CreateModelPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageName` | `string` | no |
| `ModelPackageGroupName` | `string` | no |
| `ModelPackageDescription` | `string` | no |
| `ModelPackageRegistrationType` | `string` | no |
| `InferenceSpecification` | `InferenceSpecification` | no |
| `ValidationSpecification` | `ModelPackageValidationSpecification` | no |
| `SourceAlgorithmSpecification` | `SourceAlgorithmSpecification` | no |
| `CertifyForMarketplace` | `boolean` | no |
| `Tags` | `List<Tag>` | no |
| `ModelApprovalStatus` | `string` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `ModelMetrics` | `ModelMetrics` | no |
| `ClientToken` | `string` | no |
| `Domain` | `string` | no |
| `Task` | `string` | no |
| `SamplePayloadUrl` | `string` | no |
| `CustomerMetadataProperties` | `Map<string>` | no |
| `DriftCheckBaselines` | `DriftCheckBaselines` | no |
| `AdditionalInferenceSpecifications` | `List<AdditionalInferenceSpecificationDefinition>` | no |
| `SkipModelValidation` | `string` | no |
| `SourceUri` | `string` | no |
| `SecurityConfig` | `ModelPackageSecurityConfig` | no |
| `ModelCard` | `ModelPackageModelCard` | no |
| `ModelLifeCycle` | `ModelLifeCycle` | no |
| `ManagedStorageType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageArn` | `string` | yes |

## CreateModelPackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |
| `ModelPackageGroupDescription` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `ManagedConfiguration` | `ManagedConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupArn` | `string` | yes |

## CreateModelQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |
| `ModelQualityBaselineConfig` | `ModelQualityBaselineConfig` | no |
| `ModelQualityAppSpecification` | `ModelQualityAppSpecification` | yes |
| `ModelQualityJobInput` | `ModelQualityJobInput` | yes |
| `ModelQualityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |

## CreateMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |
| `MonitoringScheduleConfig` | `MonitoringScheduleConfig` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleArn` | `string` | yes |

## CreateNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |
| `InstanceType` | `string` | yes |
| `SubnetId` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `RoleArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `LifecycleConfigName` | `string` | no |
| `DirectInternetAccess` | `string` | no |
| `VolumeSizeInGB` | `integer` | no |
| `AcceleratorTypes` | `List<string>` | no |
| `DefaultCodeRepository` | `string` | no |
| `AdditionalCodeRepositories` | `List<string>` | no |
| `RootAccess` | `string` | no |
| `PlatformIdentifier` | `string` | no |
| `InstanceMetadataServiceConfiguration` | `InstanceMetadataServiceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceArn` | `string` | no |

## CreateNotebookInstanceLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigName` | `string` | yes |
| `OnCreate` | `List<NotebookInstanceLifecycleHook>` | no |
| `OnStart` | `List<NotebookInstanceLifecycleHook>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigArn` | `string` | no |

## CreateOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `ModelSource` | `OptimizationJobModelSource` | yes |
| `DeploymentInstanceType` | `string` | yes |
| `MaxInstanceCount` | `integer` | no |
| `OptimizationEnvironment` | `Map<string>` | no |
| `OptimizationConfigs` | `List<OptimizationConfig>` | yes |
| `OutputConfig` | `OptimizationJobOutputConfig` | yes |
| `StoppingCondition` | `StoppingCondition` | yes |
| `Tags` | `List<Tag>` | no |
| `VpcConfig` | `OptimizationVpcConfig` | no |
| `TrainingPlanArns` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobArn` | `string` | yes |

## CreatePartnerApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `KmsKeyId` | `string` | no |
| `MaintenanceConfig` | `PartnerAppMaintenanceConfig` | no |
| `Tier` | `string` | yes |
| `ApplicationConfig` | `PartnerAppConfig` | no |
| `IdcConfig` | `IdcConfigInput` | no |
| `AuthType` | `string` | yes |
| `EnableIamSessionBasedIdentity` | `boolean` | no |
| `EnableAutoMinorVersionUpgrade` | `boolean` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreatePartnerAppPresignedUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ExpiresInSeconds` | `integer` | no |
| `SessionExpirationDurationInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | no |

## CreatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `PipelineDisplayName` | `string` | no |
| `PipelineDefinition` | `string` | no |
| `PipelineDefinitionS3Location` | `PipelineDefinitionS3Location` | no |
| `PipelineDescription` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |

## CreatePresignedDomainUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | yes |
| `SessionExpirationDurationInSeconds` | `integer` | no |
| `ExpiresInSeconds` | `integer` | no |
| `SpaceName` | `string` | no |
| `LandingUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrl` | `string` | no |

## CreatePresignedMlflowAppUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ExpiresInSeconds` | `integer` | no |
| `SessionExpirationDurationInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrl` | `string` | no |

## CreatePresignedMlflowTrackingServerUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |
| `ExpiresInSeconds` | `integer` | no |
| `SessionExpirationDurationInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrl` | `string` | no |

## CreatePresignedNotebookInstanceUrl

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |
| `SessionExpirationDurationInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizedUrl` | `string` | no |

## CreateProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingInputs` | `List<ProcessingInput>` | no |
| `ProcessingOutputConfig` | `ProcessingOutputConfig` | no |
| `ProcessingJobName` | `string` | yes |
| `ProcessingResources` | `ProcessingResources` | yes |
| `StoppingCondition` | `ProcessingStoppingCondition` | no |
| `AppSpecification` | `AppSpecification` | yes |
| `Environment` | `Map<string>` | no |
| `NetworkConfig` | `NetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingJobArn` | `string` | yes |

## CreateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectName` | `string` | yes |
| `ProjectDescription` | `string` | no |
| `ServiceCatalogProvisioningDetails` | `ServiceCatalogProvisioningDetails` | no |
| `Tags` | `List<Tag>` | no |
| `TemplateProviders` | `List<CreateTemplateProvider>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `ProjectId` | `string` | yes |

## CreateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpaceName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `SpaceSettings` | `SpaceSettings` | no |
| `OwnershipSettings` | `OwnershipSettings` | no |
| `SpaceSharingSettings` | `SpaceSharingSettings` | no |
| `SpaceDisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpaceArn` | `string` | no |

## CreateStudioLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioLifecycleConfigName` | `string` | yes |
| `StudioLifecycleConfigContent` | `string` | yes |
| `StudioLifecycleConfigAppType` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioLifecycleConfigArn` | `string` | no |

## CreateTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |
| `HyperParameters` | `Map<string>` | no |
| `AlgorithmSpecification` | `AlgorithmSpecification` | no |
| `RoleArn` | `string` | yes |
| `InputDataConfig` | `List<Channel>` | no |
| `OutputDataConfig` | `OutputDataConfig` | yes |
| `ResourceConfig` | `ResourceConfig` | no |
| `VpcConfig` | `VpcConfig` | no |
| `StoppingCondition` | `StoppingCondition` | no |
| `Tags` | `List<Tag>` | no |
| `EnableNetworkIsolation` | `boolean` | no |
| `EnableInterContainerTrafficEncryption` | `boolean` | no |
| `EnableManagedSpotTraining` | `boolean` | no |
| `CheckpointConfig` | `CheckpointConfig` | no |
| `DebugHookConfig` | `DebugHookConfig` | no |
| `DebugRuleConfigurations` | `List<DebugRuleConfiguration>` | no |
| `TensorBoardOutputConfig` | `TensorBoardOutputConfig` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |
| `ProfilerConfig` | `ProfilerConfig` | no |
| `ProfilerRuleConfigurations` | `List<ProfilerRuleConfiguration>` | no |
| `Environment` | `Map<string>` | no |
| `RetryStrategy` | `RetryStrategy` | no |
| `RemoteDebugConfig` | `RemoteDebugConfig` | no |
| `InfraCheckConfig` | `InfraCheckConfig` | no |
| `SessionChainingConfig` | `SessionChainingConfig` | no |
| `ServerlessJobConfig` | `ServerlessJobConfig` | no |
| `MlflowConfig` | `MlflowConfig` | no |
| `ModelPackageConfig` | `ModelPackageConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobArn` | `string` | yes |

## CreateTrainingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanName` | `string` | yes |
| `TrainingPlanOfferingId` | `string` | yes |
| `SpareInstanceCountPerUltraServer` | `integer` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanArn` | `string` | yes |

## CreateTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobName` | `string` | yes |
| `ModelName` | `string` | yes |
| `MaxConcurrentTransforms` | `integer` | no |
| `ModelClientConfig` | `ModelClientConfig` | no |
| `MaxPayloadInMB` | `integer` | no |
| `BatchStrategy` | `string` | no |
| `Environment` | `Map<string>` | no |
| `TransformInput` | `TransformInput` | yes |
| `TransformOutput` | `TransformOutput` | yes |
| `DataCaptureConfig` | `BatchDataCaptureConfig` | no |
| `TransformResources` | `TransformResources` | yes |
| `DataProcessing` | `DataProcessing` | no |
| `Tags` | `List<Tag>` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobArn` | `string` | yes |

## CreateTrial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialName` | `string` | yes |
| `DisplayName` | `string` | no |
| `ExperimentName` | `string` | yes |
| `MetadataProperties` | `MetadataProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialArn` | `string` | no |

## CreateTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Status` | `TrialComponentStatus` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Parameters` | `Map<TrialComponentParameterValue>` | no |
| `InputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `OutputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentArn` | `string` | no |

## CreateUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | yes |
| `SingleSignOnUserIdentifier` | `string` | no |
| `SingleSignOnUserValue` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `UserSettings` | `UserSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserProfileArn` | `string` | no |

## CreateWorkforce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CognitoConfig` | `CognitoConfig` | no |
| `OidcConfig` | `OidcConfig` | no |
| `SourceIpConfig` | `SourceIpConfig` | no |
| `WorkforceName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `WorkforceVpcConfig` | `WorkforceVpcConfigRequest` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkforceArn` | `string` | yes |

## CreateWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamName` | `string` | yes |
| `WorkforceName` | `string` | no |
| `MemberDefinitions` | `List<MemberDefinition>` | yes |
| `Description` | `string` | yes |
| `NotificationConfiguration` | `NotificationConfiguration` | no |
| `WorkerAccessConfiguration` | `WorkerAccessConfiguration` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamArn` | `string` | no |

## DeleteAIBenchmarkJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobArn` | `string` | no |

## DeleteAIRecommendationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobArn` | `string` | no |

## DeleteAIWorkloadConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigArn` | `string` | no |

## DeleteAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionArn` | `string` | no |

## DeleteAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | no |
| `SpaceName` | `string` | no |
| `AppType` | `string` | yes |
| `AppName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppImageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | no |
| `Source` | `ArtifactSource` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | no |

## DeleteAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | yes |
| `DestinationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `DestinationArn` | `string` | no |

## DeleteCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

## DeleteClusterSchedulerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCodeRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCompilationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteComputeQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextArn` | `string` | no |

## DeleteDataQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeviceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `RetentionPolicy` | `RetentionPolicy` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEdgeDeploymentPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEdgeDeploymentStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEndpointConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentArn` | `string` | no |

## DeleteFeatureGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFlowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHubContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentName` | `string` | yes |
| `HubContentVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHubContentReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHumanTaskUi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHyperParameterTuningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `Version` | `integer` | no |
| `Alias` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInferenceComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperimentArn` | `string` | yes |

## DeleteJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobCategory` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMlflowApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeleteMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |

## DeleteModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelBiasJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelCard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelExplainabilityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelPackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelPackageGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotebookInstanceLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePartnerApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## DeletePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |

## DeleteProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStudioLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioLifecycleConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialArn` | `string` | no |

## DeleteTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentArn` | `string` | no |

## DeleteUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkforce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkforceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Success` | `boolean` | yes |

## DeregisterDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `DeviceNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAIBenchmarkJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobName` | `string` | yes |
| `AIBenchmarkJobArn` | `string` | yes |
| `AIBenchmarkJobStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `BenchmarkTarget` | `AIBenchmarkTarget` | yes |
| `OutputConfig` | `AIBenchmarkOutputResult` | yes |
| `AIWorkloadConfigIdentifier` | `string` | yes |
| `RoleArn` | `string` | yes |
| `NetworkConfig` | `AIBenchmarkNetworkConfig` | no |
| `CreationTime` | `timestamp` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |

## DescribeAIRecommendationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobName` | `string` | yes |
| `AIRecommendationJobArn` | `string` | yes |
| `AIRecommendationJobStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `ModelSource` | `AIModelSource` | yes |
| `OutputConfig` | `AIRecommendationOutputResult` | yes |
| `InferenceSpecification` | `AIRecommendationInferenceSpecification` | no |
| `AIWorkloadConfigIdentifier` | `string` | yes |
| `OptimizeModel` | `boolean` | no |
| `PerformanceTarget` | `AIRecommendationPerformanceTarget` | no |
| `Recommendations` | `List<AIRecommendation>` | no |
| `RoleArn` | `string` | yes |
| `ComputeSpec` | `AIRecommendationComputeSpec` | no |
| `AdapterSource` | `AIAdapterSource` | no |
| `CreationTime` | `timestamp` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |

## DescribeAIWorkloadConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigName` | `string` | yes |
| `AIWorkloadConfigArn` | `string` | yes |
| `DatasetConfig` | `AIDatasetConfig` | no |
| `AIWorkloadConfigs` | `AIWorkloadConfigs` | no |
| `Tags` | `List<Tag>` | no |
| `CreationTime` | `timestamp` | yes |

## DescribeAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionName` | `string` | no |
| `ActionArn` | `string` | no |
| `Source` | `ActionSource` | no |
| `ActionType` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `Properties` | `Map<string>` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `LineageGroupArn` | `string` | no |

## DescribeAlgorithm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmName` | `string` | yes |
| `AlgorithmArn` | `string` | yes |
| `AlgorithmDescription` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `TrainingSpecification` | `TrainingSpecification` | yes |
| `InferenceSpecification` | `InferenceSpecification` | no |
| `ValidationSpecification` | `AlgorithmValidationSpecification` | no |
| `AlgorithmStatus` | `string` | yes |
| `AlgorithmStatusDetails` | `AlgorithmStatusDetails` | yes |
| `ProductId` | `string` | no |
| `CertifyForMarketplace` | `boolean` | no |

## DescribeApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | no |
| `SpaceName` | `string` | no |
| `AppType` | `string` | yes |
| `AppName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppArn` | `string` | no |
| `AppType` | `string` | no |
| `AppName` | `string` | no |
| `DomainId` | `string` | no |
| `UserProfileName` | `string` | no |
| `SpaceName` | `string` | no |
| `Status` | `string` | no |
| `EffectiveTrustedIdentityPropagationStatus` | `string` | no |
| `RecoveryMode` | `boolean` | no |
| `LastHealthCheckTimestamp` | `timestamp` | no |
| `LastUserActivityTimestamp` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `ResourceSpec` | `ResourceSpec` | no |
| `BuiltInLifecycleConfigArn` | `string` | no |

## DescribeAppImageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigArn` | `string` | no |
| `AppImageConfigName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `KernelGatewayImageConfig` | `KernelGatewayImageConfig` | no |
| `JupyterLabAppImageConfig` | `JupyterLabAppImageConfig` | no |
| `CodeEditorAppImageConfig` | `CodeEditorAppImageConfig` | no |

## DescribeArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactName` | `string` | no |
| `ArtifactArn` | `string` | no |
| `Source` | `ArtifactSource` | no |
| `ArtifactType` | `string` | no |
| `Properties` | `Map<string>` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `LineageGroupArn` | `string` | no |

## DescribeAutoMLJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |
| `AutoMLJobArn` | `string` | yes |
| `InputDataConfig` | `List<AutoMLChannel>` | yes |
| `OutputDataConfig` | `AutoMLOutputDataConfig` | yes |
| `RoleArn` | `string` | yes |
| `AutoMLJobObjective` | `AutoMLJobObjective` | no |
| `ProblemType` | `string` | no |
| `AutoMLJobConfig` | `AutoMLJobConfig` | no |
| `CreationTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `PartialFailureReasons` | `List<AutoMLPartialFailureReason>` | no |
| `BestCandidate` | `AutoMLCandidate` | no |
| `AutoMLJobStatus` | `string` | yes |
| `AutoMLJobSecondaryStatus` | `string` | yes |
| `GenerateCandidateDefinitionsOnly` | `boolean` | no |
| `AutoMLJobArtifacts` | `AutoMLJobArtifacts` | no |
| `ResolvedAttributes` | `ResolvedAttributes` | no |
| `ModelDeployConfig` | `ModelDeployConfig` | no |
| `ModelDeployResult` | `ModelDeployResult` | no |

## DescribeAutoMLJobV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |
| `AutoMLJobArn` | `string` | yes |
| `AutoMLJobInputDataConfig` | `List<AutoMLJobChannel>` | yes |
| `OutputDataConfig` | `AutoMLOutputDataConfig` | yes |
| `RoleArn` | `string` | yes |
| `AutoMLJobObjective` | `AutoMLJobObjective` | no |
| `AutoMLProblemTypeConfig` | `AutoMLProblemTypeConfig` | no |
| `AutoMLProblemTypeConfigName` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `PartialFailureReasons` | `List<AutoMLPartialFailureReason>` | no |
| `BestCandidate` | `AutoMLCandidate` | no |
| `AutoMLJobStatus` | `string` | yes |
| `AutoMLJobSecondaryStatus` | `string` | yes |
| `AutoMLJobArtifacts` | `AutoMLJobArtifacts` | no |
| `ResolvedAttributes` | `AutoMLResolvedAttributes` | no |
| `ModelDeployConfig` | `ModelDeployConfig` | no |
| `ModelDeployResult` | `ModelDeployResult` | no |
| `DataSplitConfig` | `AutoMLDataSplitConfig` | no |
| `SecurityConfig` | `AutoMLSecurityConfig` | no |
| `AutoMLComputeConfig` | `AutoMLComputeConfig` | no |

## DescribeCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `ClusterName` | `string` | no |
| `ClusterStatus` | `string` | yes |
| `CreationTime` | `timestamp` | no |
| `FailureMessage` | `string` | no |
| `InstanceGroups` | `List<ClusterInstanceGroupDetails>` | yes |
| `RestrictedInstanceGroups` | `List<ClusterRestrictedInstanceGroupDetails>` | no |
| `RestrictedInstanceGroupsConfig` | `ClusterRestrictedInstanceGroupsConfigOutput` | no |
| `VpcConfig` | `VpcConfig` | no |
| `Orchestrator` | `ClusterOrchestrator` | no |
| `TieredStorageConfig` | `ClusterTieredStorageConfig` | no |
| `NodeRecovery` | `string` | no |
| `NodeProvisioningMode` | `string` | no |
| `ClusterRole` | `string` | no |
| `AutoScaling` | `ClusterAutoScalingConfigOutput` | no |

## DescribeClusterEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventId` | `string` | yes |
| `ClusterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDetails` | `ClusterEventDetail` | no |

## DescribeClusterNode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `NodeId` | `string` | no |
| `NodeLogicalId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NodeDetails` | `ClusterNodeDetails` | yes |

## DescribeClusterSchedulerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigId` | `string` | yes |
| `ClusterSchedulerConfigVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigArn` | `string` | yes |
| `ClusterSchedulerConfigId` | `string` | yes |
| `Name` | `string` | yes |
| `ClusterSchedulerConfigVersion` | `integer` | yes |
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `StatusDetails` | `Map<string>` | no |
| `ClusterArn` | `string` | no |
| `SchedulerConfig` | `SchedulerConfig` | no |
| `Description` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |

## DescribeCodeRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryName` | `string` | yes |
| `CodeRepositoryArn` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `GitConfig` | `GitConfig` | no |

## DescribeCompilationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobName` | `string` | yes |
| `CompilationJobArn` | `string` | yes |
| `CompilationJobStatus` | `string` | yes |
| `CompilationStartTime` | `timestamp` | no |
| `CompilationEndTime` | `timestamp` | no |
| `StoppingCondition` | `StoppingCondition` | yes |
| `InferenceImage` | `string` | no |
| `ModelPackageVersionArn` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `FailureReason` | `string` | yes |
| `ModelArtifacts` | `ModelArtifacts` | yes |
| `ModelDigests` | `ModelDigests` | no |
| `RoleArn` | `string` | yes |
| `InputConfig` | `InputConfig` | yes |
| `OutputConfig` | `OutputConfig` | yes |
| `VpcConfig` | `NeoVpcConfig` | no |
| `DerivedInformation` | `DerivedInformation` | no |

## DescribeComputeQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaId` | `string` | yes |
| `ComputeQuotaVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaArn` | `string` | yes |
| `ComputeQuotaId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ComputeQuotaVersion` | `integer` | yes |
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `ClusterArn` | `string` | no |
| `ComputeQuotaConfig` | `ComputeQuotaConfig` | no |
| `ComputeQuotaTarget` | `ComputeQuotaTarget` | yes |
| `ActivationState` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |

## DescribeContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextName` | `string` | no |
| `ContextArn` | `string` | no |
| `Source` | `ContextSource` | no |
| `ContextType` | `string` | no |
| `Description` | `string` | no |
| `Properties` | `Map<string>` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `LineageGroupArn` | `string` | no |

## DescribeDataQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |
| `JobDefinitionName` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `DataQualityBaselineConfig` | `DataQualityBaselineConfig` | no |
| `DataQualityAppSpecification` | `DataQualityAppSpecification` | yes |
| `DataQualityJobInput` | `DataQualityJobInput` | yes |
| `DataQualityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |

## DescribeDevice

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DeviceName` | `string` | yes |
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceArn` | `string` | no |
| `DeviceName` | `string` | yes |
| `Description` | `string` | no |
| `DeviceFleetName` | `string` | yes |
| `IotThingName` | `string` | no |
| `RegistrationTime` | `timestamp` | yes |
| `LatestHeartbeat` | `timestamp` | no |
| `Models` | `List<EdgeModel>` | no |
| `MaxModels` | `integer` | no |
| `NextToken` | `string` | no |
| `AgentVersion` | `string` | no |

## DescribeDeviceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `DeviceFleetArn` | `string` | yes |
| `OutputConfig` | `EdgeOutputConfig` | yes |
| `Description` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `RoleArn` | `string` | no |
| `IotRoleAlias` | `string` | no |

## DescribeDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainArn` | `string` | no |
| `DomainId` | `string` | no |
| `DomainName` | `string` | no |
| `HomeEfsFileSystemId` | `string` | no |
| `SingleSignOnManagedApplicationInstanceId` | `string` | no |
| `SingleSignOnApplicationArn` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `SecurityGroupIdForDomainBoundary` | `string` | no |
| `AuthMode` | `string` | no |
| `DefaultUserSettings` | `UserSettings` | no |
| `DomainSettings` | `DomainSettings` | no |
| `AppNetworkAccessType` | `string` | no |
| `HomeEfsFileSystemKmsKeyId` | `string` | no |
| `SubnetIds` | `List<string>` | no |
| `Url` | `string` | no |
| `VpcId` | `string` | no |
| `KmsKeyId` | `string` | no |
| `AppSecurityGroupManagement` | `string` | no |
| `HomeEfsFileSystemCreation` | `string` | no |
| `TagPropagation` | `string` | no |
| `DefaultSpaceSettings` | `DefaultSpaceSettings` | no |

## DescribeEdgeDeploymentPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanArn` | `string` | yes |
| `EdgeDeploymentPlanName` | `string` | yes |
| `ModelConfigs` | `List<EdgeDeploymentModelConfig>` | yes |
| `DeviceFleetName` | `string` | yes |
| `EdgeDeploymentSuccess` | `integer` | no |
| `EdgeDeploymentPending` | `integer` | no |
| `EdgeDeploymentFailed` | `integer` | no |
| `Stages` | `List<DeploymentStageStatusSummary>` | yes |
| `NextToken` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |

## DescribeEdgePackagingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgePackagingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgePackagingJobArn` | `string` | yes |
| `EdgePackagingJobName` | `string` | yes |
| `CompilationJobName` | `string` | no |
| `ModelName` | `string` | no |
| `ModelVersion` | `string` | no |
| `RoleArn` | `string` | no |
| `OutputConfig` | `EdgeOutputConfig` | no |
| `ResourceKey` | `string` | no |
| `EdgePackagingJobStatus` | `string` | yes |
| `EdgePackagingJobStatusMessage` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `ModelArtifact` | `string` | no |
| `ModelSignature` | `string` | no |
| `PresetDeploymentOutput` | `EdgePresetDeploymentOutput` | no |

## DescribeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `EndpointArn` | `string` | yes |
| `EndpointConfigName` | `string` | no |
| `ProductionVariants` | `List<ProductionVariantSummary>` | no |
| `DataCaptureConfig` | `DataCaptureConfigSummary` | no |
| `EndpointStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `LastDeploymentConfig` | `DeploymentConfig` | no |
| `AsyncInferenceConfig` | `AsyncInferenceConfig` | no |
| `PendingDeploymentSummary` | `PendingDeploymentSummary` | no |
| `ExplainerConfig` | `ExplainerConfig` | no |
| `ShadowProductionVariants` | `List<ProductionVariantSummary>` | no |
| `MetricsConfig` | `MetricsConfig` | no |

## DescribeEndpointConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigName` | `string` | yes |
| `EndpointConfigArn` | `string` | yes |
| `ProductionVariants` | `List<ProductionVariant>` | yes |
| `DataCaptureConfig` | `DataCaptureConfig` | no |
| `KmsKeyId` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `AsyncInferenceConfig` | `AsyncInferenceConfig` | no |
| `ExplainerConfig` | `ExplainerConfig` | no |
| `ShadowProductionVariants` | `List<ProductionVariant>` | no |
| `ExecutionRoleArn` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `EnableNetworkIsolation` | `boolean` | no |
| `MetricsConfig` | `MetricsConfig` | no |

## DescribeExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | no |
| `ExperimentArn` | `string` | no |
| `DisplayName` | `string` | no |
| `Source` | `ExperimentSource` | no |
| `Description` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |

## DescribeFeatureGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupArn` | `string` | yes |
| `FeatureGroupName` | `string` | yes |
| `RecordIdentifierFeatureName` | `string` | yes |
| `EventTimeFeatureName` | `string` | yes |
| `FeatureDefinitions` | `List<FeatureDefinition>` | yes |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | no |
| `OnlineStoreConfig` | `OnlineStoreConfig` | no |
| `OfflineStoreConfig` | `OfflineStoreConfig` | no |
| `ThroughputConfig` | `ThroughputConfigDescription` | no |
| `RoleArn` | `string` | no |
| `FeatureGroupStatus` | `string` | no |
| `OfflineStoreStatus` | `OfflineStoreStatus` | no |
| `LastUpdateStatus` | `LastUpdateStatus` | no |
| `FailureReason` | `string` | no |
| `Description` | `string` | no |
| `NextToken` | `string` | yes |
| `OnlineStoreTotalSizeBytes` | `long` | no |

## DescribeFeatureMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `FeatureName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupArn` | `string` | yes |
| `FeatureGroupName` | `string` | yes |
| `FeatureName` | `string` | yes |
| `FeatureType` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `Description` | `string` | no |
| `Parameters` | `List<FeatureParameter>` | no |

## DescribeFlowDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionArn` | `string` | yes |
| `FlowDefinitionName` | `string` | yes |
| `FlowDefinitionStatus` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `HumanLoopRequestSource` | `HumanLoopRequestSource` | no |
| `HumanLoopActivationConfig` | `HumanLoopActivationConfig` | no |
| `HumanLoopConfig` | `HumanLoopConfig` | no |
| `OutputConfig` | `FlowDefinitionOutputConfig` | yes |
| `RoleArn` | `string` | yes |
| `FailureReason` | `string` | no |

## DescribeHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubArn` | `string` | yes |
| `HubDisplayName` | `string` | no |
| `HubDescription` | `string` | no |
| `HubSearchKeywords` | `List<string>` | no |
| `S3StorageConfig` | `HubS3StorageConfig` | no |
| `HubStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |

## DescribeHubContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentName` | `string` | yes |
| `HubContentVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubContentName` | `string` | yes |
| `HubContentArn` | `string` | yes |
| `HubContentVersion` | `string` | yes |
| `HubContentType` | `string` | yes |
| `DocumentSchemaVersion` | `string` | yes |
| `HubName` | `string` | yes |
| `HubArn` | `string` | yes |
| `HubContentDisplayName` | `string` | no |
| `HubContentDescription` | `string` | no |
| `HubContentMarkdown` | `string` | no |
| `HubContentDocument` | `string` | yes |
| `SageMakerPublicHubContentArn` | `string` | no |
| `ReferenceMinVersion` | `string` | no |
| `SupportStatus` | `string` | no |
| `HubContentSearchKeywords` | `List<string>` | no |
| `HubContentDependencies` | `List<HubContentDependency>` | no |
| `HubContentStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | no |

## DescribeHumanTaskUi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiArn` | `string` | yes |
| `HumanTaskUiName` | `string` | yes |
| `HumanTaskUiStatus` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `UiTemplate` | `UiTemplateInfo` | yes |

## DescribeHyperParameterTuningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |
| `HyperParameterTuningJobArn` | `string` | yes |
| `HyperParameterTuningJobConfig` | `HyperParameterTuningJobConfig` | yes |
| `TrainingJobDefinition` | `HyperParameterTrainingJobDefinition` | no |
| `TrainingJobDefinitions` | `List<HyperParameterTrainingJobDefinition>` | no |
| `HyperParameterTuningJobStatus` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `HyperParameterTuningEndTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `TrainingJobStatusCounters` | `TrainingJobStatusCounters` | yes |
| `ObjectiveStatusCounters` | `ObjectiveStatusCounters` | yes |
| `BestTrainingJob` | `HyperParameterTrainingJobSummary` | no |
| `OverallBestTrainingJob` | `HyperParameterTrainingJobSummary` | no |
| `WarmStartConfig` | `HyperParameterTuningJobWarmStartConfig` | no |
| `Autotune` | `Autotune` | no |
| `FailureReason` | `string` | no |
| `TuningJobCompletionDetails` | `HyperParameterTuningJobCompletionDetails` | no |
| `ConsumedResources` | `HyperParameterTuningJobConsumedResources` | no |

## DescribeImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTime` | `timestamp` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `FailureReason` | `string` | no |
| `ImageArn` | `string` | no |
| `ImageName` | `string` | no |
| `ImageStatus` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `RoleArn` | `string` | no |

## DescribeImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `Version` | `integer` | no |
| `Alias` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BaseImage` | `string` | no |
| `ContainerImage` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `ImageArn` | `string` | no |
| `ImageVersionArn` | `string` | no |
| `ImageVersionStatus` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `Version` | `integer` | no |
| `VendorGuidance` | `string` | no |
| `JobType` | `string` | no |
| `MLFramework` | `string` | no |
| `ProgrammingLang` | `string` | no |
| `Processor` | `string` | no |
| `Horovod` | `boolean` | no |
| `ReleaseNotes` | `string` | no |

## DescribeInferenceComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |
| `InferenceComponentArn` | `string` | yes |
| `EndpointName` | `string` | yes |
| `EndpointArn` | `string` | yes |
| `VariantName` | `string` | no |
| `FailureReason` | `string` | no |
| `Specification` | `InferenceComponentSpecificationSummary` | no |
| `Specifications` | `List<InferenceComponentSpecificationSummary>` | no |
| `RuntimeConfig` | `InferenceComponentRuntimeConfigSummary` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `InferenceComponentStatus` | `string` | no |
| `LastDeploymentConfig` | `InferenceComponentDeploymentConfig` | no |

## DescribeInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Schedule` | `InferenceExperimentSchedule` | no |
| `Status` | `string` | yes |
| `StatusReason` | `string` | no |
| `Description` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `CompletionTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `RoleArn` | `string` | no |
| `EndpointMetadata` | `EndpointMetadata` | yes |
| `ModelVariants` | `List<ModelVariantConfigSummary>` | yes |
| `DataStorageConfig` | `InferenceExperimentDataStorageConfig` | no |
| `ShadowModeConfig` | `ShadowModeConfig` | no |
| `KmsKey` | `string` | no |

## DescribeInferenceRecommendationsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobDescription` | `string` | no |
| `JobType` | `string` | yes |
| `JobArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Status` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `CompletionTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `InputConfig` | `RecommendationJobInputConfig` | yes |
| `StoppingConditions` | `RecommendationJobStoppingConditions` | no |
| `InferenceRecommendations` | `List<InferenceRecommendation>` | no |
| `EndpointPerformances` | `List<EndpointPerformance>` | no |

## DescribeJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobCategory` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobArn` | `string` | yes |
| `RoleArn` | `string` | yes |
| `JobCategory` | `string` | yes |
| `JobConfigSchemaVersion` | `string` | yes |
| `JobConfigDocument` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | no |
| `JobStatus` | `string` | yes |
| `SecondaryStatus` | `string` | yes |
| `SecondaryStatusTransitions` | `List<JobSecondaryStatusTransition>` | yes |
| `FailureReason` | `string` | no |
| `Tags` | `List<Tag>` | no |

## DescribeJobSchemaVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobCategory` | `string` | yes |
| `JobConfigSchemaVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobCategory` | `string` | yes |
| `JobConfigSchemaVersion` | `string` | yes |
| `JobConfigSchema` | `string` | yes |

## DescribeLabelingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobStatus` | `string` | yes |
| `LabelCounters` | `LabelCounters` | yes |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `JobReferenceCode` | `string` | yes |
| `LabelingJobName` | `string` | yes |
| `LabelingJobArn` | `string` | yes |
| `LabelAttributeName` | `string` | no |
| `InputConfig` | `LabelingJobInputConfig` | yes |
| `OutputConfig` | `LabelingJobOutputConfig` | yes |
| `RoleArn` | `string` | yes |
| `LabelCategoryConfigS3Uri` | `string` | no |
| `StoppingConditions` | `LabelingJobStoppingConditions` | no |
| `LabelingJobAlgorithmsConfig` | `LabelingJobAlgorithmsConfig` | no |
| `HumanTaskConfig` | `HumanTaskConfig` | yes |
| `Tags` | `List<Tag>` | no |
| `LabelingJobOutput` | `LabelingJobOutput` | no |

## DescribeLineageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LineageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LineageGroupName` | `string` | no |
| `LineageGroupArn` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |

## DescribeMlflowApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `ArtifactStoreUri` | `string` | no |
| `MlflowVersion` | `string` | no |
| `RoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `Status` | `string` | no |
| `ModelRegistrationMode` | `string` | no |
| `AccountDefaultStatus` | `string` | no |
| `DefaultDomainIdList` | `List<string>` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `MaintenanceStatus` | `string` | no |

## DescribeMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |
| `TrackingServerName` | `string` | no |
| `ArtifactStoreUri` | `string` | no |
| `TrackingServerSize` | `string` | no |
| `MlflowVersion` | `string` | no |
| `RoleArn` | `string` | no |
| `TrackingServerStatus` | `string` | no |
| `TrackingServerMaintenanceStatus` | `string` | no |
| `IsActive` | `string` | no |
| `TrackingServerUrl` | `string` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `AutomaticModelRegistration` | `boolean` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `S3BucketOwnerAccountId` | `string` | no |
| `S3BucketOwnerVerification` | `boolean` | no |

## DescribeModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelName` | `string` | yes |
| `PrimaryContainer` | `ContainerDefinition` | no |
| `Containers` | `List<ContainerDefinition>` | no |
| `InferenceExecutionConfig` | `InferenceExecutionConfig` | no |
| `ExecutionRoleArn` | `string` | no |
| `VpcConfig` | `VpcConfig` | no |
| `CreationTime` | `timestamp` | yes |
| `ModelArn` | `string` | yes |
| `EnableNetworkIsolation` | `boolean` | no |
| `DeploymentRecommendation` | `DeploymentRecommendation` | no |

## DescribeModelBiasJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |
| `JobDefinitionName` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `ModelBiasBaselineConfig` | `ModelBiasBaselineConfig` | no |
| `ModelBiasAppSpecification` | `ModelBiasAppSpecification` | yes |
| `ModelBiasJobInput` | `ModelBiasJobInput` | yes |
| `ModelBiasJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |

## DescribeModelCard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |
| `ModelCardVersion` | `integer` | no |
| `IncludedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardArn` | `string` | yes |
| `ModelCardName` | `string` | yes |
| `ModelCardVersion` | `integer` | yes |
| `Content` | `string` | yes |
| `ModelCardStatus` | `string` | yes |
| `SecurityConfig` | `ModelCardSecurityConfig` | no |
| `CreationTime` | `timestamp` | yes |
| `CreatedBy` | `UserContext` | yes |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `ModelCardProcessingStatus` | `string` | no |

## DescribeModelCardExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardExportJobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardExportJobName` | `string` | yes |
| `ModelCardExportJobArn` | `string` | yes |
| `Status` | `string` | yes |
| `ModelCardName` | `string` | yes |
| `ModelCardVersion` | `integer` | yes |
| `OutputConfig` | `ModelCardExportOutputConfig` | yes |
| `CreatedAt` | `timestamp` | yes |
| `LastModifiedAt` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `ExportArtifacts` | `ModelCardExportArtifacts` | no |

## DescribeModelExplainabilityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |
| `JobDefinitionName` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `ModelExplainabilityBaselineConfig` | `ModelExplainabilityBaselineConfig` | no |
| `ModelExplainabilityAppSpecification` | `ModelExplainabilityAppSpecification` | yes |
| `ModelExplainabilityJobInput` | `ModelExplainabilityJobInput` | yes |
| `ModelExplainabilityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |

## DescribeModelPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageName` | `string` | yes |
| `IncludedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageName` | `string` | yes |
| `ModelPackageGroupName` | `string` | no |
| `ModelPackageVersion` | `integer` | no |
| `ModelPackageRegistrationType` | `string` | no |
| `ModelPackageArn` | `string` | yes |
| `ModelPackageDescription` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `InferenceSpecification` | `InferenceSpecification` | no |
| `SourceAlgorithmSpecification` | `SourceAlgorithmSpecification` | no |
| `ValidationSpecification` | `ModelPackageValidationSpecification` | no |
| `ModelPackageStatus` | `string` | yes |
| `ModelPackageStatusDetails` | `ModelPackageStatusDetails` | yes |
| `CertifyForMarketplace` | `boolean` | no |
| `ModelApprovalStatus` | `string` | no |
| `CreatedBy` | `UserContext` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `ModelMetrics` | `ModelMetrics` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `ApprovalDescription` | `string` | no |
| `Domain` | `string` | no |
| `Task` | `string` | no |
| `SamplePayloadUrl` | `string` | no |
| `CustomerMetadataProperties` | `Map<string>` | no |
| `DriftCheckBaselines` | `DriftCheckBaselines` | no |
| `AdditionalInferenceSpecifications` | `List<AdditionalInferenceSpecificationDefinition>` | no |
| `SkipModelValidation` | `string` | no |
| `SourceUri` | `string` | no |
| `SecurityConfig` | `ModelPackageSecurityConfig` | no |
| `ModelCard` | `ModelPackageModelCard` | no |
| `ModelLifeCycle` | `ModelLifeCycle` | no |
| `ManagedStorageType` | `string` | no |

## DescribeModelPackageGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |
| `ModelPackageGroupArn` | `string` | yes |
| `ModelPackageGroupDescription` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `CreatedBy` | `UserContext` | yes |
| `ModelPackageGroupStatus` | `string` | yes |
| `ManagedConfiguration` | `ManagedConfiguration` | no |

## DescribeModelQualityJobDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionArn` | `string` | yes |
| `JobDefinitionName` | `string` | yes |
| `CreationTime` | `timestamp` | yes |
| `ModelQualityBaselineConfig` | `ModelQualityBaselineConfig` | no |
| `ModelQualityAppSpecification` | `ModelQualityAppSpecification` | yes |
| `ModelQualityJobInput` | `ModelQualityJobInput` | yes |
| `ModelQualityJobOutputConfig` | `MonitoringOutputConfig` | yes |
| `JobResources` | `MonitoringResources` | yes |
| `NetworkConfig` | `MonitoringNetworkConfig` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `MonitoringStoppingCondition` | no |

## DescribeMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleArn` | `string` | yes |
| `MonitoringScheduleName` | `string` | yes |
| `MonitoringScheduleStatus` | `string` | yes |
| `MonitoringType` | `string` | no |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `MonitoringScheduleConfig` | `MonitoringScheduleConfig` | yes |
| `EndpointName` | `string` | no |
| `LastMonitoringExecutionSummary` | `MonitoringExecutionSummary` | no |

## DescribeNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceArn` | `string` | no |
| `NotebookInstanceName` | `string` | no |
| `NotebookInstanceStatus` | `string` | no |
| `FailureReason` | `string` | no |
| `Url` | `string` | no |
| `InstanceType` | `string` | no |
| `IpAddressType` | `string` | no |
| `SubnetId` | `string` | no |
| `SecurityGroups` | `List<string>` | no |
| `RoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `NetworkInterfaceId` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `NotebookInstanceLifecycleConfigName` | `string` | no |
| `DirectInternetAccess` | `string` | no |
| `VolumeSizeInGB` | `integer` | no |
| `AcceleratorTypes` | `List<string>` | no |
| `DefaultCodeRepository` | `string` | no |
| `AdditionalCodeRepositories` | `List<string>` | no |
| `RootAccess` | `string` | no |
| `PlatformIdentifier` | `string` | no |
| `InstanceMetadataServiceConfiguration` | `InstanceMetadataServiceConfiguration` | no |

## DescribeNotebookInstanceLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigArn` | `string` | no |
| `NotebookInstanceLifecycleConfigName` | `string` | no |
| `OnCreate` | `List<NotebookInstanceLifecycleHook>` | no |
| `OnStart` | `List<NotebookInstanceLifecycleHook>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |

## DescribeOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobArn` | `string` | yes |
| `OptimizationJobStatus` | `string` | yes |
| `OptimizationStartTime` | `timestamp` | no |
| `OptimizationEndTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | yes |
| `FailureReason` | `string` | no |
| `OptimizationJobName` | `string` | yes |
| `ModelSource` | `OptimizationJobModelSource` | yes |
| `OptimizationEnvironment` | `Map<string>` | no |
| `DeploymentInstanceType` | `string` | yes |
| `MaxInstanceCount` | `integer` | no |
| `OptimizationConfigs` | `List<OptimizationConfig>` | yes |
| `OutputConfig` | `OptimizationJobOutputConfig` | yes |
| `OptimizationOutput` | `OptimizationOutput` | no |
| `RoleArn` | `string` | yes |
| `StoppingCondition` | `StoppingCondition` | yes |
| `VpcConfig` | `OptimizationVpcConfig` | no |
| `TrainingPlanArns` | `List<string>` | no |

## DescribePartnerApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `IncludeAvailableUpgrade` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Type` | `string` | no |
| `Status` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `ExecutionRoleArn` | `string` | no |
| `KmsKeyId` | `string` | no |
| `BaseUrl` | `string` | no |
| `MaintenanceConfig` | `PartnerAppMaintenanceConfig` | no |
| `Tier` | `string` | no |
| `Version` | `string` | no |
| `ApplicationConfig` | `PartnerAppConfig` | no |
| `AuthType` | `string` | no |
| `EnableIamSessionBasedIdentity` | `boolean` | no |
| `Error` | `ErrorInfo` | no |
| `EnableAutoMinorVersionUpgrade` | `boolean` | no |
| `CurrentVersionEolDate` | `timestamp` | no |
| `AvailableUpgrade` | `AvailableUpgrade` | no |
| `IdcConfig` | `IdcConfigOutput` | no |

## DescribePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `PipelineVersionId` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |
| `PipelineName` | `string` | no |
| `PipelineDisplayName` | `string` | no |
| `PipelineDefinition` | `string` | no |
| `PipelineDescription` | `string` | no |
| `RoleArn` | `string` | no |
| `PipelineStatus` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastRunTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedBy` | `UserContext` | no |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |
| `PipelineVersionDisplayName` | `string` | no |
| `PipelineVersionDescription` | `string` | no |

## DescribePipelineDefinitionForExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineDefinition` | `string` | no |
| `CreationTime` | `timestamp` | no |

## DescribePipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |
| `PipelineExecutionArn` | `string` | no |
| `PipelineExecutionDisplayName` | `string` | no |
| `PipelineExecutionStatus` | `string` | no |
| `PipelineExecutionDescription` | `string` | no |
| `PipelineExperimentConfig` | `PipelineExperimentConfig` | no |
| `FailureReason` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedBy` | `UserContext` | no |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |
| `SelectiveExecutionConfig` | `SelectiveExecutionConfig` | no |
| `PipelineVersionId` | `long` | no |
| `MLflowConfig` | `MLflowConfiguration` | no |

## DescribeProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingInputs` | `List<ProcessingInput>` | no |
| `ProcessingOutputConfig` | `ProcessingOutputConfig` | no |
| `ProcessingJobName` | `string` | yes |
| `ProcessingResources` | `ProcessingResources` | yes |
| `StoppingCondition` | `ProcessingStoppingCondition` | no |
| `AppSpecification` | `AppSpecification` | yes |
| `Environment` | `Map<string>` | no |
| `NetworkConfig` | `NetworkConfig` | no |
| `RoleArn` | `string` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |
| `ProcessingJobArn` | `string` | yes |
| `ProcessingJobStatus` | `string` | yes |
| `ExitMessage` | `string` | no |
| `FailureReason` | `string` | no |
| `ProcessingEndTime` | `timestamp` | no |
| `ProcessingStartTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | yes |
| `MonitoringScheduleArn` | `string` | no |
| `AutoMLJobArn` | `string` | no |
| `TrainingJobArn` | `string` | no |

## DescribeProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |
| `ProjectName` | `string` | yes |
| `ProjectId` | `string` | yes |
| `ProjectDescription` | `string` | no |
| `ServiceCatalogProvisioningDetails` | `ServiceCatalogProvisioningDetails` | no |
| `ServiceCatalogProvisionedProductDetails` | `ServiceCatalogProvisionedProductDetails` | no |
| `ProjectStatus` | `string` | yes |
| `TemplateProviderDetails` | `List<TemplateProviderDetail>` | no |
| `CreatedBy` | `UserContext` | no |
| `CreationTime` | `timestamp` | yes |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |

## DescribeReservedCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCapacityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCapacityArn` | `string` | yes |
| `ReservedCapacityType` | `string` | no |
| `Status` | `string` | no |
| `AvailabilityZone` | `string` | no |
| `DurationHours` | `long` | no |
| `DurationMinutes` | `long` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `InstanceType` | `string` | yes |
| `TotalInstanceCount` | `integer` | yes |
| `AvailableInstanceCount` | `integer` | no |
| `InUseInstanceCount` | `integer` | no |
| `UltraServerSummary` | `UltraServerSummary` | no |

## DescribeSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpaceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | no |
| `SpaceArn` | `string` | no |
| `SpaceName` | `string` | no |
| `HomeEfsFileSystemUid` | `string` | no |
| `Status` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `SpaceSettings` | `SpaceSettings` | no |
| `OwnershipSettings` | `OwnershipSettings` | no |
| `SpaceSharingSettings` | `SpaceSharingSettings` | no |
| `SpaceDisplayName` | `string` | no |
| `Url` | `string` | no |

## DescribeStudioLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioLifecycleConfigName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StudioLifecycleConfigArn` | `string` | no |
| `StudioLifecycleConfigName` | `string` | no |
| `CreationTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `StudioLifecycleConfigContent` | `string` | no |
| `StudioLifecycleConfigAppType` | `string` | no |

## DescribeSubscribedWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscribedWorkteam` | `SubscribedWorkteam` | yes |

## DescribeTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |
| `TrainingJobArn` | `string` | yes |
| `TuningJobArn` | `string` | no |
| `LabelingJobArn` | `string` | no |
| `AutoMLJobArn` | `string` | no |
| `ModelArtifacts` | `ModelArtifacts` | yes |
| `TrainingJobStatus` | `string` | yes |
| `SecondaryStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `HyperParameters` | `Map<string>` | no |
| `AlgorithmSpecification` | `AlgorithmSpecification` | no |
| `RoleArn` | `string` | no |
| `InputDataConfig` | `List<Channel>` | no |
| `OutputDataConfig` | `OutputDataConfig` | no |
| `ResourceConfig` | `ResourceConfig` | no |
| `WarmPoolStatus` | `WarmPoolStatus` | no |
| `VpcConfig` | `VpcConfig` | no |
| `StoppingCondition` | `StoppingCondition` | yes |
| `CreationTime` | `timestamp` | yes |
| `TrainingStartTime` | `timestamp` | no |
| `TrainingEndTime` | `timestamp` | no |
| `LastModifiedTime` | `timestamp` | no |
| `SecondaryStatusTransitions` | `List<SecondaryStatusTransition>` | no |
| `FinalMetricDataList` | `List<MetricData>` | no |
| `EnableNetworkIsolation` | `boolean` | no |
| `EnableInterContainerTrafficEncryption` | `boolean` | no |
| `EnableManagedSpotTraining` | `boolean` | no |
| `CheckpointConfig` | `CheckpointConfig` | no |
| `TrainingTimeInSeconds` | `integer` | no |
| `BillableTimeInSeconds` | `integer` | no |
| `BillableTokenCount` | `long` | no |
| `DebugHookConfig` | `DebugHookConfig` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |
| `DebugRuleConfigurations` | `List<DebugRuleConfiguration>` | no |
| `TensorBoardOutputConfig` | `TensorBoardOutputConfig` | no |
| `DebugRuleEvaluationStatuses` | `List<DebugRuleEvaluationStatus>` | no |
| `ProfilerConfig` | `ProfilerConfig` | no |
| `ProfilerRuleConfigurations` | `List<ProfilerRuleConfiguration>` | no |
| `ProfilerRuleEvaluationStatuses` | `List<ProfilerRuleEvaluationStatus>` | no |
| `ProfilingStatus` | `string` | no |
| `Environment` | `Map<string>` | no |
| `RetryStrategy` | `RetryStrategy` | no |
| `RemoteDebugConfig` | `RemoteDebugConfig` | no |
| `InfraCheckConfig` | `InfraCheckConfig` | no |
| `ServerlessJobConfig` | `ServerlessJobConfig` | no |
| `MlflowConfig` | `MlflowConfig` | no |
| `ModelPackageConfig` | `ModelPackageConfig` | no |
| `MlflowDetails` | `MlflowDetails` | no |
| `ProgressInfo` | `TrainingProgressInfo` | no |
| `OutputModelPackageArn` | `string` | no |

## DescribeTrainingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanArn` | `string` | yes |
| `TrainingPlanName` | `string` | yes |
| `Status` | `string` | yes |
| `StatusMessage` | `string` | no |
| `DurationHours` | `long` | no |
| `DurationMinutes` | `long` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `UpfrontFee` | `string` | no |
| `CurrencyCode` | `string` | no |
| `TotalInstanceCount` | `integer` | no |
| `AvailableInstanceCount` | `integer` | no |
| `InUseInstanceCount` | `integer` | no |
| `UnhealthyInstanceCount` | `integer` | no |
| `AvailableSpareInstanceCount` | `integer` | no |
| `TotalUltraServerCount` | `integer` | no |
| `TargetResources` | `List<string>` | no |
| `ReservedCapacitySummaries` | `List<ReservedCapacitySummary>` | no |

## DescribeTrainingPlanExtensionHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanExtensions` | `List<TrainingPlanExtension>` | yes |
| `NextToken` | `string` | no |

## DescribeTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobName` | `string` | yes |
| `TransformJobArn` | `string` | yes |
| `TransformJobStatus` | `string` | yes |
| `FailureReason` | `string` | no |
| `ModelName` | `string` | yes |
| `MaxConcurrentTransforms` | `integer` | no |
| `ModelClientConfig` | `ModelClientConfig` | no |
| `MaxPayloadInMB` | `integer` | no |
| `BatchStrategy` | `string` | no |
| `Environment` | `Map<string>` | no |
| `TransformInput` | `TransformInput` | yes |
| `TransformOutput` | `TransformOutput` | no |
| `DataCaptureConfig` | `BatchDataCaptureConfig` | no |
| `TransformResources` | `TransformResources` | yes |
| `CreationTime` | `timestamp` | yes |
| `TransformStartTime` | `timestamp` | no |
| `TransformEndTime` | `timestamp` | no |
| `LabelingJobArn` | `string` | no |
| `AutoMLJobArn` | `string` | no |
| `DataProcessing` | `DataProcessing` | no |
| `ExperimentConfig` | `ExperimentConfig` | no |

## DescribeTrial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialName` | `string` | no |
| `TrialArn` | `string` | no |
| `DisplayName` | `string` | no |
| `ExperimentName` | `string` | no |
| `Source` | `TrialSource` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `MetadataProperties` | `MetadataProperties` | no |

## DescribeTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | no |
| `TrialComponentArn` | `string` | no |
| `DisplayName` | `string` | no |
| `Source` | `TrialComponentSource` | no |
| `Status` | `TrialComponentStatus` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `CreatedBy` | `UserContext` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedBy` | `UserContext` | no |
| `Parameters` | `Map<TrialComponentParameterValue>` | no |
| `InputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `OutputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `MetadataProperties` | `MetadataProperties` | no |
| `Metrics` | `List<TrialComponentMetricSummary>` | no |
| `LineageGroupArn` | `string` | no |
| `Sources` | `List<TrialComponentSource>` | no |

## DescribeUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | no |
| `UserProfileArn` | `string` | no |
| `UserProfileName` | `string` | no |
| `HomeEfsFileSystemUid` | `string` | no |
| `Status` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreationTime` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `SingleSignOnUserIdentifier` | `string` | no |
| `SingleSignOnUserValue` | `string` | no |
| `UserSettings` | `UserSettings` | no |

## DescribeWorkforce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkforceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workforce` | `Workforce` | yes |

## DescribeWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workteam` | `Workteam` | yes |

## DetachClusterNodeVolume

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `NodeId` | `string` | yes |
| `VolumeId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |
| `NodeId` | `string` | yes |
| `VolumeId` | `string` | yes |
| `AttachTime` | `timestamp` | yes |
| `Status` | `string` | yes |
| `DeviceName` | `string` | yes |

## DisableSagemakerServicecatalogPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |
| `TrialName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentArn` | `string` | no |
| `TrialArn` | `string` | no |

## EnableSagemakerServicecatalogPortfolio

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExtendTrainingPlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanExtensionOfferingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanExtensions` | `List<TrainingPlanExtension>` | yes |

## GetDeviceFleetReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetArn` | `string` | yes |
| `DeviceFleetName` | `string` | yes |
| `OutputConfig` | `EdgeOutputConfig` | no |
| `Description` | `string` | no |
| `ReportGenerated` | `timestamp` | no |
| `DeviceStats` | `DeviceStats` | no |
| `AgentVersions` | `List<AgentVersion>` | no |
| `ModelStats` | `List<EdgeModelStat>` | no |

## GetLineageGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LineageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LineageGroupArn` | `string` | no |
| `ResourcePolicy` | `string` | no |

## GetModelPackageGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourcePolicy` | `string` | yes |

## GetSagemakerServicecatalogPortfolioStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## GetScalingConfigurationRecommendation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceRecommendationsJobName` | `string` | yes |
| `RecommendationId` | `string` | no |
| `EndpointName` | `string` | no |
| `TargetCpuUtilizationPerCore` | `integer` | no |
| `ScalingPolicyObjective` | `ScalingPolicyObjective` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceRecommendationsJobName` | `string` | no |
| `RecommendationId` | `string` | no |
| `EndpointName` | `string` | no |
| `TargetCpuUtilizationPerCore` | `integer` | no |
| `ScalingPolicyObjective` | `ScalingPolicyObjective` | no |
| `Metric` | `ScalingPolicyMetric` | no |
| `DynamicScalingConfiguration` | `DynamicScalingConfiguration` | no |

## GetSearchSuggestions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `SuggestionQuery` | `SuggestionQuery` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PropertyNameSuggestions` | `List<PropertyNameSuggestion>` | no |

## ImportHubContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubContentName` | `string` | yes |
| `HubContentVersion` | `string` | no |
| `HubContentType` | `string` | yes |
| `DocumentSchemaVersion` | `string` | yes |
| `HubName` | `string` | yes |
| `HubContentDisplayName` | `string` | no |
| `HubContentDescription` | `string` | no |
| `HubContentMarkdown` | `string` | no |
| `HubContentDocument` | `string` | yes |
| `SupportStatus` | `string` | no |
| `HubContentSearchKeywords` | `List<string>` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |
| `HubContentArn` | `string` | yes |

## ListAIBenchmarkJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobs` | `List<AIBenchmarkJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListAIRecommendationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobs` | `List<AIRecommendationJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListAIWorkloadConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIWorkloadConfigs` | `List<AIWorkloadConfigSummary>` | yes |
| `NextToken` | `string` | no |

## ListActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceUri` | `string` | no |
| `ActionType` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionSummaries` | `List<ActionSummary>` | no |
| `NextToken` | `string` | no |

## ListAlgorithms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AlgorithmSummaryList` | `List<AlgorithmSummary>` | yes |
| `NextToken` | `string` | no |

## ListAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `Alias` | `string` | no |
| `Version` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SageMakerImageVersionAliases` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAppImageConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `ModifiedTimeBefore` | `timestamp` | no |
| `ModifiedTimeAfter` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AppImageConfigs` | `List<AppImageConfigDetails>` | no |

## ListApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `DomainIdEquals` | `string` | no |
| `UserProfileNameEquals` | `string` | no |
| `SpaceNameEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Apps` | `List<AppDetails>` | no |
| `NextToken` | `string` | no |

## ListArtifacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceUri` | `string` | no |
| `ArtifactType` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactSummaries` | `List<ArtifactSummary>` | no |
| `NextToken` | `string` | no |

## ListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceArn` | `string` | no |
| `DestinationArn` | `string` | no |
| `SourceType` | `string` | no |
| `DestinationType` | `string` | no |
| `AssociationType` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationSummaries` | `List<AssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListAutoMLJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobSummaries` | `List<AutoMLJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListCandidatesForAutoMLJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |
| `StatusEquals` | `string` | no |
| `CandidateNameEquals` | `string` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Candidates` | `List<AutoMLCandidate>` | yes |
| `NextToken` | `string` | no |

## ListClusterEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `InstanceGroupName` | `string` | no |
| `NodeId` | `string` | no |
| `EventTimeAfter` | `timestamp` | no |
| `EventTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `ResourceType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Events` | `List<ClusterEventSummary>` | no |

## ListClusterNodes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `InstanceGroupNameContains` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `IncludeNodeLogicalIds` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ClusterNodeSummaries` | `List<ClusterNodeSummary>` | yes |

## ListClusterSchedulerConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `ClusterArn` | `string` | no |
| `Status` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigSummaries` | `List<ClusterSchedulerConfigSummary>` | no |
| `NextToken` | `string` | no |

## ListClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `TrainingPlanArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ClusterSummaries` | `List<ClusterSummary>` | yes |

## ListCodeRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositorySummaryList` | `List<CodeRepositorySummary>` | yes |
| `NextToken` | `string` | no |

## ListCompilationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobSummaries` | `List<CompilationJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListComputeQuotas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `Status` | `string` | no |
| `ClusterArn` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaSummaries` | `List<ComputeQuotaSummary>` | no |
| `NextToken` | `string` | no |

## ListContexts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceUri` | `string` | no |
| `ContextType` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextSummaries` | `List<ContextSummary>` | no |
| `NextToken` | `string` | no |

## ListDataQualityJobDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionSummaries` | `List<MonitoringJobDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListDeviceFleets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetSummaries` | `List<DeviceFleetSummary>` | yes |
| `NextToken` | `string` | no |

## ListDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LatestHeartbeatAfter` | `timestamp` | no |
| `ModelName` | `string` | no |
| `DeviceFleetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceSummaries` | `List<DeviceSummary>` | yes |
| `NextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domains` | `List<DomainDetails>` | no |
| `NextToken` | `string` | no |

## ListEdgeDeploymentPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `DeviceFleetNameContains` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanSummaries` | `List<EdgeDeploymentPlanSummary>` | yes |
| `NextToken` | `string` | no |

## ListEdgePackagingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `ModelNameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgePackagingJobSummaries` | `List<EdgePackagingJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListEndpointConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointConfigs` | `List<EndpointConfigSummary>` | yes |
| `NextToken` | `string` | no |

## ListEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<EndpointSummary>` | yes |
| `NextToken` | `string` | no |

## ListExperiments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentSummaries` | `List<ExperimentSummary>` | no |
| `NextToken` | `string` | no |

## ListFeatureGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NameContains` | `string` | no |
| `FeatureGroupStatusEquals` | `string` | no |
| `OfflineStoreStatusEquals` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupSummaries` | `List<FeatureGroupSummary>` | yes |
| `NextToken` | `string` | no |

## ListFlowDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowDefinitionSummaries` | `List<FlowDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListHubContentVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentName` | `string` | yes |
| `MinVersion` | `string` | no |
| `MaxSchemaVersion` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubContentSummaries` | `List<HubContentInfo>` | yes |
| `NextToken` | `string` | no |

## ListHubContents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `NameContains` | `string` | no |
| `MaxSchemaVersion` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubContentSummaries` | `List<HubContentInfo>` | yes |
| `NextToken` | `string` | no |

## ListHubs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubSummaries` | `List<HubInfo>` | yes |
| `NextToken` | `string` | no |

## ListHumanTaskUis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HumanTaskUiSummaries` | `List<HumanTaskUiSummary>` | yes |
| `NextToken` | `string` | no |

## ListHyperParameterTuningJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NameContains` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobSummaries` | `List<HyperParameterTuningJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListImageVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `ImageName` | `string` | yes |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageVersions` | `List<ImageVersion>` | no |
| `NextToken` | `string` | no |

## ListImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Images` | `List<Image>` | no |
| `NextToken` | `string` | no |

## ListInferenceComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |
| `EndpointNameEquals` | `string` | no |
| `VariantNameEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponents` | `List<InferenceComponentSummary>` | yes |
| `NextToken` | `string` | no |

## ListInferenceExperiments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NameContains` | `string` | no |
| `Type` | `string` | no |
| `StatusEquals` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperiments` | `List<InferenceExperimentSummary>` | no |
| `NextToken` | `string` | no |

## ListInferenceRecommendationsJobSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `Status` | `string` | no |
| `StepType` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Steps` | `List<InferenceRecommendationsJobStep>` | no |
| `NextToken` | `string` | no |

## ListInferenceRecommendationsJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ModelNameEquals` | `string` | no |
| `ModelPackageVersionArnEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceRecommendationsJobs` | `List<InferenceRecommendationsJob>` | yes |
| `NextToken` | `string` | no |

## ListJobSchemaVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobCategory` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `JobConfigSchemas` | `List<JobConfigSchemaVersionSummary>` | yes |

## ListJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobCategory` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `JobSummaries` | `List<JobSummary>` | yes |

## ListLabelingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobSummaryList` | `List<LabelingJobSummary>` | no |
| `NextToken` | `string` | no |

## ListLabelingJobsForWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `JobReferenceCodeContains` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobSummaryList` | `List<LabelingJobForWorkteamSummary>` | yes |
| `NextToken` | `string` | no |

## ListLineageGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LineageGroupSummaries` | `List<LineageGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListMlflowApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `Status` | `string` | no |
| `MlflowVersion` | `string` | no |
| `DefaultForDomainId` | `string` | no |
| `AccountDefaultStatus` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<MlflowAppSummary>` | no |
| `NextToken` | `string` | no |

## ListMlflowTrackingServers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `TrackingServerStatus` | `string` | no |
| `MlflowVersion` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerSummaries` | `List<TrackingServerSummary>` | no |
| `NextToken` | `string` | no |

## ListModelBiasJobDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionSummaries` | `List<MonitoringJobDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelCardExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |
| `ModelCardVersion` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `ModelCardExportJobNameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardExportJobSummaries` | `List<ModelCardExportJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelCardVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `ModelCardName` | `string` | yes |
| `ModelCardStatus` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardVersionSummaryList` | `List<ModelCardVersionSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelCards

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `ModelCardStatus` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardSummaries` | `List<ModelCardSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelExplainabilityJobDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionSummaries` | `List<MonitoringJobDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SearchExpression` | `ModelMetadataSearchExpression` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelMetadataSummaries` | `List<ModelMetadataSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelPackageGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `CrossAccountFilterOption` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupSummaryList` | `List<ModelPackageGroupSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelPackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `ModelApprovalStatus` | `string` | no |
| `ModelPackageGroupName` | `string` | no |
| `ModelPackageType` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageSummaryList` | `List<ModelPackageSummary>` | yes |
| `NextToken` | `string` | no |

## ListModelQualityJobDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobDefinitionSummaries` | `List<MonitoringJobDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Models` | `List<ModelSummary>` | yes |
| `NextToken` | `string` | no |

## ListMonitoringAlertHistory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | no |
| `MonitoringAlertName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringAlertHistory` | `List<MonitoringAlertHistorySummary>` | no |
| `NextToken` | `string` | no |

## ListMonitoringAlerts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringAlertSummaries` | `List<MonitoringAlertSummary>` | no |
| `NextToken` | `string` | no |

## ListMonitoringExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | no |
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ScheduledTimeBefore` | `timestamp` | no |
| `ScheduledTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |
| `MonitoringJobDefinitionName` | `string` | no |
| `MonitoringTypeEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringExecutionSummaries` | `List<MonitoringExecutionSummary>` | yes |
| `NextToken` | `string` | no |

## ListMonitoringSchedules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |
| `MonitoringJobDefinitionName` | `string` | no |
| `MonitoringTypeEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleSummaries` | `List<MonitoringScheduleSummary>` | yes |
| `NextToken` | `string` | no |

## ListNotebookInstanceLifecycleConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NotebookInstanceLifecycleConfigs` | `List<NotebookInstanceLifecycleConfigSummary>` | no |

## ListNotebookInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NameContains` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `StatusEquals` | `string` | no |
| `NotebookInstanceLifecycleConfigNameContains` | `string` | no |
| `DefaultCodeRepositoryContains` | `string` | no |
| `AdditionalCodeRepositoryEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NotebookInstances` | `List<NotebookInstanceSummary>` | no |

## ListOptimizationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `OptimizationContains` | `string` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobSummaries` | `List<OptimizationJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListPartnerApps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Summaries` | `List<PartnerAppSummary>` | no |
| `NextToken` | `string` | no |

## ListPipelineExecutionSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionSteps` | `List<PipelineExecutionStep>` | no |
| `NextToken` | `string` | no |

## ListPipelineExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionSummaries` | `List<PipelineExecutionSummary>` | no |
| `NextToken` | `string` | no |

## ListPipelineParametersForExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineParameters` | `List<Parameter>` | no |
| `NextToken` | `string` | no |

## ListPipelineVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineVersionSummaries` | `List<PipelineVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListPipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineNamePrefix` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineSummaries` | `List<PipelineSummary>` | no |
| `NextToken` | `string` | no |

## ListProcessingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingJobSummaries` | `List<ProcessingJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListProjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `MaxResults` | `integer` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectSummaryList` | `List<ProjectSummary>` | yes |
| `NextToken` | `string` | no |

## ListResourceCatalogs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NameContains` | `string` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceCatalogs` | `List<ResourceCatalog>` | no |
| `NextToken` | `string` | no |

## ListSpaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `DomainIdEquals` | `string` | no |
| `SpaceNameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Spaces` | `List<SpaceDetails>` | no |
| `NextToken` | `string` | no |

## ListStageDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `EdgeDeploymentPlanName` | `string` | yes |
| `ExcludeDevicesDeployedInOtherStage` | `boolean` | no |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceDeploymentSummaries` | `List<DeviceDeploymentSummary>` | yes |
| `NextToken` | `string` | no |

## ListStudioLifecycleConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `NameContains` | `string` | no |
| `AppTypeEquals` | `string` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `ModifiedTimeBefore` | `timestamp` | no |
| `ModifiedTimeAfter` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `StudioLifecycleConfigs` | `List<StudioLifecycleConfigDetails>` | no |

## ListSubscribedWorkteams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscribedWorkteams` | `List<SubscribedWorkteam>` | yes |
| `NextToken` | `string` | no |

## ListTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |
| `NextToken` | `string` | no |

## ListTrainingJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `WarmPoolStatusEquals` | `string` | no |
| `TrainingPlanArnEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobSummaries` | `List<TrainingJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListTrainingJobsForHyperParameterTuningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobSummaries` | `List<HyperParameterTrainingJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListTrainingPlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `StartTimeAfter` | `timestamp` | no |
| `StartTimeBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `Filters` | `List<TrainingPlanFilter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TrainingPlanSummaries` | `List<TrainingPlanSummary>` | yes |

## ListTransformJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreationTimeAfter` | `timestamp` | no |
| `CreationTimeBefore` | `timestamp` | no |
| `LastModifiedTimeAfter` | `timestamp` | no |
| `LastModifiedTimeBefore` | `timestamp` | no |
| `NameContains` | `string` | no |
| `StatusEquals` | `string` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobSummaries` | `List<TransformJobSummary>` | yes |
| `NextToken` | `string` | no |

## ListTrialComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | no |
| `TrialName` | `string` | no |
| `SourceArn` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentSummaries` | `List<TrialComponentSummary>` | no |
| `NextToken` | `string` | no |

## ListTrials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | no |
| `TrialComponentName` | `string` | no |
| `CreatedAfter` | `timestamp` | no |
| `CreatedBefore` | `timestamp` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialSummaries` | `List<TrialSummary>` | no |
| `NextToken` | `string` | no |

## ListUltraServersByReservedCapacity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReservedCapacityArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `UltraServers` | `List<UltraServer>` | yes |

## ListUserProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortOrder` | `string` | no |
| `SortBy` | `string` | no |
| `DomainIdEquals` | `string` | no |
| `UserProfileNameContains` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserProfiles` | `List<UserProfileDetails>` | no |
| `NextToken` | `string` | no |

## ListWorkforces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workforces` | `List<Workforce>` | yes |
| `NextToken` | `string` | no |

## ListWorkteams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NameContains` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workteams` | `List<Workteam>` | yes |
| `NextToken` | `string` | no |

## PutModelPackageGroupPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupName` | `string` | yes |
| `ResourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageGroupArn` | `string` | yes |

## QueryLineage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartArns` | `List<string>` | no |
| `Direction` | `string` | no |
| `IncludeEdges` | `boolean` | no |
| `Filters` | `QueryFilters` | no |
| `MaxDepth` | `integer` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Vertices` | `List<Vertex>` | no |
| `Edges` | `List<Edge>` | no |
| `NextToken` | `string` | no |

## RegisterDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `Devices` | `List<Device>` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RenderUiTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UiTemplate` | `UiTemplate` | no |
| `Task` | `RenderableTask` | yes |
| `RoleArn` | `string` | yes |
| `HumanTaskUiArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RenderedContent` | `string` | yes |
| `Errors` | `List<RenderingError>` | yes |

## RetryPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## Search

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `SearchExpression` | `SearchExpression` | no |
| `SortBy` | `string` | no |
| `SortOrder` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `CrossAccountFilterOption` | `string` | no |
| `VisibilityConditions` | `List<VisibilityConditions>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<SearchRecord>` | no |
| `NextToken` | `string` | no |
| `TotalHits` | `TotalHits` | no |

## SearchTrainingPlanOfferings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceType` | `string` | no |
| `InstanceCount` | `integer` | no |
| `UltraServerType` | `string` | no |
| `UltraServerCount` | `integer` | no |
| `StartTimeAfter` | `timestamp` | no |
| `EndTimeBefore` | `timestamp` | no |
| `DurationHours` | `long` | no |
| `TargetResources` | `List<string>` | no |
| `TrainingPlanArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingPlanOfferings` | `List<TrainingPlanOffering>` | yes |
| `TrainingPlanExtensionOfferings` | `List<TrainingPlanExtensionOffering>` | no |

## SendPipelineExecutionStepFailure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackToken` | `string` | yes |
| `FailureReason` | `string` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## SendPipelineExecutionStepSuccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CallbackToken` | `string` | yes |
| `OutputParameters` | `List<OutputParameter>` | no |
| `ClientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## StartClusterHealthCheck

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `DeepHealthCheckConfigurations` | `List<InstanceGroupHealthCheckConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

## StartEdgeDeploymentStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperimentArn` | `string` | yes |

## StartMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |

## StartMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `PipelineExecutionDisplayName` | `string` | no |
| `PipelineParameters` | `List<Parameter>` | no |
| `PipelineExecutionDescription` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |
| `SelectiveExecutionConfig` | `SelectiveExecutionConfig` | no |
| `PipelineVersionId` | `long` | no |
| `MlflowExperimentName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## StartSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SessionId` | `string` | no |
| `StreamUrl` | `string` | no |
| `TokenValue` | `string` | no |

## StopAIBenchmarkJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIBenchmarkJobArn` | `string` | yes |

## StopAIRecommendationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AIRecommendationJobArn` | `string` | yes |

## StopAutoMLJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoMLJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopCompilationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CompilationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopEdgeDeploymentStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgeDeploymentPlanName` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopEdgePackagingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EdgePackagingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopHyperParameterTuningJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HyperParameterTuningJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ModelVariantActions` | `Map<string>` | yes |
| `DesiredModelVariants` | `List<ModelVariantConfig>` | no |
| `DesiredState` | `string` | no |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperimentArn` | `string` | yes |

## StopInferenceRecommendationsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobName` | `string` | yes |
| `JobCategory` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopLabelingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LabelingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |

## StopMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OptimizationJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopPipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |
| `ClientRequestToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## StopProcessingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProcessingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTransformJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TransformJobName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionName` | `string` | yes |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `Properties` | `Map<string>` | no |
| `PropertiesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionArn` | `string` | no |

## UpdateAppImageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigName` | `string` | yes |
| `KernelGatewayImageConfig` | `KernelGatewayImageConfig` | no |
| `JupyterLabAppImageConfig` | `JupyterLabAppImageConfig` | no |
| `CodeEditorAppImageConfig` | `CodeEditorAppImageConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppImageConfigArn` | `string` | no |

## UpdateArtifact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | yes |
| `ArtifactName` | `string` | no |
| `Properties` | `Map<string>` | no |
| `PropertiesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ArtifactArn` | `string` | no |

## UpdateCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `InstanceGroups` | `List<ClusterInstanceGroupSpecification>` | no |
| `RestrictedInstanceGroups` | `List<ClusterRestrictedInstanceGroupSpecification>` | no |
| `RestrictedInstanceGroupsConfig` | `ClusterRestrictedInstanceGroupsConfig` | no |
| `TieredStorageConfig` | `ClusterTieredStorageConfig` | no |
| `NodeRecovery` | `string` | no |
| `InstanceGroupsToDelete` | `List<string>` | no |
| `NodeProvisioningMode` | `string` | no |
| `ClusterRole` | `string` | no |
| `AutoScaling` | `ClusterAutoScalingConfig` | no |
| `Orchestrator` | `ClusterOrchestrator` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

## UpdateClusterSchedulerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigId` | `string` | yes |
| `TargetVersion` | `integer` | yes |
| `SchedulerConfig` | `SchedulerConfig` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterSchedulerConfigArn` | `string` | yes |
| `ClusterSchedulerConfigVersion` | `integer` | yes |

## UpdateClusterSoftware

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterName` | `string` | yes |
| `InstanceGroups` | `List<UpdateClusterSoftwareInstanceGroupSpecification>` | no |
| `DeploymentConfig` | `DeploymentConfiguration` | no |
| `ImageId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClusterArn` | `string` | yes |

## UpdateCodeRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryName` | `string` | yes |
| `GitConfig` | `GitConfigForUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CodeRepositoryArn` | `string` | yes |

## UpdateComputeQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaId` | `string` | yes |
| `TargetVersion` | `integer` | yes |
| `ComputeQuotaConfig` | `ComputeQuotaConfig` | no |
| `ComputeQuotaTarget` | `ComputeQuotaTarget` | no |
| `ActivationState` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ComputeQuotaArn` | `string` | yes |
| `ComputeQuotaVersion` | `integer` | yes |

## UpdateContext

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextName` | `string` | yes |
| `Description` | `string` | no |
| `Properties` | `Map<string>` | no |
| `PropertiesToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContextArn` | `string` | no |

## UpdateDeviceFleet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `RoleArn` | `string` | no |
| `Description` | `string` | no |
| `OutputConfig` | `EdgeOutputConfig` | yes |
| `EnableIotRoleAlias` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDevices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeviceFleetName` | `string` | yes |
| `Devices` | `List<Device>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `DefaultUserSettings` | `UserSettings` | no |
| `DomainSettingsForUpdate` | `DomainSettingsForUpdate` | no |
| `AppSecurityGroupManagement` | `string` | no |
| `DefaultSpaceSettings` | `DefaultSpaceSettings` | no |
| `SubnetIds` | `List<string>` | no |
| `AppNetworkAccessType` | `string` | no |
| `TagPropagation` | `string` | no |
| `HomeEfsFileSystemCreation` | `string` | no |
| `VpcId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainArn` | `string` | no |

## UpdateEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `EndpointConfigName` | `string` | yes |
| `RetainAllVariantProperties` | `boolean` | no |
| `ExcludeRetainedVariantProperties` | `List<VariantProperty>` | no |
| `DeploymentConfig` | `DeploymentConfig` | no |
| `RetainDeploymentConfig` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

## UpdateEndpointWeightsAndCapacities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `DesiredWeightsAndCapacities` | `List<DesiredWeightAndCapacity>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

## UpdateExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExperimentArn` | `string` | no |

## UpdateFeatureGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `FeatureAdditions` | `List<FeatureDefinition>` | no |
| `OnlineStoreConfig` | `OnlineStoreConfigUpdate` | no |
| `ThroughputConfig` | `ThroughputConfigUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupArn` | `string` | yes |

## UpdateFeatureMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FeatureGroupName` | `string` | yes |
| `FeatureName` | `string` | yes |
| `Description` | `string` | no |
| `ParameterAdditions` | `List<FeatureParameter>` | no |
| `ParameterRemovals` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubDescription` | `string` | no |
| `HubDisplayName` | `string` | no |
| `HubSearchKeywords` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |

## UpdateHubContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `HubContentVersion` | `string` | yes |
| `HubContentDisplayName` | `string` | no |
| `HubContentDescription` | `string` | no |
| `HubContentMarkdown` | `string` | no |
| `HubContentSearchKeywords` | `List<string>` | no |
| `SupportStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |
| `HubContentArn` | `string` | yes |

## UpdateHubContentReference

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubName` | `string` | yes |
| `HubContentName` | `string` | yes |
| `HubContentType` | `string` | yes |
| `MinVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HubArn` | `string` | yes |
| `HubContentArn` | `string` | yes |

## UpdateImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeleteProperties` | `List<string>` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `ImageName` | `string` | yes |
| `RoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageArn` | `string` | no |

## UpdateImageVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageName` | `string` | yes |
| `Alias` | `string` | no |
| `Version` | `integer` | no |
| `AliasesToAdd` | `List<string>` | no |
| `AliasesToDelete` | `List<string>` | no |
| `VendorGuidance` | `string` | no |
| `JobType` | `string` | no |
| `MLFramework` | `string` | no |
| `ProgrammingLang` | `string` | no |
| `Processor` | `string` | no |
| `Horovod` | `boolean` | no |
| `ReleaseNotes` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImageVersionArn` | `string` | no |

## UpdateInferenceComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |
| `Specification` | `InferenceComponentSpecification` | no |
| `Specifications` | `List<InferenceComponentSpecification>` | no |
| `RuntimeConfig` | `InferenceComponentRuntimeConfig` | no |
| `DeploymentConfig` | `InferenceComponentDeploymentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentArn` | `string` | yes |

## UpdateInferenceComponentRuntimeConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentName` | `string` | yes |
| `DesiredRuntimeConfig` | `InferenceComponentRuntimeConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceComponentArn` | `string` | yes |

## UpdateInferenceExperiment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Schedule` | `InferenceExperimentSchedule` | no |
| `Description` | `string` | no |
| `ModelVariants` | `List<ModelVariantConfig>` | no |
| `DataStorageConfig` | `InferenceExperimentDataStorageConfig` | no |
| `ShadowModeConfig` | `ShadowModeConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InferenceExperimentArn` | `string` | yes |

## UpdateMlflowApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `Name` | `string` | no |
| `ArtifactStoreUri` | `string` | no |
| `ModelRegistrationMode` | `string` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `DefaultDomainIdList` | `List<string>` | no |
| `AccountDefaultStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## UpdateMlflowTrackingServer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerName` | `string` | yes |
| `ArtifactStoreUri` | `string` | no |
| `TrackingServerSize` | `string` | no |
| `AutomaticModelRegistration` | `boolean` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `S3BucketOwnerAccountId` | `string` | no |
| `S3BucketOwnerVerification` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrackingServerArn` | `string` | no |

## UpdateModelCard

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardName` | `string` | yes |
| `Content` | `string` | no |
| `ModelCardStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelCardArn` | `string` | yes |

## UpdateModelPackage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageArn` | `string` | yes |
| `ModelApprovalStatus` | `string` | no |
| `ModelPackageRegistrationType` | `string` | no |
| `ApprovalDescription` | `string` | no |
| `CustomerMetadataProperties` | `Map<string>` | no |
| `CustomerMetadataPropertiesToRemove` | `List<string>` | no |
| `AdditionalInferenceSpecificationsToAdd` | `List<AdditionalInferenceSpecificationDefinition>` | no |
| `InferenceSpecification` | `InferenceSpecification` | no |
| `SourceUri` | `string` | no |
| `ModelCard` | `ModelPackageModelCard` | no |
| `ModelLifeCycle` | `ModelLifeCycle` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelPackageArn` | `string` | yes |

## UpdateMonitoringAlert

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |
| `MonitoringAlertName` | `string` | yes |
| `DatapointsToAlert` | `integer` | yes |
| `EvaluationPeriod` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleArn` | `string` | yes |
| `MonitoringAlertName` | `string` | no |

## UpdateMonitoringSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleName` | `string` | yes |
| `MonitoringScheduleConfig` | `MonitoringScheduleConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MonitoringScheduleArn` | `string` | yes |

## UpdateNotebookInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceName` | `string` | yes |
| `InstanceType` | `string` | no |
| `IpAddressType` | `string` | no |
| `PlatformIdentifier` | `string` | no |
| `RoleArn` | `string` | no |
| `LifecycleConfigName` | `string` | no |
| `DisassociateLifecycleConfig` | `boolean` | no |
| `VolumeSizeInGB` | `integer` | no |
| `DefaultCodeRepository` | `string` | no |
| `AdditionalCodeRepositories` | `List<string>` | no |
| `AcceleratorTypes` | `List<string>` | no |
| `DisassociateAcceleratorTypes` | `boolean` | no |
| `DisassociateDefaultCodeRepository` | `boolean` | no |
| `DisassociateAdditionalCodeRepositories` | `boolean` | no |
| `RootAccess` | `string` | no |
| `InstanceMetadataServiceConfiguration` | `InstanceMetadataServiceConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotebookInstanceLifecycleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotebookInstanceLifecycleConfigName` | `string` | yes |
| `OnCreate` | `List<NotebookInstanceLifecycleHook>` | no |
| `OnStart` | `List<NotebookInstanceLifecycleHook>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePartnerApp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | yes |
| `MaintenanceConfig` | `PartnerAppMaintenanceConfig` | no |
| `Tier` | `string` | no |
| `ApplicationConfig` | `PartnerAppConfig` | no |
| `IdcConfig` | `IdcConfigInput` | no |
| `AuthType` | `string` | no |
| `EnableIamSessionBasedIdentity` | `boolean` | no |
| `EnableAutoMinorVersionUpgrade` | `boolean` | no |
| `AppVersion` | `string` | no |
| `ClientToken` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## UpdatePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineName` | `string` | yes |
| `PipelineDisplayName` | `string` | no |
| `PipelineDefinition` | `string` | no |
| `PipelineDefinitionS3Location` | `PipelineDefinitionS3Location` | no |
| `PipelineDescription` | `string` | no |
| `RoleArn` | `string` | no |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |
| `PipelineVersionId` | `long` | no |

## UpdatePipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | yes |
| `PipelineExecutionDescription` | `string` | no |
| `PipelineExecutionDisplayName` | `string` | no |
| `ParallelismConfiguration` | `ParallelismConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineExecutionArn` | `string` | no |

## UpdatePipelineVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | yes |
| `PipelineVersionId` | `long` | yes |
| `PipelineVersionDisplayName` | `string` | no |
| `PipelineVersionDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PipelineArn` | `string` | no |
| `PipelineVersionId` | `long` | no |

## UpdateProject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectName` | `string` | yes |
| `ProjectDescription` | `string` | no |
| `ServiceCatalogProvisioningUpdateDetails` | `ServiceCatalogProvisioningUpdateDetails` | no |
| `Tags` | `List<Tag>` | no |
| `TemplateProvidersToUpdate` | `List<UpdateTemplateProvider>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProjectArn` | `string` | yes |

## UpdateSpace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `SpaceName` | `string` | yes |
| `SpaceSettings` | `SpaceSettings` | no |
| `SpaceDisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SpaceArn` | `string` | no |

## UpdateTrainingJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobName` | `string` | yes |
| `ProfilerConfig` | `ProfilerConfigForUpdate` | no |
| `ProfilerRuleConfigurations` | `List<ProfilerRuleConfiguration>` | no |
| `ResourceConfig` | `ResourceConfigForUpdate` | no |
| `RemoteDebugConfig` | `RemoteDebugConfigForUpdate` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrainingJobArn` | `string` | yes |

## UpdateTrial

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialName` | `string` | yes |
| `DisplayName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialArn` | `string` | no |

## UpdateTrialComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Status` | `TrialComponentStatus` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Parameters` | `Map<TrialComponentParameterValue>` | no |
| `ParametersToRemove` | `List<string>` | no |
| `InputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `InputArtifactsToRemove` | `List<string>` | no |
| `OutputArtifacts` | `Map<TrialComponentArtifact>` | no |
| `OutputArtifactsToRemove` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrialComponentArn` | `string` | no |

## UpdateUserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainId` | `string` | yes |
| `UserProfileName` | `string` | yes |
| `UserSettings` | `UserSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserProfileArn` | `string` | no |

## UpdateWorkforce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkforceName` | `string` | yes |
| `SourceIpConfig` | `SourceIpConfig` | no |
| `OidcConfig` | `OidcConfig` | no |
| `WorkforceVpcConfig` | `WorkforceVpcConfigRequest` | no |
| `IpAddressType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workforce` | `Workforce` | yes |

## UpdateWorkteam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkteamName` | `string` | yes |
| `MemberDefinitions` | `List<MemberDefinition>` | no |
| `Description` | `string` | no |
| `NotificationConfiguration` | `NotificationConfiguration` | no |
| `WorkerAccessConfiguration` | `WorkerAccessConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workteam` | `Workteam` | yes |

