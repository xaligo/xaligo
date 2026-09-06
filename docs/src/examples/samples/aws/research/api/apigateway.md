# Amazon API Gateway

API version: 2015-07-09. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/apigateway/2015-07-09/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | no |
| `description` | `string` | no |
| `enabled` | `boolean` | no |
| `generateDistinctId` | `boolean` | no |
| `value` | `string` | no |
| `stageKeys` | `List<StageKey>` | no |
| `customerId` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `value` | `string` | no |
| `name` | `string` | no |
| `customerId` | `string` | no |
| `description` | `string` | no |
| `enabled` | `boolean` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `stageKeys` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## CreateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `name` | `string` | yes |
| `type` | `string` | yes |
| `providerARNs` | `List<string>` | no |
| `authType` | `string` | no |
| `authorizerUri` | `string` | no |
| `authorizerCredentials` | `string` | no |
| `identitySource` | `string` | no |
| `identityValidationExpression` | `string` | no |
| `authorizerResultTtlInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `providerARNs` | `List<string>` | no |
| `authType` | `string` | no |
| `authorizerUri` | `string` | no |
| `authorizerCredentials` | `string` | no |
| `identitySource` | `string` | no |
| `identityValidationExpression` | `string` | no |
| `authorizerResultTtlInSeconds` | `integer` | no |

## CreateBasePathMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `basePath` | `string` | no |
| `restApiId` | `string` | yes |
| `stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `basePath` | `string` | no |
| `restApiId` | `string` | no |
| `stage` | `string` | no |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | no |
| `stageDescription` | `string` | no |
| `description` | `string` | no |
| `cacheClusterEnabled` | `boolean` | no |
| `cacheClusterSize` | `string` | no |
| `variables` | `Map<string>` | no |
| `canarySettings` | `DeploymentCanarySettings` | no |
| `tracingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `apiSummary` | `Map<Map<MethodSnapshot>>` | no |

## CreateDocumentationPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `location` | `DocumentationPartLocation` | yes |
| `properties` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `location` | `DocumentationPartLocation` | no |
| `properties` | `string` | no |

## CreateDocumentationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationVersion` | `string` | yes |
| `stageName` | `string` | no |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | no |
| `createdDate` | `timestamp` | no |
| `description` | `string` | no |

## CreateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `certificateName` | `string` | no |
| `certificateBody` | `string` | no |
| `certificatePrivateKey` | `string` | no |
| `certificateChain` | `string` | no |
| `certificateArn` | `string` | no |
| `regionalCertificateName` | `string` | no |
| `regionalCertificateArn` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `tags` | `Map<string>` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `mutualTlsAuthentication` | `MutualTlsAuthenticationInput` | no |
| `ownershipVerificationCertificateArn` | `string` | no |
| `policy` | `string` | no |
| `routingMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `domainNameId` | `string` | no |
| `domainNameArn` | `string` | no |
| `certificateName` | `string` | no |
| `certificateArn` | `string` | no |
| `certificateUploadDate` | `timestamp` | no |
| `regionalDomainName` | `string` | no |
| `regionalHostedZoneId` | `string` | no |
| `regionalCertificateName` | `string` | no |
| `regionalCertificateArn` | `string` | no |
| `distributionDomainName` | `string` | no |
| `distributionHostedZoneId` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `domainNameStatus` | `string` | no |
| `domainNameStatusMessage` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `tags` | `Map<string>` | no |
| `mutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `ownershipVerificationCertificateArn` | `string` | no |
| `managementPolicy` | `string` | no |
| `policy` | `string` | no |
| `routingMode` | `string` | no |

## CreateDomainNameAccessAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameArn` | `string` | yes |
| `accessAssociationSourceType` | `string` | yes |
| `accessAssociationSource` | `string` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameAccessAssociationArn` | `string` | no |
| `domainNameArn` | `string` | no |
| `accessAssociationSourceType` | `string` | no |
| `accessAssociationSource` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | no |
| `schema` | `string` | no |
| `contentType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `schema` | `string` | no |
| `contentType` | `string` | no |

