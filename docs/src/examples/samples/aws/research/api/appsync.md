# AWS AppSync

API version: 2017-07-25. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/appsync/2017-07-25/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiAssociation` | `ApiAssociation` | no |

## AssociateMergedGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiIdentifier` | `string` | yes |
| `mergedApiIdentifier` | `string` | yes |
| `description` | `string` | no |
| `sourceApiAssociationConfig` | `SourceApiAssociationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociation` | `SourceApiAssociation` | no |

## AssociateSourceGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergedApiIdentifier` | `string` | yes |
| `sourceApiIdentifier` | `string` | yes |
| `description` | `string` | no |
| `sourceApiAssociationConfig` | `SourceApiAssociationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociation` | `SourceApiAssociation` | no |

## CreateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `ownerContact` | `string` | no |
| `tags` | `Map<string>` | no |
| `eventConfig` | `EventConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `api` | `Api` | no |

## CreateApiCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `ttl` | `long` | yes |
| `transitEncryptionEnabled` | `boolean` | no |
| `atRestEncryptionEnabled` | `boolean` | no |
| `apiCachingBehavior` | `string` | yes |
| `type` | `string` | yes |
| `healthMetricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiCache` | `ApiCache` | no |

## CreateApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `description` | `string` | no |
| `expires` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `ApiKey` | no |

## CreateChannelNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `subscribeAuthModes` | `List<AuthMode>` | no |
| `publishAuthModes` | `List<AuthMode>` | no |
| `codeHandlers` | `string` | no |
| `tags` | `Map<string>` | no |
| `handlerConfigs` | `HandlerConfigs` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelNamespace` | `ChannelNamespace` | no |

## CreateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `serviceRoleArn` | `string` | no |
| `dynamodbConfig` | `DynamodbDataSourceConfig` | no |
| `lambdaConfig` | `LambdaDataSourceConfig` | no |
| `elasticsearchConfig` | `ElasticsearchDataSourceConfig` | no |
| `openSearchServiceConfig` | `OpenSearchServiceDataSourceConfig` | no |
| `httpConfig` | `HttpDataSourceConfig` | no |
| `relationalDatabaseConfig` | `RelationalDatabaseDataSourceConfig` | no |
| `eventBridgeConfig` | `EventBridgeDataSourceConfig` | no |
| `metricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | no |

## CreateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `certificateArn` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameConfig` | `DomainNameConfig` | no |

## CreateFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `dataSourceName` | `string` | yes |
| `requestMappingTemplate` | `string` | no |
| `responseMappingTemplate` | `string` | no |
| `functionVersion` | `string` | no |
| `syncConfig` | `SyncConfig` | no |
| `maxBatchSize` | `integer` | no |
| `runtime` | `AppSyncRuntime` | no |
| `code` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionConfiguration` | `FunctionConfiguration` | no |

## CreateGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `logConfig` | `LogConfig` | no |
| `authenticationType` | `string` | yes |
| `userPoolConfig` | `UserPoolConfig` | no |
| `openIDConnectConfig` | `OpenIDConnectConfig` | no |
| `tags` | `Map<string>` | no |
| `additionalAuthenticationProviders` | `List<AdditionalAuthenticationProvider>` | no |
| `xrayEnabled` | `boolean` | no |
| `lambdaAuthorizerConfig` | `LambdaAuthorizerConfig` | no |
| `apiType` | `string` | no |
| `mergedApiExecutionRoleArn` | `string` | no |
| `visibility` | `string` | no |
| `ownerContact` | `string` | no |
| `introspectionConfig` | `string` | no |
| `queryDepthLimit` | `integer` | no |
| `resolverCountLimit` | `integer` | no |
| `enhancedMetricsConfig` | `EnhancedMetricsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphqlApi` | `GraphqlApi` | no |

## CreateResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldName` | `string` | yes |
| `dataSourceName` | `string` | no |
| `requestMappingTemplate` | `string` | no |
| `responseMappingTemplate` | `string` | no |
| `kind` | `string` | no |
| `pipelineConfig` | `PipelineConfig` | no |
| `syncConfig` | `SyncConfig` | no |
| `cachingConfig` | `CachingConfig` | no |
| `maxBatchSize` | `integer` | no |
| `runtime` | `AppSyncRuntime` | no |
| `code` | `string` | no |
| `metricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolver` | `Resolver` | no |

