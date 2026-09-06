# Amazon Bedrock AgentCore Control

API version: 2023-06-05. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/bedrock-agentcore-control/2023-06-05/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddDatasetExamples

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `clientToken` | `string` | no |
| `source` | `DataSourceType` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `addedCount` | `long` | yes |
| `updatedAt` | `timestamp` | yes |
| `exampleIds` | `List<string>` | yes |

## BatchPutGatewayRateLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `rateLimits` | `List<BatchPutLimitEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimits` | `List<GatewayRateLimitDetail>` | yes |

## CreateAgentRuntime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeName` | `string` | yes |
| `agentRuntimeArtifact` | `AgentRuntimeArtifact` | yes |
| `roleArn` | `string` | yes |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `requestHeaderConfiguration` | `RequestHeaderConfiguration` | no |
| `protocolConfiguration` | `ProtocolConfiguration` | no |
| `lifecycleConfiguration` | `LifecycleConfiguration` | no |
| `environmentVariables` | `Map<string>` | no |
| `filesystemConfigurations` | `List<FilesystemConfiguration>` | no |
| `capacityProviderConfiguration` | `CapacityProviderConfiguration` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeArn` | `string` | yes |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `agentRuntimeId` | `string` | yes |
| `agentRuntimeVersion` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateAgentRuntimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `name` | `string` | yes |
| `agentRuntimeVersion` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targetVersion` | `string` | yes |
| `agentRuntimeEndpointArn` | `string` | yes |
| `agentRuntimeArn` | `string` | yes |
| `agentRuntimeId` | `string` | no |
| `endpointName` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateApiKeyCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `apiKey` | `string` | no |
| `apiKeySecretConfig` | `SecretReference` | no |
| `apiKeySecretSource` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKeySecretArn` | `Secret` | yes |
| `apiKeySecretJsonKey` | `string` | no |
| `apiKeySecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |

## CreateBrowser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | no |
| `networkConfiguration` | `BrowserNetworkConfiguration` | yes |
| `recording` | `RecordingConfig` | no |
| `browserSigning` | `BrowserSigningConfigInput` | no |
| `enterprisePolicies` | `List<BrowserEnterprisePolicy>` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserId` | `string` | yes |
| `browserArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateBrowserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `permissionsConfiguration` | `PermissionsConfiguration` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `computeConfiguration` | `ComputeConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `capacityProviderArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |

## CreateCodeInterpreter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | no |
| `networkConfiguration` | `CodeInterpreterNetworkConfiguration` | yes |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterId` | `string` | yes |
| `codeInterpreterArn` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateConfigurationBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `bundleName` | `string` | yes |
| `description` | `string` | no |
| `components` | `Map<ComponentConfiguration>` | yes |
| `branchName` | `string` | no |
| `commitMessage` | `string` | no |
| `createdBy` | `VersionCreatedBySource` | no |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleArn` | `string` | yes |
| `bundleId` | `string` | yes |
| `versionId` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateConsentPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `executionRoleArn` | `string` | yes |
| `idpConfig` | `ConsentPortalIdpConfig` | yes |
| `name` | `string` | yes |
| `sources` | `List<ConsentPortalSource>` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<ConsentPortalSource>` | yes |
| `consentPortalArn` | `string` | yes |
| `consentPortalId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `idpConfig` | `ConsentPortalIdpConfig` | yes |
| `name` | `string` | yes |
| `portalUrl` | `string` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `updatedAt` | `timestamp` | yes |

## CreateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `datasetName` | `string` | yes |
| `description` | `string` | no |
| `source` | `DataSourceType` | yes |
| `schemaType` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateDatasetVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateEvaluator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `evaluatorName` | `string` | yes |
| `description` | `string` | no |
| `evaluatorConfig` | `EvaluatorConfig` | yes |
| `level` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorArn` | `string` | yes |
| `evaluatorId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |

## CreateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `roleArn` | `string` | yes |
| `protocolType` | `string` | no |
| `protocolConfiguration` | `GatewayProtocolConfiguration` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `kmsKeyArn` | `string` | no |
| `interceptorConfigurations` | `List<GatewayInterceptorConfiguration>` | no |
| `policyEngineConfiguration` | `GatewayPolicyEngineConfiguration` | no |
| `exceptionLevel` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `gatewayId` | `string` | yes |
| `gatewayUrl` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `protocolType` | `string` | no |
| `protocolConfiguration` | `GatewayProtocolConfiguration` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `kmsKeyArn` | `string` | no |
| `customTransformConfiguration` | `CustomTransformConfiguration` | no |
| `interceptorConfigurations` | `List<GatewayInterceptorConfiguration>` | no |
| `policyEngineConfiguration` | `GatewayPolicyEngineConfiguration` | no |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `exceptionLevel` | `string` | no |
| `webAclArn` | `string` | no |
| `wafConfiguration` | `WafConfiguration` | no |

## CreateGatewayRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `rateLimitId` | `string` | no |
| `description` | `string` | no |
| `dimensionKeys` | `List<string>` | yes |
| `entries` | `List<LimitEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimitId` | `string` | yes |
| `gatewayIdentifier` | `string` | yes |
| `description` | `string` | no |
| `dimensionKeys` | `List<string>` | yes |
| `entries` | `List<LimitEntry>` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## CreateGatewayRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `clientToken` | `string` | no |
| `priority` | `integer` | yes |
| `conditions` | `List<Condition>` | no |
| `actions` | `List<Action>` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `gatewayArn` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `List<Condition>` | no |
| `actions` | `List<Action>` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `system` | `SystemManagedBlock` | no |

## CreateGatewayTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `targetConfiguration` | `TargetConfiguration` | yes |
| `credentialProviderConfigurations` | `List<CredentialProviderConfiguration>` | no |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `privateEndpoint` | `PrivateEndpoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `targetId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `targetConfiguration` | `TargetConfiguration` | yes |
| `credentialProviderConfigurations` | `List<CredentialProviderConfiguration>` | yes |
| `lastSynchronizedAt` | `timestamp` | no |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `privateEndpoint` | `PrivateEndpoint` | no |
| `privateEndpointManagedResources` | `List<ManagedResourceDetails>` | no |
| `authorizationData` | `AuthorizationData` | no |
| `protocolType` | `string` | no |

## CreateHarness

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessName` | `string` | yes |
| `clientToken` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `environment` | `HarnessEnvironmentProviderRequest` | no |
| `environmentArtifact` | `HarnessEnvironmentArtifact` | no |
| `environmentVariables` | `Map<string>` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `model` | `HarnessModelConfiguration` | no |
| `systemPrompt` | `List<HarnessSystemContentBlock>` | no |
| `tools` | `List<HarnessTool>` | no |
| `skills` | `List<HarnessSkill>` | no |
| `allowedTools` | `List<string>` | no |
| `memory` | `HarnessMemoryConfiguration` | no |
| `truncation` | `HarnessTruncationConfiguration` | no |
| `maxIterations` | `integer` | no |
| `maxTokens` | `integer` | no |
| `timeoutSeconds` | `integer` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harness` | `Harness` | yes |

## CreateHarnessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `endpointName` | `string` | yes |
| `targetVersion` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `HarnessEndpoint` | yes |

## CreateMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `encryptionKeyArn` | `string` | no |
| `memoryExecutionRoleArn` | `string` | no |
| `eventExpiryDuration` | `integer` | yes |
| `memoryStrategies` | `List<MemoryStrategyInput>` | no |
| `indexedKeys` | `List<IndexedKey>` | no |
| `namespaceKeys` | `List<NamespaceKeyEntry>` | no |
| `streamDeliveryResources` | `StreamDeliveryResources` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memory` | `Memory` | no |

## CreateOauth2CredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `oauth2ProviderConfigInput` | `Oauth2ProviderConfigInput` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientSecretArn` | `Secret` | yes |
| `clientSecretJsonKey` | `string` | no |
| `clientSecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `callbackUrl` | `string` | no |
| `oauth2ProviderConfigOutput` | `Oauth2ProviderConfigOutput` | no |
| `status` | `string` | no |

## CreateOnlineEvaluationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `onlineEvaluationConfigName` | `string` | yes |
| `description` | `string` | no |
| `rule` | `Rule` | yes |
| `dataSourceConfig` | `DataSourceConfig` | yes |
| `evaluators` | `List<EvaluatorReference>` | no |
| `insights` | `List<Insight>` | no |
| `clusteringConfig` | `ClusteringConfig` | no |
| `outputConfig` | `OutputConfig` | no |
| `evaluationExecutionRoleArn` | `string` | yes |
| `enableOnCreate` | `boolean` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigArn` | `string` | yes |
| `onlineEvaluationConfigId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `outputConfig` | `OutputConfig` | no |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `failureReason` | `string` | no |

## CreatePaymentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `credentialProviderConfigurations` | `List<CredentialsProviderConfiguration>` | yes |
| `provisionMode` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentConnectorId` | `string` | yes |
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `credentialProviderConfigurations` | `List<CredentialsProviderConfiguration>` | yes |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `authorizationUrl` | `string` | no |

## CreatePaymentCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `providerConfigurationInput` | `PaymentProviderConfigurationInput` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `providerConfigurationOutput` | `PaymentProviderConfigurationOutput` | yes |

## CreatePaymentManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `roleArn` | `string` | yes |
| `clientToken` | `string` | no |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerArn` | `string` | yes |
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `roleArn` | `string` | yes |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

## CreatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `definition` | `PolicyDefinition` | yes |
| `description` | `string` | no |
| `validationMode` | `string` | no |
| `enforcementMode` | `string` | no |
| `policyEngineId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `name` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyArn` | `string` | yes |
| `status` | `string` | yes |
| `enforcementMode` | `string` | no |
| `definition` | `PolicyDefinition` | yes |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## CreatePolicyEngine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `clientToken` | `string` | no |
| `encryptionKeyArn` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyEngineArn` | `string` | yes |
| `status` | `string` | yes |
| `encryptionKeyArn` | `string` | no |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## CreateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `authorizerType` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `clientToken` | `string` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |

## CreateRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `descriptorType` | `string` | yes |
| `descriptors` | `Descriptors` | no |
| `recordVersion` | `string` | no |
| `synchronizationType` | `string` | no |
| `synchronizationConfiguration` | `SynchronizationConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `recordArn` | `string` | yes |
| `status` | `string` | yes |

## CreateWorkloadIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `allowedResourceOauth2ReturnUrls` | `List<string>` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `workloadIdentityArn` | `string` | yes |
| `allowedResourceOauth2ReturnUrls` | `List<string>` | no |

## DeleteAgentRuntime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `agentRuntimeVersion` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `agentRuntimeId` | `string` | no |
| `agentRuntimeVersion` | `string` | no |

## DeleteAgentRuntimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `endpointName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `agentRuntimeId` | `string` | no |
| `endpointName` | `string` | no |

## DeleteApiKeyCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBrowser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserId` | `string` | yes |
| `status` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## DeleteBrowserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `status` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `lastSavedAt` | `timestamp` | no |

## DeleteCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `status` | `string` | yes |

## DeleteCodeInterpreter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterId` | `string` | yes |
| `status` | `string` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## DeleteConfigurationBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleId` | `string` | yes |
| `status` | `string` | yes |

## DeleteConsentPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consentPortalIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteDatasetExamples

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `clientToken` | `string` | no |
| `exampleIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `deletedCount` | `long` | yes |
| `updatedAt` | `timestamp` | yes |

## DeleteEvaluator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorArn` | `string` | yes |
| `evaluatorId` | `string` | yes |
| `status` | `string` | yes |

## DeleteGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |

## DeleteGatewayRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `rateLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimitId` | `string` | yes |
| `status` | `string` | yes |

## DeleteGatewayRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `ruleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `status` | `string` | yes |

## DeleteGatewayTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `targetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `targetId` | `string` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |

## DeleteHarness

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `clientToken` | `string` | no |
| `deleteManagedMemory` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harness` | `Harness` | no |

## DeleteHarnessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `endpointName` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `HarnessEndpoint` | yes |

## DeleteMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `memoryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `status` | `string` | no |

## DeleteOauth2CredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteOnlineEvaluationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigArn` | `string` | yes |
| `onlineEvaluationConfigId` | `string` | yes |
| `status` | `string` | yes |

## DeletePaymentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `paymentConnectorId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `paymentConnectorId` | `string` | no |

## DeletePaymentCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePaymentManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `paymentManagerId` | `string` | no |

## DeletePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `name` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyArn` | `string` | yes |
| `status` | `string` | yes |
| `enforcementMode` | `string` | no |
| `definition` | `PolicyDefinition` | yes |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## DeletePolicyEngine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyEngineArn` | `string` | yes |
| `status` | `string` | yes |
| `encryptionKeyArn` | `string` | no |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## DeleteRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |

## DeleteRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

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


## DeleteWorkloadIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAgentRuntime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `agentRuntimeVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeArn` | `string` | yes |
| `agentRuntimeName` | `string` | yes |
| `agentRuntimeId` | `string` | yes |
| `agentRuntimeVersion` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `roleArn` | `string` | yes |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `status` | `string` | yes |
| `lifecycleConfiguration` | `LifecycleConfiguration` | yes |
| `failureReason` | `string` | no |
| `description` | `string` | no |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `agentRuntimeArtifact` | `AgentRuntimeArtifact` | no |
| `protocolConfiguration` | `ProtocolConfiguration` | no |
| `environmentVariables` | `Map<string>` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `requestHeaderConfiguration` | `RequestHeaderConfiguration` | no |
| `metadataConfiguration` | `RuntimeMetadataConfiguration` | no |
| `filesystemConfigurations` | `List<FilesystemConfiguration>` | no |
| `capacityProviderConfiguration` | `CapacityProviderConfiguration` | no |

## GetAgentRuntimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `endpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `liveVersion` | `string` | no |
| `targetVersion` | `string` | no |
| `agentRuntimeEndpointArn` | `string` | yes |
| `agentRuntimeArn` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `failureReason` | `string` | no |
| `name` | `string` | yes |
| `id` | `string` | yes |

## GetApiKeyCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKeySecretArn` | `Secret` | yes |
| `apiKeySecretJsonKey` | `string` | no |
| `apiKeySecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |

## GetBrowser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserId` | `string` | yes |
| `browserArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | no |
| `networkConfiguration` | `BrowserNetworkConfiguration` | yes |
| `recording` | `RecordingConfig` | no |
| `browserSigning` | `BrowserSigningConfigOutput` | no |
| `enterprisePolicies` | `List<BrowserEnterprisePolicy>` | no |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `status` | `string` | yes |
| `failureReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## GetBrowserProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileId` | `string` | yes |
| `profileArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `lastSavedAt` | `timestamp` | no |
| `lastSavedBrowserSessionId` | `string` | no |
| `lastSavedBrowserId` | `string` | no |

## GetCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `capacityProviderArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `statusCode` | `string` | no |
| `statusReason` | `string` | no |
| `permissionsConfiguration` | `PermissionsConfiguration` | yes |
| `computeConfiguration` | `ComputeConfiguration` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## GetCodeInterpreter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterId` | `string` | yes |
| `codeInterpreterArn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | no |
| `networkConfiguration` | `CodeInterpreterNetworkConfiguration` | yes |
| `status` | `string` | yes |
| `certificates` | `List<Certificate>` | no |
| `filesystemConfigurations` | `List<ToolsFileSystemConfiguration>` | no |
| `failureReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## GetConfigurationBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleId` | `string` | yes |
| `branchName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleArn` | `string` | yes |
| `bundleId` | `string` | yes |
| `bundleName` | `string` | yes |
| `description` | `string` | no |
| `versionId` | `string` | yes |
| `components` | `Map<ComponentConfiguration>` | yes |
| `lineageMetadata` | `VersionLineageMetadata` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `kmsKeyArn` | `string` | no |

## GetConfigurationBundleVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleId` | `string` | yes |
| `versionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleArn` | `string` | yes |
| `bundleId` | `string` | yes |
| `bundleName` | `string` | yes |
| `description` | `string` | no |
| `versionId` | `string` | yes |
| `components` | `Map<ComponentConfiguration>` | yes |
| `lineageMetadata` | `VersionLineageMetadata` | no |
| `createdAt` | `timestamp` | yes |
| `versionCreatedAt` | `timestamp` | yes |
| `kmsKeyArn` | `string` | no |

## GetConsentPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consentPortalIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<ConsentPortalSource>` | yes |
| `consentPortalArn` | `string` | yes |
| `consentPortalId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `idpConfig` | `ConsentPortalIdpConfig` | yes |
| `name` | `string` | yes |
| `portalUrl` | `string` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `updatedAt` | `timestamp` | yes |

## GetDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `datasetName` | `string` | yes |
| `description` | `string` | no |
| `status` | `string` | yes |
| `draftStatus` | `string` | no |
| `failureReason` | `string` | no |
| `schemaType` | `string` | yes |
| `kmsKeyArn` | `string` | no |
| `exampleCount` | `long` | yes |
| `downloadUrl` | `string` | no |
| `downloadUrlExpiresAt` | `timestamp` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetEvaluator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorId` | `string` | yes |
| `includedData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorArn` | `string` | yes |
| `evaluatorId` | `string` | yes |
| `evaluatorName` | `string` | yes |
| `description` | `string` | no |
| `evaluatorConfig` | `EvaluatorConfig` | yes |
| `evaluatorType` | `string` | no |
| `provider` | `string` | no |
| `level` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `lockedForModification` | `boolean` | no |
| `kmsKeyArn` | `string` | no |

## GetGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `gatewayId` | `string` | yes |
| `gatewayUrl` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `protocolType` | `string` | no |
| `protocolConfiguration` | `GatewayProtocolConfiguration` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `kmsKeyArn` | `string` | no |
| `customTransformConfiguration` | `CustomTransformConfiguration` | no |
| `interceptorConfigurations` | `List<GatewayInterceptorConfiguration>` | no |
| `policyEngineConfiguration` | `GatewayPolicyEngineConfiguration` | no |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `exceptionLevel` | `string` | no |
| `webAclArn` | `string` | no |
| `wafConfiguration` | `WafConfiguration` | no |

## GetGatewayRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `rateLimitId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimitId` | `string` | yes |
| `gatewayIdentifier` | `string` | yes |
| `description` | `string` | no |
| `dimensionKeys` | `List<string>` | yes |
| `entries` | `List<LimitEntry>` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetGatewayRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `ruleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `gatewayArn` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `List<Condition>` | no |
| `actions` | `List<Action>` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `system` | `SystemManagedBlock` | no |
| `updatedAt` | `timestamp` | no |

## GetGatewayTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `targetId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `targetId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `targetConfiguration` | `TargetConfiguration` | yes |
| `credentialProviderConfigurations` | `List<CredentialProviderConfiguration>` | yes |
| `lastSynchronizedAt` | `timestamp` | no |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `privateEndpoint` | `PrivateEndpoint` | no |
| `privateEndpointManagedResources` | `List<ManagedResourceDetails>` | no |
| `authorizationData` | `AuthorizationData` | no |
| `protocolType` | `string` | no |

## GetHarness

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `harnessVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harness` | `Harness` | yes |

## GetHarnessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `endpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `HarnessEndpoint` | yes |

## GetMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memoryId` | `string` | yes |
| `view` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memory` | `Memory` | yes |

## GetOauth2CredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientSecretArn` | `Secret` | yes |
| `clientSecretJsonKey` | `string` | no |
| `clientSecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `callbackUrl` | `string` | no |
| `oauth2ProviderConfigOutput` | `Oauth2ProviderConfigOutput` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |
| `status` | `string` | no |
| `failureReason` | `string` | no |

## GetOnlineEvaluationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigArn` | `string` | yes |
| `onlineEvaluationConfigId` | `string` | yes |
| `onlineEvaluationConfigName` | `string` | yes |
| `description` | `string` | no |
| `rule` | `Rule` | yes |
| `dataSourceConfig` | `DataSourceConfig` | yes |
| `evaluators` | `List<EvaluatorReference>` | no |
| `insights` | `List<Insight>` | no |
| `clusteringConfig` | `ClusteringConfig` | no |
| `outputConfig` | `OutputConfig` | no |
| `evaluationExecutionRoleArn` | `string` | no |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `failureReason` | `string` | no |

## GetPaymentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `paymentConnectorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentConnectorId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `credentialProviderConfigurations` | `List<CredentialsProviderConfiguration>` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `authorizationUrl` | `string` | no |

## GetPaymentCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `providerConfigurationOutput` | `PaymentProviderConfigurationOutput` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetPaymentManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerArn` | `string` | yes |
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `roleArn` | `string` | yes |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

## GetPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `name` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyArn` | `string` | yes |
| `status` | `string` | yes |
| `enforcementMode` | `string` | no |
| `definition` | `PolicyDefinition` | yes |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## GetPolicyEngine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyEngineArn` | `string` | yes |
| `status` | `string` | yes |
| `encryptionKeyArn` | `string` | no |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## GetPolicyEngineSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyEngineArn` | `string` | yes |
| `status` | `string` | yes |
| `encryptionKeyArn` | `string` | no |

## GetPolicyGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerationId` | `string` | yes |
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyGenerationId` | `string` | yes |
| `name` | `string` | yes |
| `policyGenerationArn` | `string` | yes |
| `resource` | `Resource` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `findings` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## GetPolicyGenerationSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerationId` | `string` | yes |
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyGenerationId` | `string` | yes |
| `name` | `string` | yes |
| `policyGenerationArn` | `string` | yes |
| `resource` | `Resource` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `findings` | `string` | no |

## GetPolicySummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `name` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyArn` | `string` | yes |
| `status` | `string` | yes |
| `enforcementMode` | `string` | no |

## GetRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `registryId` | `string` | yes |
| `registryArn` | `string` | yes |
| `authorizerType` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## GetRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `descriptorType` | `string` | yes |
| `descriptors` | `Descriptors` | yes |
| `recordVersion` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `statusReason` | `string` | no |
| `synchronizationType` | `string` | no |
| `synchronizationConfiguration` | `SynchronizationConfiguration` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | no |

## GetTokenVault

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenVaultId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenVaultId` | `string` | yes |
| `kmsConfiguration` | `KmsConfiguration` | yes |
| `lastModifiedDate` | `timestamp` | yes |

## GetWorkloadIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `workloadIdentityArn` | `string` | yes |
| `allowedResourceOauth2ReturnUrls` | `List<string>` | no |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |

## ListAgentRuntimeEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtimeEndpoints` | `List<AgentRuntimeEndpoint>` | yes |
| `nextToken` | `string` | no |

## ListAgentRuntimeVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimes` | `List<AgentRuntime>` | yes |
| `nextToken` | `string` | no |

## ListAgentRuntimeVersionsByCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimes` | `List<AgentRuntimeVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListAgentRuntimes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimes` | `List<AgentRuntime>` | yes |
| `nextToken` | `string` | no |

## ListApiKeyCredentialProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentialProviders` | `List<ApiKeyCredentialProviderItem>` | yes |
| `nextToken` | `string` | no |

## ListBrowserProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `profileSummaries` | `List<BrowserProfileSummary>` | yes |
| `nextToken` | `string` | no |

## ListBrowsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `browserSummaries` | `List<BrowserSummary>` | yes |
| `nextToken` | `string` | no |

## ListCapacityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviders` | `List<CapacityProviderSummary>` | yes |
| `nextToken` | `string` | no |

## ListCodeInterpreters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codeInterpreterSummaries` | `List<CodeInterpreterSummary>` | yes |
| `nextToken` | `string` | no |

## ListConfigurationBundleVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `filter` | `VersionFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versions` | `List<ConfigurationBundleVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListConfigurationBundles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundles` | `List<ConfigurationBundleSummary>` | yes |
| `nextToken` | `string` | no |

## ListConsentPortals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consentPortals` | `List<ConsentPortalSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasetExamples

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `datasetVersion` | `string` | yes |
| `examples` | `List<SensitiveJson>` | yes |
| `nextToken` | `string` | no |

## ListDatasetVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `versions` | `List<DatasetVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListDatasets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasets` | `List<DatasetSummary>` | yes |
| `nextToken` | `string` | no |

## ListEvaluators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluators` | `List<EvaluatorSummary>` | yes |
| `nextToken` | `string` | no |

## ListGatewayRateLimits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimits` | `List<GatewayRateLimitDetail>` | yes |
| `nextToken` | `string` | no |

## ListGatewayRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayRules` | `List<GatewayRuleDetail>` | yes |
| `nextToken` | `string` | no |

## ListGatewayTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<TargetSummary>` | yes |
| `nextToken` | `string` | no |

## ListGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `items` | `List<GatewaySummary>` | yes |
| `nextToken` | `string` | no |

## ListHarnessEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoints` | `List<HarnessEndpoint>` | yes |
| `nextToken` | `string` | no |

## ListHarnessVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessVersions` | `List<HarnessVersionSummary>` | yes |
| `nextToken` | `string` | no |

## ListHarnesses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnesses` | `List<HarnessSummary>` | yes |
| `nextToken` | `string` | no |

## ListMemories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memories` | `List<MemorySummary>` | yes |
| `nextToken` | `string` | no |

## ListOauth2CredentialProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentialProviders` | `List<Oauth2CredentialProviderItem>` | yes |
| `nextToken` | `string` | no |

## ListOnlineEvaluationConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigs` | `List<OnlineEvaluationConfigSummary>` | yes |
| `nextToken` | `string` | no |

## ListPaymentConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentConnectors` | `List<PaymentConnectorSummary>` | yes |
| `nextToken` | `string` | no |

## ListPaymentCredentialProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `credentialProviders` | `List<PaymentCredentialProviderItem>` | yes |
| `nextToken` | `string` | no |

## ListPaymentManagers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagers` | `List<PaymentManagerSummary>` | yes |
| `nextToken` | `string` | no |

## ListPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `policyEngineId` | `string` | yes |
| `targetResourceScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<Policy>` | yes |
| `nextToken` | `string` | no |

## ListPolicyEngineSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngines` | `List<PolicyEngineSummary>` | yes |
| `nextToken` | `string` | no |

## ListPolicyEngines

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngines` | `List<PolicyEngine>` | yes |
| `nextToken` | `string` | no |

## ListPolicyGenerationAssets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerationId` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerationAssets` | `List<PolicyGenerationAsset>` | no |
| `nextToken` | `string` | no |

## ListPolicyGenerationSummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerations` | `List<PolicyGenerationSummary>` | yes |
| `nextToken` | `string` | no |

## ListPolicyGenerations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `policyEngineId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyGenerations` | `List<PolicyGeneration>` | yes |
| `nextToken` | `string` | no |

## ListPolicySummaries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `policyEngineId` | `string` | yes |
| `targetResourceScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policies` | `List<PolicySummary>` | yes |
| `nextToken` | `string` | no |

## ListRegistries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `status` | `string` | no |
| `authorizerType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registries` | `List<RegistrySummary>` | yes |
| `nextToken` | `string` | no |

## ListRegistryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `name` | `string` | no |
| `status` | `string` | no |
| `descriptorType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryRecords` | `List<RegistryRecordSummary>` | yes |
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

## ListWorkloadIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `workloadIdentities` | `List<WorkloadIdentityType>` | yes |
| `nextToken` | `string` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policy` | `string` | yes |

## SetTokenVaultCMK

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenVaultId` | `string` | no |
| `kmsConfiguration` | `KmsConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tokenVaultId` | `string` | yes |
| `kmsConfiguration` | `KmsConfiguration` | yes |
| `lastModifiedDate` | `timestamp` | yes |

## StartPolicyGeneration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `resource` | `Resource` | yes |
| `content` | `Content` | yes |
| `name` | `string` | yes |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyGenerationId` | `string` | yes |
| `name` | `string` | yes |
| `policyGenerationArn` | `string` | yes |
| `resource` | `Resource` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `findings` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## SubmitRegistryRecordForApproval

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## SynchronizeGatewayTargets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `targetIdList` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `targets` | `List<GatewayTarget>` | no |

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


## UpdateAgentRuntime

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `agentRuntimeArtifact` | `AgentRuntimeArtifact` | yes |
| `roleArn` | `string` | yes |
| `networkConfiguration` | `NetworkConfiguration` | no |
| `description` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `requestHeaderConfiguration` | `RequestHeaderConfiguration` | no |
| `protocolConfiguration` | `ProtocolConfiguration` | no |
| `lifecycleConfiguration` | `LifecycleConfiguration` | no |
| `metadataConfiguration` | `RuntimeMetadataConfiguration` | no |
| `environmentVariables` | `Map<string>` | no |
| `filesystemConfigurations` | `List<FilesystemConfiguration>` | no |
| `capacityProviderConfiguration` | `CapacityProviderConfiguration` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeArn` | `string` | yes |
| `agentRuntimeId` | `string` | yes |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `agentRuntimeVersion` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## UpdateAgentRuntimeEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `agentRuntimeId` | `string` | yes |
| `endpointName` | `string` | yes |
| `agentRuntimeVersion` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `liveVersion` | `string` | no |
| `targetVersion` | `string` | no |
| `agentRuntimeEndpointArn` | `string` | yes |
| `agentRuntimeArn` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## UpdateApiKeyCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `apiKey` | `string` | no |
| `apiKeySecretConfig` | `SecretReference` | no |
| `apiKeySecretSource` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKeySecretArn` | `Secret` | yes |
| `apiKeySecretJsonKey` | `string` | no |
| `apiKeySecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |

## UpdateCapacityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `description` | `UpdatedDescription` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `capacityProviderId` | `string` | yes |
| `capacityProviderArn` | `string` | yes |
| `name` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `lastUpdatedAt` | `timestamp` | yes |

## UpdateConfigurationBundle

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `bundleId` | `string` | yes |
| `bundleName` | `string` | no |
| `description` | `string` | no |
| `components` | `Map<ComponentConfiguration>` | no |
| `parentVersionIds` | `List<string>` | yes |
| `branchName` | `string` | no |
| `commitMessage` | `string` | no |
| `createdBy` | `VersionCreatedBySource` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `bundleArn` | `string` | yes |
| `bundleId` | `string` | yes |
| `versionId` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateConsentPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `consentPortalIdentifier` | `string` | yes |
| `executionRoleArn` | `string` | no |
| `idpConfig` | `ConsentPortalIdpConfig` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<ConsentPortalSource>` | yes |
| `consentPortalArn` | `string` | yes |
| `consentPortalId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `description` | `string` | no |
| `executionRoleArn` | `string` | yes |
| `idpConfig` | `ConsentPortalIdpConfig` | yes |
| `name` | `string` | yes |
| `portalUrl` | `string` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `updatedAt` | `timestamp` | yes |

