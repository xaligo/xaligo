# Amazon Simple Email Service

API version: 2019-09-27. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/sesv2/2019-09-27/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AssociateEmailIdentityCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `FromAddress` | `string` | no |
| `CertificateArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queries` | `List<BatchGetMetricDataQuery>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<MetricDataResult>` | no |
| `Errors` | `List<MetricDataError>` | no |

## CancelExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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
| `SuppressionOptions` | `SuppressionOptions` | no |
| `VdmOptions` | `VdmOptions` | no |
| `ArchivingOptions` | `ArchivingOptions` | no |
| `MessageSecurityOptions` | `MessageSecurityOptions` | no |

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


## CreateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `EmailAddress` | `string` | yes |
| `TopicPreferences` | `List<TopicPreference>` | no |
| `UnsubscribeAll` | `boolean` | no |
| `AttributesData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateContactList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `Topics` | `List<Topic>` | no |
| `Description` | `string` | no |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateCustomVerificationEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `FromEmailAddress` | `string` | yes |
| `TemplateSubject` | `string` | yes |
| `TemplateContent` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `SuccessRedirectionURL` | `string` | yes |
| `FailureRedirectionURL` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateDedicatedIpPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `ScalingMode` | `string` | no |

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
| `DkimSigningAttributes` | `DkimSigningAttributes` | no |
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityType` | `string` | no |
| `VerifiedForSendingStatus` | `boolean` | no |
| `DkimAttributes` | `DkimAttributes` | no |

## CreateEmailIdentityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `PolicyName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `TemplateContent` | `EmailTemplateContent` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportDataSource` | `ExportDataSource` | yes |
| `ExportDestination` | `ExportDestination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## CreateImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportDestination` | `ImportDestination` | yes |
| `ImportDataSource` | `ImportDataSource` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |

## CreateMultiRegionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |
| `Details` | `Details` | yes |
| `Tags` | `List<Tag>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `EndpointId` | `string` | no |

## CreateTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |
| `Tags` | `List<Tag>` | no |
| `SuppressionAttributes` | `TenantSuppressionAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | no |
| `TenantId` | `string` | no |
| `TenantArn` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |
| `SendingStatus` | `string` | no |
| `SuppressionAttributes` | `TenantSuppressionAttributes` | no |

## CreateTenantResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## DeleteContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteCustomVerificationEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

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


## DeleteEmailIdentityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMultiRegionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |

## DeleteSuppressedDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |
| `TenantName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTenantResourceAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |
| `ResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateEmailIdentityCertificate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `FromAddress` | `string` | no |

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
| `DedicatedIpAutoWarmupEnabled` | `boolean` | no |
| `EnforcementStatus` | `string` | no |
| `ProductionAccessEnabled` | `boolean` | no |
| `SendQuota` | `SendQuota` | no |
| `SendingEnabled` | `boolean` | no |
| `SuppressionAttributes` | `SuppressionAttributes` | no |
| `Details` | `AccountDetails` | no |
| `VdmAttributes` | `VdmAttributes` | no |
| `PricingAttributes` | `PricingAttributes` | no |

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
| `SuppressionOptions` | `SuppressionOptions` | no |
| `VdmOptions` | `VdmOptions` | no |
| `ArchivingOptions` | `ArchivingOptions` | no |
| `MessageSecurityOptions` | `MessageSecurityOptions` | no |

## GetConfigurationSetEventDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventDestinations` | `List<EventDestination>` | no |

## GetContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | no |
| `EmailAddress` | `string` | no |
| `TopicPreferences` | `List<TopicPreference>` | no |
| `TopicDefaultPreferences` | `List<TopicPreference>` | no |
| `UnsubscribeAll` | `boolean` | no |
| `AttributesData` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |

## GetContactList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | no |
| `Topics` | `List<Topic>` | no |
| `Description` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |
| `Tags` | `List<Tag>` | no |

## GetCustomVerificationEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | no |
| `FromEmailAddress` | `string` | no |
| `TemplateSubject` | `string` | no |
| `TemplateContent` | `string` | no |
| `Tags` | `List<Tag>` | no |
| `SuccessRedirectionURL` | `string` | no |
| `FailureRedirectionURL` | `string` | no |

## GetDedicatedIp

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Ip` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedIp` | `DedicatedIp` | no |

