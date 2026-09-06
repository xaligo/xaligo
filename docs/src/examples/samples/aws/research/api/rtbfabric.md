# RTBFabric

API version: 2023-05-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/rtbfabric/2023-05-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | yes |
| `timeoutInMillis` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `peerGatewayId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `direction` | `string` | no |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |
| `linkId` | `string` | yes |

## AssociateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |
| `clientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |
| `status` | `string` | yes |

## CreateInboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `status` | `string` | yes |
| `domainName` | `string` | yes |

## CreateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `peerGatewayId` | `string` | yes |
| `attributes` | `LinkAttributes` | no |
| `httpResponderAllowed` | `boolean` | no |
| `tags` | `Map<string>` | no |
| `logSettings` | `LinkLogSettings` | yes |
| `timeoutInMillis` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `peerGatewayId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `direction` | `string` | no |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |
| `linkId` | `string` | yes |
| `customerProvidedId` | `string` | no |

## CreateLinkRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `RuleCondition` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |

## CreateOutboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `attributes` | `LinkAttributes` | no |
| `publicEndpoint` | `string` | yes |
| `logSettings` | `LinkLogSettings` | yes |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## CreateRequesterGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `clientToken` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `domainName` | `string` | yes |
| `status` | `string` | yes |

## CreateResponderGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `domainName` | `string` | no |
| `port` | `integer` | yes |
| `protocol` | `string` | yes |
| `listenerConfig` | `ListenerConfig` | no |
| `trustStoreConfiguration` | `TrustStoreConfiguration` | no |
| `managedEndpointConfiguration` | `ManagedEndpointConfiguration` | no |
| `clientToken` | `string` | yes |
| `description` | `string` | no |
| `tags` | `Map<string>` | no |
| `gatewayType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |
| `listenerConfig` | `ListenerConfig` | no |
| `externalInboundEndpoint` | `string` | no |

## DeleteInboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## DeleteLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## DeleteLinkRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `ruleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `status` | `string` | yes |

## DeleteOutboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## DeleteRequesterGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |

## DeleteResponderGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |

## DisassociateCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |
| `status` | `string` | yes |

## GetCertificateAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `acmCertificateArn` | `string` | yes |
| `status` | `string` | yes |
| `associatedAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |

## GetInboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `status` | `string` | yes |
| `domainName` | `string` | yes |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |

## GetLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `peerGatewayId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `direction` | `string` | no |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |
| `linkId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `httpResponderAllowed` | `boolean` | no |
| `timeoutInMillis` | `long` | no |

## GetLinkRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `ruleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `ruleId` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `RuleCondition` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `tags` | `Map<string>` | no |

## GetOutboundExternalLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `status` | `string` | yes |
| `publicEndpoint` | `string` | yes |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `tags` | `Map<string>` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |

## GetRequesterGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `status` | `string` | yes |
| `domainName` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `gatewayId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `activeLinksCount` | `integer` | no |
| `totalLinksCount` | `integer` | no |

## GetResponderGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `vpcId` | `string` | yes |
| `subnetIds` | `List<string>` | yes |
| `securityGroupIds` | `List<string>` | yes |
| `status` | `string` | yes |
| `description` | `string` | no |
| `createdAt` | `timestamp` | no |
| `updatedAt` | `timestamp` | no |
| `domainName` | `string` | no |
| `port` | `integer` | yes |
| `protocol` | `string` | yes |
| `listenerConfig` | `ListenerConfig` | no |
| `trustStoreConfiguration` | `TrustStoreConfiguration` | no |
| `managedEndpointConfiguration` | `ManagedEndpointConfiguration` | no |
| `gatewayId` | `string` | yes |
| `tags` | `Map<string>` | no |
| `activeLinksCount` | `integer` | no |
| `totalLinksCount` | `integer` | no |
| `linksRequestedCount` | `integer` | no |
| `gatewayType` | `string` | no |
| `externalInboundEndpoint` | `string` | no |

## ListCertificateAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `certificateAssociations` | `List<CertificateAssociationSummary>` | yes |
| `nextToken` | `string` | no |

## ListLinkRoutingRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `rules` | `List<LinkRoutingRuleSummary>` | no |
| `nextToken` | `string` | no |

## ListLinks

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `nextToken` | `string` | no |
| `maxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `links` | `List<ListLinksResponseStructure>` | no |
| `nextToken` | `string` | no |

## ListRequesterGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIds` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListResponderGateways

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayIds` | `List<string>` | no |
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

## RejectLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `peerGatewayId` | `string` | yes |
| `status` | `string` | yes |
| `createdAt` | `timestamp` | yes |
| `updatedAt` | `timestamp` | yes |
| `direction` | `string` | no |
| `flowModules` | `List<ModuleConfiguration>` | no |
| `pendingFlowModules` | `List<ModuleConfiguration>` | no |
| `attributes` | `LinkAttributes` | no |
| `logSettings` | `LinkLogSettings` | no |
| `connectivityType` | `string` | no |
| `linkId` | `string` | yes |

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


## UpdateLink

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `logSettings` | `LinkLogSettings` | no |
| `timeoutInMillis` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## UpdateLinkModuleFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `modules` | `List<ModuleConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `status` | `string` | yes |

## UpdateLinkRoutingRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `linkId` | `string` | yes |
| `ruleId` | `string` | yes |
| `priority` | `integer` | yes |
| `conditions` | `RuleCondition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ruleId` | `string` | yes |
| `status` | `string` | yes |
| `updatedAt` | `timestamp` | yes |

## UpdateRequesterGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |

## UpdateResponderGateway

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `domainName` | `string` | no |
| `port` | `integer` | yes |
| `protocol` | `string` | yes |
| `listenerConfig` | `ListenerConfig` | no |
| `trustStoreConfiguration` | `TrustStoreConfiguration` | no |
| `managedEndpointConfiguration` | `ManagedEndpointConfiguration` | no |
| `clientToken` | `string` | yes |
| `gatewayId` | `string` | yes |
| `description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `gatewayId` | `string` | yes |
| `status` | `string` | yes |

