# Amazon Chime SDK Identity

API version: 2021-04-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chime-sdk-identity/2021-04-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateAppInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Metadata` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | no |

## CreateAppInstanceAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceAdminArn` | `string` | yes |
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceAdmin` | `Identity` | no |
| `AppInstanceArn` | `string` | no |

## CreateAppInstanceBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `Name` | `string` | no |
| `Metadata` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `Configuration` | `Configuration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBotArn` | `string` | no |

## CreateAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `AppInstanceUserId` | `string` | yes |
| `Name` | `string` | yes |
| `Metadata` | `string` | no |
| `ClientRequestToken` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ExpirationSettings` | `ExpirationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |

## DeleteAppInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppInstanceAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceAdminArn` | `string` | yes |
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppInstanceBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBotArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterAppInstanceUserEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAppInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstance` | `AppInstance` | no |

## DescribeAppInstanceAdmin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceAdminArn` | `string` | yes |
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceAdmin` | `AppInstanceAdmin` | no |

## DescribeAppInstanceBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBotArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBot` | `AppInstanceBot` | no |

## DescribeAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUser` | `AppInstanceUser` | no |

## DescribeAppInstanceUserEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `EndpointId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserEndpoint` | `AppInstanceUserEndpoint` | no |

## GetAppInstanceRetentionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceRetentionSettings` | `AppInstanceRetentionSettings` | no |
| `InitiateDeletionTimestamp` | `timestamp` | no |

## ListAppInstanceAdmins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | no |
| `AppInstanceAdmins` | `List<AppInstanceAdminSummary>` | no |
| `NextToken` | `string` | no |

## ListAppInstanceBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | no |
| `AppInstanceBots` | `List<AppInstanceBotSummary>` | no |
| `NextToken` | `string` | no |

## ListAppInstanceUserEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserEndpoints` | `List<AppInstanceUserEndpointSummary>` | no |
| `NextToken` | `string` | no |

## ListAppInstanceUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | no |
| `AppInstanceUsers` | `List<AppInstanceUserSummary>` | no |
| `NextToken` | `string` | no |

## ListAppInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstances` | `List<AppInstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutAppInstanceRetentionSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `AppInstanceRetentionSettings` | `AppInstanceRetentionSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceRetentionSettings` | `AppInstanceRetentionSettings` | no |
| `InitiateDeletionTimestamp` | `timestamp` | no |

## PutAppInstanceUserExpirationSettings

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `ExpirationSettings` | `ExpirationSettings` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |
| `ExpirationSettings` | `ExpirationSettings` | no |

## RegisterAppInstanceUserEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `Name` | `string` | no |
| `Type` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `EndpointAttributes` | `EndpointAttributes` | yes |
| `ClientRequestToken` | `string` | yes |
| `AllowMessages` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |
| `EndpointId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAppInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | yes |
| `Name` | `string` | yes |
| `Metadata` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceArn` | `string` | no |

## UpdateAppInstanceBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBotArn` | `string` | yes |
| `Name` | `string` | yes |
| `Metadata` | `string` | yes |
| `Configuration` | `Configuration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceBotArn` | `string` | no |

## UpdateAppInstanceUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `Name` | `string` | yes |
| `Metadata` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |

## UpdateAppInstanceUserEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | yes |
| `EndpointId` | `string` | yes |
| `Name` | `string` | no |
| `AllowMessages` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AppInstanceUserArn` | `string` | no |
| `EndpointId` | `string` | no |

