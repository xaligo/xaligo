# Amazon Connect Service

API version: 2017-08-08. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/connect/2017-08-08/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## ActivateEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormArn` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

## AssociateAnalyticsDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `TargetAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataSetId` | `string` | no |
| `TargetAccountId` | `string` | no |
| `ResourceShareId` | `string` | no |
| `ResourceShareArn` | `string` | no |

## AssociateApprovedOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Origin` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `LexBot` | `LexBot` | no |
| `LexV2Bot` | `LexV2Bot` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateContactWithUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateDefaultVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `VocabularyId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateEmailAddressAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddressId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AliasConfiguration` | `AliasConfiguration` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `FlowId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `ParentHoursOfOperationConfigs` | `List<ParentHoursOfOperationConfig>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateInstanceStorageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `StorageConfig` | `InstanceStorageConfig` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |

## AssociateLambdaFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateLexBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `LexBot` | `LexBot` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociatePhoneNumberContactFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateQueueEmailAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `EmailAddressesConfig` | `List<EmailAddressConfig>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateQueueQuickConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `QuickConnectIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateRoutingProfileQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `QueueConfigs` | `List<RoutingProfileQueueConfig>` | no |
| `ManualAssignmentQueueConfigs` | `List<RoutingProfileManualAssignmentQueueConfig>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateSecurityKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Key` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociationId` | `string` | no |

## AssociateSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SecurityProfiles` | `List<SecurityProfileItem>` | yes |
| `EntityType` | `string` | yes |
| `EntityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateTrafficDistributionGroupUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroupId` | `string` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateUserProficiencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `UserId` | `string` | yes |
| `UserProficiencies` | `List<UserProficiency>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## AssociateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulList` | `List<SuccessfulBatchAssociationSummary>` | no |
| `FailedList` | `List<FailedBatchAssociationSummary>` | no |

## BatchAssociateAnalyticsDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataSetIds` | `List<string>` | yes |
| `TargetAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Created` | `List<AnalyticsDataAssociationResult>` | no |
| `Errors` | `List<ErrorResult>` | no |

## BatchCreateDataTableValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Values` | `List<DataTableValue>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<BatchCreateDataTableValueSuccessResult>` | yes |
| `Failed` | `List<BatchCreateDataTableValueFailureResult>` | yes |

## BatchDeleteDataTableValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Values` | `List<DataTableDeleteValueIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<BatchDeleteDataTableValueSuccessResult>` | yes |
| `Failed` | `List<BatchDeleteDataTableValueFailureResult>` | yes |

## BatchDescribeDataTableValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Values` | `List<DataTableValueIdentifier>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<BatchDescribeDataTableValueSuccessResult>` | yes |
| `Failed` | `List<BatchDescribeDataTableValueFailureResult>` | yes |

## BatchDisassociateAnalyticsDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataSetIds` | `List<string>` | yes |
| `TargetAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Deleted` | `List<string>` | no |
| `Errors` | `List<ErrorResult>` | no |

## BatchGetAttachedFileMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileIds` | `List<string>` | yes |
| `InstanceId` | `string` | yes |
| `AssociatedResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Files` | `List<AttachedFile>` | no |
| `Errors` | `List<AttachedFileError>` | no |

## BatchGetFlowAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceIds` | `List<string>` | yes |
| `ResourceType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowAssociationSummaryList` | `List<FlowAssociationSummary>` | no |

## BatchPutContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `ContactDataRequestList` | `List<ContactDataRequest>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulRequestList` | `List<SuccessfulRequest>` | no |
| `FailedRequestList` | `List<FailedRequest>` | no |

## BatchUpdateDataTableValue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Values` | `List<DataTableValue>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<BatchUpdateDataTableValueSuccessResult>` | yes |
| `Failed` | `List<BatchUpdateDataTableValueFailureResult>` | yes |

## ClaimPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | no |
| `InstanceId` | `string` | no |
| `PhoneNumber` | `string` | yes |
| `PhoneNumberDescription` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | no |
| `PhoneNumberArn` | `string` | no |

## CompleteAttachedFileUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FileId` | `string` | yes |
| `AssociatedResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreateAgentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `State` | `string` | yes |
| `DisplayOrder` | `integer` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentStatusARN` | `string` | no |
| `AgentStatusId` | `string` | no |

## CreateAttachedFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `FileUseCaseType` | `string` | yes |
| `FileSourceUri` | `string` | yes |
| `AssociatedResourceArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileArn` | `string` | no |
| `FileId` | `string` | no |
| `CreationTime` | `string` | no |
| `FileStatus` | `string` | no |

## CreateAuthCode

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Scope` | `AuthScope` | yes |
| `MaxSessionDurationMinutes` | `integer` | no |
| `SessionInactivityDurationMinutes` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthCode` | `string` | no |
| `SessionId` | `string` | no |
| `EntityType` | `string` | no |
| `EntityId` | `string` | no |

## CreateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `RelatedContactId` | `string` | no |
| `Attributes` | `Map<string>` | no |
| `References` | `Map<Reference>` | no |
| `Channel` | `string` | yes |
| `InitiationMethod` | `string` | yes |
| `ExpiryDurationInMinutes` | `integer` | no |
| `UserInfo` | `UserInfo` | no |
| `InitiateAs` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |
| `PreviousContactId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `ContactArn` | `string` | no |

## CreateContactFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Type` | `string` | yes |
| `Description` | `string` | no |
| `Content` | `string` | yes |
| `Status` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowId` | `string` | no |
| `ContactFlowArn` | `string` | no |
| `FlowContentSha256` | `string` | no |

## CreateContactFlowModule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Content` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `Settings` | `string` | no |
| `ExternalInvocationConfiguration` | `ExternalInvocationConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## CreateContactFlowModuleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Description` | `string` | no |
| `ContactFlowModuleId` | `string` | yes |
| `ContactFlowModuleVersion` | `long` | yes |
| `AliasName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModuleArn` | `string` | no |
| `Id` | `string` | no |

## CreateContactFlowModuleVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Description` | `string` | no |
| `ContactFlowModuleId` | `string` | yes |
| `FlowModuleContentSha256` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModuleArn` | `string` | no |
| `Version` | `long` | no |

## CreateContactFlowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Description` | `string` | no |
| `ContactFlowId` | `string` | yes |
| `FlowContentSha256` | `string` | no |
| `ContactFlowVersion` | `long` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowArn` | `string` | no |
| `Version` | `long` | no |

## CreateDataTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `TimeZone` | `string` | yes |
| `ValueLockLevel` | `string` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `LockVersion` | `DataTableLockVersion` | yes |

## CreateDataTableAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Name` | `string` | yes |
| `ValueType` | `string` | yes |
| `Description` | `string` | no |
| `Primary` | `boolean` | no |
| `Validation` | `Validation` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `AttributeId` | `string` | no |
| `LockVersion` | `DataTableLockVersion` | yes |

## CreateEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `InstanceId` | `string` | yes |
| `EmailAddress` | `string` | yes |
| `DisplayName` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddressId` | `string` | no |
| `EmailAddressArn` | `string` | no |

## CreateEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Title` | `string` | yes |
| `Description` | `string` | no |
| `Items` | `List<EvaluationFormItem>` | yes |
| `ScoringStrategy` | `EvaluationFormScoringStrategy` | no |
| `AutoEvaluationConfiguration` | `EvaluationFormAutoEvaluationConfiguration` | no |
| `ClientToken` | `string` | no |
| `AsDraft` | `boolean` | no |
| `Tags` | `Map<string>` | no |
| `ReviewConfiguration` | `EvaluationReviewConfiguration` | no |
| `TargetConfiguration` | `EvaluationFormTargetConfiguration` | no |
| `LanguageConfiguration` | `EvaluationFormLanguageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormArn` | `string` | yes |

## CreateExtractionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `ExtractionConfiguration` | `ExtractionConfiguration` | yes |
| `Display` | `ExtractionDefinitionDisplay` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtractionDefinitionArn` | `string` | yes |
| `ExtractionDefinitionId` | `string` | yes |

## CreateHoursOfOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `TimeZone` | `string` | yes |
| `Config` | `List<HoursOfOperationConfig>` | yes |
| `ParentHoursOfOperationConfigs` | `List<ParentHoursOfOperationConfig>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperationId` | `string` | no |
| `HoursOfOperationArn` | `string` | no |

## CreateHoursOfOperationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Config` | `List<HoursOfOperationOverrideConfig>` | yes |
| `EffectiveFrom` | `string` | yes |
| `EffectiveTill` | `string` | yes |
| `RecurrenceConfig` | `RecurrenceConfig` | no |
| `OverrideType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperationOverrideId` | `string` | no |

## CreateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `IdentityManagementType` | `string` | yes |
| `InstanceAlias` | `string` | no |
| `DirectoryId` | `string` | no |
| `InboundCallsEnabled` | `boolean` | yes |
| `OutboundCallsEnabled` | `boolean` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## CreateIntegrationAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationType` | `string` | yes |
| `IntegrationArn` | `string` | yes |
| `SourceApplicationUrl` | `string` | no |
| `SourceApplicationName` | `string` | no |
| `SourceType` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationAssociationId` | `string` | no |
| `IntegrationAssociationArn` | `string` | no |

## CreateMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `MetricCalculation` | `MetricCalculation` | yes |
| `Unit` | `string` | yes |
| `Status` | `string` | no |
| `ClientToken` | `string` | no |
| `Description` | `string` | no |
| `PositiveTrendIndicator` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricArn` | `string` | yes |
| `MetricId` | `string` | yes |

## CreateNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ExpiresAt` | `timestamp` | no |
| `Recipients` | `List<string>` | yes |
| `Priority` | `string` | no |
| `Content` | `Map<string>` | yes |
| `Tags` | `Map<string>` | no |
| `PredefinedNotificationId` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NotificationId` | `string` | yes |
| `NotificationArn` | `string` | yes |

## CreateParticipant

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ClientToken` | `string` | no |
| `ParticipantDetails` | `ParticipantDetailsToAdd` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ParticipantCredentials` | `ParticipantTokenCredentials` | no |
| `ParticipantId` | `string` | no |

## CreatePersistentContactAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `InitialContactId` | `string` | yes |
| `RehydrationType` | `string` | yes |
| `SourceContactId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContinuedFromContactId` | `string` | no |

## CreatePredefinedAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Values` | `PredefinedAttributeValues` | no |
| `Purposes` | `List<string>` | no |
| `AttributeConfiguration` | `InputPredefinedAttributeConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## CreatePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `S3Uri` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PromptARN` | `string` | no |
| `PromptId` | `string` | no |

## CreatePushNotificationRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `PinpointAppArn` | `string` | yes |
| `DeviceToken` | `string` | yes |
| `DeviceType` | `string` | yes |
| `ContactConfiguration` | `ContactConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RegistrationId` | `string` | yes |

## CreateQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `OutboundCallerConfig` | `OutboundCallerConfig` | no |
| `OutboundEmailConfig` | `OutboundEmailConfig` | no |
| `HoursOfOperationId` | `string` | yes |
| `MaxContacts` | `integer` | no |
| `QuickConnectIds` | `List<string>` | no |
| `EmailAddressesConfig` | `List<EmailAddressConfig>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueArn` | `string` | no |
| `QueueId` | `string` | no |

## CreateQuickConnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `QuickConnectConfig` | `QuickConnectConfig` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuickConnectARN` | `string` | no |
| `QuickConnectId` | `string` | no |

## CreateRoutingProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | yes |
| `DefaultOutboundQueueId` | `string` | yes |
| `QueueConfigs` | `List<RoutingProfileQueueConfig>` | no |
| `ManualAssignmentQueueConfigs` | `List<RoutingProfileManualAssignmentQueueConfig>` | no |
| `MediaConcurrencies` | `List<MediaConcurrency>` | yes |
| `Tags` | `Map<string>` | no |
| `AgentAvailabilityTimer` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingProfileArn` | `string` | no |
| `RoutingProfileId` | `string` | no |

## CreateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `TriggerEventSource` | `RuleTriggerEventSource` | yes |
| `Function` | `string` | yes |
| `Actions` | `List<RuleAction>` | yes |
| `PublishStatus` | `string` | yes |
| `PreEvaluationFilters` | `PreEvaluationFilters` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleArn` | `string` | yes |
| `RuleId` | `string` | yes |

## CreateSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileName` | `string` | yes |
| `Description` | `string` | no |
| `Permissions` | `List<string>` | no |
| `InstanceId` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `AllowedAccessControlTags` | `Map<string>` | no |
| `TagRestrictedResources` | `List<string>` | no |
| `Applications` | `List<Application>` | no |
| `HierarchyRestrictedResources` | `List<string>` | no |
| `AllowedAccessControlHierarchyGroupId` | `string` | no |
| `AllowedFlowModules` | `List<FlowModule>` | no |
| `GranularAccessControlConfiguration` | `GranularAccessControlConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileId` | `string` | no |
| `SecurityProfileArn` | `string` | no |

## CreateTaskTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ContactFlowId` | `string` | no |
| `SelfAssignFlowId` | `string` | no |
| `Constraints` | `TaskTemplateConstraints` | no |
| `Defaults` | `TaskTemplateDefaults` | no |
| `Status` | `string` | no |
| `Fields` | `List<TaskTemplateField>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `Arn` | `string` | yes |

## CreateTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Content` | `string` | yes |
| `EntryPoint` | `TestCaseEntryPoint` | no |
| `InitializationData` | `string` | no |
| `Status` | `string` | no |
| `TestCaseId` | `string` | no |
| `Tags` | `Map<string>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCaseId` | `string` | no |
| `TestCaseArn` | `string` | no |

## CreateTrafficDistributionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## CreateUseCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationAssociationId` | `string` | yes |
| `UseCaseType` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UseCaseId` | `string` | no |
| `UseCaseArn` | `string` | no |

## CreateUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Username` | `string` | yes |
| `Password` | `string` | no |
| `IdentityInfo` | `UserIdentityInfo` | no |
| `PhoneConfig` | `UserPhoneConfig` | no |
| `DirectoryUserId` | `string` | no |
| `SecurityProfileIds` | `List<string>` | yes |
| `RoutingProfileId` | `string` | yes |
| `HierarchyGroupId` | `string` | no |
| `InstanceId` | `string` | yes |
| `AutoAcceptConfigs` | `List<AutoAcceptConfig>` | no |
| `AfterContactWorkConfigs` | `List<AfterContactWorkConfigPerChannel>` | no |
| `PhoneNumberConfigs` | `List<PhoneNumberConfig>` | no |
| `PersistentConnectionConfigs` | `List<PersistentConnectionConfig>` | no |
| `VoiceEnhancementConfigs` | `List<VoiceEnhancementConfig>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `UserArn` | `string` | no |

## CreateUserHierarchyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `ParentGroupId` | `string` | no |
| `InstanceId` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyGroupId` | `string` | no |
| `HierarchyGroupArn` | `string` | no |

