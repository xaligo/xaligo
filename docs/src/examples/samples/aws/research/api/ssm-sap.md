# AWS Systems Manager for SAP

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm-sap/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeleteResourcePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionType` | `string` | no |
| `SourceResourceArn` | `string` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## DeregisterApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ApplicationArn` | `string` | no |
| `AppRegistryArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `Application` | no |
| `Tags` | `Map<string>` | no |

## GetComponent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ComponentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Component` | `Component` | no |
| `Tags` | `Map<string>` | no |

## GetConfigurationCheckOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationCheckOperation` | `ConfigurationCheckOperation` | no |

## GetDatabase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ComponentId` | `string` | no |
| `DatabaseId` | `string` | no |
| `DatabaseArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Database` | `Database` | no |
| `Tags` | `Map<string>` | no |

## GetOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operation` | `Operation` | no |

## GetResourcePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionType` | `string` | no |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<ApplicationSummary>` | no |
| `NextToken` | `string` | no |

## ListComponents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Components` | `List<ComponentSummary>` | no |
| `NextToken` | `string` | no |

## ListConfigurationCheckDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationChecks` | `List<ConfigurationCheckDefinition>` | no |
| `NextToken` | `string` | no |

## ListConfigurationCheckOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ListMode` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationCheckOperations` | `List<ConfigurationCheckOperation>` | no |
| `NextToken` | `string` | no |

## ListDatabases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `ComponentId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Databases` | `List<DatabaseSummary>` | no |
| `NextToken` | `string` | no |

## ListOperationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationEvents` | `List<OperationEvent>` | no |
| `NextToken` | `string` | no |

## ListOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Operations` | `List<Operation>` | no |
| `NextToken` | `string` | no |

## ListSubCheckResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubCheckResults` | `List<SubCheckResult>` | no |
| `NextToken` | `string` | no |

## ListSubCheckRuleResults

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubCheckResultId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleResults` | `List<RuleResult>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## PutResourcePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActionType` | `string` | yes |
| `SourceResourceArn` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## RegisterApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ApplicationType` | `string` | yes |
| `Instances` | `List<string>` | yes |
| `SapInstanceNumber` | `string` | no |
| `Sid` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Credentials` | `List<ApplicationCredential>` | no |
| `DatabaseArn` | `string` | no |
| `ComponentsInfo` | `List<ComponentInfo>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Application` | `Application` | no |
| `OperationId` | `string` | no |

## StartApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## StartApplicationRefresh

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

## StartConfigurationChecks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `ConfigurationCheckIds` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationCheckOperations` | `List<ConfigurationCheckOperation>` | no |

## StopApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `StopConnectedEntity` | `string` | no |
| `IncludeEc2InstanceShutdown` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OperationId` | `string` | no |

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


## UpdateApplicationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | yes |
| `CredentialsToAddOrUpdate` | `List<ApplicationCredential>` | no |
| `CredentialsToRemove` | `List<ApplicationCredential>` | no |
| `Backint` | `BackintConfig` | no |
| `DatabaseArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |
| `OperationIds` | `List<string>` | no |

