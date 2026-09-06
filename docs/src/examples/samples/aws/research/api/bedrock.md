# Amazon Bedrock

API version: 2023-04-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock/2023-04-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## BatchDeleteAdvancedPromptOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteAdvancedPromptOptimizationJobError>` | yes |
| `advancedPromptOptimizationJobs` | `List<BatchDeleteAdvancedPromptOptimizationJobItem>` | yes |

## BatchDeleteEvaluationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifiers` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `errors` | `List<BatchDeleteEvaluationJobError>` | yes |
| `evaluationJobs` | `List<BatchDeleteEvaluationJobItem>` | yes |

## CancelAutomatedReasoningPolicyBuildWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAdvancedPromptOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `jobDescription` | `string` | no |
| `clientToken` | `string` | no |
| `inputConfig` | `AdvancedPromptOptimizationInputConfig` | yes |
| `outputConfig` | `AdvancedPromptOptimizationOutputConfig` | yes |
| `encryptionKeyArn` | `string` | no |
| `tags` | `List<Tag>` | no |
| `modelConfigurations` | `List<ModelConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateAutomatedReasoningPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientRequestToken` | `string` | no |
| `policyDefinition` | `AutomatedReasoningPolicyDefinition` | no |
| `kmsKeyId` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `version` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `definitionHash` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateAutomatedReasoningPolicyTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `guardContent` | `string` | yes |
| `queryContent` | `string` | no |
| `expectedAggregatedFindingsResult` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `confidenceThreshold` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCaseId` | `string` | yes |

## CreateAutomatedReasoningPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `lastUpdatedDefinitionHash` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `version` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `definitionHash` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateCustomModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelName` | `string` | yes |
| `modelSourceConfig` | `ModelDataSource` | no |
| `customModelDataSource` | `CustomModelDataSource` | no |
| `modelKmsKeyArn` | `string` | no |
| `roleArn` | `string` | no |
| `modelTags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelArn` | `string` | yes |

## CreateCustomModelDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelDeploymentName` | `string` | yes |
| `modelArn` | `string` | yes |
| `description` | `string` | no |
| `tags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customModelDeploymentArn` | `string` | yes |

## CreateEvaluationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `jobDescription` | `string` | no |
| `clientRequestToken` | `string` | no |
| `roleArn` | `string` | yes |
| `customerEncryptionKeyId` | `string` | no |
| `jobTags` | `List<Tag>` | no |
| `applicationType` | `string` | no |
| `evaluationConfig` | `EvaluationConfig` | yes |
| `inferenceConfig` | `EvaluationInferenceConfig` | yes |
| `outputDataConfig` | `EvaluationOutputDataConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateFoundationModelAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `offerToken` | `string` | yes |
| `modelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |

## CreateGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `topicPolicyConfig` | `GuardrailTopicPolicyConfig` | no |
| `contentPolicyConfig` | `GuardrailContentPolicyConfig` | no |
| `wordPolicyConfig` | `GuardrailWordPolicyConfig` | no |
| `sensitiveInformationPolicyConfig` | `GuardrailSensitiveInformationPolicyConfig` | no |
| `contextualGroundingPolicyConfig` | `GuardrailContextualGroundingPolicyConfig` | no |
| `automatedReasoningPolicyConfig` | `GuardrailAutomatedReasoningPolicyConfig` | no |
| `crossRegionConfig` | `GuardrailCrossRegionConfig` | no |
| `blockedInputMessaging` | `string` | yes |
| `blockedOutputsMessaging` | `string` | yes |
| `kmsKeyId` | `string` | no |
| `tags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailId` | `string` | yes |
| `guardrailArn` | `string` | yes |
| `version` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateGuardrailVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | yes |
| `description` | `string` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailId` | `string` | yes |
| `version` | `string` | yes |

## CreateInferenceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileName` | `string` | yes |
| `description` | `string` | no |
| `clientRequestToken` | `string` | no |
| `modelSource` | `InferenceProfileModelSource` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileArn` | `string` | yes |
| `status` | `string` | no |

## CreateMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelSourceIdentifier` | `string` | yes |
| `endpointConfig` | `EndpointConfig` | yes |
| `acceptEula` | `boolean` | no |
| `endpointName` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceModelEndpoint` | `MarketplaceModelEndpoint` | yes |

## CreateModelCopyJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceModelArn` | `string` | yes |
| `targetModelName` | `string` | yes |
| `modelKmsKeyId` | `string` | no |
| `targetModelTags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateModelCustomizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `customModelName` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `baseModelIdentifier` | `string` | yes |
| `customizationType` | `string` | no |
| `customModelKmsKeyId` | `string` | no |
| `jobTags` | `List<Tag>` | no |
| `customModelTags` | `List<Tag>` | no |
| `trainingDataConfig` | `TrainingDataConfig` | yes |
| `validationDataConfig` | `ValidationDataConfig` | no |
| `outputDataConfig` | `OutputDataConfig` | yes |
| `hyperParameters` | `Map<string>` | no |
| `vpcConfig` | `VpcConfig` | no |
| `customizationConfig` | `CustomizationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateModelImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `importedModelName` | `string` | yes |
| `roleArn` | `string` | yes |
| `modelDataSource` | `ModelDataSource` | yes |
| `jobTags` | `List<Tag>` | no |
| `importedModelTags` | `List<Tag>` | no |
| `clientRequestToken` | `string` | no |
| `vpcConfig` | `VpcConfig` | no |
| `importedModelKmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreateModelInvocationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `roleArn` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `modelId` | `string` | yes |
| `inputDataConfig` | `ModelInvocationJobInputDataConfig` | yes |
| `outputDataConfig` | `ModelInvocationJobOutputDataConfig` | yes |
| `vpcConfig` | `VpcConfig` | no |
| `timeoutDurationInHours` | `integer` | no |
| `tags` | `List<Tag>` | no |
| `modelInvocationType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

## CreatePromptRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientRequestToken` | `string` | no |
| `promptRouterName` | `string` | yes |
| `models` | `List<PromptRouterTargetModel>` | yes |
| `description` | `string` | no |
| `routingCriteria` | `RoutingCriteria` | yes |
| `fallbackModel` | `PromptRouterTargetModel` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptRouterArn` | `string` | no |

## CreateProvisionedModelThroughput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientRequestToken` | `string` | no |
| `modelUnits` | `integer` | yes |
| `provisionedModelName` | `string` | yes |
| `modelId` | `string` | yes |
| `commitmentDuration` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provisionedModelArn` | `string` | yes |

## DeleteAutomatedReasoningPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `force` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAutomatedReasoningPolicyBuildWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAutomatedReasoningPolicyTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCaseId` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomModelDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customModelDeploymentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEnforcedGuardrailConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFoundationModelAgreement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | yes |
| `guardrailVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteImportedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInferenceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModelInvocationLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePromptRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptRouterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProvisionedModelThroughput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provisionedModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportAutomatedReasoningPolicyVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyDefinition` | `AutomatedReasoningPolicyDefinition` | yes |

## GetAccountDataRetention

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | yes |
| `updatedAt` | `timestamp` | no |

## GetAdvancedPromptOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `jobName` | `string` | yes |
| `jobDescription` | `string` | no |
| `jobStatus` | `string` | yes |
| `inputConfig` | `AdvancedPromptOptimizationInputConfig` | yes |
| `outputConfig` | `AdvancedPromptOptimizationOutputConfig` | yes |
| `encryptionKeyArn` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | no |
| `failureMessage` | `string` | no |
| `modelConfigurations` | `List<ModelConfiguration>` | yes |

## GetAutomatedReasoningPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `name` | `string` | yes |
| `version` | `string` | yes |
| `policyId` | `string` | yes |
| `description` | `string` | no |
| `definitionHash` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | yes |