## CreateView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `Status` | `string` | yes |
| `Content` | `ViewInputContent` | yes |
| `Description` | `string` | no |
| `Name` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## CreateViewVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |
| `VersionDescription` | `string` | no |
| `ViewContentSha256` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## CreateVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `VocabularyName` | `string` | yes |
| `LanguageCode` | `string` | yes |
| `Content` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyArn` | `string` | yes |
| `VocabularyId` | `string` | yes |
| `State` | `string` | yes |

## CreateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `Theme` | `WorkspaceTheme` | no |
| `Title` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkspaceId` | `string` | yes |
| `WorkspaceArn` | `string` | yes |

## CreateWorkspacePage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `ResourceArn` | `string` | yes |
| `Page` | `string` | yes |
| `Slug` | `string` | no |
| `InputData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeactivateEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormArn` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

## DeleteAttachedFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FileId` | `string` | yes |
| `AssociatedResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ContactFields` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactFlowModule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactFlowModuleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactFlowModuleVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `ContactFlowModuleVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteContactFlowVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `ContactFlowVersion` | `long` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDataTableAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `AttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LockVersion` | `DataTableLockVersion` | yes |

## DeleteEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EmailAddressId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteExtractionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ExtractionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHoursOfOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteHoursOfOperationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `HoursOfOperationOverrideId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteIntegrationAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationAssociationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MetricId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NotificationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePredefinedAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PromptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeletePushNotificationRegistration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RegistrationId` | `string` | yes |
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteQuickConnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QuickConnectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRoutingProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SecurityProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSession

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SessionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTaskTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TaskTemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteTrafficDistributionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUseCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationAssociationId` | `string` | yes |
| `UseCaseId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteUserHierarchyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyGroupId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteViewVersion

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |
| `ViewVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `VocabularyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularyArn` | `string` | yes |
| `VocabularyId` | `string` | yes |
| `State` | `string` | yes |

## DeleteWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspaceMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `MediaType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteWorkspacePage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Page` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DescribeAgentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AgentStatusId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentStatus` | `AgentStatus` | no |

## DescribeAttachedFilesConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AttachmentScope` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedFilesConfiguration` | `AttachedFilesConfiguration` | yes |

## DescribeAuthenticationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfile` | `AuthenticationProfile` | no |

## DescribeContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Contact` | `Contact` | no |

## DescribeContactEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Evaluation` | `Evaluation` | yes |
| `EvaluationForm` | `EvaluationFormContent` | yes |

## DescribeContactFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlow` | `ContactFlow` | no |

## DescribeContactFlowModule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModule` | `ContactFlowModule` | no |

## DescribeContactFlowModuleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `AliasId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModuleAlias` | `ContactFlowModuleAliasInfo` | no |

## DescribeDataTable

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataTable` | `DataTable` | yes |

## DescribeDataTableAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `AttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `DataTableAttribute` | yes |

## DescribeEmailAddress

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EmailAddressId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddressId` | `string` | no |
| `EmailAddressArn` | `string` | no |
| `EmailAddress` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `CreateTimestamp` | `string` | no |
| `ModifiedTimestamp` | `string` | no |
| `AliasConfigurations` | `List<AliasConfiguration>` | no |
| `Tags` | `Map<string>` | no |

## DescribeEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationForm` | `EvaluationForm` | yes |

## DescribeExtractionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ExtractionDefinitionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtractionDefinition` | `ExtractionDefinition` | yes |

## DescribeHoursOfOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperation` | `HoursOfOperation` | no |

## DescribeHoursOfOperationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `HoursOfOperationOverrideId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperationOverride` | `HoursOfOperationOverride` | no |

## DescribeInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Instance` | `Instance` | no |
| `ReplicationConfiguration` | `ReplicationConfiguration` | no |

## DescribeInstanceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AttributeType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attribute` | `Attribute` | no |

## DescribeInstanceStorageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageConfig` | `InstanceStorageConfig` | no |

## DescribeMetric

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MetricId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metric` | `MetricDefinition` | yes |

## DescribeNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NotificationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Notification` | `Notification` | yes |

## DescribePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClaimedPhoneNumberSummary` | `ClaimedPhoneNumberSummary` | no |

## DescribePredefinedAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredefinedAttribute` | `PredefinedAttribute` | no |

## DescribePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PromptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Prompt` | `Prompt` | no |

## DescribeQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queue` | `Queue` | no |

## DescribeQuickConnect

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QuickConnectId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuickConnect` | `QuickConnect` | no |

## DescribeRoutingProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingProfile` | `RoutingProfile` | no |

## DescribeRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RuleId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rule` | `Rule` | yes |

## DescribeSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfile` | `SecurityProfile` | no |

## DescribeTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `Status` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCase` | `TestCase` | no |

## DescribeTrafficDistributionGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroupId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroup` | `TrafficDistributionGroup` | no |

## DescribeUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `User` | `User` | no |

## DescribeUserHierarchyGroup

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyGroupId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyGroup` | `HierarchyGroup` | no |

## DescribeUserHierarchyStructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyStructure` | `HierarchyStructure` | no |

## DescribeView

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## DescribeVocabulary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `VocabularyId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Vocabulary` | `Vocabulary` | yes |

## DescribeWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Workspace` | `Workspace` | yes |

## DisassociateAnalyticsDataSet

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataSetId` | `string` | yes |
| `TargetAccountId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateApprovedOrigin

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Origin` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `LexBot` | `LexBot` | no |
| `LexV2Bot` | `LexV2Bot` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateEmailAddressAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddressId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AliasConfiguration` | `AliasConfiguration` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `ParentHoursOfOperationIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateInstanceStorageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateLambdaFunction

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FunctionArn` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateLexBot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `BotName` | `string` | yes |
| `LexRegion` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociatePhoneNumberContactFlow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateQueueEmailAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `EmailAddressesId` | `List<string>` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateQueueQuickConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `QuickConnectIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateRoutingProfileQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `QueueReferences` | `List<RoutingProfileQueueReference>` | no |
| `ManualAssignmentQueueReferences` | `List<RoutingProfileQueueReference>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSecurityKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SecurityProfiles` | `List<SecurityProfileItem>` | yes |
| `EntityType` | `string` | yes |
| `EntityArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateTrafficDistributionGroupUser

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroupId` | `string` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateUserProficiencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `UserId` | `string` | yes |
| `UserProficiencies` | `List<UserProficiencyDisassociate>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DisassociateWorkspace

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `ResourceArns` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SuccessfulList` | `List<SuccessfulBatchAssociationSummary>` | no |
| `FailedList` | `List<FailedBatchAssociationSummary>` | no |

## DismissUserContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## EvaluateDataTableValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Values` | `List<DataTableValueEvaluationSet>` | yes |
| `TimeZone` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Values` | `List<DataTableEvaluatedValue>` | yes |
| `NextToken` | `string` | no |

## GetAttachedFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FileId` | `string` | yes |
| `UrlExpiryInSeconds` | `integer` | no |
| `AssociatedResourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileArn` | `string` | no |
| `FileId` | `string` | no |
| `CreationTime` | `string` | no |
| `FileStatus` | `string` | no |
| `FileName` | `string` | no |
| `FileSizeInBytes` | `long` | yes |
| `AssociatedResourceArn` | `string` | no |
| `FileUseCaseType` | `string` | no |
| `CreatedBy` | `CreatedByInfo` | no |
| `DownloadUrlMetadata` | `DownloadUrlMetadata` | no |
| `Tags` | `Map<string>` | no |

## GetContactAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `InitialContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |

## GetContactMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `Metrics` | `List<ContactMetricInfo>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricResults` | `List<ContactMetricResult>` | no |
| `Id` | `string` | no |
| `Arn` | `string` | no |