## CreateType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `definition` | `string` | yes |
| `format` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `Type` | no |

## DeleteApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApiCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteChannelNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `functionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateMergedGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiIdentifier` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociationStatus` | `string` | no |

## DisassociateSourceGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergedApiIdentifier` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociationStatus` | `string` | no |

## EvaluateCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `runtime` | `AppSyncRuntime` | yes |
| `code` | `string` | yes |
| `context` | `string` | yes |
| `function` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluationResult` | `string` | no |
| `error` | `EvaluateCodeErrorDetail` | no |
| `logs` | `List<string>` | no |
| `stash` | `string` | no |
| `outErrors` | `string` | no |

## EvaluateMappingTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `template` | `string` | yes |
| `context` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `evaluationResult` | `string` | no |
| `error` | `ErrorDetail` | no |
| `logs` | `List<string>` | no |
| `stash` | `string` | no |
| `outErrors` | `string` | no |

## FlushApiCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `api` | `Api` | no |

## GetApiAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiAssociation` | `ApiAssociation` | no |

## GetApiCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiCache` | `ApiCache` | no |

## GetChannelNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelNamespace` | `ChannelNamespace` | no |

## GetDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | no |

## GetDataSourceIntrospection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `introspectionId` | `string` | yes |
| `includeModelsSDL` | `boolean` | no |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `introspectionId` | `string` | no |
| `introspectionStatus` | `string` | no |
| `introspectionStatusDetail` | `string` | no |
| `introspectionResult` | `DataSourceIntrospectionResult` | no |

## GetDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameConfig` | `DomainNameConfig` | no |

## GetFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `functionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionConfiguration` | `FunctionConfiguration` | no |

## GetGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphqlApi` | `GraphqlApi` | no |

## GetGraphqlApiEnvironmentVariables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentVariables` | `Map<string>` | no |

## GetIntrospectionSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `format` | `string` | yes |
| `includeDirectives` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `schema` | `blob` | no |

## GetResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolver` | `Resolver` | no |

## GetSchemaCreationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |
| `details` | `string` | no |

## GetSourceApiAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergedApiIdentifier` | `string` | yes |
| `associationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociation` | `SourceApiAssociation` | no |

## GetType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `format` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `Type` | no |

## ListApiKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKeys` | `List<ApiKey>` | no |
| `nextToken` | `string` | no |

## ListApis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apis` | `List<Api>` | no |
| `nextToken` | `string` | no |

## ListChannelNamespaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelNamespaces` | `List<ChannelNamespace>` | no |
| `nextToken` | `string` | no |

## ListDataSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSources` | `List<DataSource>` | no |
| `nextToken` | `string` | no |

## ListDomainNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameConfigs` | `List<DomainNameConfig>` | no |
| `nextToken` | `string` | no |

## ListFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functions` | `List<FunctionConfiguration>` | no |
| `nextToken` | `string` | no |

## ListGraphqlApis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |
| `apiType` | `string` | no |
| `owner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphqlApis` | `List<GraphqlApi>` | no |
| `nextToken` | `string` | no |

## ListResolvers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolvers` | `List<Resolver>` | no |
| `nextToken` | `string` | no |

## ListResolversByFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `functionId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolvers` | `List<Resolver>` | no |
| `nextToken` | `string` | no |

## ListSourceApiAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociationSummaries` | `List<SourceApiAssociationSummary>` | no |
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

## ListTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `format` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `types` | `List<Type>` | no |
| `nextToken` | `string` | no |

## ListTypesByAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `mergedApiIdentifier` | `string` | yes |
| `associationId` | `string` | yes |
| `format` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `types` | `List<Type>` | no |
| `nextToken` | `string` | no |

## PutGraphqlApiEnvironmentVariables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `environmentVariables` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `environmentVariables` | `Map<string>` | no |

## StartDataSourceIntrospection

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rdsDataApiConfig` | `RdsDataApiConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `introspectionId` | `string` | no |
| `introspectionStatus` | `string` | no |
| `introspectionStatusDetail` | `string` | no |

