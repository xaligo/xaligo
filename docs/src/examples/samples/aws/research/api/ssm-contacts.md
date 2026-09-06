# AWS Systems Manager Incident Manager Contacts

API version: 2021-05-03. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/ssm-contacts/2021-05-03/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AcceptPage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageId` | `string` | yes |
| `ContactChannelId` | `string` | no |
| `AcceptType` | `string` | yes |
| `Note` | `string` | no |
| `AcceptCode` | `string` | yes |
| `AcceptCodeValidation` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ActivateContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |
| `ActivationCode` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Alias` | `string` | yes |
| `DisplayName` | `string` | no |
| `Type` | `string` | yes |
| `Plan` | `Plan` | yes |
| `Tags` | `List<Tag>` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |

## CreateContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `DeliveryAddress` | `ContactChannelAddress` | yes |
| `DeferActivation` | `boolean` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelArn` | `string` | yes |

## CreateRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ContactIds` | `List<string>` | yes |
| `StartTime` | `timestamp` | no |
| `TimeZoneId` | `string` | yes |
| `Recurrence` | `RecurrenceSettings` | yes |
| `Tags` | `List<Tag>` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationArn` | `string` | yes |

## CreateRotationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `NewContactIds` | `List<string>` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationOverrideId` | `string` | yes |

## DeactivateContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRotationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `RotationOverrideId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |
| `EngagementArn` | `string` | yes |
| `Sender` | `string` | yes |
| `Subject` | `string` | yes |
| `Content` | `string` | yes |
| `PublicSubject` | `string` | no |
| `PublicContent` | `string` | no |
| `IncidentId` | `string` | no |
| `StartTime` | `timestamp` | no |
| `StopTime` | `timestamp` | no |

## DescribePage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageArn` | `string` | yes |
| `EngagementArn` | `string` | yes |
| `ContactArn` | `string` | yes |
| `Sender` | `string` | yes |
| `Subject` | `string` | yes |
| `Content` | `string` | yes |
| `PublicSubject` | `string` | no |
| `PublicContent` | `string` | no |
| `IncidentId` | `string` | no |
| `SentTime` | `timestamp` | no |
| `ReadTime` | `timestamp` | no |
| `DeliveryTime` | `timestamp` | no |

## GetContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |
| `Alias` | `string` | yes |
| `DisplayName` | `string` | no |
| `Type` | `string` | yes |
| `Plan` | `Plan` | yes |

## GetContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |
| `ContactChannelArn` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `DeliveryAddress` | `ContactChannelAddress` | yes |
| `ActivationStatus` | `string` | no |

## GetContactPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | no |
| `Policy` | `string` | no |

## GetRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationArn` | `string` | yes |
| `Name` | `string` | yes |
| `ContactIds` | `List<string>` | yes |
| `StartTime` | `timestamp` | yes |
| `TimeZoneId` | `string` | yes |
| `Recurrence` | `RecurrenceSettings` | yes |

## GetRotationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `RotationOverrideId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationOverrideId` | `string` | no |
| `RotationArn` | `string` | no |
| `NewContactIds` | `List<string>` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `CreateTime` | `timestamp` | no |

## ListContactChannels

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ContactChannels` | `List<ContactChannel>` | yes |

## ListContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AliasPrefix` | `string` | no |
| `Type` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Contacts` | `List<Contact>` | no |

## ListEngagements

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncidentId` | `string` | no |
| `TimeRangeValue` | `TimeRange` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Engagements` | `List<Engagement>` | yes |

## ListPageReceipts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PageId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Receipts` | `List<Receipt>` | no |

## ListPageResolutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PageResolutions` | `List<ResolutionContact>` | yes |

## ListPagesByContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Pages` | `List<Page>` | yes |

## ListPagesByEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Pages` | `List<Page>` | yes |

## ListPreviewRotationShifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationStartTime` | `timestamp` | no |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | yes |
| `Members` | `List<string>` | yes |
| `TimeZoneId` | `string` | yes |
| `Recurrence` | `RecurrenceSettings` | yes |
| `Overrides` | `List<PreviewOverride>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationShifts` | `List<RotationShift>` | no |
| `NextToken` | `string` | no |

## ListRotationOverrides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationOverrides` | `List<RotationOverride>` | no |
| `NextToken` | `string` | no |

## ListRotationShifts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationShifts` | `List<RotationShift>` | no |
| `NextToken` | `string` | no |

## ListRotations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationNamePrefix` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Rotations` | `List<Rotation>` | yes |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceARN` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<Tag>` | no |

## PutContactPolicy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactArn` | `string` | yes |
| `Policy` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendActivationCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `Sender` | `string` | yes |
| `Subject` | `string` | yes |
| `Content` | `string` | yes |
| `PublicSubject` | `string` | no |
| `PublicContent` | `string` | no |
| `IncidentId` | `string` | no |
| `IdempotencyToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementArn` | `string` | yes |

## StopEngagement

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EngagementId` | `string` | yes |
| `Reason` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


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


## UpdateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `DisplayName` | `string` | no |
| `Plan` | `Plan` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactChannel

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactChannelId` | `string` | yes |
| `Name` | `string` | no |
| `DeliveryAddress` | `ContactChannelAddress` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRotation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RotationId` | `string` | yes |
| `ContactIds` | `List<string>` | no |
| `StartTime` | `timestamp` | no |
| `TimeZoneId` | `string` | no |
| `Recurrence` | `RecurrenceSettings` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