## GetAutomatedReasoningPolicyAnnotations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `name` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `annotations` | `List<AutomatedReasoningPolicyAnnotation>` | yes |
| `annotationSetHash` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## GetAutomatedReasoningPolicyBuildWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `status` | `string` | yes |
| `buildWorkflowType` | `string` | yes |
| `documentName` | `string` | no |
| `documentContentType` | `string` | no |
| `documentDescription` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetAutomatedReasoningPolicyBuildWorkflowResultAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `assetType` | `string` | yes |
| `assetId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `buildWorkflowAssets` | `AutomatedReasoningPolicyBuildResultAssets` | no |

## GetAutomatedReasoningPolicyNextScenario

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `scenario` | `AutomatedReasoningPolicyScenario` | no |

## GetAutomatedReasoningPolicyTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCase` | `AutomatedReasoningPolicyTestCase` | yes |

## GetAutomatedReasoningPolicyTestResult

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `testCaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testResult` | `AutomatedReasoningPolicyTestResult` | yes |

## GetCustomModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelArn` | `string` | yes |
| `modelName` | `string` | yes |
| `jobName` | `string` | no |
| `jobArn` | `string` | no |
| `baseModelArn` | `string` | no |
| `customizationType` | `string` | no |
| `modelKmsKeyArn` | `string` | no |
| `hyperParameters` | `Map<string>` | no |
| `trainingDataConfig` | `TrainingDataConfig` | no |
| `validationDataConfig` | `ValidationDataConfig` | no |
| `outputDataConfig` | `OutputDataConfig` | no |
| `trainingMetrics` | `TrainingMetrics` | no |
| `validationMetrics` | `List<ValidatorMetric>` | no |
| `creationTime` | `timestamp` | yes |
| `customizationConfig` | `CustomizationConfig` | no |
| `modelStatus` | `string` | no |
| `failureMessage` | `string` | no |

## GetCustomModelDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customModelDeploymentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customModelDeploymentArn` | `string` | yes |
| `modelDeploymentName` | `string` | yes |
| `modelArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `updateDetails` | `CustomModelDeploymentUpdateDetails` | no |
| `failureMessage` | `string` | no |
| `lastUpdatedAt` | `timestamp` | no |

## GetEvaluationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobName` | `string` | yes |
| `status` | `string` | yes |
| `jobArn` | `string` | yes |
| `jobDescription` | `string` | no |
| `roleArn` | `string` | yes |
| `customerEncryptionKeyId` | `string` | no |
| `jobType` | `string` | yes |
| `applicationType` | `string` | no |
| `evaluationConfig` | `EvaluationConfig` | yes |
| `inferenceConfig` | `EvaluationInferenceConfig` | yes |
| `outputDataConfig` | `EvaluationOutputDataConfig` | yes |
| `creationTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | no |
| `failureMessages` | `List<string>` | no |

## GetFoundationModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelDetails` | `FoundationModelDetails` | no |

## GetFoundationModelAvailability

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `agreementAvailability` | `AgreementAvailability` | yes |
| `authorizationStatus` | `string` | yes |
| `entitlementAvailability` | `string` | yes |
| `regionAvailability` | `string` | yes |

## GetGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | yes |
| `guardrailVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `guardrailId` | `string` | yes |
| `guardrailArn` | `string` | yes |
| `version` | `string` | yes |
| `status` | `string` | yes |
| `topicPolicy` | `GuardrailTopicPolicy` | no |
| `contentPolicy` | `GuardrailContentPolicy` | no |
| `wordPolicy` | `GuardrailWordPolicy` | no |
| `sensitiveInformationPolicy` | `GuardrailSensitiveInformationPolicy` | no |
| `contextualGroundingPolicy` | `GuardrailContextualGroundingPolicy` | no |
| `automatedReasoningPolicy` | `GuardrailAutomatedReasoningPolicy` | no |
| `crossRegionDetails` | `GuardrailCrossRegionDetails` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `statusReasons` | `List<string>` | no |
| `failureRecommendations` | `List<string>` | no |
| `blockedInputMessaging` | `string` | yes |
| `blockedOutputsMessaging` | `string` | yes |
| `kmsKeyArn` | `string` | no |

## GetImportedModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelArn` | `string` | no |
| `modelName` | `string` | no |
| `jobName` | `string` | no |
| `jobArn` | `string` | no |
| `modelDataSource` | `ModelDataSource` | no |
| `creationTime` | `timestamp` | no |
| `modelArchitecture` | `string` | no |
| `modelKmsKeyArn` | `string` | no |
| `instructSupported` | `boolean` | no |
| `customModelUnits` | `CustomModelUnits` | no |

## GetInferenceProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileName` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `inferenceProfileArn` | `string` | yes |
| `models` | `List<InferenceProfileModel>` | yes |
| `inferenceProfileId` | `string` | yes |
| `status` | `string` | yes |
| `type` | `string` | yes |

## GetMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceModelEndpoint` | `MarketplaceModelEndpoint` | no |

## GetModelCopyJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `targetModelArn` | `string` | yes |
| `targetModelName` | `string` | no |
| `sourceAccountId` | `string` | yes |
| `sourceModelArn` | `string` | yes |
| `targetModelKmsKeyArn` | `string` | no |
| `targetModelTags` | `List<Tag>` | no |
| `failureMessage` | `string` | no |
| `sourceModelName` | `string` | no |

## GetModelCustomizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `jobName` | `string` | yes |
| `outputModelName` | `string` | yes |
| `outputModelArn` | `string` | no |
| `clientRequestToken` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | no |
| `statusDetails` | `StatusDetails` | no |
| `failureMessage` | `string` | no |
| `creationTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `baseModelArn` | `string` | yes |
| `hyperParameters` | `Map<string>` | no |
| `trainingDataConfig` | `TrainingDataConfig` | yes |
| `validationDataConfig` | `ValidationDataConfig` | yes |
| `outputDataConfig` | `OutputDataConfig` | yes |
| `customizationType` | `string` | no |
| `outputModelKmsKeyArn` | `string` | no |
| `trainingMetrics` | `TrainingMetrics` | no |
| `validationMetrics` | `List<ValidatorMetric>` | no |
| `vpcConfig` | `VpcConfig` | no |
| `customizationConfig` | `CustomizationConfig` | no |

## GetModelImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | no |
| `jobName` | `string` | no |
| `importedModelName` | `string` | no |
| `importedModelArn` | `string` | no |
| `roleArn` | `string` | no |
| `modelDataSource` | `ModelDataSource` | no |
| `status` | `string` | no |
| `failureMessage` | `string` | no |
| `creationTime` | `timestamp` | no |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `vpcConfig` | `VpcConfig` | no |
| `importedModelKmsKeyArn` | `string` | no |

## GetModelInvocationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobArn` | `string` | yes |
| `jobName` | `string` | no |
| `modelId` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `roleArn` | `string` | yes |
| `status` | `string` | no |
| `message` | `string` | no |
| `submitTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `inputDataConfig` | `ModelInvocationJobInputDataConfig` | yes |
| `outputDataConfig` | `ModelInvocationJobOutputDataConfig` | yes |
| `vpcConfig` | `VpcConfig` | no |
| `timeoutDurationInHours` | `integer` | no |
| `jobExpirationTime` | `timestamp` | no |
| `modelInvocationType` | `string` | no |
| `totalRecordCount` | `long` | no |
| `processedRecordCount` | `long` | no |
| `successRecordCount` | `long` | no |
| `errorRecordCount` | `long` | no |

## GetModelInvocationLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingConfig` | `LoggingConfig` | no |

## GetPromptRouter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptRouterArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptRouterName` | `string` | yes |
| `routingCriteria` | `RoutingCriteria` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `promptRouterArn` | `string` | yes |
| `models` | `List<PromptRouterTargetModel>` | yes |
| `fallbackModel` | `PromptRouterTargetModel` | yes |
| `status` | `string` | yes |
| `type` | `string` | yes |

