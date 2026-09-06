# AWS Migration Hub Refactor Spaces

API version: 2021-10-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/migration-hub-refactor-spaces/2021-10-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayProxy` | `ApiGatewayProxyInput` | no |
| `ClientToken` | `string` | no |
| `EnvironmentIdentifier` | `string` | yes |
| `Name` | `string` | yes |
| `ProxyType` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `VpcId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayProxy` | `ApiGatewayProxyInput` | no |
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `EnvironmentId` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `ProxyType` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `VpcId` | `string` | no |

## CreateEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `NetworkFabricType` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `Description` | `string` | no |
| `EnvironmentId` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `NetworkFabricType` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `DefaultRoute` | `DefaultRouteInput` | no |
| `EnvironmentIdentifier` | `string` | yes |
| `RouteType` | `string` | yes |
| `ServiceIdentifier` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `UriPathRoute` | `UriPathRouteInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `OwnerAccountId` | `string` | no |
| `RouteId` | `string` | no |
| `RouteType` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UriPathRoute` | `UriPathRouteInput` | no |

## CreateService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `EndpointType` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `LambdaEndpoint` | `LambdaEndpointInput` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `UrlEndpoint` | `UrlEndpointInput` | no |
| `VpcId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `Description` | `string` | no |
| `EndpointType` | `string` | no |
| `EnvironmentId` | `string` | no |
| `LambdaEndpoint` | `LambdaEndpointInput` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UrlEndpoint` | `UrlEndpointInput` | no |
| `VpcId` | `string` | no |

## DeleteApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `EnvironmentId` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## DeleteEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `EnvironmentId` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `State` | `string` | no |

## DeleteResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `RouteIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `RouteId` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |

## DeleteService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `ServiceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `EnvironmentId` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |

## GetApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApiGatewayProxy` | `ApiGatewayProxyConfig` | no |
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `EnvironmentId` | `string` | no |
| `Error` | `ErrorResponse` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `ProxyType` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `VpcId` | `string` | no |

## GetEnvironment

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Arn` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `Description` | `string` | no |
| `EnvironmentId` | `string` | no |
| `Error` | `ErrorResponse` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `NetworkFabricType` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `TransitGatewayId` | `string` | no |

## GetResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | no |

## GetRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `RouteIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppendSourcePath` | `boolean` | no |
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `EnvironmentId` | `string` | no |
| `Error` | `ErrorResponse` | no |
| `IncludeChildPaths` | `boolean` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Methods` | `List<string>` | no |
| `OwnerAccountId` | `string` | no |
| `PathResourceToId` | `Map<string>` | no |
| `RouteId` | `string` | no |
| `RouteType` | `string` | no |
| `ServiceId` | `string` | no |
| `SourcePath` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |

## GetService

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `ServiceIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `CreatedByAccountId` | `string` | no |
| `CreatedTime` | `timestamp` | no |
| `Description` | `string` | no |
| `EndpointType` | `string` | no |
| `EnvironmentId` | `string` | no |
| `Error` | `ErrorResponse` | no |
| `LambdaEndpoint` | `LambdaEndpointConfig` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `Name` | `string` | no |
| `OwnerAccountId` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |
| `Tags` | `Map<string>` | no |
| `UrlEndpoint` | `UrlEndpointConfig` | no |
| `VpcId` | `string` | no |

## ListApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationSummaryList` | `List<ApplicationSummary>` | no |
| `NextToken` | `string` | no |

## ListEnvironmentVpcs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentVpcList` | `List<EnvironmentVpc>` | no |
| `NextToken` | `string` | no |

## ListEnvironments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EnvironmentSummaryList` | `List<EnvironmentSummary>` | no |
| `NextToken` | `string` | no |

## ListRoutes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RouteSummaryList` | `List<RouteSummary>` | no |

## ListServices

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServiceSummaryList` | `List<ServiceSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## PutResourcePolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policy` | `string` | yes |
| `ResourceArn` | `string` | yes |

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
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ActivationState` | `string` | yes |
| `ApplicationIdentifier` | `string` | yes |
| `EnvironmentIdentifier` | `string` | yes |
| `RouteIdentifier` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ApplicationId` | `string` | no |
| `Arn` | `string` | no |
| `LastUpdatedTime` | `timestamp` | no |
| `RouteId` | `string` | no |
| `ServiceId` | `string` | no |
| `State` | `string` | no |

