# AmazonApiGatewayV2

API version: 2018-11-29. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/apigatewayv2/2018-11-29/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | yes |
| `ProtocolType` | `string` | yes |
| `RouteKey` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Target` | `string` | no |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiEndpoint` | `string` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `ApiId` | `string` | no |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CreatedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `ImportInfo` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `ProtocolType` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Version` | `string` | no |
| `Warnings` | `List<string>` | no |

## CreateApiMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ApiMappingKey` | `string` | no |
| `DomainName` | `string` | yes |
| `Stage` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | no |
| `ApiMappingId` | `string` | no |
| `ApiMappingKey` | `string` | no |
| `Stage` | `string` | no |

## CreateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `AuthorizerCredentialsArn` | `string` | no |
| `AuthorizerPayloadFormatVersion` | `string` | no |
| `AuthorizerResultTtlInSeconds` | `integer` | no |
| `AuthorizerType` | `string` | yes |
| `AuthorizerUri` | `string` | no |
| `EnableSimpleResponses` | `boolean` | no |
| `IdentitySource` | `List<string>` | yes |
| `IdentityValidationExpression` | `string` | no |
| `JwtConfiguration` | `JWTConfiguration` | no |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizerCredentialsArn` | `string` | no |
| `AuthorizerId` | `string` | no |
| `AuthorizerPayloadFormatVersion` | `string` | no |
| `AuthorizerResultTtlInSeconds` | `integer` | no |
| `AuthorizerType` | `string` | no |
| `AuthorizerUri` | `string` | no |
| `EnableSimpleResponses` | `boolean` | no |
| `IdentitySource` | `List<string>` | no |
| `IdentityValidationExpression` | `string` | no |
| `JwtConfiguration` | `JWTConfiguration` | no |
| `Name` | `string` | no |

## CreateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `Description` | `string` | no |
| `StageName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoDeployed` | `boolean` | no |
| `CreatedDate` | `timestamp` | no |
| `DeploymentId` | `string` | no |
| `DeploymentStatus` | `string` | no |
| `DeploymentStatusMessage` | `string` | no |
| `Description` | `string` | no |

## CreateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainNameConfigurations` | `List<DomainNameConfiguration>` | no |
| `MutualTlsAuthentication` | `MutualTlsAuthenticationInput` | no |
| `RoutingMode` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiMappingSelectionExpression` | `string` | no |
| `DomainName` | `string` | no |
| `DomainNameArn` | `string` | no |
| `DomainNameConfigurations` | `List<DomainNameConfiguration>` | no |
| `MutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `RoutingMode` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | no |
| `ContentHandlingStrategy` | `string` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `IntegrationMethod` | `string` | no |
| `IntegrationSubtype` | `string` | no |
| `IntegrationType` | `string` | yes |
| `IntegrationUri` | `string` | no |
| `PassthroughBehavior` | `string` | no |
| `PayloadFormatVersion` | `string` | no |
| `RequestParameters` | `Map<string>` | no |
| `RequestTemplates` | `Map<string>` | no |
| `ResponseParameters` | `Map<Map<string>>` | no |
| `TemplateSelectionExpression` | `string` | no |
| `TimeoutInMillis` | `integer` | no |
| `TlsConfig` | `TlsConfigInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | no |
| `ContentHandlingStrategy` | `string` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `IntegrationId` | `string` | no |
| `IntegrationMethod` | `string` | no |
| `IntegrationResponseSelectionExpression` | `string` | no |
| `IntegrationSubtype` | `string` | no |
| `IntegrationType` | `string` | no |
| `IntegrationUri` | `string` | no |
| `PassthroughBehavior` | `string` | no |
| `PayloadFormatVersion` | `string` | no |
| `RequestParameters` | `Map<string>` | no |
| `RequestTemplates` | `Map<string>` | no |
| `ResponseParameters` | `Map<Map<string>>` | no |
| `TemplateSelectionExpression` | `string` | no |
| `TimeoutInMillis` | `integer` | no |
| `TlsConfig` | `TlsConfig` | no |

## CreateIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ContentHandlingStrategy` | `string` | no |
| `IntegrationId` | `string` | yes |
| `IntegrationResponseKey` | `string` | yes |
| `ResponseParameters` | `Map<string>` | no |
| `ResponseTemplates` | `Map<string>` | no |
| `TemplateSelectionExpression` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentHandlingStrategy` | `string` | no |
| `IntegrationResponseId` | `string` | no |
| `IntegrationResponseKey` | `string` | no |
| `ResponseParameters` | `Map<string>` | no |
| `ResponseTemplates` | `Map<string>` | no |
| `TemplateSelectionExpression` | `string` | no |

## CreateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ContentType` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Schema` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Description` | `string` | no |
| `ModelId` | `string` | no |
| `Name` | `string` | no |
| `Schema` | `string` | no |

## CreatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | yes |
| `EndpointConfiguration` | `EndpointConfigurationRequest` | yes |
| `IncludedPortalProductArns` | `List<string>` | no |
| `LogoUri` | `string` | no |
| `PortalContent` | `PortalContent` | yes |
| `RumAppMonitorName` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `EndpointConfiguration` | `EndpointConfigurationResponse` | no |
| `IncludedPortalProductArns` | `List<string>` | no |
| `LastModified` | `timestamp` | no |
| `LastPublished` | `timestamp` | no |
| `LastPublishedDescription` | `string` | no |
| `PortalArn` | `string` | no |
| `PortalContent` | `PortalContent` | no |
| `PortalId` | `string` | no |
| `PublishStatus` | `string` | no |
| `RumAppMonitorName` | `string` | no |
| `StatusException` | `StatusException` | no |
| `Tags` | `Map<string>` | no |

## CreatePortalProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `DisplayOrder` | `DisplayOrder` | no |
| `LastModified` | `timestamp` | no |
| `PortalProductArn` | `string` | no |
| `PortalProductId` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateProductPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `DisplayContent` | yes |
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `DisplayContent` | no |
| `LastModified` | `timestamp` | no |
| `ProductPageArn` | `string` | no |
| `ProductPageId` | `string` | no |

## CreateProductRestEndpointPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `EndpointDisplayContent` | no |
| `PortalProductId` | `string` | yes |
| `RestEndpointIdentifier` | `RestEndpointIdentifier` | yes |
| `TryItState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `EndpointDisplayContentResponse` | no |
| `LastModified` | `timestamp` | no |
| `ProductRestEndpointPageArn` | `string` | no |
| `ProductRestEndpointPageId` | `string` | no |
| `RestEndpointIdentifier` | `RestEndpointIdentifier` | no |
| `Status` | `string` | no |
| `StatusException` | `StatusException` | no |
| `TryItState` | `string` | no |

## CreateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ApiKeyRequired` | `boolean` | no |
| `AuthorizationScopes` | `List<string>` | no |
| `AuthorizationType` | `string` | no |
| `AuthorizerId` | `string` | no |
| `ModelSelectionExpression` | `string` | no |
| `OperationName` | `string` | no |
| `RequestModels` | `Map<string>` | no |
| `RequestParameters` | `Map<ParameterConstraints>` | no |
| `RouteKey` | `string` | yes |
| `RouteResponseSelectionExpression` | `string` | no |
| `Target` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ApiKeyRequired` | `boolean` | no |
| `AuthorizationScopes` | `List<string>` | no |
| `AuthorizationType` | `string` | no |
| `AuthorizerId` | `string` | no |
| `ModelSelectionExpression` | `string` | no |
| `OperationName` | `string` | no |
| `RequestModels` | `Map<string>` | no |
| `RequestParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | no |
| `RouteKey` | `string` | no |
| `RouteResponseSelectionExpression` | `string` | no |
| `Target` | `string` | no |

## CreateRouteResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ModelSelectionExpression` | `string` | no |
| `ResponseModels` | `Map<string>` | no |
| `ResponseParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | yes |
| `RouteResponseKey` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelSelectionExpression` | `string` | no |
| `ResponseModels` | `Map<string>` | no |
| `ResponseParameters` | `Map<ParameterConstraints>` | no |
| `RouteResponseId` | `string` | no |
| `RouteResponseKey` | `string` | no |

## CreateRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<RoutingRuleAction>` | yes |
| `Conditions` | `List<RoutingRuleCondition>` | yes |
| `DomainName` | `string` | yes |
| `DomainNameId` | `string` | no |
| `Priority` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<RoutingRuleAction>` | no |
| `Conditions` | `List<RoutingRuleCondition>` | no |
| `Priority` | `integer` | no |
| `RoutingRuleArn` | `string` | no |
| `RoutingRuleId` | `string` | no |

## CreateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessLogSettings` | `AccessLogSettings` | no |
| `ApiId` | `string` | yes |
| `AutoDeploy` | `boolean` | no |
| `ClientCertificateId` | `string` | no |
| `DefaultRouteSettings` | `RouteSettings` | no |
| `DeploymentId` | `string` | no |
| `Description` | `string` | no |
| `RouteSettings` | `Map<RouteSettings>` | no |
| `StageName` | `string` | yes |
| `StageVariables` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessLogSettings` | `AccessLogSettings` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `AutoDeploy` | `boolean` | no |
| `ClientCertificateId` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `DefaultRouteSettings` | `RouteSettings` | no |
| `DeploymentId` | `string` | no |
| `Description` | `string` | no |
| `LastDeploymentStatusMessage` | `string` | no |
| `LastUpdatedDate` | `timestamp` | no |
| `RouteSettings` | `Map<RouteSettings>` | no |
| `StageName` | `string` | no |
| `StageVariables` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |

## CreateVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedDate` | `timestamp` | no |
| `Name` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `VpcLinkId` | `string` | no |
| `VpcLinkStatus` | `string` | no |
| `VpcLinkStatusMessage` | `string` | no |
| `VpcLinkVersion` | `string` | no |

## DeleteAccessLogSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteApiMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiMappingId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `AuthorizerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCorsConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `DeploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `IntegrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `IntegrationId` | `string` | yes |
| `IntegrationResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortalProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePortalProductSharingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProductPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |
| `ProductPageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteProductRestEndpointPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |
| `ProductRestEndpointPageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RouteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRouteRequestParameter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RequestParameterKey` | `string` | yes |
| `RouteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRouteResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RouteId` | `string` | yes |
| `RouteResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRouteSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RouteKey` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainNameId` | `string` | no |
| `RoutingRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ExportApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ExportVersion` | `string` | no |
| `IncludeExtensions` | `boolean` | no |
| `OutputType` | `string` | yes |
| `Specification` | `string` | yes |
| `StageName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `body` | `blob` | no |

## DisablePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResetAuthorizersCache

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiEndpoint` | `string` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `ApiId` | `string` | no |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CreatedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `ImportInfo` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `ProtocolType` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Version` | `string` | no |
| `Warnings` | `List<string>` | no |

## GetApiMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiMappingId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | no |
| `ApiMappingId` | `string` | no |
| `ApiMappingKey` | `string` | no |
| `Stage` | `string` | no |

## GetApiMappings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ApiMapping>` | no |
| `NextToken` | `string` | no |

## GetApis

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Api>` | no |
| `NextToken` | `string` | no |

## GetAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `AuthorizerId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizerCredentialsArn` | `string` | no |
| `AuthorizerId` | `string` | no |
| `AuthorizerPayloadFormatVersion` | `string` | no |
| `AuthorizerResultTtlInSeconds` | `integer` | no |
| `AuthorizerType` | `string` | no |
| `AuthorizerUri` | `string` | no |
| `EnableSimpleResponses` | `boolean` | no |
| `IdentitySource` | `List<string>` | no |
| `IdentityValidationExpression` | `string` | no |
| `JwtConfiguration` | `JWTConfiguration` | no |
| `Name` | `string` | no |

## GetAuthorizers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Authorizer>` | no |
| `NextToken` | `string` | no |

## GetDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `DeploymentId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoDeployed` | `boolean` | no |
| `CreatedDate` | `timestamp` | no |
| `DeploymentId` | `string` | no |
| `DeploymentStatus` | `string` | no |
| `DeploymentStatusMessage` | `string` | no |
| `Description` | `string` | no |

## GetDeployments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Deployment>` | no |
| `NextToken` | `string` | no |

## GetDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiMappingSelectionExpression` | `string` | no |
| `DomainName` | `string` | no |
| `DomainNameArn` | `string` | no |
| `DomainNameConfigurations` | `List<DomainNameConfiguration>` | no |
| `MutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `RoutingMode` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetDomainNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DomainName>` | no |
| `NextToken` | `string` | no |

## GetIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `IntegrationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | no |
| `ContentHandlingStrategy` | `string` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `IntegrationId` | `string` | no |
| `IntegrationMethod` | `string` | no |
| `IntegrationResponseSelectionExpression` | `string` | no |
| `IntegrationSubtype` | `string` | no |
| `IntegrationType` | `string` | no |
| `IntegrationUri` | `string` | no |
| `PassthroughBehavior` | `string` | no |
| `PayloadFormatVersion` | `string` | no |
| `RequestParameters` | `Map<string>` | no |
| `RequestTemplates` | `Map<string>` | no |
| `ResponseParameters` | `Map<Map<string>>` | no |
| `TemplateSelectionExpression` | `string` | no |
| `TimeoutInMillis` | `integer` | no |
| `TlsConfig` | `TlsConfig` | no |

## GetIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `IntegrationId` | `string` | yes |
| `IntegrationResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentHandlingStrategy` | `string` | no |
| `IntegrationResponseId` | `string` | no |
| `IntegrationResponseKey` | `string` | no |
| `ResponseParameters` | `Map<string>` | no |
| `ResponseTemplates` | `Map<string>` | no |
| `TemplateSelectionExpression` | `string` | no |

## GetIntegrationResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `IntegrationId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<IntegrationResponse>` | no |
| `NextToken` | `string` | no |

## GetIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Integration>` | no |
| `NextToken` | `string` | no |

## GetModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Description` | `string` | no |
| `ModelId` | `string` | no |
| `Name` | `string` | no |
| `Schema` | `string` | no |

## GetModelTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ModelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Value` | `string` | no |

## GetModels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Model>` | no |
| `NextToken` | `string` | no |

## GetPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `EndpointConfiguration` | `EndpointConfigurationResponse` | no |
| `IncludedPortalProductArns` | `List<string>` | no |
| `LastModified` | `timestamp` | no |
| `LastPublished` | `timestamp` | no |
| `LastPublishedDescription` | `string` | no |
| `PortalArn` | `string` | no |
| `PortalContent` | `PortalContent` | no |
| `PortalId` | `string` | no |
| `Preview` | `Preview` | no |
| `PublishStatus` | `string` | no |
| `RumAppMonitorName` | `string` | no |
| `StatusException` | `StatusException` | no |
| `Tags` | `Map<string>` | no |

## GetPortalProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |
| `ResourceOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `DisplayOrder` | `DisplayOrder` | no |
| `LastModified` | `timestamp` | no |
| `PortalProductArn` | `string` | no |
| `PortalProductId` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetPortalProductSharingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | no |
| `PortalProductId` | `string` | no |

## GetProductPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalProductId` | `string` | yes |
| `ProductPageId` | `string` | yes |
| `ResourceOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `DisplayContent` | no |
| `LastModified` | `timestamp` | no |
| `ProductPageArn` | `string` | no |
| `ProductPageId` | `string` | no |

## GetProductRestEndpointPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IncludeRawDisplayContent` | `string` | no |
| `PortalProductId` | `string` | yes |
| `ProductRestEndpointPageId` | `string` | yes |
| `ResourceOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `EndpointDisplayContentResponse` | no |
| `LastModified` | `timestamp` | no |
| `ProductRestEndpointPageArn` | `string` | no |
| `ProductRestEndpointPageId` | `string` | no |
| `RawDisplayContent` | `string` | no |
| `RestEndpointIdentifier` | `RestEndpointIdentifier` | no |
| `Status` | `string` | no |
| `StatusException` | `StatusException` | no |
| `TryItState` | `string` | no |

## GetRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RouteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ApiKeyRequired` | `boolean` | no |
| `AuthorizationScopes` | `List<string>` | no |
| `AuthorizationType` | `string` | no |
| `AuthorizerId` | `string` | no |
| `ModelSelectionExpression` | `string` | no |
| `OperationName` | `string` | no |
| `RequestModels` | `Map<string>` | no |
| `RequestParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | no |
| `RouteKey` | `string` | no |
| `RouteResponseSelectionExpression` | `string` | no |
| `Target` | `string` | no |

## GetRouteResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `RouteId` | `string` | yes |
| `RouteResponseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelSelectionExpression` | `string` | no |
| `ResponseModels` | `Map<string>` | no |
| `ResponseParameters` | `Map<ParameterConstraints>` | no |
| `RouteResponseId` | `string` | no |
| `RouteResponseKey` | `string` | no |

## GetRouteResponses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `RouteId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<RouteResponse>` | no |
| `NextToken` | `string` | no |

## GetRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Route>` | no |
| `NextToken` | `string` | no |

## GetRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainNameId` | `string` | no |
| `RoutingRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<RoutingRuleAction>` | no |
| `Conditions` | `List<RoutingRuleCondition>` | no |
| `Priority` | `integer` | no |
| `RoutingRuleArn` | `string` | no |
| `RoutingRuleId` | `string` | no |

## ListRoutingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainNameId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RoutingRules` | `List<RoutingRule>` | no |

## GetStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `StageName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessLogSettings` | `AccessLogSettings` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `AutoDeploy` | `boolean` | no |
| `ClientCertificateId` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `DefaultRouteSettings` | `RouteSettings` | no |
| `DeploymentId` | `string` | no |
| `Description` | `string` | no |
| `LastDeploymentStatusMessage` | `string` | no |
| `LastUpdatedDate` | `timestamp` | no |
| `RouteSettings` | `Map<RouteSettings>` | no |
| `StageName` | `string` | no |
| `StageVariables` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |

## GetStages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Stage>` | no |
| `NextToken` | `string` | no |

## GetTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## GetVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VpcLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedDate` | `timestamp` | no |
| `Name` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `VpcLinkId` | `string` | no |
| `VpcLinkStatus` | `string` | no |
| `VpcLinkStatusMessage` | `string` | no |
| `VpcLinkVersion` | `string` | no |

## GetVpcLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<VpcLink>` | no |
| `NextToken` | `string` | no |

## ImportApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Basepath` | `string` | no |
| `Body` | `string` | yes |
| `FailOnWarnings` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiEndpoint` | `string` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `ApiId` | `string` | no |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CreatedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `ImportInfo` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `ProtocolType` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Version` | `string` | no |
| `Warnings` | `List<string>` | no |

## ListPortalProducts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `ResourceOwner` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<PortalProductSummary>` | no |
| `NextToken` | `string` | no |

## ListPortals

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<PortalSummary>` | no |
| `NextToken` | `string` | no |

## ListProductPages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `PortalProductId` | `string` | yes |
| `ResourceOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ProductPageSummaryNoBody>` | no |
| `NextToken` | `string` | no |

## ListProductRestEndpointPages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `string` | no |
| `NextToken` | `string` | no |
| `PortalProductId` | `string` | yes |
| `ResourceOwnerAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ProductRestEndpointPageSummaryNoBody>` | no |
| `NextToken` | `string` | no |

## PreviewPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PortalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PublishPortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `PortalId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutPortalProductSharingPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyDocument` | `string` | yes |
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<RoutingRuleAction>` | yes |
| `Conditions` | `List<RoutingRuleCondition>` | yes |
| `DomainName` | `string` | yes |
| `DomainNameId` | `string` | no |
| `Priority` | `integer` | yes |
| `RoutingRuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Actions` | `List<RoutingRuleAction>` | no |
| `Conditions` | `List<RoutingRuleCondition>` | no |
| `Priority` | `integer` | no |
| `RoutingRuleArn` | `string` | no |
| `RoutingRuleId` | `string` | no |

## ReimportApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `Basepath` | `string` | no |
| `Body` | `string` | yes |
| `FailOnWarnings` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiEndpoint` | `string` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `ApiId` | `string` | no |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CreatedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `ImportInfo` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `ProtocolType` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Version` | `string` | no |
| `Warnings` | `List<string>` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateApi

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `RouteKey` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Target` | `string` | no |
| `Version` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiEndpoint` | `string` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `ApiId` | `string` | no |
| `ApiKeySelectionExpression` | `string` | no |
| `CorsConfiguration` | `Cors` | no |
| `CreatedDate` | `timestamp` | no |
| `Description` | `string` | no |
| `DisableSchemaValidation` | `boolean` | no |
| `DisableExecuteApiEndpoint` | `boolean` | no |
| `ImportInfo` | `List<string>` | no |
| `IpAddressType` | `string` | no |
| `Name` | `string` | no |
| `ProtocolType` | `string` | no |
| `RouteSelectionExpression` | `string` | no |
| `Tags` | `Map<string>` | no |
| `Version` | `string` | no |
| `Warnings` | `List<string>` | no |

## UpdateApiMapping

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ApiMappingId` | `string` | yes |
| `ApiMappingKey` | `string` | no |
| `DomainName` | `string` | yes |
| `Stage` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | no |
| `ApiMappingId` | `string` | no |
| `ApiMappingKey` | `string` | no |
| `Stage` | `string` | no |

## UpdateAuthorizer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `AuthorizerCredentialsArn` | `string` | no |
| `AuthorizerId` | `string` | yes |
| `AuthorizerPayloadFormatVersion` | `string` | no |
| `AuthorizerResultTtlInSeconds` | `integer` | no |
| `AuthorizerType` | `string` | no |
| `AuthorizerUri` | `string` | no |
| `EnableSimpleResponses` | `boolean` | no |
| `IdentitySource` | `List<string>` | no |
| `IdentityValidationExpression` | `string` | no |
| `JwtConfiguration` | `JWTConfiguration` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthorizerCredentialsArn` | `string` | no |
| `AuthorizerId` | `string` | no |
| `AuthorizerPayloadFormatVersion` | `string` | no |
| `AuthorizerResultTtlInSeconds` | `integer` | no |
| `AuthorizerType` | `string` | no |
| `AuthorizerUri` | `string` | no |
| `EnableSimpleResponses` | `boolean` | no |
| `IdentitySource` | `List<string>` | no |
| `IdentityValidationExpression` | `string` | no |
| `JwtConfiguration` | `JWTConfiguration` | no |
| `Name` | `string` | no |

## UpdateDeployment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `DeploymentId` | `string` | yes |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoDeployed` | `boolean` | no |
| `CreatedDate` | `timestamp` | no |
| `DeploymentId` | `string` | no |
| `DeploymentStatus` | `string` | no |
| `DeploymentStatusMessage` | `string` | no |
| `Description` | `string` | no |

