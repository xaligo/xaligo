# Amazon Connect Customer Profiles

API version: 2020-08-15. [Official AWS SDK source](https://github.com/boto/botocore/blob/develop/botocore/data/customer-profiles/2020-08-15/service-2.json).

These are AWS API fields, not XAL attributes. Nested type definitions are available in the linked SDK model. Availability must be verified separately.

## AddProfileKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `KeyName` | `string` | yes |
| `Values` | `List<string>` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `KeyName` | `string` | no |
| `Values` | `List<string>` | no |

## AssociateStreamForSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DestinationArn` | `string` | yes |
| `DestinationRoleArn` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## BatchGetCalculatedAttributeForProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatedAttributeName` | `string` | yes |
| `DomainName` | `string` | yes |
| `ProfileIds` | `List<string>` | yes |
| `ConditionOverrides` | `ConditionOverrides` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchGetCalculatedAttributeForProfileError>` | no |
| `CalculatedAttributeValues` | `List<CalculatedAttributeValue>` | no |
| `ConditionOverrides` | `ConditionOverrides` | no |

## BatchGetProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Errors` | `List<BatchGetProfileError>` | no |
| `Profiles` | `List<Profile>` | no |

## BatchPutProfileObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `Items` | `List<BatchPutProfileObjectRequestItem>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Successful` | `List<BatchPutProfileObjectResponseItem>` | no |
| `Failed` | `List<BatchPutProfileObjectErrorItem>` | no |

## CreateCalculatedAttributeDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `CalculatedAttributeName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `AttributeDetails` | `AttributeDetails` | yes |
| `Conditions` | `Conditions` | no |
| `Filter` | `Filter` | no |
| `Statistic` | `string` | yes |
| `UseHistoricalData` | `boolean` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatedAttributeName` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `AttributeDetails` | `AttributeDetails` | no |
| `Conditions` | `Conditions` | no |
| `Filter` | `Filter` | no |
| `Statistic` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `UseHistoricalData` | `boolean` | no |
| `Status` | `string` | no |
| `Readiness` | `Readiness` | no |
| `Tags` | `Map<string>` | no |

## CreateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DefaultExpirationDays` | `integer` | yes |
| `DefaultEncryptionKey` | `string` | no |
| `DeadLetterQueueUrl` | `string` | no |
| `Matching` | `MatchingRequest` | no |
| `RuleBasedMatching` | `RuleBasedMatchingRequest` | no |
| `DataStore` | `DataStoreRequest` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DefaultExpirationDays` | `integer` | yes |
| `DefaultEncryptionKey` | `string` | no |
| `DeadLetterQueueUrl` | `string` | no |
| `Matching` | `MatchingResponse` | no |
| `RuleBasedMatching` | `RuleBasedMatchingResponse` | no |
| `DataStore` | `DataStoreResponse` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## CreateDomainLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `LayoutDefinitionName` | `string` | yes |
| `Description` | `string` | yes |
| `DisplayName` | `string` | yes |
| `IsDefault` | `boolean` | no |
| `LayoutType` | `string` | yes |
| `Layout` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayoutDefinitionName` | `string` | yes |
| `Description` | `string` | yes |
| `DisplayName` | `string` | yes |
| `IsDefault` | `boolean` | no |
| `LayoutType` | `string` | yes |
| `Layout` | `string` | yes |
| `Version` | `string` | yes |
| `Tags` | `Map<string>` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | no |

## CreateEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | yes |
| `EventStreamName` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventStreamArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

## CreateEventTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventTriggerName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | no |
| `EventTriggerConditions` | `List<EventTriggerCondition>` | yes |
| `SegmentFilter` | `string` | no |
| `EventTriggerLimits` | `EventTriggerLimits` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventTriggerName` | `string` | no |
| `ObjectTypeName` | `string` | no |
| `Description` | `string` | no |
| `EventTriggerConditions` | `List<EventTriggerCondition>` | no |
| `SegmentFilter` | `string` | no |
| `EventTriggerLimits` | `EventTriggerLimits` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## CreateIntegrationWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `WorkflowType` | `string` | yes |
| `IntegrationConfig` | `IntegrationConfig` | yes |
| `ObjectTypeName` | `string` | yes |
| `RoleArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | yes |
| `Message` | `string` | yes |

## CreateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AccountNumber` | `string` | no |
| `AdditionalInformation` | `string` | no |
| `PartyType` | `string` | no |
| `BusinessName` | `string` | no |
| `FirstName` | `string` | no |
| `MiddleName` | `string` | no |
| `LastName` | `string` | no |
| `BirthDate` | `string` | no |
| `Gender` | `string` | no |
| `PhoneNumber` | `string` | no |
| `MobilePhoneNumber` | `string` | no |
| `HomePhoneNumber` | `string` | no |
| `BusinessPhoneNumber` | `string` | no |
| `EmailAddress` | `string` | no |
| `PersonalEmailAddress` | `string` | no |
| `BusinessEmailAddress` | `string` | no |
| `Address` | `Address` | no |
| `ShippingAddress` | `Address` | no |
| `MailingAddress` | `Address` | no |
| `BillingAddress` | `Address` | no |
| `Attributes` | `Map<string>` | no |
| `PartyTypeString` | `string` | no |
| `GenderString` | `string` | no |
| `ProfileType` | `string` | no |
| `EngagementPreferences` | `EngagementPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

## CreateRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |
| `RecommenderRecipeName` | `string` | yes |
| `RecommenderConfig` | `RecommenderConfig` | no |
| `Description` | `string` | no |
| `RecommenderSchemaName` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

## CreateRecommenderFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderFilterName` | `string` | yes |
| `RecommenderFilterExpression` | `string` | yes |
| `RecommenderSchemaName` | `string` | no |
| `Description` | `string` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderFilterArn` | `string` | yes |
| `Tags` | `Map<string>` | no |

## CreateRecommenderSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderSchemaName` | `string` | yes |
| `Fields` | `Map<List<RecommenderSchemaField>>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderSchemaArn` | `string` | yes |
| `RecommenderSchemaName` | `string` | yes |
| `Fields` | `Map<List<RecommenderSchemaField>>` | yes |
| `CreatedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `Tags` | `Map<string>` | no |

## CreateSegmentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `DisplayName` | `string` | yes |
| `Description` | `string` | no |
| `SegmentGroups` | `SegmentGroup` | no |
| `SegmentSqlQuery` | `string` | no |
| `SegmentSort` | `SegmentSort` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentDefinitionName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `SegmentDefinitionArn` | `string` | no |
| `Tags` | `Map<string>` | no |

## CreateSegmentEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentQuery` | `SegmentGroupStructure` | no |
| `SegmentSqlQuery` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `EstimateId` | `string` | no |
| `StatusCode` | `integer` | no |

## CreateSegmentSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `DataFormat` | `string` | yes |
| `EncryptionKey` | `string` | no |
| `RoleArn` | `string` | no |
| `DestinationUri` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |

## CreateUploadJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DisplayName` | `string` | yes |
| `Fields` | `Map<ObjectTypeField>` | yes |
| `UniqueKey` | `string` | yes |
| `DataExpiry` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | yes |

## DeleteCalculatedAttributeDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `CalculatedAttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteDomainLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `LayoutDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteDomainObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteEventTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventTriggerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteProfileKey

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `KeyName` | `string` | yes |
| `Values` | `List<string>` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteProfileObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |
| `ProfileObjectUniqueKey` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteProfileObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteRecommenderFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderFilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | yes |

## DeleteRecommenderSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderSchemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DeleteSegmentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteSegmentSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## DeleteWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## DetectProfileObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Objects` | `List<string>` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DetectedProfileObjectTypes` | `List<DetectedProfileObjectType>` | no |

## DisassociateStreamForSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## GetAutoMergingPreview

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Consolidation` | `Consolidation` | yes |
| `ConflictResolution` | `ConflictResolution` | yes |
| `MinAllowedConfidenceScoreForMerging` | `double` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NumberOfMatchesInSample` | `long` | no |
| `NumberOfProfilesInSample` | `long` | no |
| `NumberOfProfilesWillBeMerged` | `long` | no |

## GetCalculatedAttributeDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `CalculatedAttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatedAttributeName` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Statistic` | `string` | no |
| `Filter` | `Filter` | no |
| `Conditions` | `Conditions` | no |
| `AttributeDetails` | `AttributeDetails` | no |
| `UseHistoricalData` | `boolean` | no |
| `Status` | `string` | no |
| `Readiness` | `Readiness` | no |
| `Tags` | `Map<string>` | no |

## GetCalculatedAttributeForProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `CalculatedAttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatedAttributeName` | `string` | no |
| `DisplayName` | `string` | no |
| `IsDataPartial` | `string` | no |
| `Value` | `string` | no |
| `LastObjectTimestamp` | `timestamp` | no |

## GetDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DefaultExpirationDays` | `integer` | no |
| `DefaultEncryptionKey` | `string` | no |
| `DeadLetterQueueUrl` | `string` | no |
| `Stats` | `DomainStats` | no |
| `Matching` | `MatchingResponse` | no |
| `RuleBasedMatching` | `RuleBasedMatchingResponse` | no |
| `DataStore` | `DataStoreResponse` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## GetDomainLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `LayoutDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayoutDefinitionName` | `string` | yes |
| `Description` | `string` | yes |
| `DisplayName` | `string` | yes |
| `IsDefault` | `boolean` | no |
| `LayoutType` | `string` | yes |
| `Layout` | `string` | yes |
| `Version` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## GetDomainObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | no |
| `EncryptionKey` | `string` | no |
| `Fields` | `Map<DomainObjectTypeField>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetEventStream

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventStreamName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventStreamArn` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `State` | `string` | yes |
| `StoppedSince` | `timestamp` | no |
| `DestinationDetails` | `EventStreamDestinationDetails` | yes |
| `Tags` | `Map<string>` | no |

## GetEventTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventTriggerName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventTriggerName` | `string` | no |
| `ObjectTypeName` | `string` | no |
| `Description` | `string` | no |
| `EventTriggerConditions` | `List<EventTriggerCondition>` | no |
| `SegmentFilter` | `string` | no |
| `EventTriggerLimits` | `EventTriggerLimits` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetIdentityResolutionJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `JobId` | `string` | no |
| `Status` | `string` | no |
| `Message` | `string` | no |
| `JobStartTime` | `timestamp` | no |
| `JobEndTime` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `JobExpirationTime` | `timestamp` | no |
| `AutoMerging` | `AutoMerging` | no |
| `ExportingLocation` | `ExportingLocation` | no |
| `JobStats` | `JobStats` | no |

## GetIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | yes |
| `ObjectTypeName` | `string` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |
| `ObjectTypeNames` | `Map<string>` | no |
| `WorkflowId` | `string` | no |
| `IsUnstructured` | `boolean` | no |
| `RoleArn` | `string` | no |
| `EventTriggerNames` | `List<string>` | no |
| `Scope` | `string` | no |

## GetMatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MatchGenerationDate` | `timestamp` | no |
| `PotentialMatches` | `integer` | no |
| `Matches` | `List<MatchItem>` | no |

## GetObjectTypeAttributeStatistics

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `AttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Statistics` | `GetObjectTypeAttributeStatisticsStats` | yes |
| `CalculatedAt` | `timestamp` | yes |

## GetProfileHistoryRecord

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `Id` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Id` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | no |
| `ActionType` | `string` | yes |
| `ProfileObjectUniqueKey` | `string` | no |
| `Content` | `string` | no |
| `PerformedBy` | `string` | no |

## GetProfileObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | yes |
| `TemplateId` | `string` | no |
| `ExpirationDays` | `integer` | no |
| `EncryptionKey` | `string` | no |
| `AllowProfileCreation` | `boolean` | no |
| `SourceLastUpdatedTimestampFormat` | `string` | no |
| `MaxAvailableProfileObjectCount` | `integer` | no |
| `MaxProfileObjectCount` | `integer` | no |
| `SourcePriority` | `integer` | no |
| `Fields` | `Map<ObjectTypeField>` | no |
| `Keys` | `Map<List<ObjectTypeKey>>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## GetProfileObjectTypeTemplate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `TemplateId` | `string` | no |
| `SourceName` | `string` | no |
| `SourceObject` | `string` | no |
| `AllowProfileCreation` | `boolean` | no |
| `SourceLastUpdatedTimestampFormat` | `string` | no |
| `Fields` | `Map<ObjectTypeField>` | no |
| `Keys` | `Map<List<ObjectTypeKey>>` | no |

## GetProfileRecommendations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `RecommenderName` | `string` | yes |
| `Context` | `Map<string>` | no |
| `RecommenderFilters` | `List<RecommenderFilter>` | no |
| `RecommenderPromotionalFilters` | `List<RecommenderPromotionalFilter>` | no |
| `CandidateIds` | `List<string>` | no |
| `MaxResults` | `integer` | no |
| `MetadataConfig` | `MetadataConfig` | no |
| `DiversityConfig` | `RecommendationDiversityConfig` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Recommendations` | `List<Recommendation>` | no |

## GetRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |
| `TrainingMetricsCount` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderName` | `string` | yes |
| `RecommenderRecipeName` | `string` | yes |
| `RecommenderSchemaName` | `string` | no |
| `RecommenderConfig` | `RecommenderConfig` | no |
| `Description` | `string` | no |
| `Status` | `string` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `CreatedAt` | `timestamp` | no |
| `FailureReason` | `string` | no |
| `LatestRecommenderUpdate` | `RecommenderUpdate` | no |
| `ActiveRecommenderVersionName` | `string` | no |
| `TrainingMetrics` | `List<TrainingMetrics>` | no |
| `Tags` | `Map<string>` | no |

## GetRecommenderFilter

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderFilterName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderFilterName` | `string` | yes |
| `RecommenderFilterExpression` | `string` | yes |
| `RecommenderSchemaName` | `string` | no |
| `CreatedAt` | `timestamp` | yes |
| `Status` | `string` | yes |
| `Description` | `string` | no |
| `FailureReason` | `string` | no |
| `Tags` | `Map<string>` | yes |

## GetRecommenderSchema

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderSchemaName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderSchemaName` | `string` | yes |
| `Fields` | `Map<List<RecommenderSchemaField>>` | yes |
| `CreatedAt` | `timestamp` | yes |
| `Status` | `string` | yes |

## GetSegmentDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentDefinitionName` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `SegmentGroups` | `SegmentGroup` | no |
| `SegmentSort` | `SegmentSort` | no |
| `SegmentDefinitionArn` | `string` | yes |
| `CreatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |
| `SegmentSqlQuery` | `string` | no |
| `SegmentType` | `string` | no |

## GetSegmentEstimate

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EstimateId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `EstimateId` | `string` | no |
| `Status` | `string` | no |
| `Estimate` | `string` | no |
| `Message` | `string` | no |
| `StatusCode` | `integer` | no |

## GetSegmentMembership

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `ProfileIds` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SegmentDefinitionName` | `string` | no |
| `Profiles` | `List<ProfileQueryResult>` | no |
| `Failures` | `List<ProfileQueryFailures>` | no |
| `LastComputedAt` | `timestamp` | no |

## GetSegmentSnapshot

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `SnapshotId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `SnapshotId` | `string` | yes |
| `Status` | `string` | yes |
| `StatusMessage` | `string` | no |
| `DataFormat` | `string` | yes |
| `EncryptionKey` | `string` | no |
| `RoleArn` | `string` | no |
| `DestinationUri` | `string` | no |

## GetSegmentSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `Message` | `string` | no |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |
| `ScheduledExecutions` | `ScheduledExecutions` | no |
| `StartedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |

## GetSimilarProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `MatchType` | `string` | yes |
| `SearchKey` | `string` | yes |
| `SearchValue` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileIds` | `List<string>` | no |
| `MatchId` | `string` | no |
| `MatchType` | `string` | no |
| `RuleLevel` | `integer` | no |
| `ConfidenceScore` | `double` | no |
| `NextToken` | `string` | no |

## GetStreamForSegments

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `AssociatedAt` | `timestamp` | no |
| `AssociatedSegments` | `List<AssociatedSegment>` | no |
| `DomainName` | `string` | no |
| `DestinationArn` | `string` | no |
| `DestinationRoleArn` | `string` | no |
| `State` | `string` | no |
| `DisassociatedAt` | `timestamp` | no |
| `FailureReason` | `string` | no |

## GetUploadJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `JobId` | `string` | no |
| `DisplayName` | `string` | no |
| `Status` | `string` | no |
| `StatusReason` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `CompletedAt` | `timestamp` | no |
| `Fields` | `Map<ObjectTypeField>` | no |
| `UniqueKey` | `string` | no |
| `ResultsSummary` | `ResultsSummary` | no |
| `DataExpiry` | `integer` | no |

## GetUploadJobPath

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Url` | `string` | yes |
| `ClientToken` | `string` | no |
| `ValidUntil` | `timestamp` | no |

## GetWorkflow

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `WorkflowId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | no |
| `WorkflowType` | `string` | no |
| `Status` | `string` | no |
| `ErrorDescription` | `string` | no |
| `StartDate` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Attributes` | `WorkflowAttributes` | no |
| `Metrics` | `WorkflowMetrics` | no |

## GetWorkflowSteps

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `WorkflowId` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `WorkflowId` | `string` | no |
| `WorkflowType` | `string` | no |
| `Items` | `List<WorkflowStepItem>` | no |
| `NextToken` | `string` | no |

## ListAccountIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Uri` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeHidden` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListIntegrationItem>` | no |
| `NextToken` | `string` | no |

## ListCalculatedAttributeDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListCalculatedAttributeDefinitionItem>` | no |
| `NextToken` | `string` | no |

## ListCalculatedAttributesForProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListCalculatedAttributeForProfileItem>` | no |
| `NextToken` | `string` | no |

## ListDomainLayouts

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<LayoutItem>` | no |
| `NextToken` | `string` | no |

## ListDomainObjectTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<DomainObjectTypesListItem>` | no |
| `NextToken` | `string` | no |

## ListDomains

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListDomainItem>` | no |
| `NextToken` | `string` | no |

## ListEventStreams

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<EventStreamSummary>` | no |
| `NextToken` | `string` | no |

## ListEventTriggers

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<EventTriggerSummaryItem>` | no |
| `NextToken` | `string` | no |

## ListIdentityResolutionJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `IdentityResolutionJobsList` | `List<IdentityResolutionJob>` | no |
| `NextToken` | `string` | no |

## ListIntegrations

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `IncludeHidden` | `boolean` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListIntegrationItem>` | no |
| `NextToken` | `string` | no |

## ListObjectTypeAttributeValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `AttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListObjectTypeAttributeValuesItem>` | no |
| `NextToken` | `string` | no |

## ListObjectTypeAttributes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListObjectTypeAttributeItem>` | no |
| `NextToken` | `string` | no |

## ListProfileAttributeValues

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `AttributeName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | no |
| `AttributeName` | `string` | no |
| `Items` | `List<AttributeValueItem>` | no |
| `StatusCode` | `integer` | no |

## ListProfileHistoryRecords

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `ObjectTypeName` | `string` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `ActionType` | `string` | no |
| `PerformedBy` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileHistoryRecords` | `List<ProfileHistoryRecord>` | no |
| `NextToken` | `string` | no |

## ListProfileObjectTypeTemplates

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListProfileObjectTypeTemplateItem>` | no |
| `NextToken` | `string` | no |

## ListProfileObjectTypes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListProfileObjectTypeItem>` | no |
| `NextToken` | `string` | no |

## ListProfileObjects

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `ObjectFilter` | `ObjectFilter` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListProfileObjectsItem>` | no |
| `NextToken` | `string` | no |

## ListRecommenderFilters

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecommenderFilters` | `List<RecommenderFilterSummary>` | no |

## ListRecommenderRecipes

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecommenderRecipes` | `List<RecommenderRecipe>` | no |

## ListRecommenderSchemas

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `RecommenderSchemas` | `List<RecommenderSchemaSummary>` | no |

## ListRecommenders

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Recommenders` | `List<RecommenderSummary>` | no |

## ListRuleBasedMatches

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `MatchIds` | `List<string>` | no |
| `NextToken` | `string` | no |

## ListSegmentDefinitions

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Items` | `List<SegmentDefinitionItem>` | no |

## ListSegmentSubscriptionEvents

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Events` | `List<SubscriptionEventItem>` | no |
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

## ListUploadJobs

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MaxResults` | `integer` | no |
| `NextToken` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `Items` | `List<UploadJobItem>` | no |

## ListWorkflows

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `WorkflowType` | `string` | no |
| `Status` | `string` | no |
| `QueryStartDate` | `timestamp` | no |
| `QueryEndDate` | `timestamp` | no |
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<ListWorkflowsItem>` | no |
| `NextToken` | `string` | no |

## MergeProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `MainProfileId` | `string` | yes |
| `ProfileIdsToBeMerged` | `List<string>` | yes |
| `FieldSourceProfileIds` | `FieldSourceProfileIds` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Message` | `string` | no |

## PutDomainObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | no |
| `EncryptionKey` | `string` | no |
| `Fields` | `Map<DomainObjectTypeField>` | yes |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectTypeName` | `string` | no |
| `Description` | `string` | no |
| `EncryptionKey` | `string` | no |
| `Fields` | `Map<DomainObjectTypeField>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## PutIntegration

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | no |
| `ObjectTypeName` | `string` | no |
| `ObjectTypeNames` | `Map<string>` | no |
| `Tags` | `Map<string>` | no |
| `FlowDefinition` | `FlowDefinition` | no |
| `RoleArn` | `string` | no |
| `EventTriggerNames` | `List<string>` | no |
| `Scope` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `Uri` | `string` | yes |
| `ObjectTypeName` | `string` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |
| `ObjectTypeNames` | `Map<string>` | no |
| `WorkflowId` | `string` | no |
| `IsUnstructured` | `boolean` | no |
| `RoleArn` | `string` | no |
| `EventTriggerNames` | `List<string>` | no |
| `Scope` | `string` | no |

## PutProfileObject

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectTypeName` | `string` | yes |
| `Object` | `string` | yes |
| `DomainName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileObjectUniqueKey` | `string` | no |

## PutProfileObjectType

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | yes |
| `TemplateId` | `string` | no |
| `ExpirationDays` | `integer` | no |
| `EncryptionKey` | `string` | no |
| `AllowProfileCreation` | `boolean` | no |
| `SourceLastUpdatedTimestampFormat` | `string` | no |
| `MaxProfileObjectCount` | `integer` | no |
| `SourcePriority` | `integer` | no |
| `Fields` | `Map<ObjectTypeField>` | no |
| `Keys` | `Map<List<ObjectTypeKey>>` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ObjectTypeName` | `string` | yes |
| `Description` | `string` | yes |
| `TemplateId` | `string` | no |
| `ExpirationDays` | `integer` | no |
| `EncryptionKey` | `string` | no |
| `AllowProfileCreation` | `boolean` | no |
| `SourceLastUpdatedTimestampFormat` | `string` | no |
| `MaxProfileObjectCount` | `integer` | no |
| `MaxAvailableProfileObjectCount` | `integer` | no |
| `SourcePriority` | `integer` | no |
| `Fields` | `Map<ObjectTypeField>` | no |
| `Keys` | `Map<List<ObjectTypeKey>>` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## PutSegmentSubscription

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `SegmentDefinitionName` | `string` | yes |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Status` | `string` | no |
| `ScheduleConfiguration` | `ScheduleConfiguration` | no |
| `StartedAt` | `timestamp` | no |

## SearchProfiles

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `NextToken` | `string` | no |
| `MaxResults` | `integer` | no |
| `DomainName` | `string` | yes |
| `KeyName` | `string` | yes |
| `Values` | `List<string>` | yes |
| `AdditionalSearchKeys` | `List<AdditionalSearchKey>` | no |
| `LogicalOperator` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `Items` | `List<Profile>` | no |
| `NextToken` | `string` | no |

## StartRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StartUploadJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `JobId` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## StopUploadJob

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `JobId` | `string` | yes |

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


## UntagResource

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `resourceArn` | `string` | yes |
| `tagKeys` | `List<string>` | yes |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|


## UpdateCalculatedAttributeDefinition

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `CalculatedAttributeName` | `string` | yes |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `Conditions` | `Conditions` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `CalculatedAttributeName` | `string` | no |
| `DisplayName` | `string` | no |
| `Description` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Statistic` | `string` | no |
| `Conditions` | `Conditions` | no |
| `AttributeDetails` | `AttributeDetails` | no |
| `UseHistoricalData` | `boolean` | no |
| `Status` | `string` | no |
| `Readiness` | `Readiness` | no |
| `Tags` | `Map<string>` | no |

## UpdateDomain

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DefaultExpirationDays` | `integer` | no |
| `DefaultEncryptionKey` | `string` | no |
| `DeadLetterQueueUrl` | `string` | no |
| `Matching` | `MatchingRequest` | no |
| `RuleBasedMatching` | `RuleBasedMatchingRequest` | no |
| `DataStore` | `DataStoreRequest` | no |
| `Tags` | `Map<string>` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `DefaultExpirationDays` | `integer` | no |
| `DefaultEncryptionKey` | `string` | no |
| `DeadLetterQueueUrl` | `string` | no |
| `Matching` | `MatchingResponse` | no |
| `RuleBasedMatching` | `RuleBasedMatchingResponse` | no |
| `DataStore` | `DataStoreResponse` | no |
| `CreatedAt` | `timestamp` | yes |
| `LastUpdatedAt` | `timestamp` | yes |
| `Tags` | `Map<string>` | no |

## UpdateDomainLayout

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `LayoutDefinitionName` | `string` | yes |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `IsDefault` | `boolean` | no |
| `LayoutType` | `string` | no |
| `Layout` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `LayoutDefinitionName` | `string` | no |
| `Description` | `string` | no |
| `DisplayName` | `string` | no |
| `IsDefault` | `boolean` | no |
| `LayoutType` | `string` | no |
| `Layout` | `string` | no |
| `Version` | `string` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## UpdateEventTrigger

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `EventTriggerName` | `string` | yes |
| `ObjectTypeName` | `string` | no |
| `Description` | `string` | no |
| `EventTriggerConditions` | `List<EventTriggerCondition>` | no |
| `SegmentFilter` | `string` | no |
| `EventTriggerLimits` | `EventTriggerLimits` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `EventTriggerName` | `string` | no |
| `ObjectTypeName` | `string` | no |
| `Description` | `string` | no |
| `EventTriggerConditions` | `List<EventTriggerCondition>` | no |
| `SegmentFilter` | `string` | no |
| `EventTriggerLimits` | `EventTriggerLimits` | no |
| `CreatedAt` | `timestamp` | no |
| `LastUpdatedAt` | `timestamp` | no |
| `Tags` | `Map<string>` | no |

## UpdateProfile

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `ProfileId` | `string` | yes |
| `AdditionalInformation` | `string` | no |
| `AccountNumber` | `string` | no |
| `PartyType` | `string` | no |
| `BusinessName` | `string` | no |
| `FirstName` | `string` | no |
| `MiddleName` | `string` | no |
| `LastName` | `string` | no |
| `BirthDate` | `string` | no |
| `Gender` | `string` | no |
| `PhoneNumber` | `string` | no |
| `MobilePhoneNumber` | `string` | no |
| `HomePhoneNumber` | `string` | no |
| `BusinessPhoneNumber` | `string` | no |
| `EmailAddress` | `string` | no |
| `PersonalEmailAddress` | `string` | no |
| `BusinessEmailAddress` | `string` | no |
| `Address` | `UpdateAddress` | no |
| `ShippingAddress` | `UpdateAddress` | no |
| `MailingAddress` | `UpdateAddress` | no |
| `BillingAddress` | `UpdateAddress` | no |
| `Attributes` | `Map<string>` | no |
| `PartyTypeString` | `string` | no |
| `GenderString` | `string` | no |
| `ProfileType` | `string` | no |
| `EngagementPreferences` | `EngagementPreferences` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `ProfileId` | `string` | yes |

## UpdateRecommender

Input:

| Field | Type | Required in AWS schema |
|---|---|---|
| `DomainName` | `string` | yes |
| `RecommenderName` | `string` | yes |
| `Description` | `string` | no |
| `RecommenderConfig` | `RecommenderConfig` | no |
| `RecommenderVersionName` | `string` | no |

Output:

| Field | Type | Required in AWS schema |
|---|---|---|
| `RecommenderName` | `string` | yes |