## UpdateDataset

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `clientToken` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateDatasetExamples

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetId` | `string` | yes |
| `clientToken` | `string` | no |
| `examples` | `List<SensitiveJson>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `datasetArn` | `string` | yes |
| `datasetId` | `string` | yes |
| `status` | `string` | yes |
| `updatedCount` | `long` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateEvaluator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `evaluatorId` | `string` | yes |
| `description` | `string` | no |
| `evaluatorConfig` | `EvaluatorConfig` | no |
| `level` | `string` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluatorArn` | `string` | yes |
| `evaluatorId` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |

## UpdateGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | yes |
| `protocolType` | `string` | no |
| `protocolConfiguration` | `GatewayProtocolConfiguration` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `kmsKeyArn` | `string` | no |
| `customTransformConfiguration` | `CustomTransformConfiguration` | no |
| `interceptorConfigurations` | `List<GatewayInterceptorConfiguration>` | no |
| `policyEngineConfiguration` | `GatewayPolicyEngineConfiguration` | no |
| `exceptionLevel` | `string` | no |
| `wafConfiguration` | `WafConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `gatewayId` | `string` | yes |
| `gatewayUrl` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `roleArn` | `string` | no |
| `protocolType` | `string` | no |
| `protocolConfiguration` | `GatewayProtocolConfiguration` | no |
| `authorizerType` | `string` | yes |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `kmsKeyArn` | `string` | no |
| `customTransformConfiguration` | `CustomTransformConfiguration` | no |
| `interceptorConfigurations` | `List<GatewayInterceptorConfiguration>` | no |
| `policyEngineConfiguration` | `GatewayPolicyEngineConfiguration` | no |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `exceptionLevel` | `string` | no |
| `webAclArn` | `string` | no |
| `wafConfiguration` | `WafConfiguration` | no |

## UpdateGatewayRateLimit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `rateLimitId` | `string` | yes |
| `description` | `string` | no |
| `entries` | `List<LimitEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rateLimitId` | `string` | yes |
| `gatewayIdentifier` | `string` | yes |
| `description` | `string` | no |
| `dimensionKeys` | `List<string>` | yes |
| `entries` | `List<LimitEntry>` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateGatewayRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `ruleId` | `string` | yes |
| `priority` | `integer` | no |
| `conditions` | `List<Condition>` | no |
| `actions` | `List<Action>` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `gatewayArn` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `List<Condition>` | no |
| `actions` | `List<Action>` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `system` | `SystemManagedBlock` | no |
| `updatedAt` | `timestamp` | no |

## UpdateGatewayTarget

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIdentifier` | `string` | yes |
| `targetId` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `targetConfiguration` | `TargetConfiguration` | yes |
| `credentialProviderConfigurations` | `List<CredentialProviderConfiguration>` | no |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `privateEndpoint` | `PrivateEndpoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayArn` | `string` | yes |
| `targetId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `statusReasons` | `List<string>` | no |
| `name` | `string` | yes |
| `description` | `string` | no |
| `targetConfiguration` | `TargetConfiguration` | yes |
| `credentialProviderConfigurations` | `List<CredentialProviderConfiguration>` | yes |
| `lastSynchronizedAt` | `timestamp` | no |
| `metadataConfiguration` | `MetadataConfiguration` | no |
| `privateEndpoint` | `PrivateEndpoint` | no |
| `privateEndpointManagedResources` | `List<ManagedResourceDetails>` | no |
| `authorizationData` | `AuthorizationData` | no |
| `protocolType` | `string` | no |

## UpdateHarness

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `clientToken` | `string` | no |
| `executionRoleArn` | `string` | no |
| `environment` | `HarnessEnvironmentProviderRequest` | no |
| `environmentArtifact` | `UpdatedHarnessEnvironmentArtifact` | no |
| `environmentVariables` | `Map<string>` | no |
| `authorizerConfiguration` | `UpdatedAuthorizerConfiguration` | no |
| `model` | `HarnessModelConfiguration` | no |
| `systemPrompt` | `List<HarnessSystemContentBlock>` | no |
| `tools` | `List<HarnessTool>` | no |
| `skills` | `List<HarnessSkill>` | no |
| `allowedTools` | `List<string>` | no |
| `memory` | `UpdatedHarnessMemoryConfiguration` | no |
| `truncation` | `HarnessTruncationConfiguration` | no |
| `maxIterations` | `integer` | no |
| `maxTokens` | `integer` | no |
| `timeoutSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harness` | `Harness` | yes |

## UpdateHarnessEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `harnessId` | `string` | yes |
| `endpointName` | `string` | yes |
| `targetVersion` | `string` | no |
| `description` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `HarnessEndpoint` | yes |

