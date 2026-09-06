# AWS User Notifications

API version: 2018-05-10. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/notifications/2018-05-10/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateManagedNotificationAccountContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactIdentifier` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateManagedNotificationAdditionalChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationalUnitId` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateEventRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfigurationArn` | `string` | yes |
| `source` | `string` | yes |
| `eventType` | `string` | yes |
| `eventPattern` | `string` | no |
| `regions` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |
| `statusSummaryByRegion` | `Map<EventRuleStatusSummary>` | yes |

## CreateNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `name` | `string` | yes |
| `description` | `string` | yes |
| `aggregationDuration` | `string` | no |
| `tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `status` | `string` | yes |

## DeleteEventRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeregisterNotificationHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationHubRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationHubRegion` | `string` | yes |
| `statusSummary` | `NotificationHubStatusSummary` | yes |

## DisableNotificationsAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateManagedNotificationAccountContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `contactIdentifier` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateManagedNotificationAdditionalChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelArn` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateOrganizationalUnit

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationalUnitId` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EnableNotificationsAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetEventRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `source` | `string` | yes |
| `eventType` | `string` | yes |
| `eventPattern` | `string` | yes |
| `regions` | `List<string>` | yes |
| `managedRules` | `List<string>` | yes |
| `statusSummaryByRegion` | `Map<EventRuleStatusSummary>` | yes |

## GetManagedNotificationChildEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `content` | `ManagedNotificationChildEvent` | yes |

## GetManagedNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `category` | `string` | yes |
| `subCategory` | `string` | yes |

## GetManagedNotificationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `managedNotificationConfigurationArn` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `content` | `ManagedNotificationEvent` | yes |

## GetNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | yes |
| `description` | `string` | yes |
| `status` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `aggregationDuration` | `string` | no |
| `subtype` | `string` | no |

## GetNotificationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `locale` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |
| `creationTime` | `timestamp` | yes |
| `content` | `NotificationEvent` | yes |

## GetNotificationsAccessForOrganization

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationsAccessForOrganization` | `NotificationsAccessForOrganization` | yes |

## ListChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfigurationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `channels` | `List<string>` | yes |

## ListEventRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfigurationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `eventRules` | `List<EventRuleStructure>` | yes |

## ListManagedNotificationChannelAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `managedNotificationConfigurationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `channelAssociations` | `List<ManagedNotificationChannelAssociationSummary>` | yes |

## ListManagedNotificationChildEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `aggregateManagedNotificationEventArn` | `string` | yes |
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `locale` | `string` | no |
| `maxResults` | `integer` | no |
| `relatedAccount` | `string` | no |
| `organizationalUnitId` | `string` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `managedNotificationChildEvents` | `List<ManagedNotificationChildEventOverview>` | yes |

## ListManagedNotificationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `channelIdentifier` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `managedNotificationConfigurations` | `List<ManagedNotificationConfigurationStructure>` | yes |

## ListManagedNotificationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `locale` | `string` | no |
| `source` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `organizationalUnitId` | `string` | no |
| `relatedAccount` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `managedNotificationEvents` | `List<ManagedNotificationEventOverview>` | yes |

## ListMemberAccounts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfigurationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `memberAccount` | `string` | no |
| `status` | `string` | no |
| `organizationalUnitId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `memberAccounts` | `List<MemberAccount>` | yes |
| `nextToken` | `string` | no |

## ListNotificationConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `eventRuleSource` | `string` | no |
| `channelArn` | `string` | no |
| `status` | `string` | no |
| `subtype` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `notificationConfigurations` | `List<NotificationConfigurationStructure>` | yes |

## ListNotificationEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `startTime` | `timestamp` | no |
| `endTime` | `timestamp` | no |
| `locale` | `string` | no |
| `source` | `string` | no |
| `includeChildEvents` | `boolean` | no |
| `aggregateNotificationEventArn` | `string` | no |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |
| `organizationalUnitId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `nextToken` | `string` | no |
| `notificationEvents` | `List<NotificationEventOverview>` | yes |

## ListNotificationHubs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationHubs` | `List<NotificationHubOverview>` | yes |
| `nextToken` | `string` | no |

## ListOrganizationalUnits

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationConfigurationArn` | `string` | yes |
| `maxResults` | `integer` | no |
| `nextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `organizationalUnits` | `List<string>` | yes |
| `nextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## RegisterNotificationHub

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationHubRegion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `notificationHubRegion` | `string` | yes |
| `statusSummary` | `NotificationHubStatusSummary` | yes |
| `creationTime` | `timestamp` | yes |
| `lastActivationTime` | `timestamp` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEventRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `eventPattern` | `string` | no |
| `regions` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `notificationConfigurationArn` | `string` | yes |
| `statusSummaryByRegion` | `Map<EventRuleStatusSummary>` | yes |

## UpdateNotificationConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |
| `name` | `string` | no |
| `description` | `string` | no |
| `aggregationDuration` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `arn` | `string` | yes |