## CreateRequestValidator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `name` | `string` | no |
| `validateRequestBody` | `boolean` | no |
| `validateRequestParameters` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `validateRequestBody` | `boolean` | no |
| `validateRequestParameters` | `boolean` | no |

## CreateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `parentId` | `string` | yes |
| `pathPart` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `parentId` | `string` | no |
| `pathPart` | `string` | no |
| `path` | `string` | no |
| `resourceMethods` | `Map<Method>` | no |

## CreateRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `version` | `string` | no |
| `cloneFrom` | `string` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `warnings` | `List<string>` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `rootResourceId` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `apiStatus` | `string` | no |
| `apiStatusMessage` | `string` | no |

## CreateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |
| `deploymentId` | `string` | yes |
| `description` | `string` | no |
| `cacheClusterEnabled` | `boolean` | no |
| `cacheClusterSize` | `string` | no |
| `variables` | `Map<string>` | no |
| `documentationVersion` | `string` | no |
| `canarySettings` | `CanarySettings` | no |
| `tracingEnabled` | `boolean` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `clientCertificateId` | `string` | no |
| `stageName` | `string` | no |
| `description` | `string` | no |
| `cacheClusterEnabled` | `boolean` | no |
| `cacheClusterSize` | `string` | no |
| `cacheClusterStatus` | `string` | no |
| `methodSettings` | `Map<MethodSetting>` | no |
| `variables` | `Map<string>` | no |
| `documentationVersion` | `string` | no |
| `accessLogSettings` | `AccessLogSettings` | no |
| `canarySettings` | `CanarySettings` | no |
| `tracingEnabled` | `boolean` | no |
| `webAclArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |

## CreateUsagePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `apiStages` | `List<ApiStage>` | no |
| `throttle` | `ThrottleSettings` | no |
| `quota` | `QuotaSettings` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `apiStages` | `List<ApiStage>` | no |
| `throttle` | `ThrottleSettings` | no |
| `quota` | `QuotaSettings` | no |
| `productCode` | `string` | no |
| `tags` | `Map<string>` | no |

## CreateUsagePlanKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `keyId` | `string` | yes |
| `keyType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `type` | `string` | no |
| `value` | `string` | no |
| `name` | `string` | no |

## CreateVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | no |
| `targetArns` | `List<string>` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `targetArns` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## DeleteApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `authorizerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteBasePathMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `basePath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `deploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDocumentationPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationPartId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDocumentationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainNameAccessAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameAccessAssociationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteGatewayResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `responseType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMethodResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `modelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRequestValidator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `requestValidatorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUsagePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUsagePlanKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `keyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## FlushStageAuthorizersCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## FlushStageCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GenerateClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | no |
| `description` | `string` | no |
| `pemEncodedCertificate` | `string` | no |
| `createdDate` | `timestamp` | no |
| `expirationDate` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudwatchRoleArn` | `string` | no |
| `throttleSettings` | `ThrottleSettings` | no |
| `features` | `List<string>` | no |
| `apiKeyVersion` | `string` | no |

## GetApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `string` | yes |
| `includeValue` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `value` | `string` | no |
| `name` | `string` | no |
| `customerId` | `string` | no |
| `description` | `string` | no |
| `enabled` | `boolean` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `stageKeys` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## GetApiKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |
| `nameQuery` | `string` | no |
| `customerId` | `string` | no |
| `includeValues` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `warnings` | `List<string>` | no |
| `position` | `string` | no |
| `items` | `List<ApiKey>` | no |

## GetAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `authorizerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `providerARNs` | `List<string>` | no |
| `authType` | `string` | no |
| `authorizerUri` | `string` | no |
| `authorizerCredentials` | `string` | no |
| `identitySource` | `string` | no |
| `identityValidationExpression` | `string` | no |
| `authorizerResultTtlInSeconds` | `integer` | no |