## GetDedicatedIpPool

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DedicatedIpPool` | `DedicatedIpPool` | no |

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

## GetEmailAddressInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MailboxValidation` | `MailboxValidation` | no |

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
| `Policies` | `Map<string>` | no |
| `Tags` | `List<Tag>` | no |
| `ConfigurationSetName` | `string` | no |
| `VerificationStatus` | `string` | no |
| `VerificationInfo` | `VerificationInfo` | no |

## GetEmailIdentityPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `Map<string>` | no |

## GetEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `TemplateContent` | `EmailTemplateContent` | yes |
| `Tags` | `List<Tag>` | no |

## GetExportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `ExportSourceType` | `string` | no |
| `JobStatus` | `string` | no |
| `ExportDestination` | `ExportDestination` | no |
| `ExportDataSource` | `ExportDataSource` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `CompletedTimestamp` | `timestamp` | no |
| `FailureInfo` | `FailureInfo` | no |
| `Statistics` | `ExportStatistics` | no |

## GetImportJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `ImportDestination` | `ImportDestination` | no |
| `ImportDataSource` | `ImportDataSource` | no |
| `FailureInfo` | `FailureInfo` | no |
| `JobStatus` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `CompletedTimestamp` | `timestamp` | no |
| `ProcessedRecordsCount` | `integer` | no |
| `FailedRecordsCount` | `integer` | no |

## GetMessageInsights

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |
| `FromEmailAddress` | `string` | no |
| `Subject` | `string` | no |
| `EmailTags` | `List<MessageTag>` | no |
| `Insights` | `List<EmailInsights>` | no |

## GetMultiRegionEndpoint

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EndpointName` | `string` | no |
| `EndpointId` | `string` | no |
| `Routes` | `List<Route>` | no |
| `Status` | `string` | no |
| `CreatedTimestamp` | `timestamp` | no |
| `LastUpdatedTimestamp` | `timestamp` | no |

## GetReputationEntity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReputationEntityReference` | `string` | yes |
| `ReputationEntityType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReputationEntity` | `ReputationEntity` | no |

## GetSuppressedDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |
| `TenantName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuppressedDestination` | `SuppressedDestination` | yes |

## GetTenant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tenant` | `Tenant` | no |

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

## ListContactLists

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactLists` | `List<ContactList>` | no |
| `NextToken` | `string` | no |

## ListContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `Filter` | `ListContactsFilter` | no |
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Contacts` | `List<Contact>` | no |
| `NextToken` | `string` | no |

## ListCustomVerificationEmailTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomVerificationEmailTemplates` | `List<CustomVerificationEmailTemplateMetadata>` | no |
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

## ListEmailIdentityCertificates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Certificates` | `List<IdentityCertificate>` | no |
| `NextToken` | `string` | no |

## ListEmailTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplatesMetadata` | `List<EmailTemplateMetadata>` | no |
| `NextToken` | `string` | no |

## ListExportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |
| `ExportSourceType` | `string` | no |
| `JobStatus` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExportJobs` | `List<ExportJobSummary>` | no |
| `NextToken` | `string` | no |

## ListImportJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportDestinationType` | `string` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ImportJobs` | `List<ImportJobSummary>` | no |
| `NextToken` | `string` | no |

## ListMultiRegionEndpoints

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MultiRegionEndpoints` | `List<MultiRegionEndpoint>` | no |
| `NextToken` | `string` | no |

## ListRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `Map<string>` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recommendations` | `List<Recommendation>` | no |
| `NextToken` | `string` | no |

## ListReputationEntities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `Map<string>` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReputationEntities` | `List<ReputationEntity>` | no |
| `NextToken` | `string` | no |

## ListResourceTenants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceTenants` | `List<ResourceTenantMetadata>` | no |
| `NextToken` | `string` | no |

## ListSuppressedDestinations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | no |
| `Reasons` | `List<string>` | no |
| `StartDate` | `timestamp` | no |
| `EndDate` | `timestamp` | no |
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuppressedDestinationSummaries` | `List<SuppressedDestinationSummary>` | no |
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

## ListTenantResources

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |
| `Filter` | `Map<string>` | no |
| `PageSize` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantResources` | `List<TenantResource>` | no |
| `NextToken` | `string` | no |

## ListTenants

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageSize` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tenants` | `List<TenantInfo>` | no |
| `NextToken` | `string` | no |

## PutAccountDedicatedIpWarmupAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoWarmupEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccountDetails

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MailType` | `string` | yes |
| `WebsiteURL` | `string` | yes |
| `ContactLanguage` | `string` | no |
| `UseCaseDescription` | `string` | no |
| `AdditionalContactEmailAddresses` | `List<string>` | no |
| `ProductionAccessEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccountPricingAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Plan` | `string` | yes |

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


## PutAccountSuppressionAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuppressedReasons` | `List<string>` | no |
| `ValidationAttributes` | `SuppressionValidationAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutAccountVdmAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VdmAttributes` | `VdmAttributes` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetArchivingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `ArchiveArn` | `string` | no |

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
| `MaxDeliverySeconds` | `long` | no |

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


## PutConfigurationSetSuppressionOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `SuppressionScope` | `string` | no |
| `SuppressedReasons` | `List<string>` | no |
| `ValidationOptions` | `SuppressionValidationOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetTrackingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `CustomRedirectDomain` | `string` | no |
| `HttpsPolicy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutConfigurationSetVdmOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `VdmOptions` | `VdmOptions` | no |

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


## PutDedicatedIpPoolScalingAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PoolName` | `string` | yes |
| `ScalingMode` | `string` | yes |

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


## PutEmailIdentityConfigurationSetAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `ConfigurationSetName` | `string` | no |

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


## PutEmailIdentityDkimSigningAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `SigningAttributesOrigin` | `string` | yes |
| `SigningAttributes` | `DkimSigningAttributes` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DkimStatus` | `string` | no |
| `DkimTokens` | `List<string>` | no |
| `SigningHostedZone` | `string` | no |

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


## PutSuppressedDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |
| `Reason` | `string` | yes |
| `TenantName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutTenantSuppressionAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TenantName` | `string` | yes |
| `SuppressedReasons` | `List<string>` | no |
| `SuppressionScope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendBulkEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromEmailAddress` | `string` | no |
| `FromEmailAddressIdentityArn` | `string` | no |
| `ReplyToAddresses` | `List<string>` | no |
| `FeedbackForwardingEmailAddress` | `string` | no |
| `FeedbackForwardingEmailAddressIdentityArn` | `string` | no |
| `DefaultEmailTags` | `List<MessageTag>` | no |
| `DefaultContent` | `BulkEmailContent` | yes |
| `BulkEmailEntries` | `List<BulkEmailEntry>` | yes |
| `ConfigurationSetName` | `string` | no |
| `EndpointId` | `string` | no |
| `TenantName` | `string` | no |
| `ConfigurationOverrides` | `ConfigurationOverrides` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `BulkEmailEntryResults` | `List<BulkEmailEntryResult>` | yes |

## SendCustomVerificationEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |
| `TemplateName` | `string` | yes |
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FromEmailAddress` | `string` | no |
| `FromEmailAddressIdentityArn` | `string` | no |
| `Destination` | `Destination` | no |
| `ReplyToAddresses` | `List<string>` | no |
| `FeedbackForwardingEmailAddress` | `string` | no |
| `FeedbackForwardingEmailAddressIdentityArn` | `string` | no |
| `Content` | `EmailContent` | yes |
| `EmailTags` | `List<MessageTag>` | no |
| `ConfigurationSetName` | `string` | no |
| `EndpointId` | `string` | no |
| `TenantName` | `string` | no |
| `ListManagementOptions` | `ListManagementOptions` | no |
| `ConfigurationOverrides` | `ConfigurationOverrides` | no |

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


## TestRenderEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `TemplateData` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RenderedTemplate` | `string` | yes |

## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `MessageSecurityOptions` | `MessageSecurityOptions` | no |

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


## UpdateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `EmailAddress` | `string` | yes |
| `TopicPreferences` | `List<TopicPreference>` | no |
| `UnsubscribeAll` | `boolean` | no |
| `AttributesData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactList

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactListName` | `string` | yes |
| `Topics` | `List<Topic>` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCustomVerificationEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `FromEmailAddress` | `string` | yes |
| `TemplateSubject` | `string` | yes |
| `TemplateContent` | `string` | yes |
| `SuccessRedirectionURL` | `string` | yes |
| `FailureRedirectionURL` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEmailIdentityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailIdentity` | `string` | yes |
| `PolicyName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `TemplateContent` | `EmailTemplateContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateReputationEntityCustomerManagedStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReputationEntityType` | `string` | yes |
| `ReputationEntityReference` | `string` | yes |
| `SendingStatus` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateReputationEntityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReputationEntityType` | `string` | yes |
| `ReputationEntityReference` | `string` | yes |
| `ReputationEntityPolicy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