## GetProvisionedModelThroughput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provisionedModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelUnits` | `integer` | yes |
| `desiredModelUnits` | `integer` | yes |
| `provisionedModelName` | `string` | yes |
| `provisionedModelArn` | `string` | yes |
| `modelArn` | `string` | yes |
| `desiredModelArn` | `string` | yes |
| `foundationModelArn` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `lastModifiedTime` | `timestamp` | yes |
| `failureMessage` | `string` | no |
| `commitmentDuration` | `string` | no |
| `commitmentExpirationTime` | `timestamp` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourcePolicy` | `string` | no |

## GetUseCaseForModelAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `formData` | `blob` | yes |

## ListAdvancedPromptOptimizationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobSummaries` | `List<AdvancedPromptOptimizationJobSummary>` | no |
| `nextToken` | `string` | no |

## ListAutomatedReasoningPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automatedReasoningPolicySummaries` | `List<AutomatedReasoningPolicySummary>` | yes |
| `nextToken` | `string` | no |

## ListAutomatedReasoningPolicyBuildWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `automatedReasoningPolicyBuildWorkflowSummaries` | `List<AutomatedReasoningPolicyBuildWorkflowSummary>` | yes |
| `nextToken` | `string` | no |

## ListAutomatedReasoningPolicyTestCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testCases` | `List<AutomatedReasoningPolicyTestCase>` | yes |
| `nextToken` | `string` | no |

## ListAutomatedReasoningPolicyTestResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `testResults` | `List<AutomatedReasoningPolicyTestResult>` | yes |
| `nextToken` | `string` | no |

## ListCustomModelDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdBefore` | `timestamp` | no |
| `createdAfter` | `timestamp` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `statusEquals` | `string` | no |
| `modelArnEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelDeploymentSummaries` | `List<CustomModelDeploymentSummary>` | no |

## ListCustomModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeBefore` | `timestamp` | no |
| `creationTimeAfter` | `timestamp` | no |
| `nameContains` | `string` | no |
| `baseModelArnEquals` | `string` | no |
| `foundationModelArnEquals` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |
| `isOwned` | `boolean` | no |
| `modelStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelSummaries` | `List<CustomModelSummary>` | no |

## ListEnforcedGuardrailsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailsConfig` | `List<AccountEnforcedGuardrailOutputConfiguration>` | yes |
| `nextToken` | `string` | no |

## ListEvaluationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `applicationTypeEquals` | `string` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `jobSummaries` | `List<EvaluationSummary>` | no |

## ListFoundationModelAgreementOffers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `offerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelId` | `string` | yes |
| `offers` | `List<Offer>` | yes |

## ListFoundationModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `byProvider` | `string` | no |
| `byCustomizationType` | `string` | no |
| `byOutputModality` | `string` | no |
| `byInferenceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelSummaries` | `List<FoundationModelSummary>` | no |

## ListGuardrails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrails` | `List<GuardrailSummary>` | yes |
| `nextToken` | `string` | no |

## ListImportedModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeBefore` | `timestamp` | no |
| `creationTimeAfter` | `timestamp` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelSummaries` | `List<ImportedModelSummary>` | no |

## ListInferenceProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `typeEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `inferenceProfileSummaries` | `List<InferenceProfileSummary>` | no |
| `nextToken` | `string` | no |

## ListMarketplaceModelEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `modelSourceEquals` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceModelEndpoints` | `List<MarketplaceModelEndpointSummary>` | no |
| `nextToken` | `string` | no |

## ListModelCopyJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `sourceAccountEquals` | `string` | no |
| `sourceModelArnEquals` | `string` | no |
| `targetModelNameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelCopyJobSummaries` | `List<ModelCopyJobSummary>` | no |

## ListModelCustomizationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelCustomizationJobSummaries` | `List<ModelCustomizationJobSummary>` | no |

## ListModelImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `modelImportJobSummaries` | `List<ModelImportJobSummary>` | no |

## ListModelInvocationJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `submitTimeAfter` | `timestamp` | no |
| `submitTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `invocationJobSummaries` | `List<ModelInvocationJobSummary>` | no |

## ListPromptRouters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `promptRouterSummaries` | `List<PromptRouterSummary>` | no |
| `nextToken` | `string` | no |

## ListProvisionedModelThroughputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `creationTimeAfter` | `timestamp` | no |
| `creationTimeBefore` | `timestamp` | no |
| `statusEquals` | `string` | no |
| `modelArnEquals` | `string` | no |
| `nameContains` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `provisionedModelSummaries` | `List<ProvisionedModelSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## PutAccountDataRetention

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mode` | `string` | yes |
| `updatedAt` | `timestamp` | no |

## PutEnforcedGuardrailConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | no |
| `guardrailInferenceConfig` | `AccountEnforcedGuardrailInferenceInputConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configId` | `string` | no |
| `updatedAt` | `timestamp` | no |
| `updatedBy` | `string` | no |

## PutModelInvocationLoggingConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `loggingConfig` | `LoggingConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `resourcePolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | no |

## PutUseCaseForModelAccess

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `formData` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointIdentifier` | `string` | yes |
| `modelSourceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceModelEndpoint` | `MarketplaceModelEndpoint` | yes |

## StartAutomatedReasoningPolicyBuildWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowType` | `string` | yes |
| `clientRequestToken` | `string` | no |
| `sourceContent` | `AutomatedReasoningPolicyBuildWorkflowSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |

## StartAutomatedReasoningPolicyTestWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `testCaseIds` | `List<string>` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |

## StopAdvancedPromptOptimizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopEvaluationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopModelCustomizationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopModelInvocationJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceARN` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAutomatedReasoningPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `policyDefinition` | `AutomatedReasoningPolicyDefinition` | yes |
| `name` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `name` | `string` | yes |
| `definitionHash` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateAutomatedReasoningPolicyAnnotations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `annotations` | `List<AutomatedReasoningPolicyAnnotation>` | yes |
| `lastUpdatedAnnotationSetHash` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `buildWorkflowId` | `string` | yes |
| `annotationSetHash` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateAutomatedReasoningPolicyTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCaseId` | `string` | yes |
| `guardContent` | `string` | yes |
| `queryContent` | `string` | no |
| `lastUpdatedAt` | `timestamp` | yes |
| `expectedAggregatedFindingsResult` | `string` | yes |
| `confidenceThreshold` | `double` | no |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyArn` | `string` | yes |
| `testCaseId` | `string` | yes |

## UpdateCustomModelDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `modelArn` | `string` | yes |
| `customModelDeploymentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `customModelDeploymentArn` | `string` | yes |

## UpdateGuardrail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `topicPolicyConfig` | `GuardrailTopicPolicyConfig` | no |
| `contentPolicyConfig` | `GuardrailContentPolicyConfig` | no |
| `wordPolicyConfig` | `GuardrailWordPolicyConfig` | no |
| `sensitiveInformationPolicyConfig` | `GuardrailSensitiveInformationPolicyConfig` | no |
| `contextualGroundingPolicyConfig` | `GuardrailContextualGroundingPolicyConfig` | no |
| `automatedReasoningPolicyConfig` | `GuardrailAutomatedReasoningPolicyConfig` | no |
| `crossRegionConfig` | `GuardrailCrossRegionConfig` | no |
| `blockedInputMessaging` | `string` | yes |
| `blockedOutputsMessaging` | `string` | yes |
| `kmsKeyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `guardrailId` | `string` | yes |
| `guardrailArn` | `string` | yes |
| `version` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateMarketplaceModelEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointArn` | `string` | yes |
| `endpointConfig` | `EndpointConfig` | yes |
| `clientRequestToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `marketplaceModelEndpoint` | `MarketplaceModelEndpoint` | yes |

## UpdateProvisionedModelThroughput

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `provisionedModelId` | `string` | yes |
| `desiredProvisionedModelName` | `string` | no |
| `desiredModelId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


