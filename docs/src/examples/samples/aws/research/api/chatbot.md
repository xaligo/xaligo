# AWS Chatbot

API version: 2017-10-11. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/chatbot/2017-10-11/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateToConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `ChatConfiguration` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateChimeWebhookConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebhookDescription` | `string` | yes |
| `WebhookUrl` | `string` | yes |
| `SnsTopicArns` | `List<string>` | yes |
| `IamRoleArn` | `string` | yes |
| `ConfigurationName` | `string` | yes |
| `LoggingLevel` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebhookConfiguration` | `ChimeWebhookConfiguration` | no |

## CreateCustomAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Definition` | `CustomActionDefinition` | yes |
| `AliasName` | `string` | no |
| `Attachments` | `List<CustomActionAttachment>` | no |
| `Tags` | `List<Tag>` | no |
| `ClientToken` | `string` | no |
| `ActionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActionArn` | `string` | yes |

## CreateMicrosoftTeamsChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelId` | `string` | yes |
| `ChannelName` | `string` | no |
| `TeamId` | `string` | yes |
| `TeamName` | `string` | no |
| `TenantId` | `string` | yes |
| `SnsTopicArns` | `List<string>` | no |
| `IamRoleArn` | `string` | yes |
| `ConfigurationName` | `string` | yes |
| `LoggingLevel` | `string` | no |
| `GuardrailPolicyArns` | `List<string>` | no |
| `UserAuthorizationRequired` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelConfiguration` | `TeamsChannelConfiguration` | no |

## CreateSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SlackTeamId` | `string` | yes |
| `SlackChannelId` | `string` | yes |
| `SlackChannelName` | `string` | no |
| `SnsTopicArns` | `List<string>` | no |
| `IamRoleArn` | `string` | yes |
| `ConfigurationName` | `string` | yes |
| `LoggingLevel` | `string` | no |
| `GuardrailPolicyArns` | `List<string>` | no |
| `UserAuthorizationRequired` | `boolean` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelConfiguration` | `SlackChannelConfiguration` | no |

## DeleteChimeWebhookConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMicrosoftTeamsChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMicrosoftTeamsConfiguredTeam

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TeamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMicrosoftTeamsUserIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlackUserIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |
| `SlackTeamId` | `string` | yes |
| `SlackUserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlackWorkspaceAuthorization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SlackTeamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeChimeWebhookConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChatConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WebhookConfigurations` | `List<ChimeWebhookConfiguration>` | no |

## DescribeSlackChannelConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `ChatConfigurationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `SlackChannelConfigurations` | `List<SlackChannelConfiguration>` | no |

## DescribeSlackUserIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SlackUserIdentities` | `List<SlackUserIdentity>` | no |
| `NextToken` | `string` | no |

## DescribeSlackWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SlackWorkspaces` | `List<SlackWorkspace>` | no |
| `NextToken` | `string` | no |

## DisassociateFromConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Resource` | `string` | yes |
| `ChatConfiguration` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountPreferences` | `AccountPreferences` | no |

## GetCustomAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActionArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomAction` | `CustomAction` | no |

## GetMicrosoftTeamsChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelConfiguration` | `TeamsChannelConfiguration` | no |

## ListAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfiguration` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Associations` | `List<AssociationListing>` | yes |
| `NextToken` | `string` | no |

## ListCustomActions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActions` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListMicrosoftTeamsChannelConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `TeamId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TeamChannelConfigurations` | `List<TeamsChannelConfiguration>` | no |

## ListMicrosoftTeamsConfiguredTeams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfiguredTeams` | `List<ConfiguredTeam>` | no |
| `NextToken` | `string` | no |

## ListMicrosoftTeamsUserIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TeamsUserIdentities` | `List<TeamsUserIdentity>` | no |
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


## UpdateAccountPreferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserAuthorizationRequired` | `boolean` | no |
| `TrainingDataCollectionEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AccountPreferences` | `AccountPreferences` | no |

## UpdateChimeWebhookConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |
| `WebhookDescription` | `string` | no |
| `WebhookUrl` | `string` | no |
| `SnsTopicArns` | `List<string>` | no |
| `IamRoleArn` | `string` | no |
| `LoggingLevel` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WebhookConfiguration` | `ChimeWebhookConfiguration` | no |

## UpdateCustomAction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActionArn` | `string` | yes |
| `Definition` | `CustomActionDefinition` | yes |
| `AliasName` | `string` | no |
| `Attachments` | `List<CustomActionAttachment>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomActionArn` | `string` | yes |

## UpdateMicrosoftTeamsChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |
| `ChannelId` | `string` | yes |
| `ChannelName` | `string` | no |
| `SnsTopicArns` | `List<string>` | no |
| `IamRoleArn` | `string` | no |
| `LoggingLevel` | `string` | no |
| `GuardrailPolicyArns` | `List<string>` | no |
| `UserAuthorizationRequired` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelConfiguration` | `TeamsChannelConfiguration` | no |

## UpdateSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChatConfigurationArn` | `string` | yes |
| `SlackChannelId` | `string` | yes |
| `SlackChannelName` | `string` | no |
| `SnsTopicArns` | `List<string>` | no |
| `IamRoleArn` | `string` | no |
| `LoggingLevel` | `string` | no |
| `GuardrailPolicyArns` | `List<string>` | no |
| `UserAuthorizationRequired` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ChannelConfiguration` | `SlackChannelConfiguration` | no |