## StartSchemaCreation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `definition` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | no |

## StartSchemaMerge

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationId` | `string` | yes |
| `mergedApiIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociationStatus` | `string` | no |

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


## UpdateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `ownerContact` | `string` | no |
| `eventConfig` | `EventConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `api` | `Api` | no |

## UpdateApiCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `ttl` | `long` | yes |
| `apiCachingBehavior` | `string` | yes |
| `type` | `string` | yes |
| `healthMetricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiCache` | `ApiCache` | no |

## UpdateApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `id` | `string` | yes |
| `description` | `string` | no |
| `expires` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `ApiKey` | no |

## UpdateChannelNamespace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `subscribeAuthModes` | `List<AuthMode>` | no |
| `publishAuthModes` | `List<AuthMode>` | no |
| `codeHandlers` | `string` | no |
| `handlerConfigs` | `HandlerConfigs` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelNamespace` | `ChannelNamespace` | no |

## UpdateDataSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `type` | `string` | yes |
| `serviceRoleArn` | `string` | no |
| `dynamodbConfig` | `DynamodbDataSourceConfig` | no |
| `lambdaConfig` | `LambdaDataSourceConfig` | no |
| `elasticsearchConfig` | `ElasticsearchDataSourceConfig` | no |
| `openSearchServiceConfig` | `OpenSearchServiceDataSourceConfig` | no |
| `httpConfig` | `HttpDataSourceConfig` | no |
| `relationalDatabaseConfig` | `RelationalDatabaseDataSourceConfig` | no |
| `eventBridgeConfig` | `EventBridgeDataSourceConfig` | no |
| `metricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataSource` | `DataSource` | no |

## UpdateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameConfig` | `DomainNameConfig` | no |

## UpdateFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `functionId` | `string` | yes |
| `dataSourceName` | `string` | yes |
| `requestMappingTemplate` | `string` | no |
| `responseMappingTemplate` | `string` | no |
| `functionVersion` | `string` | no |
| `syncConfig` | `SyncConfig` | no |
| `maxBatchSize` | `integer` | no |
| `runtime` | `AppSyncRuntime` | no |
| `code` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `functionConfiguration` | `FunctionConfiguration` | no |

## UpdateGraphqlApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `name` | `string` | yes |
| `logConfig` | `LogConfig` | no |
| `authenticationType` | `string` | yes |
| `userPoolConfig` | `UserPoolConfig` | no |
| `openIDConnectConfig` | `OpenIDConnectConfig` | no |
| `additionalAuthenticationProviders` | `List<AdditionalAuthenticationProvider>` | no |
| `xrayEnabled` | `boolean` | no |
| `lambdaAuthorizerConfig` | `LambdaAuthorizerConfig` | no |
| `mergedApiExecutionRoleArn` | `string` | no |
| `ownerContact` | `string` | no |
| `introspectionConfig` | `string` | no |
| `queryDepthLimit` | `integer` | no |
| `resolverCountLimit` | `integer` | no |
| `enhancedMetricsConfig` | `EnhancedMetricsConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `graphqlApi` | `GraphqlApi` | no |

## UpdateResolver

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `fieldName` | `string` | yes |
| `dataSourceName` | `string` | no |
| `requestMappingTemplate` | `string` | no |
| `responseMappingTemplate` | `string` | no |
| `kind` | `string` | no |
| `pipelineConfig` | `PipelineConfig` | no |
| `syncConfig` | `SyncConfig` | no |
| `cachingConfig` | `CachingConfig` | no |
| `maxBatchSize` | `integer` | no |
| `runtime` | `AppSyncRuntime` | no |
| `code` | `string` | no |
| `metricsConfig` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resolver` | `Resolver` | no |

## UpdateSourceApiAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `associationId` | `string` | yes |
| `mergedApiIdentifier` | `string` | yes |
| `description` | `string` | no |
| `sourceApiAssociationConfig` | `SourceApiAssociationConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceApiAssociation` | `SourceApiAssociation` | no |

## UpdateType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiId` | `string` | yes |
| `typeName` | `string` | yes |
| `definition` | `string` | no |
| `format` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `Type` | no |

