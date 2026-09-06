# PcaConnectorAd

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pca-connector-ad/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CertificateAuthorityArn` | `string` | yes |
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `VpcInformation` | `VpcInformation` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | no |

## CreateDirectoryRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `DirectoryId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistrationArn` | `string` | no |

## CreateServicePrincipalName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectorArn` | `string` | yes |
| `DirectoryRegistrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ConnectorArn` | `string` | yes |
| `Definition` | `TemplateDefinition` | yes |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | no |

## CreateTemplateGroupAccessControlEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessRights` | `AccessRights` | yes |
| `ClientToken` | `string` | no |
| `GroupDisplayName` | `string` | yes |
| `GroupSecurityIdentifier` | `string` | yes |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDirectoryRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteServicePrincipalName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `DirectoryRegistrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTemplateGroupAccessControlEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupSecurityIdentifier` | `string` | yes |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetConnector

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connector` | `Connector` | no |

## GetDirectoryRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistration` | `DirectoryRegistration` | no |

## GetServicePrincipalName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `DirectoryRegistrationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ServicePrincipalName` | `ServicePrincipalName` | no |

## GetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Template` | `Template` | no |

## GetTemplateGroupAccessControlEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `GroupSecurityIdentifier` | `string` | yes |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessControlEntry` | `AccessControlEntry` | no |

## ListConnectors

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Connectors` | `List<ConnectorSummary>` | no |
| `NextToken` | `string` | no |

## ListDirectoryRegistrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistrations` | `List<DirectoryRegistrationSummary>` | no |
| `NextToken` | `string` | no |

## ListServicePrincipalNames

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DirectoryRegistrationArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ServicePrincipalNames` | `List<ServicePrincipalNameSummary>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `Map<string>` | no |

## ListTemplateGroupAccessControlEntries

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessControlEntries` | `List<AccessControlEntrySummary>` | no |
| `NextToken` | `string` | no |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectorArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Templates` | `List<TemplateSummary>` | no |

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


## UpdateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definition` | `TemplateDefinition` | no |
| `ReenrollAllCertificateHolders` | `boolean` | no |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTemplateGroupAccessControlEntry

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccessRights` | `AccessRights` | no |
| `GroupDisplayName` | `string` | no |
| `GroupSecurityIdentifier` | `string` | yes |
| `TemplateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


