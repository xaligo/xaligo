# Amazon EMR Containers

API version: 2020-10-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/emr-containers/2020-10-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CancelJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `virtualClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `virtualClusterId` | `string` | no |

## CreateJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `clientToken` | `string` | yes |
| `jobTemplateData` | `JobTemplateData` | yes |
| `tags` | `Map<string>` | no |
| `kmsKeyArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `createdAt` | `timestamp` | no |

## CreateManagedEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `virtualClusterId` | `string` | yes |
| `type` | `string` | yes |
| `releaseLabel` | `string` | yes |
| `executionRoleArn` | `string` | yes |
| `certificateArn` | `string` | no |
| `configurationOverrides` | `ConfigurationOverrides` | no |
| `clientToken` | `string` | yes |
| `tags` | `Map<string>` | no |
| `sessionIdleTimeoutInMinutes` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `virtualClusterId` | `string` | no |

## CreateSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `name` | `string` | yes |
| `containerProvider` | `ContainerProvider` | no |
| `securityConfigurationData` | `SecurityConfigurationData` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |

## CreateVirtualCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `containerProvider` | `ContainerProvider` | yes |
| `clientToken` | `string` | yes |
| `tags` | `Map<string>` | no |
| `securityConfigurationId` | `string` | no |
| `sessionEnabled` | `boolean` | no |
| `schedulerConfiguration` | `SchedulerConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |

## DeleteJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## DeleteManagedEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `virtualClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `virtualClusterId` | `string` | no |

## DeleteSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## DeleteVirtualCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |

## DescribeJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `virtualClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobRun` | `JobRun` | no |

## DescribeJobTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobTemplate` | `JobTemplate` | no |

## DescribeManagedEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `virtualClusterId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoint` | `Endpoint` | no |

## DescribeSecurityConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfiguration` | `SecurityConfiguration` | no |

## DescribeVirtualCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualCluster` | `VirtualCluster` | no |

## GetManagedEndpointSessionCredentials

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpointIdentifier` | `string` | yes |
| `virtualClusterIdentifier` | `string` | yes |
| `executionRoleArn` | `string` | yes |
| `credentialType` | `string` | yes |
| `durationInSeconds` | `integer` | no |
| `logContext` | `string` | no |
| `clientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `credentials` | `Credentials` | no |
| `endpointCredentials` | `Credentials` | no |
| `expiresAt` | `timestamp` | no |

## ListJobRuns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualClusterId` | `string` | yes |
| `createdBefore` | `timestamp` | no |
| `createdAfter` | `timestamp` | no |
| `name` | `string` | no |
| `states` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `jobRuns` | `List<JobRun>` | no |
| `nextToken` | `string` | no |

## ListJobTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAfter` | `timestamp` | no |
| `createdBefore` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `templates` | `List<JobTemplate>` | no |
| `nextToken` | `string` | no |

## ListManagedEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualClusterId` | `string` | yes |
| `createdBefore` | `timestamp` | no |
| `createdAfter` | `timestamp` | no |
| `types` | `List<string>` | no |
| `states` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `endpoints` | `List<Endpoint>` | no |
| `nextToken` | `string` | no |

## ListSecurityConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `createdAfter` | `timestamp` | no |
| `createdBefore` | `timestamp` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `securityConfigurations` | `List<SecurityConfiguration>` | no |
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

## ListVirtualClusters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `containerProviderId` | `string` | no |
| `containerProviderType` | `string` | no |
| `createdAfter` | `timestamp` | no |
| `createdBefore` | `timestamp` | no |
| `states` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `eksAccessEntryIntegrated` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualClusters` | `List<VirtualCluster>` | no |
| `nextToken` | `string` | no |

## StartJobRun

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `virtualClusterId` | `string` | yes |
| `clientToken` | `string` | yes |
| `executionRoleArn` | `string` | no |
| `releaseLabel` | `string` | no |
| `jobDriver` | `JobDriver` | no |
| `configurationOverrides` | `ConfigurationOverrides` | no |
| `tags` | `Map<string>` | no |
| `jobTemplateId` | `string` | no |
| `jobTemplateParameters` | `Map<string>` | no |
| `retryPolicyConfiguration` | `RetryPolicyConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `arn` | `string` | no |
| `virtualClusterId` | `string` | no |

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


## UpdateVirtualCluster

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |
| `schedulerConfiguration` | `SchedulerConfiguration` | no |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `virtualCluster` | `VirtualCluster` | no |