## GetAuthorizers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<Authorizer>` | no |

## GetBasePathMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `basePath` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `basePath` | `string` | no |
| `restApiId` | `string` | no |
| `stage` | `string` | no |

## GetBasePathMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<BasePathMapping>` | no |

## GetClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | no |
| `description` | `string` | no |
| `pemEncodedCertificate` | `string` | no |
| `createdDate` | `timestamp` | no |
| `expirationDate` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## GetClientCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<ClientCertificate>` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `deploymentId` | `string` | yes |
| `embed` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `apiSummary` | `Map<Map<MethodSnapshot>>` | no |

## GetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<Deployment>` | no |

## GetDocumentationPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationPartId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `location` | `DocumentationPartLocation` | no |
| `properties` | `string` | no |

## GetDocumentationParts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `type` | `string` | no |
| `nameQuery` | `string` | no |
| `path` | `string` | no |
| `position` | `string` | no |
| `limit` | `integer` | no |
| `locationStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<DocumentationPart>` | no |

## GetDocumentationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | no |
| `createdDate` | `timestamp` | no |
| `description` | `string` | no |

## GetDocumentationVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<DocumentationVersion>` | no |

## GetDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `domainNameId` | `string` | no |
| `domainNameArn` | `string` | no |
| `certificateName` | `string` | no |
| `certificateArn` | `string` | no |
| `certificateUploadDate` | `timestamp` | no |
| `regionalDomainName` | `string` | no |
| `regionalHostedZoneId` | `string` | no |
| `regionalCertificateName` | `string` | no |
| `regionalCertificateArn` | `string` | no |
| `distributionDomainName` | `string` | no |
| `distributionHostedZoneId` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `domainNameStatus` | `string` | no |
| `domainNameStatusMessage` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `tags` | `Map<string>` | no |
| `mutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `ownershipVerificationCertificateArn` | `string` | no |
| `managementPolicy` | `string` | no |
| `policy` | `string` | no |
| `routingMode` | `string` | no |

## GetDomainNameAccessAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |
| `resourceOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<DomainNameAccessAssociation>` | no |

## GetDomainNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |
| `resourceOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<DomainName>` | no |

## GetExport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |
| `exportType` | `string` | yes |
| `parameters` | `Map<string>` | no |
| `accepts` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `contentDisposition` | `string` | no |
| `body` | `blob` | no |

## GetGatewayResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `responseType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseType` | `string` | no |
| `statusCode` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `defaultResponse` | `boolean` | no |

## GetGatewayResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<GatewayResponse>` | no |

## GetIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | no |
| `httpMethod` | `string` | no |
| `uri` | `string` | no |
| `connectionType` | `string` | no |
| `connectionId` | `string` | no |
| `credentials` | `string` | no |
| `requestParameters` | `Map<string>` | no |
| `requestTemplates` | `Map<string>` | no |
| `passthroughBehavior` | `string` | no |
| `contentHandling` | `string` | no |
| `timeoutInMillis` | `integer` | no |
| `cacheNamespace` | `string` | no |
| `cacheKeyParameters` | `List<string>` | no |
| `integrationResponses` | `Map<IntegrationResponse>` | no |
| `tlsConfig` | `TlsConfig` | no |
| `responseTransferMode` | `string` | no |
| `integrationTarget` | `string` | no |

## GetIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `selectionPattern` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `contentHandling` | `string` | no |

## GetMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `httpMethod` | `string` | no |
| `authorizationType` | `string` | no |
| `authorizerId` | `string` | no |
| `apiKeyRequired` | `boolean` | no |
| `requestValidatorId` | `string` | no |
| `operationName` | `string` | no |
| `requestParameters` | `Map<boolean>` | no |
| `requestModels` | `Map<string>` | no |
| `methodResponses` | `Map<MethodResponse>` | no |
| `methodIntegration` | `Integration` | no |
| `authorizationScopes` | `List<string>` | no |

## GetMethodResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `responseParameters` | `Map<boolean>` | no |
| `responseModels` | `Map<string>` | no |

## GetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `modelName` | `string` | yes |
| `flatten` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `schema` | `string` | no |
| `contentType` | `string` | no |

## GetModelTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `modelName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `value` | `string` | no |

## GetModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<Model>` | no |

## GetRequestValidator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `requestValidatorId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `validateRequestBody` | `boolean` | no |
| `validateRequestParameters` | `boolean` | no |

## GetRequestValidators

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<RequestValidator>` | no |

## GetResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `embed` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `parentId` | `string` | no |
| `pathPart` | `string` | no |
| `path` | `string` | no |
| `resourceMethods` | `Map<Method>` | no |

## GetResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |
| `embed` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<Resource>` | no |

## GetRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `warnings` | `List<string>` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `rootResourceId` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `apiStatus` | `string` | no |
| `apiStatusMessage` | `string` | no |

## GetRestApis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<RestApi>` | no |

## GetSdk

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |
| `sdkType` | `string` | yes |
| `parameters` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contentType` | `string` | no |
| `contentDisposition` | `string` | no |
| `body` | `blob` | no |

## GetSdkType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `friendlyName` | `string` | no |
| `description` | `string` | no |
| `configurationProperties` | `List<SdkConfigurationProperty>` | no |

## GetSdkTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<SdkType>` | no |

## GetStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `clientCertificateId` | `string` | no |
| `stageName` | `string` | no |
| `description` | `string` | no |
| `cacheClusterEnabled` | `boolean` | no |
| `cacheClusterSize` | `string` | no |
| `cacheClusterStatus` | `string` | no |
| `methodSettings` | `Map<MethodSetting>` | no |
| `variables` | `Map<string>` | no |
| `documentationVersion` | `string` | no |
| `accessLogSettings` | `AccessLogSettings` | no |
| `canarySettings` | `CanarySettings` | no |
| `tracingEnabled` | `boolean` | no |
| `webAclArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |

## GetStages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `deploymentId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `item` | `List<Stage>` | no |

## GetTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## GetUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `keyId` | `string` | no |
| `startDate` | `string` | yes |
| `endDate` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | no |
| `startDate` | `string` | no |
| `endDate` | `string` | no |
| `position` | `string` | no |
| `items` | `Map<List<List<long>>>` | no |

## GetUsagePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `apiStages` | `List<ApiStage>` | no |
| `throttle` | `ThrottleSettings` | no |
| `quota` | `QuotaSettings` | no |
| `productCode` | `string` | no |
| `tags` | `Map<string>` | no |

## GetUsagePlanKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `keyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `type` | `string` | no |
| `value` | `string` | no |
| `name` | `string` | no |

## GetUsagePlanKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `position` | `string` | no |
| `limit` | `integer` | no |
| `nameQuery` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<UsagePlanKey>` | no |

## GetUsagePlans

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `keyId` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<UsagePlan>` | no |

## GetVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `targetArns` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `tags` | `Map<string>` | no |

## GetVpcLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `limit` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `position` | `string` | no |
| `items` | `List<VpcLink>` | no |

## ImportApiKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | yes |
| `format` | `string` | yes |
| `failOnWarnings` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `warnings` | `List<string>` | no |

## ImportDocumentationParts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `mode` | `string` | no |
| `failOnWarnings` | `boolean` | no |
| `body` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ids` | `List<string>` | no |
| `warnings` | `List<string>` | no |

## ImportRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failOnWarnings` | `boolean` | no |
| `parameters` | `Map<string>` | no |
| `body` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `warnings` | `List<string>` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `rootResourceId` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `apiStatus` | `string` | no |
| `apiStatusMessage` | `string` | no |

## PutGatewayResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `responseType` | `string` | yes |
| `statusCode` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseType` | `string` | no |
| `statusCode` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `defaultResponse` | `boolean` | no |

## PutIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `type` | `string` | yes |
| `integrationHttpMethod` | `string` | no |
| `uri` | `string` | no |
| `connectionType` | `string` | no |
| `connectionId` | `string` | no |
| `credentials` | `string` | no |
| `requestParameters` | `Map<string>` | no |
| `requestTemplates` | `Map<string>` | no |
| `passthroughBehavior` | `string` | no |
| `cacheNamespace` | `string` | no |
| `cacheKeyParameters` | `List<string>` | no |
| `contentHandling` | `string` | no |
| `timeoutInMillis` | `integer` | no |
| `tlsConfig` | `TlsConfig` | no |
| `responseTransferMode` | `string` | no |
| `integrationTarget` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | no |
| `httpMethod` | `string` | no |
| `uri` | `string` | no |
| `connectionType` | `string` | no |
| `connectionId` | `string` | no |
| `credentials` | `string` | no |
| `requestParameters` | `Map<string>` | no |
| `requestTemplates` | `Map<string>` | no |
| `passthroughBehavior` | `string` | no |
| `contentHandling` | `string` | no |
| `timeoutInMillis` | `integer` | no |
| `cacheNamespace` | `string` | no |
| `cacheKeyParameters` | `List<string>` | no |
| `integrationResponses` | `Map<IntegrationResponse>` | no |
| `tlsConfig` | `TlsConfig` | no |
| `responseTransferMode` | `string` | no |
| `integrationTarget` | `string` | no |

## PutIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |
| `selectionPattern` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `contentHandling` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `selectionPattern` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `contentHandling` | `string` | no |

## PutMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `authorizationType` | `string` | yes |
| `authorizerId` | `string` | no |
| `apiKeyRequired` | `boolean` | no |
| `operationName` | `string` | no |
| `requestParameters` | `Map<boolean>` | no |
| `requestModels` | `Map<string>` | no |
| `requestValidatorId` | `string` | no |
| `authorizationScopes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `httpMethod` | `string` | no |
| `authorizationType` | `string` | no |
| `authorizerId` | `string` | no |
| `apiKeyRequired` | `boolean` | no |
| `requestValidatorId` | `string` | no |
| `operationName` | `string` | no |
| `requestParameters` | `Map<boolean>` | no |
| `requestModels` | `Map<string>` | no |
| `methodResponses` | `Map<MethodResponse>` | no |
| `methodIntegration` | `Integration` | no |
| `authorizationScopes` | `List<string>` | no |

## PutMethodResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |
| `responseParameters` | `Map<boolean>` | no |
| `responseModels` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `responseParameters` | `Map<boolean>` | no |
| `responseModels` | `Map<string>` | no |

## PutRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `mode` | `string` | no |
| `failOnWarnings` | `boolean` | no |
| `parameters` | `Map<string>` | no |
| `body` | `blob` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `warnings` | `List<string>` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `rootResourceId` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `apiStatus` | `string` | no |
| `apiStatusMessage` | `string` | no |

## RejectDomainNameAccessAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainNameAccessAssociationArn` | `string` | yes |
| `domainNameArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestInvokeAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `authorizerId` | `string` | yes |
| `headers` | `Map<string>` | no |
| `multiValueHeaders` | `Map<List<string>>` | no |
| `pathWithQueryString` | `string` | no |
| `body` | `string` | no |
| `stageVariables` | `Map<string>` | no |
| `additionalContext` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientStatus` | `integer` | no |
| `log` | `string` | no |
| `latency` | `long` | no |
| `principalId` | `string` | no |
| `policy` | `string` | no |
| `authorization` | `Map<List<string>>` | no |
| `claims` | `Map<string>` | no |

## TestInvokeMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `pathWithQueryString` | `string` | no |
| `body` | `string` | no |
| `headers` | `Map<string>` | no |
| `multiValueHeaders` | `Map<List<string>>` | no |
| `clientCertificateId` | `string` | no |
| `stageVariables` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `integer` | no |
| `body` | `string` | no |
| `headers` | `Map<string>` | no |
| `multiValueHeaders` | `Map<List<string>>` | no |
| `log` | `string` | no |
| `latency` | `long` | no |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `cloudwatchRoleArn` | `string` | no |
| `throttleSettings` | `ThrottleSettings` | no |
| `features` | `List<string>` | no |
| `apiKeyVersion` | `string` | no |

## UpdateApiKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `apiKey` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `value` | `string` | no |
| `name` | `string` | no |
| `customerId` | `string` | no |
| `description` | `string` | no |
| `enabled` | `boolean` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |
| `stageKeys` | `List<string>` | no |
| `tags` | `Map<string>` | no |

## UpdateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `authorizerId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `type` | `string` | no |
| `providerARNs` | `List<string>` | no |
| `authType` | `string` | no |
| `authorizerUri` | `string` | no |
| `authorizerCredentials` | `string` | no |
| `identitySource` | `string` | no |
| `identityValidationExpression` | `string` | no |
| `authorizerResultTtlInSeconds` | `integer` | no |

## UpdateBasePathMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `basePath` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `basePath` | `string` | no |
| `restApiId` | `string` | no |
| `stage` | `string` | no |

## UpdateClientCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientCertificateId` | `string` | no |
| `description` | `string` | no |
| `pemEncodedCertificate` | `string` | no |
| `createdDate` | `timestamp` | no |
| `expirationDate` | `timestamp` | no |
| `tags` | `Map<string>` | no |

## UpdateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `deploymentId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `apiSummary` | `Map<Map<MethodSnapshot>>` | no |

## UpdateDocumentationPart

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationPartId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `location` | `DocumentationPartLocation` | no |
| `properties` | `string` | no |

## UpdateDocumentationVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `documentationVersion` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `version` | `string` | no |
| `createdDate` | `timestamp` | no |
| `description` | `string` | no |

## UpdateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | yes |
| `domainNameId` | `string` | no |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `domainNameId` | `string` | no |
| `domainNameArn` | `string` | no |
| `certificateName` | `string` | no |
| `certificateArn` | `string` | no |
| `certificateUploadDate` | `timestamp` | no |
| `regionalDomainName` | `string` | no |
| `regionalHostedZoneId` | `string` | no |
| `regionalCertificateName` | `string` | no |
| `regionalCertificateArn` | `string` | no |
| `distributionDomainName` | `string` | no |
| `distributionHostedZoneId` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `domainNameStatus` | `string` | no |
| `domainNameStatusMessage` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `tags` | `Map<string>` | no |
| `mutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `ownershipVerificationCertificateArn` | `string` | no |
| `managementPolicy` | `string` | no |
| `policy` | `string` | no |
| `routingMode` | `string` | no |

## UpdateGatewayResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `responseType` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `responseType` | `string` | no |
| `statusCode` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `defaultResponse` | `boolean` | no |

## UpdateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `type` | `string` | no |
| `httpMethod` | `string` | no |
| `uri` | `string` | no |
| `connectionType` | `string` | no |
| `connectionId` | `string` | no |
| `credentials` | `string` | no |
| `requestParameters` | `Map<string>` | no |
| `requestTemplates` | `Map<string>` | no |
| `passthroughBehavior` | `string` | no |
| `contentHandling` | `string` | no |
| `timeoutInMillis` | `integer` | no |
| `cacheNamespace` | `string` | no |
| `cacheKeyParameters` | `List<string>` | no |
| `integrationResponses` | `Map<IntegrationResponse>` | no |
| `tlsConfig` | `TlsConfig` | no |
| `responseTransferMode` | `string` | no |
| `integrationTarget` | `string` | no |

## UpdateIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `selectionPattern` | `string` | no |
| `responseParameters` | `Map<string>` | no |
| `responseTemplates` | `Map<string>` | no |
| `contentHandling` | `string` | no |

## UpdateMethod

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `httpMethod` | `string` | no |
| `authorizationType` | `string` | no |
| `authorizerId` | `string` | no |
| `apiKeyRequired` | `boolean` | no |
| `requestValidatorId` | `string` | no |
| `operationName` | `string` | no |
| `requestParameters` | `Map<boolean>` | no |
| `requestModels` | `Map<string>` | no |
| `methodResponses` | `Map<MethodResponse>` | no |
| `methodIntegration` | `Integration` | no |
| `authorizationScopes` | `List<string>` | no |

## UpdateMethodResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `httpMethod` | `string` | yes |
| `statusCode` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `statusCode` | `string` | no |
| `responseParameters` | `Map<boolean>` | no |
| `responseModels` | `Map<string>` | no |

## UpdateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `modelName` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `schema` | `string` | no |
| `contentType` | `string` | no |

## UpdateRequestValidator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `requestValidatorId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `validateRequestBody` | `boolean` | no |
| `validateRequestParameters` | `boolean` | no |

## UpdateResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `resourceId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `parentId` | `string` | no |
| `pathPart` | `string` | no |
| `path` | `string` | no |
| `resourceMethods` | `Map<Method>` | no |

## UpdateRestApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `createdDate` | `timestamp` | no |
| `version` | `string` | no |
| `warnings` | `List<string>` | no |
| `binaryMediaTypes` | `List<string>` | no |
| `minimumCompressionSize` | `integer` | no |
| `apiKeySource` | `string` | no |
| `endpointConfiguration` | `EndpointConfiguration` | no |
| `policy` | `string` | no |
| `tags` | `Map<string>` | no |
| `disableExecuteApiEndpoint` | `boolean` | no |
| `rootResourceId` | `string` | no |
| `securityPolicy` | `string` | no |
| `endpointAccessMode` | `string` | no |
| `apiStatus` | `string` | no |
| `apiStatusMessage` | `string` | no |

## UpdateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `restApiId` | `string` | yes |
| `stageName` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `deploymentId` | `string` | no |
| `clientCertificateId` | `string` | no |
| `stageName` | `string` | no |
| `description` | `string` | no |
| `cacheClusterEnabled` | `boolean` | no |
| `cacheClusterSize` | `string` | no |
| `cacheClusterStatus` | `string` | no |
| `methodSettings` | `Map<MethodSetting>` | no |
| `variables` | `Map<string>` | no |
| `documentationVersion` | `string` | no |
| `accessLogSettings` | `AccessLogSettings` | no |
| `canarySettings` | `CanarySettings` | no |
| `tracingEnabled` | `boolean` | no |
| `webAclArn` | `string` | no |
| `tags` | `Map<string>` | no |
| `createdDate` | `timestamp` | no |
| `lastUpdatedDate` | `timestamp` | no |

## UpdateUsage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `keyId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | no |
| `startDate` | `string` | no |
| `endDate` | `string` | no |
| `position` | `string` | no |
| `items` | `Map<List<List<long>>>` | no |

## UpdateUsagePlan

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `usagePlanId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `apiStages` | `List<ApiStage>` | no |
| `throttle` | `ThrottleSettings` | no |
| `quota` | `QuotaSettings` | no |
| `productCode` | `string` | no |
| `tags` | `Map<string>` | no |

## UpdateVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcLinkId` | `string` | yes |
| `patchOperations` | `List<PatchOperation>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `id` | `string` | no |
| `name` | `string` | no |
| `description` | `string` | no |
| `targetArns` | `List<string>` | no |
| `status` | `string` | no |
| `statusMessage` | `string` | no |
| `tags` | `Map<string>` | no |

