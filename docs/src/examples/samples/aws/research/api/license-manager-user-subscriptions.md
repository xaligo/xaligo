# AWS License Manager User Subscriptions

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/license-manager-user-subscriptions/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `InstanceId` | `string` | yes |
| `IdentityProvider` | `IdentityProvider` | yes |
| `Domain` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceUserSummary` | `InstanceUserSummary` | yes |

## CreateLicenseServerEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderArn` | `string` | yes |
| `LicenseServerSettings` | `LicenseServerSettings` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderArn` | `string` | no |
| `LicenseServerEndpointArn` | `string` | no |

## DeleteLicenseServerEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseServerEndpointArn` | `string` | yes |
| `ServerType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseServerEndpoint` | `LicenseServerEndpoint` | no |

## DeregisterIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProvider` | no |
| `Product` | `string` | no |
| `IdentityProviderArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderSummary` | `IdentityProviderSummary` | yes |

## DisassociateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | no |
| `InstanceId` | `string` | no |
| `IdentityProvider` | `IdentityProvider` | no |
| `InstanceUserArn` | `string` | no |
| `Domain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceUserSummary` | `InstanceUserSummary` | yes |

## ListIdentityProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderSummaries` | `List<IdentityProviderSummary>` | yes |
| `NextToken` | `string` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Filters` | `List<Filter>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceSummaries` | `List<InstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListLicenseServerEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LicenseServerEndpoints` | `List<LicenseServerEndpoint>` | no |
| `NextToken` | `string` | no |

## ListProductSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Product` | `string` | no |
| `IdentityProvider` | `IdentityProvider` | yes |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductUserSummaries` | `List<ProductUserSummary>` | no |
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

## ListUserAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IdentityProvider` | `IdentityProvider` | yes |
| `MaxResults` | `integer` | no |
| `Filters` | `List<Filter>` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceUserSummaries` | `List<InstanceUserSummary>` | no |
| `NextToken` | `string` | no |

## RegisterIdentityProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProvider` | yes |
| `Product` | `string` | yes |
| `Settings` | `Settings` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderSummary` | `IdentityProviderSummary` | yes |

## StartProductSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `IdentityProvider` | `IdentityProvider` | yes |
| `Product` | `string` | yes |
| `Domain` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductUserSummary` | `ProductUserSummary` | yes |

## StopProductSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | no |
| `IdentityProvider` | `IdentityProvider` | no |
| `Product` | `string` | no |
| `ProductUserArn` | `string` | no |
| `Domain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProductUserSummary` | `ProductUserSummary` | yes |

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


## UpdateIdentityProviderSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProvider` | `IdentityProvider` | no |
| `Product` | `string` | no |
| `IdentityProviderArn` | `string` | no |
| `UpdateSettings` | `UpdateSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityProviderSummary` | `IdentityProviderSummary` | yes |

