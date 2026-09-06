# Amazon Security Lake

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/securitylake/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAwsLogSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<AwsLogSourceConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failed` | `List<string>` | no |

## CreateCustomLogSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `CustomLogSourceConfiguration` | yes |
| `eventClasses` | `List<string>` | no |
| `sourceName` | `string` | yes |
| `sourceVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `source` | `CustomLogSourceResource` | no |

## CreateDataLake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<DataLakeConfiguration>` | yes |
| `metaStoreManagerRoleArn` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataLakes` | `List<DataLakeResource>` | no |

## CreateDataLakeExceptionSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exceptionTimeToLive` | `long` | no |
| `notificationEndpoint` | `string` | yes |
| `subscriptionProtocol` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDataLakeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnableNewAccount` | `List<DataLakeAutoEnableNewAccountConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accessTypes` | `List<string>` | no |
| `sources` | `List<LogSourceResource>` | yes |
| `subscriberDescription` | `string` | no |
| `subscriberIdentity` | `AwsIdentity` | yes |
| `subscriberName` | `string` | yes |
| `tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriber` | `SubscriberResource` | no |

## CreateSubscriberNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `NotificationConfiguration` | yes |
| `subscriberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriberEndpoint` | `string` | no |

## DeleteAwsLogSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<AwsLogSourceConfiguration>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `failed` | `List<string>` | no |

## DeleteCustomLogSource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sourceName` | `string` | yes |
| `sourceVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataLake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `regions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataLakeExceptionSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataLakeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnableNewAccount` | `List<DataLakeAutoEnableNewAccountConfiguration>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSubscriberNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterDataLakeDelegatedAdministrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDataLakeExceptionSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exceptionTimeToLive` | `long` | no |
| `notificationEndpoint` | `string` | no |
| `subscriptionProtocol` | `string` | no |

## GetDataLakeOrganizationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `autoEnableNewAccount` | `List<DataLakeAutoEnableNewAccountConfiguration>` | no |

## GetDataLakeSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataLakeArn` | `string` | no |
| `dataLakeSources` | `List<DataLakeSource>` | no |
| `nextToken` | `string` | no |

## GetSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriber` | `SubscriberResource` | no |

## ListDataLakeExceptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exceptions` | `List<DataLakeException>` | no |
| `nextToken` | `string` | no |

## ListDataLakes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataLakes` | `List<DataLakeResource>` | no |

## ListLogSources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accounts` | `List<string>` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `regions` | `List<string>` | no |
| `sources` | `List<LogSourceResource>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `sources` | `List<LogSource>` | no |

## ListSubscribers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `subscribers` | `List<SubscriberResource>` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `List<Tag>` | no |

## RegisterDataLakeDelegatedAdministrator

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `List<Tag>` | yes |

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


## UpdateDataLake

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configurations` | `List<DataLakeConfiguration>` | yes |
| `metaStoreManagerRoleArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `dataLakes` | `List<DataLakeResource>` | no |

## UpdateDataLakeExceptionSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `exceptionTimeToLive` | `long` | no |
| `notificationEndpoint` | `string` | yes |
| `subscriptionProtocol` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSubscriber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `sources` | `List<LogSourceResource>` | no |
| `subscriberDescription` | `string` | no |
| `subscriberId` | `string` | yes |
| `subscriberIdentity` | `AwsIdentity` | no |
| `subscriberName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriber` | `SubscriberResource` | no |

## UpdateSubscriberNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `configuration` | `NotificationConfiguration` | yes |
| `subscriberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `subscriberEndpoint` | `string` | no |

