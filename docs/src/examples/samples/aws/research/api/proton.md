# AWS Proton

API version: 2020-07-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/proton/2020-07-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | yes |

## CancelComponentDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | yes |

## CancelEnvironmentDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | yes |

## CancelServiceInstanceDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstanceName` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstance` | `ServiceInstance` | yes |

## CancelServicePipelineDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `ServicePipeline` | yes |

## CreateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `environmentName` | `string` | no |
| `manifest` | `string` | yes |
| `name` | `string` | yes |
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | no |
| `serviceSpec` | `string` | no |
| `tags` | `List<Tag>` | no |
| `templateFile` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | yes |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codebuildRoleArn` | `string` | no |
| `componentRoleArn` | `string` | no |
| `description` | `string` | no |
| `environmentAccountConnectionId` | `string` | no |
| `name` | `string` | yes |
| `protonServiceRoleArn` | `string` | no |
| `provisioningRepository` | `RepositoryBranchInput` | no |
| `spec` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `templateMajorVersion` | `string` | yes |
| `templateMinorVersion` | `string` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | yes |

## CreateEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `codebuildRoleArn` | `string` | no |
| `componentRoleArn` | `string` | no |
| `environmentName` | `string` | yes |
| `managementAccountId` | `string` | yes |
| `roleArn` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | yes |

## CreateEnvironmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `displayName` | `string` | no |
| `encryptionKey` | `string` | no |
| `name` | `string` | yes |
| `provisioning` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplate` | `EnvironmentTemplate` | yes |

## CreateEnvironmentTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `description` | `string` | no |
| `majorVersion` | `string` | no |
| `source` | `TemplateVersionSourceInput` | yes |
| `tags` | `List<Tag>` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplateVersion` | `EnvironmentTemplateVersion` | yes |

## CreateRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `connectionArn` | `string` | yes |
| `encryptionKey` | `string` | no |
| `name` | `string` | yes |
| `provider` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | yes |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branchName` | `string` | no |
| `description` | `string` | no |
| `name` | `string` | yes |
| `repositoryConnectionArn` | `string` | no |
| `repositoryId` | `string` | no |
| `spec` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `templateMajorVersion` | `string` | yes |
| `templateMinorVersion` | `string` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## CreateServiceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `name` | `string` | yes |
| `serviceName` | `string` | yes |
| `spec` | `string` | yes |
| `tags` | `List<Tag>` | no |
| `templateMajorVersion` | `string` | no |
| `templateMinorVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstance` | `ServiceInstance` | yes |

## CreateServiceSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `string` | yes |
| `filePath` | `string` | yes |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSyncConfig` | `ServiceSyncConfig` | no |

## CreateServiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `displayName` | `string` | no |
| `encryptionKey` | `string` | no |
| `name` | `string` | yes |
| `pipelineProvisioning` | `string` | no |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplate` | `ServiceTemplate` | yes |

## CreateServiceTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `compatibleEnvironmentTemplates` | `List<CompatibleEnvironmentTemplateInput>` | yes |
| `description` | `string` | no |
| `majorVersion` | `string` | no |
| `source` | `TemplateVersionSourceInput` | yes |
| `supportedComponentSources` | `List<string>` | no |
| `tags` | `List<Tag>` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplateVersion` | `ServiceTemplateVersion` | yes |

## CreateTemplateSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `string` | yes |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `subdirectory` | `string` | no |
| `templateName` | `string` | yes |
| `templateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateSyncConfig` | `TemplateSyncConfig` | no |

## DeleteComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | no |

## DeleteDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployment` | `Deployment` | no |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | no |

## DeleteEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | no |

## DeleteEnvironmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplate` | `EnvironmentTemplate` | no |

## DeleteEnvironmentTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplateVersion` | `EnvironmentTemplateVersion` | no |

## DeleteRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `provider` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | no |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | no |

## DeleteServiceSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSyncConfig` | `ServiceSyncConfig` | no |

## DeleteServiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplate` | `ServiceTemplate` | no |

## DeleteServiceTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplateVersion` | `ServiceTemplateVersion` | no |

## DeleteTemplateSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateSyncConfig` | `TemplateSyncConfig` | no |

## GetAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettings` | `AccountSettings` | no |

## GetComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | no |
| `environmentName` | `string` | no |
| `id` | `string` | yes |
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployment` | `Deployment` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | yes |

## GetEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | yes |

## GetEnvironmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplate` | `EnvironmentTemplate` | yes |

## GetEnvironmentTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplateVersion` | `EnvironmentTemplateVersion` | yes |

## GetRepository

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `provider` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `repository` | `Repository` | yes |

## GetRepositorySyncStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `string` | yes |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `syncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `latestSync` | `RepositorySyncAttempt` | no |

## GetResourcesSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `counts` | `CountsSummary` | yes |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | no |

## GetServiceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstance` | `ServiceInstance` | yes |

## GetServiceInstanceSyncStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstanceName` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `desiredState` | `Revision` | no |
| `latestSuccessfulSync` | `ResourceSyncAttempt` | no |
| `latestSync` | `ResourceSyncAttempt` | no |

## GetServiceSyncBlockerSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSyncBlockerSummary` | `ServiceSyncBlockerSummary` | no |

## GetServiceSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSyncConfig` | `ServiceSyncConfig` | no |

## GetServiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplate` | `ServiceTemplate` | yes |

## GetServiceTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplateVersion` | `ServiceTemplateVersion` | yes |

## GetTemplateSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateSyncConfig` | `TemplateSyncConfig` | no |

## GetTemplateSyncStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateName` | `string` | yes |
| `templateType` | `string` | yes |
| `templateVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `desiredState` | `Revision` | no |
| `latestSuccessfulSync` | `ResourceSyncAttempt` | no |
| `latestSync` | `ResourceSyncAttempt` | no |

## ListComponentOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | yes |
| `deploymentId` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `outputs` | `List<Output>` | yes |

## ListComponentProvisionedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `provisionedResources` | `List<ProvisionedResource>` | yes |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `components` | `List<ComponentSummary>` | yes |
| `nextToken` | `string` | no |

## ListDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `componentName` | `string` | no |
| `environmentName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deployments` | `List<DeploymentSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnvironmentAccountConnections

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentName` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `requestedBy` | `string` | yes |
| `statuses` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnections` | `List<EnvironmentAccountConnectionSummary>` | yes |
| `nextToken` | `string` | no |

## ListEnvironmentOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `outputs` | `List<Output>` | yes |

## ListEnvironmentProvisionedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentName` | `string` | yes |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `provisionedResources` | `List<ProvisionedResource>` | yes |

## ListEnvironmentTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templateVersions` | `List<EnvironmentTemplateVersionSummary>` | yes |

## ListEnvironmentTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templates` | `List<EnvironmentTemplateSummary>` | yes |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplates` | `List<EnvironmentTemplateFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environments` | `List<EnvironmentSummary>` | yes |
| `nextToken` | `string` | no |

## ListRepositories

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `repositories` | `List<RepositorySummary>` | yes |

## ListRepositorySyncDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `syncType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `syncDefinitions` | `List<RepositorySyncDefinition>` | yes |

## ListServiceInstanceOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `nextToken` | `string` | no |
| `serviceInstanceName` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `outputs` | `List<Output>` | yes |

## ListServiceInstanceProvisionedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serviceInstanceName` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `provisionedResources` | `List<ProvisionedResource>` | yes |

## ListServiceInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `filters` | `List<ListServiceInstancesFilter>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `serviceName` | `string` | no |
| `sortBy` | `string` | no |
| `sortOrder` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serviceInstances` | `List<ServiceInstanceSummary>` | yes |

## ListServicePipelineOutputs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `nextToken` | `string` | no |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `outputs` | `List<Output>` | yes |

## ListServicePipelineProvisionedResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `provisionedResources` | `List<ProvisionedResource>` | yes |

## ListServiceTemplateVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `majorVersion` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templateVersions` | `List<ServiceTemplateVersionSummary>` | yes |

## ListServiceTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `templates` | `List<ServiceTemplateSummary>` | yes |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `services` | `List<ServiceSummary>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `tags` | `List<Tag>` | yes |

## NotifyResourceDeploymentStatusChange

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `outputs` | `List<Output>` | no |
| `resourceArn` | `string` | yes |
| `status` | `string` | no |
| `statusMessage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RejectEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | yes |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateAccountSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deletePipelineProvisioningRepository` | `boolean` | no |
| `pipelineCodebuildRoleArn` | `string` | no |
| `pipelineProvisioningRepository` | `RepositoryBranchInput` | no |
| `pipelineServiceRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountSettings` | `AccountSettings` | yes |

## UpdateComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `deploymentType` | `string` | yes |
| `description` | `string` | no |
| `name` | `string` | yes |
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | no |
| `serviceSpec` | `string` | no |
| `templateFile` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `component` | `Component` | yes |

## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codebuildRoleArn` | `string` | no |
| `componentRoleArn` | `string` | no |
| `deploymentType` | `string` | yes |
| `description` | `string` | no |
| `environmentAccountConnectionId` | `string` | no |
| `name` | `string` | yes |
| `protonServiceRoleArn` | `string` | no |
| `provisioningRepository` | `RepositoryBranchInput` | no |
| `spec` | `string` | no |
| `templateMajorVersion` | `string` | no |
| `templateMinorVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environment` | `Environment` | yes |

## UpdateEnvironmentAccountConnection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `codebuildRoleArn` | `string` | no |
| `componentRoleArn` | `string` | no |
| `id` | `string` | yes |
| `roleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentAccountConnection` | `EnvironmentAccountConnection` | yes |

## UpdateEnvironmentTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `displayName` | `string` | no |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplate` | `EnvironmentTemplate` | yes |

## UpdateEnvironmentTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `status` | `string` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentTemplateVersion` | `EnvironmentTemplateVersion` | yes |

## UpdateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `name` | `string` | yes |
| `spec` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `service` | `Service` | yes |

## UpdateServiceInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | no |
| `deploymentType` | `string` | yes |
| `name` | `string` | yes |
| `serviceName` | `string` | yes |
| `spec` | `string` | no |
| `templateMajorVersion` | `string` | no |
| `templateMinorVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstance` | `ServiceInstance` | yes |

## UpdateServicePipeline

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentType` | `string` | yes |
| `serviceName` | `string` | yes |
| `spec` | `string` | yes |
| `templateMajorVersion` | `string` | no |
| `templateMinorVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `pipeline` | `ServicePipeline` | yes |

## UpdateServiceSyncBlocker

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `resolvedReason` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceInstanceName` | `string` | no |
| `serviceName` | `string` | yes |
| `serviceSyncBlocker` | `SyncBlocker` | yes |

## UpdateServiceSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `string` | yes |
| `filePath` | `string` | yes |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `serviceName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceSyncConfig` | `ServiceSyncConfig` | no |

## UpdateServiceTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `displayName` | `string` | no |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplate` | `ServiceTemplate` | yes |

## UpdateServiceTemplateVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `compatibleEnvironmentTemplates` | `List<CompatibleEnvironmentTemplateInput>` | no |
| `description` | `string` | no |
| `majorVersion` | `string` | yes |
| `minorVersion` | `string` | yes |
| `status` | `string` | no |
| `supportedComponentSources` | `List<string>` | no |
| `templateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `serviceTemplateVersion` | `ServiceTemplateVersion` | yes |

## UpdateTemplateSyncConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `branch` | `string` | yes |
| `repositoryName` | `string` | yes |
| `repositoryProvider` | `string` | yes |
| `subdirectory` | `string` | no |
| `templateName` | `string` | yes |
| `templateType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templateSyncConfig` | `TemplateSyncConfig` | no |

