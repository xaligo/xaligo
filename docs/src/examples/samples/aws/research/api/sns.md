# Amazon Simple Notification Service

API version: 2010-03-31. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sns/2010-03-31/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddPermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `Label` | `string` | yes |
| `AWSAccountId` | `List<string>` | yes |
| `ActionName` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CheckIfPhoneNumberIsOptedOut

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `phoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `isOptedOut` | `boolean` | no |

## ConfirmSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `Token` | `string` | yes |
| `AuthenticateOnUnsubscribe` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | no |

## CreatePlatformApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Platform` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | no |

## CreatePlatformEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | yes |
| `Token` | `string` | yes |
| `CustomUserData` | `string` | no |
| `Attributes` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | no |

## CreateSMSSandboxPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `string` | yes |
| `LanguageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Attributes` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `DataProtectionPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |

## DeleteEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePlatformApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSMSSandboxPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetDataProtectionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataProtectionPolicy` | `string` | no |

## GetEndpointAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## GetPlatformApplicationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## GetSMSAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `Map<string>` | no |

## GetSMSSandboxAccountStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IsInSandbox` | `boolean` | yes |

## GetSubscriptionAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## GetTopicAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## ListEndpointsByPlatformApplication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Endpoints` | `List<Endpoint>` | no |
| `NextToken` | `string` | no |

## ListOriginationNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PhoneNumbers` | `List<PhoneNumberInformation>` | no |

## ListPhoneNumbersOptedOut

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `phoneNumbers` | `List<string>` | no |
| `nextToken` | `string` | no |

## ListPlatformApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplications` | `List<PlatformApplication>` | no |
| `NextToken` | `string` | no |

## ListSMSSandboxPhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumbers` | `List<SMSSandboxPhoneNumber>` | yes |
| `NextToken` | `string` | no |

## ListSubscriptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscriptions` | `List<Subscription>` | no |
| `NextToken` | `string` | no |

## ListSubscriptionsByTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Subscriptions` | `List<Subscription>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## ListTopics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Topics` | `List<Topic>` | no |
| `NextToken` | `string` | no |

## OptInPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `phoneNumber` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Publish

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | no |
| `TargetArn` | `string` | no |
| `PhoneNumber` | `string` | no |
| `Message` | `string` | yes |
| `Subject` | `string` | no |
| `MessageStructure` | `string` | no |
| `MessageAttributes` | `Map<MessageAttributeValue>` | no |
| `MessageDeduplicationId` | `string` | no |
| `MessageGroupId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |
| `SequenceNumber` | `string` | no |

## PublishBatch

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `PublishBatchRequestEntries` | `List<PublishBatchRequestEntry>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<PublishBatchResultEntry>` | no |
| `Failed` | `List<BatchResultErrorEntry>` | no |

## PutDataProtectionPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `DataProtectionPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RemovePermission

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `Label` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetEndpointAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointArn` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetPlatformApplicationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PlatformApplicationArn` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetSMSAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetSubscriptionAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |
| `AttributeName` | `string` | yes |
| `AttributeValue` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetTopicAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `AttributeName` | `string` | yes |
| `AttributeValue` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Subscribe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TopicArn` | `string` | yes |
| `Protocol` | `string` | yes |
| `Endpoint` | `string` | no |
| `Attributes` | `Map<string>` | no |
| `ReturnSubscriptionArn` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## Unsubscribe

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SubscriptionArn` | `string` | yes |

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


## VerifySMSSandboxPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumber` | `string` | yes |
| `OneTimePassword` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


