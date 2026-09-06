# Amazon Pinpoint Email Service

API version: 2018-07-26. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/pinpoint-email/2018-07-26/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CreateConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `TrackingOptions` | `TrackingOptions` | no |
| `DeliveryOptions` | `DeliveryOptions` | no |
| `ReputationOptions` | `ReputationOptions` | no |
| `SendingOptions` | `SendingOptions` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |
| `EventDestination` | `EventDestinationDefinition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDedicatedIpPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDeliverabilityTestReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportName` | `string` | no |
| `FromEmailAddress` | `string` | yes |
| `Content` | `EmailContent` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | yes |
| `DeliverabilityTestStatus` | `string` | yes |

## CreateEmailIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityType` | `string` | no |
| `VerifiedForSendingStatus` | `boolean` | no |
| `DkimAttributes` | `DkimAttributes` | no |

## DeleteConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDedicatedIpPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEmailIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## GetAccount

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SendQuota` | `SendQuota` | no |
| `SendingEnabled` | `boolean` | no |
| `DedicatedIpAutoWarmupEnabled` | `boolean` | no |
| `EnforcementStatus` | `string` | no |
| `ProductionAccessEnabled` | `boolean` | no |

## GetBlacklistReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlacklistItemNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BlacklistReport` | `Map<List<BlacklistEntry>>` | yes |

## GetConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | no |
| `TrackingOptions` | `TrackingOptions` | no |
| `DeliveryOptions` | `DeliveryOptions` | no |
| `ReputationOptions` | `ReputationOptions` | no |
| `SendingOptions` | `SendingOptions` | no |
| `Tags` | `List<Tag>` | no |

## GetConfigurationSetEventDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDestinations` | `List<EventDestination>` | no |

## GetDedicatedIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ip` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedIp` | `DedicatedIp` | no |

## GetDedicatedIps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedIps` | `List<DedicatedIp>` | no |
| `NextToken` | `string` | no |

## GetDeliverabilityDashboardOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardEnabled` | `boolean` | yes |
| `SubscriptionExpiryDate` | `timestamp` | no |
| `AccountStatus` | `string` | no |
| `ActiveSubscribedDomains` | `List<DomainDeliverabilityTrackingOption>` | no |
| `PendingExpirationSubscribedDomains` | `List<DomainDeliverabilityTrackingOption>` | no |

## GetDeliverabilityTestReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReportId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliverabilityTestReport` | `DeliverabilityTestReport` | yes |
| `OverallPlacement` | `PlacementStatistics` | yes |
| `IspPlacements` | `List<IspPlacement>` | yes |
| `Message` | `string` | no |
| `Tags` | `List<Tag>` | no |

## GetDomainDeliverabilityCampaign

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CampaignId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainDeliverabilityCampaign` | `DomainDeliverabilityCampaign` | yes |

## GetDomainStatisticsReport

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |
| `StartDate` | `timestamp` | yes |
| `EndDate` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OverallVolume` | `OverallVolume` | yes |
| `DailyVolumes` | `List<DailyVolume>` | yes |

## GetEmailIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityType` | `string` | no |
| `FeedbackForwardingStatus` | `boolean` | no |
| `VerifiedForSendingStatus` | `boolean` | no |
| `DkimAttributes` | `DkimAttributes` | no |
| `MailFromAttributes` | `MailFromAttributes` | no |
| `Tags` | `List<Tag>` | no |

## ListConfigurationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSets` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListDedicatedIpPools

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedIpPools` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListDeliverabilityTestReports

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DeliverabilityTestReports` | `List<DeliverabilityTestReport>` | yes |
| `NextToken` | `string` | no |

## ListDomainDeliverabilityCampaigns

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartDate` | `timestamp` | yes |
| `EndDate` | `timestamp` | yes |
| `SubscribedDomain` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainDeliverabilityCampaigns` | `List<DomainDeliverabilityCampaign>` | yes |
| `NextToken` | `string` | no |

## ListEmailIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentities` | `List<IdentityInfo>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | yes |

## PutAccountDedicatedIpWarmupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoWarmupEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccountSendingAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SendingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetDeliveryOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `TlsPolicy` | `string` | no |
| `SendingPoolName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetReputationOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `ReputationMetricsEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetSendingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `SendingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetTrackingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `CustomRedirectDomain` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDedicatedIpInPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ip` | `string` | yes |
| `DestinationPoolName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDedicatedIpWarmupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ip` | `string` | yes |
| `WarmupPercentage` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutDeliverabilityDashboardOption

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DashboardEnabled` | `boolean` | yes |
| `SubscribedDomains` | `List<DomainDeliverabilityTrackingOption>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEmailIdentityDkimAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `SigningEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEmailIdentityFeedbackAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `EmailForwardingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutEmailIdentityMailFromAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `MailFromDomain` | `string` | no |
| `BehaviorOnMxFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromEmailAddress` | `string` | no |
| `Destination` | `Destination` | yes |
| `ReplyToAddresses` | `List<string>` | no |
| `FeedbackForwardingEmailAddress` | `string` | no |
| `Content` | `EmailContent` | yes |
| `EmailTags` | `List<MessageTag>` | no |
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `Tags` | `List<Tag>` | yes |

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


## UpdateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestinationName` | `string` | yes |
| `EventDestination` | `EventDestinationDefinition` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