## GetCrossRegionRouting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IsolatedRegions` | `List<string>` | no |

## GetCurrentMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Filters` | `Filters` | yes |
| `Groupings` | `List<string>` | no |
| `CurrentMetrics` | `List<CurrentMetric>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SortCriteria` | `List<CurrentMetricSortCriteria>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MetricResults` | `List<CurrentMetricResult>` | no |
| `DataSnapshotTime` | `timestamp` | no |
| `ApproximateTotalCount` | `long` | no |

## GetCurrentUserData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Filters` | `UserDataFilters` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `UserDataList` | `List<UserData>` | no |
| `ApproximateTotalCount` | `long` | no |

## GetEffectiveHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `FromDate` | `string` | yes |
| `ToDate` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EffectiveHoursOfOperationList` | `List<EffectiveHoursOfOperations>` | no |
| `EffectiveOverrideHoursList` | `List<EffectiveOverrideHours>` | no |
| `TimeZone` | `string` | no |

## GetEvaluationFormValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | yes |
| `FailureReason` | `string` | no |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |
| `StartedTime` | `timestamp` | yes |
| `Findings` | `List<EvaluationFormValidationFinding>` | no |

## GetFederationToken

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | no |
| `UserArn` | `string` | no |
| `Credentials` | `Credentials` | no |
| `SignInUrl` | `string` | no |

## GetFlowAssociation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceId` | `string` | yes |
| `ResourceType` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceId` | `string` | no |
| `FlowId` | `string` | no |
| `ResourceType` | `string` | no |

## GetMetricData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Filters` | `Filters` | yes |
| `Groupings` | `List<string>` | no |
| `HistoricalMetrics` | `List<HistoricalMetric>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MetricResults` | `List<HistoricalMetricResult>` | no |

## GetMetricDataV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ResourceArn` | `string` | yes |
| `StartTime` | `timestamp` | yes |
| `EndTime` | `timestamp` | yes |
| `Interval` | `IntervalDetails` | no |
| `Filters` | `List<FilterV2>` | yes |
| `Groupings` | `List<string>` | no |
| `Metrics` | `List<MetricV2>` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MetricResults` | `List<MetricResultV2>` | no |

## GetPromptFile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PromptId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PromptPresignedUrl` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## GetTaskTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TaskTemplateId` | `string` | yes |
| `SnapshotVersion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `Id` | `string` | yes |
| `Arn` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ContactFlowId` | `string` | no |
| `SelfAssignFlowId` | `string` | no |
| `Constraints` | `TaskTemplateConstraints` | no |
| `Defaults` | `TaskTemplateDefaults` | no |
| `Fields` | `List<TaskTemplateField>` | no |
| `Status` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreatedTime` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetTestCaseExecutionSummary

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `TestCaseExecutionId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StartTime` | `timestamp` | no |
| `EndTime` | `timestamp` | no |
| `Status` | `string` | no |
| `ObservationSummary` | `ObservationSummary` | no |

## GetTrafficDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TelephonyConfig` | `TelephonyConfig` | no |
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `SignInConfig` | `SignInConfig` | no |
| `AgentConfig` | `AgentConfig` | no |

## ImportPhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `SourcePhoneNumberArn` | `string` | yes |
| `PhoneNumberDescription` | `string` | no |
| `Tags` | `Map<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | no |
| `PhoneNumberArn` | `string` | no |

## ImportWorkspaceMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `MediaType` | `string` | yes |
| `MediaSource` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ListAgentStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `AgentStatusTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AgentStatusSummaryList` | `List<AgentStatusSummary>` | no |

## ListAnalyticsDataAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataSetId` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<AnalyticsDataAssociationResult>` | no |
| `NextToken` | `string` | no |

## ListAnalyticsDataLakeDataSets

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Results` | `List<AnalyticsDataSetsResult>` | no |
| `NextToken` | `string` | no |

## ListApprovedOrigins

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Origins` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListAssociatedContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactSummaryList` | `List<AssociatedContactSummary>` | no |
| `NextToken` | `string` | no |

## ListAttachedFilesConfigurations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AttachedFilesConfigurations` | `List<AttachedFilesConfigurationSummary>` | no |
| `NextToken` | `string` | no |

## ListAuthenticationProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileSummaryList` | `List<AuthenticationProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `LexVersion` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LexBots` | `List<LexBotConfig>` | no |
| `NextToken` | `string` | no |

## ListChildHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ChildHoursOfOperationsSummaryList` | `List<HoursOfOperationsIdentifier>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListContactEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationSummaryList` | `List<EvaluationSummary>` | yes |
| `NextToken` | `string` | no |

## ListContactFlowModuleAliases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModuleAliasSummaryList` | `List<ContactFlowModuleAliasSummary>` | no |
| `NextToken` | `string` | no |

## ListContactFlowModuleVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModuleVersionSummaryList` | `List<ContactFlowModuleVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListContactFlowModules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ContactFlowModuleState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModulesSummaryList` | `List<ContactFlowModuleSummary>` | no |
| `NextToken` | `string` | no |

## ListContactFlowVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowVersionSummaryList` | `List<ContactFlowVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListContactFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowSummaryList` | `List<ContactFlowSummary>` | no |
| `NextToken` | `string` | no |

## ListContactReferences

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ReferenceTypes` | `List<string>` | yes |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ReferenceSummaryList` | `List<ReferenceSummary>` | no |
| `NextToken` | `string` | no |

## ListDataTableAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `AttributeIds` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Attributes` | `List<DataTableAttribute>` | yes |

## ListDataTablePrimaryValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `RecordIds` | `List<string>` | no |
| `PrimaryAttributeValues` | `List<PrimaryAttributeValueFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PrimaryValuesList` | `List<RecordPrimaryValue>` | yes |

## ListDataTableValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `RecordIds` | `List<string>` | no |
| `PrimaryAttributeValues` | `List<PrimaryAttributeValueFilter>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Values` | `List<DataTableValueSummary>` | yes |

## ListDataTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `DataTableSummaryList` | `List<DataTableSummary>` | yes |

## ListDefaultVocabularies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `LanguageCode` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DefaultVocabularyList` | `List<DefaultVocabulary>` | yes |
| `NextToken` | `string` | no |

## ListEntitySecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EntityType` | `string` | yes |
| `EntityArn` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfiles` | `List<SecurityProfileItem>` | no |
| `NextToken` | `string` | no |

## ListEvaluationFormVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormVersionSummaryList` | `List<EvaluationFormVersionSummary>` | yes |
| `NextToken` | `string` | no |

## ListEvaluationForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormSummaryList` | `List<EvaluationFormSummary>` | yes |
| `NextToken` | `string` | no |

## ListExtractionDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExtractionDefinitionSummaryList` | `List<ExtractionDefinitionSummary>` | yes |
| `NextToken` | `string` | no |

## ListFlowAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FlowAssociationSummaryList` | `List<FlowAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListHoursOfOperationOverrides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `HoursOfOperationOverrideList` | `List<HoursOfOperationOverride>` | no |
| `LastModifiedRegion` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |

## ListHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperationSummaryList` | `List<HoursOfOperationSummary>` | no |
| `NextToken` | `string` | no |

## ListInstanceAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `List<Attribute>` | no |
| `NextToken` | `string` | no |

## ListInstanceStorageConfigs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StorageConfigs` | `List<InstanceStorageConfig>` | no |
| `NextToken` | `string` | no |

## ListInstances

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceSummaryList` | `List<InstanceSummary>` | no |
| `NextToken` | `string` | no |

## ListIntegrationAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationType` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IntegrationArn` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IntegrationAssociationSummaryList` | `List<IntegrationAssociationSummary>` | no |
| `NextToken` | `string` | no |

