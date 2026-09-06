# AmazonMWAA

API version: 2020-07-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/mwaa/2020-07-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateCliToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CliToken` | `string` | no |
| `WebServerHostname` | `string` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ExecutionRoleArn` | `string` | yes |
| `SourceBucketArn` | `string` | yes |
| `DagS3Path` | `string` | yes |
| `NetworkConfiguration` | `NetworkConfiguration` | yes |
| `PluginsS3Path` | `string` | no |
| `PluginsS3ObjectVersion` | `string` | no |
| `RequirementsS3Path` | `string` | no |
| `RequirementsS3ObjectVersion` | `string` | no |
| `StartupScriptS3Path` | `string` | no |
| `StartupScriptS3ObjectVersion` | `string` | no |
| `AirflowConfigurationOptions` | `Map<string>` | no |
| `EnvironmentClass` | `string` | no |
| `MaxWorkers` | `integer` | no |
| `KmsKey` | `string` | no |
| `AirflowVersion` | `string` | no |
| `LoggingConfiguration` | `LoggingConfigurationInput` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |
| `Tags` | `Map<string>` | no |
| `WebserverAccessMode` | `string` | no |
| `MinWorkers` | `integer` | no |
| `Schedulers` | `integer` | no |
| `EndpointManagement` | `string` | no |
| `MinWebservers` | `integer` | no |
| `MaxWebservers` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

## CreateWebLoginToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebToken` | `string` | no |
| `WebServerHostname` | `string` | no |
| `IamIdentity` | `string` | no |
| `AirflowIdentity` | `string` | no |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Environment` | `Environment` | no |

## InvokeRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Path` | `string` | yes |
| `Method` | `string` | yes |
| `QueryParameters` | `Document` | no |
| `Body` | `RestApiRequestBody` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RestApiStatusCode` | `integer` | no |
| `RestApiResponse` | `RestApiResponse` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Environments` | `List<string>` | yes |
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

## PublishMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentName` | `string` | yes |
| `MetricData` | `List<MetricDatum>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ExecutionRoleArn` | `string` | no |
| `AirflowConfigurationOptions` | `Map<string>` | no |
| `AirflowVersion` | `string` | no |
| `DagS3Path` | `string` | no |
| `EnvironmentClass` | `string` | no |
| `LoggingConfiguration` | `LoggingConfigurationInput` | no |
| `MaxWorkers` | `integer` | no |
| `MinWorkers` | `integer` | no |
| `MaxWebservers` | `integer` | no |
| `MinWebservers` | `integer` | no |
| `WorkerReplacementStrategy` | `string` | no |
| `NetworkConfiguration` | `UpdateNetworkConfigurationInput` | no |
| `PluginsS3Path` | `string` | no |
| `PluginsS3ObjectVersion` | `string` | no |
| `RequirementsS3Path` | `string` | no |
| `RequirementsS3ObjectVersion` | `string` | no |
| `Schedulers` | `integer` | no |
| `SourceBucketArn` | `string` | no |
| `StartupScriptS3Path` | `string` | no |
| `StartupScriptS3ObjectVersion` | `string` | no |
| `WebserverAccessMode` | `string` | no |
| `WeeklyMaintenanceWindowStart` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |

