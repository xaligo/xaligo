# EC2 Image Builder

API version: 2019-12-02. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/imagebuilder/2019-12-02/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelImageCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageBuildVersionArn` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## CancelLifecycleExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | no |

## CreateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `semanticVersion` | `string` | yes |
| `description` | `string` | no |
| `changeDescription` | `string` | no |
| `platform` | `string` | yes |
| `supportedOsVersions` | `List<string>` | no |
| `data` | `string` | no |
| `uri` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `componentBuildVersionArn` | `string` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## CreateContainerRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerType` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `semanticVersion` | `string` | yes |
| `components` | `List<ComponentConfiguration>` | no |
| `instanceConfiguration` | `InstanceConfiguration` | no |
| `dockerfileTemplateData` | `string` | no |
| `dockerfileTemplateUri` | `string` | no |
| `platformOverride` | `string` | no |
| `imageOsVersionOverride` | `string` | no |
| `parentImage` | `string` | yes |
| `tags` | `Map<string>` | no |
| `workingDirectory` | `string` | no |
| `targetRepository` | `TargetContainerRepository` | yes |
| `kmsKeyId` | `string` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `containerRecipeArn` | `string` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## CreateDistributionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `distributions` | `List<Distribution>` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `distributionConfigurationArn` | `string` | no |

## CreateImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageRecipeArn` | `string` | no |
| `containerRecipeArn` | `string` | no |
| `distributionConfigurationArn` | `string` | no |
| `infrastructureConfigurationArn` | `string` | yes |
| `imageTestsConfiguration` | `ImageTestsConfiguration` | no |
| `enhancedImageMetadataEnabled` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | no |
| `workflows` | `List<WorkflowConfiguration>` | no |
| `executionRole` | `string` | no |
| `loggingConfiguration` | `ImageLoggingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## CreateImagePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `imageRecipeArn` | `string` | no |
| `containerRecipeArn` | `string` | no |
| `infrastructureConfigurationArn` | `string` | yes |
| `distributionConfigurationArn` | `string` | no |
| `imageTestsConfiguration` | `ImageTestsConfiguration` | no |
| `enhancedImageMetadataEnabled` | `boolean` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `tags` | `Map<string>` | no |
| `imageTags` | `Map<string>` | no |
| `clientToken` | `string` | yes |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | no |
| `workflows` | `List<WorkflowConfiguration>` | no |
| `executionRole` | `string` | no |
| `loggingConfiguration` | `PipelineLoggingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imagePipelineArn` | `string` | no |

## CreateImageRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `semanticVersion` | `string` | yes |
| `components` | `List<ComponentConfiguration>` | no |
| `parentImage` | `string` | yes |
| `blockDeviceMappings` | `List<InstanceBlockDeviceMapping>` | no |
| `tags` | `Map<string>` | no |
| `workingDirectory` | `string` | no |
| `additionalInstanceConfiguration` | `AdditionalInstanceConfiguration` | no |
| `amiTags` | `Map<string>` | no |
| `amiWatermarks` | `List<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imageRecipeArn` | `string` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## CreateInfrastructureConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `instanceTypes` | `List<string>` | no |
| `instanceProfileName` | `string` | yes |
| `securityGroupIds` | `List<string>` | no |
| `subnetId` | `string` | no |
| `logging` | `Logging` | no |
| `keyPair` | `string` | no |
| `terminateInstanceOnFailure` | `boolean` | no |
| `snsTopicArn` | `string` | no |
| `resourceTags` | `Map<string>` | no |
| `instanceMetadataOptions` | `InstanceMetadataOptions` | no |
| `tags` | `Map<string>` | no |
| `placement` | `Placement` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `infrastructureConfigurationArn` | `string` | no |

## CreateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `executionRole` | `string` | yes |
| `resourceType` | `string` | yes |
| `policyDetails` | `List<LifecyclePolicyDetail>` | yes |
| `resourceSelection` | `LifecyclePolicyResourceSelection` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `lifecyclePolicyArn` | `string` | no |

## CreateWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `semanticVersion` | `string` | yes |
| `description` | `string` | no |
| `changeDescription` | `string` | no |
| `data` | `string` | no |
| `uri` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |
| `type` | `string` | yes |
| `dryRun` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `workflowBuildVersionArn` | `string` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## DeleteComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `componentBuildVersionArn` | `string` | no |

## DeleteContainerRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `containerRecipeArn` | `string` | no |

## DeleteDistributionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `distributionConfigurationArn` | `string` | no |

## DeleteImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## DeleteImagePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imagePipelineArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imagePipelineArn` | `string` | no |

## DeleteImageRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageRecipeArn` | `string` | no |

## DeleteInfrastructureConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `infrastructureConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `infrastructureConfigurationArn` | `string` | no |

## DeleteLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyArn` | `string` | no |

## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowBuildVersionArn` | `string` | no |

## DistributeImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceImage` | `string` | yes |
| `distributionConfigurationArn` | `string` | yes |
| `executionRole` | `string` | yes |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |
| `loggingConfiguration` | `ImageLoggingConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## GetComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `component` | `Component` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## GetComponentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `policy` | `string` | no |

## GetContainerRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `containerRecipe` | `ContainerRecipe` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## GetContainerRecipePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `policy` | `string` | no |

## GetDistributionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `distributionConfiguration` | `DistributionConfiguration` | no |

## GetImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `image` | `Image` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## GetImagePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imagePipelineArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imagePipeline` | `ImagePipeline` | no |

## GetImagePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `policy` | `string` | no |

## GetImageRecipe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageRecipe` | `ImageRecipe` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## GetImageRecipePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageRecipeArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `policy` | `string` | no |

## GetInfrastructureConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `infrastructureConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `infrastructureConfiguration` | `InfrastructureConfiguration` | no |

## GetLifecycleExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecution` | `LifecycleExecution` | no |

## GetLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicy` | `LifecyclePolicy` | no |

## GetMarketplaceResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceType` | `string` | yes |
| `resourceArn` | `string` | yes |
| `resourceLocation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |
| `url` | `string` | no |
| `data` | `string` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflow` | `Workflow` | no |
| `latestVersionReferences` | `LatestVersionReferences` | no |

## GetWorkflowExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `workflowBuildVersionArn` | `string` | no |
| `workflowExecutionId` | `string` | no |
| `imageBuildVersionArn` | `string` | no |
| `type` | `string` | no |
| `status` | `string` | no |
| `message` | `string` | no |
| `totalStepCount` | `integer` | no |
| `totalStepsSucceeded` | `integer` | no |
| `totalStepsFailed` | `integer` | no |
| `totalStepsSkipped` | `integer` | no |
| `startTime` | `string` | no |
| `endTime` | `string` | no |
| `parallelGroup` | `string` | no |

## GetWorkflowStepExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stepExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `stepExecutionId` | `string` | no |
| `workflowBuildVersionArn` | `string` | no |
| `workflowExecutionId` | `string` | no |
| `imageBuildVersionArn` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `action` | `string` | no |
| `status` | `string` | no |
| `rollbackStatus` | `string` | no |
| `message` | `string` | no |
| `inputs` | `string` | no |
| `outputs` | `string` | no |
| `startTime` | `string` | no |
| `endTime` | `string` | no |
| `onFailure` | `string` | no |
| `timeoutSeconds` | `integer` | no |

## ImportComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `semanticVersion` | `string` | yes |
| `description` | `string` | no |
| `changeDescription` | `string` | no |
| `type` | `string` | yes |
| `format` | `string` | yes |
| `platform` | `string` | yes |
| `data` | `string` | no |
| `uri` | `string` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `componentBuildVersionArn` | `string` | no |

## ImportDiskImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `semanticVersion` | `string` | yes |
| `description` | `string` | no |
| `platform` | `string` | yes |
| `osVersion` | `string` | yes |
| `executionRole` | `string` | no |
| `infrastructureConfigurationArn` | `string` | yes |
| `uri` | `string` | yes |
| `loggingConfiguration` | `ImageLoggingConfiguration` | no |
| `tags` | `Map<string>` | no |
| `registerImageOptions` | `RegisterImageOptions` | no |
| `windowsConfiguration` | `WindowsConfiguration` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## ImportVmImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `semanticVersion` | `string` | yes |
| `description` | `string` | no |
| `platform` | `string` | yes |
| `osVersion` | `string` | no |
| `vmImportTaskId` | `string` | yes |
| `loggingConfiguration` | `ImageLoggingConfiguration` | no |
| `tags` | `Map<string>` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageArn` | `string` | no |
| `clientToken` | `string` | no |

## ListComponentBuildVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentVersionArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `componentSummaryList` | `List<ComponentSummary>` | no |
| `nextToken` | `string` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owner` | `string` | no |
| `filters` | `List<Filter>` | no |
| `byName` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `componentVersionList` | `List<ComponentVersion>` | no |
| `nextToken` | `string` | no |

## ListContainerRecipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owner` | `string` | no |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `containerRecipeSummaryList` | `List<ContainerRecipeSummary>` | no |
| `nextToken` | `string` | no |

## ListDistributionConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `distributionConfigurationSummaryList` | `List<DistributionConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListImageBuildVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageVersionArn` | `string` | no |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageSummaryList` | `List<ImageSummary>` | no |
| `nextToken` | `string` | no |

## ListImagePackages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageBuildVersionArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imagePackageList` | `List<ImagePackage>` | no |
| `nextToken` | `string` | no |

## ListImagePipelineImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imagePipelineArn` | `string` | yes |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageSummaryList` | `List<ImageSummary>` | no |
| `nextToken` | `string` | no |

## ListImagePipelines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imagePipelineList` | `List<ImagePipeline>` | no |
| `nextToken` | `string` | no |

## ListImageRecipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owner` | `string` | no |
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageRecipeSummaryList` | `List<ImageRecipeSummary>` | no |
| `nextToken` | `string` | no |

## ListImageScanFindingAggregations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filter` | `Filter` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `aggregationType` | `string` | no |
| `responses` | `List<ImageScanFindingAggregation>` | no |
| `nextToken` | `string` | no |

## ListImageScanFindings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<ImageScanFindingsFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `findings` | `List<ImageScanFinding>` | no |
| `nextToken` | `string` | no |

## ListImages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owner` | `string` | no |
| `filters` | `List<Filter>` | no |
| `byName` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `includeDeprecated` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageVersionList` | `List<ImageVersion>` | no |
| `nextToken` | `string` | no |

## ListInfrastructureConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `infrastructureConfigurationSummaryList` | `List<InfrastructureConfigurationSummary>` | no |
| `nextToken` | `string` | no |

## ListLifecycleExecutionResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | yes |
| `parentResourceId` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | no |
| `lifecycleExecutionState` | `LifecycleExecutionState` | no |
| `resources` | `List<LifecycleExecutionResource>` | no |
| `nextToken` | `string` | no |

## ListLifecycleExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutions` | `List<LifecycleExecution>` | no |
| `nextToken` | `string` | no |

## ListLifecyclePolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<Filter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicySummaryList` | `List<LifecyclePolicySummary>` | no |
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

## ListWaitingWorkflowSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `steps` | `List<WorkflowStepExecution>` | no |
| `nextToken` | `string` | no |

## ListWorkflowBuildVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowVersionArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowSummaryList` | `List<WorkflowSummary>` | no |
| `nextToken` | `string` | no |

## ListWorkflowExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `imageBuildVersionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `workflowExecutions` | `List<WorkflowExecutionMetadata>` | no |
| `imageBuildVersionArn` | `string` | no |
| `message` | `string` | no |
| `nextToken` | `string` | no |

## ListWorkflowStepExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `workflowExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `steps` | `List<WorkflowStepMetadata>` | no |
| `workflowBuildVersionArn` | `string` | no |
| `workflowExecutionId` | `string` | no |
| `imageBuildVersionArn` | `string` | no |
| `message` | `string` | no |
| `nextToken` | `string` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `owner` | `string` | no |
| `filters` | `List<Filter>` | no |
| `byName` | `boolean` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workflowVersionList` | `List<WorkflowVersion>` | no |
| `nextToken` | `string` | no |

## PutComponentPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `componentArn` | `string` | no |

## PutContainerRecipePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerRecipeArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `containerRecipeArn` | `string` | no |

## PutImagePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageArn` | `string` | no |

## PutImageRecipePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageRecipeArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `imageRecipeArn` | `string` | no |

## RetryImage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imageBuildVersionArn` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## SendWorkflowStepAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stepExecutionId` | `string` | yes |
| `imageBuildVersionArn` | `string` | yes |
| `action` | `string` | yes |
| `reason` | `string` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `stepExecutionId` | `string` | no |
| `imageBuildVersionArn` | `string` | no |
| `clientToken` | `string` | no |

## StartImagePipelineExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imagePipelineArn` | `string` | yes |
| `clientToken` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imageBuildVersionArn` | `string` | no |

## StartResourceStateUpdate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `state` | `ResourceState` | yes |
| `executionRole` | `string` | no |
| `includeResources` | `ResourceStateUpdateIncludeResources` | no |
| `exclusionRules` | `ResourceStateUpdateExclusionRules` | no |
| `updateAt` | `timestamp` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecycleExecutionId` | `string` | no |
| `resourceArn` | `string` | no |

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


## UpdateDistributionConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `distributionConfigurationArn` | `string` | yes |
| `description` | `string` | no |
| `distributions` | `List<Distribution>` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `distributionConfigurationArn` | `string` | no |

## UpdateImagePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `imagePipelineArn` | `string` | yes |
| `description` | `string` | no |
| `imageRecipeArn` | `string` | no |
| `containerRecipeArn` | `string` | no |
| `infrastructureConfigurationArn` | `string` | yes |
| `distributionConfigurationArn` | `string` | no |
| `imageTestsConfiguration` | `ImageTestsConfiguration` | no |
| `enhancedImageMetadataEnabled` | `boolean` | no |
| `schedule` | `Schedule` | no |
| `status` | `string` | no |
| `clientToken` | `string` | yes |
| `imageScanningConfiguration` | `ImageScanningConfiguration` | no |
| `workflows` | `List<WorkflowConfiguration>` | no |
| `loggingConfiguration` | `PipelineLoggingConfiguration` | no |
| `executionRole` | `string` | no |
| `imageTags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `imagePipelineArn` | `string` | no |

## UpdateInfrastructureConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `infrastructureConfigurationArn` | `string` | yes |
| `description` | `string` | no |
| `instanceTypes` | `List<string>` | no |
| `instanceProfileName` | `string` | yes |
| `securityGroupIds` | `List<string>` | no |
| `subnetId` | `string` | no |
| `logging` | `Logging` | no |
| `keyPair` | `string` | no |
| `terminateInstanceOnFailure` | `boolean` | no |
| `snsTopicArn` | `string` | no |
| `resourceTags` | `Map<string>` | no |
| `instanceMetadataOptions` | `InstanceMetadataOptions` | no |
| `placement` | `Placement` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `requestId` | `string` | no |
| `clientToken` | `string` | no |
| `infrastructureConfigurationArn` | `string` | no |

## UpdateLifecyclePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyArn` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | no |
| `executionRole` | `string` | yes |
| `resourceType` | `string` | yes |
| `policyDetails` | `List<LifecyclePolicyDetail>` | yes |
| `resourceSelection` | `LifecyclePolicyResourceSelection` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `lifecyclePolicyArn` | `string` | no |