## UpdateMemory

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `memoryId` | `string` | yes |
| `description` | `string` | no |
| `eventExpiryDuration` | `integer` | no |
| `memoryExecutionRoleArn` | `string` | no |
| `memoryStrategies` | `ModifyMemoryStrategies` | no |
| `addIndexedKeys` | `List<IndexedKey>` | no |
| `namespaceKeys` | `List<NamespaceKeyEntry>` | no |
| `streamDeliveryResources` | `StreamDeliveryResources` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memory` | `Memory` | no |

## UpdateOauth2CredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `oauth2ProviderConfigInput` | `Oauth2ProviderConfigInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientSecretArn` | `Secret` | yes |
| `clientSecretJsonKey` | `string` | no |
| `clientSecretSource` | `string` | no |
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `callbackUrl` | `string` | no |
| `oauth2ProviderConfigOutput` | `Oauth2ProviderConfigOutput` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |
| `status` | `string` | no |

## UpdateOnlineEvaluationConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `onlineEvaluationConfigId` | `string` | yes |
| `description` | `string` | no |
| `rule` | `Rule` | no |
| `dataSourceConfig` | `DataSourceConfig` | no |
| `evaluators` | `List<EvaluatorReference>` | no |
| `insights` | `List<Insight>` | no |
| `clusteringConfig` | `ClusteringConfig` | no |
| `outputConfig` | `OutputConfig` | no |
| `evaluationExecutionRoleArn` | `string` | no |
| `executionStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `onlineEvaluationConfigArn` | `string` | yes |
| `onlineEvaluationConfigId` | `string` | yes |
| `updatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `executionStatus` | `string` | yes |
| `failureReason` | `string` | no |

## UpdatePaymentConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `paymentConnectorId` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | no |
| `credentialProviderConfigurations` | `List<CredentialsProviderConfiguration>` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentConnectorId` | `string` | yes |
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `credentialProviderConfigurations` | `List<CredentialsProviderConfiguration>` | yes |
| `lastUpdatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `authorizationUrl` | `string` | no |

## UpdatePaymentCredentialProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `providerConfigurationInput` | `PaymentProviderConfigurationInput` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `credentialProviderVendor` | `string` | yes |
| `credentialProviderArn` | `string` | yes |
| `providerConfigurationOutput` | `PaymentProviderConfigurationOutput` | yes |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |

## UpdatePaymentManager

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerId` | `string` | yes |
| `description` | `string` | no |
| `authorizerType` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `roleArn` | `string` | no |
| `clientToken` | `string` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `paymentManagerArn` | `string` | yes |
| `paymentManagerId` | `string` | yes |
| `name` | `string` | yes |
| `authorizerType` | `string` | yes |
| `roleArn` | `string` | yes |
| `workloadIdentityDetails` | `WorkloadIdentityDetails` | no |
| `lastUpdatedAt` | `timestamp` | yes |
| `status` | `string` | yes |
| `kmsKeyArn` | `string` | no |

## UpdatePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `policyId` | `string` | yes |
| `description` | `UpdatedDescription` | no |
| `definition` | `PolicyDefinition` | no |
| `validationMode` | `string` | no |
| `enforcementMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyId` | `string` | yes |
| `name` | `string` | yes |
| `policyEngineId` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyArn` | `string` | yes |
| `status` | `string` | yes |
| `enforcementMode` | `string` | no |
| `definition` | `PolicyDefinition` | yes |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## UpdatePolicyEngine

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `description` | `UpdatedDescription` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `policyEngineId` | `string` | yes |
| `name` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `policyEngineArn` | `string` | yes |
| `status` | `string` | yes |
| `encryptionKeyArn` | `string` | no |
| `description` | `string` | no |
| `statusReasons` | `List<string>` | yes |

## UpdateRegistry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `name` | `string` | no |
| `description` | `UpdatedDescription` | no |
| `authorizerConfiguration` | `UpdatedAuthorizerConfiguration` | no |
| `approvalConfiguration` | `UpdatedApprovalConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `registryId` | `string` | yes |
| `registryArn` | `string` | yes |
| `authorizerType` | `string` | no |
| `authorizerConfiguration` | `AuthorizerConfiguration` | no |
| `approvalConfiguration` | `ApprovalConfiguration` | no |
| `status` | `string` | yes |
| `statusReason` | `string` | no |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateRegistryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | no |
| `description` | `UpdatedDescription` | no |
| `descriptorType` | `string` | no |
| `descriptors` | `UpdatedDescriptors` | no |
| `recordVersion` | `string` | no |
| `synchronizationType` | `UpdatedSynchronizationType` | no |
| `synchronizationConfiguration` | `UpdatedSynchronizationConfiguration` | no |
| `triggerSynchronization` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `descriptorType` | `string` | yes |
| `descriptors` | `Descriptors` | yes |
| `recordVersion` | `string` | no |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `statusReason` | `string` | no |
| `synchronizationType` | `string` | no |
| `synchronizationConfiguration` | `SynchronizationConfiguration` | no |

## UpdateRegistryRecordStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryId` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `registryArn` | `string` | yes |
| `recordArn` | `string` | yes |
| `recordId` | `string` | yes |
| `status` | `string` | yes |
| `statusReason` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateWorkloadIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `allowedResourceOauth2ReturnUrls` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `workloadIdentityArn` | `string` | yes |
| `allowedResourceOauth2ReturnUrls` | `List<string>` | no |
| `createdTime` | `timestamp` | yes |
| `lastUpdatedTime` | `timestamp` | yes |

