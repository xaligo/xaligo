# Amazon Simple Email Service

API version: 2010-12-01. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ses/2010-12-01/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## CloneReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `OriginalRuleSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSet` | `ConfigurationSet` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestination` | `EventDestination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateConfigurationSetTrackingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `TrackingOptions` | `TrackingOptions` | yes |

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
| `SuccessRedirectionURL` | `string` | yes |
| `FailureRedirectionURL` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateReceiptFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filter` | `ReceiptFilter` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateReceiptRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `After` | `string` | no |
| `Rule` | `ReceiptRule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Template` | `Template` | yes |

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


## DeleteConfigurationSetTrackingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |

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


## DeleteIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIdentityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `PolicyName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReceiptFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReceiptRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `RuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVerifiedEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeActiveReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `ReceiptRuleSetMetadata` | no |
| `Rules` | `List<ReceiptRule>` | no |

## DescribeConfigurationSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `ConfigurationSetAttributeNames` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSet` | `ConfigurationSet` | no |
| `EventDestinations` | `List<EventDestination>` | no |
| `TrackingOptions` | `TrackingOptions` | no |
| `DeliveryOptions` | `DeliveryOptions` | no |
| `ReputationOptions` | `ReputationOptions` | no |

## DescribeReceiptRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `RuleName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `ReceiptRule` | no |

## DescribeReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metadata` | `ReceiptRuleSetMetadata` | no |
| `Rules` | `List<ReceiptRule>` | no |

## GetAccountSendingEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Enabled` | `boolean` | no |

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
| `SuccessRedirectionURL` | `string` | no |
| `FailureRedirectionURL` | `string` | no |

## GetIdentityDkimAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identities` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DkimAttributes` | `Map<IdentityDkimAttributes>` | yes |

## GetIdentityMailFromDomainAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identities` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MailFromDomainAttributes` | `Map<IdentityMailFromDomainAttributes>` | yes |

## GetIdentityNotificationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identities` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationAttributes` | `Map<IdentityNotificationAttributes>` | yes |

## GetIdentityPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `PolicyNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Policies` | `Map<string>` | yes |

## GetIdentityVerificationAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identities` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationAttributes` | `Map<IdentityVerificationAttributes>` | yes |

## GetSendQuota

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Max24HourSend` | `double` | no |
| `MaxSendRate` | `double` | no |
| `SentLast24Hours` | `double` | no |

## GetSendStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SendDataPoints` | `List<SendDataPoint>` | no |

## GetTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Template` | `Template` | no |

## ListConfigurationSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSets` | `List<ConfigurationSet>` | no |
| `NextToken` | `string` | no |

## ListCustomVerificationEmailTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CustomVerificationEmailTemplates` | `List<CustomVerificationEmailTemplate>` | no |
| `NextToken` | `string` | no |

## ListIdentities

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identities` | `List<string>` | yes |
| `NextToken` | `string` | no |

## ListIdentityPolicies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PolicyNames` | `List<string>` | yes |

## ListReceiptFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Filters` | `List<ReceiptFilter>` | no |

## ListReceiptRuleSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSets` | `List<ReceiptRuleSetMetadata>` | no |
| `NextToken` | `string` | no |

## ListTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxItems` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplatesMetadata` | `List<TemplateMetadata>` | no |
| `NextToken` | `string` | no |

## ListVerifiedEmailAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|


Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerifiedEmailAddresses` | `List<string>` | no |

## PutConfigurationSetDeliveryOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `DeliveryOptions` | `DeliveryOptions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutIdentityPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `PolicyName` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReorderReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `RuleNames` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendBounce

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `OriginalMessageId` | `string` | yes |
| `BounceSender` | `string` | yes |
| `Explanation` | `string` | no |
| `MessageDsn` | `MessageDsn` | no |
| `BouncedRecipientInfoList` | `List<BouncedRecipientInfo>` | yes |
| `BounceSenderArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | no |

## SendBulkTemplatedEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `string` | yes |
| `SourceArn` | `string` | no |
| `ReplyToAddresses` | `List<string>` | no |
| `ReturnPath` | `string` | no |
| `ReturnPathArn` | `string` | no |
| `ConfigurationSetName` | `string` | no |
| `DefaultTags` | `List<MessageTag>` | no |
| `Template` | `string` | yes |
| `TemplateArn` | `string` | no |
| `DefaultTemplateData` | `string` | yes |
| `Destinations` | `List<BulkEmailDestination>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `List<BulkEmailDestinationStatus>` | yes |

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
| `Source` | `string` | yes |
| `Destination` | `Destination` | yes |
| `Message` | `Message` | yes |
| `ReplyToAddresses` | `List<string>` | no |
| `ReturnPath` | `string` | no |
| `SourceArn` | `string` | no |
| `ReturnPathArn` | `string` | no |
| `Tags` | `List<MessageTag>` | no |
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |

## SendRawEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `string` | no |
| `Destinations` | `List<string>` | no |
| `RawMessage` | `RawMessage` | yes |
| `FromArn` | `string` | no |
| `SourceArn` | `string` | no |
| `ReturnPathArn` | `string` | no |
| `Tags` | `List<MessageTag>` | no |
| `ConfigurationSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |

## SendTemplatedEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Source` | `string` | yes |
| `Destination` | `Destination` | yes |
| `ReplyToAddresses` | `List<string>` | no |
| `ReturnPath` | `string` | no |
| `SourceArn` | `string` | no |
| `ReturnPathArn` | `string` | no |
| `Tags` | `List<MessageTag>` | no |
| `ConfigurationSetName` | `string` | no |
| `Template` | `string` | yes |
| `TemplateArn` | `string` | no |
| `TemplateData` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MessageId` | `string` | yes |

## SetActiveReceiptRuleSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityDkimEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `DkimEnabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityFeedbackForwardingEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `ForwardingEnabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityHeadersInNotificationsEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `NotificationType` | `string` | yes |
| `Enabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityMailFromDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `MailFromDomain` | `string` | no |
| `BehaviorOnMXFailure` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetIdentityNotificationTopic

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Identity` | `string` | yes |
| `NotificationType` | `string` | yes |
| `SnsTopic` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SetReceiptRulePosition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `RuleName` | `string` | yes |
| `After` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TestRenderTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `TemplateData` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RenderedTemplate` | `string` | no |

## UpdateAccountSendingEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Enabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationSetEventDestination

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `EventDestination` | `EventDestination` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationSetReputationMetricsEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `Enabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationSetSendingEnabled

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `Enabled` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateConfigurationSetTrackingOptions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConfigurationSetName` | `string` | yes |
| `TrackingOptions` | `TrackingOptions` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCustomVerificationEmailTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateName` | `string` | yes |
| `FromEmailAddress` | `string` | no |
| `TemplateSubject` | `string` | no |
| `TemplateContent` | `string` | no |
| `SuccessRedirectionURL` | `string` | no |
| `FailureRedirectionURL` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateReceiptRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSetName` | `string` | yes |
| `Rule` | `ReceiptRule` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Template` | `Template` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## VerifyDomainDkim

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DkimTokens` | `List<string>` | yes |

## VerifyDomainIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Domain` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VerificationToken` | `string` | yes |

## VerifyEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## VerifyEmailIdentity

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddress` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