## ListLambdaFunctions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LambdaFunctions` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListLexBots

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LexBots` | `List<LexBot>` | no |
| `NextToken` | `string` | no |

## ListMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Type` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MetricSummaryList` | `List<MetricSummary>` | yes |
| `NextToken` | `string` | no |

## ListNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `NotificationSummaryList` | `List<Notification>` | yes |

## ListPhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PhoneNumberTypes` | `List<string>` | no |
| `PhoneNumberCountryCodes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberSummaryList` | `List<PhoneNumberSummary>` | no |
| `NextToken` | `string` | no |

## ListPhoneNumbersV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | no |
| `InstanceId` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `PhoneNumberCountryCodes` | `List<string>` | no |
| `PhoneNumberTypes` | `List<string>` | no |
| `PhoneNumberPrefix` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `ListPhoneNumbersSummaryList` | `List<ListPhoneNumbersSummary>` | no |

## ListPredefinedAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `PredefinedAttributeSummaryList` | `List<PredefinedAttributeSummary>` | no |

## ListPrompts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PromptSummaryList` | `List<PromptSummary>` | no |
| `NextToken` | `string` | no |

## ListQueueEmailAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `EmailAddressMetadataList` | `List<EmailAddressSummary>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListQueueQuickConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `QuickConnectSummaryList` | `List<QuickConnectSummary>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QueueSummaryList` | `List<QueueSummary>` | no |
| `NextToken` | `string` | no |

## ListQuickConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `QuickConnectTypes` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuickConnectSummaryList` | `List<QuickConnectSummary>` | no |
| `NextToken` | `string` | no |

## ListRealtimeContactAnalysisSegmentsV2

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `OutputType` | `string` | yes |
| `SegmentTypes` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Channel` | `string` | yes |
| `Status` | `string` | yes |
| `Segments` | `List<RealtimeContactAnalysisSegment>` | yes |
| `NextToken` | `string` | no |

## ListRoutingProfileManualAssignmentQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RoutingProfileManualAssignmentQueueConfigSummaryList` | `List<RoutingProfileManualAssignmentQueueConfigSummary>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListRoutingProfileQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RoutingProfileQueueConfigSummaryList` | `List<RoutingProfileQueueConfigSummary>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListRoutingProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingProfileSummaryList` | `List<RoutingProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PublishStatus` | `string` | no |
| `EventSourceName` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleSummaryList` | `List<RuleSummary>` | yes |
| `NextToken` | `string` | no |

## ListSecurityKeys

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityKeys` | `List<SecurityKey>` | no |
| `NextToken` | `string` | no |

## ListSecurityProfileApplications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Applications` | `List<Application>` | no |
| `NextToken` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListSecurityProfileFlowModules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AllowedFlowModules` | `List<FlowModule>` | no |
| `NextToken` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListSecurityProfilePermissions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Permissions` | `List<string>` | no |
| `NextToken` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileSummaryList` | `List<SecurityProfileSummary>` | no |
| `NextToken` | `string` | no |

## ListTagsForResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `tags` | `Map<string>` | no |

## ListTaskTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `Status` | `string` | no |
| `Name` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskTemplates` | `List<TaskTemplateMetadata>` | no |
| `NextToken` | `string` | no |

## ListTestCaseExecutionRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `TestCaseExecutionId` | `string` | yes |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ExecutionRecords` | `List<ExecutionRecord>` | no |
| `NextToken` | `string` | no |

## ListTestCaseExecutions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | no |
| `TestCaseName` | `string` | no |
| `StartTime` | `long` | no |
| `EndTime` | `long` | no |
| `Status` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCaseExecutions` | `List<TestCaseExecution>` | no |
| `NextToken` | `string` | no |

## ListTestCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCaseSummaryList` | `List<TestCaseSummary>` | no |
| `NextToken` | `string` | no |

## ListTrafficDistributionGroupUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TrafficDistributionGroupId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TrafficDistributionGroupUserSummaryList` | `List<TrafficDistributionGroupUserSummary>` | no |

## ListTrafficDistributionGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `InstanceId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `TrafficDistributionGroupSummaryList` | `List<TrafficDistributionGroupSummary>` | no |

## ListUseCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IntegrationAssociationId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UseCaseSummaryList` | `List<UseCase>` | no |
| `NextToken` | `string` | no |

## ListUserHierarchyGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserHierarchyGroupSummaryList` | `List<HierarchyGroupSummary>` | no |
| `NextToken` | `string` | no |

## ListUserNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `UserId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserNotifications` | `List<UserNotificationSummary>` | no |
| `NextToken` | `string` | no |

## ListUserProficiencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `UserId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `UserProficiencyList` | `List<UserProficiency>` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

## ListUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserSummaryList` | `List<UserSummary>` | no |
| `NextToken` | `string` | no |

## ListViewVersions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewVersionSummaryList` | `List<ViewVersionSummary>` | no |
| `NextToken` | `string` | no |

## ListViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Type` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ViewsSummaryList` | `List<ViewSummary>` | no |
| `NextToken` | `string` | no |

## ListWorkspaceMedia

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Media` | `List<MediaItem>` | no |

## ListWorkspacePages

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WorkspacePageList` | `List<WorkspacePage>` | yes |

## ListWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WorkspaceSummaryList` | `List<WorkspaceSummary>` | yes |

## MonitorContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `UserId` | `string` | yes |
| `AllowedMonitorCapabilities` | `List<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `ContactArn` | `string` | no |

## PauseContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## PutUserStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AgentStatusId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReleasePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ReplicateInstance

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ReplicaRegion` | `string` | yes |
| `ClientToken` | `string` | no |
| `ReplicaAlias` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | no |
| `Arn` | `string` | no |

## ResumeContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## ResumeContactRecording

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `InitialContactId` | `string` | yes |
| `ContactRecordingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SearchAgentStatuses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `AgentStatusSearchFilter` | no |
| `SearchCriteria` | `AgentStatusSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AgentStatuses` | `List<AgentStatus>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchAvailablePhoneNumbers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TargetArn` | `string` | no |
| `InstanceId` | `string` | no |
| `PhoneNumberCountryCode` | `string` | yes |
| `PhoneNumberType` | `string` | yes |
| `PhoneNumberPrefix` | `string` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `AvailableNumbersList` | `List<AvailableNumberSummary>` | no |

## SearchContactEvaluations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchCriteria` | `EvaluationSearchCriteria` | no |
| `SearchFilter` | `EvaluationSearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationSearchSummaryList` | `List<EvaluationSearchSummary>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchContactFlowModules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `ContactFlowModuleSearchFilter` | no |
| `SearchCriteria` | `ContactFlowModuleSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlowModules` | `List<ContactFlowModule>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchContactFlows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `ContactFlowSearchFilter` | no |
| `SearchCriteria` | `ContactFlowSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactFlows` | `List<ContactFlow>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TimeRange` | `SearchContactsTimeRange` | yes |
| `SearchCriteria` | `SearchCriteria` | no |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `Sort` | `Sort` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Contacts` | `List<ContactSearchSummary>` | yes |
| `NextToken` | `string` | no |
| `TotalCount` | `long` | no |

## SearchDataTables

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `DataTableSearchFilter` | no |
| `SearchCriteria` | `DataTableSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DataTables` | `List<DataTable>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchEmailAddresses

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SearchCriteria` | `EmailAddressSearchCriteria` | no |
| `SearchFilter` | `EmailAddressSearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `EmailAddresses` | `List<EmailAddressMetadata>` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchEvaluationForms

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchCriteria` | `EvaluationFormSearchCriteria` | no |
| `SearchFilter` | `EvaluationFormSearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormSearchSummaryList` | `List<EvaluationFormSearchSummary>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchHoursOfOperationOverrides

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `HoursOfOperationSearchFilter` | no |
| `SearchCriteria` | `HoursOfOperationOverrideSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperationOverrides` | `List<HoursOfOperationOverride>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchHoursOfOperations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `HoursOfOperationSearchFilter` | no |
| `SearchCriteria` | `HoursOfOperationSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HoursOfOperations` | `List<HoursOfOperation>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchMetrics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `MetricSearchFilter` | no |
| `SearchCriteria` | `MetricSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Metrics` | `List<MetricDefinition>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchNotifications

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `NotificationSearchFilter` | no |
| `SearchCriteria` | `NotificationSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Notifications` | `List<NotificationSearchSummary>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchPredefinedAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchCriteria` | `PredefinedAttributeSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PredefinedAttributes` | `List<PredefinedAttribute>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchPrompts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `PromptSearchFilter` | no |
| `SearchCriteria` | `PromptSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Prompts` | `List<Prompt>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `QueueSearchFilter` | no |
| `SearchCriteria` | `QueueSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Queues` | `List<Queue>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchQuickConnects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `QuickConnectSearchFilter` | no |
| `SearchCriteria` | `QuickConnectSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `QuickConnects` | `List<QuickConnect>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchResourceTags

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ResourceTypes` | `List<string>` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchCriteria` | `ResourceTagsSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Tags` | `List<TagSet>` | no |
| `NextToken` | `string` | no |

## SearchRoutingProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `RoutingProfileSearchFilter` | no |
| `SearchCriteria` | `RoutingProfileSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingProfiles` | `List<RoutingProfile>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchRules

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `SearchCriteria` | `RulesSearchCriteria` | no |
| `SearchFilter` | `RulesSearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Rules` | `List<RuleSearchSummary>` | yes |
| `ApproximateTotalCount` | `long` | no |
| `NextToken` | `string` | no |

## SearchSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchCriteria` | `SecurityProfileSearchCriteria` | no |
| `SearchFilter` | `SecurityProfilesSearchFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfiles` | `List<SecurityProfileSearchSummary>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchTestCases

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `TestCaseSearchFilter` | no |
| `SearchCriteria` | `TestCaseSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCases` | `List<TestCase>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchUserHierarchyGroups

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `UserHierarchyGroupSearchFilter` | no |
| `SearchCriteria` | `UserHierarchyGroupSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `UserHierarchyGroups` | `List<HierarchyGroup>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchUsers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `UserSearchFilter` | no |
| `SearchCriteria` | `UserSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Users` | `List<UserSearchSummary>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchViews

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `ViewSearchFilter` | no |
| `SearchCriteria` | `ViewSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Views` | `List<View>` | no |
| `NextToken` | `string` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchVocabularies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |
| `State` | `string` | no |
| `NameStartsWith` | `string` | no |
| `LanguageCode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `VocabularySummaryList` | `List<VocabularySummary>` | no |
| `NextToken` | `string` | no |

## SearchWorkspaceAssociations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `WorkspaceAssociationSearchFilter` | no |
| `SearchCriteria` | `WorkspaceAssociationSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `WorkspaceAssociations` | `List<WorkspaceAssociationSearchSummary>` | no |
| `ApproximateTotalCount` | `long` | no |

## SearchWorkspaces

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `SearchFilter` | `WorkspaceSearchFilter` | no |
| `SearchCriteria` | `WorkspaceSearchCriteria` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Workspaces` | `List<WorkspaceSearchSummary>` | no |
| `ApproximateTotalCount` | `long` | no |

## SendChatIntegrationEvent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceId` | `string` | yes |
| `DestinationId` | `string` | yes |
| `Subtype` | `string` | no |
| `Event` | `ChatEvent` | yes |
| `NewSessionDetails` | `NewSessionDetails` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InitialContactId` | `string` | no |
| `NewChatCreated` | `boolean` | no |

## SendOutboundEmail

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FromEmailAddress` | `EmailAddressInfo` | yes |
| `DestinationEmailAddress` | `EmailAddressInfo` | yes |
| `AdditionalRecipients` | `OutboundAdditionalRecipients` | no |
| `EmailMessage` | `OutboundEmailContent` | yes |
| `TrafficType` | `string` | yes |
| `SourceCampaign` | `SourceCampaign` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SendOutboundWebNotification

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `BrowserId` | `string` | yes |
| `SessionId` | `string` | yes |
| `ExpiresAt` | `timestamp` | yes |
| `Source` | `WebNotificationSource` | yes |
| `Destination` | `WidgetDestination` | yes |
| `Content` | `WebNotificationContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartAssistantContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AiAgent` | `AiAgentInput` | yes |
| `ParticipantDetails` | `ParticipantDetails` | yes |
| `InitialMessage` | `ChatMessage` | no |
| `Attributes` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `PersistentChat` | `PersistentChat` | no |
| `RelatedContactId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `ParticipantId` | `string` | no |
| `ParticipantToken` | `string` | no |
| `ContinuedFromContactId` | `string` | no |

## StartAttachedFileUpload

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `FileName` | `string` | yes |
| `FileSizeInBytes` | `long` | yes |
| `UrlExpiryInSeconds` | `integer` | no |
| `FileUseCaseType` | `string` | yes |
| `AssociatedResourceArn` | `string` | yes |
| `CreatedBy` | `CreatedByInfo` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `FileArn` | `string` | no |
| `FileId` | `string` | no |
| `CreationTime` | `string` | no |
| `FileStatus` | `string` | no |
| `CreatedBy` | `CreatedByInfo` | no |
| `UploadUrlMetadata` | `UploadUrlMetadata` | no |

## StartChatContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `Attributes` | `Map<string>` | no |
| `ParticipantDetails` | `ParticipantDetails` | yes |
| `ParticipantConfiguration` | `ParticipantConfiguration` | no |
| `InitialMessage` | `ChatMessage` | no |
| `ClientToken` | `string` | no |
| `ChatDurationInMinutes` | `integer` | no |
| `SupportedMessagingContentTypes` | `List<string>` | no |
| `PersistentChat` | `PersistentChat` | no |
| `RelatedContactId` | `string` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |
| `CustomerId` | `string` | no |
| `DisconnectOnCustomerExit` | `List<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `ParticipantId` | `string` | no |
| `ParticipantToken` | `string` | no |
| `ContinuedFromContactId` | `string` | no |

## StartContactConversationalAnalyticsJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `AnalyticsModes` | `List<string>` | yes |
| `AnalyticsConfiguration` | `AnalyticsConfiguration` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `ContactId` | `string` | no |

## StartContactEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `AutoEvaluationConfiguration` | `AutoEvaluationConfiguration` | no |
| `ClientToken` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |
| `EvaluationArn` | `string` | yes |

## StartContactMediaProcessing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `ContactId` | `string` | no |
| `ProcessorArn` | `string` | no |
| `FailureMode` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartContactRecording

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `InitialContactId` | `string` | yes |
| `VoiceRecordingConfiguration` | `VoiceRecordingConfiguration` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartContactStreaming

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ChatStreamingConfiguration` | `ChatStreamingConfiguration` | yes |
| `ClientToken` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `StreamingId` | `string` | yes |

## StartEmailContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `FromEmailAddress` | `EmailAddressInfo` | yes |
| `DestinationEmailAddress` | `string` | yes |
| `Description` | `string` | no |
| `References` | `Map<Reference>` | no |
| `Name` | `string` | no |
| `EmailMessage` | `InboundEmailContent` | yes |
| `AdditionalRecipients` | `InboundAdditionalRecipients` | no |
| `Attachments` | `List<EmailAttachment>` | no |
| `ContactFlowId` | `string` | no |
| `RelatedContactId` | `string` | no |
| `Attributes` | `Map<string>` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |

## StartEvaluationFormValidation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormArn` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

## StartOutboundChatContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SourceEndpoint` | `Endpoint` | yes |
| `DestinationEndpoint` | `Endpoint` | yes |
| `InstanceId` | `string` | yes |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | yes |
| `Attributes` | `Map<string>` | no |
| `ContactFlowId` | `string` | yes |
| `ChatDurationInMinutes` | `integer` | no |
| `ParticipantDetails` | `ParticipantDetails` | no |
| `InitialSystemMessage` | `ChatMessage` | no |
| `InitialTemplatedSystemMessage` | `TemplatedMessageConfig` | no |
| `RelatedContactId` | `string` | no |
| `SupportedMessagingContentTypes` | `List<string>` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |

## StartOutboundEmailContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `FromEmailAddress` | `EmailAddressInfo` | no |
| `DestinationEmailAddress` | `EmailAddressInfo` | yes |
| `AdditionalRecipients` | `OutboundAdditionalRecipients` | no |
| `EmailMessage` | `OutboundEmailContent` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |

## StartOutboundVoiceContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | no |
| `Description` | `string` | no |
| `References` | `Map<Reference>` | no |
| `RelatedContactId` | `string` | no |
| `DestinationPhoneNumber` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `ClientToken` | `string` | no |
| `SourcePhoneNumber` | `string` | no |
| `QueueId` | `string` | no |
| `Attributes` | `Map<string>` | no |
| `AnswerMachineDetectionConfig` | `AnswerMachineDetectionConfig` | no |
| `CampaignId` | `string` | no |
| `TrafficType` | `string` | no |
| `OutboundStrategy` | `OutboundStrategy` | no |
| `RingTimeoutInSeconds` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |

## StartScreenSharing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartTaskContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PreviousContactId` | `string` | no |
| `ContactFlowId` | `string` | no |
| `Attributes` | `Map<string>` | no |
| `Name` | `string` | yes |
| `References` | `Map<Reference>` | no |
| `Description` | `string` | no |
| `ClientToken` | `string` | no |
| `ScheduledTime` | `timestamp` | no |
| `TaskTemplateId` | `string` | no |
| `QuickConnectId` | `string` | no |
| `RelatedContactId` | `string` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |
| `Attachments` | `List<TaskAttachment>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |

## StartTestCaseExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TestCaseExecutionId` | `string` | no |
| `TestCaseId` | `string` | no |
| `Status` | `string` | no |

## StartWebRTCContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Attributes` | `Map<string>` | no |
| `ClientToken` | `string` | no |
| `ContactFlowId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AllowedCapabilities` | `AllowedCapabilities` | no |
| `ParticipantDetails` | `ParticipantDetails` | yes |
| `RelatedContactId` | `string` | no |
| `References` | `Map<Reference>` | no |
| `Description` | `string` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ConnectionData` | `ConnectionData` | no |
| `ContactId` | `string` | no |
| `ParticipantId` | `string` | no |
| `ParticipantToken` | `string` | no |

## StopContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `DisconnectReason` | `DisconnectReason` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopContactMediaProcessing

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `ContactId` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopContactRecording

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `InitialContactId` | `string` | yes |
| `ContactRecordingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopContactStreaming

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `StreamingId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopTestCaseExecution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseExecutionId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## SubmitContactEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationId` | `string` | yes |
| `Answers` | `Map<EvaluationAnswerInput>` | no |
| `Notes` | `Map<EvaluationNote>` | no |
| `SubmittedBy` | `EvaluatorUserUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |
| `EvaluationArn` | `string` | yes |

## SuspendContactRecording

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `InitialContactId` | `string` | yes |
| `ContactRecordingType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tags` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## TransferContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `QueueId` | `string` | no |
| `UserId` | `string` | no |
| `ContactFlowId` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | no |
| `ContactArn` | `string` | no |

## UntagContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `TagKeys` | `List<string>` | yes |

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


## UpdateAgentStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AgentStatusId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |
| `DisplayOrder` | `integer` | no |
| `ResetOrderNumber` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateAttachedFilesConfiguration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AttachmentScope` | `string` | yes |
| `MaximumSizeLimitInBytes` | `long` | no |
| `ExtensionConfiguration` | `ExtensionConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AttachmentScope` | `string` | yes |
| `MaximumSizeLimitInBytes` | `long` | no |
| `ExtensionConfiguration` | `ExtensionConfiguration` | no |
| `LastModifiedTime` | `timestamp` | no |

## UpdateAuthenticationProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AuthenticationProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `AllowedIps` | `List<string>` | no |
| `BlockedIps` | `List<string>` | no |
| `PeriodicSessionDuration` | `integer` | no |
| `SessionInactivityDuration` | `integer` | no |
| `SessionInactivityHandlingEnabled` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContact

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `References` | `Map<Reference>` | no |
| `SegmentAttributes` | `Map<SegmentAttributeValue>` | no |
| `QueueInfo` | `QueueInfoInput` | no |
| `UserInfo` | `UserInfo` | no |
| `CustomerEndpoint` | `Endpoint` | no |
| `SystemEndpoint` | `Endpoint` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InitialContactId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Attributes` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactEvaluation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationId` | `string` | yes |
| `Answers` | `Map<EvaluationAnswerInput>` | no |
| `Notes` | `Map<EvaluationNote>` | no |
| `UpdatedBy` | `EvaluatorUserUnion` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationId` | `string` | yes |
| `EvaluationArn` | `string` | yes |

## UpdateContactFlowContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `Content` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactFlowMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ContactFlowState` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactFlowModuleAlias

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `AliasId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ContactFlowModuleVersion` | `long` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactFlowModuleContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `Content` | `string` | no |
| `Settings` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactFlowModuleMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowModuleId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `State` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactFlowName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactFlowId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactRoutingData

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `QueueTimeAdjustmentSeconds` | `integer` | no |
| `QueuePriority` | `long` | no |
| `RoutingCriteria` | `RoutingCriteriaInput` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactSchedule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ScheduledTime` | `timestamp` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateContactTaskTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TaskTemplateId` | `string` | yes |
| `ContactId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCrossRegionRouting

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `IsolatedAll` | `boolean` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateDataTableAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `AttributeName` | `string` | yes |
| `Name` | `string` | yes |
| `ValueType` | `string` | yes |
| `Description` | `string` | no |
| `Primary` | `boolean` | no |
| `Validation` | `Validation` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `LockVersion` | `DataTableLockVersion` | yes |

## UpdateDataTableMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `Name` | `string` | yes |
| `Description` | `string` | no |
| `ValueLockLevel` | `string` | yes |
| `TimeZone` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LockVersion` | `DataTableLockVersion` | yes |

## UpdateDataTablePrimaryValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `DataTableId` | `string` | yes |
| `PrimaryValues` | `List<PrimaryValue>` | yes |
| `NewPrimaryValues` | `List<PrimaryValue>` | yes |
| `LockVersion` | `DataTableLockVersion` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LockVersion` | `DataTableLockVersion` | yes |

## UpdateEmailAddressMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EmailAddressId` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EmailAddressId` | `string` | no |
| `EmailAddressArn` | `string` | no |

## UpdateEvaluationForm

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |
| `CreateNewVersion` | `boolean` | no |
| `Title` | `string` | yes |
| `Description` | `string` | no |
| `Items` | `List<EvaluationFormItem>` | yes |
| `ScoringStrategy` | `EvaluationFormScoringStrategy` | no |
| `AutoEvaluationConfiguration` | `EvaluationFormAutoEvaluationConfiguration` | no |
| `ReviewConfiguration` | `EvaluationReviewConfiguration` | no |
| `AsDraft` | `boolean` | no |
| `ClientToken` | `string` | no |
| `TargetConfiguration` | `EvaluationFormTargetConfiguration` | no |
| `LanguageConfiguration` | `EvaluationFormLanguageConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EvaluationFormId` | `string` | yes |
| `EvaluationFormArn` | `string` | yes |
| `EvaluationFormVersion` | `integer` | yes |

## UpdateExtractionDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ClientToken` | `string` | no |
| `ExtractionDefinitionId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `ExtractionConfiguration` | `ExtractionConfiguration` | yes |
| `Display` | `ExtractionDefinitionDisplay` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHoursOfOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `TimeZone` | `string` | no |
| `Config` | `List<HoursOfOperationConfig>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateHoursOfOperationOverride

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |
| `HoursOfOperationOverrideId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Config` | `List<HoursOfOperationOverrideConfig>` | no |
| `EffectiveFrom` | `string` | no |
| `EffectiveTill` | `string` | no |
| `RecurrenceConfig` | `RecurrenceConfig` | no |
| `OverrideType` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInstanceAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AttributeType` | `string` | yes |
| `Value` | `string` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateInstanceStorageConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `AssociationId` | `string` | yes |
| `ResourceType` | `string` | yes |
| `StorageConfig` | `InstanceStorageConfig` | yes |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMetricContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MetricId` | `string` | yes |
| `MetricCalculation` | `MetricCalculation` | no |
| `Unit` | `string` | no |
| `PositiveTrendIndicator` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateMetricMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `MetricId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateNotificationContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NotificationId` | `string` | yes |
| `Content` | `Map<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateParticipantAuthentication

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `State` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Code` | `string` | no |
| `Error` | `string` | no |
| `ErrorDescription` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateParticipantRoleConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ContactId` | `string` | yes |
| `ChannelConfiguration` | `UpdateParticipantRoleConfigChannelInfo` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePhoneNumber

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `TargetArn` | `string` | no |
| `InstanceId` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | no |
| `PhoneNumberArn` | `string` | no |

## UpdatePhoneNumberMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneNumberId` | `string` | yes |
| `PhoneNumberDescription` | `string` | no |
| `ClientToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePredefinedAttribute

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Values` | `PredefinedAttributeValues` | no |
| `Purposes` | `List<string>` | no |
| `AttributeConfiguration` | `InputPredefinedAttributeConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdatePrompt

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `PromptId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `S3Uri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PromptARN` | `string` | no |
| `PromptId` | `string` | no |

## UpdateQueueHoursOfOperation

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `HoursOfOperationId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueMaxContacts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `MaxContacts` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueOutboundCallerConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `OutboundCallerConfig` | `OutboundCallerConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueOutboundEmailConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `OutboundEmailConfig` | `OutboundEmailConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQueueStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QueueId` | `string` | yes |
| `Status` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQuickConnectConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QuickConnectId` | `string` | yes |
| `QuickConnectConfig` | `QuickConnectConfig` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateQuickConnectName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `QuickConnectId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingProfileAgentAvailabilityTimer

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `AgentAvailabilityTimer` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingProfileConcurrency

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `MediaConcurrencies` | `List<MediaConcurrency>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingProfileDefaultOutboundQueue

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `DefaultOutboundQueueId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingProfileName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRoutingProfileQueues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `RoutingProfileId` | `string` | yes |
| `QueueConfigs` | `List<RoutingProfileQueueConfig>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateRule

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RuleId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Name` | `string` | yes |
| `Function` | `string` | yes |
| `Actions` | `List<RuleAction>` | yes |
| `PublishStatus` | `string` | yes |
| `PreEvaluationFilters` | `PreEvaluationFilters` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateSecurityProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Description` | `string` | no |
| `Permissions` | `List<string>` | no |
| `SecurityProfileId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `AllowedAccessControlTags` | `Map<string>` | no |
| `TagRestrictedResources` | `List<string>` | no |
| `Applications` | `List<Application>` | no |
| `HierarchyRestrictedResources` | `List<string>` | no |
| `AllowedAccessControlHierarchyGroupId` | `string` | no |
| `AllowedFlowModules` | `List<FlowModule>` | no |
| `GranularAccessControlConfiguration` | `GranularAccessControlConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTaskTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TaskTemplateId` | `string` | yes |
| `InstanceId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ContactFlowId` | `string` | no |
| `SelfAssignFlowId` | `string` | no |
| `Constraints` | `TaskTemplateConstraints` | no |
| `Defaults` | `TaskTemplateDefaults` | no |
| `Status` | `string` | no |
| `Fields` | `List<TaskTemplateField>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | no |
| `Id` | `string` | no |
| `Arn` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `ContactFlowId` | `string` | no |
| `SelfAssignFlowId` | `string` | no |
| `Constraints` | `TaskTemplateConstraints` | no |
| `Defaults` | `TaskTemplateDefaults` | no |
| `Fields` | `List<TaskTemplateField>` | no |
| `Status` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `CreatedTime` | `timestamp` | no |

## UpdateTestCase

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `TestCaseId` | `string` | yes |
| `Content` | `string` | no |
| `EntryPoint` | `TestCaseEntryPoint` | no |
| `InitializationData` | `string` | no |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateTrafficDistribution

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `TelephonyConfig` | `TelephonyConfig` | no |
| `SignInConfig` | `SignInConfig` | no |
| `AgentConfig` | `AgentConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AutoAcceptConfigs` | `List<AutoAcceptConfig>` | no |
| `AfterContactWorkConfigs` | `List<AfterContactWorkConfigPerChannel>` | no |
| `PhoneNumberConfigs` | `List<PhoneNumberConfig>` | no |
| `PersistentConnectionConfigs` | `List<PersistentConnectionConfig>` | no |
| `VoiceEnhancementConfigs` | `List<VoiceEnhancementConfig>` | no |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserHierarchy

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyGroupId` | `string` | no |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserHierarchyGroupName

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Name` | `string` | yes |
| `HierarchyGroupId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserHierarchyStructure

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `HierarchyStructure` | `HierarchyStructureUpdate` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserIdentityInfo

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityInfo` | `UserIdentityInfo` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserNotificationStatus

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `NotificationId` | `string` | yes |
| `UserId` | `string` | yes |
| `Status` | `string` | yes |
| `LastModifiedTime` | `timestamp` | no |
| `LastModifiedRegion` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserPhoneConfig

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `PhoneConfig` | `UserPhoneConfig` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserProficiencies

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `UserId` | `string` | yes |
| `UserProficiencies` | `List<UserProficiency>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserRoutingProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RoutingProfileId` | `string` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateUserSecurityProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SecurityProfileIds` | `List<string>` | yes |
| `UserId` | `string` | yes |
| `InstanceId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateViewContent

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |
| `Status` | `string` | yes |
| `Content` | `ViewInputContent` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `View` | `View` | no |

## UpdateViewMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `ViewId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceMetadata

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Name` | `string` | no |
| `Description` | `string` | no |
| `Title` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspacePage

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Page` | `string` | yes |
| `NewPage` | `string` | no |
| `ResourceArn` | `string` | no |
| `Slug` | `string` | no |
| `InputData` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceTheme

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Theme` | `WorkspaceTheme` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateWorkspaceVisibility

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `InstanceId` | `string` | yes |
| `WorkspaceId` | `string` | yes |
| `Visibility` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


