# AWS License Manager Linux Subscriptions

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/license-manager-linux-subscriptions/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## DeregisterSubscriptionProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetRegisteredSubscriptionProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionProviderArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LastSuccessfulDataRetrievalTime` | `string` | no |
| `SecretArn` | `string` | no |
| `SubscriptionProviderArn` | `string` | no |
| `SubscriptionProviderSource` | `string` | no |
| `SubscriptionProviderStatus` | `string` | no |
| `SubscriptionProviderStatusMessage` | `string` | no |

## GetServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegions` | `List<string>` | no |
| `LinuxSubscriptionsDiscovery` | `string` | no |
| `LinuxSubscriptionsDiscoverySettings` | `LinuxSubscriptionsDiscoverySettings` | no |
| `Status` | `string` | no |
| `StatusMessage` | `Map<string>` | no |

## ListLinuxSubscriptionInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instances` | `List<Instance>` | no |
| `NextToken` | `string` | no |

## ListLinuxSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<Filter>` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Subscriptions` | `List<Subscription>` | no |

## ListRegisteredSubscriptionProviders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SubscriptionProviderSources` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RegisteredSubscriptionProviders` | `List<RegisteredSubscriptionProvider>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RegisterSubscriptionProvider

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecretArn` | `string` | yes |
| `SubscriptionProviderSource` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionProviderArn` | `string` | no |
| `SubscriptionProviderSource` | `string` | no |
| `SubscriptionProviderStatus` | `string` | no |

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


## UpdateServiceSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowUpdate` | `boolean` | no |
| `LinuxSubscriptionsDiscovery` | `string` | yes |
| `LinuxSubscriptionsDiscoverySettings` | `LinuxSubscriptionsDiscoverySettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HomeRegions` | `List<string>` | no |
| `LinuxSubscriptionsDiscovery` | `string` | no |
| `LinuxSubscriptionsDiscoverySettings` | `LinuxSubscriptionsDiscoverySettings` | no |
| `Status` | `string` | no |
| `StatusMessage` | `Map<string>` | no |

