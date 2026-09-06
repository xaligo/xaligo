# AWS Support App

API version: 2021-08-20. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/support-app/2021-08-20/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | yes |
| `channelName` | `string` | no |
| `channelRoleArn` | `string` | yes |
| `notifyOnAddCorrespondenceToCase` | `boolean` | no |
| `notifyOnCaseSeverity` | `string` | yes |
| `notifyOnCreateOrReopenCase` | `boolean` | no |
| `notifyOnResolveCase` | `boolean` | no |
| `teamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteAccountAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | yes |
| `teamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSlackWorkspaceConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `teamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccountAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountAlias` | `string` | no |

## ListSlackChannelConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `slackChannelConfigurations` | `List<SlackChannelConfiguration>` | yes |

## ListSlackWorkspaceConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `slackWorkspaceConfigurations` | `List<SlackWorkspaceConfiguration>` | no |

## PutAccountAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## RegisterSlackWorkspaceForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `teamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `accountType` | `string` | no |
| `teamId` | `string` | no |
| `teamName` | `string` | no |

## UpdateSlackChannelConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | yes |
| `channelName` | `string` | no |
| `channelRoleArn` | `string` | no |
| `notifyOnAddCorrespondenceToCase` | `boolean` | no |
| `notifyOnCaseSeverity` | `string` | no |
| `notifyOnCreateOrReopenCase` | `boolean` | no |
| `notifyOnResolveCase` | `boolean` | no |
| `teamId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelId` | `string` | no |
| `channelName` | `string` | no |
| `channelRoleArn` | `string` | no |
| `notifyOnAddCorrespondenceToCase` | `boolean` | no |
| `notifyOnCaseSeverity` | `string` | no |
| `notifyOnCreateOrReopenCase` | `boolean` | no |
| `notifyOnResolveCase` | `boolean` | no |
| `teamId` | `string` | no |