## UpdateDomainName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DomainNameConfigurations` | `List<DomainNameConfiguration>` | no |
| `MutualTlsAuthentication` | `MutualTlsAuthenticationInput` | no |
| `RoutingMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiMappingSelectionExpression` | `string` | no |
| `DomainName` | `string` | no |
| `DomainNameArn` | `string` | no |
| `DomainNameConfigurations` | `List<DomainNameConfiguration>` | no |
| `MutualTlsAuthentication` | `MutualTlsAuthentication` | no |
| `RoutingMode` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | no |
| `ContentHandlingStrategy` | `string` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `IntegrationId` | `string` | yes |
| `IntegrationMethod` | `string` | no |
| `IntegrationSubtype` | `string` | no |
| `IntegrationType` | `string` | no |
| `IntegrationUri` | `string` | no |
| `PassthroughBehavior` | `string` | no |
| `PayloadFormatVersion` | `string` | no |
| `RequestParameters` | `Map<string>` | no |
| `RequestTemplates` | `Map<string>` | no |
| `ResponseParameters` | `Map<Map<string>>` | no |
| `TemplateSelectionExpression` | `string` | no |
| `TimeoutInMillis` | `integer` | no |
| `TlsConfig` | `TlsConfigInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ConnectionId` | `string` | no |
| `ConnectionType` | `string` | no |
| `ContentHandlingStrategy` | `string` | no |
| `CredentialsArn` | `string` | no |
| `Description` | `string` | no |
| `IntegrationId` | `string` | no |
| `IntegrationMethod` | `string` | no |
| `IntegrationResponseSelectionExpression` | `string` | no |
| `IntegrationSubtype` | `string` | no |
| `IntegrationType` | `string` | no |
| `IntegrationUri` | `string` | no |
| `PassthroughBehavior` | `string` | no |
| `PayloadFormatVersion` | `string` | no |
| `RequestParameters` | `Map<string>` | no |
| `RequestTemplates` | `Map<string>` | no |
| `ResponseParameters` | `Map<Map<string>>` | no |
| `TemplateSelectionExpression` | `string` | no |
| `TimeoutInMillis` | `integer` | no |
| `TlsConfig` | `TlsConfig` | no |

## UpdateIntegrationResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ContentHandlingStrategy` | `string` | no |
| `IntegrationId` | `string` | yes |
| `IntegrationResponseId` | `string` | yes |
| `IntegrationResponseKey` | `string` | no |
| `ResponseParameters` | `Map<string>` | no |
| `ResponseTemplates` | `Map<string>` | no |
| `TemplateSelectionExpression` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentHandlingStrategy` | `string` | no |
| `IntegrationResponseId` | `string` | no |
| `IntegrationResponseKey` | `string` | no |
| `ResponseParameters` | `Map<string>` | no |
| `ResponseTemplates` | `Map<string>` | no |
| `TemplateSelectionExpression` | `string` | no |

## UpdateModel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ContentType` | `string` | no |
| `Description` | `string` | no |
| `ModelId` | `string` | yes |
| `Name` | `string` | no |
| `Schema` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContentType` | `string` | no |
| `Description` | `string` | no |
| `ModelId` | `string` | no |
| `Name` | `string` | no |
| `Schema` | `string` | no |

## UpdatePortal

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `EndpointConfiguration` | `EndpointConfigurationRequest` | no |
| `IncludedPortalProductArns` | `List<string>` | no |
| `LogoUri` | `string` | no |
| `PortalContent` | `PortalContent` | no |
| `PortalId` | `string` | yes |
| `RumAppMonitorName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Authorization` | `Authorization` | no |
| `EndpointConfiguration` | `EndpointConfigurationResponse` | no |
| `IncludedPortalProductArns` | `List<string>` | no |
| `LastModified` | `timestamp` | no |
| `LastPublished` | `timestamp` | no |
| `LastPublishedDescription` | `string` | no |
| `PortalArn` | `string` | no |
| `PortalContent` | `PortalContent` | no |
| `PortalId` | `string` | no |
| `Preview` | `Preview` | no |
| `PublishStatus` | `string` | no |
| `RumAppMonitorName` | `string` | no |
| `StatusException` | `StatusException` | no |
| `Tags` | `Map<string>` | no |

## UpdatePortalProduct

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `DisplayOrder` | `DisplayOrder` | no |
| `PortalProductId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `DisplayOrder` | `DisplayOrder` | no |
| `LastModified` | `timestamp` | no |
| `PortalProductArn` | `string` | no |
| `PortalProductId` | `string` | no |
| `Tags` | `Map<string>` | no |

## UpdateProductPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `DisplayContent` | no |
| `PortalProductId` | `string` | yes |
| `ProductPageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `DisplayContent` | no |
| `LastModified` | `timestamp` | no |
| `ProductPageArn` | `string` | no |
| `ProductPageId` | `string` | no |

## UpdateProductRestEndpointPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `EndpointDisplayContent` | no |
| `PortalProductId` | `string` | yes |
| `ProductRestEndpointPageId` | `string` | yes |
| `TryItState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DisplayContent` | `EndpointDisplayContentResponse` | no |
| `LastModified` | `timestamp` | no |
| `ProductRestEndpointPageArn` | `string` | no |
| `ProductRestEndpointPageId` | `string` | no |
| `RestEndpointIdentifier` | `RestEndpointIdentifier` | no |
| `Status` | `string` | no |
| `StatusException` | `StatusException` | no |
| `TryItState` | `string` | no |

## UpdateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ApiKeyRequired` | `boolean` | no |
| `AuthorizationScopes` | `List<string>` | no |
| `AuthorizationType` | `string` | no |
| `AuthorizerId` | `string` | no |
| `ModelSelectionExpression` | `string` | no |
| `OperationName` | `string` | no |
| `RequestModels` | `Map<string>` | no |
| `RequestParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | yes |
| `RouteKey` | `string` | no |
| `RouteResponseSelectionExpression` | `string` | no |
| `Target` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayManaged` | `boolean` | no |
| `ApiKeyRequired` | `boolean` | no |
| `AuthorizationScopes` | `List<string>` | no |
| `AuthorizationType` | `string` | no |
| `AuthorizerId` | `string` | no |
| `ModelSelectionExpression` | `string` | no |
| `OperationName` | `string` | no |
| `RequestModels` | `Map<string>` | no |
| `RequestParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | no |
| `RouteKey` | `string` | no |
| `RouteResponseSelectionExpression` | `string` | no |
| `Target` | `string` | no |

## UpdateRouteResponse

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiId` | `string` | yes |
| `ModelSelectionExpression` | `string` | no |
| `ResponseModels` | `Map<string>` | no |
| `ResponseParameters` | `Map<ParameterConstraints>` | no |
| `RouteId` | `string` | yes |
| `RouteResponseId` | `string` | yes |
| `RouteResponseKey` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ModelSelectionExpression` | `string` | no |
| `ResponseModels` | `Map<string>` | no |
| `ResponseParameters` | `Map<ParameterConstraints>` | no |
| `RouteResponseId` | `string` | no |
| `RouteResponseKey` | `string` | no |

## UpdateStage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessLogSettings` | `AccessLogSettings` | no |
| `ApiId` | `string` | yes |
| `AutoDeploy` | `boolean` | no |
| `ClientCertificateId` | `string` | no |
| `DefaultRouteSettings` | `RouteSettings` | no |
| `DeploymentId` | `string` | no |
| `Description` | `string` | no |
| `RouteSettings` | `Map<RouteSettings>` | no |
| `StageName` | `string` | yes |
| `StageVariables` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessLogSettings` | `AccessLogSettings` | no |
| `ApiGatewayManaged` | `boolean` | no |
| `AutoDeploy` | `boolean` | no |
| `ClientCertificateId` | `string` | no |
| `CreatedDate` | `timestamp` | no |
| `DefaultRouteSettings` | `RouteSettings` | no |
| `DeploymentId` | `string` | no |
| `Description` | `string` | no |
| `LastDeploymentStatusMessage` | `string` | no |
| `LastUpdatedDate` | `timestamp` | no |
| `RouteSettings` | `Map<RouteSettings>` | no |
| `StageName` | `string` | no |
| `StageVariables` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |

## UpdateVpcLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `VpcLinkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CreatedDate` | `timestamp` | no |
| `Name` | `string` | no |
| `SecurityGroupIds` | `List<string>` | no |
| `SubnetIds` | `List<string>` | no |
| `Tags` | `Map<string>` | no |
| `VpcLinkId` | `string` | no |
| `VpcLinkStatus` | `string` | no |
| `VpcLinkStatusMessage` | `string` | no |
| `VpcLinkVersion` | `string` | no |

